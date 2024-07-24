// Copyright 2018 The go-ethereum Authors
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
	"bytes"
	"crypto/sha256"
	"math/big"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/clique"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/txpool/legacypool"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
)

const (
	// testCode is the testing contract binary code which will initialises some
	// variables in constructor
	testCode = "0x60806040527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060005534801561003457600080fd5b5060fc806100436000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c80630c4dae8814603757806398a213cf146053575b600080fd5b603d607e565b6040518082815260200191505060405180910390f35b607c60048036036020811015606757600080fd5b81019080803590602001909291905050506084565b005b60005481565b806000819055507fe9e44f9f7da8c559de847a3232b57364adc0354f15a2cd8dc636d54396f9587a6000546040518082815260200191505060405180910390a15056fea265627a7a723058208ae31d9424f2d0bc2a3da1a5dd659db2d71ec322a17db8f87e19e209e3a1ff4a64736f6c634300050a0032"

	// testGas is the gas required for contract deployment.
	testGas = 144109
)

var (
	// Test chain configurations
	testTxPoolConfig  legacypool.Config
	ethashChainConfig *params.ChainConfig
	cliqueChainConfig *params.ChainConfig

	// Test accounts
	testBankKey, _  = crypto.GenerateKey()
	testBankAddress = crypto.PubkeyToAddress(testBankKey.PublicKey)
	testBankFunds   = big.NewInt(1000000000000000000)

	testUserKey, _  = crypto.GenerateKey()
	testUserAddress = crypto.PubkeyToAddress(testUserKey.PublicKey)

	// Test transactions
	pendingTxs []*types.Transaction
	newTxs     []*types.Transaction

	testConfig = &Config{
		Recommit: time.Second,
		GasCeil:  params.GenesisGasLimit,
	}
)

func init() {
	testTxPoolConfig = legacypool.DefaultConfig
	testTxPoolConfig.Journal = ""
	ethashChainConfig = new(params.ChainConfig)
	*ethashChainConfig = *params.TestChainConfig
	ethashChainConfig.CancunBlock = nil
	cliqueChainConfig = new(params.ChainConfig)
	*cliqueChainConfig = *params.TestChainConfig
	cliqueChainConfig.CancunBlock = nil
	cliqueChainConfig.Clique = &params.CliqueConfig{
		Period: 10,
		Epoch:  30000,
	}

	signer := types.LatestSigner(params.TestChainConfig)
	tx1 := types.MustSignNewTx(testBankKey, signer, &types.AccessListTx{
		ChainID:  params.TestChainConfig.ChainID,
		Nonce:    0,
		To:       &testUserAddress,
		Value:    big.NewInt(1000),
		Gas:      params.TxGas,
		GasPrice: big.NewInt(params.InitialBaseFee),
	})
	pendingTxs = append(pendingTxs, tx1)

	tx2 := types.MustSignNewTx(testBankKey, signer, &types.LegacyTx{
		Nonce:    1,
		To:       &testUserAddress,
		Value:    big.NewInt(1000),
		Gas:      params.TxGas,
		GasPrice: big.NewInt(params.InitialBaseFee),
	})
	newTxs = append(newTxs, tx2)

	rand.Seed(time.Now().UnixNano())
}

// testWorkerBackend implements worker.Backend interfaces and wraps all information needed during the testing.
type testWorkerBackend struct {
	db         ethdb.Database
	txPool     *txpool.TxPool
	chain      *core.BlockChain
	testTxFeed event.Feed
	genesis    *core.Genesis
	uncleBlock *types.Block
}

func newTestWorkerBackend(t *testing.T, chainConfig *params.ChainConfig, engine consensus.Engine, db ethdb.Database, n int) *testWorkerBackend {
	var gspec = core.Genesis{
		Config: chainConfig,
		Alloc:  core.GenesisAlloc{testBankAddress: {Balance: testBankFunds}},
	}

	switch e := engine.(type) {
	case *clique.Clique:
		gspec.ExtraData = make([]byte, 32+common.AddressLength+crypto.SignatureLength)
		copy(gspec.ExtraData[32:32+common.AddressLength], testBankAddress.Bytes())
		e.Authorize(testBankAddress, func(account accounts.Account, s string, data []byte) ([]byte, error) {
			return crypto.Sign(crypto.Keccak256(data), testBankKey)
		})
	case *ethash.Ethash:
	default:
		t.Fatalf("unexpected consensus engine type: %T", engine)
	}
	genesis := gspec.MustCommit(db)

	chain, _ := core.NewBlockChain(db, &core.CacheConfig{TrieDirtyDisabled: true}, gspec.Config, engine, vm.Config{}, nil, nil)
	legacyPool := legacypool.New(testTxPoolConfig, chainConfig, chain)
	txpool, err := txpool.New(testTxPoolConfig.PriceLimit, chain, []txpool.SubPool{legacyPool})
	if err != nil {
		t.Fatal(err)
	}

	// Generate a small n-block chain and an uncle block for it
	if n > 0 {
		blocks, _ := core.GenerateChain(chainConfig, genesis, engine, db, n, func(i int, gen *core.BlockGen) {
			gen.SetCoinbase(testBankAddress)
		}, true)
		if _, err := chain.InsertChain(blocks); err != nil {
			t.Fatalf("failed to insert origin chain: %v", err)
		}
	}
	parent := genesis
	if n > 0 {
		parent = chain.GetBlockByHash(chain.CurrentBlock().ParentHash())
	}
	blocks, _ := core.GenerateChain(chainConfig, parent, engine, db, 1, func(i int, gen *core.BlockGen) {
		gen.SetCoinbase(testUserAddress)
	}, true)

	return &testWorkerBackend{
		db:         db,
		chain:      chain,
		txPool:     txpool,
		genesis:    &gspec,
		uncleBlock: blocks[0],
	}
}

func (b *testWorkerBackend) BlockChain() *core.BlockChain { return b.chain }
func (b *testWorkerBackend) TxPool() *txpool.TxPool       { return b.txPool }

func (b *testWorkerBackend) newRandomUncle() *types.Block {
	var parent *types.Block
	cur := b.chain.CurrentBlock()
	if cur.NumberU64() == 0 {
		parent = b.chain.Genesis()
	} else {
		parent = b.chain.GetBlockByHash(b.chain.CurrentBlock().ParentHash())
	}
	blocks, _ := core.GenerateChain(b.chain.Config(), parent, b.chain.Engine(), b.db, 1, func(i int, gen *core.BlockGen) {
		var addr = make([]byte, common.AddressLength)
		rand.Read(addr)
		gen.SetCoinbase(common.BytesToAddress(addr))
	}, true)
	return blocks[0]
}

func (b *testWorkerBackend) newRandomTx(creation bool) *types.Transaction {
	var tx *types.Transaction
	gasPrice := big.NewInt(10 * params.InitialBaseFee)
	if creation {
		tx, _ = types.SignTx(types.NewContractCreation(b.txPool.Nonce(testBankAddress), big.NewInt(0), testGas, gasPrice, common.FromHex(testCode)), types.HomesteadSigner{}, testBankKey)
	} else {
		tx, _ = types.SignTx(types.NewTransaction(b.txPool.Nonce(testBankAddress), testUserAddress, big.NewInt(1000), params.TxGas, gasPrice, nil), types.HomesteadSigner{}, testBankKey)
	}
	return tx
}

func newTestWorker(t *testing.T, chainConfig *params.ChainConfig, engine consensus.Engine, db ethdb.Database, blocks int) (*worker, *testWorkerBackend) {
	backend := newTestWorkerBackend(t, chainConfig, engine, db, blocks)
	backend.txPool.Add(pendingTxs, true, false)
	w := newWorker(testConfig, chainConfig, engine, backend, new(event.TypeMux), nil, false)
	w.setEtherbase(testBankAddress)
	return w, backend
}

func TestGenerateBlockAndImportEthash(t *testing.T) {
	testGenerateBlockAndImport(t, false)
}

func TestGenerateBlockAndImportClique(t *testing.T) {
	testGenerateBlockAndImport(t, true)
}

func testGenerateBlockAndImport(t *testing.T, isClique bool) {
	var (
		engine      consensus.Engine
		chainConfig *params.ChainConfig
		db          = rawdb.NewMemoryDatabase()
	)
	if isClique {
		chainConfig = params.AllCliqueProtocolChanges
		chainConfig.Clique = &params.CliqueConfig{Period: 1, Epoch: 30000}
		engine = clique.New(chainConfig.Clique, db)
	} else {
		chainConfig = params.AllEthashProtocolChanges
		engine = ethash.NewFaker()
	}

	chainConfig.LondonBlock = big.NewInt(0)
	w, b := newTestWorker(t, chainConfig, engine, db, 0)
	defer w.close()

	// This test chain imports the mined blocks.
	db2 := rawdb.NewMemoryDatabase()
	b.genesis.MustCommit(db2)
	chain, _ := core.NewBlockChain(db2, nil, b.chain.Config(), engine, vm.Config{}, nil, nil)
	defer chain.Stop()

	// Ignore empty commit here for less noise.
	w.skipSealHook = func(task *task) bool {
		return len(task.receipts) == 0
	}

	// Wait for mined blocks.
	sub := w.mux.Subscribe(core.NewMinedBlockEvent{})
	defer sub.Unsubscribe()

	// Start mining!
	w.start()

	for i := 0; i < 5; i++ {
		b.txPool.Add([]*types.Transaction{b.newRandomTx(true)}, true, false)
		b.txPool.Add([]*types.Transaction{b.newRandomTx(false)}, true, false)
		w.postSideBlock(core.ChainSideEvent{Block: b.newRandomUncle()})
		w.postSideBlock(core.ChainSideEvent{Block: b.newRandomUncle()})

		select {
		case ev := <-sub.Chan():
			block := ev.Data.(core.NewMinedBlockEvent).Block
			if _, err := chain.InsertChain([]*types.Block{block}); err != nil {
				t.Fatalf("failed to insert new mined block %d: %v", block.NumberU64(), err)
			}
		case <-time.After(3 * time.Second): // Worker needs 1s to include new changes.
			t.Fatalf("timeout")
		}
	}
}

func TestEmptyWorkEthash(t *testing.T) {
	testEmptyWork(t, ethashChainConfig, ethash.NewFaker())
}
func TestEmptyWorkClique(t *testing.T) {
	testEmptyWork(t, cliqueChainConfig, clique.New(cliqueChainConfig.Clique, rawdb.NewMemoryDatabase()))
}

func testEmptyWork(t *testing.T, chainConfig *params.ChainConfig, engine consensus.Engine) {
	defer engine.Close()

	w, _ := newTestWorker(t, chainConfig, engine, rawdb.NewMemoryDatabase(), 0)
	defer w.close()

	var (
		taskIndex int
		taskCh    = make(chan struct{}, 2)
	)
	checkEqual := func(t *testing.T, task *task, index int) {
		// The first empty work without any txs included
		receiptLen, balance := 0, big.NewInt(0)
		if index == 1 {
			// The second full work with 1 tx included
			receiptLen, balance = 1, big.NewInt(1000)
		}
		if len(task.receipts) != receiptLen {
			t.Fatalf("receipt number mismatch: have %d, want %d", len(task.receipts), receiptLen)
		}
		if task.state.GetBalance(testUserAddress).Cmp(balance) != 0 {
			t.Fatalf("account balance mismatch: have %d, want %d", task.state.GetBalance(testUserAddress), balance)
		}
	}
	w.newTaskHook = func(task *task) {
		if task.block.NumberU64() == 1 {
			checkEqual(t, task, taskIndex)
			taskIndex += 1
			taskCh <- struct{}{}
		}
	}
	w.skipSealHook = func(task *task) bool { return true }
	w.fullTaskHook = func() {
		time.Sleep(100 * time.Millisecond)
	}
	w.start() // Start mining!
	for i := 0; i < 2; i += 1 {
		select {
		case <-taskCh:
		case <-time.NewTimer(3 * time.Second).C:
			t.Error("new task timeout")
		}
	}
}

func TestStreamUncleBlock(t *testing.T) {
	ethash := ethash.NewFaker()
	defer ethash.Close()

	w, b := newTestWorker(t, ethashChainConfig, ethash, rawdb.NewMemoryDatabase(), 1)
	defer w.close()

	var taskCh = make(chan struct{})

	taskIndex := 0
	w.newTaskHook = func(task *task) {
		if task.block.NumberU64() == 2 {
			// The first task is an empty task, the second
			// one has 1 pending tx, the third one has 1 tx
			// and 1 uncle.
			if taskIndex == 2 {
				have := task.block.Header().UncleHash
				want := types.CalcUncleHash([]*types.Header{b.uncleBlock.Header()})
				if have != want {
					t.Errorf("uncle hash mismatch: have %s, want %s", have.Hex(), want.Hex())
				}
			}
			taskCh <- struct{}{}
			taskIndex += 1
		}
	}
	w.skipSealHook = func(task *task) bool {
		return true
	}
	w.fullTaskHook = func() {
		time.Sleep(100 * time.Millisecond)
	}
	w.start()

	for i := 0; i < 2; i += 1 {
		select {
		case <-taskCh:
		case <-time.NewTimer(time.Second).C:
			t.Error("new task timeout")
		}
	}

	w.postSideBlock(core.ChainSideEvent{Block: b.uncleBlock})

	select {
	case <-taskCh:
	case <-time.NewTimer(time.Second).C:
		t.Error("new task timeout")
	}
}

func TestRegenerateMiningBlockEthash(t *testing.T) {
	testRegenerateMiningBlock(t, ethashChainConfig, ethash.NewFaker())
}

func TestRegenerateMiningBlockClique(t *testing.T) {
	testRegenerateMiningBlock(t, cliqueChainConfig, clique.New(cliqueChainConfig.Clique, rawdb.NewMemoryDatabase()))
}

func testRegenerateMiningBlock(t *testing.T, chainConfig *params.ChainConfig, engine consensus.Engine) {
	defer engine.Close()

	w, b := newTestWorker(t, chainConfig, engine, rawdb.NewMemoryDatabase(), 0)
	defer w.close()

	var taskCh = make(chan struct{})

	taskIndex := 0
	w.newTaskHook = func(task *task) {
		if task.block.NumberU64() == 1 {
			// The first task is an empty task, the second
			// one has 1 pending tx, the third one has 2 txs
			if taskIndex == 2 {
				receiptLen, balance := 2, big.NewInt(2000)
				if len(task.receipts) != receiptLen {
					t.Errorf("receipt number mismatch: have %d, want %d", len(task.receipts), receiptLen)
				}
				if task.state.GetBalance(testUserAddress).Cmp(balance) != 0 {
					t.Errorf("account balance mismatch: have %d, want %d", task.state.GetBalance(testUserAddress), balance)
				}
			}
			taskCh <- struct{}{}
			taskIndex += 1
		}
	}
	w.skipSealHook = func(task *task) bool {
		return true
	}
	w.fullTaskHook = func() {
		time.Sleep(100 * time.Millisecond)
	}

	w.start()
	// Ignore the first two works
	for i := 0; i < 2; i += 1 {
		select {
		case <-taskCh:
		case <-time.NewTimer(time.Second).C:
			t.Error("new task timeout")
		}
	}
	b.txPool.Add(newTxs, true, false)
	time.Sleep(time.Second)

	select {
	case <-taskCh:
	case <-time.NewTimer(time.Second).C:
		t.Error("new task timeout")
	}
}

func TestAdjustIntervalEthash(t *testing.T) {
	testAdjustInterval(t, ethashChainConfig, ethash.NewFaker())
}

func TestAdjustIntervalClique(t *testing.T) {
	testAdjustInterval(t, cliqueChainConfig, clique.New(cliqueChainConfig.Clique, rawdb.NewMemoryDatabase()))
}

func testAdjustInterval(t *testing.T, chainConfig *params.ChainConfig, engine consensus.Engine) {
	defer engine.Close()

	w, _ := newTestWorker(t, chainConfig, engine, rawdb.NewMemoryDatabase(), 0)
	defer w.close()

	w.skipSealHook = func(task *task) bool {
		return true
	}
	w.fullTaskHook = func() {
		time.Sleep(100 * time.Millisecond)
	}
	var (
		progress = make(chan struct{}, 10)
		result   = make([]float64, 0, 10)
		index    = 0
		start    uint32
	)
	w.resubmitHook = func(minInterval time.Duration, recommitInterval time.Duration) {
		// Short circuit if interval checking hasn't started.
		if atomic.LoadUint32(&start) == 0 {
			return
		}
		var wantMinInterval, wantRecommitInterval time.Duration

		switch index {
		case 0:
			wantMinInterval, wantRecommitInterval = 3*time.Second, 3*time.Second
		case 1:
			origin := float64(3 * time.Second.Nanoseconds())
			estimate := origin*(1-intervalAdjustRatio) + intervalAdjustRatio*(origin/0.8+intervalAdjustBias)
			wantMinInterval, wantRecommitInterval = 3*time.Second, time.Duration(estimate)*time.Nanosecond
		case 2:
			estimate := result[index-1]
			min := float64(3 * time.Second.Nanoseconds())
			estimate = estimate*(1-intervalAdjustRatio) + intervalAdjustRatio*(min-intervalAdjustBias)
			wantMinInterval, wantRecommitInterval = 3*time.Second, time.Duration(estimate)*time.Nanosecond
		case 3:
			wantMinInterval, wantRecommitInterval = time.Second, time.Second
		}

		// Check interval
		if minInterval != wantMinInterval {
			t.Errorf("resubmit min interval mismatch: have %v, want %v ", minInterval, wantMinInterval)
		}
		if recommitInterval != wantRecommitInterval {
			t.Errorf("resubmit interval mismatch: have %v, want %v", recommitInterval, wantRecommitInterval)
		}
		result = append(result, float64(recommitInterval.Nanoseconds()))
		index += 1
		progress <- struct{}{}
	}
	w.start()

	time.Sleep(time.Second) // Ensure two tasks have been summitted due to start opt
	atomic.StoreUint32(&start, 1)

	w.setRecommitInterval(3 * time.Second)
	select {
	case <-progress:
	case <-time.NewTimer(time.Second).C:
		t.Error("interval reset timeout")
	}

	w.resubmitAdjustCh <- &intervalAdjust{inc: true, ratio: 0.8}
	select {
	case <-progress:
	case <-time.NewTimer(time.Second).C:
		t.Error("interval reset timeout")
	}

	w.resubmitAdjustCh <- &intervalAdjust{inc: false}
	select {
	case <-progress:
	case <-time.NewTimer(time.Second).C:
		t.Error("interval reset timeout")
	}

	w.setRecommitInterval(500 * time.Millisecond)
	select {
	case <-progress:
	case <-time.NewTimer(time.Second).C:
		t.Error("interval reset timeout")
	}
}

func TestDoubleSignPrevention(t *testing.T) {
	chainConfig := cliqueChainConfig
	/* The double sign prevention logic only applies in consortium v2 */
	chainConfig.ConsortiumV2Block = common.Big0
	chainConfig.Clique.Period = 1

	w, b := newTestWorker(t, chainConfig, clique.New(cliqueChainConfig.Clique, rawdb.NewMemoryDatabase()), rawdb.NewMemoryDatabase(), 0)
	defer w.close()

	chainEvent := make(chan core.ChainEvent)
	chainEventSub := b.chain.SubscribeChainEvent(chainEvent)
	defer chainEventSub.Unsubscribe()

	chainSideEvent := make(chan core.ChainSideEvent)
	chainSideEventSub := b.chain.SubscribeChainSideEvent(chainSideEvent)
	defer chainSideEventSub.Unsubscribe()

	w.start()

	done := make(chan struct{})
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		for {
			select {
			case <-ticker.C:
				// Spam transactions to trigger recommit
				b.txPool.Add([]*types.Transaction{b.newRandomTx(false)}, true, false)
			case <-done:
				return
			}
		}
	}()

	timer := time.NewTimer(10 * time.Second)
	generatedBlocks := make(map[uint64][]*types.Block)
	for {
		select {
		case ev := <-chainEvent:
			newBlock := ev.Block
			if blocks, ok := generatedBlocks[newBlock.NumberU64()]; ok {
				for _, block := range blocks {
					if bytes.Equal(block.ParentHash().Bytes(), newBlock.ParentHash().Bytes()) {
						t.Errorf("Block %d is double signed", newBlock.NumberU64())
						return
					}
				}
				blocks = append(blocks, newBlock)
				generatedBlocks[newBlock.NumberU64()] = blocks
			} else {
				generatedBlocks[newBlock.NumberU64()] = []*types.Block{newBlock}
			}
		case ev := <-chainSideEvent:
			newBlock := ev.Block
			if blocks, ok := generatedBlocks[newBlock.NumberU64()]; ok {
				for _, block := range blocks {
					if bytes.Equal(block.ParentHash().Bytes(), newBlock.ParentHash().Bytes()) {
						t.Errorf("Block %d is double signed", newBlock.NumberU64())
						return
					}
				}
				blocks = append(blocks, newBlock)
				generatedBlocks[newBlock.NumberU64()] = blocks
			} else {
				generatedBlocks[newBlock.NumberU64()] = []*types.Block{newBlock}
			}
		case <-timer.C:
			done <- struct{}{}
			return
		}
	}
}

func toLazyTransaction(tx *types.Transaction) *txpool.LazyTransaction {
	return &txpool.LazyTransaction{
		Tx:        tx,
		Time:      time.Now(),
		GasFeeCap: uint256.MustFromBig(tx.GasFeeCap()),
		GasTipCap: uint256.MustFromBig(tx.GasTipCap()),
		Gas:       tx.Gas(),
		BlobGas:   tx.BlobGas(),
	}
}

func newCurrent(t *testing.T, signer types.Signer, chain *core.BlockChain) *environment {
	state, err := chain.StateAt(common.Hash{})
	if err != nil {
		t.Fatal(err)
	}

	excessBlobGas := uint64(0)
	header := &types.Header{
		GasLimit:      100_000_000,
		Number:        common.Big0,
		Difficulty:    big.NewInt(7),
		BaseFee:       common.Big0,
		ExcessBlobGas: &excessBlobGas,
	}
	env := &environment{
		signer:             signer,
		state:              state,
		ancestors:          mapset.NewSet(),
		family:             mapset.NewSet(),
		uncles:             mapset.NewSet(),
		header:             header,
		estimatedBlockSize: 0,
		tcount:             0,
	}
	return env
}

func TestCommitBlobTransaction(t *testing.T) {
	var (
		emptyBlob          = new(kzg4844.Blob)
		emptyBlobCommit, _ = kzg4844.BlobToCommitment(emptyBlob)
		emptyBlobProof, _  = kzg4844.ComputeBlobProof(emptyBlob, emptyBlobCommit)
		emptyBlobVHash     = kzg4844.CalcBlobHashV1(sha256.New(), &emptyBlobCommit)
	)

	db := rawdb.NewMemoryDatabase()
	chainConfig := params.AllEthashProtocolChanges
	chainConfig.ChainID = big.NewInt(2020)
	chainConfig.RoninTreasuryAddress = &common.Address{0x88}
	engine := clique.New(cliqueChainConfig.Clique, db)

	backend := newTestWorkerBackend(t, chainConfig, engine, db, 0)
	w := newWorker(testConfig, chainConfig, engine, backend, new(event.TypeMux), nil, false)

	signer := types.NewCancunSigner(big.NewInt(2020))
	senderKey1, _ := crypto.GenerateKey()
	senderAddress1 := crypto.PubkeyToAddress(senderKey1.PublicKey)
	senderKey2, _ := crypto.GenerateKey()
	senderAddress2 := crypto.PubkeyToAddress(senderKey2.PublicKey)

	legacyTx, err := types.SignNewTx(senderKey1, signer, &types.LegacyTx{
		Nonce:    0,
		GasPrice: big.NewInt(20_000_000_000),
		Gas:      21000,
		To:       &senderAddress1,
	})
	if err != nil {
		t.Fatal(err)
	}

	legacyTxs := make(map[common.Address][]*txpool.LazyTransaction)
	legacyTxs[senderAddress1] = append(legacyTxs[senderAddress1], toLazyTransaction(legacyTx))
	plainTxsByPrice := NewTransactionsByPriceAndNonce(signer, legacyTxs, common.Big0)

	blobTx, err := types.SignNewTx(senderKey2, signer, &types.BlobTx{
		ChainID:    uint256.NewInt(2020),
		Nonce:      0,
		GasTipCap:  uint256.NewInt(20_000_000_000),
		GasFeeCap:  uint256.NewInt(20_000_000_000),
		Gas:        21000,
		To:         senderAddress2,
		BlobHashes: []common.Hash{emptyBlobVHash},
		BlobFeeCap: uint256.MustFromBig(common.Big1),
		Sidecar: &types.BlobTxSidecar{
			Blobs:       []kzg4844.Blob{*emptyBlob},
			Commitments: []kzg4844.Commitment{emptyBlobCommit},
			Proofs:      []kzg4844.Proof{emptyBlobProof},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	blobTxs := make(map[common.Address][]*txpool.LazyTransaction)
	blobTxs[senderAddress2] = append(blobTxs[senderAddress2], toLazyTransaction(blobTx))
	blobTxsByPrice := NewTransactionsByPriceAndNonce(signer, blobTxs, common.Big0)

	w.current = newCurrent(t, signer, w.chain)
	w.current.state.AddBalance(senderAddress1, new(big.Int).Exp(big.NewInt(10), big.NewInt(20), nil))
	w.current.state.AddBalance(senderAddress2, new(big.Int).Exp(big.NewInt(10), big.NewInt(20), nil))

	// Case 1: Not Cancun, blob transactions are not committed but plain transactions are committed
	// normally without error
	failed := w.commitTransactions(plainTxsByPrice, blobTxsByPrice, senderAddress1, nil)
	if failed {
		t.Fatal("Commit transaction failed")
	}
	if len(w.current.txs) != 1 || w.current.header.GasUsed != 21000 {
		t.Fatalf(
			"Unexpected mined block, number of txs %d, gas used %d",
			len(w.current.txs), w.current.header.GasUsed,
		)
	}

	// Case 2: Higher blob transaction tip is prioritized
	w.current = newCurrent(t, signer, w.chain)
	w.current.state.AddBalance(senderAddress1, new(big.Int).Exp(big.NewInt(10), big.NewInt(20), nil))
	w.current.state.AddBalance(senderAddress2, new(big.Int).Exp(big.NewInt(10), big.NewInt(20), nil))

	w.current.header.BlobGasUsed = new(uint64)
	chainConfig.CancunBlock = common.Big0
	w.chainConfig = chainConfig

	plainTxsByPrice = NewTransactionsByPriceAndNonce(signer, make(map[common.Address][]*txpool.LazyTransaction), common.Big0)

	blobTx1, err := types.SignNewTx(senderKey1, signer, &types.BlobTx{
		ChainID:    uint256.NewInt(2020),
		Nonce:      0,
		GasTipCap:  uint256.NewInt(30_000_000_000),
		GasFeeCap:  uint256.NewInt(30_000_000_000),
		Gas:        21000,
		To:         senderAddress1,
		BlobHashes: []common.Hash{emptyBlobVHash, emptyBlobVHash, emptyBlobVHash, emptyBlobVHash, emptyBlobVHash, emptyBlobVHash},
		BlobFeeCap: uint256.MustFromBig(common.Big1),
		Sidecar: &types.BlobTxSidecar{
			Blobs:       []kzg4844.Blob{*emptyBlob, *emptyBlob, *emptyBlob, *emptyBlob, *emptyBlob, *emptyBlob},
			Commitments: []kzg4844.Commitment{emptyBlobCommit, emptyBlobCommit, emptyBlobCommit, emptyBlobCommit, emptyBlobCommit, emptyBlobCommit},
			Proofs:      []kzg4844.Proof{emptyBlobProof, emptyBlobProof, emptyBlobProof, emptyBlobProof, emptyBlobProof, emptyBlobProof},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	blobTx2, err := types.SignNewTx(senderKey2, signer, &types.BlobTx{
		ChainID:    uint256.NewInt(2020),
		Nonce:      0,
		GasTipCap:  uint256.NewInt(20_000_000_000),
		GasFeeCap:  uint256.NewInt(20_000_000_000),
		Gas:        21000,
		To:         senderAddress2,
		BlobHashes: []common.Hash{emptyBlobVHash},
		BlobFeeCap: uint256.MustFromBig(common.Big1),
		Sidecar: &types.BlobTxSidecar{
			Blobs:       []kzg4844.Blob{*emptyBlob},
			Commitments: []kzg4844.Commitment{emptyBlobCommit},
			Proofs:      []kzg4844.Proof{emptyBlobProof},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	blobTxs = make(map[common.Address][]*txpool.LazyTransaction)
	blobTxs[senderAddress1] = append(blobTxs[senderAddress1], toLazyTransaction(blobTx1))
	blobTxs[senderAddress2] = append(blobTxs[senderAddress2], toLazyTransaction(blobTx2))
	blobTxsByPrice = NewTransactionsByPriceAndNonce(signer, blobTxs, common.Big0)
	failed = w.commitTransactions(plainTxsByPrice, blobTxsByPrice, senderAddress1, nil)
	if failed {
		t.Fatal("Commit transaction failed")
	}
	if len(w.current.txs) != 1 || w.current.header.GasUsed != 21000 || *w.current.header.BlobGasUsed != 6*params.BlobTxBlobGasPerBlob {
		t.Fatalf(
			"Unexpected mined block, number of txs %d, gas used %d, blob gas used %d",
			len(w.current.txs), w.current.header.GasUsed, *w.current.header.BlobGasUsed,
		)
	}

	// Case 3: Choose the blob transaction that does not make the blob gas used exceed the limit
	w.current = newCurrent(t, signer, w.chain)
	w.current.state.AddBalance(senderAddress1, new(big.Int).Exp(big.NewInt(10), big.NewInt(20), nil))
	w.current.state.AddBalance(senderAddress2, new(big.Int).Exp(big.NewInt(10), big.NewInt(20), nil))

	w.current.header.BlobGasUsed = new(uint64)
	*w.current.header.BlobGasUsed = 3 * params.BlobTxBlobGasPerBlob
	chainConfig.CancunBlock = common.Big0
	w.chainConfig = chainConfig

	blobTxs = make(map[common.Address][]*txpool.LazyTransaction)
	blobTxs[senderAddress1] = append(blobTxs[senderAddress1], toLazyTransaction(blobTx1))
	blobTxs[senderAddress2] = append(blobTxs[senderAddress2], toLazyTransaction(blobTx2))
	blobTxsByPrice = NewTransactionsByPriceAndNonce(signer, blobTxs, common.Big0)
	failed = w.commitTransactions(plainTxsByPrice, blobTxsByPrice, senderAddress1, nil)
	if failed {
		t.Fatal("Commit transaction failed")
	}
	if len(w.current.txs) != 1 || w.current.header.GasUsed != 21000 || *w.current.header.BlobGasUsed != 4*params.BlobTxBlobGasPerBlob {
		t.Fatalf(
			"Unexpected mined block, number of txs %d, gas used %d, blob gas used %d",
			len(w.current.txs), w.current.header.GasUsed, *w.current.header.BlobGasUsed,
		)
	}
}
