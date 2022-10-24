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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"doubleSigningConstrainBlocks\",\"type\":\"uint256\"}],\"name\":\"DoubleSigningConstrainBlocksUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"doubleSigningJailUntilBlock\",\"type\":\"uint256\"}],\"name\":\"DoubleSigningJailUntilBlockUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"felonyJailDuration\",\"type\":\"uint256\"}],\"name\":\"FelonyJailDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"MaintenanceContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashDoubleSignAmount\",\"type\":\"uint256\"}],\"name\":\"SlashDoubleSignAmountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashFelonyAmount\",\"type\":\"uint256\"}],\"name\":\"SlashFelonyAmountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"felonyThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"misdemeanorThreshold\",\"type\":\"uint256\"}],\"name\":\"SlashThresholdsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumISlashIndicator.SlashType\",\"name\":\"slashType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"UnavailabilitySlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ValidatorContractUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"currentUnavailabilityIndicator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"doubleSigningConstrainBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"doubleSigningJailUntilBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"felonyJailDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"felonyThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"getUnavailabilityIndicator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"getUnavailabilitySlashType\",\"outputs\":[{\"internalType\":\"enumISlashIndicator.SlashType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUnavailabilityThresholds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__validatorContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__maintenanceContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_misdemeanorThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_felonyThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slashFelonyAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slashDoubleSignAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_felonyJailBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_doubleSigningConstrainBlocks\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastSlashedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maintenanceContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"misdemeanorThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"precompileValidateDoubleSignAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_felonyJailDuration\",\"type\":\"uint256\"}],\"name\":\"setFelonyJailDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setMaintenanceContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slashDoubleSignAmount\",\"type\":\"uint256\"}],\"name\":\"setSlashDoubleSignAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slashFelonyAmount\",\"type\":\"uint256\"}],\"name\":\"setSlashFelonyAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_felonyThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_misdemeanorThreshold\",\"type\":\"uint256\"}],\"name\":\"setSlashThresholds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setValidatorContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_header1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_header2\",\"type\":\"bytes\"}],\"name\":\"slashDoubleSign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashDoubleSignAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashFelonyAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"unavailabilityThresholdsOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_misdemeanorThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_felonyThreshold\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061001961001e565b6100eb565b600154600160a81b900460ff161561008c5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60015460ff600160a01b909104811610156100e9576001805460ff60a01b191660ff60a01b17905560405160ff81527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6117e6806100fa6000396000f3fe608060405234801561001057600080fd5b50600436106101735760003560e01c80637c2b55a0116100de578063d180ecb011610097578063dfe484cb11610071578063dfe484cb14610347578063e37a14d314610350578063e4f9c4be14610359578063f82bdd011461036c57600080fd5b8063d180ecb0146102e5578063d2cb215e1461032b578063d91935871461033c57600080fd5b80637c2b55a0146102775780638ba796af1461029257806399439089146102a5578063aef250be146102b6578063c96be4cb146102bf578063cdf64a76146102d257600080fd5b8063567a372d11610130578063567a372d146102005780635786e9ee1461020957806362ffe6cb1461021c5780636b79f95f146102525780636d14c4e51461025b5780636e91bec51461026457600080fd5b8063082e7420146101785780631e90b2a01461019e578063389f4f71146101b357806346fe9311146101bc5780634829fa5f146101cf5780634d961e18146101d8575b600080fd5b61018b6101863660046112f6565b61037f565b6040519081526020015b60405180910390f35b6101b16101ac366004611361565b6103f8565b005b61018b60075481565b6101b16101ca3660046112f6565b6105b9565b61018b60055481565b6101eb6101e63660046113e2565b6105fd565b60408051928352602083019190915201610195565b61018b60065481565b6101b161021736600461140c565b610875565b61018b61022a3660046113e2565b6001600160a01b03919091166000908152600260209081526040808320938352929052205490565b61018b600a5481565b61018b60095481565b6101b161027236600461140c565b6108b6565b60675b6040516001600160a01b039091168152602001610195565b6101b16102a0366004611425565b6108f7565b6000546001600160a01b031661027a565b61018b60045481565b6101b16102cd3660046112f6565b610a6a565b6101b16102e03660046112f6565b610da0565b61031e6102f33660046113e2565b6001600160a01b03919091166000908152600360209081526040808320938352929052205460ff1690565b60405161019591906114c4565b6001546001600160a01b031661027a565b6006546007546101eb565b61018b60085481565b61018b600b5481565b6101b16103673660046114d2565b610de1565b6101b161037a36600461140c565b610e27565b6000805460405163f8549af960e01b81524360048201526103f29184916001600160a01b039091169063f8549af990602401602060405180830381865afa1580156103ce573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061022a91906114f4565b92915050565b3341146104205760405162461bcd60e51b81526004016104179061150d565b60405180910390fd5b60045443116104415760405162461bcd60e51b81526004016104179061155b565b61044a85610e68565b156105ae5761045b84848484610f6e565b156105ae576000805460405163f8549af960e01b81524360048201526001600160a01b039091169063f8549af990602401602060405180830381865afa1580156104a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104cd91906114f4565b6001600160a01b038716600081815260036020818152604080842086855290915291829020805460ff191682179055905192935090917f8c2c2bfe532ccdb4523fa2392954fd58445929e7f260d786d0ab93cd981cde54916105309185906115de565b60405180910390a2600054600b54600954604051631c3e07db60e21b81526001600160a01b038a81166004830152602482019390935260448101919091529116906370f81f6c90606401600060405180830381600087803b15801561059457600080fd5b505af11580156105a8573d6000803e3d6000fd5b50505050505b505043600455505050565b6105c161106a565b6001600160a01b0316336001600160a01b0316146105f15760405162461bcd60e51b8152600401610417906115f9565b6105fa81611098565b50565b60008060008060009054906101000a90046001600160a01b03166001600160a01b0316635186dc7e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610654573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061067891906114f4565b60008054906101000a90046001600160a01b03166001600160a01b0316636aa1c2ef6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106c9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106ed91906114f4565b6106f79190611651565b90506000816107068187611668565b6107109190611651565b905060006001610720848461168a565b61072a919061169d565b6001546040516334e7fb8d60e21b81526001600160a01b038a811660048301529293506000929091169063d39fee3490602401606060405180830381865afa15801561077a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061079e91906116b0565b80519091506000906107b19085856110ed565b60208301519091506000906107c79086866110ed565b9050858280156107d45750815b1561080657835160208501516107ea919061169d565b6107f590600161168a565b6107ff908261169d565b9050610846565b82156108185783516107ea908661169d565b81156108465785846020015161082e919061169d565b61083990600161168a565b610843908261169d565b90505b600654610854908289611107565b600754909950610865908289611107565b9750505050505050509250929050565b61087d61106a565b6001600160a01b0316336001600160a01b0316146108ad5760405162461bcd60e51b8152600401610417906115f9565b6105fa8161111e565b6108be61106a565b6001600160a01b0316336001600160a01b0316146108ee5760405162461bcd60e51b8152600401610417906115f9565b6105fa81611153565b600154600160a81b900460ff161580801561091d575060018054600160a01b900460ff16105b8061093d5750303b15801561093d575060018054600160a01b900460ff16145b6109a05760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610417565b6001805460ff60a01b1916600160a01b17905580156109cd576001805460ff60a81b1916600160a81b1790555b6109d689611188565b6109df88611098565b6109e986886111d6565b6109f285611153565b6109fb8461121d565b610a048361111e565b610a0d82611252565b610a18600019611287565b8015610a5f576001805460ff60a81b191681556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050505050565b334114610a895760405162461bcd60e51b81526004016104179061150d565b6004544311610aaa5760405162461bcd60e51b81526004016104179061155b565b610ab381610e68565b15610d99576000805460405163f8549af960e01b81524360048201526001600160a01b039091169063f8549af990602401602060405180830381865afa158015610b01573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b2591906114f4565b6001600160a01b03831660009081526002602090815260408083208484529091528120805492935090918290610b5a9061171a565b91829055509050600080610b6e85436105fd565b6001600160a01b0387166000908152600360209081526040808320898452909152902054919350915060ff16818410801590610bbb57506002816003811115610bb957610bb961148c565b105b15610cc2576001600160a01b038616600081815260036020908152604080832089845290915290819020805460ff1916600290811790915590517f8c2c2bfe532ccdb4523fa2392954fd58445929e7f260d786d0ab93cd981cde5491610c229189906115de565b60405180910390a2600054600a546001600160a01b03909116906370f81f6c908890610c4e904361168a565b6008546040516001600160e01b031960e086901b1681526001600160a01b039093166004840152602483019190915260448201526064015b600060405180830381600087803b158015610ca057600080fd5b505af1158015610cb4573d6000803e3d6000fd5b505050505050505050610d99565b828410158015610ce357506001816003811115610ce157610ce161148c565b105b15610d93576001600160a01b038616600081815260036020908152604080832089845290915290819020805460ff1916600190811790915590517f8c2c2bfe532ccdb4523fa2392954fd58445929e7f260d786d0ab93cd981cde5491610d4a9189906115de565b60405180910390a260008054604051631c3e07db60e21b81526001600160a01b0389811660048301526024820184905260448201939093529116906370f81f6c90606401610c86565b50505050505b5043600455565b610da861106a565b6001600160a01b0316336001600160a01b031614610dd85760405162461bcd60e51b8152600401610417906115f9565b6105fa81611188565b610de961106a565b6001600160a01b0316336001600160a01b031614610e195760405162461bcd60e51b8152600401610417906115f9565b610e2382826111d6565b5050565b610e2f61106a565b6001600160a01b0316336001600160a01b031614610e5f5760405162461bcd60e51b8152600401610417906115f9565b6105fa8161121d565b6000336001600160a01b03831614801590610eec575060005460405163facd743b60e01b81526001600160a01b0384811660048301529091169063facd743b90602401602060405180830381865afa158015610ec8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610eec9190611733565b80156103f2575060015460405163934e9a0360e01b81526001600160a01b0384811660048301524360248301529091169063934e9a0390604401602060405180830381865afa158015610f43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f679190611733565b1592915050565b6040516000906067906001908390610f9090899089908990899060240161177e565b60408051601f198184030181529190526020810180516001600160e01b031663580a316360e01b1790528051909150610fc76112bc565b602083016020828483895afa610fdc57600094505b503d610fe757600093505b8361105a5760405162461bcd60e51b815260206004820152603b60248201527f507265636f6d70696c65557361676556616c6964617465446f75626c6553696760448201527f6e3a2063616c6c20746f20707265636f6d70696c65206661696c7300000000006064820152608401610417565b5115159998505050505050505050565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103546001600160a01b031690565b600180546001600160a01b0319166001600160a01b0383169081179091556040519081527f31a33f126a5bae3c5bdf6cfc2cd6dcfffe2fe9634bdb09e21c44762993889e3b906020015b60405180910390a150565b60008383111580156110ff5750818411155b949350505050565b6000816111148486611651565b6110ff9190611668565b600a8190556040518181527f3092a623ebcf71b79f9f68801b081a4d0c839dfb4c0e6f9ff118f4fb870375dc906020016110e2565b60088190556040518181527f71491762d43dbadd94f85cf7a0322192a7c22f3ca4fb30de8895420607b56712906020016110e2565b600080546001600160a01b0319166001600160a01b0383169081179091556040519081527fef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169906020016110e2565b6007829055600681905560408051838152602081018390527f25a70459f14137ea646efb17b806c2e75632b931d69e642a8a809b6bfd7217a6910160405180910390a15050565b60098190556040518181527f42458d43c99c17f034dde22f3a1cf003e4bc27372c1ff14cb9e68a79a1b1e8ed906020016110e2565b60058190556040518181527f55fe8ed087353640949b09246027afc346464699a647dda2a478a11028725b79906020016110e2565b600b8190556040518181527f1c0135655cb9101f41e405c4063b1eee093f2238031a2593443ee3ee815ff35a906020016110e2565b60405180602001604052806001906020820280368337509192915050565b80356001600160a01b03811681146112f157600080fd5b919050565b60006020828403121561130857600080fd5b611311826112da565b9392505050565b60008083601f84011261132a57600080fd5b50813567ffffffffffffffff81111561134257600080fd5b60208301915083602082850101111561135a57600080fd5b9250929050565b60008060008060006060868803121561137957600080fd5b611382866112da565b9450602086013567ffffffffffffffff8082111561139f57600080fd5b6113ab89838a01611318565b909650945060408801359150808211156113c457600080fd5b506113d188828901611318565b969995985093965092949392505050565b600080604083850312156113f557600080fd5b6113fe836112da565b946020939093013593505050565b60006020828403121561141e57600080fd5b5035919050565b600080600080600080600080610100898b03121561144257600080fd5b61144b896112da565b975061145960208a016112da565b979a9799505050506040860135956060810135956080820135955060a0820135945060c0820135935060e0909101359150565b634e487b7160e01b600052602160045260246000fd5b600481106114c057634e487b7160e01b600052602160045260246000fd5b9052565b602081016103f282846114a2565b600080604083850312156114e557600080fd5b50508035926020909101359150565b60006020828403121561150657600080fd5b5051919050565b6020808252602e908201527f536c617368496e64696361746f723a206d6574686f642063616c6c6572206d7560408201526d737420626520636f696e6261736560901b606082015260800190565b6020808252605c908201527f536c617368496e64696361746f723a2063616e6e6f7420736c6173682061207660408201527f616c696461746f72207477696365206f7220736c617368206d6f72652074686160608201527f6e206f6e652076616c696461746f7220696e206f6e6520626c6f636b00000000608082015260a00190565b604081016115ec82856114a2565b8260208301529392505050565b60208082526022908201527f48617350726f787941646d696e3a20756e617574686f72697a65642073656e6460408201526132b960f11b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b80820281158282048414176103f2576103f261163b565b60008261168557634e487b7160e01b600052601260045260246000fd5b500490565b808201808211156103f2576103f261163b565b818103818111156103f2576103f261163b565b6000606082840312156116c257600080fd5b6040516060810181811067ffffffffffffffff821117156116f357634e487b7160e01b600052604160045260246000fd5b80604052508251815260208301516020820152604083015160408201528091505092915050565b60006001820161172c5761172c61163b565b5060010190565b60006020828403121561174557600080fd5b8151801515811461131157600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b604081526000611792604083018688611755565b82810360208401526117a5818587611755565b97965050505050505056fea26469706673582212205b31a55f6da3d5f9024073a9755dee46b2dcb6545345e68646e70fcb7e88b26364736f6c63430008110033",
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

// DoubleSigningConstrainBlocks is a free data retrieval call binding the contract method 0x4829fa5f.
//
// Solidity: function doubleSigningConstrainBlocks() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCaller) DoubleSigningConstrainBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "doubleSigningConstrainBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DoubleSigningConstrainBlocks is a free data retrieval call binding the contract method 0x4829fa5f.
//
// Solidity: function doubleSigningConstrainBlocks() view returns(uint256)
func (_SlashIndicator *SlashIndicatorSession) DoubleSigningConstrainBlocks() (*big.Int, error) {
	return _SlashIndicator.Contract.DoubleSigningConstrainBlocks(&_SlashIndicator.CallOpts)
}

// DoubleSigningConstrainBlocks is a free data retrieval call binding the contract method 0x4829fa5f.
//
// Solidity: function doubleSigningConstrainBlocks() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCallerSession) DoubleSigningConstrainBlocks() (*big.Int, error) {
	return _SlashIndicator.Contract.DoubleSigningConstrainBlocks(&_SlashIndicator.CallOpts)
}

// DoubleSigningJailUntilBlock is a free data retrieval call binding the contract method 0xe37a14d3.
//
// Solidity: function doubleSigningJailUntilBlock() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCaller) DoubleSigningJailUntilBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "doubleSigningJailUntilBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DoubleSigningJailUntilBlock is a free data retrieval call binding the contract method 0xe37a14d3.
//
// Solidity: function doubleSigningJailUntilBlock() view returns(uint256)
func (_SlashIndicator *SlashIndicatorSession) DoubleSigningJailUntilBlock() (*big.Int, error) {
	return _SlashIndicator.Contract.DoubleSigningJailUntilBlock(&_SlashIndicator.CallOpts)
}

// DoubleSigningJailUntilBlock is a free data retrieval call binding the contract method 0xe37a14d3.
//
// Solidity: function doubleSigningJailUntilBlock() view returns(uint256)
func (_SlashIndicator *SlashIndicatorCallerSession) DoubleSigningJailUntilBlock() (*big.Int, error) {
	return _SlashIndicator.Contract.DoubleSigningJailUntilBlock(&_SlashIndicator.CallOpts)
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

// PrecompileValidateDoubleSignAddress is a free data retrieval call binding the contract method 0x7c2b55a0.
//
// Solidity: function precompileValidateDoubleSignAddress() view returns(address)
func (_SlashIndicator *SlashIndicatorCaller) PrecompileValidateDoubleSignAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlashIndicator.contract.Call(opts, &out, "precompileValidateDoubleSignAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrecompileValidateDoubleSignAddress is a free data retrieval call binding the contract method 0x7c2b55a0.
//
// Solidity: function precompileValidateDoubleSignAddress() view returns(address)
func (_SlashIndicator *SlashIndicatorSession) PrecompileValidateDoubleSignAddress() (common.Address, error) {
	return _SlashIndicator.Contract.PrecompileValidateDoubleSignAddress(&_SlashIndicator.CallOpts)
}

// PrecompileValidateDoubleSignAddress is a free data retrieval call binding the contract method 0x7c2b55a0.
//
// Solidity: function precompileValidateDoubleSignAddress() view returns(address)
func (_SlashIndicator *SlashIndicatorCallerSession) PrecompileValidateDoubleSignAddress() (common.Address, error) {
	return _SlashIndicator.Contract.PrecompileValidateDoubleSignAddress(&_SlashIndicator.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0x8ba796af.
//
// Solidity: function initialize(address __validatorContract, address __maintenanceContract, uint256 _misdemeanorThreshold, uint256 _felonyThreshold, uint256 _slashFelonyAmount, uint256 _slashDoubleSignAmount, uint256 _felonyJailBlocks, uint256 _doubleSigningConstrainBlocks) returns()
func (_SlashIndicator *SlashIndicatorTransactor) Initialize(opts *bind.TransactOpts, __validatorContract common.Address, __maintenanceContract common.Address, _misdemeanorThreshold *big.Int, _felonyThreshold *big.Int, _slashFelonyAmount *big.Int, _slashDoubleSignAmount *big.Int, _felonyJailBlocks *big.Int, _doubleSigningConstrainBlocks *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.contract.Transact(opts, "initialize", __validatorContract, __maintenanceContract, _misdemeanorThreshold, _felonyThreshold, _slashFelonyAmount, _slashDoubleSignAmount, _felonyJailBlocks, _doubleSigningConstrainBlocks)
}

// Initialize is a paid mutator transaction binding the contract method 0x8ba796af.
//
// Solidity: function initialize(address __validatorContract, address __maintenanceContract, uint256 _misdemeanorThreshold, uint256 _felonyThreshold, uint256 _slashFelonyAmount, uint256 _slashDoubleSignAmount, uint256 _felonyJailBlocks, uint256 _doubleSigningConstrainBlocks) returns()
func (_SlashIndicator *SlashIndicatorSession) Initialize(__validatorContract common.Address, __maintenanceContract common.Address, _misdemeanorThreshold *big.Int, _felonyThreshold *big.Int, _slashFelonyAmount *big.Int, _slashDoubleSignAmount *big.Int, _felonyJailBlocks *big.Int, _doubleSigningConstrainBlocks *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.Contract.Initialize(&_SlashIndicator.TransactOpts, __validatorContract, __maintenanceContract, _misdemeanorThreshold, _felonyThreshold, _slashFelonyAmount, _slashDoubleSignAmount, _felonyJailBlocks, _doubleSigningConstrainBlocks)
}

// Initialize is a paid mutator transaction binding the contract method 0x8ba796af.
//
// Solidity: function initialize(address __validatorContract, address __maintenanceContract, uint256 _misdemeanorThreshold, uint256 _felonyThreshold, uint256 _slashFelonyAmount, uint256 _slashDoubleSignAmount, uint256 _felonyJailBlocks, uint256 _doubleSigningConstrainBlocks) returns()
func (_SlashIndicator *SlashIndicatorTransactorSession) Initialize(__validatorContract common.Address, __maintenanceContract common.Address, _misdemeanorThreshold *big.Int, _felonyThreshold *big.Int, _slashFelonyAmount *big.Int, _slashDoubleSignAmount *big.Int, _felonyJailBlocks *big.Int, _doubleSigningConstrainBlocks *big.Int) (*types.Transaction, error) {
	return _SlashIndicator.Contract.Initialize(&_SlashIndicator.TransactOpts, __validatorContract, __maintenanceContract, _misdemeanorThreshold, _felonyThreshold, _slashFelonyAmount, _slashDoubleSignAmount, _felonyJailBlocks, _doubleSigningConstrainBlocks)
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

// SlashDoubleSign is a paid mutator transaction binding the contract method 0x1e90b2a0.
//
// Solidity: function slashDoubleSign(address _validatorAddr, bytes _header1, bytes _header2) returns()
func (_SlashIndicator *SlashIndicatorTransactor) SlashDoubleSign(opts *bind.TransactOpts, _validatorAddr common.Address, _header1 []byte, _header2 []byte) (*types.Transaction, error) {
	return _SlashIndicator.contract.Transact(opts, "slashDoubleSign", _validatorAddr, _header1, _header2)
}

// SlashDoubleSign is a paid mutator transaction binding the contract method 0x1e90b2a0.
//
// Solidity: function slashDoubleSign(address _validatorAddr, bytes _header1, bytes _header2) returns()
func (_SlashIndicator *SlashIndicatorSession) SlashDoubleSign(_validatorAddr common.Address, _header1 []byte, _header2 []byte) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SlashDoubleSign(&_SlashIndicator.TransactOpts, _validatorAddr, _header1, _header2)
}

// SlashDoubleSign is a paid mutator transaction binding the contract method 0x1e90b2a0.
//
// Solidity: function slashDoubleSign(address _validatorAddr, bytes _header1, bytes _header2) returns()
func (_SlashIndicator *SlashIndicatorTransactorSession) SlashDoubleSign(_validatorAddr common.Address, _header1 []byte, _header2 []byte) (*types.Transaction, error) {
	return _SlashIndicator.Contract.SlashDoubleSign(&_SlashIndicator.TransactOpts, _validatorAddr, _header1, _header2)
}

// SlashIndicatorDoubleSigningConstrainBlocksUpdatedIterator is returned from FilterDoubleSigningConstrainBlocksUpdated and is used to iterate over the raw logs and unpacked data for DoubleSigningConstrainBlocksUpdated events raised by the SlashIndicator contract.
type SlashIndicatorDoubleSigningConstrainBlocksUpdatedIterator struct {
	Event *SlashIndicatorDoubleSigningConstrainBlocksUpdated // Event containing the contract specifics and raw log

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
func (it *SlashIndicatorDoubleSigningConstrainBlocksUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashIndicatorDoubleSigningConstrainBlocksUpdated)
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
		it.Event = new(SlashIndicatorDoubleSigningConstrainBlocksUpdated)
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
func (it *SlashIndicatorDoubleSigningConstrainBlocksUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashIndicatorDoubleSigningConstrainBlocksUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashIndicatorDoubleSigningConstrainBlocksUpdated represents a DoubleSigningConstrainBlocksUpdated event raised by the SlashIndicator contract.
type SlashIndicatorDoubleSigningConstrainBlocksUpdated struct {
	DoubleSigningConstrainBlocks *big.Int
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterDoubleSigningConstrainBlocksUpdated is a free log retrieval operation binding the contract event 0x55fe8ed087353640949b09246027afc346464699a647dda2a478a11028725b79.
//
// Solidity: event DoubleSigningConstrainBlocksUpdated(uint256 doubleSigningConstrainBlocks)
func (_SlashIndicator *SlashIndicatorFilterer) FilterDoubleSigningConstrainBlocksUpdated(opts *bind.FilterOpts) (*SlashIndicatorDoubleSigningConstrainBlocksUpdatedIterator, error) {

	logs, sub, err := _SlashIndicator.contract.FilterLogs(opts, "DoubleSigningConstrainBlocksUpdated")
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorDoubleSigningConstrainBlocksUpdatedIterator{contract: _SlashIndicator.contract, event: "DoubleSigningConstrainBlocksUpdated", logs: logs, sub: sub}, nil
}

// WatchDoubleSigningConstrainBlocksUpdated is a free log subscription operation binding the contract event 0x55fe8ed087353640949b09246027afc346464699a647dda2a478a11028725b79.
//
// Solidity: event DoubleSigningConstrainBlocksUpdated(uint256 doubleSigningConstrainBlocks)
func (_SlashIndicator *SlashIndicatorFilterer) WatchDoubleSigningConstrainBlocksUpdated(opts *bind.WatchOpts, sink chan<- *SlashIndicatorDoubleSigningConstrainBlocksUpdated) (event.Subscription, error) {

	logs, sub, err := _SlashIndicator.contract.WatchLogs(opts, "DoubleSigningConstrainBlocksUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashIndicatorDoubleSigningConstrainBlocksUpdated)
				if err := _SlashIndicator.contract.UnpackLog(event, "DoubleSigningConstrainBlocksUpdated", log); err != nil {
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

// ParseDoubleSigningConstrainBlocksUpdated is a log parse operation binding the contract event 0x55fe8ed087353640949b09246027afc346464699a647dda2a478a11028725b79.
//
// Solidity: event DoubleSigningConstrainBlocksUpdated(uint256 doubleSigningConstrainBlocks)
func (_SlashIndicator *SlashIndicatorFilterer) ParseDoubleSigningConstrainBlocksUpdated(log types.Log) (*SlashIndicatorDoubleSigningConstrainBlocksUpdated, error) {
	event := new(SlashIndicatorDoubleSigningConstrainBlocksUpdated)
	if err := _SlashIndicator.contract.UnpackLog(event, "DoubleSigningConstrainBlocksUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashIndicatorDoubleSigningJailUntilBlockUpdatedIterator is returned from FilterDoubleSigningJailUntilBlockUpdated and is used to iterate over the raw logs and unpacked data for DoubleSigningJailUntilBlockUpdated events raised by the SlashIndicator contract.
type SlashIndicatorDoubleSigningJailUntilBlockUpdatedIterator struct {
	Event *SlashIndicatorDoubleSigningJailUntilBlockUpdated // Event containing the contract specifics and raw log

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
func (it *SlashIndicatorDoubleSigningJailUntilBlockUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashIndicatorDoubleSigningJailUntilBlockUpdated)
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
		it.Event = new(SlashIndicatorDoubleSigningJailUntilBlockUpdated)
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
func (it *SlashIndicatorDoubleSigningJailUntilBlockUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashIndicatorDoubleSigningJailUntilBlockUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashIndicatorDoubleSigningJailUntilBlockUpdated represents a DoubleSigningJailUntilBlockUpdated event raised by the SlashIndicator contract.
type SlashIndicatorDoubleSigningJailUntilBlockUpdated struct {
	DoubleSigningJailUntilBlock *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterDoubleSigningJailUntilBlockUpdated is a free log retrieval operation binding the contract event 0x1c0135655cb9101f41e405c4063b1eee093f2238031a2593443ee3ee815ff35a.
//
// Solidity: event DoubleSigningJailUntilBlockUpdated(uint256 doubleSigningJailUntilBlock)
func (_SlashIndicator *SlashIndicatorFilterer) FilterDoubleSigningJailUntilBlockUpdated(opts *bind.FilterOpts) (*SlashIndicatorDoubleSigningJailUntilBlockUpdatedIterator, error) {

	logs, sub, err := _SlashIndicator.contract.FilterLogs(opts, "DoubleSigningJailUntilBlockUpdated")
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorDoubleSigningJailUntilBlockUpdatedIterator{contract: _SlashIndicator.contract, event: "DoubleSigningJailUntilBlockUpdated", logs: logs, sub: sub}, nil
}

// WatchDoubleSigningJailUntilBlockUpdated is a free log subscription operation binding the contract event 0x1c0135655cb9101f41e405c4063b1eee093f2238031a2593443ee3ee815ff35a.
//
// Solidity: event DoubleSigningJailUntilBlockUpdated(uint256 doubleSigningJailUntilBlock)
func (_SlashIndicator *SlashIndicatorFilterer) WatchDoubleSigningJailUntilBlockUpdated(opts *bind.WatchOpts, sink chan<- *SlashIndicatorDoubleSigningJailUntilBlockUpdated) (event.Subscription, error) {

	logs, sub, err := _SlashIndicator.contract.WatchLogs(opts, "DoubleSigningJailUntilBlockUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashIndicatorDoubleSigningJailUntilBlockUpdated)
				if err := _SlashIndicator.contract.UnpackLog(event, "DoubleSigningJailUntilBlockUpdated", log); err != nil {
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

// ParseDoubleSigningJailUntilBlockUpdated is a log parse operation binding the contract event 0x1c0135655cb9101f41e405c4063b1eee093f2238031a2593443ee3ee815ff35a.
//
// Solidity: event DoubleSigningJailUntilBlockUpdated(uint256 doubleSigningJailUntilBlock)
func (_SlashIndicator *SlashIndicatorFilterer) ParseDoubleSigningJailUntilBlockUpdated(log types.Log) (*SlashIndicatorDoubleSigningJailUntilBlockUpdated, error) {
	event := new(SlashIndicatorDoubleSigningJailUntilBlockUpdated)
	if err := _SlashIndicator.contract.UnpackLog(event, "DoubleSigningJailUntilBlockUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
