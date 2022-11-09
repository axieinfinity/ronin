package vm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"math/big"
	"testing"
)

type TestOpEvent struct {
	feed *event.Feed
}

func (tx *TestOpEvent) Publish(opcode OpCode, order uint64, stateDB StateDB, blockHeight uint64,
	blockHash common.Hash, blockTime uint64, hash common.Hash, from, to common.Address, value *big.Int, input []byte, err error) error {
	tx.feed.Send(true)
	return nil
}

func TestPublishEvents(t *testing.T) {
	var (
		internalTxFeed event.Feed
		scope          event.SubscriptionScope
		rs             bool
	)

	ch := make(chan bool, 1)
	scope.Track(internalTxFeed.Subscribe(ch))

	ctx := BlockContext{
		PublishEvents: map[OpCode]OpEvent{
			CALL: &TestOpEvent{feed: &internalTxFeed},
		},
		CurrentTransaction: types.NewTx(&types.LegacyTx{
			Nonce:    1,
			To:       nil,
			Value:    big.NewInt(0),
			Gas:      0,
			GasPrice: big.NewInt(0),
			Data:     []byte(""),
		}),
	}

	evm := &EVM{Context: ctx}
	evm.PublishEvent(CALL, 1, common.Address{}, common.Address{}, big.NewInt(0), []byte(""), nil)
	select {
	case rs = <-ch:
		if !rs {
			t.Fatal("Publish Event failed")
		}
	}
}
