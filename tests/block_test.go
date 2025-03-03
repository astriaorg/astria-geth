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

package tests

//
//func TestBlockchain(t *testing.T) {
//	bt := new(testMatcher)
//
//	// We are running most of GeneralStatetests to tests witness support, even
//	// though they are ran as state tests too. Still, the performance tests are
//	// less about state andmore about EVM number crunching, so skip those.
//	bt.skipLoad(`^GeneralStateTests/VMTests/vmPerformance`)
//
//	// Skip random failures due to selfish mining test
//	bt.skipLoad(`.*bcForgedTest/bcForkUncle\.json`)
//
//	// Slow tests
//	bt.slow(`.*bcExploitTest/DelegateCallSpam.json`)
//	bt.slow(`.*bcExploitTest/ShanghaiLove.json`)
//	bt.slow(`.*bcExploitTest/SuicideIssue.json`)
//	bt.slow(`.*/bcForkStressTest/`)
//	bt.slow(`.*/bcGasPricerTest/RPC_API_Test.json`)
//	bt.slow(`.*/bcWalletTest/`)
//
//	// Very slow test
//	bt.skipLoad(`.*/stTimeConsuming/.*`)
//	// test takes a lot for time and goes easily OOM because of sha3 calculation on a huge range,
//	// using 4.6 TGas
//	bt.skipLoad(`.*randomStatetest94.json.*`)
//
//	// After the merge we would accept side chains as canonical even if they have lower td
//	bt.skipLoad(`.*bcMultiChainTest/ChainAtoChainB_difficultyB.json`)
//	bt.skipLoad(`.*bcMultiChainTest/CallContractFromNotBestBlock.json`)
//	bt.skipLoad(`.*bcTotalDifficultyTest/uncleBlockAtBlock3afterBlock4.json`)
//	bt.skipLoad(`.*bcTotalDifficultyTest/lotsOfBranchesOverrideAtTheMiddle.json`)
//	bt.skipLoad(`.*bcTotalDifficultyTest/sideChainWithMoreTransactions.json`)
//	bt.skipLoad(`.*bcForkStressTest/ForkStressTest.json`)
//	bt.skipLoad(`.*bcMultiChainTest/lotsOfLeafs.json`)
//	bt.skipLoad(`.*bcFrontierToHomestead/blockChainFrontierWithLargerTDvsHomesteadBlockchain.json`)
//	bt.skipLoad(`.*bcFrontierToHomestead/blockChainFrontierWithLargerTDvsHomesteadBlockchain2.json`)
//
//	bt.walk(t, blockTestDir, func(t *testing.T, name string, test *BlockTest) {
//		execBlockTest(t, bt, test)
//	})
//	// There is also a LegacyTests folder, containing blockchain tests generated
//	// prior to Istanbul. However, they are all derived from GeneralStateTests,
//	// which run natively, so there's no reason to run them here.
//}

// TODO - get these tests to work. Tests mainly fail due to https://github.com/astriaorg/astria-geth/pull/5
// where we add the basefee balance to the coinbase address. This causes the state root to change, we will have to
// update the expected state roots in the tests
// TestExecutionSpec runs the test fixtures from execution-spec-tests.
//func TestExecutionSpec(t *testing.T) {
//	if !common.FileExist(executionSpecDir) {
//		t.Skipf("directory %s does not exist", executionSpecDir)
//	}
//	bt := new(testMatcher)
//
//	bt.walk(t, executionSpecDir, func(t *testing.T, name string, test *BlockTest) {
//		execBlockTest(t, bt, test)
//	})
//}

// TestExecutionSpecBlocktests runs the test fixtures from execution-spec-tests.
//func TestExecutionSpecBlocktests(t *testing.T) {
//	if !common.FileExist(executionSpecBlockchainTestDir) {
//		t.Skipf("directory %s does not exist", executionSpecBlockchainTestDir)
//	}
//	bt := new(testMatcher)
//
//	bt.walk(t, executionSpecBlockchainTestDir, func(t *testing.T, name string, test *BlockTest) {
//		execBlockTest(t, bt, test)
//	})
//}
//
//func execBlockTest(t *testing.T, bt *testMatcher, test *BlockTest) {
//	// If -short flag is used, we don't execute all four permutations, only one.
//	executionMask := 0xf
//	if testing.Short() {
//		executionMask = (1 << (rand.Int63() & 4))
//	}
//	if executionMask&0x1 != 0 {
//		if err := bt.checkFailure(t, test.Run(false, rawdb.HashScheme, nil, nil)); err != nil {
//			t.Errorf("test in hash mode without snapshotter failed: %v", err)
//			return
//		}
//	}
//	if executionMask&0x2 != 0 {
//		if err := bt.checkFailure(t, test.Run(true, rawdb.HashScheme, nil, nil)); err != nil {
//			t.Errorf("test in hash mode with snapshotter failed: %v", err)
//			return
//		}
//	}
//	if executionMask&0x4 != 0 {
//		if err := bt.checkFailure(t, test.Run(false, rawdb.PathScheme, nil, nil)); err != nil {
//			t.Errorf("test in path mode without snapshotter failed: %v", err)
//			return
//		}
//	}
//	if executionMask&0x8 != 0 {
//		if err := bt.checkFailure(t, test.Run(true, rawdb.PathScheme, nil, nil)); err != nil {
//			t.Errorf("test in path mode with snapshotter failed: %v", err)
//			return
//		}
//	}
//}
