package live

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/tracers"
	"github.com/ethereum/go-ethereum/params"
)

func init() {
	tracers.LiveDirectory.Register("tiatracer", newTiaTracer)
}

// noop is a no-op live tracer. It's there to
// catch changes in the tracing interface, as well as
// for testing live tracing performance. Can be removed
// as soon as we have a real live tracer.
type TiaTracer struct {
	TotalDeposited       *big.Int
	TotalWithdrawn       *big.Int
	TotalDepositedStaged *big.Int
	TotalWithdrawnStaged *big.Int
}

func newTiaTracer(_ json.RawMessage) (*tracing.Hooks, error) {
	t := &TiaTracer{
		TotalDeposited:       big.NewInt(0),
		TotalWithdrawn:       big.NewInt(0),
		TotalDepositedStaged: big.NewInt(0),
		TotalWithdrawnStaged: big.NewInt(0),
	}

	return &tracing.Hooks{
		OnTxStart:        t.OnTxStart,
		OnTxEnd:          t.OnTxEnd,
		OnEnter:          t.OnEnter,
		OnExit:           t.OnExit,
		OnOpcode:         t.OnOpcode,
		OnFault:          t.OnFault,
		OnGasChange:      t.OnGasChange,
		OnBlockchainInit: t.OnBlockchainInit,
		OnBlockStart:     t.OnBlockStart,
		OnBlockEnd:       t.OnBlockEnd,
		OnSkippedBlock:   t.OnSkippedBlock,
		OnGenesisBlock:   t.OnGenesisBlock,
		OnBalanceChange:  t.OnBalanceChange,
		OnNonceChange:    t.OnNonceChange,
		OnCodeChange:     t.OnCodeChange,
		OnStorageChange:  t.OnStorageChange,
		OnLog:            t.OnLog,
	}, nil
}

func (t *TiaTracer) OnOpcode(pc uint64, op byte, gas, cost uint64, scope tracing.OpContext, rData []byte, depth int, err error) {
}

func (t *TiaTracer) OnFault(pc uint64, op byte, gas, cost uint64, _ tracing.OpContext, depth int, err error) {
}

func (t *TiaTracer) OnEnter(depth int, typ byte, from common.Address, to common.Address, input []byte, gas uint64, value *big.Int) {
}

func (t *TiaTracer) OnExit(depth int, output []byte, gasUsed uint64, err error, reverted bool) {
}

func (t *TiaTracer) OnTxStart(vm *tracing.VMContext, tx *types.Transaction, from common.Address) {
}

func (t *TiaTracer) OnTxEnd(receipt *types.Receipt, err error) {}

func (t *TiaTracer) OnBlockStart(ev tracing.BlockEvent) {
	// Compute the increase and decrease in bridged asset here
	// These values are staged over here

	for _, tx := range ev.Block.Transactions() {
		// TODO - include erc20 deposit and withdrawals too here
		if tx.Type() == types.DepositTxType {
			t.TotalDepositedStaged = t.TotalDepositedStaged.Add(t.TotalDepositedStaged, tx.Value())
		}
	}
}

func (t *TiaTracer) OnBlockEnd(err error) {
	// if err == nil, then we need to commit the values we have staged.
	if err == nil {
		t.TotalDeposited = t.TotalDeposited.Add(t.TotalDeposited, t.TotalDepositedStaged)
		t.TotalWithdrawn = t.TotalWithdrawn.Add(t.TotalWithdrawn, t.TotalWithdrawnStaged)

		// TODO - we should sent this to prometheus? or can we?
		// if multiple
	}

	t.TotalDepositedStaged = big.NewInt(0)
	t.TotalWithdrawnStaged = big.NewInt(0)
}

func (t *TiaTracer) OnSkippedBlock(ev tracing.BlockEvent) {}

func (t *TiaTracer) OnBlockchainInit(chainConfig *params.ChainConfig) {
}

func (t *TiaTracer) OnGenesisBlock(b *types.Block, alloc types.GenesisAlloc) {
}

func (t *TiaTracer) OnBalanceChange(a common.Address, prev, new *big.Int, reason tracing.BalanceChangeReason) {
}

func (t *TiaTracer) OnNonceChange(a common.Address, prev, new uint64) {
}

func (t *TiaTracer) OnCodeChange(a common.Address, prevCodeHash common.Hash, prev []byte, codeHash common.Hash, code []byte) {
}

func (t *TiaTracer) OnStorageChange(a common.Address, k, prev, new common.Hash) {
}

func (t *TiaTracer) OnLog(l *types.Log) {

}

func (t *TiaTracer) OnGasChange(old, new uint64, reason tracing.GasChangeReason) {
}
