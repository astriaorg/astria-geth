package execution

import (
	astriaPb "buf.build/gen/go/astria/execution-apis/protocolbuffers/go/astria/execution/v1alpha2"
	primitivev1 "buf.build/gen/go/astria/primitives/protocolbuffers/go/astria/primitive/v1"
	sequencerblockv1alpha1 "buf.build/gen/go/astria/sequencerblock-apis/protocolbuffers/go/astria/sequencerblock/v1alpha1"
	"bytes"
	"context"
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math/big"
	"testing"
)

func TestExecutionService_GetGenesisInfo(t *testing.T) {
	ethservice, serviceV1Alpha1 := setupExecutionService(t, 10)

	genesisInfo, err := serviceV1Alpha1.GetGenesisInfo(context.Background(), &astriaPb.GetGenesisInfoRequest{})
	require.Nil(t, err, "GetGenesisInfo failed")

	hashedRollupId := sha256.Sum256([]byte(ethservice.BlockChain().Config().AstriaRollupName))

	require.True(t, bytes.Equal(genesisInfo.RollupId, hashedRollupId[:]), "RollupId is not correct")
	require.Equal(t, genesisInfo.GetSequencerGenesisBlockHeight(), ethservice.BlockChain().Config().AstriaSequencerInitialHeight, "SequencerInitialHeight is not correct")
	require.Equal(t, genesisInfo.GetCelestiaBaseBlockHeight(), ethservice.BlockChain().Config().AstriaCelestiaInitialHeight, "CelestiaInitialHeight is not correct")
	require.Equal(t, genesisInfo.GetCelestiaBlockVariance(), ethservice.BlockChain().Config().AstriaCelestiaHeightVariance, "CelestiaHeightVariance is not correct")
	require.True(t, serviceV1Alpha1.genesisInfoCalled, "GetGenesisInfo should be called")
}

func TestExecutionServiceServerV1Alpha2_GetCommitmentState(t *testing.T) {
	ethservice, serviceV1Alpha1 := setupExecutionService(t, 10)

	commitmentState, err := serviceV1Alpha1.GetCommitmentState(context.Background(), &astriaPb.GetCommitmentStateRequest{})
	require.Nil(t, err, "GetCommitmentState failed")

	require.NotNil(t, commitmentState, "CommitmentState is nil")

	softBlock := ethservice.BlockChain().CurrentSafeBlock()
	require.NotNil(t, softBlock, "SoftBlock is nil")

	firmBlock := ethservice.BlockChain().CurrentFinalBlock()
	require.NotNil(t, firmBlock, "FirmBlock is nil")

	require.True(t, bytes.Equal(commitmentState.Soft.Hash, softBlock.Hash().Bytes()), "Soft Block Hashes do not match")
	require.True(t, bytes.Equal(commitmentState.Soft.ParentBlockHash, softBlock.ParentHash.Bytes()), "Soft Block Parent Hash do not match")
	require.Equal(t, uint64(commitmentState.Soft.Number), softBlock.Number.Uint64(), "Soft Block Number do not match")

	require.True(t, bytes.Equal(commitmentState.Firm.Hash, firmBlock.Hash().Bytes()), "Firm Block Hashes do not match")
	require.True(t, bytes.Equal(commitmentState.Firm.ParentBlockHash, firmBlock.ParentHash.Bytes()), "Firm Block Parent Hash do not match")
	require.Equal(t, uint64(commitmentState.Firm.Number), firmBlock.Number.Uint64(), "Firm Block Number do not match")

	require.True(t, serviceV1Alpha1.getCommitmentStateCalled, "GetCommitmentState should be called")
}

func TestExecutionService_GetBlock(t *testing.T) {
	ethservice, serviceV1Alpha1 := setupExecutionService(t, 10)

	tests := []struct {
		description        string
		getBlockRequst     *astriaPb.GetBlockRequest
		expectedReturnCode codes.Code
	}{
		{
			description: "Get block by block number 1",
			getBlockRequst: &astriaPb.GetBlockRequest{
				Identifier: &astriaPb.BlockIdentifier{Identifier: &astriaPb.BlockIdentifier_BlockNumber{BlockNumber: 1}},
			},
			expectedReturnCode: 0,
		},
		{
			description: "Get block by block hash",
			getBlockRequst: &astriaPb.GetBlockRequest{
				Identifier: &astriaPb.BlockIdentifier{Identifier: &astriaPb.BlockIdentifier_BlockHash{BlockHash: ethservice.BlockChain().GetBlockByNumber(4).Hash().Bytes()}},
			},
			expectedReturnCode: 0,
		},
		{
			description: "Get block which is not present",
			getBlockRequst: &astriaPb.GetBlockRequest{
				Identifier: &astriaPb.BlockIdentifier{Identifier: &astriaPb.BlockIdentifier_BlockNumber{BlockNumber: 100}},
			},
			expectedReturnCode: codes.NotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			blockInfo, err := serviceV1Alpha1.GetBlock(context.Background(), tt.getBlockRequst)
			if tt.expectedReturnCode > 0 {
				require.NotNil(t, err, "GetBlock should return an error")
				require.Equal(t, tt.expectedReturnCode, status.Code(err), "GetBlock failed")
			}
			if err == nil {
				require.NotNil(t, blockInfo, "Block not found")
				var block *types.Block
				if tt.getBlockRequst.Identifier.GetBlockNumber() != 0 {
					// get block by number
					block = ethservice.BlockChain().GetBlockByNumber(uint64(tt.getBlockRequst.Identifier.GetBlockNumber()))
				}
				if tt.getBlockRequst.Identifier.GetBlockHash() != nil {
					block = ethservice.BlockChain().GetBlockByHash(common.Hash(tt.getBlockRequst.Identifier.GetBlockHash()))
				}
				require.NotNil(t, block, "Block not found")

				require.Equal(t, uint64(blockInfo.Number), block.NumberU64(), "Block number is not correct")
				require.Equal(t, block.ParentHash().Bytes(), blockInfo.ParentBlockHash, "Parent Block Hash is not correct")
				require.Equal(t, block.Hash().Bytes(), blockInfo.Hash, "BlockHash is not correct")
			}
		})

	}
}

func TestExecutionServiceServerV1Alpha2_BatchGetBlocks(t *testing.T) {
	ethservice, serviceV1Alpha1 := setupExecutionService(t, 10)

	tests := []struct {
		description          string
		batchGetBlockRequest *astriaPb.BatchGetBlocksRequest
		expectedReturnCode   codes.Code
	}{
		{
			description: "BatchGetBlocks with block hashes",
			batchGetBlockRequest: &astriaPb.BatchGetBlocksRequest{
				Identifiers: []*astriaPb.BlockIdentifier{
					{Identifier: &astriaPb.BlockIdentifier_BlockHash{BlockHash: ethservice.BlockChain().GetBlockByNumber(1).Hash().Bytes()}},
					{Identifier: &astriaPb.BlockIdentifier_BlockHash{BlockHash: ethservice.BlockChain().GetBlockByNumber(2).Hash().Bytes()}},
					{Identifier: &astriaPb.BlockIdentifier_BlockHash{BlockHash: ethservice.BlockChain().GetBlockByNumber(3).Hash().Bytes()}},
					{Identifier: &astriaPb.BlockIdentifier_BlockHash{BlockHash: ethservice.BlockChain().GetBlockByNumber(4).Hash().Bytes()}},
					{Identifier: &astriaPb.BlockIdentifier_BlockHash{BlockHash: ethservice.BlockChain().GetBlockByNumber(5).Hash().Bytes()}},
				},
			},
			expectedReturnCode: 0,
		},
		{
			description: "BatchGetBlocks with block numbers",
			batchGetBlockRequest: &astriaPb.BatchGetBlocksRequest{
				Identifiers: []*astriaPb.BlockIdentifier{
					{Identifier: &astriaPb.BlockIdentifier_BlockNumber{BlockNumber: 1}},
					{Identifier: &astriaPb.BlockIdentifier_BlockNumber{BlockNumber: 2}},
					{Identifier: &astriaPb.BlockIdentifier_BlockNumber{BlockNumber: 3}},
					{Identifier: &astriaPb.BlockIdentifier_BlockNumber{BlockNumber: 4}},
					{Identifier: &astriaPb.BlockIdentifier_BlockNumber{BlockNumber: 5}},
				},
			},
			expectedReturnCode: 0,
		},
		{
			description: "BatchGetBlocks block not found",
			batchGetBlockRequest: &astriaPb.BatchGetBlocksRequest{
				Identifiers: []*astriaPb.BlockIdentifier{
					{Identifier: &astriaPb.BlockIdentifier_BlockNumber{BlockNumber: 1}},
					{Identifier: &astriaPb.BlockIdentifier_BlockNumber{BlockNumber: 2}},
					{Identifier: &astriaPb.BlockIdentifier_BlockNumber{BlockNumber: 3}},
					{Identifier: &astriaPb.BlockIdentifier_BlockNumber{BlockNumber: 4}},
					{Identifier: &astriaPb.BlockIdentifier_BlockNumber{BlockNumber: 100}},
				},
			},
			expectedReturnCode: codes.NotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			batchBlocksRes, err := serviceV1Alpha1.BatchGetBlocks(context.Background(), tt.batchGetBlockRequest)
			if tt.expectedReturnCode > 0 {
				require.NotNil(t, err, "BatchGetBlocks should return an error")
				require.Equal(t, tt.expectedReturnCode, status.Code(err), "BatchGetBlocks failed")
			}

			for _, batchBlock := range batchBlocksRes.GetBlocks() {
				require.NotNil(t, batchBlock, "Block not found in batch blocks response")

				block := ethservice.BlockChain().GetBlockByNumber(uint64(batchBlock.Number))
				require.NotNil(t, block, "Block not found in blockchain")

				require.Equal(t, uint64(batchBlock.Number), block.NumberU64(), "Block number is not correct")
				require.Equal(t, block.ParentHash().Bytes(), batchBlock.ParentBlockHash, "Parent Block Hash is not correct")
				require.Equal(t, block.Hash().Bytes(), batchBlock.Hash, "BlockHash is not correct")
			}
		})
	}
}

func bigIntToProtoU128(i *big.Int) *primitivev1.Uint128 {
	lo := i.Uint64()
	hi := new(big.Int).Rsh(i, 64).Uint64()
	return &primitivev1.Uint128{Lo: lo, Hi: hi}
}

func TestExecutionServiceServerV1Alpha2_ExecuteBlock(t *testing.T) {
	ethservice, _ := setupExecutionService(t, 10)

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
			description:                          "ExecuteBlock without calling GetGenesisInfo and GetCommitmentState",
			callGenesisInfoAndGetCommitmentState: false,
			numberOfTxs:                          5,
			prevBlockHash:                        ethservice.BlockChain().GetBlockByNumber(2).Hash().Bytes(),
			timestamp:                            ethservice.BlockChain().GetBlockByNumber(2).Time() + 2,
			depositTxAmount:                      big.NewInt(0),
			expectedReturnCode:                   codes.PermissionDenied,
		},
		{
			description:                          "ExecuteBlock with 5 txs and no deposit tx",
			callGenesisInfoAndGetCommitmentState: true,
			numberOfTxs:                          5,
			prevBlockHash:                        ethservice.BlockChain().CurrentSafeBlock().Hash().Bytes(),
			timestamp:                            ethservice.BlockChain().CurrentSafeBlock().Time + 2,
			depositTxAmount:                      big.NewInt(0),
			expectedReturnCode:                   0,
		},
		{
			description:                          "ExecuteBlock with 5 txs and a deposit tx",
			callGenesisInfoAndGetCommitmentState: true,
			numberOfTxs:                          5,
			prevBlockHash:                        ethservice.BlockChain().CurrentSafeBlock().Hash().Bytes(),
			timestamp:                            ethservice.BlockChain().CurrentSafeBlock().Time + 2,
			depositTxAmount:                      big.NewInt(1000000000000000000),
			expectedReturnCode:                   0,
		},
		{
			description:                          "ExecuteBlock with incorrect previous block hash",
			callGenesisInfoAndGetCommitmentState: true,
			numberOfTxs:                          5,
			prevBlockHash:                        ethservice.BlockChain().GetBlockByNumber(2).Hash().Bytes(),
			timestamp:                            ethservice.BlockChain().GetBlockByNumber(2).Time() + 2,
			depositTxAmount:                      big.NewInt(0),
			expectedReturnCode:                   codes.FailedPrecondition,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			// reset the blockchain with each test
			ethservice, serviceV1Alpha1 := setupExecutionService(t, 10)

			var err error // adding this to prevent shadowing of genesisInfo in the below if branch
			var genesisInfo *astriaPb.GenesisInfo
			var commitmentStateBeforeExecuteBlock *astriaPb.CommitmentState
			if tt.callGenesisInfoAndGetCommitmentState {
				// call getGenesisInfo and getCommitmentState before calling executeBlock
				genesisInfo, err = serviceV1Alpha1.GetGenesisInfo(context.Background(), &astriaPb.GetGenesisInfoRequest{})
				require.Nil(t, err, "GetGenesisInfo failed")
				require.NotNil(t, genesisInfo, "GenesisInfo is nil")

				commitmentStateBeforeExecuteBlock, err = serviceV1Alpha1.GetCommitmentState(context.Background(), &astriaPb.GetCommitmentStateRequest{})
				require.Nil(t, err, "GetCommitmentState failed")
				require.NotNil(t, commitmentStateBeforeExecuteBlock, "CommitmentState is nil")
			}

			// create the txs to send
			// create 5 txs
			txs := []*types.Transaction{}
			marshalledTxs := []*sequencerblockv1alpha1.RollupData{}
			for i := 0; i < 5; i++ {
				unsignedTx := types.NewTransaction(uint64(i), testToAddress, big.NewInt(1), params.TxGas, big.NewInt(params.InitialBaseFee*2), nil)
				tx, err := types.SignTx(unsignedTx, types.LatestSigner(ethservice.BlockChain().Config()), testKey)
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
				depositAmount := bigIntToProtoU128(tt.depositTxAmount)
				bridgeAddress := ethservice.BlockChain().Config().AstriaBridgeAddressConfigs[0].BridgeAddress
				bridgeAssetDenom := sha256.Sum256([]byte(ethservice.BlockChain().Config().AstriaBridgeAddressConfigs[0].AssetDenom))

				// create new chain destination address for better testing
				chainDestinationAddressPrivKey, err := crypto.GenerateKey()
				require.Nil(t, err, "Failed to generate chain destination address")

				chainDestinationAddress := crypto.PubkeyToAddress(chainDestinationAddressPrivKey.PublicKey)

				depositTx := &sequencerblockv1alpha1.RollupData{Value: &sequencerblockv1alpha1.RollupData_Deposit{Deposit: &sequencerblockv1alpha1.Deposit{
					BridgeAddress: &primitivev1.Address{
						Inner: bridgeAddress,
					},
					AssetId:                 bridgeAssetDenom[:],
					Amount:                  depositAmount,
					RollupId:                &primitivev1.RollupId{Inner: genesisInfo.RollupId},
					DestinationChainAddress: chainDestinationAddress.String(),
				}}}

				marshalledTxs = append(marshalledTxs, depositTx)
			}

			executeBlockReq := &astriaPb.ExecuteBlockRequest{
				PrevBlockHash: tt.prevBlockHash,
				Timestamp: &timestamppb.Timestamp{
					Seconds: int64(tt.timestamp),
				},
				Transactions: marshalledTxs,
			}

			executeBlockRes, err := serviceV1Alpha1.ExecuteBlock(context.Background(), executeBlockReq)
			if tt.expectedReturnCode > 0 {
				require.NotNil(t, err, "ExecuteBlock should return an error")
				require.Equal(t, tt.expectedReturnCode, status.Code(err), "ExecuteBlock failed")
			}
			if err == nil {
				require.NotNil(t, executeBlockRes, "ExecuteBlock response is nil")

				astriaOrdered := ethservice.TxPool().AstriaOrdered()
				require.Equal(t, 0, astriaOrdered.Len(), "AstriaOrdered should be empty")

				// check if commitment state is not updated
				commitmentStateAfterExecuteBlock, err := serviceV1Alpha1.GetCommitmentState(context.Background(), &astriaPb.GetCommitmentStateRequest{})
				require.Nil(t, err, "GetCommitmentState failed")

				require.Exactly(t, commitmentStateBeforeExecuteBlock, commitmentStateAfterExecuteBlock, "Commitment state should not be updated")
			}

		})
	}
}

func TestExecutionServiceServerV1Alpha2_ExecuteBlockAndUpdateCommitment(t *testing.T) {
	ethservice, serviceV1Alpha1 := setupExecutionService(t, 10)

	// call genesis info
	genesisInfo, err := serviceV1Alpha1.GetGenesisInfo(context.Background(), &astriaPb.GetGenesisInfoRequest{})
	require.Nil(t, err, "GetGenesisInfo failed")
	require.NotNil(t, genesisInfo, "GenesisInfo is nil")

	// call get commitment state
	commitmentState, err := serviceV1Alpha1.GetCommitmentState(context.Background(), &astriaPb.GetCommitmentStateRequest{})
	require.Nil(t, err, "GetCommitmentState failed")
	require.NotNil(t, commitmentState, "CommitmentState is nil")

	// get previous block hash
	previousBlock := ethservice.BlockChain().CurrentSafeBlock()
	require.NotNil(t, previousBlock, "Previous block not found")

	// create 5 txs
	txs := []*types.Transaction{}
	marshalledTxs := []*sequencerblockv1alpha1.RollupData{}
	for i := 0; i < 5; i++ {
		unsignedTx := types.NewTransaction(uint64(i), testToAddress, big.NewInt(1), params.TxGas, big.NewInt(params.InitialBaseFee*2), nil)
		tx, err := types.SignTx(unsignedTx, types.LatestSigner(ethservice.BlockChain().Config()), testKey)
		require.Nil(t, err, "Failed to sign tx")
		txs = append(txs, tx)

		marshalledTx, err := tx.MarshalBinary()
		require.Nil(t, err, "Failed to marshal tx")
		marshalledTxs = append(marshalledTxs, &sequencerblockv1alpha1.RollupData{
			Value: &sequencerblockv1alpha1.RollupData_SequencedData{SequencedData: marshalledTx},
		})
	}

	amountToDeposit := big.NewInt(1000000000000000000)
	depositAmount := bigIntToProtoU128(amountToDeposit)
	bridgeAddress := ethservice.BlockChain().Config().AstriaBridgeAddressConfigs[0].BridgeAddress
	bridgeAssetDenom := sha256.Sum256([]byte(ethservice.BlockChain().Config().AstriaBridgeAddressConfigs[0].AssetDenom))

	// create new chain destination address for better testing
	chainDestinationAddressPrivKey, err := crypto.GenerateKey()
	require.Nil(t, err, "Failed to generate chain destination address")

	chainDestinationAddress := crypto.PubkeyToAddress(chainDestinationAddressPrivKey.PublicKey)

	stateDb, err := ethservice.BlockChain().State()
	require.Nil(t, err, "Failed to get state db")
	require.NotNil(t, stateDb, "State db is nil")

	chainDestinationAddressBalanceBefore := stateDb.GetBalance(chainDestinationAddress)

	depositTx := &sequencerblockv1alpha1.RollupData{Value: &sequencerblockv1alpha1.RollupData_Deposit{Deposit: &sequencerblockv1alpha1.Deposit{
		BridgeAddress: &primitivev1.Address{
			Inner: bridgeAddress,
		},
		AssetId:                 bridgeAssetDenom[:],
		Amount:                  depositAmount,
		RollupId:                &primitivev1.RollupId{Inner: genesisInfo.RollupId},
		DestinationChainAddress: chainDestinationAddress.String(),
	}}}

	marshalledTxs = append(marshalledTxs, depositTx)

	executeBlockReq := &astriaPb.ExecuteBlockRequest{
		PrevBlockHash: previousBlock.Hash().Bytes(),
		Timestamp: &timestamppb.Timestamp{
			Seconds: int64(previousBlock.Time + 2),
		},
		Transactions: marshalledTxs,
	}

	executeBlockRes, err := serviceV1Alpha1.ExecuteBlock(context.Background(), executeBlockReq)
	require.Nil(t, err, "ExecuteBlock failed")

	require.NotNil(t, executeBlockRes, "ExecuteBlock response is nil")

	// check if astria ordered txs are cleared
	astriaOrdered := ethservice.TxPool().AstriaOrdered()
	require.Equal(t, 0, astriaOrdered.Len(), "AstriaOrdered should be empty")

	// call update commitment state to set the block we executed as soft and firm
	updateCommitmentStateReq := &astriaPb.UpdateCommitmentStateRequest{
		CommitmentState: &astriaPb.CommitmentState{
			Soft: &astriaPb.Block{
				Hash:            executeBlockRes.Hash,
				ParentBlockHash: executeBlockRes.ParentBlockHash,
				Number:          executeBlockRes.Number,
				Timestamp:       executeBlockRes.Timestamp,
			},
			Firm: &astriaPb.Block{
				Hash:            executeBlockRes.Hash,
				ParentBlockHash: executeBlockRes.ParentBlockHash,
				Number:          executeBlockRes.Number,
				Timestamp:       executeBlockRes.Timestamp,
			},
		},
	}

	updateCommitmentStateRes, err := serviceV1Alpha1.UpdateCommitmentState(context.Background(), updateCommitmentStateReq)
	require.Nil(t, err, "UpdateCommitmentState failed")
	require.NotNil(t, updateCommitmentStateRes, "UpdateCommitmentState response should not be nil")

	// get the soft and firm block
	softBlock := ethservice.BlockChain().CurrentSafeBlock()
	require.NotNil(t, softBlock, "SoftBlock is nil")
	firmBlock := ethservice.BlockChain().CurrentFinalBlock()
	require.NotNil(t, firmBlock, "FirmBlock is nil")

	// check if the soft and firm block are set correctly
	require.True(t, bytes.Equal(softBlock.Hash().Bytes(), updateCommitmentStateRes.Soft.Hash), "Soft Block Hashes do not match")
	require.True(t, bytes.Equal(softBlock.ParentHash.Bytes(), updateCommitmentStateRes.Soft.ParentBlockHash), "Soft Block Parent Hash do not match")
	require.Equal(t, softBlock.Number.Uint64(), uint64(updateCommitmentStateRes.Soft.Number), "Soft Block Number do not match")

	require.True(t, bytes.Equal(firmBlock.Hash().Bytes(), updateCommitmentStateRes.Firm.Hash), "Firm Block Hashes do not match")
	require.True(t, bytes.Equal(firmBlock.ParentHash.Bytes(), updateCommitmentStateRes.Firm.ParentBlockHash), "Firm Block Parent Hash do not match")
	require.Equal(t, firmBlock.Number.Uint64(), uint64(updateCommitmentStateRes.Firm.Number), "Firm Block Number do not match")

	// check the difference in balances after deposit tx
	stateDb, err = ethservice.BlockChain().State()
	require.Nil(t, err, "Failed to get state db")
	require.NotNil(t, stateDb, "State db is nil")
	chainDestinationAddressBalanceAfter := stateDb.GetBalance(chainDestinationAddress)

	balanceDiff := new(big.Int).Sub(chainDestinationAddressBalanceAfter, chainDestinationAddressBalanceBefore)
	require.True(t, balanceDiff.Cmp(big.NewInt(1000000000000000000)) == 0, "Chain destination address balance is not correct")
}
