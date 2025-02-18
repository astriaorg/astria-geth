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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_requireCurrencyPairAuthorization\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"currencyPair\",\"type\":\"bytes32\"}],\"name\":\"UnauthorizedCurrencyPair\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"UninitializedCurrencyPair\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"currencyPair\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"CurrencyPairInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"currencyPair\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"price\",\"type\":\"uint128\"}],\"name\":\"PriceDataUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ORACLE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_currencyPair\",\"type\":\"bytes32\"}],\"name\":\"authorizeCurrencyPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"authorizedCurrencyPairs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"currencyPairInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_currencyPair\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"name\":\"initializeCurrencyPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"priceData\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requireCurrencyPairAuthorization\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_currencyPair\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"_price\",\"type\":\"uint128\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_currencyPairs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint128[]\",\"name\":\"_prices\",\"type\":\"uint128[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_requireCurrencyPairAuthorization\",\"type\":\"bool\"}],\"name\":\"setRequireCurrencyPairAuthorization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f80fd5b50604051610caf380380610caf83398101604081905261002e916100da565b338061005357604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b61005c8161008b565b506001600160a01b039091166080525f8054911515600160a01b0260ff60a01b19909216919091179055610121565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f80604083850312156100eb575f80fd5b82516001600160a01b0381168114610101575f80fd5b60208401519092508015158114610116575f80fd5b809150509250929050565b608051610b6261014d5f395f818160ee015281816102c0015281816103d7015261051c0152610b625ff3fe608060405234801561000f575f80fd5b50600436106100e5575f3560e01c80638da5cb5b11610088578063abcdd06911610063578063abcdd06914610211578063dad84e9d14610233578063f2fde38b1461028f578063f6dad934146102a2575f80fd5b80638da5cb5b146101cb57806395009a53146101db5780639c97f6d8146101ee575f80fd5b806348832f3c116100c357806348832f3c14610159578063715018a61461016c578063859bd5b514610174578063888810c9146101b8575f80fd5b806338013f02146100e95780633a6a71621461012d5780634599c78814610142575b5f80fd5b6101107f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b61014061013b366004610846565b6102b5565b005b61014b60045481565b604051908152602001610124565b610140610167366004610870565b6103cc565b6101406104d9565b61019f6101823660046108a3565b60026020525f908152604090205460ff8082169161010090041682565b60408051921515835260ff909116602083015201610124565b6101406101c63660046108ba565b6104ec565b5f546001600160a01b0316610110565b6101406101e93660046109b7565b610511565b5f5461020190600160a01b900460ff1681565b6040519015158152602001610124565b61020161021f3660046108a3565b60016020525f908152604090205460ff1681565b610270610241366004610a6a565b600360209081525f9283526040808420909152908252902080546001909101546001600160801b039091169082565b604080516001600160801b039093168352602083019190915201610124565b61014061029d366004610a8a565b61074e565b6101406102b03660046108a3565b61078b565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146103065760405162461bcd60e51b81526004016102fd90610ab0565b60405180910390fd5b436004555f8281526002602052604090205460ff1661033a5760405163183c4ba760e01b81525f60048201526024016102fd565b6040805180820182526001600160801b038381168083524260208085019182526004545f908152600382528681208982528252869020945185546001600160801b0319169416939093178455516001909301929092558251858152908101919091527fd616ae5f8d378c1264fdbbbc72af91e16e3645564d7eae37e267ef1c67bf5cee91015b60405180910390a15050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146104145760405162461bcd60e51b81526004016102fd90610ab0565b5f54600160a01b900460ff16801561043a57505f8281526001602052604090205460ff16155b1561045b5760405163401e592360e01b8152600481018390526024016102fd565b6040805180820182526001815260ff83811660208084018281525f8881526002835286902094518554915161ffff1990921690151561ff001916176101009190941602929092179092558251858152908101919091527f675b5c62c7826a107baf315a10339c41c59f32f58ac6431f359e9ac89c64a01b91016103c0565b6104e16107b0565b6104ea5f6107dc565b565b6104f46107b0565b5f8054911515600160a01b0260ff60a01b19909216919091179055565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146105595760405162461bcd60e51b81526004016102fd90610ab0565b80518251146105ba5760405162461bcd60e51b815260206004820152602760248201527f63757272656e6379207061697220616e64207072696365206c656e677468206d6044820152660d2e6dac2e8c6d60cb1b60648201526084016102fd565b436004555f5b82518110156107495760025f8483815181106105de576105de610af4565b60209081029190910181015182528101919091526040015f205460ff1661061b5760405163183c4ba760e01b8152600481018290526024016102fd565b604051806040016040528083838151811061063857610638610af4565b60200260200101516001600160801b031681526020014281525060035f60045481526020019081526020015f205f85848151811061067857610678610af4565b6020908102919091018101518252818101929092526040015f20825181546001600160801b0319166001600160801b0390911617815591015160019091015582517fd616ae5f8d378c1264fdbbbc72af91e16e3645564d7eae37e267ef1c67bf5cee908490839081106106ed576106ed610af4565b602002602001015183838151811061070757610707610af4565b602002602001015160405161072f9291909182526001600160801b0316602082015260400190565b60405180910390a18061074181610b08565b9150506105c0565b505050565b6107566107b0565b6001600160a01b03811661077f57604051631e4fbdf760e01b81525f60048201526024016102fd565b610788816107dc565b50565b6107936107b0565b5f908152600160208190526040909120805460ff19169091179055565b5f546001600160a01b031633146104ea5760405163118cdaa760e01b81523360048201526024016102fd565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80356001600160801b0381168114610841575f80fd5b919050565b5f8060408385031215610857575f80fd5b823591506108676020840161082b565b90509250929050565b5f8060408385031215610881575f80fd5b82359150602083013560ff81168114610898575f80fd5b809150509250929050565b5f602082840312156108b3575f80fd5b5035919050565b5f602082840312156108ca575f80fd5b813580151581146108d9575f80fd5b9392505050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561091d5761091d6108e0565b604052919050565b5f67ffffffffffffffff82111561093e5761093e6108e0565b5060051b60200190565b5f82601f830112610957575f80fd5b8135602061096c61096783610925565b6108f4565b82815260059290921b8401810191818101908684111561098a575f80fd5b8286015b848110156109ac5761099f8161082b565b835291830191830161098e565b509695505050505050565b5f80604083850312156109c8575f80fd5b823567ffffffffffffffff808211156109df575f80fd5b818501915085601f8301126109f2575f80fd5b81356020610a0261096783610925565b82815260059290921b84018101918181019089841115610a20575f80fd5b948201945b83861015610a3e57853582529482019490820190610a25565b96505086013592505080821115610a53575f80fd5b50610a6085828601610948565b9150509250929050565b5f8060408385031215610a7b575f80fd5b50508035926020909101359150565b5f60208284031215610a9a575f80fd5b81356001600160a01b03811681146108d9575f80fd5b60208082526024908201527f4173747269614f7261636c653a206f6e6c79206f7261636c652063616e2075706040820152636461746560e01b606082015260800190565b634e487b7160e01b5f52603260045260245ffd5b5f60018201610b2557634e487b7160e01b5f52601160045260245ffd5b506001019056fea26469706673582212206447e8e6f4c0e75dc74fe9f33088774c0f21634d98ed787ad4d4638ecc965ff864736f6c63430008150033",
}

// AstriaOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use AstriaOracleMetaData.ABI instead.
var AstriaOracleABI = AstriaOracleMetaData.ABI

// AstriaOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AstriaOracleMetaData.Bin instead.
var AstriaOracleBin = AstriaOracleMetaData.Bin

// DeployAstriaOracle deploys a new Ethereum contract, binding an instance of AstriaOracle to it.
func DeployAstriaOracle(auth *bind.TransactOpts, backend bind.ContractBackend, _oracle common.Address, _requireCurrencyPairAuthorization bool) (common.Address, *types.Transaction, *AstriaOracle, error) {
	parsed, err := AstriaOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AstriaOracleBin), backend, _oracle, _requireCurrencyPairAuthorization)
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

// AuthorizedCurrencyPairs is a free data retrieval call binding the contract method 0xabcdd069.
//
// Solidity: function authorizedCurrencyPairs(bytes32 ) view returns(bool)
func (_AstriaOracle *AstriaOracleCaller) AuthorizedCurrencyPairs(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _AstriaOracle.contract.Call(opts, &out, "authorizedCurrencyPairs", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuthorizedCurrencyPairs is a free data retrieval call binding the contract method 0xabcdd069.
//
// Solidity: function authorizedCurrencyPairs(bytes32 ) view returns(bool)
func (_AstriaOracle *AstriaOracleSession) AuthorizedCurrencyPairs(arg0 [32]byte) (bool, error) {
	return _AstriaOracle.Contract.AuthorizedCurrencyPairs(&_AstriaOracle.CallOpts, arg0)
}

// AuthorizedCurrencyPairs is a free data retrieval call binding the contract method 0xabcdd069.
//
// Solidity: function authorizedCurrencyPairs(bytes32 ) view returns(bool)
func (_AstriaOracle *AstriaOracleCallerSession) AuthorizedCurrencyPairs(arg0 [32]byte) (bool, error) {
	return _AstriaOracle.Contract.AuthorizedCurrencyPairs(&_AstriaOracle.CallOpts, arg0)
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AstriaOracle *AstriaOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AstriaOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AstriaOracle *AstriaOracleSession) Owner() (common.Address, error) {
	return _AstriaOracle.Contract.Owner(&_AstriaOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AstriaOracle *AstriaOracleCallerSession) Owner() (common.Address, error) {
	return _AstriaOracle.Contract.Owner(&_AstriaOracle.CallOpts)
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

// RequireCurrencyPairAuthorization is a free data retrieval call binding the contract method 0x9c97f6d8.
//
// Solidity: function requireCurrencyPairAuthorization() view returns(bool)
func (_AstriaOracle *AstriaOracleCaller) RequireCurrencyPairAuthorization(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AstriaOracle.contract.Call(opts, &out, "requireCurrencyPairAuthorization")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequireCurrencyPairAuthorization is a free data retrieval call binding the contract method 0x9c97f6d8.
//
// Solidity: function requireCurrencyPairAuthorization() view returns(bool)
func (_AstriaOracle *AstriaOracleSession) RequireCurrencyPairAuthorization() (bool, error) {
	return _AstriaOracle.Contract.RequireCurrencyPairAuthorization(&_AstriaOracle.CallOpts)
}

// RequireCurrencyPairAuthorization is a free data retrieval call binding the contract method 0x9c97f6d8.
//
// Solidity: function requireCurrencyPairAuthorization() view returns(bool)
func (_AstriaOracle *AstriaOracleCallerSession) RequireCurrencyPairAuthorization() (bool, error) {
	return _AstriaOracle.Contract.RequireCurrencyPairAuthorization(&_AstriaOracle.CallOpts)
}

// AuthorizeCurrencyPair is a paid mutator transaction binding the contract method 0xf6dad934.
//
// Solidity: function authorizeCurrencyPair(bytes32 _currencyPair) returns()
func (_AstriaOracle *AstriaOracleTransactor) AuthorizeCurrencyPair(opts *bind.TransactOpts, _currencyPair [32]byte) (*types.Transaction, error) {
	return _AstriaOracle.contract.Transact(opts, "authorizeCurrencyPair", _currencyPair)
}

// AuthorizeCurrencyPair is a paid mutator transaction binding the contract method 0xf6dad934.
//
// Solidity: function authorizeCurrencyPair(bytes32 _currencyPair) returns()
func (_AstriaOracle *AstriaOracleSession) AuthorizeCurrencyPair(_currencyPair [32]byte) (*types.Transaction, error) {
	return _AstriaOracle.Contract.AuthorizeCurrencyPair(&_AstriaOracle.TransactOpts, _currencyPair)
}

// AuthorizeCurrencyPair is a paid mutator transaction binding the contract method 0xf6dad934.
//
// Solidity: function authorizeCurrencyPair(bytes32 _currencyPair) returns()
func (_AstriaOracle *AstriaOracleTransactorSession) AuthorizeCurrencyPair(_currencyPair [32]byte) (*types.Transaction, error) {
	return _AstriaOracle.Contract.AuthorizeCurrencyPair(&_AstriaOracle.TransactOpts, _currencyPair)
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

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AstriaOracle *AstriaOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AstriaOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AstriaOracle *AstriaOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _AstriaOracle.Contract.RenounceOwnership(&_AstriaOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AstriaOracle *AstriaOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AstriaOracle.Contract.RenounceOwnership(&_AstriaOracle.TransactOpts)
}

// SetPrice is a paid mutator transaction binding the contract method 0x3a6a7162.
//
// Solidity: function setPrice(bytes32 _currencyPair, uint128 _price) returns()
func (_AstriaOracle *AstriaOracleTransactor) SetPrice(opts *bind.TransactOpts, _currencyPair [32]byte, _price *big.Int) (*types.Transaction, error) {
	return _AstriaOracle.contract.Transact(opts, "setPrice", _currencyPair, _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x3a6a7162.
//
// Solidity: function setPrice(bytes32 _currencyPair, uint128 _price) returns()
func (_AstriaOracle *AstriaOracleSession) SetPrice(_currencyPair [32]byte, _price *big.Int) (*types.Transaction, error) {
	return _AstriaOracle.Contract.SetPrice(&_AstriaOracle.TransactOpts, _currencyPair, _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x3a6a7162.
//
// Solidity: function setPrice(bytes32 _currencyPair, uint128 _price) returns()
func (_AstriaOracle *AstriaOracleTransactorSession) SetPrice(_currencyPair [32]byte, _price *big.Int) (*types.Transaction, error) {
	return _AstriaOracle.Contract.SetPrice(&_AstriaOracle.TransactOpts, _currencyPair, _price)
}

// SetPrices is a paid mutator transaction binding the contract method 0x95009a53.
//
// Solidity: function setPrices(bytes32[] _currencyPairs, uint128[] _prices) returns()
func (_AstriaOracle *AstriaOracleTransactor) SetPrices(opts *bind.TransactOpts, _currencyPairs [][32]byte, _prices []*big.Int) (*types.Transaction, error) {
	return _AstriaOracle.contract.Transact(opts, "setPrices", _currencyPairs, _prices)
}

// SetPrices is a paid mutator transaction binding the contract method 0x95009a53.
//
// Solidity: function setPrices(bytes32[] _currencyPairs, uint128[] _prices) returns()
func (_AstriaOracle *AstriaOracleSession) SetPrices(_currencyPairs [][32]byte, _prices []*big.Int) (*types.Transaction, error) {
	return _AstriaOracle.Contract.SetPrices(&_AstriaOracle.TransactOpts, _currencyPairs, _prices)
}

// SetPrices is a paid mutator transaction binding the contract method 0x95009a53.
//
// Solidity: function setPrices(bytes32[] _currencyPairs, uint128[] _prices) returns()
func (_AstriaOracle *AstriaOracleTransactorSession) SetPrices(_currencyPairs [][32]byte, _prices []*big.Int) (*types.Transaction, error) {
	return _AstriaOracle.Contract.SetPrices(&_AstriaOracle.TransactOpts, _currencyPairs, _prices)
}

// SetRequireCurrencyPairAuthorization is a paid mutator transaction binding the contract method 0x888810c9.
//
// Solidity: function setRequireCurrencyPairAuthorization(bool _requireCurrencyPairAuthorization) returns()
func (_AstriaOracle *AstriaOracleTransactor) SetRequireCurrencyPairAuthorization(opts *bind.TransactOpts, _requireCurrencyPairAuthorization bool) (*types.Transaction, error) {
	return _AstriaOracle.contract.Transact(opts, "setRequireCurrencyPairAuthorization", _requireCurrencyPairAuthorization)
}

// SetRequireCurrencyPairAuthorization is a paid mutator transaction binding the contract method 0x888810c9.
//
// Solidity: function setRequireCurrencyPairAuthorization(bool _requireCurrencyPairAuthorization) returns()
func (_AstriaOracle *AstriaOracleSession) SetRequireCurrencyPairAuthorization(_requireCurrencyPairAuthorization bool) (*types.Transaction, error) {
	return _AstriaOracle.Contract.SetRequireCurrencyPairAuthorization(&_AstriaOracle.TransactOpts, _requireCurrencyPairAuthorization)
}

// SetRequireCurrencyPairAuthorization is a paid mutator transaction binding the contract method 0x888810c9.
//
// Solidity: function setRequireCurrencyPairAuthorization(bool _requireCurrencyPairAuthorization) returns()
func (_AstriaOracle *AstriaOracleTransactorSession) SetRequireCurrencyPairAuthorization(_requireCurrencyPairAuthorization bool) (*types.Transaction, error) {
	return _AstriaOracle.Contract.SetRequireCurrencyPairAuthorization(&_AstriaOracle.TransactOpts, _requireCurrencyPairAuthorization)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AstriaOracle *AstriaOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AstriaOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AstriaOracle *AstriaOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AstriaOracle.Contract.TransferOwnership(&_AstriaOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AstriaOracle *AstriaOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AstriaOracle.Contract.TransferOwnership(&_AstriaOracle.TransactOpts, newOwner)
}

// AstriaOracleCurrencyPairInitializedIterator is returned from FilterCurrencyPairInitialized and is used to iterate over the raw logs and unpacked data for CurrencyPairInitialized events raised by the AstriaOracle contract.
type AstriaOracleCurrencyPairInitializedIterator struct {
	Event *AstriaOracleCurrencyPairInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AstriaOracleCurrencyPairInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AstriaOracleCurrencyPairInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AstriaOracleCurrencyPairInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AstriaOracleCurrencyPairInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AstriaOracleCurrencyPairInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AstriaOracleCurrencyPairInitialized represents a CurrencyPairInitialized event raised by the AstriaOracle contract.
type AstriaOracleCurrencyPairInitialized struct {
	CurrencyPair [32]byte
	Decimals     uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCurrencyPairInitialized is a free log retrieval operation binding the contract event 0x675b5c62c7826a107baf315a10339c41c59f32f58ac6431f359e9ac89c64a01b.
//
// Solidity: event CurrencyPairInitialized(bytes32 currencyPair, uint8 decimals)
func (_AstriaOracle *AstriaOracleFilterer) FilterCurrencyPairInitialized(opts *bind.FilterOpts) (*AstriaOracleCurrencyPairInitializedIterator, error) {

	logs, sub, err := _AstriaOracle.contract.FilterLogs(opts, "CurrencyPairInitialized")
	if err != nil {
		return nil, err
	}
	return &AstriaOracleCurrencyPairInitializedIterator{contract: _AstriaOracle.contract, event: "CurrencyPairInitialized", logs: logs, sub: sub}, nil
}

// WatchCurrencyPairInitialized is a free log subscription operation binding the contract event 0x675b5c62c7826a107baf315a10339c41c59f32f58ac6431f359e9ac89c64a01b.
//
// Solidity: event CurrencyPairInitialized(bytes32 currencyPair, uint8 decimals)
func (_AstriaOracle *AstriaOracleFilterer) WatchCurrencyPairInitialized(opts *bind.WatchOpts, sink chan<- *AstriaOracleCurrencyPairInitialized) (event.Subscription, error) {

	logs, sub, err := _AstriaOracle.contract.WatchLogs(opts, "CurrencyPairInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AstriaOracleCurrencyPairInitialized)
				if err := _AstriaOracle.contract.UnpackLog(event, "CurrencyPairInitialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCurrencyPairInitialized is a log parse operation binding the contract event 0x675b5c62c7826a107baf315a10339c41c59f32f58ac6431f359e9ac89c64a01b.
//
// Solidity: event CurrencyPairInitialized(bytes32 currencyPair, uint8 decimals)
func (_AstriaOracle *AstriaOracleFilterer) ParseCurrencyPairInitialized(log types.Log) (*AstriaOracleCurrencyPairInitialized, error) {
	event := new(AstriaOracleCurrencyPairInitialized)
	if err := _AstriaOracle.contract.UnpackLog(event, "CurrencyPairInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AstriaOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AstriaOracle contract.
type AstriaOracleOwnershipTransferredIterator struct {
	Event *AstriaOracleOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AstriaOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AstriaOracleOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AstriaOracleOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AstriaOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AstriaOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AstriaOracleOwnershipTransferred represents a OwnershipTransferred event raised by the AstriaOracle contract.
type AstriaOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AstriaOracle *AstriaOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AstriaOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AstriaOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AstriaOracleOwnershipTransferredIterator{contract: _AstriaOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AstriaOracle *AstriaOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AstriaOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AstriaOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AstriaOracleOwnershipTransferred)
				if err := _AstriaOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AstriaOracle *AstriaOracleFilterer) ParseOwnershipTransferred(log types.Log) (*AstriaOracleOwnershipTransferred, error) {
	event := new(AstriaOracleOwnershipTransferred)
	if err := _AstriaOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AstriaOraclePriceDataUpdatedIterator is returned from FilterPriceDataUpdated and is used to iterate over the raw logs and unpacked data for PriceDataUpdated events raised by the AstriaOracle contract.
type AstriaOraclePriceDataUpdatedIterator struct {
	Event *AstriaOraclePriceDataUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AstriaOraclePriceDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AstriaOraclePriceDataUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AstriaOraclePriceDataUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AstriaOraclePriceDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AstriaOraclePriceDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AstriaOraclePriceDataUpdated represents a PriceDataUpdated event raised by the AstriaOracle contract.
type AstriaOraclePriceDataUpdated struct {
	CurrencyPair [32]byte
	Price        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPriceDataUpdated is a free log retrieval operation binding the contract event 0xd616ae5f8d378c1264fdbbbc72af91e16e3645564d7eae37e267ef1c67bf5cee.
//
// Solidity: event PriceDataUpdated(bytes32 currencyPair, uint128 price)
func (_AstriaOracle *AstriaOracleFilterer) FilterPriceDataUpdated(opts *bind.FilterOpts) (*AstriaOraclePriceDataUpdatedIterator, error) {

	logs, sub, err := _AstriaOracle.contract.FilterLogs(opts, "PriceDataUpdated")
	if err != nil {
		return nil, err
	}
	return &AstriaOraclePriceDataUpdatedIterator{contract: _AstriaOracle.contract, event: "PriceDataUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceDataUpdated is a free log subscription operation binding the contract event 0xd616ae5f8d378c1264fdbbbc72af91e16e3645564d7eae37e267ef1c67bf5cee.
//
// Solidity: event PriceDataUpdated(bytes32 currencyPair, uint128 price)
func (_AstriaOracle *AstriaOracleFilterer) WatchPriceDataUpdated(opts *bind.WatchOpts, sink chan<- *AstriaOraclePriceDataUpdated) (event.Subscription, error) {

	logs, sub, err := _AstriaOracle.contract.WatchLogs(opts, "PriceDataUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AstriaOraclePriceDataUpdated)
				if err := _AstriaOracle.contract.UnpackLog(event, "PriceDataUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePriceDataUpdated is a log parse operation binding the contract event 0xd616ae5f8d378c1264fdbbbc72af91e16e3645564d7eae37e267ef1c67bf5cee.
//
// Solidity: event PriceDataUpdated(bytes32 currencyPair, uint128 price)
func (_AstriaOracle *AstriaOracleFilterer) ParsePriceDataUpdated(log types.Log) (*AstriaOraclePriceDataUpdated, error) {
	event := new(AstriaOraclePriceDataUpdated)
	if err := _AstriaOracle.contract.UnpackLog(event, "PriceDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
