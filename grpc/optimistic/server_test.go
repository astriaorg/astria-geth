package optimistic

import (
	optimsticPb "buf.build/gen/go/astria/execution-apis/protocolbuffers/go/astria/bundle/v1alpha1"
	astriaPb "buf.build/gen/go/astria/execution-apis/protocolbuffers/go/astria/execution/v1alpha2"
	primitivev1 "buf.build/gen/go/astria/primitives/protocolbuffers/go/astria/primitive/v1"
	sequencerblockv1alpha1 "buf.build/gen/go/astria/sequencerblock-apis/protocolbuffers/go/astria/sequencerblock/v1alpha1"
	"bytes"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/grpc/execution"
	"github.com/ethereum/go-ethereum/grpc/shared"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math/big"
	"testing"
	"time"
)

func TestExecutionServiceServerV1Alpha2_ExecuteOptimisticBlock(t *testing.T) {
	ethService, _ := shared.SetupSharedService(t, 10)

	tests := []struct {
		description                          string
		callGenesisInfoAndGetCommitmentState bool
		numberOfTxs                          int
		prevBlockHash                        []byte
		timestamp                            uint64
		depositTxAmount                      *big.Int // if this is non zero then we send a deposit tx
		expectedReturnCode                   codes.Code
	}{
		{
			description:                          "ExecuteOptimisticBlock without calling GetGenesisInfo and GetCommitmentState",
			callGenesisInfoAndGetCommitmentState: false,
			numberOfTxs:                          5,
			prevBlockHash:                        ethService.BlockChain().GetBlockByNumber(2).Hash().Bytes(),
			timestamp:                            ethService.BlockChain().GetBlockByNumber(2).Time() + 2,
			depositTxAmount:                      big.NewInt(0),
			expectedReturnCode:                   codes.PermissionDenied,
		},
		{
			description:                          "ExecuteOptimisticBlock with 5 txs and no deposit tx",
			callGenesisInfoAndGetCommitmentState: true,
			numberOfTxs:                          5,
			prevBlockHash:                        ethService.BlockChain().CurrentSafeBlock().Hash().Bytes(),
			timestamp:                            ethService.BlockChain().CurrentSafeBlock().Time + 2,
			depositTxAmount:                      big.NewInt(0),
			expectedReturnCode:                   0,
		},
		{
			description:                          "ExecuteOptimisticBlock with 5 txs and a deposit tx",
			callGenesisInfoAndGetCommitmentState: true,
			numberOfTxs:                          5,
			prevBlockHash:                        ethService.BlockChain().CurrentSafeBlock().Hash().Bytes(),
			timestamp:                            ethService.BlockChain().CurrentSafeBlock().Time + 2,
			depositTxAmount:                      big.NewInt(1000000000000000000),
			expectedReturnCode:                   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			ethservice, sharedService := shared.SetupSharedService(t, 10)

			// reset the blockchain with each test
			optimisticServiceV1Alpha1 := SetupOptimisticService(t, sharedService)
			executionServiceV1Alpha1 := execution.SetupExecutionService(t, sharedService)

			var err error // adding this to prevent shadowing of genesisInfo in the below if branch
			var genesisInfo *astriaPb.GenesisInfo
			var commitmentStateBeforeExecuteBlock *astriaPb.CommitmentState
			if tt.callGenesisInfoAndGetCommitmentState {
				// call getGenesisInfo and getCommitmentState before calling executeBlock
				genesisInfo, err = executionServiceV1Alpha1.GetGenesisInfo(context.Background(), &astriaPb.GetGenesisInfoRequest{})
				require.Nil(t, err, "GetGenesisInfo failed")
				require.NotNil(t, genesisInfo, "GenesisInfo is nil")

				commitmentStateBeforeExecuteBlock, err = executionServiceV1Alpha1.GetCommitmentState(context.Background(), &astriaPb.GetCommitmentStateRequest{})
				require.Nil(t, err, "GetCommitmentState failed")
				require.NotNil(t, commitmentStateBeforeExecuteBlock, "CommitmentState is nil")
			}

			// create the txs to send
			// create 5 txs
			txs := []*types.Transaction{}
			marshalledTxs := []*sequencerblockv1alpha1.RollupData{}
			for i := 0; i < 5; i++ {
				unsignedTx := types.NewTransaction(uint64(i), shared.TestToAddress, big.NewInt(1), params.TxGas, big.NewInt(params.InitialBaseFee*2), nil)
				tx, err := types.SignTx(unsignedTx, types.LatestSigner(ethservice.BlockChain().Config()), shared.TestKey)
				require.Nil(t, err, "Failed to sign tx")
				txs = append(txs, tx)

				marshalledTx, err := tx.MarshalBinary()
				require.Nil(t, err, "Failed to marshal tx")
				marshalledTxs = append(marshalledTxs, &sequencerblockv1alpha1.RollupData{
					Value: &sequencerblockv1alpha1.RollupData_SequencedData{SequencedData: marshalledTx},
				})
			}

			// create deposit tx if depositTxAmount is non zero
			if tt.depositTxAmount.Cmp(big.NewInt(0)) != 0 {
				depositAmount := shared.BigIntToProtoU128(tt.depositTxAmount)
				bridgeAddress := ethservice.BlockChain().Config().AstriaBridgeAddressConfigs[0].BridgeAddress
				bridgeAssetDenom := ethservice.BlockChain().Config().AstriaBridgeAddressConfigs[0].AssetDenom

				// create new chain destination address for better testing
				chainDestinationAddressPrivKey, err := crypto.GenerateKey()
				require.Nil(t, err, "Failed to generate chain destination address")

				chainDestinationAddress := crypto.PubkeyToAddress(chainDestinationAddressPrivKey.PublicKey)

				depositTx := &sequencerblockv1alpha1.RollupData{Value: &sequencerblockv1alpha1.RollupData_Deposit{Deposit: &sequencerblockv1alpha1.Deposit{
					BridgeAddress: &primitivev1.Address{
						Bech32M: bridgeAddress,
					},
					Asset:                   bridgeAssetDenom,
					Amount:                  depositAmount,
					RollupId:                genesisInfo.RollupId,
					DestinationChainAddress: chainDestinationAddress.String(),
					SourceTransactionId: &primitivev1.TransactionId{
						Inner: "test_tx_hash",
					},
					SourceActionIndex: 0,
				}}}

				marshalledTxs = append(marshalledTxs, depositTx)
			}

			optimisticHeadCh := make(chan core.ChainOptimisticHeadEvent, 1)
			optimsticHeadSub := ethservice.BlockChain().SubscribeChainOptimisticHeadEvent(optimisticHeadCh)
			defer optimsticHeadSub.Unsubscribe()

			baseBlockReq := &optimsticPb.BaseBlock{
				Timestamp: &timestamppb.Timestamp{
					Seconds: int64(tt.timestamp),
				},
				Transactions:       marshalledTxs,
				SequencerBlockHash: []byte("test_hash"),
			}

			res, err := optimisticServiceV1Alpha1.ExecuteOptimisticBlock(context.Background(), baseBlockReq)
			if tt.expectedReturnCode > 0 {
				require.NotNil(t, err, "ExecuteOptimisticBlock should return an error")
				require.Equal(t, tt.expectedReturnCode, status.Code(err), "ExecuteOptimisticBlock failed")
			} else {
				require.Nil(t, err, "ExecuteOptimisticBlock failed")
			}
			if err == nil {
				require.NotNil(t, res, "ExecuteOptimisticBlock response is nil")

				astriaOrdered := ethservice.TxPool().AstriaOrdered()
				require.Equal(t, 0, astriaOrdered.Len(), "AstriaOrdered should be empty")

				// check if commitment state is not updated
				commitmentStateAfterExecuteBlock, err := executionServiceV1Alpha1.GetCommitmentState(context.Background(), &astriaPb.GetCommitmentStateRequest{})
				require.Nil(t, err, "GetCommitmentState failed")

				require.Exactly(t, commitmentStateBeforeExecuteBlock, commitmentStateAfterExecuteBlock, "Commitment state should not be updated")

				// check if the optimistic block is set
				optimisticBlock := ethservice.BlockChain().CurrentOptimisticBlock()
				require.NotNil(t, optimisticBlock, "Optimistic block is not set")

				// check if the optimistic block is correct
				require.Equal(t, common.BytesToHash(res.Hash), optimisticBlock.Hash(), "Optimistic block hashes do not match")
				require.Equal(t, common.BytesToHash(res.ParentBlockHash), optimisticBlock.ParentHash, "Optimistic block parent hashes do not match")
				require.Equal(t, uint64(res.Number), optimisticBlock.Number.Uint64(), "Optimistic block numbers do not match")

				// check if optimistic block is inserted into chain
				block := ethservice.BlockChain().GetBlockByHash(optimisticBlock.Hash())
				require.NotNil(t, block, "Optimistic block not found in blockchain")
				require.Equal(t, uint64(res.Number), block.NumberU64(), "Block number is not correct")

				// timeout for optimistic head event
				select {
				case blockEvent := <-optimisticHeadCh:
					require.NotNil(t, blockEvent, "Optimistic head event not received")
					require.Equal(t, block.Hash(), blockEvent.Block.Hash(), "Optimistic head event block hash is not correct")
					require.Equal(t, block.NumberU64(), blockEvent.Block.NumberU64(), "Optimistic head event block number is not correct")
				case <-time.After(2 * time.Second):
					require.FailNow(t, "Optimistic head event not received")
				case err := <-optimsticHeadSub.Err():
					require.Nil(t, err, "Optimistic head event subscription failed")
				}
			}
		})
	}
}

func TestNewExecutionServiceServerV1Alpha2_StreamBundles(t *testing.T) {
	ethservice, sharedService := shared.SetupSharedService(t, 10)

	optimisticServiceV1Alpha1 := SetupOptimisticService(t, sharedService)
	executionServiceV1Alpha1 := execution.SetupExecutionService(t, sharedService)

	// call genesis info
	genesisInfo, err := executionServiceV1Alpha1.GetGenesisInfo(context.Background(), &astriaPb.GetGenesisInfoRequest{})
	require.Nil(t, err, "GetGenesisInfo failed")
	require.NotNil(t, genesisInfo, "GenesisInfo is nil")

	// call get commitment state
	commitmentState, err := executionServiceV1Alpha1.GetCommitmentState(context.Background(), &astriaPb.GetCommitmentStateRequest{})
	require.Nil(t, err, "GetCommitmentState failed")
	require.NotNil(t, commitmentState, "CommitmentState is nil")

	// get previous block hash
	previousBlock := ethservice.BlockChain().CurrentSafeBlock()
	require.NotNil(t, previousBlock, "Previous block not found")

	// create the optimistic block via the StreamExecuteOptimisticBlock rpc
	requestStreams := []*optimsticPb.StreamExecuteOptimisticBlockRequest{}
	sequencerBlockHash := []byte("sequencer_block_hash")

	// create 1 stream item with 5 txs
	txs := []*types.Transaction{}
	marshalledTxs := []*sequencerblockv1alpha1.RollupData{}
	for i := 0; i < 5; i++ {
		unsignedTx := types.NewTransaction(uint64(i), shared.TestToAddress, big.NewInt(1), params.TxGas, big.NewInt(params.InitialBaseFee*2), nil)
		tx, err := types.SignTx(unsignedTx, types.LatestSigner(ethservice.BlockChain().Config()), shared.TestKey)
		require.Nil(t, err, "Failed to sign tx")
		txs = append(txs, tx)

		marshalledTx, err := tx.MarshalBinary()
		require.Nil(t, err, "Failed to marshal tx")
		marshalledTxs = append(marshalledTxs, &sequencerblockv1alpha1.RollupData{
			Value: &sequencerblockv1alpha1.RollupData_SequencedData{SequencedData: marshalledTx},
		})
	}

	txErrors := ethservice.TxPool().Add(txs, true, false)
	for _, txErr := range txErrors {
		require.Nil(t, txErr, "Failed to add tx to mempool")
	}

	pending, queued := ethservice.TxPool().Stats()
	require.Equal(t, pending, 5, "Mempool should have 5 pending txs")
	require.Equal(t, queued, 0, "Mempool should have 0 queued txs")

	req := optimsticPb.StreamExecuteOptimisticBlockRequest{Block: &optimsticPb.BaseBlock{
		SequencerBlockHash: sequencerBlockHash,
		Transactions:       marshalledTxs,
		Timestamp: &timestamppb.Timestamp{
			Seconds: int64(previousBlock.Time + 2),
		},
	}}

	requestStreams = append(requestStreams, &req)

	mockBidirectionalStream := &MockBidirectionalStreaming[optimsticPb.StreamExecuteOptimisticBlockRequest, optimsticPb.StreamExecuteOptimisticBlockResponse]{
		requestStream:        requestStreams,
		accumulatedResponses: []*optimsticPb.StreamExecuteOptimisticBlockResponse{},
		requestCounter:       0,
	}

	errorCh := make(chan error)
	go func(errorCh chan error) {
		errorCh <- optimisticServiceV1Alpha1.StreamExecuteOptimisticBlock(mockBidirectionalStream)
	}(errorCh)

	select {
	// stream either errors out of gets closed
	case err := <-errorCh:
		require.Nil(t, err, "StreamExecuteOptimisticBlock failed")
	}

	require.Len(t, mockBidirectionalStream.accumulatedResponses, 1, "Number of responses should match the number of requests")
	accumulatedResponse := mockBidirectionalStream.accumulatedResponses[0]

	currentOptimisticBlock := ethservice.BlockChain().CurrentOptimisticBlock()
	require.NotNil(t, currentOptimisticBlock, "Optimistic block is not set")
	require.True(t, bytes.Equal(accumulatedResponse.GetBlock().Hash, currentOptimisticBlock.Hash().Bytes()), "Optimistic block hashes do not match")
	require.True(t, bytes.Equal(accumulatedResponse.GetBlock().ParentBlockHash, currentOptimisticBlock.ParentHash.Bytes()), "Optimistic block parent hashes do not match")
	require.Equal(t, uint64(accumulatedResponse.GetBlock().Number), currentOptimisticBlock.Number.Uint64(), "Optimistic block numbers do not match")

	// assert mempool is cleared
	astriaOrdered := ethservice.TxPool().AstriaOrdered()
	require.Equal(t, 0, astriaOrdered.Len(), "AstriaOrdered should be empty")

	pending, queued = ethservice.TxPool().Stats()
	require.Equal(t, pending, 0, "Mempool should have 0 pending txs")
	require.Equal(t, queued, 0, "Mempool should have 0 queued txs")

	mockServerSideStreaming := MockServerSideStreaming[optimsticPb.Bundle]{
		sentResponses: []*optimsticPb.Bundle{},
	}

	errorCh = make(chan error)
	go func() {
		errorCh <- optimisticServiceV1Alpha1.StreamBundles(&optimsticPb.StreamBundlesRequest{}, &mockServerSideStreaming)
	}()

	stateDb, err := ethservice.BlockChain().StateAt(currentOptimisticBlock.Root)
	require.Nil(t, err, "Failed to get state db")

	latestNonce := stateDb.GetNonce(shared.TestAddr)

	// optimistic block is created, we can now add txs and check if they get streamed
	// create 5 txs
	txs = []*types.Transaction{}
	for i := 0; i < 5; i++ {
		unsignedTx := types.NewTransaction(latestNonce+uint64(i), shared.TestToAddress, big.NewInt(1), params.TxGas, big.NewInt(params.InitialBaseFee*2), nil)
		tx, err := types.SignTx(unsignedTx, types.LatestSigner(ethservice.BlockChain().Config()), shared.TestKey)
		require.Nil(t, err, "Failed to sign tx")
		txs = append(txs, tx)

		marshalledTx, err := tx.MarshalBinary()
		require.Nil(t, err, "Failed to marshal tx")
		marshalledTxs = append(marshalledTxs, &sequencerblockv1alpha1.RollupData{
			Value: &sequencerblockv1alpha1.RollupData_SequencedData{SequencedData: marshalledTx},
		})
	}

	txErrors = ethservice.TxPool().Add(txs, true, false)
	for _, txErr := range txErrors {
		require.Nil(t, txErr, "Failed to add tx to mempool")
	}

	pending, queued = ethservice.TxPool().Stats()
	require.Equal(t, pending, 5, "Mempool should have 5 pending txs")
	require.Equal(t, queued, 0, "Mempool should have 0 queued txs")

	// give some time for the txs to stream
	time.Sleep(5 * time.Second)

	// close the mempool to error the method out
	err = ethservice.TxPool().Close()
	require.Nil(t, err, "Failed to close mempool")

	select {
	case err := <-errorCh:
		require.ErrorContains(t, err, "error waiting for pending transactions")
	}

	require.Len(t, mockServerSideStreaming.sentResponses, 5, "Number of responses should match the number of requests")

	txIndx := 0
	for _, resp := range mockServerSideStreaming.sentResponses {
		require.Len(t, resp.Transactions, 1, "Bundle should have 1 tx")

		receivedTx := resp.Transactions[0]
		sentTx := txs[txIndx]
		marshalledSentTx, err := sentTx.MarshalBinary()
		require.Nil(t, err, "Failed to marshal tx")
		require.True(t, bytes.Equal(receivedTx, marshalledSentTx), "Received tx does not match sent tx")
		txIndx += 1

		require.True(t, bytes.Equal(resp.PrevRollupBlockHash, currentOptimisticBlock.Hash().Bytes()), "PrevRollupBlockHash should match the current optimistic block hash")
		require.True(t, bytes.Equal(resp.BaseSequencerBlockHash, *optimisticServiceV1Alpha1.currentOptimisticSequencerBlock.Load()), "BaseSequencerBlockHash should match the current optimistic sequencer block hash")
	}
}

func TestExecutionServiceServerV1Alpha2_StreamExecuteOptimisticBlock(t *testing.T) {
	ethservice, sharedService := shared.SetupSharedService(t, 10)

	optimisticServiceV1Alpha1 := SetupOptimisticService(t, sharedService)
	executionServiceV1Alpha1 := execution.SetupExecutionService(t, sharedService)

	// call genesis info
	genesisInfo, err := executionServiceV1Alpha1.GetGenesisInfo(context.Background(), &astriaPb.GetGenesisInfoRequest{})
	require.Nil(t, err, "GetGenesisInfo failed")
	require.NotNil(t, genesisInfo, "GenesisInfo is nil")

	// call get commitment state
	commitmentState, err := executionServiceV1Alpha1.GetCommitmentState(context.Background(), &astriaPb.GetCommitmentStateRequest{})
	require.Nil(t, err, "GetCommitmentState failed")
	require.NotNil(t, commitmentState, "CommitmentState is nil")

	// get previous block hash
	previousBlock := ethservice.BlockChain().CurrentSafeBlock()
	require.NotNil(t, previousBlock, "Previous block not found")

	requestStreams := []*optimsticPb.StreamExecuteOptimisticBlockRequest{}
	sequencerBlockHash := []byte("sequencer_block_hash")

	// create 1 stream item with 5 txs
	txs := []*types.Transaction{}
	marshalledTxs := []*sequencerblockv1alpha1.RollupData{}
	for i := 0; i < 5; i++ {
		unsignedTx := types.NewTransaction(uint64(i), shared.TestToAddress, big.NewInt(1), params.TxGas, big.NewInt(params.InitialBaseFee*2), nil)
		tx, err := types.SignTx(unsignedTx, types.LatestSigner(ethservice.BlockChain().Config()), shared.TestKey)
		require.Nil(t, err, "Failed to sign tx")
		txs = append(txs, tx)

		marshalledTx, err := tx.MarshalBinary()
		require.Nil(t, err, "Failed to marshal tx")
		marshalledTxs = append(marshalledTxs, &sequencerblockv1alpha1.RollupData{
			Value: &sequencerblockv1alpha1.RollupData_SequencedData{SequencedData: marshalledTx},
		})
	}

	errs := ethservice.TxPool().Add(txs, true, false)
	for _, err := range errs {
		require.Nil(t, err, "Failed to add tx to mempool")
	}

	pending, queued := ethservice.TxPool().Stats()
	require.Equal(t, pending, 5, "Mempool should have 5 pending txs")
	require.Equal(t, queued, 0, "Mempool should have 0 queued txs")

	req := optimsticPb.StreamExecuteOptimisticBlockRequest{Block: &optimsticPb.BaseBlock{
		SequencerBlockHash: sequencerBlockHash,
		Transactions:       marshalledTxs,
		Timestamp: &timestamppb.Timestamp{
			Seconds: int64(previousBlock.Time + 2),
		},
	}}

	requestStreams = append(requestStreams, &req)

	mockStream := &MockBidirectionalStreaming[optimsticPb.StreamExecuteOptimisticBlockRequest, optimsticPb.StreamExecuteOptimisticBlockResponse]{
		requestStream:        requestStreams,
		accumulatedResponses: []*optimsticPb.StreamExecuteOptimisticBlockResponse{},
		requestCounter:       0,
	}

	errorCh := make(chan error)
	go func(errorCh chan error) {
		errorCh <- optimisticServiceV1Alpha1.StreamExecuteOptimisticBlock(mockStream)
	}(errorCh)

	select {
	// the stream will either errors out or gets closed
	case err := <-errorCh:
		require.Nil(t, err, "StreamExecuteOptimisticBlock failed")
	}

	accumulatedResponses := mockStream.accumulatedResponses

	require.Equal(t, len(accumulatedResponses), len(mockStream.requestStream), "Number of responses should match the number of requests")

	blockCounter := 1
	for _, response := range accumulatedResponses {
		require.True(t, bytes.Equal(response.GetBaseSequencerBlockHash(), sequencerBlockHash), "Sequencer block hash does not match")
		block := response.GetBlock()
		require.True(t, bytes.Equal(block.ParentBlockHash, previousBlock.Hash().Bytes()), "Parent block hash does not match")
		requiredBlockNumber := big.NewInt(0).Add(previousBlock.Number, big.NewInt(int64(blockCounter)))
		require.Equal(t, requiredBlockNumber.Uint64(), uint64(block.Number), "Block number is not correct")
		blockCounter += 1
	}

	// ensure mempool is cleared
	astriaOrdered := ethservice.TxPool().AstriaOrdered()
	require.Equal(t, 0, astriaOrdered.Len(), "AstriaOrdered should be empty")

	pending, queued = ethservice.TxPool().Stats()
	require.Equal(t, pending, 0, "Mempool should have 0 pending txs")
	require.Equal(t, queued, 0, "Mempool should have 0 queued txs")
}
