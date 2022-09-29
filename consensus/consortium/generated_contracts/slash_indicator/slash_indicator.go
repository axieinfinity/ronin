// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package slashIndicator

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SlashIndicatorMetaData contains all meta data concerning the SlashIndicator contract.
var SlashIndicatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"felonyJailDuration\",\"type\":\"uint256\"}],\"name\":\"FelonyJailDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"MaintenanceContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashDoubleSignAmount\",\"type\":\"uint256\"}],\"name\":\"SlashDoubleSignAmountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashFelonyAmount\",\"type\":\"uint256\"}],\"name\":\"SlashFelonyAmountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"felonyThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"misdemeanorThreshold\",\"type\":\"uint256\"}],\"name\":\"SlashThresholdsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumISlashIndicator.SlashType\",\"name\":\"slashType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"UnavailabilitySlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ValidatorContractUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"currentUnavailabilityIndicator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"felonyJailDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"felonyThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"getUnavailabilityIndicator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"getUnavailabilitySlashType\",\"outputs\":[{\"internalType\":\"enumISlashIndicator.SlashType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUnavailabilityThresholds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__validatorContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__maintenanceContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_misdemeanorThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_felonyThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slashFelonyAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slashDoubleSignAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_felonyJailBlocks\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastSlashedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maintenanceContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"misdemeanorThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_felonyJailDuration\",\"type\":\"uint256\"}],\"name\":\"setFelonyJailDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setMaintenanceContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slashDoubleSignAmount\",\"type\":\"uint256\"}],\"name\":\"setSlashDoubleSignAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slashFelonyAmount\",\"type\":\"uint256\"}],\"name\":\"setSlashFelonyAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_felonyThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_misdemeanorThreshold\",\"type\":\"uint256\"}],\"name\":\"setSlashThresholds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setValidatorContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"slashDoubleSign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashDoubleSignAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashFelonyAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"unavailabilityThresholdsOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_misdemeanorThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_felonyThreshold\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50620000226200002860201b60201c565b620001d6565b600160159054906101000a900460ff16156200007b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000729062000179565b60405180910390fd5b60ff8016600160149054906101000a900460ff1660ff161015620000f05760ff600160146101000a81548160ff021916908360ff1602179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860ff604051620000e79190620001b9565b60405180910390a15b565b600082825260208201905092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320696e69746960008201527f616c697a696e6700000000000000000000000000000000000000000000000000602082015250565b600062000161602783620000f2565b91506200016e8262000103565b604082019050919050565b60006020820190508181036000830152620001948162000152565b9050919050565b600060ff82169050919050565b620001b3816200019b565b82525050565b6000602082019050620001d06000830184620001a8565b92915050565b61232680620001e66000396000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c806399439089116100b8578063d2cb215e1161007c578063d2cb215e14610364578063d919358714610382578063dfe484cb146103a1578063e4f9c4be146103bf578063edbf4ac2146103db578063f82bdd01146103f757610142565b806399439089146102c0578063aef250be146102de578063c96be4cb146102fc578063cdf64a7614610318578063d180ecb01461033457610142565b8063567a372d1161010a578063567a372d146101fe5780635786e9ee1461021c57806362ffe6cb146102385780636b79f95f146102685780636d14c4e5146102865780636e91bec5146102a457610142565b8063082e742014610147578063389f4f711461017757806346fe9311146101955780634d961e18146101b1578063518e463a146101e2575b600080fd5b610161600480360381019061015c9190611821565b610413565b60405161016e9190611867565b60405180910390f35b61017f6104bf565b60405161018c9190611867565b60405180910390f35b6101af60048036038101906101aa9190611821565b6104c5565b005b6101cb60048036038101906101c691906118ae565b610546565b6040516101d99291906118ee565b60405180910390f35b6101fc60048036038101906101f7919061197c565b610873565b005b6102066109a1565b6040516102139190611867565b60405180910390f35b610236600480360381019061023191906119dc565b6109a7565b005b610252600480360381019061024d91906118ae565b610a28565b60405161025f9190611867565b60405180910390f35b610270610a83565b60405161027d9190611867565b60405180910390f35b61028e610a89565b60405161029b9190611867565b60405180910390f35b6102be60048036038101906102b991906119dc565b610a8f565b005b6102c8610b10565b6040516102d59190611a18565b60405180910390f35b6102e6610b39565b6040516102f39190611867565b60405180910390f35b61031660048036038101906103119190611821565b610b3f565b005b610332600480360381019061032d9190611821565b611146565b005b61034e600480360381019061034991906118ae565b6111c7565b60405161035b9190611aaa565b60405180910390f35b61036c61122f565b6040516103799190611a18565b60405180910390f35b61038a611259565b6040516103989291906118ee565b60405180910390f35b6103a961126a565b6040516103b69190611867565b60405180910390f35b6103d960048036038101906103d49190611ac5565b611270565b005b6103f560048036038101906103f09190611b05565b6112f3565b005b610411600480360381019061040c91906119dc565b61146a565b005b60006104b88260008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f8549af9436040518263ffffffff1660e01b81526004016104729190611867565b602060405180830381865afa15801561048f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104b39190611bbc565b610a28565b9050919050565b60065481565b6104cd6114eb565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461053a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053190611c6c565b60405180910390fd5b61054381611542565b50565b60008060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635186dc7e6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105db9190611bbc565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636aa1c2ef6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610646573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061066a9190611bbc565b6106749190611cbb565b905060008182866106859190611d44565b61068f9190611cbb565b90506000600183836106a19190611d75565b6106ab9190611da9565b90506000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d39fee34896040518263ffffffff1660e01b815260040161070a9190611a18565b606060405180830381865afa158015610727573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061074b9190611ed2565b90506000610768848484600001516115bd9092919063ffffffff16565b90506000610785858585602001516115bd9092919063ffffffff16565b905060008690508280156107965750815b156107cd576001846000015185602001516107b19190611da9565b6107bb9190611d75565b816107c69190611da9565b9050610833565b82156108015760018460000151866107e59190611da9565b6107ef9190611d75565b816107fa9190611da9565b9050610832565b81156108315760018685602001516108199190611da9565b6108239190611d75565b8161082e9190611da9565b90505b5b5b61084a81886005546115d89092919063ffffffff16565b985061086381886006546115d89092919063ffffffff16565b9750505050505050509250929050565b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d890611f71565b60405180910390fd5b6000801561099b5760008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370f81f6c857fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6008546040518463ffffffff1660e01b815260040161096893929190611f91565b600060405180830381600087803b15801561098257600080fd5b505af1158015610996573d6000803e3d6000fd5b505050505b50505050565b60055481565b6109af6114eb565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a1c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a1390611c6c565b60405180910390fd5b610a25816115fa565b50565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600083815260200190815260200160002054905092915050565b60095481565b60085481565b610a976114eb565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b04576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610afb90611c6c565b60405180910390fd5b610b0d8161163b565b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60045481565b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610bad576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ba490611f71565b60405180910390fd5b6004544311610bf1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610be890612060565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480610cc55750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663934e9a0382436040518363ffffffff1660e01b8152600401610c83929190612080565b602060405180830381865afa158015610ca0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cc491906120e1565b5b61113c5760008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f8549af9436040518263ffffffff1660e01b8152600401610d259190611867565b602060405180830381865afa158015610d42573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d669190611bbc565b90506000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600083815260200190815260200160002060008154610dc79061210e565b9190508190559050600080610ddc8543610546565b915091506000610dec86866111c7565b9050818410158015610e22575060026003811115610e0d57610e0c611a33565b5b816003811115610e2057610e1f611a33565b5b105b15610f99576002600360008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600087815260200190815260200160002060006101000a81548160ff02191690836003811115610e9b57610e9a611a33565b5b02179055508573ffffffffffffffffffffffffffffffffffffffff167f8c2c2bfe532ccdb4523fa2392954fd58445929e7f260d786d0ab93cd981cde54600287604051610ee9929190612156565b60405180910390a260008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370f81f6c8760095443610f3c9190611d75565b6007546040518463ffffffff1660e01b8152600401610f5d93929190611f91565b600060405180830381600087803b158015610f7757600080fd5b505af1158015610f8b573d6000803e3d6000fd5b50505050505050505061113c565b828410158015610fcd575060016003811115610fb857610fb7611a33565b5b816003811115610fcb57610fca611a33565b5b105b15611136576001600360008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600087815260200190815260200160002060006101000a81548160ff0219169083600381111561104657611045611a33565b5b02179055508573ffffffffffffffffffffffffffffffffffffffff167f8c2c2bfe532ccdb4523fa2392954fd58445929e7f260d786d0ab93cd981cde54600187604051611094929190612156565b60405180910390a260008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370f81f6c876000806040518463ffffffff1660e01b81526004016110fa939291906121c4565b600060405180830381600087803b15801561111457600080fd5b505af1158015611128573d6000803e3d6000fd5b50505050505050505061113c565b50505050505b4360048190555050565b61114e6114eb565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146111bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111b290611c6c565b60405180910390fd5b6111c48161167c565b50565b6000600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600083815260200190815260200160002060009054906101000a900460ff16905092915050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600080600554600654915091509091565b60075481565b6112786114eb565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146112e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112dc90611c6c565b60405180910390fd5b6112ef82826116f6565b5050565b6000600160159054906101000a900460ff16159050808015611326575060018060149054906101000a900460ff1660ff16105b80611354575061133530611741565b158015611353575060018060149054906101000a900460ff1660ff16145b5b611393576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161138a9061226d565b60405180910390fd5b60018060146101000a81548160ff021916908360ff16021790555080156113cf5760018060156101000a81548160ff0219169083151502179055505b6113d88861167c565b6113e187611542565b6113eb85876116f6565b6113f48461163b565b6113fd83611764565b611406826115fa565b8015611460576000600160156101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498600160405161145791906122d5565b60405180910390a15b5050505050505050565b6114726114eb565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146114df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114d690611c6c565b60405180910390fd5b6114e881611764565b50565b60006115197fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610360001b6117a5565b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f31a33f126a5bae3c5bdf6cfc2cd6dcfffe2fe9634bdb09e21c44762993889e3b816040516115b29190611a18565b60405180910390a150565b60008383111580156115cf5750818411155b90509392505050565b60008183856115e79190611cbb565b6115f19190611d44565b90509392505050565b806009819055507f3092a623ebcf71b79f9f68801b081a4d0c839dfb4c0e6f9ff118f4fb870375dc816040516116309190611867565b60405180910390a150565b806007819055507f71491762d43dbadd94f85cf7a0322192a7c22f3ca4fb30de8895420607b56712816040516116719190611867565b60405180910390a150565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169816040516116eb9190611a18565b60405180910390a150565b81600681905550806005819055507f25a70459f14137ea646efb17b806c2e75632b931d69e642a8a809b6bfd7217a682826040516117359291906118ee565b60405180910390a15050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b806008819055507f42458d43c99c17f034dde22f3a1cf003e4bc27372c1ff14cb9e68a79a1b1e8ed8160405161179a9190611867565b60405180910390a150565b6000819050919050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006117ee826117c3565b9050919050565b6117fe816117e3565b811461180957600080fd5b50565b60008135905061181b816117f5565b92915050565b600060208284031215611837576118366117b9565b5b60006118458482850161180c565b91505092915050565b6000819050919050565b6118618161184e565b82525050565b600060208201905061187c6000830184611858565b92915050565b61188b8161184e565b811461189657600080fd5b50565b6000813590506118a881611882565b92915050565b600080604083850312156118c5576118c46117b9565b5b60006118d38582860161180c565b92505060206118e485828601611899565b9150509250929050565b60006040820190506119036000830185611858565b6119106020830184611858565b9392505050565b600080fd5b600080fd5b600080fd5b60008083601f84011261193c5761193b611917565b5b8235905067ffffffffffffffff8111156119595761195861191c565b5b60208301915083600182028301111561197557611974611921565b5b9250929050565b600080600060408486031215611995576119946117b9565b5b60006119a38682870161180c565b935050602084013567ffffffffffffffff8111156119c4576119c36117be565b5b6119d086828701611926565b92509250509250925092565b6000602082840312156119f2576119f16117b9565b5b6000611a0084828501611899565b91505092915050565b611a12816117e3565b82525050565b6000602082019050611a2d6000830184611a09565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60048110611a7357611a72611a33565b5b50565b6000819050611a8482611a62565b919050565b6000611a9482611a76565b9050919050565b611aa481611a89565b82525050565b6000602082019050611abf6000830184611a9b565b92915050565b60008060408385031215611adc57611adb6117b9565b5b6000611aea85828601611899565b9250506020611afb85828601611899565b9150509250929050565b600080600080600080600060e0888a031215611b2457611b236117b9565b5b6000611b328a828b0161180c565b9750506020611b438a828b0161180c565b9650506040611b548a828b01611899565b9550506060611b658a828b01611899565b9450506080611b768a828b01611899565b93505060a0611b878a828b01611899565b92505060c0611b988a828b01611899565b91505092959891949750929550565b600081519050611bb681611882565b92915050565b600060208284031215611bd257611bd16117b9565b5b6000611be084828501611ba7565b91505092915050565b600082825260208201905092915050565b7f48617350726f787941646d696e3a20756e617574686f72697a65642073656e6460008201527f6572000000000000000000000000000000000000000000000000000000000000602082015250565b6000611c56602283611be9565b9150611c6182611bfa565b604082019050919050565b60006020820190508181036000830152611c8581611c49565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611cc68261184e565b9150611cd18361184e565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611d0a57611d09611c8c565b5b828202905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611d4f8261184e565b9150611d5a8361184e565b925082611d6a57611d69611d15565b5b828204905092915050565b6000611d808261184e565b9150611d8b8361184e565b9250828201905080821115611da357611da2611c8c565b5b92915050565b6000611db48261184e565b9150611dbf8361184e565b9250828203905081811115611dd757611dd6611c8c565b5b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b611e2b82611de2565b810181811067ffffffffffffffff82111715611e4a57611e49611df3565b5b80604052505050565b6000611e5d6117af565b9050611e698282611e22565b919050565b600060608284031215611e8457611e83611ddd565b5b611e8e6060611e53565b90506000611e9e84828501611ba7565b6000830152506020611eb284828501611ba7565b6020830152506040611ec684828501611ba7565b60408301525092915050565b600060608284031215611ee857611ee76117b9565b5b6000611ef684828501611e6e565b91505092915050565b7f536c617368496e64696361746f723a206d6574686f642063616c6c6572206d7560008201527f737420626520636f696e62617365000000000000000000000000000000000000602082015250565b6000611f5b602e83611be9565b9150611f6682611eff565b604082019050919050565b60006020820190508181036000830152611f8a81611f4e565b9050919050565b6000606082019050611fa66000830186611a09565b611fb36020830185611858565b611fc06040830184611858565b949350505050565b7f536c617368496e64696361746f723a2063616e6e6f7420736c6173682061207660008201527f616c696461746f72207477696365206f7220736c617368206d6f72652074686160208201527f6e206f6e652076616c696461746f7220696e206f6e6520626c6f636b00000000604082015250565b600061204a605c83611be9565b915061205582611fc8565b606082019050919050565b600060208201905081810360008301526120798161203d565b9050919050565b60006040820190506120956000830185611a09565b6120a26020830184611858565b9392505050565b60008115159050919050565b6120be816120a9565b81146120c957600080fd5b50565b6000815190506120db816120b5565b92915050565b6000602082840312156120f7576120f66117b9565b5b6000612105848285016120cc565b91505092915050565b60006121198261184e565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361214b5761214a611c8c565b5b600182019050919050565b600060408201905061216b6000830185611a9b565b6121786020830184611858565b9392505050565b6000819050919050565b6000819050919050565b60006121ae6121a96121a48461217f565b612189565b61184e565b9050919050565b6121be81612193565b82525050565b60006060820190506121d96000830186611a09565b6121e660208301856121b5565b6121f360408301846121b5565b949350505050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b6000612257602e83611be9565b9150612262826121fb565b604082019050919050565b600060208201905081810360008301526122868161224a565b9050919050565b6000819050919050565b600060ff82169050919050565b60006122bf6122ba6122b58461228d565b612189565b612297565b9050919050565b6122cf816122a4565b82525050565b60006020820190506122ea60008301846122c6565b9291505056fea2646970667358221220596bf6931e5e1ac6d2a596b7742c7865bf6a8c1a711329b22248ebcace8f47d064736f6c63430008100033",
}

// SlashIndicatorABI is the input ABI used to generate the binding from.
// Deprecated: Use SlashIndicatorMetaData.ABI instead.
var SlashIndicatorABI = SlashIndicatorMetaData.ABI

// SlashIndicatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SlashIndicatorMetaData.Bin instead.
var SlashIndicatorBin = SlashIndicatorMetaData.Bin

// DeploySlashIndicator deploys a new Ethereum contract, binding an instance of SlashIndicator to it.
func DeploySlashIndicator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SlashIndicator, error) {
	parsed, err := SlashIndicatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SlashIndicatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SlashIndicator{SlashIndicatorCaller: SlashIndicatorCaller{contract: contract}, SlashIndicatorTransactor: SlashIndicatorTransactor{contract: contract}, SlashIndicatorFilterer: SlashIndicatorFilterer{contract: contract}}, nil
}

// SlashIndicator is an auto generated Go binding around an Ethereum contract.
type SlashIndicator struct {
	SlashIndicatorCaller     // Read-only binding to the contract
	SlashIndicatorTransactor // Write-only binding to the contract
	SlashIndicatorFilterer   // Log filterer for contract events
}

// SlashIndicatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type SlashIndicatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashIndicatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SlashIndicatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashIndicatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SlashIndicatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashIndicatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SlashIndicatorSession struct {
	Contract     *SlashIndicator   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SlashIndicatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SlashIndicatorCallerSession struct {
	Contract *SlashIndicatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SlashIndicatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SlashIndicatorTransactorSession struct {
	Contract     *SlashIndicatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SlashIndicatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type SlashIndicatorRaw struct {
	Contract *SlashIndicator // Generic contract binding to access the raw methods on
}

// SlashIndicatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SlashIndicatorCallerRaw struct {
	Contract *SlashIndicatorCaller // Generic read-only contract binding to access the raw methods on
}

// SlashIndicatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SlashIndicatorTransactorRaw struct {
	Contract *SlashIndicatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSlashIndicator creates a new instance of SlashIndicator, bound to a specific deployed contract.
func NewSlashIndicator(address common.Address, backend bind.ContractBackend) (*SlashIndicator, error) {
	contract, err := bindSlashIndicator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SlashIndicator{SlashIndicatorCaller: SlashIndicatorCaller{contract: contract}, SlashIndicatorTransactor: SlashIndicatorTransactor{contract: contract}, SlashIndicatorFilterer: SlashIndicatorFilterer{contract: contract}}, nil
}

// NewSlashIndicatorCaller creates a new read-only instance of SlashIndicator, bound to a specific deployed contract.
func NewSlashIndicatorCaller(address common.Address, caller bind.ContractCaller) (*SlashIndicatorCaller, error) {
	contract, err := bindSlashIndicator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorCaller{contract: contract}, nil
}

// NewSlashIndicatorTransactor creates a new write-only instance of SlashIndicator, bound to a specific deployed contract.
func NewSlashIndicatorTransactor(address common.Address, transactor bind.ContractTransactor) (*SlashIndicatorTransactor, error) {
	contract, err := bindSlashIndicator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorTransactor{contract: contract}, nil
}

// NewSlashIndicatorFilterer creates a new log filterer instance of SlashIndicator, bound to a specific deployed contract.
func NewSlashIndicatorFilterer(address common.Address, filterer bind.ContractFilterer) (*SlashIndicatorFilterer, error) {
	contract, err := bindSlashIndicator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorFilterer{contract: contract}, nil
}

// bindSlashIndicator binds a generic wrapper to an already deployed contract.
func bindSlashIndicator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SlashIndicatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlashIndicator *SlashIndicatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlashIndicator.Contract.SlashIndicatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlashIndicator *SlashIndicatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SlashIndicatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlashIndicator *SlashIndicatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SlashIndicatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlashIndicator *SlashIndicatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlashIndicator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlashIndicator *SlashIndicatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlashIndicator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlashIndicator *SlashIndicatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlashIndicator.Contract.contract.Transact(opts, method, params...)
}

// CurrentUnavailabilityIndicator is a free data retrieval call binding the contract method 0x082e7420.
//
// Solidity: function currentUnavailabilityIndicator(address _validator) view returns(uint256)
func (_SlashIndicator *SlashIndicatorCaller) CurrentUnavailabilityIndicator(opts *bind.CallOpts, _validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "currentUnavailabilityIndicator", _validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentUnavailabilityIndicator is a free data retrieval call binding the contract method 0x082e7420.
//
// Solidity: function currentUnavailabilityIndicator(address _validator) view returns(uint256)
func (_SlashIndicator *SlashIndicatorSession) CurrentUnavailabilityIndicator(_validator common.Address) (*big.Int, error) {
	return _SlashIndicator.Contract.CurrentUnavailabilityIndicator(&_SlashIndicator.CallOpts, _validator)
}

// CurrentUnavailabilityIndicator is a free data retrieval call binding the contract method 0x082e7420.
//
// Solidity: function currentUnavailabilityIndicator(address _validator) view returns(uint256)
func (_SlashIndicator *SlashIndicatorCallerSession) CurrentUnavailabilityIndicator(_validator common.Address) (*big.Int, error) {
	return _SlashIndicator.Contract.CurrentUnavailabilityIndicator(&_SlashIndicator.CallOpts, _validator)
}

// FelonyJailDuration is a free data retrieval call binding the contract method 0x6b79f95f.
//
// Solidity: function felonyJailDuration() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCaller) FelonyJailDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "felonyJailDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FelonyJailDuration is a free data retrieval call binding the contract method 0x6b79f95f.
//
// Solidity: function felonyJailDuration() view returns(uint256)
func (_SlashIndicator *SlashIndicatorSession) FelonyJailDuration() (*big.Int, error) {
	return _SlashIndicator.Contract.FelonyJailDuration(&_SlashIndicator.CallOpts)
}

// FelonyJailDuration is a free data retrieval call binding the contract method 0x6b79f95f.
//
// Solidity: function felonyJailDuration() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCallerSession) FelonyJailDuration() (*big.Int, error) {
	return _SlashIndicator.Contract.FelonyJailDuration(&_SlashIndicator.CallOpts)
}

// FelonyThreshold is a free data retrieval call binding the contract method 0x389f4f71.
//
// Solidity: function felonyThreshold() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCaller) FelonyThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "felonyThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FelonyThreshold is a free data retrieval call binding the contract method 0x389f4f71.
//
// Solidity: function felonyThreshold() view returns(uint256)
func (_SlashIndicator *SlashIndicatorSession) FelonyThreshold() (*big.Int, error) {
	return _SlashIndicator.Contract.FelonyThreshold(&_SlashIndicator.CallOpts)
}

// FelonyThreshold is a free data retrieval call binding the contract method 0x389f4f71.
//
// Solidity: function felonyThreshold() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCallerSession) FelonyThreshold() (*big.Int, error) {
	return _SlashIndicator.Contract.FelonyThreshold(&_SlashIndicator.CallOpts)
}

// GetUnavailabilityIndicator is a free data retrieval call binding the contract method 0x62ffe6cb.
//
// Solidity: function getUnavailabilityIndicator(address _validator, uint256 _period) view returns(uint256)
func (_SlashIndicator *SlashIndicatorCaller) GetUnavailabilityIndicator(opts *bind.CallOpts, _validator common.Address, _period *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "getUnavailabilityIndicator", _validator, _period)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnavailabilityIndicator is a free data retrieval call binding the contract method 0x62ffe6cb.
//
// Solidity: function getUnavailabilityIndicator(address _validator, uint256 _period) view returns(uint256)
func (_SlashIndicator *SlashIndicatorSession) GetUnavailabilityIndicator(_validator common.Address, _period *big.Int) (*big.Int, error) {
	return _SlashIndicator.Contract.GetUnavailabilityIndicator(&_SlashIndicator.CallOpts, _validator, _period)
}

// GetUnavailabilityIndicator is a free data retrieval call binding the contract method 0x62ffe6cb.
//
// Solidity: function getUnavailabilityIndicator(address _validator, uint256 _period) view returns(uint256)
func (_SlashIndicator *SlashIndicatorCallerSession) GetUnavailabilityIndicator(_validator common.Address, _period *big.Int) (*big.Int, error) {
	return _SlashIndicator.Contract.GetUnavailabilityIndicator(&_SlashIndicator.CallOpts, _validator, _period)
}

// GetUnavailabilitySlashType is a free data retrieval call binding the contract method 0xd180ecb0.
//
// Solidity: function getUnavailabilitySlashType(address _validatorAddr, uint256 _period) view returns(uint8)
func (_SlashIndicator *SlashIndicatorCaller) GetUnavailabilitySlashType(opts *bind.CallOpts, _validatorAddr common.Address, _period *big.Int) (uint8, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "getUnavailabilitySlashType", _validatorAddr, _period)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetUnavailabilitySlashType is a free data retrieval call binding the contract method 0xd180ecb0.
//
// Solidity: function getUnavailabilitySlashType(address _validatorAddr, uint256 _period) view returns(uint8)
func (_SlashIndicator *SlashIndicatorSession) GetUnavailabilitySlashType(_validatorAddr common.Address, _period *big.Int) (uint8, error) {
	return _SlashIndicator.Contract.GetUnavailabilitySlashType(&_SlashIndicator.CallOpts, _validatorAddr, _period)
}

// GetUnavailabilitySlashType is a free data retrieval call binding the contract method 0xd180ecb0.
//
// Solidity: function getUnavailabilitySlashType(address _validatorAddr, uint256 _period) view returns(uint8)
func (_SlashIndicator *SlashIndicatorCallerSession) GetUnavailabilitySlashType(_validatorAddr common.Address, _period *big.Int) (uint8, error) {
	return _SlashIndicator.Contract.GetUnavailabilitySlashType(&_SlashIndicator.CallOpts, _validatorAddr, _period)
}

// GetUnavailabilityThresholds is a free data retrieval call binding the contract method 0xd9193587.
//
// Solidity: function getUnavailabilityThresholds() view returns(uint256, uint256)
func (_SlashIndicator *SlashIndicatorCaller) GetUnavailabilityThresholds(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "getUnavailabilityThresholds")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetUnavailabilityThresholds is a free data retrieval call binding the contract method 0xd9193587.
//
// Solidity: function getUnavailabilityThresholds() view returns(uint256, uint256)
func (_SlashIndicator *SlashIndicatorSession) GetUnavailabilityThresholds() (*big.Int, *big.Int, error) {
	return _SlashIndicator.Contract.GetUnavailabilityThresholds(&_SlashIndicator.CallOpts)
}

// GetUnavailabilityThresholds is a free data retrieval call binding the contract method 0xd9193587.
//
// Solidity: function getUnavailabilityThresholds() view returns(uint256, uint256)
func (_SlashIndicator *SlashIndicatorCallerSession) GetUnavailabilityThresholds() (*big.Int, *big.Int, error) {
	return _SlashIndicator.Contract.GetUnavailabilityThresholds(&_SlashIndicator.CallOpts)
}

// LastSlashedBlock is a free data retrieval call binding the contract method 0xaef250be.
//
// Solidity: function lastSlashedBlock() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCaller) LastSlashedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "lastSlashedBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastSlashedBlock is a free data retrieval call binding the contract method 0xaef250be.
//
// Solidity: function lastSlashedBlock() view returns(uint256)
func (_SlashIndicator *SlashIndicatorSession) LastSlashedBlock() (*big.Int, error) {
	return _SlashIndicator.Contract.LastSlashedBlock(&_SlashIndicator.CallOpts)
}

// LastSlashedBlock is a free data retrieval call binding the contract method 0xaef250be.
//
// Solidity: function lastSlashedBlock() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCallerSession) LastSlashedBlock() (*big.Int, error) {
	return _SlashIndicator.Contract.LastSlashedBlock(&_SlashIndicator.CallOpts)
}

// MaintenanceContract is a free data retrieval call binding the contract method 0xd2cb215e.
//
// Solidity: function maintenanceContract() view returns(address)
func (_SlashIndicator *SlashIndicatorCaller) MaintenanceContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "maintenanceContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MaintenanceContract is a free data retrieval call binding the contract method 0xd2cb215e.
//
// Solidity: function maintenanceContract() view returns(address)
func (_SlashIndicator *SlashIndicatorSession) MaintenanceContract() (common.Address, error) {
	return _SlashIndicator.Contract.MaintenanceContract(&_SlashIndicator.CallOpts)
}

// MaintenanceContract is a free data retrieval call binding the contract method 0xd2cb215e.
//
// Solidity: function maintenanceContract() view returns(address)
func (_SlashIndicator *SlashIndicatorCallerSession) MaintenanceContract() (common.Address, error) {
	return _SlashIndicator.Contract.MaintenanceContract(&_SlashIndicator.CallOpts)
}

// MisdemeanorThreshold is a free data retrieval call binding the contract method 0x567a372d.
//
// Solidity: function misdemeanorThreshold() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCaller) MisdemeanorThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "misdemeanorThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MisdemeanorThreshold is a free data retrieval call binding the contract method 0x567a372d.
//
// Solidity: function misdemeanorThreshold() view returns(uint256)
func (_SlashIndicator *SlashIndicatorSession) MisdemeanorThreshold() (*big.Int, error) {
	return _SlashIndicator.Contract.MisdemeanorThreshold(&_SlashIndicator.CallOpts)
}

// MisdemeanorThreshold is a free data retrieval call binding the contract method 0x567a372d.
//
// Solidity: function misdemeanorThreshold() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCallerSession) MisdemeanorThreshold() (*big.Int, error) {
	return _SlashIndicator.Contract.MisdemeanorThreshold(&_SlashIndicator.CallOpts)
}

// SlashDoubleSignAmount is a free data retrieval call binding the contract method 0x6d14c4e5.
//
// Solidity: function slashDoubleSignAmount() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCaller) SlashDoubleSignAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "slashDoubleSignAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashDoubleSignAmount is a free data retrieval call binding the contract method 0x6d14c4e5.
//
// Solidity: function slashDoubleSignAmount() view returns(uint256)
func (_SlashIndicator *SlashIndicatorSession) SlashDoubleSignAmount() (*big.Int, error) {
	return _SlashIndicator.Contract.SlashDoubleSignAmount(&_SlashIndicator.CallOpts)
}

// SlashDoubleSignAmount is a free data retrieval call binding the contract method 0x6d14c4e5.
//
// Solidity: function slashDoubleSignAmount() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCallerSession) SlashDoubleSignAmount() (*big.Int, error) {
	return _SlashIndicator.Contract.SlashDoubleSignAmount(&_SlashIndicator.CallOpts)
}

// SlashFelonyAmount is a free data retrieval call binding the contract method 0xdfe484cb.
//
// Solidity: function slashFelonyAmount() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCaller) SlashFelonyAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "slashFelonyAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashFelonyAmount is a free data retrieval call binding the contract method 0xdfe484cb.
//
// Solidity: function slashFelonyAmount() view returns(uint256)
func (_SlashIndicator *SlashIndicatorSession) SlashFelonyAmount() (*big.Int, error) {
	return _SlashIndicator.Contract.SlashFelonyAmount(&_SlashIndicator.CallOpts)
}

// SlashFelonyAmount is a free data retrieval call binding the contract method 0xdfe484cb.
//
// Solidity: function slashFelonyAmount() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCallerSession) SlashFelonyAmount() (*big.Int, error) {
	return _SlashIndicator.Contract.SlashFelonyAmount(&_SlashIndicator.CallOpts)
}

// UnavailabilityThresholdsOf is a free data retrieval call binding the contract method 0x4d961e18.
//
// Solidity: function unavailabilityThresholdsOf(address _addr, uint256 _block) view returns(uint256 _misdemeanorThreshold, uint256 _felonyThreshold)
func (_SlashIndicator *SlashIndicatorCaller) UnavailabilityThresholdsOf(opts *bind.CallOpts, _addr common.Address, _block *big.Int) (struct {
	MisdemeanorThreshold *big.Int
	FelonyThreshold      *big.Int
}, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "unavailabilityThresholdsOf", _addr, _block)

	outstruct := new(struct {
		MisdemeanorThreshold *big.Int
		FelonyThreshold      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MisdemeanorThreshold = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FelonyThreshold = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UnavailabilityThresholdsOf is a free data retrieval call binding the contract method 0x4d961e18.
//
// Solidity: function unavailabilityThresholdsOf(address _addr, uint256 _block) view returns(uint256 _misdemeanorThreshold, uint256 _felonyThreshold)
func (_SlashIndicator *SlashIndicatorSession) UnavailabilityThresholdsOf(_addr common.Address, _block *big.Int) (struct {
	MisdemeanorThreshold *big.Int
	FelonyThreshold      *big.Int
}, error) {
	return _SlashIndicator.Contract.UnavailabilityThresholdsOf(&_SlashIndicator.CallOpts, _addr, _block)
}

// UnavailabilityThresholdsOf is a free data retrieval call binding the contract method 0x4d961e18.
//
// Solidity: function unavailabilityThresholdsOf(address _addr, uint256 _block) view returns(uint256 _misdemeanorThreshold, uint256 _felonyThreshold)
func (_SlashIndicator *SlashIndicatorCallerSession) UnavailabilityThresholdsOf(_addr common.Address, _block *big.Int) (struct {
	MisdemeanorThreshold *big.Int
	FelonyThreshold      *big.Int
}, error) {
	return _SlashIndicator.Contract.UnavailabilityThresholdsOf(&_SlashIndicator.CallOpts, _addr, _block)
}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_SlashIndicator *SlashIndicatorCaller) ValidatorContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "validatorContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_SlashIndicator *SlashIndicatorSession) ValidatorContract() (common.Address, error) {
	return _SlashIndicator.Contract.ValidatorContract(&_SlashIndicator.CallOpts)
}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_SlashIndicator *SlashIndicatorCallerSession) ValidatorContract() (common.Address, error) {
	return _SlashIndicator.Contract.ValidatorContract(&_SlashIndicator.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xedbf4ac2.
//
// Solidity: function initialize(address __validatorContract, address __maintenanceContract, uint256 _misdemeanorThreshold, uint256 _felonyThreshold, uint256 _slashFelonyAmount, uint256 _slashDoubleSignAmount, uint256 _felonyJailBlocks) returns()
func (_SlashIndicator *SlashIndicatorTransactor) Initialize(opts *bind.TransactOpts, __validatorContract common.Address, __maintenanceContract common.Address, _misdemeanorThreshold *big.Int, _felonyThreshold *big.Int, _slashFelonyAmount *big.Int, _slashDoubleSignAmount *big.Int, _felonyJailBlocks *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.contract.Transact(opts, "initialize", __validatorContract, __maintenanceContract, _misdemeanorThreshold, _felonyThreshold, _slashFelonyAmount, _slashDoubleSignAmount, _felonyJailBlocks)
}

// Initialize is a paid mutator transaction binding the contract method 0xedbf4ac2.
//
// Solidity: function initialize(address __validatorContract, address __maintenanceContract, uint256 _misdemeanorThreshold, uint256 _felonyThreshold, uint256 _slashFelonyAmount, uint256 _slashDoubleSignAmount, uint256 _felonyJailBlocks) returns()
func (_SlashIndicator *SlashIndicatorSession) Initialize(__validatorContract common.Address, __maintenanceContract common.Address, _misdemeanorThreshold *big.Int, _felonyThreshold *big.Int, _slashFelonyAmount *big.Int, _slashDoubleSignAmount *big.Int, _felonyJailBlocks *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.Contract.Initialize(&_SlashIndicator.TransactOpts, __validatorContract, __maintenanceContract, _misdemeanorThreshold, _felonyThreshold, _slashFelonyAmount, _slashDoubleSignAmount, _felonyJailBlocks)
}

// Initialize is a paid mutator transaction binding the contract method 0xedbf4ac2.
//
// Solidity: function initialize(address __validatorContract, address __maintenanceContract, uint256 _misdemeanorThreshold, uint256 _felonyThreshold, uint256 _slashFelonyAmount, uint256 _slashDoubleSignAmount, uint256 _felonyJailBlocks) returns()
func (_SlashIndicator *SlashIndicatorTransactorSession) Initialize(__validatorContract common.Address, __maintenanceContract common.Address, _misdemeanorThreshold *big.Int, _felonyThreshold *big.Int, _slashFelonyAmount *big.Int, _slashDoubleSignAmount *big.Int, _felonyJailBlocks *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.Contract.Initialize(&_SlashIndicator.TransactOpts, __validatorContract, __maintenanceContract, _misdemeanorThreshold, _felonyThreshold, _slashFelonyAmount, _slashDoubleSignAmount, _felonyJailBlocks)
}

// SetFelonyJailDuration is a paid mutator transaction binding the contract method 0x5786e9ee.
//
// Solidity: function setFelonyJailDuration(uint256 _felonyJailDuration) returns()
func (_SlashIndicator *SlashIndicatorTransactor) SetFelonyJailDuration(opts *bind.TransactOpts, _felonyJailDuration *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.contract.Transact(opts, "setFelonyJailDuration", _felonyJailDuration)
}

// SetFelonyJailDuration is a paid mutator transaction binding the contract method 0x5786e9ee.
//
// Solidity: function setFelonyJailDuration(uint256 _felonyJailDuration) returns()
func (_SlashIndicator *SlashIndicatorSession) SetFelonyJailDuration(_felonyJailDuration *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetFelonyJailDuration(&_SlashIndicator.TransactOpts, _felonyJailDuration)
}

// SetFelonyJailDuration is a paid mutator transaction binding the contract method 0x5786e9ee.
//
// Solidity: function setFelonyJailDuration(uint256 _felonyJailDuration) returns()
func (_SlashIndicator *SlashIndicatorTransactorSession) SetFelonyJailDuration(_felonyJailDuration *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetFelonyJailDuration(&_SlashIndicator.TransactOpts, _felonyJailDuration)
}

// SetMaintenanceContract is a paid mutator transaction binding the contract method 0x46fe9311.
//
// Solidity: function setMaintenanceContract(address _addr) returns()
func (_SlashIndicator *SlashIndicatorTransactor) SetMaintenanceContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.contract.Transact(opts, "setMaintenanceContract", _addr)
}

// SetMaintenanceContract is a paid mutator transaction binding the contract method 0x46fe9311.
//
// Solidity: function setMaintenanceContract(address _addr) returns()
func (_SlashIndicator *SlashIndicatorSession) SetMaintenanceContract(_addr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetMaintenanceContract(&_SlashIndicator.TransactOpts, _addr)
}

// SetMaintenanceContract is a paid mutator transaction binding the contract method 0x46fe9311.
//
// Solidity: function setMaintenanceContract(address _addr) returns()
func (_SlashIndicator *SlashIndicatorTransactorSession) SetMaintenanceContract(_addr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetMaintenanceContract(&_SlashIndicator.TransactOpts, _addr)
}

// SetSlashDoubleSignAmount is a paid mutator transaction binding the contract method 0xf82bdd01.
//
// Solidity: function setSlashDoubleSignAmount(uint256 _slashDoubleSignAmount) returns()
func (_SlashIndicator *SlashIndicatorTransactor) SetSlashDoubleSignAmount(opts *bind.TransactOpts, _slashDoubleSignAmount *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.contract.Transact(opts, "setSlashDoubleSignAmount", _slashDoubleSignAmount)
}

// SetSlashDoubleSignAmount is a paid mutator transaction binding the contract method 0xf82bdd01.
//
// Solidity: function setSlashDoubleSignAmount(uint256 _slashDoubleSignAmount) returns()
func (_SlashIndicator *SlashIndicatorSession) SetSlashDoubleSignAmount(_slashDoubleSignAmount *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetSlashDoubleSignAmount(&_SlashIndicator.TransactOpts, _slashDoubleSignAmount)
}

// SetSlashDoubleSignAmount is a paid mutator transaction binding the contract method 0xf82bdd01.
//
// Solidity: function setSlashDoubleSignAmount(uint256 _slashDoubleSignAmount) returns()
func (_SlashIndicator *SlashIndicatorTransactorSession) SetSlashDoubleSignAmount(_slashDoubleSignAmount *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetSlashDoubleSignAmount(&_SlashIndicator.TransactOpts, _slashDoubleSignAmount)
}

// SetSlashFelonyAmount is a paid mutator transaction binding the contract method 0x6e91bec5.
//
// Solidity: function setSlashFelonyAmount(uint256 _slashFelonyAmount) returns()
func (_SlashIndicator *SlashIndicatorTransactor) SetSlashFelonyAmount(opts *bind.TransactOpts, _slashFelonyAmount *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.contract.Transact(opts, "setSlashFelonyAmount", _slashFelonyAmount)
}

// SetSlashFelonyAmount is a paid mutator transaction binding the contract method 0x6e91bec5.
//
// Solidity: function setSlashFelonyAmount(uint256 _slashFelonyAmount) returns()
func (_SlashIndicator *SlashIndicatorSession) SetSlashFelonyAmount(_slashFelonyAmount *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetSlashFelonyAmount(&_SlashIndicator.TransactOpts, _slashFelonyAmount)
}

// SetSlashFelonyAmount is a paid mutator transaction binding the contract method 0x6e91bec5.
//
// Solidity: function setSlashFelonyAmount(uint256 _slashFelonyAmount) returns()
func (_SlashIndicator *SlashIndicatorTransactorSession) SetSlashFelonyAmount(_slashFelonyAmount *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetSlashFelonyAmount(&_SlashIndicator.TransactOpts, _slashFelonyAmount)
}

// SetSlashThresholds is a paid mutator transaction binding the contract method 0xe4f9c4be.
//
// Solidity: function setSlashThresholds(uint256 _felonyThreshold, uint256 _misdemeanorThreshold) returns()
func (_SlashIndicator *SlashIndicatorTransactor) SetSlashThresholds(opts *bind.TransactOpts, _felonyThreshold *big.Int, _misdemeanorThreshold *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.contract.Transact(opts, "setSlashThresholds", _felonyThreshold, _misdemeanorThreshold)
}

// SetSlashThresholds is a paid mutator transaction binding the contract method 0xe4f9c4be.
//
// Solidity: function setSlashThresholds(uint256 _felonyThreshold, uint256 _misdemeanorThreshold) returns()
func (_SlashIndicator *SlashIndicatorSession) SetSlashThresholds(_felonyThreshold *big.Int, _misdemeanorThreshold *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetSlashThresholds(&_SlashIndicator.TransactOpts, _felonyThreshold, _misdemeanorThreshold)
}

// SetSlashThresholds is a paid mutator transaction binding the contract method 0xe4f9c4be.
//
// Solidity: function setSlashThresholds(uint256 _felonyThreshold, uint256 _misdemeanorThreshold) returns()
func (_SlashIndicator *SlashIndicatorTransactorSession) SetSlashThresholds(_felonyThreshold *big.Int, _misdemeanorThreshold *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetSlashThresholds(&_SlashIndicator.TransactOpts, _felonyThreshold, _misdemeanorThreshold)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _addr) returns()
func (_SlashIndicator *SlashIndicatorTransactor) SetValidatorContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.contract.Transact(opts, "setValidatorContract", _addr)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _addr) returns()
func (_SlashIndicator *SlashIndicatorSession) SetValidatorContract(_addr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetValidatorContract(&_SlashIndicator.TransactOpts, _addr)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _addr) returns()
func (_SlashIndicator *SlashIndicatorTransactorSession) SetValidatorContract(_addr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SetValidatorContract(&_SlashIndicator.TransactOpts, _addr)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address _validatorAddr) returns()
func (_SlashIndicator *SlashIndicatorTransactor) Slash(opts *bind.TransactOpts, _validatorAddr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.contract.Transact(opts, "slash", _validatorAddr)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address _validatorAddr) returns()
func (_SlashIndicator *SlashIndicatorSession) Slash(_validatorAddr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.Contract.Slash(&_SlashIndicator.TransactOpts, _validatorAddr)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address _validatorAddr) returns()
func (_SlashIndicator *SlashIndicatorTransactorSession) Slash(_validatorAddr common.Address) (*types.Transaction, error) {
	return _SlashIndicator.Contract.Slash(&_SlashIndicator.TransactOpts, _validatorAddr)
}

// SlashDoubleSign is a paid mutator transaction binding the contract method 0x518e463a.
//
// Solidity: function slashDoubleSign(address _validatorAddr, bytes ) returns()
func (_SlashIndicator *SlashIndicatorTransactor) SlashDoubleSign(opts *bind.TransactOpts, _validatorAddr common.Address, arg1 []byte) (*types.Transaction, error) {
	return _SlashIndicator.contract.Transact(opts, "slashDoubleSign", _validatorAddr, arg1)
}

// SlashDoubleSign is a paid mutator transaction binding the contract method 0x518e463a.
//
// Solidity: function slashDoubleSign(address _validatorAddr, bytes ) returns()
func (_SlashIndicator *SlashIndicatorSession) SlashDoubleSign(_validatorAddr common.Address, arg1 []byte) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SlashDoubleSign(&_SlashIndicator.TransactOpts, _validatorAddr, arg1)
}

// SlashDoubleSign is a paid mutator transaction binding the contract method 0x518e463a.
//
// Solidity: function slashDoubleSign(address _validatorAddr, bytes ) returns()
func (_SlashIndicator *SlashIndicatorTransactorSession) SlashDoubleSign(_validatorAddr common.Address, arg1 []byte) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SlashDoubleSign(&_SlashIndicator.TransactOpts, _validatorAddr, arg1)
}

// SlashIndicatorFelonyJailDurationUpdatedIterator is returned from FilterFelonyJailDurationUpdated and is used to iterate over the raw logs and unpacked data for FelonyJailDurationUpdated events raised by the SlashIndicator contract.
type SlashIndicatorFelonyJailDurationUpdatedIterator struct {
	Event *SlashIndicatorFelonyJailDurationUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SlashIndicatorFelonyJailDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashIndicatorFelonyJailDurationUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SlashIndicatorFelonyJailDurationUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SlashIndicatorFelonyJailDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashIndicatorFelonyJailDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashIndicatorFelonyJailDurationUpdated represents a FelonyJailDurationUpdated event raised by the SlashIndicator contract.
type SlashIndicatorFelonyJailDurationUpdated struct {
	FelonyJailDuration *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterFelonyJailDurationUpdated is a free log retrieval operation binding the contract event 0x3092a623ebcf71b79f9f68801b081a4d0c839dfb4c0e6f9ff118f4fb870375dc.
//
// Solidity: event FelonyJailDurationUpdated(uint256 felonyJailDuration)
func (_SlashIndicator *SlashIndicatorFilterer) FilterFelonyJailDurationUpdated(opts *bind.FilterOpts) (*SlashIndicatorFelonyJailDurationUpdatedIterator, error) {

	logs, sub, err := _SlashIndicator.contract.FilterLogs(opts, "FelonyJailDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorFelonyJailDurationUpdatedIterator{contract: _SlashIndicator.contract, event: "FelonyJailDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchFelonyJailDurationUpdated is a free log subscription operation binding the contract event 0x3092a623ebcf71b79f9f68801b081a4d0c839dfb4c0e6f9ff118f4fb870375dc.
//
// Solidity: event FelonyJailDurationUpdated(uint256 felonyJailDuration)
func (_SlashIndicator *SlashIndicatorFilterer) WatchFelonyJailDurationUpdated(opts *bind.WatchOpts, sink chan<- *SlashIndicatorFelonyJailDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _SlashIndicator.contract.WatchLogs(opts, "FelonyJailDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashIndicatorFelonyJailDurationUpdated)
				if err := _SlashIndicator.contract.UnpackLog(event, "FelonyJailDurationUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFelonyJailDurationUpdated is a log parse operation binding the contract event 0x3092a623ebcf71b79f9f68801b081a4d0c839dfb4c0e6f9ff118f4fb870375dc.
//
// Solidity: event FelonyJailDurationUpdated(uint256 felonyJailDuration)
func (_SlashIndicator *SlashIndicatorFilterer) ParseFelonyJailDurationUpdated(log types.Log) (*SlashIndicatorFelonyJailDurationUpdated, error) {
	event := new(SlashIndicatorFelonyJailDurationUpdated)
	if err := _SlashIndicator.contract.UnpackLog(event, "FelonyJailDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashIndicatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SlashIndicator contract.
type SlashIndicatorInitializedIterator struct {
	Event *SlashIndicatorInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SlashIndicatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashIndicatorInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SlashIndicatorInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SlashIndicatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashIndicatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashIndicatorInitialized represents a Initialized event raised by the SlashIndicator contract.
type SlashIndicatorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SlashIndicator *SlashIndicatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*SlashIndicatorInitializedIterator, error) {

	logs, sub, err := _SlashIndicator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorInitializedIterator{contract: _SlashIndicator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SlashIndicator *SlashIndicatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SlashIndicatorInitialized) (event.Subscription, error) {

	logs, sub, err := _SlashIndicator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashIndicatorInitialized)
				if err := _SlashIndicator.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SlashIndicator *SlashIndicatorFilterer) ParseInitialized(log types.Log) (*SlashIndicatorInitialized, error) {
	event := new(SlashIndicatorInitialized)
	if err := _SlashIndicator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashIndicatorMaintenanceContractUpdatedIterator is returned from FilterMaintenanceContractUpdated and is used to iterate over the raw logs and unpacked data for MaintenanceContractUpdated events raised by the SlashIndicator contract.
type SlashIndicatorMaintenanceContractUpdatedIterator struct {
	Event *SlashIndicatorMaintenanceContractUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SlashIndicatorMaintenanceContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashIndicatorMaintenanceContractUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SlashIndicatorMaintenanceContractUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SlashIndicatorMaintenanceContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashIndicatorMaintenanceContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashIndicatorMaintenanceContractUpdated represents a MaintenanceContractUpdated event raised by the SlashIndicator contract.
type SlashIndicatorMaintenanceContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMaintenanceContractUpdated is a free log retrieval operation binding the contract event 0x31a33f126a5bae3c5bdf6cfc2cd6dcfffe2fe9634bdb09e21c44762993889e3b.
//
// Solidity: event MaintenanceContractUpdated(address arg0)
func (_SlashIndicator *SlashIndicatorFilterer) FilterMaintenanceContractUpdated(opts *bind.FilterOpts) (*SlashIndicatorMaintenanceContractUpdatedIterator, error) {

	logs, sub, err := _SlashIndicator.contract.FilterLogs(opts, "MaintenanceContractUpdated")
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorMaintenanceContractUpdatedIterator{contract: _SlashIndicator.contract, event: "MaintenanceContractUpdated", logs: logs, sub: sub}, nil
}

// WatchMaintenanceContractUpdated is a free log subscription operation binding the contract event 0x31a33f126a5bae3c5bdf6cfc2cd6dcfffe2fe9634bdb09e21c44762993889e3b.
//
// Solidity: event MaintenanceContractUpdated(address arg0)
func (_SlashIndicator *SlashIndicatorFilterer) WatchMaintenanceContractUpdated(opts *bind.WatchOpts, sink chan<- *SlashIndicatorMaintenanceContractUpdated) (event.Subscription, error) {

	logs, sub, err := _SlashIndicator.contract.WatchLogs(opts, "MaintenanceContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashIndicatorMaintenanceContractUpdated)
				if err := _SlashIndicator.contract.UnpackLog(event, "MaintenanceContractUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMaintenanceContractUpdated is a log parse operation binding the contract event 0x31a33f126a5bae3c5bdf6cfc2cd6dcfffe2fe9634bdb09e21c44762993889e3b.
//
// Solidity: event MaintenanceContractUpdated(address arg0)
func (_SlashIndicator *SlashIndicatorFilterer) ParseMaintenanceContractUpdated(log types.Log) (*SlashIndicatorMaintenanceContractUpdated, error) {
	event := new(SlashIndicatorMaintenanceContractUpdated)
	if err := _SlashIndicator.contract.UnpackLog(event, "MaintenanceContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashIndicatorSlashDoubleSignAmountUpdatedIterator is returned from FilterSlashDoubleSignAmountUpdated and is used to iterate over the raw logs and unpacked data for SlashDoubleSignAmountUpdated events raised by the SlashIndicator contract.
type SlashIndicatorSlashDoubleSignAmountUpdatedIterator struct {
	Event *SlashIndicatorSlashDoubleSignAmountUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SlashIndicatorSlashDoubleSignAmountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashIndicatorSlashDoubleSignAmountUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SlashIndicatorSlashDoubleSignAmountUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SlashIndicatorSlashDoubleSignAmountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashIndicatorSlashDoubleSignAmountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashIndicatorSlashDoubleSignAmountUpdated represents a SlashDoubleSignAmountUpdated event raised by the SlashIndicator contract.
type SlashIndicatorSlashDoubleSignAmountUpdated struct {
	SlashDoubleSignAmount *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSlashDoubleSignAmountUpdated is a free log retrieval operation binding the contract event 0x42458d43c99c17f034dde22f3a1cf003e4bc27372c1ff14cb9e68a79a1b1e8ed.
//
// Solidity: event SlashDoubleSignAmountUpdated(uint256 slashDoubleSignAmount)
func (_SlashIndicator *SlashIndicatorFilterer) FilterSlashDoubleSignAmountUpdated(opts *bind.FilterOpts) (*SlashIndicatorSlashDoubleSignAmountUpdatedIterator, error) {

	logs, sub, err := _SlashIndicator.contract.FilterLogs(opts, "SlashDoubleSignAmountUpdated")
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorSlashDoubleSignAmountUpdatedIterator{contract: _SlashIndicator.contract, event: "SlashDoubleSignAmountUpdated", logs: logs, sub: sub}, nil
}

// WatchSlashDoubleSignAmountUpdated is a free log subscription operation binding the contract event 0x42458d43c99c17f034dde22f3a1cf003e4bc27372c1ff14cb9e68a79a1b1e8ed.
//
// Solidity: event SlashDoubleSignAmountUpdated(uint256 slashDoubleSignAmount)
func (_SlashIndicator *SlashIndicatorFilterer) WatchSlashDoubleSignAmountUpdated(opts *bind.WatchOpts, sink chan<- *SlashIndicatorSlashDoubleSignAmountUpdated) (event.Subscription, error) {

	logs, sub, err := _SlashIndicator.contract.WatchLogs(opts, "SlashDoubleSignAmountUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashIndicatorSlashDoubleSignAmountUpdated)
				if err := _SlashIndicator.contract.UnpackLog(event, "SlashDoubleSignAmountUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSlashDoubleSignAmountUpdated is a log parse operation binding the contract event 0x42458d43c99c17f034dde22f3a1cf003e4bc27372c1ff14cb9e68a79a1b1e8ed.
//
// Solidity: event SlashDoubleSignAmountUpdated(uint256 slashDoubleSignAmount)
func (_SlashIndicator *SlashIndicatorFilterer) ParseSlashDoubleSignAmountUpdated(log types.Log) (*SlashIndicatorSlashDoubleSignAmountUpdated, error) {
	event := new(SlashIndicatorSlashDoubleSignAmountUpdated)
	if err := _SlashIndicator.contract.UnpackLog(event, "SlashDoubleSignAmountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashIndicatorSlashFelonyAmountUpdatedIterator is returned from FilterSlashFelonyAmountUpdated and is used to iterate over the raw logs and unpacked data for SlashFelonyAmountUpdated events raised by the SlashIndicator contract.
type SlashIndicatorSlashFelonyAmountUpdatedIterator struct {
	Event *SlashIndicatorSlashFelonyAmountUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SlashIndicatorSlashFelonyAmountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashIndicatorSlashFelonyAmountUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SlashIndicatorSlashFelonyAmountUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SlashIndicatorSlashFelonyAmountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashIndicatorSlashFelonyAmountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashIndicatorSlashFelonyAmountUpdated represents a SlashFelonyAmountUpdated event raised by the SlashIndicator contract.
type SlashIndicatorSlashFelonyAmountUpdated struct {
	SlashFelonyAmount *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSlashFelonyAmountUpdated is a free log retrieval operation binding the contract event 0x71491762d43dbadd94f85cf7a0322192a7c22f3ca4fb30de8895420607b56712.
//
// Solidity: event SlashFelonyAmountUpdated(uint256 slashFelonyAmount)
func (_SlashIndicator *SlashIndicatorFilterer) FilterSlashFelonyAmountUpdated(opts *bind.FilterOpts) (*SlashIndicatorSlashFelonyAmountUpdatedIterator, error) {

	logs, sub, err := _SlashIndicator.contract.FilterLogs(opts, "SlashFelonyAmountUpdated")
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorSlashFelonyAmountUpdatedIterator{contract: _SlashIndicator.contract, event: "SlashFelonyAmountUpdated", logs: logs, sub: sub}, nil
}

// WatchSlashFelonyAmountUpdated is a free log subscription operation binding the contract event 0x71491762d43dbadd94f85cf7a0322192a7c22f3ca4fb30de8895420607b56712.
//
// Solidity: event SlashFelonyAmountUpdated(uint256 slashFelonyAmount)
func (_SlashIndicator *SlashIndicatorFilterer) WatchSlashFelonyAmountUpdated(opts *bind.WatchOpts, sink chan<- *SlashIndicatorSlashFelonyAmountUpdated) (event.Subscription, error) {

	logs, sub, err := _SlashIndicator.contract.WatchLogs(opts, "SlashFelonyAmountUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashIndicatorSlashFelonyAmountUpdated)
				if err := _SlashIndicator.contract.UnpackLog(event, "SlashFelonyAmountUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSlashFelonyAmountUpdated is a log parse operation binding the contract event 0x71491762d43dbadd94f85cf7a0322192a7c22f3ca4fb30de8895420607b56712.
//
// Solidity: event SlashFelonyAmountUpdated(uint256 slashFelonyAmount)
func (_SlashIndicator *SlashIndicatorFilterer) ParseSlashFelonyAmountUpdated(log types.Log) (*SlashIndicatorSlashFelonyAmountUpdated, error) {
	event := new(SlashIndicatorSlashFelonyAmountUpdated)
	if err := _SlashIndicator.contract.UnpackLog(event, "SlashFelonyAmountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashIndicatorSlashThresholdsUpdatedIterator is returned from FilterSlashThresholdsUpdated and is used to iterate over the raw logs and unpacked data for SlashThresholdsUpdated events raised by the SlashIndicator contract.
type SlashIndicatorSlashThresholdsUpdatedIterator struct {
	Event *SlashIndicatorSlashThresholdsUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SlashIndicatorSlashThresholdsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashIndicatorSlashThresholdsUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SlashIndicatorSlashThresholdsUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SlashIndicatorSlashThresholdsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashIndicatorSlashThresholdsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashIndicatorSlashThresholdsUpdated represents a SlashThresholdsUpdated event raised by the SlashIndicator contract.
type SlashIndicatorSlashThresholdsUpdated struct {
	FelonyThreshold      *big.Int
	MisdemeanorThreshold *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSlashThresholdsUpdated is a free log retrieval operation binding the contract event 0x25a70459f14137ea646efb17b806c2e75632b931d69e642a8a809b6bfd7217a6.
//
// Solidity: event SlashThresholdsUpdated(uint256 felonyThreshold, uint256 misdemeanorThreshold)
func (_SlashIndicator *SlashIndicatorFilterer) FilterSlashThresholdsUpdated(opts *bind.FilterOpts) (*SlashIndicatorSlashThresholdsUpdatedIterator, error) {

	logs, sub, err := _SlashIndicator.contract.FilterLogs(opts, "SlashThresholdsUpdated")
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorSlashThresholdsUpdatedIterator{contract: _SlashIndicator.contract, event: "SlashThresholdsUpdated", logs: logs, sub: sub}, nil
}

// WatchSlashThresholdsUpdated is a free log subscription operation binding the contract event 0x25a70459f14137ea646efb17b806c2e75632b931d69e642a8a809b6bfd7217a6.
//
// Solidity: event SlashThresholdsUpdated(uint256 felonyThreshold, uint256 misdemeanorThreshold)
func (_SlashIndicator *SlashIndicatorFilterer) WatchSlashThresholdsUpdated(opts *bind.WatchOpts, sink chan<- *SlashIndicatorSlashThresholdsUpdated) (event.Subscription, error) {

	logs, sub, err := _SlashIndicator.contract.WatchLogs(opts, "SlashThresholdsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashIndicatorSlashThresholdsUpdated)
				if err := _SlashIndicator.contract.UnpackLog(event, "SlashThresholdsUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSlashThresholdsUpdated is a log parse operation binding the contract event 0x25a70459f14137ea646efb17b806c2e75632b931d69e642a8a809b6bfd7217a6.
//
// Solidity: event SlashThresholdsUpdated(uint256 felonyThreshold, uint256 misdemeanorThreshold)
func (_SlashIndicator *SlashIndicatorFilterer) ParseSlashThresholdsUpdated(log types.Log) (*SlashIndicatorSlashThresholdsUpdated, error) {
	event := new(SlashIndicatorSlashThresholdsUpdated)
	if err := _SlashIndicator.contract.UnpackLog(event, "SlashThresholdsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashIndicatorUnavailabilitySlashedIterator is returned from FilterUnavailabilitySlashed and is used to iterate over the raw logs and unpacked data for UnavailabilitySlashed events raised by the SlashIndicator contract.
type SlashIndicatorUnavailabilitySlashedIterator struct {
	Event *SlashIndicatorUnavailabilitySlashed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SlashIndicatorUnavailabilitySlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashIndicatorUnavailabilitySlashed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SlashIndicatorUnavailabilitySlashed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SlashIndicatorUnavailabilitySlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashIndicatorUnavailabilitySlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashIndicatorUnavailabilitySlashed represents a UnavailabilitySlashed event raised by the SlashIndicator contract.
type SlashIndicatorUnavailabilitySlashed struct {
	Validator common.Address
	SlashType uint8
	Period    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnavailabilitySlashed is a free log retrieval operation binding the contract event 0x8c2c2bfe532ccdb4523fa2392954fd58445929e7f260d786d0ab93cd981cde54.
//
// Solidity: event UnavailabilitySlashed(address indexed validator, uint8 slashType, uint256 period)
func (_SlashIndicator *SlashIndicatorFilterer) FilterUnavailabilitySlashed(opts *bind.FilterOpts, validator []common.Address) (*SlashIndicatorUnavailabilitySlashedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _SlashIndicator.contract.FilterLogs(opts, "UnavailabilitySlashed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorUnavailabilitySlashedIterator{contract: _SlashIndicator.contract, event: "UnavailabilitySlashed", logs: logs, sub: sub}, nil
}

// WatchUnavailabilitySlashed is a free log subscription operation binding the contract event 0x8c2c2bfe532ccdb4523fa2392954fd58445929e7f260d786d0ab93cd981cde54.
//
// Solidity: event UnavailabilitySlashed(address indexed validator, uint8 slashType, uint256 period)
func (_SlashIndicator *SlashIndicatorFilterer) WatchUnavailabilitySlashed(opts *bind.WatchOpts, sink chan<- *SlashIndicatorUnavailabilitySlashed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _SlashIndicator.contract.WatchLogs(opts, "UnavailabilitySlashed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashIndicatorUnavailabilitySlashed)
				if err := _SlashIndicator.contract.UnpackLog(event, "UnavailabilitySlashed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnavailabilitySlashed is a log parse operation binding the contract event 0x8c2c2bfe532ccdb4523fa2392954fd58445929e7f260d786d0ab93cd981cde54.
//
// Solidity: event UnavailabilitySlashed(address indexed validator, uint8 slashType, uint256 period)
func (_SlashIndicator *SlashIndicatorFilterer) ParseUnavailabilitySlashed(log types.Log) (*SlashIndicatorUnavailabilitySlashed, error) {
	event := new(SlashIndicatorUnavailabilitySlashed)
	if err := _SlashIndicator.contract.UnpackLog(event, "UnavailabilitySlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashIndicatorValidatorContractUpdatedIterator is returned from FilterValidatorContractUpdated and is used to iterate over the raw logs and unpacked data for ValidatorContractUpdated events raised by the SlashIndicator contract.
type SlashIndicatorValidatorContractUpdatedIterator struct {
	Event *SlashIndicatorValidatorContractUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SlashIndicatorValidatorContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashIndicatorValidatorContractUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SlashIndicatorValidatorContractUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SlashIndicatorValidatorContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashIndicatorValidatorContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashIndicatorValidatorContractUpdated represents a ValidatorContractUpdated event raised by the SlashIndicator contract.
type SlashIndicatorValidatorContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterValidatorContractUpdated is a free log retrieval operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_SlashIndicator *SlashIndicatorFilterer) FilterValidatorContractUpdated(opts *bind.FilterOpts) (*SlashIndicatorValidatorContractUpdatedIterator, error) {

	logs, sub, err := _SlashIndicator.contract.FilterLogs(opts, "ValidatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorValidatorContractUpdatedIterator{contract: _SlashIndicator.contract, event: "ValidatorContractUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorContractUpdated is a free log subscription operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_SlashIndicator *SlashIndicatorFilterer) WatchValidatorContractUpdated(opts *bind.WatchOpts, sink chan<- *SlashIndicatorValidatorContractUpdated) (event.Subscription, error) {

	logs, sub, err := _SlashIndicator.contract.WatchLogs(opts, "ValidatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashIndicatorValidatorContractUpdated)
				if err := _SlashIndicator.contract.UnpackLog(event, "ValidatorContractUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorContractUpdated is a log parse operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_SlashIndicator *SlashIndicatorFilterer) ParseValidatorContractUpdated(log types.Log) (*SlashIndicatorValidatorContractUpdated, error) {
	event := new(SlashIndicatorValidatorContractUpdated)
	if err := _SlashIndicator.contract.UnpackLog(event, "ValidatorContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
