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
