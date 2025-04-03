package precompile

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/holiman/uint256"
)

type statefulContext struct {
	state     StateDB
	address   common.Address
	msgSender common.Address
	msgValue  *uint256.Int
	readOnly  bool
}

func NewStatefulContext(
	state StateDB,
	address common.Address,
	msgSender common.Address,
	msgValue *uint256.Int,
) StatefulContext {
	return &statefulContext{
		state:     state,
		address:   address,
		msgSender: msgSender,
		msgValue:  msgValue,
		readOnly:  false,
	}
}

func (sc *statefulContext) SetState(key common.Hash, value common.Hash) error {
	if sc.readOnly {
		return ErrWriteProtection
	}
	sc.state.SetState(sc.address, key, value)
	return nil
}

func (sc *statefulContext) GetState(key common.Hash) common.Hash {
	return sc.state.GetState(sc.address, key)
}

func (sc *statefulContext) GetCommittedState(key common.Hash) common.Hash {
	return sc.state.GetCommittedState(sc.address, key)
}

func (sc *statefulContext) SubBalance(address common.Address, amount *uint256.Int, reason tracing.BalanceChangeReason) error {
	if sc.readOnly {
		return ErrWriteProtection
	}
	sc.state.SubBalance(address, amount, reason)
	return nil
}

func (sc *statefulContext) AddBalance(address common.Address, amount *uint256.Int, reason tracing.BalanceChangeReason) error {
	if sc.readOnly {
		return ErrWriteProtection
	}
	sc.state.AddBalance(address, amount, reason)
	return nil
}

func (sc *statefulContext) GetBalance(address common.Address) *uint256.Int {
	return sc.state.GetBalance(address)
}

func (sc *statefulContext) Address() common.Address {
	return sc.address
}

func (sc *statefulContext) MsgSender() common.Address {
	return sc.msgSender
}

func (sc *statefulContext) MsgValue() *uint256.Int {
	return sc.msgValue
}

func (sc *statefulContext) IsReadOnly() bool {
	return sc.readOnly
}

func (sc *statefulContext) SetReadOnly(readOnly bool) {
	sc.readOnly = readOnly
}
