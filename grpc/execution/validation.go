package execution

import (
	sequencerblockv1alpha1 "buf.build/gen/go/astria/sequencerblock-apis/protocolbuffers/go/astria/sequencerblock/v1alpha1"
	"crypto/sha256"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

// `validateAndUnmarshalSequencerTx` validates and unmarshals the given rollup sequencer transaction.
// If the sequencer transaction is a deposit tx, we ensure that the asset ID is allowed and the bridge address is known.
// If the sequencer transaction is not a deposit tx, we unmarshal the sequenced data into an Ethereum transaction. We ensure that the
// tx is not a blob tx or a deposit tx.
func validateAndUnmarshalSequencerTx(tx *sequencerblockv1alpha1.RollupData, bridgeAddresses map[string]*params.AstriaBridgeAddressConfig, bridgeAllowedAssetIDs map[[32]byte]struct{}) (*types.Transaction, error) {
	if deposit := tx.GetDeposit(); deposit != nil {
		bridgeAddress := string(deposit.BridgeAddress.GetInner())
		bac, ok := bridgeAddresses[bridgeAddress]
		if !ok {
			return nil, fmt.Errorf("unknown bridge address: %s", bridgeAddress)
		}

		if len(deposit.AssetId) != 32 {
			return nil, fmt.Errorf("invalid asset ID: %x", deposit.AssetId)
		}
		assetID := [32]byte{}
		copy(assetID[:], deposit.AssetId[:32])
		if _, ok := bridgeAllowedAssetIDs[assetID]; !ok {
			return nil, fmt.Errorf("disallowed asset ID: %x", deposit.AssetId)
		}

		amount := protoU128ToBigInt(deposit.Amount)
		address := common.HexToAddress(deposit.DestinationChainAddress)
		txdata := types.DepositTx{
			From:  address,
			Value: bac.ScaledDepositAmount(amount),
			Gas:   0,
		}

		tx := types.NewTx(&txdata)
		return tx, nil
	} else {
		ethTx := new(types.Transaction)
		err := ethTx.UnmarshalBinary(tx.GetSequencedData())
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal sequenced data into transaction: %w. tx hash: %s", err, sha256.Sum256(tx.GetSequencedData()))
		}

		if ethTx.Type() == types.DepositTxType {
			return nil, fmt.Errorf("deposit tx not allowed in sequenced data. tx hash: %s", sha256.Sum256(tx.GetSequencedData()))
		}

		if ethTx.Type() == types.BlobTxType {
			return nil, fmt.Errorf("blob tx not allowed in sequenced data. tx hash: %s", sha256.Sum256(tx.GetSequencedData()))
		}

		return ethTx, nil
	}
}
