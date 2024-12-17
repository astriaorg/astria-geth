package execution

import (
	"context"
	"crypto/sha256"
	"fmt"
	"math/big"

	astriaPb "buf.build/gen/go/astria/execution-apis/protocolbuffers/go/astria/execution/v1"
	primitivev1 "buf.build/gen/go/astria/primitives/protocolbuffers/go/astria/primitive/v1"
	sequencerblockv1 "buf.build/gen/go/astria/sequencerblock-apis/protocolbuffers/go/astria/sequencerblock/v1"
	connecttypesv2 "buf.build/gen/go/astria/vendored/protocolbuffers/go/connect/types/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/holiman/uint256"
)

func hashCurrencyPair(currencyPair *connecttypesv2.CurrencyPair) [32]byte {
	cpStr := fmt.Sprintf("%s/%s", currencyPair.Base, currencyPair.Quote)
	return sha256.Sum256([]byte(cpStr))
}

func validateAndConvertOracleDataTx(
	ctx context.Context,
	height uint64,
	oracleData *sequencerblockv1.OracleData,
	cfg *conversionConfig,
) ([]*types.Transaction, error) {
	txs := make([]*types.Transaction, 0)

	log.Info("creating oracle data update tx, price count: %d", len(oracleData.Prices))
	abi, err := contracts.AstriaOracleMetaData.GetAbi()
	if err != nil {
		// this should never happen, as the abi is hardcoded in the contract bindings
		return nil, fmt.Errorf("failed to get abi for AstriaOracle: %w", err)
	}

	state, header, err := cfg.api.StateAndHeaderByNumber(ctx, rpc.BlockNumber(int64(height-1)))
	if err != nil {
		return nil, fmt.Errorf("failed to get state and header for height %d: %w", height-1, err)
	}

	// arguments for calling the `updatePriceData()` function on the oracle contract
	currencyPairs := make([][32]byte, len(oracleData.Prices))
	prices := make([]*big.Int, len(oracleData.Prices))
	for i, price := range oracleData.Prices {
		currencyPairs[i] = hashCurrencyPair(price.CurrencyPair)
		prices[i] = protoU128ToBigInt(price.Price)

		// see if currency pair was already initialized; if not, initialize it
		//
		// to check if it was initialized, we call `currencyPairInfo()` on the parent state; since oracle data is always top of block,
		// if the currency pair is not initialized in the parent state, then we need to initialize it here
		// as it has never been initialized before.
		evm := cfg.api.GetEVM(ctx, &core.Message{GasPrice: big.NewInt(1)}, state, header, &vm.Config{NoBaseFee: true}, nil)
		args := []interface{}{currencyPairs[i]}
		calldata, err := abi.Pack("currencyPairInfo", args...)
		if err != nil {
			return nil, fmt.Errorf("failed to pack args for currencyPairInfo: %w", err)
		}
		ret, _, err := evm.Call(vm.AccountRef(cfg.oracleCallerAddress), cfg.oracleContractAddress, calldata, 100000, uint256.NewInt(0)) // gas is arbitrary
		if err != nil {
			return nil, fmt.Errorf("failed to call currencyPairInfo: %w", err)
		}

		// result should be abi-packed `(bool initialized, uint8 decimals)`
		res, err := abi.Unpack("currencyPairInfo", ret)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack currencyPairInfo: %w", err)
		}
		if len(res) != 2 {
			return nil, fmt.Errorf("unexpected result length from currencyPairInfo: %d", len(res))
		}

		init, ok := res[0].(bool)
		if !ok {
			return nil, fmt.Errorf("unexpected type for initialized: %T", res[0])
		}
		if init {
			continue
		}

		// pack arguments for calling the `initializeCurrencyPair` function on the oracle contract
		args = []interface{}{currencyPairs[i], uint8(price.Decimals)}
		calldata, err = abi.Pack("initializeCurrencyPair", args...)
		if err != nil {
			return nil, fmt.Errorf("failed to pack args for initializeCurrencyPair: %w", err)
		}

		txdata := types.InjectedTx{
			From:                   cfg.oracleCallerAddress,
			Value:                  new(big.Int),
			Gas:                    500000, // TODO
			To:                     &cfg.oracleContractAddress,
			Data:                   calldata,
			SourceTransactionId:    primitivev1.TransactionId{},
			SourceTransactionIndex: 0,
		}
		tx := types.NewTx(&txdata)
		txs = append(txs, tx)
		log.Info("created initializeCurrencyPair tx for currency pair %s", currencyPairs[i])
	}

	args := []interface{}{}
	calldata, err := abi.Pack("updatePriceData", args...)
	if err != nil {
		return nil, err
	}

	txdata := types.InjectedTx{
		From:  cfg.oracleCallerAddress,
		Value: new(big.Int),
		// TODO: max gas costs?
		Gas:                    500000,
		To:                     &cfg.oracleContractAddress,
		Data:                   calldata,
		SourceTransactionId:    primitivev1.TransactionId{}, // not relevant
		SourceTransactionIndex: 0,                           // not relevant
	}
	log.Info("created updatePriceData tx")
	tx := types.NewTx(&txdata)
	txs = append(txs, tx)
	return txs, nil
}

func validateAndConvertDepositTx(
	height uint64,
	deposit *sequencerblockv1.Deposit,
	cfg *conversionConfig,
) ([]*types.Transaction, error) {
	bridgeAddress := deposit.BridgeAddress.GetBech32M()
	bac, ok := cfg.bridgeAddresses[bridgeAddress]
	if !ok {
		return nil, fmt.Errorf("unknown bridge address: %s", bridgeAddress)
	}

	if height < uint64(bac.StartHeight) {
		return nil, fmt.Errorf("bridging asset %s from bridge %s not allowed before height %d", bac.AssetDenom, bridgeAddress, bac.StartHeight)
	}

	if _, ok := cfg.bridgeAllowedAssets[deposit.Asset]; !ok {
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

		txdata := types.InjectedTx{
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

		return []*types.Transaction{types.NewTx(&txdata)}, nil
	}

	txdata := types.InjectedTx{
		From:                   bac.SenderAddress,
		To:                     &recipient,
		Value:                  amount,
		Gas:                    0,
		SourceTransactionId:    *deposit.SourceTransactionId,
		SourceTransactionIndex: deposit.SourceActionIndex,
	}
	return []*types.Transaction{types.NewTx(&txdata)}, nil
}

func validateAndConvertSequencedDataTx(sequencedData []byte) ([]*types.Transaction, error) {
	ethTx := new(types.Transaction)
	err := ethTx.UnmarshalBinary(sequencedData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal sequenced data into transaction: %w. tx hash: %s", err, sha256.Sum256(sequencedData))
	}

	if ethTx.Type() == types.InjectedTxType {
		return nil, fmt.Errorf("injected tx not allowed in sequenced data. tx hash: %s", sha256.Sum256(sequencedData))
	}

	if ethTx.Type() == types.BlobTxType {
		return nil, fmt.Errorf("blob tx not allowed in sequenced data. tx hash: %s", sha256.Sum256(sequencedData))
	}

	return []*types.Transaction{ethTx}, nil
}

// `validateAndConvertSequencerTx` validates and unmarshals the given rollup sequencer transaction and converts it into
// an EVM transaction.
// If the sequencer transaction is an oracle data update tx, we create a transaction to update the oracle data.
// If the sequencer transaction is a deposit tx, we ensure that the asset ID is allowed and the bridge address is known.
// If the sequencer transaction is a normal user tx, we unmarshal the sequenced data into an Ethereum transaction. We ensure that the
// tx is not a blob tx or a deposit tx.
func validateAndConvertSequencerTx(
	ctx context.Context,
	height uint64,
	tx *sequencerblockv1.RollupData,
	cfg *conversionConfig,
) ([]*types.Transaction, error) {
	switch {
	case tx.GetOracleData() != nil:
		return validateAndConvertOracleDataTx(ctx, height, tx.GetOracleData(), cfg)
	case tx.GetDeposit() != nil:
		return validateAndConvertDepositTx(height, tx.GetDeposit(), cfg)
	case tx.GetSequencedData() != nil:
		return validateAndConvertSequencedDataTx(tx.GetSequencedData())
	default:
		return nil, fmt.Errorf("unknown sequencer tx type %v", tx)
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
