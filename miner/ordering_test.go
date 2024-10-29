package miner

import (
	"crypto/ecdsa"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
)

func TestTransactionPriceNonceSortLegacy(t *testing.T) {
	testTransactionPriceNonceSort(t, nil)
}

func TestTransactionPriceNonceSort1559(t *testing.T) {
	testTransactionPriceNonceSort(t, big.NewInt(0))
	testTransactionPriceNonceSort(t, big.NewInt(5))
	testTransactionPriceNonceSort(t, big.NewInt(50))
}

// Tests that transactions can be correctly sorted according to their price in
// decreasing order, but at the same time with increasing nonces when issued by
// the same account.
func testTransactionPriceNonceSort(t *testing.T, baseFee *big.Int) {
	// Generate a batch of accounts to start with
	keys := make([]*ecdsa.PrivateKey, 25)
	for i := 0; i < len(keys); i++ {
		keys[i], _ = crypto.GenerateKey()
	}
	signer := types.LatestSignerForChainID(common.Big1)

	// Generate a batch of transactions with overlapping values, but shifted nonces
	groups := map[common.Address][]*txpool.LazyTransaction{}
	expectedCount := 0
	for start, key := range keys {
		addr := crypto.PubkeyToAddress(key.PublicKey)
		count := 25
		for i := 0; i < 25; i++ {
			var tx *types.Transaction
			gasFeeCap := rand.Intn(50)
			if baseFee == nil {
				tx = types.NewTx(&types.LegacyTx{
					Nonce:    uint64(start + i),
					To:       &common.Address{},
					Value:    big.NewInt(100),
					Gas:      100,
					GasPrice: big.NewInt(int64(gasFeeCap)),
					Data:     nil,
				})
			} else {
				tx = types.NewTx(&types.DynamicFeeTx{
					Nonce:     uint64(start + i),
					To:        &common.Address{},
					Value:     big.NewInt(100),
					Gas:       100,
					GasFeeCap: big.NewInt(int64(gasFeeCap)),
					GasTipCap: big.NewInt(int64(rand.Intn(gasFeeCap + 1))),
					Data:      nil,
				})
				if count == 25 && int64(gasFeeCap) < baseFee.Int64() {
					count = i
				}
			}
			tx, err := types.SignTx(tx, signer, key)
			if err != nil {
				t.Fatalf("failed to sign tx: %s", err)
			}
			groups[addr] = append(groups[addr], &txpool.LazyTransaction{
				Tx:        tx,
				Time:      tx.Time(),
				GasFeeCap: uint256.MustFromBig(tx.GasFeeCap()),
				GasTipCap: uint256.MustFromBig(tx.GasTipCap()),
			})
		}
		expectedCount += count
	}
	// Sort the transactions and cross check the nonce ordering
	txset := NewTransactionsByPriceAndNonce(signer, groups, baseFee)

	txs := types.Transactions{}
	for tx, _ := txset.Peek(); tx != nil; tx, _ = txset.Peek() {
		txs = append(txs, tx.Resolve())
		txset.Shift()
	}
	if len(txs) != expectedCount {
		t.Errorf("expected %d transactions, found %d", expectedCount, len(txs))
	}
	for i, txi := range txs {
		fromi, _ := types.Sender(signer, txi)

		// Make sure the nonce order is valid
		for j, txj := range txs[i+1:] {
			fromj, _ := types.Sender(signer, txj)
			if fromi == fromj && txi.Nonce() > txj.Nonce() {
				t.Errorf("invalid nonce ordering: tx #%d (A=%x N=%v) < tx #%d (A=%x N=%v)", i, fromi[:4], txi.Nonce(), i+j, fromj[:4], txj.Nonce())
			}
		}
		// If the next tx has different from account, the price must be lower than the current one
		if i+1 < len(txs) {
			next := txs[i+1]
			fromNext, _ := types.Sender(signer, next)
			tip, err := txi.EffectiveGasTip(baseFee)
			nextTip, nextErr := next.EffectiveGasTip(baseFee)
			if err != nil || nextErr != nil {
				t.Errorf("error calculating effective tip, err %v nextErr %v", err, nextErr)
			}
			if fromi != fromNext && tip.Cmp(nextTip) < 0 {
				t.Errorf("invalid gasprice ordering: tx #%d (A=%x P=%v) < tx #%d (A=%x P=%v)", i, fromi[:4], txi.GasPrice(), i+1, fromNext[:4], next.GasPrice())
			}
		}
	}
}

// Tests that if multiple transactions have the same price, the ones seen earlier
// are prioritized to avoid network spam attacks aiming for a specific ordering.
func TestTransactionTimeSort(t *testing.T) {
	// Generate a batch of accounts to start with
	keys := make([]*ecdsa.PrivateKey, 5)
	for i := 0; i < len(keys); i++ {
		keys[i], _ = crypto.GenerateKey()
	}
	signer := types.HomesteadSigner{}

	// Generate a batch of transactions with overlapping prices, but different creation times
	groups := map[common.Address][]*txpool.LazyTransaction{}
	for start, key := range keys {
		addr := crypto.PubkeyToAddress(key.PublicKey)

		tx, _ := types.SignTx(types.NewTransaction(0, common.Address{}, big.NewInt(100), 100, big.NewInt(1), nil), signer, key)
		tx.SetTime(time.Unix(0, int64(len(keys)-start)))

		groups[addr] = append(groups[addr], &txpool.LazyTransaction{
			Tx:        tx,
			Time:      tx.Time(),
			GasFeeCap: uint256.MustFromBig(tx.GasFeeCap()),
			GasTipCap: uint256.MustFromBig(tx.GasTipCap()),
		})
	}
	// Sort the transactions and cross check the nonce ordering
	txset := NewTransactionsByPriceAndNonce(signer, groups, nil)

	txs := types.Transactions{}
	for tx, _ := txset.Peek(); tx != nil; tx, _ = txset.Peek() {
		txs = append(txs, tx.Resolve())
		txset.Shift()
	}
	if len(txs) != len(keys) {
		t.Errorf("expected %d transactions, found %d", len(keys), len(txs))
	}
	for i, txi := range txs {
		fromi, _ := types.Sender(signer, txi)
		if i+1 < len(txs) {
			next := txs[i+1]
			fromNext, _ := types.Sender(signer, next)

			if txi.GasPrice().Cmp(next.GasPrice()) < 0 {
				t.Errorf("invalid gasprice ordering: tx #%d (A=%x P=%v) < tx #%d (A=%x P=%v)", i, fromi[:4], txi.GasPrice(), i+1, fromNext[:4], next.GasPrice())
			}
			// Make sure time order is ascending if the txs have the same gas price
			if txi.GasPrice().Cmp(next.GasPrice()) == 0 && txi.Time().After(next.Time()) {
				t.Errorf("invalid received time ordering: tx #%d (A=%x T=%v) > tx #%d (A=%x T=%v)", i, fromi[:4], txi.Time(), i+1, fromNext[:4], next.Time())
			}
		}
	}
}

func TestTransactionsByPriceAndNonceNext(t *testing.T) {
	keys := make([]*ecdsa.PrivateKey, 4)
	addresses := make([]common.Address, 4)
	for i := range keys {
		keys[i], _ = crypto.GenerateKey()
		addresses[i] = crypto.PubkeyToAddress(keys[i].PublicKey)
	}

	signer := types.NewMikoSigner(big.NewInt(2020))

	transactions := make([]*txpool.LazyTransaction, 4)
	for i := range transactions {
		tx, _ := types.SignTx(types.NewTransaction(0, common.Address{}, common.Big0, 21000, big.NewInt(int64(i)), nil), signer, keys[i])
		transactions[i] = &txpool.LazyTransaction{
			Tx:        tx,
			Hash:      tx.Hash(),
			Time:      tx.Time(),
			GasFeeCap: uint256.MustFromBig(tx.GasFeeCap()),
			GasTipCap: uint256.MustFromBig(tx.GasTipCap()),
			Gas:       tx.Gas(),
			BlobGas:   tx.BlobGas(),
		}
	}

	// There is no next transaction
	groups := make(map[common.Address][]*txpool.LazyTransaction)
	groups[addresses[0]] = append(groups[addresses[0]], transactions[0])

	txs := NewTransactionsByPriceAndNonce(signer, groups, nil)
	lazyTx, _ := txs.Next()
	if lazyTx != nil {
		t.Fatalf("Expect no next transaction, got %v", lazyTx)
	}

	// In heap, head transaction has gas price 1, child transaction has gas price 0
	groups = make(map[common.Address][]*txpool.LazyTransaction)
	groups[addresses[0]] = append(groups[addresses[0]], transactions[0])
	groups[addresses[1]] = append(groups[addresses[1]], transactions[1])
	txs = NewTransactionsByPriceAndNonce(signer, groups, nil)
	lazyTx, _ = txs.Next()
	if lazyTx == nil || lazyTx.GasFeeCap.ToBig().Cmp(common.Big0) != 0 {
		t.Fatalf("Expect to have next transaction of gas price 0, got: %v", lazyTx)
	}

	// In heap, head transaction has gas price 2, children have gas price 1, 0
	// Next transaction must have gas price 1
	groups = make(map[common.Address][]*txpool.LazyTransaction)
	groups[addresses[0]] = append(groups[addresses[0]], transactions[0])
	groups[addresses[1]] = append(groups[addresses[1]], transactions[1])
	groups[addresses[2]] = append(groups[addresses[2]], transactions[2])
	txs = NewTransactionsByPriceAndNonce(signer, groups, nil)
	lazyTx, _ = txs.Next()
	if lazyTx == nil || lazyTx.GasFeeCap.ToBig().Cmp(common.Big1) != 0 {
		t.Fatalf("Expect to have next transaction of gas price 1, got: %v", lazyTx)
	}

	// In heap, head transaction has gas price 3, children have gas price 2, 1
	// Next transaction must have gas price 2
	groups = make(map[common.Address][]*txpool.LazyTransaction)
	groups[addresses[0]] = append(groups[addresses[0]], transactions[0])
	groups[addresses[1]] = append(groups[addresses[1]], transactions[1])
	groups[addresses[2]] = append(groups[addresses[2]], transactions[2])
	groups[addresses[3]] = append(groups[addresses[3]], transactions[3])
	txs = NewTransactionsByPriceAndNonce(signer, groups, nil)
	lazyTx, _ = txs.Next()
	if lazyTx == nil || lazyTx.GasFeeCap.ToBig().Cmp(common.Big2) != 0 {
		t.Fatalf("Expect to have next transaction of gas price 2, got: %v", lazyTx)
	}

	// The next transaction of head transaction has higher price than children in heap
	nextTx, _ := types.SignTx(types.NewTransaction(0, common.Address{}, common.Big0, 21000, big.NewInt(100), nil), signer, keys[3])
	tx := &txpool.LazyTransaction{
		Tx:        nextTx,
		Hash:      nextTx.Hash(),
		Time:      nextTx.Time(),
		GasFeeCap: uint256.MustFromBig(nextTx.GasFeeCap()),
		GasTipCap: uint256.MustFromBig(nextTx.GasTipCap()),
		Gas:       nextTx.Gas(),
		BlobGas:   nextTx.BlobGas(),
	}

	groups = make(map[common.Address][]*txpool.LazyTransaction)
	groups[addresses[0]] = append(groups[addresses[0]], transactions[0])
	groups[addresses[1]] = append(groups[addresses[1]], transactions[1])
	groups[addresses[2]] = append(groups[addresses[2]], transactions[2])
	groups[addresses[3]] = append(groups[addresses[3]], transactions[3], tx)
	txs = NewTransactionsByPriceAndNonce(signer, groups, nil)
	lazyTx, _ = txs.Next()
	if lazyTx == nil || lazyTx.GasFeeCap.ToBig().Cmp(big.NewInt(100)) != 0 {
		t.Fatalf("Expect to have next transaction of gas price 100, got: %v", tx)
	}
}
