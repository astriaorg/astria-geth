package execution

import (
	optimsticPb "buf.build/gen/go/astria/execution-apis/protocolbuffers/go/astria/bundle/v1alpha1"
	astriaPb "buf.build/gen/go/astria/execution-apis/protocolbuffers/go/astria/execution/v1"
	primitivev1 "buf.build/gen/go/astria/primitives/protocolbuffers/go/astria/primitive/v1"
	sequencerblockv1 "buf.build/gen/go/astria/sequencerblock-apis/protocolbuffers/go/astria/sequencerblock/v1"
	"bytes"
	"context"
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math/big"
	"testing"
	"time"
)

func TestExecutionService_GetGenesisInfo(t *testing.T) {
	ethservice, serviceV1Alpha1 := setupExecutionService(t, 10)

	genesisInfo, err := serviceV1Alpha1.GetGenesisInfo(context.Background(), &astriaPb.GetGenesisInfoRequest{})
	require.Nil(t, err, "GetGenesisInfo failed")

	hashedRollupId := sha256.Sum256([]byte(ethservice.BlockChain().Config().AstriaRollupName))

	require.True(t, bytes.Equal(genesisInfo.RollupId.Inner, hashedRollupId[:]), "RollupId is not correct")
	require.Equal(t, genesisInfo.GetSequencerGenesisBlockHeight(), ethservice.BlockChain().Config().AstriaSequencerInitialHeight, "SequencerInitialHeight is not correct")
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
	require.Equal(t, commitmentState.BaseCelestiaHeight, ethservice.BlockChain().Config().AstriaCelestiaInitialHeight, "BaseCelestiaHeight is not correct")

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
			marshalledTxs := []*sequencerblockv1.RollupData{}
			for i := 0; i < 5; i++ {
				unsignedTx := types.NewTransaction(uint64(i), testToAddress, big.NewInt(1), params.TxGas, big.NewInt(params.InitialBaseFee*2), nil)
				tx, err := types.SignTx(unsignedTx, types.LatestSigner(ethservice.BlockChain().Config()), testKey)
				require.Nil(t, err, "Failed to sign tx")
				txs = append(txs, tx)

				marshalledTx, err := tx.MarshalBinary()
				require.Nil(t, err, "Failed to marshal tx")
				marshalledTxs = append(marshalledTxs, &sequencerblockv1.RollupData{
					Value: &sequencerblockv1.RollupData_SequencedData{SequencedData: marshalledTx},
				})
			}

			// create deposit tx if depositTxAmount is non zero
			if tt.depositTxAmount.Cmp(big.NewInt(0)) != 0 {
				depositAmount := bigIntToProtoU128(tt.depositTxAmount)
				bridgeAddress := ethservice.BlockChain().Config().AstriaBridgeAddressConfigs[0].BridgeAddress
				bridgeAssetDenom := ethservice.BlockChain().Config().AstriaBridgeAddressConfigs[0].AssetDenom

				// create new chain destination address for better testing
				chainDestinationAddressPrivKey, err := crypto.GenerateKey()
				require.Nil(t, err, "Failed to generate chain destination address")

				chainDestinationAddress := crypto.PubkeyToAddress(chainDestinationAddressPrivKey.PublicKey)

				depositTx := &sequencerblockv1.RollupData{Value: &sequencerblockv1.RollupData_Deposit{Deposit: &sequencerblockv1.Deposit{
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
	marshalledTxs := []*sequencerblockv1.RollupData{}
	for i := 0; i < 5; i++ {
		unsignedTx := types.NewTransaction(uint64(i), testToAddress, big.NewInt(1), params.TxGas, big.NewInt(params.InitialBaseFee*2), nil)
		tx, err := types.SignTx(unsignedTx, types.LatestSigner(ethservice.BlockChain().Config()), testKey)
		require.Nil(t, err, "Failed to sign tx")
		txs = append(txs, tx)

		marshalledTx, err := tx.MarshalBinary()
		require.Nil(t, err, "Failed to marshal tx")
		marshalledTxs = append(marshalledTxs, &sequencerblockv1.RollupData{
			Value: &sequencerblockv1.RollupData_SequencedData{SequencedData: marshalledTx},
		})
	}

	amountToDeposit := big.NewInt(1000000000000000000)
	depositAmount := bigIntToProtoU128(amountToDeposit)
	bridgeAddress := ethservice.BlockChain().Config().AstriaBridgeAddressConfigs[0].BridgeAddress
	bridgeAssetDenom := ethservice.BlockChain().Config().AstriaBridgeAddressConfigs[0].AssetDenom

	// create new chain destination address for better testing
	chainDestinationAddressPrivKey, err := crypto.GenerateKey()
	require.Nil(t, err, "Failed to generate chain destination address")

	chainDestinationAddress := crypto.PubkeyToAddress(chainDestinationAddressPrivKey.PublicKey)

	stateDb, err := ethservice.BlockChain().State()
	require.Nil(t, err, "Failed to get state db")
	require.NotNil(t, stateDb, "State db is nil")

	chainDestinationAddressBalanceBefore := stateDb.GetBalance(chainDestinationAddress)

	depositTx := &sequencerblockv1.RollupData{Value: &sequencerblockv1.RollupData_Deposit{Deposit: &sequencerblockv1.Deposit{
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
			BaseCelestiaHeight: commitmentState.BaseCelestiaHeight + 1,
		},
	}

	updateCommitmentStateRes, err := serviceV1Alpha1.UpdateCommitmentState(context.Background(), updateCommitmentStateReq)
	require.Nil(t, err, "UpdateCommitmentState failed")
	require.NotNil(t, updateCommitmentStateRes, "UpdateCommitmentState response should not be nil")
	require.Equal(t, updateCommitmentStateRes, updateCommitmentStateReq.CommitmentState, "CommitmentState response should match request")

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

	celestiaBaseHeight := ethservice.BlockChain().CurrentBaseCelestiaHeight()
	require.Equal(t, celestiaBaseHeight, updateCommitmentStateRes.BaseCelestiaHeight, "BaseCelestiaHeight should be updated in db")

	// check the difference in balances after deposit tx
	stateDb, err = ethservice.BlockChain().State()
	require.Nil(t, err, "Failed to get state db")
	require.NotNil(t, stateDb, "State db is nil")
	chainDestinationAddressBalanceAfter := stateDb.GetBalance(chainDestinationAddress)

	balanceDiff := new(uint256.Int).Sub(chainDestinationAddressBalanceAfter, chainDestinationAddressBalanceBefore)
	require.True(t, balanceDiff.Cmp(uint256.NewInt(1000000000000000000)) == 0, "Chain destination address balance is not correct")
}

func TestExecutionServiceServerV1Alpha2_ExecuteOptimisticBlock(t *testing.T) {
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
			description:                          "ExecuteOptimisticBlock without calling GetGenesisInfo and GetCommitmentState",
			callGenesisInfoAndGetCommitmentState: false,
			numberOfTxs:                          5,
			prevBlockHash:                        ethservice.BlockChain().GetBlockByNumber(2).Hash().Bytes(),
			timestamp:                            ethservice.BlockChain().GetBlockByNumber(2).Time() + 2,
			depositTxAmount:                      big.NewInt(0),
			expectedReturnCode:                   codes.PermissionDenied,
		},
		{
			description:                          "ExecuteOptimisticBlock with 5 txs and no deposit tx",
			callGenesisInfoAndGetCommitmentState: true,
			numberOfTxs:                          5,
			prevBlockHash:                        ethservice.BlockChain().CurrentSafeBlock().Hash().Bytes(),
			timestamp:                            ethservice.BlockChain().CurrentSafeBlock().Time + 2,
			depositTxAmount:                      big.NewInt(0),
			expectedReturnCode:                   0,
		},
		{
			description:                          "ExecuteOptimisticBlock with 5 txs and a deposit tx",
			callGenesisInfoAndGetCommitmentState: true,
			numberOfTxs:                          5,
			prevBlockHash:                        ethservice.BlockChain().CurrentSafeBlock().Hash().Bytes(),
			timestamp:                            ethservice.BlockChain().CurrentSafeBlock().Time + 2,
			depositTxAmount:                      big.NewInt(1000000000000000000),
			expectedReturnCode:                   0,
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

			res, err := serviceV1Alpha1.ExecuteOptimisticBlock(context.Background(), baseBlockReq)
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
				commitmentStateAfterExecuteBlock, err := serviceV1Alpha1.GetCommitmentState(context.Background(), &astriaPb.GetCommitmentStateRequest{})
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

func TestExecutionServiceServerV1Alpha2_StreamExecuteOptimisticBlock(t *testing.T) {
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

	requestStreams := []*optimsticPb.StreamExecuteOptimisticBlockRequest{}
	sequencerBlockHash := []byte("sequencer_block_hash")

	// create 1 stream item with 5 txs
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

	errs := ethservice.TxPool().Add(txs, true, false)
	for _, err := range errs {
		require.Nil(t, err, "Failed to add tx to mempool")
	}

	pendingTxs := ethservice.TxPool().Pending(txpool.PendingFilter{OnlyPlainTxs: true})
	require.Len(t, pendingTxs, 1, "Mempool should have 1 tx")
	addrTxs := pendingTxs[testAddr]
	require.Len(t, addrTxs, 5, "Mempool should have 5 txs for test address")

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
		errorCh <- serviceV1Alpha1.StreamExecuteOptimisticBlock(mockStream)
	}(errorCh)

	select {
	// stream either errors out of gets closed
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

	pending := ethservice.TxPool().Pending(txpool.PendingFilter{
		OnlyPlainTxs: true,
	})
	require.Len(t, pending, 0, "Mempool should be empty")
}

// Check that invalid transactions are not added into a block and are removed from the mempool
func TestExecutionServiceServerV1Alpha2_ExecuteBlockAndUpdateCommitmentWithInvalidTransactions(t *testing.T) {
	ethservice, serviceV1Alpha1 := setupExecutionService(t, 10)

	// call genesis info
	genesisInfo, err := serviceV1Alpha1.GetGenesisInfo(context.Background(), &astriaPb.GetGenesisInfoRequest{})
	require.Nil(t, err, "GetGenesisInfo failed")
	require.NotNil(t, genesisInfo, "GenesisInfo is nil")

	// call get commitment state
	commitmentState, err := serviceV1Alpha1.GetCommitmentState(context.Background(), &astriaPb.GetCommitmentStateRequest{})
	require.Nil(t, err, "GetCommitmentState failed")
	require.NotNil(t, commitmentState, "CommitmentState is nil")

	previousBlockHeader := ethservice.BlockChain().CurrentBlock()
	previousBlock := ethservice.BlockChain().GetBlockByHash(previousBlockHeader.Hash())

	ethservice.BlockChain().SetOptimistic(previousBlock)
	ethservice.BlockChain().SetSafe(previousBlockHeader)

	require.NotNil(t, previousBlock, "Previous block not found")

	stateDb, err := ethservice.BlockChain().StateAt(previousBlock.Root())
	require.Nil(t, err, "Failed to get state db")

	latestNonce := stateDb.GetNonce(testAddr)

	// create 5 txs
	txs := []*types.Transaction{}
	marshalledTxs := []*sequencerblockv1.RollupData{}
	for i := 0; i < 5; i++ {
		unsignedTx := types.NewTransaction(latestNonce+uint64(i), testToAddress, big.NewInt(1), params.TxGas, big.NewInt(params.InitialBaseFee*2), nil)
		tx, err := types.SignTx(unsignedTx, types.LatestSigner(ethservice.BlockChain().Config()), testKey)
		require.Nil(t, err, "Failed to sign tx")
		txs = append(txs, tx)

		marshalledTx, err := tx.MarshalBinary()
		require.Nil(t, err, "Failed to marshal tx")
		marshalledTxs = append(marshalledTxs, &sequencerblockv1.RollupData{
			Value: &sequencerblockv1.RollupData_SequencedData{SequencedData: marshalledTx},
		})
	}

	// add a tx with lesser gas than the base gas
	unsignedTx := types.NewTransaction(latestNonce+uint64(5), testToAddress, big.NewInt(1), ethservice.BlockChain().GasLimit(), big.NewInt(params.InitialBaseFee*2), nil)
	tx, err := types.SignTx(unsignedTx, types.LatestSigner(ethservice.BlockChain().Config()), testKey)
	require.Nil(t, err, "Failed to sign tx")
	txs = append(txs, tx)

	marshalledTx, err := tx.MarshalBinary()
	require.Nil(t, err, "Failed to marshal tx")
	marshalledTxs = append(marshalledTxs, &sequencerblockv1.RollupData{
		Value: &sequencerblockv1.RollupData_SequencedData{SequencedData: marshalledTx},
	})

	executeBlockReq := &astriaPb.ExecuteBlockRequest{
		PrevBlockHash: previousBlock.Hash().Bytes(),
		Timestamp: &timestamppb.Timestamp{
			Seconds: int64(previousBlock.Time() + 2),
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
			BaseCelestiaHeight: commitmentState.BaseCelestiaHeight + 1,
		},
	}

	updateCommitmentStateRes, err := serviceV1Alpha1.UpdateCommitmentState(context.Background(), updateCommitmentStateReq)
	require.Nil(t, err, "UpdateCommitmentState failed")
	require.NotNil(t, updateCommitmentStateRes, "UpdateCommitmentState response should not be nil")
	require.Equal(t, updateCommitmentStateRes, updateCommitmentStateReq.CommitmentState, "CommitmentState response should match request")

	// get the soft and firm block
	softBlock := ethservice.BlockChain().CurrentSafeBlock()
	require.NotNil(t, softBlock, "SoftBlock is nil")
	firmBlock := ethservice.BlockChain().CurrentFinalBlock()
	require.NotNil(t, firmBlock, "FirmBlock is nil")

	block := ethservice.BlockChain().GetBlockByNumber(softBlock.Number.Uint64())
	require.NotNil(t, block, "Soft Block not found")
	require.Equal(t, block.Transactions().Len(), 5, "Soft Block should have 5 txs")

	// check if the soft and firm block are set correctly
	require.True(t, bytes.Equal(softBlock.Hash().Bytes(), updateCommitmentStateRes.Soft.Hash), "Soft Block Hashes do not match")
	require.True(t, bytes.Equal(softBlock.ParentHash.Bytes(), updateCommitmentStateRes.Soft.ParentBlockHash), "Soft Block Parent Hash do not match")
	require.Equal(t, softBlock.Number.Uint64(), uint64(updateCommitmentStateRes.Soft.Number), "Soft Block Number do not match")

	require.True(t, bytes.Equal(firmBlock.Hash().Bytes(), updateCommitmentStateRes.Firm.Hash), "Firm Block Hashes do not match")
	require.True(t, bytes.Equal(firmBlock.ParentHash.Bytes(), updateCommitmentStateRes.Firm.ParentBlockHash), "Firm Block Parent Hash do not match")
	require.Equal(t, firmBlock.Number.Uint64(), uint64(updateCommitmentStateRes.Firm.Number), "Firm Block Number do not match")

	celestiaBaseHeight := ethservice.BlockChain().CurrentBaseCelestiaHeight()
	require.Equal(t, celestiaBaseHeight, updateCommitmentStateRes.BaseCelestiaHeight, "BaseCelestiaHeight should be updated in db")
}
