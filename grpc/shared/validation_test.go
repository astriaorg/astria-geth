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

func TestUnmarshallAllocationTxs(t *testing.T) {
	ethService, serviceV1Alpha1, auctioneerPrivKey, auctioneerPubKey := SetupSharedService(t, 10)
	addressPrefix := ethService.BlockChain().Config().AstriaSequencerAddressPrefix

	tx1 := transaction(0, 1000, TestKey)
	validMarshalledTx1, err := tx1.MarshalBinary()
	require.NoError(t, err, "failed to marshal valid tx: %v", err)

	tx2 := transaction(1, 1000, TestKey)
	validMarshalledTx2, err := tx2.MarshalBinary()
	require.NoError(t, err, "failed to marshal valid tx: %v", err)

	tx3 := transaction(2, 1000, TestKey)
	validMarshalledTx3, err := tx3.MarshalBinary()
	require.NoError(t, err, "failed to marshal valid tx: %v", err)

	validPayload := &bundlev1alpha1.Bundle{
		Fee:                    100,
		Transactions:           [][]byte{validMarshalledTx1, validMarshalledTx2, validMarshalledTx3},
		BaseSequencerBlockHash: []byte("sequencer block hash"),
		PrevRollupBlockHash:    []byte("prev rollup block hash"),
	}

	marshalledAllocation, err := proto.Marshal(validPayload)
	require.NoError(t, err, "failed to marshal payload: %v", err)

	signedAllocation, err := auctioneerPrivKey.Sign(nil, marshalledAllocation, &ed25519.Options{
		Hash:    0,
		Context: "",
	})
	require.NoError(t, err, "failed to sign allocation: %v", err)

	tests := []struct {
		description    string
		allocation     *bundlev1alpha1.Allocation
		prevBlockHash  []byte
		expectedOutput types.Transactions
		// just check if error contains the string since error contains other details
		wantErr string
	}{
		{
			description: "previous block hash mismatch",
			allocation: &bundlev1alpha1.Allocation{
				// TODO - add signature and public key validation
				Signature: make([]byte, 0),
				PublicKey: make([]byte, 0),
				Payload: &bundlev1alpha1.Bundle{
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
			description: "public key doesn't match",
			allocation: &bundlev1alpha1.Allocation{
				Signature: []byte("invalid signature"),
				PublicKey: []byte("invalid public key"),
				Payload: &bundlev1alpha1.Bundle{
					Fee:                    100,
					Transactions:           [][]byte{[]byte("unmarshallable tx")},
					BaseSequencerBlockHash: []byte("sequencer block hash"),
					PrevRollupBlockHash:    []byte("prev rollup block hash"),
				},
			},
			prevBlockHash:  []byte("prev rollup block hash"),
			expectedOutput: types.Transactions{},
			wantErr:        "address in allocation does not match auctioneer address",
		},
		{
			description: "invalid signature",
			allocation: &bundlev1alpha1.Allocation{
				Signature: []byte("invalid signature"),
				PublicKey: auctioneerPubKey,
				Payload: &bundlev1alpha1.Bundle{
					Fee:                    100,
					Transactions:           [][]byte{[]byte("unmarshallable tx")},
					BaseSequencerBlockHash: []byte("sequencer block hash"),
					PrevRollupBlockHash:    []byte("prev rollup block hash"),
				},
			},
			prevBlockHash:  []byte("prev rollup block hash"),
			expectedOutput: types.Transactions{},
			wantErr:        "signature in allocation does not match the public key",
		},
		{
			description: "valid allocation",
			allocation: &bundlev1alpha1.Allocation{
				Signature: signedAllocation,
				PublicKey: auctioneerPubKey,
				Payload:   validPayload,
			},
			prevBlockHash:  []byte("prev rollup block hash"),
			expectedOutput: types.Transactions{tx1, tx2, tx3},
			wantErr:        "",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			finalTxs, err := unmarshallAllocationTxs(test.allocation, test.prevBlockHash, serviceV1Alpha1.AuctioneerAddress(), addressPrefix)
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
	ethservice, serviceV1Alpha1, _, _ := SetupSharedService(t, 10)

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
	ethservice, serviceV1Alpha1, auctioneerPrivKey, auctioneerPubKey := SetupSharedService(t, 10)

	addressPrefix := ethservice.BlockChain().Config().AstriaSequencerAddressPrefix

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

	payload := &bundlev1alpha1.Bundle{
		Fee:                    100,
		Transactions:           [][]byte{validMarshalledTx1, validMarshalledTx2, validMarshalledTx3},
		BaseSequencerBlockHash: baseSequencerBlockHash,
		PrevRollupBlockHash:    prevRollupBlockHash,
	}

	marshalledPayload, err := proto.Marshal(payload)
	require.NoError(t, err, "failed to marshal payload: %v", err)
	signedPayload, err := auctioneerPrivKey.Sign(nil, marshalledPayload, &ed25519.Options{
		Hash:    0,
		Context: "",
	})
	require.NoError(t, err, "failed to sign payload: %v", err)

	allocation := &bundlev1alpha1.Allocation{
		Signature: signedPayload,
		PublicKey: auctioneerPubKey,
		Payload:   payload,
	}

	marshalledAllocation, err := proto.Marshal(allocation)
	require.NoError(t, err, "failed to marshal allocation: %v", err)
	allocationSequenceData := &sequencerblockv1.RollupData{
		Value: &sequencerblockv1.RollupData_SequencedData{
			SequencedData: marshalledAllocation,
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

	finalTxs := []*sequencerblockv1.RollupData{seqData1, seqData2, allocationSequenceData, depositTx}

	txsToProcess := UnbundleRollupDataTransactions(finalTxs, 2, serviceV1Alpha1.BridgeAddresses(), serviceV1Alpha1.BridgeAllowedAssets(), prevRollupBlockHash, serviceV1Alpha1.AuctioneerAddress(), addressPrefix)

	require.Equal(t, txsToProcess.Len(), 6, "expected 6 txs to process")

	// allocation txs should be the first 3
	require.True(t, bytes.Equal(txsToProcess[0].Hash().Bytes(), tx1.Hash().Bytes()), "expected tx1 to be first")
	require.True(t, bytes.Equal(txsToProcess[1].Hash().Bytes(), tx2.Hash().Bytes()), "expected tx2 to be second")
	require.True(t, bytes.Equal(txsToProcess[2].Hash().Bytes(), tx3.Hash().Bytes()), "expected tx3 to be third")
	require.True(t, bytes.Equal(txsToProcess[3].Hash().Bytes(), tx4.Hash().Bytes()), "expected tx4 to be fourth")
	require.True(t, bytes.Equal(txsToProcess[4].Hash().Bytes(), tx5.Hash().Bytes()), "expected tx5 to be fifth")
}

func TestUnbundleRollupDataWithDuplicateAllocations(t *testing.T) {
	ethservice, serviceV1Alpha1, auctioneerPrivKey, auctioneerPubKey := SetupSharedService(t, 10)
	addressPrefix := ethservice.BlockChain().Config().AstriaSequencerAddressPrefix

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

	payload := &bundlev1alpha1.Bundle{
		Fee:                    100,
		Transactions:           [][]byte{validMarshalledTx1, validMarshalledTx2, validMarshalledTx3},
		BaseSequencerBlockHash: baseSequencerBlockHash,
		PrevRollupBlockHash:    prevRollupBlockHash,
	}

	marshalledPayload, err := proto.Marshal(payload)
	require.NoError(t, err, "failed to marshal payload: %v", err)
	signedPayload, err := auctioneerPrivKey.Sign(nil, marshalledPayload, &ed25519.Options{
		Hash:    0,
		Context: "",
	})
	require.NoError(t, err, "failed to sign payload: %v", err)

	allocation := &bundlev1alpha1.Allocation{
		Signature: signedPayload,
		PublicKey: auctioneerPubKey,
		Payload:   payload,
	}

	marshalledAllocation, err := proto.Marshal(allocation)
	require.NoError(t, err, "failed to marshal allocation: %v", err)
	allocationSequenceData := &sequencerblockv1.RollupData{
		Value: &sequencerblockv1.RollupData_SequencedData{
			SequencedData: marshalledAllocation,
		},
	}
	// this allocation should be ignored
	allocationSequenceData2 := &sequencerblockv1.RollupData{
		Value: &sequencerblockv1.RollupData_SequencedData{
			SequencedData: marshalledAllocation,
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

	finalTxs := []*sequencerblockv1.RollupData{seqData1, seqData2, allocationSequenceData, allocationSequenceData2, depositTx}

	txsToProcess := UnbundleRollupDataTransactions(finalTxs, 2, serviceV1Alpha1.BridgeAddresses(), serviceV1Alpha1.BridgeAllowedAssets(), prevRollupBlockHash, serviceV1Alpha1.AuctioneerAddress(), addressPrefix)

	require.Equal(t, txsToProcess.Len(), 6, "expected 6 txs to process")

	// allocation txs should be the first 3
	require.True(t, bytes.Equal(txsToProcess[0].Hash().Bytes(), tx1.Hash().Bytes()), "expected tx1 to be first")
	require.True(t, bytes.Equal(txsToProcess[1].Hash().Bytes(), tx2.Hash().Bytes()), "expected tx2 to be second")
	require.True(t, bytes.Equal(txsToProcess[2].Hash().Bytes(), tx3.Hash().Bytes()), "expected tx3 to be third")
	require.True(t, bytes.Equal(txsToProcess[3].Hash().Bytes(), tx4.Hash().Bytes()), "expected tx4 to be fourth")
	require.True(t, bytes.Equal(txsToProcess[4].Hash().Bytes(), tx5.Hash().Bytes()), "expected tx5 to be fifth")
}

func TestUnbundleRollupDataWithDuplicateInvalidAllocations(t *testing.T) {
	ethservice, serviceV1Alpha1, auctioneerPrivKey, auctioneerPubKey := SetupSharedService(t, 10)
	addressPrefix := ethservice.BlockChain().Config().AstriaSequencerAddressPrefix

	baseSequencerBlockHash := []byte("sequencer block hash")
	prevRollupBlockHash := []byte("prev rollup block hash")

	_, invalidAuctioneerprivkey, err := ed25519.GenerateKey(nil)
	require.Nil(t, err, "failed to generate invalid auctioneer key: %v", err)

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

	// transactions that the attacker is trying to get into the top of block
	invalidTx1 := transaction(5, 1000, TestKey)
	invalidMarshalledTx1, err := invalidTx1.MarshalBinary()
	require.NoError(t, err, "failed to marshal valid tx: %v", err)
	invalidTx2 := transaction(6, 1000, TestKey)
	invalidMarshalledTx2, err := invalidTx2.MarshalBinary()
	require.NoError(t, err, "failed to marshal valid tx: %v", err)

	payload := &bundlev1alpha1.Bundle{
		Fee:                    100,
		Transactions:           [][]byte{validMarshalledTx1, validMarshalledTx2, validMarshalledTx3},
		BaseSequencerBlockHash: baseSequencerBlockHash,
		PrevRollupBlockHash:    prevRollupBlockHash,
	}

	marshalledPayload, err := proto.Marshal(payload)
	require.NoError(t, err, "failed to marshal allocation: %v", err)
	signedPayload, err := auctioneerPrivKey.Sign(nil, marshalledPayload, &ed25519.Options{
		Hash:    0,
		Context: "",
	})
	require.NoError(t, err, "failed to sign allocation: %v", err)

	invalidPayload := &bundlev1alpha1.Bundle{
		Fee:                    100,
		Transactions:           [][]byte{invalidMarshalledTx1, invalidMarshalledTx2},
		BaseSequencerBlockHash: baseSequencerBlockHash,
		PrevRollupBlockHash:    prevRollupBlockHash,
	}
	marshalledInvalidPayload, err := proto.Marshal(invalidPayload)
	require.NoError(t, err, "failed to marshal invalid allocation: %v", err)

	signedInvalidPayload, err := invalidAuctioneerprivkey.Sign(nil, marshalledInvalidPayload, &ed25519.Options{
		Hash:    0,
		Context: "",
	})
	require.NoError(t, err, "failed to sign allocation: %v", err)

	allocation := &bundlev1alpha1.Allocation{
		Signature: signedPayload,
		PublicKey: auctioneerPubKey,
		Payload:   payload,
	}

	marshalledAllocation, err := proto.Marshal(allocation)
	require.NoError(t, err, "failed to marshal allocation: %v", err)

	invalidAllocation := &bundlev1alpha1.Allocation{
		Signature: signedInvalidPayload,
		// trying to spoof the actual auctioneer key
		PublicKey: auctioneerPubKey,
		Payload:   invalidPayload,
	}
	marshalledInvalidAllocation, err := proto.Marshal(invalidAllocation)
	require.NoError(t, err, "failed to marshal invalid allocation: %v", err)

	allocationSequenceData := &sequencerblockv1.RollupData{
		Value: &sequencerblockv1.RollupData_SequencedData{
			SequencedData: marshalledAllocation,
		},
	}
	// this allocation should be ignored
	invalidAllocationSequenceData := &sequencerblockv1.RollupData{
		Value: &sequencerblockv1.RollupData_SequencedData{
			SequencedData: marshalledInvalidAllocation,
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

	finalTxs := []*sequencerblockv1.RollupData{seqData1, seqData2, allocationSequenceData, invalidAllocationSequenceData, depositTx}

	txsToProcess := UnbundleRollupDataTransactions(finalTxs, 2, serviceV1Alpha1.BridgeAddresses(), serviceV1Alpha1.BridgeAllowedAssets(), prevRollupBlockHash, serviceV1Alpha1.AuctioneerAddress(), addressPrefix)

	require.Equal(t, txsToProcess.Len(), 6, "expected 6 txs to process")

	// allocation txs should be the first 3
	require.True(t, bytes.Equal(txsToProcess[0].Hash().Bytes(), tx1.Hash().Bytes()), "expected tx1 to be first")
	require.True(t, bytes.Equal(txsToProcess[1].Hash().Bytes(), tx2.Hash().Bytes()), "expected tx2 to be second")
	require.True(t, bytes.Equal(txsToProcess[2].Hash().Bytes(), tx3.Hash().Bytes()), "expected tx3 to be third")
	require.True(t, bytes.Equal(txsToProcess[3].Hash().Bytes(), tx4.Hash().Bytes()), "expected tx4 to be fourth")
	require.True(t, bytes.Equal(txsToProcess[4].Hash().Bytes(), tx5.Hash().Bytes()), "expected tx5 to be fifth")
}
