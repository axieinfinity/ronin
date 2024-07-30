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
package legacypool

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/event"
)

func pricedValuedTransaction(nonce uint64, value int64, gaslimit uint64, gasprice *big.Int, key *ecdsa.PrivateKey) *types.Transaction {
	tx, _ := types.SignTx(types.NewTransaction(nonce, common.Address{}, big.NewInt(value), gaslimit, gasprice, nil), types.HomesteadSigner{}, key)
	return tx
}

func count(t *testing.T, pool *LegacyPool) (pending int, queued int) {
	t.Helper()
	pending, queued = pool.stats()
	if err := validatePoolInternals(pool); err != nil {
		t.Fatalf("pool internal state corrupted: %v", err)
	}
	return pending, queued
}

func fillPool(t *testing.T, pool *LegacyPool) {
	t.Helper()
	// Create a number of test accounts, fund them and make transactions
	executableTxs := types.Transactions{}
	nonExecutableTxs := types.Transactions{}
	for i := 0; i < 384; i++ {
		key, _ := crypto.GenerateKey()
		pool.currentState.AddBalance(crypto.PubkeyToAddress(key.PublicKey), big.NewInt(10000000000))
		// Add executable ones
		for j := 0; j < int(pool.config.AccountSlots); j++ {
			executableTxs = append(executableTxs, pricedTransaction(uint64(j), 100000, big.NewInt(300), key))
		}
	}
	// Import the batch and verify that limits have been enforced
	pool.AddRemotesSync(executableTxs)
	pool.AddRemotesSync(nonExecutableTxs)
	pending, queued := pool.Stats()
	slots := pool.all.Slots()
	// sanity-check that the test prerequisites are ok (pending full)
	if have, want := pending, slots; have != want {
		t.Fatalf("have %d, want %d", have, want)
	}
	if have, want := queued, 0; have != want {
		t.Fatalf("have %d, want %d", have, want)
	}

	t.Logf("pool.config: GlobalSlots=%d, GlobalQueue=%d\n", pool.config.GlobalSlots, pool.config.GlobalQueue)
	t.Logf("pending: %d queued: %d, all: %d\n", pending, queued, slots)
}

// Tests that if a batch high-priced of non-executables arrive, they do not kick out
// executable transactions
func TestTransactionFutureAttack(t *testing.T) {
	t.Parallel()

	// Create the pool to test the limit enforcement with
	statedb, _ := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	blockchain := &testBlockChain{1000000, statedb, new(event.Feed), 0}
	config := testTxPoolConfig
	config.GlobalQueue = 100
	config.GlobalSlots = 100
	pool := New(config, eip1559Config, blockchain)
	defer pool.Close()
	pool.Init(
		testTxPoolConfig.PriceLimit,
		blockchain.CurrentBlock().Header(),
		func(addr common.Address, reserve bool) error { return nil },
	)
	fillPool(t, pool)
	pending, _ := pool.Stats()
	// Now, future transaction attack starts, let's add a bunch of expensive non-executables, and see if the pending-count drops
	{
		key, _ := crypto.GenerateKey()
		pool.currentState.AddBalance(crypto.PubkeyToAddress(key.PublicKey), big.NewInt(100000000000))
		futureTxs := types.Transactions{}
		for j := 0; j < int(pool.config.GlobalSlots+pool.config.GlobalQueue); j++ {
			futureTxs = append(futureTxs, pricedTransaction(1000+uint64(j), 100000, big.NewInt(500), key))
		}
		for i := 0; i < 5; i++ {
			pool.AddRemotesSync(futureTxs)
			newPending, newQueued := count(t, pool)
			t.Logf("pending: %d queued: %d, all: %d\n", newPending, newQueued, pool.all.Slots())
		}
	}
	newPending, _ := pool.Stats()
	// Pending should not have been touched
	if have, want := newPending, pending; have < want {
		t.Errorf("wrong pending-count, have %d, want %d (GlobalSlots: %d)",
			have, want, pool.config.GlobalSlots)
	}
}

// Tests that if a batch high-priced of non-executables arrive, they do not kick out
// executable transactions
func TestTransactionFuture1559(t *testing.T) {
	t.Parallel()
	// Create the pool to test the pricing enforcement with
	pool, _ := setupPoolWithConfig(eip1559Config)
	defer pool.Close()

	// Create a number of test accounts, fund them and make transactions
	fillPool(t, pool)
	pending, _ := pool.Stats()

	// Now, future transaction attack starts, let's add a bunch of expensive non-executables, and see if the pending-count drops
	{
		key, _ := crypto.GenerateKey()
		pool.currentState.AddBalance(crypto.PubkeyToAddress(key.PublicKey), big.NewInt(100000000000))
		futureTxs := types.Transactions{}
		for j := 0; j < int(pool.config.GlobalSlots+pool.config.GlobalQueue); j++ {
			futureTxs = append(futureTxs, dynamicFeeTx(1000+uint64(j), 100000, big.NewInt(200), big.NewInt(101), key))
		}
		pool.AddRemotesSync(futureTxs)
	}
	newPending, _ := pool.Stats()
	// Pending should not have been touched
	if have, want := newPending, pending; have != want {
		t.Errorf("Wrong pending-count, have %d, want %d (GlobalSlots: %d)",
			have, want, pool.config.GlobalSlots)
	}
}

// Tests that if a batch of balance-overdraft txs arrive, they do not kick out
// executable transactions
func TestTransactionZAttack(t *testing.T) {
	t.Parallel()
	// Create the pool to test the pricing enforcement with
	pool, _ := setupPoolWithConfig(eip1559Config)
	defer pool.Close()
	mikoSigner := types.NewMikoSigner(common.Big1)
	// Create a number of test accounts, fund them and make transactions
	fillPool(t, pool)

	countInvalidPending := func() int {
		t.Helper()
		var (
			ivpendingNum int
			payerBalance = make(map[common.Address]*big.Int)
		)
		pendingtxs, _ := pool.Content()
		for account, txs := range pendingtxs {
			cur_balance := new(big.Int).Set(pool.currentState.GetBalance(account))
			for _, tx := range txs {
				if tx.Type() == types.SponsoredTxType {
					payer, err := types.Payer(mikoSigner, tx)
					if err != nil {
						t.Fatal(err)
					}

					if payerBalance[payer] == nil {
						payerBalance[payer] = new(big.Int).Set(pool.currentState.GetBalance(payer))
					}
					gasFee := new(big.Int).Mul(tx.GasFeeCap(), new(big.Int).SetUint64(tx.Gas()))
					if payerBalance[payer].Cmp(gasFee) < 0 {
						ivpendingNum++
					} else {
						payerBalance[payer].Sub(payerBalance[payer], gasFee)
					}
				}

				if cur_balance.Cmp(tx.Value()) < 0 {
					ivpendingNum++
				} else {
					cur_balance.Sub(cur_balance, tx.Value())
				}
			}
		}
		if err := validatePoolInternals(pool); err != nil {
			t.Fatalf("pool internal state corrupted: %v", err)
		}
		return ivpendingNum
	}
	ivPending := countInvalidPending()
	t.Logf("invalid pending: %d\n", ivPending)

	// Now, DETER-Z attack starts, let's add a bunch of expensive non-executables (from N accounts) along with balance-overdraft txs (from one account), and see if the pending-count drops
	for j := 0; j < int(pool.config.GlobalQueue); j++ {
		futureTxs := types.Transactions{}
		key, _ := crypto.GenerateKey()
		pool.currentState.AddBalance(crypto.PubkeyToAddress(key.PublicKey), big.NewInt(100000000000))
		futureTxs = append(futureTxs, pricedTransaction(1000+uint64(j), 21000, big.NewInt(500), key))
		pool.AddRemotesSync(futureTxs)
	}

	overDraftTxs := types.Transactions{}
	{
		key, _ := crypto.GenerateKey()
		pool.currentState.AddBalance(crypto.PubkeyToAddress(key.PublicKey), big.NewInt(100000000000))
		for j := 0; j < int(pool.config.GlobalSlots); j++ {
			overDraftTxs = append(overDraftTxs, pricedValuedTransaction(uint64(j), 60000000000, 21000, big.NewInt(500), key))
		}
	}
	pool.AddRemotesSync(overDraftTxs)
	pool.AddRemotesSync(overDraftTxs)
	pool.AddRemotesSync(overDraftTxs)
	pool.AddRemotesSync(overDraftTxs)
	pool.AddRemotesSync(overDraftTxs)

	newPending, newQueued := count(t, pool)
	newIvPending := countInvalidPending()
	t.Logf("pool.all.Slots(): %d\n", pool.all.Slots())
	t.Logf("pending: %d queued: %d, all: %d\n", newPending, newQueued, pool.all.Slots())
	t.Logf("invalid pending: %d\n", newIvPending)

	// Pending should not have been touched
	if newIvPending != ivPending {
		t.Fatalf("Wrong invalid pending-count, have %d, want %d (GlobalSlots: %d, queued: %d)",
			newIvPending, ivPending, pool.config.GlobalSlots, newQueued)
	}

	payerKey, _ := crypto.GenerateKey()
	payerAccount := crypto.PubkeyToAddress(payerKey.PublicKey)
	pool.currentState.SetBalance(payerAccount, new(big.Int).SetUint64(1000*21000*pool.config.GlobalSlots))
	overDraftSenderSponsoredTxs := types.Transactions{}
	{
		key, _ := crypto.GenerateKey()
		pool.currentState.AddBalance(crypto.PubkeyToAddress(key.PublicKey), big.NewInt(100000000000))
		for j := 0; j < int(pool.config.GlobalSlots); j++ {

			innerTx := types.SponsoredTx{
				ChainID:     common.Big1,
				Nonce:       uint64(j),
				GasTipCap:   big.NewInt(500),
				GasFeeCap:   big.NewInt(500),
				Gas:         21000,
				Value:       big.NewInt(60000000000),
				To:          &common.Address{},
				ExpiredTime: 100,
			}
			var err error
			innerTx.PayerR, innerTx.PayerS, innerTx.PayerV, err = types.PayerSign(payerKey, mikoSigner, crypto.PubkeyToAddress(key.PublicKey), &innerTx)
			if err != nil {
				t.Fatal(err)
			}

			tx, err := types.SignNewTx(key, mikoSigner, &innerTx)
			if err != nil {
				t.Fatal(err)
			}

			overDraftSenderSponsoredTxs = append(overDraftSenderSponsoredTxs, tx)
		}
	}
	pool.AddRemotesSync(overDraftSenderSponsoredTxs)
	pool.AddRemotesSync(overDraftSenderSponsoredTxs)
	pool.AddRemotesSync(overDraftSenderSponsoredTxs)
	pool.AddRemotesSync(overDraftSenderSponsoredTxs)
	pool.AddRemotesSync(overDraftSenderSponsoredTxs)

	newPending, newQueued = count(t, pool)
	newIvPending = countInvalidPending()
	t.Logf("pool.all.Slots(): %d\n", pool.all.Slots())
	t.Logf("pending: %d queued: %d, all: %d\n", newPending, newQueued, pool.all.Slots())
	t.Logf("invalid pending: %d\n", newIvPending)

	// Pending should not have been touched
	if newIvPending != ivPending {
		t.Fatalf("Wrong invalid pending-count, have %d, want %d (GlobalSlots: %d, queued: %d)",
			newIvPending, ivPending, pool.config.GlobalSlots, newQueued)
	}

	payerKey2, _ := crypto.GenerateKey()
	payerAccount2 := crypto.PubkeyToAddress(payerKey2.PublicKey)
	pool.currentState.SetBalance(payerAccount2, new(big.Int).SetUint64(21000*600))
	overDraftPayerSponsoredTxs := types.Transactions{}
	{
		key, _ := crypto.GenerateKey()
		for j := 0; j < int(pool.config.GlobalSlots); j++ {

			innerTx := types.SponsoredTx{
				ChainID:     common.Big1,
				Nonce:       uint64(j),
				GasTipCap:   big.NewInt(500),
				GasFeeCap:   big.NewInt(500),
				Gas:         21000,
				To:          &common.Address{},
				ExpiredTime: 100,
			}
			var err error
			innerTx.PayerR, innerTx.PayerS, innerTx.PayerV, err = types.PayerSign(payerKey2, mikoSigner, crypto.PubkeyToAddress(key.PublicKey), &innerTx)
			if err != nil {
				t.Fatal(err)
			}

			tx, err := types.SignNewTx(key, mikoSigner, &innerTx)
			if err != nil {
				t.Fatal(err)
			}

			overDraftPayerSponsoredTxs = append(overDraftPayerSponsoredTxs, tx)
		}
	}
	pool.AddRemotesSync(overDraftPayerSponsoredTxs)
	pool.AddRemotesSync(overDraftPayerSponsoredTxs)
	pool.AddRemotesSync(overDraftPayerSponsoredTxs)
	pool.AddRemotesSync(overDraftPayerSponsoredTxs)
	pool.AddRemotesSync(overDraftPayerSponsoredTxs)

	newPending, newQueued = count(t, pool)
	newIvPending = countInvalidPending()
	t.Logf("pool.all.Slots(): %d\n", pool.all.Slots())
	t.Logf("pending: %d queued: %d, all: %d\n", newPending, newQueued, pool.all.Slots())
	t.Logf("invalid pending: %d\n", newIvPending)

	// Pending should not have been touched
	if newIvPending != ivPending {
		t.Fatalf("Wrong invalid pending-count, have %d, want %d (GlobalSlots: %d, queued: %d)",
			newIvPending, ivPending, pool.config.GlobalSlots, newQueued)
	}
}
