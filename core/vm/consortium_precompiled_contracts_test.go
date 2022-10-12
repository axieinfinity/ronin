package vm

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"strings"
	"testing"
)

/*
*
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

/*
*
wrapupCode is used to call sortValidators precompiled contract with the following code
```
// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.0 <0.9.0;

	contract Wrapup {
	    constructor() {}

	    function sortValidators()
	        public
	        view
	        returns (address[21] memory _validators)
	    {
	        address[21] memory validators = [
	            0x61A56247A283aA5a923CaAFec56509ccbb639290,
	            0xa7A4Be6067bB62b4D3879D3Df88e47e65609A7C7,
	            0x54f5D1849fbe650fAff0DF6A19dF3669fe70588A,
	            0xB74CaE2BC84555e009fb8093Fb0aad62281E8060,
	            0x65788575E220996bDEac700AA2ffaed0B7b539A6,
	            0x0870B1b8fB3800bc6a55ACce882A6C0d5b363E8A,
	            0x7eD04EAD6f31ba4DBE1371AF5Ee9f90f87CD96B5,
	            0x2b2a31b6CE5c0486455bb3119C95ADb3f8b5b1a6,
	            0xdE7d06682f3e11e7E9d1BE28339868d1aFDb7798,
	            0xc9f3A945cDbBBFe92c1c8F020B7dbeb783fE1f57,
	            0x10258AF6ae03e1B170e9bA8336349F85F9C63e27,
	            0x6c9A0E83AD142c082597054197aEb9b797785dAE,
	            0xCdC1Ca0D4a2c3A8289ff88f4D8eDbB153A1347eA,
	            0x5A482Ea3CCDe98600663c7E1D74679f1dc1b21f2,
	            0x825Fa8d5203637a4c1C3a6E7a379615fC69f57fB,
	            0xe0c06007DcFF4CD06064E44B91100b24d336a482,
	            0x63e750b6B1c38B149fA9457f380DdBE4a7021150,
	            0x10e336E30C422D25634fb16D6Ce87A0ec9CCC956,
	            0xa2c8C6Ca751C6bE33578814D3298d4ba37101B13,
	            0xBA780F5694E144c3765425E3CBb97760494F285b,
	            0xBD28217b06edB63e679d722F400EbAC89521Af59
	        ];
	        uint16[21] memory weights = [
	            1000,
	            2000,
	            3000,
	            4000,
	            5000,
	            6000,
	            7000,
	            8000,
	            9000,
	            10000,
	            11000,
	            12000,
	            13000,
	            14000,
	            15000,
	            16000,
	            17000,
	            18000,
	            19000,
	            20000,
	            21000
	        ];
	        bytes memory payload = abi.encodeWithSignature(
	            "sortValidators(address[],uint256[],256)",
	            validators,
	            weights,
	            21
	        );
	        uint256 payloadLength = payload.length;
	        address _smc = address(0x66);
	        assembly {
	            let payloadStart := add(payload, 32)
	            if iszero(
	                staticcall(
	                    0,
	                    _smc,
	                    payloadStart,
	                    payloadLength,
	                    _validators,
	                    0x2e0
	                )
	            ) {
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
	wrapupCode = `608060405234801561001057600080fd5b506109ef806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806327167aec14610030575b600080fd5b61003861004e565b60405161004591906108cd565b60405180910390f35b61005661079f565b6000604051806102a001604052807361a56247a283aa5a923caafec56509ccbb63929073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200173a7a4be6067bb62b4d3879d3df88e47e65609a7c773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020017354f5d1849fbe650faff0df6a19df3669fe70588a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200173b74cae2bc84555e009fb8093fb0aad62281e806073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020017365788575e220996bdeac700aa2ffaed0b7b539a673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001730870b1b8fb3800bc6a55acce882a6c0d5b363e8a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001737ed04ead6f31ba4dbe1371af5ee9f90f87cd96b573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001732b2a31b6ce5c0486455bb3119c95adb3f8b5b1a673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200173de7d06682f3e11e7e9d1be28339868d1afdb779873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200173c9f3a945cdbbbfe92c1c8f020b7dbeb783fe1f5773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020017310258af6ae03e1b170e9ba8336349f85f9c63e2773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001736c9a0e83ad142c082597054197aeb9b797785dae73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200173cdc1ca0d4a2c3a8289ff88f4d8edbb153a1347ea73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001735a482ea3ccde98600663c7e1d74679f1dc1b21f273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200173825fa8d5203637a4c1c3a6e7a379615fc69f57fb73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200173e0c06007dcff4cd06064e44b91100b24d336a48273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020017363e750b6b1c38b149fa9457f380ddbe4a702115073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020017310e336e30c422d25634fb16d6ce87a0ec9ccc95673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200173a2c8c6ca751c6be33578814d3298d4ba37101b1373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200173ba780f5694e144c3765425e3cbb97760494f285b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200173bd28217b06edb63e679d722f400ebac89521af5973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681525090506000604051806102a001604052806103e881526020016107d08152602001610bb88152602001610fa0815260200161138881526020016117708152602001611b588152602001611f40815260200161232881526020016127108152602001612af88152602001612ee081526020016132c881526020016136b08152602001613a988152602001613e80815260200161426881526020016146508152602001614a388152602001614e20815260200161520881525090506000828260156040516024016106f0939291906108e9565b6040516020818303038152906040527f4cb312b1000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050509050600081519050600060669050602083016102e0878483856000fa61079257600080fd5b6102a06040883e6102a087f35b604051806102a00160405280601590602082028036833780820191505090505090565b60006107ce83836107f2565b60208301905092915050565b60006107e683836108af565b60208301905092915050565b6107fb8161097d565b82525050565b61080a81610937565b6108148184610967565b925061081f82610923565b8060005b8381101561085057815161083787826107c2565b96506108428361094d565b925050600181019050610823565b505050505050565b61086181610942565b61086b8184610972565b92506108768261092d565b8060005b838110156108a757815161088e87826107da565b96506108998361095a565b92505060018101905061087a565b505050505050565b6108b8816109af565b82525050565b6108c7816109af565b82525050565b60006102a0820190506108e36000830184610801565b92915050565b6000610560820190506108ff6000830186610801565b61090d6102a0830185610858565b61091b6105408301846108be565b949350505050565b6000819050919050565b6000819050919050565b600060159050919050565b600060159050919050565b6000602082019050919050565b6000602082019050919050565b600081905092915050565b600081905092915050565b60006109888261098f565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600081905091905056fea26469706673582212206fb1aaa6c0418cfde9941c20876037a489b196ebb3eb0cc6c53cb9d0fff75d4564736f6c63430008070033`
	wrapupAbi  = `[{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"sortValidators","outputs":[{"internalType":"address[21]","name":"_validators","type":"address[21]"}],"stateMutability":"view","type":"function"}]`
)

var (
	caller             = common.BytesToAddress([]byte("sender"))
	expectedValidators = []common.Address{
		common.BytesToAddress([]byte{120}),
		common.BytesToAddress([]byte{119}),
		common.BytesToAddress([]byte{118}),
		common.BytesToAddress([]byte{117}),
		common.BytesToAddress([]byte{116}),
		common.BytesToAddress([]byte{115}),
		common.BytesToAddress([]byte{114}),
		common.BytesToAddress([]byte{113}),
		common.BytesToAddress([]byte{112}),
		common.BytesToAddress([]byte{111}),
		common.BytesToAddress([]byte{110}),
		common.BytesToAddress([]byte{109}),
		common.BytesToAddress([]byte{108}),
		common.BytesToAddress([]byte{107}),
		common.BytesToAddress([]byte{106}),
		common.BytesToAddress([]byte{105}),
		common.BytesToAddress([]byte{104}),
		common.BytesToAddress([]byte{103}),
		common.BytesToAddress([]byte{102}),
		common.BytesToAddress([]byte{101}),
		common.BytesToAddress([]byte{100}),
	}
)

/*
*
verifyHeadersTestCode represents the following smart contract code
// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.0 <0.9.0;

contract VerifyHeaderTestContract {

	    constructor() {}

	    function verify(bytes memory header1, bytes memory header2) public view returns (bool) {
	        bytes memory payload = abi.encodeWithSignature("validatingDoubleSignProof(bytes,bytes)", header1, header2);
	        uint payloadLength = payload.length;
	        address _smc = address(0x67);
	        uint[1] memory _output;
	        assembly {
	            let payloadStart := add(payload, 32)
	            if iszero(staticcall(0, _smc, payloadStart, payloadLength, _output, 0x20)) {
	                revert(0, 0)
	            }
	        }
	        return (_output[0] != 0);
	    }
	}
*/
var (
	verifyHeadersTestCode = "608060405234801561001057600080fd5b50610299806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063f7e83aee14610030575b600080fd5b61004361003e36600461018b565b610057565b604051901515815260200160405180910390f35b600080838360405160240161006d929190610235565b60408051601f198184030181529190526020810180516001600160e01b031663580a316360e01b179052805190915060676100a66100ca565b602084016020828583866000fa6100bc57600080fd5b505115159695505050505050565b60405180602001604052806001906020820280368337509192915050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261010f57600080fd5b813567ffffffffffffffff8082111561012a5761012a6100e8565b604051601f8301601f19908116603f01168101908282118183101715610152576101526100e8565b8160405283815286602085880101111561016b57600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806040838503121561019e57600080fd5b823567ffffffffffffffff808211156101b657600080fd5b6101c2868387016100fe565b935060208501359150808211156101d857600080fd5b506101e5858286016100fe565b9150509250929050565b6000815180845260005b81811015610215576020818501810151868301820152016101f9565b506000602082860101526020601f19601f83011685010191505092915050565b60408152600061024860408301856101ef565b828103602084015261025a81856101ef565b9594505050505056fea2646970667358221220e689890bbe17c2e97389470ed4baa21af25fd9cd6348d7511924615440d967d364736f6c63430008110033"
	verifyHeadersTestAbi  = `[{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[{"internalType":"bytes","name":"header1","type":"bytes"},{"internalType":"bytes","name":"header2","type":"bytes"}],"name":"verify","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"}]`
)

var (
	addressesTest = []common.Address{
		common.BytesToAddress([]byte{100}),
		common.BytesToAddress([]byte{101}),
		common.BytesToAddress([]byte{102}),
		common.BytesToAddress([]byte{103}),
		common.BytesToAddress([]byte{104}),
		common.BytesToAddress([]byte{105}),
		common.BytesToAddress([]byte{106}),
		common.BytesToAddress([]byte{107}),
		common.BytesToAddress([]byte{108}),
		common.BytesToAddress([]byte{109}),
		common.BytesToAddress([]byte{110}),
		common.BytesToAddress([]byte{111}),
		common.BytesToAddress([]byte{112}),
		common.BytesToAddress([]byte{113}),
		common.BytesToAddress([]byte{114}),
		common.BytesToAddress([]byte{115}),
		common.BytesToAddress([]byte{116}),
		common.BytesToAddress([]byte{117}),
		common.BytesToAddress([]byte{118}),
		common.BytesToAddress([]byte{119}),
		common.BytesToAddress([]byte{120}),
	}
	weightsTest = []*big.Int{
		big.NewInt(1_000_000),
		big.NewInt(2_000_000),
		big.NewInt(3_000_000),
		big.NewInt(4_000_000),
		big.NewInt(5_000_000),
		big.NewInt(6_000_000),
		big.NewInt(7_000_000),
		big.NewInt(8_000_000),
		big.NewInt(9_000_000),
		big.NewInt(10_000_000),
		big.NewInt(11_000_000),
		big.NewInt(12_000_000),
		big.NewInt(13_000_000),
		big.NewInt(14_000_000),
		big.NewInt(15_000_000),
		big.NewInt(16_000_000),
		big.NewInt(17_000_000),
		big.NewInt(18_000_000),
		big.NewInt(19_000_000),
		big.NewInt(20_000_000),
		big.NewInt(22_000_000),
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

// TestConsortiumValidatorSorting_Sort21ValidatorsSuccessfully sorts 21 validators successfully
func TestConsortiumValidatorSorting_Sort21ValidatorsSuccessfully(t *testing.T) {
	const limit = 21
	var (
		statedb, _ = state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	)

	smcAbi, err := abi.JSON(strings.NewReader(consortiumSortValidatorAbi))
	if err != nil {
		t.Fatal(err)
	}
	input, err := smcAbi.Pack(sortValidatorsMethod, addressesTest, weightsTest, big.NewInt(limit))
	fmt.Println(input)
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
	//println(common.Bytes2Hex(output))

	res, err := smcAbi.Methods[sortValidatorsMethod].Outputs.Unpack(output)
	if err != nil {
		t.Fatal(err)
	}
	sortedValidators := *abi.ConvertType(res[0], new([limit]common.Address)).(*[limit]common.Address)
	if len(expectedValidators) != len(sortedValidators) {
		t.Fatal(fmt.Sprintf("expected len %d, got %v", limit, len(sortedValidators)))
	}
	for i, addr := range sortedValidators {
		//println(addr.Hex())
		if expectedValidators[i].Hex() != addr.Hex() {
			t.Fatal(fmt.Sprintf("mismatched addr at %d, expected:%s got:%s", i, expectedValidators[i].Hex(), addr.Hex()))
		}
	}
}

// TestConsortiumValidatorSorting_Run2 simulates a call from a user who trigger system contract to call `sort` precompiled contract
func TestConsortiumValidatorSorting_Run2(t *testing.T) {
	const limit = 21
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
	fmt.Println(input)

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

	ret, _, err := evm.StaticCall(AccountRef(caller), contract, input, 1000000000)
	if err != nil {
		t.Fatal(err)
	}

	res, err := smcAbi.Methods[sortValidatorsMethod].Outputs.Unpack(ret)
	if err != nil {
		t.Fatal(err)
	}

	sortedValidators := *abi.ConvertType(res[0], new([limit]common.Address)).(*[limit]common.Address)
	if len(expectedValidators) != len(sortedValidators) {
		t.Fatal(fmt.Sprintf("expected len 21, got %v", len(sortedValidators)))
	}
	for i, addr := range sortedValidators {
		if expectedValidators[i].Hex() != addr.Hex() {
			t.Fatal(fmt.Sprintf("mismatched addr at %d, expected:%s got:%s", i, expectedValidators[i].Hex(), addr.Hex()))
		}
	}
}

// TestConsortiumVerifyHeaders_verify tests verify function
func TestConsortiumVerifyHeaders_verify(t *testing.T) {
	header1, header2, err := prepareHeader(big1)
	if err != nil {
		t.Fatal(err)
	}
	c := &consortiumVerifyHeaders{evm: &EVM{chainConfig: &params.ChainConfig{ChainID: big1}}}
	if !c.verify(*fromHeader(header1, big1), *fromHeader(header2, big1)) {
		t.Fatal("expected true, got false")
	}
}

// TestConsortiumVerifyHeaders_Run init 2 headers, pack them and call `Run` function directly
func TestConsortiumVerifyHeaders_Run(t *testing.T) {
	var (
		statedb, _ = state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	)
	smcAbi, err := abi.JSON(strings.NewReader(consortiumVerifyHeadersAbi))
	if err != nil {
		t.Fatal(err)
	}
	evm, err := newEVM(caller, statedb)
	if err != nil {
		t.Fatal(err)
	}
	header1, header2, err := prepareHeader(evm.chainConfig.ChainID)
	if err != nil {
		t.Fatal(err)
	}
	encodedHeader1, err := fromHeader(header1, big1).Bytes()
	if err != nil {
		t.Fatal(err)
	}
	encodedHeader2, err := fromHeader(header2, big1).Bytes()
	if err != nil {
		t.Fatal(err)
	}
	input, err := smcAbi.Pack(verifyHeaders, encodedHeader1, encodedHeader2)
	if err != nil {
		t.Fatal(err)
	}
	c := &consortiumVerifyHeaders{evm: evm, caller: AccountRef(caller)}
	result, err := c.Run(input)
	if err != nil {
		t.Fatal(err)
	}
	if len(result) != 32 {
		t.Fatal(fmt.Sprintf("expected len 32 got %d", len(result)))
	}
	if result[len(result)-1] != 1 {
		t.Fatal(fmt.Sprintf("expected 1 (true) got %d", result[len(result)-1]))
	}
}

// TestConsortiumVerifyHeaders_Run2 deploys smart contract and call precompiled contracts via this contract
func TestConsortiumVerifyHeaders_Run2(t *testing.T) {
	var (
		statedb, _ = state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	)
	smcAbi, err := abi.JSON(strings.NewReader(verifyHeadersTestAbi))
	if err != nil {
		t.Fatal(err)
	}
	header1, header2, err := prepareHeader(big1)
	if err != nil {
		t.Fatal(err)
	}
	encodedHeader1, err := fromHeader(header1, big1).Bytes()
	if err != nil {
		t.Fatal(err)
	}
	encodedHeader2, err := fromHeader(header2, big1).Bytes()
	if err != nil {
		t.Fatal(err)
	}
	input, err := smcAbi.Pack("verify", encodedHeader1, encodedHeader2)
	if err != nil {
		t.Fatal(err)
	}
	evm, err := newEVM(caller, statedb)
	if err != nil {
		t.Fatal(err)
	}
	_, contract, _, err := evm.Create(AccountRef(caller), common.FromHex(verifyHeadersTestCode), math.MaxUint64/2, big0)
	if err != nil {
		t.Fatal(err)
	}

	// set this contract into consortiumV2Contracts to make it become a system contract
	evm.chainConfig.ConsortiumV2Contracts.SlashIndicator = contract

	ret, _, err := evm.StaticCall(AccountRef(caller), contract, input, 1000000)
	if err != nil {
		t.Fatal(err)
	}

	res, err := smcAbi.Methods["verify"].Outputs.Unpack(ret)
	if err != nil {
		t.Fatal(err)
	}
	result := *abi.ConvertType(res[0], new(bool)).(*bool)
	if !result {
		t.Fatal("expected true got false")
	}
}

func prepareHeader(chainId *big.Int) (*types.Header, *types.Header, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, nil, err
	}
	// init extraData with extraVanity
	extraData := bytes.Repeat([]byte{0x00}, extraVanity)

	// append to extraData with validators set
	extraData = append(extraData, common.BytesToAddress([]byte("validator1")).Bytes()...)
	extraData = append(extraData, common.BytesToAddress([]byte("validator2")).Bytes()...)

	// add extra seal space
	extraData = append(extraData, make([]byte, crypto.SignatureLength)...)

	// create header1
	header1 := &types.Header{
		ParentHash:  common.BytesToHash([]byte("11")),
		UncleHash:   common.Hash{},
		Coinbase:    crypto.PubkeyToAddress(privateKey.PublicKey),
		Root:        common.BytesToHash([]byte("123")),
		TxHash:      common.BytesToHash([]byte("abc")),
		ReceiptHash: common.BytesToHash([]byte("def")),
		Bloom:       types.Bloom{},
		Difficulty:  big.NewInt(1000),
		Number:      big.NewInt(1000),
		GasLimit:    100000000,
		GasUsed:     0,
		Time:        1000,
		Extra:       make([]byte, len(extraData)),
		MixDigest:   common.Hash{},
		Nonce:       types.EncodeNonce(1000),
	}

	// create header2
	header2 := &types.Header{
		ParentHash:  common.BytesToHash([]byte("11")),
		UncleHash:   common.Hash{},
		Coinbase:    crypto.PubkeyToAddress(privateKey.PublicKey),
		Root:        common.BytesToHash([]byte("1232")),
		TxHash:      common.BytesToHash([]byte("abcd")),
		ReceiptHash: common.BytesToHash([]byte("defd")),
		Bloom:       types.Bloom{},
		Difficulty:  big.NewInt(1000),
		Number:      big.NewInt(1000),
		GasLimit:    100000000,
		GasUsed:     0,
		Time:        1000,
		Extra:       make([]byte, len(extraData)),
		MixDigest:   common.Hash{},
		Nonce:       types.EncodeNonce(1000),
	}

	// copy extraData
	copy(header1.Extra[:], extraData)
	copy(header2.Extra[:], extraData)

	// signing and add to extraData
	sig1, err := crypto.Sign(crypto.Keccak256(consortiumRlp(header1, chainId)), privateKey)
	if err != nil {
		return nil, nil, err
	}
	sig2, err := crypto.Sign(crypto.Keccak256(consortiumRlp(header2, chainId)), privateKey)
	if err != nil {
		return nil, nil, err
	}

	copy(header1.Extra[len(header1.Extra)-crypto.SignatureLength:], sig1)
	copy(header2.Extra[len(header2.Extra)-crypto.SignatureLength:], sig2)

	return header1, header2, nil
}

func consortiumRlp(header *types.Header, chainId *big.Int) []byte {
	b := new(bytes.Buffer)
	encodeSigHeader(b, header, chainId)
	return b.Bytes()
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
