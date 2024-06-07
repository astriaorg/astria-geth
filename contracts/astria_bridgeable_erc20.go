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

// AstriaBridgeableERC20MetaData contains all meta data concerning the AstriaBridgeableERC20 contract.
var AstriaBridgeableERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_baseChainAssetPrecision\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destinationChainAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"Ics20Withdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationChainAddress\",\"type\":\"address\"}],\"name\":\"SequencerWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASE_CHAIN_ASSET_PRECISION\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BRIDGE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_destinationChainAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_memo\",\"type\":\"string\"}],\"name\":\"withdrawToIbcChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_destinationChainAddress\",\"type\":\"address\"}],\"name\":\"withdrawToSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801562000010575f80fd5b50604051620011fe380380620011fe833981016040819052620000339162000214565b8181600362000043838262000341565b50600462000052828262000341565b5050505f620000666200015160201b60201c565b90508060ff168463ffffffff161115620001125760405162461bcd60e51b815260206004820152605e60248201527f41737472696142726964676561626c6545524332303a2062617365206368616960448201527f6e20617373657420707265636973696f6e206d757374206265206c657373207460648201527f68616e206f7220657175616c20746f20746f6b656e20646563696d616c730000608482015260a40160405180910390fd5b63ffffffff84166080526200012b8460ff83166200041d565b6200013890600a6200053f565b60c052505050506001600160a01b031660a05262000559565b601290565b634e487b7160e01b5f52604160045260245ffd5b5f82601f8301126200017a575f80fd5b81516001600160401b038082111562000197576200019762000156565b604051601f8301601f19908116603f01168101908282118183101715620001c257620001c262000156565b81604052838152602092508683858801011115620001de575f80fd5b5f91505b83821015620002015785820183015181830184015290820190620001e2565b5f93810190920192909252949350505050565b5f805f806080858703121562000228575f80fd5b84516001600160a01b03811681146200023f575f80fd5b602086015190945063ffffffff8116811462000259575f80fd5b60408601519093506001600160401b038082111562000276575f80fd5b62000284888389016200016a565b935060608701519150808211156200029a575f80fd5b50620002a9878288016200016a565b91505092959194509250565b600181811c90821680620002ca57607f821691505b602082108103620002e957634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156200033c575f81815260208120601f850160051c81016020861015620003175750805b601f850160051c820191505b81811015620003385782815560010162000323565b5050505b505050565b81516001600160401b038111156200035d576200035d62000156565b62000375816200036e8454620002b5565b84620002ef565b602080601f831160018114620003ab575f8415620003935750858301515b5f19600386901b1c1916600185901b17855562000338565b5f85815260208120601f198616915b82811015620003db57888601518255948401946001909101908401620003ba565b5085821015620003f957878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b5f52601160045260245ffd5b63ffffffff8281168282160390808211156200043d576200043d62000409565b5092915050565b600181815b808511156200048457815f190482111562000468576200046862000409565b808516156200047657918102915b93841c939080029062000449565b509250929050565b5f826200049c5750600162000539565b81620004aa57505f62000539565b8160018114620004c35760028114620004ce57620004ee565b600191505062000539565b60ff841115620004e257620004e262000409565b50506001821b62000539565b5060208310610133831016604e8410600b841016171562000513575081810a62000539565b6200051f838362000444565b805f190482111562000535576200053562000409565b0290505b92915050565b5f6200055263ffffffff8416836200048c565b9392505050565b60805160a05160c051610c6c620005925f395f818161044401526104e701525f8181610255015261036601525f6101c60152610c6c5ff3fe608060405234801561000f575f80fd5b50600436106100e5575f3560e01c806370a082311161008857806395d89b411161006357806395d89b41146101fd578063a9059cbb14610205578063dd62ed3e14610218578063ee9a31a214610250575f80fd5b806370a0823114610186578063757e9874146101ae5780637eb6dec7146101c1575f80fd5b806323b872dd116100c357806323b872dd1461013c578063313ce5671461014f57806340c10f191461015e5780635fe56b0914610173575f80fd5b806306fdde03146100e9578063095ea7b31461010757806318160ddd1461012a575b5f80fd5b6100f161028f565b6040516100fe91906108d6565b60405180910390f35b61011a61011536600461093c565b61031f565b60405190151581526020016100fe565b6002545b6040519081526020016100fe565b61011a61014a366004610964565b610338565b604051601281526020016100fe565b61017161016c36600461093c565b61035b565b005b6101716101813660046109e2565b61043d565b61012e610194366004610a56565b6001600160a01b03165f9081526020819052604090205490565b6101716101bc366004610a76565b6104e0565b6101e87f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016100fe565b6100f1610579565b61011a61021336600461093c565b610588565b61012e610226366004610aa0565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b6102777f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100fe565b60606003805461029e90610ac8565b80601f01602080910402602001604051908101604052809291908181526020018280546102ca90610ac8565b80156103155780601f106102ec57610100808354040283529160200191610315565b820191905f5260205f20905b8154815290600101906020018083116102f857829003601f168201915b5050505050905090565b5f3361032c818585610595565b60019150505b92915050565b5f336103458582856105a7565b610350858585610622565b506001949350505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146103ec5760405162461bcd60e51b815260206004820152602b60248201527f41737472696142726964676561626c6545524332303a206f6e6c79206272696460448201526a19d94818d85b881b5a5b9d60aa1b60648201526084015b60405180910390fd5b6103f6828261067f565b816001600160a01b03167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968858260405161043191815260200190565b60405180910390a25050565b845f6104697f000000000000000000000000000000000000000000000000000000000000000083610b00565b116104865760405162461bcd60e51b81526004016103e390610b1f565b61049033876106b7565b85336001600160a01b03167f0c64e29a5254a71c7f4e52b3d2d236348c80e00a00ba2e1961962bd2827c03fb878787876040516104d09493929190610be6565b60405180910390a3505050505050565b815f61050c7f000000000000000000000000000000000000000000000000000000000000000083610b00565b116105295760405162461bcd60e51b81526004016103e390610b1f565b61053333846106b7565b6040516001600160a01b0383168152839033907fae8e66664d108544509c9a5b6a9f33c3b5fef3f88e5d3fa680706a6feb1360e3906020015b60405180910390a3505050565b60606004805461029e90610ac8565b5f3361032c818585610622565b6105a283838360016106eb565b505050565b6001600160a01b038381165f908152600160209081526040808320938616835292905220545f19811461061c578181101561060e57604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064016103e3565b61061c84848484035f6106eb565b50505050565b6001600160a01b03831661064b57604051634b637e8f60e11b81525f60048201526024016103e3565b6001600160a01b0382166106745760405163ec442f0560e01b81525f60048201526024016103e3565b6105a28383836107bd565b6001600160a01b0382166106a85760405163ec442f0560e01b81525f60048201526024016103e3565b6106b35f83836107bd565b5050565b6001600160a01b0382166106e057604051634b637e8f60e11b81525f60048201526024016103e3565b6106b3825f836107bd565b6001600160a01b0384166107145760405163e602df0560e01b81525f60048201526024016103e3565b6001600160a01b03831661073d57604051634a1406b160e11b81525f60048201526024016103e3565b6001600160a01b038085165f908152600160209081526040808320938716835292905220829055801561061c57826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516107af91815260200190565b60405180910390a350505050565b6001600160a01b0383166107e7578060025f8282546107dc9190610c17565b909155506108579050565b6001600160a01b0383165f90815260208190526040902054818110156108395760405163391434e360e21b81526001600160a01b038516600482015260248101829052604481018390526064016103e3565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b03821661087357600280548290039055610891565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161056c91815260200190565b5f6020808352835180828501525f5b81811015610901578581018301518582016040015282016108e5565b505f604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b0381168114610937575f80fd5b919050565b5f806040838503121561094d575f80fd5b61095683610921565b946020939093013593505050565b5f805f60608486031215610976575f80fd5b61097f84610921565b925061098d60208501610921565b9150604084013590509250925092565b5f8083601f8401126109ad575f80fd5b50813567ffffffffffffffff8111156109c4575f80fd5b6020830191508360208285010111156109db575f80fd5b9250929050565b5f805f805f606086880312156109f6575f80fd5b85359450602086013567ffffffffffffffff80821115610a14575f80fd5b610a2089838a0161099d565b90965094506040880135915080821115610a38575f80fd5b50610a458882890161099d565b969995985093965092949392505050565b5f60208284031215610a66575f80fd5b610a6f82610921565b9392505050565b5f8060408385031215610a87575f80fd5b82359150610a9760208401610921565b90509250929050565b5f8060408385031215610ab1575f80fd5b610aba83610921565b9150610a9760208401610921565b600181811c90821680610adc57607f821691505b602082108103610afa57634e487b7160e01b5f52602260045260245ffd5b50919050565b5f82610b1a57634e487b7160e01b5f52601260045260245ffd5b500490565b60208082526073908201527f41737472696142726964676561626c6545524332303a20696e7375666669636960408201527f656e742076616c75652c206d7573742062652067726561746572207468616e2060608201527f3130202a2a2028544f4b454e5f444543494d414c53202d20424153455f434841608082015272494e5f41535345545f505245434953494f4e2960681b60a082015260c00190565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b604081525f610bf9604083018688610bbe565b8281036020840152610c0c818587610bbe565b979650505050505050565b8082018082111561033257634e487b7160e01b5f52601160045260245ffdfea26469706673582212202de14742c4900725b016ad188cb5ed05ada890c96df6755671b0914095ccb68e64736f6c63430008150033",
}

// AstriaBridgeableERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use AstriaBridgeableERC20MetaData.ABI instead.
var AstriaBridgeableERC20ABI = AstriaBridgeableERC20MetaData.ABI

// AstriaBridgeableERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AstriaBridgeableERC20MetaData.Bin instead.
var AstriaBridgeableERC20Bin = AstriaBridgeableERC20MetaData.Bin

// DeployAstriaBridgeableERC20 deploys a new Ethereum contract, binding an instance of AstriaBridgeableERC20 to it.
func DeployAstriaBridgeableERC20(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address, _baseChainAssetPrecision uint32, _name string, _symbol string) (common.Address, *types.Transaction, *AstriaBridgeableERC20, error) {
	parsed, err := AstriaBridgeableERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AstriaBridgeableERC20Bin), backend, _bridge, _baseChainAssetPrecision, _name, _symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AstriaBridgeableERC20{AstriaBridgeableERC20Caller: AstriaBridgeableERC20Caller{contract: contract}, AstriaBridgeableERC20Transactor: AstriaBridgeableERC20Transactor{contract: contract}, AstriaBridgeableERC20Filterer: AstriaBridgeableERC20Filterer{contract: contract}}, nil
}

// AstriaBridgeableERC20 is an auto generated Go binding around an Ethereum contract.
type AstriaBridgeableERC20 struct {
	AstriaBridgeableERC20Caller     // Read-only binding to the contract
	AstriaBridgeableERC20Transactor // Write-only binding to the contract
	AstriaBridgeableERC20Filterer   // Log filterer for contract events
}

// AstriaBridgeableERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type AstriaBridgeableERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AstriaBridgeableERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AstriaBridgeableERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AstriaBridgeableERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AstriaBridgeableERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AstriaBridgeableERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AstriaBridgeableERC20Session struct {
	Contract     *AstriaBridgeableERC20 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AstriaBridgeableERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AstriaBridgeableERC20CallerSession struct {
	Contract *AstriaBridgeableERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// AstriaBridgeableERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AstriaBridgeableERC20TransactorSession struct {
	Contract     *AstriaBridgeableERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// AstriaBridgeableERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type AstriaBridgeableERC20Raw struct {
	Contract *AstriaBridgeableERC20 // Generic contract binding to access the raw methods on
}

// AstriaBridgeableERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AstriaBridgeableERC20CallerRaw struct {
	Contract *AstriaBridgeableERC20Caller // Generic read-only contract binding to access the raw methods on
}

// AstriaBridgeableERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AstriaBridgeableERC20TransactorRaw struct {
	Contract *AstriaBridgeableERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAstriaBridgeableERC20 creates a new instance of AstriaBridgeableERC20, bound to a specific deployed contract.
func NewAstriaBridgeableERC20(address common.Address, backend bind.ContractBackend) (*AstriaBridgeableERC20, error) {
	contract, err := bindAstriaBridgeableERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AstriaBridgeableERC20{AstriaBridgeableERC20Caller: AstriaBridgeableERC20Caller{contract: contract}, AstriaBridgeableERC20Transactor: AstriaBridgeableERC20Transactor{contract: contract}, AstriaBridgeableERC20Filterer: AstriaBridgeableERC20Filterer{contract: contract}}, nil
}

// NewAstriaBridgeableERC20Caller creates a new read-only instance of AstriaBridgeableERC20, bound to a specific deployed contract.
func NewAstriaBridgeableERC20Caller(address common.Address, caller bind.ContractCaller) (*AstriaBridgeableERC20Caller, error) {
	contract, err := bindAstriaBridgeableERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AstriaBridgeableERC20Caller{contract: contract}, nil
}

// NewAstriaBridgeableERC20Transactor creates a new write-only instance of AstriaBridgeableERC20, bound to a specific deployed contract.
func NewAstriaBridgeableERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*AstriaBridgeableERC20Transactor, error) {
	contract, err := bindAstriaBridgeableERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AstriaBridgeableERC20Transactor{contract: contract}, nil
}

// NewAstriaBridgeableERC20Filterer creates a new log filterer instance of AstriaBridgeableERC20, bound to a specific deployed contract.
func NewAstriaBridgeableERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*AstriaBridgeableERC20Filterer, error) {
	contract, err := bindAstriaBridgeableERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AstriaBridgeableERC20Filterer{contract: contract}, nil
}

// bindAstriaBridgeableERC20 binds a generic wrapper to an already deployed contract.
func bindAstriaBridgeableERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AstriaBridgeableERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AstriaBridgeableERC20.Contract.AstriaBridgeableERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AstriaBridgeableERC20.Contract.AstriaBridgeableERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AstriaBridgeableERC20.Contract.AstriaBridgeableERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AstriaBridgeableERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AstriaBridgeableERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AstriaBridgeableERC20.Contract.contract.Transact(opts, method, params...)
}

// BASECHAINASSETPRECISION is a free data retrieval call binding the contract method 0x7eb6dec7.
//
// Solidity: function BASE_CHAIN_ASSET_PRECISION() view returns(uint32)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Caller) BASECHAINASSETPRECISION(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AstriaBridgeableERC20.contract.Call(opts, &out, "BASE_CHAIN_ASSET_PRECISION")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// BASECHAINASSETPRECISION is a free data retrieval call binding the contract method 0x7eb6dec7.
//
// Solidity: function BASE_CHAIN_ASSET_PRECISION() view returns(uint32)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Session) BASECHAINASSETPRECISION() (uint32, error) {
	return _AstriaBridgeableERC20.Contract.BASECHAINASSETPRECISION(&_AstriaBridgeableERC20.CallOpts)
}

// BASECHAINASSETPRECISION is a free data retrieval call binding the contract method 0x7eb6dec7.
//
// Solidity: function BASE_CHAIN_ASSET_PRECISION() view returns(uint32)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20CallerSession) BASECHAINASSETPRECISION() (uint32, error) {
	return _AstriaBridgeableERC20.Contract.BASECHAINASSETPRECISION(&_AstriaBridgeableERC20.CallOpts)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Caller) BRIDGE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AstriaBridgeableERC20.contract.Call(opts, &out, "BRIDGE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Session) BRIDGE() (common.Address, error) {
	return _AstriaBridgeableERC20.Contract.BRIDGE(&_AstriaBridgeableERC20.CallOpts)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20CallerSession) BRIDGE() (common.Address, error) {
	return _AstriaBridgeableERC20.Contract.BRIDGE(&_AstriaBridgeableERC20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AstriaBridgeableERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AstriaBridgeableERC20.Contract.Allowance(&_AstriaBridgeableERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AstriaBridgeableERC20.Contract.Allowance(&_AstriaBridgeableERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AstriaBridgeableERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _AstriaBridgeableERC20.Contract.BalanceOf(&_AstriaBridgeableERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _AstriaBridgeableERC20.Contract.BalanceOf(&_AstriaBridgeableERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AstriaBridgeableERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Session) Decimals() (uint8, error) {
	return _AstriaBridgeableERC20.Contract.Decimals(&_AstriaBridgeableERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20CallerSession) Decimals() (uint8, error) {
	return _AstriaBridgeableERC20.Contract.Decimals(&_AstriaBridgeableERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AstriaBridgeableERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Session) Name() (string, error) {
	return _AstriaBridgeableERC20.Contract.Name(&_AstriaBridgeableERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20CallerSession) Name() (string, error) {
	return _AstriaBridgeableERC20.Contract.Name(&_AstriaBridgeableERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AstriaBridgeableERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Session) Symbol() (string, error) {
	return _AstriaBridgeableERC20.Contract.Symbol(&_AstriaBridgeableERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20CallerSession) Symbol() (string, error) {
	return _AstriaBridgeableERC20.Contract.Symbol(&_AstriaBridgeableERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AstriaBridgeableERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Session) TotalSupply() (*big.Int, error) {
	return _AstriaBridgeableERC20.Contract.TotalSupply(&_AstriaBridgeableERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _AstriaBridgeableERC20.Contract.TotalSupply(&_AstriaBridgeableERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AstriaBridgeableERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AstriaBridgeableERC20.Contract.Approve(&_AstriaBridgeableERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AstriaBridgeableERC20.Contract.Approve(&_AstriaBridgeableERC20.TransactOpts, spender, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Transactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AstriaBridgeableERC20.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Session) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AstriaBridgeableERC20.Contract.Mint(&_AstriaBridgeableERC20.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20TransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AstriaBridgeableERC20.Contract.Mint(&_AstriaBridgeableERC20.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AstriaBridgeableERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AstriaBridgeableERC20.Contract.Transfer(&_AstriaBridgeableERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AstriaBridgeableERC20.Contract.Transfer(&_AstriaBridgeableERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AstriaBridgeableERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AstriaBridgeableERC20.Contract.TransferFrom(&_AstriaBridgeableERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AstriaBridgeableERC20.Contract.TransferFrom(&_AstriaBridgeableERC20.TransactOpts, from, to, value)
}

// WithdrawToIbcChain is a paid mutator transaction binding the contract method 0x5fe56b09.
//
// Solidity: function withdrawToIbcChain(uint256 _amount, string _destinationChainAddress, string _memo) returns()
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Transactor) WithdrawToIbcChain(opts *bind.TransactOpts, _amount *big.Int, _destinationChainAddress string, _memo string) (*types.Transaction, error) {
	return _AstriaBridgeableERC20.contract.Transact(opts, "withdrawToIbcChain", _amount, _destinationChainAddress, _memo)
}

// WithdrawToIbcChain is a paid mutator transaction binding the contract method 0x5fe56b09.
//
// Solidity: function withdrawToIbcChain(uint256 _amount, string _destinationChainAddress, string _memo) returns()
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Session) WithdrawToIbcChain(_amount *big.Int, _destinationChainAddress string, _memo string) (*types.Transaction, error) {
	return _AstriaBridgeableERC20.Contract.WithdrawToIbcChain(&_AstriaBridgeableERC20.TransactOpts, _amount, _destinationChainAddress, _memo)
}

// WithdrawToIbcChain is a paid mutator transaction binding the contract method 0x5fe56b09.
//
// Solidity: function withdrawToIbcChain(uint256 _amount, string _destinationChainAddress, string _memo) returns()
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20TransactorSession) WithdrawToIbcChain(_amount *big.Int, _destinationChainAddress string, _memo string) (*types.Transaction, error) {
	return _AstriaBridgeableERC20.Contract.WithdrawToIbcChain(&_AstriaBridgeableERC20.TransactOpts, _amount, _destinationChainAddress, _memo)
}

// WithdrawToSequencer is a paid mutator transaction binding the contract method 0x757e9874.
//
// Solidity: function withdrawToSequencer(uint256 _amount, address _destinationChainAddress) returns()
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Transactor) WithdrawToSequencer(opts *bind.TransactOpts, _amount *big.Int, _destinationChainAddress common.Address) (*types.Transaction, error) {
	return _AstriaBridgeableERC20.contract.Transact(opts, "withdrawToSequencer", _amount, _destinationChainAddress)
}

// WithdrawToSequencer is a paid mutator transaction binding the contract method 0x757e9874.
//
// Solidity: function withdrawToSequencer(uint256 _amount, address _destinationChainAddress) returns()
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Session) WithdrawToSequencer(_amount *big.Int, _destinationChainAddress common.Address) (*types.Transaction, error) {
	return _AstriaBridgeableERC20.Contract.WithdrawToSequencer(&_AstriaBridgeableERC20.TransactOpts, _amount, _destinationChainAddress)
}

// WithdrawToSequencer is a paid mutator transaction binding the contract method 0x757e9874.
//
// Solidity: function withdrawToSequencer(uint256 _amount, address _destinationChainAddress) returns()
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20TransactorSession) WithdrawToSequencer(_amount *big.Int, _destinationChainAddress common.Address) (*types.Transaction, error) {
	return _AstriaBridgeableERC20.Contract.WithdrawToSequencer(&_AstriaBridgeableERC20.TransactOpts, _amount, _destinationChainAddress)
}

// AstriaBridgeableERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the AstriaBridgeableERC20 contract.
type AstriaBridgeableERC20ApprovalIterator struct {
	Event *AstriaBridgeableERC20Approval // Event containing the contract specifics and raw log

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
func (it *AstriaBridgeableERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AstriaBridgeableERC20Approval)
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
		it.Event = new(AstriaBridgeableERC20Approval)
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
func (it *AstriaBridgeableERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AstriaBridgeableERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AstriaBridgeableERC20Approval represents a Approval event raised by the AstriaBridgeableERC20 contract.
type AstriaBridgeableERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AstriaBridgeableERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AstriaBridgeableERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AstriaBridgeableERC20ApprovalIterator{contract: _AstriaBridgeableERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AstriaBridgeableERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AstriaBridgeableERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AstriaBridgeableERC20Approval)
				if err := _AstriaBridgeableERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Filterer) ParseApproval(log types.Log) (*AstriaBridgeableERC20Approval, error) {
	event := new(AstriaBridgeableERC20Approval)
	if err := _AstriaBridgeableERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AstriaBridgeableERC20Ics20WithdrawalIterator is returned from FilterIcs20Withdrawal and is used to iterate over the raw logs and unpacked data for Ics20Withdrawal events raised by the AstriaBridgeableERC20 contract.
type AstriaBridgeableERC20Ics20WithdrawalIterator struct {
	Event *AstriaBridgeableERC20Ics20Withdrawal // Event containing the contract specifics and raw log

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
func (it *AstriaBridgeableERC20Ics20WithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AstriaBridgeableERC20Ics20Withdrawal)
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
		it.Event = new(AstriaBridgeableERC20Ics20Withdrawal)
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
func (it *AstriaBridgeableERC20Ics20WithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AstriaBridgeableERC20Ics20WithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AstriaBridgeableERC20Ics20Withdrawal represents a Ics20Withdrawal event raised by the AstriaBridgeableERC20 contract.
type AstriaBridgeableERC20Ics20Withdrawal struct {
	Sender                  common.Address
	Amount                  *big.Int
	DestinationChainAddress string
	Memo                    string
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterIcs20Withdrawal is a free log retrieval operation binding the contract event 0x0c64e29a5254a71c7f4e52b3d2d236348c80e00a00ba2e1961962bd2827c03fb.
//
// Solidity: event Ics20Withdrawal(address indexed sender, uint256 indexed amount, string destinationChainAddress, string memo)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Filterer) FilterIcs20Withdrawal(opts *bind.FilterOpts, sender []common.Address, amount []*big.Int) (*AstriaBridgeableERC20Ics20WithdrawalIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _AstriaBridgeableERC20.contract.FilterLogs(opts, "Ics20Withdrawal", senderRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &AstriaBridgeableERC20Ics20WithdrawalIterator{contract: _AstriaBridgeableERC20.contract, event: "Ics20Withdrawal", logs: logs, sub: sub}, nil
}

// WatchIcs20Withdrawal is a free log subscription operation binding the contract event 0x0c64e29a5254a71c7f4e52b3d2d236348c80e00a00ba2e1961962bd2827c03fb.
//
// Solidity: event Ics20Withdrawal(address indexed sender, uint256 indexed amount, string destinationChainAddress, string memo)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Filterer) WatchIcs20Withdrawal(opts *bind.WatchOpts, sink chan<- *AstriaBridgeableERC20Ics20Withdrawal, sender []common.Address, amount []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _AstriaBridgeableERC20.contract.WatchLogs(opts, "Ics20Withdrawal", senderRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AstriaBridgeableERC20Ics20Withdrawal)
				if err := _AstriaBridgeableERC20.contract.UnpackLog(event, "Ics20Withdrawal", log); err != nil {
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
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Filterer) ParseIcs20Withdrawal(log types.Log) (*AstriaBridgeableERC20Ics20Withdrawal, error) {
	event := new(AstriaBridgeableERC20Ics20Withdrawal)
	if err := _AstriaBridgeableERC20.contract.UnpackLog(event, "Ics20Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AstriaBridgeableERC20MintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the AstriaBridgeableERC20 contract.
type AstriaBridgeableERC20MintIterator struct {
	Event *AstriaBridgeableERC20Mint // Event containing the contract specifics and raw log

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
func (it *AstriaBridgeableERC20MintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AstriaBridgeableERC20Mint)
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
		it.Event = new(AstriaBridgeableERC20Mint)
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
func (it *AstriaBridgeableERC20MintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AstriaBridgeableERC20MintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AstriaBridgeableERC20Mint represents a Mint event raised by the AstriaBridgeableERC20 contract.
type AstriaBridgeableERC20Mint struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed account, uint256 amount)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Filterer) FilterMint(opts *bind.FilterOpts, account []common.Address) (*AstriaBridgeableERC20MintIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AstriaBridgeableERC20.contract.FilterLogs(opts, "Mint", accountRule)
	if err != nil {
		return nil, err
	}
	return &AstriaBridgeableERC20MintIterator{contract: _AstriaBridgeableERC20.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed account, uint256 amount)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Filterer) WatchMint(opts *bind.WatchOpts, sink chan<- *AstriaBridgeableERC20Mint, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AstriaBridgeableERC20.contract.WatchLogs(opts, "Mint", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AstriaBridgeableERC20Mint)
				if err := _AstriaBridgeableERC20.contract.UnpackLog(event, "Mint", log); err != nil {
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
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Filterer) ParseMint(log types.Log) (*AstriaBridgeableERC20Mint, error) {
	event := new(AstriaBridgeableERC20Mint)
	if err := _AstriaBridgeableERC20.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AstriaBridgeableERC20SequencerWithdrawalIterator is returned from FilterSequencerWithdrawal and is used to iterate over the raw logs and unpacked data for SequencerWithdrawal events raised by the AstriaBridgeableERC20 contract.
type AstriaBridgeableERC20SequencerWithdrawalIterator struct {
	Event *AstriaBridgeableERC20SequencerWithdrawal // Event containing the contract specifics and raw log

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
func (it *AstriaBridgeableERC20SequencerWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AstriaBridgeableERC20SequencerWithdrawal)
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
		it.Event = new(AstriaBridgeableERC20SequencerWithdrawal)
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
func (it *AstriaBridgeableERC20SequencerWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AstriaBridgeableERC20SequencerWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AstriaBridgeableERC20SequencerWithdrawal represents a SequencerWithdrawal event raised by the AstriaBridgeableERC20 contract.
type AstriaBridgeableERC20SequencerWithdrawal struct {
	Sender                  common.Address
	Amount                  *big.Int
	DestinationChainAddress common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterSequencerWithdrawal is a free log retrieval operation binding the contract event 0xae8e66664d108544509c9a5b6a9f33c3b5fef3f88e5d3fa680706a6feb1360e3.
//
// Solidity: event SequencerWithdrawal(address indexed sender, uint256 indexed amount, address destinationChainAddress)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Filterer) FilterSequencerWithdrawal(opts *bind.FilterOpts, sender []common.Address, amount []*big.Int) (*AstriaBridgeableERC20SequencerWithdrawalIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _AstriaBridgeableERC20.contract.FilterLogs(opts, "SequencerWithdrawal", senderRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &AstriaBridgeableERC20SequencerWithdrawalIterator{contract: _AstriaBridgeableERC20.contract, event: "SequencerWithdrawal", logs: logs, sub: sub}, nil
}

// WatchSequencerWithdrawal is a free log subscription operation binding the contract event 0xae8e66664d108544509c9a5b6a9f33c3b5fef3f88e5d3fa680706a6feb1360e3.
//
// Solidity: event SequencerWithdrawal(address indexed sender, uint256 indexed amount, address destinationChainAddress)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Filterer) WatchSequencerWithdrawal(opts *bind.WatchOpts, sink chan<- *AstriaBridgeableERC20SequencerWithdrawal, sender []common.Address, amount []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _AstriaBridgeableERC20.contract.WatchLogs(opts, "SequencerWithdrawal", senderRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AstriaBridgeableERC20SequencerWithdrawal)
				if err := _AstriaBridgeableERC20.contract.UnpackLog(event, "SequencerWithdrawal", log); err != nil {
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
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Filterer) ParseSequencerWithdrawal(log types.Log) (*AstriaBridgeableERC20SequencerWithdrawal, error) {
	event := new(AstriaBridgeableERC20SequencerWithdrawal)
	if err := _AstriaBridgeableERC20.contract.UnpackLog(event, "SequencerWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AstriaBridgeableERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the AstriaBridgeableERC20 contract.
type AstriaBridgeableERC20TransferIterator struct {
	Event *AstriaBridgeableERC20Transfer // Event containing the contract specifics and raw log

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
func (it *AstriaBridgeableERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AstriaBridgeableERC20Transfer)
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
		it.Event = new(AstriaBridgeableERC20Transfer)
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
func (it *AstriaBridgeableERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AstriaBridgeableERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AstriaBridgeableERC20Transfer represents a Transfer event raised by the AstriaBridgeableERC20 contract.
type AstriaBridgeableERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AstriaBridgeableERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AstriaBridgeableERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AstriaBridgeableERC20TransferIterator{contract: _AstriaBridgeableERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AstriaBridgeableERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AstriaBridgeableERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AstriaBridgeableERC20Transfer)
				if err := _AstriaBridgeableERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Filterer) ParseTransfer(log types.Log) (*AstriaBridgeableERC20Transfer, error) {
	event := new(AstriaBridgeableERC20Transfer)
	if err := _AstriaBridgeableERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
