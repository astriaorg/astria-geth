package execution

import (
	"math/big"
	"testing"

	astriaPb "buf.build/gen/go/astria/execution-apis/protocolbuffers/go/astria/execution/v2"
	primitivev1 "buf.build/gen/go/astria/primitives/protocolbuffers/go/astria/primitive/v1"
	sequencerblockv1 "buf.build/gen/go/astria/sequencerblock-apis/protocolbuffers/go/astria/sequencerblock/v1"
	"github.com/btcsuite/btcd/btcutil/bech32"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func testBlobTx() *types.Transaction {
	return types.NewTx(&types.BlobTx{
		Nonce: 1,
		To:    testAddr,
		Value: uint256.NewInt(1000),
		Gas:   1000,
		Data:  []byte("data"),
	})
}

func testDepositTx() *types.Transaction {
	return types.NewTx(&types.DepositTx{
		From:  testAddr,
		Value: big.NewInt(1000),
		Gas:   1000,
	})
}

func generateBech32MAddress() string {
	addressKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	bridgeAddress := crypto.PubkeyToAddress(addressKey.PublicKey)
	bridgeAddressBytes, err := bech32.ConvertBits(bridgeAddress.Bytes(), 8, 5, false)
	if err != nil {
		panic(err)
	}

	bech32m, err := bech32.EncodeM("astria", bridgeAddressBytes)
	if err != nil {
		panic(err)
	}

	return bech32m
}

func TestSequenceTxValidation(t *testing.T) {
	ethservice, _ := setupExecutionService(t, 10)

	blobTx, err := testBlobTx().MarshalBinary()
	require.Nil(t, err, "failed to marshal random blob tx: %v", err)

	depositTx, err := testDepositTx().MarshalBinary()
	require.Nil(t, err, "failed to marshal random deposit tx: %v", err)

	unsignedTx := types.NewTransaction(uint64(0), common.HexToAddress("0x9a9070028361F7AAbeB3f2F2Dc07F82C4a98A02a"), big.NewInt(1), params.TxGas, big.NewInt(params.InitialBaseFee*2), nil)
	tx, err := types.SignTx(unsignedTx, types.LatestSigner(ethservice.BlockChain().Config()), testKey)
	require.Nil(t, err, "failed to sign tx: %v", err)

	validMarshalledTx, err := tx.MarshalBinary()
	require.Nil(t, err, "failed to marshal valid tx: %v", err)

	chainDestinationKey, err := crypto.GenerateKey()
	require.Nil(t, err, "failed to generate chain destination key: %v", err)
	chainDestinationAddress := crypto.PubkeyToAddress(chainDestinationKey.PublicKey)

	fork := ethservice.BlockChain().Config().AstriaForks.GetForkAtHeight(1)

	var bridgeAssetDenom string
	var bridgeAddress string
	for _, bridgeCfg := range fork.BridgeAddresses {
		bridgeAssetDenom = bridgeCfg.AssetDenom
		bridgeAddress = bridgeCfg.BridgeAddress
		break
	}
	require.NotEmpty(t, bridgeAssetDenom, "bridgeAssetDenom not found")

	invalidBridgeAssetDenom := "invalid-asset-denom"

	invalidHeightBridgeAssetDenom := "invalid-height-asset-denom"
	invalidHeightBridgeAddressBech32m := generateBech32MAddress()
	fork.BridgeAddresses[invalidHeightBridgeAddressBech32m] = &params.AstriaBridgeAddressConfig{
		AssetDenom: invalidHeightBridgeAssetDenom,
	}

	tests := []struct {
		description string
		sequencerTx *sequencerblockv1.RollupData
		// just check if error contains the string since error contains other details
		wantErr string
	}{
		{
			description: "unmarshallable sequencer tx",
			sequencerTx: &sequencerblockv1.RollupData{
				Value: &sequencerblockv1.RollupData_SequencedData{
					SequencedData: []byte("unmarshallable tx"),
				},
			},
			wantErr: "failed to unmarshal sequenced data into transaction",
		},
		{
			description: "blob type sequence tx",
			sequencerTx: &sequencerblockv1.RollupData{
				Value: &sequencerblockv1.RollupData_SequencedData{
					SequencedData: blobTx,
				},
			},
			wantErr: "blob tx not allowed in sequenced data",
		},
		{
			description: "deposit type sequence tx",
			sequencerTx: &sequencerblockv1.RollupData{
				Value: &sequencerblockv1.RollupData_SequencedData{
					SequencedData: depositTx,
				},
			},
			wantErr: "deposit tx not allowed in sequenced data",
		},
		{
			description: "deposit tx with an unknown bridge address",
			sequencerTx: &sequencerblockv1.RollupData{Value: &sequencerblockv1.RollupData_Deposit{Deposit: &sequencerblockv1.Deposit{
				BridgeAddress: &primitivev1.Address{
					Bech32M: generateBech32MAddress(),
				},
				Asset:                   bridgeAssetDenom,
				Amount:                  bigIntToProtoU128(big.NewInt(1000000000000000000)),
				RollupId:                &primitivev1.RollupId{Inner: make([]byte, 0)},
				DestinationChainAddress: chainDestinationAddress.String(),
				SourceTransactionId: &primitivev1.TransactionId{
					Inner: "test_tx_hash",
				},
				SourceActionIndex: 0,
			}}},
			wantErr: "unknown bridge address",
		},
		{
			description: "deposit tx with a disallowed asset id",
			sequencerTx: &sequencerblockv1.RollupData{Value: &sequencerblockv1.RollupData_Deposit{Deposit: &sequencerblockv1.Deposit{
				BridgeAddress: &primitivev1.Address{
					Bech32M: bridgeAddress,
				},
				Asset:                   invalidBridgeAssetDenom,
				Amount:                  bigIntToProtoU128(big.NewInt(1000000000000000000)),
				RollupId:                &primitivev1.RollupId{Inner: make([]byte, 0)},
				DestinationChainAddress: chainDestinationAddress.String(),
				SourceTransactionId: &primitivev1.TransactionId{
					Inner: "test_tx_hash",
				},
				SourceActionIndex: 0,
			}}},
			wantErr: "disallowed asset",
		},
		{
			description: "deposit tx with a height and asset below the bridge start height",
			sequencerTx: &sequencerblockv1.RollupData{Value: &sequencerblockv1.RollupData_Deposit{Deposit: &sequencerblockv1.Deposit{
				BridgeAddress: &primitivev1.Address{
					Bech32M: invalidHeightBridgeAddressBech32m,
				},
				Asset:                   invalidHeightBridgeAssetDenom,
				Amount:                  bigIntToProtoU128(big.NewInt(1000000000000000000)),
				RollupId:                &primitivev1.RollupId{Inner: make([]byte, 0)},
				DestinationChainAddress: chainDestinationAddress.String(),
				SourceTransactionId: &primitivev1.TransactionId{
					Inner: "test_tx_hash",
				},
				SourceActionIndex: 0,
			}}},
			wantErr: "not allowed before height",
		},
		{
			description: "valid deposit tx",
			sequencerTx: &sequencerblockv1.RollupData{Value: &sequencerblockv1.RollupData_Deposit{Deposit: &sequencerblockv1.Deposit{
				BridgeAddress: &primitivev1.Address{
					Bech32M: bridgeAddress,
				},
				Asset:                   bridgeAssetDenom,
				Amount:                  bigIntToProtoU128(big.NewInt(1000000000000000000)),
				RollupId:                &primitivev1.RollupId{Inner: make([]byte, 0)},
				DestinationChainAddress: chainDestinationAddress.String(),
				SourceTransactionId: &primitivev1.TransactionId{
					Inner: "test_tx_hash",
				},
				SourceActionIndex: 0,
			}}},
			wantErr: "",
		},
		{
			description: "valid sequencer tx",
			sequencerTx: &sequencerblockv1.RollupData{
				Value: &sequencerblockv1.RollupData_SequencedData{SequencedData: validMarshalledTx},
			},
			wantErr: "",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			_, err := validateAndUnmarshalSequencerTx(test.sequencerTx, &fork)
			if test.wantErr == "" && err == nil {
				return
			}
			require.False(t, test.wantErr == "" && err != nil, "expected error, got nil")
			require.Contains(t, err.Error(), test.wantErr)
		})
	}
}

func TestValidateStaticCommitmentState(t *testing.T) {
	// Valid CommitmentState
	validState := &astriaPb.CommitmentState{
		SoftExecutedBlockMetadata: &astriaPb.ExecutedBlockMetadata{
			Number:     10,
			Hash:       "0x123456",
			ParentHash: "0x654321",
			Timestamp:  &timestamppb.Timestamp{Seconds: 1234567890},
		},
		FirmExecutedBlockMetadata: &astriaPb.ExecutedBlockMetadata{
			Number:     9,
			Hash:       "0xabcdef",
			ParentHash: "0xfedcba",
			Timestamp:  &timestamppb.Timestamp{Seconds: 1234567880},
		},
		LowestCelestiaSearchHeight: 100,
	}

	err := validateStaticCommitmentState(validState)
	require.Nil(t, err, "Valid CommitmentState should pass validation")

	// Test missing SoftExecutedBlockMetadata
	invalidState1 := &astriaPb.CommitmentState{
		SoftExecutedBlockMetadata:  nil,
		FirmExecutedBlockMetadata:  validState.FirmExecutedBlockMetadata,
		LowestCelestiaSearchHeight: 100,
	}
	err = validateStaticCommitmentState(invalidState1)
	require.NotNil(t, err, "CommitmentState without SoftExecutedBlockMetadata should fail validation")
	require.Contains(t, err.Error(), "SoftExecutedBlockMetadata cannot be nil", "Error should mention SoftExecutedBlockMetadata")

	// Test missing FirmExecutedBlockMetadata
	invalidState2 := &astriaPb.CommitmentState{
		SoftExecutedBlockMetadata:  validState.SoftExecutedBlockMetadata,
		FirmExecutedBlockMetadata:  nil,
		LowestCelestiaSearchHeight: 100,
	}
	err = validateStaticCommitmentState(invalidState2)
	require.NotNil(t, err, "CommitmentState without FirmExecutedBlockMetadata should fail validation")
	require.Contains(t, err.Error(), "FirmExecutedBlockMetadata cannot be nil", "Error should mention FirmExecutedBlockMetadata")

	// Test invalid SoftExecutedBlockMetadata
	invalidState3 := &astriaPb.CommitmentState{
		SoftExecutedBlockMetadata: &astriaPb.ExecutedBlockMetadata{
			Number:     0, // Invalid number
			Hash:       "0x123456",
			ParentHash: "0x654321",
			Timestamp:  &timestamppb.Timestamp{Seconds: 1234567890},
		},
		FirmExecutedBlockMetadata:  validState.FirmExecutedBlockMetadata,
		LowestCelestiaSearchHeight: 100,
	}
	err = validateStaticCommitmentState(invalidState3)
	require.NotNil(t, err, "CommitmentState with invalid SoftExecutedBlockMetadata should fail validation")

	// Test firm block newer than soft block
	invalidState4 := &astriaPb.CommitmentState{
		SoftExecutedBlockMetadata: validState.SoftExecutedBlockMetadata,
		FirmExecutedBlockMetadata: &astriaPb.ExecutedBlockMetadata{
			Number:     11, // Higher than soft block
			Hash:       "0xabcdef",
			ParentHash: "0xfedcba",
			Timestamp:  &timestamppb.Timestamp{Seconds: 1234567880},
		},
		LowestCelestiaSearchHeight: 100,
	}
	err = validateStaticCommitmentState(invalidState4)
	require.NotNil(t, err, "CommitmentState with firm block newer than soft block should fail validation")
	require.Contains(t, err.Error(), "FirmExecutedBlockMetadata number", "Error should mention FirmExecutedBlockMetadata number")
}

func TestValidateStaticExecutedBlockMetadata(t *testing.T) {
	// Valid ExecutedBlockMetadata
	validMetadata := &astriaPb.ExecutedBlockMetadata{
		Number:     10,
		Hash:       "0x123456",
		ParentHash: "0x654321",
		Timestamp:  &timestamppb.Timestamp{Seconds: 1234567890},
	}

	err := validateStaticExecutedBlockMetadata(validMetadata)
	require.Nil(t, err, "Valid ExecutedBlockMetadata should pass validation")

	// Test block number 0
	invalidMetadata1 := &astriaPb.ExecutedBlockMetadata{
		Number:     0,
		Hash:       "0x123456",
		ParentHash: "0x654321",
		Timestamp:  &timestamppb.Timestamp{Seconds: 1234567890},
	}
	err = validateStaticExecutedBlockMetadata(invalidMetadata1)
	require.NotNil(t, err, "ExecutedBlockMetadata with block number 0 should fail validation")
	require.Contains(t, err.Error(), "block number cannot be 0", "Error should mention block number")

	// Test empty hash
	invalidMetadata2 := &astriaPb.ExecutedBlockMetadata{
		Number:     10,
		Hash:       "",
		ParentHash: "0x654321",
		Timestamp:  &timestamppb.Timestamp{Seconds: 1234567890},
	}
	err = validateStaticExecutedBlockMetadata(invalidMetadata2)
	require.NotNil(t, err, "ExecutedBlockMetadata with empty hash should fail validation")
	require.Contains(t, err.Error(), "block hash cannot be empty", "Error should mention block hash")

	// Test empty parent hash
	invalidMetadata3 := &astriaPb.ExecutedBlockMetadata{
		Number:     10,
		Hash:       "0x123456",
		ParentHash: "",
		Timestamp:  &timestamppb.Timestamp{Seconds: 1234567890},
	}
	err = validateStaticExecutedBlockMetadata(invalidMetadata3)
	require.NotNil(t, err, "ExecutedBlockMetadata with empty parent hash should fail validation")
	require.Contains(t, err.Error(), "parent hash cannot be empty", "Error should mention parent hash")

	// Test missing timestamp
	invalidMetadata4 := &astriaPb.ExecutedBlockMetadata{
		Number:     10,
		Hash:       "0x123456",
		ParentHash: "0x654321",
		Timestamp:  nil,
	}
	err = validateStaticExecutedBlockMetadata(invalidMetadata4)
	require.NotNil(t, err, "ExecutedBlockMetadata with nil timestamp should fail validation")
	require.Contains(t, err.Error(), "timestamp cannot be nil", "Error should mention timestamp")
}

func TestValidateStaticExecuteBlockRequest(t *testing.T) {
	// Valid ExecuteBlockRequest
	validRequest := &astriaPb.ExecuteBlockRequest{
		SessionId:  "valid-session-id",
		ParentHash: "0x123456",
		Timestamp:  &timestamppb.Timestamp{Seconds: 1234567890},
		Transactions: []*sequencerblockv1.RollupData{
			{
				Value: &sequencerblockv1.RollupData_SequencedData{
					SequencedData: []byte("valid-data"),
				},
			},
		},
	}

	err := validateStaticExecuteBlockRequest(validRequest)
	require.Nil(t, err, "Valid ExecuteBlockRequest should pass validation")

	// Test empty session ID
	invalidRequest1 := &astriaPb.ExecuteBlockRequest{
		SessionId:    "",
		ParentHash:   "0x123456",
		Timestamp:    &timestamppb.Timestamp{Seconds: 1234567890},
		Transactions: validRequest.Transactions,
	}
	err = validateStaticExecuteBlockRequest(invalidRequest1)
	require.NotNil(t, err, "ExecuteBlockRequest with empty session ID should fail validation")
	require.Contains(t, err.Error(), "session_id cannot be empty", "Error should mention session_id")

	// Test empty parent hash
	invalidRequest2 := &astriaPb.ExecuteBlockRequest{
		SessionId:    "valid-session-id",
		ParentHash:   "",
		Timestamp:    &timestamppb.Timestamp{Seconds: 1234567890},
		Transactions: validRequest.Transactions,
	}
	err = validateStaticExecuteBlockRequest(invalidRequest2)
	require.NotNil(t, err, "ExecuteBlockRequest with empty parent hash should fail validation")
	require.Contains(t, err.Error(), "parent_hash cannot be empty", "Error should mention parent_hash")

	// Test nil timestamp
	invalidRequest3 := &astriaPb.ExecuteBlockRequest{
		SessionId:    "valid-session-id",
		ParentHash:   "0x123456",
		Timestamp:    nil,
		Transactions: validRequest.Transactions,
	}
	err = validateStaticExecuteBlockRequest(invalidRequest3)
	require.NotNil(t, err, "ExecuteBlockRequest with nil timestamp should fail validation")
	require.Contains(t, err.Error(), "timestamp cannot be nil", "Error should mention timestamp")
}
