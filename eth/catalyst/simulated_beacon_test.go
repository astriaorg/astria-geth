// Copyright 2023 The go-ethereum Authors
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
package catalyst

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/miner"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
)

func startSimulatedBeaconEthService(t *testing.T, genesis *core.Genesis, period uint64) (*node.Node, *eth.Ethereum, *SimulatedBeacon) {
	t.Helper()

	n, err := node.New(&node.Config{
		P2P: p2p.Config{
			ListenAddr:  "127.0.0.1:0",
			NoDiscovery: true,
			MaxPeers:    0,
		},
	})
	if err != nil {
		t.Fatal("can't create node:", err)
	}

	ethcfg := &ethconfig.Config{Genesis: genesis, SyncMode: downloader.FullSync, TrieTimeout: time.Minute, TrieDirtyCache: 256, TrieCleanCache: 256, Miner: miner.DefaultConfig}
	ethservice, err := eth.New(n, ethcfg)
	if err != nil {
		t.Fatal("can't create eth service:", err)
	}

	simBeacon, err := NewSimulatedBeacon(period, ethservice)
	if err != nil {
		t.Fatal("can't create simulated beacon:", err)
	}

	n.RegisterLifecycle(simBeacon)

	if err := n.Start(); err != nil {
		t.Fatal("can't start node:", err)
	}

	ethservice.SetSynced()
	return n, ethservice, simBeacon
}

// NOTE(bharath): The `TestSimulatedBeaconSendWithdrawals` are not relevant for astria since we do not use the ethereum
// beacon chain for consensus.
// send 20 transactions, >10 withdrawals and ensure they are included in order
// send enough transactions to fill multiple blocks
//func TestSimulatedBeaconSendWithdrawals(t *testing.T) {
//	var withdrawals []types.Withdrawal
//	txs := make(map[common.Hash]*types.Transaction)
//
//	var (
//		// testKey is a private key to use for funding a tester account.
//		testKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
//
//		// testAddr is the Ethereum address of the tester account.
//		testAddr = crypto.PubkeyToAddress(testKey.PublicKey)
//	)
//
//	// short period (1 second) for testing purposes
//	var gasLimit uint64 = 10_000_000
//	genesis := core.DeveloperGenesisBlock(gasLimit, &testAddr)
//	node, ethService, mock := startSimulatedBeaconEthService(t, genesis, 1)
//	_ = mock
//	defer node.Close()
//
//	chainHeadCh := make(chan core.ChainHeadEvent, 10)
//	subscription := ethService.BlockChain().SubscribeChainHeadEvent(chainHeadCh)
//	defer subscription.Unsubscribe()
//
//	// generate some withdrawals
//	for i := 0; i < 20; i++ {
//		withdrawals = append(withdrawals, types.Withdrawal{Index: uint64(i)})
//		if err := mock.withdrawals.add(&withdrawals[i]); err != nil {
//			t.Fatal("addWithdrawal failed", err)
//		}
//	}
//
//	// generate a bunch of transactions
//	signer := types.NewEIP155Signer(ethService.BlockChain().Config().ChainID)
//	for i := 0; i < 20; i++ {
//		tx, err := types.SignTx(types.NewTransaction(uint64(i), common.Address{}, big.NewInt(1000), params.TxGas, big.NewInt(params.InitialBaseFee), nil), signer, testKey)
//		if err != nil {
//			t.Fatalf("error signing transaction, err=%v", err)
//		}
//		txs[tx.Hash()] = tx
//
//		//if err := ethService.APIBackend.SendTx(context.Background(), tx); err != nil {
//		//	t.Fatal("SendTx failed", err)
//		//}
//	}
//
//	txsAdded := types.Transactions{}
//	for _, v := range txs {
//		txsAdded = append(txsAdded, v)
//	}
//
//	ethService.TxPool().SetAstriaOrdered(txsAdded)
//
//	includedTxs := make(map[common.Hash]struct{})
//	var includedWithdrawals []uint64
//
//	timer := time.NewTimer(12 * time.Second)
//	for {
//		select {
//		case ev := <-chainHeadCh:
//			block := ethService.BlockChain().GetBlock(ev.Header.Hash(), ev.Header.Number.Uint64())
//			for _, includedTx := range block.Transactions() {
//				includedTxs[includedTx.Hash()] = struct{}{}
//			}
//			for _, includedWithdrawal := range block.Withdrawals() {
//				includedWithdrawals = append(includedWithdrawals, includedWithdrawal.Index)
//			}
//			// ensure all withdrawals/txs included. this will take two blocks b/c number of withdrawals > 10
//			if len(includedTxs) == len(txs) && len(includedWithdrawals) == len(withdrawals) && ev.Header.Number.Cmp(big.NewInt(2)) == 0 {
//				return
//			}
//		case <-timer.C:
//			t.Fatal("timed out without including all withdrawals/txs")
//		}
//	}
//}

// TODO(bharath) - These tests fail because we do not drive execution through the beacon chain but through conductor.
// Tests that zero-period dev mode can handle a lot of simultaneous
// transactions/withdrawals
//func TestOnDemandSpam(t *testing.T) {
//	var (
//		withdrawals     []types.Withdrawal
//		txs                    = make(map[common.Hash]*types.Transaction)
//		testKey, _             = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
//		testAddr               = crypto.PubkeyToAddress(testKey.PublicKey)
//		gasLimit        uint64 = 10_000_000
//		genesis                = core.DeveloperGenesisBlock(gasLimit, &testAddr)
//		node, eth, mock        = startSimulatedBeaconEthService(t, genesis, 0)
//		_                      = newSimulatedBeaconAPI(mock)
//		signer                 = types.LatestSigner(eth.BlockChain().Config())
//		chainHeadCh            = make(chan core.ChainHeadEvent, 100)
//		sub                    = eth.BlockChain().SubscribeChainHeadEvent(chainHeadCh)
//	)
//	defer node.Close()
//	defer sub.Unsubscribe()
//
//	// generate some withdrawals
//	for i := 0; i < 20; i++ {
//		withdrawals = append(withdrawals, types.Withdrawal{Index: uint64(i)})
//		if err := mock.withdrawals.add(&withdrawals[i]); err != nil {
//			t.Fatal("addWithdrawal failed", err)
//		}
//	}
//
//	// generate a bunch of transactions
//	allTxs := types.Transactions{}
//	for i := 0; i < 20000; i++ {
//		tx, err := types.SignTx(types.NewTransaction(uint64(i), common.Address{byte(i), byte(1)}, big.NewInt(1000), params.TxGas, big.NewInt(params.InitialBaseFee*2), nil), signer, testKey)
//		if err != nil {
//			t.Fatal("error signing transaction", err)
//		}
//		txs[tx.Hash()] = tx
//		allTxs = append(allTxs, tx)
//		//if err := eth.APIBackend.SendTx(context.Background(), tx); err != nil {
//		//	t.Fatal("error adding txs to pool", err)
//		//}
//	}
//
//	eth.TxPool().SetAstriaOrdered(allTxs)
//
//	var (
//		includedTxs = make(map[common.Hash]struct{})
//		includedWxs []uint64
//	)
//	for {
//		select {
//		case ev := <-chainHeadCh:
//			block := eth.BlockChain().GetBlock(ev.Header.Hash(), ev.Header.Number.Uint64())
//			for _, itx := range block.Transactions() {
//				includedTxs[itx.Hash()] = struct{}{}
//			}
//			for _, iwx := range block.Withdrawals() {
//				includedWxs = append(includedWxs, iwx.Index)
//			}
//			// ensure all withdrawals/txs included. this will take two blocks b/c number of withdrawals > 10
//			if len(includedTxs) == len(txs) && len(includedWxs) == len(withdrawals) {
//				return
//			}
//		case <-time.After(10 * time.Second):
//			t.Fatalf("timed out without including all withdrawals/txs: have txs %d, want %d, have wxs %d, want %d", len(includedTxs), len(txs), len(includedWxs), len(withdrawals))
//		}
//	}
//}
