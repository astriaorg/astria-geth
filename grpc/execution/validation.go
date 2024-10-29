package execution

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	astriaPb "buf.build/gen/go/astria/execution-apis/protocolbuffers/go/astria/execution/v1"
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
func validateAndUnmarshalSequencerTx(
	height uint64,
	tx *sequencerblockv1.RollupData,
	bridgeAddresses map[string]*params.AstriaBridgeAddressConfig,
	bridgeAllowedAssets map[string]struct{},
) (*types.Transaction, error) {
	if deposit := tx.GetDeposit(); deposit != nil {
		bridgeAddress := deposit.BridgeAddress.GetBech32M()
		bac, ok := bridgeAddresses[bridgeAddress]
		if !ok {
			return nil, fmt.Errorf("unknown bridge address: %s", bridgeAddress)
		}

		if height < uint64(bac.StartHeight) {
			return nil, fmt.Errorf("bridging asset %s from bridge %s not allowed before height %d", bac.AssetDenom, bridgeAddress, bac.StartHeight)
		}

		if _, ok := bridgeAllowedAssets[deposit.Asset]; !ok {
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
	if req.PrevBlockHash == nil {
		return fmt.Errorf("PrevBlockHash cannot be nil")
	}
	if req.Timestamp == nil {
		return fmt.Errorf("Timestamp cannot be nil")
	}

	return nil
}

func validateStaticExecuteOptimisticBlockRequest(req *sequencerblockv1alpha1.BaseBlock) error {
	if req.Timestamp == nil {
		return fmt.Errorf("Timestamp cannot be nil")
	}
	if len(req.SequencerBlockHash) == 0 {
		return fmt.Errorf("SequencerBlockHash cannot be empty")
	}

	return nil
}

// `validateStaticCommitment` validates the given commitment without regard to the current state of the system.
func validateStaticCommitmentState(commitmentState *astriaPb.CommitmentState) error {
	if commitmentState == nil {
		return fmt.Errorf("commitment state is nil")
	}
	if commitmentState.Soft == nil {
		return fmt.Errorf("soft block is nil")
	}
	if commitmentState.Firm == nil {
		return fmt.Errorf("firm block is nil")
	}
	if commitmentState.BaseCelestiaHeight == 0 {
		return fmt.Errorf("base celestia height of 0 is not valid")
	}

	if err := validateStaticBlock(commitmentState.Soft); err != nil {
		return fmt.Errorf("soft block invalid: %w", err)
	}
	if err := validateStaticBlock(commitmentState.Firm); err != nil {
		return fmt.Errorf("firm block invalid: %w", err)
	}

	return nil
}

// `validateStaticBlock` validates the given block as a  without regard to the current state of the system.
func validateStaticBlock(block *astriaPb.Block) error {
	if block.ParentBlockHash == nil {
		return fmt.Errorf("parent block hash is nil")
	}
	if block.Hash == nil {
		return fmt.Errorf("block hash is nil")
	}
	if block.Timestamp == nil {
		return fmt.Errorf("timestamp is 0")
	}

	return nil
}
