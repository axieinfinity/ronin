package vm

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type TestOpEvent struct {
}

func (tx *TestOpEvent) Publish(
	opcode OpCode,
	order, blockHeight uint64,
	blockHash common.Hash,
	blockTime uint64,
	hash common.Hash,
	from, to common.Address,
	value *big.Int,
	input, output []byte,
	err error,
) *types.InternalTransaction {
	return &types.InternalTransaction{
		Opcode:  opcode.String(),
		Type:    "test",
		Success: err == nil,
		Error:   "",
		InternalTransactionBody: &types.InternalTransactionBody{
			Order:           order,
			TransactionHash: hash,
			Value:           value,
			Input:           input,
			Output:          output,
			From:            from,
			To:              to,
			Height:          blockHeight,
			BlockHash:       blockHash,
			BlockTime:       blockTime,
		},
	}
}

func TestPublishEvents(t *testing.T) {
	ctx := BlockContext{
		PublishEvents: map[OpCode]OpEvent{
			CALL: &TestOpEvent{},
		},
		CurrentTransaction: types.NewTx(&types.LegacyTx{
			Nonce:    1,
			To:       nil,
			Value:    big.NewInt(0),
			Gas:      0,
			GasPrice: big.NewInt(0),
			Data:     []byte(""),
		}),
		BlockNumber:          common.Big0,
		Time:                 common.Big0,
		InternalTransactions: &[]*types.InternalTransaction{},
	}

	evm := &EVM{Context: ctx}
	evm.PublishEvent(CALL, 1, common.Address{}, common.Address{}, big.NewInt(0), []byte(""), []byte(""), nil)
	if len(*evm.Context.InternalTransactions) != 1 || (*evm.Context.InternalTransactions)[0].Type != "test" {
		t.Error("Failed to publish opcode event")
	}
}
