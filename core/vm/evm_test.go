package vm

import (
	"math/big"
	"math/rand"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
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

func TestInternalTransactionOrderMultipleCalls(t *testing.T) {
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
		contract Test1{
			fallback() external {
				address test2 = address(0x202);
				address test3 = address(0x203);
				address test4 = address(0x204);
				address test5 = address(0x205);
				test2.call{gas: 200000}("");
				test3.call{gas: 200000}("");
				test4.call{gas: 200000}("");
				test5.call{gas: 200000}("");
			}
		}
	*/
	test1Contract := common.Hex2Bytes("6080604052348015600f57600080fd5b5061025c8061001f6000396000f3fe608060405234801561001057600080fd5b50600061020290506000610203905060006102049050600061020590508373ffffffffffffffffffffffffffffffffffffffff1662030d4060405161005490610211565b60006040518083038160008787f1925050503d8060008114610092576040519150601f19603f3d011682016040523d82523d6000602084013e610097565b606091505b5050508273ffffffffffffffffffffffffffffffffffffffff1662030d406040516100c190610211565b60006040518083038160008787f1925050503d80600081146100ff576040519150601f19603f3d011682016040523d82523d6000602084013e610104565b606091505b5050508173ffffffffffffffffffffffffffffffffffffffff1662030d4060405161012e90610211565b60006040518083038160008787f1925050503d806000811461016c576040519150601f19603f3d011682016040523d82523d6000602084013e610171565b606091505b5050508073ffffffffffffffffffffffffffffffffffffffff1662030d4060405161019b90610211565b60006040518083038160008787f1925050503d80600081146101d9576040519150601f19603f3d011682016040523d82523d6000602084013e6101de565b606091505b005b600081905092915050565b50565b60006101fb6000836101e0565b9150610206826101eb565b600082019050919050565b600061021c826101ee565b915081905091905056fea26469706673582212203366ca55156e44ab6b493b487c644cc8d5399a6979db8d103b81a6c51a11fb0964736f6c63430008190033")
	test1Address := common.BigToAddress(big.NewInt(0x201))

	/*
		pragma solidity ^0.8.18;
		contract Test2{
			fallback() external {

			}
		}
	*/
	test2Contract := common.Hex2Bytes("6080604052348015600f57600080fd5b50604780601d6000396000f3fe6080604052348015600f57600080fd5b00fea2646970667358221220e616bd0cc3f68dfe8757db1aedc6efacbfb3abe75e9f1b63101cffaa05945aa564736f6c63430008190033")
	test2Address := common.BigToAddress(big.NewInt(0x202))

	/*
		pragma solidity ^0.8.18;
		contract Test3{
			fallback() external {

			}
		}
	*/
	test3Contract := common.Hex2Bytes("6080604052348015600f57600080fd5b50604780601d6000396000f3fe6080604052348015600f57600080fd5b00fea264697066735822122059d526ccd076eb695b6fe0924ce842d0a615674b72bf487fdbedd44164e8e82864736f6c63430008190033")
	test3Address := common.BigToAddress(big.NewInt(0x203))

	/*
		pragma solidity ^0.8.18;
		contract Test4{
			fallback() external {

			}
		}
	*/
	test4Contract := common.Hex2Bytes("6080604052348015600f57600080fd5b50604780601d6000396000f3fe6080604052348015600f57600080fd5b00fea2646970667358221220434e4ebedd1ee41ba291b229f7af765a1d7e7b3a986343899c1396f5cbb4804964736f6c63430008190033")
	test4Address := common.BigToAddress(big.NewInt(0x204))

	/*
		pragma solidity ^0.8.18;
		contract Test5{
			fallback() external {

			}
		}
	*/
	test5Contract := common.Hex2Bytes("6080604052348015600f57600080fd5b50604780601d6000396000f3fe6080604052348015600f57600080fd5b00fea2646970667358221220d8555b379f58cf3ef2c11bdd95a5e4e6b9924243fa2bb2907ab0ae9cee8545f264736f6c63430008190033")
	test5Address := common.BigToAddress(big.NewInt(0x205))

	statedb.SetCode(test1Address, test1Contract)
	statedb.SetCode(test2Address, test2Contract)
	statedb.SetCode(test3Address, test3Contract)
	statedb.SetCode(test4Address, test4Contract)
	statedb.SetCode(test5Address, test5Contract)

	deployedTest1, _, err := evm.Call(AccountRef(test1Address), test1Address, []byte{}, 1_000_000, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	deployedTest2, _, err := evm.Call(AccountRef(test2Address), test2Address, []byte{}, 1_000_000, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	deployedTest3, _, err := evm.Call(AccountRef(test3Address), test3Address, []byte{}, 1_000_000, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	deployedTest4, _, err := evm.Call(AccountRef(test4Address), test4Address, []byte{}, 1_000_000, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	deployedTest5, _, err := evm.Call(AccountRef(test5Address), test5Address, []byte{}, 1_000_000, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	statedb.SetCode(test1Address, deployedTest1)
	statedb.SetCode(test2Address, deployedTest2)
	statedb.SetCode(test3Address, deployedTest3)
	statedb.SetCode(test4Address, deployedTest4)
	statedb.SetCode(test5Address, deployedTest5)
	_, _, err = evm.Call(AccountRef(test1Address), test1Address, []byte{}, 1_000_000, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	/*
		Calls tree:
			Test1 -> Test2
			Test1 -> Test3
			Test1 -> Test4
			Test1 -> Test5
	*/
	var expectedTransactions []struct {
		from, to common.Address
	}
	expectedTransactions = append(expectedTransactions, struct {
		from, to common.Address
	}{test1Address, test2Address})
	expectedTransactions = append(expectedTransactions, struct {
		from, to common.Address
	}{test1Address, test3Address})
	expectedTransactions = append(expectedTransactions, struct {
		from, to common.Address
	}{test1Address, test4Address})
	expectedTransactions = append(expectedTransactions, struct {
		from, to common.Address
	}{test1Address, test5Address})
	internalTxs := *evm.Context.InternalTransactions
	if len(internalTxs) != len(expectedTransactions) {
		t.Fatalf("Internal transactions length mismatches, got %d expect %d", len(*evm.Context.InternalTransactions), len(expectedTransactions))
	}
	for i, tx := range internalTxs {
		if tx.From != expectedTransactions[i].from || tx.To != expectedTransactions[i].to {
			t.Fatalf("Unexpected internal transaction #%d, %+v, body: %+v", i, tx, tx.InternalTransactionBody)
		}
	}
}

func TestInternalTransactionOrderComplexCalls(t *testing.T) {
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
		contract Test1{
			fallback() external {
				address test2 = address(0x202);
				address test5 = address(0x205);
				address test4 = address(0x204);
				test2.call{gas: 500000}("");
				test5.call{gas: 100000}("");
				test4.call{gas: 300000}("");
			}
		}
	*/
	test1Contract := common.Hex2Bytes("6080604052348015600f57600080fd5b506101e88061001f6000396000f3fe608060405234801561001057600080fd5b506000610202905060006102059050600061020490508273ffffffffffffffffffffffffffffffffffffffff166207a12060405161004d9061019d565b60006040518083038160008787f1925050503d806000811461008b576040519150601f19603f3d011682016040523d82523d6000602084013e610090565b606091505b5050508173ffffffffffffffffffffffffffffffffffffffff16620186a06040516100ba9061019d565b60006040518083038160008787f1925050503d80600081146100f8576040519150601f19603f3d011682016040523d82523d6000602084013e6100fd565b606091505b5050508073ffffffffffffffffffffffffffffffffffffffff16620493e06040516101279061019d565b60006040518083038160008787f1925050503d8060008114610165576040519150601f19603f3d011682016040523d82523d6000602084013e61016a565b606091505b005b600081905092915050565b50565b600061018760008361016c565b915061019282610177565b600082019050919050565b60006101a88261017a565b915081905091905056fea264697066735822122066247615e92ffcf7dbd2b40c7c943709ba6e0d50701c0604f6fa5e4aeff736d764736f6c63430008190033")
	test1Address := common.BigToAddress(big.NewInt(0x201))

	/*
		pragma solidity ^0.8.18;
		contract Test2{
			fallback() external {
				address test3 = address(0x203);
				address test4 = address(0x204);
				test3.call{gas: 100000}("");
				test4.call{gas: 300000}("");
			}
		}
	*/
	test2Contract := common.Hex2Bytes("6080604052348015600f57600080fd5b506101748061001f6000396000f3fe608060405234801561001057600080fd5b5060006102039050600061020490508173ffffffffffffffffffffffffffffffffffffffff16620186a060405161004690610129565b60006040518083038160008787f1925050503d8060008114610084576040519150601f19603f3d011682016040523d82523d6000602084013e610089565b606091505b5050508073ffffffffffffffffffffffffffffffffffffffff16620493e06040516100b390610129565b60006040518083038160008787f1925050503d80600081146100f1576040519150601f19603f3d011682016040523d82523d6000602084013e6100f6565b606091505b005b600081905092915050565b50565b60006101136000836100f8565b915061011e82610103565b600082019050919050565b600061013482610106565b915081905091905056fea264697066735822122083ac0cb0ba6a17a50b958c370f1edad081cce6d5983e681e45e85e4bf6b2959464736f6c63430008190033")
	test2Address := common.BigToAddress(big.NewInt(0x202))

	/*
		pragma solidity ^0.8.18;
		contract Test3{
			fallback() external {

			}
		}
	*/
	test3Contract := common.Hex2Bytes("6080604052348015600f57600080fd5b50604780601d6000396000f3fe6080604052348015600f57600080fd5b00fea264697066735822122059d526ccd076eb695b6fe0924ce842d0a615674b72bf487fdbedd44164e8e82864736f6c63430008190033")
	test3Address := common.BigToAddress(big.NewInt(0x203))

	/*
		pragma solidity ^0.8.18;
		contract Test4{
			fallback() external {
				address test3 = address(0x203);
				address test5 = address(0x205);
				test3.call{gas: 100000}("");
				test5.call{gas: 100000}("");
			}
		}
	*/
	test4Contract := common.Hex2Bytes("6080604052348015600f57600080fd5b506101748061001f6000396000f3fe608060405234801561001057600080fd5b5060006102039050600061020590508173ffffffffffffffffffffffffffffffffffffffff16620186a060405161004690610129565b60006040518083038160008787f1925050503d8060008114610084576040519150601f19603f3d011682016040523d82523d6000602084013e610089565b606091505b5050508073ffffffffffffffffffffffffffffffffffffffff16620186a06040516100b390610129565b60006040518083038160008787f1925050503d80600081146100f1576040519150601f19603f3d011682016040523d82523d6000602084013e6100f6565b606091505b005b600081905092915050565b50565b60006101136000836100f8565b915061011e82610103565b600082019050919050565b600061013482610106565b915081905091905056fea26469706673582212202f9f7beae47464e7507130b103b1996029e44ded5fa09102bde9f62bbe7a363164736f6c63430008190033")
	test4Address := common.BigToAddress(big.NewInt(0x204))

	/*
		pragma solidity ^0.8.18;
		contract Test5{
			fallback() external {

			}
		}
	*/
	test5Contract := common.Hex2Bytes("6080604052348015600f57600080fd5b50604780601d6000396000f3fe6080604052348015600f57600080fd5b00fea2646970667358221220d8555b379f58cf3ef2c11bdd95a5e4e6b9924243fa2bb2907ab0ae9cee8545f264736f6c63430008190033")
	test5Address := common.BigToAddress(big.NewInt(0x205))

	statedb.SetCode(test1Address, test1Contract)
	statedb.SetCode(test2Address, test2Contract)
	statedb.SetCode(test3Address, test3Contract)
	statedb.SetCode(test4Address, test4Contract)
	statedb.SetCode(test5Address, test5Contract)

	deployedTest1, _, err := evm.Call(AccountRef(test1Address), test1Address, []byte{}, 2_000_000, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	deployedTest2, _, err := evm.Call(AccountRef(test2Address), test2Address, []byte{}, 1_000_000, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	deployedTest3, _, err := evm.Call(AccountRef(test3Address), test3Address, []byte{}, 1_000_000, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	deployedTest4, _, err := evm.Call(AccountRef(test4Address), test4Address, []byte{}, 1_000_000, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	deployedTest5, _, err := evm.Call(AccountRef(test5Address), test5Address, []byte{}, 1_000_000, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	statedb.SetCode(test1Address, deployedTest1)
	statedb.SetCode(test2Address, deployedTest2)
	statedb.SetCode(test3Address, deployedTest3)
	statedb.SetCode(test4Address, deployedTest4)
	statedb.SetCode(test5Address, deployedTest5)
	_, _, err = evm.Call(AccountRef(test1Address), test1Address, []byte{}, 1_000_000, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	/*
		Calls tree:
			Test1 -> Test2
						Test2 -> Test3
						Test2 -> Test4
								Test4 -> Test3
								Test4 -> Test5
			Test1 -> Test5
			Test1 -> Test4
						Test4 -> Test3
						Test4 -> Test5
	*/
	var expectedTransactions []struct {
		from, to common.Address
	}
	expectedTransactions = append(expectedTransactions, struct {
		from, to common.Address
	}{test1Address, test2Address})
	expectedTransactions = append(expectedTransactions, struct {
		from, to common.Address
	}{test2Address, test3Address})
	expectedTransactions = append(expectedTransactions, struct {
		from, to common.Address
	}{test2Address, test4Address})
	expectedTransactions = append(expectedTransactions, struct {
		from, to common.Address
	}{test4Address, test3Address})
	expectedTransactions = append(expectedTransactions, struct {
		from, to common.Address
	}{test4Address, test5Address})
	expectedTransactions = append(expectedTransactions, struct {
		from, to common.Address
	}{test1Address, test5Address})
	expectedTransactions = append(expectedTransactions, struct {
		from, to common.Address
	}{test1Address, test4Address})
	expectedTransactions = append(expectedTransactions, struct {
		from, to common.Address
	}{test4Address, test3Address})
	expectedTransactions = append(expectedTransactions, struct {
		from, to common.Address
	}{test4Address, test5Address})

	internalTxs := *evm.Context.InternalTransactions
	if len(internalTxs) != len(expectedTransactions) {
		t.Fatalf("Internal transactions length mismatches, got %d expect %d", len(*evm.Context.InternalTransactions), len(expectedTransactions))
	}
	for i, tx := range internalTxs {
		if tx.From != expectedTransactions[i].from || tx.To != expectedTransactions[i].to {
			t.Fatalf("Unexpected internal transaction #%d, %+v, body: %+v", i, tx, tx.InternalTransactionBody)
		}
	}
}

func TestInternalTransactionOrderDelegateCalls(t *testing.T) {
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
		contract Test1{
			fallback() external {
				address test2 = address(0x202);
				address test3 = address(0x203);
				test2.delegatecall{gas: 300000}("");
				test3.delegatecall{gas: 100000}("");
			}
		}
	*/
	test1Contract := common.Hex2Bytes("6080604052348015600f57600080fd5b506101708061001f6000396000f3fe608060405234801561001057600080fd5b5060006102029050600061020390508173ffffffffffffffffffffffffffffffffffffffff16620493e060405161004690610125565b6000604051808303818686f4925050503d8060008114610082576040519150601f19603f3d011682016040523d82523d6000602084013e610087565b606091505b5050508073ffffffffffffffffffffffffffffffffffffffff16620186a06040516100b190610125565b6000604051808303818686f4925050503d80600081146100ed576040519150601f19603f3d011682016040523d82523d6000602084013e6100f2565b606091505b005b600081905092915050565b50565b600061010f6000836100f4565b915061011a826100ff565b600082019050919050565b600061013082610102565b915081905091905056fea26469706673582212203c4f1657e18d5425a246604c111a278e3ce104da92558848e0fb2e1f4e8427cf64736f6c63430008190033")
	test1Address := common.BigToAddress(big.NewInt(0x201))
	/*
		pragma solidity ^0.8.18;
		contract Test2{
			fallback() external {
				address test3 = address(0x203);
				address test4 = address(0x204);
				test3.call{gas: 100000}("");
				test4.call{gas: 100000}("");
			}
		}
	*/
	test2Contract := common.Hex2Bytes("6080604052348015600f57600080fd5b506101748061001f6000396000f3fe608060405234801561001057600080fd5b5060006102039050600061020490508173ffffffffffffffffffffffffffffffffffffffff16620186a060405161004690610129565b60006040518083038160008787f1925050503d8060008114610084576040519150601f19603f3d011682016040523d82523d6000602084013e610089565b606091505b5050508073ffffffffffffffffffffffffffffffffffffffff16620186a06040516100b390610129565b60006040518083038160008787f1925050503d80600081146100f1576040519150601f19603f3d011682016040523d82523d6000602084013e6100f6565b606091505b005b600081905092915050565b50565b60006101136000836100f8565b915061011e82610103565b600082019050919050565b600061013482610106565b915081905091905056fea26469706673582212204274bcacfee9e5be08621b8e035937748edea34c5d5ab44e57715b13e2e8522f64736f6c63430008190033")
	test2Address := common.BigToAddress(big.NewInt(0x202))

	/*
		pragma solidity ^0.8.18;
		contract Test3{
			fallback() external {

			}
		}
	*/
	test3Contract := common.Hex2Bytes("6080604052348015600f57600080fd5b50604780601d6000396000f3fe6080604052348015600f57600080fd5b00fea264697066735822122059d526ccd076eb695b6fe0924ce842d0a615674b72bf487fdbedd44164e8e82864736f6c63430008190033")
	test3Address := common.BigToAddress(big.NewInt(0x203))

	/*
		pragma solidity ^0.8.18;
		contract Test4{
			fallback() external {

			}
		}
	*/
	test4Contract := common.Hex2Bytes("6080604052348015600f57600080fd5b50604780601d6000396000f3fe6080604052348015600f57600080fd5b00fea2646970667358221220e100e37006806f03bf4d8b68c6d5a3e892cbb17d9cdb52a7403d71beaafce27864736f6c63430008190033")
	test4Address := common.BigToAddress(big.NewInt(0x204))

	statedb.SetCode(test1Address, test1Contract)
	statedb.SetCode(test2Address, test2Contract)
	statedb.SetCode(test3Address, test3Contract)
	statedb.SetCode(test4Address, test4Contract)

	deployedTest1, _, err := evm.Call(AccountRef(test1Address), test1Address, []byte{}, 1_000_000, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	deployedTest2, _, err := evm.Call(AccountRef(test2Address), test2Address, []byte{}, 1_000_000, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	deployedTest3, _, err := evm.Call(AccountRef(test3Address), test3Address, []byte{}, 1_000_000, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	deployedTest4, _, err := evm.Call(AccountRef(test4Address), test4Address, []byte{}, 1_000_000, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}

	statedb.SetCode(test1Address, deployedTest1)
	statedb.SetCode(test2Address, deployedTest2)
	statedb.SetCode(test3Address, deployedTest3)
	statedb.SetCode(test4Address, deployedTest4)

	_, _, err = evm.Call(AccountRef(test1Address), test1Address, []byte{}, 1_000_000, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	/*
		Calls tree:
			Test1 -> Test2
			Test1(-> Test2)-> Test3
			Test1(-> Test2)-> Test4
			Test1 -> Test3
	*/
	var expectedTransactions []struct {
		from, to common.Address
	}
	expectedTransactions = append(expectedTransactions, struct {
		from, to common.Address
	}{test1Address, test2Address})
	expectedTransactions = append(expectedTransactions, struct {
		from, to common.Address
	}{test1Address, test3Address})
	expectedTransactions = append(expectedTransactions, struct {
		from, to common.Address
	}{test1Address, test4Address})
	expectedTransactions = append(expectedTransactions, struct {
		from, to common.Address
	}{test1Address, test3Address})

	internalTxs := *evm.Context.InternalTransactions
	if len(internalTxs) != len(expectedTransactions) {
		t.Fatalf("Internal transactions length mismatches, got %d expect %d", len(*evm.Context.InternalTransactions), len(expectedTransactions))
	}
	for i, tx := range internalTxs {
		if tx.From != expectedTransactions[i].from || tx.To != expectedTransactions[i].to {
			t.Fatalf("Unexpected internal transaction #%d, %+v, body: %+v", i, tx, tx.InternalTransactionBody)
		}
	}
}

func BenchmarkRBTreeContract(b *testing.B) {
	// Create random array of numbers and options
	// The "random" values are generated using a fixed seed for accurate comparison
	n := 10000
	bound := 1000
	var numbers []*big.Int
	var options []*big.Int
	var args []interface{}
	for i := 0; i < n; i++ {
		rng := rand.New(rand.NewSource(int64(i)))
		numbers = append(numbers, big.NewInt(int64(rng.Intn(bound)) + 1))
		options = append(options, big.NewInt(int64(rng.Intn(2)+1)))
	}
	args = append(args, numbers)
	args = append(args, options)

	// Pack data for ABI
    parsedABI, err := abi.JSON(strings.NewReader(
	`
	[
		{
			"inputs": [
				{
					"internalType": "uint256[]",
					"name": "numbers",
					"type": "uint256[]"
				},
				{
					"internalType": "uint256[]",
					"name": "options",
					"type": "uint256[]"
				}
			],
			"name": "process",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		}
	]
	`,
	))
	if err != nil {
        b.Fatalf("Failed to parse ABI: %v", err)
    }
    data, err := parsedABI.Pack("process", args...)
    if err != nil {
        b.Fatalf("Failed to pack data for ABI: %v", err)
    }

	// Setup EVM to run contract
	/*
		pragma solidity ^0.8.18;
		...
		// RBTree library - https://github.com/Vectorized/solady/blob/29d61c504425519c6deddc3e12c2e039ad43e8e3/test/RedBlackTree.t.sol#L4
		...
		contract OpcodeTest {
			using RedBlackTreeLib for *;
			RedBlackTreeLib.Tree tree;

			function process(uint256[] memory numbers, uint256[] memory options) external {
				for (uint256 i = 0; i < numbers.length; i++) {
					uint256 option = options[i];
					uint256 number = numbers[i];
					if (option == 1) {
						if (!tree.exists(number)) {
							tree.insert(number);
						}
					} else if (option == 2) {
						if (tree.exists(number)) {
							tree.remove(number);
						}
					}
				}
			}
		}
	*/
	contractString :=
		"608060405234801561001057600080fd5b506004361061002b5760003560e01c80636e476ea014610030575b600080fd5b61004361003e366004610987565b610045565b005b60005b82518110156100e1576000828281518110610065576100656109f0565b602002602001015190506000848381518110610083576100836109f0565b60200260200101519050816001036100b4576100a06000826100e6565b6100af576100af6000826100ff565b6100d7565b816002036100d7576100c76000826100e6565b156100d7576100d760008261011c565b5050600101610048565b505050565b6000806100f38484610128565b15159695505050505050565b600061010b83836101b0565b905080156100e1576100e1816101de565b600061010b83836101e8565b600080808361013e5761013e63c94f18776101de565b846020526801dc27bb5462fdadcb600052604060002060201b9250601f600152825460801c5b80156101a857809250808417548060601c806101865750848217638000000017545b8681036101975784935050506101a8565b8611511c637fffffff169050610164565b509250925092565b6000806000806101c08686610128565b9250925092506101d4838383886000610215565b9695505050505050565b806000526004601cfd5b60008060006101f78585610128565b925050915061020c8260008360006001610215565b95945050505050565b6000610896565b8082175480851c637fffffff16603e82901c637fffffff168382175480871c637fffffff16801561025f578086178054637fffffff603e1b1916603e89901b1790555b8261026d57836000526102a5565b8583178054808a1c637fffffff16890361029557637fffffff8a1b1916858a1b1790556102a5565b637fffffff8b1b1916858b1b1790555b637fffffff603e1b19637fffffff808b1b19969096169190991b178816603e84811b919091178787175593871b199716921b91909117949094169190921b17911755565b600083156102fc575063bb33e6ac6104b4565b600160205160801c01637fffffff81111561031e5763ed732d0c9150506104b4565b8060801b60205283603e1b6001605d1b178184176001600160a01b03881161034d57818860601b179150610357565b8763800000008217555b558361036657806000526103ab565b83831780548060601c8061037e575063800000008217545b80891061039e5750673fffffff800000001916601f83901b1790556103ab565b50637fffffff1916821790555b93506001605d1b5b60005185146104a55782851754603e1c637fffffff16808417548281166103db5750506104a5565b603e81901c637fffffff1685811754601f637fffffff8216851402601f811882821c637fffffff16808a175480891661047e5786841c637fffffff168d0361042c57879c5061042c84848f8e61021c565b8a8d1754603e1c637fffffff169750878b1754965088198716888c175561045a603e88637fffffff911c1690565b8b811780548b17905595506104718385888e61021c565b50505050505050506103b3565b88198716888c175588198116828c17555050505084811782881755508097505050506103b3565b60005183178119815416815550505b949350505050565b6001605d1b5b60005183146105fc5782821754808216156104dd57506105fc565b603e81901c637fffffff1683811754909150601f637fffffff8216861402601f811882821c637fffffff1680871754808716156105425786198116828917558685178689175561052f8484888b61021c565b505085841754821c637fffffff16808717545b80831c637fffffff168881175482861c637fffffff16808b17548083178b1661057b5750505050861790871755509194506104c2915050565b808b166105c4578a198316848d17558a8517868d175561059d8789888f61021c565b8b8a1754881c637fffffff168c811754909650945084881c637fffffff169150818c175490505b898c175498508885188b168518868d17558a1989168a8d17558a198116828d17555050505050506105f78282868961021c565b505050505b9117805491199091169055565b8161061657505060005250565b178054637fffffff601f81831695909514159490940293841b19169190921b179055565b600060205160801c831115610654575063ccd52fbc610890565b82610664575063b113638a610890565b818317548390637fffffff601f82901c81169116808202156106a0578192505b84831754637fffffff168061069957506106a0565b9250610684565b505082811754637fffffff81168015601f0282901c637fffffff169050603e82901c637fffffff168186178054637fffffff603e1b1916603e83901b1790556106eb84838389610609565b50858314610765578486175461070d8785603e84901c637fffffff1689610609565b637fffffff811686178054637fffffff603e1b1916603e86901b179055601f81901c637fffffff1686178054637fffffff603e1b1916603e86901b17905582186bffffffffffffffffffffffff168218858417559194915b816001605d1b1661077a5761077a81866104bc565b505060205160801c808417548060601c60008161079f57505063800000008583171754805b8487175460601c6000816107bb57505063800000008786171754805b8184146108665784878a17558083146107da5782878a17638000000017555b603e85901c637fffffff16806107f35787600052610814565b8981178054637fffffff8082168a1415601f028b811b91901b199091161790555b50601f85901c637fffffff16801561083e57808a178054637fffffff603e1b1916603e8a901b1790555b50637fffffff8516801561086457808a178054637fffffff603e1b1916603e8a901b1790555b505b50506000848817558015610881576000848817638000000017555b5050506000190160801b602052505b92915050565b386000528554601052816108b7576108b0838587896102e9565b90506108c4565b6108c1848761063a565b90505b601051865595945050505050565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126108f957600080fd5b813567ffffffffffffffff811115610913576109136108d2565b8060051b604051601f19603f830116810181811067ffffffffffffffff82111715610940576109406108d2565b60405291825260208185018101929081018684111561095e57600080fd5b6020860192505b8383101561097d578235815260209283019201610965565b5095945050505050565b6000806040838503121561099a57600080fd5b823567ffffffffffffffff8111156109b157600080fd5b6109bd858286016108e8565b925050602083013567ffffffffffffffff8111156109da57600080fd5b6109e6858286016108e8565b9150509250929050565b634e487b7160e01b600052603260045260246000fdfea264697066735822122060f7fd286cb5b1331bc5d1520a171917f649e24017f29882e087085a042e6a0c64736f6c634300081a0033"
	statedb, err := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	if err != nil {
		b.Fatal(err)
	}
	evm := NewEVM(
		BlockContext{
			BlockNumber:        common.Big0,
			Transfer:           func(_ StateDB, _, _ common.Address, _ *big.Int) {},
			PublishEvents:      make(PublishEventsMap),
			CurrentTransaction: types.NewTx(&types.LegacyTx{}),
		},
		TxContext{},
		statedb,
		&params.ChainConfig{
			LondonBlock: common.Big0,
		},
		Config{},
	)

	testContract := common.Hex2Bytes(contractString)
	testAddress := common.BigToAddress(big.NewInt(0x201))

	statedb.SetCode(testAddress, testContract)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, err = evm.Call(AccountRef(testAddress), testAddress, data, 100_000_000_000, big.NewInt(0))
		if err != nil {
			b.Fatal(err)
		}
	}
}

