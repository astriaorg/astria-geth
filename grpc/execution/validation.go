package execution

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	astriaPb "buf.build/gen/go/astria/execution-apis/protocolbuffers/go/astria/execution/v2"
	sequencerblockv1 "buf.build/gen/go/astria/sequencerblock-apis/protocolbuffers/go/astria/sequencerblock/v1"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

// `validateAndUnmarshalSequencerTx` validates and unmarshals the given rollup sequencer transaction.
// If the sequencer transaction is a deposit tx, we ensure that the asset ID is allowed and the bridge address is known.
// If the sequencer transaction is not a deposit tx, we unmarshal the sequenced data into an Ethereum transaction. We ensure that the
// tx is not a blob tx or a deposit tx.
func validateAndUnmarshalSequencerTx(tx *sequencerblockv1.RollupData, fork *params.AstriaForkData) (*types.Transaction, error) {
	if deposit := tx.GetDeposit(); deposit != nil {
		bridgeAddress := deposit.BridgeAddress.GetBech32M()
		bac, ok := fork.BridgeAddresses[bridgeAddress]
		if !ok {
			return nil, fmt.Errorf("unknown bridge address: %s", bridgeAddress)
		}

		if _, ok := fork.BridgeAllowedAssets[deposit.Asset]; !ok {
			return nil, fmt.Errorf("disallowed asset %s in deposit tx", deposit.Asset)
		}

		if deposit.Asset != bac.AssetDenom {
			return nil, fmt.Errorf("asset %s does not match bridge address %s asset", deposit.Asset, bridgeAddress)
		}

		recipient := common.HexToAddress(deposit.DestinationChainAddress)
		amount := bac.ScaledDepositAmount(protoU128ToBigInt(deposit.Amount))

		if bac.Erc20Asset != nil {
			log.Debug("creating deposit tx to mint ERC20 asset", "token", bac.AssetDenom, "erc20Address", bac.Erc20Asset.ContractAddress)
			abi, err := contracts.AstriaBridgeableERC20MetaData.GetAbi()
			if err != nil {
				// this should never happen, as the abi is hardcoded in the contract bindings
				return nil, fmt.Errorf("failed to get abi for erc20 contract for asset %s: %w", bac.AssetDenom, err)
			}

			// pack arguments for calling the `mint` function on the ERC20 contract
			args := []interface{}{recipient, amount}
			calldata, err := abi.Pack("mint", args...)
			if err != nil {
				return nil, err
			}

			txdata := types.DepositTx{
				From:  bac.SenderAddress,
				Value: new(big.Int), // don't need to set this, as we aren't minting the native asset
				// mints cost ~14k gas, however this can vary based on existing storage, so we add a little extra as buffer.
				//
				// the fees are spent from the "bridge account" which is not actually a real account, but is instead some
				// address defined by consensus, so the gas cost is not actually deducted from any account.
				Gas:                    64000,
				To:                     &bac.Erc20Asset.ContractAddress,
				Data:                   calldata,
				SourceTransactionId:    *deposit.SourceTransactionId,
				SourceTransactionIndex: deposit.SourceActionIndex,
			}

			tx := types.NewTx(&txdata)
			return tx, nil
		}

		txdata := types.DepositTx{
			From:                   bac.SenderAddress,
			To:                     &recipient,
			Value:                  amount,
			Gas:                    0,
			SourceTransactionId:    *deposit.SourceTransactionId,
			SourceTransactionIndex: deposit.SourceActionIndex,
		}
		return types.NewTx(&txdata), nil
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

// `validateStaticExecuteBlockRequest` validates the given execute block request without regard
// to the current state of the system. This is useful for validating the request before any
// state changes or reads are made as a basic guard.
func validateStaticExecuteBlockRequest(req *astriaPb.ExecuteBlockRequest) error {
	if req.SessionId == "" {
		return fmt.Errorf("session_id cannot be empty")
	}
	if req.ParentHash == "" {
		return fmt.Errorf("parent_hash cannot be empty")
	}
	if req.Timestamp == nil {
		return fmt.Errorf("timestamp cannot be nil")
	}

	return nil
}

// `validateStaticCommitment` validates the given commitment without regard to the current state of the system.
func validateStaticCommitmentState(commitmentState *astriaPb.CommitmentState) error {
	if commitmentState == nil {
		return fmt.Errorf("commitment state is nil")
	}
	if commitmentState.SoftExecutedBlockMetadata == nil {
		return fmt.Errorf("SoftExecutedBlockMetadata cannot be nil")
	}
	if commitmentState.FirmExecutedBlockMetadata == nil {
		return fmt.Errorf("FirmExecutedBlockMetadata cannot be nil")
	}
	if err := validateStaticExecutedBlockMetadata(commitmentState.SoftExecutedBlockMetadata, false); err != nil {
		return fmt.Errorf("soft block invalid: %w", err)
	}
	if err := validateStaticExecutedBlockMetadata(commitmentState.FirmExecutedBlockMetadata, true); err != nil {
		return fmt.Errorf("firm block invalid: %w", err)
	}

	return nil
}

// `validateStaticExecutedBlockMetadata` validates the given block metadata without regard to the current state of the system.
func validateStaticExecutedBlockMetadata(block *astriaPb.ExecutedBlockMetadata, firm bool) error {
	if !firm && block.Number == 0 {
		return fmt.Errorf("block number cannot be 0")
	}
	if block.ParentHash == "" {
		return fmt.Errorf("parent hash cannot be empty")
	}
	if block.Hash == "" {
		return fmt.Errorf("block hash cannot be empty")
	}
	if block.Timestamp == nil {
		return fmt.Errorf("timestamp cannot be nil")
	}

	return nil
}
