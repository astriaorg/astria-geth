package execution

import (
	primitivev1 "buf.build/gen/go/astria/primitives/protocolbuffers/go/astria/primitive/v1"
	sequencerblockv1alpha1 "buf.build/gen/go/astria/sequencerblock-apis/protocolbuffers/go/astria/sequencerblock/v1alpha1"
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
	"math/big"
	"strings"
	"testing"
)

func randomBlobTx() *types.Transaction {
	return types.NewTx(&types.BlobTx{
		Nonce: 1,
		To:    testAddr,
		Value: uint256.NewInt(1000),
		Gas:   1000,
		Data:  []byte("data"),
	})
}

func randomDepositTx() *types.Transaction {
	return types.NewTx(&types.DepositTx{
		From:  testAddr,
		Value: big.NewInt(1000),
		Gas:   1000,
	})
}

func TestSequenceTxValidation(t *testing.T) {
	_, ethservice, serviceV1Alpha1 := setupExecutionService(t, 10)

	blobTx, err := randomBlobTx().MarshalBinary()
	if err != nil {
		t.Fatalf("failed to marshal random blob tx: %v", err)
	}

	depositTx, err := randomDepositTx().MarshalBinary()
	if err != nil {
		t.Fatalf("failed to marshal random deposit tx: %v", err)
	}

	unsignedTx := types.NewTransaction(uint64(0), common.HexToAddress("0x9a9070028361F7AAbeB3f2F2Dc07F82C4a98A02a"), big.NewInt(1), params.TxGas, big.NewInt(params.InitialBaseFee*2), nil)
	tx, err := types.SignTx(unsignedTx, types.LatestSigner(ethservice.BlockChain().Config()), testKey)
	if err != nil {
		t.Fatalf("Failed to sign tx: %v", err)
	}

	validMarshalledTx, err := tx.MarshalBinary()
	if err != nil {
		t.Fatalf("failed to marshal valid tx: %v", err)
	}

	chainDestinationKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("failed to generate chain destination key: %v", err)
	}
	chainDestinationAddress := crypto.PubkeyToAddress(chainDestinationKey.PublicKey)

	bridgeAssetDenom := sha256.Sum256([]byte(ethservice.BlockChain().Config().AstriaBridgeAddressConfigs[0].AssetDenom))
	invalidBridgeAssetDenom := sha256.Sum256([]byte("invalid-asset-denom"))

	bridgeAddress := ethservice.BlockChain().Config().AstriaBridgeAddressConfigs[0].BridgeAddress

	tests := []struct {
		description string
		sequencerTx *sequencerblockv1alpha1.RollupData
		// just check if error contains the string since error contains other details
		wantErr string
	}{
		{
			description: "unmarshallable sequencer tx",
			sequencerTx: &sequencerblockv1alpha1.RollupData{
				Value: &sequencerblockv1alpha1.RollupData_SequencedData{
					SequencedData: []byte("unmarshallable tx"),
				},
			},
			wantErr: "failed to unmarshal sequenced data into transaction",
		},
		{
			description: "blob type sequence tx",
			sequencerTx: &sequencerblockv1alpha1.RollupData{
				Value: &sequencerblockv1alpha1.RollupData_SequencedData{
					SequencedData: blobTx,
				},
			},
			wantErr: "blob tx not allowed in sequenced data",
		},
		{
			description: "deposit type sequence tx",
			sequencerTx: &sequencerblockv1alpha1.RollupData{
				Value: &sequencerblockv1alpha1.RollupData_SequencedData{
					SequencedData: depositTx,
				},
			},
			wantErr: "deposit tx not allowed in sequenced data",
		},
		{
			description: "deposit tx with an unknown bridge address",
			sequencerTx: &sequencerblockv1alpha1.RollupData{Value: &sequencerblockv1alpha1.RollupData_Deposit{Deposit: &sequencerblockv1alpha1.Deposit{
				BridgeAddress: &primitivev1.Address{
					Inner: []byte("unknown-bridge-address"),
				},
				AssetId:                 bridgeAssetDenom[:],
				Amount:                  bigIntToProtoU128(big.NewInt(1000000000000000000)),
				RollupId:                &primitivev1.RollupId{Inner: make([]byte, 0)},
				DestinationChainAddress: chainDestinationAddress.String(),
			}}},
			wantErr: "unknown bridge address",
		},
		{
			description: "deposit tx with an invalid asset id",
			sequencerTx: &sequencerblockv1alpha1.RollupData{Value: &sequencerblockv1alpha1.RollupData_Deposit{Deposit: &sequencerblockv1alpha1.Deposit{
				BridgeAddress: &primitivev1.Address{
					Inner: bridgeAddress,
				},
				AssetId:                 []byte("invalid-asset-id"),
				Amount:                  bigIntToProtoU128(big.NewInt(1000000000000000000)),
				RollupId:                &primitivev1.RollupId{Inner: make([]byte, 0)},
				DestinationChainAddress: chainDestinationAddress.String(),
			}}},
			wantErr: "invalid asset ID",
		},
		{
			description: "deposit tx with a disallowed asset id",
			sequencerTx: &sequencerblockv1alpha1.RollupData{Value: &sequencerblockv1alpha1.RollupData_Deposit{Deposit: &sequencerblockv1alpha1.Deposit{
				BridgeAddress: &primitivev1.Address{
					Inner: bridgeAddress,
				},
				AssetId:                 invalidBridgeAssetDenom[:],
				Amount:                  bigIntToProtoU128(big.NewInt(1000000000000000000)),
				RollupId:                &primitivev1.RollupId{Inner: make([]byte, 0)},
				DestinationChainAddress: chainDestinationAddress.String(),
			}}},
			wantErr: "disallowed asset ID",
		},
		{
			description: "valid deposit tx",
			sequencerTx: &sequencerblockv1alpha1.RollupData{Value: &sequencerblockv1alpha1.RollupData_Deposit{Deposit: &sequencerblockv1alpha1.Deposit{
				BridgeAddress: &primitivev1.Address{
					Inner: bridgeAddress,
				},
				AssetId:                 bridgeAssetDenom[:],
				Amount:                  bigIntToProtoU128(big.NewInt(1000000000000000000)),
				RollupId:                &primitivev1.RollupId{Inner: make([]byte, 0)},
				DestinationChainAddress: chainDestinationAddress.String(),
			}}},
			wantErr: "",
		},
		{
			description: "valid sequencer tx",
			sequencerTx: &sequencerblockv1alpha1.RollupData{
				Value: &sequencerblockv1alpha1.RollupData_SequencedData{SequencedData: validMarshalledTx},
			},
			wantErr: "",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			_, err := validateAndUnmarshalSequencerTx(test.sequencerTx, serviceV1Alpha1.bridgeAddresses, serviceV1Alpha1.bridgeAllowedAssetIDs)
			if test.wantErr != "" && err == nil {
				t.Errorf("expected error, got nil")
			}
			if test.wantErr == "" && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			// check if wantErr is in err.Error()
			if test.wantErr != "" && !strings.Contains(err.Error(), test.wantErr) {
				t.Errorf("expected error to contain %q, got %q", test.wantErr, err.Error())
			}
		})
	}
}
