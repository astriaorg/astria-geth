// Package execution provides the gRPC server for the execution layer.
//
// Its procedures will be called from the conductor. It is responsible
// for immediately executing lists of ordered transactions that come from the shared sequencer.
package execution

import (
	"context"
	"fmt"
	"sync"

	astriaGrpc "buf.build/gen/go/astria/astria/grpc/go/astria/execution/v1alpha2/executionv1alpha2grpc"
	astriaPb "buf.build/gen/go/astria/astria/protocolbuffers/go/astria/execution/v1alpha2"
	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/miner"
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

	commitementUpdateLock sync.Mutex // Lock for the forkChoiceUpdated method
	blockExecutionLock    sync.Mutex // Lock for the NewPayload method
}

func NewExecutionServiceServerV1Alpha2(eth *eth.Ethereum) *ExecutionServiceServerV1Alpha2 {
	bc := eth.BlockChain()

	if merger := eth.Merger(); !merger.PoSFinalized() {
		merger.FinalizePoS()
	}

	return &ExecutionServiceServerV1Alpha2{
		eth: eth,
		bc:  bc,
	}
}

// GetBlock will return a block given an identifier.
func (s *ExecutionServiceServerV1Alpha2) GetBlock(ctx context.Context, req *astriaPb.GetBlockRequest) (*astriaPb.Block, error) {
	log.Info("GetBlock called", "request", req)

	res, err := s.getBlockFromIdentifier(req.GetIdentifier())
	if err != nil {
		log.Error("failed finding block", err)
		return nil, err
	}

	log.Info("GetBlock completed", "request", req, "response", res)
	return res, nil
}

// BatchGetBlocks will return an array of Blocks given an array of block
// identifiers.
func (s *ExecutionServiceServerV1Alpha2) BatchGetBlocks(ctx context.Context, req *astriaPb.BatchGetBlocksRequest) (*astriaPb.BatchGetBlocksResponse, error) {
	log.Info("BatchGetBlocks called", "request", req)
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

	log.Info("BatchGetBlocks completed", "request", req, "response", res)
	return res, nil
}

// ExecuteBlock drives deterministic derivation of a rollup block from sequencer
// block data
func (s *ExecutionServiceServerV1Alpha2) ExecuteBlock(ctx context.Context, req *astriaPb.ExecuteBlockRequest) (*astriaPb.Block, error) {
	log.Info("ExecuteBlock called", "request", req)

	s.blockExecutionLock.Lock()
	defer s.blockExecutionLock.Unlock()

	// Validate block being created has valid previous hash
	prevHeadHash := common.BytesToHash(req.PrevBlockHash)
	softHash := s.bc.CurrentSafeBlock().Hash()
	if prevHeadHash != softHash {
		return nil, status.Error(codes.FailedPrecondition, "Block can only be created on top of soft block.")
	}

	// This set of ordered TXs on the TxPool is has been configured to be used by
	// the Miner when building a payload.
	s.eth.TxPool().SetAstriaOrdered(req.Transactions)

	// Build a payload to add to the chain
	payloadAttributes := &miner.BuildPayloadArgs{
		Parent:       prevHeadHash,
		Timestamp:    uint64(req.GetTimestamp().GetSeconds()),
		Random:       common.Hash{},
		FeeRecipient: common.Address{},
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
		Number: uint32(block.NumberU64()),
		Hash:   block.Hash().Bytes(),
		ParentBlockHash: block.ParentHash().Bytes(),
		Timestamp: &timestamppb.Timestamp{
			Seconds: int64(block.Time()),
		},
	}

	log.Info("ExecuteBlock completed", "request", req, "response", res)
	return res, nil
}

// GetCommitmentState fetches the current CommitmentState of the chain.
func (s *ExecutionServiceServerV1Alpha2) GetCommitmentState(ctx context.Context, req *astriaPb.GetCommitmentStateRequest) (*astriaPb.CommitmentState, error) {
	log.Info("GetCommitmentState called", "request", req)

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

	res := &astriaPb.CommitmentState{
		Soft: softBlock,
		Firm: firmBlock,
	}

	log.Info("GetCommitmentState completed", "request", req, "response", res)
	return res, nil
}

// UpdateCommitmentState replaces the whole CommitmentState with a new
// CommitmentState.
func (s *ExecutionServiceServerV1Alpha2) UpdateCommitmentState(ctx context.Context, req *astriaPb.UpdateCommitmentStateRequest) (*astriaPb.CommitmentState, error) {
	log.Info("UpdateCommitmentState called", "request", req)

	s.commitementUpdateLock.Lock()
	defer s.commitementUpdateLock.Unlock()

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
		s.bc.SetFinalized(firmBlock.Header())
	}

	log.Info("UpdateCommitmentState completed", "request", req)
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
