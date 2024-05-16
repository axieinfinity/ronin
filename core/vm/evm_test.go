package vm

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
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
		Output:  output,
		InternalTransactionBody: &types.InternalTransactionBody{
			Order:           order,
			TransactionHash: hash,
			Value:           value,
			Input:           input,
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
		Time:                 0,
		InternalTransactions: &[]*types.InternalTransaction{},
	}

	evm := &EVM{Context: ctx}
	evm.PublishEvent(CALL, 1, common.Address{}, common.Address{}, big.NewInt(0), []byte(""), []byte(""), nil)
	if len(*evm.Context.InternalTransactions) != 1 || (*evm.Context.InternalTransactions)[0].Type != "test" {
		t.Error("Failed to publish opcode event")
	}
}

type InternalTransactionEvent struct{}

func (tx *InternalTransactionEvent) Publish(
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
	var msgType string
	if opcode == CALL || opcode == DELEGATECALL {
		msgType = types.InternalTransactionContractCall
	} else {
		msgType = types.InternalTransactionContractCreation
	}

	internal := &types.InternalTransaction{
		Opcode:  opcode.String(),
		Type:    msgType,
		Success: err == nil,
		Error:   "",
		Output:  output,
		InternalTransactionBody: &types.InternalTransactionBody{
			Order:           order,
			TransactionHash: hash,
			Value:           value,
			Input:           input,
			From:            from,
			To:              to,
			Height:          blockHeight,
			BlockHash:       blockHash,
			BlockTime:       blockTime,
		},
	}
	if err != nil {
		internal.Error = err.Error()
	}
	return internal
}

func TestInternalTransactionOrder(t *testing.T) {
	statedb, err := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	if err != nil {
		t.Fatal(err)
	}

	publishEvent := PublishEvent{
		OpCodes: []OpCode{
			CALL,
			DELEGATECALL,
			CREATE,
			CREATE2,
		},
		Event: &InternalTransactionEvent{},
	}

	internalTransactions := make([]*types.InternalTransaction, 0)
	evm := NewEVM(
		BlockContext{
			BlockNumber:          common.Big0,
			InternalTransactions: &internalTransactions,
			Transfer:             func(_ StateDB, _, _ common.Address, _ *big.Int) {},
			PublishEvents:        make(PublishEventsMap),
			CurrentTransaction:   types.NewTx(&types.LegacyTx{}),
		},
		TxContext{},
		statedb,
		&params.ChainConfig{
			IstanbulBlock: common.Big0,
		},
		Config{},
	)

	for _, opcode := range publishEvent.OpCodes {
		evm.Context.PublishEvents[opcode] = publishEvent.Event
	}

	/*
		pragma solidity ^0.8.18;

		contract Test1 {
			fallback() external {
				address test2 = address(0x202);
				test2.call{gas: 200000}("");
			}
		}
	*/
	test1Contract := common.Hex2Bytes("608060405234801561001057600080fd5b5060f58061001f6000396000f3fe6080604052348015600f57600080fd5b50600061020290508073ffffffffffffffffffffffffffffffffffffffff1662030d40604051603c9060ac565b60006040518083038160008787f1925050503d80600081146078576040519150601f19603f3d011682016040523d82523d6000602084013e607d565b606091505b005b600081905092915050565b50565b60006098600083607f565b915060a182608a565b600082019050919050565b600060b582608d565b915081905091905056fea26469706673582212208ba7f80d6a96ebc77318dbf151ea79ff647cf772581f7c4acb3fd1d2babfdfb464736f6c63430008120033")
	test1Address := common.BigToAddress(big.NewInt(0x201))

	/*
		pragma solidity ^0.8.18;

		contract Test2 {
			fallback() external {
				address test3 = address(0x203);
				test3.call{gas: 100000}("");
			}
		}
	*/

	test2Contract := common.Hex2Bytes("608060405234801561001057600080fd5b5060f58061001f6000396000f3fe6080604052348015600f57600080fd5b50600061020390508073ffffffffffffffffffffffffffffffffffffffff16620186a0604051603c9060ac565b60006040518083038160008787f1925050503d80600081146078576040519150601f19603f3d011682016040523d82523d6000602084013e607d565b606091505b005b600081905092915050565b50565b60006098600083607f565b915060a182608a565b600082019050919050565b600060b582608d565b915081905091905056fea2646970667358221220dc64986e2e2b84dfc27218360968609146254d5606999d4de5187b18fc1ca7a664736f6c63430008120033")
	test2Address := common.BigToAddress(big.NewInt(0x202))

	/*
		pragma solidity ^0.8.18;

		contract Test3 {
			fallback() external {
			}
		}
	*/
	test3Contract := common.Hex2Bytes("6080604052348015600f57600080fd5b50604780601d6000396000f3fe6080604052348015600f57600080fd5b00fea2646970667358221220e499609e32f510161ba4b0c5e900173a55b3457a85347d20c83099f6d20b2ace64736f6c63430008120033")
	test3Address := common.BigToAddress(big.NewInt(0x203))

	statedb.SetCode(test1Address, test1Contract)
	statedb.SetCode(test2Address, test2Contract)
	statedb.SetCode(test3Address, test3Contract)

	deployedTest1, _, err := evm.Call(AccountRef(test1Address), test1Address, []byte{}, 100_000, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}

	deployedTest2, _, err := evm.Call(AccountRef(test2Address), test2Address, []byte{}, 100_000, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}

	deployedTest3, _, err := evm.Call(AccountRef(test3Address), test3Address, []byte{}, 100_000, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}

	statedb.SetCode(test1Address, deployedTest1)
	statedb.SetCode(test2Address, deployedTest2)
	statedb.SetCode(test3Address, deployedTest3)

	_, _, err = evm.Call(AccountRef(test1Address), test1Address, []byte{}, 1_000_000, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}

	internalTxs := *evm.Context.InternalTransactions
	if len(internalTxs) != 2 {
		t.Fatalf("Internal transactions length mismatches, got %d expect %d", len(*evm.Context.InternalTransactions), 2)
	}

	if internalTxs[0].From != test1Address || internalTxs[0].To != test2Address {
		t.Fatalf("Unexpected internal transaction #0, %+v, body: %+v", internalTxs[0], internalTxs[0].InternalTransactionBody)
	}

	if internalTxs[1].From != test2Address || internalTxs[1].To != test3Address {
		t.Fatalf("Unexpected internal transaction #1, %+v, body: %+v", internalTxs[1], internalTxs[1].InternalTransactionBody)
	}
}
