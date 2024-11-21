package shared

import (
	bundlev1alpha1 "buf.build/gen/go/astria/execution-apis/protocolbuffers/go/astria/bundle/v1alpha1"
	primitivev1 "buf.build/gen/go/astria/primitives/protocolbuffers/go/astria/primitive/v1"
	sequencerblockv1 "buf.build/gen/go/astria/sequencerblock-apis/protocolbuffers/go/astria/sequencerblock/v1"
	"bytes"
	"crypto/ed25519"
	"crypto/sha256"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"math/big"
)

func protoU128ToBigInt(u128 *primitivev1.Uint128) *big.Int {
	lo := big.NewInt(0).SetUint64(u128.Lo)
	hi := big.NewInt(0).SetUint64(u128.Hi)
	hi.Lsh(hi, 64)
	return lo.Add(lo, hi)
}

func validateAndUnmarshalDepositTx(
	deposit *sequencerblockv1.Deposit,
	height uint64,
	bridgeAddresses map[string]*params.AstriaBridgeAddressConfig,
	bridgeAllowedAssets map[string]struct{}) (*types.Transaction, error) {
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
}

func validateAndUnmarshallSequenceAction(tx *sequencerblockv1.RollupData) (*types.Transaction, error) {
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

func unmarshallAllocationTxs(allocation *bundlev1alpha1.Allocation, prevBlockHash []byte, auctioneerBech32Address string, addressPrefix string) (types.Transactions, error) {
	processedTxs := types.Transactions{}
	payload := allocation.GetPayload()

	if !bytes.Equal(payload.PrevRollupBlockHash, prevBlockHash) {
		return nil, errors.New("prev block hash do not match in allocation")
	}

	publicKey := ed25519.PublicKey(allocation.GetPublicKey())
	bech32Address, err := EncodeFromPublicKey(addressPrefix, publicKey)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to encode public key to bech32m address: %s", publicKey)
	}
	if auctioneerBech32Address != bech32Address {
		return nil, errors.Errorf("address in allocation does not match auctioneer address. expected: %s, got: %s", auctioneerBech32Address, bech32Address)
	}

	message, err := proto.Marshal(allocation.GetPayload())
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal allocation")
	}

	signature := allocation.GetSignature()
	if !ed25519.Verify(publicKey, message, signature) {
		return nil, errors.New("failed to verify signature")
	}

	// unmarshall the transactions in the bundle
	for _, allocationTx := range payload.GetTransactions() {
		ethtx := new(types.Transaction)
		err := ethtx.UnmarshalBinary(allocationTx)
		if err != nil {
			return nil, errors.Wrap(err, "failed to unmarshall allocation transaction")
		}
		processedTxs = append(processedTxs, ethtx)
	}

	return processedTxs, nil

}

// `UnbundleRollupDataTransactions` takes in a list of rollup data transactions and returns a list of Ethereum transactions.
// TODO - this function has become too big. we should start breaking it down
func UnbundleRollupDataTransactions(txs []*sequencerblockv1.RollupData, height uint64, bridgeAddresses map[string]*params.AstriaBridgeAddressConfig,
	bridgeAllowedAssets map[string]struct{}, prevBlockHash []byte, auctioneerBech32Address string, addressPrefix string) types.Transactions {
	processedTxs := types.Transactions{}
	allocationTxs := types.Transactions{}
	// we just return the allocation here and do not unmarshall the transactions in the bundle if we find it
	var allocation *bundlev1alpha1.Allocation
	for _, tx := range txs {
		if deposit := tx.GetDeposit(); deposit != nil {
			depositTx, err := validateAndUnmarshalDepositTx(deposit, height, bridgeAddresses, bridgeAllowedAssets)
			if err != nil {
				log.Error("failed to validate and unmarshal deposit tx", "error", err)
				continue
			}

			processedTxs = append(processedTxs, depositTx)
		} else {
			sequenceData := tx.GetSequencedData()
			// check if sequence data is of type Allocation
			if allocation == nil {
				// TODO - check if we can avoid a temp value
				tempAllocation := &bundlev1alpha1.Allocation{}
				err := proto.Unmarshal(sequenceData, tempAllocation)
				if err == nil {
					unmarshalledAllocationTxs, err := unmarshallAllocationTxs(tempAllocation, prevBlockHash, auctioneerBech32Address, addressPrefix)
					if err != nil {
						log.Error("failed to unmarshall allocation transactions", "error", err)
						continue
					}

					allocation = tempAllocation
					allocationTxs = unmarshalledAllocationTxs
				} else {
					ethtx, err := validateAndUnmarshallSequenceAction(tx)
					if err != nil {
						log.Error("failed to unmarshall sequence action", "error", err)
						continue
					}
					processedTxs = append(processedTxs, ethtx)
				}
			} else {
				ethtx, err := validateAndUnmarshallSequenceAction(tx)
				if err != nil {
					log.Error("failed to unmarshall sequence action", "error", err)
					continue
				}
				processedTxs = append(processedTxs, ethtx)
			}
		}
	}

	// prepend allocation txs to processedTxs
	processedTxs = append(allocationTxs, processedTxs...)

	return processedTxs
}
