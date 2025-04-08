package execution

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"

	astriaPb "buf.build/gen/go/astria/execution-apis/protocolbuffers/go/astria/execution/v2"
	sequencerblockv1 "buf.build/gen/go/astria/sequencerblock-apis/protocolbuffers/go/astria/sequencerblock/v1"
	connecttypesv2 "buf.build/gen/go/astria/vendored/protocolbuffers/go/connect/types/v2"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/holiman/uint256"
)

func hashCurrencyPair(currencyPair *connecttypesv2.CurrencyPair) [32]byte {
	cpStr := fmt.Sprintf("%s/%s", currencyPair.Base, currencyPair.Quote)
	var bytes [32]byte
	copy(bytes[:], crypto.Keccak256([]byte(cpStr)))
	return bytes
}

func validateAndConvertPriceFeedDataTx(
	ctx context.Context,
	height uint64,
	priceFeedData *sequencerblockv1.PriceFeedData,
	cfg *conversionConfig,
) ([]*types.Transaction, error) {
	txs := make([]*types.Transaction, 0)

	log.Debug("creating price feed data update tx", "price count", len(priceFeedData.Prices))
	abi, err := contracts.AstriaOracleMetaData.GetAbi()
	if err != nil {
		// this should never happen, as the abi is hardcoded in the contract bindings
		return nil, fmt.Errorf("failed to get abi for AstriaOracle: %w", err)
	}

	// subtract 1 from the height to get the parent height (height parameter is the height of the current block being executed)
	state, header, err := cfg.api.StateAndHeaderByNumber(ctx, rpc.BlockNumber(int64(height-1)))
	if err != nil {
		return nil, fmt.Errorf("failed to get state and header for height %d: %w", height-1, err)
	}

	// arguments for calling the `setPrices()` function on the oracle contract
	currencyPairsToSet := make([][32]byte, 0)
	pricesToSet := make([]*big.Int, 0)

	// see if contract requires currency pair authorization
	evm := cfg.api.GetEVM(ctx, &core.Message{GasPrice: big.NewInt(0)}, state, header, &vm.Config{NoBaseFee: true}, nil)
	res, err := callContract(abi, evm, cfg.oracleCallerAddress, cfg.oracleContractAddress, "requireCurrencyPairAuthorization")
	if err != nil {
		return nil, fmt.Errorf("failed to call requireCurrencyPairAuthorization: %w", err)
	}

	// result should be `(bool requireCurrencyPairAuthorization)`
	if len(res) != 1 {
		return nil, fmt.Errorf("unexpected result length from requireCurrencyPairAuthorization: %d", len(res))
	}

	requireCurrencyPairAuthorization, ok := res[0].(bool)
	if !ok {
		return nil, fmt.Errorf("unexpected return type for requireCurrencyPairAuthorization: %T", res[0])
	}

	for _, price := range priceFeedData.Prices {
		currencyPairHash := hashCurrencyPair(price.CurrencyPair)

		// see if currency pair was already initialized; if not, initialize it
		//
		// to check if it was initialized, we call `currencyPairInfo()` on the parent state; since oracle data is always top of block,
		// if the currency pair is not initialized in the parent state, then we need to initialize it here
		// as it has never been initialized before.
		res, err := callContract(abi, evm, cfg.oracleCallerAddress, cfg.oracleContractAddress, "currencyPairInfo", currencyPairHash)
		if err != nil {
			return nil, fmt.Errorf("failed to call currencyPairInfo: %w", err)
		}

		// result should be `(bool initialized, uint8 decimals)`
		if len(res) != 2 {
			return nil, fmt.Errorf("unexpected result length from currencyPairInfo: %d", len(res))
		}

		init, ok := res[0].(bool)
		if !ok {
			return nil, fmt.Errorf("unexpected return type for initialized: %T", res[0])
		}
		if init {
			currencyPairsToSet = append(currencyPairsToSet, currencyPairHash)
			pricesToSet = append(pricesToSet, protoI128ToBigInt(price.Price))
			continue
		}

		// if contract requires currency pair authorization to initialize,
		// check if pair is authorized and skip it if not
		if requireCurrencyPairAuthorization {
			res, err := callContract(abi, evm, cfg.oracleCallerAddress, cfg.oracleContractAddress, "authorizedCurrencyPairs", currencyPairHash)
			if err != nil {
				return nil, fmt.Errorf("failed to call authorizedCurrencyPairs: %w", err)
			}

			// result should be `(bool)`
			if len(res) != 1 {
				return nil, fmt.Errorf("unexpected result length from authorizedCurrencyPairs: %d", len(res))
			}

			authorized, ok := res[0].(bool)
			if !ok {
				return nil, fmt.Errorf("unexpected return type for authorizedCurrencyPairs: %T", res[0])
			}
			if !authorized {
				continue
			}
		}

		// pack arguments for calling the `initializeCurrencyPair` function on the oracle contract
		args := []interface{}{currencyPairHash, uint8(price.Decimals)}
		calldata, err := abi.Pack("initializeCurrencyPair", args...)
		if err != nil {
			return nil, fmt.Errorf("failed to pack args for initializeCurrencyPair: %w", err)
		}

		txdata := types.InjectedTx{
			From:                   cfg.oracleCallerAddress,
			Value:                  new(big.Int),
			Gas:                    100000,
			To:                     &cfg.oracleContractAddress,
			Data:                   calldata,
			SourceTransactionId:    "",
			SourceTransactionIndex: 0,
		}
		tx := types.NewTx(&txdata)
		txs = append(txs, tx)
		currencyPairsToSet = append(currencyPairsToSet, currencyPairHash)
		pricesToSet = append(pricesToSet, protoI128ToBigInt(price.Price))
		log.Debug("created initializeCurrencyPair tx for currency pair", "pair", price.CurrencyPair, "hash", hex.EncodeToString(currencyPairHash[:]))
	}

	args := []interface{}{currencyPairsToSet, pricesToSet}
	calldata, err := abi.Pack("setPrices", args...)
	if err != nil {
		return nil, err
	}

	txdata := types.InjectedTx{
		From:  cfg.oracleCallerAddress,
		Value: new(big.Int),
		// TODO: max gas costs; proportional to the amount of pairs being updated
		Gas:                    900000,
		To:                     &cfg.oracleContractAddress,
		Data:                   calldata,
		SourceTransactionId:    "", // not relevant
		SourceTransactionIndex: 0,  // not relevant
	}
	log.Debug("created setPrices tx", "pairs", priceFeedData.Prices)
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
			SourceTransactionId:    deposit.SourceTransactionId.Inner,
			SourceTransactionIndex: deposit.SourceActionIndex,
		}

		return []*types.Transaction{types.NewTx(&txdata)}, nil
	}

	txdata := types.InjectedTx{
		From:                   bac.SenderAddress,
		To:                     &recipient,
		Value:                  amount,
		Gas:                    0,
		SourceTransactionId:    deposit.SourceTransactionId.Inner,
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
// If the sequencer transaction is an price feed data update tx, we create a transaction to update the price feed data.
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
	case tx.GetPriceFeedData() != nil:
		return validateAndConvertPriceFeedDataTx(ctx, height, tx.GetPriceFeedData(), cfg)
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

func callContract(abi *abi.ABI, evm *vm.EVM, from common.Address, to common.Address, methodName string, args ...interface{}) ([]interface{}, error) {
	const CONTRACT_CALL_GAS = 100000 // gas is arbitrary

	calldata, err := abi.Pack(methodName, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to pack args for %s: %w", methodName, err)
	}

	ret, _, err := evm.Call(vm.AccountRef(from), to, calldata, CONTRACT_CALL_GAS, uint256.NewInt(0))
	if err != nil {
		return nil, fmt.Errorf("failed to call %s: %w", methodName, err)
	}

	res, err := abi.Unpack(methodName, ret)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack %s: %w", methodName, err)
	}

	return res, nil
}
