package execution

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	beaconConsensus "github.com/ethereum/go-ethereum/consensus/beacon"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"net"
	"testing"
	"time"
)

var (
	// testKey is a private key to use for funding a tester account.
	testKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")

	// testAddr is the Ethereum address of the tester account.
	testAddr = crypto.PubkeyToAddress(testKey.PublicKey)

	testBalance = big.NewInt(2e18)
)

func generateMergeChain(n int, merged bool) (*core.Genesis, []*types.Block, *ecdsa.PrivateKey, *ecdsa.PrivateKey) {
	config := *params.AllEthashProtocolChanges
	engine := consensus.Engine(beaconConsensus.New(ethash.NewFaker()))
	if merged {
		config.TerminalTotalDifficulty = common.Big0
		config.TerminalTotalDifficultyPassed = true
		engine = beaconConsensus.NewFaker()
	}

	bridgeAddressKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	bridgeAddress := crypto.PubkeyToAddress(bridgeAddressKey.PublicKey)

	config.AstriaConfig.RollupName = "astria"
	config.AstriaConfig.SequencerInitialHeight = 10
	config.AstriaConfig.CelestiaInitialHeight = 10
	config.AstriaConfig.CelestiaHeightVariance = 10
	config.AstriaConfig.BridgeAddressConfigs = []params.AstriaBridgeAddressConfig{
		{
			BridgeAddress:  bridgeAddress.Bytes(),
			StartHeight:    2,
			AssetDenom:     "nria",
			AssetPrecision: 18,
			Erc20Asset:     nil,
		},
	}

	feeCollectorKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	feeCollector := crypto.PubkeyToAddress(feeCollectorKey.PublicKey)

	astriaFeeCollectors := make(map[uint32]common.Address)
	astriaFeeCollectors[1] = feeCollector
	config.AstriaConfig.FeeCollectors = astriaFeeCollectors

	genesis := &core.Genesis{
		Config: &config,
		Alloc: core.GenesisAlloc{
			testAddr:                         {Balance: testBalance},
			params.BeaconRootsStorageAddress: {Balance: common.Big0, Code: common.Hex2Bytes("3373fffffffffffffffffffffffffffffffffffffffe14604457602036146024575f5ffd5b620180005f350680545f35146037575f5ffd5b6201800001545f5260205ff35b6201800042064281555f359062018000015500")},
		},
		ExtraData:  []byte("test genesis"),
		Timestamp:  9000,
		BaseFee:    big.NewInt(params.InitialBaseFee),
		Difficulty: big.NewInt(0),
	}
	testNonce := uint64(0)
	generate := func(i int, g *core.BlockGen) {
		g.OffsetTime(5)
		g.SetExtra([]byte("test"))
		tx, _ := types.SignTx(types.NewTransaction(testNonce, common.HexToAddress("0x9a9070028361F7AAbeB3f2F2Dc07F82C4a98A02a"), big.NewInt(1), params.TxGas, big.NewInt(params.InitialBaseFee*2), nil), types.LatestSigner(&config), testKey)
		g.AddTx(tx)
		testNonce++
	}
	_, blocks, _ := core.GenerateChainWithGenesis(genesis, engine, n, generate)

	if !merged {
		totalDifficulty := big.NewInt(0)
		for _, b := range blocks {
			totalDifficulty.Add(totalDifficulty, b.Difficulty())
		}
		config.TerminalTotalDifficulty = totalDifficulty
	}

	return genesis, blocks, bridgeAddressKey, feeCollectorKey
}

func getFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

// startEthService creates a full node instance for testing.
func startEthService(t *testing.T, genesis *core.Genesis) (*node.Node, *eth.Ethereum) {
	freePort, err := getFreePort()
	if err != nil {
		t.Fatal("can't get free port:", err)
	}
	n, err := node.New(&node.Config{
		P2P: p2p.Config{
			ListenAddr:  "0.0.0.0:0",
			NoDiscovery: true,
			MaxPeers:    25,
		},
		GRPCHost: "127.0.0.1",
		GRPCPort: freePort,
	})
	if err != nil {
		t.Fatal("can't create node:", err)
	}

	ethcfg := &ethconfig.Config{Genesis: genesis, SyncMode: downloader.FullSync, TrieTimeout: time.Minute, TrieDirtyCache: 256, TrieCleanCache: 256}
	ethservice, err := eth.New(n, ethcfg)
	if err != nil {
		t.Fatal("can't create eth service:", err)
	}

	ethservice.SetEtherbase(testAddr)
	ethservice.SetSynced()

	return n, ethservice
}

func setupExecutionService(t *testing.T, noOfBlocksToGenerate int) (*node.Node, *eth.Ethereum, *ExecutionServiceServerV1Alpha2) {
	t.Helper()
	genesis, blocks, bridgeAddressKey, feeCollectorKey := generateMergeChain(noOfBlocksToGenerate, true)
	n, ethservice := startEthService(t, genesis)

	serviceV1Alpha1, err := NewExecutionServiceServerV1Alpha2(ethservice)
	if err != nil {
		t.Fatal("can't create execution service:", err)
	}

	feeCollector := crypto.PubkeyToAddress(feeCollectorKey.PublicKey)
	if serviceV1Alpha1.nextFeeRecipient != feeCollector {
		t.Fatalf("nextFeeRecipient not set correctly")
	}

	bridgeAsset := sha256.Sum256([]byte(genesis.Config.AstriaBridgeAddressConfigs()[0].AssetDenom))
	_, ok := serviceV1Alpha1.bridgeAllowedAssetIDs[bridgeAsset]
	if !ok {
		t.Fatalf("bridgeAllowedAssetIDs does not contain bridge asset id")
	}

	bridgeAddress := crypto.PubkeyToAddress(bridgeAddressKey.PublicKey)
	_, ok = serviceV1Alpha1.bridgeAddresses[string(bridgeAddress.Bytes())]
	if !ok {
		t.Fatalf("bridgeAddress not set correctly")
	}

	utils.RegisterGRPCExecutionService(n, serviceV1Alpha1, n.Config())

	if err := n.Start(); err != nil {
		t.Fatal("can't start node:", err)
	}
	if _, err := ethservice.BlockChain().InsertChain(blocks); err != nil {
		n.Close()
		t.Fatal("can't import test blocks:", err)
	}

	return n, ethservice, serviceV1Alpha1

}

func GrpcEndpointWithoutPrefix(n *node.Node) string {
	grpcEndpoint := n.GRPCEndpoint()
	// remove the http:// prefix
	return grpcEndpoint[7:]
}
