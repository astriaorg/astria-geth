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

// AstriaMintableERC20MetaData contains all meta data concerning the AstriaMintableERC20 contract.
var AstriaMintableERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BRIDGE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801562000010575f80fd5b5060405162000c9138038062000c9183398101604081905262000033916200012a565b818160036200004383826200023a565b5060046200005282826200023a565b5050506001600160a01b0390921660805250620003029050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f83011262000090575f80fd5b81516001600160401b0380821115620000ad57620000ad6200006c565b604051601f8301601f19908116603f01168101908282118183101715620000d857620000d86200006c565b81604052838152602092508683858801011115620000f4575f80fd5b5f91505b83821015620001175785820183015181830184015290820190620000f8565b5f93810190920192909252949350505050565b5f805f606084860312156200013d575f80fd5b83516001600160a01b038116811462000154575f80fd5b60208501519093506001600160401b038082111562000171575f80fd5b6200017f8783880162000080565b9350604086015191508082111562000195575f80fd5b50620001a48682870162000080565b9150509250925092565b600181811c90821680620001c357607f821691505b602082108103620001e257634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111562000235575f81815260208120601f850160051c81016020861015620002105750805b601f850160051c820191505b8181101562000231578281556001016200021c565b5050505b505050565b81516001600160401b038111156200025657620002566200006c565b6200026e81620002678454620001ae565b84620001e8565b602080601f831160018114620002a4575f84156200028c5750858301515b5f19600386901b1c1916600185901b17855562000231565b5f85815260208120601f198616915b82811015620002d457888601518255948401946001909101908401620002b3565b5085821015620002f257878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b608051610968620003295f395f81816101d2015281816102e3015261039401526109685ff3fe608060405234801561000f575f80fd5b50600436106100b1575f3560e01c806370a082311161006e57806370a082311461013f57806395d89b41146101675780639dc29fac1461016f578063a9059cbb14610182578063dd62ed3e14610195578063ee9a31a2146101cd575f80fd5b806306fdde03146100b5578063095ea7b3146100d357806318160ddd146100f657806323b872dd14610108578063313ce5671461011b57806340c10f191461012a575b5f80fd5b6100bd61020c565b6040516100ca9190610771565b60405180910390f35b6100e66100e13660046107d7565b61029c565b60405190151581526020016100ca565b6002545b6040519081526020016100ca565b6100e66101163660046107ff565b6102b5565b604051601281526020016100ca565b61013d6101383660046107d7565b6102d8565b005b6100fa61014d366004610838565b6001600160a01b03165f9081526020819052604090205490565b6100bd61037a565b61013d61017d3660046107d7565b610389565b6100e66101903660046107d7565b610416565b6100fa6101a3366004610858565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b6101f47f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100ca565b60606003805461021b90610889565b80601f016020809104026020016040519081016040528092919081815260200182805461024790610889565b80156102925780601f1061026957610100808354040283529160200191610292565b820191905f5260205f20905b81548152906001019060200180831161027557829003601f168201915b5050505050905090565b5f336102a9818585610423565b60019150505b92915050565b5f336102c2858285610435565b6102cd8585856104b0565b506001949350505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146103295760405162461bcd60e51b8152600401610320906108c1565b60405180910390fd5b610333828261050d565b816001600160a01b03167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968858260405161036e91815260200190565b60405180910390a25050565b60606004805461021b90610889565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146103d15760405162461bcd60e51b8152600401610320906108c1565b6103db8282610545565b816001600160a01b03167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca58260405161036e91815260200190565b5f336102a98185856104b0565b6104308383836001610579565b505050565b6001600160a01b038381165f908152600160209081526040808320938616835292905220545f1981146104aa578181101561049c57604051637dc7a0d960e11b81526001600160a01b03841660048201526024810182905260448101839052606401610320565b6104aa84848484035f610579565b50505050565b6001600160a01b0383166104d957604051634b637e8f60e11b81525f6004820152602401610320565b6001600160a01b0382166105025760405163ec442f0560e01b81525f6004820152602401610320565b61043083838361064b565b6001600160a01b0382166105365760405163ec442f0560e01b81525f6004820152602401610320565b6105415f838361064b565b5050565b6001600160a01b03821661056e57604051634b637e8f60e11b81525f6004820152602401610320565b610541825f8361064b565b6001600160a01b0384166105a25760405163e602df0560e01b81525f6004820152602401610320565b6001600160a01b0383166105cb57604051634a1406b160e11b81525f6004820152602401610320565b6001600160a01b038085165f90815260016020908152604080832093871683529290522082905580156104aa57826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161063d91815260200190565b60405180910390a350505050565b6001600160a01b038316610675578060025f82825461066a9190610913565b909155506106e59050565b6001600160a01b0383165f90815260208190526040902054818110156106c75760405163391434e360e21b81526001600160a01b03851660048201526024810182905260448101839052606401610320565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b0382166107015760028054829003905561071f565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161076491815260200190565b60405180910390a3505050565b5f6020808352835180828501525f5b8181101561079c57858101830151858201604001528201610780565b505f604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b03811681146107d2575f80fd5b919050565b5f80604083850312156107e8575f80fd5b6107f1836107bc565b946020939093013593505050565b5f805f60608486031215610811575f80fd5b61081a846107bc565b9250610828602085016107bc565b9150604084013590509250925092565b5f60208284031215610848575f80fd5b610851826107bc565b9392505050565b5f8060408385031215610869575f80fd5b610872836107bc565b9150610880602084016107bc565b90509250929050565b600181811c9082168061089d57607f821691505b6020821081036108bb57634e487b7160e01b5f52602260045260245ffd5b50919050565b60208082526032908201527f4173747269614d696e7461626c6545524332303a206f6e6c79206272696467656040820152711031b0b71036b4b73a1030b73210313ab93760711b606082015260800190565b808201808211156102af57634e487b7160e01b5f52601160045260245ffdfea264697066735822122014c66589442b8d35131eb4d14de96201d92d56c50cece5e9151d0ca979053a8864736f6c63430008150033",
}

// AstriaMintableERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use AstriaMintableERC20MetaData.ABI instead.
var AstriaMintableERC20ABI = AstriaMintableERC20MetaData.ABI

// AstriaMintableERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AstriaMintableERC20MetaData.Bin instead.
var AstriaMintableERC20Bin = AstriaMintableERC20MetaData.Bin

// DeployAstriaMintableERC20 deploys a new Ethereum contract, binding an instance of AstriaMintableERC20 to it.
func DeployAstriaMintableERC20(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address, _name string, _symbol string) (common.Address, *types.Transaction, *AstriaMintableERC20, error) {
	parsed, err := AstriaMintableERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AstriaMintableERC20Bin), backend, _bridge, _name, _symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AstriaMintableERC20{AstriaMintableERC20Caller: AstriaMintableERC20Caller{contract: contract}, AstriaMintableERC20Transactor: AstriaMintableERC20Transactor{contract: contract}, AstriaMintableERC20Filterer: AstriaMintableERC20Filterer{contract: contract}}, nil
}

// AstriaMintableERC20 is an auto generated Go binding around an Ethereum contract.
type AstriaMintableERC20 struct {
	AstriaMintableERC20Caller     // Read-only binding to the contract
	AstriaMintableERC20Transactor // Write-only binding to the contract
	AstriaMintableERC20Filterer   // Log filterer for contract events
}

// AstriaMintableERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type AstriaMintableERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AstriaMintableERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AstriaMintableERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AstriaMintableERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AstriaMintableERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AstriaMintableERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AstriaMintableERC20Session struct {
	Contract     *AstriaMintableERC20 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AstriaMintableERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AstriaMintableERC20CallerSession struct {
	Contract *AstriaMintableERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// AstriaMintableERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AstriaMintableERC20TransactorSession struct {
	Contract     *AstriaMintableERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// AstriaMintableERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type AstriaMintableERC20Raw struct {
	Contract *AstriaMintableERC20 // Generic contract binding to access the raw methods on
}

// AstriaMintableERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AstriaMintableERC20CallerRaw struct {
	Contract *AstriaMintableERC20Caller // Generic read-only contract binding to access the raw methods on
}

// AstriaMintableERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AstriaMintableERC20TransactorRaw struct {
	Contract *AstriaMintableERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAstriaMintableERC20 creates a new instance of AstriaMintableERC20, bound to a specific deployed contract.
func NewAstriaMintableERC20(address common.Address, backend bind.ContractBackend) (*AstriaMintableERC20, error) {
	contract, err := bindAstriaMintableERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AstriaMintableERC20{AstriaMintableERC20Caller: AstriaMintableERC20Caller{contract: contract}, AstriaMintableERC20Transactor: AstriaMintableERC20Transactor{contract: contract}, AstriaMintableERC20Filterer: AstriaMintableERC20Filterer{contract: contract}}, nil
}

// NewAstriaMintableERC20Caller creates a new read-only instance of AstriaMintableERC20, bound to a specific deployed contract.
func NewAstriaMintableERC20Caller(address common.Address, caller bind.ContractCaller) (*AstriaMintableERC20Caller, error) {
	contract, err := bindAstriaMintableERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AstriaMintableERC20Caller{contract: contract}, nil
}

// NewAstriaMintableERC20Transactor creates a new write-only instance of AstriaMintableERC20, bound to a specific deployed contract.
func NewAstriaMintableERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*AstriaMintableERC20Transactor, error) {
	contract, err := bindAstriaMintableERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AstriaMintableERC20Transactor{contract: contract}, nil
}

// NewAstriaMintableERC20Filterer creates a new log filterer instance of AstriaMintableERC20, bound to a specific deployed contract.
func NewAstriaMintableERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*AstriaMintableERC20Filterer, error) {
	contract, err := bindAstriaMintableERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AstriaMintableERC20Filterer{contract: contract}, nil
}

// bindAstriaMintableERC20 binds a generic wrapper to an already deployed contract.
func bindAstriaMintableERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AstriaMintableERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AstriaMintableERC20 *AstriaMintableERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AstriaMintableERC20.Contract.AstriaMintableERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AstriaMintableERC20 *AstriaMintableERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AstriaMintableERC20.Contract.AstriaMintableERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AstriaMintableERC20 *AstriaMintableERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AstriaMintableERC20.Contract.AstriaMintableERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AstriaMintableERC20 *AstriaMintableERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AstriaMintableERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AstriaMintableERC20 *AstriaMintableERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AstriaMintableERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AstriaMintableERC20 *AstriaMintableERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AstriaMintableERC20.Contract.contract.Transact(opts, method, params...)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_AstriaMintableERC20 *AstriaMintableERC20Caller) BRIDGE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AstriaMintableERC20.contract.Call(opts, &out, "BRIDGE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_AstriaMintableERC20 *AstriaMintableERC20Session) BRIDGE() (common.Address, error) {
	return _AstriaMintableERC20.Contract.BRIDGE(&_AstriaMintableERC20.CallOpts)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_AstriaMintableERC20 *AstriaMintableERC20CallerSession) BRIDGE() (common.Address, error) {
	return _AstriaMintableERC20.Contract.BRIDGE(&_AstriaMintableERC20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AstriaMintableERC20 *AstriaMintableERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AstriaMintableERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AstriaMintableERC20 *AstriaMintableERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AstriaMintableERC20.Contract.Allowance(&_AstriaMintableERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AstriaMintableERC20 *AstriaMintableERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AstriaMintableERC20.Contract.Allowance(&_AstriaMintableERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AstriaMintableERC20 *AstriaMintableERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AstriaMintableERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AstriaMintableERC20 *AstriaMintableERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _AstriaMintableERC20.Contract.BalanceOf(&_AstriaMintableERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AstriaMintableERC20 *AstriaMintableERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _AstriaMintableERC20.Contract.BalanceOf(&_AstriaMintableERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AstriaMintableERC20 *AstriaMintableERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AstriaMintableERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AstriaMintableERC20 *AstriaMintableERC20Session) Decimals() (uint8, error) {
	return _AstriaMintableERC20.Contract.Decimals(&_AstriaMintableERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AstriaMintableERC20 *AstriaMintableERC20CallerSession) Decimals() (uint8, error) {
	return _AstriaMintableERC20.Contract.Decimals(&_AstriaMintableERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AstriaMintableERC20 *AstriaMintableERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AstriaMintableERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AstriaMintableERC20 *AstriaMintableERC20Session) Name() (string, error) {
	return _AstriaMintableERC20.Contract.Name(&_AstriaMintableERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AstriaMintableERC20 *AstriaMintableERC20CallerSession) Name() (string, error) {
	return _AstriaMintableERC20.Contract.Name(&_AstriaMintableERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AstriaMintableERC20 *AstriaMintableERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AstriaMintableERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AstriaMintableERC20 *AstriaMintableERC20Session) Symbol() (string, error) {
	return _AstriaMintableERC20.Contract.Symbol(&_AstriaMintableERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AstriaMintableERC20 *AstriaMintableERC20CallerSession) Symbol() (string, error) {
	return _AstriaMintableERC20.Contract.Symbol(&_AstriaMintableERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AstriaMintableERC20 *AstriaMintableERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AstriaMintableERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AstriaMintableERC20 *AstriaMintableERC20Session) TotalSupply() (*big.Int, error) {
	return _AstriaMintableERC20.Contract.TotalSupply(&_AstriaMintableERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AstriaMintableERC20 *AstriaMintableERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _AstriaMintableERC20.Contract.TotalSupply(&_AstriaMintableERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AstriaMintableERC20 *AstriaMintableERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AstriaMintableERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AstriaMintableERC20 *AstriaMintableERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AstriaMintableERC20.Contract.Approve(&_AstriaMintableERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AstriaMintableERC20 *AstriaMintableERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AstriaMintableERC20.Contract.Approve(&_AstriaMintableERC20.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_AstriaMintableERC20 *AstriaMintableERC20Transactor) Burn(opts *bind.TransactOpts, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AstriaMintableERC20.contract.Transact(opts, "burn", _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_AstriaMintableERC20 *AstriaMintableERC20Session) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AstriaMintableERC20.Contract.Burn(&_AstriaMintableERC20.TransactOpts, _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_AstriaMintableERC20 *AstriaMintableERC20TransactorSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AstriaMintableERC20.Contract.Burn(&_AstriaMintableERC20.TransactOpts, _from, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_AstriaMintableERC20 *AstriaMintableERC20Transactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AstriaMintableERC20.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_AstriaMintableERC20 *AstriaMintableERC20Session) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AstriaMintableERC20.Contract.Mint(&_AstriaMintableERC20.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_AstriaMintableERC20 *AstriaMintableERC20TransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AstriaMintableERC20.Contract.Mint(&_AstriaMintableERC20.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_AstriaMintableERC20 *AstriaMintableERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AstriaMintableERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_AstriaMintableERC20 *AstriaMintableERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AstriaMintableERC20.Contract.Transfer(&_AstriaMintableERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_AstriaMintableERC20 *AstriaMintableERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AstriaMintableERC20.Contract.Transfer(&_AstriaMintableERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_AstriaMintableERC20 *AstriaMintableERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AstriaMintableERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_AstriaMintableERC20 *AstriaMintableERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AstriaMintableERC20.Contract.TransferFrom(&_AstriaMintableERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_AstriaMintableERC20 *AstriaMintableERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AstriaMintableERC20.Contract.TransferFrom(&_AstriaMintableERC20.TransactOpts, from, to, value)
}

// AstriaMintableERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the AstriaMintableERC20 contract.
type AstriaMintableERC20ApprovalIterator struct {
	Event *AstriaMintableERC20Approval // Event containing the contract specifics and raw log

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
func (it *AstriaMintableERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AstriaMintableERC20Approval)
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
		it.Event = new(AstriaMintableERC20Approval)
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
func (it *AstriaMintableERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AstriaMintableERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AstriaMintableERC20Approval represents a Approval event raised by the AstriaMintableERC20 contract.
type AstriaMintableERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AstriaMintableERC20 *AstriaMintableERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AstriaMintableERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AstriaMintableERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AstriaMintableERC20ApprovalIterator{contract: _AstriaMintableERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AstriaMintableERC20 *AstriaMintableERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AstriaMintableERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AstriaMintableERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AstriaMintableERC20Approval)
				if err := _AstriaMintableERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AstriaMintableERC20 *AstriaMintableERC20Filterer) ParseApproval(log types.Log) (*AstriaMintableERC20Approval, error) {
	event := new(AstriaMintableERC20Approval)
	if err := _AstriaMintableERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AstriaMintableERC20BurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the AstriaMintableERC20 contract.
type AstriaMintableERC20BurnIterator struct {
	Event *AstriaMintableERC20Burn // Event containing the contract specifics and raw log

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
func (it *AstriaMintableERC20BurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AstriaMintableERC20Burn)
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
		it.Event = new(AstriaMintableERC20Burn)
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
func (it *AstriaMintableERC20BurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AstriaMintableERC20BurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AstriaMintableERC20Burn represents a Burn event raised by the AstriaMintableERC20 contract.
type AstriaMintableERC20Burn struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed account, uint256 amount)
func (_AstriaMintableERC20 *AstriaMintableERC20Filterer) FilterBurn(opts *bind.FilterOpts, account []common.Address) (*AstriaMintableERC20BurnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AstriaMintableERC20.contract.FilterLogs(opts, "Burn", accountRule)
	if err != nil {
		return nil, err
	}
	return &AstriaMintableERC20BurnIterator{contract: _AstriaMintableERC20.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed account, uint256 amount)
func (_AstriaMintableERC20 *AstriaMintableERC20Filterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *AstriaMintableERC20Burn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AstriaMintableERC20.contract.WatchLogs(opts, "Burn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AstriaMintableERC20Burn)
				if err := _AstriaMintableERC20.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed account, uint256 amount)
func (_AstriaMintableERC20 *AstriaMintableERC20Filterer) ParseBurn(log types.Log) (*AstriaMintableERC20Burn, error) {
	event := new(AstriaMintableERC20Burn)
	if err := _AstriaMintableERC20.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AstriaMintableERC20MintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the AstriaMintableERC20 contract.
type AstriaMintableERC20MintIterator struct {
	Event *AstriaMintableERC20Mint // Event containing the contract specifics and raw log

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
func (it *AstriaMintableERC20MintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AstriaMintableERC20Mint)
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
		it.Event = new(AstriaMintableERC20Mint)
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
func (it *AstriaMintableERC20MintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AstriaMintableERC20MintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AstriaMintableERC20Mint represents a Mint event raised by the AstriaMintableERC20 contract.
type AstriaMintableERC20Mint struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed account, uint256 amount)
func (_AstriaMintableERC20 *AstriaMintableERC20Filterer) FilterMint(opts *bind.FilterOpts, account []common.Address) (*AstriaMintableERC20MintIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AstriaMintableERC20.contract.FilterLogs(opts, "Mint", accountRule)
	if err != nil {
		return nil, err
	}
	return &AstriaMintableERC20MintIterator{contract: _AstriaMintableERC20.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed account, uint256 amount)
func (_AstriaMintableERC20 *AstriaMintableERC20Filterer) WatchMint(opts *bind.WatchOpts, sink chan<- *AstriaMintableERC20Mint, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AstriaMintableERC20.contract.WatchLogs(opts, "Mint", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AstriaMintableERC20Mint)
				if err := _AstriaMintableERC20.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed account, uint256 amount)
func (_AstriaMintableERC20 *AstriaMintableERC20Filterer) ParseMint(log types.Log) (*AstriaMintableERC20Mint, error) {
	event := new(AstriaMintableERC20Mint)
	if err := _AstriaMintableERC20.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AstriaMintableERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the AstriaMintableERC20 contract.
type AstriaMintableERC20TransferIterator struct {
	Event *AstriaMintableERC20Transfer // Event containing the contract specifics and raw log

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
func (it *AstriaMintableERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AstriaMintableERC20Transfer)
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
		it.Event = new(AstriaMintableERC20Transfer)
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
func (it *AstriaMintableERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AstriaMintableERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AstriaMintableERC20Transfer represents a Transfer event raised by the AstriaMintableERC20 contract.
type AstriaMintableERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AstriaMintableERC20 *AstriaMintableERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AstriaMintableERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AstriaMintableERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AstriaMintableERC20TransferIterator{contract: _AstriaMintableERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AstriaMintableERC20 *AstriaMintableERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AstriaMintableERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AstriaMintableERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AstriaMintableERC20Transfer)
				if err := _AstriaMintableERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AstriaMintableERC20 *AstriaMintableERC20Filterer) ParseTransfer(log types.Log) (*AstriaMintableERC20Transfer, error) {
	event := new(AstriaMintableERC20Transfer)
	if err := _AstriaMintableERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
