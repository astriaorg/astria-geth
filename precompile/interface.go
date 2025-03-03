package precompile

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/holiman/uint256"
)

type BalanceChangeReason byte

type StateDB interface {
	SubBalance(common.Address, *uint256.Int, tracing.BalanceChangeReason) uint256.Int
	AddBalance(common.Address, *uint256.Int, tracing.BalanceChangeReason) uint256.Int
	GetBalance(common.Address) *uint256.Int
	GetState(common.Address, common.Hash) common.Hash
	SetState(common.Address, common.Hash, common.Hash) common.Hash
	GetCommittedState(common.Address, common.Hash) common.Hash
}

type StatefulContext interface {
	SetState(common.Hash, common.Hash) error
	GetState(common.Hash) common.Hash
	GetCommittedState(common.Hash) common.Hash
	SubBalance(common.Address, *uint256.Int, tracing.BalanceChangeReason) error
	AddBalance(common.Address, *uint256.Int, tracing.BalanceChangeReason) error
	GetBalance(common.Address) *uint256.Int
	Address() common.Address
	MsgSender() common.Address
	MsgValue() *uint256.Int
	IsReadOnly() bool
	SetReadOnly(bool)
}

type StatefulPrecompiledContract interface {
	GetABI() abi.ABI
	DefaultGas(input []byte) uint64
}

type PrecompileMap map[common.Address]StatefulPrecompiledContract
