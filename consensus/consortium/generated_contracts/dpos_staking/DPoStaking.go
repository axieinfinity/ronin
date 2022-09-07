// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dposStaking

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

// IStakingValidatorCandidate is an auto generated low-level Go binding around an user-defined struct.
type IStakingValidatorCandidate struct {
	CandidateAdmin  common.Address
	ConsensusAddr   common.Address
	TreasuryAddr    common.Address
	CommissionRate  *big.Int
	StakedAmount    *big.Int
	DelegatedAmount *big.Int
	Governing       bool
	State           uint8
	Gap             [20]*big.Int
}

// DposStakingMetaData contains all meta data concerning the DposStaking contract.
var DposStakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Delegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"GovernanceAdminContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"MaxValidatorCandidateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"MinValidatorBalanceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accumulatedRps\",\"type\":\"uint256\"}],\"name\":\"PendingPoolUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debited\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"credited\",\"type\":\"uint256\"}],\"name\":\"PendingRewardUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"poolAddress\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"accumulatedRps\",\"type\":\"uint256[]\"}],\"name\":\"SettledPoolsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debited\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accumulatedRps\",\"type\":\"uint256\"}],\"name\":\"SettledRewardUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Undelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ValidatorContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"candidateIdx\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"candidateAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"treasuryAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatedAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"governing\",\"type\":\"bool\"},{\"internalType\":\"enumIStaking.ValidatorState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256[20]\",\"name\":\"____gap\",\"type\":\"uint256[20]\"}],\"indexed\":false,\"internalType\":\"structIStaking.ValidatorCandidate\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"ValidatorProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ValidatorRenounceFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ValidatorRenounceRequested\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"_claimReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_consensusAddrList\",\"type\":\"address[]\"}],\"name\":\"claimRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"}],\"name\":\"commissionRateOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deductStakingAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_consensusAddrList\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_consensusAddrDst\",\"type\":\"address\"}],\"name\":\"delegateRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCandidateWeights\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_candidates\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_weights\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getClaimableReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getPendingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_poolAddrList\",\"type\":\"address[]\"}],\"name\":\"getRewards\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_pendings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_claimables\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getTotalReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorCandidates\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"candidateAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"treasuryAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatedAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"governing\",\"type\":\"bool\"},{\"internalType\":\"enumIStaking.ValidatorState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256[20]\",\"name\":\"____gap\",\"type\":\"uint256[20]\"}],\"internalType\":\"structIStaking.ValidatorCandidate[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceAdminContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__validatorContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__governanceAdminContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"__maxValidatorCandidate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"__minValidatorBalance\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxValidatorCandidate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minValidatorBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_treasuryAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_commissionRate\",\"type\":\"uint256\"}],\"name\":\"proposeValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_candidateIdx\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"}],\"name\":\"recordReward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddrSrc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_consensusAddrDst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"redelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"}],\"name\":\"renounce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"setMaxValidatorCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"setMinValidatorBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_consensusAddrs\",\"type\":\"address[]\"}],\"name\":\"settleRewardPools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"}],\"name\":\"sinkPendingReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolAddr\",\"type\":\"address\"}],\"name\":\"totalBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"}],\"name\":\"treasuryAddressOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatorCandidates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"candidateAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"treasuryAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatedAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"governing\",\"type\":\"bool\"},{\"internalType\":\"enumIStaking.ValidatorState\",\"name\":\"state\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50620000226200002860201b60201c565b620001d6565b600560019054906101000a900460ff16156200007b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000729062000179565b60405180910390fd5b60ff8016600560009054906101000a900460ff1660ff161015620000f05760ff600560006101000a81548160ff021916908360ff1602179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860ff604051620000e79190620001b9565b60405180910390a15b565b600082825260208201905092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320696e69746960008201527f616c697a696e6700000000000000000000000000000000000000000000000000602082015250565b600062000161602783620000f2565b91506200016e8262000103565b604082019050919050565b60006020820190508181036000830152620001948162000152565b9050919050565b600060ff82169050919050565b620001b3816200019b565b82525050565b6000602082019050620001d06000830184620001a8565b92915050565b6158b080620001e66000396000f3fe6080604052600436106101dc5760003560e01c8063a5885f1d11610102578063e8a22a8a11610095578063f7888aec11610064578063f7888aec146107ff578063f9f031df1461083c578063fe0f3a1314610879578063fe7732f4146108b657610273565b8063e8a22a8a14610742578063ea11bf831461077f578063ea82f784146107ab578063eb990c59146107d657610273565b8063c905bb35116100d1578063c905bb3514610688578063ce99b586146106b1578063d45e6273146106dc578063dac8ef491461070557610273565b8063a5885f1d146105ef578063b863d71014610618578063ba77b06c14610634578063c2a672e01461065f57610273565b8063470126b01161017a5780635c19a95c116101495780635c19a95c14610542578063605239a11461055e5780636bd8f804146105895780636eacd398146105b257610273565b8063470126b0146104835780634d99dd16146104b35780634f2a693f146104dc5780635a7836561461050557610273565b80631f76a7af116101b65780631f76a7af146103c357806321e91dea146103ec57806326476204146104295780633d8e846e1461044557610273565b8063097e4a9d14610305578063099feccb146103425780631104b7501461037f57610273565b3661027357600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610271576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026890614055565b60405180910390fd5b005b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610303576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102fa90614055565b60405180910390fd5b005b34801561031157600080fd5b5061032c60048036038101906103279190614142565b6108df565b60405161033991906141bb565b60405180910390f35b34801561034e57600080fd5b50610369600480360381019061036491906141d6565b610b3f565b6040516103769190614212565b60405180910390f35b34801561038b57600080fd5b506103a660048036038101906103a19190614259565b610b7a565b6040516103ba989796959493929190614339565b60405180910390f35b3480156103cf57600080fd5b506103ea60048036038101906103e591906141d6565b610c4c565b005b3480156103f857600080fd5b50610413600480360381019061040e91906143b7565b610c87565b60405161042091906141bb565b60405180910390f35b610443600480360381019061043e91906141d6565b610f36565b005b34801561045157600080fd5b5061046c600480360381019061046791906143f7565b610f87565b60405161047a929190614515565b60405180910390f35b61049d60048036038101906104989190614578565b61104d565b6040516104aa91906141bb565b60405180910390f35b3480156104bf57600080fd5b506104da60048036038101906104d591906145cb565b61107b565b005b3480156104e857600080fd5b5061050360048036038101906104fe9190614259565b61134d565b005b34801561051157600080fd5b5061052c600480360381019061052791906141d6565b6113e9565b60405161053991906141bb565b60405180910390f35b61055c600480360381019061055791906141d6565b611404565b005b34801561056a57600080fd5b5061057361169e565b60405161058091906141bb565b60405180910390f35b34801561059557600080fd5b506105b060048036038101906105ab919061460b565b6116a8565b005b3480156105be57600080fd5b506105d960048036038101906105d491906141d6565b611912565b6040516105e691906141bb565b60405180910390f35b3480156105fb57600080fd5b50610616600480360381019061061191906141d6565b61192d565b005b610632600480360381019061062d91906145cb565b6119b0565b005b34801561064057600080fd5b50610649611a4e565b6040516106569190614885565b60405180910390f35b34801561066b57600080fd5b50610686600480360381019061068191906145cb565b611c69565b005b34801561069457600080fd5b506106af60048036038101906106aa91906145cb565b611cf2565b005b3480156106bd57600080fd5b506106c6611df2565b6040516106d391906141bb565b60405180910390f35b3480156106e857600080fd5b5061070360048036038101906106fe91906148a7565b611dfc565b005b34801561071157600080fd5b5061072c600480360381019061072791906143b7565b611ea7565b60405161073991906141bb565b60405180910390f35b34801561074e57600080fd5b50610769600480360381019061076491906143b7565b612122565b60405161077691906141bb565b60405180910390f35b34801561078b57600080fd5b506107946123bc565b6040516107a29291906149a3565b60405180910390f35b3480156107b757600080fd5b506107c0612540565b6040516107cd9190614212565b60405180910390f35b3480156107e257600080fd5b506107fd60048036038101906107f891906149da565b61256a565b005b34801561080b57600080fd5b50610826600480360381019061082191906143b7565b6126cf565b60405161083391906141bb565b60405180910390f35b34801561084857600080fd5b50610863600480360381019061085e91906148a7565b612756565b60405161087091906141bb565b60405180910390f35b34801561088557600080fd5b506108a0600480360381019061089b91906143b7565b6127df565b6040516108ad91906141bb565b60405180910390f35b3480156108c257600080fd5b506108dd60048036038101906108d89190614259565b612807565b005b60008160006108ed826128a3565b604051806101200160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820154815260200160048201548152602001600582015481526020016006820160009054906101000a900460ff161515151581526020016006820160019054906101000a900460ff166003811115610a5957610a586142c2565b5b6003811115610a6b57610a6a6142c2565b5b815260200160078201601480602002604051908101604052809291908260148015610aab576020028201915b815481526020019060010190808311610a97575b5050505050815250509050806000015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1603610b28576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b1f90614ab3565b60405180910390fd5b610b3433878787612956565b925050509392505050565b600080610b4b836128a3565b90508060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16915050919050565b600b8181548110610b8a57600080fd5b90600052602060002090601b02016000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154908060040154908060050154908060060160009054906101000a900460ff16908060060160019054906101000a900460ff16905088565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c7e90614b1f565b60405180910390fd5b600080600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060600160405290816000820154815260200160018201548152602001600282015481525050905060008060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806060016040529081600082015481526020016001820154815260200160028201548152505090506000600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060400160405290816000820154815260200160018201548152505090508060000151836040015111610ebc576000610e5687876126cf565b905083602001518460000151670de0b6b3a7640000846020015184610e7b9190614b6e565b610e859190614bf7565b610e8f9190614c28565b610e999190614c5c565b836020018181525050808360000181815250508160200151836040018181525050505b6000826000015190506000670de0b6b3a7640000846040015183610ee09190614b6e565b610eea9190614bf7565b9050808460200151670de0b6b3a7640000856020015185610f0b9190614b6e565b610f159190614bf7565b610f1f9190614c28565b610f299190614c5c565b9550505050505092915050565b60003411610f79576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f7090614d02565b60405180910390fd5b610f84813334612978565b50565b606080600080600090505b8585905081101561104357858582818110610fb057610faf614d22565b5b9050602002016020810190610fc591906141d6565b91506000610fd38389612122565b90506000610fe1848a610c87565b90508082610fef9190614c5c565b86848151811061100257611001614d22565b5b6020026020010181815250508085848151811061102257611021614d22565b5b6020026020010181815250505050808061103b90614d51565b915050610f92565b5050935093915050565b60008034905060003390506110658686868585612a91565b9250611072868284612978565b50509392505050565b816000611087826128a3565b604051806101200160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820154815260200160048201548152602001600582015481526020016006820160009054906101000a900460ff161515151581526020016006820160019054906101000a900460ff1660038111156111f3576111f26142c2565b5b6003811115611205576112046142c2565b5b815260200160078201601480602002604051908101604052809291908260148015611245576020028201915b815481526020019060010190808311611231575b5050505050815250509050806000015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16036112c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112b990614ab3565b60405180910390fd5b60003390506112d2858286612c81565b8073ffffffffffffffffffffffffffffffffffffffff166108fc859081150290604051600060405180830381858888f19350505050611346576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161133d90614e0b565b60405180910390fd5b5050505050565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146113dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113d490614e9d565b60405180910390fd5b6113e681612eec565b50565b6000806113f5836128a3565b90508060030154915050919050565b60003411611447576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161143e90614d02565b60405180910390fd5b806000611453826128a3565b604051806101200160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820154815260200160048201548152602001600582015481526020016006820160009054906101000a900460ff161515151581526020016006820160019054906101000a900460ff1660038111156115bf576115be6142c2565b5b60038111156115d1576115d06142c2565b5b815260200160078201601480602002604051908101604052809291908260148015611611576020028201915b8154815260200190600101908083116115fd575b5050505050815250509050806000015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff160361168e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161168590614ab3565b60405180910390fd5b611699833334612f2d565b505050565b6000600754905090565b8160006116b4826128a3565b604051806101200160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820154815260200160048201548152602001600582015481526020016006820160009054906101000a900460ff161515151581526020016006820160019054906101000a900460ff1660038111156118205761181f6142c2565b5b6003811115611832576118316142c2565b5b815260200160078201601480602002604051908101604052809291908260148015611872576020028201915b81548152602001906001019080831161185e575b5050505050815250509050806000015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16036118ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118e690614ab3565b60405180910390fd5b60003390506118ff868286612c81565b61190a858286612f2d565b505050505050565b60008061191e836128a3565b90508060050154915050919050565b6000611938436130d9565b90506001600c60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600083815260200190815260200160002060006101000a81548160ff0219169083151502179055506119ac8261317e565b5050565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611a40576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a3790614055565b60405180910390fd5b611a4a8282613258565b5050565b6060600b805480602002602001604051908101604052809291908181526020016000905b82821015611c6057838290600052602060002090601b0201604051806101200160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820154815260200160048201548152602001600582015481526020016006820160009054906101000a900460ff161515151581526020016006820160019054906101000a900460ff166003811115611bf657611bf56142c2565b5b6003811115611c0857611c076142c2565b5b815260200160078201601480602002604051908101604052809291908260148015611c48576020028201915b815481526020019060010190808311611c34575b50505050508152505081526020019060010190611a72565b50505050905090565b6000339050611c79838284613326565b3373ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f19350505050611ced576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ce490614e0b565b60405180910390fd5b505050565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611d82576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d7990614055565b60405180910390fd5b6000611d8d836128a3565b9050611dbe838260000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684613326565b611ded838260000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684612c81565b505050565b6000600654905090565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611e8c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e8390614055565b60405180910390fd5b6000828290500315611ea357611ea282826134e5565b5b5050565b6000611eb38383610c87565b90507f0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7838383604051611ee893929190614ebd565b60405180910390a16000600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060400160405290816000820154815260200160018201548152505090506000600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905082816001016000828254611fe69190614c28565b925050819055504381600201819055507fd6faab37181ad960bdaaf04de2913a71085ce9224f4575265ec261d57419d4778585836000015484600101546040516120339493929190614ef4565b60405180910390a160008060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060008160010181905550826020015181600201819055507f0b74380c3e212861270a8670337fc8147462c92a81a228d1dc50d3caa0ab62c58686836000015460008760200151604051612111959493929190614f7e565b60405180910390a150505092915050565b600080600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806060016040529081600082015481526020016001820154815260200160028201548152505090506000600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806040016040529081600082015481526020016001820154815250509050600061223a86866126cf565b90506122528661224d85604001516130d9565b61367f565b1561236f5760008060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806060016040529081600082015481526020016001820154815260200160028201548152505090506000670de0b6b3a764000083836040015161231a9190614b6e565b6123249190614bf7565b9050808260200151670de0b6b3a76400008660200151866123459190614b6e565b61234f9190614bf7565b6123599190614c28565b6123639190614c5c565b955050505050506123b6565b82602001518360000151670de0b6b3a76400008460200151846123929190614b6e565b61239c9190614bf7565b6123a69190614c28565b6123b09190614c5c565b93505050505b92915050565b6060806000600b8054905090508067ffffffffffffffff8111156123e3576123e2614fd1565b5b6040519080825280602002602001820160405280156124115781602001602082028036833780820191505090505b5092508067ffffffffffffffff81111561242e5761242d614fd1565b5b60405190808252806020026020018201604052801561245c5781602001602082028036833780820191505090505b50915060005b8181101561253a576000600b82815481106124805761247f614d22565b5b90600052602060002090601b020190508060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168583815181106124c8576124c7614d22565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050806005015484838151811061251a57612519614d22565b5b60200260200101818152505050808061253290614d51565b915050612462565b50509091565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600560019054906101000a900460ff1615905080801561259e57506001600560009054906101000a900460ff1660ff16105b806125cd57506125ad306136e7565b1580156125cc57506001600560009054906101000a900460ff1660ff16145b5b61260c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161260390615072565b60405180910390fd5b6001600560006101000a81548160ff021916908360ff160217905550801561264a576001600560016101000a81548160ff0219169083151502179055505b6126538561370a565b61265c84613785565b61266583612eec565b61266e82613800565b80156126c8576000600560016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516126bf91906150da565b60405180910390a15b5050505050565b6000600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6000612763338484613841565b90503373ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050506127d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016127d090614e0b565b60405180910390fd5b92915050565b60006127eb8383610c87565b6127f58484612122565b6127ff9190614c5c565b905092915050565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612897576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161288e90614e9d565b60405180910390fd5b6128a081613800565b50565b600080600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000811161292b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161292290615167565b60405180910390fd5b600b8119815481106129405761293f614d22565b5b90600052602060002090601b0201915050919050565b6000612963858585613841565b9050612970828683612f2d565b949350505050565b6000612983846128a3565b90508273ffffffffffffffffffffffffffffffffffffffff168160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614612a17576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a0e906151f9565b60405180910390fd5b81816004016000828254612a2b9190614c28565b925050819055508373ffffffffffffffffffffffffffffffffffffffff167f9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d83604051612a7891906141bb565b60405180910390a2612a8b848484612f2d565b50505050565b600080612a9c6138ad565b9050612aa661169e565b8110612ae7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612ade9061528b565b60405180910390fd5b6000612af2886138ba565b14612b32576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b299061531d565b60405180910390fd5b612b3a611df2565b841015612b7c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b73906153af565b60405180910390fd5b8573ffffffffffffffffffffffffffffffffffffffff166108fc60009081150290604051600060405180830381858888f19350505050612bf1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612be890615441565b60405180910390fd5b80199150612bff8783613903565b6000612c0d8885898961394b565b90508373ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f48a4728ce2a28ba5a3060eafbd15c8f72563d2dda3e74c24cbf14bc6c68e03db8784604051612c6e929190615518565b60405180910390a3505095945050505050565b80600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015612d40576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d37906155b4565b60405180910390fd5b600081600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054612dca9190614c5c565b9050612dd7848483613c1c565b6000612de2856128a3565b905082816005016000828254612df89190614c5c565b9250508190555081600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508473ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c85604051612edd91906141bb565b60405180910390a35050505050565b806007819055507f82d5dc32d1b741512ad09c32404d7e7921e8934c6222343d95f55f7a2b9b2ab481604051612f2291906141bb565b60405180910390a150565b600081600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054612fb79190614c28565b9050612fc4848483613c1c565b6000612fcf856128a3565b905082816005016000828254612fe59190614c28565b9250508190555081600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508473ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b856040516130ca91906141bb565b60405180910390a35050505050565b6000600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f8549af9836040518263ffffffff1660e01b815260040161313691906141bb565b602060405180830381865afa158015613153573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061317791906155e9565b9050919050565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015490506000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508181600101819055504381600001819055507f5791918d066f160d5cfcc4ccd4df1f7f4f0a0f48b486a2ef33e238c1369f2224838360405161324b929190615616565b60405180910390a1505050565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060006132a684611912565b670de0b6b3a7640000846132ba9190614b6e565b6132c49190614bf7565b82600101546132d39190614c28565b90508082600101819055504382600001819055507f5791918d066f160d5cfcc4ccd4df1f7f4f0a0f48b486a2ef33e238c1369f22248482604051613318929190615616565b60405180910390a150505050565b6000613331846128a3565b90508273ffffffffffffffffffffffffffffffffffffffff168160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146133c5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133bc906151f9565b60405180910390fd5b806004015482111561340c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613403906156b1565b60405180910390fd5b600082826004015461341e9190614c5c565b9050613428611df2565b81101561346a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161346190615743565b60405180910390fd5b8282600401600082825461347e9190614c5c565b925050819055508473ffffffffffffffffffffffffffffffffffffffff167f0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75846040516134cb91906141bb565b60405180910390a26134de858585612c81565b5050505050565b60008282905067ffffffffffffffff81111561350457613503614fd1565b5b6040519080825280602002602001820160405280156135325781602001602082028036833780820191505090505b50905060008060005b8585905081101561363c5785858281811061355957613558614d22565b5b905060200201602081019061356e91906141d6565b9250600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101548482815181106135c6576135c5614d22565b5b602002602001018181525091506000600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905082816001018190555043816000018190555050808061363490614d51565b91505061353b565b507fed065952215f1a6f0ff936480b7ae27e204860bd6e35aca6a678cce83dc0c490858585604051613670939291906157ee565b60405180910390a15050505050565b6000600c60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600083815260200190815260200160002060009054906101000a900460ff16905092915050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b80600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b68361698160405161377a9190614212565b60405180910390a150565b80600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f2374c4e57a97e792cc22ffb7ddfc5f526897b2e22738b569da0e8f5dd17e1f34816040516137f59190614212565b60405180910390a150565b806006819055507f8364d8933fc5aad4a5ae33139fea4cbc485c4f1931e822d06cd3b873c13a07e68160405161383691906141bb565b60405180910390a150565b600080600090505b838390508110156138a55761388584848381811061386a57613869614d22565b5b905060200201602081019061387f91906141d6565b86611ea7565b826138909190614c28565b9150808061389d90614d51565b915050613849565b509392505050565b6000600b80549050905090565b6000600a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b80600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050565b613953613f07565b6000600b6001816001815401808255809150500390600052602060002090601b02019050858160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550848160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082816003018190555080604051806101200160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820154815260200160048201548152602001600582015481526020016006820160009054906101000a900460ff161515151581526020016006820160019054906101000a900460ff166003811115613bb657613bb56142c2565b5b6003811115613bc857613bc76142c2565b5b815260200160078201601480602002604051908101604052809291908260148015613c08576020028201915b815481526020019060010190808311613bf4575b505050505081525050915050949350505050565b6000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060400160405290816000820154815260200160018201548152505090508060000151826002015411613e0e576000613d198686610c87565b90506000613d2787876126cf565b905060008060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050818160000181905550828160010181905550836020015181600201819055507f0b74380c3e212861270a8670337fc8147462c92a81a228d1dc50d3caa0ab62c5888884868860200151604051613e02959493929190615827565b60405180910390a15050505b6000600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060400160405290816000820154815260200160018201548152505090506000613e7c8787612122565b90506000670de0b6b3a7640000836020015187613e999190614b6e565b613ea39190614bf7565b90508185600001819055508085600101819055504385600201819055507fd6faab37181ad960bdaaf04de2913a71085ce9224f4575265ec261d57419d47788888484604051613ef59493929190614ef4565b60405180910390a15050505050505050565b604051806101200160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081526020016000815260200160001515815260200160006003811115613f9c57613f9b6142c2565b5b8152602001613fa9613faf565b81525090565b604051806102800160405280601490602082028036833780820191505090505090565b600082825260208201905092915050565b7f44506f5374616b696e673a206d6574686f642063616c6c6572206973206e6f7460008201527f207468652076616c696461746f7220636f6e7472616374000000000000000000602082015250565b600061403f603783613fd2565b915061404a82613fe3565b604082019050919050565b6000602082019050818103600083015261406e81614032565b9050919050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b60008083601f8401126140a4576140a361407f565b5b8235905067ffffffffffffffff8111156140c1576140c0614084565b5b6020830191508360208202830111156140dd576140dc614089565b5b9250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061410f826140e4565b9050919050565b61411f81614104565b811461412a57600080fd5b50565b60008135905061413c81614116565b92915050565b60008060006040848603121561415b5761415a614075565b5b600084013567ffffffffffffffff8111156141795761417861407a565b5b6141858682870161408e565b935093505060206141988682870161412d565b9150509250925092565b6000819050919050565b6141b5816141a2565b82525050565b60006020820190506141d060008301846141ac565b92915050565b6000602082840312156141ec576141eb614075565b5b60006141fa8482850161412d565b91505092915050565b61420c81614104565b82525050565b60006020820190506142276000830184614203565b92915050565b614236816141a2565b811461424157600080fd5b50565b6000813590506142538161422d565b92915050565b60006020828403121561426f5761426e614075565b5b600061427d84828501614244565b91505092915050565b6000614291826140e4565b9050919050565b6142a181614286565b82525050565b60008115159050919050565b6142bc816142a7565b82525050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60048110614302576143016142c2565b5b50565b6000819050614313826142f1565b919050565b600061432382614305565b9050919050565b61433381614318565b82525050565b60006101008201905061434f600083018b614203565b61435c602083018a614203565b6143696040830189614298565b61437660608301886141ac565b61438360808301876141ac565b61439060a08301866141ac565b61439d60c08301856142b3565b6143aa60e083018461432a565b9998505050505050505050565b600080604083850312156143ce576143cd614075565b5b60006143dc8582860161412d565b92505060206143ed8582860161412d565b9150509250929050565b6000806000604084860312156144105761440f614075565b5b600061441e8682870161412d565b935050602084013567ffffffffffffffff81111561443f5761443e61407a565b5b61444b8682870161408e565b92509250509250925092565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61448c816141a2565b82525050565b600061449e8383614483565b60208301905092915050565b6000602082019050919050565b60006144c282614457565b6144cc8185614462565b93506144d783614473565b8060005b838110156145085781516144ef8882614492565b97506144fa836144aa565b9250506001810190506144db565b5085935050505092915050565b6000604082019050818103600083015261452f81856144b7565b9050818103602083015261454381846144b7565b90509392505050565b61455581614286565b811461456057600080fd5b50565b6000813590506145728161454c565b92915050565b60008060006060848603121561459157614590614075565b5b600061459f8682870161412d565b93505060206145b086828701614563565b92505060406145c186828701614244565b9150509250925092565b600080604083850312156145e2576145e1614075565b5b60006145f08582860161412d565b925050602061460185828601614244565b9150509250929050565b60008060006060848603121561462457614623614075565b5b60006146328682870161412d565b93505060206146438682870161412d565b925050604061465486828701614244565b9150509250925092565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61469381614104565b82525050565b6146a281614286565b82525050565b6146b1816142a7565b82525050565b6146c081614318565b82525050565b600060149050919050565b600081905092915050565b6000819050919050565b6000602082019050919050565b6146fc816146c6565b61470681846146d1565b9250614711826146dc565b8060005b838110156147425781516147298782614492565b9650614734836146e6565b925050600181019050614715565b505050505050565b61038082016000820151614761600085018261468a565b506020820151614774602085018261468a565b5060408201516147876040850182614699565b50606082015161479a6060850182614483565b5060808201516147ad6080850182614483565b5060a08201516147c060a0850182614483565b5060c08201516147d360c08501826146a8565b5060e08201516147e660e08501826146b7565b506101008201516147fb6101008501826146f3565b50505050565b600061480d838361474a565b6103808301905092915050565b6000602082019050919050565b60006148328261465e565b61483c8185614669565b93506148478361467a565b8060005b8381101561487857815161485f8882614801565b975061486a8361481a565b92505060018101905061484b565b5085935050505092915050565b6000602082019050818103600083015261489f8184614827565b905092915050565b600080602083850312156148be576148bd614075565b5b600083013567ffffffffffffffff8111156148dc576148db61407a565b5b6148e88582860161408e565b92509250509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600061492c838361468a565b60208301905092915050565b6000602082019050919050565b6000614950826148f4565b61495a81856148ff565b935061496583614910565b8060005b8381101561499657815161497d8882614920565b975061498883614938565b925050600181019050614969565b5085935050505092915050565b600060408201905081810360008301526149bd8185614945565b905081810360208301526149d181846144b7565b90509392505050565b600080600080608085870312156149f4576149f3614075565b5b6000614a028782880161412d565b9450506020614a138782880161412d565b9350506040614a2487828801614244565b9250506060614a3587828801614244565b91505092959194509250565b7f5374616b696e674d616e616765723a206d6574686f642063616c6c6572206d7560008201527f7374206e6f74206265207468652063616e6469646174652061646d696e000000602082015250565b6000614a9d603d83613fd2565b9150614aa882614a41565b604082019050919050565b60006020820190508181036000830152614acc81614a90565b9050919050565b7f756e696d706c656d656e74656400000000000000000000000000000000000000600082015250565b6000614b09600d83613fd2565b9150614b1482614ad3565b602082019050919050565b60006020820190508181036000830152614b3881614afc565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000614b79826141a2565b9150614b84836141a2565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615614bbd57614bbc614b3f565b5b828202905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000614c02826141a2565b9150614c0d836141a2565b925082614c1d57614c1c614bc8565b5b828204905092915050565b6000614c33826141a2565b9150614c3e836141a2565b9250828201905080821115614c5657614c55614b3f565b5b92915050565b6000614c67826141a2565b9150614c72836141a2565b9250828203905081811115614c8a57614c89614b3f565b5b92915050565b7f5374616b696e674d616e616765723a207175657279207769746820656d70747960008201527f2076616c75650000000000000000000000000000000000000000000000000000602082015250565b6000614cec602683613fd2565b9150614cf782614c90565b604082019050919050565b60006020820190508181036000830152614d1b81614cdf565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000614d5c826141a2565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614d8e57614d8d614b3f565b5b600182019050919050565b7f5374616b696e674d616e616765723a20636f756c64206e6f74207472616e736660008201527f657220524f4e0000000000000000000000000000000000000000000000000000602082015250565b6000614df5602683613fd2565b9150614e0082614d99565b604082019050919050565b60006020820190508181036000830152614e2481614de8565b9050919050565b7f44506f5374616b696e673a206d6574686f642063616c6c6572206973206e6f7460008201527f20676f7665726e616e63652061646d696e20636f6e7472616374000000000000602082015250565b6000614e87603a83613fd2565b9150614e9282614e2b565b604082019050919050565b60006020820190508181036000830152614eb681614e7a565b9050919050565b6000606082019050614ed26000830186614203565b614edf6020830185614203565b614eec60408301846141ac565b949350505050565b6000608082019050614f096000830187614203565b614f166020830186614203565b614f2360408301856141ac565b614f3060608301846141ac565b95945050505050565b6000819050919050565b6000819050919050565b6000614f68614f63614f5e84614f39565b614f43565b6141a2565b9050919050565b614f7881614f4d565b82525050565b600060a082019050614f936000830188614203565b614fa06020830187614203565b614fad60408301866141ac565b614fba6060830185614f6f565b614fc760808301846141ac565b9695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b600061505c602e83613fd2565b915061506782615000565b604082019050919050565b6000602082019050818103600083015261508b8161504f565b9050919050565b6000819050919050565b600060ff82169050919050565b60006150c46150bf6150ba84615092565b614f43565b61509c565b9050919050565b6150d4816150a9565b82525050565b60006020820190506150ef60008301846150cb565b92915050565b7f44506f5374616b696e673a20717565727920666f72206e6f6e6578697374656e60008201527f742063616e646964617465000000000000000000000000000000000000000000602082015250565b6000615151602b83613fd2565b915061515c826150f5565b604082019050919050565b6000602082019050818103600083015261518081615144565b9050919050565b7f5374616b696e674d616e616765723a2075736572206973206e6f74207468652060008201527f63616e6469646174652061646d696e0000000000000000000000000000000000602082015250565b60006151e3602f83613fd2565b91506151ee82615187565b604082019050919050565b60006020820190508181036000830152615212816151d6565b9050919050565b7f5374616b696e674d616e616765723a20717565727920666f722065786365656460008201527f65642076616c696461746f72206172726179206c656e67746800000000000000602082015250565b6000615275603983613fd2565b915061528082615219565b604082019050919050565b600060208201905081810360008301526152a481615268565b9050919050565b7f5374616b696e674d616e616765723a20717565727920666f722065786973746560008201527f642063616e646964617465000000000000000000000000000000000000000000602082015250565b6000615307602b83613fd2565b9150615312826152ab565b604082019050919050565b60006020820190508181036000830152615336816152fa565b9050919050565b7f5374616b696e674d616e616765723a20696e73756666696369656e7420616d6f60008201527f756e740000000000000000000000000000000000000000000000000000000000602082015250565b6000615399602383613fd2565b91506153a48261533d565b604082019050919050565b600060208201905081810360008301526153c88161538c565b9050919050565b7f5374616b696e674d616e616765723a20696e76616c696420747265617375727960008201527f2061646472657373000000000000000000000000000000000000000000000000602082015250565b600061542b602883613fd2565b9150615436826153cf565b604082019050919050565b6000602082019050818103600083015261545a8161541e565b9050919050565b61038082016000820151615478600085018261468a565b50602082015161548b602085018261468a565b50604082015161549e6040850182614699565b5060608201516154b16060850182614483565b5060808201516154c46080850182614483565b5060a08201516154d760a0850182614483565b5060c08201516154ea60c08501826146a8565b5060e08201516154fd60e08501826146b7565b506101008201516155126101008501826146f3565b50505050565b60006103a08201905061552e60008301856141ac565b61553b6020830184615461565b9392505050565b7f5374616b696e674d616e616765723a20696e73756666696369656e7420616d6f60008201527f756e7420746f20756e64656c6567617465000000000000000000000000000000602082015250565b600061559e603183613fd2565b91506155a982615542565b604082019050919050565b600060208201905081810360008301526155cd81615591565b9050919050565b6000815190506155e38161422d565b92915050565b6000602082840312156155ff576155fe614075565b5b600061560d848285016155d4565b91505092915050565b600060408201905061562b6000830185614203565b61563860208301846141ac565b9392505050565b7f5374616b696e674d616e616765723a20696e73756666696369656e742073746160008201527f6b656420616d6f756e7400000000000000000000000000000000000000000000602082015250565b600061569b602a83613fd2565b91506156a68261563f565b604082019050919050565b600060208201905081810360008301526156ca8161568e565b9050919050565b7f5374616b696e674d616e616765723a20696e76616c6964207374616b6564206160008201527f6d6f756e74206c65667400000000000000000000000000000000000000000000602082015250565b600061572d602a83613fd2565b9150615738826156d1565b604082019050919050565b6000602082019050818103600083015261575c81615720565b9050919050565b6000819050919050565b600061577c602084018461412d565b905092915050565b6000602082019050919050565b600061579d83856148ff565b93506157a882615763565b8060005b858110156157e1576157be828461576d565b6157c88882614920565b97506157d383615784565b9250506001810190506157ac565b5085925050509392505050565b60006040820190508181036000830152615809818587615791565b9050818103602083015261581d81846144b7565b9050949350505050565b600060a08201905061583c6000830188614203565b6158496020830187614203565b61585660408301866141ac565b61586360608301856141ac565b61587060808301846141ac565b969550505050505056fea2646970667358221220d43d422832d159e161c917d754bb5dd838887074d4392d03029e2994dfba854464736f6c63430008100033",
}

// DposStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use DposStakingMetaData.ABI instead.
var DposStakingABI = DposStakingMetaData.ABI

// DposStakingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DposStakingMetaData.Bin instead.
var DposStakingBin = DposStakingMetaData.Bin

// DeployDposStaking deploys a new Ethereum contract, binding an instance of DposStaking to it.
func DeployDposStaking(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DposStaking, error) {
	parsed, err := DposStakingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DposStakingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DposStaking{DposStakingCaller: DposStakingCaller{contract: contract}, DposStakingTransactor: DposStakingTransactor{contract: contract}, DposStakingFilterer: DposStakingFilterer{contract: contract}}, nil
}

// DposStaking is an auto generated Go binding around an Ethereum contract.
type DposStaking struct {
	DposStakingCaller     // Read-only binding to the contract
	DposStakingTransactor // Write-only binding to the contract
	DposStakingFilterer   // Log filterer for contract events
}

// DposStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type DposStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DposStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DposStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DposStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DposStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DposStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DposStakingSession struct {
	Contract     *DposStaking      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DposStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DposStakingCallerSession struct {
	Contract *DposStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DposStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DposStakingTransactorSession struct {
	Contract     *DposStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DposStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type DposStakingRaw struct {
	Contract *DposStaking // Generic contract binding to access the raw methods on
}

// DposStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DposStakingCallerRaw struct {
	Contract *DposStakingCaller // Generic read-only contract binding to access the raw methods on
}

// DposStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DposStakingTransactorRaw struct {
	Contract *DposStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDposStaking creates a new instance of DposStaking, bound to a specific deployed contract.
func NewDposStaking(address common.Address, backend bind.ContractBackend) (*DposStaking, error) {
	contract, err := bindDposStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DposStaking{DposStakingCaller: DposStakingCaller{contract: contract}, DposStakingTransactor: DposStakingTransactor{contract: contract}, DposStakingFilterer: DposStakingFilterer{contract: contract}}, nil
}

// NewDposStakingCaller creates a new read-only instance of DposStaking, bound to a specific deployed contract.
func NewDposStakingCaller(address common.Address, caller bind.ContractCaller) (*DposStakingCaller, error) {
	contract, err := bindDposStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DposStakingCaller{contract: contract}, nil
}

// NewDposStakingTransactor creates a new write-only instance of DposStaking, bound to a specific deployed contract.
func NewDposStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*DposStakingTransactor, error) {
	contract, err := bindDposStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DposStakingTransactor{contract: contract}, nil
}

// NewDposStakingFilterer creates a new log filterer instance of DposStaking, bound to a specific deployed contract.
func NewDposStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*DposStakingFilterer, error) {
	contract, err := bindDposStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DposStakingFilterer{contract: contract}, nil
}

// bindDposStaking binds a generic wrapper to an already deployed contract.
func bindDposStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DposStakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DposStaking *DposStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DposStaking.Contract.DposStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DposStaking *DposStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DposStaking.Contract.DposStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DposStaking *DposStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DposStaking.Contract.DposStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DposStaking *DposStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DposStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DposStaking *DposStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DposStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DposStaking *DposStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DposStaking.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address _poolAddr, address _user) view returns(uint256)
func (_DposStaking *DposStakingCaller) BalanceOf(opts *bind.CallOpts, _poolAddr common.Address, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DposStaking.contract.Call(opts, &out, "balanceOf", _poolAddr, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address _poolAddr, address _user) view returns(uint256)
func (_DposStaking *DposStakingSession) BalanceOf(_poolAddr common.Address, _user common.Address) (*big.Int, error) {
	return _DposStaking.Contract.BalanceOf(&_DposStaking.CallOpts, _poolAddr, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address _poolAddr, address _user) view returns(uint256)
func (_DposStaking *DposStakingCallerSession) BalanceOf(_poolAddr common.Address, _user common.Address) (*big.Int, error) {
	return _DposStaking.Contract.BalanceOf(&_DposStaking.CallOpts, _poolAddr, _user)
}

// CommissionRateOf is a free data retrieval call binding the contract method 0x5a783656.
//
// Solidity: function commissionRateOf(address _consensusAddr) view returns(uint256 _rate)
func (_DposStaking *DposStakingCaller) CommissionRateOf(opts *bind.CallOpts, _consensusAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DposStaking.contract.Call(opts, &out, "commissionRateOf", _consensusAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CommissionRateOf is a free data retrieval call binding the contract method 0x5a783656.
//
// Solidity: function commissionRateOf(address _consensusAddr) view returns(uint256 _rate)
func (_DposStaking *DposStakingSession) CommissionRateOf(_consensusAddr common.Address) (*big.Int, error) {
	return _DposStaking.Contract.CommissionRateOf(&_DposStaking.CallOpts, _consensusAddr)
}

// CommissionRateOf is a free data retrieval call binding the contract method 0x5a783656.
//
// Solidity: function commissionRateOf(address _consensusAddr) view returns(uint256 _rate)
func (_DposStaking *DposStakingCallerSession) CommissionRateOf(_consensusAddr common.Address) (*big.Int, error) {
	return _DposStaking.Contract.CommissionRateOf(&_DposStaking.CallOpts, _consensusAddr)
}

// GetCandidateWeights is a free data retrieval call binding the contract method 0xea11bf83.
//
// Solidity: function getCandidateWeights() view returns(address[] _candidates, uint256[] _weights)
func (_DposStaking *DposStakingCaller) GetCandidateWeights(opts *bind.CallOpts) (struct {
	Candidates []common.Address
	Weights    []*big.Int
}, error) {
	var out []interface{}
	err := _DposStaking.contract.Call(opts, &out, "getCandidateWeights")

	outstruct := new(struct {
		Candidates []common.Address
		Weights    []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Candidates = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Weights = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetCandidateWeights is a free data retrieval call binding the contract method 0xea11bf83.
//
// Solidity: function getCandidateWeights() view returns(address[] _candidates, uint256[] _weights)
func (_DposStaking *DposStakingSession) GetCandidateWeights() (struct {
	Candidates []common.Address
	Weights    []*big.Int
}, error) {
	return _DposStaking.Contract.GetCandidateWeights(&_DposStaking.CallOpts)
}

// GetCandidateWeights is a free data retrieval call binding the contract method 0xea11bf83.
//
// Solidity: function getCandidateWeights() view returns(address[] _candidates, uint256[] _weights)
func (_DposStaking *DposStakingCallerSession) GetCandidateWeights() (struct {
	Candidates []common.Address
	Weights    []*big.Int
}, error) {
	return _DposStaking.Contract.GetCandidateWeights(&_DposStaking.CallOpts)
}

// GetClaimableReward is a free data retrieval call binding the contract method 0x21e91dea.
//
// Solidity: function getClaimableReward(address _poolAddr, address _user) view returns(uint256)
func (_DposStaking *DposStakingCaller) GetClaimableReward(opts *bind.CallOpts, _poolAddr common.Address, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DposStaking.contract.Call(opts, &out, "getClaimableReward", _poolAddr, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimableReward is a free data retrieval call binding the contract method 0x21e91dea.
//
// Solidity: function getClaimableReward(address _poolAddr, address _user) view returns(uint256)
func (_DposStaking *DposStakingSession) GetClaimableReward(_poolAddr common.Address, _user common.Address) (*big.Int, error) {
	return _DposStaking.Contract.GetClaimableReward(&_DposStaking.CallOpts, _poolAddr, _user)
}

// GetClaimableReward is a free data retrieval call binding the contract method 0x21e91dea.
//
// Solidity: function getClaimableReward(address _poolAddr, address _user) view returns(uint256)
func (_DposStaking *DposStakingCallerSession) GetClaimableReward(_poolAddr common.Address, _user common.Address) (*big.Int, error) {
	return _DposStaking.Contract.GetClaimableReward(&_DposStaking.CallOpts, _poolAddr, _user)
}

// GetPendingReward is a free data retrieval call binding the contract method 0xfe0f3a13.
//
// Solidity: function getPendingReward(address _poolAddr, address _user) view returns(uint256 _amount)
func (_DposStaking *DposStakingCaller) GetPendingReward(opts *bind.CallOpts, _poolAddr common.Address, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DposStaking.contract.Call(opts, &out, "getPendingReward", _poolAddr, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingReward is a free data retrieval call binding the contract method 0xfe0f3a13.
//
// Solidity: function getPendingReward(address _poolAddr, address _user) view returns(uint256 _amount)
func (_DposStaking *DposStakingSession) GetPendingReward(_poolAddr common.Address, _user common.Address) (*big.Int, error) {
	return _DposStaking.Contract.GetPendingReward(&_DposStaking.CallOpts, _poolAddr, _user)
}

// GetPendingReward is a free data retrieval call binding the contract method 0xfe0f3a13.
//
// Solidity: function getPendingReward(address _poolAddr, address _user) view returns(uint256 _amount)
func (_DposStaking *DposStakingCallerSession) GetPendingReward(_poolAddr common.Address, _user common.Address) (*big.Int, error) {
	return _DposStaking.Contract.GetPendingReward(&_DposStaking.CallOpts, _poolAddr, _user)
}

// GetRewards is a free data retrieval call binding the contract method 0x3d8e846e.
//
// Solidity: function getRewards(address _user, address[] _poolAddrList) view returns(uint256[] _pendings, uint256[] _claimables)
func (_DposStaking *DposStakingCaller) GetRewards(opts *bind.CallOpts, _user common.Address, _poolAddrList []common.Address) (struct {
	Pendings   []*big.Int
	Claimables []*big.Int
}, error) {
	var out []interface{}
	err := _DposStaking.contract.Call(opts, &out, "getRewards", _user, _poolAddrList)

	outstruct := new(struct {
		Pendings   []*big.Int
		Claimables []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pendings = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Claimables = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetRewards is a free data retrieval call binding the contract method 0x3d8e846e.
//
// Solidity: function getRewards(address _user, address[] _poolAddrList) view returns(uint256[] _pendings, uint256[] _claimables)
func (_DposStaking *DposStakingSession) GetRewards(_user common.Address, _poolAddrList []common.Address) (struct {
	Pendings   []*big.Int
	Claimables []*big.Int
}, error) {
	return _DposStaking.Contract.GetRewards(&_DposStaking.CallOpts, _user, _poolAddrList)
}

// GetRewards is a free data retrieval call binding the contract method 0x3d8e846e.
//
// Solidity: function getRewards(address _user, address[] _poolAddrList) view returns(uint256[] _pendings, uint256[] _claimables)
func (_DposStaking *DposStakingCallerSession) GetRewards(_user common.Address, _poolAddrList []common.Address) (struct {
	Pendings   []*big.Int
	Claimables []*big.Int
}, error) {
	return _DposStaking.Contract.GetRewards(&_DposStaking.CallOpts, _user, _poolAddrList)
}

// GetTotalReward is a free data retrieval call binding the contract method 0xe8a22a8a.
//
// Solidity: function getTotalReward(address _poolAddr, address _user) view returns(uint256)
func (_DposStaking *DposStakingCaller) GetTotalReward(opts *bind.CallOpts, _poolAddr common.Address, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DposStaking.contract.Call(opts, &out, "getTotalReward", _poolAddr, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalReward is a free data retrieval call binding the contract method 0xe8a22a8a.
//
// Solidity: function getTotalReward(address _poolAddr, address _user) view returns(uint256)
func (_DposStaking *DposStakingSession) GetTotalReward(_poolAddr common.Address, _user common.Address) (*big.Int, error) {
	return _DposStaking.Contract.GetTotalReward(&_DposStaking.CallOpts, _poolAddr, _user)
}

// GetTotalReward is a free data retrieval call binding the contract method 0xe8a22a8a.
//
// Solidity: function getTotalReward(address _poolAddr, address _user) view returns(uint256)
func (_DposStaking *DposStakingCallerSession) GetTotalReward(_poolAddr common.Address, _user common.Address) (*big.Int, error) {
	return _DposStaking.Contract.GetTotalReward(&_DposStaking.CallOpts, _poolAddr, _user)
}

// GetValidatorCandidates is a free data retrieval call binding the contract method 0xba77b06c.
//
// Solidity: function getValidatorCandidates() view returns((address,address,address,uint256,uint256,uint256,bool,uint8,uint256[20])[])
func (_DposStaking *DposStakingCaller) GetValidatorCandidates(opts *bind.CallOpts) ([]IStakingValidatorCandidate, error) {
	var out []interface{}
	err := _DposStaking.contract.Call(opts, &out, "getValidatorCandidates")

	if err != nil {
		return *new([]IStakingValidatorCandidate), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStakingValidatorCandidate)).(*[]IStakingValidatorCandidate)

	return out0, err

}

// GetValidatorCandidates is a free data retrieval call binding the contract method 0xba77b06c.
//
// Solidity: function getValidatorCandidates() view returns((address,address,address,uint256,uint256,uint256,bool,uint8,uint256[20])[])
func (_DposStaking *DposStakingSession) GetValidatorCandidates() ([]IStakingValidatorCandidate, error) {
	return _DposStaking.Contract.GetValidatorCandidates(&_DposStaking.CallOpts)
}

// GetValidatorCandidates is a free data retrieval call binding the contract method 0xba77b06c.
//
// Solidity: function getValidatorCandidates() view returns((address,address,address,uint256,uint256,uint256,bool,uint8,uint256[20])[])
func (_DposStaking *DposStakingCallerSession) GetValidatorCandidates() ([]IStakingValidatorCandidate, error) {
	return _DposStaking.Contract.GetValidatorCandidates(&_DposStaking.CallOpts)
}

// GovernanceAdminContract is a free data retrieval call binding the contract method 0xea82f784.
//
// Solidity: function governanceAdminContract() view returns(address)
func (_DposStaking *DposStakingCaller) GovernanceAdminContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DposStaking.contract.Call(opts, &out, "governanceAdminContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceAdminContract is a free data retrieval call binding the contract method 0xea82f784.
//
// Solidity: function governanceAdminContract() view returns(address)
func (_DposStaking *DposStakingSession) GovernanceAdminContract() (common.Address, error) {
	return _DposStaking.Contract.GovernanceAdminContract(&_DposStaking.CallOpts)
}

// GovernanceAdminContract is a free data retrieval call binding the contract method 0xea82f784.
//
// Solidity: function governanceAdminContract() view returns(address)
func (_DposStaking *DposStakingCallerSession) GovernanceAdminContract() (common.Address, error) {
	return _DposStaking.Contract.GovernanceAdminContract(&_DposStaking.CallOpts)
}

// MaxValidatorCandidate is a free data retrieval call binding the contract method 0x605239a1.
//
// Solidity: function maxValidatorCandidate() view returns(uint256)
func (_DposStaking *DposStakingCaller) MaxValidatorCandidate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DposStaking.contract.Call(opts, &out, "maxValidatorCandidate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxValidatorCandidate is a free data retrieval call binding the contract method 0x605239a1.
//
// Solidity: function maxValidatorCandidate() view returns(uint256)
func (_DposStaking *DposStakingSession) MaxValidatorCandidate() (*big.Int, error) {
	return _DposStaking.Contract.MaxValidatorCandidate(&_DposStaking.CallOpts)
}

// MaxValidatorCandidate is a free data retrieval call binding the contract method 0x605239a1.
//
// Solidity: function maxValidatorCandidate() view returns(uint256)
func (_DposStaking *DposStakingCallerSession) MaxValidatorCandidate() (*big.Int, error) {
	return _DposStaking.Contract.MaxValidatorCandidate(&_DposStaking.CallOpts)
}

// MinValidatorBalance is a free data retrieval call binding the contract method 0xce99b586.
//
// Solidity: function minValidatorBalance() view returns(uint256)
func (_DposStaking *DposStakingCaller) MinValidatorBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DposStaking.contract.Call(opts, &out, "minValidatorBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinValidatorBalance is a free data retrieval call binding the contract method 0xce99b586.
//
// Solidity: function minValidatorBalance() view returns(uint256)
func (_DposStaking *DposStakingSession) MinValidatorBalance() (*big.Int, error) {
	return _DposStaking.Contract.MinValidatorBalance(&_DposStaking.CallOpts)
}

// MinValidatorBalance is a free data retrieval call binding the contract method 0xce99b586.
//
// Solidity: function minValidatorBalance() view returns(uint256)
func (_DposStaking *DposStakingCallerSession) MinValidatorBalance() (*big.Int, error) {
	return _DposStaking.Contract.MinValidatorBalance(&_DposStaking.CallOpts)
}

// TotalBalance is a free data retrieval call binding the contract method 0x6eacd398.
//
// Solidity: function totalBalance(address _poolAddr) view returns(uint256)
func (_DposStaking *DposStakingCaller) TotalBalance(opts *bind.CallOpts, _poolAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DposStaking.contract.Call(opts, &out, "totalBalance", _poolAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBalance is a free data retrieval call binding the contract method 0x6eacd398.
//
// Solidity: function totalBalance(address _poolAddr) view returns(uint256)
func (_DposStaking *DposStakingSession) TotalBalance(_poolAddr common.Address) (*big.Int, error) {
	return _DposStaking.Contract.TotalBalance(&_DposStaking.CallOpts, _poolAddr)
}

// TotalBalance is a free data retrieval call binding the contract method 0x6eacd398.
//
// Solidity: function totalBalance(address _poolAddr) view returns(uint256)
func (_DposStaking *DposStakingCallerSession) TotalBalance(_poolAddr common.Address) (*big.Int, error) {
	return _DposStaking.Contract.TotalBalance(&_DposStaking.CallOpts, _poolAddr)
}

// TreasuryAddressOf is a free data retrieval call binding the contract method 0x099feccb.
//
// Solidity: function treasuryAddressOf(address _consensusAddr) view returns(address)
func (_DposStaking *DposStakingCaller) TreasuryAddressOf(opts *bind.CallOpts, _consensusAddr common.Address) (common.Address, error) {
	var out []interface{}
	err := _DposStaking.contract.Call(opts, &out, "treasuryAddressOf", _consensusAddr)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryAddressOf is a free data retrieval call binding the contract method 0x099feccb.
//
// Solidity: function treasuryAddressOf(address _consensusAddr) view returns(address)
func (_DposStaking *DposStakingSession) TreasuryAddressOf(_consensusAddr common.Address) (common.Address, error) {
	return _DposStaking.Contract.TreasuryAddressOf(&_DposStaking.CallOpts, _consensusAddr)
}

// TreasuryAddressOf is a free data retrieval call binding the contract method 0x099feccb.
//
// Solidity: function treasuryAddressOf(address _consensusAddr) view returns(address)
func (_DposStaking *DposStakingCallerSession) TreasuryAddressOf(_consensusAddr common.Address) (common.Address, error) {
	return _DposStaking.Contract.TreasuryAddressOf(&_DposStaking.CallOpts, _consensusAddr)
}

// ValidatorCandidates is a free data retrieval call binding the contract method 0x1104b750.
//
// Solidity: function validatorCandidates(uint256 ) view returns(address candidateAdmin, address consensusAddr, address treasuryAddr, uint256 commissionRate, uint256 stakedAmount, uint256 delegatedAmount, bool governing, uint8 state)
func (_DposStaking *DposStakingCaller) ValidatorCandidates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	CandidateAdmin  common.Address
	ConsensusAddr   common.Address
	TreasuryAddr    common.Address
	CommissionRate  *big.Int
	StakedAmount    *big.Int
	DelegatedAmount *big.Int
	Governing       bool
	State           uint8
}, error) {
	var out []interface{}
	err := _DposStaking.contract.Call(opts, &out, "validatorCandidates", arg0)

	outstruct := new(struct {
		CandidateAdmin  common.Address
		ConsensusAddr   common.Address
		TreasuryAddr    common.Address
		CommissionRate  *big.Int
		StakedAmount    *big.Int
		DelegatedAmount *big.Int
		Governing       bool
		State           uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CandidateAdmin = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ConsensusAddr = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TreasuryAddr = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.CommissionRate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.StakedAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.DelegatedAmount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Governing = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.State = *abi.ConvertType(out[7], new(uint8)).(*uint8)

	return *outstruct, err

}

// ValidatorCandidates is a free data retrieval call binding the contract method 0x1104b750.
//
// Solidity: function validatorCandidates(uint256 ) view returns(address candidateAdmin, address consensusAddr, address treasuryAddr, uint256 commissionRate, uint256 stakedAmount, uint256 delegatedAmount, bool governing, uint8 state)
func (_DposStaking *DposStakingSession) ValidatorCandidates(arg0 *big.Int) (struct {
	CandidateAdmin  common.Address
	ConsensusAddr   common.Address
	TreasuryAddr    common.Address
	CommissionRate  *big.Int
	StakedAmount    *big.Int
	DelegatedAmount *big.Int
	Governing       bool
	State           uint8
}, error) {
	return _DposStaking.Contract.ValidatorCandidates(&_DposStaking.CallOpts, arg0)
}

// ValidatorCandidates is a free data retrieval call binding the contract method 0x1104b750.
//
// Solidity: function validatorCandidates(uint256 ) view returns(address candidateAdmin, address consensusAddr, address treasuryAddr, uint256 commissionRate, uint256 stakedAmount, uint256 delegatedAmount, bool governing, uint8 state)
func (_DposStaking *DposStakingCallerSession) ValidatorCandidates(arg0 *big.Int) (struct {
	CandidateAdmin  common.Address
	ConsensusAddr   common.Address
	TreasuryAddr    common.Address
	CommissionRate  *big.Int
	StakedAmount    *big.Int
	DelegatedAmount *big.Int
	Governing       bool
	State           uint8
}, error) {
	return _DposStaking.Contract.ValidatorCandidates(&_DposStaking.CallOpts, arg0)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xdac8ef49.
//
// Solidity: function _claimReward(address _poolAddr, address _user) returns(uint256 _amount)
func (_DposStaking *DposStakingTransactor) ClaimReward(opts *bind.TransactOpts, _poolAddr common.Address, _user common.Address) (*types.Transaction, error) {
	return _DposStaking.contract.Transact(opts, "_claimReward", _poolAddr, _user)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xdac8ef49.
//
// Solidity: function _claimReward(address _poolAddr, address _user) returns(uint256 _amount)
func (_DposStaking *DposStakingSession) ClaimReward(_poolAddr common.Address, _user common.Address) (*types.Transaction, error) {
	return _DposStaking.Contract.ClaimReward(&_DposStaking.TransactOpts, _poolAddr, _user)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xdac8ef49.
//
// Solidity: function _claimReward(address _poolAddr, address _user) returns(uint256 _amount)
func (_DposStaking *DposStakingTransactorSession) ClaimReward(_poolAddr common.Address, _user common.Address) (*types.Transaction, error) {
	return _DposStaking.Contract.ClaimReward(&_DposStaking.TransactOpts, _poolAddr, _user)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xf9f031df.
//
// Solidity: function claimRewards(address[] _consensusAddrList) returns(uint256 _amount)
func (_DposStaking *DposStakingTransactor) ClaimRewards(opts *bind.TransactOpts, _consensusAddrList []common.Address) (*types.Transaction, error) {
	return _DposStaking.contract.Transact(opts, "claimRewards", _consensusAddrList)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xf9f031df.
//
// Solidity: function claimRewards(address[] _consensusAddrList) returns(uint256 _amount)
func (_DposStaking *DposStakingSession) ClaimRewards(_consensusAddrList []common.Address) (*types.Transaction, error) {
	return _DposStaking.Contract.ClaimRewards(&_DposStaking.TransactOpts, _consensusAddrList)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xf9f031df.
//
// Solidity: function claimRewards(address[] _consensusAddrList) returns(uint256 _amount)
func (_DposStaking *DposStakingTransactorSession) ClaimRewards(_consensusAddrList []common.Address) (*types.Transaction, error) {
	return _DposStaking.Contract.ClaimRewards(&_DposStaking.TransactOpts, _consensusAddrList)
}

// DeductStakingAmount is a paid mutator transaction binding the contract method 0xc905bb35.
//
// Solidity: function deductStakingAmount(address _consensusAddr, uint256 _amount) returns()
func (_DposStaking *DposStakingTransactor) DeductStakingAmount(opts *bind.TransactOpts, _consensusAddr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DposStaking.contract.Transact(opts, "deductStakingAmount", _consensusAddr, _amount)
}

// DeductStakingAmount is a paid mutator transaction binding the contract method 0xc905bb35.
//
// Solidity: function deductStakingAmount(address _consensusAddr, uint256 _amount) returns()
func (_DposStaking *DposStakingSession) DeductStakingAmount(_consensusAddr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DposStaking.Contract.DeductStakingAmount(&_DposStaking.TransactOpts, _consensusAddr, _amount)
}

// DeductStakingAmount is a paid mutator transaction binding the contract method 0xc905bb35.
//
// Solidity: function deductStakingAmount(address _consensusAddr, uint256 _amount) returns()
func (_DposStaking *DposStakingTransactorSession) DeductStakingAmount(_consensusAddr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DposStaking.Contract.DeductStakingAmount(&_DposStaking.TransactOpts, _consensusAddr, _amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address _consensusAddr) payable returns()
func (_DposStaking *DposStakingTransactor) Delegate(opts *bind.TransactOpts, _consensusAddr common.Address) (*types.Transaction, error) {
	return _DposStaking.contract.Transact(opts, "delegate", _consensusAddr)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address _consensusAddr) payable returns()
func (_DposStaking *DposStakingSession) Delegate(_consensusAddr common.Address) (*types.Transaction, error) {
	return _DposStaking.Contract.Delegate(&_DposStaking.TransactOpts, _consensusAddr)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address _consensusAddr) payable returns()
func (_DposStaking *DposStakingTransactorSession) Delegate(_consensusAddr common.Address) (*types.Transaction, error) {
	return _DposStaking.Contract.Delegate(&_DposStaking.TransactOpts, _consensusAddr)
}

// DelegateRewards is a paid mutator transaction binding the contract method 0x097e4a9d.
//
// Solidity: function delegateRewards(address[] _consensusAddrList, address _consensusAddrDst) returns(uint256 _amount)
func (_DposStaking *DposStakingTransactor) DelegateRewards(opts *bind.TransactOpts, _consensusAddrList []common.Address, _consensusAddrDst common.Address) (*types.Transaction, error) {
	return _DposStaking.contract.Transact(opts, "delegateRewards", _consensusAddrList, _consensusAddrDst)
}

// DelegateRewards is a paid mutator transaction binding the contract method 0x097e4a9d.
//
// Solidity: function delegateRewards(address[] _consensusAddrList, address _consensusAddrDst) returns(uint256 _amount)
func (_DposStaking *DposStakingSession) DelegateRewards(_consensusAddrList []common.Address, _consensusAddrDst common.Address) (*types.Transaction, error) {
	return _DposStaking.Contract.DelegateRewards(&_DposStaking.TransactOpts, _consensusAddrList, _consensusAddrDst)
}

// DelegateRewards is a paid mutator transaction binding the contract method 0x097e4a9d.
//
// Solidity: function delegateRewards(address[] _consensusAddrList, address _consensusAddrDst) returns(uint256 _amount)
func (_DposStaking *DposStakingTransactorSession) DelegateRewards(_consensusAddrList []common.Address, _consensusAddrDst common.Address) (*types.Transaction, error) {
	return _DposStaking.Contract.DelegateRewards(&_DposStaking.TransactOpts, _consensusAddrList, _consensusAddrDst)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address __validatorContract, address __governanceAdminContract, uint256 __maxValidatorCandidate, uint256 __minValidatorBalance) returns()
func (_DposStaking *DposStakingTransactor) Initialize(opts *bind.TransactOpts, __validatorContract common.Address, __governanceAdminContract common.Address, __maxValidatorCandidate *big.Int, __minValidatorBalance *big.Int) (*types.Transaction, error) {
	return _DposStaking.contract.Transact(opts, "initialize", __validatorContract, __governanceAdminContract, __maxValidatorCandidate, __minValidatorBalance)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address __validatorContract, address __governanceAdminContract, uint256 __maxValidatorCandidate, uint256 __minValidatorBalance) returns()
func (_DposStaking *DposStakingSession) Initialize(__validatorContract common.Address, __governanceAdminContract common.Address, __maxValidatorCandidate *big.Int, __minValidatorBalance *big.Int) (*types.Transaction, error) {
	return _DposStaking.Contract.Initialize(&_DposStaking.TransactOpts, __validatorContract, __governanceAdminContract, __maxValidatorCandidate, __minValidatorBalance)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address __validatorContract, address __governanceAdminContract, uint256 __maxValidatorCandidate, uint256 __minValidatorBalance) returns()
func (_DposStaking *DposStakingTransactorSession) Initialize(__validatorContract common.Address, __governanceAdminContract common.Address, __maxValidatorCandidate *big.Int, __minValidatorBalance *big.Int) (*types.Transaction, error) {
	return _DposStaking.Contract.Initialize(&_DposStaking.TransactOpts, __validatorContract, __governanceAdminContract, __maxValidatorCandidate, __minValidatorBalance)
}

// ProposeValidator is a paid mutator transaction binding the contract method 0x470126b0.
//
// Solidity: function proposeValidator(address _consensusAddr, address _treasuryAddr, uint256 _commissionRate) payable returns(uint256 _candidateIdx)
func (_DposStaking *DposStakingTransactor) ProposeValidator(opts *bind.TransactOpts, _consensusAddr common.Address, _treasuryAddr common.Address, _commissionRate *big.Int) (*types.Transaction, error) {
	return _DposStaking.contract.Transact(opts, "proposeValidator", _consensusAddr, _treasuryAddr, _commissionRate)
}

// ProposeValidator is a paid mutator transaction binding the contract method 0x470126b0.
//
// Solidity: function proposeValidator(address _consensusAddr, address _treasuryAddr, uint256 _commissionRate) payable returns(uint256 _candidateIdx)
func (_DposStaking *DposStakingSession) ProposeValidator(_consensusAddr common.Address, _treasuryAddr common.Address, _commissionRate *big.Int) (*types.Transaction, error) {
	return _DposStaking.Contract.ProposeValidator(&_DposStaking.TransactOpts, _consensusAddr, _treasuryAddr, _commissionRate)
}

// ProposeValidator is a paid mutator transaction binding the contract method 0x470126b0.
//
// Solidity: function proposeValidator(address _consensusAddr, address _treasuryAddr, uint256 _commissionRate) payable returns(uint256 _candidateIdx)
func (_DposStaking *DposStakingTransactorSession) ProposeValidator(_consensusAddr common.Address, _treasuryAddr common.Address, _commissionRate *big.Int) (*types.Transaction, error) {
	return _DposStaking.Contract.ProposeValidator(&_DposStaking.TransactOpts, _consensusAddr, _treasuryAddr, _commissionRate)
}

// RecordReward is a paid mutator transaction binding the contract method 0xb863d710.
//
// Solidity: function recordReward(address _consensusAddr, uint256 _reward) payable returns()
func (_DposStaking *DposStakingTransactor) RecordReward(opts *bind.TransactOpts, _consensusAddr common.Address, _reward *big.Int) (*types.Transaction, error) {
	return _DposStaking.contract.Transact(opts, "recordReward", _consensusAddr, _reward)
}

// RecordReward is a paid mutator transaction binding the contract method 0xb863d710.
//
// Solidity: function recordReward(address _consensusAddr, uint256 _reward) payable returns()
func (_DposStaking *DposStakingSession) RecordReward(_consensusAddr common.Address, _reward *big.Int) (*types.Transaction, error) {
	return _DposStaking.Contract.RecordReward(&_DposStaking.TransactOpts, _consensusAddr, _reward)
}

// RecordReward is a paid mutator transaction binding the contract method 0xb863d710.
//
// Solidity: function recordReward(address _consensusAddr, uint256 _reward) payable returns()
func (_DposStaking *DposStakingTransactorSession) RecordReward(_consensusAddr common.Address, _reward *big.Int) (*types.Transaction, error) {
	return _DposStaking.Contract.RecordReward(&_DposStaking.TransactOpts, _consensusAddr, _reward)
}

// Redelegate is a paid mutator transaction binding the contract method 0x6bd8f804.
//
// Solidity: function redelegate(address _consensusAddrSrc, address _consensusAddrDst, uint256 _amount) returns()
func (_DposStaking *DposStakingTransactor) Redelegate(opts *bind.TransactOpts, _consensusAddrSrc common.Address, _consensusAddrDst common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DposStaking.contract.Transact(opts, "redelegate", _consensusAddrSrc, _consensusAddrDst, _amount)
}

// Redelegate is a paid mutator transaction binding the contract method 0x6bd8f804.
//
// Solidity: function redelegate(address _consensusAddrSrc, address _consensusAddrDst, uint256 _amount) returns()
func (_DposStaking *DposStakingSession) Redelegate(_consensusAddrSrc common.Address, _consensusAddrDst common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DposStaking.Contract.Redelegate(&_DposStaking.TransactOpts, _consensusAddrSrc, _consensusAddrDst, _amount)
}

// Redelegate is a paid mutator transaction binding the contract method 0x6bd8f804.
//
// Solidity: function redelegate(address _consensusAddrSrc, address _consensusAddrDst, uint256 _amount) returns()
func (_DposStaking *DposStakingTransactorSession) Redelegate(_consensusAddrSrc common.Address, _consensusAddrDst common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DposStaking.Contract.Redelegate(&_DposStaking.TransactOpts, _consensusAddrSrc, _consensusAddrDst, _amount)
}

// Renounce is a paid mutator transaction binding the contract method 0x1f76a7af.
//
// Solidity: function renounce(address consensusAddr) returns()
func (_DposStaking *DposStakingTransactor) Renounce(opts *bind.TransactOpts, consensusAddr common.Address) (*types.Transaction, error) {
	return _DposStaking.contract.Transact(opts, "renounce", consensusAddr)
}

// Renounce is a paid mutator transaction binding the contract method 0x1f76a7af.
//
// Solidity: function renounce(address consensusAddr) returns()
func (_DposStaking *DposStakingSession) Renounce(consensusAddr common.Address) (*types.Transaction, error) {
	return _DposStaking.Contract.Renounce(&_DposStaking.TransactOpts, consensusAddr)
}

// Renounce is a paid mutator transaction binding the contract method 0x1f76a7af.
//
// Solidity: function renounce(address consensusAddr) returns()
func (_DposStaking *DposStakingTransactorSession) Renounce(consensusAddr common.Address) (*types.Transaction, error) {
	return _DposStaking.Contract.Renounce(&_DposStaking.TransactOpts, consensusAddr)
}

// SetMaxValidatorCandidate is a paid mutator transaction binding the contract method 0x4f2a693f.
//
// Solidity: function setMaxValidatorCandidate(uint256 _threshold) returns()
func (_DposStaking *DposStakingTransactor) SetMaxValidatorCandidate(opts *bind.TransactOpts, _threshold *big.Int) (*types.Transaction, error) {
	return _DposStaking.contract.Transact(opts, "setMaxValidatorCandidate", _threshold)
}

// SetMaxValidatorCandidate is a paid mutator transaction binding the contract method 0x4f2a693f.
//
// Solidity: function setMaxValidatorCandidate(uint256 _threshold) returns()
func (_DposStaking *DposStakingSession) SetMaxValidatorCandidate(_threshold *big.Int) (*types.Transaction, error) {
	return _DposStaking.Contract.SetMaxValidatorCandidate(&_DposStaking.TransactOpts, _threshold)
}

// SetMaxValidatorCandidate is a paid mutator transaction binding the contract method 0x4f2a693f.
//
// Solidity: function setMaxValidatorCandidate(uint256 _threshold) returns()
func (_DposStaking *DposStakingTransactorSession) SetMaxValidatorCandidate(_threshold *big.Int) (*types.Transaction, error) {
	return _DposStaking.Contract.SetMaxValidatorCandidate(&_DposStaking.TransactOpts, _threshold)
}

// SetMinValidatorBalance is a paid mutator transaction binding the contract method 0xfe7732f4.
//
// Solidity: function setMinValidatorBalance(uint256 _threshold) returns()
func (_DposStaking *DposStakingTransactor) SetMinValidatorBalance(opts *bind.TransactOpts, _threshold *big.Int) (*types.Transaction, error) {
	return _DposStaking.contract.Transact(opts, "setMinValidatorBalance", _threshold)
}

// SetMinValidatorBalance is a paid mutator transaction binding the contract method 0xfe7732f4.
//
// Solidity: function setMinValidatorBalance(uint256 _threshold) returns()
func (_DposStaking *DposStakingSession) SetMinValidatorBalance(_threshold *big.Int) (*types.Transaction, error) {
	return _DposStaking.Contract.SetMinValidatorBalance(&_DposStaking.TransactOpts, _threshold)
}

// SetMinValidatorBalance is a paid mutator transaction binding the contract method 0xfe7732f4.
//
// Solidity: function setMinValidatorBalance(uint256 _threshold) returns()
func (_DposStaking *DposStakingTransactorSession) SetMinValidatorBalance(_threshold *big.Int) (*types.Transaction, error) {
	return _DposStaking.Contract.SetMinValidatorBalance(&_DposStaking.TransactOpts, _threshold)
}

// SettleRewardPools is a paid mutator transaction binding the contract method 0xd45e6273.
//
// Solidity: function settleRewardPools(address[] _consensusAddrs) returns()
func (_DposStaking *DposStakingTransactor) SettleRewardPools(opts *bind.TransactOpts, _consensusAddrs []common.Address) (*types.Transaction, error) {
	return _DposStaking.contract.Transact(opts, "settleRewardPools", _consensusAddrs)
}

// SettleRewardPools is a paid mutator transaction binding the contract method 0xd45e6273.
//
// Solidity: function settleRewardPools(address[] _consensusAddrs) returns()
func (_DposStaking *DposStakingSession) SettleRewardPools(_consensusAddrs []common.Address) (*types.Transaction, error) {
	return _DposStaking.Contract.SettleRewardPools(&_DposStaking.TransactOpts, _consensusAddrs)
}

// SettleRewardPools is a paid mutator transaction binding the contract method 0xd45e6273.
//
// Solidity: function settleRewardPools(address[] _consensusAddrs) returns()
func (_DposStaking *DposStakingTransactorSession) SettleRewardPools(_consensusAddrs []common.Address) (*types.Transaction, error) {
	return _DposStaking.Contract.SettleRewardPools(&_DposStaking.TransactOpts, _consensusAddrs)
}

// SinkPendingReward is a paid mutator transaction binding the contract method 0xa5885f1d.
//
// Solidity: function sinkPendingReward(address _consensusAddr) returns()
func (_DposStaking *DposStakingTransactor) SinkPendingReward(opts *bind.TransactOpts, _consensusAddr common.Address) (*types.Transaction, error) {
	return _DposStaking.contract.Transact(opts, "sinkPendingReward", _consensusAddr)
}

// SinkPendingReward is a paid mutator transaction binding the contract method 0xa5885f1d.
//
// Solidity: function sinkPendingReward(address _consensusAddr) returns()
func (_DposStaking *DposStakingSession) SinkPendingReward(_consensusAddr common.Address) (*types.Transaction, error) {
	return _DposStaking.Contract.SinkPendingReward(&_DposStaking.TransactOpts, _consensusAddr)
}

// SinkPendingReward is a paid mutator transaction binding the contract method 0xa5885f1d.
//
// Solidity: function sinkPendingReward(address _consensusAddr) returns()
func (_DposStaking *DposStakingTransactorSession) SinkPendingReward(_consensusAddr common.Address) (*types.Transaction, error) {
	return _DposStaking.Contract.SinkPendingReward(&_DposStaking.TransactOpts, _consensusAddr)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address _consensusAddr) payable returns()
func (_DposStaking *DposStakingTransactor) Stake(opts *bind.TransactOpts, _consensusAddr common.Address) (*types.Transaction, error) {
	return _DposStaking.contract.Transact(opts, "stake", _consensusAddr)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address _consensusAddr) payable returns()
func (_DposStaking *DposStakingSession) Stake(_consensusAddr common.Address) (*types.Transaction, error) {
	return _DposStaking.Contract.Stake(&_DposStaking.TransactOpts, _consensusAddr)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address _consensusAddr) payable returns()
func (_DposStaking *DposStakingTransactorSession) Stake(_consensusAddr common.Address) (*types.Transaction, error) {
	return _DposStaking.Contract.Stake(&_DposStaking.TransactOpts, _consensusAddr)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address _consensusAddr, uint256 _amount) returns()
func (_DposStaking *DposStakingTransactor) Undelegate(opts *bind.TransactOpts, _consensusAddr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DposStaking.contract.Transact(opts, "undelegate", _consensusAddr, _amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address _consensusAddr, uint256 _amount) returns()
func (_DposStaking *DposStakingSession) Undelegate(_consensusAddr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DposStaking.Contract.Undelegate(&_DposStaking.TransactOpts, _consensusAddr, _amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address _consensusAddr, uint256 _amount) returns()
func (_DposStaking *DposStakingTransactorSession) Undelegate(_consensusAddr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DposStaking.Contract.Undelegate(&_DposStaking.TransactOpts, _consensusAddr, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address _consensusAddr, uint256 _amount) returns()
func (_DposStaking *DposStakingTransactor) Unstake(opts *bind.TransactOpts, _consensusAddr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DposStaking.contract.Transact(opts, "unstake", _consensusAddr, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address _consensusAddr, uint256 _amount) returns()
func (_DposStaking *DposStakingSession) Unstake(_consensusAddr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DposStaking.Contract.Unstake(&_DposStaking.TransactOpts, _consensusAddr, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address _consensusAddr, uint256 _amount) returns()
func (_DposStaking *DposStakingTransactorSession) Unstake(_consensusAddr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DposStaking.Contract.Unstake(&_DposStaking.TransactOpts, _consensusAddr, _amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_DposStaking *DposStakingTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _DposStaking.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_DposStaking *DposStakingSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _DposStaking.Contract.Fallback(&_DposStaking.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_DposStaking *DposStakingTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _DposStaking.Contract.Fallback(&_DposStaking.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DposStaking *DposStakingTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DposStaking.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DposStaking *DposStakingSession) Receive() (*types.Transaction, error) {
	return _DposStaking.Contract.Receive(&_DposStaking.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DposStaking *DposStakingTransactorSession) Receive() (*types.Transaction, error) {
	return _DposStaking.Contract.Receive(&_DposStaking.TransactOpts)
}

// DposStakingDelegatedIterator is returned from FilterDelegated and is used to iterate over the raw logs and unpacked data for Delegated events raised by the DposStaking contract.
type DposStakingDelegatedIterator struct {
	Event *DposStakingDelegated // Event containing the contract specifics and raw log

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
func (it *DposStakingDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposStakingDelegated)
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
		it.Event = new(DposStakingDelegated)
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
func (it *DposStakingDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposStakingDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposStakingDelegated represents a Delegated event raised by the DposStaking contract.
type DposStakingDelegated struct {
	Delegator common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegated is a free log retrieval operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed delegator, address indexed validator, uint256 amount)
func (_DposStaking *DposStakingFilterer) FilterDelegated(opts *bind.FilterOpts, delegator []common.Address, validator []common.Address) (*DposStakingDelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposStaking.contract.FilterLogs(opts, "Delegated", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &DposStakingDelegatedIterator{contract: _DposStaking.contract, event: "Delegated", logs: logs, sub: sub}, nil
}

// WatchDelegated is a free log subscription operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed delegator, address indexed validator, uint256 amount)
func (_DposStaking *DposStakingFilterer) WatchDelegated(opts *bind.WatchOpts, sink chan<- *DposStakingDelegated, delegator []common.Address, validator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposStaking.contract.WatchLogs(opts, "Delegated", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposStakingDelegated)
				if err := _DposStaking.contract.UnpackLog(event, "Delegated", log); err != nil {
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

// ParseDelegated is a log parse operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed delegator, address indexed validator, uint256 amount)
func (_DposStaking *DposStakingFilterer) ParseDelegated(log types.Log) (*DposStakingDelegated, error) {
	event := new(DposStakingDelegated)
	if err := _DposStaking.contract.UnpackLog(event, "Delegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposStakingGovernanceAdminContractUpdatedIterator is returned from FilterGovernanceAdminContractUpdated and is used to iterate over the raw logs and unpacked data for GovernanceAdminContractUpdated events raised by the DposStaking contract.
type DposStakingGovernanceAdminContractUpdatedIterator struct {
	Event *DposStakingGovernanceAdminContractUpdated // Event containing the contract specifics and raw log

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
func (it *DposStakingGovernanceAdminContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposStakingGovernanceAdminContractUpdated)
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
		it.Event = new(DposStakingGovernanceAdminContractUpdated)
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
func (it *DposStakingGovernanceAdminContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposStakingGovernanceAdminContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposStakingGovernanceAdminContractUpdated represents a GovernanceAdminContractUpdated event raised by the DposStaking contract.
type DposStakingGovernanceAdminContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterGovernanceAdminContractUpdated is a free log retrieval operation binding the contract event 0x2374c4e57a97e792cc22ffb7ddfc5f526897b2e22738b569da0e8f5dd17e1f34.
//
// Solidity: event GovernanceAdminContractUpdated(address arg0)
func (_DposStaking *DposStakingFilterer) FilterGovernanceAdminContractUpdated(opts *bind.FilterOpts) (*DposStakingGovernanceAdminContractUpdatedIterator, error) {

	logs, sub, err := _DposStaking.contract.FilterLogs(opts, "GovernanceAdminContractUpdated")
	if err != nil {
		return nil, err
	}
	return &DposStakingGovernanceAdminContractUpdatedIterator{contract: _DposStaking.contract, event: "GovernanceAdminContractUpdated", logs: logs, sub: sub}, nil
}

// WatchGovernanceAdminContractUpdated is a free log subscription operation binding the contract event 0x2374c4e57a97e792cc22ffb7ddfc5f526897b2e22738b569da0e8f5dd17e1f34.
//
// Solidity: event GovernanceAdminContractUpdated(address arg0)
func (_DposStaking *DposStakingFilterer) WatchGovernanceAdminContractUpdated(opts *bind.WatchOpts, sink chan<- *DposStakingGovernanceAdminContractUpdated) (event.Subscription, error) {

	logs, sub, err := _DposStaking.contract.WatchLogs(opts, "GovernanceAdminContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposStakingGovernanceAdminContractUpdated)
				if err := _DposStaking.contract.UnpackLog(event, "GovernanceAdminContractUpdated", log); err != nil {
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

// ParseGovernanceAdminContractUpdated is a log parse operation binding the contract event 0x2374c4e57a97e792cc22ffb7ddfc5f526897b2e22738b569da0e8f5dd17e1f34.
//
// Solidity: event GovernanceAdminContractUpdated(address arg0)
func (_DposStaking *DposStakingFilterer) ParseGovernanceAdminContractUpdated(log types.Log) (*DposStakingGovernanceAdminContractUpdated, error) {
	event := new(DposStakingGovernanceAdminContractUpdated)
	if err := _DposStaking.contract.UnpackLog(event, "GovernanceAdminContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposStakingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DposStaking contract.
type DposStakingInitializedIterator struct {
	Event *DposStakingInitialized // Event containing the contract specifics and raw log

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
func (it *DposStakingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposStakingInitialized)
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
		it.Event = new(DposStakingInitialized)
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
func (it *DposStakingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposStakingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposStakingInitialized represents a Initialized event raised by the DposStaking contract.
type DposStakingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DposStaking *DposStakingFilterer) FilterInitialized(opts *bind.FilterOpts) (*DposStakingInitializedIterator, error) {

	logs, sub, err := _DposStaking.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DposStakingInitializedIterator{contract: _DposStaking.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DposStaking *DposStakingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DposStakingInitialized) (event.Subscription, error) {

	logs, sub, err := _DposStaking.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposStakingInitialized)
				if err := _DposStaking.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_DposStaking *DposStakingFilterer) ParseInitialized(log types.Log) (*DposStakingInitialized, error) {
	event := new(DposStakingInitialized)
	if err := _DposStaking.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposStakingMaxValidatorCandidateUpdatedIterator is returned from FilterMaxValidatorCandidateUpdated and is used to iterate over the raw logs and unpacked data for MaxValidatorCandidateUpdated events raised by the DposStaking contract.
type DposStakingMaxValidatorCandidateUpdatedIterator struct {
	Event *DposStakingMaxValidatorCandidateUpdated // Event containing the contract specifics and raw log

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
func (it *DposStakingMaxValidatorCandidateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposStakingMaxValidatorCandidateUpdated)
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
		it.Event = new(DposStakingMaxValidatorCandidateUpdated)
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
func (it *DposStakingMaxValidatorCandidateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposStakingMaxValidatorCandidateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposStakingMaxValidatorCandidateUpdated represents a MaxValidatorCandidateUpdated event raised by the DposStaking contract.
type DposStakingMaxValidatorCandidateUpdated struct {
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMaxValidatorCandidateUpdated is a free log retrieval operation binding the contract event 0x82d5dc32d1b741512ad09c32404d7e7921e8934c6222343d95f55f7a2b9b2ab4.
//
// Solidity: event MaxValidatorCandidateUpdated(uint256 threshold)
func (_DposStaking *DposStakingFilterer) FilterMaxValidatorCandidateUpdated(opts *bind.FilterOpts) (*DposStakingMaxValidatorCandidateUpdatedIterator, error) {

	logs, sub, err := _DposStaking.contract.FilterLogs(opts, "MaxValidatorCandidateUpdated")
	if err != nil {
		return nil, err
	}
	return &DposStakingMaxValidatorCandidateUpdatedIterator{contract: _DposStaking.contract, event: "MaxValidatorCandidateUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxValidatorCandidateUpdated is a free log subscription operation binding the contract event 0x82d5dc32d1b741512ad09c32404d7e7921e8934c6222343d95f55f7a2b9b2ab4.
//
// Solidity: event MaxValidatorCandidateUpdated(uint256 threshold)
func (_DposStaking *DposStakingFilterer) WatchMaxValidatorCandidateUpdated(opts *bind.WatchOpts, sink chan<- *DposStakingMaxValidatorCandidateUpdated) (event.Subscription, error) {

	logs, sub, err := _DposStaking.contract.WatchLogs(opts, "MaxValidatorCandidateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposStakingMaxValidatorCandidateUpdated)
				if err := _DposStaking.contract.UnpackLog(event, "MaxValidatorCandidateUpdated", log); err != nil {
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
func (_DposStaking *DposStakingFilterer) ParseMaxValidatorCandidateUpdated(log types.Log) (*DposStakingMaxValidatorCandidateUpdated, error) {
	event := new(DposStakingMaxValidatorCandidateUpdated)
	if err := _DposStaking.contract.UnpackLog(event, "MaxValidatorCandidateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposStakingMinValidatorBalanceUpdatedIterator is returned from FilterMinValidatorBalanceUpdated and is used to iterate over the raw logs and unpacked data for MinValidatorBalanceUpdated events raised by the DposStaking contract.
type DposStakingMinValidatorBalanceUpdatedIterator struct {
	Event *DposStakingMinValidatorBalanceUpdated // Event containing the contract specifics and raw log

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
func (it *DposStakingMinValidatorBalanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposStakingMinValidatorBalanceUpdated)
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
		it.Event = new(DposStakingMinValidatorBalanceUpdated)
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
func (it *DposStakingMinValidatorBalanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposStakingMinValidatorBalanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposStakingMinValidatorBalanceUpdated represents a MinValidatorBalanceUpdated event raised by the DposStaking contract.
type DposStakingMinValidatorBalanceUpdated struct {
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinValidatorBalanceUpdated is a free log retrieval operation binding the contract event 0x8364d8933fc5aad4a5ae33139fea4cbc485c4f1931e822d06cd3b873c13a07e6.
//
// Solidity: event MinValidatorBalanceUpdated(uint256 threshold)
func (_DposStaking *DposStakingFilterer) FilterMinValidatorBalanceUpdated(opts *bind.FilterOpts) (*DposStakingMinValidatorBalanceUpdatedIterator, error) {

	logs, sub, err := _DposStaking.contract.FilterLogs(opts, "MinValidatorBalanceUpdated")
	if err != nil {
		return nil, err
	}
	return &DposStakingMinValidatorBalanceUpdatedIterator{contract: _DposStaking.contract, event: "MinValidatorBalanceUpdated", logs: logs, sub: sub}, nil
}

// WatchMinValidatorBalanceUpdated is a free log subscription operation binding the contract event 0x8364d8933fc5aad4a5ae33139fea4cbc485c4f1931e822d06cd3b873c13a07e6.
//
// Solidity: event MinValidatorBalanceUpdated(uint256 threshold)
func (_DposStaking *DposStakingFilterer) WatchMinValidatorBalanceUpdated(opts *bind.WatchOpts, sink chan<- *DposStakingMinValidatorBalanceUpdated) (event.Subscription, error) {

	logs, sub, err := _DposStaking.contract.WatchLogs(opts, "MinValidatorBalanceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposStakingMinValidatorBalanceUpdated)
				if err := _DposStaking.contract.UnpackLog(event, "MinValidatorBalanceUpdated", log); err != nil {
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

// ParseMinValidatorBalanceUpdated is a log parse operation binding the contract event 0x8364d8933fc5aad4a5ae33139fea4cbc485c4f1931e822d06cd3b873c13a07e6.
//
// Solidity: event MinValidatorBalanceUpdated(uint256 threshold)
func (_DposStaking *DposStakingFilterer) ParseMinValidatorBalanceUpdated(log types.Log) (*DposStakingMinValidatorBalanceUpdated, error) {
	event := new(DposStakingMinValidatorBalanceUpdated)
	if err := _DposStaking.contract.UnpackLog(event, "MinValidatorBalanceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposStakingPendingPoolUpdatedIterator is returned from FilterPendingPoolUpdated and is used to iterate over the raw logs and unpacked data for PendingPoolUpdated events raised by the DposStaking contract.
type DposStakingPendingPoolUpdatedIterator struct {
	Event *DposStakingPendingPoolUpdated // Event containing the contract specifics and raw log

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
func (it *DposStakingPendingPoolUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposStakingPendingPoolUpdated)
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
		it.Event = new(DposStakingPendingPoolUpdated)
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
func (it *DposStakingPendingPoolUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposStakingPendingPoolUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposStakingPendingPoolUpdated represents a PendingPoolUpdated event raised by the DposStaking contract.
type DposStakingPendingPoolUpdated struct {
	PoolAddress    common.Address
	AccumulatedRps *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPendingPoolUpdated is a free log retrieval operation binding the contract event 0x5791918d066f160d5cfcc4ccd4df1f7f4f0a0f48b486a2ef33e238c1369f2224.
//
// Solidity: event PendingPoolUpdated(address poolAddress, uint256 accumulatedRps)
func (_DposStaking *DposStakingFilterer) FilterPendingPoolUpdated(opts *bind.FilterOpts) (*DposStakingPendingPoolUpdatedIterator, error) {

	logs, sub, err := _DposStaking.contract.FilterLogs(opts, "PendingPoolUpdated")
	if err != nil {
		return nil, err
	}
	return &DposStakingPendingPoolUpdatedIterator{contract: _DposStaking.contract, event: "PendingPoolUpdated", logs: logs, sub: sub}, nil
}

// WatchPendingPoolUpdated is a free log subscription operation binding the contract event 0x5791918d066f160d5cfcc4ccd4df1f7f4f0a0f48b486a2ef33e238c1369f2224.
//
// Solidity: event PendingPoolUpdated(address poolAddress, uint256 accumulatedRps)
func (_DposStaking *DposStakingFilterer) WatchPendingPoolUpdated(opts *bind.WatchOpts, sink chan<- *DposStakingPendingPoolUpdated) (event.Subscription, error) {

	logs, sub, err := _DposStaking.contract.WatchLogs(opts, "PendingPoolUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposStakingPendingPoolUpdated)
				if err := _DposStaking.contract.UnpackLog(event, "PendingPoolUpdated", log); err != nil {
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

// ParsePendingPoolUpdated is a log parse operation binding the contract event 0x5791918d066f160d5cfcc4ccd4df1f7f4f0a0f48b486a2ef33e238c1369f2224.
//
// Solidity: event PendingPoolUpdated(address poolAddress, uint256 accumulatedRps)
func (_DposStaking *DposStakingFilterer) ParsePendingPoolUpdated(log types.Log) (*DposStakingPendingPoolUpdated, error) {
	event := new(DposStakingPendingPoolUpdated)
	if err := _DposStaking.contract.UnpackLog(event, "PendingPoolUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposStakingPendingRewardUpdatedIterator is returned from FilterPendingRewardUpdated and is used to iterate over the raw logs and unpacked data for PendingRewardUpdated events raised by the DposStaking contract.
type DposStakingPendingRewardUpdatedIterator struct {
	Event *DposStakingPendingRewardUpdated // Event containing the contract specifics and raw log

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
func (it *DposStakingPendingRewardUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposStakingPendingRewardUpdated)
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
		it.Event = new(DposStakingPendingRewardUpdated)
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
func (it *DposStakingPendingRewardUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposStakingPendingRewardUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposStakingPendingRewardUpdated represents a PendingRewardUpdated event raised by the DposStaking contract.
type DposStakingPendingRewardUpdated struct {
	PoolAddress common.Address
	User        common.Address
	Debited     *big.Int
	Credited    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPendingRewardUpdated is a free log retrieval operation binding the contract event 0xd6faab37181ad960bdaaf04de2913a71085ce9224f4575265ec261d57419d477.
//
// Solidity: event PendingRewardUpdated(address poolAddress, address user, uint256 debited, uint256 credited)
func (_DposStaking *DposStakingFilterer) FilterPendingRewardUpdated(opts *bind.FilterOpts) (*DposStakingPendingRewardUpdatedIterator, error) {

	logs, sub, err := _DposStaking.contract.FilterLogs(opts, "PendingRewardUpdated")
	if err != nil {
		return nil, err
	}
	return &DposStakingPendingRewardUpdatedIterator{contract: _DposStaking.contract, event: "PendingRewardUpdated", logs: logs, sub: sub}, nil
}

// WatchPendingRewardUpdated is a free log subscription operation binding the contract event 0xd6faab37181ad960bdaaf04de2913a71085ce9224f4575265ec261d57419d477.
//
// Solidity: event PendingRewardUpdated(address poolAddress, address user, uint256 debited, uint256 credited)
func (_DposStaking *DposStakingFilterer) WatchPendingRewardUpdated(opts *bind.WatchOpts, sink chan<- *DposStakingPendingRewardUpdated) (event.Subscription, error) {

	logs, sub, err := _DposStaking.contract.WatchLogs(opts, "PendingRewardUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposStakingPendingRewardUpdated)
				if err := _DposStaking.contract.UnpackLog(event, "PendingRewardUpdated", log); err != nil {
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

// ParsePendingRewardUpdated is a log parse operation binding the contract event 0xd6faab37181ad960bdaaf04de2913a71085ce9224f4575265ec261d57419d477.
//
// Solidity: event PendingRewardUpdated(address poolAddress, address user, uint256 debited, uint256 credited)
func (_DposStaking *DposStakingFilterer) ParsePendingRewardUpdated(log types.Log) (*DposStakingPendingRewardUpdated, error) {
	event := new(DposStakingPendingRewardUpdated)
	if err := _DposStaking.contract.UnpackLog(event, "PendingRewardUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposStakingRewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the DposStaking contract.
type DposStakingRewardClaimedIterator struct {
	Event *DposStakingRewardClaimed // Event containing the contract specifics and raw log

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
func (it *DposStakingRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposStakingRewardClaimed)
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
		it.Event = new(DposStakingRewardClaimed)
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
func (it *DposStakingRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposStakingRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposStakingRewardClaimed represents a RewardClaimed event raised by the DposStaking contract.
type DposStakingRewardClaimed struct {
	PoolAddress common.Address
	User        common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimed is a free log retrieval operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address poolAddress, address user, uint256 amount)
func (_DposStaking *DposStakingFilterer) FilterRewardClaimed(opts *bind.FilterOpts) (*DposStakingRewardClaimedIterator, error) {

	logs, sub, err := _DposStaking.contract.FilterLogs(opts, "RewardClaimed")
	if err != nil {
		return nil, err
	}
	return &DposStakingRewardClaimedIterator{contract: _DposStaking.contract, event: "RewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address poolAddress, address user, uint256 amount)
func (_DposStaking *DposStakingFilterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *DposStakingRewardClaimed) (event.Subscription, error) {

	logs, sub, err := _DposStaking.contract.WatchLogs(opts, "RewardClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposStakingRewardClaimed)
				if err := _DposStaking.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
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

// ParseRewardClaimed is a log parse operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address poolAddress, address user, uint256 amount)
func (_DposStaking *DposStakingFilterer) ParseRewardClaimed(log types.Log) (*DposStakingRewardClaimed, error) {
	event := new(DposStakingRewardClaimed)
	if err := _DposStaking.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposStakingSettledPoolsUpdatedIterator is returned from FilterSettledPoolsUpdated and is used to iterate over the raw logs and unpacked data for SettledPoolsUpdated events raised by the DposStaking contract.
type DposStakingSettledPoolsUpdatedIterator struct {
	Event *DposStakingSettledPoolsUpdated // Event containing the contract specifics and raw log

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
func (it *DposStakingSettledPoolsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposStakingSettledPoolsUpdated)
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
		it.Event = new(DposStakingSettledPoolsUpdated)
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
func (it *DposStakingSettledPoolsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposStakingSettledPoolsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposStakingSettledPoolsUpdated represents a SettledPoolsUpdated event raised by the DposStaking contract.
type DposStakingSettledPoolsUpdated struct {
	PoolAddress    []common.Address
	AccumulatedRps []*big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSettledPoolsUpdated is a free log retrieval operation binding the contract event 0xed065952215f1a6f0ff936480b7ae27e204860bd6e35aca6a678cce83dc0c490.
//
// Solidity: event SettledPoolsUpdated(address[] poolAddress, uint256[] accumulatedRps)
func (_DposStaking *DposStakingFilterer) FilterSettledPoolsUpdated(opts *bind.FilterOpts) (*DposStakingSettledPoolsUpdatedIterator, error) {

	logs, sub, err := _DposStaking.contract.FilterLogs(opts, "SettledPoolsUpdated")
	if err != nil {
		return nil, err
	}
	return &DposStakingSettledPoolsUpdatedIterator{contract: _DposStaking.contract, event: "SettledPoolsUpdated", logs: logs, sub: sub}, nil
}

// WatchSettledPoolsUpdated is a free log subscription operation binding the contract event 0xed065952215f1a6f0ff936480b7ae27e204860bd6e35aca6a678cce83dc0c490.
//
// Solidity: event SettledPoolsUpdated(address[] poolAddress, uint256[] accumulatedRps)
func (_DposStaking *DposStakingFilterer) WatchSettledPoolsUpdated(opts *bind.WatchOpts, sink chan<- *DposStakingSettledPoolsUpdated) (event.Subscription, error) {

	logs, sub, err := _DposStaking.contract.WatchLogs(opts, "SettledPoolsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposStakingSettledPoolsUpdated)
				if err := _DposStaking.contract.UnpackLog(event, "SettledPoolsUpdated", log); err != nil {
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

// ParseSettledPoolsUpdated is a log parse operation binding the contract event 0xed065952215f1a6f0ff936480b7ae27e204860bd6e35aca6a678cce83dc0c490.
//
// Solidity: event SettledPoolsUpdated(address[] poolAddress, uint256[] accumulatedRps)
func (_DposStaking *DposStakingFilterer) ParseSettledPoolsUpdated(log types.Log) (*DposStakingSettledPoolsUpdated, error) {
	event := new(DposStakingSettledPoolsUpdated)
	if err := _DposStaking.contract.UnpackLog(event, "SettledPoolsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposStakingSettledRewardUpdatedIterator is returned from FilterSettledRewardUpdated and is used to iterate over the raw logs and unpacked data for SettledRewardUpdated events raised by the DposStaking contract.
type DposStakingSettledRewardUpdatedIterator struct {
	Event *DposStakingSettledRewardUpdated // Event containing the contract specifics and raw log

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
func (it *DposStakingSettledRewardUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposStakingSettledRewardUpdated)
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
		it.Event = new(DposStakingSettledRewardUpdated)
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
func (it *DposStakingSettledRewardUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposStakingSettledRewardUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposStakingSettledRewardUpdated represents a SettledRewardUpdated event raised by the DposStaking contract.
type DposStakingSettledRewardUpdated struct {
	PoolAddress    common.Address
	User           common.Address
	Balance        *big.Int
	Debited        *big.Int
	AccumulatedRps *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSettledRewardUpdated is a free log retrieval operation binding the contract event 0x0b74380c3e212861270a8670337fc8147462c92a81a228d1dc50d3caa0ab62c5.
//
// Solidity: event SettledRewardUpdated(address poolAddress, address user, uint256 balance, uint256 debited, uint256 accumulatedRps)
func (_DposStaking *DposStakingFilterer) FilterSettledRewardUpdated(opts *bind.FilterOpts) (*DposStakingSettledRewardUpdatedIterator, error) {

	logs, sub, err := _DposStaking.contract.FilterLogs(opts, "SettledRewardUpdated")
	if err != nil {
		return nil, err
	}
	return &DposStakingSettledRewardUpdatedIterator{contract: _DposStaking.contract, event: "SettledRewardUpdated", logs: logs, sub: sub}, nil
}

// WatchSettledRewardUpdated is a free log subscription operation binding the contract event 0x0b74380c3e212861270a8670337fc8147462c92a81a228d1dc50d3caa0ab62c5.
//
// Solidity: event SettledRewardUpdated(address poolAddress, address user, uint256 balance, uint256 debited, uint256 accumulatedRps)
func (_DposStaking *DposStakingFilterer) WatchSettledRewardUpdated(opts *bind.WatchOpts, sink chan<- *DposStakingSettledRewardUpdated) (event.Subscription, error) {

	logs, sub, err := _DposStaking.contract.WatchLogs(opts, "SettledRewardUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposStakingSettledRewardUpdated)
				if err := _DposStaking.contract.UnpackLog(event, "SettledRewardUpdated", log); err != nil {
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

// ParseSettledRewardUpdated is a log parse operation binding the contract event 0x0b74380c3e212861270a8670337fc8147462c92a81a228d1dc50d3caa0ab62c5.
//
// Solidity: event SettledRewardUpdated(address poolAddress, address user, uint256 balance, uint256 debited, uint256 accumulatedRps)
func (_DposStaking *DposStakingFilterer) ParseSettledRewardUpdated(log types.Log) (*DposStakingSettledRewardUpdated, error) {
	event := new(DposStakingSettledRewardUpdated)
	if err := _DposStaking.contract.UnpackLog(event, "SettledRewardUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposStakingStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the DposStaking contract.
type DposStakingStakedIterator struct {
	Event *DposStakingStaked // Event containing the contract specifics and raw log

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
func (it *DposStakingStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposStakingStaked)
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
		it.Event = new(DposStakingStaked)
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
func (it *DposStakingStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposStakingStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposStakingStaked represents a Staked event raised by the DposStaking contract.
type DposStakingStaked struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed validator, uint256 amount)
func (_DposStaking *DposStakingFilterer) FilterStaked(opts *bind.FilterOpts, validator []common.Address) (*DposStakingStakedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposStaking.contract.FilterLogs(opts, "Staked", validatorRule)
	if err != nil {
		return nil, err
	}
	return &DposStakingStakedIterator{contract: _DposStaking.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed validator, uint256 amount)
func (_DposStaking *DposStakingFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *DposStakingStaked, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposStaking.contract.WatchLogs(opts, "Staked", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposStakingStaked)
				if err := _DposStaking.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed validator, uint256 amount)
func (_DposStaking *DposStakingFilterer) ParseStaked(log types.Log) (*DposStakingStaked, error) {
	event := new(DposStakingStaked)
	if err := _DposStaking.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposStakingUndelegatedIterator is returned from FilterUndelegated and is used to iterate over the raw logs and unpacked data for Undelegated events raised by the DposStaking contract.
type DposStakingUndelegatedIterator struct {
	Event *DposStakingUndelegated // Event containing the contract specifics and raw log

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
func (it *DposStakingUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposStakingUndelegated)
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
		it.Event = new(DposStakingUndelegated)
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
func (it *DposStakingUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposStakingUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposStakingUndelegated represents a Undelegated event raised by the DposStaking contract.
type DposStakingUndelegated struct {
	Delegator common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUndelegated is a free log retrieval operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed delegator, address indexed validator, uint256 amount)
func (_DposStaking *DposStakingFilterer) FilterUndelegated(opts *bind.FilterOpts, delegator []common.Address, validator []common.Address) (*DposStakingUndelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposStaking.contract.FilterLogs(opts, "Undelegated", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &DposStakingUndelegatedIterator{contract: _DposStaking.contract, event: "Undelegated", logs: logs, sub: sub}, nil
}

// WatchUndelegated is a free log subscription operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed delegator, address indexed validator, uint256 amount)
func (_DposStaking *DposStakingFilterer) WatchUndelegated(opts *bind.WatchOpts, sink chan<- *DposStakingUndelegated, delegator []common.Address, validator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposStaking.contract.WatchLogs(opts, "Undelegated", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposStakingUndelegated)
				if err := _DposStaking.contract.UnpackLog(event, "Undelegated", log); err != nil {
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

// ParseUndelegated is a log parse operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed delegator, address indexed validator, uint256 amount)
func (_DposStaking *DposStakingFilterer) ParseUndelegated(log types.Log) (*DposStakingUndelegated, error) {
	event := new(DposStakingUndelegated)
	if err := _DposStaking.contract.UnpackLog(event, "Undelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposStakingUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the DposStaking contract.
type DposStakingUnstakedIterator struct {
	Event *DposStakingUnstaked // Event containing the contract specifics and raw log

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
func (it *DposStakingUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposStakingUnstaked)
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
		it.Event = new(DposStakingUnstaked)
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
func (it *DposStakingUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposStakingUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposStakingUnstaked represents a Unstaked event raised by the DposStaking contract.
type DposStakingUnstaked struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed validator, uint256 amount)
func (_DposStaking *DposStakingFilterer) FilterUnstaked(opts *bind.FilterOpts, validator []common.Address) (*DposStakingUnstakedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposStaking.contract.FilterLogs(opts, "Unstaked", validatorRule)
	if err != nil {
		return nil, err
	}
	return &DposStakingUnstakedIterator{contract: _DposStaking.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed validator, uint256 amount)
func (_DposStaking *DposStakingFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *DposStakingUnstaked, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposStaking.contract.WatchLogs(opts, "Unstaked", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposStakingUnstaked)
				if err := _DposStaking.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed validator, uint256 amount)
func (_DposStaking *DposStakingFilterer) ParseUnstaked(log types.Log) (*DposStakingUnstaked, error) {
	event := new(DposStakingUnstaked)
	if err := _DposStaking.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposStakingValidatorContractUpdatedIterator is returned from FilterValidatorContractUpdated and is used to iterate over the raw logs and unpacked data for ValidatorContractUpdated events raised by the DposStaking contract.
type DposStakingValidatorContractUpdatedIterator struct {
	Event *DposStakingValidatorContractUpdated // Event containing the contract specifics and raw log

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
func (it *DposStakingValidatorContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposStakingValidatorContractUpdated)
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
		it.Event = new(DposStakingValidatorContractUpdated)
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
func (it *DposStakingValidatorContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposStakingValidatorContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposStakingValidatorContractUpdated represents a ValidatorContractUpdated event raised by the DposStaking contract.
type DposStakingValidatorContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterValidatorContractUpdated is a free log retrieval operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_DposStaking *DposStakingFilterer) FilterValidatorContractUpdated(opts *bind.FilterOpts) (*DposStakingValidatorContractUpdatedIterator, error) {

	logs, sub, err := _DposStaking.contract.FilterLogs(opts, "ValidatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return &DposStakingValidatorContractUpdatedIterator{contract: _DposStaking.contract, event: "ValidatorContractUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorContractUpdated is a free log subscription operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_DposStaking *DposStakingFilterer) WatchValidatorContractUpdated(opts *bind.WatchOpts, sink chan<- *DposStakingValidatorContractUpdated) (event.Subscription, error) {

	logs, sub, err := _DposStaking.contract.WatchLogs(opts, "ValidatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposStakingValidatorContractUpdated)
				if err := _DposStaking.contract.UnpackLog(event, "ValidatorContractUpdated", log); err != nil {
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
func (_DposStaking *DposStakingFilterer) ParseValidatorContractUpdated(log types.Log) (*DposStakingValidatorContractUpdated, error) {
	event := new(DposStakingValidatorContractUpdated)
	if err := _DposStaking.contract.UnpackLog(event, "ValidatorContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposStakingValidatorProposedIterator is returned from FilterValidatorProposed and is used to iterate over the raw logs and unpacked data for ValidatorProposed events raised by the DposStaking contract.
type DposStakingValidatorProposedIterator struct {
	Event *DposStakingValidatorProposed // Event containing the contract specifics and raw log

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
func (it *DposStakingValidatorProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposStakingValidatorProposed)
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
		it.Event = new(DposStakingValidatorProposed)
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
func (it *DposStakingValidatorProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposStakingValidatorProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposStakingValidatorProposed represents a ValidatorProposed event raised by the DposStaking contract.
type DposStakingValidatorProposed struct {
	ConsensusAddr common.Address
	CandidateIdx  common.Address
	Amount        *big.Int
	Info          IStakingValidatorCandidate
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterValidatorProposed is a free log retrieval operation binding the contract event 0x48a4728ce2a28ba5a3060eafbd15c8f72563d2dda3e74c24cbf14bc6c68e03db.
//
// Solidity: event ValidatorProposed(address indexed consensusAddr, address indexed candidateIdx, uint256 amount, (address,address,address,uint256,uint256,uint256,bool,uint8,uint256[20]) _info)
func (_DposStaking *DposStakingFilterer) FilterValidatorProposed(opts *bind.FilterOpts, consensusAddr []common.Address, candidateIdx []common.Address) (*DposStakingValidatorProposedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var candidateIdxRule []interface{}
	for _, candidateIdxItem := range candidateIdx {
		candidateIdxRule = append(candidateIdxRule, candidateIdxItem)
	}

	logs, sub, err := _DposStaking.contract.FilterLogs(opts, "ValidatorProposed", consensusAddrRule, candidateIdxRule)
	if err != nil {
		return nil, err
	}
	return &DposStakingValidatorProposedIterator{contract: _DposStaking.contract, event: "ValidatorProposed", logs: logs, sub: sub}, nil
}

// WatchValidatorProposed is a free log subscription operation binding the contract event 0x48a4728ce2a28ba5a3060eafbd15c8f72563d2dda3e74c24cbf14bc6c68e03db.
//
// Solidity: event ValidatorProposed(address indexed consensusAddr, address indexed candidateIdx, uint256 amount, (address,address,address,uint256,uint256,uint256,bool,uint8,uint256[20]) _info)
func (_DposStaking *DposStakingFilterer) WatchValidatorProposed(opts *bind.WatchOpts, sink chan<- *DposStakingValidatorProposed, consensusAddr []common.Address, candidateIdx []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var candidateIdxRule []interface{}
	for _, candidateIdxItem := range candidateIdx {
		candidateIdxRule = append(candidateIdxRule, candidateIdxItem)
	}

	logs, sub, err := _DposStaking.contract.WatchLogs(opts, "ValidatorProposed", consensusAddrRule, candidateIdxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposStakingValidatorProposed)
				if err := _DposStaking.contract.UnpackLog(event, "ValidatorProposed", log); err != nil {
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

// ParseValidatorProposed is a log parse operation binding the contract event 0x48a4728ce2a28ba5a3060eafbd15c8f72563d2dda3e74c24cbf14bc6c68e03db.
//
// Solidity: event ValidatorProposed(address indexed consensusAddr, address indexed candidateIdx, uint256 amount, (address,address,address,uint256,uint256,uint256,bool,uint8,uint256[20]) _info)
func (_DposStaking *DposStakingFilterer) ParseValidatorProposed(log types.Log) (*DposStakingValidatorProposed, error) {
	event := new(DposStakingValidatorProposed)
	if err := _DposStaking.contract.UnpackLog(event, "ValidatorProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposStakingValidatorRenounceFinalizedIterator is returned from FilterValidatorRenounceFinalized and is used to iterate over the raw logs and unpacked data for ValidatorRenounceFinalized events raised by the DposStaking contract.
type DposStakingValidatorRenounceFinalizedIterator struct {
	Event *DposStakingValidatorRenounceFinalized // Event containing the contract specifics and raw log

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
func (it *DposStakingValidatorRenounceFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposStakingValidatorRenounceFinalized)
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
		it.Event = new(DposStakingValidatorRenounceFinalized)
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
func (it *DposStakingValidatorRenounceFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposStakingValidatorRenounceFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposStakingValidatorRenounceFinalized represents a ValidatorRenounceFinalized event raised by the DposStaking contract.
type DposStakingValidatorRenounceFinalized struct {
	ConsensusAddr common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterValidatorRenounceFinalized is a free log retrieval operation binding the contract event 0xab42748c6c0d9912e327fa961ef2a00d422f5eda666d9fb9625fb50def39c9aa.
//
// Solidity: event ValidatorRenounceFinalized(address indexed consensusAddr, uint256 amount)
func (_DposStaking *DposStakingFilterer) FilterValidatorRenounceFinalized(opts *bind.FilterOpts, consensusAddr []common.Address) (*DposStakingValidatorRenounceFinalizedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _DposStaking.contract.FilterLogs(opts, "ValidatorRenounceFinalized", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return &DposStakingValidatorRenounceFinalizedIterator{contract: _DposStaking.contract, event: "ValidatorRenounceFinalized", logs: logs, sub: sub}, nil
}

// WatchValidatorRenounceFinalized is a free log subscription operation binding the contract event 0xab42748c6c0d9912e327fa961ef2a00d422f5eda666d9fb9625fb50def39c9aa.
//
// Solidity: event ValidatorRenounceFinalized(address indexed consensusAddr, uint256 amount)
func (_DposStaking *DposStakingFilterer) WatchValidatorRenounceFinalized(opts *bind.WatchOpts, sink chan<- *DposStakingValidatorRenounceFinalized, consensusAddr []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _DposStaking.contract.WatchLogs(opts, "ValidatorRenounceFinalized", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposStakingValidatorRenounceFinalized)
				if err := _DposStaking.contract.UnpackLog(event, "ValidatorRenounceFinalized", log); err != nil {
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

// ParseValidatorRenounceFinalized is a log parse operation binding the contract event 0xab42748c6c0d9912e327fa961ef2a00d422f5eda666d9fb9625fb50def39c9aa.
//
// Solidity: event ValidatorRenounceFinalized(address indexed consensusAddr, uint256 amount)
func (_DposStaking *DposStakingFilterer) ParseValidatorRenounceFinalized(log types.Log) (*DposStakingValidatorRenounceFinalized, error) {
	event := new(DposStakingValidatorRenounceFinalized)
	if err := _DposStaking.contract.UnpackLog(event, "ValidatorRenounceFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposStakingValidatorRenounceRequestedIterator is returned from FilterValidatorRenounceRequested and is used to iterate over the raw logs and unpacked data for ValidatorRenounceRequested events raised by the DposStaking contract.
type DposStakingValidatorRenounceRequestedIterator struct {
	Event *DposStakingValidatorRenounceRequested // Event containing the contract specifics and raw log

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
func (it *DposStakingValidatorRenounceRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposStakingValidatorRenounceRequested)
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
		it.Event = new(DposStakingValidatorRenounceRequested)
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
func (it *DposStakingValidatorRenounceRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposStakingValidatorRenounceRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposStakingValidatorRenounceRequested represents a ValidatorRenounceRequested event raised by the DposStaking contract.
type DposStakingValidatorRenounceRequested struct {
	ConsensusAddr common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterValidatorRenounceRequested is a free log retrieval operation binding the contract event 0x7846cc8ba1aebec179d3e03448547b26ec8c875f510695bf1a494c837140655f.
//
// Solidity: event ValidatorRenounceRequested(address indexed consensusAddr, uint256 amount)
func (_DposStaking *DposStakingFilterer) FilterValidatorRenounceRequested(opts *bind.FilterOpts, consensusAddr []common.Address) (*DposStakingValidatorRenounceRequestedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _DposStaking.contract.FilterLogs(opts, "ValidatorRenounceRequested", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return &DposStakingValidatorRenounceRequestedIterator{contract: _DposStaking.contract, event: "ValidatorRenounceRequested", logs: logs, sub: sub}, nil
}

// WatchValidatorRenounceRequested is a free log subscription operation binding the contract event 0x7846cc8ba1aebec179d3e03448547b26ec8c875f510695bf1a494c837140655f.
//
// Solidity: event ValidatorRenounceRequested(address indexed consensusAddr, uint256 amount)
func (_DposStaking *DposStakingFilterer) WatchValidatorRenounceRequested(opts *bind.WatchOpts, sink chan<- *DposStakingValidatorRenounceRequested, consensusAddr []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _DposStaking.contract.WatchLogs(opts, "ValidatorRenounceRequested", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposStakingValidatorRenounceRequested)
				if err := _DposStaking.contract.UnpackLog(event, "ValidatorRenounceRequested", log); err != nil {
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

// ParseValidatorRenounceRequested is a log parse operation binding the contract event 0x7846cc8ba1aebec179d3e03448547b26ec8c875f510695bf1a494c837140655f.
//
// Solidity: event ValidatorRenounceRequested(address indexed consensusAddr, uint256 amount)
func (_DposStaking *DposStakingFilterer) ParseValidatorRenounceRequested(log types.Log) (*DposStakingValidatorRenounceRequested, error) {
	event := new(DposStakingValidatorRenounceRequested)
	if err := _DposStaking.contract.UnpackLog(event, "ValidatorRenounceRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
