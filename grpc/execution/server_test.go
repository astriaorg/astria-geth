package execution

import (
	astriaGrpc "buf.build/gen/go/astria/execution-apis/grpc/go/astria/execution/v1alpha2/executionv1alpha2grpc"
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
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math/big"
	"testing"
)

func TestExecutionService_GetGenesisInfo(t *testing.T) {
	n, ethservice, serviceV1Alpha1 := setupExecutionService(t, 10)

	conn, err := grpc.Dial(GrpcEndpointWithoutPrefix(n), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial gRPC: %v", err)
	}

	client := astriaGrpc.NewExecutionServiceClient(conn)

	genesisInfo, err := client.GetGenesisInfo(context.Background(), &astriaPb.GetGenesisInfoRequest{})
	if err != nil {
		t.Fatalf("GetGenesisInfo failed: %v", err)
	}

	hashedRollupId := sha256.Sum256([]byte(ethservice.BlockChain().Config().AstriaRollupName))

	if bytes.Compare(genesisInfo.RollupId, hashedRollupId[:]) != 0 {
		t.Fatalf("RollupId is not correct")
	}
	if genesisInfo.GetSequencerGenesisBlockHeight() != ethservice.BlockChain().Config().AstriaSequencerInitialHeight {
		t.Fatalf("SequencerInitialHeight is not correct")
	}
	if genesisInfo.GetCelestiaBaseBlockHeight() != ethservice.BlockChain().Config().AstriaCelestiaInitialHeight {
		t.Fatalf("CelestiaInitialHeight is not correct")
	}
	if genesisInfo.GetCelestiaBlockVariance() != ethservice.BlockChain().Config().AstriaCelestiaHeightVariance {
		t.Fatalf("CelestiaHeightVariance is not correct")
	}

	if serviceV1Alpha1.genesisInfoCalled != true {
		t.Fatalf("GetGenesisInfo should be called")
	}

	res := ethservice.BlockChain().GetBlockByNumber(3)
	if res == nil {
		t.Fatalf("Block not found")
	}
}

func TestExecutionService_GetBlockByBlockNumber(t *testing.T) {
	n, ethservice, _ := setupExecutionService(t, 10)

	conn, err := grpc.Dial(GrpcEndpointWithoutPrefix(n), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial gRPC: %v", err)
	}

	client := astriaGrpc.NewExecutionServiceClient(conn)

	block := ethservice.BlockChain().GetBlockByNumber(4)
	if block == nil {
		t.Fatalf("Block not found")
	}

	blockInfo, err := client.GetBlock(context.Background(), &astriaPb.GetBlockRequest{
		Identifier: &astriaPb.BlockIdentifier{Identifier: &astriaPb.BlockIdentifier_BlockNumber{BlockNumber: 4}},
	})
	if err != nil {
		t.Fatalf("GetGenesisInfo failed: %v", err)
	}

	if blockInfo == nil {
		t.Fatalf("Block not found")
	}

	if uint64(blockInfo.Number) != block.NumberU64() {
		t.Fatalf("Block number is not correct")
	}
	if bytes.Compare(blockInfo.ParentBlockHash, block.ParentHash().Bytes()) != 0 {
		t.Fatalf("ParentBlockHash is not correct")
	}
	if bytes.Compare(blockInfo.Hash, block.Hash().Bytes()) != 0 {
		t.Fatalf("BlockHash is not correct")
	}
}

func TestExecutionService_GetBlockByBlockHash(t *testing.T) {
	n, ethservice, _ := setupExecutionService(t, 10)

	conn, err := grpc.Dial(GrpcEndpointWithoutPrefix(n), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial gRPC: %v", err)
	}

	client := astriaGrpc.NewExecutionServiceClient(conn)

	block := ethservice.BlockChain().GetBlockByNumber(4)
	if block == nil {
		t.Fatalf("Block not found")
	}

	blockInfo, err := client.GetBlock(context.Background(), &astriaPb.GetBlockRequest{
		Identifier: &astriaPb.BlockIdentifier{Identifier: &astriaPb.BlockIdentifier_BlockHash{BlockHash: block.Hash().Bytes()}},
	})
	if err != nil {
		t.Fatalf("GetGenesisInfo failed: %v", err)
	}

	if blockInfo == nil {
		t.Fatalf("Block not found")
	}

	if uint64(blockInfo.Number) != block.NumberU64() {
		t.Fatalf("Block number is not correct")
	}
	if bytes.Compare(blockInfo.ParentBlockHash, block.ParentHash().Bytes()) != 0 {
		t.Fatalf("ParentBlockHash is not correct")
	}
	if bytes.Compare(blockInfo.Hash, block.Hash().Bytes()) != 0 {
		t.Fatalf("BlockHash is not correct")
	}
}

func TestExecutionServiceServerV1Alpha2_BatchGetBlocksByBlockNumber(t *testing.T) {
	n, ethservice, _ := setupExecutionService(t, 10)

	conn, err := grpc.Dial(GrpcEndpointWithoutPrefix(n), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial gRPC: %v", err)
	}

	client := astriaGrpc.NewExecutionServiceClient(conn)

	batchGetBlocksRequest := &astriaPb.BatchGetBlocksRequest{
		Identifiers: []*astriaPb.BlockIdentifier{
			{Identifier: &astriaPb.BlockIdentifier_BlockNumber{BlockNumber: 1}},
			{Identifier: &astriaPb.BlockIdentifier_BlockNumber{BlockNumber: 2}},
			{Identifier: &astriaPb.BlockIdentifier_BlockNumber{BlockNumber: 3}},
			{Identifier: &astriaPb.BlockIdentifier_BlockNumber{BlockNumber: 4}},
			{Identifier: &astriaPb.BlockIdentifier_BlockNumber{BlockNumber: 5}},
		},
	}

	batchBlocksRes, err := client.BatchGetBlocks(context.Background(), batchGetBlocksRequest)
	if err != nil {
		t.Fatalf("BatchGetBlocks failed: %v", err)
	}

	for _, batchBlock := range batchBlocksRes.GetBlocks() {
		if batchBlock == nil {
			t.Fatalf("Block not found in batch blocks response")
		}

		block := ethservice.BlockChain().GetBlockByNumber(uint64(batchBlock.Number))
		if block == nil {
			t.Fatalf("Block not found in blockchain")
		}

		if uint64(batchBlock.Number) != block.NumberU64() {
			t.Fatalf("Block number is not correct")
		}
		if bytes.Compare(batchBlock.ParentBlockHash, block.ParentHash().Bytes()) != 0 {
			t.Fatalf("ParentBlockHash is not correct")
		}
		if bytes.Compare(batchBlock.Hash, block.Hash().Bytes()) != 0 {
			t.Fatalf("BlockHash is not correct")
		}
	}
}

func TestExecutionServiceServerV1Alpha2_BatchGetBlocksByBlockHash(t *testing.T) {
	n, ethservice, _ := setupExecutionService(t, 10)

	conn, err := grpc.Dial(GrpcEndpointWithoutPrefix(n), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial gRPC: %v", err)
	}

	client := astriaGrpc.NewExecutionServiceClient(conn)

	batchGetBlocksRequest := &astriaPb.BatchGetBlocksRequest{
		Identifiers: []*astriaPb.BlockIdentifier{},
	}

	for i := 1; i <= 5; i++ {
		block := ethservice.BlockChain().GetBlockByNumber(uint64(i))
		if block == nil {
			t.Fatalf("Block not found in blockchain")
		}
		batchGetBlocksRequest.Identifiers = append(batchGetBlocksRequest.Identifiers, &astriaPb.BlockIdentifier{
			Identifier: &astriaPb.BlockIdentifier_BlockHash{BlockHash: block.Hash().Bytes()},
		})
	}

	batchBlocksRes, err := client.BatchGetBlocks(context.Background(), batchGetBlocksRequest)
	if err != nil {
		t.Fatalf("BatchGetBlocks failed: %v", err)
	}

	for _, batchBlock := range batchBlocksRes.GetBlocks() {
		if batchBlock == nil {
			t.Fatalf("Block not found in batch blocks response")
		}

		block := ethservice.BlockChain().GetBlockByNumber(uint64(batchBlock.Number))
		if block == nil {
			t.Fatalf("Block not found in blockchain")
		}

		if uint64(batchBlock.Number) != block.NumberU64() {
			t.Fatalf("Block number is not correct")
		}
		if bytes.Compare(batchBlock.ParentBlockHash, block.ParentHash().Bytes()) != 0 {
			t.Fatalf("ParentBlockHash is not correct")
		}
		if bytes.Compare(batchBlock.Hash, block.Hash().Bytes()) != 0 {
			t.Fatalf("BlockHash is not correct")
		}
	}
}

func TestExecutionServiceServerV1Alpha2_GetCommitmentState(t *testing.T) {
	n, ethservice, serviceV1Alpha1 := setupExecutionService(t, 10)

	conn, err := grpc.Dial(GrpcEndpointWithoutPrefix(n), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial gRPC: %v", err)
	}

	client := astriaGrpc.NewExecutionServiceClient(conn)

	commitmentState, err := client.GetCommitmentState(context.Background(), &astriaPb.GetCommitmentStateRequest{})
	if err != nil {
		t.Fatalf("GetCommitmentState failed: %v", err)
	}

	if commitmentState == nil {
		t.Fatalf("CommitmentState is nil")
	}

	softBlock := ethservice.BlockChain().CurrentSafeBlock()
	if softBlock == nil {
		t.Fatalf("SoftBlock is nil")
	}
	firmBlock := ethservice.BlockChain().CurrentFinalBlock()
	if firmBlock == nil {
		t.Fatalf("FirmBlock is nil")
	}

	if bytes.Compare(commitmentState.Soft.Hash, softBlock.Hash().Bytes()) != 0 {
		t.Fatalf("Soft Block Hashes do not match")
	}
	if bytes.Compare(commitmentState.Soft.ParentBlockHash, softBlock.ParentHash.Bytes()) != 0 {
		t.Fatalf("Soft Block Parent Hash do not match")
	}
	if uint64(commitmentState.Soft.Number) != softBlock.Number.Uint64() {
		t.Fatalf("Soft Block Number do not match")
	}

	if bytes.Compare(commitmentState.Firm.Hash, firmBlock.Hash().Bytes()) != 0 {
		t.Fatalf("Firm Block Hashes do not match")
	}
	if bytes.Compare(commitmentState.Firm.ParentBlockHash, firmBlock.ParentHash.Bytes()) != 0 {
		t.Fatalf("Firm Block Parent Hash do not match")
	}
	if uint64(commitmentState.Firm.Number) != firmBlock.Number.Uint64() {
		t.Fatalf("Firm Block Number do not match")
	}

	if serviceV1Alpha1.getCommitmentStateCalled != true {
		t.Fatalf("GetCommitmentState should be called")
	}
}

func TestExecutionServiceServerV1Alpha2_ExecuteBlockAndUpdateCommitment(t *testing.T) {
	n, ethservice, _ := setupExecutionService(t, 10)
	
	conn, err := grpc.Dial(GrpcEndpointWithoutPrefix(n), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial gRPC: %v", err)
	}

	client := astriaGrpc.NewExecutionServiceClient(conn)

	// call genesis info
	genesisInfo, err := client.GetGenesisInfo(context.Background(), &astriaPb.GetGenesisInfoRequest{})
	if err != nil {
		t.Fatalf("GetGenesisInfo failed: %v", err)
	}
	if genesisInfo == nil {
		t.Fatalf("GenesisInfo is nil")
	}

	// call get commitment state
	commitmentState, err := client.GetCommitmentState(context.Background(), &astriaPb.GetCommitmentStateRequest{})
	if err != nil {
		t.Fatalf("GetCommitmentState failed: %v", err)
	}
	if commitmentState == nil {
		t.Fatalf("CommitmentState is nil")
	}

	// get previous block hash
	previousBlock := ethservice.BlockChain().CurrentSafeBlock()
	if previousBlock == nil {
		t.Fatalf("Previous block not found")
	}

	// create 5 txs
	txs := []*types.Transaction{}
	marshalledTxs := []*sequencerblockv1alpha1.RollupData{}
	for i := 0; i < 5; i++ {
		tx, err := types.SignTx(types.NewTransaction(uint64(i), common.HexToAddress("0x9a9070028361F7AAbeB3f2F2Dc07F82C4a98A02a"), big.NewInt(1), params.TxGas, big.NewInt(params.InitialBaseFee*2), nil), types.LatestSigner(ethservice.BlockChain().Config()), testKey)
		if err != nil {
			t.Fatalf("Failed to sign tx: %v", err)
		}
		txs = append(txs, tx)

		marshalledTx, err := tx.MarshalBinary()
		if err != nil {
			t.Fatalf("Failed to marshal tx: %v", err)
		}
		marshalledTxs = append(marshalledTxs, &sequencerblockv1alpha1.RollupData{
			Value: &sequencerblockv1alpha1.RollupData_SequencedData{SequencedData: marshalledTx},
		})
	}

	executeBlockReq := &astriaPb.ExecuteBlockRequest{
		PrevBlockHash: previousBlock.Hash().Bytes(),
		Timestamp: &timestamppb.Timestamp{
			Seconds: int64(previousBlock.Time + 2),
		},
		Transactions: marshalledTxs,
	}

	executeBlockRes, err := client.ExecuteBlock(context.Background(), executeBlockReq)
	if err != nil {
		t.Fatalf("ExecuteBlock failed: %v", err)
	}

	if executeBlockRes == nil {
		t.Fatalf("ExecuteBlock response is nil")
	}

	// check astria ordered is clear
	astriaOrdered := ethservice.TxPool().AstriaOrdered()
	if astriaOrdered.Len() != 0 {
		t.Fatalf("AstriaOrdered should be empty")
	}

	// call update commitment state to set our block as soft and firm
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

	updateCommitmentStateRes, err := client.UpdateCommitmentState(context.Background(), updateCommitmentStateReq)
	if err != nil {
		t.Fatalf("UpdateCommitmentState failed: %v", err)
	}
	if updateCommitmentStateRes == nil {
		t.Fatalf("UpdateCommitmentState response should not be nil")
	}

	// get the soft and firm block
	softBlock := ethservice.BlockChain().CurrentSafeBlock()
	if softBlock == nil {
		t.Fatalf("SoftBlock is nil")
	}
	firmBlock := ethservice.BlockChain().CurrentFinalBlock()
	if firmBlock == nil {
		t.Fatalf("FirmBlock is nil")
	}

	if bytes.Compare(softBlock.Hash().Bytes(), executeBlockRes.Hash) != 0 {
		t.Fatalf("Soft Block Hashes do not match")
	}
	if bytes.Compare(softBlock.ParentHash.Bytes(), executeBlockRes.ParentBlockHash) != 0 {
		t.Fatalf("Soft Block Parent Hashes do not match")
	}
	if softBlock.Number.Uint64() != uint64(executeBlockRes.Number) {
		t.Fatalf("Soft Block Numbers do not match")
	}

	if bytes.Compare(firmBlock.Hash().Bytes(), executeBlockRes.Hash) != 0 {
		t.Fatalf("Firm Block Hashes do not match")
	}
	if bytes.Compare(firmBlock.ParentHash.Bytes(), executeBlockRes.ParentBlockHash) != 0 {
		t.Fatalf("Firm Block Parent Hashes do not match")
	}
	if firmBlock.Number.Uint64() != uint64(executeBlockRes.Number) {
		t.Fatalf("Firm Block Numbers do not match")
	}

}

func bigIntToProtoU128(i *big.Int) *primitivev1.Uint128 {
	lo := i.Uint64()
	hi := new(big.Int).Rsh(i, 64).Uint64()
	return &primitivev1.Uint128{Lo: lo, Hi: hi}
}

func TestExecutionServiceServerV1Alpha2_ExecuteBlockAndUpdateCommitmentWithDepositTx(t *testing.T) {
	n, ethservice, _ := setupExecutionService(t, 10)

	conn, err := grpc.Dial(GrpcEndpointWithoutPrefix(n), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial gRPC: %v", err)
	}

	client := astriaGrpc.NewExecutionServiceClient(conn)

	// call genesis info
	genesisInfo, err := client.GetGenesisInfo(context.Background(), &astriaPb.GetGenesisInfoRequest{})
	if err != nil {
		t.Fatalf("GetGenesisInfo failed: %v", err)
	}
	if genesisInfo == nil {
		t.Fatalf("GenesisInfo is nil")
	}

	// call get commitment state
	commitmentState, err := client.GetCommitmentState(context.Background(), &astriaPb.GetCommitmentStateRequest{})
	if err != nil {
		t.Fatalf("GetCommitmentState failed: %v", err)
	}
	if commitmentState == nil {
		t.Fatalf("CommitmentState is nil")
	}

	// get previous block hash
	previousBlock := ethservice.BlockChain().CurrentSafeBlock()
	if previousBlock == nil {
		t.Fatalf("Previous block not found")
	}

	// create 5 txs
	txs := []*types.Transaction{}
	marshalledTxs := []*sequencerblockv1alpha1.RollupData{}
	for i := 0; i < 5; i++ {
		tx, err := types.SignTx(types.NewTransaction(uint64(i), common.HexToAddress("0x9a9070028361F7AAbeB3f2F2Dc07F82C4a98A02a"), big.NewInt(1), params.TxGas, big.NewInt(params.InitialBaseFee*2), nil), types.LatestSigner(ethservice.BlockChain().Config()), testKey)
		if err != nil {
			t.Fatalf("Failed to sign tx: %v", err)
		}
		txs = append(txs, tx)

		marshalledTx, err := tx.MarshalBinary()
		if err != nil {
			t.Fatalf("Failed to marshal tx: %v", err)
		}
		marshalledTxs = append(marshalledTxs, &sequencerblockv1alpha1.RollupData{
			Value: &sequencerblockv1alpha1.RollupData_SequencedData{SequencedData: marshalledTx},
		})
	}

	amountToDeposit := big.NewInt(1000000000000000000)
	depositAmount := bigIntToProtoU128(big.NewInt(1000000000000000000))
	bridgeAddress := ethservice.BlockChain().Config().AstriaBridgeAddressConfigs[0].BridgeAddress
	bridgeAssetDenom := sha256.Sum256([]byte(ethservice.BlockChain().Config().AstriaBridgeAddressConfigs[0].AssetDenom))

	// create new chain destination address for better testing
	chainDestinationAddressPrivKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Failed to generate chain destination address: %v", err)
	}

	chainDestinationAddress := crypto.PubkeyToAddress(chainDestinationAddressPrivKey.PublicKey)

	stateDb, err := ethservice.BlockChain().State()
	if err != nil {
		t.Fatalf("Failed to get state db: %v", err)
	}

	chainDestinationAddressBalanceBefore := stateDb.GetBalance(chainDestinationAddress)

	depositTx := &sequencerblockv1alpha1.RollupData{Value: &sequencerblockv1alpha1.RollupData_Deposit{Deposit: &sequencerblockv1alpha1.Deposit{
		BridgeAddress: &primitivev1.Address{
			Inner: bridgeAddress,
		},
		AssetId:                 bridgeAssetDenom[:],
		Amount:                  depositAmount,
		RollupId:                &primitivev1.RollupId{Inner: make([]byte, 0)},
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

	executeBlockRes, err := client.ExecuteBlock(context.Background(), executeBlockReq)
	if err != nil {
		t.Fatalf("ExecuteBlock failed: %v", err)
	}

	if executeBlockRes == nil {
		t.Fatalf("ExecuteBlock response is nil")
	}

	// check astria ordered is clear
	astriaOrdered := ethservice.TxPool().AstriaOrdered()
	if astriaOrdered.Len() != 0 {
		t.Fatalf("AstriaOrdered should be empty")
	}

	// call update commitment state to set our block as soft and firm
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

	updateCommitmentStateRes, err := client.UpdateCommitmentState(context.Background(), updateCommitmentStateReq)
	if err != nil {
		t.Fatalf("UpdateCommitmentState failed: %v", err)
	}
	if updateCommitmentStateRes == nil {
		t.Fatalf("UpdateCommitmentState response should not be nil")
	}

	// get the soft and firm block
	softBlock := ethservice.BlockChain().CurrentSafeBlock()
	if softBlock == nil {
		t.Fatalf("SoftBlock is nil")
	}
	firmBlock := ethservice.BlockChain().CurrentFinalBlock()
	if firmBlock == nil {
		t.Fatalf("FirmBlock is nil")
	}

	if bytes.Compare(softBlock.Hash().Bytes(), executeBlockRes.Hash) != 0 {
		t.Fatalf("SoftBlock Hashes do not match")
	}
	if bytes.Compare(softBlock.ParentHash.Bytes(), executeBlockRes.ParentBlockHash) != 0 {
		t.Fatalf("Soft Block Parent Hashes do not match")
	}
	if softBlock.Number.Uint64() != uint64(executeBlockRes.Number) {
		t.Fatalf("Soft Block Numbers do not match")
	}

	if bytes.Compare(firmBlock.Hash().Bytes(), executeBlockRes.Hash) != 0 {
		t.Fatalf("Firm Block Hashes do not match")
	}
	if bytes.Compare(firmBlock.ParentHash.Bytes(), executeBlockRes.ParentBlockHash) != 0 {
		t.Fatalf("Firm Block Parent Hashes do not match")
	}
	if firmBlock.Number.Uint64() != uint64(executeBlockRes.Number) {
		t.Fatalf("Firm Block Numbers do not match")
	}

	stateDb, err = ethservice.BlockChain().State()
	if err != nil {
		t.Fatalf("Failed to get state db: %v", err)
	}
	chainDestinationAddressBalanceAfter := stateDb.GetBalance(chainDestinationAddress)

	balanceDiff := new(big.Int).Sub(chainDestinationAddressBalanceAfter, chainDestinationAddressBalanceBefore)
	if balanceDiff.Cmp(amountToDeposit) != 0 {
		t.Fatalf("Chain destination address balance is not correct")
	}
}
