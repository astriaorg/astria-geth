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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_assetWithdrawalDecimals\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destinationChainAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"Ics20Withdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationChainAddress\",\"type\":\"address\"}],\"name\":\"SequencerWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ASSET_WITHDRAWAL_DECIMALS\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BRIDGE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_destinationChainAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_memo\",\"type\":\"string\"}],\"name\":\"withdrawToIbcChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_destinationChainAddress\",\"type\":\"address\"}],\"name\":\"withdrawToSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801562000010575f80fd5b5060405162000e7438038062000e74833981016040819052620000339162000132565b818160036200004383826200025f565b5060046200005282826200025f565b5050506001600160a01b0390931660a0525063ffffffff166080525062000327565b634e487b7160e01b5f52604160045260245ffd5b5f82601f83011262000098575f80fd5b81516001600160401b0380821115620000b557620000b562000074565b604051601f8301601f19908116603f01168101908282118183101715620000e057620000e062000074565b81604052838152602092508683858801011115620000fc575f80fd5b5f91505b838210156200011f578582018301518183018401529082019062000100565b5f93810190920192909252949350505050565b5f805f806080858703121562000146575f80fd5b84516001600160a01b03811681146200015d575f80fd5b602086015190945063ffffffff8116811462000177575f80fd5b60408601519093506001600160401b038082111562000194575f80fd5b620001a28883890162000088565b93506060870151915080821115620001b8575f80fd5b50620001c78782880162000088565b91505092959194509250565b600181811c90821680620001e857607f821691505b6020821081036200020757634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156200025a575f81815260208120601f850160051c81016020861015620002355750805b601f850160051c820191505b81811015620002565782815560010162000241565b5050505b505050565b81516001600160401b038111156200027b576200027b62000074565b62000293816200028c8454620001d3565b846200020d565b602080601f831160018114620002c9575f8415620002b15750858301515b5f19600386901b1c1916600185901b17855562000256565b5f85815260208120601f198616915b82811015620002f957888601518255948401946001909101908401620002d8565b50858210156200031757878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b60805160a051610b24620003505f395f8181610255015261036601525f6101c60152610b245ff3fe608060405234801561000f575f80fd5b50600436106100e5575f3560e01c806370a082311161008857806395d89b411161006357806395d89b41146101fd578063a9059cbb14610205578063dd62ed3e14610218578063ee9a31a214610250575f80fd5b806370a0823114610186578063757e9874146101ae5780638f2d8cb8146101c1575f80fd5b806323b872dd116100c357806323b872dd1461013c578063313ce5671461014f57806340c10f191461015e5780635fe56b0914610173575f80fd5b806306fdde03146100e9578063095ea7b31461010757806318160ddd1461012a575b5f80fd5b6100f161028f565b6040516100fe919061084c565b60405180910390f35b61011a6101153660046108b2565b61031f565b60405190151581526020016100fe565b6002545b6040519081526020016100fe565b61011a61014a3660046108da565b610338565b604051601281526020016100fe565b61017161016c3660046108b2565b61035b565b005b610171610181366004610958565b61043b565b61012e6101943660046109cc565b6001600160a01b03165f9081526020819052604090205490565b6101716101bc3660046109ec565b610494565b6101e87f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016100fe565b6100f16104e2565b61011a6102133660046108b2565b6104f1565b61012e610226366004610a16565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b6102777f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100fe565b60606003805461029e90610a3e565b80601f01602080910402602001604051908101604052809291908181526020018280546102ca90610a3e565b80156103155780601f106102ec57610100808354040283529160200191610315565b820191905f5260205f20905b8154815290600101906020018083116102f857829003601f168201915b5050505050905090565b5f3361032c8185856104fe565b60019150505b92915050565b5f33610345858285610510565b61035085858561058b565b506001949350505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146103ea5760405162461bcd60e51b815260206004820152602960248201527f4173747269614d696e7461626c6545524332303a206f6e6c79206272696467656044820152680818d85b881b5a5b9d60ba1b60648201526084015b60405180910390fd5b6103f482826105e8565b816001600160a01b03167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968858260405161042f91815260200190565b60405180910390a25050565b6104453386610620565b84336001600160a01b03167f0c64e29a5254a71c7f4e52b3d2d236348c80e00a00ba2e1961962bd2827c03fb868686866040516104859493929190610a9e565b60405180910390a35050505050565b61049e3383610620565b6040516001600160a01b0382168152829033907fae8e66664d108544509c9a5b6a9f33c3b5fef3f88e5d3fa680706a6feb1360e39060200160405180910390a35050565b60606004805461029e90610a3e565b5f3361032c81858561058b565b61050b8383836001610654565b505050565b6001600160a01b038381165f908152600160209081526040808320938616835292905220545f198114610585578181101561057757604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064016103e1565b61058584848484035f610654565b50505050565b6001600160a01b0383166105b457604051634b637e8f60e11b81525f60048201526024016103e1565b6001600160a01b0382166105dd5760405163ec442f0560e01b81525f60048201526024016103e1565b61050b838383610726565b6001600160a01b0382166106115760405163ec442f0560e01b81525f60048201526024016103e1565b61061c5f8383610726565b5050565b6001600160a01b03821661064957604051634b637e8f60e11b81525f60048201526024016103e1565b61061c825f83610726565b6001600160a01b03841661067d5760405163e602df0560e01b81525f60048201526024016103e1565b6001600160a01b0383166106a657604051634a1406b160e11b81525f60048201526024016103e1565b6001600160a01b038085165f908152600160209081526040808320938716835292905220829055801561058557826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161071891815260200190565b60405180910390a350505050565b6001600160a01b038316610750578060025f8282546107459190610acf565b909155506107c09050565b6001600160a01b0383165f90815260208190526040902054818110156107a25760405163391434e360e21b81526001600160a01b038516600482015260248101829052604481018390526064016103e1565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b0382166107dc576002805482900390556107fa565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161083f91815260200190565b60405180910390a3505050565b5f6020808352835180828501525f5b818110156108775785810183015185820160400152820161085b565b505f604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b03811681146108ad575f80fd5b919050565b5f80604083850312156108c3575f80fd5b6108cc83610897565b946020939093013593505050565b5f805f606084860312156108ec575f80fd5b6108f584610897565b925061090360208501610897565b9150604084013590509250925092565b5f8083601f840112610923575f80fd5b50813567ffffffffffffffff81111561093a575f80fd5b602083019150836020828501011115610951575f80fd5b9250929050565b5f805f805f6060868803121561096c575f80fd5b85359450602086013567ffffffffffffffff8082111561098a575f80fd5b61099689838a01610913565b909650945060408801359150808211156109ae575f80fd5b506109bb88828901610913565b969995985093965092949392505050565b5f602082840312156109dc575f80fd5b6109e582610897565b9392505050565b5f80604083850312156109fd575f80fd5b82359150610a0d60208401610897565b90509250929050565b5f8060408385031215610a27575f80fd5b610a3083610897565b9150610a0d60208401610897565b600181811c90821680610a5257607f821691505b602082108103610a7057634e487b7160e01b5f52602260045260245ffd5b50919050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b604081525f610ab1604083018688610a76565b8281036020840152610ac4818587610a76565b979650505050505050565b8082018082111561033257634e487b7160e01b5f52601160045260245ffdfea264697066735822122098489cb57d7515d4a1193b40331dc88740c1eff35a207983f03283dc375c139364736f6c63430008150033",
}

// AstriaMintableERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use AstriaMintableERC20MetaData.ABI instead.
var AstriaMintableERC20ABI = AstriaMintableERC20MetaData.ABI

// AstriaMintableERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AstriaMintableERC20MetaData.Bin instead.
var AstriaMintableERC20Bin = AstriaMintableERC20MetaData.Bin

// DeployAstriaMintableERC20 deploys a new Ethereum contract, binding an instance of AstriaMintableERC20 to it.
func DeployAstriaMintableERC20(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address, _assetWithdrawalDecimals uint32, _name string, _symbol string) (common.Address, *types.Transaction, *AstriaMintableERC20, error) {
	parsed, err := AstriaMintableERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AstriaMintableERC20Bin), backend, _bridge, _assetWithdrawalDecimals, _name, _symbol)
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

// ASSETWITHDRAWALDECIMALS is a free data retrieval call binding the contract method 0x8f2d8cb8.
//
// Solidity: function ASSET_WITHDRAWAL_DECIMALS() view returns(uint32)
func (_AstriaMintableERC20 *AstriaMintableERC20Caller) ASSETWITHDRAWALDECIMALS(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AstriaMintableERC20.contract.Call(opts, &out, "ASSET_WITHDRAWAL_DECIMALS")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ASSETWITHDRAWALDECIMALS is a free data retrieval call binding the contract method 0x8f2d8cb8.
//
// Solidity: function ASSET_WITHDRAWAL_DECIMALS() view returns(uint32)
func (_AstriaMintableERC20 *AstriaMintableERC20Session) ASSETWITHDRAWALDECIMALS() (uint32, error) {
	return _AstriaMintableERC20.Contract.ASSETWITHDRAWALDECIMALS(&_AstriaMintableERC20.CallOpts)
}

// ASSETWITHDRAWALDECIMALS is a free data retrieval call binding the contract method 0x8f2d8cb8.
//
// Solidity: function ASSET_WITHDRAWAL_DECIMALS() view returns(uint32)
func (_AstriaMintableERC20 *AstriaMintableERC20CallerSession) ASSETWITHDRAWALDECIMALS() (uint32, error) {
	return _AstriaMintableERC20.Contract.ASSETWITHDRAWALDECIMALS(&_AstriaMintableERC20.CallOpts)
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
