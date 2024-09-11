// Copyright 2021 The go-ethereum Authors
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

package tracers

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"sort"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/tracers/logger"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	errStateNotFound = errors.New("state not found")
	errBlockNotFound = errors.New("block not found")
)

type testBackend struct {
	chainConfig *params.ChainConfig
	engine      consensus.Engine
	chaindb     ethdb.Database
	chain       *core.BlockChain

	refHook func() // Hook is invoked when the requested state is referenced
	relHook func() // Hook is invoked when the requested state is released
}

// testBackend creates a new test backend. OBS: After test is done, teardown must be
// invoked in order to release associated resources.
func newTestBackend(t *testing.T, n int, gspec *core.Genesis, generator func(i int, b *core.BlockGen)) *testBackend {
	backend := &testBackend{
		chainConfig: gspec.Config,
		engine:      ethash.NewFaker(),
		chaindb:     rawdb.NewMemoryDatabase(),
	}
	var blocks []*types.Block
	// Generate blocks for testing
	_, blocks, _ = core.GenerateChainWithGenesis(gspec, backend.engine, n, generator)

	// Import the canonical chain
	cacheConfig := &core.CacheConfig{
		TrieCleanLimit:    256,
		TrieDirtyLimit:    256,
		TrieTimeLimit:     5 * time.Minute,
		SnapshotLimit:     0,
		TrieDirtyDisabled: true, // Archive mode
	}

	_, _, genesisErr := core.SetupGenesisBlockWithOverride(backend.chaindb, gspec, nil, true)
	if _, ok := genesisErr.(*params.ConfigCompatError); genesisErr != nil && !ok {
		t.Fatal(genesisErr.Error())
	}

	//chainDb, cacheConfig, chainConfig, eth.engine, vmConfig, eth.shouldPreserve, &config.TxLookupLimit
	chain, err := core.NewBlockChain(backend.chaindb, cacheConfig, gspec, nil, backend.engine, vm.Config{}, nil, nil)
	if err != nil {
		t.Fatalf("failed to create tester chain: %v", err)
	}
	if n, err := chain.InsertChain(blocks); err != nil {
		t.Fatalf("block %d: failed to insert into chain: %v", n, err)
	}
	backend.chain = chain
	return backend
}

func (b *testBackend) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	return b.chain.GetHeaderByHash(hash), nil
}

func (b *testBackend) HeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Header, error) {
	if number == rpc.PendingBlockNumber || number == rpc.LatestBlockNumber {
		return b.chain.CurrentHeader(), nil
	}
	return b.chain.GetHeaderByNumber(uint64(number)), nil
}

func (b *testBackend) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	return b.chain.GetBlockByHash(hash), nil
}

func (b *testBackend) BlockByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Block, error) {
	if number == rpc.PendingBlockNumber || number == rpc.LatestBlockNumber {
		return b.chain.GetBlockByNumber(b.chain.CurrentBlock().Number().Uint64()), nil
	}
	return b.chain.GetBlockByNumber(uint64(number)), nil
}

func (b *testBackend) GetTransaction(ctx context.Context, txHash common.Hash) (*types.Transaction, common.Hash, uint64, uint64, error) {
	tx, hash, blockNumber, index := rawdb.ReadTransaction(b.chaindb, txHash)
	return tx, hash, blockNumber, index, nil
}

func (b *testBackend) RPCGasCap() uint64 {
	return 25000000
}

func (b *testBackend) ChainConfig() *params.ChainConfig {
	return b.chainConfig
}

func (b *testBackend) Engine() consensus.Engine {
	return b.engine
}

func (b *testBackend) ChainDb() ethdb.Database {
	return b.chaindb
}

// teardown releases the associated resources.
func (b *testBackend) teardown() {
	b.chain.Stop()
}

func (b *testBackend) StateAtBlock(ctx context.Context, block *types.Block, reexec uint64, base *state.StateDB, readOnly bool, preferDisk bool) (*state.StateDB, StateReleaseFunc, error) {
	statedb, err := b.chain.StateAt(block.Root())
	if err != nil {
		return nil, nil, errStateNotFound
	}
	if b.refHook != nil {
		b.refHook()
	}
	release := func() {
		if b.relHook != nil {
			b.relHook()
		}
	}
	return statedb, release, nil
}

func (b *testBackend) StateAtTransaction(ctx context.Context, block *types.Block, txIndex int, reexec uint64) (core.Message, vm.BlockContext, *state.StateDB, StateReleaseFunc, error) {
	parent := b.chain.GetBlock(block.ParentHash(), block.NumberU64()-1)
	if parent == nil {
		return nil, vm.BlockContext{}, nil, nil, errBlockNotFound
	}
	statedb, release, err := b.StateAtBlock(ctx, parent, reexec, nil, true, false)
	if err != nil {
		return nil, vm.BlockContext{}, nil, nil, errStateNotFound
	}
	if txIndex == 0 && len(block.Transactions()) == 0 {
		return nil, vm.BlockContext{}, statedb, release, nil
	}
	// Recompute transactions up to the target index.
	signer := types.MakeSigner(b.chainConfig, block.Number())
	for idx, tx := range block.Transactions() {
		msg, _ := tx.AsMessage(signer, block.BaseFee())
		txContext := core.NewEVMTxContext(msg)
		context := core.NewEVMBlockContext(block.Header(), b.chain, nil)
		if idx == txIndex {
			return msg, context, statedb, release, nil
		}
		vmenv := vm.NewEVM(context, txContext, statedb, b.chainConfig, vm.Config{})
		if _, err := core.ApplyMessage(vmenv, msg, new(core.GasPool).AddGas(tx.Gas())); err != nil {
			return nil, vm.BlockContext{}, nil, nil, fmt.Errorf("transaction %#x failed: %v", tx.Hash(), err)
		}
		statedb.Finalise(vmenv.ChainConfig().IsEIP158(block.Number()))
	}
	return nil, vm.BlockContext{}, nil, nil, fmt.Errorf("transaction index %d out of range for block %#x", txIndex, block.Hash())
}

func TestTraceCall(t *testing.T) {
	t.Parallel()

	// Initialize test accounts
	accounts := newAccounts(3)
	genesis := &core.Genesis{
		Config: params.TestChainConfig,
		Alloc: core.GenesisAlloc{
			accounts[0].addr: {Balance: big.NewInt(params.Ether)},
			accounts[1].addr: {Balance: big.NewInt(params.Ether)},
			accounts[2].addr: {Balance: big.NewInt(params.Ether)},
		},
	}
	genBlocks := 10
	signer := types.HomesteadSigner{}
	backend := newTestBackend(t, genBlocks, genesis, func(i int, b *core.BlockGen) {
		// Transfer from account[0] to account[1]
		//    value: 1000 wei
		//    fee:   0 wei
		tx, _ := types.SignTx(types.NewTransaction(uint64(i), accounts[1].addr, big.NewInt(1000), params.TxGas, b.BaseFee(), nil), signer, accounts[0].key)
		b.AddTx(tx)
	})
	defer backend.teardown()
	api := NewAPI(backend)
	var testSuite = []struct {
		blockNumber rpc.BlockNumber
		call        ethapi.TransactionArgs
		config      *TraceCallConfig
		expectErr   error
		expect      string
	}{
		// Standard JSON trace upon the genesis, plain transfer.
		{
			blockNumber: rpc.BlockNumber(0),
			call: ethapi.TransactionArgs{
				From:  &accounts[0].addr,
				To:    &accounts[1].addr,
				Value: (*hexutil.Big)(big.NewInt(1000)),
			},
			config:    nil,
			expectErr: nil,
			expect:    `{"gas":21000,"failed":false,"returnValue":"","structLogs":[]}`,
		},
		// Standard JSON trace upon the head, plain transfer.
		{
			blockNumber: rpc.BlockNumber(genBlocks),
			call: ethapi.TransactionArgs{
				From:  &accounts[0].addr,
				To:    &accounts[1].addr,
				Value: (*hexutil.Big)(big.NewInt(1000)),
			},
			config:    nil,
			expectErr: nil,
			expect:    `{"gas":21000,"failed":false,"returnValue":"","structLogs":[]}`,
		},
		// Standard JSON trace upon the non-existent block, error expects
		{
			blockNumber: rpc.BlockNumber(genBlocks + 1),
			call: ethapi.TransactionArgs{
				From:  &accounts[0].addr,
				To:    &accounts[1].addr,
				Value: (*hexutil.Big)(big.NewInt(1000)),
			},
			config:    nil,
			expectErr: fmt.Errorf("block #%d not found", genBlocks+1),
			//expect:    nil,
		},
		// Standard JSON trace upon the latest block
		{
			blockNumber: rpc.LatestBlockNumber,
			call: ethapi.TransactionArgs{
				From:  &accounts[0].addr,
				To:    &accounts[1].addr,
				Value: (*hexutil.Big)(big.NewInt(1000)),
			},
			config:    nil,
			expectErr: nil,
			expect:    `{"gas":21000,"failed":false,"returnValue":"","structLogs":[]}`,
		},
		// Tracing on 'pending' should fail:
		{
			blockNumber: rpc.PendingBlockNumber,
			call: ethapi.TransactionArgs{
				From:  &accounts[0].addr,
				To:    &accounts[1].addr,
				Value: (*hexutil.Big)(big.NewInt(1000)),
			},
			config:    nil,
			expectErr: errors.New("tracing on top of pending is not supported"),
		},
		{
			blockNumber: rpc.LatestBlockNumber,
			call: ethapi.TransactionArgs{
				From:  &accounts[0].addr,
				Input: &hexutil.Bytes{0x43}, // blocknumber
			},
			config: &TraceCallConfig{
				BlockOverrides: &ethapi.BlockOverrides{Number: (*hexutil.Big)(big.NewInt(0x1337))},
			},
			expectErr: nil,
			expect: ` {"gas":53018,"failed":false,"returnValue":"","structLogs":[
		{"pc":0,"op":"NUMBER","gas":24946984,"gasCost":2,"depth":1,"stack":[]},
		{"pc":1,"op":"STOP","gas":24946982,"gasCost":0,"depth":1,"stack":["0x1337"]}]}`,
		},
	}
	for i, testspec := range testSuite {
		result, err := api.TraceCall(context.Background(), testspec.call, rpc.BlockNumberOrHash{BlockNumber: &testspec.blockNumber}, testspec.config)
		if testspec.expectErr != nil {
			if err == nil {
				t.Errorf("test %d: expect error %v, got nothing", i, testspec.expectErr)
				continue
			}
			if !reflect.DeepEqual(err, testspec.expectErr) {
				t.Errorf("test %d: error mismatch, want %v, git %v", i, testspec.expectErr, err)
			}
		} else {
			if err != nil {
				t.Errorf("test %d: expect no error, got %v", i, err)
				continue
			}
			var have *logger.ExecutionResult
			if err := json.Unmarshal(result.(json.RawMessage), &have); err != nil {
				t.Errorf("test %d: failed to unmarshal result %v", i, err)
			}
			var want *logger.ExecutionResult
			if err := json.Unmarshal([]byte(testspec.expect), &want); err != nil {
				t.Errorf("test %d: failed to unmarshal result %v", i, err)
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("test %d: result mismatch, want %v, got %v", i, testspec.expect, string(result.(json.RawMessage)))
			}
		}
	}
}

func TestTraceTransaction(t *testing.T) {
	t.Parallel()

	// Initialize test accounts
	accounts := newAccounts(2)
	genesis := &core.Genesis{
		Config: params.TestChainConfig,
		Alloc: core.GenesisAlloc{
			accounts[0].addr: {Balance: big.NewInt(params.Ether)},
			accounts[1].addr: {Balance: big.NewInt(params.Ether)},
		},
	}
	target := common.Hash{}
	signer := types.HomesteadSigner{}
	backend := newTestBackend(t, 1, genesis, func(i int, b *core.BlockGen) {
		// Transfer from account[0] to account[1]
		//    value: 1000 wei
		//    fee:   0 wei
		tx, _ := types.SignTx(types.NewTransaction(uint64(i), accounts[1].addr, big.NewInt(1000), params.TxGas, b.BaseFee(), nil), signer, accounts[0].key)
		b.AddTx(tx)
		target = tx.Hash()
	})
	defer backend.chain.Stop()
	api := NewAPI(backend)
	result, err := api.TraceTransaction(context.Background(), target, nil)
	if err != nil {
		t.Errorf("Failed to trace transaction %v", err)
	}
	var have *logger.ExecutionResult
	if err := json.Unmarshal(result.(json.RawMessage), &have); err != nil {
		t.Errorf("failed to unmarshal result %v", err)
	}
	if !reflect.DeepEqual(have, &logger.ExecutionResult{
		Gas:         params.TxGas,
		Failed:      false,
		ReturnValue: "",
		StructLogs:  []logger.StructLogRes{},
	}) {
		t.Error("Transaction tracing result is different")
	}

	// Test non-existent transaction
	_, err = api.TraceTransaction(context.Background(), common.Hash{42}, nil)
	if !errors.Is(err, errTxNotFound) {
		t.Fatalf("want %v, have %v", errTxNotFound, err)
	}
}

func TestTraceBlock(t *testing.T) {
	t.Parallel()

	// Initialize test accounts
	accounts := newAccounts(3)
	genesis := &core.Genesis{
		Config: params.TestChainConfig,
		Alloc: core.GenesisAlloc{
			accounts[0].addr: {Balance: big.NewInt(params.Ether)},
			accounts[1].addr: {Balance: big.NewInt(params.Ether)},
			accounts[2].addr: {Balance: big.NewInt(params.Ether)},
		},
	}
	genBlocks := 10
	signer := types.HomesteadSigner{}
	backend := newTestBackend(t, genBlocks, genesis, func(i int, b *core.BlockGen) {
		// Transfer from account[0] to account[1]
		//    value: 1000 wei
		//    fee:   0 wei
		tx, _ := types.SignTx(types.NewTransaction(uint64(i), accounts[1].addr, big.NewInt(1000), params.TxGas, b.BaseFee(), nil), signer, accounts[0].key)
		b.AddTx(tx)
	})
	defer backend.chain.Stop()
	api := NewAPI(backend)

	var testSuite = []struct {
		blockNumber rpc.BlockNumber
		config      *TraceConfig
		want        string
		expectErr   error
	}{
		// Trace genesis block, expect error
		{
			blockNumber: rpc.BlockNumber(0),
			expectErr:   errors.New("genesis is not traceable"),
		},
		// Trace head block
		{
			blockNumber: rpc.BlockNumber(genBlocks),
			want:        `[{"result":{"gas":21000,"failed":false,"returnValue":"","structLogs":[]}}]`,
		},
		// Trace non-existent block
		{
			blockNumber: rpc.BlockNumber(genBlocks + 1),
			expectErr:   fmt.Errorf("block #%d not found", genBlocks+1),
		},
		// Trace latest block
		{
			blockNumber: rpc.LatestBlockNumber,
			want:        `[{"result":{"gas":21000,"failed":false,"returnValue":"","structLogs":[]}}]`,
		},
		// Trace pending block
		{
			blockNumber: rpc.PendingBlockNumber,
			want:        `[{"result":{"gas":21000,"failed":false,"returnValue":"","structLogs":[]}}]`,
		},
	}
	for i, tc := range testSuite {
		result, err := api.TraceBlockByNumber(context.Background(), tc.blockNumber, tc.config)
		if tc.expectErr != nil {
			if err == nil {
				t.Errorf("test %d, want error %v", i, tc.expectErr)
				continue
			}
			if !reflect.DeepEqual(err, tc.expectErr) {
				t.Errorf("test %d: error mismatch, want %v, get %v", i, tc.expectErr, err)
			}
			continue
		}
		if err != nil {
			t.Errorf("test %d, want no error, have %v", i, err)
			continue
		}
		have, _ := json.Marshal(result)
		want := tc.want
		if strings.Contains(string(have), want) {
			t.Errorf("test %d, result mismatch, have\n%v\n, want\n%v\n", i, string(have), want)
		}
	}
}

func TestTraceInternalsAndAccounts_BatchTransferAccounts(t *testing.T) {
	accounts := newAccounts(2)
	genesis := &core.Genesis{Alloc: core.GenesisAlloc{
		accounts[0].addr: {
			// Contract code
			// // SPDX-License-Identifier: MIT
			//  pragma solidity ^0.8.14;
			//
			//  contract MultiTransfer {
			//    event Transfer(address indexed from, address indexed to, uint256 value);
			//
			//    constructor() {}
			//
			//    function transferFunds(address payable[] memory recipients, uint256 amount)
			//        public
			//        payable
			//    {
			//        require(
			//            amount * recipients.length <= address(this).balance,
			//            "Insufficient funds in the contract"
			//        );
			//
			//        for (uint256 i = 0; i < recipients.length; i++) {
			//            recipients[i].transfer(amount);
			//            emit Transfer(msg.sender, recipients[i], amount);
			//        }
			//    }
			//
			//    receive() external payable {}
			//
			//    fallback() external payable {}
			//  }
			Code:    common.FromHex("0x6080604052600436106100225760003560e01c8063876e58611461002b57610029565b3661002957005b005b6100456004803603810190610040919061039b565b610047565b005b478251826100559190610426565b1115610096576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161008d906104eb565b60405180910390fd5b60005b8251811015610195578281815181106100b5576100b461050b565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f19350505050158015610102573d6000803e3d6000fd5b508281815181106101165761011561050b565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161017a9190610549565b60405180910390a3808061018d90610564565b915050610099565b505050565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6101fc826101b3565b810181811067ffffffffffffffff8211171561021b5761021a6101c4565b5b80604052505050565b600061022e61019a565b905061023a82826101f3565b919050565b600067ffffffffffffffff82111561025a576102596101c4565b5b602082029050602081019050919050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061029b82610270565b9050919050565b6102ab81610290565b81146102b657600080fd5b50565b6000813590506102c8816102a2565b92915050565b60006102e16102dc8461023f565b610224565b905080838252602082019050602084028301858111156103045761030361026b565b5b835b8181101561032d578061031988826102b9565b845260208401935050602081019050610306565b5050509392505050565b600082601f83011261034c5761034b6101ae565b5b813561035c8482602086016102ce565b91505092915050565b6000819050919050565b61037881610365565b811461038357600080fd5b50565b6000813590506103958161036f565b92915050565b600080604083850312156103b2576103b16101a4565b5b600083013567ffffffffffffffff8111156103d0576103cf6101a9565b5b6103dc85828601610337565b92505060206103ed85828601610386565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061043182610365565b915061043c83610365565b925082820261044a81610365565b91508282048414831517610461576104606103f7565b5b5092915050565b600082825260208201905092915050565b7f496e73756666696369656e742066756e647320696e2074686520636f6e74726160008201527f6374000000000000000000000000000000000000000000000000000000000000602082015250565b60006104d5602283610468565b91506104e082610479565b604082019050919050565b60006020820190508181036000830152610504816104c8565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b61054381610365565b82525050565b600060208201905061055e600083018461053a565b92915050565b600061056f82610365565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036105a1576105a06103f7565b5b60018201905091905056fea2646970667358221220cf97fc5287a954e32220197048875abb6e6e7b645fd1f06e12b25c456d86218d64736f6c63430008110033"),
			Balance: big.NewInt(params.Ether),
		},
		accounts[1].addr: {Balance: big.NewInt(params.Ether)},
	}, Config: params.TestChainConfig}
	genBlocks := 3
	signer := types.HomesteadSigner{}
	api := NewAPI(newTestBackend(t, genBlocks, genesis, func(i int, b *core.BlockGen) {
		if i == 1 {
			// Batch send to this account:
			// - "0x05ba56c60ceb54f53294bf60d606a919eea4282e"
			data1 := common.FromHex("0x876e586100000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000100000000000000000000000005ba56c60ceb54f53294bf60d606a919eea4282e")
			tx1, _ := types.SignTx(types.NewTransaction(uint64(0), accounts[0].addr, big.NewInt(0), 4*params.TxGas, big.NewInt(0), data1), signer, accounts[1].key)
			b.AddTx(tx1)

			// Batch send to these accounts:
			// - "0x359aef78ffa9807889258d0dd398172ca3b77eb1"
			// - "0x45f60b415111e3e7abb4c79fc659d3dee430dff5"
			data2 := common.FromHex("0x876e5861000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002000000000000000000000000359aef78ffa9807889258d0dd398172ca3b77eb100000000000000000000000045f60b415111e3e7abb4c79fc659d3dee430dff5")
			tx2, _ := types.SignTx(types.NewTransaction(uint64(1), accounts[0].addr, big.NewInt(0), 10*params.TxGas, big.NewInt(0), data2), signer, accounts[1].key)
			b.AddTx(tx2)
		}
	}))

	ctx := context.Background()
	block, err := api.blockByNumber(ctx, rpc.BlockNumber(2))
	if err != nil {
		panic(err)
	}

	result, err := api.TraceInternalsAndAccountsByBlockHash(context.Background(), block.Hash(), nil)
	if err != nil {
		panic(err)
	}

	callOps := make([]string, 0)
	for _, itx := range result.InternalTxs {
		// unmarshal result
		var rs ethapi.ExecutionResult
		if err = json.Unmarshal(itx.Result.(json.RawMessage), &rs); err != nil {
			panic(err)
		}
		if rs.Failed {
			b := hexutil.MustDecode("0x" + rs.ReturnValue)
			t.Fatal(string(b))
		}

		structLogs := rs.StructLogs
		for _, s := range structLogs {
			if s.Op == "CALL" {
				callOps = append(callOps, s.Op)
			}
			if s.Error != "" {
				t.Fatalf("unable to process op: %s - %v", s.Op, s.Error)
			}
		}
	}

	txsLen := block.Transactions().Len()
	expectedTxsLen := 2
	if txsLen != expectedTxsLen {
		t.Errorf("got %v, wanted %v", txsLen, expectedTxsLen)
	}

	callOpsLen := len(callOps)
	expectedCallOpsLen := 3
	if callOpsLen != expectedCallOpsLen {
		t.Errorf("got %v, wanted %v", callOpsLen, expectedCallOpsLen)
	}

	expectedDirtyAccounts := map[common.Address]*big.Int{}
	expectedDirtyAccounts[common.HexToAddress("0x05bA56C60ceb54f53294bF60d606a919eeA4282E")] = big.NewInt(1)
	expectedDirtyAccounts[common.HexToAddress("0x359aEf78fFa9807889258D0DD398172CA3B77eB1")] = big.NewInt(1)
	expectedDirtyAccounts[common.HexToAddress("0x45F60B415111e3E7aBB4c79FC659d3DeE430dfF5")] = big.NewInt(1)
	expectedDirtyAccounts[accounts[1].addr] = big.NewInt(1000000000000000000)
	expectedDirtyAccounts[accounts[0].addr] = big.NewInt(999999999999999997)

	for _, actual := range result.DirtyAccounts {
		expectedBalance, exists := expectedDirtyAccounts[actual.Address]
		if !exists {
			t.Errorf("account %v not found", actual.Address)
		}

		if actual.Balance.ToInt().Cmp(expectedBalance) != 0 {
			t.Errorf("account balance is not match got %v, wanted %v", actual.Balance, expectedBalance)
		}
	}
}

func TestTraceInternalsAndAccounts_CreateContract(t *testing.T) {
	accounts := newAccounts(2)
	genesis := &core.Genesis{Alloc: core.GenesisAlloc{
		accounts[0].addr: {
			// Contract code
			// // SPDX-License-Identifier: MIT
			//  pragma solidity ^0.8.14;
			//
			//  contract Factory {
			//    function deploy() public {
			//        new TestContract();
			//    }
			//  }
			//
			//  contract TestContract {}
			Code:    common.FromHex("0x6080604052348015600f57600080fd5b506004361060285760003560e01c8063775c300c14602d575b600080fd5b60336035565b005b604051603f90605e565b604051809103906000f080158015605a573d6000803e3d6000fd5b5050565b605c8061006b8339019056fe6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea2646970667358221220f97f5d6b859eee3b1716ad536a9f65496d508465d92a1b7f3bcf53843bd6499e64736f6c63430008110033a26469706673582212206e55fec0de5b157c817dc80a9135e13a7b49421106e446a3359ff3cee8ea4b6664736f6c63430008110033"),
			Balance: big.NewInt(params.Ether),
		},
		accounts[1].addr: {Balance: big.NewInt(params.Ether)},
	}, Config: params.TestChainConfig}

	genBlocks := 3
	signer := types.HomesteadSigner{}
	backend := newTestBackend(t, genBlocks, genesis, func(i int, b *core.BlockGen) {
		if i == 1 {
			// Call `deploy` method to create a new contract
			data1 := common.FromHex("0x775c300c")
			tx1, _ := types.SignTx(types.NewTransaction(uint64(0), accounts[0].addr, big.NewInt(0), 100*params.TxGas, big.NewInt(0), data1), signer, accounts[1].key)
			b.AddTx(tx1)

			// Call `deploy` method to create a new contract
			data2 := common.FromHex("0x775c300c")
			tx2, _ := types.SignTx(types.NewTransaction(uint64(1), accounts[0].addr, big.NewInt(0), 100*params.TxGas, big.NewInt(0), data2), signer, accounts[1].key)
			b.AddTx(tx2)
		}
	})
	api := NewAPI(backend)

	ctx := context.Background()
	block, err := api.blockByNumber(ctx, rpc.BlockNumber(2))
	if err != nil {
		panic(err)
	}

	result, err := api.TraceInternalsAndAccountsByBlockHash(context.Background(), block.Hash(), nil)
	if err != nil {
		panic(err)
	}

	callOps := make([]string, 0)
	for _, itx := range result.InternalTxs {
		var rs ethapi.ExecutionResult
		if err = json.Unmarshal(itx.Result.(json.RawMessage), &rs); err != nil {
			panic(err)
		}
		if rs.Failed {
			b := hexutil.MustDecode("0x" + rs.ReturnValue)
			t.Fatal(string(b))
		}

		structLogs := rs.StructLogs
		for _, s := range structLogs {
			if s.Op == "CREATE" {
				callOps = append(callOps, s.Op)
			}
			if s.Error != "" {
				t.Fatalf("unable to process op: %s - %v", s.Op, s.Error)
			}
		}
	}

	receiptsLen := len(block.Transactions())
	expectedTxsLen := 2
	if receiptsLen != expectedTxsLen {
		t.Errorf("got %v, wanted %v", receiptsLen, expectedTxsLen)
	}

	callOpsLen := len(callOps)
	expectedCallOpsLen := 2
	if callOpsLen != expectedCallOpsLen {
		t.Errorf("got %v, wanted %v", callOpsLen, expectedCallOpsLen)
	}

	// Only check the dirty accounts len since can not get logs in the test
	dirtyAccountsLen := len(result.DirtyAccounts)
	expectedDirtyAccountsLen := 4
	if dirtyAccountsLen != expectedDirtyAccountsLen {
		t.Errorf("got %v, wanted %v", dirtyAccountsLen, expectedDirtyAccountsLen)
	}

	expectedDirtyAccounts := map[common.Address]*big.Int{}
	expectedDirtyAccounts[accounts[1].addr] = big.NewInt(1000000000000000000)
	expectedDirtyAccounts[accounts[0].addr] = big.NewInt(1000000000000000000)

	for _, actual := range result.DirtyAccounts {
		expectedBalance, exists := expectedDirtyAccounts[actual.Address]
		if !exists {
			continue
		}

		if actual.Balance.ToInt().Cmp(expectedBalance) != 0 {
			t.Errorf("account balance is not match got %v, wanted %v", actual.Balance, expectedBalance)
		}
	}
}

func TestTraceInternalsAndAccounts_Create2Contract(t *testing.T) {
	accounts := newAccounts(2)
	genesis := &core.Genesis{Alloc: core.GenesisAlloc{
		accounts[0].addr: {
			// Contract code
			// // SPDX-License-Identifier: MIT
			//  pragma solidity ^0.8.14;
			//
			//  contract Factory {
			//    function deploy(bytes32 _salt) public payable returns (address) {
			//        return address(new TestContract{salt: _salt}());
			//    }
			//  }
			//
			//  contract TestContract {}
			Code:    common.FromHex("0x60806040526004361061001e5760003560e01c80632b85ba3814610023575b600080fd5b61003d600480360381019061003891906100d1565b610053565b60405161004a919061013f565b60405180910390f35b6000816040516100629061008a565b8190604051809103906000f5905080158015610082573d6000803e3d6000fd5b509050919050565b605c8061015b83390190565b600080fd5b6000819050919050565b6100ae8161009b565b81146100b957600080fd5b50565b6000813590506100cb816100a5565b92915050565b6000602082840312156100e7576100e6610096565b5b60006100f5848285016100bc565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610129826100fe565b9050919050565b6101398161011e565b82525050565b60006020820190506101546000830184610130565b9291505056fe6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea2646970667358221220a26d261e0c14f9ad05b19e2aa40e3a052d7c46ddddd59800c2623c7aa01d4f7f64736f6c63430008110033a26469706673582212205471489b98d070ef7f3ffb9bf428f79b32bb7e3673ce270709cda2f1ccac8b3864736f6c63430008110033"),
			Balance: big.NewInt(params.Ether),
		},
		accounts[1].addr: {Balance: big.NewInt(params.Ether)},
	}, Config: params.TestChainConfig}

	genBlocks := 3
	signer := types.HomesteadSigner{}
	backend := newTestBackend(t, genBlocks, genesis, func(i int, b *core.BlockGen) {
		if i == 1 {
			// Call `deploy` method to create a new contract
			data1 := common.FromHex("0x2b85ba3800000000000000000000000000000000000000000000000000000000686f6c61")
			tx1, _ := types.SignTx(types.NewTransaction(uint64(0), accounts[0].addr, big.NewInt(0), 100*params.TxGas, big.NewInt(0), data1), signer, accounts[1].key)
			b.AddTx(tx1)

			// Call `deploy` method to create a new contract
			data2 := common.FromHex("0x2b85ba3800000000000000000000000000000000000000000000000000000000686f6c62")
			tx2, _ := types.SignTx(types.NewTransaction(uint64(1), accounts[0].addr, big.NewInt(0), 100*params.TxGas, big.NewInt(0), data2), signer, accounts[1].key)
			b.AddTx(tx2)
		}
	})
	api := NewAPI(backend)

	ctx := context.Background()
	block, err := api.blockByNumber(ctx, rpc.BlockNumber(2))
	if err != nil {
		panic(err)
	}

	result, err := api.TraceInternalsAndAccountsByBlockHash(context.Background(), block.Hash(), nil)
	if err != nil {
		panic(err)
	}

	callOps := make([]string, 0)
	for _, itx := range result.InternalTxs {
		var res ethapi.ExecutionResult
		if err = json.Unmarshal(itx.Result.(json.RawMessage), &res); err != nil {
			panic(err)
		}
		if res.Failed {
			b := hexutil.MustDecode("0x" + res.ReturnValue)
			t.Fatal(string(b))
		}

		structLogs := res.StructLogs
		for _, s := range structLogs {
			if s.Op == "CREATE2" {
				callOps = append(callOps, s.Op)
			}
			if s.Error != "" {
				t.Fatalf("unable to process op: %s - %v", s.Op, s.Error)
			}
		}
	}

	receiptsLen := len(block.Transactions())
	expectedTxsLen := 2
	if receiptsLen != expectedTxsLen {
		t.Errorf("got %v, wanted %v", receiptsLen, expectedTxsLen)
	}

	callOpsLen := len(callOps)
	expectedCallOpsLen := 2
	if callOpsLen != expectedCallOpsLen {
		t.Errorf("got %v, wanted %v", callOpsLen, expectedCallOpsLen)
	}

	// Only check the dirty accounts len since can not get logs in the test
	dirtyAccountsLen := len(result.DirtyAccounts)
	expectedDirtyAccountsLen := 4
	if dirtyAccountsLen != expectedDirtyAccountsLen {
		t.Errorf("got %v, wanted %v", dirtyAccountsLen, expectedDirtyAccountsLen)
	}

	expectedDirtyAccounts := map[common.Address]*big.Int{}
	expectedDirtyAccounts[accounts[1].addr] = big.NewInt(1000000000000000000)
	expectedDirtyAccounts[accounts[0].addr] = big.NewInt(1000000000000000000)

	for _, actual := range result.DirtyAccounts {
		expectedBalance, exists := expectedDirtyAccounts[actual.Address]
		if !exists {
			continue
		}

		if actual.Balance.ToInt().Cmp(expectedBalance) != 0 {
			t.Errorf("account balance is not match got %v, wanted %v", actual.Balance, expectedBalance)
		}
	}
}

func TestTracingWithOverrides(t *testing.T) {
	t.Parallel()
	// Initialize test accounts
	accounts := newAccounts(3)
	storageAccount := common.Address{0x13, 37}
	genesis := &core.Genesis{
		Config: params.TestChainConfig,
		Alloc: core.GenesisAlloc{
			accounts[0].addr: {Balance: big.NewInt(params.Ether)},
			accounts[1].addr: {Balance: big.NewInt(params.Ether)},
			accounts[2].addr: {Balance: big.NewInt(params.Ether)},
			// An account with existing storage
			storageAccount: {
				Balance: new(big.Int),
				Storage: map[common.Hash]common.Hash{
					common.HexToHash("0x03"): common.HexToHash("0x33"),
					common.HexToHash("0x04"): common.HexToHash("0x44"),
				},
			},
		},
	}
	genBlocks := 10
	signer := types.HomesteadSigner{}
	backend := newTestBackend(t, genBlocks, genesis, func(i int, b *core.BlockGen) {
		// Transfer from account[0] to account[1]
		//    value: 1000 wei
		//    fee:   0 wei
		tx, _ := types.SignTx(types.NewTransaction(uint64(i), accounts[1].addr, big.NewInt(1000), params.TxGas, b.BaseFee(), nil), signer, accounts[0].key)
		b.AddTx(tx)
	})
	defer backend.chain.Stop()
	api := NewAPI(backend)
	randomAccounts := newAccounts(3)
	type res struct {
		Gas         int
		Failed      bool
		ReturnValue string
	}
	var testSuite = []struct {
		blockNumber rpc.BlockNumber
		call        ethapi.TransactionArgs
		config      *TraceCallConfig
		expectErr   error
		want        string
	}{
		// Call which can only succeed if state is state overridden
		{
			blockNumber: rpc.LatestBlockNumber,
			call: ethapi.TransactionArgs{
				From:  &randomAccounts[0].addr,
				To:    &randomAccounts[1].addr,
				Value: (*hexutil.Big)(big.NewInt(1000)),
			},
			config: &TraceCallConfig{
				StateOverrides: &ethapi.StateOverride{
					randomAccounts[0].addr: ethapi.OverrideAccount{Balance: newRPCBalance(new(big.Int).Mul(big.NewInt(1), big.NewInt(params.Ether)))},
				},
			},
			want: `{"gas":21000,"failed":false,"returnValue":""}`,
		},
		// Invalid call without state overriding
		{
			blockNumber: rpc.LatestBlockNumber,
			call: ethapi.TransactionArgs{
				From:  &randomAccounts[0].addr,
				To:    &randomAccounts[1].addr,
				Value: (*hexutil.Big)(big.NewInt(1000)),
			},
			config:    &TraceCallConfig{},
			expectErr: core.ErrInsufficientFunds,
		},
		// Successful simple contract call
		//
		// // SPDX-License-Identifier: GPL-3.0
		//
		//  pragma solidity >=0.7.0 <0.8.0;
		//
		//  /**
		//   * @title Storage
		//   * @dev Store & retrieve value in a variable
		//   */
		//  contract Storage {
		//      uint256 public number;
		//      constructor() {
		//          number = block.number;
		//      }
		//  }
		{
			blockNumber: rpc.LatestBlockNumber,
			call: ethapi.TransactionArgs{
				From: &randomAccounts[0].addr,
				To:   &randomAccounts[2].addr,
				Data: newRPCBytes(common.Hex2Bytes("8381f58a")), // call number()
			},
			config: &TraceCallConfig{
				//Tracer: &tracer,
				StateOverrides: &ethapi.StateOverride{
					randomAccounts[2].addr: ethapi.OverrideAccount{
						Code:      newRPCBytes(common.Hex2Bytes("6080604052348015600f57600080fd5b506004361060285760003560e01c80638381f58a14602d575b600080fd5b60336049565b6040518082815260200191505060405180910390f35b6000548156fea2646970667358221220eab35ffa6ab2adfe380772a48b8ba78e82a1b820a18fcb6f59aa4efb20a5f60064736f6c63430007040033")),
						StateDiff: newStates([]common.Hash{{}}, []common.Hash{common.BigToHash(big.NewInt(123))}),
					},
				},
			},
			want: `{"gas":23347,"failed":false,"returnValue":"000000000000000000000000000000000000000000000000000000000000007b"}`,
		},
		{ // Override blocknumber
			blockNumber: rpc.LatestBlockNumber,
			call: ethapi.TransactionArgs{
				From: &accounts[0].addr,
				// BLOCKNUMBER PUSH1 MSTORE
				Input: newRPCBytes(common.Hex2Bytes("4360005260206000f3")),
				//&hexutil.Bytes{0x43}, // blocknumber
			},
			config: &TraceCallConfig{
				BlockOverrides: &ethapi.BlockOverrides{Number: (*hexutil.Big)(big.NewInt(0x1337))},
			},
			want: `{"gas":59537,"failed":false,"returnValue":"0000000000000000000000000000000000000000000000000000000000001337"}`,
		},
		{ // Override blocknumber, and query a blockhash
			blockNumber: rpc.LatestBlockNumber,
			call: ethapi.TransactionArgs{
				From: &accounts[0].addr,
				Input: &hexutil.Bytes{
					0x60, 0x00, 0x40, // BLOCKHASH(0)
					0x60, 0x00, 0x52, // STORE memory offset 0
					0x61, 0x13, 0x36, 0x40, // BLOCKHASH(0x1336)
					0x60, 0x20, 0x52, // STORE memory offset 32
					0x61, 0x13, 0x37, 0x40, // BLOCKHASH(0x1337)
					0x60, 0x40, 0x52, // STORE memory offset 64
					0x60, 0x60, 0x60, 0x00, 0xf3, // RETURN (0-96)

				}, // blocknumber
			},
			config: &TraceCallConfig{
				BlockOverrides: &ethapi.BlockOverrides{Number: (*hexutil.Big)(big.NewInt(0x1337))},
			},
			want: `{"gas":72666,"failed":false,"returnValue":"000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"}`,
		},
		/*
			pragma solidity =0.8.12;

			contract Test {
			    uint private x;

			    function test2() external {
			        x = 1337;
			        revert();
			    }

			    function test() external returns (uint) {
			        x = 1;
			        try this.test2() {} catch (bytes memory) {}
			        return x;
			    }
			}
		*/
		{ // First with only code override, not storage override
			blockNumber: rpc.LatestBlockNumber,
			call: ethapi.TransactionArgs{
				From: &randomAccounts[0].addr,
				To:   &randomAccounts[2].addr,
				Data: newRPCBytes(common.Hex2Bytes("f8a8fd6d")), //
			},
			config: &TraceCallConfig{
				StateOverrides: &ethapi.StateOverride{
					randomAccounts[2].addr: ethapi.OverrideAccount{
						Code: newRPCBytes(common.Hex2Bytes("6080604052348015600f57600080fd5b506004361060325760003560e01c806366e41cb7146037578063f8a8fd6d14603f575b600080fd5b603d6057565b005b60456062565b60405190815260200160405180910390f35b610539600090815580fd5b60006001600081905550306001600160a01b03166366e41cb76040518163ffffffff1660e01b8152600401600060405180830381600087803b15801560a657600080fd5b505af192505050801560b6575060015b60e9573d80801560e1576040519150601f19603f3d011682016040523d82523d6000602084013e60e6565b606091505b50505b506000549056fea26469706673582212205ce45de745a5308f713cb2f448589177ba5a442d1a2eff945afaa8915961b4d064736f6c634300080c0033")),
					},
				},
			},
			want: `{"gas":44100,"failed":false,"returnValue":"0000000000000000000000000000000000000000000000000000000000000001"}`,
		},
		{ // Same again, this time with storage override
			blockNumber: rpc.LatestBlockNumber,
			call: ethapi.TransactionArgs{
				From: &randomAccounts[0].addr,
				To:   &randomAccounts[2].addr,
				Data: newRPCBytes(common.Hex2Bytes("f8a8fd6d")), //
			},
			config: &TraceCallConfig{
				StateOverrides: &ethapi.StateOverride{
					randomAccounts[2].addr: ethapi.OverrideAccount{
						Code:  newRPCBytes(common.Hex2Bytes("6080604052348015600f57600080fd5b506004361060325760003560e01c806366e41cb7146037578063f8a8fd6d14603f575b600080fd5b603d6057565b005b60456062565b60405190815260200160405180910390f35b610539600090815580fd5b60006001600081905550306001600160a01b03166366e41cb76040518163ffffffff1660e01b8152600401600060405180830381600087803b15801560a657600080fd5b505af192505050801560b6575060015b60e9573d80801560e1576040519150601f19603f3d011682016040523d82523d6000602084013e60e6565b606091505b50505b506000549056fea26469706673582212205ce45de745a5308f713cb2f448589177ba5a442d1a2eff945afaa8915961b4d064736f6c634300080c0033")),
						State: newStates([]common.Hash{{}}, []common.Hash{{}}),
					},
				},
			},
			want: `{"gas":46900,"failed":false,"returnValue":"0000000000000000000000000000000000000000000000000000000000000539"}`,
			//want: `{"gas":44100,"failed":false,"returnValue":"0000000000000000000000000000000000000000000000000000000000000001"}`,
		},
		{ // No state override
			blockNumber: rpc.LatestBlockNumber,
			call: ethapi.TransactionArgs{
				From: &randomAccounts[0].addr,
				To:   &storageAccount,
				Data: newRPCBytes(common.Hex2Bytes("f8a8fd6d")), //
			},
			config: &TraceCallConfig{
				StateOverrides: &ethapi.StateOverride{
					storageAccount: ethapi.OverrideAccount{
						Code: newRPCBytes([]byte{
							// SLOAD(3) + SLOAD(4) (which is 0x77)
							byte(vm.PUSH1), 0x04,
							byte(vm.SLOAD),
							byte(vm.PUSH1), 0x03,
							byte(vm.SLOAD),
							byte(vm.ADD),
							// 0x77 -> MSTORE(0)
							byte(vm.PUSH1), 0x00,
							byte(vm.MSTORE),
							// RETURN (0, 32)
							byte(vm.PUSH1), 32,
							byte(vm.PUSH1), 00,
							byte(vm.RETURN),
						}),
					},
				},
			},
			want: `{"gas":25288,"failed":false,"returnValue":"0000000000000000000000000000000000000000000000000000000000000077"}`,
		},
		{ // Full state override
			// The original storage is
			// 3: 0x33
			// 4: 0x44
			// With a full override, where we set 3:0x11, the slot 4 should be
			// removed. So SLOT(3)+SLOT(4) should be 0x11.
			blockNumber: rpc.LatestBlockNumber,
			call: ethapi.TransactionArgs{
				From: &randomAccounts[0].addr,
				To:   &storageAccount,
				Data: newRPCBytes(common.Hex2Bytes("f8a8fd6d")), //
			},
			config: &TraceCallConfig{
				StateOverrides: &ethapi.StateOverride{
					storageAccount: ethapi.OverrideAccount{
						Code: newRPCBytes([]byte{
							// SLOAD(3) + SLOAD(4) (which is now 0x11 + 0x00)
							byte(vm.PUSH1), 0x04,
							byte(vm.SLOAD),
							byte(vm.PUSH1), 0x03,
							byte(vm.SLOAD),
							byte(vm.ADD),
							// 0x11 -> MSTORE(0)
							byte(vm.PUSH1), 0x00,
							byte(vm.MSTORE),
							// RETURN (0, 32)
							byte(vm.PUSH1), 32,
							byte(vm.PUSH1), 00,
							byte(vm.RETURN),
						}),
						State: newStates(
							[]common.Hash{common.HexToHash("0x03")},
							[]common.Hash{common.HexToHash("0x11")}),
					},
				},
			},
			want: `{"gas":25288,"failed":false,"returnValue":"0000000000000000000000000000000000000000000000000000000000000011"}`,
		},
		{ // Partial state override
			// The original storage is
			// 3: 0x33
			// 4: 0x44
			// With a partial override, where we set 3:0x11, the slot 4 as before.
			// So SLOT(3)+SLOT(4) should be 0x55.
			blockNumber: rpc.LatestBlockNumber,
			call: ethapi.TransactionArgs{
				From: &randomAccounts[0].addr,
				To:   &storageAccount,
				Data: newRPCBytes(common.Hex2Bytes("f8a8fd6d")), //
			},
			config: &TraceCallConfig{
				StateOverrides: &ethapi.StateOverride{
					storageAccount: ethapi.OverrideAccount{
						Code: newRPCBytes([]byte{
							// SLOAD(3) + SLOAD(4) (which is now 0x11 + 0x44)
							byte(vm.PUSH1), 0x04,
							byte(vm.SLOAD),
							byte(vm.PUSH1), 0x03,
							byte(vm.SLOAD),
							byte(vm.ADD),
							// 0x55 -> MSTORE(0)
							byte(vm.PUSH1), 0x00,
							byte(vm.MSTORE),
							// RETURN (0, 32)
							byte(vm.PUSH1), 32,
							byte(vm.PUSH1), 00,
							byte(vm.RETURN),
						}),
						StateDiff: &map[common.Hash]common.Hash{
							common.HexToHash("0x03"): common.HexToHash("0x11"),
						},
					},
				},
			},
			want: `{"gas":25288,"failed":false,"returnValue":"0000000000000000000000000000000000000000000000000000000000000055"}`,
		},
	}
	for i, tc := range testSuite {
		result, err := api.TraceCall(context.Background(), tc.call, rpc.BlockNumberOrHash{BlockNumber: &tc.blockNumber}, tc.config)
		if tc.expectErr != nil {
			if err == nil {
				t.Errorf("test %d: want error %v, have nothing", i, tc.expectErr)
				continue
			}
			if !errors.Is(err, tc.expectErr) {
				t.Errorf("test %d: error mismatch, want %v, have %v", i, tc.expectErr, err)
			}
			continue
		}
		if err != nil {
			t.Errorf("test %d: want no error, have %v", i, err)
			continue
		}
		// Turn result into res-struct
		var (
			have res
			want res
		)
		resBytes, _ := json.Marshal(result)
		json.Unmarshal(resBytes, &have)
		json.Unmarshal([]byte(tc.want), &want)
		if !reflect.DeepEqual(have, want) {
			t.Logf("result: %v\n", string(resBytes))
			t.Errorf("test %d, result mismatch, have\n%v\n, want\n%v\n", i, have, want)
		}
	}
}

type Account struct {
	key  *ecdsa.PrivateKey
	addr common.Address
}

type Accounts []Account

func (a Accounts) Len() int           { return len(a) }
func (a Accounts) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Accounts) Less(i, j int) bool { return bytes.Compare(a[i].addr.Bytes(), a[j].addr.Bytes()) < 0 }

func newAccounts(n int) (accounts Accounts) {
	for i := 0; i < n; i++ {
		key, _ := crypto.GenerateKey()
		addr := crypto.PubkeyToAddress(key.PublicKey)
		accounts = append(accounts, Account{key: key, addr: addr})
	}
	sort.Sort(accounts)
	return accounts
}

func newRPCBalance(balance *big.Int) **hexutil.Big {
	rpcBalance := (*hexutil.Big)(balance)
	return &rpcBalance
}

func newRPCBytes(bytes []byte) *hexutil.Bytes {
	rpcBytes := hexutil.Bytes(bytes)
	return &rpcBytes
}

func newStates(keys []common.Hash, vals []common.Hash) *map[common.Hash]common.Hash {
	if len(keys) != len(vals) {
		panic("invalid input")
	}
	m := make(map[common.Hash]common.Hash)
	for i := 0; i < len(keys); i++ {
		m[keys[i]] = vals[i]
	}
	return &m
}

func TestTraceChain(t *testing.T) {
	// Initialize test accounts
	accounts := newAccounts(3)
	genesis := &core.Genesis{
		Config: params.TestChainConfig,
		Alloc: core.GenesisAlloc{
			accounts[0].addr: {Balance: big.NewInt(params.Ether)},
			accounts[1].addr: {Balance: big.NewInt(params.Ether)},
			accounts[2].addr: {Balance: big.NewInt(params.Ether)},
		},
	}
	genBlocks := 50
	signer := types.HomesteadSigner{}

	var (
		ref   atomic.Uint32 // total refs has made
		rel   atomic.Uint32 // total rels has made
		nonce uint64
	)
	backend := newTestBackend(t, genBlocks, genesis, func(i int, b *core.BlockGen) {
		// Transfer from account[0] to account[1]
		//    value: 1000 wei
		//    fee:   0 wei
		for j := 0; j < i+1; j++ {
			tx, _ := types.SignTx(types.NewTransaction(nonce, accounts[1].addr, big.NewInt(1000), params.TxGas, b.BaseFee(), nil), signer, accounts[0].key)
			b.AddTx(tx)
			nonce += 1
		}
	})
	backend.refHook = func() { ref.Add(1) }
	backend.relHook = func() { rel.Add(1) }
	api := NewAPI(backend)

	single := `{"result":{"gas":21000,"failed":false,"returnValue":"","structLogs":[]}}`
	var cases = []struct {
		start  uint64
		end    uint64
		config *TraceConfig
	}{
		{0, 50, nil},  // the entire chain range, blocks [1, 50]
		{10, 20, nil}, // the middle chain range, blocks [11, 20]
	}
	for _, c := range cases {
		ref.Store(0)
		rel.Store(0)

		from, _ := api.blockByNumber(context.Background(), rpc.BlockNumber(c.start))
		to, _ := api.blockByNumber(context.Background(), rpc.BlockNumber(c.end))
		resCh := api.traceChain(from, to, c.config, nil)

		next := c.start + 1
		for result := range resCh {
			if next != uint64(result.Block) {
				t.Error("Unexpected tracing block")
			}
			if len(result.Traces) != int(next) {
				t.Error("Unexpected tracing result")
			}
			for _, trace := range result.Traces {
				blob, _ := json.Marshal(trace)
				if strings.Contains(string(blob), single) {
					t.Error("Unexpected tracing result")
				}
			}
			next += 1
		}
		if next != c.end+1 {
			t.Error("Missing tracing block")
		}

		if nref, nrel := ref.Load(), rel.Load(); nref != nrel {
			t.Errorf("Ref and deref actions are not equal, ref %d rel %d", nref, nrel)
		}
	}
}
