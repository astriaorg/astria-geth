// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package miner

import (
	"errors"
	"fmt"
	"math/big"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/misc/eip1559"
	"github.com/ethereum/go-ethereum/consensus/misc/eip4844"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

var (
	errBlockInterruptedByNewHead  = errors.New("new head arrived while building block")
	errBlockInterruptedByRecommit = errors.New("recommit interrupt while building block")
	errBlockInterruptedByTimeout  = errors.New("timeout while building block")
)

// environment is the worker's current environment and holds all
// information of the sealing block generation.
type environment struct {
	signer   types.Signer
	state    *state.StateDB // apply state changes here
	tcount   int            // tx count in cycle
	gasPool  *core.GasPool  // available gas used to pack transactions
	coinbase common.Address

	header   *types.Header
	txs      []*types.Transaction
	receipts []*types.Receipt
	sidecars []*types.BlobTxSidecar
	blobs    int
}

const (
	commitInterruptNone int32 = iota
	commitInterruptNewHead
	commitInterruptResubmit
	commitInterruptTimeout
)

// newPayloadResult is the result of payload generation.
type newPayloadResult struct {
	err      error
	block    *types.Block
	fees     *big.Int               // total block fees
	sidecars []*types.BlobTxSidecar // collected blobs of blob transactions
	stateDB  *state.StateDB         // StateDB after executing the transactions
	receipts []*types.Receipt       // Receipts collected during construction
}

// generateParams wraps various of settings for generating sealing task.
type generateParams struct {
	timestamp   uint64            // The timstamp for sealing task
	forceTime   bool              // Flag whether the given timestamp is immutable or not
	parentHash  common.Hash       // Parent block hash, empty means the latest chain head
	coinbase    common.Address    // The fee recipient address for including transaction
	random      common.Hash       // The randomness generated by beacon chain, empty before the merge
	withdrawals types.Withdrawals // List of withdrawals to include in block.
	beaconRoot  *common.Hash      // The beacon root (cancun field).
	noTxs       bool              // Flag whether an empty block without any transaction is expected
}

// prepareWork constructs the sealing task according to the given parameters,
// either based on the last chain head or specified parent. In this function
// the pending transactions are not filled yet, only the empty task returned.
func (miner *Miner) prepareWork(genParams *generateParams) (*environment, error) {
	miner.confMu.RLock()
	defer miner.confMu.RUnlock()

	// Find the parent block for sealing task
	parent := miner.chain.CurrentBlock()
	if genParams.parentHash != (common.Hash{}) {
		block := miner.chain.GetBlockByHash(genParams.parentHash)
		if block == nil {
			return nil, errors.New("missing parent")
		}
		parent = block.Header()
	}
	// Sanity check the timestamp correctness, recap the timestamp
	// to parent+1 if the mutation is allowed.
	timestamp := genParams.timestamp
	if parent.Time >= timestamp {
		if genParams.forceTime {
			return nil, fmt.Errorf("invalid timestamp, parent %d given %d", parent.Time, timestamp)
		}
		timestamp = parent.Time + 1
	}
	// Construct the sealing block header.
	header := &types.Header{
		ParentHash: parent.Hash(),
		Number:     new(big.Int).Add(parent.Number, common.Big1),
		GasLimit:   core.CalcGasLimit(parent.GasLimit, miner.config.GasCeil),
		Time:       timestamp,
		Coinbase:   genParams.coinbase,
	}
	// Set the extra field.
	if len(miner.config.ExtraData) != 0 {
		header.Extra = miner.config.ExtraData
	}
	// Set the randomness field from the beacon chain if it's available.
	if genParams.random != (common.Hash{}) {
		header.MixDigest = genParams.random
	}
	// Set baseFee and GasLimit if we are on an EIP-1559 chain
	if miner.chainConfig.IsLondon(header.Number) {
		header.BaseFee = eip1559.CalcBaseFee(miner.chainConfig, parent)
		if !miner.chainConfig.IsLondon(parent.Number) {
			parentGasLimit := parent.GasLimit * miner.chainConfig.GetAstriaForks().ElasticityMultiplierAt(parent.Number.Uint64())
			header.GasLimit = core.CalcGasLimit(parentGasLimit, miner.config.GasCeil)
		}
	}
	// Run the consensus preparation with the default or customized consensus engine.
	// Note that the `header.Time` may be changed.
	if err := miner.engine.Prepare(miner.chain, header); err != nil {
		log.Error("Failed to prepare header for sealing", "err", err)
		return nil, err
	}
	// Apply EIP-4844, EIP-4788.
	if miner.chainConfig.IsCancun(header.Number, header.Time) {
		var excessBlobGas uint64
		if miner.chainConfig.IsCancun(parent.Number, parent.Time) {
			excessBlobGas = eip4844.CalcExcessBlobGas(*parent.ExcessBlobGas, *parent.BlobGasUsed)
		} else {
			// For the first post-fork block, both parent.data_gas_used and parent.excess_data_gas are evaluated as 0
			excessBlobGas = eip4844.CalcExcessBlobGas(0, 0)
		}
		header.BlobGasUsed = new(uint64)
		header.ExcessBlobGas = &excessBlobGas
		header.ParentBeaconRoot = genParams.beaconRoot
	}
	// Could potentially happen if starting to mine in an odd state.
	// Note genParams.coinbase can be different with header.Coinbase
	// since clique algorithm can modify the coinbase field in header.
	env, err := miner.makeEnv(parent, header, genParams.coinbase)
	if err != nil {
		log.Error("Failed to create sealing context", "err", err)
		return nil, err
	}
	if header.ParentBeaconRoot != nil {
		context := core.NewEVMBlockContext(header, miner.chain, nil)
		vmenv := vm.NewEVM(context, vm.TxContext{}, env.state, miner.chainConfig, vm.Config{})
		core.ProcessBeaconBlockRoot(*header.ParentBeaconRoot, vmenv, env.state)
	}
	return env, nil
}

// makeEnv creates a new environment for the sealing block.
func (miner *Miner) makeEnv(parent *types.Header, header *types.Header, coinbase common.Address) (*environment, error) {
	// Retrieve the parent state to execute on top and start a prefetcher for
	// the miner to speed block sealing up a bit.
	state, err := miner.chain.StateAt(parent.Root)
	if err != nil {
		return nil, err
	}
	// Note the passed coinbase may be different with header.Coinbase.
	return &environment{
		signer:   types.MakeSigner(miner.chainConfig, header.Number, header.Time),
		state:    state,
		coinbase: coinbase,
		header:   header,
	}, nil
}

func (miner *Miner) commitTransaction(env *environment, tx *types.Transaction) error {
	if tx.Type() == types.BlobTxType {
		return miner.commitBlobTransaction(env, tx)
	}
	receipt, err := miner.applyTransaction(env, tx)
	if err != nil {
		return err
	}
	env.txs = append(env.txs, tx)
	env.receipts = append(env.receipts, receipt)
	env.tcount++
	return nil
}

func (miner *Miner) commitBlobTransaction(env *environment, tx *types.Transaction) error {
	sc := tx.BlobTxSidecar()
	if sc == nil {
		panic("blob transaction without blobs in miner")
	}
	// Checking against blob gas limit: It's kind of ugly to perform this check here, but there
	// isn't really a better place right now. The blob gas limit is checked at block validation time
	// and not during execution. This means core.ApplyTransaction will not return an error if the
	// tx has too many blobs. So we have to explicitly check it here.
	if (env.blobs+len(sc.Blobs))*params.BlobTxBlobGasPerBlob > params.MaxBlobGasPerBlock {
		return errors.New("max data blobs reached")
	}
	receipt, err := miner.applyTransaction(env, tx)
	if err != nil {
		return err
	}
	env.txs = append(env.txs, tx.WithoutBlobTxSidecar())
	env.receipts = append(env.receipts, receipt)
	env.sidecars = append(env.sidecars, sc)
	env.blobs += len(sc.Blobs)
	*env.header.BlobGasUsed += receipt.BlobGasUsed
	env.tcount++
	return nil
}

// applyTransaction runs the transaction. If execution fails, state and gas pool are reverted.
func (miner *Miner) applyTransaction(env *environment, tx *types.Transaction) (*types.Receipt, error) {
	var (
		snap = env.state.Snapshot()
		gp   = env.gasPool.Gas()
	)
	receipt, err := core.ApplyTransaction(miner.chainConfig, miner.chain, &env.coinbase, env.gasPool, env.state, env.header, tx, &env.header.GasUsed, vm.Config{})
	if err != nil {
		env.state.RevertToSnapshot(snap)
		env.gasPool.SetGas(gp)
	}
	return receipt, err
}

// This is a copy of commitTransactions, but updated to take a list of txs instead of using heap
func (miner *Miner) commitAstriaTransactions(env *environment, txs *types.Transactions, interrupt *atomic.Int32) error {
	gasLimit := env.header.GasLimit
	if env.gasPool == nil {
		env.gasPool = new(core.GasPool).AddGas(gasLimit)
	}

	for i, tx := range *txs {
		// Check interruption signal and abort building if it's fired.
		if interrupt != nil {
			if signal := interrupt.Load(); signal != commitInterruptNone {
				// remove the subsequent txs from the mempool if block building has been interrupted
				for _, txToRemove := range (*txs)[i:] {
					miner.txpool.AddToAstriaExcludedFromBlock(txToRemove)
				}
				return signalToErr(signal)
			}
		}
		// If we don't have enough gas for any further transactions then we're done.
		if env.gasPool.Gas() < params.TxGas && tx.Type() != types.InjectedTxType {
			log.Trace("Not enough gas for further transactions", "have", env.gasPool, "want", params.TxGas)
			// remove txs from the mempool if they are too big for this block
			for _, txToRemove := range (*txs)[i:] {
				miner.txpool.AddToAstriaExcludedFromBlock(txToRemove)
			}
			break
		}

		// Error may be ignored here. The error has already been checked
		// during transaction acceptance is the transaction pool.
		from, _ := types.Sender(env.signer, tx)

		// Check whether the tx is replay protected. If we're not in the EIP155 hf
		// phase, start ignoring the sender until we do.
		if tx.Protected() && !miner.chainConfig.IsEIP155(env.header.Number) {
			log.Trace("Ignoring reply protected transaction", "hash", tx.Hash(), "eip155", miner.chainConfig.EIP155Block)
			miner.txpool.AddToAstriaExcludedFromBlock(tx)
			continue
		}
		// Start executing the transaction
		env.state.SetTxContext(tx.Hash(), env.tcount)

		err := miner.commitTransaction(env, tx)
		switch {
		case errors.Is(err, core.ErrNonceTooLow):
			// New head notification data race between the transaction pool and miner, shift
			log.Trace("Skipping transaction with low nonce", "sender", from, "nonce", tx.Nonce())
		case errors.Is(err, nil):
			// Transaction successfully applied, continue
		default:
			// Strange error, discard the transaction and get the next in line (note, the
			// nonce-too-high clause will prevent us from executing in vain).
			log.Debug("Transaction failed, account skipped", "hash", tx.Hash(), "err", err)
		}
		if err != nil {
			log.Trace("Marking transaction as invalid", "hash", tx.Hash(), "err", err)
			miner.txpool.AddToAstriaExcludedFromBlock(tx)
		}
	}

	return nil
}

func (miner *Miner) fillAstriaTransactions(interrupt *atomic.Int32, env *environment) error {
	// Use pre ordered array of txs
	astriaTxs := miner.txpool.AstriaOrdered()
	if len(*astriaTxs) > 0 {
		if err := miner.commitAstriaTransactions(env, astriaTxs, interrupt); err != nil {
			return err
		}
	}

	return nil
}

// generateWork generates a sealing block based on the given parameters.
func (miner *Miner) generateWork(params *generateParams) *newPayloadResult {
	work, err := miner.prepareWork(params)
	if err != nil {
		return &newPayloadResult{err: err}
	}
	if !params.noTxs {
		interrupt := new(atomic.Int32)

		err := miner.fillAstriaTransactions(interrupt, work)
		if errors.Is(err, errBlockInterruptedByTimeout) {
			log.Error("Block building is interrupted", "allowance", common.PrettyDuration(miner.config.Recommit))
		}
	}
	body := types.Body{Transactions: work.txs, Withdrawals: params.withdrawals}
	block, err := miner.engine.FinalizeAndAssemble(miner.chain, work.header, work.state, &body, work.receipts)
	if err != nil {
		return &newPayloadResult{err: err}
	}
	return &newPayloadResult{
		block:    block,
		fees:     totalFees(block, work.receipts),
		sidecars: work.sidecars,
		stateDB:  work.state,
		receipts: work.receipts,
	}
}

// totalFees computes total consumed miner fees in Wei. Block transactions and receipts have to have the same order.
func totalFees(block *types.Block, receipts []*types.Receipt) *big.Int {
	feesWei := new(big.Int)
	for i, tx := range block.Transactions() {
		minerFee, _ := tx.EffectiveGasTip(block.BaseFee())
		feesWei.Add(feesWei, new(big.Int).Mul(new(big.Int).SetUint64(receipts[i].GasUsed), minerFee))
	}
	return feesWei
}

// signalToErr converts the interruption signal to a concrete error type for return.
// The given signal must be a valid interruption signal.
func signalToErr(signal int32) error {
	switch signal {
	case commitInterruptNewHead:
		return errBlockInterruptedByNewHead
	case commitInterruptResubmit:
		return errBlockInterruptedByRecommit
	case commitInterruptTimeout:
		return errBlockInterruptedByTimeout
	default:
		panic(fmt.Errorf("undefined signal %d", signal))
	}
}
