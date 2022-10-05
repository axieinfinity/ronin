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
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"os"
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

/**
verifyHeadersTestCode represents the following smart contract code

// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.0 <0.9.0;

contract VerifyHeaderTestContract {

    // Convert an hexadecimal character to their value
    function fromHexChar(uint8 c) public pure returns (uint8) {
        if (bytes1(c) >= bytes1('0') && bytes1(c) <= bytes1('9')) {
            return c - uint8(bytes1('0'));
        }
        if (bytes1(c) >= bytes1('a') && bytes1(c) <= bytes1('f')) {
            return 10 + c - uint8(bytes1('a'));
        }
        if (bytes1(c) >= bytes1('A') && bytes1(c) <= bytes1('F')) {
            return 10 + c - uint8(bytes1('A'));
        }
        revert("fail");
    }

    // Convert an hexadecimal string to raw bytes
    function fromHex(string memory s) public pure returns (bytes memory) {
        bytes memory ss = bytes(s);
        require(ss.length%2 == 0); // length must be even
        bytes memory r = new bytes(ss.length/2);
        for (uint i=0; i<ss.length/2; ++i) {
            r[i] = bytes1(fromHexChar(uint8(ss[2*i])) * 16 + fromHexChar(uint8(ss[2*i+1])));
        }
        return r;
    }

    struct BlockHeader {
        bytes32 parentHash;
        bytes32 ommersHash;
        bytes32 stateRoot;
        bytes32 transactionsRoot;
        bytes32 receiptsRoot;
        uint8[256] logsBloom;
        uint256 difficulty;
        uint64 number;
        uint64 gasLimit;
        uint64 gasUsed;
        uint64 timestamp;
        bytes extraData;
        bytes32 mixHash;
        uint64 nonce;
    }

    constructor() {}

    function verify() public view returns (bool) {
        uint8[256] memory bloom;
        bytes memory ex1 = fromHex("00000000000000000000000000000000000000000000000000000000000000000000000000000000000076616c696461746f72310000000000000000000076616c696461746f72321b8e2142d31f1b0fe686285ba79f7fe652ae3e0dada8450f7c1bacb22bbc1db06d48e0cae9392612797660c54d0a002c394a899d78869864b8cb07d926aac33301");
        bytes memory ex2 = fromHex("00000000000000000000000000000000000000000000000000000000000000000000000000000000000076616c696461746f72310000000000000000000076616c696461746f723251ed4084e48b910519506e98b829ae733590707a03264ae1b096a3668874163c036a7033caf7693821d9724cba1c409a6e3bafdb25817ebfab42e7de95d1fd0f01");
        BlockHeader memory header1 = BlockHeader("11", "", "123", "abc", "def", bloom, 1000, 1000, 100000000, 0, 1000, ex1, "", 1000);
        BlockHeader memory header2 = BlockHeader("11", "", "1232", "abcd", "defd", bloom, 1000, 1000, 100000000, 0, 1000, ex2, "", 1000);

        bytes memory data1 = _packBlockHeader(header1);
        bytes memory data2 = _packBlockHeader(header2);

        bytes memory payload = abi.encodeWithSignature("verifyHeaders(bytes,bytes)", data1, data2);
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

    function _packBlockHeader(BlockHeader memory _header) private pure returns (bytes memory) {
    return
    abi.encodePacked(
      abi.encode(
        _header.parentHash,
        _header.ommersHash,
        _header.stateRoot,
        _header.transactionsRoot,
        _header.receiptsRoot,
        _header.logsBloom,
        _header.difficulty,
        _header.number,
        _header.gasLimit,
        _header.gasUsed,
        _header.timestamp,
        _header.extraData
      ),
      abi.encode(
        _header.mixHash,
        _header.nonce
      )
    );
  }
}

*/
var (
	verifyHeadersTestCode = "608060405234801561001057600080fd5b50610c2a806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80632ecb20d3146100465780638e7e34d714610070578063fc735e9914610090575b600080fd5b61005961005436600461065d565b6100a8565b60405160ff90911681526020015b60405180910390f35b61008361007e36600461069d565b6101c4565b604051610067919061079d565b6100986102ec565b6040519015158152602001610067565b6000600360fc1b60f883901b6001600160f81b031916108015906100de5750603960f81b60f883901b6001600160f81b03191611155b156100f4576100ee6030836107c6565b92915050565b606160f81b60f883901b6001600160f81b031916108015906101285750603360f91b60f883901b6001600160f81b03191611155b1561014457606161013a83600a6107df565b6100ee91906107c6565b604160f81b60f883901b6001600160f81b031916108015906101785750602360f91b60f883901b6001600160f81b03191611155b1561018a57604161013a83600a6107df565b60405162461bcd60e51b81526004016101bb9060208082526004908201526319985a5b60e21b604082015260600190565b60405180910390fd5b805160609082906101d79060029061080e565b156101e157600080fd5b6000600282516101f19190610822565b6001600160401b0381111561020857610208610687565b6040519080825280601f01601f191660200182016040528015610232576020820181803683370190505b50905060005b600283516102469190610822565b8110156102e4576102848361025c836002610836565b61026790600161084d565b8151811061027757610277610860565b016020015160f81c6100a8565b61029384610267846002610836565b61029e906010610876565b6102a891906107df565b60f81b8282815181106102bd576102bd610860565b60200101906001600160f81b031916908160001a9053506102dd81610899565b9050610238565b509392505050565b60006102f661061f565b600061031c6040518061014001604052806101128152602001610ae361011291396101c4565b9050600061034460405180610140016040528061011281526020016109d161011291396101c4565b90506000604051806101c0016040528061313160f01b8152602001600080191681526020016231323360e81b81526020016261626360e81b8152602001623232b360e91b81526020018581526020016103e881526020016103e86001600160401b031681526020016305f5e1006001600160401b0316815260200160006001600160401b031681526020016103e86001600160401b03168152602001848152602001600080191681526020016103e86001600160401b031681525090506000604051806101c0016040528061313160f01b815260200160008019168152602001631899199960e11b815260200163185898d960e21b8152602001631919599960e21b81526020018681526020016103e881526020016103e86001600160401b031681526020016305f5e1006001600160401b0316815260200160006001600160401b031681526020016103e86001600160401b03168152602001848152602001600080191681526020016103e86001600160401b0316815250905060006104ca83610550565b905060006104d783610550565b9050600082826040516024016104ee9291906108b2565b60408051601f198184030181529190526020810180516001600160e01b03166344ce662b60e01b1790528051909150606761052761063f565b602084016020828583866000fa61053d57600080fd5b505115159b9a5050505050505050505050565b6060816000015182602001518360400151846060015185608001518660a001518760c001518860e001518961010001518a61012001518b61014001518c61016001516040516020016105ad9c9b9a999897969594939291906108e0565b604051602081830303815290604052826101800151836101a001516040516020016105eb9291909182526001600160401b0316602082015260400190565b60408051601f198184030181529082905261060992916020016109a1565b6040516020818303038152906040529050919050565b604051806120000160405280610100906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b60006020828403121561066f57600080fd5b813560ff8116811461068057600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b6000602082840312156106af57600080fd5b81356001600160401b03808211156106c657600080fd5b818401915084601f8301126106da57600080fd5b8135818111156106ec576106ec610687565b604051601f8201601f19908116603f0116810190838211818310171561071457610714610687565b8160405282815287602084870101111561072d57600080fd5b826020860160208301376000928101602001929092525095945050505050565b60005b83811015610768578181015183820152602001610750565b50506000910152565b6000815180845261078981602086016020860161074d565b601f01601f19169290920160200192915050565b6020815260006106806020830184610771565b634e487b7160e01b600052601160045260246000fd5b60ff82811682821603908111156100ee576100ee6107b0565b60ff81811683821601908111156100ee576100ee6107b0565b634e487b7160e01b600052601260045260246000fd5b60008261081d5761081d6107f8565b500690565b600082610831576108316107f8565b500490565b80820281158282048414176100ee576100ee6107b0565b808201808211156100ee576100ee6107b0565b634e487b7160e01b600052603260045260246000fd5b60ff8181168382160290811690818114610892576108926107b0565b5092915050565b6000600182016108ab576108ab6107b0565b5060010190565b6040815260006108c56040830185610771565b82810360208401526108d78185610771565b95945050505050565b60006121608e835260208e818501528d60408501528c60608501528b608085015260a084018b60005b61010081101561092a57815160ff1683529183019190830190600101610909565b50505050886120a084015261094b6120c08401896001600160401b03169052565b6001600160401b0387166120e08401526001600160401b0386166121008401526001600160401b0385166121208401528061214084015261098e81840185610771565b9f9e505050505050505050505050505050565b600083516109b381846020880161074d565b8351908301906109c781836020880161074d565b0194935050505056fe3030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303736363136633639363436313734366637323331303030303030303030303030303030303030303037363631366336393634363137343666373233323531656434303834653438623931303531393530366539386238323961653733333539303730376130333236346165316230393661333636383837343136336330333661373033336361663736393338323164393732346362613163343039613665336261666462323538313765626661623432653764653935643166643066303130303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303037363631366336393634363137343666373233313030303030303030303030303030303030303030373636313663363936343631373436663732333231623865323134326433316631623066653638363238356261373966376665363532616533653064616461383435306637633162616362323262626331646230366434386530636165393339323631323739373636306335346430613030326333393461383939643738383639383634623863623037643932366161633333333031a2646970667358221220276adff4e07a2fce591f03f80dc4c293983aaff0314861cd1d54318dfd77ed0f64736f6c63430008110033"
	verifyHeadersTestAbi  = `[{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"verify","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"}]`
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

// TestConsortiumVerifyHeaders_verify tests verify function
func TestConsortiumVerifyHeaders_verify(t *testing.T) {
	header1, header2, err := prepareHeader(big1)
	if err != nil {
		t.Fatal(err)
	}
	c := &consortiumVerifyHeaders{evm: &EVM{chainConfig: &params.ChainConfig{ChainID: big1}}}
	if !c.verify(*fromHeader(header1), *fromHeader(header2)) {
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
	encodedHeader1, err := fromHeader(header1).Bytes()
	if err != nil {
		t.Fatal(err)
	}
	encodedHeader2, err := fromHeader(header2).Bytes()
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
	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.LvlInfo)
	log.Root().SetHandler(glogger)
	var (
		statedb, _ = state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	)
	smcAbi, err := abi.JSON(strings.NewReader(verifyHeadersTestAbi))
	if err != nil {
		t.Fatal(err)
	}
	input, err := smcAbi.Pack("verify")
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
