package execution

import (
	"crypto/ecdsa"
	"math/big"
	"testing"
	"time"

	"github.com/btcsuite/btcd/btcutil/bech32"
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
	"github.com/ethereum/go-ethereum/miner"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"
)

var (
	// testKey is a private key to use for funding a tester account.
	testKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")

	// testAddr is the Ethereum address of the tester account.
	testAddr = crypto.PubkeyToAddress(testKey.PublicKey)

	testToAddress = common.HexToAddress("0x9a9070028361F7AAbeB3f2F2Dc07F82C4a98A02a")

	testBalance = big.NewInt(2e18)
)

func generateMergeChain(n int, merged bool) (*core.Genesis, []*types.Block, string, *ecdsa.PrivateKey) {
	config := *params.AllEthashProtocolChanges
	engine := consensus.Engine(beaconConsensus.New(ethash.NewFaker()))
	if merged {
		config.TerminalTotalDifficulty = common.Big0
		engine = beaconConsensus.NewFaker()
	}

	bridgeAddressKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	bridgeAddress := crypto.PubkeyToAddress(bridgeAddressKey.PublicKey)
	bridgeAddressBytes, err := bech32.ConvertBits(bridgeAddress.Bytes(), 8, 5, false)
	if err != nil {
		panic(err)
	}

	config.AstriaRollupName = "astria"
	config.AstriaSequencerAddressPrefix = "astria"
	config.AstriaSequencerInitialHeight = 10
	config.AstriaCelestiaInitialHeight = 10
	config.AstriaCelestiaHeightVariance = 10

	bech32mBridgeAddress, err := bech32.EncodeM(config.AstriaSequencerAddressPrefix, bridgeAddressBytes)
	if err != nil {
		panic(err)
	}
	config.AstriaBridgeAddressConfigs = []params.AstriaBridgeAddressConfig{
		{
			BridgeAddress:  bech32mBridgeAddress,
			SenderAddress:  common.Address{},
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
	config.AstriaFeeCollectors = astriaFeeCollectors

	genesis := &core.Genesis{
		Config: &config,
		Alloc: core.GenesisAlloc{
			testAddr: {Balance: testBalance},
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
		tx, _ := types.SignTx(types.NewTransaction(testNonce, testToAddress, big.NewInt(1), params.TxGas, big.NewInt(params.InitialBaseFee*2), nil), types.LatestSigner(&config), testKey)
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

	return genesis, blocks, bech32mBridgeAddress, feeCollectorKey
}

// startEthService creates a full node instance for testing.
func startEthService(t *testing.T, genesis *core.Genesis) *eth.Ethereum {
	n, err := node.New(&node.Config{})
	require.Nil(t, err, "can't create node")
	mcfg := miner.DefaultConfig
	mcfg.PendingFeeRecipient = testAddr
	ethcfg := &ethconfig.Config{Genesis: genesis, SyncMode: downloader.FullSync, TrieTimeout: time.Minute, TrieDirtyCache: 256, TrieCleanCache: 256, Miner: mcfg}
	ethservice, err := eth.New(n, ethcfg)
	require.Nil(t, err, "can't create eth service")

	ethservice.SetSynced()
	return ethservice
}

func setupExecutionService(t *testing.T, noOfBlocksToGenerate int) (*eth.Ethereum, *ExecutionServiceServerV1) {
	t.Helper()
	genesis, blocks, bridgeAddress, feeCollectorKey := generateMergeChain(noOfBlocksToGenerate, true)
	ethservice := startEthService(t, genesis)

	serviceV1Alpha1, err := NewExecutionServiceServerV1(ethservice)
	require.Nil(t, err, "can't create execution service")

	feeCollector := crypto.PubkeyToAddress(feeCollectorKey.PublicKey)
	require.Equal(t, feeCollector, serviceV1Alpha1.nextFeeRecipient, "nextFeeRecipient not set correctly")

	bridgeAsset := genesis.Config.AstriaBridgeAddressConfigs[0].AssetDenom
	_, ok := serviceV1Alpha1.bridgeAllowedAssets[bridgeAsset]
	require.True(t, ok, "bridgeAllowedAssetIDs does not contain bridge asset id")

	_, ok = serviceV1Alpha1.bridgeAddresses[bridgeAddress]
	require.True(t, ok, "bridgeAddress not set correctly")

	_, err = ethservice.BlockChain().InsertChain(blocks)
	require.Nil(t, err, "can't insert blocks")

	return ethservice, serviceV1Alpha1
}
