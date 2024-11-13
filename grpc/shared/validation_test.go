package shared

import (
	bundlev1alpha1 "buf.build/gen/go/astria/execution-apis/protocolbuffers/go/astria/bundle/v1alpha1"
	"bytes"
	"crypto/ecdsa"
	"crypto/ed25519"
	"github.com/golang/protobuf/proto"
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

func transaction(nonce uint64, gaslimit uint64, key *ecdsa.PrivateKey) *types.Transaction {
	return pricedTransaction(nonce, gaslimit, big.NewInt(1), key)
}

func pricedTransaction(nonce uint64, gaslimit uint64, gasprice *big.Int, key *ecdsa.PrivateKey) *types.Transaction {
	tx, _ := types.SignTx(types.NewTransaction(nonce, common.Address{}, big.NewInt(100), gaslimit, gasprice, nil), types.HomesteadSigner{}, key)
	return tx
}

func bigIntToProtoU128(i *big.Int) *primitivev1.Uint128 {
	lo := i.Uint64()
	hi := new(big.Int).Rsh(i, 64).Uint64()
	return &primitivev1.Uint128{Lo: lo, Hi: hi}
}

func testBlobTx() *types.Transaction {
	return types.NewTx(&types.BlobTx{
		Nonce: 1,
		To:    TestAddr,
		Value: uint256.NewInt(1000),
		Gas:   1000,
		Data:  []byte("data"),
	})
}

func testDepositTx() *types.Transaction {
	return types.NewTx(&types.DepositTx{
		From:  TestAddr,
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

func TestUnmarshallAuctionResultTxs(t *testing.T) {
	tx1 := transaction(0, 1000, TestKey)
	validMarshalledTx1, err := tx1.MarshalBinary()
	require.NoError(t, err, "failed to marshal valid tx: %v", err)

	tx2 := transaction(1, 1000, TestKey)
	validMarshalledTx2, err := tx2.MarshalBinary()
	require.NoError(t, err, "failed to marshal valid tx: %v", err)

	tx3 := transaction(2, 1000, TestKey)
	validMarshalledTx3, err := tx3.MarshalBinary()
	require.NoError(t, err, "failed to marshal valid tx: %v", err)

	pubkey, privkey, err := ed25519.GenerateKey(nil)
	require.NoError(t, err, "failed to generate public and private key")

	validAllocation := &bundlev1alpha1.Bundle{
		Fee:                    100,
		Transactions:           [][]byte{validMarshalledTx1, validMarshalledTx2, validMarshalledTx3},
		BaseSequencerBlockHash: []byte("sequencer block hash"),
		PrevRollupBlockHash:    []byte("prev rollup block hash"),
	}

	marshalledAllocation, err := proto.Marshal(validAllocation)
	require.NoError(t, err, "failed to marshal allocation: %v", err)

	signedAllocation, err := privkey.Sign(nil, marshalledAllocation, &ed25519.Options{
		Hash:    0,
		Context: "",
	})
	require.NoError(t, err, "failed to sign allocation: %v", err)

	tests := []struct {
		description    string
		auctionResult  *bundlev1alpha1.AuctionResult
		prevBlockHash  []byte
		expectedOutput types.Transactions
		// just check if error contains the string since error contains other details
		wantErr string
	}{
		{
			description: "previous block hash mismatch",
			auctionResult: &bundlev1alpha1.AuctionResult{
				// TODO - add signature and public key validation
				Signature: make([]byte, 0),
				PublicKey: make([]byte, 0),
				Allocation: &bundlev1alpha1.Bundle{
					Fee:                    100,
					Transactions:           [][]byte{[]byte("unmarshallable tx")},
					BaseSequencerBlockHash: []byte("sequencer block hash"),
					PrevRollupBlockHash:    []byte("prev rollup block hash"),
				},
			},
			prevBlockHash:  []byte("not prev rollup block hash"),
			expectedOutput: types.Transactions{},
			wantErr:        "prev block hash do not match in allocation",
		},
		{
			description: "invalid signature",
			auctionResult: &bundlev1alpha1.AuctionResult{
				Signature: []byte("invalid signature"),
				PublicKey: pubkey,
				Allocation: &bundlev1alpha1.Bundle{
					Fee:                    100,
					Transactions:           [][]byte{[]byte("unmarshallable tx")},
					BaseSequencerBlockHash: []byte("sequencer block hash"),
					PrevRollupBlockHash:    []byte("prev rollup block hash"),
				},
			},
			prevBlockHash:  []byte("prev rollup block hash"),
			expectedOutput: types.Transactions{},
			wantErr:        "failed to verify signature",
		},
		{
			description: "valid auction result",
			auctionResult: &bundlev1alpha1.AuctionResult{
				Signature:  signedAllocation,
				PublicKey:  pubkey,
				Allocation: validAllocation,
			},
			prevBlockHash:  []byte("prev rollup block hash"),
			expectedOutput: types.Transactions{tx1, tx2, tx3},
			wantErr:        "",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			finalTxs, err := unmarshallAuctionResultTxs(test.auctionResult, test.prevBlockHash)
			if test.wantErr == "" && err == nil {
				for _, tx := range test.expectedOutput {
					foundTx := false
					for _, finalTx := range finalTxs {
						if bytes.Equal(finalTx.Hash().Bytes(), tx.Hash().Bytes()) {
							foundTx = true
						}
					}

					require.True(t, foundTx, "expected tx not found in final txs")
				}
				return
			}
			require.False(t, test.wantErr == "" && err != nil, "expected error, got nil")
			require.Contains(t, err.Error(), test.wantErr)
		})
	}
}

func TestValidateAndUnmarshallDepositTx(t *testing.T) {
	ethservice, serviceV1Alpha1 := SetupSharedService(t, 10)

	chainDestinationKey, err := crypto.GenerateKey()
	require.Nil(t, err, "failed to generate chain destination key: %v", err)
	chainDestinationAddress := crypto.PubkeyToAddress(chainDestinationKey.PublicKey)

	bridgeAssetDenom := ethservice.BlockChain().Config().AstriaBridgeAddressConfigs[0].AssetDenom
	invalidBridgeAssetDenom := "invalid-asset-denom"

	invalidHeightBridgeAssetDenom := "invalid-height-asset-denom"
	invalidHeightBridgeAddressBech32m := generateBech32MAddress()
	serviceV1Alpha1.BridgeAddresses()[invalidHeightBridgeAddressBech32m] = &params.AstriaBridgeAddressConfig{
		AssetDenom:  invalidHeightBridgeAssetDenom,
		StartHeight: 100,
	}

	bridgeAddress := ethservice.BlockChain().Config().AstriaBridgeAddressConfigs[0].BridgeAddress

	tests := []struct {
		description string
		sequencerTx *sequencerblockv1.Deposit
		// just check if error contains the string since error contains other details
		wantErr string
	}{
		{
			description: "deposit tx with an unknown bridge address",
			sequencerTx: &sequencerblockv1.Deposit{
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
			},
			wantErr: "unknown bridge address",
		},
		{
			description: "deposit tx with a disallowed asset id",
			sequencerTx: &sequencerblockv1.Deposit{
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
			},
			wantErr: "disallowed asset",
		},
		{
			description: "deposit tx with a height and asset below the bridge start height",
			sequencerTx: &sequencerblockv1.Deposit{
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
			},
			wantErr: "not allowed before height",
		},
		{
			description: "valid deposit tx",
			sequencerTx: &sequencerblockv1.Deposit{
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
			},
			wantErr: "",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			_, err := validateAndUnmarshalDepositTx(test.sequencerTx, 2, serviceV1Alpha1.BridgeAddresses(), serviceV1Alpha1.BridgeAllowedAssets())
			if test.wantErr == "" && err == nil {
				return
			}
			require.False(t, test.wantErr == "" && err != nil, "expected error, got nil")
			require.Contains(t, err.Error(), test.wantErr)
		})
	}
}

func TestValidateAndUnmarshallSequenceAction(t *testing.T) {
	blobTx, err := testBlobTx().MarshalBinary()
	require.Nil(t, err, "failed to marshal random blob tx: %v", err)

	depositTx, err := testDepositTx().MarshalBinary()
	require.Nil(t, err, "failed to marshal random deposit tx: %v", err)

	tx1 := transaction(0, 1000, TestKey)
	validMarshalledTx, err := tx1.MarshalBinary()
	require.NoError(t, err, "failed to marshal valid tx: %v", err)

	tests := []struct {
		description string
		sequencerTx *sequencerblockv1.RollupData
		// just check if error contains the string since errors can contains other details
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
			description: "valid sequencer tx",
			sequencerTx: &sequencerblockv1.RollupData{
				Value: &sequencerblockv1.RollupData_SequencedData{SequencedData: validMarshalledTx},
			},
			wantErr: "",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			_, err := validateAndUnmarshallSequenceAction(test.sequencerTx)
			if test.wantErr == "" && err == nil {
				return
			}
			require.False(t, test.wantErr == "" && err != nil, "expected error, got nil")
			require.Contains(t, err.Error(), test.wantErr)
		})
	}
}

func TestUnbundleRollupData(t *testing.T) {
	ethservice, serviceV1Alpha1 := SetupSharedService(t, 10)

	baseSequencerBlockHash := []byte("sequencer block hash")
	prevRollupBlockHash := []byte("prev rollup block hash")

	// txs in
	tx1 := transaction(0, 1000, TestKey)
	tx2 := transaction(1, 1000, TestKey)
	tx3 := transaction(2, 1000, TestKey)
	tx4 := transaction(3, 1000, TestKey)
	tx5 := transaction(4, 1000, TestKey)

	validMarshalledTx1, err := tx1.MarshalBinary()
	require.NoError(t, err, "failed to marshal valid tx: %v", err)
	validMarshalledTx2, err := tx2.MarshalBinary()
	require.NoError(t, err, "failed to marshal valid tx: %v", err)
	validMarshalledTx3, err := tx3.MarshalBinary()
	require.NoError(t, err, "failed to marshal valid tx: %v", err)
	validMarshalledTx4, err := tx4.MarshalBinary()
	require.NoError(t, err, "failed to marshal valid tx: %v", err)
	validMarshalledTx5, err := tx5.MarshalBinary()
	require.NoError(t, err, "failed to marshal valid tx: %v", err)

	pubKey, privKey, err := ed25519.GenerateKey(nil)
	require.NoError(t, err, "failed to generate ed25519 key")

	allocation := &bundlev1alpha1.Bundle{
		Fee:                    100,
		Transactions:           [][]byte{validMarshalledTx1, validMarshalledTx2, validMarshalledTx3},
		BaseSequencerBlockHash: baseSequencerBlockHash,
		PrevRollupBlockHash:    prevRollupBlockHash,
	}

	marshalledAllocation, err := proto.Marshal(allocation)
	require.NoError(t, err, "failed to marshal allocation: %v", err)
	signedAllocation, err := privKey.Sign(nil, marshalledAllocation, &ed25519.Options{
		Hash:    0,
		Context: "",
	})
	require.NoError(t, err, "failed to sign allocation: %v", err)

	auctionResult := &bundlev1alpha1.AuctionResult{
		Signature:  signedAllocation,
		PublicKey:  pubKey,
		Allocation: allocation,
	}

	marshalledAuctionResult, err := proto.Marshal(auctionResult)
	require.NoError(t, err, "failed to marshal auction result: %v", err)
	auctionResultSequenceData := &sequencerblockv1.RollupData{
		Value: &sequencerblockv1.RollupData_SequencedData{
			SequencedData: marshalledAuctionResult,
		},
	}
	seqData1 := &sequencerblockv1.RollupData{
		Value: &sequencerblockv1.RollupData_SequencedData{
			SequencedData: validMarshalledTx4,
		},
	}
	seqData2 := &sequencerblockv1.RollupData{
		Value: &sequencerblockv1.RollupData_SequencedData{
			SequencedData: validMarshalledTx5,
		},
	}

	bridgeAddress := ethservice.BlockChain().Config().AstriaBridgeAddressConfigs[0].BridgeAddress
	bridgeAssetDenom := ethservice.BlockChain().Config().AstriaBridgeAddressConfigs[0].AssetDenom
	chainDestinationKey, err := crypto.GenerateKey()
	require.Nil(t, err, "failed to generate chain destination key: %v", err)
	chainDestinationAddress := crypto.PubkeyToAddress(chainDestinationKey.PublicKey)

	depositTx := &sequencerblockv1.RollupData{Value: &sequencerblockv1.RollupData_Deposit{Deposit: &sequencerblockv1.Deposit{
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
	}}}

	finalTxs := []*sequencerblockv1.RollupData{seqData1, seqData2, auctionResultSequenceData, depositTx}

	txsToProcess, err := UnbundleRollupData(finalTxs, 2, serviceV1Alpha1.BridgeAddresses(), serviceV1Alpha1.BridgeAllowedAssets(), prevRollupBlockHash)
	require.NoError(t, err, "failed to unbundle rollup data: %v", err)

	require.Equal(t, txsToProcess.Len(), 6, "expected 6 txs to process")

	// auction result txs should be the first 3
	require.True(t, bytes.Equal(txsToProcess[0].Hash().Bytes(), tx1.Hash().Bytes()), "expected tx1 to be first")
	require.True(t, bytes.Equal(txsToProcess[1].Hash().Bytes(), tx2.Hash().Bytes()), "expected tx2 to be second")
	require.True(t, bytes.Equal(txsToProcess[2].Hash().Bytes(), tx3.Hash().Bytes()), "expected tx3 to be third")
}
