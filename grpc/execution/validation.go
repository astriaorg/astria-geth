package execution

import (
	composerv1alpha1 "buf.build/gen/go/astria/composer-apis/protocolbuffers/go/astria/composer/v1alpha1"
	sequencerblockv1alpha1 "buf.build/gen/go/astria/sequencerblock-apis/protocolbuffers/go/astria/sequencerblock/v1alpha1"
	"crypto/sha256"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/golang/protobuf/proto"
	"math/big"
)

// `validateAndUnmarshalTx` validates and unmarshals the given rollup sequencer transaction.
// If the sequencer transaction is a deposit tx, we ensure that the asset ID is allowed and the bridge address is known.
// If the sequencer transaction is not a deposit tx, we unmarshal the sequenced data into an Ethereum transaction. We ensure that the
// tx is not a blob tx or a deposit tx.
func validateAndUnmarshalTx(
	height uint64,
	tx *sequencerblockv1alpha1.RollupData,
	bridgeAddresses map[string]*params.AstriaBridgeAddressConfig,
	bridgeAllowedAssets map[string]struct{},
	bridgeSenderAddress common.Address,
) (*types.Transaction, error) {
	if deposit := tx.GetDeposit(); deposit != nil {
		return validateAndUnmarshallDepositTx(height, tx, bridgeAddresses, bridgeAllowedAssets, bridgeSenderAddress)
	} else {
		return validateAndUnmarshallSequenceTx(tx)
	}
}

func validateAndUnmarshallDepositTx(
	height uint64,
	tx *sequencerblockv1alpha1.RollupData,
	bridgeAddresses map[string]*params.AstriaBridgeAddressConfig,
	bridgeAllowedAssets map[string]struct{},
	bridgeSenderAddress common.Address,
) (*types.Transaction, error) {
	deposit := tx.GetDeposit()
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
			From:  bridgeSenderAddress,
			Value: new(big.Int), // don't need to set this, as we aren't minting the native asset
			// mints cost ~14k gas, however this can vary based on existing storage, so we add a little extra as buffer.
			//
			// the fees are spent from the "bridge account" which is not actually a real account, but is instead some
			// address defined by consensus, so the gas cost is not actually deducted from any account.
			Gas:  16000,
			To:   &bac.Erc20Asset.ContractAddress,
			Data: calldata,
		}

		tx := types.NewTx(&txdata)
		return tx, nil
	}

	txdata := types.DepositTx{
		From:  bridgeSenderAddress,
		To:    &recipient,
		Value: amount,
		Gas:   0,
	}
	return types.NewTx(&txdata), nil
}

func validateAndUnmarshallSequenceTx(
	tx *sequencerblockv1alpha1.RollupData,
) (*types.Transaction, error) {
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

func extractBuilderBundleAndTxs(txs []*sequencerblockv1alpha1.RollupData, height uint64,
	bridgeAddresses map[string]*params.AstriaBridgeAddressConfig,
	bridgeAllowedAssets map[string]struct{},
	bridgeSenderAddress common.Address) (*composerv1alpha1.BuilderBundlePacket, types.Transactions, map[string]*sequencerblockv1alpha1.Deposit, error) {
	// Extract the builder bundle from the sequencer txs
	var builderBundle *composerv1alpha1.BuilderBundlePacket
	ethTxs := types.Transactions{}
	// this is a mapping from tx hash to depositTx, we use this when we need to simulate the txs to map the tx hash to the depositTx
	depositTxMapping := map[string]*sequencerblockv1alpha1.Deposit{}
	for _, tx := range txs {
		if deposit := tx.GetDeposit(); deposit != nil {
			tx, err := validateAndUnmarshallDepositTx(height, tx, bridgeAddresses, bridgeAllowedAssets, bridgeSenderAddress)
			if err != nil {
				return nil, nil, nil, err
			}
			depositTxMapping[tx.Hash().Hex()] = deposit
			ethTxs = append(ethTxs, tx)
		} else {
			// check if we can unmarshall the sequence data to a BuilderBundlePacket
			tempBuilderBundle := &composerv1alpha1.BuilderBundlePacket{}
			err := proto.Unmarshal(tx.GetSequencedData(), tempBuilderBundle)
			if err == nil {
				// we found a builder bundle, we first check if we got a duplicate builder bundle. if we did
				// we throw an error
				// TODO - we could just ignore the duplicate builder bundle??
				if builderBundle != nil {
					return nil, nil, nil, fmt.Errorf("duplicate builder bundle found in sequencer txs")
				}

				// duplicate builder bundle not found, we set the builder bundle to the temp builder bundle
				builderBundle = tempBuilderBundle

			} else {
				// if its not a builder bundle, then it should be a regular eth tx
				ethTx, err := validateAndUnmarshallSequenceTx(tx)
				if err != nil {
					return nil, nil, nil, err
				}

				ethTxs = append(ethTxs, ethTx)
			}

		}

	}

	return builderBundle, ethTxs, depositTxMapping, nil
}
