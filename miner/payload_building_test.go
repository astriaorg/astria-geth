// Copyright 2022 The go-ethereum Authors
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
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>

package miner

import (
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

func TestBuildPayload(t *testing.T) {
	var (
		db        = rawdb.NewMemoryDatabase()
		recipient = common.HexToAddress("0xdeadbeef")
	)
	w, b := newTestWorker(t, params.TestChainConfig, ethash.NewFaker(), db, 0)
	defer w.close()

	timestamp := uint64(time.Now().Unix())

	verify := func(outer *engine.ExecutionPayloadEnvelope, txs int) {
		payload := outer.ExecutionPayload
		if payload.ParentHash != b.chain.CurrentBlock().Hash() {
			t.Fatal("Unexpect parent hash")
		}
		if payload.Random != (common.Hash{}) {
			t.Fatal("Unexpect random value")
		}
		if payload.Timestamp != timestamp {
			t.Fatal("Unexpect timestamp")
		}
		if payload.FeeRecipient != recipient {
			t.Fatal("Unexpect fee recipient")
		}
		if len(payload.Transactions) != txs {
			t.Fatal("Unexpect transaction set")
		}
	}

	txGasPrice := big.NewInt(10 * params.InitialBaseFee)

	tests := []struct {
		name                 string
		txsToBuildPayload    types.Transactions
		expectedTxsInPayload types.Transactions
		invalidTxs           types.Transactions
	}{
		{
			name:                 "empty",
			txsToBuildPayload:    types.Transactions{},
			expectedTxsInPayload: types.Transactions{},
			invalidTxs:           types.Transactions{},
		},
		{
			name: "transactions with gas enough to fit into a single block",
			txsToBuildPayload: types.Transactions{
				types.NewTransaction(b.txPool.Nonce(testBankAddress), testUserAddress, big.NewInt(1000), params.TxGas, txGasPrice, nil),
				types.NewTransaction(b.txPool.Nonce(testBankAddress)+1, testUserAddress, big.NewInt(2000), params.TxGas, txGasPrice, nil),
			},
			expectedTxsInPayload: types.Transactions{
				types.NewTransaction(b.txPool.Nonce(testBankAddress), testUserAddress, big.NewInt(1000), params.TxGas, txGasPrice, nil),
				types.NewTransaction(b.txPool.Nonce(testBankAddress)+1, testUserAddress, big.NewInt(2000), params.TxGas, txGasPrice, nil),
			},
			invalidTxs: types.Transactions{},
		},
		{
			name: "transactions with gas which doesn't fit in a single block",
			txsToBuildPayload: types.Transactions{
				types.NewTransaction(b.txPool.Nonce(testBankAddress), testUserAddress, big.NewInt(1000), b.BlockChain().GasLimit()-10000, txGasPrice, nil),
				types.NewTransaction(b.txPool.Nonce(testBankAddress)+1, testUserAddress, big.NewInt(1000), b.BlockChain().GasLimit()-10000, txGasPrice, nil),
			},
			expectedTxsInPayload: types.Transactions{
				types.NewTransaction(b.txPool.Nonce(testBankAddress), testUserAddress, big.NewInt(1000), b.BlockChain().GasLimit()-10000, txGasPrice, nil),
			},
			invalidTxs: types.Transactions{
				types.NewTransaction(b.txPool.Nonce(testBankAddress)+1, testUserAddress, big.NewInt(1000), b.BlockChain().GasLimit()-10000, txGasPrice, nil),
			},
		},
		{
			name: "transactions with nonce too high",
			txsToBuildPayload: types.Transactions{
				types.NewTransaction(b.txPool.Nonce(testBankAddress), testUserAddress, big.NewInt(1000), params.TxGas, txGasPrice, nil),
				types.NewTransaction(b.txPool.Nonce(testBankAddress)+4, testUserAddress, big.NewInt(2000), params.TxGas, txGasPrice, nil),
			},
			expectedTxsInPayload: types.Transactions{
				types.NewTransaction(b.txPool.Nonce(testBankAddress), testUserAddress, big.NewInt(1000), params.TxGas, txGasPrice, nil),
			},
			invalidTxs: types.Transactions{
				types.NewTransaction(b.txPool.Nonce(testBankAddress)+4, testUserAddress, big.NewInt(2000), params.TxGas, txGasPrice, nil),
			},
		},
		{
			name: "transactions with nonce too low",
			txsToBuildPayload: types.Transactions{
				types.NewTransaction(b.txPool.Nonce(testBankAddress), testUserAddress, big.NewInt(1000), params.TxGas, txGasPrice, nil),
				types.NewTransaction(b.txPool.Nonce(testBankAddress)-1, testUserAddress, big.NewInt(2000), params.TxGas, txGasPrice, nil),
			},
			expectedTxsInPayload: types.Transactions{
				types.NewTransaction(b.txPool.Nonce(testBankAddress), testUserAddress, big.NewInt(1000), params.TxGas, txGasPrice, nil),
			},
			invalidTxs: types.Transactions{
				types.NewTransaction(b.txPool.Nonce(testBankAddress)-1, testUserAddress, big.NewInt(2000), params.TxGas, txGasPrice, nil),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			signedTxs := types.Transactions{}
			signedInvalidTxs := types.Transactions{}

			for _, tx := range tt.txsToBuildPayload {
				signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, testBankKey)
				if err != nil {
					t.Fatalf("Failed to sign tx %v", err)
				}
				signedTxs = append(signedTxs, signedTx)
			}

			for _, tx := range tt.invalidTxs {
				signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, testBankKey)
				if err != nil {
					t.Fatalf("Failed to sign tx %v", err)
				}
				signedInvalidTxs = append(signedInvalidTxs, signedTx)
			}

			// set the astria ordered txsToBuildPayload
			b.TxPool().SetAstriaOrdered(signedTxs)
			astriaTxs := b.TxPool().AstriaOrdered()

			if astriaTxs.Len() != len(tt.txsToBuildPayload) {
				t.Fatalf("Unexpected number of astria ordered transactions: %d", astriaTxs.Len())
			}

			txs := types.TxDifference(*astriaTxs, signedTxs)
			if txs.Len() != 0 {
				t.Fatalf("Unexpected transactions in astria ordered transactions: %v", txs)
			}

			args := &BuildPayloadArgs{
				Parent:       b.chain.CurrentBlock().Hash(),
				Timestamp:    timestamp,
				Random:       common.Hash{},
				FeeRecipient: recipient,
			}

			payload, err := w.buildPayload(args)
			if err != nil {
				t.Fatalf("Failed to build payload %v", err)
			}
			full := payload.ResolveFull()
			verify(full, len(tt.expectedTxsInPayload))

			// Ensure resolve can be called multiple times and the
			// result should be unchanged
			dataOne := payload.Resolve()
			dataTwo := payload.Resolve()
			if !reflect.DeepEqual(dataOne, dataTwo) {
				t.Fatal("Unexpected payload data")
			}

			// Ensure invalid transactions are stored
			if len(tt.invalidTxs) > 0 {
				invalidTxs := b.TxPool().AstriaExcludedFromBlock()
				txDifference := types.TxDifference(*invalidTxs, signedInvalidTxs)
				if txDifference.Len() != 0 {
					t.Fatalf("Unexpected invalid transactions in astria invalid transactions: %v", txDifference)
				}
			}
		})
	}
}

func TestPayloadId(t *testing.T) {
	ids := make(map[string]int)
	for i, tt := range []*BuildPayloadArgs{
		{
			Parent:       common.Hash{1},
			Timestamp:    1,
			Random:       common.Hash{0x1},
			FeeRecipient: common.Address{0x1},
		},
		// Different parent
		{
			Parent:       common.Hash{2},
			Timestamp:    1,
			Random:       common.Hash{0x1},
			FeeRecipient: common.Address{0x1},
		},
		// Different timestamp
		{
			Parent:       common.Hash{2},
			Timestamp:    2,
			Random:       common.Hash{0x1},
			FeeRecipient: common.Address{0x1},
		},
		// Different Random
		{
			Parent:       common.Hash{2},
			Timestamp:    2,
			Random:       common.Hash{0x2},
			FeeRecipient: common.Address{0x1},
		},
		// Different fee-recipient
		{
			Parent:       common.Hash{2},
			Timestamp:    2,
			Random:       common.Hash{0x2},
			FeeRecipient: common.Address{0x2},
		},
		// Different withdrawals (non-empty)
		{
			Parent:       common.Hash{2},
			Timestamp:    2,
			Random:       common.Hash{0x2},
			FeeRecipient: common.Address{0x2},
			Withdrawals: []*types.Withdrawal{
				{
					Index:     0,
					Validator: 0,
					Address:   common.Address{},
					Amount:    0,
				},
			},
		},
		// Different withdrawals (non-empty)
		{
			Parent:       common.Hash{2},
			Timestamp:    2,
			Random:       common.Hash{0x2},
			FeeRecipient: common.Address{0x2},
			Withdrawals: []*types.Withdrawal{
				{
					Index:     2,
					Validator: 0,
					Address:   common.Address{},
					Amount:    0,
				},
			},
		},
	} {
		id := tt.Id().String()
		if prev, exists := ids[id]; exists {
			t.Errorf("ID collision, case %d and case %d: id %v", prev, i, id)
		}
		ids[id] = i
	}
}
