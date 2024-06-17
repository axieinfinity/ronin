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
	"github.com/ethereum/go-ethereum/params"
)

type testSuite struct {
	setupTransactions   []transaction // The transactions run before benchmark loop
	prepareTransactions []transaction // The transactions run before each loop iteration
	benchTransactions   []transaction // The actual transactions need to be benchmarked
	contracts           []contractCode
}

type contractCode struct {
	code    []byte
	address common.Address
}

type transaction struct {
	from     common.Address
	to       common.Address
	input    []byte
	gasLimit uint64
	value    *big.Int
}

func benchmarkEVM(b *testing.B, suite *testSuite) {
	statedb, err := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	if err != nil {
		b.Fatal(err)
	}

	for _, contract := range suite.contracts {
		statedb.SetCode(contract.address, contract.code)
	}

	evm := NewEVM(
		BlockContext{
			BlockNumber:   common.Big0,
			Transfer:      func(_ StateDB, _, _ common.Address, _ *big.Int) {},
			PublishEvents: make(PublishEventsMap),
		},
		TxContext{},
		statedb,
		&params.ChainConfig{
			LondonBlock: common.Big0,
		},
		Config{},
	)

	for _, tx := range suite.setupTransactions {
		_, _, err := evm.Call(AccountRef(tx.from), tx.to, tx.input, tx.gasLimit, tx.value)
		if err != nil {
			b.Fatal(err)
		}
	}

	hasPrepareTranction := len(suite.prepareTransactions) != 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if hasPrepareTranction {
			b.StopTimer()
			for _, tx := range suite.prepareTransactions {
				_, _, err := evm.Call(AccountRef(tx.from), tx.to, tx.input, tx.gasLimit, tx.value)
				if err != nil {
					b.Fatal(err)
				}
			}
			b.StartTimer()
		}
		for _, tx := range suite.benchTransactions {
			_, _, err := evm.Call(AccountRef(tx.to), tx.to, tx.input, tx.gasLimit, tx.value)
			if err != nil {
				b.Fatal(err)
			}
		}
	}

}

func BenchmarkEvmInsertionSort(b *testing.B) {
	contractAbi, _ := abi.JSON(strings.NewReader(`[{"inputs": [{"internalType": "uint256[]","name": "a","type": "uint256[]"}],"name": "insertionSort","outputs": [],"stateMutability": "pure","type": "function"}]`))
	rand.New(rand.NewSource(0))
	const inputLen = 1_000
	input := make([]*big.Int, inputLen)
	for i := 0; i < inputLen; i++ {
		input[i] = big.NewInt(int64(rand.Int31()))
	}
	data, err := contractAbi.Pack("insertionSort", input)
	if err != nil {
		b.Fatal(err)
	}

	testAddress := common.BigToAddress(big.NewInt(0x204))
	suite := testSuite{
		benchTransactions: []transaction{
			{
				to:       testAddress,
				input:    data,
				gasLimit: 1_000_000_000,
				value:    common.Big0,
			},
		},
		contracts: []contractCode{
			{
				// https://github.com/Vectorized/solady/blob/678c9163550810b08f0ffb09624c9f7532392303/src/utils/LibSort.sol#L17
				code:    common.Hex2Bytes("608060405234801561001057600080fd5b506004361061002b5760003560e01c80636297206f14610030575b600080fd5b61004a60048036038101906100459190610264565b61004c565b005b8051600082528060051b82016020601f198185015b6001156100b65782810190508381116100b6578051828201805182811161008a575050506100b1565b5b6001156100a857808683015284820191508151905082811161008b575b82868301525050505b610061565b508385525050505050565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610123826100da565b810181811067ffffffffffffffff82111715610142576101416100eb565b5b80604052505050565b60006101556100c1565b9050610161828261011a565b919050565b600067ffffffffffffffff821115610181576101806100eb565b5b602082029050602081019050919050565b600080fd5b6000819050919050565b6101aa81610197565b81146101b557600080fd5b50565b6000813590506101c7816101a1565b92915050565b60006101e06101db84610166565b61014b565b9050808382526020820190506020840283018581111561020357610202610192565b5b835b8181101561022c578061021888826101b8565b845260208401935050602081019050610205565b5050509392505050565b600082601f83011261024b5761024a6100d5565b5b813561025b8482602086016101cd565b91505092915050565b60006020828403121561027a576102796100cb565b5b600082013567ffffffffffffffff811115610298576102976100d0565b5b6102a484828501610236565b9150509291505056fea2646970667358221220c05eb12b713fde4b945bd660ddeda4dc91b0d11b4749c8bf02617b25bf3c022064736f6c63430008120033"),
				address: testAddress,
			},
		},
	}

	benchmarkEVM(b, &suite)
}
