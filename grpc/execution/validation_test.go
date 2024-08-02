package execution

import (
	composerv1alpha1 "buf.build/gen/go/astria/composer-apis/protocolbuffers/go/astria/composer/v1alpha1"
	"github.com/golang/protobuf/proto"
	"math/big"
	"testing"

	primitivev1 "buf.build/gen/go/astria/primitives/protocolbuffers/go/astria/primitive/v1"
	sequencerblockv1alpha1 "buf.build/gen/go/astria/sequencerblock-apis/protocolbuffers/go/astria/sequencerblock/v1alpha1"
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

func TestExtractBuilderBundleAndTxs(t *testing.T) {
	// we need to consider the following cases
	// 1: empty list of transactions
	// 2. list of transactions without builderBundlePacket
	// 3. list of transactions with builderBundlePacket
	// 4. one transaction with builderBundlePacket

	ethservice, serviceV1Alpha1 := setupExecutionService(t, 10)

	//chainDestinationKey, err := crypto.GenerateKey()
	//require.Nil(t, err, "failed to generate chain destination key: %v", err)
	//chainDestinationAddress := crypto.PubkeyToAddress(chainDestinationKey.PublicKey)
	//
	//bridgeAssetDenom := ethservice.BlockChain().Config().AstriaBridgeAddressConfigs[0].AssetDenom
	//invalidBridgeAssetDenom := "invalid-asset-denom"

	invalidHeightBridgeAssetDenom := "invalid-height-asset-denom"
	invalidHeightBridgeAddressBech32m := generateBech32MAddress()
	serviceV1Alpha1.bridgeAddresses[invalidHeightBridgeAddressBech32m] = &params.AstriaBridgeAddressConfig{
		AssetDenom:  invalidHeightBridgeAssetDenom,
		StartHeight: 100,
	}

	tests := []struct {
		description            string
		noOfTxsInBuilderBundle int
		noOfOtherTxs           int
	}{
		{
			description:            "empty list of transactions",
			noOfTxsInBuilderBundle: 0,
			noOfOtherTxs:           0,
		},
		{
			description:            "list of transactions without builderBundlePacket",
			noOfTxsInBuilderBundle: 0,
			noOfOtherTxs:           5,
		},
		{
			description:            "list of transactions with builderBundlePacket",
			noOfTxsInBuilderBundle: 5,
			noOfOtherTxs:           5,
		},
		{
			description:            "one transaction with builderBundlePacket",
			noOfTxsInBuilderBundle: 5,
			noOfOtherTxs:           0,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			marshalledTxs := []*sequencerblockv1alpha1.RollupData{}

			// create the other txs
			nonce := 0
			for i := nonce; i < test.noOfOtherTxs; i++ {
				unsignedTx := types.NewTransaction(uint64(i), testToAddress, big.NewInt(1), params.TxGas, big.NewInt(params.InitialBaseFee*2), nil)
				tx, err := types.SignTx(unsignedTx, types.LatestSigner(ethservice.BlockChain().Config()), testKey)
				require.Nil(t, err, "Failed to sign tx")

				marshalledTx, err := tx.MarshalBinary()
				require.Nil(t, err, "Failed to marshal tx")
				marshalledTxs = append(marshalledTxs, &sequencerblockv1alpha1.RollupData{
					Value: &sequencerblockv1alpha1.RollupData_SequencedData{SequencedData: marshalledTx},
				})
			}

			// create the BuilderBundlePacket
			builderBundle := &composerv1alpha1.BuilderBundlePacket{
				Bundle: &composerv1alpha1.BuilderBundle{
					Transactions: []*sequencerblockv1alpha1.RollupData{},
					ParentHash:   nil,
				},
			}
			if test.noOfTxsInBuilderBundle > 0 {
				// create noOfTxsInBuilderBundle txs
				for i := nonce; i < nonce+test.noOfTxsInBuilderBundle; i++ {
					unsignedTx := types.NewTransaction(uint64(i), testToAddress, big.NewInt(1), params.TxGas, big.NewInt(params.InitialBaseFee*2), nil)
					tx, err := types.SignTx(unsignedTx, types.LatestSigner(ethservice.BlockChain().Config()), testKey)
					require.Nil(t, err, "Failed to sign tx")

					marshalledTx, err := tx.MarshalBinary()
					require.Nil(t, err, "Failed to marshal tx")
					builderBundle.Bundle.Transactions = append(builderBundle.Bundle.Transactions, &sequencerblockv1alpha1.RollupData{
						Value: &sequencerblockv1alpha1.RollupData_SequencedData{SequencedData: marshalledTx},
					})
				}

				// add the builderBundle to the list of transactions
				marshalledProto, err := proto.Marshal(builderBundle)
				require.Nil(t, err, "Failed to marshal builder bundle")
				marshalledTxs = append(marshalledTxs, &sequencerblockv1alpha1.RollupData{
					Value: &sequencerblockv1alpha1.RollupData_SequencedData{SequencedData: marshalledProto},
				})
			}

			builderBundlePacket, otherTxs, err := extractBuilderBundleAndTxs(marshalledTxs, 2, serviceV1Alpha1.bridgeAddresses, serviceV1Alpha1.bridgeAllowedAssets, common.Address{})
			require.Nil(t, err, "Failed to extract builder bundle and txs")
			require.Equal(t, test.noOfTxsInBuilderBundle, len(builderBundlePacket.GetBundle().GetTransactions()), "Incorrect number of txs in builder bundle")
			require.Equal(t, test.noOfOtherTxs, len(otherTxs), "Incorrect number of other txs")
		})
	}

}

func TestSequenceTxValidation(t *testing.T) {
	ethservice, serviceV1Alpha1 := setupExecutionService(t, 10)

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
					Bech32M: generateBech32MAddress(),
				},
				Asset:                   bridgeAssetDenom,
				Amount:                  bigIntToProtoU128(big.NewInt(1000000000000000000)),
				RollupId:                &primitivev1.RollupId{Inner: make([]byte, 0)},
				DestinationChainAddress: chainDestinationAddress.String(),
			}}},
			wantErr: "unknown bridge address",
		},
		{
			description: "deposit tx with a disallowed asset id",
			sequencerTx: &sequencerblockv1alpha1.RollupData{Value: &sequencerblockv1alpha1.RollupData_Deposit{Deposit: &sequencerblockv1alpha1.Deposit{
				BridgeAddress: &primitivev1.Address{
					Bech32M: bridgeAddress,
				},
				Asset:                   invalidBridgeAssetDenom,
				Amount:                  bigIntToProtoU128(big.NewInt(1000000000000000000)),
				RollupId:                &primitivev1.RollupId{Inner: make([]byte, 0)},
				DestinationChainAddress: chainDestinationAddress.String(),
			}}},
			wantErr: "disallowed asset",
		},
		{
			description: "deposit tx with a height and asset below the bridge start height",
			sequencerTx: &sequencerblockv1alpha1.RollupData{Value: &sequencerblockv1alpha1.RollupData_Deposit{Deposit: &sequencerblockv1alpha1.Deposit{
				BridgeAddress: &primitivev1.Address{
					Bech32M: invalidHeightBridgeAddressBech32m,
				},
				Asset:                   invalidHeightBridgeAssetDenom,
				Amount:                  bigIntToProtoU128(big.NewInt(1000000000000000000)),
				RollupId:                &primitivev1.RollupId{Inner: make([]byte, 0)},
				DestinationChainAddress: chainDestinationAddress.String(),
			}}},
			wantErr: "not allowed before height",
		},
		{
			description: "valid deposit tx",
			sequencerTx: &sequencerblockv1alpha1.RollupData{Value: &sequencerblockv1alpha1.RollupData_Deposit{Deposit: &sequencerblockv1alpha1.Deposit{
				BridgeAddress: &primitivev1.Address{
					Bech32M: bridgeAddress,
				},
				Asset:                   bridgeAssetDenom,
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
			_, err := validateAndUnmarshalTx(2, test.sequencerTx, serviceV1Alpha1.bridgeAddresses, serviceV1Alpha1.bridgeAllowedAssets, common.Address{})
			if test.wantErr == "" && err == nil {
				return
			}
			require.False(t, test.wantErr == "" && err != nil, "expected error, got nil")
			require.Contains(t, err.Error(), test.wantErr)
		})
	}
}
