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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destinationChainAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"Ics20Withdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationChainAddress\",\"type\":\"address\"}],\"name\":\"SequencerWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BRIDGE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_destinationChainAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_memo\",\"type\":\"string\"}],\"name\":\"withdrawToIbcChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_destinationChainAddress\",\"type\":\"address\"}],\"name\":\"withdrawToSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801562000010575f80fd5b5060405162000df038038062000df083398101604081905262000033916200012a565b818160036200004383826200023a565b5060046200005282826200023a565b5050506001600160a01b0390921660805250620003029050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f83011262000090575f80fd5b81516001600160401b0380821115620000ad57620000ad6200006c565b604051601f8301601f19908116603f01168101908282118183101715620000d857620000d86200006c565b81604052838152602092508683858801011115620000f4575f80fd5b5f91505b83821015620001175785820183015181830184015290820190620000f8565b5f93810190920192909252949350505050565b5f805f606084860312156200013d575f80fd5b83516001600160a01b038116811462000154575f80fd5b60208501519093506001600160401b038082111562000171575f80fd5b6200017f8783880162000080565b9350604086015191508082111562000195575f80fd5b50620001a48682870162000080565b9150509250925092565b600181811c90821680620001c357607f821691505b602082108103620001e257634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111562000235575f81815260208120601f850160051c81016020861015620002105750805b601f850160051c820191505b8181101562000231578281556001016200021c565b5050505b505050565b81516001600160401b038111156200025657620002566200006c565b6200026e81620002678454620001ae565b84620001e8565b602080601f831160018114620002a4575f84156200028c5750858301515b5f19600386901b1c1916600185901b17855562000231565b5f85815260208120601f198616915b82811015620002d457888601518255948401946001909101908401620002b3565b5085821015620002f257878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b608051610ace620003225f395f81816101ff01526103100152610ace5ff3fe608060405234801561000f575f80fd5b50600436106100cb575f3560e01c80635fe56b091161008857806395d89b411161006357806395d89b41146101a7578063a9059cbb146101af578063dd62ed3e146101c2578063ee9a31a2146101fa575f80fd5b80635fe56b091461015957806370a082311461016c578063757e987414610194575f80fd5b806306fdde03146100cf578063095ea7b3146100ed57806318160ddd1461011057806323b872dd14610122578063313ce5671461013557806340c10f1914610144575b5f80fd5b6100d7610239565b6040516100e491906107f6565b60405180910390f35b6101006100fb36600461085c565b6102c9565b60405190151581526020016100e4565b6002545b6040519081526020016100e4565b610100610130366004610884565b6102e2565b604051601281526020016100e4565b61015761015236600461085c565b610305565b005b610157610167366004610902565b6103e5565b61011461017a366004610976565b6001600160a01b03165f9081526020819052604090205490565b6101576101a2366004610996565b61043e565b6100d761048c565b6101006101bd36600461085c565b61049b565b6101146101d03660046109c0565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b6102217f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100e4565b606060038054610248906109e8565b80601f0160208091040260200160405190810160405280929190818152602001828054610274906109e8565b80156102bf5780601f10610296576101008083540402835291602001916102bf565b820191905f5260205f20905b8154815290600101906020018083116102a257829003601f168201915b5050505050905090565b5f336102d68185856104a8565b60019150505b92915050565b5f336102ef8582856104ba565b6102fa858585610535565b506001949350505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146103945760405162461bcd60e51b815260206004820152602960248201527f4173747269614d696e7461626c6545524332303a206f6e6c79206272696467656044820152680818d85b881b5a5b9d60ba1b60648201526084015b60405180910390fd5b61039e8282610592565b816001600160a01b03167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885826040516103d991815260200190565b60405180910390a25050565b6103ef33866105ca565b84336001600160a01b03167f0c64e29a5254a71c7f4e52b3d2d236348c80e00a00ba2e1961962bd2827c03fb8686868660405161042f9493929190610a48565b60405180910390a35050505050565b61044833836105ca565b6040516001600160a01b0382168152829033907fae8e66664d108544509c9a5b6a9f33c3b5fef3f88e5d3fa680706a6feb1360e39060200160405180910390a35050565b606060048054610248906109e8565b5f336102d6818585610535565b6104b583838360016105fe565b505050565b6001600160a01b038381165f908152600160209081526040808320938616835292905220545f19811461052f578181101561052157604051637dc7a0d960e11b81526001600160a01b0384166004820152602481018290526044810183905260640161038b565b61052f84848484035f6105fe565b50505050565b6001600160a01b03831661055e57604051634b637e8f60e11b81525f600482015260240161038b565b6001600160a01b0382166105875760405163ec442f0560e01b81525f600482015260240161038b565b6104b58383836106d0565b6001600160a01b0382166105bb5760405163ec442f0560e01b81525f600482015260240161038b565b6105c65f83836106d0565b5050565b6001600160a01b0382166105f357604051634b637e8f60e11b81525f600482015260240161038b565b6105c6825f836106d0565b6001600160a01b0384166106275760405163e602df0560e01b81525f600482015260240161038b565b6001600160a01b03831661065057604051634a1406b160e11b81525f600482015260240161038b565b6001600160a01b038085165f908152600160209081526040808320938716835292905220829055801561052f57826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516106c291815260200190565b60405180910390a350505050565b6001600160a01b0383166106fa578060025f8282546106ef9190610a79565b9091555061076a9050565b6001600160a01b0383165f908152602081905260409020548181101561074c5760405163391434e360e21b81526001600160a01b0385166004820152602481018290526044810183905260640161038b565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b038216610786576002805482900390556107a4565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516107e991815260200190565b60405180910390a3505050565b5f6020808352835180828501525f5b8181101561082157858101830151858201604001528201610805565b505f604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b0381168114610857575f80fd5b919050565b5f806040838503121561086d575f80fd5b61087683610841565b946020939093013593505050565b5f805f60608486031215610896575f80fd5b61089f84610841565b92506108ad60208501610841565b9150604084013590509250925092565b5f8083601f8401126108cd575f80fd5b50813567ffffffffffffffff8111156108e4575f80fd5b6020830191508360208285010111156108fb575f80fd5b9250929050565b5f805f805f60608688031215610916575f80fd5b85359450602086013567ffffffffffffffff80821115610934575f80fd5b61094089838a016108bd565b90965094506040880135915080821115610958575f80fd5b50610965888289016108bd565b969995985093965092949392505050565b5f60208284031215610986575f80fd5b61098f82610841565b9392505050565b5f80604083850312156109a7575f80fd5b823591506109b760208401610841565b90509250929050565b5f80604083850312156109d1575f80fd5b6109da83610841565b91506109b760208401610841565b600181811c908216806109fc57607f821691505b602082108103610a1a57634e487b7160e01b5f52602260045260245ffd5b50919050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b604081525f610a5b604083018688610a20565b8281036020840152610a6e818587610a20565b979650505050505050565b808201808211156102dc57634e487b7160e01b5f52601160045260245ffdfea264697066735822122011621167ab4de45ec6d8cbd0d773af54a01b2fe1b6017481c3c7a5eadf62a1d364736f6c63430008150033",
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

// WithdrawToIbcChain is a paid mutator transaction binding the contract method 0x5fe56b09.
//
// Solidity: function withdrawToIbcChain(uint256 _amount, string _destinationChainAddress, string _memo) returns()
func (_AstriaMintableERC20 *AstriaMintableERC20Transactor) WithdrawToIbcChain(opts *bind.TransactOpts, _amount *big.Int, _destinationChainAddress string, _memo string) (*types.Transaction, error) {
	return _AstriaMintableERC20.contract.Transact(opts, "withdrawToIbcChain", _amount, _destinationChainAddress, _memo)
}

// WithdrawToIbcChain is a paid mutator transaction binding the contract method 0x5fe56b09.
//
// Solidity: function withdrawToIbcChain(uint256 _amount, string _destinationChainAddress, string _memo) returns()
func (_AstriaMintableERC20 *AstriaMintableERC20Session) WithdrawToIbcChain(_amount *big.Int, _destinationChainAddress string, _memo string) (*types.Transaction, error) {
	return _AstriaMintableERC20.Contract.WithdrawToIbcChain(&_AstriaMintableERC20.TransactOpts, _amount, _destinationChainAddress, _memo)
}

// WithdrawToIbcChain is a paid mutator transaction binding the contract method 0x5fe56b09.
//
// Solidity: function withdrawToIbcChain(uint256 _amount, string _destinationChainAddress, string _memo) returns()
func (_AstriaMintableERC20 *AstriaMintableERC20TransactorSession) WithdrawToIbcChain(_amount *big.Int, _destinationChainAddress string, _memo string) (*types.Transaction, error) {
	return _AstriaMintableERC20.Contract.WithdrawToIbcChain(&_AstriaMintableERC20.TransactOpts, _amount, _destinationChainAddress, _memo)
}

// WithdrawToSequencer is a paid mutator transaction binding the contract method 0x757e9874.
//
// Solidity: function withdrawToSequencer(uint256 _amount, address _destinationChainAddress) returns()
func (_AstriaMintableERC20 *AstriaMintableERC20Transactor) WithdrawToSequencer(opts *bind.TransactOpts, _amount *big.Int, _destinationChainAddress common.Address) (*types.Transaction, error) {
	return _AstriaMintableERC20.contract.Transact(opts, "withdrawToSequencer", _amount, _destinationChainAddress)
}

// WithdrawToSequencer is a paid mutator transaction binding the contract method 0x757e9874.
//
// Solidity: function withdrawToSequencer(uint256 _amount, address _destinationChainAddress) returns()
func (_AstriaMintableERC20 *AstriaMintableERC20Session) WithdrawToSequencer(_amount *big.Int, _destinationChainAddress common.Address) (*types.Transaction, error) {
	return _AstriaMintableERC20.Contract.WithdrawToSequencer(&_AstriaMintableERC20.TransactOpts, _amount, _destinationChainAddress)
}

// WithdrawToSequencer is a paid mutator transaction binding the contract method 0x757e9874.
//
// Solidity: function withdrawToSequencer(uint256 _amount, address _destinationChainAddress) returns()
func (_AstriaMintableERC20 *AstriaMintableERC20TransactorSession) WithdrawToSequencer(_amount *big.Int, _destinationChainAddress common.Address) (*types.Transaction, error) {
	return _AstriaMintableERC20.Contract.WithdrawToSequencer(&_AstriaMintableERC20.TransactOpts, _amount, _destinationChainAddress)
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

// AstriaMintableERC20Ics20WithdrawalIterator is returned from FilterIcs20Withdrawal and is used to iterate over the raw logs and unpacked data for Ics20Withdrawal events raised by the AstriaMintableERC20 contract.
type AstriaMintableERC20Ics20WithdrawalIterator struct {
	Event *AstriaMintableERC20Ics20Withdrawal // Event containing the contract specifics and raw log

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
func (it *AstriaMintableERC20Ics20WithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AstriaMintableERC20Ics20Withdrawal)
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
		it.Event = new(AstriaMintableERC20Ics20Withdrawal)
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
func (it *AstriaMintableERC20Ics20WithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AstriaMintableERC20Ics20WithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AstriaMintableERC20Ics20Withdrawal represents a Ics20Withdrawal event raised by the AstriaMintableERC20 contract.
type AstriaMintableERC20Ics20Withdrawal struct {
	Sender                  common.Address
	Amount                  *big.Int
	DestinationChainAddress string
	Memo                    string
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterIcs20Withdrawal is a free log retrieval operation binding the contract event 0x0c64e29a5254a71c7f4e52b3d2d236348c80e00a00ba2e1961962bd2827c03fb.
//
// Solidity: event Ics20Withdrawal(address indexed sender, uint256 indexed amount, string destinationChainAddress, string memo)
func (_AstriaMintableERC20 *AstriaMintableERC20Filterer) FilterIcs20Withdrawal(opts *bind.FilterOpts, sender []common.Address, amount []*big.Int) (*AstriaMintableERC20Ics20WithdrawalIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _AstriaMintableERC20.contract.FilterLogs(opts, "Ics20Withdrawal", senderRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &AstriaMintableERC20Ics20WithdrawalIterator{contract: _AstriaMintableERC20.contract, event: "Ics20Withdrawal", logs: logs, sub: sub}, nil
}

// WatchIcs20Withdrawal is a free log subscription operation binding the contract event 0x0c64e29a5254a71c7f4e52b3d2d236348c80e00a00ba2e1961962bd2827c03fb.
//
// Solidity: event Ics20Withdrawal(address indexed sender, uint256 indexed amount, string destinationChainAddress, string memo)
func (_AstriaMintableERC20 *AstriaMintableERC20Filterer) WatchIcs20Withdrawal(opts *bind.WatchOpts, sink chan<- *AstriaMintableERC20Ics20Withdrawal, sender []common.Address, amount []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _AstriaMintableERC20.contract.WatchLogs(opts, "Ics20Withdrawal", senderRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AstriaMintableERC20Ics20Withdrawal)
				if err := _AstriaMintableERC20.contract.UnpackLog(event, "Ics20Withdrawal", log); err != nil {
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

// ParseIcs20Withdrawal is a log parse operation binding the contract event 0x0c64e29a5254a71c7f4e52b3d2d236348c80e00a00ba2e1961962bd2827c03fb.
//
// Solidity: event Ics20Withdrawal(address indexed sender, uint256 indexed amount, string destinationChainAddress, string memo)
func (_AstriaMintableERC20 *AstriaMintableERC20Filterer) ParseIcs20Withdrawal(log types.Log) (*AstriaMintableERC20Ics20Withdrawal, error) {
	event := new(AstriaMintableERC20Ics20Withdrawal)
	if err := _AstriaMintableERC20.contract.UnpackLog(event, "Ics20Withdrawal", log); err != nil {
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

// AstriaMintableERC20SequencerWithdrawalIterator is returned from FilterSequencerWithdrawal and is used to iterate over the raw logs and unpacked data for SequencerWithdrawal events raised by the AstriaMintableERC20 contract.
type AstriaMintableERC20SequencerWithdrawalIterator struct {
	Event *AstriaMintableERC20SequencerWithdrawal // Event containing the contract specifics and raw log

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
func (it *AstriaMintableERC20SequencerWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AstriaMintableERC20SequencerWithdrawal)
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
		it.Event = new(AstriaMintableERC20SequencerWithdrawal)
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
func (it *AstriaMintableERC20SequencerWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AstriaMintableERC20SequencerWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AstriaMintableERC20SequencerWithdrawal represents a SequencerWithdrawal event raised by the AstriaMintableERC20 contract.
type AstriaMintableERC20SequencerWithdrawal struct {
	Sender                  common.Address
	Amount                  *big.Int
	DestinationChainAddress common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterSequencerWithdrawal is a free log retrieval operation binding the contract event 0xae8e66664d108544509c9a5b6a9f33c3b5fef3f88e5d3fa680706a6feb1360e3.
//
// Solidity: event SequencerWithdrawal(address indexed sender, uint256 indexed amount, address destinationChainAddress)
func (_AstriaMintableERC20 *AstriaMintableERC20Filterer) FilterSequencerWithdrawal(opts *bind.FilterOpts, sender []common.Address, amount []*big.Int) (*AstriaMintableERC20SequencerWithdrawalIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _AstriaMintableERC20.contract.FilterLogs(opts, "SequencerWithdrawal", senderRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &AstriaMintableERC20SequencerWithdrawalIterator{contract: _AstriaMintableERC20.contract, event: "SequencerWithdrawal", logs: logs, sub: sub}, nil
}

// WatchSequencerWithdrawal is a free log subscription operation binding the contract event 0xae8e66664d108544509c9a5b6a9f33c3b5fef3f88e5d3fa680706a6feb1360e3.
//
// Solidity: event SequencerWithdrawal(address indexed sender, uint256 indexed amount, address destinationChainAddress)
func (_AstriaMintableERC20 *AstriaMintableERC20Filterer) WatchSequencerWithdrawal(opts *bind.WatchOpts, sink chan<- *AstriaMintableERC20SequencerWithdrawal, sender []common.Address, amount []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _AstriaMintableERC20.contract.WatchLogs(opts, "SequencerWithdrawal", senderRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AstriaMintableERC20SequencerWithdrawal)
				if err := _AstriaMintableERC20.contract.UnpackLog(event, "SequencerWithdrawal", log); err != nil {
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

// ParseSequencerWithdrawal is a log parse operation binding the contract event 0xae8e66664d108544509c9a5b6a9f33c3b5fef3f88e5d3fa680706a6feb1360e3.
//
// Solidity: event SequencerWithdrawal(address indexed sender, uint256 indexed amount, address destinationChainAddress)
func (_AstriaMintableERC20 *AstriaMintableERC20Filterer) ParseSequencerWithdrawal(log types.Log) (*AstriaMintableERC20SequencerWithdrawal, error) {
	event := new(AstriaMintableERC20SequencerWithdrawal)
	if err := _AstriaMintableERC20.contract.UnpackLog(event, "SequencerWithdrawal", log); err != nil {
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
