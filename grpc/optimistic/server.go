package optimistic

import (
	optimisticGrpc "buf.build/gen/go/astria/execution-apis/grpc/go/astria/bundle/v1alpha1/bundlev1alpha1grpc"
	optimsticPb "buf.build/gen/go/astria/execution-apis/protocolbuffers/go/astria/bundle/v1alpha1"
	astriaPb "buf.build/gen/go/astria/execution-apis/protocolbuffers/go/astria/execution/v1"
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	cmath "github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/grpc/shared"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/miner"
	"github.com/ethereum/go-ethereum/params"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"math/big"
	"sync"
	"sync/atomic"
	"time"
)

type OptimisticServiceV1Alpha1 struct {
	optimisticGrpc.UnimplementedOptimisticExecutionServiceServer
	optimisticGrpc.UnimplementedBundleServiceServer

	sharedServiceContainer *shared.SharedServiceContainer

	currentOptimisticSequencerBlock atomic.Pointer[[]byte]
}

var (
	executeOptimisticBlockRequestCount = metrics.GetOrRegisterCounter("astria/optimistic/execute_optimistic_block_requests", nil)
	executeOptimisticBlockSuccessCount = metrics.GetOrRegisterCounter("astria/optimistic/execute_optimistic_block_success", nil)
	optimisticBlockHeight              = metrics.GetOrRegisterGauge("astria/execution/optimistic_block_height", nil)
	txsStreamedCount                   = metrics.GetOrRegisterCounter("astria/optimistic/txs_streamed", nil)

	executionOptimisticBlockTimer = metrics.GetOrRegisterTimer("astria/optimistic/execute_optimistic_block_time", nil)
)

func NewOptimisticServiceV1Alpha(sharedServiceContainer *shared.SharedServiceContainer) *OptimisticServiceV1Alpha1 {
	optimisticService := &OptimisticServiceV1Alpha1{
		sharedServiceContainer: sharedServiceContainer,
	}

	optimisticService.currentOptimisticSequencerBlock.Store(&[]byte{})

	return optimisticService
}

func (o *OptimisticServiceV1Alpha1) GetBundleStream(_ *optimsticPb.GetBundleStreamRequest, stream optimisticGrpc.BundleService_GetBundleStreamServer) error {
	log.Debug("GetBundleStream called")

	pendingTxEventCh := make(chan core.NewTxsEvent)
	pendingTxEvent := o.Eth().TxPool().SubscribeTransactions(pendingTxEventCh, false)
	defer pendingTxEvent.Unsubscribe()

	for {
		select {
		case pendingTxs := <-pendingTxEventCh:
			// get the optimistic block
			// this is an in-memory read, so there shouldn't be a lot of concerns on speed
			optimisticBlock := o.Eth().BlockChain().CurrentOptimisticBlock()

			for _, pendingTx := range pendingTxs.Txs {
				bundle := optimsticPb.Bundle{}

				totalCost := big.NewInt(0)
				effectiveTip := cmath.BigMin(pendingTx.GasTipCap(), new(big.Int).Sub(pendingTx.GasFeeCap(), optimisticBlock.BaseFee))
				totalCost.Add(totalCost, effectiveTip)

				marshalledTxs := [][]byte{}
				marshalledTx, err := pendingTx.MarshalBinary()
				if err != nil {
					return status.Errorf(codes.Internal, shared.WrapError(err, "error marshalling tx").Error())
				}
				marshalledTxs = append(marshalledTxs, marshalledTx)

				bundle.Fee = totalCost.Uint64()
				bundle.Transactions = marshalledTxs
				bundle.BaseSequencerBlockHash = *o.currentOptimisticSequencerBlock.Load()
				bundle.PrevRollupBlockHash = optimisticBlock.Hash().Bytes()

				txsStreamedCount.Inc(1)
				err = stream.Send(&optimsticPb.GetBundleStreamResponse{Bundle: &bundle})
				if err != nil {
					log.Error("error sending bundle over stream", "err", err)
					return status.Error(codes.Internal, shared.WrapError(err, "error sending bundle over stream").Error())
				}
			}

		case err := <-pendingTxEvent.Err():
			if err != nil {
				log.Error("error waiting for pending transactions", "err", err)
				return status.Error(codes.Internal, shared.WrapError(err, "error waiting for pending transactions").Error())
			} else {
				// TODO - what is the right error code here?
				return status.Error(codes.Internal, "tx pool subscription closed")
			}

		case <-stream.Context().Done():
			return stream.Context().Err()
		}
	}
}

func (o *OptimisticServiceV1Alpha1) ExecuteOptimisticBlockStream(stream optimisticGrpc.OptimisticExecutionService_ExecuteOptimisticBlockStreamServer) error {
	log.Debug("ExecuteOptimisticBlockStream called")

	mempoolClearingEventCh := make(chan core.NewMempoolCleared)
	mempoolClearingEvent := o.Eth().TxPool().SubscribeMempoolClearance(mempoolClearingEventCh)
	defer mempoolClearingEvent.Unsubscribe()

	for {
		msg, err := stream.Recv()
		// stream has been closed
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}

		executeOptimisticBlockRequestCount.Inc(1)

		baseBlock := msg.GetBaseBlock()

		// execute the optimistic block and wait for the mempool clearing event
		optimisticBlock, err := o.ExecuteOptimisticBlock(stream.Context(), baseBlock)
		if err != nil {
			return status.Errorf(codes.Internal, shared.WrapError(err, "failed to execute optimistic block").Error())
		}
		optimisticBlockHash := common.BytesToHash(optimisticBlock.Hash)

		// listen to the mempool clearing event and send the response back to the auctioneer when the mempool is cleared
		select {
		case event := <-mempoolClearingEventCh:
			if event.NewHead.Hash() != optimisticBlockHash {
				return status.Error(codes.Internal, "failed to clear mempool after optimistic block execution")
			}
			o.currentOptimisticSequencerBlock.Store(&baseBlock.SequencerBlockHash)
			executeOptimisticBlockSuccessCount.Inc(1)
			err = stream.Send(&optimsticPb.ExecuteOptimisticBlockStreamResponse{
				Block:                  optimisticBlock,
				BaseSequencerBlockHash: baseBlock.SequencerBlockHash,
			})
		case <-time.After(500 * time.Millisecond):
			log.Error("timed out waiting for mempool to clear after optimistic block execution")
			return status.Error(codes.DeadlineExceeded, "timed out waiting for mempool to clear after optimistic block execution")
		case err := <-mempoolClearingEvent.Err():
			if err != nil {
				log.Error("error waiting for mempool clearing event", "err", err)
				return status.Errorf(codes.Internal, shared.WrapError(err, "error waiting for mempool clearing event").Error())
			} else {
				// TODO - what is the right error code here?
				return status.Error(codes.Internal, "mempool clearance subscription closed")
			}
		case <-stream.Context().Done():
			return stream.Context().Err()
		}
	}
}

func (o *OptimisticServiceV1Alpha1) ExecuteOptimisticBlock(ctx context.Context, req *optimsticPb.BaseBlock) (*astriaPb.Block, error) {
	// we need to execute the optimistic block
	log.Debug("ExecuteOptimisticBlock called", "timestamp", req.Timestamp, "sequencer_block_hash", req.SequencerBlockHash)

	// Deliberately called after lock, to more directly measure the time spent executing
	executionStart := time.Now()
	defer executionOptimisticBlockTimer.UpdateSince(executionStart)

	if err := validateStaticExecuteOptimisticBlockRequest(req); err != nil {
		log.Error("ExecuteOptimisticBlock called with invalid BaseBlock", "err", err)
		return nil, status.Error(codes.InvalidArgument, shared.WrapError(err, "invalid BaseBlock").Error())
	}

	if !o.SyncMethodsCalled() {
		return nil, status.Error(codes.PermissionDenied, "Cannot execute block until GetGenesisInfo && GetCommitmentState methods are called")
	}

	softBlock := o.Bc().CurrentSafeBlock()

	nextFeeRecipient := o.NextFeeRecipient()

	// the height that this block will be at
	height := o.Bc().CurrentBlock().Number.Uint64() + 1

	addressPrefix := o.Bc().Config().AstriaSequencerAddressPrefix

	txsToProcess := shared.UnbundleRollupDataTransactions(req.Transactions, height, o.BridgeAddresses(), o.BridgeAllowedAssets(), softBlock.Hash().Bytes(), o.AuctioneerAddress(), addressPrefix)

	// Build a payload to add to the chain
	payloadAttributes := &miner.BuildPayloadArgs{
		Parent:                softBlock.Hash(),
		Timestamp:             uint64(req.GetTimestamp().GetSeconds()),
		Random:                common.Hash{},
		FeeRecipient:          nextFeeRecipient,
		OverrideTransactions:  txsToProcess,
		IsOptimisticExecution: true,
	}
	payload, err := o.Eth().Miner().BuildPayload(payloadAttributes)
	if err != nil {
		log.Error("failed to build payload", "err", err)
		return nil, status.Errorf(codes.InvalidArgument, shared.WrapError(err, "failed to build payload").Error())
	}

	block, err := engine.ExecutableDataToBlock(*payload.Resolve().ExecutionPayload, nil, nil)
	if err != nil {
		log.Error("failed to convert executable data to block", err)
		return nil, status.Error(codes.Internal, shared.WrapError(err, "failed to convert executable data to block").Error())
	}

	// this will insert the optimistic block into the chain and persist it's state without
	// setting it as the HEAD.
	err = o.Bc().InsertBlockWithoutSetHead(block)
	if err != nil {
		log.Error("failed to insert block to chain", "hash", block.Hash(), "prevHash", block.ParentHash(), "err", err)
		return nil, status.Error(codes.Internal, shared.WrapError(err, "failed to insert block to chain").Error())
	}

	// we store a pointer to the optimistic block in the chain so that we can use it
	// to retrieve the state of the optimistic block
	// this method also sends an event which indicates that a new optimistic block has been set
	// the mempool clearing logic is triggered when this event is received
	o.Bc().SetOptimistic(block)

	res := &astriaPb.Block{
		Number:          uint32(block.NumberU64()),
		Hash:            block.Hash().Bytes(),
		ParentBlockHash: block.ParentHash().Bytes(),
		Timestamp: &timestamppb.Timestamp{
			Seconds: int64(block.Time()),
		},
	}

	optimisticBlockHeight.Update(int64(block.NumberU64()))

	log.Info("ExecuteOptimisticBlock completed", "block_num", res.Number, "timestamp", res.Timestamp)

	return res, nil
}

func (o *OptimisticServiceV1Alpha1) Eth() *eth.Ethereum {
	return o.sharedServiceContainer.Eth()
}

func (o *OptimisticServiceV1Alpha1) Bc() *core.BlockChain {
	return o.sharedServiceContainer.Bc()
}

func (o *OptimisticServiceV1Alpha1) SetGenesisInfoCalled(value bool) {
	o.sharedServiceContainer.SetGenesisInfoCalled(value)
}

func (o *OptimisticServiceV1Alpha1) GenesisInfoCalled() bool {
	return o.sharedServiceContainer.GenesisInfoCalled()
}

func (o *OptimisticServiceV1Alpha1) SetGetCommitmentStateCalled(value bool) {
	o.sharedServiceContainer.SetGetCommitmentStateCalled(value)
}

func (o *OptimisticServiceV1Alpha1) CommitmentStateCalled() bool {
	return o.sharedServiceContainer.CommitmentStateCalled()
}

func (o *OptimisticServiceV1Alpha1) CommitmentUpdateLock() *sync.Mutex {
	return o.sharedServiceContainer.CommitmentUpdateLock()
}

func (o *OptimisticServiceV1Alpha1) BlockExecutionLock() *sync.Mutex {
	return o.sharedServiceContainer.BlockExecutionLock()
}

func (o *OptimisticServiceV1Alpha1) NextFeeRecipient() common.Address {
	return o.sharedServiceContainer.NextFeeRecipient()
}

func (o *OptimisticServiceV1Alpha1) SetNextFeeRecipient(feeRecipient common.Address) {
	o.sharedServiceContainer.SetNextFeeRecipient(feeRecipient)
}

func (s *OptimisticServiceV1Alpha1) BridgeAddresses() map[string]*params.AstriaBridgeAddressConfig {
	return s.sharedServiceContainer.BridgeAddresses()
}

func (s *OptimisticServiceV1Alpha1) BridgeAllowedAssets() map[string]struct{} {
	return s.sharedServiceContainer.BridgeAllowedAssets()
}

func (s *OptimisticServiceV1Alpha1) SyncMethodsCalled() bool {
	return s.sharedServiceContainer.SyncMethodsCalled()
}

func (s *OptimisticServiceV1Alpha1) AuctioneerAddress() string {
	return s.sharedServiceContainer.AuctioneerAddress()
}
