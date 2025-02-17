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

package ethapi

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/bloombits"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/slices"
)

func testTransactionMarshal(t *testing.T, tests []txData, config *params.ChainConfig) {
	t.Parallel()
	var (
		signer = types.LatestSigner(config)
		key, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	)

	for i, tt := range tests {
		var tx2 types.Transaction
		tx, err := types.SignNewTx(key, signer, tt.Tx)
		if err != nil {
			t.Fatalf("test %d: signing failed: %v", i, err)
		}
		// Regular transaction
		if data, err := json.Marshal(tx); err != nil {
			t.Fatalf("test %d: marshalling failed; %v", i, err)
		} else if err = tx2.UnmarshalJSON(data); err != nil {
			t.Fatalf("test %d: sunmarshal failed: %v", i, err)
		} else if want, have := tx.Hash(), tx2.Hash(); want != have {
			t.Fatalf("test %d: stx changed, want %x have %x", i, want, have)
		}

		// rpcTransaction
		rpcTx := newRPCTransaction(tx, common.Hash{}, 0, 0, nil, config)
		if data, err := json.Marshal(rpcTx); err != nil {
			t.Fatalf("test %d: marshalling failed; %v", i, err)
		} else if err = tx2.UnmarshalJSON(data); err != nil {
			t.Fatalf("test %d: unmarshal failed: %v", i, err)
		} else if want, have := tx.Hash(), tx2.Hash(); want != have {
			t.Fatalf("test %d: tx changed, want %x have %x", i, want, have)
		} else {
			want, have := tt.Want, string(data)
			require.JSONEqf(t, want, have, "test %d: rpc json not match, want %s have %s", i, want, have)
		}
	}
}

func TestTransaction_RoundTripRpcJSON(t *testing.T) {
	var (
		config = params.AllEthashProtocolChanges
		tests  = allTransactionTypes(common.Address{0xde, 0xad}, config)
	)
	config.CancunBlock = common.Big0
	testTransactionMarshal(t, tests, config)
}

type txData struct {
	Tx   types.TxData
	Want string
}

func allTransactionTypes(addr common.Address, config *params.ChainConfig) []txData {
	return []txData{
		{
			Tx: &types.LegacyTx{
				Nonce:    5,
				GasPrice: big.NewInt(6),
				Gas:      7,
				To:       &addr,
				Value:    big.NewInt(8),
				Data:     []byte{0, 1, 2, 3, 4},
				V:        big.NewInt(9),
				R:        big.NewInt(10),
				S:        big.NewInt(11),
			},
			Want: `{
				"chainId": "0x539",
				"blockHash": null,
				"blockNumber": null,
				"from": "0x71562b71999873db5b286df957af199ec94617f7",
				"gas": "0x7",
				"gasPrice": "0x6",
				"hash": "0x5f3240454cd09a5d8b1c5d651eefae7a339262875bcd2d0e6676f3d989967008",
				"input": "0x0001020304",
				"nonce": "0x5",
				"to": "0xdead000000000000000000000000000000000000",
				"transactionIndex": null,
				"value": "0x8",
				"type": "0x0",
				"v": "0xa96",
				"r": "0xbc85e96592b95f7160825d837abb407f009df9ebe8f1b9158a4b8dd093377f75",
				"s": "0x1b55ea3af5574c536967b039ba6999ef6c89cf22fc04bcb296e0e8b0b9b576f5"
				}`,
		}, {
			Tx: &types.LegacyTx{
				Nonce:    5,
				GasPrice: big.NewInt(6),
				Gas:      7,
				To:       nil,
				Value:    big.NewInt(8),
				Data:     []byte{0, 1, 2, 3, 4},
				V:        big.NewInt(32),
				R:        big.NewInt(10),
				S:        big.NewInt(11),
			},
			Want: `{
				"chainId": "0x539",
				"blockHash": null,
				"blockNumber": null,
				"from": "0x71562b71999873db5b286df957af199ec94617f7",
				"gas": "0x7",
				"gasPrice": "0x6",
				"hash": "0x806e97f9d712b6cb7e781122001380a2837531b0fc1e5f5d78174ad4cb699873",
				"input": "0x0001020304",
				"nonce": "0x5",
				"to": null,
				"transactionIndex": null,
				"value": "0x8",
				"type": "0x0",
				"v": "0xa96",
				"r": "0x9dc28b267b6ad4e4af6fe9289668f9305c2eb7a3241567860699e478af06835a",
				"s": "0xa0b51a071aa9bed2cd70aedea859779dff039e3630ea38497d95202e9b1fec7"
				}`,
		},
		{
			Tx: &types.AccessListTx{
				ChainID:  config.ChainID,
				Nonce:    5,
				GasPrice: big.NewInt(6),
				Gas:      7,
				To:       &addr,
				Value:    big.NewInt(8),
				Data:     []byte{0, 1, 2, 3, 4},
				AccessList: types.AccessList{
					types.AccessTuple{
						Address:     common.Address{0x2},
						StorageKeys: []common.Hash{types.EmptyRootHash},
					},
				},
				V: big.NewInt(32),
				R: big.NewInt(10),
				S: big.NewInt(11),
			},
			Want: `{
				"blockHash": null,
				"blockNumber": null,
				"from": "0x71562b71999873db5b286df957af199ec94617f7",
				"gas": "0x7",
				"gasPrice": "0x6",
				"hash": "0x121347468ee5fe0a29f02b49b4ffd1c8342bc4255146bb686cd07117f79e7129",
				"input": "0x0001020304",
				"nonce": "0x5",
				"to": "0xdead000000000000000000000000000000000000",
				"transactionIndex": null,
				"value": "0x8",
				"type": "0x1",
				"accessList": [
					{
						"address": "0x0200000000000000000000000000000000000000",
						"storageKeys": [
							"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"
						]
					}
				],
				"chainId": "0x539",
				"v": "0x0",
				"r": "0xf372ad499239ae11d91d34c559ffc5dab4daffc0069e03afcabdcdf231a0c16b",
				"s": "0x28573161d1f9472fa0fd4752533609e72f06414f7ab5588699a7141f65d2abf"
				}`,
			// "yParity": "0x0"
		}, {
			Tx: &types.AccessListTx{
				ChainID:  config.ChainID,
				Nonce:    5,
				GasPrice: big.NewInt(6),
				Gas:      7,
				To:       nil,
				Value:    big.NewInt(8),
				Data:     []byte{0, 1, 2, 3, 4},
				AccessList: types.AccessList{
					types.AccessTuple{
						Address:     common.Address{0x2},
						StorageKeys: []common.Hash{types.EmptyRootHash},
					},
				},
				V: big.NewInt(32),
				R: big.NewInt(10),
				S: big.NewInt(11),
			},
			Want: `{
				"blockHash": null,
				"blockNumber": null,
				"from": "0x71562b71999873db5b286df957af199ec94617f7",
				"gas": "0x7",
				"gasPrice": "0x6",
				"hash": "0x067c3baebede8027b0f828a9d933be545f7caaec623b00684ac0659726e2055b",
				"input": "0x0001020304",
				"nonce": "0x5",
				"to": null,
				"transactionIndex": null,
				"value": "0x8",
				"type": "0x1",
				"accessList": [
					{
						"address": "0x0200000000000000000000000000000000000000",
						"storageKeys": [
							"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"
						]
					}
				],
				"chainId": "0x539",
				"v": "0x1",
				"r": "0x542981b5130d4613897fbab144796cb36d3cb3d7807d47d9c7f89ca7745b085c",
				"s": "0x7425b9dd6c5deaa42e4ede35d0c4570c4624f68c28d812c10d806ffdf86ce63"
				}`,
		}, {
			Tx: &types.DynamicFeeTx{
				ChainID:   config.ChainID,
				Nonce:     5,
				GasTipCap: big.NewInt(6),
				GasFeeCap: big.NewInt(9),
				Gas:       7,
				To:        &addr,
				Value:     big.NewInt(8),
				Data:      []byte{0, 1, 2, 3, 4},
				AccessList: types.AccessList{
					types.AccessTuple{
						Address:     common.Address{0x2},
						StorageKeys: []common.Hash{types.EmptyRootHash},
					},
				},
				V: big.NewInt(32),
				R: big.NewInt(10),
				S: big.NewInt(11),
			},
			Want: `{
				"blockHash": null,
				"blockNumber": null,
				"from": "0x71562b71999873db5b286df957af199ec94617f7",
				"gas": "0x7",
				"gasPrice": "0x9",
				"maxFeePerGas": "0x9",
				"maxPriorityFeePerGas": "0x6",
				"hash": "0xb63e0b146b34c3e9cb7fbabb5b3c081254a7ded6f1b65324b5898cc0545d79ff",
				"input": "0x0001020304",
				"nonce": "0x5",
				"to": "0xdead000000000000000000000000000000000000",
				"transactionIndex": null,
				"value": "0x8",
				"type": "0x2",
				"accessList": [
					{
						"address": "0x0200000000000000000000000000000000000000",
						"storageKeys": [
							"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"
						]
					}
				],
				"chainId": "0x539",
				"v": "0x1",
				"r": "0x3b167e05418a8932cd53d7578711fe1a76b9b96c48642402bb94978b7a107e80",
				"s": "0x22f98a332d15ea2cc80386c1ebaa31b0afebfa79ebc7d039a1e0074418301fef"
				}`,
		}, {
			Tx: &types.DynamicFeeTx{
				ChainID:    config.ChainID,
				Nonce:      5,
				GasTipCap:  big.NewInt(6),
				GasFeeCap:  big.NewInt(9),
				Gas:        7,
				To:         nil,
				Value:      big.NewInt(8),
				Data:       []byte{0, 1, 2, 3, 4},
				AccessList: types.AccessList{},
				V:          big.NewInt(32),
				R:          big.NewInt(10),
				S:          big.NewInt(11),
			},
			Want: `{
				"blockHash": null,
				"blockNumber": null,
				"from": "0x71562b71999873db5b286df957af199ec94617f7",
				"gas": "0x7",
				"gasPrice": "0x9",
				"maxFeePerGas": "0x9",
				"maxPriorityFeePerGas": "0x6",
				"hash": "0xcbab17ee031a9d5b5a09dff909f0a28aedb9b295ac0635d8710d11c7b806ec68",
				"input": "0x0001020304",
				"nonce": "0x5",
				"to": null,
				"transactionIndex": null,
				"value": "0x8",
				"type": "0x2",
				"accessList": [],
				"chainId": "0x539",
				"v": "0x0",
				"r": "0x6446b8a682db7e619fc6b4f6d1f708f6a17351a41c7fbd63665f469bc78b41b9",
				"s": "0x7626abc15834f391a117c63450047309dbf84c5ce3e8e609b607062641e2de43"
				}`,
		},
		{
			Tx: &types.BlobTx{
				Nonce:      6,
				GasTipCap:  uint256.NewInt(1),
				GasFeeCap:  uint256.NewInt(5),
				Gas:        6,
				To:         addr,
				BlobFeeCap: uint256.NewInt(1),
				BlobHashes: []common.Hash{{1}},
				Value:      new(uint256.Int),
				V:          uint256.NewInt(32),
				R:          uint256.NewInt(10),
				S:          uint256.NewInt(11),
			},
			Want: `{
                "blockHash": null,
                "blockNumber": null,
                "from": "0x71562b71999873db5b286df957af199ec94617f7",
                "gas": "0x6",
                "gasPrice": "0x5",
                "maxFeePerGas": "0x5",
                "maxPriorityFeePerGas": "0x1",	
                "maxFeePerBlobGas": "0x1",
                "hash": "0x3a65e97c6dadf3b09016abdaee312954ab5d20a95939bdf90f626afe1832854c",
                "input": "0x",
                "nonce": "0x6",
                "to": "0xdead000000000000000000000000000000000000",
                "transactionIndex": null,
                "value": "0x0",
                "type": "0x3",
                "accessList": [],
                "chainId": "0x539",
                "blobVersionedHashes": [
                    "0x0100000000000000000000000000000000000000000000000000000000000000"
                ],
                "v": "0x0",
                "r": "0xa21180f1de7bb3180c8979bc92aca17fdc0391d14211e9737ad1aa6d785b38df",
                "s": "0x2d87612bccb7bd71265b13154cce4ea8d66c376d452d52f44ab0bfb8e0082ee9"
				}`,
		},
	}
}

type testBackend struct {
	db      ethdb.Database
	chain   *core.BlockChain
	pending *types.Block
	accman  *accounts.Manager
	acc     accounts.Account
}

func newTestAccountManager(t *testing.T) (*accounts.Manager, accounts.Account) {
	var (
		dir = t.TempDir()
		am  = accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: true, EnableSigningMethods: true})
		b   = keystore.NewKeyStore(dir, 2, 1)
		// testPassphrase =ethkey
		testKey, _ = crypto.GenerateKey()
	)
	acc, err := b.ImportECDSA(testKey, "")
	if err != nil {
		t.Fatalf("failed to create test account: %v", err)
	}
	if err := b.Unlock(acc, ""); err != nil {
		t.Fatalf("failed to unlock account: %v\n", err)
	}
	am.AddBackend(b)
	return am, acc
}

func newTestBackend(t *testing.T, n int, gspec *core.Genesis, engine consensus.Engine, generator func(i int, b *core.BlockGen)) *testBackend {
	var (
		cacheConfig = &core.CacheConfig{
			TrieCleanLimit:    256,
			TrieDirtyLimit:    256,
			TrieTimeLimit:     5 * time.Minute,
			SnapshotLimit:     0,
			TrieDirtyDisabled: true, // Archive mode
		}
	)
	accman, acc := newTestAccountManager(t)
	gspec.Alloc[acc.Address] = core.GenesisAccount{Balance: big.NewInt(params.Ether)}
	// Generate blocks for testing
	db, blocks, _ := core.GenerateChainWithGenesis(gspec, engine, n, generator)
	txlookupLimit := uint64(0)
	chain, err := core.NewBlockChain(db, cacheConfig, gspec, nil, engine, vm.Config{}, nil, &txlookupLimit)
	if err != nil {
		t.Fatalf("failed to create tester chain: %v", err)
	}
	if n, err := chain.InsertChain(blocks, nil); err != nil {
		t.Fatalf("block %d: failed to insert into chain: %v", n, err)
	}

	backend := &testBackend{db: db, chain: chain, accman: accman, acc: acc}
	return backend
}

func (b *testBackend) setPendingBlock(block *types.Block) {
	b.pending = block
}

func (b testBackend) SyncProgress() ethereum.SyncProgress { return ethereum.SyncProgress{} }
func (b testBackend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return big.NewInt(0), nil
}
func (b testBackend) FeeHistory(ctx context.Context, blockCount int, lastBlock rpc.BlockNumber, rewardPercentiles []float64) (*big.Int, [][]*big.Int, []*big.Int, []float64, []*big.Int, []float64, error) {
	return nil, nil, nil, nil, nil, nil, nil
}
func (b testBackend) BlobBaseFee(ctx context.Context) *big.Int {
	return new(big.Int)
}
func (b testBackend) ChainDb() ethdb.Database           { return b.db }
func (b testBackend) AccountManager() *accounts.Manager { return b.accman }
func (b testBackend) ExtRPCEnabled() bool               { return false }
func (b testBackend) RPCGasCap() uint64                 { return 10000000 }
func (b testBackend) RPCEVMTimeout() time.Duration      { return time.Second }
func (b testBackend) RPCTxFeeCap() float64              { return 0 }
func (b testBackend) UnprotectedAllowed() bool          { return false }
func (b testBackend) SetHead(number uint64)             {}
func (b testBackend) HeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Header, error) {
	if number == rpc.LatestBlockNumber {
		return b.chain.CurrentHeader(), nil
	}
	if number == rpc.PendingBlockNumber && b.pending != nil {
		return b.pending.Header(), nil
	}
	return b.chain.GetHeaderByNumber(uint64(number)), nil
}
func (b testBackend) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	return b.chain.GetHeaderByHash(hash), nil
}
func (b testBackend) HeaderByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*types.Header, error) {
	if blockNr, ok := blockNrOrHash.Number(); ok {
		return b.HeaderByNumber(ctx, blockNr)
	}
	if blockHash, ok := blockNrOrHash.Hash(); ok {
		return b.HeaderByHash(ctx, blockHash)
	}
	panic("unknown type rpc.BlockNumberOrHash")
}
func (b testBackend) CurrentHeader() *types.Header { return b.chain.CurrentHeader() }
func (b testBackend) CurrentBlock() *types.Block   { return b.chain.CurrentBlock() }
func (b testBackend) BlockByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Block, error) {
	if number == rpc.LatestBlockNumber {
		head := b.chain.CurrentBlock()
		return b.chain.GetBlock(head.Hash(), head.NumberU64()), nil
	}
	if number == rpc.PendingBlockNumber {
		return b.pending, nil
	}
	return b.chain.GetBlockByNumber(uint64(number)), nil
}
func (b testBackend) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	return b.chain.GetBlockByHash(hash), nil
}
func (b testBackend) BlockByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*types.Block, error) {
	if blockNr, ok := blockNrOrHash.Number(); ok {
		return b.BlockByNumber(ctx, blockNr)
	}
	if blockHash, ok := blockNrOrHash.Hash(); ok {
		return b.BlockByHash(ctx, blockHash)
	}
	panic("unknown type rpc.BlockNumberOrHash")
}
func (b testBackend) GetBody(ctx context.Context, hash common.Hash, number rpc.BlockNumber) (*types.Body, error) {
	return b.chain.GetBlock(hash, uint64(number.Int64())).Body(), nil
}
func (b testBackend) StateAndHeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*state.StateDB, *types.Header, error) {
	var (
		header *types.Header
		err    error
	)
	if number == rpc.PendingBlockNumber {
		if b.pending == nil {
			panic("pending state not found")
		}
		header = b.pending.Header()
	} else {
		header, err = b.HeaderByNumber(ctx, number)
		if err != nil {
			return nil, nil, err
		}
	}
	if header == nil {
		return nil, nil, errors.New("header not found")
	}
	stateDb, err := b.chain.StateAt(header.Root)
	return stateDb, header, err
}
func (b testBackend) StateAndHeaderByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*state.StateDB, *types.Header, error) {
	if blockNr, ok := blockNrOrHash.Number(); ok {
		return b.StateAndHeaderByNumber(ctx, blockNr)
	}
	panic("only implemented for number")
}
func (b testBackend) PendingBlockAndReceipts() (*types.Block, types.Receipts) { panic("implement me") }
func (b testBackend) GetReceipts(ctx context.Context, hash common.Hash) (types.Receipts, error) {
	header, err := b.HeaderByHash(ctx, hash)
	if header == nil || err != nil {
		return nil, err
	}
	receipts := rawdb.ReadReceipts(b.db, hash, header.Number.Uint64(), b.chain.Config())
	return receipts, nil
}
func (b testBackend) GetTd(ctx context.Context, hash common.Hash) *big.Int {
	if b.pending != nil && hash == b.pending.Hash() {
		return nil
	}
	return big.NewInt(1)
}
func (b testBackend) GetEVM(ctx context.Context, msg core.Message, state *state.StateDB, header *types.Header, vmConfig *vm.Config, blockContext *vm.BlockContext) (*vm.EVM, func() error, error) {
	vmError := func() error { return nil }
	if vmConfig == nil {
		vmConfig = b.chain.GetVMConfig()
	}
	txContext := core.NewEVMTxContext(msg)
	context := core.NewEVMBlockContext(header, b.chain, nil)
	if blockContext != nil {
		context = *blockContext
	}
	return vm.NewEVM(context, txContext, state, b.chain.Config(), *vmConfig), vmError, nil
}
func (b testBackend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	panic("implement me")
}
func (b testBackend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	panic("implement me")
}
func (b testBackend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription {
	panic("implement me")
}
func (b testBackend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	panic("implement me")
}
func (b testBackend) GetTransaction(ctx context.Context, txHash common.Hash) (*types.Transaction, common.Hash, uint64, uint64, error) {
	tx, blockHash, blockNumber, index := rawdb.ReadTransaction(b.db, txHash)
	return tx, blockHash, blockNumber, index, nil
}
func (b testBackend) GetPoolTransactions() (types.Transactions, error)         { panic("implement me") }
func (b testBackend) GetPoolTransaction(txHash common.Hash) *types.Transaction { panic("implement me") }
func (b testBackend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	return 0, nil
}
func (b testBackend) Stats() (pending int, queued int) { panic("implement me") }
func (b testBackend) TxPoolContent() (map[common.Address][]*types.Transaction, map[common.Address][]*types.Transaction) {
	panic("implement me")
}
func (b testBackend) TxPoolContentFrom(addr common.Address) ([]*types.Transaction, []*types.Transaction) {
	panic("implement me")
}
func (b testBackend) SubscribeNewTxsEvent(events chan<- core.NewTxsEvent) event.Subscription {
	panic("implement me")
}
func (b testBackend) ChainConfig() *params.ChainConfig { return b.chain.Config() }
func (b testBackend) Engine() consensus.Engine         { return b.chain.Engine() }
func (b testBackend) GetLogs(ctx context.Context, blockHash common.Hash) ([][]*types.Log, error) {
	panic("implement me")
}
func (b testBackend) SubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent) event.Subscription {
	panic("implement me")
}
func (b testBackend) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
	panic("implement me")
}
func (b testBackend) SubscribePendingLogsEvent(ch chan<- []*types.Log) event.Subscription {
	panic("implement me")
}
func (b testBackend) SubscribeReorgEvent(ch chan<- core.ReorgEvent) event.Subscription {
	panic("implement me")
}
func (b testBackend) SubscribeInternalTransactionEvent(ch chan<- []*types.InternalTransaction) event.Subscription {
	panic("implement me")
}
func (b testBackend) SubscribeDirtyAccountEvent(ch chan<- []*types.DirtyStateAccount) event.Subscription {
	panic("implement me")
}
func (b testBackend) BloomStatus() (uint64, uint64) { panic("implement me") }
func (b testBackend) ServiceFilter(ctx context.Context, session *bloombits.MatcherSession) {
	panic("implement me")
}

func (b testBackend) BlobSidecarsByHash(ctx context.Context, hash common.Hash) (types.BlobSidecars, error) {
	panic("implement me")
}
func (b testBackend) BlobSidecarsByNumber(ctx context.Context, number rpc.BlockNumber) (types.BlobSidecars, error) {
	panic("implement me")
}
func (b testBackend) BlobSidecarsByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (types.BlobSidecars, error) {
	panic("implement me")
}

func TestEstimateGas(t *testing.T) {
	t.Parallel()
	// Initialize test accounts
	var (
		accounts = newAccounts(2)
		genesis  = &core.Genesis{
			Config: params.TestChainConfig,
			Alloc: core.GenesisAlloc{
				accounts[0].addr: {Balance: big.NewInt(params.Ether)},
				accounts[1].addr: {Balance: big.NewInt(params.Ether)},
			},
		}
		genBlocks      = 10
		signer         = types.HomesteadSigner{}
		randomAccounts = newAccounts(2)
	)
	api := NewPublicBlockChainAPI(newTestBackend(t, genBlocks, genesis, ethash.NewFaker(), func(i int, b *core.BlockGen) {
		// Transfer from account[0] to account[1]
		//    value: 1000 wei
		//    fee:   0 wei
		tx, _ := types.SignTx(types.NewTx(&types.LegacyTx{Nonce: uint64(i), To: &accounts[1].addr, Value: big.NewInt(1000), Gas: params.TxGas, GasPrice: b.BaseFee(), Data: nil}), signer, accounts[0].key)
		b.AddTx(tx)
	}))
	var testSuite = []struct {
		blockNumber rpc.BlockNumber
		call        TransactionArgs
		overrides   StateOverride
		expectErr   error
		want        uint64
	}{
		// simple transfer on latest block
		{
			blockNumber: rpc.LatestBlockNumber,
			call: TransactionArgs{
				From:  &accounts[0].addr,
				To:    &accounts[1].addr,
				Value: (*hexutil.Big)(big.NewInt(1000)),
			},
			expectErr: nil,
			want:      21000,
		},
		// simple transfer with insufficient funds on latest block
		{
			blockNumber: rpc.LatestBlockNumber,
			call: TransactionArgs{
				From:  &randomAccounts[0].addr,
				To:    &accounts[1].addr,
				Value: (*hexutil.Big)(big.NewInt(1000)),
			},
			expectErr: core.ErrInsufficientFunds,
			want:      21000,
		},
		// empty create
		{
			blockNumber: rpc.LatestBlockNumber,
			call:        TransactionArgs{},
			expectErr:   nil,
			want:        53000,
		},
		{
			blockNumber: rpc.LatestBlockNumber,
			call:        TransactionArgs{},
			overrides: StateOverride{
				randomAccounts[0].addr: OverrideAccount{Balance: newRPCBalance(new(big.Int).Mul(big.NewInt(1), big.NewInt(params.Ether)))},
			},
			expectErr: nil,
			want:      53000,
		},
		{
			blockNumber: rpc.LatestBlockNumber,
			call: TransactionArgs{
				From:  &randomAccounts[0].addr,
				To:    &randomAccounts[1].addr,
				Value: (*hexutil.Big)(big.NewInt(1000)),
			},
			overrides: StateOverride{
				randomAccounts[0].addr: OverrideAccount{Balance: newRPCBalance(big.NewInt(0))},
			},
			expectErr: core.ErrInsufficientFunds,
		},
		// Blobs should have no effect on gas estimate
		{
			blockNumber: rpc.LatestBlockNumber,
			call: TransactionArgs{
				From:       &accounts[0].addr,
				To:         &accounts[1].addr,
				Value:      (*hexutil.Big)(big.NewInt(1)),
				BlobHashes: []common.Hash{{0x01, 0x22}},
				BlobFeeCap: (*hexutil.Big)(big.NewInt(1)),
			},
			want: 21000,
		},
	}
	for i, tc := range testSuite {
		result, err := api.EstimateGas(context.Background(), tc.call, &rpc.BlockNumberOrHash{BlockNumber: &tc.blockNumber}, &tc.overrides)
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
		if float64(result) > float64(tc.want)*(1+estimateGasErrorRatio) {
			t.Errorf("test %d, result mismatch, have\n%v\n, want\n%v\n", i, uint64(result), tc.want)
		}
	}
}

func TestCall(t *testing.T) {
	t.Parallel()
	// Initialize test accounts
	var (
		accounts = newAccounts(3)
		genesis  = &core.Genesis{
			Config: params.TestChainConfig,
			Alloc: core.GenesisAlloc{
				accounts[0].addr: {Balance: big.NewInt(params.Ether)},
				accounts[1].addr: {Balance: big.NewInt(params.Ether)},
				accounts[2].addr: {Balance: big.NewInt(params.Ether)},
			},
		}
		genBlocks = 10
		signer    = types.HomesteadSigner{}
	)
	api := NewPublicBlockChainAPI(newTestBackend(t, genBlocks, genesis, ethash.NewFaker(), func(i int, b *core.BlockGen) {
		// Transfer from account[0] to account[1]
		//    value: 1000 wei
		//    fee:   0 wei
		tx, _ := types.SignTx(types.NewTx(&types.LegacyTx{Nonce: uint64(i), To: &accounts[1].addr, Value: big.NewInt(1000), Gas: params.TxGas, GasPrice: b.BaseFee(), Data: nil}), signer, accounts[0].key)
		b.AddTx(tx)
	}))
	randomAccounts := newAccounts(3)
	var testSuite = []struct {
		blockNumber    rpc.BlockNumber
		overrides      StateOverride
		call           TransactionArgs
		blockOverrides BlockOverrides
		expectErr      error
		want           string
	}{
		// transfer on genesis
		{
			blockNumber: rpc.BlockNumber(0),
			call: TransactionArgs{
				From:  &accounts[0].addr,
				To:    &accounts[1].addr,
				Value: (*hexutil.Big)(big.NewInt(1000)),
			},
			expectErr: nil,
			want:      "0x",
		},
		// transfer on the head
		{
			blockNumber: rpc.BlockNumber(genBlocks),
			call: TransactionArgs{
				From:  &accounts[0].addr,
				To:    &accounts[1].addr,
				Value: (*hexutil.Big)(big.NewInt(1000)),
			},
			expectErr: nil,
			want:      "0x",
		},
		// transfer on a non-existent block, error expects
		{
			blockNumber: rpc.BlockNumber(genBlocks + 1),
			call: TransactionArgs{
				From:  &accounts[0].addr,
				To:    &accounts[1].addr,
				Value: (*hexutil.Big)(big.NewInt(1000)),
			},
			expectErr: errors.New("header not found"),
		},
		// transfer on the latest block
		{
			blockNumber: rpc.LatestBlockNumber,
			call: TransactionArgs{
				From:  &accounts[0].addr,
				To:    &accounts[1].addr,
				Value: (*hexutil.Big)(big.NewInt(1000)),
			},
			expectErr: nil,
			want:      "0x",
		},
		// Call which can only succeed if state is state overridden
		{
			blockNumber: rpc.LatestBlockNumber,
			call: TransactionArgs{
				From:  &randomAccounts[0].addr,
				To:    &randomAccounts[1].addr,
				Value: (*hexutil.Big)(big.NewInt(1000)),
			},
			overrides: StateOverride{
				randomAccounts[0].addr: OverrideAccount{Balance: newRPCBalance(new(big.Int).Mul(big.NewInt(1), big.NewInt(params.Ether)))},
			},
			want: "0x",
		},
		// Invalid call without state overriding
		{
			blockNumber: rpc.LatestBlockNumber,
			call: TransactionArgs{
				From:  &randomAccounts[0].addr,
				To:    &randomAccounts[1].addr,
				Value: (*hexutil.Big)(big.NewInt(1000)),
			},
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
			call: TransactionArgs{
				From: &randomAccounts[0].addr,
				To:   &randomAccounts[2].addr,
				Data: hex2Bytes("8381f58a"), // call number()
			},
			overrides: StateOverride{
				randomAccounts[2].addr: OverrideAccount{
					Code:      hex2Bytes("6080604052348015600f57600080fd5b506004361060285760003560e01c80638381f58a14602d575b600080fd5b60336049565b6040518082815260200191505060405180910390f35b6000548156fea2646970667358221220eab35ffa6ab2adfe380772a48b8ba78e82a1b820a18fcb6f59aa4efb20a5f60064736f6c63430007040033"),
					StateDiff: &map[common.Hash]common.Hash{{}: common.BigToHash(big.NewInt(123))},
				},
			},
			want: "0x000000000000000000000000000000000000000000000000000000000000007b",
		},
		// Block overrides should work
		{
			blockNumber: rpc.LatestBlockNumber,
			call: TransactionArgs{
				From: &accounts[1].addr,
				Input: &hexutil.Bytes{
					0x43,             // NUMBER
					0x60, 0x00, 0x52, // MSTORE offset 0
					0x60, 0x20, 0x60, 0x00, 0xf3,
				},
			},
			blockOverrides: BlockOverrides{Number: (*hexutil.Big)(big.NewInt(11))},
			want:           "0x000000000000000000000000000000000000000000000000000000000000000b",
		},
		// Invalid blob tx
		{
			blockNumber: rpc.LatestBlockNumber,
			call: TransactionArgs{
				From:       &accounts[1].addr,
				Input:      &hexutil.Bytes{0x00},
				BlobHashes: []common.Hash{},
			},
			expectErr: core.ErrBlobTxCreate,
		},
		// BLOBHASH opcode
		{
			blockNumber: rpc.LatestBlockNumber,
			call: TransactionArgs{
				From:       &accounts[1].addr,
				To:         &randomAccounts[2].addr,
				BlobHashes: []common.Hash{{0x01, 0x22}},
				BlobFeeCap: (*hexutil.Big)(big.NewInt(1)),
			},
			overrides: StateOverride{
				// override this bytecode, which do return the first blob hash if any
				randomAccounts[2].addr: {
					// Code: hex2Bytes("60004960005260206000f3"),
					Code: newRPCBytes([]byte{byte(vm.PUSH1), byte(0), byte(vm.BLOBHASH), byte(vm.PUSH1), byte(0), byte(vm.MSTORE), byte(vm.PUSH1), byte(0x20), byte(vm.PUSH1), byte(0), byte(vm.RETURN)}),
				},
			},
			want: "0x0122000000000000000000000000000000000000000000000000000000000000",
		},
	}
	for i, tc := range testSuite {
		result, err := api.Call(context.Background(), tc.call, &rpc.BlockNumberOrHash{BlockNumber: &tc.blockNumber}, &tc.overrides, &tc.blockOverrides)
		if tc.expectErr != nil {
			if err == nil {
				t.Errorf("test %d: want error %v, have nothing", i, tc.expectErr)
				continue
			}
			if !errors.Is(err, tc.expectErr) {
				// Second try
				if !reflect.DeepEqual(err, tc.expectErr) {
					t.Errorf("test %d: error mismatch, want %v, have %v", i, tc.expectErr, err)
				}
			}
			continue
		}
		if err != nil {
			t.Errorf("test %d: want no error, have %v", i, err)
			continue
		}
		if !reflect.DeepEqual(result.String(), tc.want) {
			t.Errorf("test %d, result mismatch, have\n%v\n, want\n%v\n", i, result.String(), tc.want)
		}
	}
}

type Account struct {
	key  *ecdsa.PrivateKey
	addr common.Address
}

func newAccounts(n int) (accounts []Account) {
	for i := 0; i < n; i++ {
		key, _ := crypto.GenerateKey()
		addr := crypto.PubkeyToAddress(key.PublicKey)
		accounts = append(accounts, Account{key: key, addr: addr})
	}
	slices.SortFunc(accounts, func(a, b Account) int { return bytes.Compare(a.addr[:], b.addr[:]) })
	return accounts
}

func newRPCBalance(balance *big.Int) **hexutil.Big {
	rpcBalance := (*hexutil.Big)(balance)
	return &rpcBalance
}

func hex2Bytes(str string) *hexutil.Bytes {
	rpcBytes := hexutil.Bytes(common.Hex2Bytes(str))
	return &rpcBytes
}

func TestHeader4844MarshalJson(t *testing.T) {
	header := types.Header{
		Number:     big.NewInt(100),
		Difficulty: big.NewInt(7),
	}
	data, err := json.Marshal(RPCMarshalHeader(&header))
	if err != nil {
		t.Fatal(err)
	}

	expect := `{"difficulty":"0x7","extraData":"0x","gasLimit":"0x0","gasUsed":"0x0","hash":"0x7638fef16ccc17d30038b807c09ca0f0bb47a6132d81253799448855504ed217","logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","miner":"0x0000000000000000000000000000000000000000","mixHash":"0x0000000000000000000000000000000000000000000000000000000000000000","nonce":"0x0000000000000000","number":"0x64","parentHash":"0x0000000000000000000000000000000000000000000000000000000000000000","receiptsRoot":"0x0000000000000000000000000000000000000000000000000000000000000000","sha3Uncles":"0x0000000000000000000000000000000000000000000000000000000000000000","size":"0x239","stateRoot":"0x0000000000000000000000000000000000000000000000000000000000000000","timestamp":"0x0","transactionsRoot":"0x0000000000000000000000000000000000000000000000000000000000000000"}`
	if string(data) != expect {
		t.Fatalf("Header mismatches, expect: %s\n got: %s", expect, string(data))
	}

	blobGasUsed := uint64(1 << 17)
	excessBlobGas := 2 * blobGasUsed
	header.BlobGasUsed = &blobGasUsed
	header.ExcessBlobGas = &excessBlobGas

	data, err = json.Marshal(RPCMarshalHeader(&header))
	if err != nil {
		t.Fatal(err)
	}

	expect = `{"blobGasUsed":"0x20000","difficulty":"0x7","excessBlobGas":"0x40000","extraData":"0x","gasLimit":"0x0","gasUsed":"0x0","hash":"0xd2bae9d64fe00db8bc637990b38432d8281604d1caf81bfe7c0b46ecc1dfd1ca","logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","miner":"0x0000000000000000000000000000000000000000","mixHash":"0x0000000000000000000000000000000000000000000000000000000000000000","nonce":"0x0000000000000000","number":"0x64","parentHash":"0x0000000000000000000000000000000000000000000000000000000000000000","receiptsRoot":"0x0000000000000000000000000000000000000000000000000000000000000000","sha3Uncles":"0x0000000000000000000000000000000000000000000000000000000000000000","size":"0x239","stateRoot":"0x0000000000000000000000000000000000000000000000000000000000000000","timestamp":"0x0","transactionsRoot":"0x0000000000000000000000000000000000000000000000000000000000000000"}`
	if string(data) != expect {
		t.Fatalf("Header mismatches, expect: %s\n got: %s", expect, string(data))
	}
}

func argsFromTransaction(tx *types.Transaction, from common.Address) TransactionArgs {
	var (
		gas        = tx.Gas()
		nonce      = tx.Nonce()
		input      = tx.Data()
		accessList *types.AccessList
	)
	if acl := tx.AccessList(); acl != nil {
		accessList = &acl
	}
	return TransactionArgs{
		From:                 &from,
		To:                   tx.To(),
		Gas:                  (*hexutil.Uint64)(&gas),
		MaxFeePerGas:         (*hexutil.Big)(tx.GasFeeCap()),
		MaxPriorityFeePerGas: (*hexutil.Big)(tx.GasTipCap()),
		Value:                (*hexutil.Big)(tx.Value()),
		Nonce:                (*hexutil.Uint64)(&nonce),
		Input:                (*hexutil.Bytes)(&input),
		ChainID:              (*hexutil.Big)(tx.ChainId()),
		AccessList:           accessList,
		BlobFeeCap:           (*hexutil.Big)(tx.BlobGasFeeCap()),
		BlobHashes:           tx.BlobHashes(),
	}
}

var (
	emptyBlob          = kzg4844.Blob{}
	emptyBlobs         = []kzg4844.Blob{emptyBlob}
	emptyBlobCommit, _ = kzg4844.BlobToCommitment(&emptyBlob)
	emptyBlobProof, _  = kzg4844.ComputeBlobProof(&emptyBlob, emptyBlobCommit)
	emptyBlobHash      = kzg4844.CalcBlobHashV1(sha256.New(), &emptyBlobCommit)
)

func TestFillBlobTransaction(t *testing.T) {
	t.Parallel()
	var (
		height   = 5
		accounts = newAccounts(2)
		to       = accounts[1].addr
		genesis  = &core.Genesis{
			Config: params.TestChainConfig,
			Alloc: core.GenesisAlloc{
				accounts[0].addr: {Balance: big.NewInt(params.Ether)},
				accounts[1].addr: {Balance: big.NewInt(params.Ether)},
			},
		}
	)

	b := newTestBackend(t, height, genesis, ethash.NewFaker(), nil)
	api := NewPublicTransactionPoolAPI(b, nil)

	type result struct {
		Hashes  []common.Hash
		Sidecar *types.BlobTxSidecar
	}
	suite := []struct {
		name string
		args TransactionArgs
		err  string
		want *result
	}{
		{
			name: "TestInvalidParamsCombination1",
			args: TransactionArgs{
				From:   &b.acc.Address,
				To:     &to,
				Value:  (*hexutil.Big)(big.NewInt(1)),
				Blobs:  []kzg4844.Blob{{}},
				Proofs: []kzg4844.Proof{{}},
			},
			err: `blob proofs provided while commitments were not`,
		},
		{
			name: "TestInvalidParamsCombination2",
			args: TransactionArgs{
				From:        &b.acc.Address,
				To:          &to,
				Value:       (*hexutil.Big)(big.NewInt(1)),
				Blobs:       []kzg4844.Blob{{}},
				Commitments: []kzg4844.Commitment{{}},
			},
			err: `blob commitments provided while proofs were not`,
		},
		{
			name: "TestInvalidParamsCount1",
			args: TransactionArgs{
				From:        &b.acc.Address,
				To:          &to,
				Value:       (*hexutil.Big)(big.NewInt(1)),
				Blobs:       []kzg4844.Blob{{}},
				Commitments: []kzg4844.Commitment{{}, {}},
				Proofs:      []kzg4844.Proof{{}, {}},
			},
			err: `number of blobs and commitments mismatch (have=2, want=1)`,
		},
		{
			name: "TestInvalidParamsCount2",
			args: TransactionArgs{
				From:        &b.acc.Address,
				To:          &to,
				Value:       (*hexutil.Big)(big.NewInt(1)),
				Blobs:       []kzg4844.Blob{{}, {}},
				Commitments: []kzg4844.Commitment{{}, {}},
				Proofs:      []kzg4844.Proof{{}},
			},
			err: `number of blobs and proofs mismatch (have=1, want=2)`,
		},
		{
			name: "TestInvalidProofVerification",
			args: TransactionArgs{
				From:        &b.acc.Address,
				To:          &to,
				Value:       (*hexutil.Big)(big.NewInt(1)),
				Blobs:       []kzg4844.Blob{{}, {}},
				Commitments: []kzg4844.Commitment{{}, {}},
				Proofs:      []kzg4844.Proof{{}, {}},
			},
			err: `failed to verify blob proof: short buffer`,
		},
		{
			name: "TestGenerateBlobHashes",
			args: TransactionArgs{
				From:        &b.acc.Address,
				To:          &to,
				Value:       (*hexutil.Big)(big.NewInt(1)),
				Blobs:       emptyBlobs,
				Commitments: []kzg4844.Commitment{emptyBlobCommit},
				Proofs:      []kzg4844.Proof{emptyBlobProof},
				Gas:         (*hexutil.Uint64)(new(uint64)),
			},
			want: &result{
				Hashes: []common.Hash{emptyBlobHash},
				Sidecar: &types.BlobTxSidecar{
					Blobs:       emptyBlobs,
					Commitments: []kzg4844.Commitment{emptyBlobCommit},
					Proofs:      []kzg4844.Proof{emptyBlobProof},
				},
			},
		},
		{
			name: "TestValidBlobHashes",
			args: TransactionArgs{
				From:        &b.acc.Address,
				To:          &to,
				Value:       (*hexutil.Big)(big.NewInt(1)),
				BlobHashes:  []common.Hash{emptyBlobHash},
				Blobs:       emptyBlobs,
				Commitments: []kzg4844.Commitment{emptyBlobCommit},
				Proofs:      []kzg4844.Proof{emptyBlobProof},
				Gas:         (*hexutil.Uint64)(new(uint64)),
			},
			want: &result{
				Hashes: []common.Hash{emptyBlobHash},
				Sidecar: &types.BlobTxSidecar{
					Blobs:       emptyBlobs,
					Commitments: []kzg4844.Commitment{emptyBlobCommit},
					Proofs:      []kzg4844.Proof{emptyBlobProof},
				},
			},
		},
		{
			name: "TestInvalidBlobHashes",
			args: TransactionArgs{
				From:        &b.acc.Address,
				To:          &to,
				Value:       (*hexutil.Big)(big.NewInt(1)),
				BlobHashes:  []common.Hash{{0x01, 0x22}},
				Blobs:       emptyBlobs,
				Commitments: []kzg4844.Commitment{emptyBlobCommit},
				Proofs:      []kzg4844.Proof{emptyBlobProof},
				Gas:         (*hexutil.Uint64)(new(uint64)),
			},
			err: fmt.Sprintf("blob hash verification failed (have=%s, want=%s)", common.Hash{0x01, 0x22}, common.BytesToHash(emptyBlobHash[:])),
		},
		{
			name: "TestGenerateBlobProofs",
			args: TransactionArgs{
				From:  &b.acc.Address,
				To:    &to,
				Value: (*hexutil.Big)(big.NewInt(1)),
				Blobs: emptyBlobs,
				Gas:   (*hexutil.Uint64)(new(uint64)),
			},
			want: &result{
				Hashes: []common.Hash{emptyBlobHash},
				Sidecar: &types.BlobTxSidecar{
					Blobs:       emptyBlobs,
					Commitments: []kzg4844.Commitment{emptyBlobCommit},
					Proofs:      []kzg4844.Proof{emptyBlobProof},
				},
			},
		},
		{
			name: "TestZeroBlobFeeCap",
			args: TransactionArgs{
				BlobFeeCap: (*hexutil.Big)(common.Big0),
			},
			err: "maxFeePerBlobGas, if specified, must be non-zero",
		},
		{
			name: "TestInvalidSidecarsProvided1",
			args: TransactionArgs{
				Blobs:              []kzg4844.Blob{emptyBlob},
				Commitments:        []kzg4844.Commitment{emptyBlobCommit},
				Proofs:             []kzg4844.Proof{emptyBlobProof},
				BlobHashes:         []common.Hash{},
				blobSidecarAllowed: true,
			},
			err: "number of blobs and hashes mismatch (have=0, want=1)",
		},
		{
			name: "TestMissingToField",
			args: TransactionArgs{
				To:                 nil,
				Nonce:              (*hexutil.Uint64)(new(uint64)),
				Blobs:              []kzg4844.Blob{emptyBlob},
				BlobFeeCap:         (*hexutil.Big)(big.NewInt(1)),
				blobSidecarAllowed: true,
			},
			err: `missing "to" in blob transaction`,
		},
		{
			name: "TestTooManyBlobs",
			args: TransactionArgs{
				Nonce:              (*hexutil.Uint64)(new(uint64)),
				Blobs:              []kzg4844.Blob{emptyBlob, emptyBlob, emptyBlob, emptyBlob, emptyBlob, emptyBlob, emptyBlob},
				BlobFeeCap:         (*hexutil.Big)(big.NewInt(1)),
				blobSidecarAllowed: true,
			},
			err: `too many blobs in transaction (have=7, max=6)`,
		},
	}
	for _, tc := range suite {
		t.Run(tc.name, func(t *testing.T) {
			res, err := api.FillTransaction(context.Background(), tc.args)
			if len(tc.err) > 0 {
				if err == nil {
					t.Fatalf("missing error. want: %s", tc.err)
				} else if err.Error() != tc.err {
					t.Fatalf("error mismatch. want: %s, have: %s", tc.err, err.Error())
				}
				return
			}
			if err != nil && len(tc.err) == 0 {
				t.Fatalf("expected no error. have: %s", err)
			}
			if res == nil {
				t.Fatal("result missing")
			}
			want, err := json.Marshal(tc.want)
			if err != nil {
				t.Fatalf("failed to encode expected: %v", err)
			}
			have, err := json.Marshal(result{Hashes: res.Tx.BlobHashes(), Sidecar: res.Tx.BlobTxSidecar()})
			if err != nil {
				t.Fatalf("failed to encode computed sidecar: %v", err)
			}
			if !bytes.Equal(have, want) {
				t.Errorf("blob sidecar mismatch. Have: %s, want: %s", have, want)
			}
		})
	}
}

func TestBlobTransactionApi(t *testing.T) {
	t.Parallel()
	// Initialize test accounts
	var (
		accounts = newAccounts(2)
		genesis  = &core.Genesis{
			Config: params.TestChainConfig,
			Alloc: core.GenesisAlloc{
				accounts[0].addr: {Balance: big.NewInt(params.Ether)},
				accounts[1].addr: {Balance: big.NewInt(params.Ether)},
			},
		}
		height = 5
	)
	b := newTestBackend(t, height, genesis, ethash.NewFaker(), func(i int, b *core.BlockGen) {})
	b.setPendingBlock(b.CurrentBlock())
	api := NewPublicTransactionPoolAPI(b, nil)
	res, err := api.FillTransaction(context.Background(), TransactionArgs{
		Nonce:      (*hexutil.Uint64)(new(uint64)),
		From:       &accounts[0].addr,
		To:         &accounts[1].addr,
		Value:      (*hexutil.Big)(big.NewInt(1)),
		BlobHashes: []common.Hash{{0x01, 0x22}},
	})
	if err != nil {
		t.Fatalf("failed to fill tx defaults: %v\n", err)
	}

	t.Run("TestSignBlobTransaction", func(t *testing.T) {
		// Test sign transaction
		_, err = api.SignTransaction(context.Background(), argsFromTransaction(res.Tx, b.acc.Address))
		if err != nil {
			t.Fatalf("should not fail on blob transaction %s", err)
		}
	})

	t.Run("TestSendBlobTransaction", func(t *testing.T) {
		_, err = api.SendTransaction(context.Background(), argsFromTransaction(res.Tx, b.acc.Address))
		if err == nil {
			t.Errorf("sending tx should have failed")
		} else if !errors.Is(err, errBlobTxNotSupported) {
			t.Errorf("unexpected error. Have %v, want %v\n", err, errBlobTxNotSupported)
		}
	})
}

func newRPCBytes(bytes []byte) *hexutil.Bytes {
	rpcBytes := hexutil.Bytes(bytes)
	return &rpcBytes
}
