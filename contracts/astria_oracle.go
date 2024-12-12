// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AstriaOracleMetaData contains all meta data concerning the AstriaOracle contract.
var AstriaOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ORACLE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"currencyPairInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_currencyPair\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"name\":\"initializeCurrencyPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"priceData\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_currencyPairs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint128[]\",\"name\":\"_prices\",\"type\":\"uint128[]\"}],\"name\":\"updatePriceData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f80fd5b506040516106d73803806106d783398101604081905261002e9161003f565b6001600160a01b031660805261006c565b5f6020828403121561004f575f80fd5b81516001600160a01b0381168114610065575f80fd5b9392505050565b6080516106466100915f395f8181607e0152818161018f015261030b01526106465ff3fe608060405234801561000f575f80fd5b5060043610610060575f3560e01c80633595f6911461006457806338013f02146100795780634599c788146100bd57806348832f3c146100d4578063859bd5b5146100e7578063dad84e9d1461012b575b5f80fd5b610077610072366004610477565b610184565b005b6100a07f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b6100c660025481565b6040519081526020016100b4565b6100776100e236600461052a565b610300565b6101126100f536600461055d565b60016020525f908152604090205460ff8082169161010090041682565b60408051921515835260ff9091166020830152016100b4565b610165610139366004610574565b5f602081815292815260408082209093529081522080546001909101546001600160801b039091169082565b604080516001600160801b0390931683526020830191909152016100b4565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146101d55760405162461bcd60e51b81526004016101cc90610594565b60405180910390fd5b80518251146102365760405162461bcd60e51b815260206004820152602760248201527f63757272656e6379207061697220616e64207072696365206c656e677468206d6044820152660d2e6dac2e8c6d60cb1b60648201526084016101cc565b436002555f5b82518110156102fb576040518060400160405280838381518110610262576102626105d8565b60200260200101516001600160801b03168152602001428152505f8060025481526020019081526020015f205f8584815181106102a1576102a16105d8565b6020908102919091018101518252818101929092526040015f20825181546fffffffffffffffffffffffffffffffff19166001600160801b03909116178155910151600190910155806102f3816105ec565b91505061023c565b505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146103485760405162461bcd60e51b81526004016101cc90610594565b604080518082018252600180825260ff93841660208084019182525f968752919091529190932092518354915161ffff1990921690151561ff001916176101009190921602179055565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff811182821017156103cf576103cf610392565b604052919050565b5f67ffffffffffffffff8211156103f0576103f0610392565b5060051b60200190565b5f82601f830112610409575f80fd5b8135602061041e610419836103d7565b6103a6565b82815260059290921b8401810191818101908684111561043c575f80fd5b8286015b8481101561046c5780356001600160801b038116811461045f575f8081fd5b8352918301918301610440565b509695505050505050565b5f8060408385031215610488575f80fd5b823567ffffffffffffffff8082111561049f575f80fd5b818501915085601f8301126104b2575f80fd5b813560206104c2610419836103d7565b82815260059290921b840181019181810190898411156104e0575f80fd5b948201945b838610156104fe578535825294820194908201906104e5565b96505086013592505080821115610513575f80fd5b50610520858286016103fa565b9150509250929050565b5f806040838503121561053b575f80fd5b82359150602083013560ff81168114610552575f80fd5b809150509250929050565b5f6020828403121561056d575f80fd5b5035919050565b5f8060408385031215610585575f80fd5b50508035926020909101359150565b60208082526024908201527f4173747269614f7261636c653a206f6e6c79206f7261636c652063616e2075706040820152636461746560e01b606082015260800190565b634e487b7160e01b5f52603260045260245ffd5b5f6001820161060957634e487b7160e01b5f52601160045260245ffd5b506001019056fea2646970667358221220ea51ee7bfc144111b13d6d504a4639f5b68c5675339ceff47b9eda652b6cd8e664736f6c63430008150033",
}

// AstriaOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use AstriaOracleMetaData.ABI instead.
var AstriaOracleABI = AstriaOracleMetaData.ABI

// AstriaOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AstriaOracleMetaData.Bin instead.
var AstriaOracleBin = AstriaOracleMetaData.Bin

// DeployAstriaOracle deploys a new Ethereum contract, binding an instance of AstriaOracle to it.
func DeployAstriaOracle(auth *bind.TransactOpts, backend bind.ContractBackend, _oracle common.Address) (common.Address, *types.Transaction, *AstriaOracle, error) {
	parsed, err := AstriaOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AstriaOracleBin), backend, _oracle)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AstriaOracle{AstriaOracleCaller: AstriaOracleCaller{contract: contract}, AstriaOracleTransactor: AstriaOracleTransactor{contract: contract}, AstriaOracleFilterer: AstriaOracleFilterer{contract: contract}}, nil
}

// AstriaOracle is an auto generated Go binding around an Ethereum contract.
type AstriaOracle struct {
	AstriaOracleCaller     // Read-only binding to the contract
	AstriaOracleTransactor // Write-only binding to the contract
	AstriaOracleFilterer   // Log filterer for contract events
}

// AstriaOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type AstriaOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AstriaOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AstriaOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AstriaOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AstriaOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AstriaOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AstriaOracleSession struct {
	Contract     *AstriaOracle     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AstriaOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AstriaOracleCallerSession struct {
	Contract *AstriaOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AstriaOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AstriaOracleTransactorSession struct {
	Contract     *AstriaOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AstriaOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type AstriaOracleRaw struct {
	Contract *AstriaOracle // Generic contract binding to access the raw methods on
}

// AstriaOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AstriaOracleCallerRaw struct {
	Contract *AstriaOracleCaller // Generic read-only contract binding to access the raw methods on
}

// AstriaOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AstriaOracleTransactorRaw struct {
	Contract *AstriaOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAstriaOracle creates a new instance of AstriaOracle, bound to a specific deployed contract.
func NewAstriaOracle(address common.Address, backend bind.ContractBackend) (*AstriaOracle, error) {
	contract, err := bindAstriaOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AstriaOracle{AstriaOracleCaller: AstriaOracleCaller{contract: contract}, AstriaOracleTransactor: AstriaOracleTransactor{contract: contract}, AstriaOracleFilterer: AstriaOracleFilterer{contract: contract}}, nil
}

// NewAstriaOracleCaller creates a new read-only instance of AstriaOracle, bound to a specific deployed contract.
func NewAstriaOracleCaller(address common.Address, caller bind.ContractCaller) (*AstriaOracleCaller, error) {
	contract, err := bindAstriaOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AstriaOracleCaller{contract: contract}, nil
}

// NewAstriaOracleTransactor creates a new write-only instance of AstriaOracle, bound to a specific deployed contract.
func NewAstriaOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*AstriaOracleTransactor, error) {
	contract, err := bindAstriaOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AstriaOracleTransactor{contract: contract}, nil
}

// NewAstriaOracleFilterer creates a new log filterer instance of AstriaOracle, bound to a specific deployed contract.
func NewAstriaOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*AstriaOracleFilterer, error) {
	contract, err := bindAstriaOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AstriaOracleFilterer{contract: contract}, nil
}

// bindAstriaOracle binds a generic wrapper to an already deployed contract.
func bindAstriaOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AstriaOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AstriaOracle *AstriaOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AstriaOracle.Contract.AstriaOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AstriaOracle *AstriaOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AstriaOracle.Contract.AstriaOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AstriaOracle *AstriaOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AstriaOracle.Contract.AstriaOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AstriaOracle *AstriaOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AstriaOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AstriaOracle *AstriaOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AstriaOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AstriaOracle *AstriaOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AstriaOracle.Contract.contract.Transact(opts, method, params...)
}

// ORACLE is a free data retrieval call binding the contract method 0x38013f02.
//
// Solidity: function ORACLE() view returns(address)
func (_AstriaOracle *AstriaOracleCaller) ORACLE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AstriaOracle.contract.Call(opts, &out, "ORACLE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ORACLE is a free data retrieval call binding the contract method 0x38013f02.
//
// Solidity: function ORACLE() view returns(address)
func (_AstriaOracle *AstriaOracleSession) ORACLE() (common.Address, error) {
	return _AstriaOracle.Contract.ORACLE(&_AstriaOracle.CallOpts)
}

// ORACLE is a free data retrieval call binding the contract method 0x38013f02.
//
// Solidity: function ORACLE() view returns(address)
func (_AstriaOracle *AstriaOracleCallerSession) ORACLE() (common.Address, error) {
	return _AstriaOracle.Contract.ORACLE(&_AstriaOracle.CallOpts)
}

// CurrencyPairInfo is a free data retrieval call binding the contract method 0x859bd5b5.
//
// Solidity: function currencyPairInfo(bytes32 ) view returns(bool initialized, uint8 decimals)
func (_AstriaOracle *AstriaOracleCaller) CurrencyPairInfo(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Initialized bool
	Decimals    uint8
}, error) {
	var out []interface{}
	err := _AstriaOracle.contract.Call(opts, &out, "currencyPairInfo", arg0)

	outstruct := new(struct {
		Initialized bool
		Decimals    uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Initialized = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Decimals = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// CurrencyPairInfo is a free data retrieval call binding the contract method 0x859bd5b5.
//
// Solidity: function currencyPairInfo(bytes32 ) view returns(bool initialized, uint8 decimals)
func (_AstriaOracle *AstriaOracleSession) CurrencyPairInfo(arg0 [32]byte) (struct {
	Initialized bool
	Decimals    uint8
}, error) {
	return _AstriaOracle.Contract.CurrencyPairInfo(&_AstriaOracle.CallOpts, arg0)
}

// CurrencyPairInfo is a free data retrieval call binding the contract method 0x859bd5b5.
//
// Solidity: function currencyPairInfo(bytes32 ) view returns(bool initialized, uint8 decimals)
func (_AstriaOracle *AstriaOracleCallerSession) CurrencyPairInfo(arg0 [32]byte) (struct {
	Initialized bool
	Decimals    uint8
}, error) {
	return _AstriaOracle.Contract.CurrencyPairInfo(&_AstriaOracle.CallOpts, arg0)
}

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_AstriaOracle *AstriaOracleCaller) LatestBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AstriaOracle.contract.Call(opts, &out, "latestBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_AstriaOracle *AstriaOracleSession) LatestBlockNumber() (*big.Int, error) {
	return _AstriaOracle.Contract.LatestBlockNumber(&_AstriaOracle.CallOpts)
}

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_AstriaOracle *AstriaOracleCallerSession) LatestBlockNumber() (*big.Int, error) {
	return _AstriaOracle.Contract.LatestBlockNumber(&_AstriaOracle.CallOpts)
}

// PriceData is a free data retrieval call binding the contract method 0xdad84e9d.
//
// Solidity: function priceData(uint256 , bytes32 ) view returns(uint128 price, uint256 timestamp)
func (_AstriaOracle *AstriaOracleCaller) PriceData(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _AstriaOracle.contract.Call(opts, &out, "priceData", arg0, arg1)

	outstruct := new(struct {
		Price     *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PriceData is a free data retrieval call binding the contract method 0xdad84e9d.
//
// Solidity: function priceData(uint256 , bytes32 ) view returns(uint128 price, uint256 timestamp)
func (_AstriaOracle *AstriaOracleSession) PriceData(arg0 *big.Int, arg1 [32]byte) (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	return _AstriaOracle.Contract.PriceData(&_AstriaOracle.CallOpts, arg0, arg1)
}

// PriceData is a free data retrieval call binding the contract method 0xdad84e9d.
//
// Solidity: function priceData(uint256 , bytes32 ) view returns(uint128 price, uint256 timestamp)
func (_AstriaOracle *AstriaOracleCallerSession) PriceData(arg0 *big.Int, arg1 [32]byte) (struct {
	Price     *big.Int
	Timestamp *big.Int
}, error) {
	return _AstriaOracle.Contract.PriceData(&_AstriaOracle.CallOpts, arg0, arg1)
}

// InitializeCurrencyPair is a paid mutator transaction binding the contract method 0x48832f3c.
//
// Solidity: function initializeCurrencyPair(bytes32 _currencyPair, uint8 _decimals) returns()
func (_AstriaOracle *AstriaOracleTransactor) InitializeCurrencyPair(opts *bind.TransactOpts, _currencyPair [32]byte, _decimals uint8) (*types.Transaction, error) {
	return _AstriaOracle.contract.Transact(opts, "initializeCurrencyPair", _currencyPair, _decimals)
}

// InitializeCurrencyPair is a paid mutator transaction binding the contract method 0x48832f3c.
//
// Solidity: function initializeCurrencyPair(bytes32 _currencyPair, uint8 _decimals) returns()
func (_AstriaOracle *AstriaOracleSession) InitializeCurrencyPair(_currencyPair [32]byte, _decimals uint8) (*types.Transaction, error) {
	return _AstriaOracle.Contract.InitializeCurrencyPair(&_AstriaOracle.TransactOpts, _currencyPair, _decimals)
}

// InitializeCurrencyPair is a paid mutator transaction binding the contract method 0x48832f3c.
//
// Solidity: function initializeCurrencyPair(bytes32 _currencyPair, uint8 _decimals) returns()
func (_AstriaOracle *AstriaOracleTransactorSession) InitializeCurrencyPair(_currencyPair [32]byte, _decimals uint8) (*types.Transaction, error) {
	return _AstriaOracle.Contract.InitializeCurrencyPair(&_AstriaOracle.TransactOpts, _currencyPair, _decimals)
}

// UpdatePriceData is a paid mutator transaction binding the contract method 0x3595f691.
//
// Solidity: function updatePriceData(bytes32[] _currencyPairs, uint128[] _prices) returns()
func (_AstriaOracle *AstriaOracleTransactor) UpdatePriceData(opts *bind.TransactOpts, _currencyPairs [][32]byte, _prices []*big.Int) (*types.Transaction, error) {
	return _AstriaOracle.contract.Transact(opts, "updatePriceData", _currencyPairs, _prices)
}

// UpdatePriceData is a paid mutator transaction binding the contract method 0x3595f691.
//
// Solidity: function updatePriceData(bytes32[] _currencyPairs, uint128[] _prices) returns()
func (_AstriaOracle *AstriaOracleSession) UpdatePriceData(_currencyPairs [][32]byte, _prices []*big.Int) (*types.Transaction, error) {
	return _AstriaOracle.Contract.UpdatePriceData(&_AstriaOracle.TransactOpts, _currencyPairs, _prices)
}

// UpdatePriceData is a paid mutator transaction binding the contract method 0x3595f691.
//
// Solidity: function updatePriceData(bytes32[] _currencyPairs, uint128[] _prices) returns()
func (_AstriaOracle *AstriaOracleTransactorSession) UpdatePriceData(_currencyPairs [][32]byte, _prices []*big.Int) (*types.Transaction, error) {
	return _AstriaOracle.Contract.UpdatePriceData(&_AstriaOracle.TransactOpts, _currencyPairs, _prices)
}
