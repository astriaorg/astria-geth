package precompile_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/precompile"
	"github.com/ethereum/go-ethereum/precompile/mocks"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/assert"
)

func setupStatefulContext() precompile.StatefulContext {
	state := mocks.NewMockStateDB()
	address := common.HexToAddress("0x123")
	msgSender := common.HexToAddress("0x456")
	msgValue := uint256.NewInt(1000)

	return precompile.NewStatefulContext(state, address, msgSender, msgValue)
}

func TestAddress(t *testing.T) {
	ctx := setupStatefulContext()
	assert.Equal(t, common.HexToAddress("0x123"), ctx.Address())
}

func TestMsgSender(t *testing.T) {
	ctx := setupStatefulContext()
	assert.Equal(t, common.HexToAddress("0x456"), ctx.MsgSender())
}

func TestMsgValue(t *testing.T) {
	ctx := setupStatefulContext()
	assert.Equal(t, uint256.NewInt(1000), ctx.MsgValue())
}

func TestIsReadOnly(t *testing.T) {
	ctx := setupStatefulContext()

	assert.False(t, ctx.IsReadOnly())

	ctx.SetReadOnly(true)
	assert.True(t, ctx.IsReadOnly())
}

func TestSetState(t *testing.T) {
	ctx := setupStatefulContext()

	key := common.HexToHash("0x789")
	value := common.HexToHash("0xabc")

	assert.Equal(t, common.HexToHash("0x00"), ctx.GetState(key))

	ctx.SetReadOnly(true)
	err := ctx.SetState(key, value)
	assert.Error(t, err)

	ctx.SetReadOnly(false)
	err = ctx.SetState(key, value)
	assert.NoError(t, err)

	assert.Equal(t, value, ctx.GetState(key))
}

func TestBalance(t *testing.T) {
	ctx := setupStatefulContext()

	initialBalance := ctx.GetBalance(common.HexToAddress("0x123"))
	assert.Equal(t, uint256.NewInt(0), initialBalance)

	amount := uint256.NewInt(500)

	err := ctx.AddBalance(common.HexToAddress("0x123"), amount, tracing.BalanceChangeUnspecified)
	assert.NoError(t, err)
	assert.Equal(t, uint256.NewInt(500), ctx.GetBalance(common.HexToAddress("0x123")))

	err = ctx.AddBalance(common.HexToAddress("0x123"), amount, tracing.BalanceChangeUnspecified)
	assert.NoError(t, err)
	assert.Equal(t, uint256.NewInt(1000), ctx.GetBalance(common.HexToAddress("0x123")))

	err = ctx.SubBalance(common.HexToAddress("0x123"), amount, tracing.BalanceChangeUnspecified)
	assert.NoError(t, err)
	assert.Equal(t, uint256.NewInt(500), ctx.GetBalance(common.HexToAddress("0x123")))

	ctx.SetReadOnly(true)

	err = ctx.AddBalance(common.HexToAddress("0x123"), amount, tracing.BalanceChangeUnspecified)
	assert.Error(t, err)

	err = ctx.SubBalance(common.HexToAddress("0x123"), amount, tracing.BalanceChangeUnspecified)
	assert.Error(t, err)
}
