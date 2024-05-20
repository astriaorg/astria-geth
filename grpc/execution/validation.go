package execution

import (
	sequencerblockv1alpha1 "buf.build/gen/go/astria/sequencerblock-apis/protocolbuffers/go/astria/sequencerblock/v1alpha1"
	"crypto/sha256"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

// if sequencer tx is valid, then a unmarshalled ethereum transaction is returned. if not valid, nil is returned
func (s *ExecutionServiceServerV1Alpha2) SequencerTxValidation(tx *sequencerblockv1alpha1.RollupData) (*types.Transaction, error) {
	if deposit := tx.GetDeposit(); deposit != nil {
		bridgeAddress := string(deposit.BridgeAddress.GetInner())
		bac, ok := s.bridgeAddresses[bridgeAddress]
		if !ok {
			log.Debug("ignoring deposit tx from unknown bridge", "bridgeAddress", bridgeAddress)
			return nil, fmt.Errorf("unknown bridge address: %s", bridgeAddress)
		}

		if len(deposit.AssetId) != 32 {
			log.Debug("ignoring deposit tx with invalid asset ID", "assetID", deposit.AssetId)
			return nil, fmt.Errorf("invalid asset ID: %x", deposit.AssetId)
		}
		assetID := [32]byte{}
		copy(assetID[:], deposit.AssetId[:32])
		if _, ok := s.bridgeAllowedAssetIDs[assetID]; !ok {
			log.Debug("ignoring deposit tx with disallowed asset ID", "assetID", deposit.AssetId)
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
			log.Error("failed to unmarshal sequenced data into transaction, ignoring", "tx hash", sha256.Sum256(tx.GetSequencedData()), "err", err)
			return nil, fmt.Errorf("failed to unmarshal sequenced data into transaction: %w. tx hash: %s", err, sha256.Sum256(tx.GetSequencedData()))
		}

		if ethTx.Type() == types.DepositTxType {
			log.Debug("ignoring deposit tx in sequenced data", "tx hash", sha256.Sum256(tx.GetSequencedData()))
			return nil, fmt.Errorf("deposit tx not allowed in sequenced data. tx hash: %s", sha256.Sum256(tx.GetSequencedData()))
		}

		if ethTx.Type() == types.BlobTxType {
			log.Debug("ignoring blob tx in sequenced data", "tx hash", sha256.Sum256(tx.GetSequencedData()))
			return nil, fmt.Errorf("blob tx not allowed in sequenced data. tx hash: %s", sha256.Sum256(tx.GetSequencedData()))
		}
	}

	return nil, fmt.Errorf("unknown sequencer tx type")
}
