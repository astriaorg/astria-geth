// Package execution provides the gRPC server for the execution layer.
//
// Its procedures will be called from the conductor. It is responsible
// for immediately executing lists of ordered transactions that come from the shared sequencer.
package execution

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	astriaGrpc "buf.build/gen/go/astria/execution-apis/grpc/go/astria/execution/v2/executionv2grpc"
	astriaPb "buf.build/gen/go/astria/execution-apis/protocolbuffers/go/astria/execution/v2"
	primitivev1 "buf.build/gen/go/astria/primitives/protocolbuffers/go/astria/primitive/v1"
	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/miner"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ExecutionServiceServerV2 is the implementation of the
// ExecutionServiceServer interface.
type ExecutionServiceServerV2 struct {
	// NOTE - from the generated code: All implementations must embed
	// UnimplementedExecutionServiceServer for forward compatibility
	astriaGrpc.UnimplementedExecutionServiceServer

	eth *eth.Ethereum
	bc  *core.BlockChain

	commitmentUpdateLock sync.Mutex // Lock for the forkChoiceUpdated method
	blockExecutionLock   sync.Mutex // Lock for the NewPayload method

	softAsFirm softAsFirmConfig

	activeSessionId string
	activeFork      *params.AstriaForkData
}

var (
	createExecutionSessionRequestCount   = metrics.GetOrRegisterCounter("astria/execution/create_execution_session_requests", nil)
	createExecutionSessionSuccessCount   = metrics.GetOrRegisterCounter("astria/execution/create_execution_session_success", nil)
	getExecutedBlockMetadataRequestCount = metrics.GetOrRegisterCounter("astria/execution/get_executed_block_metadata_requests", nil)
	getExecutedBlockMetadataSuccessCount = metrics.GetOrRegisterCounter("astria/execution/get_executed_block_metadata_success", nil)
	executeBlockRequestCount             = metrics.GetOrRegisterCounter("astria/execution/execute_block_requests", nil)
	executeBlockSuccessCount             = metrics.GetOrRegisterCounter("astria/execution/execute_block_success", nil)
	updateCommitmentStateRequestCount    = metrics.GetOrRegisterCounter("astria/execution/update_commitment_state_requests", nil)
	updateCommitmentStateSuccessCount    = metrics.GetOrRegisterCounter("astria/execution/update_commitment_state_success", nil)

	softCommitmentHeight = metrics.GetOrRegisterGauge("astria/execution/soft_commitment_height", nil)
	firmCommitmentHeight = metrics.GetOrRegisterGauge("astria/execution/firm_commitment_height", nil)
	totalExecutedTxCount = metrics.GetOrRegisterCounter("astria/execution/total_executed_tx", nil)

	executeBlockTimer          = metrics.GetOrRegisterTimer("astria/execution/execute_block_time", nil)
	commitmentStateUpdateTimer = metrics.GetOrRegisterTimer("astria/execution/commitment", nil)
)

func NewExecutionServiceServerV2(eth *eth.Ethereum, softAsFirm bool, softAsFirmMaxHeight uint64) (*ExecutionServiceServerV2, error) {
	bc := eth.BlockChain()

	if bc.Config().AstriaRollupName == "" {
		return nil, errors.New("rollup name not set")
	}

	softAsFirmConfig := softAsFirmConfig{
		enabled:   softAsFirm,
		maxHeight: softAsFirmMaxHeight,
	}
	return &ExecutionServiceServerV2{
		eth: eth,
		bc:  bc,

		softAsFirm: softAsFirmConfig,
	}, nil
}

func (s *ExecutionServiceServerV2) CreateExecutionSession(ctx context.Context, req *astriaPb.CreateExecutionSessionRequest) (*astriaPb.ExecutionSession, error) {
	log.Debug("CreateExecutionSession called")
	createExecutionSessionRequestCount.Inc(1)

	// We shouldn't create a new session if we are actively executing within one.
	s.blockExecutionLock.Lock()
	defer s.blockExecutionLock.Unlock()
	s.commitmentUpdateLock.Lock()
	defer s.commitmentUpdateLock.Unlock()

	rollupHash := sha256.Sum256([]byte(s.bc.Config().AstriaRollupName))
	rollupId := primitivev1.RollupId{Inner: rollupHash[:]}

	fork := s.bc.Config().GetAstriaForks().GetForkAtHeight(s.bc.CurrentFinalBlock().Number.Uint64() + 1)

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

func (s *ExecutionServiceServerV2) GetExecutedBlockMetadata(ctx context.Context, req *astriaPb.GetExecutedBlockMetadataRequest) (*astriaPb.ExecutedBlockMetadata, error) {
	log.Debug("GetExecutedBlockMetadata called", "request", req)
	getExecutedBlockMetadataRequestCount.Inc(1)

	if req.GetIdentifier() == nil {
		return nil, status.Error(codes.InvalidArgument, "block identifier cannot be empty")
	}

	res, err := s.getExecutedBlockMetadataFromIdentifier(req.GetIdentifier())
	if err != nil {
		log.Error("failed finding block", err)
		return nil, err
	}

	log.Debug("GetExecutedBlockMetadata completed", "request", req, "response", res)
	getExecutedBlockMetadataSuccessCount.Inc(1)

	return res, nil
}

func protoU128ToBigInt(u128 *primitivev1.Uint128) *big.Int {
	lo := big.NewInt(0).SetUint64(u128.Lo)
	hi := big.NewInt(0).SetUint64(u128.Hi)
	hi.Lsh(hi, 64)
	return lo.Add(lo, hi)
}

func protoI128ToBigInt(i128 *primitivev1.Int128) *big.Int {
	lo := big.NewInt(0).SetUint64(i128.Lo)
	hi := big.NewInt(0).SetUint64(i128.Hi)
	hi.Lsh(hi, 64)
	return lo.Add(lo, hi)
}

type conversionConfig struct {
	bridgeAddresses       map[string]*params.AstriaBridgeAddressConfig
	bridgeAllowedAssets   map[string]struct{}
	api                   *eth.EthAPIBackend
	oracleContractAddress common.Address
	oracleCallerAddress   common.Address
}

// ExecuteBlock drives deterministic derivation of a rollup block from sequencer
// block data
func (s *ExecutionServiceServerV2) ExecuteBlock(ctx context.Context, req *astriaPb.ExecuteBlockRequest) (*astriaPb.ExecuteBlockResponse, error) {
	log.Debug("ExecuteBlock called", "parentHash", req.ParentHash, "tx_count", len(req.Transactions), "timestamp", req.Timestamp)
	executeBlockRequestCount.Inc(1)

	s.blockExecutionLock.Lock()
	defer s.blockExecutionLock.Unlock()

	// Deliberately called after lock, to more directly measure the time spent executing
	executionStart := time.Now()
	defer executeBlockTimer.UpdateSince(executionStart)

	// Check for valid session first
	if s.activeSessionId == "" || req.GetSessionId() != s.activeSessionId {
		return nil, status.Error(codes.PermissionDenied, "Cannot execute block until a valid ExecutionSession is created")
	}

	// Then validate the details of the request
	if err := validateStaticExecuteBlockRequest(req); err != nil {
		log.Error("ExecuteBlock called with invalid ExecuteBlockRequest", "err", err)
		return nil, status.Error(codes.InvalidArgument, "ExecuteBlockRequest is invalid")
	}

	// this fork halts the chain
	if s.activeFork.Halt {
		return nil, status.Error(codes.FailedPrecondition, "Block cannot be created at halted fork")
	}

	// the height that this block will be at
	height := s.bc.CurrentBlock().Number.Uint64() + 1

	// Session is out of range
	// If StopHeight is 0, there is no upper limit
	if (s.activeFork.StopHeight > 0 && height > s.activeFork.StopHeight) || height < s.activeFork.Height {
		return nil, status.Error(codes.OutOfRange, "Session is out of range")
	}

	// Validate block being created has valid previous hash
	prevHeadHash := common.HexToHash(req.ParentHash)
	softHash := s.bc.CurrentSafeBlock().Hash()
	if prevHeadHash != softHash {
		return nil, status.Error(codes.FailedPrecondition, "Block can only be created on top of soft block.")
	}

	blockTimestamp := uint64(req.GetTimestamp().GetSeconds())
	var sequencerHashRef *common.Hash
	if s.bc.Config().IsCancun(big.NewInt(int64(height)), blockTimestamp) {
		if req.SequencerBlockHash == "" {
			return nil, status.Error(codes.InvalidArgument, "Sequencer block hash must be set for Cancun block")
		}
		sequencerHash := common.HexToHash(req.SequencerBlockHash)
		sequencerHashRef = &sequencerHash
	}

	txsToProcess := types.Transactions{}
	txsByType := map[params.AstriaTransactionType][]*types.Transaction{}
	conversionConfig := &conversionConfig{
		bridgeAddresses:       s.activeFork.BridgeAddresses,
		bridgeAllowedAssets:   s.activeFork.BridgeAllowedAssets,
		api:                   s.eth.APIBackend,
		oracleContractAddress: s.activeFork.Oracle.ContractAddress,
		oracleCallerAddress:   s.activeFork.Oracle.CallerAddress,
	}

	for _, tx := range req.Transactions {
		astriaTx, err := validateAndConvertSequencerTx(ctx, height, tx, conversionConfig)
		if err != nil {
			log.Info("failed to validate sequencer tx, ignoring", "tx", tx, "err", err)
			continue
		}
		if s.activeFork.AppSpecificOrdering == nil {
			txsToProcess = append(txsToProcess, astriaTx.Transactions...)
		} else {
			txsByType[astriaTx.TransactionType] = append(txsByType[astriaTx.TransactionType], astriaTx.Transactions...)
		}
	}
	if s.activeFork.AppSpecificOrdering != nil {
		for _, txType := range s.activeFork.AppSpecificOrdering {
			txsToProcess = append(txsToProcess, txsByType[txType]...)
		}
	}

	// This set of ordered TXs on the TxPool is used when building a payload.
	s.eth.TxPool().SetAstriaOrdered(txsToProcess)

	// Set extra data deterministically based on the current block and genesis.
	err := s.eth.Miner().SetExtra(s.activeFork.ExtraData)
	if err != nil {
		log.Error("failed to set extra data", "err", err)
		return nil, status.Error(codes.Internal, "could not set extra data")
	}

	// Build a payload to add to the chain
	payloadAttributes := &miner.BuildPayloadArgs{
		Parent:       prevHeadHash,
		Timestamp:    uint64(req.GetTimestamp().GetSeconds()),
		Random:       common.Hash{},
		FeeRecipient: s.activeFork.FeeCollector,
		BeaconRoot:   sequencerHashRef,
	}
	payload, err := s.eth.Miner().BuildPayload(payloadAttributes)
	if err != nil {
		log.Error("failed to build payload", "err", err)
		return nil, status.Error(codes.InvalidArgument, "Could not build block with provided txs")
	}

	// call blockchain.InsertChain to actually execute and write the blocks to
	// state
	block, err := engine.ExecutableDataToBlock(*payload.Resolve().ExecutionPayload, nil, sequencerHashRef)
	if err != nil {
		log.Error("failed to convert executable data to block", err)
		return nil, status.Error(codes.Internal, "failed to execute block")
	}

	err = s.bc.InsertBlockWithoutSetHead(block)
	if err != nil {
		log.Error("failed to insert block to chain", "hash", block.Hash(), "parentHash", req.ParentHash, "err", err)
		return nil, status.Error(codes.Internal, "failed to insert block to chain")
	}

	// Reset the building pool, this also clears out any excluded from txs from the main mempool
	s.eth.TxPool().ClearAstriaOrdered()

	resBlockMetadata, _ := ethHeaderToExecutedBlockMetadata(block.Header())
	res := &astriaPb.ExecuteBlockResponse{
		ExecutedBlockMetadata: resBlockMetadata,
	}

	log.Info("ExecuteBlock completed", "block_num", res.ExecutedBlockMetadata.Number, "timestamp", res.ExecutedBlockMetadata.Timestamp)
	totalExecutedTxCount.Inc(int64(len(block.Transactions())))
	executeBlockSuccessCount.Inc(1)
	return res, nil
}

// UpdateCommitmentState replaces the whole CommitmentState with a new
// CommitmentState.
func (s *ExecutionServiceServerV2) UpdateCommitmentState(ctx context.Context, req *astriaPb.UpdateCommitmentStateRequest) (*astriaPb.CommitmentState, error) {
	log.Debug("UpdateCommitmentState called", "request_soft_height", req.CommitmentState.SoftExecutedBlockMetadata.Number, "request_firm_height", req.CommitmentState.FirmExecutedBlockMetadata.Number)
	updateCommitmentStateRequestCount.Inc(1)

	s.commitmentUpdateLock.Lock()
	defer s.commitmentUpdateLock.Unlock()

	if err := validateStaticCommitmentState(req.CommitmentState); err != nil {
		log.Error("UpdateCommitmentState called with invalid CommitmentState", "err", err)
		return nil, status.Error(codes.InvalidArgument, "CommitmentState is invalid")
	}

	if s.activeSessionId == "" || req.GetSessionId() != s.activeSessionId {
		return nil, status.Error(codes.PermissionDenied, "Cannot update commitment state until a valid ExecutionSession is created")
	}

	// Soft commitment is out of range
	// If StopHeight is 0, there is no upper limit
	if (s.activeFork.StopHeight > 0 && req.CommitmentState.SoftExecutedBlockMetadata.Number > s.activeFork.StopHeight) || req.CommitmentState.SoftExecutedBlockMetadata.Number < s.activeFork.Height {
		return nil, status.Error(codes.OutOfRange, "Soft commitment is out of range")
	}

	// If softAsFirm is true, firm commitment state is ignored. If the firm commitment
	// state is unchanged, we assume the stored firm block is correct and do not
	// perform these height checks.
	softAsFirm := s.softAsFirm.useHeightAsFirm(req.CommitmentState.SoftExecutedBlockMetadata.Number)
	if !softAsFirm && (req.CommitmentState.FirmExecutedBlockMetadata.Number != s.bc.CurrentFinalBlock().Number.Uint64()) {
		// Firm commitment is out of range
		// If StopHeight is 0, there is no upper limit
		if s.activeFork.StopHeight > 0 && req.CommitmentState.FirmExecutedBlockMetadata.Number > s.activeFork.StopHeight {
			return nil, status.Error(codes.OutOfRange, fmt.Sprintf("Firm number %d is greater than stop height %d", req.CommitmentState.FirmExecutedBlockMetadata.Number, s.activeFork.StopHeight))
		}
		// The firm commitment must be greater than or equal to the fork height
		if req.CommitmentState.FirmExecutedBlockMetadata.Number < s.activeFork.Height {
			return nil, status.Error(codes.OutOfRange, fmt.Sprintf("Firm number %d is less than current fork height %d", req.CommitmentState.FirmExecutedBlockMetadata.Number, s.activeFork.Height))
		}
	}

	commitmentUpdateStart := time.Now()
	defer commitmentStateUpdateTimer.UpdateSince(commitmentUpdateStart)

	if s.bc.CurrentBaseCelestiaHeight() > req.CommitmentState.LowestCelestiaSearchHeight {
		errStr := fmt.Sprintf("Base Celestia height cannot be decreased, current_base_celestia_height: %d, new_base_celestia_height: %d", s.bc.CurrentBaseCelestiaHeight(), req.CommitmentState.LowestCelestiaSearchHeight)
		return nil, status.Error(codes.InvalidArgument, errStr)
	}

	softEthHash := common.HexToHash(req.CommitmentState.SoftExecutedBlockMetadata.Hash)

	var firmEthHash common.Hash
	if softAsFirm {
		firmEthHash = softEthHash
	} else {
		firmEthHash = common.HexToHash(req.CommitmentState.FirmExecutedBlockMetadata.Hash)
	}

	// Validate that the firm and soft blocks exist before going further
	softBlock := s.bc.GetBlockByHash(softEthHash)
	if softBlock == nil {
		return nil, status.Error(codes.InvalidArgument, "Soft block specified does not exist")
	}
	if softBlock.NumberU64() != req.CommitmentState.SoftExecutedBlockMetadata.Number {
		return nil, status.Error(codes.InvalidArgument, "Soft block number specified does not match the block number identified by hash")
	}
	if softBlock.ParentHash() != common.HexToHash(req.CommitmentState.SoftExecutedBlockMetadata.ParentHash) {
		return nil, status.Error(codes.InvalidArgument, "Soft block parent hash specified does not match the block parent hash identified by hash")
	}

	firmBlock := s.bc.GetBlockByHash(firmEthHash)
	if firmBlock == nil {
		return nil, status.Error(codes.InvalidArgument, "Firm block specified does not exist")
	}
	// Only validate matches if softAsFirm is false, otherwise we would expect will match the soft block and already verified.
	if !softAsFirm {
		if firmBlock.NumberU64() != req.CommitmentState.FirmExecutedBlockMetadata.Number {
			return nil, status.Error(codes.InvalidArgument, "Firm block number specified does not match the block number identified by hash")
		}
		if firmBlock.ParentHash() != common.HexToHash(req.CommitmentState.FirmExecutedBlockMetadata.ParentHash) {
			return nil, status.Error(codes.InvalidArgument, "Firm block parent hash specified does not match the block parent hash identified by hash")
		}
	}

	currentHead := s.bc.CurrentBlock().Hash()

	// Update the canonical chain to soft block. We must do this before last
	// validation step since there is no way to check if firm block descends from
	// anything but the canonical chain
	if currentHead != softEthHash {
		if _, err := s.bc.SetCanonical(softBlock); err != nil {
			log.Error("failed updating canonical chain to soft block", err)
			return nil, status.Error(codes.Internal, "Could not update head to safe hash")
		}
	}

	// Once head is updated validate that firm belongs to chain
	rollbackBlock := s.bc.GetBlockByHash(currentHead)
	if s.bc.GetCanonicalHash(firmBlock.NumberU64()) != firmEthHash {
		log.Error("firm block not found in canonical chain defined by soft block, rolling back")

		if _, err := s.bc.SetCanonical(rollbackBlock); err != nil {
			panic("rollback to previous head after failed validation failed")
		}

		return nil, status.Error(codes.InvalidArgument, "soft block in request is not a descendant of the current firmly committed block")
	}

	s.eth.SetSynced()

	// Updating the safe and final after everything validated
	currentSafe := s.bc.CurrentSafeBlock().Hash()
	if currentSafe != softEthHash {
		s.bc.SetSafe(softBlock.Header())
	}

	currentFirm := s.bc.CurrentFinalBlock().Hash()
	if currentFirm != firmEthHash {
		s.bc.SetCelestiaFinalized(firmBlock.Header(), req.CommitmentState.LowestCelestiaSearchHeight)
	}

	log.Info("UpdateCommitmentState completed", "soft_height", softBlock.NumberU64(), "firm_height", firmBlock.NumberU64())
	softCommitmentHeight.Update(int64(softBlock.NumberU64()))
	firmCommitmentHeight.Update(int64(firmBlock.NumberU64()))
	updateCommitmentStateSuccessCount.Inc(1)
	return req.CommitmentState, nil
}

func (s *ExecutionServiceServerV2) getExecutedBlockMetadataFromIdentifier(identifier *astriaPb.ExecutedBlockIdentifier) (*astriaPb.ExecutedBlockMetadata, error) {
	var header *types.Header

	// Grab the header based on the identifier provided
	switch idType := identifier.Identifier.(type) {
	case *astriaPb.ExecutedBlockIdentifier_Number:
		header = s.bc.GetHeaderByNumber(identifier.GetNumber())
	case *astriaPb.ExecutedBlockIdentifier_Hash:
		header = s.bc.GetHeaderByHash(common.HexToHash(identifier.GetHash()))
	default:
		return nil, status.Errorf(codes.InvalidArgument, "identifier has unexpected type %T", idType)
	}

	if header == nil {
		return nil, status.Errorf(codes.NotFound, "Couldn't locate block with identifier %s", identifier.Identifier)
	}

	res, err := ethHeaderToExecutedBlockMetadata(header)
	if err != nil {
		// This should never happen since we validate header exists above.
		return nil, status.Error(codes.Internal, "internal error")
	}

	return res, nil
}

func ethHeaderToExecutedBlockMetadata(header *types.Header) (*astriaPb.ExecutedBlockMetadata, error) {
	if header == nil {
		return nil, fmt.Errorf("cannot convert nil header to executed block metadata")
	}

	var sequencerHash common.Hash
	if header.ParentBeaconRoot != nil {
		sequencerHash = *header.ParentBeaconRoot
	}

	return &astriaPb.ExecutedBlockMetadata{
		Number:             header.Number.Uint64(),
		Hash:               header.Hash().Hex(),
		ParentHash:         header.ParentHash.Hex(),
		SequencerBlockHash: sequencerHash.Hex(),
		Timestamp: &timestamppb.Timestamp{
			Seconds: int64(header.Time),
		},
	}, nil
}

type softAsFirmConfig struct {
	enabled   bool
	maxHeight uint64
}

func (sfc *softAsFirmConfig) useHeightAsFirm(blockNum uint64) bool {
	if !sfc.enabled {
		return false
	}

	return sfc.maxHeight == 0 || blockNum <= sfc.maxHeight
}
