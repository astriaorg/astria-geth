package execution

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"math/big"
	"testing"
	"time"

	astriaPb "buf.build/gen/go/astria/execution-apis/protocolbuffers/go/astria/execution/v2"
	primitivev1 "buf.build/gen/go/astria/primitives/protocolbuffers/go/astria/primitive/v1"
	"github.com/btcsuite/btcd/btcutil/bech32"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	beaconConsensus "github.com/ethereum/go-ethereum/consensus/beacon"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/miner"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// testKey is a private key to use for funding a tester account.
	testKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")

	// testAddr is the Ethereum address of the tester account.
	testAddr = crypto.PubkeyToAddress(testKey.PublicKey)

	testToAddress = common.HexToAddress("0x9a9070028361F7AAbeB3f2F2Dc07F82C4a98A02a")

	testBalance = big.NewInt(2e18)
)

func bigIntToProtoU128(i *big.Int) *primitivev1.Uint128 {
	lo := i.Uint64()
	hi := new(big.Int).Rsh(i, 64).Uint64()
	return &primitivev1.Uint128{Lo: lo, Hi: hi}
}

func generateMergeChain(n int, merged bool, gasLimit uint64, halted ...bool) (*core.Genesis, []*types.Block, string, *ecdsa.PrivateKey) {
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
	bridgeAddressBytes, err := bech32.ConvertBits(bridgeAddress.Bytes(), 8, 5, false)
	if err != nil {
		panic(err)
	}

	bech32mBridgeAddress, err := bech32.EncodeM("astria", bridgeAddressBytes)
	if err != nil {
		panic(err)
	}

	feeCollectorKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	feeCollector := crypto.PubkeyToAddress(feeCollectorKey.PublicKey)

	// Check if halted parameter was provided
	isHalted := false
	if len(halted) > 0 {
		isHalted = halted[0]
	}

	config.AstriaRollupName = "astria"
	config.AstriaForks, _ = params.NewAstriaForks(map[string]params.AstriaForkConfig{
		"genesis": {
			Height:       1,
			FeeCollector: &feeCollector,
			Halt:         isHalted, // Set the halt flag based on the parameter
			Sequencer: &params.AstriaSequencerConfig{
				ChainID:       "astria",
				AddressPrefix: "astria",
				StartHeight:   10,
			},
			Celestia: &params.AstriaCelestiaConfig{
				ChainID:                  "celestia",
				StartHeight:              10,
				SearchHeightMaxLookAhead: 10,
			},
			BridgeAddresses: []params.AstriaBridgeAddressConfig{
				{
					BridgeAddress:  bech32mBridgeAddress,
					SenderAddress:  common.Address{},
					AssetDenom:     "nria",
					AssetPrecision: 18,
					Erc20Asset:     nil,
				},
			},
		},
	})

	genesis := &core.Genesis{
		Config: &config,
		Alloc: types.GenesisAlloc{
			testAddr: {Balance: testBalance},
		},
		ExtraData:  []byte("test genesis"),
		Timestamp:  9000,
		BaseFee:    big.NewInt(params.InitialBaseFee),
		Difficulty: big.NewInt(0),
		GasLimit:   gasLimit,
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

// setupExecutionService creates an execution service for testing
func setupExecutionService(t *testing.T, noOfBlocksToGenerate int, lowGasLimit bool, halted ...bool) (*eth.Ethereum, *ExecutionServiceServerV2) {
	t.Helper()
	var gasLimit uint64
	if lowGasLimit {
		gasLimit = 5000
	} else {
		gasLimit = 50000000
	}
	// Check if halted parameter was provided
	isHalted := false
	if len(halted) > 0 {
		isHalted = halted[0]
	}
	genesis, blocks, bridgeAddress, feeCollectorKey := generateMergeChain(noOfBlocksToGenerate, true, gasLimit, isHalted)
	ethservice := startEthService(t, genesis)

	serviceV2, err := NewExecutionServiceServerV2(ethservice, false, 0, 0)
	require.Nil(t, err, "can't create execution service")

	fork := genesis.Config.AstriaForks.GetForkAtHeight(1)

	feeCollector := crypto.PubkeyToAddress(feeCollectorKey.PublicKey)
	require.Equal(t, feeCollector, fork.FeeCollector, "feeCollector not set correctly")

	// If the fork is expected to be halted, verify it
	if isHalted {
		require.True(t, fork.Halt, "fork should be halted")
	}

	bridgeCfg, ok := fork.BridgeAddresses[bridgeAddress]
	require.True(t, ok, "bridgeAddress not set correctly")

	bridgeAsset := bridgeCfg.AssetDenom
	_, ok = fork.BridgeAllowedAssets[bridgeAsset]
	require.True(t, ok, "bridgeAllowedAssetIDs does not contain bridge asset id")

	_, err = ethservice.BlockChain().InsertChain(blocks)
	require.Nil(t, err, "can't insert blocks")

	return ethservice, serviceV2
}

// setupExecutionServiceWithHaltedFork sets up an execution service with a halted fork
func setupExecutionServiceWithHaltedFork(t *testing.T, noOfBlocksToGenerate int) (*eth.Ethereum, *ExecutionServiceServerV2) {
	return setupExecutionService(t, noOfBlocksToGenerate, false, true)
}

func (s *ExecutionServiceServerV2) createExecutionSessionWithForkOverride(ctx context.Context, req *astriaPb.CreateExecutionSessionRequest, fork params.AstriaForkData) (*astriaPb.ExecutionSession, error) {
	log.Debug("CreateExecutionSession called")
	createExecutionSessionRequestCount.Inc(1)

	// We shouldn't create a new session if we are actively executing within one.
	s.blockExecutionLock.Lock()
	defer s.blockExecutionLock.Unlock()
	s.commitmentUpdateLock.Lock()
	defer s.commitmentUpdateLock.Unlock()

	rollupHash := sha256.Sum256([]byte(s.bc.Config().AstriaRollupName))
	rollupId := primitivev1.RollupId{Inner: rollupHash[:]}

	if fork.Halt {
		log.Error("CreateExecutionSession called at halted fork", "fork", fork.Name)
		return nil, status.Error(codes.FailedPrecondition, "Execution session cannot be created at halted fork")
	}

	s.activeSessionId = uuid.NewString()
	s.activeFork = &fork

	softBlock, err := ethHeaderToExecutedBlockMetadata(s.bc.CurrentSafeBlock())
	if err != nil {
		log.Error("error finding safe block", err)
		return nil, status.Error(codes.Internal, "Could not locate soft block")
	}

	firmBlock, err := ethHeaderToExecutedBlockMetadata(s.bc.CurrentFinalBlock())
	if err != nil {
		log.Error("error finding final block", err)
		return nil, status.Error(codes.Internal, "Could not locate firm block")
	}

	// sanity code check for oracle contract address
	if fork.Oracle.ContractAddress != (common.Address{}) {
		height := s.bc.CurrentFinalBlock().Number.Uint64() // consider should this be the current final block, safe block, or the fork start height - 1?
		state, header, err := s.eth.APIBackend.StateAndHeaderByNumber(context.Background(), rpc.BlockNumber(height))
		if err != nil {
			log.Error("failed to get state and header for height", "height", height, "error", err)
			return nil, status.Error(codes.Internal, "Failed to get state and header for height")
		}

		evm := s.eth.APIBackend.GetEVM(context.Background(), &core.Message{GasPrice: big.NewInt(0)}, state, header, &vm.Config{NoBaseFee: true}, nil)
		code := evm.StateDB.GetCode(fork.Oracle.ContractAddress)
		if len(code) == 0 {
			log.Error("oracle contract address has no code", "address", fork.Oracle.ContractAddress)
			return nil, status.Error(codes.FailedPrecondition, "Oracle contract address has no code")
		}
	}

	res := &astriaPb.ExecutionSession{
		SessionId: s.activeSessionId,
		ExecutionSessionParameters: &astriaPb.ExecutionSessionParameters{
			RollupId:                         &rollupId,
			RollupStartBlockNumber:           fork.Height,
			RollupEndBlockNumber:             fork.StopHeight,
			SequencerChainId:                 fork.Sequencer.ChainID,
			SequencerStartBlockHeight:        fork.Sequencer.StartHeight,
			CelestiaChainId:                  fork.Celestia.ChainID,
			CelestiaSearchHeightMaxLookAhead: fork.Celestia.SearchHeightMaxLookAhead,
		},
		CommitmentState: &astriaPb.CommitmentState{
			SoftExecutedBlockMetadata:  softBlock,
			FirmExecutedBlockMetadata:  firmBlock,
			LowestCelestiaSearchHeight: max(s.bc.CurrentBaseCelestiaHeight(), fork.Celestia.StartHeight),
		},
	}

	log.Info("CreateExecutionSession completed", "response", res)
	createExecutionSessionSuccessCount.Inc(1)

	return res, nil
}
