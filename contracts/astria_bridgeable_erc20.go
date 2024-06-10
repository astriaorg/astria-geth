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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_baseChainAssetPrecision\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"_baseChainBridgeAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_baseChainAssetDenomination\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destinationChainAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"Ics20Withdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destinationChainAddress\",\"type\":\"string\"}],\"name\":\"SequencerWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASE_CHAIN_ASSET_DENOMINATION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASE_CHAIN_ASSET_PRECISION\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASE_CHAIN_BRIDGE_ADDRESS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BRIDGE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_destinationChainAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_memo\",\"type\":\"string\"}],\"name\":\"withdrawToIbcChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_destinationChainAddress\",\"type\":\"string\"}],\"name\":\"withdrawToSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801562000010575f80fd5b506040516200138838038062001388833981016040819052620000339162000266565b81816005620000438382620003c8565b506006620000528282620003c8565b5050505f620000666200017360201b60201c565b90508060ff168663ffffffff161115620001125760405162461bcd60e51b815260206004820152605e60248201527f41737472696142726964676561626c6545524332303a2062617365206368616960448201527f6e20617373657420707265636973696f6e206d757374206265206c657373207460648201527f68616e206f7220657175616c20746f20746f6b656e20646563696d616c730000608482015260a40160405180910390fd5b63ffffffff86166080525f620001298682620003c8565b506001620001388582620003c8565b50620001488660ff8316620004a4565b6200015590600a620005c6565b60c0525050506001600160a01b0390931660a05250620005e0915050565b601290565b80516001600160a01b03811681146200018f575f80fd5b919050565b805163ffffffff811681146200018f575f80fd5b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112620001cc575f80fd5b81516001600160401b0380821115620001e957620001e9620001a8565b604051601f8301601f19908116603f01168101908282118183101715620002145762000214620001a8565b8160405283815260209250868385880101111562000230575f80fd5b5f91505b8382101562000253578582018301518183018401529082019062000234565b5f93810190920192909252949350505050565b5f805f805f8060c087890312156200027c575f80fd5b620002878762000178565b9550620002976020880162000194565b60408801519095506001600160401b0380821115620002b4575f80fd5b620002c28a838b01620001bc565b95506060890151915080821115620002d8575f80fd5b620002e68a838b01620001bc565b94506080890151915080821115620002fc575f80fd5b6200030a8a838b01620001bc565b935060a089015191508082111562000320575f80fd5b506200032f89828a01620001bc565b9150509295509295509295565b600181811c908216806200035157607f821691505b6020821081036200037057634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115620003c3575f81815260208120601f850160051c810160208610156200039e5750805b601f850160051c820191505b81811015620003bf57828155600101620003aa565b5050505b505050565b81516001600160401b03811115620003e457620003e4620001a8565b620003fc81620003f584546200033c565b8462000376565b602080601f83116001811462000432575f84156200041a5750858301515b5f19600386901b1c1916600185901b178555620003bf565b5f85815260208120601f198616915b82811015620004625788860151825594840194600190910190840162000441565b50858210156200048057878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b5f52601160045260245ffd5b63ffffffff828116828216039080821115620004c457620004c462000490565b5092915050565b600181815b808511156200050b57815f1904821115620004ef57620004ef62000490565b80851615620004fd57918102915b93841c9390800290620004d0565b509250929050565b5f826200052357506001620005c0565b816200053157505f620005c0565b81600181146200054a5760028114620005555762000575565b6001915050620005c0565b60ff84111562000569576200056962000490565b50506001821b620005c0565b5060208310610133831016604e8410600b84101617156200059a575081810a620005c0565b620005a68383620004cb565b805f1904821115620005bc57620005bc62000490565b0290505b92915050565b5f620005d963ffffffff84168362000513565b9392505050565b60805160a05160c051610d6f620006195f395f818161046a01526105b501525f818161027b015261038c01525f6101c90152610d6f5ff3fe608060405234801561000f575f80fd5b50600436106100fb575f3560e01c80637eb6dec711610093578063d38fe9a711610063578063d38fe9a714610223578063db97dc9814610236578063dd62ed3e1461023e578063ee9a31a214610276575f80fd5b80637eb6dec7146101c457806395d89b4114610200578063a9059cbb14610208578063b6476c7e1461021b575f80fd5b8063313ce567116100ce578063313ce5671461016557806340c10f19146101745780635fe56b091461018957806370a082311461019c575f80fd5b806306fdde03146100ff578063095ea7b31461011d57806318160ddd1461014057806323b872dd14610152575b5f80fd5b6101076102b5565b6040516101149190610997565b60405180910390f35b61013061012b3660046109fd565b610345565b6040519015158152602001610114565b6004545b604051908152602001610114565b610130610160366004610a25565b61035e565b60405160128152602001610114565b6101876101823660046109fd565b610381565b005b610187610197366004610aa3565b610463565b6101446101aa366004610b17565b6001600160a01b03165f9081526002602052604090205490565b6101eb7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610114565b610107610506565b6101306102163660046109fd565b610515565b610107610522565b610187610231366004610b37565b6105ae565b61010761064b565b61014461024c366004610b7f565b6001600160a01b039182165f90815260036020908152604080832093909416825291909152205490565b61029d7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610114565b6060600580546102c490610bb0565b80601f01602080910402602001604051908101604052809291908181526020018280546102f090610bb0565b801561033b5780601f106103125761010080835404028352916020019161033b565b820191905f5260205f20905b81548152906001019060200180831161031e57829003601f168201915b5050505050905090565b5f33610352818585610657565b60019150505b92915050565b5f3361036b858285610669565b6103768585856106e4565b506001949350505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146104125760405162461bcd60e51b815260206004820152602b60248201527f41737472696142726964676561626c6545524332303a206f6e6c79206272696460448201526a19d94818d85b881b5a5b9d60aa1b60648201526084015b60405180910390fd5b61041c8282610741565b816001600160a01b03167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968858260405161045791815260200190565b60405180910390a25050565b845f61048f7f000000000000000000000000000000000000000000000000000000000000000083610be8565b116104ac5760405162461bcd60e51b815260040161040990610c07565b6104b63387610779565b85336001600160a01b03167f0c64e29a5254a71c7f4e52b3d2d236348c80e00a00ba2e1961962bd2827c03fb878787876040516104f69493929190610cce565b60405180910390a3505050505050565b6060600680546102c490610bb0565b5f336103528185856106e4565b6001805461052f90610bb0565b80601f016020809104026020016040519081016040528092919081815260200182805461055b90610bb0565b80156105a65780601f1061057d576101008083540402835291602001916105a6565b820191905f5260205f20905b81548152906001019060200180831161058957829003601f168201915b505050505081565b825f6105da7f000000000000000000000000000000000000000000000000000000000000000083610be8565b116105f75760405162461bcd60e51b815260040161040990610c07565b6106013385610779565b83336001600160a01b03167f0f4961cab7530804898499aa89f5ec81d1a73102e2e4a1f30f88e5ae3513ba2a858560405161063d929190610cff565b60405180910390a350505050565b5f805461052f90610bb0565b61066483838360016107ad565b505050565b6001600160a01b038381165f908152600360209081526040808320938616835292905220545f1981146106de57818110156106d057604051637dc7a0d960e11b81526001600160a01b03841660048201526024810182905260448101839052606401610409565b6106de84848484035f6107ad565b50505050565b6001600160a01b03831661070d57604051634b637e8f60e11b81525f6004820152602401610409565b6001600160a01b0382166107365760405163ec442f0560e01b81525f6004820152602401610409565b610664838383610871565b6001600160a01b03821661076a5760405163ec442f0560e01b81525f6004820152602401610409565b6107755f8383610871565b5050565b6001600160a01b0382166107a257604051634b637e8f60e11b81525f6004820152602401610409565b610775825f83610871565b6001600160a01b0384166107d65760405163e602df0560e01b81525f6004820152602401610409565b6001600160a01b0383166107ff57604051634a1406b160e11b81525f6004820152602401610409565b6001600160a01b038085165f90815260036020908152604080832093871683529290522082905580156106de57826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161063d91815260200190565b6001600160a01b03831661089b578060045f8282546108909190610d1a565b9091555061090b9050565b6001600160a01b0383165f90815260026020526040902054818110156108ed5760405163391434e360e21b81526001600160a01b03851660048201526024810182905260448101839052606401610409565b6001600160a01b0384165f9081526002602052604090209082900390555b6001600160a01b03821661092757600480548290039055610945565b6001600160a01b0382165f9081526002602052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161098a91815260200190565b60405180910390a3505050565b5f6020808352835180828501525f5b818110156109c2578581018301518582016040015282016109a6565b505f604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b03811681146109f8575f80fd5b919050565b5f8060408385031215610a0e575f80fd5b610a17836109e2565b946020939093013593505050565b5f805f60608486031215610a37575f80fd5b610a40846109e2565b9250610a4e602085016109e2565b9150604084013590509250925092565b5f8083601f840112610a6e575f80fd5b50813567ffffffffffffffff811115610a85575f80fd5b602083019150836020828501011115610a9c575f80fd5b9250929050565b5f805f805f60608688031215610ab7575f80fd5b85359450602086013567ffffffffffffffff80821115610ad5575f80fd5b610ae189838a01610a5e565b90965094506040880135915080821115610af9575f80fd5b50610b0688828901610a5e565b969995985093965092949392505050565b5f60208284031215610b27575f80fd5b610b30826109e2565b9392505050565b5f805f60408486031215610b49575f80fd5b83359250602084013567ffffffffffffffff811115610b66575f80fd5b610b7286828701610a5e565b9497909650939450505050565b5f8060408385031215610b90575f80fd5b610b99836109e2565b9150610ba7602084016109e2565b90509250929050565b600181811c90821680610bc457607f821691505b602082108103610be257634e487b7160e01b5f52602260045260245ffd5b50919050565b5f82610c0257634e487b7160e01b5f52601260045260245ffd5b500490565b60208082526073908201527f41737472696142726964676561626c6545524332303a20696e7375666669636960408201527f656e742076616c75652c206d7573742062652067726561746572207468616e2060608201527f3130202a2a2028544f4b454e5f444543494d414c53202d20424153455f434841608082015272494e5f41535345545f505245434953494f4e2960681b60a082015260c00190565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b604081525f610ce1604083018688610ca6565b8281036020840152610cf4818587610ca6565b979650505050505050565b602081525f610d12602083018486610ca6565b949350505050565b8082018082111561035857634e487b7160e01b5f52601160045260245ffdfea2646970667358221220839fde846bdf8b562d3e7e8b39bc0acde2c69b1f2a7bba30ec27b96ba437f6a564736f6c63430008150033",
}

// AstriaBridgeableERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use AstriaBridgeableERC20MetaData.ABI instead.
var AstriaBridgeableERC20ABI = AstriaBridgeableERC20MetaData.ABI

// AstriaBridgeableERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AstriaBridgeableERC20MetaData.Bin instead.
var AstriaBridgeableERC20Bin = AstriaBridgeableERC20MetaData.Bin

// DeployAstriaBridgeableERC20 deploys a new Ethereum contract, binding an instance of AstriaBridgeableERC20 to it.
func DeployAstriaBridgeableERC20(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address, _baseChainAssetPrecision uint32, _baseChainBridgeAddress string, _baseChainAssetDenomination string, _name string, _symbol string) (common.Address, *types.Transaction, *AstriaBridgeableERC20, error) {
	parsed, err := AstriaBridgeableERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AstriaBridgeableERC20Bin), backend, _bridge, _baseChainAssetPrecision, _baseChainBridgeAddress, _baseChainAssetDenomination, _name, _symbol)
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

// BASECHAINASSETDENOMINATION is a free data retrieval call binding the contract method 0xb6476c7e.
//
// Solidity: function BASE_CHAIN_ASSET_DENOMINATION() view returns(string)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Caller) BASECHAINASSETDENOMINATION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AstriaBridgeableERC20.contract.Call(opts, &out, "BASE_CHAIN_ASSET_DENOMINATION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BASECHAINASSETDENOMINATION is a free data retrieval call binding the contract method 0xb6476c7e.
//
// Solidity: function BASE_CHAIN_ASSET_DENOMINATION() view returns(string)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Session) BASECHAINASSETDENOMINATION() (string, error) {
	return _AstriaBridgeableERC20.Contract.BASECHAINASSETDENOMINATION(&_AstriaBridgeableERC20.CallOpts)
}

// BASECHAINASSETDENOMINATION is a free data retrieval call binding the contract method 0xb6476c7e.
//
// Solidity: function BASE_CHAIN_ASSET_DENOMINATION() view returns(string)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20CallerSession) BASECHAINASSETDENOMINATION() (string, error) {
	return _AstriaBridgeableERC20.Contract.BASECHAINASSETDENOMINATION(&_AstriaBridgeableERC20.CallOpts)
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

// BASECHAINBRIDGEADDRESS is a free data retrieval call binding the contract method 0xdb97dc98.
//
// Solidity: function BASE_CHAIN_BRIDGE_ADDRESS() view returns(string)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Caller) BASECHAINBRIDGEADDRESS(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AstriaBridgeableERC20.contract.Call(opts, &out, "BASE_CHAIN_BRIDGE_ADDRESS")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BASECHAINBRIDGEADDRESS is a free data retrieval call binding the contract method 0xdb97dc98.
//
// Solidity: function BASE_CHAIN_BRIDGE_ADDRESS() view returns(string)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Session) BASECHAINBRIDGEADDRESS() (string, error) {
	return _AstriaBridgeableERC20.Contract.BASECHAINBRIDGEADDRESS(&_AstriaBridgeableERC20.CallOpts)
}

// BASECHAINBRIDGEADDRESS is a free data retrieval call binding the contract method 0xdb97dc98.
//
// Solidity: function BASE_CHAIN_BRIDGE_ADDRESS() view returns(string)
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20CallerSession) BASECHAINBRIDGEADDRESS() (string, error) {
	return _AstriaBridgeableERC20.Contract.BASECHAINBRIDGEADDRESS(&_AstriaBridgeableERC20.CallOpts)
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

// WithdrawToSequencer is a paid mutator transaction binding the contract method 0xd38fe9a7.
//
// Solidity: function withdrawToSequencer(uint256 _amount, string _destinationChainAddress) returns()
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Transactor) WithdrawToSequencer(opts *bind.TransactOpts, _amount *big.Int, _destinationChainAddress string) (*types.Transaction, error) {
	return _AstriaBridgeableERC20.contract.Transact(opts, "withdrawToSequencer", _amount, _destinationChainAddress)
}

// WithdrawToSequencer is a paid mutator transaction binding the contract method 0xd38fe9a7.
//
// Solidity: function withdrawToSequencer(uint256 _amount, string _destinationChainAddress) returns()
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20Session) WithdrawToSequencer(_amount *big.Int, _destinationChainAddress string) (*types.Transaction, error) {
	return _AstriaBridgeableERC20.Contract.WithdrawToSequencer(&_AstriaBridgeableERC20.TransactOpts, _amount, _destinationChainAddress)
}

// WithdrawToSequencer is a paid mutator transaction binding the contract method 0xd38fe9a7.
//
// Solidity: function withdrawToSequencer(uint256 _amount, string _destinationChainAddress) returns()
func (_AstriaBridgeableERC20 *AstriaBridgeableERC20TransactorSession) WithdrawToSequencer(_amount *big.Int, _destinationChainAddress string) (*types.Transaction, error) {
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
	DestinationChainAddress string
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterSequencerWithdrawal is a free log retrieval operation binding the contract event 0x0f4961cab7530804898499aa89f5ec81d1a73102e2e4a1f30f88e5ae3513ba2a.
//
// Solidity: event SequencerWithdrawal(address indexed sender, uint256 indexed amount, string destinationChainAddress)
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

// WatchSequencerWithdrawal is a free log subscription operation binding the contract event 0x0f4961cab7530804898499aa89f5ec81d1a73102e2e4a1f30f88e5ae3513ba2a.
//
// Solidity: event SequencerWithdrawal(address indexed sender, uint256 indexed amount, string destinationChainAddress)
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

// ParseSequencerWithdrawal is a log parse operation binding the contract event 0x0f4961cab7530804898499aa89f5ec81d1a73102e2e4a1f30f88e5ae3513ba2a.
//
// Solidity: event SequencerWithdrawal(address indexed sender, uint256 indexed amount, string destinationChainAddress)
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
