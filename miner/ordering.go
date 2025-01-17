package miner

import (
	"container/heap"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/types"
)

// TxWithMinerFee wraps a transaction with its gas price or effective miner gasTipCap
type TxWithMinerFee struct {
	from     common.Address
	tx       *txpool.LazyTransaction
	minerFee *big.Int
}

// newTxWithMinerFee creates a wrapped transaction, calculating the effective
// miner gasTipCap if a base fee is provided.
// Returns error in case of a negative effective miner gasTipCap.
func newTxWithMinerFee(tx *txpool.LazyTransaction, from common.Address, baseFee *big.Int) (*TxWithMinerFee, error) {
	var (
		minerFee *big.Int
		tipCap   = tx.GasTipCap.ToBig()
		feeCap   = tx.GasFeeCap.ToBig()
	)
	if baseFee == nil {
		minerFee = tipCap
	} else {
		if tx.GasFeeCap.ToBig().Cmp(baseFee) < 0 {
			return nil, types.ErrGasFeeCapTooLow
		}
		minerFee = math.BigMin(tipCap, new(big.Int).Sub(feeCap, baseFee))
	}
	return &TxWithMinerFee{
		from:     from,
		tx:       tx,
		minerFee: minerFee,
	}, nil
}

// TxByPriceAndTime implements both the sort and the heap interface, making it useful
// for all at once sorting as well as individually adding and removing elements.
type TxByPriceAndTime []*TxWithMinerFee

func (s TxByPriceAndTime) Len() int { return len(s) }
func (s TxByPriceAndTime) Less(i, j int) bool {
	return cmpPriceAndTime(s[i], s[j])
}
func (s TxByPriceAndTime) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s *TxByPriceAndTime) Push(x interface{}) {
	*s = append(*s, x.(*TxWithMinerFee))
}

func (s *TxByPriceAndTime) Pop() interface{} {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[0 : n-1]
	return x
}

// cmpPriceAndTime compares 2 transactions by their miner fee and
// time first seen in txpool.
// Returns true if `a` has higher miner fee or appears in txpool
// before `b`.
func cmpPriceAndTime(a *TxWithMinerFee, b *TxWithMinerFee) bool {
	// If the prices are equal, use the time the transaction was first seen for
	// deterministic sorting
	cmp := a.minerFee.Cmp(b.minerFee)
	if cmp == 0 {
		return a.tx.Time.Before(b.tx.Time)
	}
	return cmp > 0
}

// TransactionsByPriceAndNonce represents a set of transactions that can return
// transactions in a profit-maximizing sorted order, while supporting removing
// entire batches of transactions for non-executable accounts.
type TransactionsByPriceAndNonce struct {
	txs     map[common.Address][]*txpool.LazyTransaction // Per account nonce-sorted list of transactions
	heads   TxByPriceAndTime                             // Next transaction for each unique account (price heap)
	signer  types.Signer                                 // Signer for the set of transactions
	baseFee *big.Int                                     // Current base fee
}

// NewTransactionsByPriceAndNonce creates a transaction set that can retrieve
// price sorted transactions in a nonce-honouring way.
//
// Note, the input map is reowned so the caller should not interact any more with
// if after providing it to the constructor.
func NewTransactionsByPriceAndNonce(signer types.Signer, txs map[common.Address][]*txpool.LazyTransaction, baseFee *big.Int) *TransactionsByPriceAndNonce {
	// Initialize a price and received time based heap with the head transactions
	heads := make(TxByPriceAndTime, 0, len(txs))
	for from, accTxs := range txs {
		wrapped, err := newTxWithMinerFee(accTxs[0], from, baseFee)
		if err != nil {
			delete(txs, from)
			continue
		}
		heads = append(heads, wrapped)
		txs[from] = accTxs[1:]
	}
	heap.Init(&heads)

	// Assemble and return the transaction set
	return &TransactionsByPriceAndNonce{
		txs:     txs,
		heads:   heads,
		signer:  signer,
		baseFee: baseFee,
	}
}

// Peek returns the next transaction by price and the miner fee.
func (t *TransactionsByPriceAndNonce) Peek() (*txpool.LazyTransaction, *big.Int) {
	if len(t.heads) == 0 {
		return nil, nil
	}
	return t.heads[0].tx, t.heads[0].minerFee
}

// Shift replaces the current best head with the next one from the same account.
func (t *TransactionsByPriceAndNonce) Shift() {
	acc := t.heads[0].from
	if txs, ok := t.txs[acc]; ok && len(txs) > 0 {
		if wrapped, err := newTxWithMinerFee(txs[0], acc, t.baseFee); err == nil {
			t.heads[0], t.txs[acc] = wrapped, txs[1:]
			heap.Fix(&t.heads, 0)
			return
		}
	}
	heap.Pop(&t.heads)
}

// Pop removes the best transaction, *not* replacing it with the next one from
// the same account. This should be used when a transaction cannot be executed
// and hence all subsequent ones should be discarded from the same account.
func (t *TransactionsByPriceAndNonce) Pop() {
	heap.Pop(&t.heads)
}

func (t *TransactionsByPriceAndNonce) Size() int {
	return t.heads.Len()
}

func (t *TransactionsByPriceAndNonce) Clear() {
	t.heads = TxByPriceAndTime{}
}

// Next return the potential next committed transaction so that we can speculative
// execute that transaction. As we don't know the result of current transaction yet,
// we don't know that the next operation is shift or pop. This function returns the
// largest among right, left children and the next transaction from the account at
// the head node.
func (t *TransactionsByPriceAndNonce) Next() (*txpool.LazyTransaction, *big.Int) {
	heapSize := len(t.heads)
	acc := t.heads[0].from

	var candidateTx *TxWithMinerFee
	if txs, ok := t.txs[acc]; ok && len(txs) > 0 {
		if tx, err := newTxWithMinerFee(txs[0], acc, t.baseFee); err == nil {
			candidateTx = tx
		}
	}

	if heapSize >= 2 {
		// left child
		if candidateTx == nil || cmpPriceAndTime(t.heads[1], candidateTx) {
			candidateTx = t.heads[1]
		}

		if heapSize >= 3 {
			// right child
			if cmpPriceAndTime(t.heads[2], candidateTx) {
				candidateTx = t.heads[2]
			}
		}
	}

	if candidateTx != nil {
		return candidateTx.tx, candidateTx.minerFee
	} else {
		return nil, nil
	}
}
