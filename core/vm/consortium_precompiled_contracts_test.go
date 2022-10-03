package vm

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"strings"
	"testing"
)

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
    function getValidatorCandidates() public view returns (address[] memory _validatorList) {
        _validatorList = _validators;
    }
    function totalBalances(address[] calldata _poolList) public view returns (uint256[] memory _balances) {
        _balances = _totalBalances;
    }
}
```
*/
var testSortCode = `6080604052604051806102a00160405280601073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001601173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001601273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001601373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001601473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001601573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001601673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001601773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001601873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001601973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001602073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001602173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001602273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001602373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001602473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001602573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001602673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001602773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001602873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001602973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001603073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250600090601561044d92919061054d565b50604051806102a00160405280600160ff168152602001600460ff168152602001600660ff168152602001600260ff168152602001600860ff168152602001600960ff168152602001600a60ff168152602001600360ff168152602001601060ff168152602001601460ff168152602001606460ff168152602001600c60ff168152602001601660ff168152602001601e60ff168152602001603260ff168152602001603c60ff168152602001600560ff168152602001601260ff168152602001601060ff168152602001601660ff168152602001601560ff16815250600190601561053a9291906105d7565b5034801561054757600080fd5b50610646565b8280548282559060005260206000209081019282156105c6579160200282015b828111156105c55782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509160200191906001019061056d565b5b5090506105d39190610629565b5090565b828054828255906000526020600020908101928215610618579160200282015b82811115610617578251829060ff169055916020019190600101906105f7565b5b5090506106259190610629565b5090565b5b8082111561064257600081600090555060010161062a565b5090565b610460806106556000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80634a5d76cd1461003b578063ba77b06c1461006b575b600080fd5b610055600480360381019061005091906101e1565b610089565b60405161006291906102f6565b60405180910390f35b6100736100e4565b6040516100809190610408565b60405180910390f35b606060018054806020026020016040519081016040528092919081815260200182805480156100d757602002820191906000526020600020905b8154815260200190600101908083116100c3575b5050505050905092915050565b6060600080548060200260200160405190810160405280929190818152602001828054801561016857602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161011e575b5050505050905090565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b60008083601f8401126101a1576101a061017c565b5b8235905067ffffffffffffffff8111156101be576101bd610181565b5b6020830191508360208202830111156101da576101d9610186565b5b9250929050565b600080602083850312156101f8576101f7610172565b5b600083013567ffffffffffffffff81111561021657610215610177565b5b6102228582860161018b565b92509250509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000819050919050565b61026d8161025a565b82525050565b600061027f8383610264565b60208301905092915050565b6000602082019050919050565b60006102a38261022e565b6102ad8185610239565b93506102b88361024a565b8060005b838110156102e95781516102d08882610273565b97506102db8361028b565b9250506001810190506102bc565b5085935050505092915050565b600060208201905081810360008301526103108184610298565b905092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061036f82610344565b9050919050565b61037f81610364565b82525050565b60006103918383610376565b60208301905092915050565b6000602082019050919050565b60006103b582610318565b6103bf8185610323565b93506103ca83610334565b8060005b838110156103fb5781516103e28882610385565b97506103ed8361039d565b9250506001810190506103ce565b5085935050505092915050565b6000602082019050818103600083015261042281846103aa565b90509291505056fea26469706673582212204aab6725297603e4a73daf57fffdd55330e60293dfde78dcbbddb9178144583064736f6c63430008110033`

/**
wrapupCode is used to call sortValidators precompiled contract with the following code
```
// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.0 <0.9.0;

contract Wrapup {
    constructor() {}

    function sortValidators() public view returns (address[21] memory _validators) {
        bytes memory payload = abi.encodeWithSignature("sortValidators(uint256)", 21);
        uint payloadLength = payload.length;
        address _smc = address(0x66);
        assembly {
            let payloadStart := add(payload, 32)
            if iszero(staticcall(0, _smc, payloadStart, payloadLength, _validators, 0x2e0)) {
                revert(0, 0)
            }
            returndatacopy(_validators, 64, 672)
            return(_validators, 672)
        }
    }
}
```
*/
var (
	wrapupCode = `608060405234801561001057600080fd5b506102d8806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806327167aec14610030575b600080fd5b61003861004e565b6040516100459190610219565b60405180910390f35b610056610119565b6000601560405160240161006a9190610287565b6040516020818303038152906040527f2e73ae11000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050509050600081519050600060669050602083016102e0858483856000fa61010c57600080fd5b6102a06040863e6102a085f35b604051806102a00160405280601590602082028036833780820191505090505090565b600060159050919050565b600081905092915050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006101878261015c565b9050919050565b6101978161017c565b82525050565b60006101a9838361018e565b60208301905092915050565b6000602082019050919050565b6101cb8161013c565b6101d58184610147565b92506101e082610152565b8060005b838110156102115781516101f8878261019d565b9650610203836101b5565b9250506001810190506101e4565b505050505050565b60006102a08201905061022f60008301846101c2565b92915050565b6000819050919050565b600060ff82169050919050565b6000819050919050565b600061027161026c61026784610235565b61024c565b61023f565b9050919050565b61028181610256565b82525050565b600060208201905061029c6000830184610278565b9291505056fea2646970667358221220d9e99c6aa93156ae399b93dff44f456409a5715d61d7f761ea1ed96a42bba43864736f6c63430008110033`
	wrapupAbi  = `[{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"sortValidators","outputs":[{"internalType":"address[21]","name":"_validators","type":"address[21]"}],"stateMutability":"view","type":"function"}]`
)

var (
	caller             = common.BytesToAddress([]byte("sender"))
	expectedValidators = []common.Address{
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
		if expectedBalances[i].Cmp(totalBalances[i]) != 0 {
			t.Fatal(fmt.Sprintf("mismatched balance at %d, expected:%s got:%s", i, expectedBalances[i].String(), totalBalances[i].String()))
		}
		if expectedAddrs[i].Hex() != val.Hex() {
			t.Fatal(fmt.Sprintf("mismatched addr at %d, expected:%s got:%s", i, expectedValidators[i].Hex(), val.Hex()))
		}
	}
}

func TestConsortiumValidatorSorting_Run(t *testing.T) {
	var (
		statedb, _ = state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
		limit      = 21
	)

	smcAbi, err := abi.JSON(strings.NewReader(consortiumSortValidatorAbi))
	if err != nil {
		t.Fatal(err)
	}
	input, err := smcAbi.Pack(sortValidatorsMethod, big.NewInt(int64(limit)))
	if err != nil {
		t.Fatal(err)
	}

	evm, err := newEVM(caller, statedb)
	if err != nil {
		t.Fatal(err)
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
	if len(expectedValidators) != len(sortedValidators) {
		t.Fatal(fmt.Sprintf("expected len %d, got %v", limit, len(sortedValidators)))
	}
	for i, addr := range sortedValidators {
		println(addr.Hex())
		if expectedValidators[i].Hex() != addr.Hex() {
			t.Fatal(fmt.Sprintf("mismatched addr at %d, expected:%s got:%s", i, expectedValidators[i].Hex(), addr.Hex()))
		}
	}
}

// TestConsortiumValidatorSorting_Run2 simulates a call from a user who trigger system contract to call `sort` precompiled contract
func TestConsortiumValidatorSorting_Run2(t *testing.T) {
	var (
		statedb, _ = state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	)
	smcAbi, err := abi.JSON(strings.NewReader(wrapupAbi))
	if err != nil {
		t.Fatal(err)
	}
	input, err := smcAbi.Pack(sortValidatorsMethod)
	if err != nil {
		t.Fatal(err)
	}
	evm, err := newEVM(caller, statedb)
	if err != nil {
		t.Fatal(err)
	}
	_, contract, _, err := evm.Create(AccountRef(caller), common.FromHex(wrapupCode), math.MaxUint64/2, big0)
	if err != nil {
		t.Fatal(err)
	}

	// set this contract into consortiumV2Contracts to make it become a system contract
	evm.chainConfig.ConsortiumV2Contracts.SlashIndicator = contract

	ret, _, err := evm.StaticCall(AccountRef(caller), contract, input, 1000000)
	if err != nil {
		t.Fatal(err)
	}

	res, err := smcAbi.Methods[sortValidatorsMethod].Outputs.Unpack(ret)
	if err != nil {
		t.Fatal(err)
	}

	sortedValidators := *abi.ConvertType(res[0], new([21]common.Address)).(*[21]common.Address)
	if len(expectedValidators) != len(sortedValidators) {
		t.Fatal(fmt.Sprintf("expected len 21, got %v", len(sortedValidators)))
	}
	for i, addr := range sortedValidators {
		println(addr.Hex())
		if expectedValidators[i].Hex() != addr.Hex() {
			t.Fatal(fmt.Sprintf("mismatched addr at %d, expected:%s got:%s", i, expectedValidators[i].Hex(), addr.Hex()))
		}
	}
}

func newEVM(caller common.Address, statedb StateDB) (*EVM, error) {
	evm := &EVM{
		Context: BlockContext{
			CanTransfer: func(state StateDB, addr common.Address, value *big.Int) bool { return true },
			Transfer:    func(StateDB, common.Address, common.Address, *big.Int) {},
		},
		chainConfig: params.TestChainConfig,
		StateDB:     statedb,
		chainRules:  params.Rules{IsIstanbul: true},
	}
	evm.interpreter = NewEVMInterpreter(evm, Config{NoBaseFee: true})
	_, contract, _, err := evm.Create(AccountRef(caller), common.FromHex(testSortCode), math.MaxUint64/2, big0)
	if err != nil {
		return nil, err
	}
	evm.chainConfig.ConsortiumV2Contracts = &params.ConsortiumV2Contracts{
		StakingContract:   contract,
		RoninValidatorSet: contract,
		SlashIndicator:    caller,
	}
	return evm, nil
}
