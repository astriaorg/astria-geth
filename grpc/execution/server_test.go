package execution

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"math/big"
	"testing"

	astriaPb "buf.build/gen/go/astria/execution-apis/protocolbuffers/go/astria/execution/v2"
	primitivev1 "buf.build/gen/go/astria/primitives/protocolbuffers/go/astria/primitive/v1"
	sequencerblockv1 "buf.build/gen/go/astria/sequencerblock-apis/protocolbuffers/go/astria/sequencerblock/v1"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestExecutionServiceServerV2_CreateExecutionSession(t *testing.T) {
	ethservice, serviceV2 := setupExecutionService(t, 10, false)

	session, err := serviceV2.CreateExecutionSession(context.Background(), &astriaPb.CreateExecutionSessionRequest{})
	require.Nil(t, err, "CreateExecutionSession failed")
	require.NotEmpty(t, serviceV2.activeSessionId, "Active session ID should not be empty")
	require.NotNil(t, serviceV2.activeFork, "Active fork should not be nil")

	bc := ethservice.BlockChain()
	hashedRollupId := sha256.Sum256([]byte(bc.Config().AstriaRollupName))

	require.NotEmpty(t, session.SessionId, "SessionId should not be empty")
	require.Equal(t, session.SessionId, serviceV2.activeSessionId, "Active session ID is not set correctly")

	activeFork := bc.Config().AstriaForks.GetForkAtHeight(1)

	require.True(t, bytes.Equal(session.ExecutionSessionParameters.RollupId.Inner, hashedRollupId[:]), "RollupId is not correct")
	require.Equal(t, session.ExecutionSessionParameters.RollupStartBlockNumber, activeFork.Height, "RollupStartBlockNumber is not correct")
	require.Equal(t, session.ExecutionSessionParameters.RollupEndBlockNumber, activeFork.StopHeight, "RollupEndBlockNumber is not correct")
	require.Equal(t, session.ExecutionSessionParameters.SequencerStartBlockHeight, activeFork.Sequencer.StartHeight, "SequencerStartBlockHeight is not correct")
	require.Equal(t, session.ExecutionSessionParameters.SequencerChainId, activeFork.Sequencer.ChainID, "SequencerChainId is not correct")
	require.Equal(t, session.ExecutionSessionParameters.CelestiaChainId, activeFork.Celestia.ChainID, "CelestiaChainId is not correct")
	require.Equal(t, session.ExecutionSessionParameters.CelestiaSearchHeightMaxLookAhead, activeFork.Celestia.SearchHeightMaxLookAhead, "CelestiaSearchHeightMaxLookAhead is not correct")

	require.NotNil(t, session.CommitmentState, "CommitmentState is nil")

	softBlock := ethservice.BlockChain().CurrentSafeBlock()
	require.NotNil(t, softBlock, "SoftBlock is nil")

	firmBlock := ethservice.BlockChain().CurrentFinalBlock()
	require.NotNil(t, firmBlock, "FirmBlock is nil")

	require.Equal(t, softBlock.Hash().Hex(), session.CommitmentState.SoftExecutedBlockMetadata.Hash, "Soft Block Hashes do not match")
	require.Equal(t, softBlock.ParentHash.Hex(), session.CommitmentState.SoftExecutedBlockMetadata.ParentHash, "Soft Block Parent Hash do not match")
	require.Equal(t, softBlock.Number.Uint64(), session.CommitmentState.SoftExecutedBlockMetadata.Number, "Soft Block Number do not match")

	require.Equal(t, firmBlock.Hash().Hex(), session.CommitmentState.FirmExecutedBlockMetadata.Hash, "Firm Block Hashes do not match")
	require.Equal(t, firmBlock.ParentHash.Hex(), session.CommitmentState.FirmExecutedBlockMetadata.ParentHash, "Firm Block Parent Hash do not match")
	require.Equal(t, firmBlock.Number.Uint64(), session.CommitmentState.FirmExecutedBlockMetadata.Number, "Firm Block Number do not match")
	require.Equal(t, session.CommitmentState.LowestCelestiaSearchHeight, ethservice.BlockChain().Config().AstriaForks.GetForkAtHeight(1).Celestia.StartHeight, "LowestCelestiaSearchHeight is not correct")
}

func TestExecutionServiceServerV2_GetExecutedBlockMetadata(t *testing.T) {
	ethservice, serviceV2 := setupExecutionService(t, 10, false)

	tests := []struct {
		description        string
		request            *astriaPb.GetExecutedBlockMetadataRequest
		expectedReturnCode codes.Code
	}{
		{
			description: "Get block by block number 1",
			request: &astriaPb.GetExecutedBlockMetadataRequest{
				Identifier: &astriaPb.ExecutedBlockIdentifier{Identifier: &astriaPb.ExecutedBlockIdentifier_Number{Number: 1}},
			},
			expectedReturnCode: 0,
		},
		{
			description: "Get block by block hash",
			request: &astriaPb.GetExecutedBlockMetadataRequest{
				Identifier: &astriaPb.ExecutedBlockIdentifier{Identifier: &astriaPb.ExecutedBlockIdentifier_Hash{Hash: ethservice.BlockChain().GetBlockByNumber(4).Hash().Hex()}},
			},
			expectedReturnCode: 0,
		},
		{
			description: "Get block which is not present",
			request: &astriaPb.GetExecutedBlockMetadataRequest{
				Identifier: &astriaPb.ExecutedBlockIdentifier{Identifier: &astriaPb.ExecutedBlockIdentifier_Number{Number: 100}},
			},
			expectedReturnCode: codes.NotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			blockMetadata, err := serviceV2.GetExecutedBlockMetadata(context.Background(), tt.request)
			if tt.expectedReturnCode > 0 {
				require.NotNil(t, err, "GetExecutedBlockMetadata should return an error")
				require.Equal(t, tt.expectedReturnCode, status.Code(err), "GetExecutedBlockMetadata failed")
			}
			if err == nil {
				require.NotNil(t, blockMetadata, "Block metadata not found")
				var block *types.Block
				if tt.request.Identifier.GetNumber() != 0 {
					// get block by number
					block = ethservice.BlockChain().GetBlockByNumber(uint64(tt.request.Identifier.GetNumber()))
				}
				if tt.request.Identifier.GetHash() != "" {
					block = ethservice.BlockChain().GetBlockByHash(common.HexToHash(tt.request.Identifier.GetHash()))
				}
				require.NotNil(t, block, "Block not found")

				require.Equal(t, uint64(blockMetadata.Number), block.NumberU64(), "Block number is not correct")
				require.Equal(t, block.ParentHash().Hex(), blockMetadata.ParentHash, "Parent Block Hash is not correct")
				require.Equal(t, block.Hash().Hex(), blockMetadata.Hash, "BlockHash is not correct")
			}
		})
	}
}

func TestExecutionServiceServerV2_ExecuteBlock(t *testing.T) {
	ethservice, _ := setupExecutionService(t, 10, false)

	tests := []struct {
		description        string
		createSessionFirst bool
		numberOfTxs        int
		prevBlockHash      string
		timestamp          uint64
		depositTxAmount    *big.Int // if this is non zero then we send a deposit tx
		expectedReturnCode codes.Code
	}{
		{
			description:        "ExecuteBlock without creating session first",
			createSessionFirst: false,
			numberOfTxs:        5,
			prevBlockHash:      ethservice.BlockChain().GetBlockByNumber(2).Hash().Hex(),
			timestamp:          ethservice.BlockChain().GetBlockByNumber(2).Time() + 2,
			depositTxAmount:    big.NewInt(0),
			expectedReturnCode: codes.PermissionDenied,
		},
		{
			description:        "ExecuteBlock with 5 txs and no deposit tx",
			createSessionFirst: true,
			numberOfTxs:        5,
			prevBlockHash:      ethservice.BlockChain().CurrentSafeBlock().Hash().Hex(),
			timestamp:          ethservice.BlockChain().CurrentSafeBlock().Time + 2,
			depositTxAmount:    big.NewInt(0),
			expectedReturnCode: 0,
		},
		{
			description:        "ExecuteBlock with 5 txs and a deposit tx",
			createSessionFirst: true,
			numberOfTxs:        5,
			prevBlockHash:      ethservice.BlockChain().CurrentSafeBlock().Hash().Hex(),
			timestamp:          ethservice.BlockChain().CurrentSafeBlock().Time + 2,
			depositTxAmount:    big.NewInt(1000000000000000000),
			expectedReturnCode: 0,
		},
		{
			description:        "ExecuteBlock with incorrect previous block hash",
			createSessionFirst: true,
			numberOfTxs:        5,
			prevBlockHash:      common.BytesToHash([]byte("incorrect_hash")).Hex(),
			timestamp:          ethservice.BlockChain().CurrentSafeBlock().Time + 2,
			depositTxAmount:    big.NewInt(0),
			expectedReturnCode: codes.FailedPrecondition,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			// reset the blockchain with each test
			ethservice, serviceV2 := setupExecutionService(t, 10, false)

			fork := ethservice.BlockChain().Config().AstriaForks.GetForkAtHeight(1)
			var bridgeConfig params.AstriaBridgeAddressConfig
			for _, cfg := range fork.BridgeAddresses {
				bridgeConfig = *cfg
				break
			}

			expectedTransactionCount := tt.numberOfTxs
			if tt.depositTxAmount.Cmp(big.NewInt(0)) != 0 {
				expectedTransactionCount = expectedTransactionCount + 1
			}

			var err error
			var session *astriaPb.ExecutionSession
			if tt.createSessionFirst {
				// Create execution session before calling executeBlock
				session, err = serviceV2.CreateExecutionSession(context.Background(), &astriaPb.CreateExecutionSessionRequest{})
				require.Nil(t, err, "CreateExecutionSession failed")
				require.NotNil(t, session, "ExecutionSession is nil")

				// Debug information for the "incorrect previous block hash" test
				if tt.description == "ExecuteBlock with incorrect previous block hash" {
					currentBlock := ethservice.BlockChain().CurrentBlock()
					currentHeight := currentBlock.Number.Uint64() + 1
					currentFork := ethservice.BlockChain().Config().AstriaForks.GetForkAtHeight(currentHeight)

					t.Logf("Debug - Current height: %d, Fork height: %d, Fork stop height: %d",
						currentHeight, currentFork.Height, currentFork.StopHeight)
					t.Logf("Debug - Current safe block hash: %s, Test previous hash: %s",
						ethservice.BlockChain().CurrentSafeBlock().Hash().Hex(), tt.prevBlockHash)
				}
			}

			// create 5 txs
			marshalledTxs := []*sequencerblockv1.RollupData{}
			for i := 0; i < 5; i++ {
				unsignedTx := types.NewTransaction(uint64(i), testToAddress, big.NewInt(1), params.TxGas, big.NewInt(params.InitialBaseFee*2), nil)
				tx, err := types.SignTx(unsignedTx, types.LatestSigner(ethservice.BlockChain().Config()), testKey)
				require.Nil(t, err, "Failed to sign tx")

				marshalledTx, err := tx.MarshalBinary()
				require.Nil(t, err, "Failed to marshal tx")
				marshalledTxs = append(marshalledTxs, &sequencerblockv1.RollupData{
					Value: &sequencerblockv1.RollupData_SequencedData{SequencedData: marshalledTx},
				})
			}

			// create deposit tx if depositTxAmount is non zero
			if tt.depositTxAmount.Cmp(big.NewInt(0)) != 0 && tt.createSessionFirst {
				depositAmount := bigIntToProtoU128(tt.depositTxAmount)
				bridgeAddress := bridgeConfig.BridgeAddress
				bridgeAssetDenom := bridgeConfig.AssetDenom

				// create new chain destination address for better testing
				chainDestinationAddressPrivKey, err := crypto.GenerateKey()
				require.Nil(t, err, "Failed to generate chain destination address")

				chainDestinationAddress := crypto.PubkeyToAddress(chainDestinationAddressPrivKey.PublicKey)

				hashedRollupId := sha256.Sum256([]byte(ethservice.BlockChain().Config().AstriaRollupName))
				rollupId := primitivev1.RollupId{
					Inner: hashedRollupId[:],
				}

				depositTx := &sequencerblockv1.RollupData{Value: &sequencerblockv1.RollupData_Deposit{Deposit: &sequencerblockv1.Deposit{
					BridgeAddress: &primitivev1.Address{
						Bech32M: bridgeAddress,
					},
					Asset:                   bridgeAssetDenom,
					Amount:                  depositAmount,
					RollupId:                &rollupId,
					DestinationChainAddress: chainDestinationAddress.String(),
					SourceTransactionId: &primitivev1.TransactionId{
						Inner: "test_tx_hash",
					},
					SourceActionIndex: 0,
				}}}

				marshalledTxs = append(marshalledTxs, depositTx)
			}

			sessionId := ""
			if session != nil {
				sessionId = session.SessionId
			}

			executeBlockReq := &astriaPb.ExecuteBlockRequest{
				SessionId:  sessionId,
				ParentHash: tt.prevBlockHash,
				Timestamp: &timestamppb.Timestamp{
					Seconds: int64(tt.timestamp),
				},
				Transactions: marshalledTxs,
			}

			executeBlockRes, err := serviceV2.ExecuteBlock(context.Background(), executeBlockReq)
			if tt.expectedReturnCode > 0 {
				require.NotNil(t, err, "ExecuteBlock should return an error")
				require.Equal(t, tt.expectedReturnCode, status.Code(err), "ExecuteBlock failed")
			}
			if err == nil {
				require.NotNil(t, executeBlockRes, "ExecuteBlock response is nil")

				astriaOrdered := ethservice.TxPool().AstriaOrdered()
				require.Equal(t, 0, astriaOrdered.Len(), "AstriaOrdered should be empty")

				blockhash := common.HexToHash(executeBlockRes.ExecutedBlockMetadata.Hash)
				block := ethservice.BlockChain().GetBlockByHash(blockhash)
				require.Equal(t, expectedTransactionCount, len(block.Transactions()), "Transaction count is not correct")
			}
		})
	}
}

func TestExecutionServiceServerV2_ExecuteBlockWithGasOverrun(t *testing.T) {
	ethservice, _ := setupExecutionService(t, 0, true)

	tests := []struct {
		description              string
		createSessionFirst       bool
		numberOfTxs              int
		prevBlockHash            string
		timestamp                uint64
		depositTxAmount          *big.Int // if this is non zero then we send a deposit tx
		expectedReturnCode       codes.Code
		expectedTransactionCount int
	}{
		{
			description:              "ExecuteBlock with 5 txs and no deposit tx",
			createSessionFirst:       true,
			numberOfTxs:              5,
			prevBlockHash:            ethservice.BlockChain().CurrentSafeBlock().Hash().Hex(),
			timestamp:                ethservice.BlockChain().CurrentSafeBlock().Time + 2,
			depositTxAmount:          big.NewInt(0),
			expectedReturnCode:       0,
			expectedTransactionCount: 0,
		},
		{
			description:              "ExecuteBlock with 5 txs and a deposit tx",
			createSessionFirst:       true,
			numberOfTxs:              5,
			prevBlockHash:            ethservice.BlockChain().CurrentSafeBlock().Hash().Hex(),
			timestamp:                ethservice.BlockChain().CurrentSafeBlock().Time + 2,
			depositTxAmount:          big.NewInt(1000000000000000000),
			expectedReturnCode:       0,
			expectedTransactionCount: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			// reset the blockchain with each test
			ethservice, serviceV2 := setupExecutionService(t, 0, true)

			fork := ethservice.BlockChain().Config().AstriaForks.GetForkAtHeight(1)
			var bridgeConfig params.AstriaBridgeAddressConfig
			for _, cfg := range fork.BridgeAddresses {
				bridgeConfig = *cfg
				break
			}

			var err error
			var session *astriaPb.ExecutionSession
			if tt.createSessionFirst {
				// Create execution session before calling executeBlock
				session, err = serviceV2.CreateExecutionSession(context.Background(), &astriaPb.CreateExecutionSessionRequest{})
				require.Nil(t, err, "CreateExecutionSession failed")
				require.NotNil(t, session, "ExecutionSession is nil")

				// Debug information for the "incorrect previous block hash" test
				if tt.description == "ExecuteBlock with incorrect previous block hash" {
					currentBlock := ethservice.BlockChain().CurrentBlock()
					currentHeight := currentBlock.Number.Uint64() + 1
					currentFork := ethservice.BlockChain().Config().AstriaForks.GetForkAtHeight(currentHeight)

					t.Logf("Debug - Current height: %d, Fork height: %d, Fork stop height: %d",
						currentHeight, currentFork.Height, currentFork.StopHeight)
					t.Logf("Debug - Current safe block hash: %s, Test previous hash: %s",
						ethservice.BlockChain().CurrentSafeBlock().Hash().Hex(), tt.prevBlockHash)
				}
			}

			// create 5 txs
			marshalledTxs := []*sequencerblockv1.RollupData{}
			for i := 0; i < 5; i++ {
				unsignedTx := types.NewTransaction(uint64(i), testToAddress, big.NewInt(1), params.TxGas, big.NewInt(params.InitialBaseFee*2), nil)
				tx, err := types.SignTx(unsignedTx, types.LatestSigner(ethservice.BlockChain().Config()), testKey)
				require.Nil(t, err, "Failed to sign tx")

				marshalledTx, err := tx.MarshalBinary()
				require.Nil(t, err, "Failed to marshal tx")
				marshalledTxs = append(marshalledTxs, &sequencerblockv1.RollupData{
					Value: &sequencerblockv1.RollupData_SequencedData{SequencedData: marshalledTx},
				})
			}

			// create deposit tx if depositTxAmount is non zero
			if tt.depositTxAmount.Cmp(big.NewInt(0)) != 0 && tt.createSessionFirst {
				depositAmount := bigIntToProtoU128(tt.depositTxAmount)
				bridgeAddress := bridgeConfig.BridgeAddress
				bridgeAssetDenom := bridgeConfig.AssetDenom

				// create new chain destination address for better testing
				chainDestinationAddressPrivKey, err := crypto.GenerateKey()
				require.Nil(t, err, "Failed to generate chain destination address")

				chainDestinationAddress := crypto.PubkeyToAddress(chainDestinationAddressPrivKey.PublicKey)

				hashedRollupId := sha256.Sum256([]byte(ethservice.BlockChain().Config().AstriaRollupName))
				rollupId := primitivev1.RollupId{
					Inner: hashedRollupId[:],
				}

				depositTx := &sequencerblockv1.RollupData{Value: &sequencerblockv1.RollupData_Deposit{Deposit: &sequencerblockv1.Deposit{
					BridgeAddress: &primitivev1.Address{
						Bech32M: bridgeAddress,
					},
					Asset:                   bridgeAssetDenom,
					Amount:                  depositAmount,
					RollupId:                &rollupId,
					DestinationChainAddress: chainDestinationAddress.String(),
					SourceTransactionId: &primitivev1.TransactionId{
						Inner: "test_tx_hash",
					},
					SourceActionIndex: 0,
				}}}

				marshalledTxs = append(marshalledTxs, depositTx)
			}

			sessionId := ""
			if session != nil {
				sessionId = session.SessionId
			}

			executeBlockReq := &astriaPb.ExecuteBlockRequest{
				SessionId:  sessionId,
				ParentHash: tt.prevBlockHash,
				Timestamp: &timestamppb.Timestamp{
					Seconds: int64(tt.timestamp),
				},
				Transactions: marshalledTxs,
			}

			executeBlockRes, err := serviceV2.ExecuteBlock(context.Background(), executeBlockReq)
			if tt.expectedReturnCode > 0 {
				require.NotNil(t, err, "ExecuteBlock should return an error")
				require.Equal(t, tt.expectedReturnCode, status.Code(err), "ExecuteBlock failed")
			}
			if err == nil {
				require.NotNil(t, executeBlockRes, "ExecuteBlock response is nil")

				astriaOrdered := ethservice.TxPool().AstriaOrdered()
				require.Equal(t, 0, astriaOrdered.Len(), "AstriaOrdered should be empty")

				blockhash := common.HexToHash(executeBlockRes.ExecutedBlockMetadata.Hash)
				block := ethservice.BlockChain().GetBlockByHash(blockhash)
				require.Equal(t, tt.expectedTransactionCount, len(block.Transactions()), "Transaction count is not correct")
			}
		})
	}
}

func TestExecutionServiceServerV2_UpdateCommitmentState(t *testing.T) {
	ethservice, serviceV2 := setupExecutionService(t, 10, false)

	// Create execution session
	session, err := serviceV2.CreateExecutionSession(context.Background(), &astriaPb.CreateExecutionSessionRequest{})
	require.Nil(t, err, "CreateExecutionSession failed")
	require.NotNil(t, session, "ExecutionSession is nil")

	// Execute a block
	softBlock := ethservice.BlockChain().CurrentSafeBlock()
	executeBlockReq := &astriaPb.ExecuteBlockRequest{
		SessionId:  session.SessionId,
		ParentHash: softBlock.Hash().Hex(),
		Timestamp: &timestamppb.Timestamp{
			Seconds: int64(softBlock.Time + 2),
		},
		Transactions: []*sequencerblockv1.RollupData{},
	}

	executeBlockRes, err := serviceV2.ExecuteBlock(context.Background(), executeBlockReq)
	require.Nil(t, err, "ExecuteBlock failed")
	require.NotNil(t, executeBlockRes, "ExecuteBlock response is nil")

	// Test invalid session ID
	invalidSessionReq := &astriaPb.UpdateCommitmentStateRequest{
		SessionId: "invalid-session-id",
		CommitmentState: &astriaPb.CommitmentState{
			SoftExecutedBlockMetadata:  executeBlockRes.ExecutedBlockMetadata,
			FirmExecutedBlockMetadata:  executeBlockRes.ExecutedBlockMetadata,
			LowestCelestiaSearchHeight: session.CommitmentState.LowestCelestiaSearchHeight + 1,
		},
	}

	_, err = serviceV2.UpdateCommitmentState(context.Background(), invalidSessionReq)
	require.NotNil(t, err, "UpdateCommitmentState should fail with invalid session ID")
	require.Equal(t, codes.PermissionDenied, status.Code(err), "Should get PermissionDenied error")

	// Test decreasing Celestia height
	decreasedHeightReq := &astriaPb.UpdateCommitmentStateRequest{
		SessionId: session.SessionId,
		CommitmentState: &astriaPb.CommitmentState{
			SoftExecutedBlockMetadata:  executeBlockRes.ExecutedBlockMetadata,
			FirmExecutedBlockMetadata:  executeBlockRes.ExecutedBlockMetadata,
			LowestCelestiaSearchHeight: session.CommitmentState.LowestCelestiaSearchHeight - 1, // Decrease height
		},
	}

	_, err = serviceV2.UpdateCommitmentState(context.Background(), decreasedHeightReq)
	require.NotNil(t, err, "UpdateCommitmentState should fail with decreased Celestia height")
	require.Equal(t, codes.InvalidArgument, status.Code(err), "Should get InvalidArgument error")
	require.Contains(t, err.Error(), "Base Celestia height cannot be decreased")

	// Test block hash that doesn't exist
	nonExistentBlockReq := &astriaPb.UpdateCommitmentStateRequest{
		SessionId: session.SessionId,
		CommitmentState: &astriaPb.CommitmentState{
			SoftExecutedBlockMetadata: &astriaPb.ExecutedBlockMetadata{
				Number:     100,
				Hash:       common.BytesToHash([]byte("non-existent-hash")).Hex(),
				ParentHash: common.BytesToHash([]byte("non-existent-parent")).Hex(),
				Timestamp:  &timestamppb.Timestamp{Seconds: int64(softBlock.Time + 10)},
			},
			FirmExecutedBlockMetadata:  executeBlockRes.ExecutedBlockMetadata,
			LowestCelestiaSearchHeight: session.CommitmentState.LowestCelestiaSearchHeight + 1,
		},
	}

	_, err = serviceV2.UpdateCommitmentState(context.Background(), nonExistentBlockReq)
	require.NotNil(t, err, "UpdateCommitmentState should fail with non-existent block hash")
	require.Equal(t, codes.InvalidArgument, status.Code(err), "Should get InvalidArgument error")
	require.Contains(t, err.Error(), "Soft block specified does not exist")

	// Test basic successful case
	updateCommitmentStateReq := &astriaPb.UpdateCommitmentStateRequest{
		SessionId: session.SessionId,
		CommitmentState: &astriaPb.CommitmentState{
			SoftExecutedBlockMetadata:  executeBlockRes.ExecutedBlockMetadata,
			FirmExecutedBlockMetadata:  executeBlockRes.ExecutedBlockMetadata,
			LowestCelestiaSearchHeight: session.CommitmentState.LowestCelestiaSearchHeight + 1,
		},
	}

	commitmentStateRes, err := serviceV2.UpdateCommitmentState(context.Background(), updateCommitmentStateReq)
	require.Nil(t, err, "UpdateCommitmentState failed")
	require.NotNil(t, commitmentStateRes, "UpdateCommitmentState response is nil")

	// Verify commitment state
	softBlockAfter := ethservice.BlockChain().CurrentSafeBlock()
	firmBlockAfter := ethservice.BlockChain().CurrentFinalBlock()

	require.Equal(t, executeBlockRes.ExecutedBlockMetadata.Hash, softBlockAfter.Hash().Hex(), "Soft block hash incorrect")
	require.Equal(t, executeBlockRes.ExecutedBlockMetadata.Hash, firmBlockAfter.Hash().Hex(), "Firm block hash incorrect")

	// Check celestia height
	celestiaHeight := ethservice.BlockChain().CurrentBaseCelestiaHeight()
	require.Equal(t, updateCommitmentStateReq.CommitmentState.LowestCelestiaSearchHeight, celestiaHeight, "Celestia height not updated correctly")
}

func TestEthHeaderToExecutedBlockMetadata(t *testing.T) {
	ethservice, _ := setupExecutionService(t, 10, false)

	// Test with valid header
	header := ethservice.BlockChain().CurrentBlock()
	metadata, err := ethHeaderToExecutedBlockMetadata(header)
	require.Nil(t, err, "ethHeaderToExecutedBlockMetadata should not fail with valid header")
	require.Equal(t, header.Number.Uint64(), metadata.Number, "Block number doesn't match")
	require.Equal(t, header.Hash().Hex(), metadata.Hash, "Block hash doesn't match")
	require.Equal(t, header.ParentHash.Hex(), metadata.ParentHash, "Parent hash doesn't match")
	require.Equal(t, int64(header.Time), metadata.Timestamp.Seconds, "Timestamp doesn't match")

	// Test with nil header
	metadata, err = ethHeaderToExecutedBlockMetadata(nil)
	require.NotNil(t, err, "ethHeaderToExecutedBlockMetadata should fail with nil header")
	require.Nil(t, metadata, "Metadata should be nil when header is nil")
}

func TestGetExecutedBlockMetadataFromIdentifier(t *testing.T) {
	ethservice, serviceV2 := setupExecutionService(t, 10, false)
	bc := ethservice.BlockChain()

	// Get a block by number
	block2 := bc.GetBlockByNumber(2)
	require.NotNil(t, block2, "Test setup should have a block at number 2")

	identifier1 := &astriaPb.ExecutedBlockIdentifier{
		Identifier: &astriaPb.ExecutedBlockIdentifier_Number{
			Number: 2,
		},
	}

	metadata1, err := serviceV2.getExecutedBlockMetadataFromIdentifier(identifier1)
	require.Nil(t, err, "Should successfully get metadata by block number")
	require.Equal(t, uint64(2), metadata1.Number, "Block number should match")
	require.Equal(t, block2.Hash().Hex(), metadata1.Hash, "Block hash should match")

	// Get a block by hash
	block3 := bc.GetBlockByNumber(3)
	require.NotNil(t, block3, "Test setup should have a block at number 3")

	identifier2 := &astriaPb.ExecutedBlockIdentifier{
		Identifier: &astriaPb.ExecutedBlockIdentifier_Hash{
			Hash: block3.Hash().Hex(),
		},
	}

	metadata2, err := serviceV2.getExecutedBlockMetadataFromIdentifier(identifier2)
	require.Nil(t, err, "Should successfully get metadata by block hash")
	require.Equal(t, uint64(3), metadata2.Number, "Block number should match")
	require.Equal(t, block3.Hash().Hex(), metadata2.Hash, "Block hash should match")

	// non-existent block by number
	identifier3 := &astriaPb.ExecutedBlockIdentifier{
		Identifier: &astriaPb.ExecutedBlockIdentifier_Number{
			Number: 100,
		},
	}

	metadata3, err := serviceV2.getExecutedBlockMetadataFromIdentifier(identifier3)
	require.NotNil(t, err, "Should fail for non-existent block")
	require.Nil(t, metadata3, "Metadata should be nil for non-existent block")
	require.Equal(t, codes.NotFound, status.Code(err), "Should return NotFound error")

	// non-existent block by hash
	identifier4 := &astriaPb.ExecutedBlockIdentifier{
		Identifier: &astriaPb.ExecutedBlockIdentifier_Hash{
			Hash: common.BytesToHash([]byte("non-existent-hash")).Hex(),
		},
	}

	metadata4, err := serviceV2.getExecutedBlockMetadataFromIdentifier(identifier4)
	require.NotNil(t, err, "Should fail for non-existent block")
	require.Nil(t, metadata4, "Metadata should be nil for non-existent block")
	require.Equal(t, codes.NotFound, status.Code(err), "Should return NotFound error")
}

func TestExecutionServiceServerV2_CreateExecutionSession_HaltedFork(t *testing.T) {
	ethservice, serviceV2 := setupExecutionServiceWithHaltedFork(t, 10)

	currentFork := ethservice.BlockChain().Config().AstriaForks.GetForkAtHeight(1)
	require.True(t, currentFork.Halt, "Fork should be halted")

	_, err := serviceV2.CreateExecutionSession(context.Background(), &astriaPb.CreateExecutionSessionRequest{})
	require.NotNil(t, err, "CreateExecutionSession should fail with halted fork")
	require.Equal(t, codes.FailedPrecondition, status.Code(err), "Should get FailedPrecondition error")
}

func TestExecutionServiceServerV2_ExecuteBlock_InvalidSession(t *testing.T) {
	ethservice, serviceV2 := setupExecutionService(t, 10, false)

	softBlock := ethservice.BlockChain().CurrentSafeBlock()
	executeBlockReq := &astriaPb.ExecuteBlockRequest{
		SessionId:  "invalid-session-id",
		ParentHash: softBlock.Hash().Hex(),
		Timestamp: &timestamppb.Timestamp{
			Seconds: int64(softBlock.Time + 2),
		},
		Transactions: []*sequencerblockv1.RollupData{},
	}

	_, err := serviceV2.ExecuteBlock(context.Background(), executeBlockReq)
	require.NotNil(t, err, "ExecuteBlock should fail without a valid session")
	require.Equal(t, codes.PermissionDenied, status.Code(err), "Should get PermissionDenied error")
}

func TestExecutionServiceServerV2_ExecuteBlock_HaltedFork(t *testing.T) {
	ethservice, serviceV2 := setupExecutionServiceWithHaltedFork(t, 10)

	currentFork := ethservice.BlockChain().Config().AstriaForks.GetForkAtHeight(1)
	require.True(t, currentFork.Halt, "Fork should be halted")

	_, err := serviceV2.CreateExecutionSession(context.Background(), &astriaPb.CreateExecutionSessionRequest{})
	require.NotNil(t, err, "CreateExecutionSession should fail with halted fork")
	require.Equal(t, codes.FailedPrecondition, status.Code(err), "Should get FailedPrecondition error for halted fork")

	// Directly set the activeSessionId and activeFork to simulate a session that was created
	// before the fork was halted
	serviceV2.activeSessionId = "simulated-session-id"
	serviceV2.activeFork = &currentFork

	softBlock := ethservice.BlockChain().CurrentSafeBlock()
	executeBlockReq := &astriaPb.ExecuteBlockRequest{
		SessionId:  serviceV2.activeSessionId,
		ParentHash: softBlock.Hash().Hex(),
		Timestamp: &timestamppb.Timestamp{
			Seconds: int64(softBlock.Time + 2),
		},
		Transactions: []*sequencerblockv1.RollupData{},
	}

	_, err = serviceV2.ExecuteBlock(context.Background(), executeBlockReq)
	require.NotNil(t, err, "ExecuteBlock should fail with halted fork")
	require.Equal(t, codes.FailedPrecondition, status.Code(err), "Should get FailedPrecondition error")
}

func TestExecutionServiceServerV2_ExecuteBlock_OutOfRange(t *testing.T) {
	ethservice, serviceV2 := setupExecutionService(t, 10, false)

	session, err := serviceV2.CreateExecutionSession(context.Background(), &astriaPb.CreateExecutionSessionRequest{})
	require.Nil(t, err, "CreateExecutionSession failed")
	require.NotNil(t, session, "Session should not be nil")

	fork := ethservice.BlockChain().Config().AstriaForks.GetForkAtHeight(1)

	fork.StopHeight = 5
	serviceV2.activeFork = &fork

	softBlock := ethservice.BlockChain().GetBlockByNumber(6)
	executeBlockReq := &astriaPb.ExecuteBlockRequest{
		SessionId:  session.SessionId,
		ParentHash: softBlock.Hash().Hex(),
		Timestamp: &timestamppb.Timestamp{
			Seconds: int64(softBlock.Time() + 2),
		},
		Transactions: []*sequencerblockv1.RollupData{},
	}

	_, err = serviceV2.ExecuteBlock(context.Background(), executeBlockReq)
	require.NotNil(t, err, "ExecuteBlock should fail with out-of-range fork")
	require.Equal(t, codes.OutOfRange, status.Code(err), "Should get OutOfRange error")
}

func TestExecutionServiceServerV2_UpdateCommitmentState_OutOfRange(t *testing.T) {
	ethservice, serviceV2 := setupExecutionService(t, 10, false)

	session, err := serviceV2.CreateExecutionSession(context.Background(), &astriaPb.CreateExecutionSessionRequest{})
	require.Nil(t, err, "CreateExecutionSession failed")
	require.NotNil(t, session, "ExecutionSession is nil")

	fork := ethservice.BlockChain().Config().AstriaForks.GetForkAtHeight(1)

	fork.StopHeight = 5
	serviceV2.activeFork = &fork

	outOfRangeBlock := ethservice.BlockChain().GetBlockByNumber(9)
	require.NotNil(t, outOfRangeBlock, "Block 9 should exist")

	outOfRangeSoftBlockMetadata, err := ethHeaderToExecutedBlockMetadata(outOfRangeBlock.Header())
	require.Nil(t, err, "Failed to convert block header to metadata")

	firmBlock := ethservice.BlockChain().CurrentFinalBlock()
	firmBlockMetadata, err := ethHeaderToExecutedBlockMetadata(firmBlock)
	require.Nil(t, err, "Failed to convert firm block header to metadata")

	updateReq1 := &astriaPb.UpdateCommitmentStateRequest{
		SessionId: session.SessionId,
		CommitmentState: &astriaPb.CommitmentState{
			SoftExecutedBlockMetadata:  outOfRangeSoftBlockMetadata,
			FirmExecutedBlockMetadata:  firmBlockMetadata,
			LowestCelestiaSearchHeight: session.CommitmentState.LowestCelestiaSearchHeight + 1,
		},
	}

	_, err = serviceV2.UpdateCommitmentState(context.Background(), updateReq1)
	require.NotNil(t, err, "UpdateCommitmentState should fail with soft block out of range")
	require.Equal(t, codes.OutOfRange, status.Code(err), "Should get OutOfRange error")
}

func TestExecutionServiceServerV2_UpdateCommitmentState_NonCanonicalBlocks(t *testing.T) {
	_, serviceV2 := setupExecutionService(t, 10, false)

	session, err := serviceV2.CreateExecutionSession(context.Background(), &astriaPb.CreateExecutionSessionRequest{})
	require.Nil(t, err, "CreateExecutionSession failed")
	require.NotNil(t, session, "ExecutionSession is nil")

	nonExistentSoftBlock := &astriaPb.ExecutedBlockMetadata{
		Number:     100,
		Hash:       common.BytesToHash([]byte("non-existent-hash-1")).Hex(),
		ParentHash: common.BytesToHash([]byte("non-existent-parent-1")).Hex(),
		Timestamp:  &timestamppb.Timestamp{Seconds: int64(1000)},
	}

	nonExistentFirmBlock := &astriaPb.ExecutedBlockMetadata{
		Number:     99,
		Hash:       common.BytesToHash([]byte("non-existent-hash-2")).Hex(),
		ParentHash: common.BytesToHash([]byte("non-existent-parent-2")).Hex(),
		Timestamp:  &timestamppb.Timestamp{Seconds: int64(990)},
	}

	updateReq := &astriaPb.UpdateCommitmentStateRequest{
		SessionId: session.SessionId,
		CommitmentState: &astriaPb.CommitmentState{
			SoftExecutedBlockMetadata:  nonExistentSoftBlock,
			FirmExecutedBlockMetadata:  nonExistentFirmBlock,
			LowestCelestiaSearchHeight: session.CommitmentState.LowestCelestiaSearchHeight + 1,
		},
	}

	_, err = serviceV2.UpdateCommitmentState(context.Background(), updateReq)
	require.NotNil(t, err, "UpdateCommitmentState should fail with non-existent blocks")
	require.Equal(t, codes.InvalidArgument, status.Code(err), "Should get InvalidArgument error")
}

func TestExecutionServiceServerV2_ExecuteBlockTransactionOrdering(t *testing.T) {
	tests := []struct {
		description        string
		forkConfig         params.AstriaForkConfig
		numberOfRegularTxs int
		numberOfDepositTxs int
		expectedOrder      []params.AstriaTransactionType
	}{
		{
			description: "Set order (deposit before sequenced)",
			forkConfig: params.AstriaForkConfig{
				Height:              2,
				AppSpecificOrdering: []string{"deposit", "sequencedData", "priceFeedData"},
			},
			numberOfRegularTxs: 3,
			numberOfDepositTxs: 2,
			expectedOrder:      []params.AstriaTransactionType{params.Deposit, params.Deposit, params.SequencedData, params.SequencedData, params.SequencedData},
		},
		{
			description: "Reversed order (sequenced before deposit)",
			forkConfig: params.AstriaForkConfig{
				Height:              3,
				AppSpecificOrdering: []string{"sequencedData", "deposit", "priceFeedData"},
			},
			numberOfRegularTxs: 3,
			numberOfDepositTxs: 2,
			expectedOrder:      []params.AstriaTransactionType{params.SequencedData, params.SequencedData, params.SequencedData, params.Deposit, params.Deposit},
		},
		{
			description: "No ordering",
			forkConfig: params.AstriaForkConfig{
				Height:              4,
				AppSpecificOrdering: nil,
			},
			numberOfRegularTxs: 3,
			numberOfDepositTxs: 2,
			// This order is the order that the test txs are created in
			expectedOrder: []params.AstriaTransactionType{params.SequencedData, params.SequencedData, params.SequencedData, params.Deposit, params.Deposit},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			// Setup execution service with custom fork configuration
			ethservice, serviceV2 := setupExecutionService(t, 10, false)

			// Create a fork with the app specific ordering
			var aso []params.AstriaTransactionType
			if tt.forkConfig.AppSpecificOrdering != nil {
				for _, txType := range tt.forkConfig.AppSpecificOrdering {
					aso = append(aso, params.AstriaTransactionTypeMap[txType])
				}
			} else {
				aso = nil
			}
			// Get the default fork at height 1 to copy bridge addresses and other required fields
			defaultFork := ethservice.BlockChain().Config().AstriaForks.GetForkAtHeight(1)
			fork := params.AstriaForkData{
				Height:              tt.forkConfig.Height,
				AppSpecificOrdering: aso,
				BridgeAddresses:     defaultFork.BridgeAddresses,
				BridgeAllowedAssets: defaultFork.BridgeAllowedAssets,
				Sequencer:           defaultFork.Sequencer,
				Celestia:            defaultFork.Celestia,
				EIP1559Params:       defaultFork.EIP1559Params,
				FeeCollector:        defaultFork.FeeCollector,
				Oracle:              defaultFork.Oracle,
				Precompiles:         defaultFork.Precompiles,
			}

			// Create execution session
			session, err := serviceV2.createExecutionSessionWithForkOverride(context.Background(), &astriaPb.CreateExecutionSessionRequest{}, fork)
			require.Nil(t, err, "CreateExecutionSession failed")
			require.NotNil(t, session, "ExecutionSession is nil")

			// Create regular transactions
			marshalledTxs := []*sequencerblockv1.RollupData{}
			for i := 0; i < tt.numberOfRegularTxs; i++ {
				unsignedTx := types.NewTransaction(uint64(i), testToAddress, big.NewInt(1), params.TxGas, big.NewInt(params.InitialBaseFee*2), nil)
				tx, err := types.SignTx(unsignedTx, types.LatestSigner(ethservice.BlockChain().Config()), testKey)
				require.Nil(t, err, "Failed to sign tx")

				marshalledTx, err := tx.MarshalBinary()
				require.Nil(t, err, "Failed to marshal tx")
				marshalledTxs = append(marshalledTxs, &sequencerblockv1.RollupData{
					Value: &sequencerblockv1.RollupData_SequencedData{SequencedData: marshalledTx},
				})
			}

			// Create deposit transactions
			bf := ethservice.BlockChain().Config().AstriaForks.GetForkAtHeight(1)
			var bridgeConfig params.AstriaBridgeAddressConfig
			for _, cfg := range bf.BridgeAddresses {
				bridgeConfig = *cfg
				break
			}

			for i := 0; i < tt.numberOfDepositTxs; i++ {
				chainDestinationAddressPrivKey, err := crypto.GenerateKey()
				require.Nil(t, err, "Failed to generate chain destination address")
				chainDestinationAddress := crypto.PubkeyToAddress(chainDestinationAddressPrivKey.PublicKey)

				hashedRollupId := sha256.Sum256([]byte(ethservice.BlockChain().Config().AstriaRollupName))
				rollupId := primitivev1.RollupId{
					Inner: hashedRollupId[:],
				}

				depositTx := &sequencerblockv1.RollupData{
					Value: &sequencerblockv1.RollupData_Deposit{
						Deposit: &sequencerblockv1.Deposit{
							BridgeAddress: &primitivev1.Address{
								Bech32M: bridgeConfig.BridgeAddress,
							},
							Asset:                   bridgeConfig.AssetDenom,
							Amount:                  bigIntToProtoU128(big.NewInt(1000000000000000000)),
							RollupId:                &rollupId,
							DestinationChainAddress: chainDestinationAddress.String(),
							SourceTransactionId: &primitivev1.TransactionId{
								Inner: fmt.Sprintf("test_tx_hash_%d", i),
							},
							SourceActionIndex: 0,
						},
					},
				}
				marshalledTxs = append(marshalledTxs, depositTx)
			}

			// Execute block
			softBlock := ethservice.BlockChain().CurrentSafeBlock()
			executeBlockReq := &astriaPb.ExecuteBlockRequest{
				SessionId:  session.SessionId,
				ParentHash: softBlock.Hash().Hex(),
				Timestamp: &timestamppb.Timestamp{
					Seconds: int64(softBlock.Time + 2),
				},
				Transactions: marshalledTxs,
			}

			executeBlockRes, err := serviceV2.ExecuteBlock(context.Background(), executeBlockReq)
			require.Nil(t, err, "ExecuteBlock failed")
			require.NotNil(t, executeBlockRes, "ExecuteBlock response is nil")

			// Verify transaction ordering in the block
			block := ethservice.BlockChain().GetBlockByHash(common.HexToHash(executeBlockRes.ExecutedBlockMetadata.Hash))
			require.NotNil(t, block, "Block not found")
			require.Equal(t, tt.numberOfRegularTxs+tt.numberOfDepositTxs, len(block.Transactions()), "Transaction count mismatch")

			// Verify the transaction type ordering
			for i, tx := range block.Transactions() {
				var txType params.AstriaTransactionType
				if tx.Type() == types.InjectedTxType {
					txType = params.Deposit
				} else {
					txType = params.SequencedData
				}
				require.Equal(t, tt.expectedOrder[i], txType, "Transaction at index %d has incorrect type", i)
			}
		})
	}
}
