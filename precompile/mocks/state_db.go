package mocks

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/precompile"
	"github.com/holiman/uint256"
)

type mockStateDB struct {
	balances map[common.Address]*uint256.Int
	states   map[common.Address]map[common.Hash]common.Hash
}

func NewMockStateDB() precompile.StateDB {
	return &mockStateDB{
		balances: make(map[common.Address]*uint256.Int),
		states:   make(map[common.Address]map[common.Hash]common.Hash),
	}
}

func (m *mockStateDB) SubBalance(address common.Address, amount *uint256.Int, reason tracing.BalanceChangeReason) {
	if _, ok := m.balances[address]; !ok {
		m.balances[address] = uint256.NewInt(0)
	}
	m.balances[address].Sub(m.balances[address], amount)
}

func (m *mockStateDB) AddBalance(address common.Address, amount *uint256.Int, reason tracing.BalanceChangeReason) {
	if _, ok := m.balances[address]; !ok {
		m.balances[address] = uint256.NewInt(0)
	}
	m.balances[address].Add(m.balances[address], amount)
}

func (m *mockStateDB) GetBalance(address common.Address) *uint256.Int {
	if _, ok := m.balances[address]; !ok {
		m.balances[address] = uint256.NewInt(0)
	}
	return new(uint256.Int).Set(m.balances[address])
}

func (m *mockStateDB) GetState(address common.Address, hash common.Hash) common.Hash {
	if _, ok := m.states[address]; !ok {
		m.states[address] = make(map[common.Hash]common.Hash)
	}
	return m.states[address][hash]
}

func (m *mockStateDB) SetState(address common.Address, hash common.Hash, value common.Hash) {
	if _, ok := m.states[address]; !ok {
		m.states[address] = make(map[common.Hash]common.Hash)
	}
	m.states[address][hash] = value
}

func (m *mockStateDB) GetCommittedState(address common.Address, hash common.Hash) common.Hash {
	if _, ok := m.states[address]; !ok {
		m.states[address] = make(map[common.Hash]common.Hash)
	}
	return m.states[address][hash]
}
