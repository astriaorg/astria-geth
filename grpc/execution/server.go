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

	astriaGrpc "buf.build/gen/go/astria/execution-apis/grpc/go/astria/execution/v1alpha2/executionv1alpha2grpc"
	astriaPb "buf.build/gen/go/astria/execution-apis/protocolbuffers/go/astria/execution/v1alpha2"
	primitivev1 "buf.build/gen/go/astria/primitives/protocolbuffers/go/astria/primitive/v1"
	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/miner"
	"github.com/ethereum/go-ethereum/params"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ExecutionServiceServerV1Alpha2 is the implementation of the
// ExecutionServiceServer interface.
type ExecutionServiceServerV1Alpha2 struct {
	// NOTE - from the generated code: All implementations must embed
	// UnimplementedExecutionServiceServer for forward compatibility
	astriaGrpc.UnimplementedExecutionServiceServer

	eth *eth.Ethereum
	bc  *core.BlockChain

	commitmentUpdateLock sync.Mutex // Lock for the forkChoiceUpdated method
	blockExecutionLock   sync.Mutex // Lock for the NewPayload method

	genesisInfoCalled        bool
	getCommitmentStateCalled bool

	bridgeAddresses     map[string]*params.AstriaBridgeAddressConfig // astria bridge addess to config for that bridge account
	bridgeAllowedAssets map[string]struct{}                          // a set of allowed asset IDs structs are left empty
	bridgeSenderAddress common.Address                               // address from which AstriaBridgeableERC20 contracts are called

	nextFeeRecipient common.Address // Fee recipient for the next block
}

var (
	getGenesisInfoRequestCount        = metrics.GetOrRegisterCounter("astria/execution/get_genesis_info_requests", nil)
	getGenesisInfoSuccessCount        = metrics.GetOrRegisterCounter("astria/execution/get_genesis_info_success", nil)
	getBlockRequestCount              = metrics.GetOrRegisterCounter("astria/execution/get_block_requests", nil)
	getBlockSuccessCount              = metrics.GetOrRegisterCounter("astria/execution/get_block_success", nil)
	batchGetBlockRequestCount         = metrics.GetOrRegisterCounter("astria/execution/batch_get_block_requests", nil)
	batchGetBlockSuccessCount         = metrics.GetOrRegisterCounter("astria/execution/batch_get_block_success", nil)
	executeBlockRequestCount          = metrics.GetOrRegisterCounter("astria/execution/execute_block_requests", nil)
	executeBlockSuccessCount          = metrics.GetOrRegisterCounter("astria/execution/execute_block_success", nil)
	getCommitmentStateRequestCount    = metrics.GetOrRegisterCounter("astria/execution/get_commitment_state_requests", nil)
	getCommitmentStateSuccessCount    = metrics.GetOrRegisterCounter("astria/execution/get_commitment_state_success", nil)
	updateCommitmentStateRequestCount = metrics.GetOrRegisterCounter("astria/execution/update_commitment_state_requests", nil)
	updateCommitmentStateSuccessCount = metrics.GetOrRegisterCounter("astria/execution/update_commitment_state_success", nil)

	softCommitmentHeight = metrics.GetOrRegisterGauge("astria/execution/soft_commitment_height", nil)
	firmCommitmentHeight = metrics.GetOrRegisterGauge("astria/execution/firm_commitment_height", nil)
	totalExecutedTxCount = metrics.GetOrRegisterCounter("astria/execution/total_executed_tx", nil)

	executeBlockTimer          = metrics.GetOrRegisterTimer("astria/execution/execute_block_time", nil)
	commitmentStateUpdateTimer = metrics.GetOrRegisterTimer("astria/execution/commitment", nil)
)

func NewExecutionServiceServerV1Alpha2(eth *eth.Ethereum) (*ExecutionServiceServerV1Alpha2, error) {
	bc := eth.BlockChain()

	if bc.Config().AstriaRollupName == "" {
		return nil, errors.New("rollup name not set")
	}

	if bc.Config().AstriaSequencerInitialHeight == 0 {
		return nil, errors.New("sequencer initial height not set")
	}

	if bc.Config().AstriaCelestiaInitialHeight == 0 {
		return nil, errors.New("celestia initial height not set")
	}

	if bc.Config().AstriaCelestiaHeightVariance == 0 {
		return nil, errors.New("celestia height variance not set")
	}

	bridgeAddresses := make(map[string]*params.AstriaBridgeAddressConfig)
	bridgeAllowedAssets := make(map[string]struct{})
	if bc.Config().AstriaBridgeAddressConfigs == nil {
		log.Warn("bridge addresses not set")
	} else {
		nativeBridgeSeen := false
		for _, cfg := range bc.Config().AstriaBridgeAddressConfigs {
			err := cfg.Validate(bc.Config().AstriaSequencerAddressPrefix)
			if err != nil {
				return nil, fmt.Errorf("invalid bridge address config: %w", err)
			}

			if cfg.Erc20Asset == nil {
				if nativeBridgeSeen {
					return nil, errors.New("only one native bridge address is allowed")
				}
				nativeBridgeSeen = true
			}

			if cfg.Erc20Asset != nil && bc.Config().AstriaBridgeSenderAddress == (common.Address{}) {
				return nil, errors.New("astria bridge sender address must be set for bridged ERC20 assets")
			}

			bridgeAddresses[cfg.BridgeAddress] = &cfg
			bridgeAllowedAssets[cfg.AssetDenom] = struct{}{}
			if cfg.Erc20Asset == nil {
				log.Info("bridge for sequencer native asset initialized", "bridgeAddress", cfg.BridgeAddress, "assetDenom", cfg.AssetDenom)
			} else {
				log.Info("bridge for ERC20 asset initialized", "bridgeAddress", cfg.BridgeAddress, "assetDenom", cfg.AssetDenom, "contractAddress", cfg.Erc20Asset.ContractAddress)
			}
		}
	}

	// To decrease compute cost, we identify the next fee recipient at the start
	// and update it as we execute blocks.
	nextFeeRecipient := common.Address{}
	if bc.Config().AstriaFeeCollectors == nil {
		log.Warn("fee asset collectors not set, assets will be burned")
	} else {
		maxHeightCollectorMatch := uint32(0)
		nextBlock := uint32(bc.CurrentBlock().Number.Int64()) + 1
		for height, collector := range bc.Config().AstriaFeeCollectors {
			if height <= nextBlock && height > maxHeightCollectorMatch {
				maxHeightCollectorMatch = height
				nextFeeRecipient = collector
			}
		}
	}

	return &ExecutionServiceServerV1Alpha2{
		eth:                 eth,
		bc:                  bc,
		bridgeAddresses:     bridgeAddresses,
		bridgeAllowedAssets: bridgeAllowedAssets,
		bridgeSenderAddress: bc.Config().AstriaBridgeSenderAddress,
		nextFeeRecipient:    nextFeeRecipient,
	}, nil
}

func (s *ExecutionServiceServerV1Alpha2) GetGenesisInfo(ctx context.Context, req *astriaPb.GetGenesisInfoRequest) (*astriaPb.GenesisInfo, error) {
	log.Debug("GetGenesisInfo called")
	getGenesisInfoRequestCount.Inc(1)

	rollupHash := sha256.Sum256([]byte(s.bc.Config().AstriaRollupName))
	rollupId := primitivev1.RollupId{Inner: rollupHash[:]}

	res := &astriaPb.GenesisInfo{
		RollupId:                    &rollupId,
		SequencerGenesisBlockHeight: s.bc.Config().AstriaSequencerInitialHeight,
		CelestiaBlockVariance:       s.bc.Config().AstriaCelestiaHeightVariance,
	}

	log.Info("GetGenesisInfo completed", "response", res)
	getGenesisInfoSuccessCount.Inc(1)
	s.genesisInfoCalled = true
	return res, nil
}

// GetBlock will return a block given an identifier.
func (s *ExecutionServiceServerV1Alpha2) GetBlock(ctx context.Context, req *astriaPb.GetBlockRequest) (*astriaPb.Block, error) {
	log.Debug("GetBlock called", "request", req)
	getBlockRequestCount.Inc(1)

	res, err := s.getBlockFromIdentifier(req.GetIdentifier())
	if err != nil {
		log.Error("failed finding block", err)
		return nil, err
	}

	log.Debug("GetBlock completed", "request", req, "response", res)
	getBlockSuccessCount.Inc(1)
	return res, nil
}

// BatchGetBlocks will return an array of Blocks given an array of block
// identifiers.
func (s *ExecutionServiceServerV1Alpha2) BatchGetBlocks(ctx context.Context, req *astriaPb.BatchGetBlocksRequest) (*astriaPb.BatchGetBlocksResponse, error) {
	log.Debug("BatchGetBlocks called", "first block", req.Identifiers[0], "last block", req.Identifiers[len(req.Identifiers)-1], "total blocks", len(req.Identifiers))
	batchGetBlockRequestCount.Inc(1)
	var blocks []*astriaPb.Block

	ids := req.GetIdentifiers()
	for _, id := range ids {
		block, err := s.getBlockFromIdentifier(id)
		if err != nil {
			log.Error("failed finding block with id", id, "error", err)
			return nil, err
		}

		blocks = append(blocks, block)
	}

	res := &astriaPb.BatchGetBlocksResponse{
		Blocks: blocks,
	}

	log.Info("BatchGetBlocks completed")
	batchGetBlockSuccessCount.Inc(1)
	return res, nil
}

func protoU128ToBigInt(u128 *primitivev1.Uint128) *big.Int {
	lo := big.NewInt(0).SetUint64(u128.Lo)
	hi := big.NewInt(0).SetUint64(u128.Hi)
	hi.Lsh(hi, 64)
	return lo.Add(lo, hi)
}

// ExecuteBlock drives deterministic derivation of a rollup block from sequencer
// block data
func (s *ExecutionServiceServerV1Alpha2) ExecuteBlock(ctx context.Context, req *astriaPb.ExecuteBlockRequest) (*astriaPb.Block, error) {
	log.Debug("ExecuteBlock called", "prevBlockHash", common.BytesToHash(req.PrevBlockHash), "tx_count", len(req.Transactions), "timestamp", req.Timestamp)
	executeBlockRequestCount.Inc(1)

	s.blockExecutionLock.Lock()
	defer s.blockExecutionLock.Unlock()
	// Deliberately called after lock, to more directly measure the time spent executing
	executionStart := time.Now()
	defer executeBlockTimer.UpdateSince(executionStart)

	if !s.syncMethodsCalled() {
		return nil, status.Error(codes.PermissionDenied, "Cannot execute block until GetGenesisInfo && GetCommitmentState methods are called")
	}

	// Validate block being created has valid previous hash
	prevHeadHash := common.BytesToHash(req.PrevBlockHash)
	softHash := s.bc.CurrentSafeBlock().Hash()
	if prevHeadHash != softHash {
		return nil, status.Error(codes.FailedPrecondition, "Block can only be created on top of soft block.")
	}

	// the height that this block will be at
	height := s.bc.CurrentBlock().Number.Uint64() + 1

	txsToProcess := types.Transactions{}
	for _, tx := range req.Transactions {
		unmarshalledTx, err := validateAndUnmarshalSequencerTx(height, tx, s.bridgeAddresses, s.bridgeAllowedAssets, s.bridgeSenderAddress)
		if err != nil {
			log.Debug("failed to validate sequencer tx, ignoring", "tx", tx, "err", err)
			continue
		}
		txsToProcess = append(txsToProcess, unmarshalledTx)
	}

	// This set of ordered TXs on the TxPool is has been configured to be used by
	// the Miner when building a payload.
	s.eth.TxPool().SetAstriaOrdered(txsToProcess)

	// Build a payload to add to the chain
	payloadAttributes := &miner.BuildPayloadArgs{
		Parent:       prevHeadHash,
		Timestamp:    uint64(req.GetTimestamp().GetSeconds()),
		Random:       common.Hash{},
		FeeRecipient: s.nextFeeRecipient,
	}
	payload, err := s.eth.Miner().BuildPayload(payloadAttributes)
	if err != nil {
		log.Error("failed to build payload", "err", err)
		return nil, status.Error(codes.InvalidArgument, "Could not build block with provided txs")
	}

	// call blockchain.InsertChain to actually execute and write the blocks to
	// state
	block, err := engine.ExecutableDataToBlock(*payload.Resolve().ExecutionPayload, nil, nil)
	if err != nil {
		log.Error("failed to convert executable data to block", err)
		return nil, status.Error(codes.Internal, "failed to execute block")
	}
	err = s.bc.InsertBlockWithoutSetHead(block)
	if err != nil {
		log.Error("failed to insert block to chain", "hash", block.Hash(), "prevHash", req.PrevBlockHash, "err", err)
		return nil, status.Error(codes.Internal, "failed to insert block to chain")
	}

	// remove txs from original mempool
	s.eth.TxPool().ClearAstriaOrdered()

	res := &astriaPb.Block{
		Number:          uint32(block.NumberU64()),
		Hash:            block.Hash().Bytes(),
		ParentBlockHash: block.ParentHash().Bytes(),
		Timestamp: &timestamppb.Timestamp{
			Seconds: int64(block.Time()),
		},
	}

	if next, ok := s.bc.Config().AstriaFeeCollectors[res.Number+1]; ok {
		s.nextFeeRecipient = next
	}

	log.Info("ExecuteBlock completed", "block_num", res.Number, "timestamp", res.Timestamp)
	totalExecutedTxCount.Inc(int64(len(block.Transactions())))
	executeBlockSuccessCount.Inc(1)
	return res, nil
}

// GetCommitmentState fetches the current CommitmentState of the chain.
func (s *ExecutionServiceServerV1Alpha2) GetCommitmentState(ctx context.Context, req *astriaPb.GetCommitmentStateRequest) (*astriaPb.CommitmentState, error) {
	log.Info("GetCommitmentState called")
	getCommitmentStateRequestCount.Inc(1)

	softBlock, err := ethHeaderToExecutionBlock(s.bc.CurrentSafeBlock())
	if err != nil {
		log.Error("error finding safe block", err)
		return nil, status.Error(codes.Internal, "could not locate soft block")
	}
	firmBlock, err := ethHeaderToExecutionBlock(s.bc.CurrentFinalBlock())
	if err != nil {
		log.Error("error finding final block", err)
		return nil, status.Error(codes.Internal, "could not locate firm block")
	}

	celestiaBlock := s.bc.CurrentBaseCelestiaHeight()

	res := &astriaPb.CommitmentState{
		Soft:               softBlock,
		Firm:               firmBlock,
		BaseCelestiaHeight: celestiaBlock,
	}

	log.Info("GetCommitmentState completed", "soft_height", res.Soft.Number, "firm_height", res.Firm.Number, "base_celestia_height", res.BaseCelestiaHeight)
	getCommitmentStateSuccessCount.Inc(1)
	s.getCommitmentStateCalled = true
	return res, nil
}

// UpdateCommitmentState replaces the whole CommitmentState with a new
// CommitmentState.
func (s *ExecutionServiceServerV1Alpha2) UpdateCommitmentState(ctx context.Context, req *astriaPb.UpdateCommitmentStateRequest) (*astriaPb.CommitmentState, error) {
	log.Debug("UpdateCommitmentState called", "request_soft_height", req.CommitmentState.Soft.Number, "request_firm_height", req.CommitmentState.Firm.Number)
	updateCommitmentStateRequestCount.Inc(1)
	commitmentUpdateStart := time.Now()
	defer commitmentStateUpdateTimer.UpdateSince(commitmentUpdateStart)

	s.commitmentUpdateLock.Lock()
	defer s.commitmentUpdateLock.Unlock()

	if !s.syncMethodsCalled() {
		return nil, status.Error(codes.PermissionDenied, "Cannot update commitment state until GetGenesisInfo && GetCommitmentState methods are called")
	}

	if s.bc.CurrentBaseCelestiaHeight() > req.CommitmentState.BaseCelestiaHeight {
		errStr := fmt.Sprintf("Base Celestia height cannot be decreased, current_base_celestia_height: %d, new_base_celestia_height: %d", s.bc.CurrentBaseCelestiaHeight(), req.CommitmentState.BaseCelestiaHeight)
		return nil, status.Error(codes.InvalidArgument, errStr)
	}

	softEthHash := common.BytesToHash(req.CommitmentState.Soft.Hash)
	firmEthHash := common.BytesToHash(req.CommitmentState.Firm.Hash)

	// Validate that the firm and soft blocks exist before going further
	softBlock := s.bc.GetBlockByHash(softEthHash)
	if softBlock == nil {
		return nil, status.Error(codes.InvalidArgument, "Soft block specified does not exist")
	}
	firmBlock := s.bc.GetBlockByHash(firmEthHash)
	if firmBlock == nil {
		return nil, status.Error(codes.InvalidArgument, "Firm block specified does not exist")
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
		s.bc.SetCelestiaFinalized(firmBlock.Header(), req.CommitmentState.BaseCelestiaHeight)
	}

	log.Info("UpdateCommitmentState completed", "soft_height", softBlock.NumberU64(), "firm_height", firmBlock.NumberU64())
	softCommitmentHeight.Update(int64(softBlock.NumberU64()))
	firmCommitmentHeight.Update(int64(firmBlock.NumberU64()))
	updateCommitmentStateSuccessCount.Inc(1)
	return req.CommitmentState, nil
}

func (s *ExecutionServiceServerV1Alpha2) getBlockFromIdentifier(identifier *astriaPb.BlockIdentifier) (*astriaPb.Block, error) {
	var header *types.Header

	// Grab the header based on the identifier provided
	switch idType := identifier.Identifier.(type) {
	case *astriaPb.BlockIdentifier_BlockNumber:
		header = s.bc.GetHeaderByNumber(uint64(identifier.GetBlockNumber()))
	case *astriaPb.BlockIdentifier_BlockHash:
		header = s.bc.GetHeaderByHash(common.BytesToHash(identifier.GetBlockHash()))
	default:
		return nil, status.Errorf(codes.InvalidArgument, "identifier has unexpected type %T", idType)
	}

	if header == nil {
		return nil, status.Errorf(codes.NotFound, "Couldn't locate block with identifier %s", identifier.Identifier)
	}

	res, err := ethHeaderToExecutionBlock(header)
	if err != nil {
		// This should never happen since we validate header exists above.
		return nil, status.Error(codes.Internal, "internal error")
	}

	return res, nil
}

func ethHeaderToExecutionBlock(header *types.Header) (*astriaPb.Block, error) {
	if header == nil {
		return nil, fmt.Errorf("cannot convert nil header to execution block")
	}

	return &astriaPb.Block{
		Number:          uint32(header.Number.Int64()),
		Hash:            header.Hash().Bytes(),
		ParentBlockHash: header.ParentHash.Bytes(),
		Timestamp: &timestamppb.Timestamp{
			Seconds: int64(header.Time),
		},
	}, nil
}

func (s *ExecutionServiceServerV1Alpha2) syncMethodsCalled() bool {
	return s.genesisInfoCalled && s.getCommitmentStateCalled
}
