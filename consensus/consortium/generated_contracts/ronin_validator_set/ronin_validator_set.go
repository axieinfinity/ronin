// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package roninValidatorSet

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

// ICandidateManagerValidatorCandidate is an auto generated low-level Go binding around an user-defined struct.
type ICandidateManagerValidatorCandidate struct {
	Admin          common.Address
	ConsensusAddr  common.Address
	TreasuryAddr   common.Address
	CommissionRate *big.Int
	ExtraData      []byte
}

// RoninValidatorSetMetaData contains all meta data concerning the RoninValidatorSet contract.
var RoninValidatorSetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"name\":\"AddressesPriorityStatusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coinbaseAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"submittedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bonusAmount\",\"type\":\"uint256\"}],\"name\":\"BlockRewardSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"MaintenanceContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MaxPrioritizedValidatorNumberUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"MaxValidatorCandidateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MaxValidatorNumberUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MiningRewardDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"NumberOfBlocksInEpochUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"NumberOfEpochsInPeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coinbaseAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"name\":\"RewardDeprecated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"SlashIndicatorContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"StakingContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakingRewardDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"StakingVestingContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"treasuryAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"candidateIdx\",\"type\":\"uint256\"}],\"name\":\"ValidatorCandidateAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"}],\"name\":\"ValidatorCandidateRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"jailedUntil\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deductedStakingAmount\",\"type\":\"uint256\"}],\"name\":\"ValidatorPunished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"name\":\"ValidatorSetUpdated\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_treasuryAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_commissionRate\",\"type\":\"uint256\"}],\"name\":\"addValidatorCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"epochEndingAt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"epochOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCandidateInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"treasuryAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structICandidateManager.ValidatorCandidate[]\",\"name\":\"_list\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastUpdatedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getPriorityStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorCandidates\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_validatorList\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__slashIndicatorContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__stakingContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__stakingVestingContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__maintenanceContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"__maxValidatorNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"__maxValidatorCandidate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"__maxPrioritizedValidatorNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"__numberOfBlocksInEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"__numberOfEpochsInPeriod\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"isCandidateAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isValidatorCandidate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addrList\",\"type\":\"address[]\"}],\"name\":\"jailed\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"_result\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maintenanceContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPrioritizedValidatorNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_maximumPrioritizedValidatorNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxValidatorCandidate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxValidatorNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_maximumValidatorNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberOfBlocksInEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_numberOfBlocks\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberOfEpochsInPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_numberOfEpochs\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"periodEndingAt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"periodOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addrList\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"rewardDeprecated\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"_result\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setMaintenanceContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"setMaxValidatorCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"__maxValidatorNumber\",\"type\":\"uint256\"}],\"name\":\"setMaxValidatorNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"__numberOfBlocksInEpoch\",\"type\":\"uint256\"}],\"name\":\"setNumberOfBlocksInEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"__numberOfEpochsInPeriod\",\"type\":\"uint256\"}],\"name\":\"setNumberOfEpochsInPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addrs\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"_statuses\",\"type\":\"bool[]\"}],\"name\":\"setPrioritizedAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setSlashIndicatorContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setStakingContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setStakingVestingContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newJailedUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slashAmount\",\"type\":\"uint256\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashIndicatorContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingVestingContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitBlockReward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"syncCandidates\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_balances\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrapUpEpoch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50620000226200002860201b60201c565b620001d6565b600860019054906101000a900460ff16156200007b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000729062000179565b60405180910390fd5b60ff8016600860009054906101000a900460ff1660ff161015620000f05760ff600860006101000a81548160ff021916908360ff1602179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860ff604051620000e79190620001b9565b60405180910390a15b565b600082825260208201905092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320696e69746960008201527f616c697a696e6700000000000000000000000000000000000000000000000000602082015250565b600062000161602783620000f2565b91506200016e8262000103565b604082019050919050565b60006020820190508181036000830152620001948162000152565b9050919050565b600060ff82169050919050565b620001b3816200019b565b82525050565b6000602082019050620001d06000830184620001a8565b92915050565b6162ca80620001e66000396000f3fe6080604052600436106102345760003560e01c806387c891bd1161012e578063d2cb215e116100ab578063e18572bf1161006f578063e18572bf1461083b578063ee99205c14610864578063eeb629a81461088f578063f8549af9146108ba578063facd743b146108f757610243565b8063d2cb215e1461076a578063d33a5ca214610795578063d6fa322c146107be578063d72733fc146107e7578063de7702fb1461081057610243565b8063ad295783116100f2578063ad29578314610697578063b7ab4db5146106c0578063ba77b06c146106eb578063beb3e38214610716578063d09f1ab41461073f57610243565b806387c891bd1461058c5780639dd373b9146105b7578063a0c3f2d2146105e0578063a3d545f51461061d578063ac00125f1461065a57610243565b806352091f17116101bc57806370f81f6c1161018057806370f81f6c146104b657806372e46810146104df5780637593ff71146104e9578063823a7b9c14610526578063865231931461054f57610243565b806352091f17146104005780635248184a1461040a5780635a08482d14610435578063605239a1146104605780636aa1c2ef1461048b57610243565b80633529214b116102035780633529214b1461031b5780634454af9d1461034657806346fe9311146103835780634f2a693f146103ac5780635186dc7e146103d557610243565b806304d971ab1461024d5780630f43a6771461028a57806325a6b529146102b55780632bcf3d15146102f257610243565b3661024357610241610934565b005b61024b610934565b005b34801561025957600080fd5b50610274600480360381019061026f9190614227565b6109ab565b6040516102819190614282565b60405180910390f35b34801561029657600080fd5b5061029f610a46565b6040516102ac91906142b6565b60405180910390f35b3480156102c157600080fd5b506102dc60048036038101906102d791906142fd565b610a4c565b6040516102e99190614282565b60405180910390f35b3480156102fe57600080fd5b506103196004803603810190610314919061432a565b610a83565b005b34801561032757600080fd5b50610330610b04565b60405161033d9190614366565b60405180910390f35b34801561035257600080fd5b5061036d600480360381019061036891906144da565b610b2e565b60405161037a91906145e1565b60405180910390f35b34801561038f57600080fd5b506103aa60048036038101906103a5919061432a565b610b9d565b005b3480156103b857600080fd5b506103d360048036038101906103ce91906142fd565b610c1e565b005b3480156103e157600080fd5b506103ea610c9f565b6040516103f791906142b6565b60405180910390f35b610408610ca9565b005b34801561041657600080fd5b5061041f611047565b60405161042c91906147f9565b60405180910390f35b34801561044157600080fd5b5061044a61130b565b6040516104579190614366565b60405180910390f35b34801561046c57600080fd5b50610475611335565b60405161048291906142b6565b60405180910390f35b34801561049757600080fd5b506104a061133f565b6040516104ad91906142b6565b60405180910390f35b3480156104c257600080fd5b506104dd60048036038101906104d8919061481b565b611349565b005b6104e76116ec565b005b3480156104f557600080fd5b50610510600480360381019061050b91906142fd565b611bc8565b60405161051d9190614282565b60405180910390f35b34801561053257600080fd5b5061054d600480360381019061054891906142fd565b611bef565b005b34801561055b57600080fd5b506105766004803603810190610571919061432a565b611c70565b6040516105839190614282565b60405180910390f35b34801561059857600080fd5b506105a1611cc6565b6040516105ae91906142b6565b60405180910390f35b3480156105c357600080fd5b506105de60048036038101906105d9919061432a565b611cd0565b005b3480156105ec57600080fd5b506106076004803603810190610602919061432a565b611d51565b6040516106149190614282565b60405180910390f35b34801561062957600080fd5b50610644600480360381019061063f91906142fd565b611d9d565b60405161065191906142b6565b60405180910390f35b34801561066657600080fd5b50610681600480360381019061067c919061486e565b611dc0565b60405161068e91906145e1565b60405180910390f35b3480156106a357600080fd5b506106be60048036038101906106b9919061432a565b611e31565b005b3480156106cc57600080fd5b506106d5611eb2565b6040516106e29190614979565b60405180910390f35b3480156106f757600080fd5b50610700611fa5565b60405161070d9190614979565b60405180910390f35b34801561072257600080fd5b5061073d6004803603810190610738919061499b565b612033565b005b34801561074b57600080fd5b506107546121ca565b60405161076191906142b6565b60405180910390f35b34801561077657600080fd5b5061077f6121d4565b60405161078c9190614366565b60405180910390f35b3480156107a157600080fd5b506107bc60048036038101906107b79190614b54565b6121fe565b005b3480156107ca57600080fd5b506107e560048036038101906107e091906142fd565b61246f565b005b3480156107f357600080fd5b5061080e600480360381019061080991906142fd565b6124f0565b005b34801561081c57600080fd5b50610825612571565b6040516108329190614c7b565b60405180910390f35b34801561084757600080fd5b50610862600480360381019061085d9190614cc9565b612770565b005b34801561087057600080fd5b50610879612bb7565b6040516108869190614366565b60405180910390f35b34801561089b57600080fd5b506108a4612be0565b6040516108b191906142b6565b60405180910390f35b3480156108c657600080fd5b506108e160048036038101906108dc91906142fd565b612bea565b6040516108ee91906142b6565b60405180910390f35b34801561090357600080fd5b5061091e6004803603810190610919919061432a565b612c1a565b60405161092b9190614282565b60405180910390f35b61093c610b04565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109a090614dd9565b60405180910390fd5b565b60008173ffffffffffffffffffffffffffffffffffffffff16600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614905092915050565b600d5481565b600080600b54600a54610a5f9190614e28565b9050600181610a6e9190614e82565b8184610a7a9190614ee5565b14915050919050565b610a8b612c70565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610af8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aef90614f88565b60405180910390fd5b610b0181612cc7565b50565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b606060005b8251811015610b9757610b5f838281518110610b5257610b51614fa8565b5b6020026020010151612d42565b828281518110610b7257610b71614fa8565b5b6020026020010190151590811515815250508080610b8f90614fd7565b915050610b33565b50919050565b610ba5612c70565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c12576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c0990614f88565b60405180910390fd5b610c1b81612d8e565b50565b610c26612c70565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c93576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c8a90614f88565b60405180910390fd5b610c9c81612e09565b50565b6000600b54905090565b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610d17576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d0e90615091565b60405180910390fd5b600034905060008103610d2a5750611045565b6000339050610d3881612c1a565b1580610d495750610d4881612d42565b5b80610d625750610d6181610d5c43612bea565b612e4a565b5b15610da7577f2439a6ac441f1d6b3dbb7827ef6e056822e2261f900cad468012eee4f1f7f31c8183604051610d989291906150b1565b60405180910390a15050611045565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166379ad52846040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610e18573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e3c91906150ef565b905060008184610e4c919061511c565b905060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000600760008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030154905060006127108483610ecc9190614e28565b610ed69190615150565b905060008185610ee69190614e82565b905081601460008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610f37919061511c565b9250508190555080601560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610f8d919061511c565b925050819055508373ffffffffffffffffffffffffffffffffffffffff1663b863d71088836040518363ffffffff1660e01b8152600401610fcf9291906150b1565b600060405180830381600087803b158015610fe957600080fd5b505af1158015610ffd573d6000803e3d6000fd5b505050507f0ede5c3be8625943fa64003cd4b91230089411249f3059bac6500873543ca9b187898860405161103493929190615181565b60405180910390a150505050505050505b565b606060058054905067ffffffffffffffff81111561106857611067614397565b5b6040519080825280602002602001820160405280156110a157816020015b61108e6140cd565b8152602001906001900390816110865790505b50905060005b81518110156113075760076000600583815481106110c8576110c7614fa8565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060a00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160038201548152602001600482018054611254906151e7565b80601f0160208091040260200160405190810160405280929190818152602001828054611280906151e7565b80156112cd5780601f106112a2576101008083540402835291602001916112cd565b820191906000526020600020905b8154815290600101906020018083116112b057829003601f168201915b5050505050815250508282815181106112e9576112e8614fa8565b5b602002602001018190525080806112ff90614fd7565b9150506110a7565b5090565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600454905090565b6000600a54905090565b3373ffffffffffffffffffffffffffffffffffffffff1661136861130b565b73ffffffffffffffffffffffffffffffffffffffff16146113be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113b5906152b0565b60405180910390fd5b6001601260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600061140a43612bea565b815260200190815260200160002060006101000a81548160ff021916908315150217905550601460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009055601560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000905560008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a5885f1d846040518263ffffffff1660e01b815260040161150e9190614366565b600060405180830381600087803b15801561152857600080fd5b505af115801561153c573d6000803e3d6000fd5b5050505060008211156115d65761159282601360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054612eb2565b601360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b600081111561166d5760008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c905bb3584836040518363ffffffff1660e01b815260040161163a9291906150b1565b600060405180830381600087803b15801561165457600080fd5b505af1158015611668573d6000803e3d6000fd5b505050505b7f69284547cc931ff0e04d5a21cdfb0748f22a3788269711028ce1d4833900e47483601360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054836040516116df93929190615181565b60405180910390a1505050565b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461175a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161175190615091565b60405180910390fd5b61176343611bc8565b6117a2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161179990615342565b60405180910390fd5b6117ab43611d9d565b6117b6600c54611d9d565b106117f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117ed906153d4565b60405180910390fd5b43600c8190555060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600080600061183143612bea565b9050600061183e43610a4c565b9050600061184a611eb2565b905060005b8151811015611aba5781818151811061186b5761186a614fa8565b5b6020026020010151955061187e86612d42565b8061188f575061188e8685612e4a565b5b611aa7578215611a16576000601460008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050601460008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600090556000811115611a14576000600760008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905061199a8183612ecc565b6119d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119d090615466565b60405180910390fd5b7f9bae506d1374d366cdfa60105473d52dfdbaaae60e55de77bf6ee07f2add0cfb8883604051611a0a9291906150b1565b60405180910390a1505b505b601560008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205485611a61919061511c565b9450601560008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600090555b8080611ab290614fd7565b91505061184f565b508115611bb8578573ffffffffffffffffffffffffffffffffffffffff1663d45e6273826040518263ffffffff1660e01b8152600401611afa9190614979565b600060405180830381600087803b158015611b1457600080fd5b505af1158015611b28573d6000803e3d6000fd5b505050506000841115611bb757611b40866000612ecc565b611b7f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b76906154f8565b60405180910390fd5b7feb09b8cc1cefa77cd4ec30003e6364cf60afcedd20be8c09f26e717788baf13984604051611bae91906142b6565b60405180910390a15b5b611bc0612f83565b505050505050565b60006001600a54611bd99190614e82565b600a5483611be79190614ee5565b149050919050565b611bf7612c70565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611c64576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c5b90614f88565b60405180910390fd5b611c6d81613394565b50565b6000601060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b6000600c54905090565b611cd8612c70565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611d45576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d3c90614f88565b60405180910390fd5b611d4e816133d5565b50565b600080600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414159050919050565b60006001600a5483611daf9190615150565b611db9919061511c565b9050919050565b606060005b8351811015611e2a57611df2848281518110611de457611de3614fa8565b5b602002602001015184612e4a565b828281518110611e0557611e04614fa8565b5b6020026020010190151590811515815250508080611e2290614fd7565b915050611dc5565b5092915050565b611e39612c70565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611ea6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e9d90614f88565b60405180910390fd5b611eaf8161344f565b50565b6060600d5467ffffffffffffffff811115611ed057611ecf614397565b5b604051908082528060200260200182016040528015611efe5781602001602082028036833780820191505090505b50905060005b8151811015611fa157600e600082815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16828281518110611f5457611f53614fa8565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508080611f9990614fd7565b915050611f04565b5090565b6060600580548060200260200160405190810160405280929190818152602001828054801561202957602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611fdf575b5050505050905090565b6000600860019054906101000a900460ff1615905080801561206757506001600860009054906101000a900460ff1660ff16105b806120965750612076306134ca565b15801561209557506001600860009054906101000a900460ff1660ff16145b5b6120d5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120cc9061558a565b60405180910390fd5b6001600860006101000a81548160ff021916908360ff1602179055508015612113576001600860016101000a81548160ff0219169083151502179055505b61211c8a612cc7565b612125896133d5565b61212e8861344f565b61213787612d8e565b61214086613394565b61214985612e09565b612152846134ed565b61215b83613573565b612164826135b4565b80156121be576000600860016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516121b591906155fc565b60405180910390a15b50505050505050505050565b6000600954905090565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b612206612c70565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612273576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161226a90614f88565b60405180910390fd5b60008251036122b7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122ae90615663565b60405180910390fd5b80518251146122fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122f2906156f5565b60405180910390fd5b60005b82518110156124315781818151811061231a57612319614fa8565b5b602002602001015115156010600085848151811061233b5761233a614fa8565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615151461241e578181815181106123a5576123a4614fa8565b5b6020026020010151601060008584815181106123c4576123c3614fa8565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b808061242990614fd7565b9150506122fe565b507fa52c766fffd3af2ed65a8599973f96a0713c68566068cc057ad25cabd88ed4668282604051612463929190615715565b60405180910390a15050565b612477612c70565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146124e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016124db90614f88565b60405180910390fd5b6124ed816135b4565b50565b6124f8612c70565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612565576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161255c90614f88565b60405180910390fd5b61256e81613573565b50565b606060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff1663ce99b5866040518163ffffffff1660e01b8152600401602060405180830381865afa1580156125e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061260a91906150ef565b90508173ffffffffffffffffffffffffffffffffffffffff16634a5d76cd60056040518263ffffffff1660e01b81526004016126469190615838565b600060405180830381865afa158015612663573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061268c919061591d565b92506000600580549050905060005b8181101561276657828582815181106126b7576126b6614fa8565b5b602002602001015110156127535784826126d090615966565b925082815181106126e4576126e3614fa8565b5b60200260200101518582815181106126ff576126fe614fa8565b5b6020026020010181815250506127526005828154811061272257612721614fa8565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166135f5565b5b808061275e90614fd7565b91505061269b565b5080845250505090565b3373ffffffffffffffffffffffffffffffffffffffff1661278f612bb7565b73ffffffffffffffffffffffffffffffffffffffff16146127e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016127dc90615a01565b60405180910390fd5b600060058054905090506127f7611335565b8110612838576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161282f90615a93565b60405180910390fd5b61284184611d51565b15612881576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161287890615b25565b60405180910390fd5b8019600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506005849080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040518060a001604052808673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff168152602001838152602001600067ffffffffffffffff8111156129a9576129a8614397565b5b6040519080825280601f01601f1916602001820160405280156129db5781602001600182028036833780820191505090505b50815250600760008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550606082015181600301556080820151816004019081612b129190615ce7565b50905050600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f5ea0ccc37694ce2ce1e44e06663c8caff77c1ec661d991e2ece3a6195f879acf60405160405180910390a45050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000601154905090565b60006001600b54600a54612bfe9190614e28565b83612c099190615150565b612c13919061511c565b9050919050565b6000600f60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b6000612c9e7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610360001b61389c565b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507faa5b07dd43aa44c69b70a6a2b9c3fcfed12b6e5f6323596ba7ac91035ab80a4f81604051612d379190614366565b60405180910390a150565b6000601360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020544311159050919050565b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f31a33f126a5bae3c5bdf6cfc2cd6dcfffe2fe9634bdb09e21c44762993889e3b81604051612dfe9190614366565b60405180910390a150565b806004819055507f82d5dc32d1b741512ad09c32404d7e7921e8934c6222343d95f55f7a2b9b2ab481604051612e3f91906142b6565b60405180910390a150565b6000601260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600083815260200190815260200160002060009054906101000a900460ff16905092915050565b600081831015612ec25781612ec4565b825b905092915050565b600081471015612f11576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612f0890615e2b565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff1682604051612f3590615e7c565b60006040518083038185875af1925050503d8060008114612f72576040519150601f19603f3d011682016040523d82523d6000602084013e612f77565b606091505b50508091505092915050565b6000612f8d6138a6565b90506000612f9e6009548351613a53565b9050612faa8282613a6c565b80825260008190505b600d5481101561308857600e600082815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600f6000600e600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff0219169055808061308090614fd7565b915050612fb3565b50600080600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f0a46709856001436130d8919061511c565b6040518363ffffffff1660e01b81526004016130f5929190615e91565b600060405180830381865afa158015613112573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061313b9190615f6d565b905060005b8381101561334f5781818151811061315b5761315a614fa8565b5b602002602001015161333c57600085828151811061317c5761317b614fa8565b5b60200260200101519050600e600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036132005783806131f790614fd7565b9450505061333c565b600f6000600e600087815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff02191690556001600f60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555080600e600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838061333790614fd7565b945050505b808061334790614fd7565b915050613140565b5081600d819055507f6120448f3d4245c1ff708d970c34e1b6484ee22a794ede0bfca2317a97aa8ced846040516133869190614979565b60405180910390a150505050565b806009819055507fb5464c05fd0e0f000c535850116cda2742ee1f7b34384cb920ad7b8e802138b5816040516133ca91906142b6565b60405180910390a150565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f6397f5b135542bb3f477cb346cfab5abdec1251d08dc8f8d4efb4ffe122ea0bf816040516134449190614366565b60405180910390a150565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fc328090a37d855191ab58469296f98f87a851ca57d5cdfd1e9ac3c83e9e7096d816040516134bf9190614366565b60405180910390a150565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600954811115613532576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016135299061604e565b60405180910390fd5b806011819055507fa9588dc77416849bd922605ce4fc806712281ad8a8f32d4238d6c8cca548e15e8160405161356891906142b6565b60405180910390a150565b80600a819055507fbfd285a38b782d8a00e424fb824320ff3d1a698534358d02da611468d59b7808816040516135a991906142b6565b60405180910390a150565b80600b819055507f1d01baa2db15fced4f4e5fcfd4245e65ad9b083c110d26542f4a5f78d5425e77816040516135ea91906142b6565b60405180910390a150565b6000600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050600081036136475750613899565b600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556003820160009055600482016000613714919061413e565b5050600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000905560006005600160058054905061376f9190614e82565b815481106137805761377f614fa8565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555080600583198154811061380757613806614fa8565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060058054806138615761386061606e565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055905550505b50565b6000819050919050565b606060006138b2612571565b9050600580548060200260200160405190810160405280929190818152602001828054801561393657602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116138ec575b5050505050915060008251905060005b8351811015613a3b5761397284828151811061396557613964614fa8565b5b6020026020010151612d42565b15613a2857818061398290615966565b92505083828151811061399857613997614fa8565b5b60200260200101518482815181106139b3576139b2614fa8565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050828281518110613a00576139ff614fa8565b5b6020026020010151838281518110613a1b57613a1a614fa8565b5b6020026020010181815250505b8080613a3390614fd7565b915050613946565b50808352808252613a4c8383613ce3565b9250505090565b6000818310613a625781613a64565b825b905092915050565b6000825167ffffffffffffffff811115613a8957613a88614397565b5b604051908082528060200260200182016040528015613ab75781602001602082028036833780820191505090505b50905060008060005b8551811015613c415760106000878381518110613ae057613adf614fa8565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615613bb957601154821015613bb857858181518110613b5257613b51614fa8565b5b6020026020010151868380613b6690614fd7565b945081518110613b7957613b78614fa8565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050613c2e565b5b858181518110613bcc57613bcb614fa8565b5b6020026020010151848480613be090614fd7565b955081518110613bf357613bf2614fa8565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250505b8080613c3990614fd7565b915050613ac0565b506000915060008190505b84811015613cdb57838380613c6090614fd7565b945081518110613c7357613c72614fa8565b5b6020026020010151868281518110613c8e57613c8d614fa8565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508080613cd390614fd7565b915050613c4c565b505050505050565b60608251825114613d29576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613d20906160e9565b60405180910390fd5b6000835103613d3a57829050613edb565b6000835167ffffffffffffffff811115613d5757613d56614397565b5b604051908082528060200260200182016040528015613d9057816020015b613d7d61417e565b815260200190600190039081613d755790505b50905060005b8151811015613e2f576040518060400160405280868381518110613dbd57613dbc614fa8565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff168152602001858381518110613df357613df2614fa8565b5b6020026020010151815250828281518110613e1157613e10614fa8565b5b60200260200101819052508080613e2790614fd7565b915050613d96565b50613e4981600060018451613e449190614e82565b613ee1565b5060005b8151811015613ed557818181518110613e6957613e68614fa8565b5b602002602001015160000151858281518110613e8857613e87614fa8565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508080613ecd90614fd7565b915050613e4d565b50839150505b92915050565b606060008390506000839050808203613efe578592505050614099565b60008660028787613f0f9190616113565b613f199190616156565b87613f2491906161c0565b81518110613f3557613f34614fa8565b5b602002602001015190505b818313614066575b8060200151878481518110613f6057613f5f614fa8565b5b6020026020010151602001511115613f85578280613f7d90616204565b935050613f48565b5b868281518110613f9957613f98614fa8565b5b60200260200101516020015181602001511115613fc3578180613fbb9061624c565b925050613f86565b81831361406157614008878481518110613fe057613fdf614fa8565b5b6020026020010151888481518110613ffb57613ffa614fa8565b5b60200260200101516140a0565b88858151811061401b5761401a614fa8565b5b6020026020010189858151811061403557614034614fa8565b5b602002602001018290528290525050828061404f90616204565b935050818061405d9061624c565b9250505b613f40565b8186121561407c57614079878784613ee1565b96505b848312156140925761408f878487613ee1565b96505b8693505050505b9392505050565b6140a861417e565b6140b061417e565b600084905083818095508196505050848492509250509250929050565b6040518060a00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001606081525090565b50805461414a906151e7565b6000825580601f1061415c575061417b565b601f01602090049060005260206000209081019061417a9190614198565b5b50565b604051806040016040528060008152602001600081525090565b5b808211156141b1576000816000905550600101614199565b5090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006141f4826141c9565b9050919050565b614204816141e9565b811461420f57600080fd5b50565b600081359050614221816141fb565b92915050565b6000806040838503121561423e5761423d6141bf565b5b600061424c85828601614212565b925050602061425d85828601614212565b9150509250929050565b60008115159050919050565b61427c81614267565b82525050565b60006020820190506142976000830184614273565b92915050565b6000819050919050565b6142b08161429d565b82525050565b60006020820190506142cb60008301846142a7565b92915050565b6142da8161429d565b81146142e557600080fd5b50565b6000813590506142f7816142d1565b92915050565b600060208284031215614313576143126141bf565b5b6000614321848285016142e8565b91505092915050565b6000602082840312156143405761433f6141bf565b5b600061434e84828501614212565b91505092915050565b614360816141e9565b82525050565b600060208201905061437b6000830184614357565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6143cf82614386565b810181811067ffffffffffffffff821117156143ee576143ed614397565b5b80604052505050565b60006144016141b5565b905061440d82826143c6565b919050565b600067ffffffffffffffff82111561442d5761442c614397565b5b602082029050602081019050919050565b600080fd5b600061445661445184614412565b6143f7565b905080838252602082019050602084028301858111156144795761447861443e565b5b835b818110156144a2578061448e8882614212565b84526020840193505060208101905061447b565b5050509392505050565b600082601f8301126144c1576144c0614381565b5b81356144d1848260208601614443565b91505092915050565b6000602082840312156144f0576144ef6141bf565b5b600082013567ffffffffffffffff81111561450e5761450d6141c4565b5b61451a848285016144ac565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61455881614267565b82525050565b600061456a838361454f565b60208301905092915050565b6000602082019050919050565b600061458e82614523565b614598818561452e565b93506145a38361453f565b8060005b838110156145d45781516145bb888261455e565b97506145c683614576565b9250506001810190506145a7565b5085935050505092915050565b600060208201905081810360008301526145fb8184614583565b905092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b614638816141e9565b82525050565b6000614649826141c9565b9050919050565b6146598161463e565b82525050565b6146688161429d565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b838110156146a857808201518184015260208101905061468d565b60008484015250505050565b60006146bf8261466e565b6146c98185614679565b93506146d981856020860161468a565b6146e281614386565b840191505092915050565b600060a083016000830151614705600086018261462f565b506020830151614718602086018261462f565b50604083015161472b6040860182614650565b50606083015161473e606086018261465f565b506080830151848203608086015261475682826146b4565b9150508091505092915050565b600061476f83836146ed565b905092915050565b6000602082019050919050565b600061478f82614603565b614799818561460e565b9350836020820285016147ab8561461f565b8060005b858110156147e757848403895281516147c88582614763565b94506147d383614777565b925060208a019950506001810190506147af565b50829750879550505050505092915050565b600060208201905081810360008301526148138184614784565b905092915050565b600080600060608486031215614834576148336141bf565b5b600061484286828701614212565b9350506020614853868287016142e8565b9250506040614864868287016142e8565b9150509250925092565b60008060408385031215614885576148846141bf565b5b600083013567ffffffffffffffff8111156148a3576148a26141c4565b5b6148af858286016144ac565b92505060206148c0858286016142e8565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000614902838361462f565b60208301905092915050565b6000602082019050919050565b6000614926826148ca565b61493081856148d5565b935061493b836148e6565b8060005b8381101561496c57815161495388826148f6565b975061495e8361490e565b92505060018101905061493f565b5085935050505092915050565b60006020820190508181036000830152614993818461491b565b905092915050565b60008060008060008060008060006101208a8c0312156149be576149bd6141bf565b5b60006149cc8c828d01614212565b99505060206149dd8c828d01614212565b98505060406149ee8c828d01614212565b97505060606149ff8c828d01614212565b9650506080614a108c828d016142e8565b95505060a0614a218c828d016142e8565b94505060c0614a328c828d016142e8565b93505060e0614a438c828d016142e8565b925050610100614a558c828d016142e8565b9150509295985092959850929598565b600067ffffffffffffffff821115614a8057614a7f614397565b5b602082029050602081019050919050565b614a9a81614267565b8114614aa557600080fd5b50565b600081359050614ab781614a91565b92915050565b6000614ad0614acb84614a65565b6143f7565b90508083825260208201905060208402830185811115614af357614af261443e565b5b835b81811015614b1c5780614b088882614aa8565b845260208401935050602081019050614af5565b5050509392505050565b600082601f830112614b3b57614b3a614381565b5b8135614b4b848260208601614abd565b91505092915050565b60008060408385031215614b6b57614b6a6141bf565b5b600083013567ffffffffffffffff811115614b8957614b886141c4565b5b614b95858286016144ac565b925050602083013567ffffffffffffffff811115614bb657614bb56141c4565b5b614bc285828601614b26565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000614c04838361465f565b60208301905092915050565b6000602082019050919050565b6000614c2882614bcc565b614c328185614bd7565b9350614c3d83614be8565b8060005b83811015614c6e578151614c558882614bf8565b9750614c6083614c10565b925050600181019050614c41565b5085935050505092915050565b60006020820190508181036000830152614c958184614c1d565b905092915050565b614ca68161463e565b8114614cb157600080fd5b50565b600081359050614cc381614c9d565b92915050565b60008060008060808587031215614ce357614ce26141bf565b5b6000614cf187828801614212565b9450506020614d0287828801614212565b9350506040614d1387828801614cb4565b9250506060614d24878288016142e8565b91505092959194509250565b600082825260208201905092915050565b7f526f6e696e56616c696461746f725365743a206f6e6c7920726563656976657360008201527f20524f4e2066726f6d207374616b696e672076657374696e6720636f6e74726160208201527f6374000000000000000000000000000000000000000000000000000000000000604082015250565b6000614dc3604283614d30565b9150614dce82614d41565b606082019050919050565b60006020820190508181036000830152614df281614db6565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000614e338261429d565b9150614e3e8361429d565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615614e7757614e76614df9565b5b828202905092915050565b6000614e8d8261429d565b9150614e988361429d565b9250828203905081811115614eb057614eaf614df9565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000614ef08261429d565b9150614efb8361429d565b925082614f0b57614f0a614eb6565b5b828206905092915050565b7f48617350726f787941646d696e3a20756e617574686f72697a65642073656e6460008201527f6572000000000000000000000000000000000000000000000000000000000000602082015250565b6000614f72602283614d30565b9150614f7d82614f16565b604082019050919050565b60006020820190508181036000830152614fa181614f65565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000614fe28261429d565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361501457615013614df9565b5b600182019050919050565b7f526f6e696e56616c696461746f725365743a206d6574686f642063616c6c657260008201527f206d75737420626520636f696e62617365000000000000000000000000000000602082015250565b600061507b603183614d30565b91506150868261501f565b604082019050919050565b600060208201905081810360008301526150aa8161506e565b9050919050565b60006040820190506150c66000830185614357565b6150d360208301846142a7565b9392505050565b6000815190506150e9816142d1565b92915050565b600060208284031215615105576151046141bf565b5b6000615113848285016150da565b91505092915050565b60006151278261429d565b91506151328361429d565b925082820190508082111561514a57615149614df9565b5b92915050565b600061515b8261429d565b91506151668361429d565b92508261517657615175614eb6565b5b828204905092915050565b60006060820190506151966000830186614357565b6151a360208301856142a7565b6151b060408301846142a7565b949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806151ff57607f821691505b602082108103615212576152116151b8565b5b50919050565b7f486173536c617368496e64696361746f72436f6e74726163743a206d6574686f60008201527f642063616c6c6572206d75737420626520736c61736820696e64696361746f7260208201527f20636f6e74726163740000000000000000000000000000000000000000000000604082015250565b600061529a604983614d30565b91506152a582615218565b606082019050919050565b600060208201905081810360008301526152c98161528d565b9050919050565b7f526f6e696e56616c696461746f725365743a206f6e6c7920616c6c6f7765642060008201527f61742074686520656e64206f662065706f636800000000000000000000000000602082015250565b600061532c603383614d30565b9150615337826152d0565b604082019050919050565b6000602082019050818103600083015261535b8161531f565b9050919050565b7f526f6e696e56616c696461746f725365743a20717565727920666f7220616c7260008201527f6561647920777261707065642075702065706f63680000000000000000000000602082015250565b60006153be603583614d30565b91506153c982615362565b604082019050919050565b600060208201905081810360008301526153ed816153b1565b9050919050565b7f526f6e696e56616c696461746f725365743a20636f756c64206e6f742074726160008201527f6e7366657220524f4e2074726561737572792061646472657373000000000000602082015250565b6000615450603a83614d30565b915061545b826153f4565b604082019050919050565b6000602082019050818103600083015261547f81615443565b9050919050565b7f526f6e696e56616c696461746f725365743a20636f756c64206e6f742074726160008201527f6e7366657220524f4e20746f207374616b696e6720636f6e7472616374000000602082015250565b60006154e2603d83614d30565b91506154ed82615486565b604082019050919050565b60006020820190508181036000830152615511816154d5565b9050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b6000615574602e83614d30565b915061557f82615518565b604082019050919050565b600060208201905081810360008301526155a381615567565b9050919050565b6000819050919050565b600060ff82169050919050565b6000819050919050565b60006155e66155e16155dc846155aa565b6155c1565b6155b4565b9050919050565b6155f6816155cb565b82525050565b600060208201905061561160008301846155ed565b92915050565b7f526f6e696e56616c696461746f725365743a20656d7074792061727261790000600082015250565b600061564d601e83614d30565b915061565882615617565b602082019050919050565b6000602082019050818103600083015261567c81615640565b9050919050565b7f526f6e696e56616c696461746f725365743a206c656e677468206f662074776f60008201527f20696e70757420617272617973206d69736d6174636865730000000000000000602082015250565b60006156df603883614d30565b91506156ea82615683565b604082019050919050565b6000602082019050818103600083015261570e816156d2565b9050919050565b6000604082019050818103600083015261572f818561491b565b905081810360208301526157438184614583565b90509392505050565b600081549050919050565b60008190508160005260206000209050919050565b60008160001c9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006157ac6157a78361576c565b615779565b9050919050565b60006157bf8254615799565b9050919050565b6000600182019050919050565b60006157de8261574c565b6157e881856148d5565b93506157f383615757565b8060005b8381101561582b57615808826157b3565b61581288826148f6565b975061581d836157c6565b9250506001810190506157f7565b5085935050505092915050565b6000602082019050818103600083015261585281846157d3565b905092915050565b600067ffffffffffffffff82111561587557615874614397565b5b602082029050602081019050919050565b60006158996158948461585a565b6143f7565b905080838252602082019050602084028301858111156158bc576158bb61443e565b5b835b818110156158e557806158d188826150da565b8452602084019350506020810190506158be565b5050509392505050565b600082601f83011261590457615903614381565b5b8151615914848260208601615886565b91505092915050565b600060208284031215615933576159326141bf565b5b600082015167ffffffffffffffff811115615951576159506141c4565b5b61595d848285016158ef565b91505092915050565b60006159718261429d565b91506000820361598457615983614df9565b5b600182039050919050565b7f4861735374616b696e674d616e616765723a206d6574686f642063616c6c657260008201527f206d757374206265207374616b696e6720636f6e747261637400000000000000602082015250565b60006159eb603983614d30565b91506159f68261598f565b604082019050919050565b60006020820190508181036000830152615a1a816159de565b9050919050565b7f43616e6469646174654d616e616765723a2065786365656473206d6178696d7560008201527f6d206e756d626572206f662063616e6469646174657300000000000000000000602082015250565b6000615a7d603683614d30565b9150615a8882615a21565b604082019050919050565b60006020820190508181036000830152615aac81615a70565b9050919050565b7f43616e6469646174654d616e616765723a20717565727920666f7220616c726560008201527f616479206578697374656e742063616e64696461746500000000000000000000602082015250565b6000615b0f603683614d30565b9150615b1a82615ab3565b604082019050919050565b60006020820190508181036000830152615b3e81615b02565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302615ba77fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82615b6a565b615bb18683615b6a565b95508019841693508086168417925050509392505050565b6000615be4615bdf615bda8461429d565b6155c1565b61429d565b9050919050565b6000819050919050565b615bfe83615bc9565b615c12615c0a82615beb565b848454615b77565b825550505050565b600090565b615c27615c1a565b615c32818484615bf5565b505050565b5b81811015615c5657615c4b600082615c1f565b600181019050615c38565b5050565b601f821115615c9b57615c6c81615b45565b615c7584615b5a565b81016020851015615c84578190505b615c98615c9085615b5a565b830182615c37565b50505b505050565b600082821c905092915050565b6000615cbe60001984600802615ca0565b1980831691505092915050565b6000615cd78383615cad565b9150826002028217905092915050565b615cf08261466e565b67ffffffffffffffff811115615d0957615d08614397565b5b615d1382546151e7565b615d1e828285615c5a565b600060209050601f831160018114615d515760008415615d3f578287015190505b615d498582615ccb565b865550615db1565b601f198416615d5f86615b45565b60005b82811015615d8757848901518255600182019150602085019450602081019050615d62565b86831015615da45784890151615da0601f891682615cad565b8355505b6001600288020188555050505b505050505050565b7f524f4e5472616e736665723a20696e73756666696369656e742062616c616e6360008201527f6500000000000000000000000000000000000000000000000000000000000000602082015250565b6000615e15602183614d30565b9150615e2082615db9565b604082019050919050565b60006020820190508181036000830152615e4481615e08565b9050919050565b600081905092915050565b50565b6000615e66600083615e4b565b9150615e7182615e56565b600082019050919050565b6000615e8782615e59565b9150819050919050565b60006040820190508181036000830152615eab818561491b565b9050615eba60208301846142a7565b9392505050565b600081519050615ed081614a91565b92915050565b6000615ee9615ee484614a65565b6143f7565b90508083825260208201905060208402830185811115615f0c57615f0b61443e565b5b835b81811015615f355780615f218882615ec1565b845260208401935050602081019050615f0e565b5050509392505050565b600082601f830112615f5457615f53614381565b5b8151615f64848260208601615ed6565b91505092915050565b600060208284031215615f8357615f826141bf565b5b600082015167ffffffffffffffff811115615fa157615fa06141c4565b5b615fad84828501615f3f565b91505092915050565b7f526f6e696e56616c696461746f725365743a2063616e6e6f7420736574206e7560008201527f6d626572206f66207072696f726974697a65642067726561746572207468616e60208201527f206e756d626572206f66206d61782076616c696461746f727300000000000000604082015250565b6000616038605983614d30565b915061604382615fb6565b606082019050919050565b600060208201905081810360008301526160678161602b565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b7f536f7274696e673a20696e76616c6964206172726179206c656e677468000000600082015250565b60006160d3601d83614d30565b91506160de8261609d565b602082019050919050565b60006020820190508181036000830152616102816160c6565b9050919050565b6000819050919050565b600061611e82616109565b915061612983616109565b92508282039050818112600084121682821360008512151617156161505761614f614df9565b5b92915050565b600061616182616109565b915061616c83616109565b92508261617c5761617b614eb6565b5b600160000383147f8000000000000000000000000000000000000000000000000000000000000000831416156161b5576161b4614df9565b5b828205905092915050565b60006161cb82616109565b91506161d683616109565b9250828201905082811215600083121683821260008412151617156161fe576161fd614df9565b5b92915050565b600061620f82616109565b91507f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361624157616240614df9565b5b600182019050919050565b600061625782616109565b91507f8000000000000000000000000000000000000000000000000000000000000000820361628957616288614df9565b5b60018203905091905056fea26469706673582212203680cab58b9c20f6796a239eb78e31c55056236c013ebbadaa9a3f08059d033c64736f6c63430008100033",
}

// RoninValidatorSetABI is the input ABI used to generate the binding from.
// Deprecated: Use RoninValidatorSetMetaData.ABI instead.
var RoninValidatorSetABI = RoninValidatorSetMetaData.ABI

// RoninValidatorSetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RoninValidatorSetMetaData.Bin instead.
var RoninValidatorSetBin = RoninValidatorSetMetaData.Bin

// DeployRoninValidatorSet deploys a new Ethereum contract, binding an instance of RoninValidatorSet to it.
func DeployRoninValidatorSet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RoninValidatorSet, error) {
	parsed, err := RoninValidatorSetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RoninValidatorSetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RoninValidatorSet{RoninValidatorSetCaller: RoninValidatorSetCaller{contract: contract}, RoninValidatorSetTransactor: RoninValidatorSetTransactor{contract: contract}, RoninValidatorSetFilterer: RoninValidatorSetFilterer{contract: contract}}, nil
}

// RoninValidatorSet is an auto generated Go binding around an Ethereum contract.
type RoninValidatorSet struct {
	RoninValidatorSetCaller     // Read-only binding to the contract
	RoninValidatorSetTransactor // Write-only binding to the contract
	RoninValidatorSetFilterer   // Log filterer for contract events
}

// RoninValidatorSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoninValidatorSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoninValidatorSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoninValidatorSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoninValidatorSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoninValidatorSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoninValidatorSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoninValidatorSetSession struct {
	Contract     *RoninValidatorSet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RoninValidatorSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoninValidatorSetCallerSession struct {
	Contract *RoninValidatorSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// RoninValidatorSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoninValidatorSetTransactorSession struct {
	Contract     *RoninValidatorSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// RoninValidatorSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoninValidatorSetRaw struct {
	Contract *RoninValidatorSet // Generic contract binding to access the raw methods on
}

// RoninValidatorSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoninValidatorSetCallerRaw struct {
	Contract *RoninValidatorSetCaller // Generic read-only contract binding to access the raw methods on
}

// RoninValidatorSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoninValidatorSetTransactorRaw struct {
	Contract *RoninValidatorSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoninValidatorSet creates a new instance of RoninValidatorSet, bound to a specific deployed contract.
func NewRoninValidatorSet(address common.Address, backend bind.ContractBackend) (*RoninValidatorSet, error) {
	contract, err := bindRoninValidatorSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSet{RoninValidatorSetCaller: RoninValidatorSetCaller{contract: contract}, RoninValidatorSetTransactor: RoninValidatorSetTransactor{contract: contract}, RoninValidatorSetFilterer: RoninValidatorSetFilterer{contract: contract}}, nil
}

// NewRoninValidatorSetCaller creates a new read-only instance of RoninValidatorSet, bound to a specific deployed contract.
func NewRoninValidatorSetCaller(address common.Address, caller bind.ContractCaller) (*RoninValidatorSetCaller, error) {
	contract, err := bindRoninValidatorSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetCaller{contract: contract}, nil
}

// NewRoninValidatorSetTransactor creates a new write-only instance of RoninValidatorSet, bound to a specific deployed contract.
func NewRoninValidatorSetTransactor(address common.Address, transactor bind.ContractTransactor) (*RoninValidatorSetTransactor, error) {
	contract, err := bindRoninValidatorSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetTransactor{contract: contract}, nil
}

// NewRoninValidatorSetFilterer creates a new log filterer instance of RoninValidatorSet, bound to a specific deployed contract.
func NewRoninValidatorSetFilterer(address common.Address, filterer bind.ContractFilterer) (*RoninValidatorSetFilterer, error) {
	contract, err := bindRoninValidatorSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetFilterer{contract: contract}, nil
}

// bindRoninValidatorSet binds a generic wrapper to an already deployed contract.
func bindRoninValidatorSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RoninValidatorSetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoninValidatorSet *RoninValidatorSetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoninValidatorSet.Contract.RoninValidatorSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoninValidatorSet *RoninValidatorSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.RoninValidatorSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoninValidatorSet *RoninValidatorSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.RoninValidatorSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoninValidatorSet *RoninValidatorSetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoninValidatorSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoninValidatorSet *RoninValidatorSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoninValidatorSet *RoninValidatorSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.contract.Transact(opts, method, params...)
}

// EpochEndingAt is a free data retrieval call binding the contract method 0x7593ff71.
//
// Solidity: function epochEndingAt(uint256 _block) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCaller) EpochEndingAt(opts *bind.CallOpts, _block *big.Int) (bool, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "epochEndingAt", _block)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EpochEndingAt is a free data retrieval call binding the contract method 0x7593ff71.
//
// Solidity: function epochEndingAt(uint256 _block) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetSession) EpochEndingAt(_block *big.Int) (bool, error) {
	return _RoninValidatorSet.Contract.EpochEndingAt(&_RoninValidatorSet.CallOpts, _block)
}

// EpochEndingAt is a free data retrieval call binding the contract method 0x7593ff71.
//
// Solidity: function epochEndingAt(uint256 _block) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) EpochEndingAt(_block *big.Int) (bool, error) {
	return _RoninValidatorSet.Contract.EpochEndingAt(&_RoninValidatorSet.CallOpts, _block)
}

// EpochOf is a free data retrieval call binding the contract method 0xa3d545f5.
//
// Solidity: function epochOf(uint256 _block) view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCaller) EpochOf(opts *bind.CallOpts, _block *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "epochOf", _block)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochOf is a free data retrieval call binding the contract method 0xa3d545f5.
//
// Solidity: function epochOf(uint256 _block) view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetSession) EpochOf(_block *big.Int) (*big.Int, error) {
	return _RoninValidatorSet.Contract.EpochOf(&_RoninValidatorSet.CallOpts, _block)
}

// EpochOf is a free data retrieval call binding the contract method 0xa3d545f5.
//
// Solidity: function epochOf(uint256 _block) view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) EpochOf(_block *big.Int) (*big.Int, error) {
	return _RoninValidatorSet.Contract.EpochOf(&_RoninValidatorSet.CallOpts, _block)
}

// GetCandidateInfos is a free data retrieval call binding the contract method 0x5248184a.
//
// Solidity: function getCandidateInfos() view returns((address,address,address,uint256,bytes)[] _list)
func (_RoninValidatorSet *RoninValidatorSetCaller) GetCandidateInfos(opts *bind.CallOpts) ([]ICandidateManagerValidatorCandidate, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "getCandidateInfos")

	if err != nil {
		return *new([]ICandidateManagerValidatorCandidate), err
	}

	out0 := *abi.ConvertType(out[0], new([]ICandidateManagerValidatorCandidate)).(*[]ICandidateManagerValidatorCandidate)

	return out0, err

}

// GetCandidateInfos is a free data retrieval call binding the contract method 0x5248184a.
//
// Solidity: function getCandidateInfos() view returns((address,address,address,uint256,bytes)[] _list)
func (_RoninValidatorSet *RoninValidatorSetSession) GetCandidateInfos() ([]ICandidateManagerValidatorCandidate, error) {
	return _RoninValidatorSet.Contract.GetCandidateInfos(&_RoninValidatorSet.CallOpts)
}

// GetCandidateInfos is a free data retrieval call binding the contract method 0x5248184a.
//
// Solidity: function getCandidateInfos() view returns((address,address,address,uint256,bytes)[] _list)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) GetCandidateInfos() ([]ICandidateManagerValidatorCandidate, error) {
	return _RoninValidatorSet.Contract.GetCandidateInfos(&_RoninValidatorSet.CallOpts)
}

// GetLastUpdatedBlock is a free data retrieval call binding the contract method 0x87c891bd.
//
// Solidity: function getLastUpdatedBlock() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCaller) GetLastUpdatedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "getLastUpdatedBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastUpdatedBlock is a free data retrieval call binding the contract method 0x87c891bd.
//
// Solidity: function getLastUpdatedBlock() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetSession) GetLastUpdatedBlock() (*big.Int, error) {
	return _RoninValidatorSet.Contract.GetLastUpdatedBlock(&_RoninValidatorSet.CallOpts)
}

// GetLastUpdatedBlock is a free data retrieval call binding the contract method 0x87c891bd.
//
// Solidity: function getLastUpdatedBlock() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) GetLastUpdatedBlock() (*big.Int, error) {
	return _RoninValidatorSet.Contract.GetLastUpdatedBlock(&_RoninValidatorSet.CallOpts)
}

// GetPriorityStatus is a free data retrieval call binding the contract method 0x86523193.
//
// Solidity: function getPriorityStatus(address _addr) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCaller) GetPriorityStatus(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "getPriorityStatus", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetPriorityStatus is a free data retrieval call binding the contract method 0x86523193.
//
// Solidity: function getPriorityStatus(address _addr) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetSession) GetPriorityStatus(_addr common.Address) (bool, error) {
	return _RoninValidatorSet.Contract.GetPriorityStatus(&_RoninValidatorSet.CallOpts, _addr)
}

// GetPriorityStatus is a free data retrieval call binding the contract method 0x86523193.
//
// Solidity: function getPriorityStatus(address _addr) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) GetPriorityStatus(_addr common.Address) (bool, error) {
	return _RoninValidatorSet.Contract.GetPriorityStatus(&_RoninValidatorSet.CallOpts, _addr)
}

// GetValidatorCandidates is a free data retrieval call binding the contract method 0xba77b06c.
//
// Solidity: function getValidatorCandidates() view returns(address[])
func (_RoninValidatorSet *RoninValidatorSetCaller) GetValidatorCandidates(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "getValidatorCandidates")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidatorCandidates is a free data retrieval call binding the contract method 0xba77b06c.
//
// Solidity: function getValidatorCandidates() view returns(address[])
func (_RoninValidatorSet *RoninValidatorSetSession) GetValidatorCandidates() ([]common.Address, error) {
	return _RoninValidatorSet.Contract.GetValidatorCandidates(&_RoninValidatorSet.CallOpts)
}

// GetValidatorCandidates is a free data retrieval call binding the contract method 0xba77b06c.
//
// Solidity: function getValidatorCandidates() view returns(address[])
func (_RoninValidatorSet *RoninValidatorSetCallerSession) GetValidatorCandidates() ([]common.Address, error) {
	return _RoninValidatorSet.Contract.GetValidatorCandidates(&_RoninValidatorSet.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[] _validatorList)
func (_RoninValidatorSet *RoninValidatorSetCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[] _validatorList)
func (_RoninValidatorSet *RoninValidatorSetSession) GetValidators() ([]common.Address, error) {
	return _RoninValidatorSet.Contract.GetValidators(&_RoninValidatorSet.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[] _validatorList)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) GetValidators() ([]common.Address, error) {
	return _RoninValidatorSet.Contract.GetValidators(&_RoninValidatorSet.CallOpts)
}

// IsCandidateAdmin is a free data retrieval call binding the contract method 0x04d971ab.
//
// Solidity: function isCandidateAdmin(address _candidate, address _admin) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCaller) IsCandidateAdmin(opts *bind.CallOpts, _candidate common.Address, _admin common.Address) (bool, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "isCandidateAdmin", _candidate, _admin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCandidateAdmin is a free data retrieval call binding the contract method 0x04d971ab.
//
// Solidity: function isCandidateAdmin(address _candidate, address _admin) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetSession) IsCandidateAdmin(_candidate common.Address, _admin common.Address) (bool, error) {
	return _RoninValidatorSet.Contract.IsCandidateAdmin(&_RoninValidatorSet.CallOpts, _candidate, _admin)
}

// IsCandidateAdmin is a free data retrieval call binding the contract method 0x04d971ab.
//
// Solidity: function isCandidateAdmin(address _candidate, address _admin) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) IsCandidateAdmin(_candidate common.Address, _admin common.Address) (bool, error) {
	return _RoninValidatorSet.Contract.IsCandidateAdmin(&_RoninValidatorSet.CallOpts, _candidate, _admin)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _addr) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCaller) IsValidator(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "isValidator", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _addr) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetSession) IsValidator(_addr common.Address) (bool, error) {
	return _RoninValidatorSet.Contract.IsValidator(&_RoninValidatorSet.CallOpts, _addr)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _addr) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) IsValidator(_addr common.Address) (bool, error) {
	return _RoninValidatorSet.Contract.IsValidator(&_RoninValidatorSet.CallOpts, _addr)
}

// IsValidatorCandidate is a free data retrieval call binding the contract method 0xa0c3f2d2.
//
// Solidity: function isValidatorCandidate(address _addr) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCaller) IsValidatorCandidate(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "isValidatorCandidate", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorCandidate is a free data retrieval call binding the contract method 0xa0c3f2d2.
//
// Solidity: function isValidatorCandidate(address _addr) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetSession) IsValidatorCandidate(_addr common.Address) (bool, error) {
	return _RoninValidatorSet.Contract.IsValidatorCandidate(&_RoninValidatorSet.CallOpts, _addr)
}

// IsValidatorCandidate is a free data retrieval call binding the contract method 0xa0c3f2d2.
//
// Solidity: function isValidatorCandidate(address _addr) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) IsValidatorCandidate(_addr common.Address) (bool, error) {
	return _RoninValidatorSet.Contract.IsValidatorCandidate(&_RoninValidatorSet.CallOpts, _addr)
}

// Jailed is a free data retrieval call binding the contract method 0x4454af9d.
//
// Solidity: function jailed(address[] _addrList) view returns(bool[] _result)
func (_RoninValidatorSet *RoninValidatorSetCaller) Jailed(opts *bind.CallOpts, _addrList []common.Address) ([]bool, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "jailed", _addrList)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// Jailed is a free data retrieval call binding the contract method 0x4454af9d.
//
// Solidity: function jailed(address[] _addrList) view returns(bool[] _result)
func (_RoninValidatorSet *RoninValidatorSetSession) Jailed(_addrList []common.Address) ([]bool, error) {
	return _RoninValidatorSet.Contract.Jailed(&_RoninValidatorSet.CallOpts, _addrList)
}

// Jailed is a free data retrieval call binding the contract method 0x4454af9d.
//
// Solidity: function jailed(address[] _addrList) view returns(bool[] _result)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) Jailed(_addrList []common.Address) ([]bool, error) {
	return _RoninValidatorSet.Contract.Jailed(&_RoninValidatorSet.CallOpts, _addrList)
}

// MaintenanceContract is a free data retrieval call binding the contract method 0xd2cb215e.
//
// Solidity: function maintenanceContract() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetCaller) MaintenanceContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "maintenanceContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MaintenanceContract is a free data retrieval call binding the contract method 0xd2cb215e.
//
// Solidity: function maintenanceContract() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetSession) MaintenanceContract() (common.Address, error) {
	return _RoninValidatorSet.Contract.MaintenanceContract(&_RoninValidatorSet.CallOpts)
}

// MaintenanceContract is a free data retrieval call binding the contract method 0xd2cb215e.
//
// Solidity: function maintenanceContract() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) MaintenanceContract() (common.Address, error) {
	return _RoninValidatorSet.Contract.MaintenanceContract(&_RoninValidatorSet.CallOpts)
}

// MaxPrioritizedValidatorNumber is a free data retrieval call binding the contract method 0xeeb629a8.
//
// Solidity: function maxPrioritizedValidatorNumber() view returns(uint256 _maximumPrioritizedValidatorNumber)
func (_RoninValidatorSet *RoninValidatorSetCaller) MaxPrioritizedValidatorNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "maxPrioritizedValidatorNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPrioritizedValidatorNumber is a free data retrieval call binding the contract method 0xeeb629a8.
//
// Solidity: function maxPrioritizedValidatorNumber() view returns(uint256 _maximumPrioritizedValidatorNumber)
func (_RoninValidatorSet *RoninValidatorSetSession) MaxPrioritizedValidatorNumber() (*big.Int, error) {
	return _RoninValidatorSet.Contract.MaxPrioritizedValidatorNumber(&_RoninValidatorSet.CallOpts)
}

// MaxPrioritizedValidatorNumber is a free data retrieval call binding the contract method 0xeeb629a8.
//
// Solidity: function maxPrioritizedValidatorNumber() view returns(uint256 _maximumPrioritizedValidatorNumber)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) MaxPrioritizedValidatorNumber() (*big.Int, error) {
	return _RoninValidatorSet.Contract.MaxPrioritizedValidatorNumber(&_RoninValidatorSet.CallOpts)
}

// MaxValidatorCandidate is a free data retrieval call binding the contract method 0x605239a1.
//
// Solidity: function maxValidatorCandidate() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCaller) MaxValidatorCandidate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "maxValidatorCandidate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxValidatorCandidate is a free data retrieval call binding the contract method 0x605239a1.
//
// Solidity: function maxValidatorCandidate() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetSession) MaxValidatorCandidate() (*big.Int, error) {
	return _RoninValidatorSet.Contract.MaxValidatorCandidate(&_RoninValidatorSet.CallOpts)
}

// MaxValidatorCandidate is a free data retrieval call binding the contract method 0x605239a1.
//
// Solidity: function maxValidatorCandidate() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) MaxValidatorCandidate() (*big.Int, error) {
	return _RoninValidatorSet.Contract.MaxValidatorCandidate(&_RoninValidatorSet.CallOpts)
}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() view returns(uint256 _maximumValidatorNumber)
func (_RoninValidatorSet *RoninValidatorSetCaller) MaxValidatorNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "maxValidatorNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() view returns(uint256 _maximumValidatorNumber)
func (_RoninValidatorSet *RoninValidatorSetSession) MaxValidatorNumber() (*big.Int, error) {
	return _RoninValidatorSet.Contract.MaxValidatorNumber(&_RoninValidatorSet.CallOpts)
}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() view returns(uint256 _maximumValidatorNumber)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) MaxValidatorNumber() (*big.Int, error) {
	return _RoninValidatorSet.Contract.MaxValidatorNumber(&_RoninValidatorSet.CallOpts)
}

// NumberOfBlocksInEpoch is a free data retrieval call binding the contract method 0x6aa1c2ef.
//
// Solidity: function numberOfBlocksInEpoch() view returns(uint256 _numberOfBlocks)
func (_RoninValidatorSet *RoninValidatorSetCaller) NumberOfBlocksInEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "numberOfBlocksInEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberOfBlocksInEpoch is a free data retrieval call binding the contract method 0x6aa1c2ef.
//
// Solidity: function numberOfBlocksInEpoch() view returns(uint256 _numberOfBlocks)
func (_RoninValidatorSet *RoninValidatorSetSession) NumberOfBlocksInEpoch() (*big.Int, error) {
	return _RoninValidatorSet.Contract.NumberOfBlocksInEpoch(&_RoninValidatorSet.CallOpts)
}

// NumberOfBlocksInEpoch is a free data retrieval call binding the contract method 0x6aa1c2ef.
//
// Solidity: function numberOfBlocksInEpoch() view returns(uint256 _numberOfBlocks)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) NumberOfBlocksInEpoch() (*big.Int, error) {
	return _RoninValidatorSet.Contract.NumberOfBlocksInEpoch(&_RoninValidatorSet.CallOpts)
}

// NumberOfEpochsInPeriod is a free data retrieval call binding the contract method 0x5186dc7e.
//
// Solidity: function numberOfEpochsInPeriod() view returns(uint256 _numberOfEpochs)
func (_RoninValidatorSet *RoninValidatorSetCaller) NumberOfEpochsInPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "numberOfEpochsInPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberOfEpochsInPeriod is a free data retrieval call binding the contract method 0x5186dc7e.
//
// Solidity: function numberOfEpochsInPeriod() view returns(uint256 _numberOfEpochs)
func (_RoninValidatorSet *RoninValidatorSetSession) NumberOfEpochsInPeriod() (*big.Int, error) {
	return _RoninValidatorSet.Contract.NumberOfEpochsInPeriod(&_RoninValidatorSet.CallOpts)
}

// NumberOfEpochsInPeriod is a free data retrieval call binding the contract method 0x5186dc7e.
//
// Solidity: function numberOfEpochsInPeriod() view returns(uint256 _numberOfEpochs)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) NumberOfEpochsInPeriod() (*big.Int, error) {
	return _RoninValidatorSet.Contract.NumberOfEpochsInPeriod(&_RoninValidatorSet.CallOpts)
}

// PeriodEndingAt is a free data retrieval call binding the contract method 0x25a6b529.
//
// Solidity: function periodEndingAt(uint256 _block) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCaller) PeriodEndingAt(opts *bind.CallOpts, _block *big.Int) (bool, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "periodEndingAt", _block)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PeriodEndingAt is a free data retrieval call binding the contract method 0x25a6b529.
//
// Solidity: function periodEndingAt(uint256 _block) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetSession) PeriodEndingAt(_block *big.Int) (bool, error) {
	return _RoninValidatorSet.Contract.PeriodEndingAt(&_RoninValidatorSet.CallOpts, _block)
}

// PeriodEndingAt is a free data retrieval call binding the contract method 0x25a6b529.
//
// Solidity: function periodEndingAt(uint256 _block) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) PeriodEndingAt(_block *big.Int) (bool, error) {
	return _RoninValidatorSet.Contract.PeriodEndingAt(&_RoninValidatorSet.CallOpts, _block)
}

// PeriodOf is a free data retrieval call binding the contract method 0xf8549af9.
//
// Solidity: function periodOf(uint256 _block) view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCaller) PeriodOf(opts *bind.CallOpts, _block *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "periodOf", _block)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodOf is a free data retrieval call binding the contract method 0xf8549af9.
//
// Solidity: function periodOf(uint256 _block) view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetSession) PeriodOf(_block *big.Int) (*big.Int, error) {
	return _RoninValidatorSet.Contract.PeriodOf(&_RoninValidatorSet.CallOpts, _block)
}

// PeriodOf is a free data retrieval call binding the contract method 0xf8549af9.
//
// Solidity: function periodOf(uint256 _block) view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) PeriodOf(_block *big.Int) (*big.Int, error) {
	return _RoninValidatorSet.Contract.PeriodOf(&_RoninValidatorSet.CallOpts, _block)
}

// RewardDeprecated is a free data retrieval call binding the contract method 0xac00125f.
//
// Solidity: function rewardDeprecated(address[] _addrList, uint256 _period) view returns(bool[] _result)
func (_RoninValidatorSet *RoninValidatorSetCaller) RewardDeprecated(opts *bind.CallOpts, _addrList []common.Address, _period *big.Int) ([]bool, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "rewardDeprecated", _addrList, _period)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// RewardDeprecated is a free data retrieval call binding the contract method 0xac00125f.
//
// Solidity: function rewardDeprecated(address[] _addrList, uint256 _period) view returns(bool[] _result)
func (_RoninValidatorSet *RoninValidatorSetSession) RewardDeprecated(_addrList []common.Address, _period *big.Int) ([]bool, error) {
	return _RoninValidatorSet.Contract.RewardDeprecated(&_RoninValidatorSet.CallOpts, _addrList, _period)
}

// RewardDeprecated is a free data retrieval call binding the contract method 0xac00125f.
//
// Solidity: function rewardDeprecated(address[] _addrList, uint256 _period) view returns(bool[] _result)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) RewardDeprecated(_addrList []common.Address, _period *big.Int) ([]bool, error) {
	return _RoninValidatorSet.Contract.RewardDeprecated(&_RoninValidatorSet.CallOpts, _addrList, _period)
}

// SlashIndicatorContract is a free data retrieval call binding the contract method 0x5a08482d.
//
// Solidity: function slashIndicatorContract() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetCaller) SlashIndicatorContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "slashIndicatorContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SlashIndicatorContract is a free data retrieval call binding the contract method 0x5a08482d.
//
// Solidity: function slashIndicatorContract() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetSession) SlashIndicatorContract() (common.Address, error) {
	return _RoninValidatorSet.Contract.SlashIndicatorContract(&_RoninValidatorSet.CallOpts)
}

// SlashIndicatorContract is a free data retrieval call binding the contract method 0x5a08482d.
//
// Solidity: function slashIndicatorContract() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) SlashIndicatorContract() (common.Address, error) {
	return _RoninValidatorSet.Contract.SlashIndicatorContract(&_RoninValidatorSet.CallOpts)
}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetCaller) StakingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "stakingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetSession) StakingContract() (common.Address, error) {
	return _RoninValidatorSet.Contract.StakingContract(&_RoninValidatorSet.CallOpts)
}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) StakingContract() (common.Address, error) {
	return _RoninValidatorSet.Contract.StakingContract(&_RoninValidatorSet.CallOpts)
}

// StakingVestingContract is a free data retrieval call binding the contract method 0x3529214b.
//
// Solidity: function stakingVestingContract() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetCaller) StakingVestingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "stakingVestingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingVestingContract is a free data retrieval call binding the contract method 0x3529214b.
//
// Solidity: function stakingVestingContract() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetSession) StakingVestingContract() (common.Address, error) {
	return _RoninValidatorSet.Contract.StakingVestingContract(&_RoninValidatorSet.CallOpts)
}

// StakingVestingContract is a free data retrieval call binding the contract method 0x3529214b.
//
// Solidity: function stakingVestingContract() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) StakingVestingContract() (common.Address, error) {
	return _RoninValidatorSet.Contract.StakingVestingContract(&_RoninValidatorSet.CallOpts)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCaller) ValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "validatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetSession) ValidatorCount() (*big.Int, error) {
	return _RoninValidatorSet.Contract.ValidatorCount(&_RoninValidatorSet.CallOpts)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) ValidatorCount() (*big.Int, error) {
	return _RoninValidatorSet.Contract.ValidatorCount(&_RoninValidatorSet.CallOpts)
}

// AddValidatorCandidate is a paid mutator transaction binding the contract method 0xe18572bf.
//
// Solidity: function addValidatorCandidate(address _admin, address _consensusAddr, address _treasuryAddr, uint256 _commissionRate) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) AddValidatorCandidate(opts *bind.TransactOpts, _admin common.Address, _consensusAddr common.Address, _treasuryAddr common.Address, _commissionRate *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "addValidatorCandidate", _admin, _consensusAddr, _treasuryAddr, _commissionRate)
}

// AddValidatorCandidate is a paid mutator transaction binding the contract method 0xe18572bf.
//
// Solidity: function addValidatorCandidate(address _admin, address _consensusAddr, address _treasuryAddr, uint256 _commissionRate) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) AddValidatorCandidate(_admin common.Address, _consensusAddr common.Address, _treasuryAddr common.Address, _commissionRate *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.AddValidatorCandidate(&_RoninValidatorSet.TransactOpts, _admin, _consensusAddr, _treasuryAddr, _commissionRate)
}

// AddValidatorCandidate is a paid mutator transaction binding the contract method 0xe18572bf.
//
// Solidity: function addValidatorCandidate(address _admin, address _consensusAddr, address _treasuryAddr, uint256 _commissionRate) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) AddValidatorCandidate(_admin common.Address, _consensusAddr common.Address, _treasuryAddr common.Address, _commissionRate *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.AddValidatorCandidate(&_RoninValidatorSet.TransactOpts, _admin, _consensusAddr, _treasuryAddr, _commissionRate)
}

// Initialize is a paid mutator transaction binding the contract method 0xbeb3e382.
//
// Solidity: function initialize(address __slashIndicatorContract, address __stakingContract, address __stakingVestingContract, address __maintenanceContract, uint256 __maxValidatorNumber, uint256 __maxValidatorCandidate, uint256 __maxPrioritizedValidatorNumber, uint256 __numberOfBlocksInEpoch, uint256 __numberOfEpochsInPeriod) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) Initialize(opts *bind.TransactOpts, __slashIndicatorContract common.Address, __stakingContract common.Address, __stakingVestingContract common.Address, __maintenanceContract common.Address, __maxValidatorNumber *big.Int, __maxValidatorCandidate *big.Int, __maxPrioritizedValidatorNumber *big.Int, __numberOfBlocksInEpoch *big.Int, __numberOfEpochsInPeriod *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "initialize", __slashIndicatorContract, __stakingContract, __stakingVestingContract, __maintenanceContract, __maxValidatorNumber, __maxValidatorCandidate, __maxPrioritizedValidatorNumber, __numberOfBlocksInEpoch, __numberOfEpochsInPeriod)
}

// Initialize is a paid mutator transaction binding the contract method 0xbeb3e382.
//
// Solidity: function initialize(address __slashIndicatorContract, address __stakingContract, address __stakingVestingContract, address __maintenanceContract, uint256 __maxValidatorNumber, uint256 __maxValidatorCandidate, uint256 __maxPrioritizedValidatorNumber, uint256 __numberOfBlocksInEpoch, uint256 __numberOfEpochsInPeriod) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) Initialize(__slashIndicatorContract common.Address, __stakingContract common.Address, __stakingVestingContract common.Address, __maintenanceContract common.Address, __maxValidatorNumber *big.Int, __maxValidatorCandidate *big.Int, __maxPrioritizedValidatorNumber *big.Int, __numberOfBlocksInEpoch *big.Int, __numberOfEpochsInPeriod *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.Initialize(&_RoninValidatorSet.TransactOpts, __slashIndicatorContract, __stakingContract, __stakingVestingContract, __maintenanceContract, __maxValidatorNumber, __maxValidatorCandidate, __maxPrioritizedValidatorNumber, __numberOfBlocksInEpoch, __numberOfEpochsInPeriod)
}

// Initialize is a paid mutator transaction binding the contract method 0xbeb3e382.
//
// Solidity: function initialize(address __slashIndicatorContract, address __stakingContract, address __stakingVestingContract, address __maintenanceContract, uint256 __maxValidatorNumber, uint256 __maxValidatorCandidate, uint256 __maxPrioritizedValidatorNumber, uint256 __numberOfBlocksInEpoch, uint256 __numberOfEpochsInPeriod) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) Initialize(__slashIndicatorContract common.Address, __stakingContract common.Address, __stakingVestingContract common.Address, __maintenanceContract common.Address, __maxValidatorNumber *big.Int, __maxValidatorCandidate *big.Int, __maxPrioritizedValidatorNumber *big.Int, __numberOfBlocksInEpoch *big.Int, __numberOfEpochsInPeriod *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.Initialize(&_RoninValidatorSet.TransactOpts, __slashIndicatorContract, __stakingContract, __stakingVestingContract, __maintenanceContract, __maxValidatorNumber, __maxValidatorCandidate, __maxPrioritizedValidatorNumber, __numberOfBlocksInEpoch, __numberOfEpochsInPeriod)
}

// SetMaintenanceContract is a paid mutator transaction binding the contract method 0x46fe9311.
//
// Solidity: function setMaintenanceContract(address _addr) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) SetMaintenanceContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "setMaintenanceContract", _addr)
}

// SetMaintenanceContract is a paid mutator transaction binding the contract method 0x46fe9311.
//
// Solidity: function setMaintenanceContract(address _addr) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) SetMaintenanceContract(_addr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetMaintenanceContract(&_RoninValidatorSet.TransactOpts, _addr)
}

// SetMaintenanceContract is a paid mutator transaction binding the contract method 0x46fe9311.
//
// Solidity: function setMaintenanceContract(address _addr) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) SetMaintenanceContract(_addr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetMaintenanceContract(&_RoninValidatorSet.TransactOpts, _addr)
}

// SetMaxValidatorCandidate is a paid mutator transaction binding the contract method 0x4f2a693f.
//
// Solidity: function setMaxValidatorCandidate(uint256 _number) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) SetMaxValidatorCandidate(opts *bind.TransactOpts, _number *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "setMaxValidatorCandidate", _number)
}

// SetMaxValidatorCandidate is a paid mutator transaction binding the contract method 0x4f2a693f.
//
// Solidity: function setMaxValidatorCandidate(uint256 _number) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) SetMaxValidatorCandidate(_number *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetMaxValidatorCandidate(&_RoninValidatorSet.TransactOpts, _number)
}

// SetMaxValidatorCandidate is a paid mutator transaction binding the contract method 0x4f2a693f.
//
// Solidity: function setMaxValidatorCandidate(uint256 _number) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) SetMaxValidatorCandidate(_number *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetMaxValidatorCandidate(&_RoninValidatorSet.TransactOpts, _number)
}

// SetMaxValidatorNumber is a paid mutator transaction binding the contract method 0x823a7b9c.
//
// Solidity: function setMaxValidatorNumber(uint256 __maxValidatorNumber) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) SetMaxValidatorNumber(opts *bind.TransactOpts, __maxValidatorNumber *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "setMaxValidatorNumber", __maxValidatorNumber)
}

// SetMaxValidatorNumber is a paid mutator transaction binding the contract method 0x823a7b9c.
//
// Solidity: function setMaxValidatorNumber(uint256 __maxValidatorNumber) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) SetMaxValidatorNumber(__maxValidatorNumber *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetMaxValidatorNumber(&_RoninValidatorSet.TransactOpts, __maxValidatorNumber)
}

// SetMaxValidatorNumber is a paid mutator transaction binding the contract method 0x823a7b9c.
//
// Solidity: function setMaxValidatorNumber(uint256 __maxValidatorNumber) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) SetMaxValidatorNumber(__maxValidatorNumber *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetMaxValidatorNumber(&_RoninValidatorSet.TransactOpts, __maxValidatorNumber)
}

// SetNumberOfBlocksInEpoch is a paid mutator transaction binding the contract method 0xd72733fc.
//
// Solidity: function setNumberOfBlocksInEpoch(uint256 __numberOfBlocksInEpoch) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) SetNumberOfBlocksInEpoch(opts *bind.TransactOpts, __numberOfBlocksInEpoch *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "setNumberOfBlocksInEpoch", __numberOfBlocksInEpoch)
}

// SetNumberOfBlocksInEpoch is a paid mutator transaction binding the contract method 0xd72733fc.
//
// Solidity: function setNumberOfBlocksInEpoch(uint256 __numberOfBlocksInEpoch) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) SetNumberOfBlocksInEpoch(__numberOfBlocksInEpoch *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetNumberOfBlocksInEpoch(&_RoninValidatorSet.TransactOpts, __numberOfBlocksInEpoch)
}

// SetNumberOfBlocksInEpoch is a paid mutator transaction binding the contract method 0xd72733fc.
//
// Solidity: function setNumberOfBlocksInEpoch(uint256 __numberOfBlocksInEpoch) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) SetNumberOfBlocksInEpoch(__numberOfBlocksInEpoch *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetNumberOfBlocksInEpoch(&_RoninValidatorSet.TransactOpts, __numberOfBlocksInEpoch)
}

// SetNumberOfEpochsInPeriod is a paid mutator transaction binding the contract method 0xd6fa322c.
//
// Solidity: function setNumberOfEpochsInPeriod(uint256 __numberOfEpochsInPeriod) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) SetNumberOfEpochsInPeriod(opts *bind.TransactOpts, __numberOfEpochsInPeriod *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "setNumberOfEpochsInPeriod", __numberOfEpochsInPeriod)
}

// SetNumberOfEpochsInPeriod is a paid mutator transaction binding the contract method 0xd6fa322c.
//
// Solidity: function setNumberOfEpochsInPeriod(uint256 __numberOfEpochsInPeriod) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) SetNumberOfEpochsInPeriod(__numberOfEpochsInPeriod *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetNumberOfEpochsInPeriod(&_RoninValidatorSet.TransactOpts, __numberOfEpochsInPeriod)
}

// SetNumberOfEpochsInPeriod is a paid mutator transaction binding the contract method 0xd6fa322c.
//
// Solidity: function setNumberOfEpochsInPeriod(uint256 __numberOfEpochsInPeriod) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) SetNumberOfEpochsInPeriod(__numberOfEpochsInPeriod *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetNumberOfEpochsInPeriod(&_RoninValidatorSet.TransactOpts, __numberOfEpochsInPeriod)
}

// SetPrioritizedAddresses is a paid mutator transaction binding the contract method 0xd33a5ca2.
//
// Solidity: function setPrioritizedAddresses(address[] _addrs, bool[] _statuses) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) SetPrioritizedAddresses(opts *bind.TransactOpts, _addrs []common.Address, _statuses []bool) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "setPrioritizedAddresses", _addrs, _statuses)
}

// SetPrioritizedAddresses is a paid mutator transaction binding the contract method 0xd33a5ca2.
//
// Solidity: function setPrioritizedAddresses(address[] _addrs, bool[] _statuses) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) SetPrioritizedAddresses(_addrs []common.Address, _statuses []bool) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetPrioritizedAddresses(&_RoninValidatorSet.TransactOpts, _addrs, _statuses)
}

// SetPrioritizedAddresses is a paid mutator transaction binding the contract method 0xd33a5ca2.
//
// Solidity: function setPrioritizedAddresses(address[] _addrs, bool[] _statuses) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) SetPrioritizedAddresses(_addrs []common.Address, _statuses []bool) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetPrioritizedAddresses(&_RoninValidatorSet.TransactOpts, _addrs, _statuses)
}

// SetSlashIndicatorContract is a paid mutator transaction binding the contract method 0x2bcf3d15.
//
// Solidity: function setSlashIndicatorContract(address _addr) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) SetSlashIndicatorContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "setSlashIndicatorContract", _addr)
}

// SetSlashIndicatorContract is a paid mutator transaction binding the contract method 0x2bcf3d15.
//
// Solidity: function setSlashIndicatorContract(address _addr) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) SetSlashIndicatorContract(_addr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetSlashIndicatorContract(&_RoninValidatorSet.TransactOpts, _addr)
}

// SetSlashIndicatorContract is a paid mutator transaction binding the contract method 0x2bcf3d15.
//
// Solidity: function setSlashIndicatorContract(address _addr) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) SetSlashIndicatorContract(_addr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetSlashIndicatorContract(&_RoninValidatorSet.TransactOpts, _addr)
}

// SetStakingContract is a paid mutator transaction binding the contract method 0x9dd373b9.
//
// Solidity: function setStakingContract(address _addr) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) SetStakingContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "setStakingContract", _addr)
}

// SetStakingContract is a paid mutator transaction binding the contract method 0x9dd373b9.
//
// Solidity: function setStakingContract(address _addr) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) SetStakingContract(_addr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetStakingContract(&_RoninValidatorSet.TransactOpts, _addr)
}

// SetStakingContract is a paid mutator transaction binding the contract method 0x9dd373b9.
//
// Solidity: function setStakingContract(address _addr) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) SetStakingContract(_addr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetStakingContract(&_RoninValidatorSet.TransactOpts, _addr)
}

// SetStakingVestingContract is a paid mutator transaction binding the contract method 0xad295783.
//
// Solidity: function setStakingVestingContract(address _addr) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) SetStakingVestingContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "setStakingVestingContract", _addr)
}

// SetStakingVestingContract is a paid mutator transaction binding the contract method 0xad295783.
//
// Solidity: function setStakingVestingContract(address _addr) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) SetStakingVestingContract(_addr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetStakingVestingContract(&_RoninValidatorSet.TransactOpts, _addr)
}

// SetStakingVestingContract is a paid mutator transaction binding the contract method 0xad295783.
//
// Solidity: function setStakingVestingContract(address _addr) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) SetStakingVestingContract(_addr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetStakingVestingContract(&_RoninValidatorSet.TransactOpts, _addr)
}

// Slash is a paid mutator transaction binding the contract method 0x70f81f6c.
//
// Solidity: function slash(address _validatorAddr, uint256 _newJailedUntil, uint256 _slashAmount) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) Slash(opts *bind.TransactOpts, _validatorAddr common.Address, _newJailedUntil *big.Int, _slashAmount *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "slash", _validatorAddr, _newJailedUntil, _slashAmount)
}

// Slash is a paid mutator transaction binding the contract method 0x70f81f6c.
//
// Solidity: function slash(address _validatorAddr, uint256 _newJailedUntil, uint256 _slashAmount) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) Slash(_validatorAddr common.Address, _newJailedUntil *big.Int, _slashAmount *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.Slash(&_RoninValidatorSet.TransactOpts, _validatorAddr, _newJailedUntil, _slashAmount)
}

// Slash is a paid mutator transaction binding the contract method 0x70f81f6c.
//
// Solidity: function slash(address _validatorAddr, uint256 _newJailedUntil, uint256 _slashAmount) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) Slash(_validatorAddr common.Address, _newJailedUntil *big.Int, _slashAmount *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.Slash(&_RoninValidatorSet.TransactOpts, _validatorAddr, _newJailedUntil, _slashAmount)
}

// SubmitBlockReward is a paid mutator transaction binding the contract method 0x52091f17.
//
// Solidity: function submitBlockReward() payable returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) SubmitBlockReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "submitBlockReward")
}

// SubmitBlockReward is a paid mutator transaction binding the contract method 0x52091f17.
//
// Solidity: function submitBlockReward() payable returns()
func (_RoninValidatorSet *RoninValidatorSetSession) SubmitBlockReward() (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SubmitBlockReward(&_RoninValidatorSet.TransactOpts)
}

// SubmitBlockReward is a paid mutator transaction binding the contract method 0x52091f17.
//
// Solidity: function submitBlockReward() payable returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) SubmitBlockReward() (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SubmitBlockReward(&_RoninValidatorSet.TransactOpts)
}

// SyncCandidates is a paid mutator transaction binding the contract method 0xde7702fb.
//
// Solidity: function syncCandidates() returns(uint256[] _balances)
func (_RoninValidatorSet *RoninValidatorSetTransactor) SyncCandidates(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "syncCandidates")
}

// SyncCandidates is a paid mutator transaction binding the contract method 0xde7702fb.
//
// Solidity: function syncCandidates() returns(uint256[] _balances)
func (_RoninValidatorSet *RoninValidatorSetSession) SyncCandidates() (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SyncCandidates(&_RoninValidatorSet.TransactOpts)
}

// SyncCandidates is a paid mutator transaction binding the contract method 0xde7702fb.
//
// Solidity: function syncCandidates() returns(uint256[] _balances)
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) SyncCandidates() (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SyncCandidates(&_RoninValidatorSet.TransactOpts)
}

// WrapUpEpoch is a paid mutator transaction binding the contract method 0x72e46810.
//
// Solidity: function wrapUpEpoch() payable returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) WrapUpEpoch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "wrapUpEpoch")
}

// WrapUpEpoch is a paid mutator transaction binding the contract method 0x72e46810.
//
// Solidity: function wrapUpEpoch() payable returns()
func (_RoninValidatorSet *RoninValidatorSetSession) WrapUpEpoch() (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.WrapUpEpoch(&_RoninValidatorSet.TransactOpts)
}

// WrapUpEpoch is a paid mutator transaction binding the contract method 0x72e46810.
//
// Solidity: function wrapUpEpoch() payable returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) WrapUpEpoch() (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.WrapUpEpoch(&_RoninValidatorSet.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RoninValidatorSet *RoninValidatorSetSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.Fallback(&_RoninValidatorSet.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.Fallback(&_RoninValidatorSet.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RoninValidatorSet *RoninValidatorSetSession) Receive() (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.Receive(&_RoninValidatorSet.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) Receive() (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.Receive(&_RoninValidatorSet.TransactOpts)
}

// RoninValidatorSetAddressesPriorityStatusUpdatedIterator is returned from FilterAddressesPriorityStatusUpdated and is used to iterate over the raw logs and unpacked data for AddressesPriorityStatusUpdated events raised by the RoninValidatorSet contract.
type RoninValidatorSetAddressesPriorityStatusUpdatedIterator struct {
	Event *RoninValidatorSetAddressesPriorityStatusUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetAddressesPriorityStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetAddressesPriorityStatusUpdated)
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
		it.Event = new(RoninValidatorSetAddressesPriorityStatusUpdated)
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
func (it *RoninValidatorSetAddressesPriorityStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetAddressesPriorityStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetAddressesPriorityStatusUpdated represents a AddressesPriorityStatusUpdated event raised by the RoninValidatorSet contract.
type RoninValidatorSetAddressesPriorityStatusUpdated struct {
	Arg0 []common.Address
	Arg1 []bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddressesPriorityStatusUpdated is a free log retrieval operation binding the contract event 0xa52c766fffd3af2ed65a8599973f96a0713c68566068cc057ad25cabd88ed466.
//
// Solidity: event AddressesPriorityStatusUpdated(address[] arg0, bool[] arg1)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterAddressesPriorityStatusUpdated(opts *bind.FilterOpts) (*RoninValidatorSetAddressesPriorityStatusUpdatedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "AddressesPriorityStatusUpdated")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetAddressesPriorityStatusUpdatedIterator{contract: _RoninValidatorSet.contract, event: "AddressesPriorityStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchAddressesPriorityStatusUpdated is a free log subscription operation binding the contract event 0xa52c766fffd3af2ed65a8599973f96a0713c68566068cc057ad25cabd88ed466.
//
// Solidity: event AddressesPriorityStatusUpdated(address[] arg0, bool[] arg1)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchAddressesPriorityStatusUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetAddressesPriorityStatusUpdated) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "AddressesPriorityStatusUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetAddressesPriorityStatusUpdated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "AddressesPriorityStatusUpdated", log); err != nil {
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

// ParseAddressesPriorityStatusUpdated is a log parse operation binding the contract event 0xa52c766fffd3af2ed65a8599973f96a0713c68566068cc057ad25cabd88ed466.
//
// Solidity: event AddressesPriorityStatusUpdated(address[] arg0, bool[] arg1)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseAddressesPriorityStatusUpdated(log types.Log) (*RoninValidatorSetAddressesPriorityStatusUpdated, error) {
	event := new(RoninValidatorSetAddressesPriorityStatusUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "AddressesPriorityStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetBlockRewardSubmittedIterator is returned from FilterBlockRewardSubmitted and is used to iterate over the raw logs and unpacked data for BlockRewardSubmitted events raised by the RoninValidatorSet contract.
type RoninValidatorSetBlockRewardSubmittedIterator struct {
	Event *RoninValidatorSetBlockRewardSubmitted // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetBlockRewardSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetBlockRewardSubmitted)
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
		it.Event = new(RoninValidatorSetBlockRewardSubmitted)
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
func (it *RoninValidatorSetBlockRewardSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetBlockRewardSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetBlockRewardSubmitted represents a BlockRewardSubmitted event raised by the RoninValidatorSet contract.
type RoninValidatorSetBlockRewardSubmitted struct {
	CoinbaseAddr    common.Address
	SubmittedAmount *big.Int
	BonusAmount     *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBlockRewardSubmitted is a free log retrieval operation binding the contract event 0x0ede5c3be8625943fa64003cd4b91230089411249f3059bac6500873543ca9b1.
//
// Solidity: event BlockRewardSubmitted(address coinbaseAddr, uint256 submittedAmount, uint256 bonusAmount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterBlockRewardSubmitted(opts *bind.FilterOpts) (*RoninValidatorSetBlockRewardSubmittedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "BlockRewardSubmitted")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetBlockRewardSubmittedIterator{contract: _RoninValidatorSet.contract, event: "BlockRewardSubmitted", logs: logs, sub: sub}, nil
}

// WatchBlockRewardSubmitted is a free log subscription operation binding the contract event 0x0ede5c3be8625943fa64003cd4b91230089411249f3059bac6500873543ca9b1.
//
// Solidity: event BlockRewardSubmitted(address coinbaseAddr, uint256 submittedAmount, uint256 bonusAmount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchBlockRewardSubmitted(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetBlockRewardSubmitted) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "BlockRewardSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetBlockRewardSubmitted)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "BlockRewardSubmitted", log); err != nil {
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

// ParseBlockRewardSubmitted is a log parse operation binding the contract event 0x0ede5c3be8625943fa64003cd4b91230089411249f3059bac6500873543ca9b1.
//
// Solidity: event BlockRewardSubmitted(address coinbaseAddr, uint256 submittedAmount, uint256 bonusAmount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseBlockRewardSubmitted(log types.Log) (*RoninValidatorSetBlockRewardSubmitted, error) {
	event := new(RoninValidatorSetBlockRewardSubmitted)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "BlockRewardSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RoninValidatorSet contract.
type RoninValidatorSetInitializedIterator struct {
	Event *RoninValidatorSetInitialized // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetInitialized)
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
		it.Event = new(RoninValidatorSetInitialized)
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
func (it *RoninValidatorSetInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetInitialized represents a Initialized event raised by the RoninValidatorSet contract.
type RoninValidatorSetInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterInitialized(opts *bind.FilterOpts) (*RoninValidatorSetInitializedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetInitializedIterator{contract: _RoninValidatorSet.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetInitialized) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetInitialized)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseInitialized(log types.Log) (*RoninValidatorSetInitialized, error) {
	event := new(RoninValidatorSetInitialized)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetMaintenanceContractUpdatedIterator is returned from FilterMaintenanceContractUpdated and is used to iterate over the raw logs and unpacked data for MaintenanceContractUpdated events raised by the RoninValidatorSet contract.
type RoninValidatorSetMaintenanceContractUpdatedIterator struct {
	Event *RoninValidatorSetMaintenanceContractUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetMaintenanceContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetMaintenanceContractUpdated)
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
		it.Event = new(RoninValidatorSetMaintenanceContractUpdated)
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
func (it *RoninValidatorSetMaintenanceContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetMaintenanceContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetMaintenanceContractUpdated represents a MaintenanceContractUpdated event raised by the RoninValidatorSet contract.
type RoninValidatorSetMaintenanceContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMaintenanceContractUpdated is a free log retrieval operation binding the contract event 0x31a33f126a5bae3c5bdf6cfc2cd6dcfffe2fe9634bdb09e21c44762993889e3b.
//
// Solidity: event MaintenanceContractUpdated(address arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterMaintenanceContractUpdated(opts *bind.FilterOpts) (*RoninValidatorSetMaintenanceContractUpdatedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "MaintenanceContractUpdated")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetMaintenanceContractUpdatedIterator{contract: _RoninValidatorSet.contract, event: "MaintenanceContractUpdated", logs: logs, sub: sub}, nil
}

// WatchMaintenanceContractUpdated is a free log subscription operation binding the contract event 0x31a33f126a5bae3c5bdf6cfc2cd6dcfffe2fe9634bdb09e21c44762993889e3b.
//
// Solidity: event MaintenanceContractUpdated(address arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchMaintenanceContractUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetMaintenanceContractUpdated) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "MaintenanceContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetMaintenanceContractUpdated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "MaintenanceContractUpdated", log); err != nil {
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
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseMaintenanceContractUpdated(log types.Log) (*RoninValidatorSetMaintenanceContractUpdated, error) {
	event := new(RoninValidatorSetMaintenanceContractUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "MaintenanceContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetMaxPrioritizedValidatorNumberUpdatedIterator is returned from FilterMaxPrioritizedValidatorNumberUpdated and is used to iterate over the raw logs and unpacked data for MaxPrioritizedValidatorNumberUpdated events raised by the RoninValidatorSet contract.
type RoninValidatorSetMaxPrioritizedValidatorNumberUpdatedIterator struct {
	Event *RoninValidatorSetMaxPrioritizedValidatorNumberUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetMaxPrioritizedValidatorNumberUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetMaxPrioritizedValidatorNumberUpdated)
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
		it.Event = new(RoninValidatorSetMaxPrioritizedValidatorNumberUpdated)
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
func (it *RoninValidatorSetMaxPrioritizedValidatorNumberUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetMaxPrioritizedValidatorNumberUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetMaxPrioritizedValidatorNumberUpdated represents a MaxPrioritizedValidatorNumberUpdated event raised by the RoninValidatorSet contract.
type RoninValidatorSetMaxPrioritizedValidatorNumberUpdated struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMaxPrioritizedValidatorNumberUpdated is a free log retrieval operation binding the contract event 0xa9588dc77416849bd922605ce4fc806712281ad8a8f32d4238d6c8cca548e15e.
//
// Solidity: event MaxPrioritizedValidatorNumberUpdated(uint256 arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterMaxPrioritizedValidatorNumberUpdated(opts *bind.FilterOpts) (*RoninValidatorSetMaxPrioritizedValidatorNumberUpdatedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "MaxPrioritizedValidatorNumberUpdated")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetMaxPrioritizedValidatorNumberUpdatedIterator{contract: _RoninValidatorSet.contract, event: "MaxPrioritizedValidatorNumberUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxPrioritizedValidatorNumberUpdated is a free log subscription operation binding the contract event 0xa9588dc77416849bd922605ce4fc806712281ad8a8f32d4238d6c8cca548e15e.
//
// Solidity: event MaxPrioritizedValidatorNumberUpdated(uint256 arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchMaxPrioritizedValidatorNumberUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetMaxPrioritizedValidatorNumberUpdated) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "MaxPrioritizedValidatorNumberUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetMaxPrioritizedValidatorNumberUpdated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "MaxPrioritizedValidatorNumberUpdated", log); err != nil {
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

// ParseMaxPrioritizedValidatorNumberUpdated is a log parse operation binding the contract event 0xa9588dc77416849bd922605ce4fc806712281ad8a8f32d4238d6c8cca548e15e.
//
// Solidity: event MaxPrioritizedValidatorNumberUpdated(uint256 arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseMaxPrioritizedValidatorNumberUpdated(log types.Log) (*RoninValidatorSetMaxPrioritizedValidatorNumberUpdated, error) {
	event := new(RoninValidatorSetMaxPrioritizedValidatorNumberUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "MaxPrioritizedValidatorNumberUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetMaxValidatorCandidateUpdatedIterator is returned from FilterMaxValidatorCandidateUpdated and is used to iterate over the raw logs and unpacked data for MaxValidatorCandidateUpdated events raised by the RoninValidatorSet contract.
type RoninValidatorSetMaxValidatorCandidateUpdatedIterator struct {
	Event *RoninValidatorSetMaxValidatorCandidateUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetMaxValidatorCandidateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetMaxValidatorCandidateUpdated)
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
		it.Event = new(RoninValidatorSetMaxValidatorCandidateUpdated)
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
func (it *RoninValidatorSetMaxValidatorCandidateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetMaxValidatorCandidateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetMaxValidatorCandidateUpdated represents a MaxValidatorCandidateUpdated event raised by the RoninValidatorSet contract.
type RoninValidatorSetMaxValidatorCandidateUpdated struct {
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMaxValidatorCandidateUpdated is a free log retrieval operation binding the contract event 0x82d5dc32d1b741512ad09c32404d7e7921e8934c6222343d95f55f7a2b9b2ab4.
//
// Solidity: event MaxValidatorCandidateUpdated(uint256 threshold)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterMaxValidatorCandidateUpdated(opts *bind.FilterOpts) (*RoninValidatorSetMaxValidatorCandidateUpdatedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "MaxValidatorCandidateUpdated")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetMaxValidatorCandidateUpdatedIterator{contract: _RoninValidatorSet.contract, event: "MaxValidatorCandidateUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxValidatorCandidateUpdated is a free log subscription operation binding the contract event 0x82d5dc32d1b741512ad09c32404d7e7921e8934c6222343d95f55f7a2b9b2ab4.
//
// Solidity: event MaxValidatorCandidateUpdated(uint256 threshold)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchMaxValidatorCandidateUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetMaxValidatorCandidateUpdated) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "MaxValidatorCandidateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetMaxValidatorCandidateUpdated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "MaxValidatorCandidateUpdated", log); err != nil {
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

// ParseMaxValidatorCandidateUpdated is a log parse operation binding the contract event 0x82d5dc32d1b741512ad09c32404d7e7921e8934c6222343d95f55f7a2b9b2ab4.
//
// Solidity: event MaxValidatorCandidateUpdated(uint256 threshold)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseMaxValidatorCandidateUpdated(log types.Log) (*RoninValidatorSetMaxValidatorCandidateUpdated, error) {
	event := new(RoninValidatorSetMaxValidatorCandidateUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "MaxValidatorCandidateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetMaxValidatorNumberUpdatedIterator is returned from FilterMaxValidatorNumberUpdated and is used to iterate over the raw logs and unpacked data for MaxValidatorNumberUpdated events raised by the RoninValidatorSet contract.
type RoninValidatorSetMaxValidatorNumberUpdatedIterator struct {
	Event *RoninValidatorSetMaxValidatorNumberUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetMaxValidatorNumberUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetMaxValidatorNumberUpdated)
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
		it.Event = new(RoninValidatorSetMaxValidatorNumberUpdated)
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
func (it *RoninValidatorSetMaxValidatorNumberUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetMaxValidatorNumberUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetMaxValidatorNumberUpdated represents a MaxValidatorNumberUpdated event raised by the RoninValidatorSet contract.
type RoninValidatorSetMaxValidatorNumberUpdated struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMaxValidatorNumberUpdated is a free log retrieval operation binding the contract event 0xb5464c05fd0e0f000c535850116cda2742ee1f7b34384cb920ad7b8e802138b5.
//
// Solidity: event MaxValidatorNumberUpdated(uint256 arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterMaxValidatorNumberUpdated(opts *bind.FilterOpts) (*RoninValidatorSetMaxValidatorNumberUpdatedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "MaxValidatorNumberUpdated")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetMaxValidatorNumberUpdatedIterator{contract: _RoninValidatorSet.contract, event: "MaxValidatorNumberUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxValidatorNumberUpdated is a free log subscription operation binding the contract event 0xb5464c05fd0e0f000c535850116cda2742ee1f7b34384cb920ad7b8e802138b5.
//
// Solidity: event MaxValidatorNumberUpdated(uint256 arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchMaxValidatorNumberUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetMaxValidatorNumberUpdated) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "MaxValidatorNumberUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetMaxValidatorNumberUpdated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "MaxValidatorNumberUpdated", log); err != nil {
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

// ParseMaxValidatorNumberUpdated is a log parse operation binding the contract event 0xb5464c05fd0e0f000c535850116cda2742ee1f7b34384cb920ad7b8e802138b5.
//
// Solidity: event MaxValidatorNumberUpdated(uint256 arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseMaxValidatorNumberUpdated(log types.Log) (*RoninValidatorSetMaxValidatorNumberUpdated, error) {
	event := new(RoninValidatorSetMaxValidatorNumberUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "MaxValidatorNumberUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetMiningRewardDistributedIterator is returned from FilterMiningRewardDistributed and is used to iterate over the raw logs and unpacked data for MiningRewardDistributed events raised by the RoninValidatorSet contract.
type RoninValidatorSetMiningRewardDistributedIterator struct {
	Event *RoninValidatorSetMiningRewardDistributed // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetMiningRewardDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetMiningRewardDistributed)
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
		it.Event = new(RoninValidatorSetMiningRewardDistributed)
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
func (it *RoninValidatorSetMiningRewardDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetMiningRewardDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetMiningRewardDistributed represents a MiningRewardDistributed event raised by the RoninValidatorSet contract.
type RoninValidatorSetMiningRewardDistributed struct {
	ValidatorAddr common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMiningRewardDistributed is a free log retrieval operation binding the contract event 0x9bae506d1374d366cdfa60105473d52dfdbaaae60e55de77bf6ee07f2add0cfb.
//
// Solidity: event MiningRewardDistributed(address validatorAddr, uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterMiningRewardDistributed(opts *bind.FilterOpts) (*RoninValidatorSetMiningRewardDistributedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "MiningRewardDistributed")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetMiningRewardDistributedIterator{contract: _RoninValidatorSet.contract, event: "MiningRewardDistributed", logs: logs, sub: sub}, nil
}

// WatchMiningRewardDistributed is a free log subscription operation binding the contract event 0x9bae506d1374d366cdfa60105473d52dfdbaaae60e55de77bf6ee07f2add0cfb.
//
// Solidity: event MiningRewardDistributed(address validatorAddr, uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchMiningRewardDistributed(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetMiningRewardDistributed) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "MiningRewardDistributed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetMiningRewardDistributed)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "MiningRewardDistributed", log); err != nil {
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

// ParseMiningRewardDistributed is a log parse operation binding the contract event 0x9bae506d1374d366cdfa60105473d52dfdbaaae60e55de77bf6ee07f2add0cfb.
//
// Solidity: event MiningRewardDistributed(address validatorAddr, uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseMiningRewardDistributed(log types.Log) (*RoninValidatorSetMiningRewardDistributed, error) {
	event := new(RoninValidatorSetMiningRewardDistributed)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "MiningRewardDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetNumberOfBlocksInEpochUpdatedIterator is returned from FilterNumberOfBlocksInEpochUpdated and is used to iterate over the raw logs and unpacked data for NumberOfBlocksInEpochUpdated events raised by the RoninValidatorSet contract.
type RoninValidatorSetNumberOfBlocksInEpochUpdatedIterator struct {
	Event *RoninValidatorSetNumberOfBlocksInEpochUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetNumberOfBlocksInEpochUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetNumberOfBlocksInEpochUpdated)
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
		it.Event = new(RoninValidatorSetNumberOfBlocksInEpochUpdated)
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
func (it *RoninValidatorSetNumberOfBlocksInEpochUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetNumberOfBlocksInEpochUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetNumberOfBlocksInEpochUpdated represents a NumberOfBlocksInEpochUpdated event raised by the RoninValidatorSet contract.
type RoninValidatorSetNumberOfBlocksInEpochUpdated struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNumberOfBlocksInEpochUpdated is a free log retrieval operation binding the contract event 0xbfd285a38b782d8a00e424fb824320ff3d1a698534358d02da611468d59b7808.
//
// Solidity: event NumberOfBlocksInEpochUpdated(uint256 arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterNumberOfBlocksInEpochUpdated(opts *bind.FilterOpts) (*RoninValidatorSetNumberOfBlocksInEpochUpdatedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "NumberOfBlocksInEpochUpdated")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetNumberOfBlocksInEpochUpdatedIterator{contract: _RoninValidatorSet.contract, event: "NumberOfBlocksInEpochUpdated", logs: logs, sub: sub}, nil
}

// WatchNumberOfBlocksInEpochUpdated is a free log subscription operation binding the contract event 0xbfd285a38b782d8a00e424fb824320ff3d1a698534358d02da611468d59b7808.
//
// Solidity: event NumberOfBlocksInEpochUpdated(uint256 arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchNumberOfBlocksInEpochUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetNumberOfBlocksInEpochUpdated) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "NumberOfBlocksInEpochUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetNumberOfBlocksInEpochUpdated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "NumberOfBlocksInEpochUpdated", log); err != nil {
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

// ParseNumberOfBlocksInEpochUpdated is a log parse operation binding the contract event 0xbfd285a38b782d8a00e424fb824320ff3d1a698534358d02da611468d59b7808.
//
// Solidity: event NumberOfBlocksInEpochUpdated(uint256 arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseNumberOfBlocksInEpochUpdated(log types.Log) (*RoninValidatorSetNumberOfBlocksInEpochUpdated, error) {
	event := new(RoninValidatorSetNumberOfBlocksInEpochUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "NumberOfBlocksInEpochUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetNumberOfEpochsInPeriodUpdatedIterator is returned from FilterNumberOfEpochsInPeriodUpdated and is used to iterate over the raw logs and unpacked data for NumberOfEpochsInPeriodUpdated events raised by the RoninValidatorSet contract.
type RoninValidatorSetNumberOfEpochsInPeriodUpdatedIterator struct {
	Event *RoninValidatorSetNumberOfEpochsInPeriodUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetNumberOfEpochsInPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetNumberOfEpochsInPeriodUpdated)
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
		it.Event = new(RoninValidatorSetNumberOfEpochsInPeriodUpdated)
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
func (it *RoninValidatorSetNumberOfEpochsInPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetNumberOfEpochsInPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetNumberOfEpochsInPeriodUpdated represents a NumberOfEpochsInPeriodUpdated event raised by the RoninValidatorSet contract.
type RoninValidatorSetNumberOfEpochsInPeriodUpdated struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNumberOfEpochsInPeriodUpdated is a free log retrieval operation binding the contract event 0x1d01baa2db15fced4f4e5fcfd4245e65ad9b083c110d26542f4a5f78d5425e77.
//
// Solidity: event NumberOfEpochsInPeriodUpdated(uint256 arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterNumberOfEpochsInPeriodUpdated(opts *bind.FilterOpts) (*RoninValidatorSetNumberOfEpochsInPeriodUpdatedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "NumberOfEpochsInPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetNumberOfEpochsInPeriodUpdatedIterator{contract: _RoninValidatorSet.contract, event: "NumberOfEpochsInPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchNumberOfEpochsInPeriodUpdated is a free log subscription operation binding the contract event 0x1d01baa2db15fced4f4e5fcfd4245e65ad9b083c110d26542f4a5f78d5425e77.
//
// Solidity: event NumberOfEpochsInPeriodUpdated(uint256 arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchNumberOfEpochsInPeriodUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetNumberOfEpochsInPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "NumberOfEpochsInPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetNumberOfEpochsInPeriodUpdated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "NumberOfEpochsInPeriodUpdated", log); err != nil {
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

// ParseNumberOfEpochsInPeriodUpdated is a log parse operation binding the contract event 0x1d01baa2db15fced4f4e5fcfd4245e65ad9b083c110d26542f4a5f78d5425e77.
//
// Solidity: event NumberOfEpochsInPeriodUpdated(uint256 arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseNumberOfEpochsInPeriodUpdated(log types.Log) (*RoninValidatorSetNumberOfEpochsInPeriodUpdated, error) {
	event := new(RoninValidatorSetNumberOfEpochsInPeriodUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "NumberOfEpochsInPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetRewardDeprecatedIterator is returned from FilterRewardDeprecated and is used to iterate over the raw logs and unpacked data for RewardDeprecated events raised by the RoninValidatorSet contract.
type RoninValidatorSetRewardDeprecatedIterator struct {
	Event *RoninValidatorSetRewardDeprecated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetRewardDeprecatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetRewardDeprecated)
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
		it.Event = new(RoninValidatorSetRewardDeprecated)
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
func (it *RoninValidatorSetRewardDeprecatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetRewardDeprecatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetRewardDeprecated represents a RewardDeprecated event raised by the RoninValidatorSet contract.
type RoninValidatorSetRewardDeprecated struct {
	CoinbaseAddr common.Address
	RewardAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRewardDeprecated is a free log retrieval operation binding the contract event 0x2439a6ac441f1d6b3dbb7827ef6e056822e2261f900cad468012eee4f1f7f31c.
//
// Solidity: event RewardDeprecated(address coinbaseAddr, uint256 rewardAmount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterRewardDeprecated(opts *bind.FilterOpts) (*RoninValidatorSetRewardDeprecatedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "RewardDeprecated")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetRewardDeprecatedIterator{contract: _RoninValidatorSet.contract, event: "RewardDeprecated", logs: logs, sub: sub}, nil
}

// WatchRewardDeprecated is a free log subscription operation binding the contract event 0x2439a6ac441f1d6b3dbb7827ef6e056822e2261f900cad468012eee4f1f7f31c.
//
// Solidity: event RewardDeprecated(address coinbaseAddr, uint256 rewardAmount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchRewardDeprecated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetRewardDeprecated) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "RewardDeprecated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetRewardDeprecated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "RewardDeprecated", log); err != nil {
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

// ParseRewardDeprecated is a log parse operation binding the contract event 0x2439a6ac441f1d6b3dbb7827ef6e056822e2261f900cad468012eee4f1f7f31c.
//
// Solidity: event RewardDeprecated(address coinbaseAddr, uint256 rewardAmount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseRewardDeprecated(log types.Log) (*RoninValidatorSetRewardDeprecated, error) {
	event := new(RoninValidatorSetRewardDeprecated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "RewardDeprecated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetSlashIndicatorContractUpdatedIterator is returned from FilterSlashIndicatorContractUpdated and is used to iterate over the raw logs and unpacked data for SlashIndicatorContractUpdated events raised by the RoninValidatorSet contract.
type RoninValidatorSetSlashIndicatorContractUpdatedIterator struct {
	Event *RoninValidatorSetSlashIndicatorContractUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetSlashIndicatorContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetSlashIndicatorContractUpdated)
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
		it.Event = new(RoninValidatorSetSlashIndicatorContractUpdated)
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
func (it *RoninValidatorSetSlashIndicatorContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetSlashIndicatorContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetSlashIndicatorContractUpdated represents a SlashIndicatorContractUpdated event raised by the RoninValidatorSet contract.
type RoninValidatorSetSlashIndicatorContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSlashIndicatorContractUpdated is a free log retrieval operation binding the contract event 0xaa5b07dd43aa44c69b70a6a2b9c3fcfed12b6e5f6323596ba7ac91035ab80a4f.
//
// Solidity: event SlashIndicatorContractUpdated(address arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterSlashIndicatorContractUpdated(opts *bind.FilterOpts) (*RoninValidatorSetSlashIndicatorContractUpdatedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "SlashIndicatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetSlashIndicatorContractUpdatedIterator{contract: _RoninValidatorSet.contract, event: "SlashIndicatorContractUpdated", logs: logs, sub: sub}, nil
}

// WatchSlashIndicatorContractUpdated is a free log subscription operation binding the contract event 0xaa5b07dd43aa44c69b70a6a2b9c3fcfed12b6e5f6323596ba7ac91035ab80a4f.
//
// Solidity: event SlashIndicatorContractUpdated(address arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchSlashIndicatorContractUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetSlashIndicatorContractUpdated) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "SlashIndicatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetSlashIndicatorContractUpdated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "SlashIndicatorContractUpdated", log); err != nil {
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

// ParseSlashIndicatorContractUpdated is a log parse operation binding the contract event 0xaa5b07dd43aa44c69b70a6a2b9c3fcfed12b6e5f6323596ba7ac91035ab80a4f.
//
// Solidity: event SlashIndicatorContractUpdated(address arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseSlashIndicatorContractUpdated(log types.Log) (*RoninValidatorSetSlashIndicatorContractUpdated, error) {
	event := new(RoninValidatorSetSlashIndicatorContractUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "SlashIndicatorContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetStakingContractUpdatedIterator is returned from FilterStakingContractUpdated and is used to iterate over the raw logs and unpacked data for StakingContractUpdated events raised by the RoninValidatorSet contract.
type RoninValidatorSetStakingContractUpdatedIterator struct {
	Event *RoninValidatorSetStakingContractUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetStakingContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetStakingContractUpdated)
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
		it.Event = new(RoninValidatorSetStakingContractUpdated)
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
func (it *RoninValidatorSetStakingContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetStakingContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetStakingContractUpdated represents a StakingContractUpdated event raised by the RoninValidatorSet contract.
type RoninValidatorSetStakingContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterStakingContractUpdated is a free log retrieval operation binding the contract event 0x6397f5b135542bb3f477cb346cfab5abdec1251d08dc8f8d4efb4ffe122ea0bf.
//
// Solidity: event StakingContractUpdated(address arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterStakingContractUpdated(opts *bind.FilterOpts) (*RoninValidatorSetStakingContractUpdatedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "StakingContractUpdated")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetStakingContractUpdatedIterator{contract: _RoninValidatorSet.contract, event: "StakingContractUpdated", logs: logs, sub: sub}, nil
}

// WatchStakingContractUpdated is a free log subscription operation binding the contract event 0x6397f5b135542bb3f477cb346cfab5abdec1251d08dc8f8d4efb4ffe122ea0bf.
//
// Solidity: event StakingContractUpdated(address arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchStakingContractUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetStakingContractUpdated) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "StakingContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetStakingContractUpdated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "StakingContractUpdated", log); err != nil {
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

// ParseStakingContractUpdated is a log parse operation binding the contract event 0x6397f5b135542bb3f477cb346cfab5abdec1251d08dc8f8d4efb4ffe122ea0bf.
//
// Solidity: event StakingContractUpdated(address arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseStakingContractUpdated(log types.Log) (*RoninValidatorSetStakingContractUpdated, error) {
	event := new(RoninValidatorSetStakingContractUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "StakingContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetStakingRewardDistributedIterator is returned from FilterStakingRewardDistributed and is used to iterate over the raw logs and unpacked data for StakingRewardDistributed events raised by the RoninValidatorSet contract.
type RoninValidatorSetStakingRewardDistributedIterator struct {
	Event *RoninValidatorSetStakingRewardDistributed // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetStakingRewardDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetStakingRewardDistributed)
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
		it.Event = new(RoninValidatorSetStakingRewardDistributed)
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
func (it *RoninValidatorSetStakingRewardDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetStakingRewardDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetStakingRewardDistributed represents a StakingRewardDistributed event raised by the RoninValidatorSet contract.
type RoninValidatorSetStakingRewardDistributed struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakingRewardDistributed is a free log retrieval operation binding the contract event 0xeb09b8cc1cefa77cd4ec30003e6364cf60afcedd20be8c09f26e717788baf139.
//
// Solidity: event StakingRewardDistributed(uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterStakingRewardDistributed(opts *bind.FilterOpts) (*RoninValidatorSetStakingRewardDistributedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "StakingRewardDistributed")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetStakingRewardDistributedIterator{contract: _RoninValidatorSet.contract, event: "StakingRewardDistributed", logs: logs, sub: sub}, nil
}

// WatchStakingRewardDistributed is a free log subscription operation binding the contract event 0xeb09b8cc1cefa77cd4ec30003e6364cf60afcedd20be8c09f26e717788baf139.
//
// Solidity: event StakingRewardDistributed(uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchStakingRewardDistributed(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetStakingRewardDistributed) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "StakingRewardDistributed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetStakingRewardDistributed)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "StakingRewardDistributed", log); err != nil {
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

// ParseStakingRewardDistributed is a log parse operation binding the contract event 0xeb09b8cc1cefa77cd4ec30003e6364cf60afcedd20be8c09f26e717788baf139.
//
// Solidity: event StakingRewardDistributed(uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseStakingRewardDistributed(log types.Log) (*RoninValidatorSetStakingRewardDistributed, error) {
	event := new(RoninValidatorSetStakingRewardDistributed)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "StakingRewardDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetStakingVestingContractUpdatedIterator is returned from FilterStakingVestingContractUpdated and is used to iterate over the raw logs and unpacked data for StakingVestingContractUpdated events raised by the RoninValidatorSet contract.
type RoninValidatorSetStakingVestingContractUpdatedIterator struct {
	Event *RoninValidatorSetStakingVestingContractUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetStakingVestingContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetStakingVestingContractUpdated)
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
		it.Event = new(RoninValidatorSetStakingVestingContractUpdated)
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
func (it *RoninValidatorSetStakingVestingContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetStakingVestingContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetStakingVestingContractUpdated represents a StakingVestingContractUpdated event raised by the RoninValidatorSet contract.
type RoninValidatorSetStakingVestingContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterStakingVestingContractUpdated is a free log retrieval operation binding the contract event 0xc328090a37d855191ab58469296f98f87a851ca57d5cdfd1e9ac3c83e9e7096d.
//
// Solidity: event StakingVestingContractUpdated(address arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterStakingVestingContractUpdated(opts *bind.FilterOpts) (*RoninValidatorSetStakingVestingContractUpdatedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "StakingVestingContractUpdated")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetStakingVestingContractUpdatedIterator{contract: _RoninValidatorSet.contract, event: "StakingVestingContractUpdated", logs: logs, sub: sub}, nil
}

// WatchStakingVestingContractUpdated is a free log subscription operation binding the contract event 0xc328090a37d855191ab58469296f98f87a851ca57d5cdfd1e9ac3c83e9e7096d.
//
// Solidity: event StakingVestingContractUpdated(address arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchStakingVestingContractUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetStakingVestingContractUpdated) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "StakingVestingContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetStakingVestingContractUpdated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "StakingVestingContractUpdated", log); err != nil {
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

// ParseStakingVestingContractUpdated is a log parse operation binding the contract event 0xc328090a37d855191ab58469296f98f87a851ca57d5cdfd1e9ac3c83e9e7096d.
//
// Solidity: event StakingVestingContractUpdated(address arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseStakingVestingContractUpdated(log types.Log) (*RoninValidatorSetStakingVestingContractUpdated, error) {
	event := new(RoninValidatorSetStakingVestingContractUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "StakingVestingContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetValidatorCandidateAddedIterator is returned from FilterValidatorCandidateAdded and is used to iterate over the raw logs and unpacked data for ValidatorCandidateAdded events raised by the RoninValidatorSet contract.
type RoninValidatorSetValidatorCandidateAddedIterator struct {
	Event *RoninValidatorSetValidatorCandidateAdded // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetValidatorCandidateAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetValidatorCandidateAdded)
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
		it.Event = new(RoninValidatorSetValidatorCandidateAdded)
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
func (it *RoninValidatorSetValidatorCandidateAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetValidatorCandidateAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetValidatorCandidateAdded represents a ValidatorCandidateAdded event raised by the RoninValidatorSet contract.
type RoninValidatorSetValidatorCandidateAdded struct {
	ConsensusAddr common.Address
	TreasuryAddr  common.Address
	CandidateIdx  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterValidatorCandidateAdded is a free log retrieval operation binding the contract event 0x5ea0ccc37694ce2ce1e44e06663c8caff77c1ec661d991e2ece3a6195f879acf.
//
// Solidity: event ValidatorCandidateAdded(address indexed consensusAddr, address indexed treasuryAddr, uint256 indexed candidateIdx)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterValidatorCandidateAdded(opts *bind.FilterOpts, consensusAddr []common.Address, treasuryAddr []common.Address, candidateIdx []*big.Int) (*RoninValidatorSetValidatorCandidateAddedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var treasuryAddrRule []interface{}
	for _, treasuryAddrItem := range treasuryAddr {
		treasuryAddrRule = append(treasuryAddrRule, treasuryAddrItem)
	}
	var candidateIdxRule []interface{}
	for _, candidateIdxItem := range candidateIdx {
		candidateIdxRule = append(candidateIdxRule, candidateIdxItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "ValidatorCandidateAdded", consensusAddrRule, treasuryAddrRule, candidateIdxRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetValidatorCandidateAddedIterator{contract: _RoninValidatorSet.contract, event: "ValidatorCandidateAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorCandidateAdded is a free log subscription operation binding the contract event 0x5ea0ccc37694ce2ce1e44e06663c8caff77c1ec661d991e2ece3a6195f879acf.
//
// Solidity: event ValidatorCandidateAdded(address indexed consensusAddr, address indexed treasuryAddr, uint256 indexed candidateIdx)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchValidatorCandidateAdded(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetValidatorCandidateAdded, consensusAddr []common.Address, treasuryAddr []common.Address, candidateIdx []*big.Int) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var treasuryAddrRule []interface{}
	for _, treasuryAddrItem := range treasuryAddr {
		treasuryAddrRule = append(treasuryAddrRule, treasuryAddrItem)
	}
	var candidateIdxRule []interface{}
	for _, candidateIdxItem := range candidateIdx {
		candidateIdxRule = append(candidateIdxRule, candidateIdxItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "ValidatorCandidateAdded", consensusAddrRule, treasuryAddrRule, candidateIdxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetValidatorCandidateAdded)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "ValidatorCandidateAdded", log); err != nil {
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

// ParseValidatorCandidateAdded is a log parse operation binding the contract event 0x5ea0ccc37694ce2ce1e44e06663c8caff77c1ec661d991e2ece3a6195f879acf.
//
// Solidity: event ValidatorCandidateAdded(address indexed consensusAddr, address indexed treasuryAddr, uint256 indexed candidateIdx)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseValidatorCandidateAdded(log types.Log) (*RoninValidatorSetValidatorCandidateAdded, error) {
	event := new(RoninValidatorSetValidatorCandidateAdded)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "ValidatorCandidateAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetValidatorCandidateRemovedIterator is returned from FilterValidatorCandidateRemoved and is used to iterate over the raw logs and unpacked data for ValidatorCandidateRemoved events raised by the RoninValidatorSet contract.
type RoninValidatorSetValidatorCandidateRemovedIterator struct {
	Event *RoninValidatorSetValidatorCandidateRemoved // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetValidatorCandidateRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetValidatorCandidateRemoved)
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
		it.Event = new(RoninValidatorSetValidatorCandidateRemoved)
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
func (it *RoninValidatorSetValidatorCandidateRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetValidatorCandidateRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetValidatorCandidateRemoved represents a ValidatorCandidateRemoved event raised by the RoninValidatorSet contract.
type RoninValidatorSetValidatorCandidateRemoved struct {
	ConsensusAddr common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterValidatorCandidateRemoved is a free log retrieval operation binding the contract event 0xec1ae6fd172c65ee5804ae3dafed7d57ec0f1e1183f910a8b7bd4abe31110f15.
//
// Solidity: event ValidatorCandidateRemoved(address indexed consensusAddr)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterValidatorCandidateRemoved(opts *bind.FilterOpts, consensusAddr []common.Address) (*RoninValidatorSetValidatorCandidateRemovedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "ValidatorCandidateRemoved", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetValidatorCandidateRemovedIterator{contract: _RoninValidatorSet.contract, event: "ValidatorCandidateRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorCandidateRemoved is a free log subscription operation binding the contract event 0xec1ae6fd172c65ee5804ae3dafed7d57ec0f1e1183f910a8b7bd4abe31110f15.
//
// Solidity: event ValidatorCandidateRemoved(address indexed consensusAddr)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchValidatorCandidateRemoved(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetValidatorCandidateRemoved, consensusAddr []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "ValidatorCandidateRemoved", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetValidatorCandidateRemoved)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "ValidatorCandidateRemoved", log); err != nil {
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

// ParseValidatorCandidateRemoved is a log parse operation binding the contract event 0xec1ae6fd172c65ee5804ae3dafed7d57ec0f1e1183f910a8b7bd4abe31110f15.
//
// Solidity: event ValidatorCandidateRemoved(address indexed consensusAddr)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseValidatorCandidateRemoved(log types.Log) (*RoninValidatorSetValidatorCandidateRemoved, error) {
	event := new(RoninValidatorSetValidatorCandidateRemoved)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "ValidatorCandidateRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetValidatorPunishedIterator is returned from FilterValidatorPunished and is used to iterate over the raw logs and unpacked data for ValidatorPunished events raised by the RoninValidatorSet contract.
type RoninValidatorSetValidatorPunishedIterator struct {
	Event *RoninValidatorSetValidatorPunished // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetValidatorPunishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetValidatorPunished)
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
		it.Event = new(RoninValidatorSetValidatorPunished)
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
func (it *RoninValidatorSetValidatorPunishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetValidatorPunishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetValidatorPunished represents a ValidatorPunished event raised by the RoninValidatorSet contract.
type RoninValidatorSetValidatorPunished struct {
	ValidatorAddr         common.Address
	JailedUntil           *big.Int
	DeductedStakingAmount *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterValidatorPunished is a free log retrieval operation binding the contract event 0x69284547cc931ff0e04d5a21cdfb0748f22a3788269711028ce1d4833900e474.
//
// Solidity: event ValidatorPunished(address validatorAddr, uint256 jailedUntil, uint256 deductedStakingAmount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterValidatorPunished(opts *bind.FilterOpts) (*RoninValidatorSetValidatorPunishedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "ValidatorPunished")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetValidatorPunishedIterator{contract: _RoninValidatorSet.contract, event: "ValidatorPunished", logs: logs, sub: sub}, nil
}

// WatchValidatorPunished is a free log subscription operation binding the contract event 0x69284547cc931ff0e04d5a21cdfb0748f22a3788269711028ce1d4833900e474.
//
// Solidity: event ValidatorPunished(address validatorAddr, uint256 jailedUntil, uint256 deductedStakingAmount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchValidatorPunished(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetValidatorPunished) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "ValidatorPunished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetValidatorPunished)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "ValidatorPunished", log); err != nil {
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

// ParseValidatorPunished is a log parse operation binding the contract event 0x69284547cc931ff0e04d5a21cdfb0748f22a3788269711028ce1d4833900e474.
//
// Solidity: event ValidatorPunished(address validatorAddr, uint256 jailedUntil, uint256 deductedStakingAmount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseValidatorPunished(log types.Log) (*RoninValidatorSetValidatorPunished, error) {
	event := new(RoninValidatorSetValidatorPunished)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "ValidatorPunished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetValidatorSetUpdatedIterator is returned from FilterValidatorSetUpdated and is used to iterate over the raw logs and unpacked data for ValidatorSetUpdated events raised by the RoninValidatorSet contract.
type RoninValidatorSetValidatorSetUpdatedIterator struct {
	Event *RoninValidatorSetValidatorSetUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetValidatorSetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetValidatorSetUpdated)
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
		it.Event = new(RoninValidatorSetValidatorSetUpdated)
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
func (it *RoninValidatorSetValidatorSetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetValidatorSetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetValidatorSetUpdated represents a ValidatorSetUpdated event raised by the RoninValidatorSet contract.
type RoninValidatorSetValidatorSetUpdated struct {
	Arg0 []common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetUpdated is a free log retrieval operation binding the contract event 0x6120448f3d4245c1ff708d970c34e1b6484ee22a794ede0bfca2317a97aa8ced.
//
// Solidity: event ValidatorSetUpdated(address[] arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterValidatorSetUpdated(opts *bind.FilterOpts) (*RoninValidatorSetValidatorSetUpdatedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "ValidatorSetUpdated")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetValidatorSetUpdatedIterator{contract: _RoninValidatorSet.contract, event: "ValidatorSetUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorSetUpdated is a free log subscription operation binding the contract event 0x6120448f3d4245c1ff708d970c34e1b6484ee22a794ede0bfca2317a97aa8ced.
//
// Solidity: event ValidatorSetUpdated(address[] arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchValidatorSetUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetValidatorSetUpdated) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "ValidatorSetUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetValidatorSetUpdated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "ValidatorSetUpdated", log); err != nil {
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

// ParseValidatorSetUpdated is a log parse operation binding the contract event 0x6120448f3d4245c1ff708d970c34e1b6484ee22a794ede0bfca2317a97aa8ced.
//
// Solidity: event ValidatorSetUpdated(address[] arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseValidatorSetUpdated(log types.Log) (*RoninValidatorSetValidatorSetUpdated, error) {
	event := new(RoninValidatorSetValidatorSetUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "ValidatorSetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
