package execution

import (
	"context"
	"math/big"
	"testing"

	primitivev1 "buf.build/gen/go/astria/primitives/protocolbuffers/go/astria/primitive/v1"
	sequencerblockv1 "buf.build/gen/go/astria/sequencerblock-apis/protocolbuffers/go/astria/sequencerblock/v1"
	"github.com/btcsuite/btcd/btcutil/bech32"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
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

func testInjectedTx() *types.Transaction {
	return types.NewTx(&types.InjectedTx{
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
	ethservice, serviceV1Alpha1 := setupExecutionService(t, 10)

	blobTx, err := testBlobTx().MarshalBinary()
	require.Nil(t, err, "failed to marshal random blob tx: %v", err)

	injectedTx, err := testInjectedTx().MarshalBinary()
	require.Nil(t, err, "failed to marshal random injected tx: %v", err)

	unsignedTx := types.NewTransaction(uint64(0), common.HexToAddress("0x9a9070028361F7AAbeB3f2F2Dc07F82C4a98A02a"), big.NewInt(1), params.TxGas, big.NewInt(params.InitialBaseFee*2), nil)
	tx, err := types.SignTx(unsignedTx, types.LatestSigner(ethservice.BlockChain().Config()), testKey)
	require.Nil(t, err, "failed to sign tx: %v", err)

	validMarshalledTx, err := tx.MarshalBinary()
	require.Nil(t, err, "failed to marshal valid tx: %v", err)

	chainDestinationKey, err := crypto.GenerateKey()
	require.Nil(t, err, "failed to generate chain destination key: %v", err)
	chainDestinationAddress := crypto.PubkeyToAddress(chainDestinationKey.PublicKey)

	bridgeAssetDenom := ethservice.BlockChain().Config().AstriaBridgeAddressConfigs[0].AssetDenom
	invalidBridgeAssetDenom := "invalid-asset-denom"

	invalidHeightBridgeAssetDenom := "invalid-height-asset-denom"
	invalidHeightBridgeAddressBech32m := generateBech32MAddress()
	serviceV1Alpha1.bridgeAddresses[invalidHeightBridgeAddressBech32m] = &params.AstriaBridgeAddressConfig{
		AssetDenom:  invalidHeightBridgeAssetDenom,
		StartHeight: 100,
	}

	bridgeAddress := ethservice.BlockChain().Config().AstriaBridgeAddressConfigs[0].BridgeAddress

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
			description: "injected type sequence tx",
			sequencerTx: &sequencerblockv1.RollupData{
				Value: &sequencerblockv1.RollupData_SequencedData{
					SequencedData: injectedTx,
				},
			},
			wantErr: "injected tx not allowed in sequenced data",
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
			cfg := &conversionConfig{
				bridgeAddresses:     serviceV1Alpha1.bridgeAddresses,
				bridgeAllowedAssets: serviceV1Alpha1.bridgeAllowedAssets,
			}
			_, err := validateAndConvertSequencerTx(context.Background(), 2, test.sequencerTx, cfg)
			if test.wantErr == "" && err == nil {
				return
			}
			require.False(t, test.wantErr == "" && err != nil, "expected error, got nil")
			require.Contains(t, err.Error(), test.wantErr)
		})
	}
}
