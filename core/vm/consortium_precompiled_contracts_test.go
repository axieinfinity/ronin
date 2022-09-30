package vm

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"
	"math/big"
	"strings"
	"testing"
)

func TestSort(t *testing.T) {
	addrs := []common.Address{
		common.BytesToAddress([]byte{102}),
		common.BytesToAddress([]byte{103}),
		common.BytesToAddress([]byte{104}),
		common.BytesToAddress([]byte{105}),
		common.BytesToAddress([]byte{106}),
	}

	totalBalances := []*big.Int{
		big.NewInt(10),
		big.NewInt(2),
		big.NewInt(11),
		big.NewInt(15),
		big.NewInt(1),
	}

	expectedAddrs := []common.Address{
		common.BytesToAddress([]byte{105}),
		common.BytesToAddress([]byte{104}),
		common.BytesToAddress([]byte{102}),
		common.BytesToAddress([]byte{103}),
		common.BytesToAddress([]byte{106}),
	}

	expectedBalances := []*big.Int{
		big.NewInt(15),
		big.NewInt(11),
		big.NewInt(10),
		big.NewInt(2),
		big.NewInt(1),
	}

	sortValidators(addrs, totalBalances)
	for i, val := range addrs {
		println(fmt.Sprintf("%s:%s", val.Hex(), totalBalances[i].String()))
		require.Equal(t, expectedAddrs[i].Hex(), val.Hex())
		require.Equal(t, expectedBalances[i].String(), totalBalances[i].String())
	}
}

/**
testSortCodes is generated based on the following code
```
// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.0 <0.9.0;

contract Validators {
    address[] _validators = [
    0x0000000000000000000000000000000000000010,
    0x0000000000000000000000000000000000000011,
    0x0000000000000000000000000000000000000012,
    0x0000000000000000000000000000000000000013,
    0x0000000000000000000000000000000000000014,
    0x0000000000000000000000000000000000000015,
    0x0000000000000000000000000000000000000016,
    0x0000000000000000000000000000000000000017,
    0x0000000000000000000000000000000000000018,
    0x0000000000000000000000000000000000000019,
    0x0000000000000000000000000000000000000020,
    0x0000000000000000000000000000000000000021,
    0x0000000000000000000000000000000000000022,
    0x0000000000000000000000000000000000000023,
    0x0000000000000000000000000000000000000024,
    0x0000000000000000000000000000000000000025,
    0x0000000000000000000000000000000000000026,
    0x0000000000000000000000000000000000000027,
    0x0000000000000000000000000000000000000028,
    0x0000000000000000000000000000000000000029,
    0x0000000000000000000000000000000000000030
    ];
    uint256[] _totalBalances = [1,4,6,2,8,9,10,3,16,20,100,12,22,30,50,60,5,18,16,22,21];
	constructor() {}
    function getValidators() public view returns (address[] memory _validatorList) {
        _validatorList = _validators;
    }
    function totalBalances(address[] calldata _poolList) public view returns (uint256[] memory _balances) {
        _balances = _totalBalances;
    }
}
```
*/
var testSortCode = `6080604052604051806102a00160405280601073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001601173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001601273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001601373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001601473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001601573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001601673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001601773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001601873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001601973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001602073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001602173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001602273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001602373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001602473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001602573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001602673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001602773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001602873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001602973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001603073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250600090601561044d92919061054d565b50604051806102a00160405280600160ff168152602001600460ff168152602001600660ff168152602001600260ff168152602001600860ff168152602001600960ff168152602001600a60ff168152602001600360ff168152602001601060ff168152602001601460ff168152602001606460ff168152602001600c60ff168152602001601660ff168152602001601e60ff168152602001603260ff168152602001603c60ff168152602001600560ff168152602001601260ff168152602001601060ff168152602001601660ff168152602001601560ff16815250600190601561053a9291906105d7565b5034801561054757600080fd5b50610646565b8280548282559060005260206000209081019282156105c6579160200282015b828111156105c55782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509160200191906001019061056d565b5b5090506105d39190610629565b5090565b828054828255906000526020600020908101928215610618579160200282015b82811115610617578251829060ff169055916020019190600101906105f7565b5b5090506106259190610629565b5090565b5b8082111561064257600081600090555060010161062a565b5090565b610460806106556000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80634a5d76cd1461003b578063b7ab4db51461006b575b600080fd5b610055600480360381019061005091906101e1565b610089565b60405161006291906102f6565b60405180910390f35b6100736100e4565b6040516100809190610408565b60405180910390f35b606060018054806020026020016040519081016040528092919081815260200182805480156100d757602002820191906000526020600020905b8154815260200190600101908083116100c3575b5050505050905092915050565b6060600080548060200260200160405190810160405280929190818152602001828054801561016857602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161011e575b5050505050905090565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b60008083601f8401126101a1576101a061017c565b5b8235905067ffffffffffffffff8111156101be576101bd610181565b5b6020830191508360208202830111156101da576101d9610186565b5b9250929050565b600080602083850312156101f8576101f7610172565b5b600083013567ffffffffffffffff81111561021657610215610177565b5b6102228582860161018b565b92509250509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000819050919050565b61026d8161025a565b82525050565b600061027f8383610264565b60208301905092915050565b6000602082019050919050565b60006102a38261022e565b6102ad8185610239565b93506102b88361024a565b8060005b838110156102e95781516102d08882610273565b97506102db8361028b565b9250506001810190506102bc565b5085935050505092915050565b600060208201905081810360008301526103108184610298565b905092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061036f82610344565b9050919050565b61037f81610364565b82525050565b60006103918383610376565b60208301905092915050565b6000602082019050919050565b60006103b582610318565b6103bf8185610323565b93506103ca83610334565b8060005b838110156103fb5781516103e28882610385565b97506103ed8361039d565b9250506001810190506103ce565b5085935050505092915050565b6000602082019050818103600083015261042281846103aa565b90509291505056fea26469706673582212204c2898e5390859d5425b29eb0ddb24a9e9c0d9dbeb67c0d2d7d7216ff51867d764736f6c63430008110033`
var expectedValidators = []common.Address{
	common.HexToAddress("0x0000000000000000000000000000000000000020"),
	common.HexToAddress("0x0000000000000000000000000000000000000025"),
	common.HexToAddress("0x0000000000000000000000000000000000000024"),
	common.HexToAddress("0x0000000000000000000000000000000000000023"),
	common.HexToAddress("0x0000000000000000000000000000000000000029"),
	common.HexToAddress("0x0000000000000000000000000000000000000022"),
	common.HexToAddress("0x0000000000000000000000000000000000000030"),
	common.HexToAddress("0x0000000000000000000000000000000000000019"),
	common.HexToAddress("0x0000000000000000000000000000000000000027"),
	common.HexToAddress("0x0000000000000000000000000000000000000028"),
	common.HexToAddress("0x0000000000000000000000000000000000000018"),
	common.HexToAddress("0x0000000000000000000000000000000000000021"),
	common.HexToAddress("0x0000000000000000000000000000000000000016"),
	common.HexToAddress("0x0000000000000000000000000000000000000015"),
	common.HexToAddress("0x0000000000000000000000000000000000000014"),
	common.HexToAddress("0x0000000000000000000000000000000000000012"),
	common.HexToAddress("0x0000000000000000000000000000000000000026"),
	common.HexToAddress("0x0000000000000000000000000000000000000011"),
	common.HexToAddress("0x0000000000000000000000000000000000000017"),
	common.HexToAddress("0x0000000000000000000000000000000000000013"),
	common.HexToAddress("0x0000000000000000000000000000000000000010"),
}

func TestConsortiumValidatorSorting_Run(t *testing.T) {

	var (
		caller      = common.BytesToAddress([]byte("sender"))
		statedb, _  = state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
		chainConfig = params.TestChainConfig
		limit       = 21
	)

	smcAbi, err := abi.JSON(strings.NewReader(consortiumSortValidatorAbi))
	if err != nil {
		t.Fatal(err)
	}
	input, err := smcAbi.Pack(sortValidatorsMethod, big.NewInt(int64(limit)))
	if err != nil {
		t.Fatal(err)
	}

	statedb.SetBalance(caller, math.BigPow(10, 18))
	statedb.SetNonce(caller, 1)

	evm := &EVM{
		Context: BlockContext{
			CurrentTransaction: types.NewTx(&types.LegacyTx{
				Nonce:    1,
				To:       nil,
				Value:    big.NewInt(0),
				Gas:      0,
				GasPrice: big.NewInt(0),
				Data:     []byte(""),
			}),
			CanTransfer: func(state StateDB, addr common.Address, value *big.Int) bool { return true },
			Transfer:    func(StateDB, common.Address, common.Address, *big.Int) {},
		},
		chainConfig: chainConfig,
		StateDB:     statedb,
		chainRules:  params.Rules{IsLondon: true},
	}
	evm.interpreter = NewEVMInterpreter(evm, Config{NoBaseFee: true})
	_, contract, _, err := evm.Create(AccountRef(caller), common.FromHex(testSortCode), math.MaxUint64/2, big0)
	if err != nil {
		t.Fatal(err)
	}
	chainConfig.ConsortiumV2Contracts = &params.ConsortiumV2Contracts{
		StakingContract:   contract,
		RoninValidatorSet: contract,
		SlashIndicator:    caller,
	}
	c := &consortiumValidatorSorting{caller: AccountRef(caller), evm: evm}
	output, err := c.Run(input)
	if err != nil {
		t.Fatal(err)
	}
	println(common.Bytes2Hex(output))

	res, err := smcAbi.Methods[sortValidatorsMethod].Outputs.Unpack(output)
	if err != nil {
		t.Fatal(err)
	}
	sortedValidators := *abi.ConvertType(res[0], new([]common.Address)).(*[]common.Address)
	require.Len(t, expectedValidators, limit, sortedValidators)
	for i, addr := range sortedValidators {
		println(addr.Hex())
		require.Equal(t, expectedValidators[i].Hex(), addr.Hex())
	}
}
