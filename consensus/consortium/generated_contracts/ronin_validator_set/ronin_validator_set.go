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
	Admin              common.Address
	ConsensusAddr      common.Address
	TreasuryAddr       common.Address
	BridgeOperatorAddr common.Address
	CommissionRate     *big.Int
	RevokedPeriod      *big.Int
	ExtraData          []byte
}

// RoninValidatorSetMetaData contains all meta data concerning the RoninValidatorSet contract.
var RoninValidatorSetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"consensusAddrs\",\"type\":\"address[]\"}],\"name\":\"BlockProducerSetUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"coinbaseAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"name\":\"BlockRewardRewardDeprecated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"coinbaseAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"submittedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bonusAmount\",\"type\":\"uint256\"}],\"name\":\"BlockRewardSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeOperator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeOperatorRewardDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeOperator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractBalance\",\"type\":\"uint256\"}],\"name\":\"BridgeOperatorRewardDistributionFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"bridgeOperators\",\"type\":\"address[]\"}],\"name\":\"BridgeOperatorSetUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"BridgeTrackingContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"treasuryAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridgeOperator\",\"type\":\"address\"}],\"name\":\"CandidateGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"revokedPeriod\",\"type\":\"uint256\"}],\"name\":\"CandidateRevokedPeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"consensusAddrs\",\"type\":\"address[]\"}],\"name\":\"CandidatesRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"MaintenanceContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MaxPrioritizedValidatorNumberUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"MaxValidatorCandidateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MaxValidatorNumberUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MiningRewardDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractBalance\",\"type\":\"uint256\"}],\"name\":\"MiningRewardDistributionFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"NumberOfBlocksInEpochUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"RoninTrustedOrganizationContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"SlashIndicatorContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"StakingContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakingRewardDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractBalance\",\"type\":\"uint256\"}],\"name\":\"StakingRewardDistributionFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"StakingVestingContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorLiberated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"jailedUntil\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deductedStakingAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"blockProducerRewardDeprecated\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"bridgeOperatorRewardDeprecated\",\"type\":\"bool\"}],\"name\":\"ValidatorPunished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"consensusAddrs\",\"type\":\"address[]\"}],\"name\":\"ValidatorSetUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"periodNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"periodEnding\",\"type\":\"bool\"}],\"name\":\"WrappedUpEpoch\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPeriod\",\"type\":\"uint256\"}],\"name\":\"_isPeriodEnding\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"}],\"name\":\"bailOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeTrackingContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addrList\",\"type\":\"address[]\"}],\"name\":\"bulkJailed\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"_result\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentPeriodStartAtBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"epochEndingAt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"epochOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockProducers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_result\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_bridgeOperatorList\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getCandidateInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"treasuryAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridgeOperatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revokedPeriod\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structICandidateManager.ValidatorCandidate\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCandidateInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"treasuryAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridgeOperatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revokedPeriod\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structICandidateManager.ValidatorCandidate[]\",\"name\":\"_list\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastUpdatedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorCandidates\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_validatorList\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_treasuryAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeOperatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_commissionRate\",\"type\":\"uint256\"}],\"name\":\"grantValidatorCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__slashIndicatorContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__stakingContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__stakingVestingContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__maintenanceContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__roninTrustedOrganizationContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__bridgeTrackingContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"__maxValidatorNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"__maxValidatorCandidate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"__maxPrioritizedValidatorNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"__numberOfBlocksInEpoch\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isBlockProducer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeOperatorAddr\",\"type\":\"address\"}],\"name\":\"isBridgeOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"isCandidateAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPeriodEnding\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isValidatorCandidate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"jailed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNum\",\"type\":\"uint256\"}],\"name\":\"jailedAtBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"jailedTimeLeft\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isJailed_\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"blockLeft_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochLeft_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNum\",\"type\":\"uint256\"}],\"name\":\"jailedTimeLeftAtBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isJailed_\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"blockLeft_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochLeft_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maintenanceContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPrioritizedValidatorNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_maximumPrioritizedValidatorNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxValidatorCandidate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxValidatorNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_maximumValidatorNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_blockProducers\",\"type\":\"address[]\"}],\"name\":\"miningRewardDeprecated\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"_result\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_blockProducers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"miningRewardDeprecatedAtPeriod\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"_result\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberOfBlocksInEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_numberOfBlocks\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"precompilePickValidatorSetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"precompileSortValidatorsAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"}],\"name\":\"requestRevokeCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roninTrustedOrganizationContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setBridgeTrackingContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setMaintenanceContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"setMaxValidatorCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_max\",\"type\":\"uint256\"}],\"name\":\"setMaxValidatorNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"setNumberOfBlocksInEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setRoninTrustedOrganizationContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setSlashIndicatorContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setStakingContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setStakingVestingContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newJailedUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slashAmount\",\"type\":\"uint256\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashIndicatorContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingVestingContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitBlockReward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBlockProducers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBridgeOperators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrapUpEpoch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RoninValidatorSetABI is the input ABI used to generate the binding from.
// Deprecated: Use RoninValidatorSetMetaData.ABI instead.
var RoninValidatorSetABI = RoninValidatorSetMetaData.ABI

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

// IsPeriodEnding is a free data retrieval call binding the contract method 0x9b8c334b.
//
// Solidity: function _isPeriodEnding(uint256 _newPeriod) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCaller) IsPeriodEnding(opts *bind.CallOpts, _newPeriod *big.Int) (bool, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "_isPeriodEnding", _newPeriod)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPeriodEnding is a free data retrieval call binding the contract method 0x9b8c334b.
//
// Solidity: function _isPeriodEnding(uint256 _newPeriod) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetSession) IsPeriodEnding(_newPeriod *big.Int) (bool, error) {
	return _RoninValidatorSet.Contract.IsPeriodEnding(&_RoninValidatorSet.CallOpts, _newPeriod)
}

// IsPeriodEnding is a free data retrieval call binding the contract method 0x9b8c334b.
//
// Solidity: function _isPeriodEnding(uint256 _newPeriod) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) IsPeriodEnding(_newPeriod *big.Int) (bool, error) {
	return _RoninValidatorSet.Contract.IsPeriodEnding(&_RoninValidatorSet.CallOpts, _newPeriod)
}

// BridgeTrackingContract is a free data retrieval call binding the contract method 0x4493421e.
//
// Solidity: function bridgeTrackingContract() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetCaller) BridgeTrackingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "bridgeTrackingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeTrackingContract is a free data retrieval call binding the contract method 0x4493421e.
//
// Solidity: function bridgeTrackingContract() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetSession) BridgeTrackingContract() (common.Address, error) {
	return _RoninValidatorSet.Contract.BridgeTrackingContract(&_RoninValidatorSet.CallOpts)
}

// BridgeTrackingContract is a free data retrieval call binding the contract method 0x4493421e.
//
// Solidity: function bridgeTrackingContract() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) BridgeTrackingContract() (common.Address, error) {
	return _RoninValidatorSet.Contract.BridgeTrackingContract(&_RoninValidatorSet.CallOpts)
}

// BulkJailed is a free data retrieval call binding the contract method 0x428483c3.
//
// Solidity: function bulkJailed(address[] _addrList) view returns(bool[] _result)
func (_RoninValidatorSet *RoninValidatorSetCaller) BulkJailed(opts *bind.CallOpts, _addrList []common.Address) ([]bool, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "bulkJailed", _addrList)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// BulkJailed is a free data retrieval call binding the contract method 0x428483c3.
//
// Solidity: function bulkJailed(address[] _addrList) view returns(bool[] _result)
func (_RoninValidatorSet *RoninValidatorSetSession) BulkJailed(_addrList []common.Address) ([]bool, error) {
	return _RoninValidatorSet.Contract.BulkJailed(&_RoninValidatorSet.CallOpts, _addrList)
}

// BulkJailed is a free data retrieval call binding the contract method 0x428483c3.
//
// Solidity: function bulkJailed(address[] _addrList) view returns(bool[] _result)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) BulkJailed(_addrList []common.Address) ([]bool, error) {
	return _RoninValidatorSet.Contract.BulkJailed(&_RoninValidatorSet.CallOpts, _addrList)
}

// CurrentPeriod is a free data retrieval call binding the contract method 0x06040618.
//
// Solidity: function currentPeriod() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCaller) CurrentPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "currentPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPeriod is a free data retrieval call binding the contract method 0x06040618.
//
// Solidity: function currentPeriod() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetSession) CurrentPeriod() (*big.Int, error) {
	return _RoninValidatorSet.Contract.CurrentPeriod(&_RoninValidatorSet.CallOpts)
}

// CurrentPeriod is a free data retrieval call binding the contract method 0x06040618.
//
// Solidity: function currentPeriod() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) CurrentPeriod() (*big.Int, error) {
	return _RoninValidatorSet.Contract.CurrentPeriod(&_RoninValidatorSet.CallOpts)
}

// CurrentPeriodStartAtBlock is a free data retrieval call binding the contract method 0x297a8fca.
//
// Solidity: function currentPeriodStartAtBlock() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCaller) CurrentPeriodStartAtBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "currentPeriodStartAtBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPeriodStartAtBlock is a free data retrieval call binding the contract method 0x297a8fca.
//
// Solidity: function currentPeriodStartAtBlock() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetSession) CurrentPeriodStartAtBlock() (*big.Int, error) {
	return _RoninValidatorSet.Contract.CurrentPeriodStartAtBlock(&_RoninValidatorSet.CallOpts)
}

// CurrentPeriodStartAtBlock is a free data retrieval call binding the contract method 0x297a8fca.
//
// Solidity: function currentPeriodStartAtBlock() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) CurrentPeriodStartAtBlock() (*big.Int, error) {
	return _RoninValidatorSet.Contract.CurrentPeriodStartAtBlock(&_RoninValidatorSet.CallOpts)
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

// GetBlockProducers is a free data retrieval call binding the contract method 0x49096d26.
//
// Solidity: function getBlockProducers() view returns(address[] _result)
func (_RoninValidatorSet *RoninValidatorSetCaller) GetBlockProducers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "getBlockProducers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBlockProducers is a free data retrieval call binding the contract method 0x49096d26.
//
// Solidity: function getBlockProducers() view returns(address[] _result)
func (_RoninValidatorSet *RoninValidatorSetSession) GetBlockProducers() ([]common.Address, error) {
	return _RoninValidatorSet.Contract.GetBlockProducers(&_RoninValidatorSet.CallOpts)
}

// GetBlockProducers is a free data retrieval call binding the contract method 0x49096d26.
//
// Solidity: function getBlockProducers() view returns(address[] _result)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) GetBlockProducers() ([]common.Address, error) {
	return _RoninValidatorSet.Contract.GetBlockProducers(&_RoninValidatorSet.CallOpts)
}

// GetBridgeOperators is a free data retrieval call binding the contract method 0x9b19dbfd.
//
// Solidity: function getBridgeOperators() view returns(address[] _bridgeOperatorList)
func (_RoninValidatorSet *RoninValidatorSetCaller) GetBridgeOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "getBridgeOperators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBridgeOperators is a free data retrieval call binding the contract method 0x9b19dbfd.
//
// Solidity: function getBridgeOperators() view returns(address[] _bridgeOperatorList)
func (_RoninValidatorSet *RoninValidatorSetSession) GetBridgeOperators() ([]common.Address, error) {
	return _RoninValidatorSet.Contract.GetBridgeOperators(&_RoninValidatorSet.CallOpts)
}

// GetBridgeOperators is a free data retrieval call binding the contract method 0x9b19dbfd.
//
// Solidity: function getBridgeOperators() view returns(address[] _bridgeOperatorList)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) GetBridgeOperators() ([]common.Address, error) {
	return _RoninValidatorSet.Contract.GetBridgeOperators(&_RoninValidatorSet.CallOpts)
}

// GetCandidateInfo is a free data retrieval call binding the contract method 0x28bde1e1.
//
// Solidity: function getCandidateInfo(address _candidate) view returns((address,address,address,address,uint256,uint256,bytes))
func (_RoninValidatorSet *RoninValidatorSetCaller) GetCandidateInfo(opts *bind.CallOpts, _candidate common.Address) (ICandidateManagerValidatorCandidate, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "getCandidateInfo", _candidate)

	if err != nil {
		return *new(ICandidateManagerValidatorCandidate), err
	}

	out0 := *abi.ConvertType(out[0], new(ICandidateManagerValidatorCandidate)).(*ICandidateManagerValidatorCandidate)

	return out0, err

}

// GetCandidateInfo is a free data retrieval call binding the contract method 0x28bde1e1.
//
// Solidity: function getCandidateInfo(address _candidate) view returns((address,address,address,address,uint256,uint256,bytes))
func (_RoninValidatorSet *RoninValidatorSetSession) GetCandidateInfo(_candidate common.Address) (ICandidateManagerValidatorCandidate, error) {
	return _RoninValidatorSet.Contract.GetCandidateInfo(&_RoninValidatorSet.CallOpts, _candidate)
}

// GetCandidateInfo is a free data retrieval call binding the contract method 0x28bde1e1.
//
// Solidity: function getCandidateInfo(address _candidate) view returns((address,address,address,address,uint256,uint256,bytes))
func (_RoninValidatorSet *RoninValidatorSetCallerSession) GetCandidateInfo(_candidate common.Address) (ICandidateManagerValidatorCandidate, error) {
	return _RoninValidatorSet.Contract.GetCandidateInfo(&_RoninValidatorSet.CallOpts, _candidate)
}

// GetCandidateInfos is a free data retrieval call binding the contract method 0x5248184a.
//
// Solidity: function getCandidateInfos() view returns((address,address,address,address,uint256,uint256,bytes)[] _list)
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
// Solidity: function getCandidateInfos() view returns((address,address,address,address,uint256,uint256,bytes)[] _list)
func (_RoninValidatorSet *RoninValidatorSetSession) GetCandidateInfos() ([]ICandidateManagerValidatorCandidate, error) {
	return _RoninValidatorSet.Contract.GetCandidateInfos(&_RoninValidatorSet.CallOpts)
}

// GetCandidateInfos is a free data retrieval call binding the contract method 0x5248184a.
//
// Solidity: function getCandidateInfos() view returns((address,address,address,address,uint256,uint256,bytes)[] _list)
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

// IsBlockProducer is a free data retrieval call binding the contract method 0x65244ece.
//
// Solidity: function isBlockProducer(address _addr) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCaller) IsBlockProducer(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "isBlockProducer", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlockProducer is a free data retrieval call binding the contract method 0x65244ece.
//
// Solidity: function isBlockProducer(address _addr) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetSession) IsBlockProducer(_addr common.Address) (bool, error) {
	return _RoninValidatorSet.Contract.IsBlockProducer(&_RoninValidatorSet.CallOpts, _addr)
}

// IsBlockProducer is a free data retrieval call binding the contract method 0x65244ece.
//
// Solidity: function isBlockProducer(address _addr) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) IsBlockProducer(_addr common.Address) (bool, error) {
	return _RoninValidatorSet.Contract.IsBlockProducer(&_RoninValidatorSet.CallOpts, _addr)
}

// IsBridgeOperator is a free data retrieval call binding the contract method 0xb405aaf2.
//
// Solidity: function isBridgeOperator(address _bridgeOperatorAddr) view returns(bool _result)
func (_RoninValidatorSet *RoninValidatorSetCaller) IsBridgeOperator(opts *bind.CallOpts, _bridgeOperatorAddr common.Address) (bool, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "isBridgeOperator", _bridgeOperatorAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBridgeOperator is a free data retrieval call binding the contract method 0xb405aaf2.
//
// Solidity: function isBridgeOperator(address _bridgeOperatorAddr) view returns(bool _result)
func (_RoninValidatorSet *RoninValidatorSetSession) IsBridgeOperator(_bridgeOperatorAddr common.Address) (bool, error) {
	return _RoninValidatorSet.Contract.IsBridgeOperator(&_RoninValidatorSet.CallOpts, _bridgeOperatorAddr)
}

// IsBridgeOperator is a free data retrieval call binding the contract method 0xb405aaf2.
//
// Solidity: function isBridgeOperator(address _bridgeOperatorAddr) view returns(bool _result)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) IsBridgeOperator(_bridgeOperatorAddr common.Address) (bool, error) {
	return _RoninValidatorSet.Contract.IsBridgeOperator(&_RoninValidatorSet.CallOpts, _bridgeOperatorAddr)
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

// IsCurrentPeriodEnding is a free data retrieval call binding the contract method 0x217f35c2.
//
// Solidity: function isPeriodEnding() view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCaller) IsCurrentPeriodEnding(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "isPeriodEnding")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCurrentPeriodEnding is a free data retrieval call binding the contract method 0x217f35c2.
//
// Solidity: function isPeriodEnding() view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetSession) IsCurrentPeriodEnding() (bool, error) {
	return _RoninValidatorSet.Contract.IsCurrentPeriodEnding(&_RoninValidatorSet.CallOpts)
}

// IsCurrentPeriodEnding is a free data retrieval call binding the contract method 0x217f35c2.
//
// Solidity: function isPeriodEnding() view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) IsCurrentPeriodEnding() (bool, error) {
	return _RoninValidatorSet.Contract.IsCurrentPeriodEnding(&_RoninValidatorSet.CallOpts)
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

// Jailed is a free data retrieval call binding the contract method 0x7043e5dd.
//
// Solidity: function jailed(address _addr) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCaller) Jailed(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "jailed", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Jailed is a free data retrieval call binding the contract method 0x7043e5dd.
//
// Solidity: function jailed(address _addr) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetSession) Jailed(_addr common.Address) (bool, error) {
	return _RoninValidatorSet.Contract.Jailed(&_RoninValidatorSet.CallOpts, _addr)
}

// Jailed is a free data retrieval call binding the contract method 0x7043e5dd.
//
// Solidity: function jailed(address _addr) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) Jailed(_addr common.Address) (bool, error) {
	return _RoninValidatorSet.Contract.Jailed(&_RoninValidatorSet.CallOpts, _addr)
}

// JailedAtBlock is a free data retrieval call binding the contract method 0x2607d919.
//
// Solidity: function jailedAtBlock(address _addr, uint256 _blockNum) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCaller) JailedAtBlock(opts *bind.CallOpts, _addr common.Address, _blockNum *big.Int) (bool, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "jailedAtBlock", _addr, _blockNum)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// JailedAtBlock is a free data retrieval call binding the contract method 0x2607d919.
//
// Solidity: function jailedAtBlock(address _addr, uint256 _blockNum) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetSession) JailedAtBlock(_addr common.Address, _blockNum *big.Int) (bool, error) {
	return _RoninValidatorSet.Contract.JailedAtBlock(&_RoninValidatorSet.CallOpts, _addr, _blockNum)
}

// JailedAtBlock is a free data retrieval call binding the contract method 0x2607d919.
//
// Solidity: function jailedAtBlock(address _addr, uint256 _blockNum) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) JailedAtBlock(_addr common.Address, _blockNum *big.Int) (bool, error) {
	return _RoninValidatorSet.Contract.JailedAtBlock(&_RoninValidatorSet.CallOpts, _addr, _blockNum)
}

// JailedTimeLeft is a free data retrieval call binding the contract method 0x81f9535f.
//
// Solidity: function jailedTimeLeft(address _addr) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_RoninValidatorSet *RoninValidatorSetCaller) JailedTimeLeft(opts *bind.CallOpts, _addr common.Address) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "jailedTimeLeft", _addr)

	outstruct := new(struct {
		IsJailed  bool
		BlockLeft *big.Int
		EpochLeft *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsJailed = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.BlockLeft = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EpochLeft = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// JailedTimeLeft is a free data retrieval call binding the contract method 0x81f9535f.
//
// Solidity: function jailedTimeLeft(address _addr) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_RoninValidatorSet *RoninValidatorSetSession) JailedTimeLeft(_addr common.Address) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	return _RoninValidatorSet.Contract.JailedTimeLeft(&_RoninValidatorSet.CallOpts, _addr)
}

// JailedTimeLeft is a free data retrieval call binding the contract method 0x81f9535f.
//
// Solidity: function jailedTimeLeft(address _addr) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) JailedTimeLeft(_addr common.Address) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	return _RoninValidatorSet.Contract.JailedTimeLeft(&_RoninValidatorSet.CallOpts, _addr)
}

// JailedTimeLeftAtBlock is a free data retrieval call binding the contract method 0x85ad5aec.
//
// Solidity: function jailedTimeLeftAtBlock(address _addr, uint256 _blockNum) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_RoninValidatorSet *RoninValidatorSetCaller) JailedTimeLeftAtBlock(opts *bind.CallOpts, _addr common.Address, _blockNum *big.Int) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "jailedTimeLeftAtBlock", _addr, _blockNum)

	outstruct := new(struct {
		IsJailed  bool
		BlockLeft *big.Int
		EpochLeft *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsJailed = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.BlockLeft = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EpochLeft = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// JailedTimeLeftAtBlock is a free data retrieval call binding the contract method 0x85ad5aec.
//
// Solidity: function jailedTimeLeftAtBlock(address _addr, uint256 _blockNum) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_RoninValidatorSet *RoninValidatorSetSession) JailedTimeLeftAtBlock(_addr common.Address, _blockNum *big.Int) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	return _RoninValidatorSet.Contract.JailedTimeLeftAtBlock(&_RoninValidatorSet.CallOpts, _addr, _blockNum)
}

// JailedTimeLeftAtBlock is a free data retrieval call binding the contract method 0x85ad5aec.
//
// Solidity: function jailedTimeLeftAtBlock(address _addr, uint256 _blockNum) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) JailedTimeLeftAtBlock(_addr common.Address, _blockNum *big.Int) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	return _RoninValidatorSet.Contract.JailedTimeLeftAtBlock(&_RoninValidatorSet.CallOpts, _addr, _blockNum)
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

// MiningRewardDeprecated is a free data retrieval call binding the contract method 0x4a68f8c6.
//
// Solidity: function miningRewardDeprecated(address[] _blockProducers) view returns(bool[] _result)
func (_RoninValidatorSet *RoninValidatorSetCaller) MiningRewardDeprecated(opts *bind.CallOpts, _blockProducers []common.Address) ([]bool, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "miningRewardDeprecated", _blockProducers)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// MiningRewardDeprecated is a free data retrieval call binding the contract method 0x4a68f8c6.
//
// Solidity: function miningRewardDeprecated(address[] _blockProducers) view returns(bool[] _result)
func (_RoninValidatorSet *RoninValidatorSetSession) MiningRewardDeprecated(_blockProducers []common.Address) ([]bool, error) {
	return _RoninValidatorSet.Contract.MiningRewardDeprecated(&_RoninValidatorSet.CallOpts, _blockProducers)
}

// MiningRewardDeprecated is a free data retrieval call binding the contract method 0x4a68f8c6.
//
// Solidity: function miningRewardDeprecated(address[] _blockProducers) view returns(bool[] _result)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) MiningRewardDeprecated(_blockProducers []common.Address) ([]bool, error) {
	return _RoninValidatorSet.Contract.MiningRewardDeprecated(&_RoninValidatorSet.CallOpts, _blockProducers)
}

// MiningRewardDeprecatedAtPeriod is a free data retrieval call binding the contract method 0x92a8c2e8.
//
// Solidity: function miningRewardDeprecatedAtPeriod(address[] _blockProducers, uint256 _period) view returns(bool[] _result)
func (_RoninValidatorSet *RoninValidatorSetCaller) MiningRewardDeprecatedAtPeriod(opts *bind.CallOpts, _blockProducers []common.Address, _period *big.Int) ([]bool, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "miningRewardDeprecatedAtPeriod", _blockProducers, _period)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// MiningRewardDeprecatedAtPeriod is a free data retrieval call binding the contract method 0x92a8c2e8.
//
// Solidity: function miningRewardDeprecatedAtPeriod(address[] _blockProducers, uint256 _period) view returns(bool[] _result)
func (_RoninValidatorSet *RoninValidatorSetSession) MiningRewardDeprecatedAtPeriod(_blockProducers []common.Address, _period *big.Int) ([]bool, error) {
	return _RoninValidatorSet.Contract.MiningRewardDeprecatedAtPeriod(&_RoninValidatorSet.CallOpts, _blockProducers, _period)
}

// MiningRewardDeprecatedAtPeriod is a free data retrieval call binding the contract method 0x92a8c2e8.
//
// Solidity: function miningRewardDeprecatedAtPeriod(address[] _blockProducers, uint256 _period) view returns(bool[] _result)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) MiningRewardDeprecatedAtPeriod(_blockProducers []common.Address, _period *big.Int) ([]bool, error) {
	return _RoninValidatorSet.Contract.MiningRewardDeprecatedAtPeriod(&_RoninValidatorSet.CallOpts, _blockProducers, _period)
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

// PrecompilePickValidatorSetAddress is a free data retrieval call binding the contract method 0x3b3159b6.
//
// Solidity: function precompilePickValidatorSetAddress() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetCaller) PrecompilePickValidatorSetAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "precompilePickValidatorSetAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrecompilePickValidatorSetAddress is a free data retrieval call binding the contract method 0x3b3159b6.
//
// Solidity: function precompilePickValidatorSetAddress() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetSession) PrecompilePickValidatorSetAddress() (common.Address, error) {
	return _RoninValidatorSet.Contract.PrecompilePickValidatorSetAddress(&_RoninValidatorSet.CallOpts)
}

// PrecompilePickValidatorSetAddress is a free data retrieval call binding the contract method 0x3b3159b6.
//
// Solidity: function precompilePickValidatorSetAddress() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) PrecompilePickValidatorSetAddress() (common.Address, error) {
	return _RoninValidatorSet.Contract.PrecompilePickValidatorSetAddress(&_RoninValidatorSet.CallOpts)
}

// PrecompileSortValidatorsAddress is a free data retrieval call binding the contract method 0x8d559c38.
//
// Solidity: function precompileSortValidatorsAddress() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetCaller) PrecompileSortValidatorsAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "precompileSortValidatorsAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrecompileSortValidatorsAddress is a free data retrieval call binding the contract method 0x8d559c38.
//
// Solidity: function precompileSortValidatorsAddress() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetSession) PrecompileSortValidatorsAddress() (common.Address, error) {
	return _RoninValidatorSet.Contract.PrecompileSortValidatorsAddress(&_RoninValidatorSet.CallOpts)
}

// PrecompileSortValidatorsAddress is a free data retrieval call binding the contract method 0x8d559c38.
//
// Solidity: function precompileSortValidatorsAddress() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) PrecompileSortValidatorsAddress() (common.Address, error) {
	return _RoninValidatorSet.Contract.PrecompileSortValidatorsAddress(&_RoninValidatorSet.CallOpts)
}

// RoninTrustedOrganizationContract is a free data retrieval call binding the contract method 0x5511cde1.
//
// Solidity: function roninTrustedOrganizationContract() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetCaller) RoninTrustedOrganizationContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "roninTrustedOrganizationContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoninTrustedOrganizationContract is a free data retrieval call binding the contract method 0x5511cde1.
//
// Solidity: function roninTrustedOrganizationContract() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetSession) RoninTrustedOrganizationContract() (common.Address, error) {
	return _RoninValidatorSet.Contract.RoninTrustedOrganizationContract(&_RoninValidatorSet.CallOpts)
}

// RoninTrustedOrganizationContract is a free data retrieval call binding the contract method 0x5511cde1.
//
// Solidity: function roninTrustedOrganizationContract() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) RoninTrustedOrganizationContract() (common.Address, error) {
	return _RoninValidatorSet.Contract.RoninTrustedOrganizationContract(&_RoninValidatorSet.CallOpts)
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

// TotalBlockProducers is a free data retrieval call binding the contract method 0x9e94b9ec.
//
// Solidity: function totalBlockProducers() view returns(uint256 _total)
func (_RoninValidatorSet *RoninValidatorSetCaller) TotalBlockProducers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "totalBlockProducers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBlockProducers is a free data retrieval call binding the contract method 0x9e94b9ec.
//
// Solidity: function totalBlockProducers() view returns(uint256 _total)
func (_RoninValidatorSet *RoninValidatorSetSession) TotalBlockProducers() (*big.Int, error) {
	return _RoninValidatorSet.Contract.TotalBlockProducers(&_RoninValidatorSet.CallOpts)
}

// TotalBlockProducers is a free data retrieval call binding the contract method 0x9e94b9ec.
//
// Solidity: function totalBlockProducers() view returns(uint256 _total)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) TotalBlockProducers() (*big.Int, error) {
	return _RoninValidatorSet.Contract.TotalBlockProducers(&_RoninValidatorSet.CallOpts)
}

// TotalBridgeOperators is a free data retrieval call binding the contract method 0x562d5304.
//
// Solidity: function totalBridgeOperators() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCaller) TotalBridgeOperators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "totalBridgeOperators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBridgeOperators is a free data retrieval call binding the contract method 0x562d5304.
//
// Solidity: function totalBridgeOperators() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetSession) TotalBridgeOperators() (*big.Int, error) {
	return _RoninValidatorSet.Contract.TotalBridgeOperators(&_RoninValidatorSet.CallOpts)
}

// TotalBridgeOperators is a free data retrieval call binding the contract method 0x562d5304.
//
// Solidity: function totalBridgeOperators() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) TotalBridgeOperators() (*big.Int, error) {
	return _RoninValidatorSet.Contract.TotalBridgeOperators(&_RoninValidatorSet.CallOpts)
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

// BailOut is a paid mutator transaction binding the contract method 0xd1f992f7.
//
// Solidity: function bailOut(address _validatorAddr) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) BailOut(opts *bind.TransactOpts, _validatorAddr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "bailOut", _validatorAddr)
}

// BailOut is a paid mutator transaction binding the contract method 0xd1f992f7.
//
// Solidity: function bailOut(address _validatorAddr) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) BailOut(_validatorAddr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.BailOut(&_RoninValidatorSet.TransactOpts, _validatorAddr)
}

// BailOut is a paid mutator transaction binding the contract method 0xd1f992f7.
//
// Solidity: function bailOut(address _validatorAddr) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) BailOut(_validatorAddr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.BailOut(&_RoninValidatorSet.TransactOpts, _validatorAddr)
}

// GrantValidatorCandidate is a paid mutator transaction binding the contract method 0x733ec970.
//
// Solidity: function grantValidatorCandidate(address _admin, address _consensusAddr, address _treasuryAddr, address _bridgeOperatorAddr, uint256 _commissionRate) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) GrantValidatorCandidate(opts *bind.TransactOpts, _admin common.Address, _consensusAddr common.Address, _treasuryAddr common.Address, _bridgeOperatorAddr common.Address, _commissionRate *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "grantValidatorCandidate", _admin, _consensusAddr, _treasuryAddr, _bridgeOperatorAddr, _commissionRate)
}

// GrantValidatorCandidate is a paid mutator transaction binding the contract method 0x733ec970.
//
// Solidity: function grantValidatorCandidate(address _admin, address _consensusAddr, address _treasuryAddr, address _bridgeOperatorAddr, uint256 _commissionRate) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) GrantValidatorCandidate(_admin common.Address, _consensusAddr common.Address, _treasuryAddr common.Address, _bridgeOperatorAddr common.Address, _commissionRate *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.GrantValidatorCandidate(&_RoninValidatorSet.TransactOpts, _admin, _consensusAddr, _treasuryAddr, _bridgeOperatorAddr, _commissionRate)
}

// GrantValidatorCandidate is a paid mutator transaction binding the contract method 0x733ec970.
//
// Solidity: function grantValidatorCandidate(address _admin, address _consensusAddr, address _treasuryAddr, address _bridgeOperatorAddr, uint256 _commissionRate) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) GrantValidatorCandidate(_admin common.Address, _consensusAddr common.Address, _treasuryAddr common.Address, _bridgeOperatorAddr common.Address, _commissionRate *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.GrantValidatorCandidate(&_RoninValidatorSet.TransactOpts, _admin, _consensusAddr, _treasuryAddr, _bridgeOperatorAddr, _commissionRate)
}

// Initialize is a paid mutator transaction binding the contract method 0x3986de6a.
//
// Solidity: function initialize(address __slashIndicatorContract, address __stakingContract, address __stakingVestingContract, address __maintenanceContract, address __roninTrustedOrganizationContract, address __bridgeTrackingContract, uint256 __maxValidatorNumber, uint256 __maxValidatorCandidate, uint256 __maxPrioritizedValidatorNumber, uint256 __numberOfBlocksInEpoch) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) Initialize(opts *bind.TransactOpts, __slashIndicatorContract common.Address, __stakingContract common.Address, __stakingVestingContract common.Address, __maintenanceContract common.Address, __roninTrustedOrganizationContract common.Address, __bridgeTrackingContract common.Address, __maxValidatorNumber *big.Int, __maxValidatorCandidate *big.Int, __maxPrioritizedValidatorNumber *big.Int, __numberOfBlocksInEpoch *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "initialize", __slashIndicatorContract, __stakingContract, __stakingVestingContract, __maintenanceContract, __roninTrustedOrganizationContract, __bridgeTrackingContract, __maxValidatorNumber, __maxValidatorCandidate, __maxPrioritizedValidatorNumber, __numberOfBlocksInEpoch)
}

// Initialize is a paid mutator transaction binding the contract method 0x3986de6a.
//
// Solidity: function initialize(address __slashIndicatorContract, address __stakingContract, address __stakingVestingContract, address __maintenanceContract, address __roninTrustedOrganizationContract, address __bridgeTrackingContract, uint256 __maxValidatorNumber, uint256 __maxValidatorCandidate, uint256 __maxPrioritizedValidatorNumber, uint256 __numberOfBlocksInEpoch) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) Initialize(__slashIndicatorContract common.Address, __stakingContract common.Address, __stakingVestingContract common.Address, __maintenanceContract common.Address, __roninTrustedOrganizationContract common.Address, __bridgeTrackingContract common.Address, __maxValidatorNumber *big.Int, __maxValidatorCandidate *big.Int, __maxPrioritizedValidatorNumber *big.Int, __numberOfBlocksInEpoch *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.Initialize(&_RoninValidatorSet.TransactOpts, __slashIndicatorContract, __stakingContract, __stakingVestingContract, __maintenanceContract, __roninTrustedOrganizationContract, __bridgeTrackingContract, __maxValidatorNumber, __maxValidatorCandidate, __maxPrioritizedValidatorNumber, __numberOfBlocksInEpoch)
}

// Initialize is a paid mutator transaction binding the contract method 0x3986de6a.
//
// Solidity: function initialize(address __slashIndicatorContract, address __stakingContract, address __stakingVestingContract, address __maintenanceContract, address __roninTrustedOrganizationContract, address __bridgeTrackingContract, uint256 __maxValidatorNumber, uint256 __maxValidatorCandidate, uint256 __maxPrioritizedValidatorNumber, uint256 __numberOfBlocksInEpoch) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) Initialize(__slashIndicatorContract common.Address, __stakingContract common.Address, __stakingVestingContract common.Address, __maintenanceContract common.Address, __roninTrustedOrganizationContract common.Address, __bridgeTrackingContract common.Address, __maxValidatorNumber *big.Int, __maxValidatorCandidate *big.Int, __maxPrioritizedValidatorNumber *big.Int, __numberOfBlocksInEpoch *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.Initialize(&_RoninValidatorSet.TransactOpts, __slashIndicatorContract, __stakingContract, __stakingVestingContract, __maintenanceContract, __roninTrustedOrganizationContract, __bridgeTrackingContract, __maxValidatorNumber, __maxValidatorCandidate, __maxPrioritizedValidatorNumber, __numberOfBlocksInEpoch)
}

// RequestRevokeCandidate is a paid mutator transaction binding the contract method 0x86b60e1a.
//
// Solidity: function requestRevokeCandidate(address _consensusAddr) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) RequestRevokeCandidate(opts *bind.TransactOpts, _consensusAddr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "requestRevokeCandidate", _consensusAddr)
}

// RequestRevokeCandidate is a paid mutator transaction binding the contract method 0x86b60e1a.
//
// Solidity: function requestRevokeCandidate(address _consensusAddr) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) RequestRevokeCandidate(_consensusAddr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.RequestRevokeCandidate(&_RoninValidatorSet.TransactOpts, _consensusAddr)
}

// RequestRevokeCandidate is a paid mutator transaction binding the contract method 0x86b60e1a.
//
// Solidity: function requestRevokeCandidate(address _consensusAddr) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) RequestRevokeCandidate(_consensusAddr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.RequestRevokeCandidate(&_RoninValidatorSet.TransactOpts, _consensusAddr)
}

// SetBridgeTrackingContract is a paid mutator transaction binding the contract method 0x9c8d98da.
//
// Solidity: function setBridgeTrackingContract(address _addr) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) SetBridgeTrackingContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "setBridgeTrackingContract", _addr)
}

// SetBridgeTrackingContract is a paid mutator transaction binding the contract method 0x9c8d98da.
//
// Solidity: function setBridgeTrackingContract(address _addr) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) SetBridgeTrackingContract(_addr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetBridgeTrackingContract(&_RoninValidatorSet.TransactOpts, _addr)
}

// SetBridgeTrackingContract is a paid mutator transaction binding the contract method 0x9c8d98da.
//
// Solidity: function setBridgeTrackingContract(address _addr) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) SetBridgeTrackingContract(_addr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetBridgeTrackingContract(&_RoninValidatorSet.TransactOpts, _addr)
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
// Solidity: function setMaxValidatorNumber(uint256 _max) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) SetMaxValidatorNumber(opts *bind.TransactOpts, _max *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "setMaxValidatorNumber", _max)
}

// SetMaxValidatorNumber is a paid mutator transaction binding the contract method 0x823a7b9c.
//
// Solidity: function setMaxValidatorNumber(uint256 _max) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) SetMaxValidatorNumber(_max *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetMaxValidatorNumber(&_RoninValidatorSet.TransactOpts, _max)
}

// SetMaxValidatorNumber is a paid mutator transaction binding the contract method 0x823a7b9c.
//
// Solidity: function setMaxValidatorNumber(uint256 _max) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) SetMaxValidatorNumber(_max *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetMaxValidatorNumber(&_RoninValidatorSet.TransactOpts, _max)
}

// SetNumberOfBlocksInEpoch is a paid mutator transaction binding the contract method 0xd72733fc.
//
// Solidity: function setNumberOfBlocksInEpoch(uint256 _number) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) SetNumberOfBlocksInEpoch(opts *bind.TransactOpts, _number *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "setNumberOfBlocksInEpoch", _number)
}

// SetNumberOfBlocksInEpoch is a paid mutator transaction binding the contract method 0xd72733fc.
//
// Solidity: function setNumberOfBlocksInEpoch(uint256 _number) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) SetNumberOfBlocksInEpoch(_number *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetNumberOfBlocksInEpoch(&_RoninValidatorSet.TransactOpts, _number)
}

// SetNumberOfBlocksInEpoch is a paid mutator transaction binding the contract method 0xd72733fc.
//
// Solidity: function setNumberOfBlocksInEpoch(uint256 _number) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) SetNumberOfBlocksInEpoch(_number *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetNumberOfBlocksInEpoch(&_RoninValidatorSet.TransactOpts, _number)
}

// SetRoninTrustedOrganizationContract is a paid mutator transaction binding the contract method 0xb5e337de.
//
// Solidity: function setRoninTrustedOrganizationContract(address _addr) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) SetRoninTrustedOrganizationContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "setRoninTrustedOrganizationContract", _addr)
}

// SetRoninTrustedOrganizationContract is a paid mutator transaction binding the contract method 0xb5e337de.
//
// Solidity: function setRoninTrustedOrganizationContract(address _addr) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) SetRoninTrustedOrganizationContract(_addr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetRoninTrustedOrganizationContract(&_RoninValidatorSet.TransactOpts, _addr)
}

// SetRoninTrustedOrganizationContract is a paid mutator transaction binding the contract method 0xb5e337de.
//
// Solidity: function setRoninTrustedOrganizationContract(address _addr) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) SetRoninTrustedOrganizationContract(_addr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetRoninTrustedOrganizationContract(&_RoninValidatorSet.TransactOpts, _addr)
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

// RoninValidatorSetBlockProducerSetUpdatedIterator is returned from FilterBlockProducerSetUpdated and is used to iterate over the raw logs and unpacked data for BlockProducerSetUpdated events raised by the RoninValidatorSet contract.
type RoninValidatorSetBlockProducerSetUpdatedIterator struct {
	Event *RoninValidatorSetBlockProducerSetUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetBlockProducerSetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetBlockProducerSetUpdated)
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
		it.Event = new(RoninValidatorSetBlockProducerSetUpdated)
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
func (it *RoninValidatorSetBlockProducerSetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetBlockProducerSetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetBlockProducerSetUpdated represents a BlockProducerSetUpdated event raised by the RoninValidatorSet contract.
type RoninValidatorSetBlockProducerSetUpdated struct {
	Period         *big.Int
	ConsensusAddrs []common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlockProducerSetUpdated is a free log retrieval operation binding the contract event 0x60324bb9c8b0d077621d76762c52d6cc937043427992a2f6a602b449315922ef.
//
// Solidity: event BlockProducerSetUpdated(uint256 indexed period, address[] consensusAddrs)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterBlockProducerSetUpdated(opts *bind.FilterOpts, period []*big.Int) (*RoninValidatorSetBlockProducerSetUpdatedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "BlockProducerSetUpdated", periodRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetBlockProducerSetUpdatedIterator{contract: _RoninValidatorSet.contract, event: "BlockProducerSetUpdated", logs: logs, sub: sub}, nil
}

// WatchBlockProducerSetUpdated is a free log subscription operation binding the contract event 0x60324bb9c8b0d077621d76762c52d6cc937043427992a2f6a602b449315922ef.
//
// Solidity: event BlockProducerSetUpdated(uint256 indexed period, address[] consensusAddrs)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchBlockProducerSetUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetBlockProducerSetUpdated, period []*big.Int) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "BlockProducerSetUpdated", periodRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetBlockProducerSetUpdated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "BlockProducerSetUpdated", log); err != nil {
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

// ParseBlockProducerSetUpdated is a log parse operation binding the contract event 0x60324bb9c8b0d077621d76762c52d6cc937043427992a2f6a602b449315922ef.
//
// Solidity: event BlockProducerSetUpdated(uint256 indexed period, address[] consensusAddrs)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseBlockProducerSetUpdated(log types.Log) (*RoninValidatorSetBlockProducerSetUpdated, error) {
	event := new(RoninValidatorSetBlockProducerSetUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "BlockProducerSetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetBlockRewardRewardDeprecatedIterator is returned from FilterBlockRewardRewardDeprecated and is used to iterate over the raw logs and unpacked data for BlockRewardRewardDeprecated events raised by the RoninValidatorSet contract.
type RoninValidatorSetBlockRewardRewardDeprecatedIterator struct {
	Event *RoninValidatorSetBlockRewardRewardDeprecated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetBlockRewardRewardDeprecatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetBlockRewardRewardDeprecated)
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
		it.Event = new(RoninValidatorSetBlockRewardRewardDeprecated)
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
func (it *RoninValidatorSetBlockRewardRewardDeprecatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetBlockRewardRewardDeprecatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetBlockRewardRewardDeprecated represents a BlockRewardRewardDeprecated event raised by the RoninValidatorSet contract.
type RoninValidatorSetBlockRewardRewardDeprecated struct {
	CoinbaseAddr common.Address
	RewardAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBlockRewardRewardDeprecated is a free log retrieval operation binding the contract event 0xdbb176bd1de20fc558f30c5d00c1f7818d3a9f6835105a2b5b9c2beeffc2d394.
//
// Solidity: event BlockRewardRewardDeprecated(address indexed coinbaseAddr, uint256 rewardAmount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterBlockRewardRewardDeprecated(opts *bind.FilterOpts, coinbaseAddr []common.Address) (*RoninValidatorSetBlockRewardRewardDeprecatedIterator, error) {

	var coinbaseAddrRule []interface{}
	for _, coinbaseAddrItem := range coinbaseAddr {
		coinbaseAddrRule = append(coinbaseAddrRule, coinbaseAddrItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "BlockRewardRewardDeprecated", coinbaseAddrRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetBlockRewardRewardDeprecatedIterator{contract: _RoninValidatorSet.contract, event: "BlockRewardRewardDeprecated", logs: logs, sub: sub}, nil
}

// WatchBlockRewardRewardDeprecated is a free log subscription operation binding the contract event 0xdbb176bd1de20fc558f30c5d00c1f7818d3a9f6835105a2b5b9c2beeffc2d394.
//
// Solidity: event BlockRewardRewardDeprecated(address indexed coinbaseAddr, uint256 rewardAmount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchBlockRewardRewardDeprecated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetBlockRewardRewardDeprecated, coinbaseAddr []common.Address) (event.Subscription, error) {

	var coinbaseAddrRule []interface{}
	for _, coinbaseAddrItem := range coinbaseAddr {
		coinbaseAddrRule = append(coinbaseAddrRule, coinbaseAddrItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "BlockRewardRewardDeprecated", coinbaseAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetBlockRewardRewardDeprecated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "BlockRewardRewardDeprecated", log); err != nil {
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

// ParseBlockRewardRewardDeprecated is a log parse operation binding the contract event 0xdbb176bd1de20fc558f30c5d00c1f7818d3a9f6835105a2b5b9c2beeffc2d394.
//
// Solidity: event BlockRewardRewardDeprecated(address indexed coinbaseAddr, uint256 rewardAmount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseBlockRewardRewardDeprecated(log types.Log) (*RoninValidatorSetBlockRewardRewardDeprecated, error) {
	event := new(RoninValidatorSetBlockRewardRewardDeprecated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "BlockRewardRewardDeprecated", log); err != nil {
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
// Solidity: event BlockRewardSubmitted(address indexed coinbaseAddr, uint256 submittedAmount, uint256 bonusAmount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterBlockRewardSubmitted(opts *bind.FilterOpts, coinbaseAddr []common.Address) (*RoninValidatorSetBlockRewardSubmittedIterator, error) {

	var coinbaseAddrRule []interface{}
	for _, coinbaseAddrItem := range coinbaseAddr {
		coinbaseAddrRule = append(coinbaseAddrRule, coinbaseAddrItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "BlockRewardSubmitted", coinbaseAddrRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetBlockRewardSubmittedIterator{contract: _RoninValidatorSet.contract, event: "BlockRewardSubmitted", logs: logs, sub: sub}, nil
}

// WatchBlockRewardSubmitted is a free log subscription operation binding the contract event 0x0ede5c3be8625943fa64003cd4b91230089411249f3059bac6500873543ca9b1.
//
// Solidity: event BlockRewardSubmitted(address indexed coinbaseAddr, uint256 submittedAmount, uint256 bonusAmount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchBlockRewardSubmitted(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetBlockRewardSubmitted, coinbaseAddr []common.Address) (event.Subscription, error) {

	var coinbaseAddrRule []interface{}
	for _, coinbaseAddrItem := range coinbaseAddr {
		coinbaseAddrRule = append(coinbaseAddrRule, coinbaseAddrItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "BlockRewardSubmitted", coinbaseAddrRule)
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
// Solidity: event BlockRewardSubmitted(address indexed coinbaseAddr, uint256 submittedAmount, uint256 bonusAmount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseBlockRewardSubmitted(log types.Log) (*RoninValidatorSetBlockRewardSubmitted, error) {
	event := new(RoninValidatorSetBlockRewardSubmitted)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "BlockRewardSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetBridgeOperatorRewardDistributedIterator is returned from FilterBridgeOperatorRewardDistributed and is used to iterate over the raw logs and unpacked data for BridgeOperatorRewardDistributed events raised by the RoninValidatorSet contract.
type RoninValidatorSetBridgeOperatorRewardDistributedIterator struct {
	Event *RoninValidatorSetBridgeOperatorRewardDistributed // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetBridgeOperatorRewardDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetBridgeOperatorRewardDistributed)
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
		it.Event = new(RoninValidatorSetBridgeOperatorRewardDistributed)
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
func (it *RoninValidatorSetBridgeOperatorRewardDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetBridgeOperatorRewardDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetBridgeOperatorRewardDistributed represents a BridgeOperatorRewardDistributed event raised by the RoninValidatorSet contract.
type RoninValidatorSetBridgeOperatorRewardDistributed struct {
	ConsensusAddr  common.Address
	BridgeOperator common.Address
	RecipientAddr  common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeOperatorRewardDistributed is a free log retrieval operation binding the contract event 0x72a57dc38837a1cba7881b7b1a5594d9e6b65cec6a985b54e2cee3e89369691c.
//
// Solidity: event BridgeOperatorRewardDistributed(address indexed consensusAddr, address indexed bridgeOperator, address indexed recipientAddr, uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterBridgeOperatorRewardDistributed(opts *bind.FilterOpts, consensusAddr []common.Address, bridgeOperator []common.Address, recipientAddr []common.Address) (*RoninValidatorSetBridgeOperatorRewardDistributedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var bridgeOperatorRule []interface{}
	for _, bridgeOperatorItem := range bridgeOperator {
		bridgeOperatorRule = append(bridgeOperatorRule, bridgeOperatorItem)
	}
	var recipientAddrRule []interface{}
	for _, recipientAddrItem := range recipientAddr {
		recipientAddrRule = append(recipientAddrRule, recipientAddrItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "BridgeOperatorRewardDistributed", consensusAddrRule, bridgeOperatorRule, recipientAddrRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetBridgeOperatorRewardDistributedIterator{contract: _RoninValidatorSet.contract, event: "BridgeOperatorRewardDistributed", logs: logs, sub: sub}, nil
}

// WatchBridgeOperatorRewardDistributed is a free log subscription operation binding the contract event 0x72a57dc38837a1cba7881b7b1a5594d9e6b65cec6a985b54e2cee3e89369691c.
//
// Solidity: event BridgeOperatorRewardDistributed(address indexed consensusAddr, address indexed bridgeOperator, address indexed recipientAddr, uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchBridgeOperatorRewardDistributed(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetBridgeOperatorRewardDistributed, consensusAddr []common.Address, bridgeOperator []common.Address, recipientAddr []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var bridgeOperatorRule []interface{}
	for _, bridgeOperatorItem := range bridgeOperator {
		bridgeOperatorRule = append(bridgeOperatorRule, bridgeOperatorItem)
	}
	var recipientAddrRule []interface{}
	for _, recipientAddrItem := range recipientAddr {
		recipientAddrRule = append(recipientAddrRule, recipientAddrItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "BridgeOperatorRewardDistributed", consensusAddrRule, bridgeOperatorRule, recipientAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetBridgeOperatorRewardDistributed)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "BridgeOperatorRewardDistributed", log); err != nil {
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

// ParseBridgeOperatorRewardDistributed is a log parse operation binding the contract event 0x72a57dc38837a1cba7881b7b1a5594d9e6b65cec6a985b54e2cee3e89369691c.
//
// Solidity: event BridgeOperatorRewardDistributed(address indexed consensusAddr, address indexed bridgeOperator, address indexed recipientAddr, uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseBridgeOperatorRewardDistributed(log types.Log) (*RoninValidatorSetBridgeOperatorRewardDistributed, error) {
	event := new(RoninValidatorSetBridgeOperatorRewardDistributed)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "BridgeOperatorRewardDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetBridgeOperatorRewardDistributionFailedIterator is returned from FilterBridgeOperatorRewardDistributionFailed and is used to iterate over the raw logs and unpacked data for BridgeOperatorRewardDistributionFailed events raised by the RoninValidatorSet contract.
type RoninValidatorSetBridgeOperatorRewardDistributionFailedIterator struct {
	Event *RoninValidatorSetBridgeOperatorRewardDistributionFailed // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetBridgeOperatorRewardDistributionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetBridgeOperatorRewardDistributionFailed)
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
		it.Event = new(RoninValidatorSetBridgeOperatorRewardDistributionFailed)
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
func (it *RoninValidatorSetBridgeOperatorRewardDistributionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetBridgeOperatorRewardDistributionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetBridgeOperatorRewardDistributionFailed represents a BridgeOperatorRewardDistributionFailed event raised by the RoninValidatorSet contract.
type RoninValidatorSetBridgeOperatorRewardDistributionFailed struct {
	ConsensusAddr   common.Address
	BridgeOperator  common.Address
	Recipient       common.Address
	Amount          *big.Int
	ContractBalance *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBridgeOperatorRewardDistributionFailed is a free log retrieval operation binding the contract event 0xd35d76d87d51ed89407fc7ceaaccf32cf72784b94530892ce33546540e141b72.
//
// Solidity: event BridgeOperatorRewardDistributionFailed(address indexed consensusAddr, address indexed bridgeOperator, address indexed recipient, uint256 amount, uint256 contractBalance)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterBridgeOperatorRewardDistributionFailed(opts *bind.FilterOpts, consensusAddr []common.Address, bridgeOperator []common.Address, recipient []common.Address) (*RoninValidatorSetBridgeOperatorRewardDistributionFailedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var bridgeOperatorRule []interface{}
	for _, bridgeOperatorItem := range bridgeOperator {
		bridgeOperatorRule = append(bridgeOperatorRule, bridgeOperatorItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "BridgeOperatorRewardDistributionFailed", consensusAddrRule, bridgeOperatorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetBridgeOperatorRewardDistributionFailedIterator{contract: _RoninValidatorSet.contract, event: "BridgeOperatorRewardDistributionFailed", logs: logs, sub: sub}, nil
}

// WatchBridgeOperatorRewardDistributionFailed is a free log subscription operation binding the contract event 0xd35d76d87d51ed89407fc7ceaaccf32cf72784b94530892ce33546540e141b72.
//
// Solidity: event BridgeOperatorRewardDistributionFailed(address indexed consensusAddr, address indexed bridgeOperator, address indexed recipient, uint256 amount, uint256 contractBalance)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchBridgeOperatorRewardDistributionFailed(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetBridgeOperatorRewardDistributionFailed, consensusAddr []common.Address, bridgeOperator []common.Address, recipient []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var bridgeOperatorRule []interface{}
	for _, bridgeOperatorItem := range bridgeOperator {
		bridgeOperatorRule = append(bridgeOperatorRule, bridgeOperatorItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "BridgeOperatorRewardDistributionFailed", consensusAddrRule, bridgeOperatorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetBridgeOperatorRewardDistributionFailed)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "BridgeOperatorRewardDistributionFailed", log); err != nil {
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

// ParseBridgeOperatorRewardDistributionFailed is a log parse operation binding the contract event 0xd35d76d87d51ed89407fc7ceaaccf32cf72784b94530892ce33546540e141b72.
//
// Solidity: event BridgeOperatorRewardDistributionFailed(address indexed consensusAddr, address indexed bridgeOperator, address indexed recipient, uint256 amount, uint256 contractBalance)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseBridgeOperatorRewardDistributionFailed(log types.Log) (*RoninValidatorSetBridgeOperatorRewardDistributionFailed, error) {
	event := new(RoninValidatorSetBridgeOperatorRewardDistributionFailed)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "BridgeOperatorRewardDistributionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetBridgeOperatorSetUpdatedIterator is returned from FilterBridgeOperatorSetUpdated and is used to iterate over the raw logs and unpacked data for BridgeOperatorSetUpdated events raised by the RoninValidatorSet contract.
type RoninValidatorSetBridgeOperatorSetUpdatedIterator struct {
	Event *RoninValidatorSetBridgeOperatorSetUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetBridgeOperatorSetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetBridgeOperatorSetUpdated)
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
		it.Event = new(RoninValidatorSetBridgeOperatorSetUpdated)
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
func (it *RoninValidatorSetBridgeOperatorSetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetBridgeOperatorSetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetBridgeOperatorSetUpdated represents a BridgeOperatorSetUpdated event raised by the RoninValidatorSet contract.
type RoninValidatorSetBridgeOperatorSetUpdated struct {
	Period          *big.Int
	BridgeOperators []common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBridgeOperatorSetUpdated is a free log retrieval operation binding the contract event 0x8d7d519e81c2b8dc67b44fd645fd2c8805110d9ab1d643e3dd68b622bde331ff.
//
// Solidity: event BridgeOperatorSetUpdated(uint256 indexed period, address[] bridgeOperators)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterBridgeOperatorSetUpdated(opts *bind.FilterOpts, period []*big.Int) (*RoninValidatorSetBridgeOperatorSetUpdatedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "BridgeOperatorSetUpdated", periodRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetBridgeOperatorSetUpdatedIterator{contract: _RoninValidatorSet.contract, event: "BridgeOperatorSetUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeOperatorSetUpdated is a free log subscription operation binding the contract event 0x8d7d519e81c2b8dc67b44fd645fd2c8805110d9ab1d643e3dd68b622bde331ff.
//
// Solidity: event BridgeOperatorSetUpdated(uint256 indexed period, address[] bridgeOperators)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchBridgeOperatorSetUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetBridgeOperatorSetUpdated, period []*big.Int) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "BridgeOperatorSetUpdated", periodRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetBridgeOperatorSetUpdated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "BridgeOperatorSetUpdated", log); err != nil {
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

// ParseBridgeOperatorSetUpdated is a log parse operation binding the contract event 0x8d7d519e81c2b8dc67b44fd645fd2c8805110d9ab1d643e3dd68b622bde331ff.
//
// Solidity: event BridgeOperatorSetUpdated(uint256 indexed period, address[] bridgeOperators)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseBridgeOperatorSetUpdated(log types.Log) (*RoninValidatorSetBridgeOperatorSetUpdated, error) {
	event := new(RoninValidatorSetBridgeOperatorSetUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "BridgeOperatorSetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetBridgeTrackingContractUpdatedIterator is returned from FilterBridgeTrackingContractUpdated and is used to iterate over the raw logs and unpacked data for BridgeTrackingContractUpdated events raised by the RoninValidatorSet contract.
type RoninValidatorSetBridgeTrackingContractUpdatedIterator struct {
	Event *RoninValidatorSetBridgeTrackingContractUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetBridgeTrackingContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetBridgeTrackingContractUpdated)
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
		it.Event = new(RoninValidatorSetBridgeTrackingContractUpdated)
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
func (it *RoninValidatorSetBridgeTrackingContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetBridgeTrackingContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetBridgeTrackingContractUpdated represents a BridgeTrackingContractUpdated event raised by the RoninValidatorSet contract.
type RoninValidatorSetBridgeTrackingContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBridgeTrackingContractUpdated is a free log retrieval operation binding the contract event 0x034c8da497df28467c79ddadbba1cc3cdd41f510ea73faae271e6f16a6111621.
//
// Solidity: event BridgeTrackingContractUpdated(address arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterBridgeTrackingContractUpdated(opts *bind.FilterOpts) (*RoninValidatorSetBridgeTrackingContractUpdatedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "BridgeTrackingContractUpdated")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetBridgeTrackingContractUpdatedIterator{contract: _RoninValidatorSet.contract, event: "BridgeTrackingContractUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeTrackingContractUpdated is a free log subscription operation binding the contract event 0x034c8da497df28467c79ddadbba1cc3cdd41f510ea73faae271e6f16a6111621.
//
// Solidity: event BridgeTrackingContractUpdated(address arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchBridgeTrackingContractUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetBridgeTrackingContractUpdated) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "BridgeTrackingContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetBridgeTrackingContractUpdated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "BridgeTrackingContractUpdated", log); err != nil {
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

// ParseBridgeTrackingContractUpdated is a log parse operation binding the contract event 0x034c8da497df28467c79ddadbba1cc3cdd41f510ea73faae271e6f16a6111621.
//
// Solidity: event BridgeTrackingContractUpdated(address arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseBridgeTrackingContractUpdated(log types.Log) (*RoninValidatorSetBridgeTrackingContractUpdated, error) {
	event := new(RoninValidatorSetBridgeTrackingContractUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "BridgeTrackingContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetCandidateGrantedIterator is returned from FilterCandidateGranted and is used to iterate over the raw logs and unpacked data for CandidateGranted events raised by the RoninValidatorSet contract.
type RoninValidatorSetCandidateGrantedIterator struct {
	Event *RoninValidatorSetCandidateGranted // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetCandidateGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetCandidateGranted)
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
		it.Event = new(RoninValidatorSetCandidateGranted)
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
func (it *RoninValidatorSetCandidateGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetCandidateGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetCandidateGranted represents a CandidateGranted event raised by the RoninValidatorSet contract.
type RoninValidatorSetCandidateGranted struct {
	ConsensusAddr  common.Address
	TreasuryAddr   common.Address
	Admin          common.Address
	BridgeOperator common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCandidateGranted is a free log retrieval operation binding the contract event 0xd690f592ed983cfbc05717fbcf06c4e10ae328432c309fe49246cf4a4be69fcd.
//
// Solidity: event CandidateGranted(address indexed consensusAddr, address indexed treasuryAddr, address indexed admin, address bridgeOperator)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterCandidateGranted(opts *bind.FilterOpts, consensusAddr []common.Address, treasuryAddr []common.Address, admin []common.Address) (*RoninValidatorSetCandidateGrantedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var treasuryAddrRule []interface{}
	for _, treasuryAddrItem := range treasuryAddr {
		treasuryAddrRule = append(treasuryAddrRule, treasuryAddrItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "CandidateGranted", consensusAddrRule, treasuryAddrRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetCandidateGrantedIterator{contract: _RoninValidatorSet.contract, event: "CandidateGranted", logs: logs, sub: sub}, nil
}

// WatchCandidateGranted is a free log subscription operation binding the contract event 0xd690f592ed983cfbc05717fbcf06c4e10ae328432c309fe49246cf4a4be69fcd.
//
// Solidity: event CandidateGranted(address indexed consensusAddr, address indexed treasuryAddr, address indexed admin, address bridgeOperator)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchCandidateGranted(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetCandidateGranted, consensusAddr []common.Address, treasuryAddr []common.Address, admin []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var treasuryAddrRule []interface{}
	for _, treasuryAddrItem := range treasuryAddr {
		treasuryAddrRule = append(treasuryAddrRule, treasuryAddrItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "CandidateGranted", consensusAddrRule, treasuryAddrRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetCandidateGranted)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "CandidateGranted", log); err != nil {
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

// ParseCandidateGranted is a log parse operation binding the contract event 0xd690f592ed983cfbc05717fbcf06c4e10ae328432c309fe49246cf4a4be69fcd.
//
// Solidity: event CandidateGranted(address indexed consensusAddr, address indexed treasuryAddr, address indexed admin, address bridgeOperator)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseCandidateGranted(log types.Log) (*RoninValidatorSetCandidateGranted, error) {
	event := new(RoninValidatorSetCandidateGranted)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "CandidateGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetCandidateRevokedPeriodUpdatedIterator is returned from FilterCandidateRevokedPeriodUpdated and is used to iterate over the raw logs and unpacked data for CandidateRevokedPeriodUpdated events raised by the RoninValidatorSet contract.
type RoninValidatorSetCandidateRevokedPeriodUpdatedIterator struct {
	Event *RoninValidatorSetCandidateRevokedPeriodUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetCandidateRevokedPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetCandidateRevokedPeriodUpdated)
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
		it.Event = new(RoninValidatorSetCandidateRevokedPeriodUpdated)
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
func (it *RoninValidatorSetCandidateRevokedPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetCandidateRevokedPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetCandidateRevokedPeriodUpdated represents a CandidateRevokedPeriodUpdated event raised by the RoninValidatorSet contract.
type RoninValidatorSetCandidateRevokedPeriodUpdated struct {
	ConsensusAddr common.Address
	RevokedPeriod *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCandidateRevokedPeriodUpdated is a free log retrieval operation binding the contract event 0xcc3e2d110665b49361ca48964d502edcb29866be463a5aafcff5b6aac69359a3.
//
// Solidity: event CandidateRevokedPeriodUpdated(address indexed consensusAddr, uint256 revokedPeriod)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterCandidateRevokedPeriodUpdated(opts *bind.FilterOpts, consensusAddr []common.Address) (*RoninValidatorSetCandidateRevokedPeriodUpdatedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "CandidateRevokedPeriodUpdated", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetCandidateRevokedPeriodUpdatedIterator{contract: _RoninValidatorSet.contract, event: "CandidateRevokedPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchCandidateRevokedPeriodUpdated is a free log subscription operation binding the contract event 0xcc3e2d110665b49361ca48964d502edcb29866be463a5aafcff5b6aac69359a3.
//
// Solidity: event CandidateRevokedPeriodUpdated(address indexed consensusAddr, uint256 revokedPeriod)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchCandidateRevokedPeriodUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetCandidateRevokedPeriodUpdated, consensusAddr []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "CandidateRevokedPeriodUpdated", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetCandidateRevokedPeriodUpdated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "CandidateRevokedPeriodUpdated", log); err != nil {
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

// ParseCandidateRevokedPeriodUpdated is a log parse operation binding the contract event 0xcc3e2d110665b49361ca48964d502edcb29866be463a5aafcff5b6aac69359a3.
//
// Solidity: event CandidateRevokedPeriodUpdated(address indexed consensusAddr, uint256 revokedPeriod)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseCandidateRevokedPeriodUpdated(log types.Log) (*RoninValidatorSetCandidateRevokedPeriodUpdated, error) {
	event := new(RoninValidatorSetCandidateRevokedPeriodUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "CandidateRevokedPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetCandidatesRevokedIterator is returned from FilterCandidatesRevoked and is used to iterate over the raw logs and unpacked data for CandidatesRevoked events raised by the RoninValidatorSet contract.
type RoninValidatorSetCandidatesRevokedIterator struct {
	Event *RoninValidatorSetCandidatesRevoked // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetCandidatesRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetCandidatesRevoked)
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
		it.Event = new(RoninValidatorSetCandidatesRevoked)
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
func (it *RoninValidatorSetCandidatesRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetCandidatesRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetCandidatesRevoked represents a CandidatesRevoked event raised by the RoninValidatorSet contract.
type RoninValidatorSetCandidatesRevoked struct {
	ConsensusAddrs []common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCandidatesRevoked is a free log retrieval operation binding the contract event 0x4eaf233b9dc25a5552c1927feee1412eea69add17c2485c831c2e60e234f3c91.
//
// Solidity: event CandidatesRevoked(address[] consensusAddrs)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterCandidatesRevoked(opts *bind.FilterOpts) (*RoninValidatorSetCandidatesRevokedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "CandidatesRevoked")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetCandidatesRevokedIterator{contract: _RoninValidatorSet.contract, event: "CandidatesRevoked", logs: logs, sub: sub}, nil
}

// WatchCandidatesRevoked is a free log subscription operation binding the contract event 0x4eaf233b9dc25a5552c1927feee1412eea69add17c2485c831c2e60e234f3c91.
//
// Solidity: event CandidatesRevoked(address[] consensusAddrs)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchCandidatesRevoked(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetCandidatesRevoked) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "CandidatesRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetCandidatesRevoked)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "CandidatesRevoked", log); err != nil {
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

// ParseCandidatesRevoked is a log parse operation binding the contract event 0x4eaf233b9dc25a5552c1927feee1412eea69add17c2485c831c2e60e234f3c91.
//
// Solidity: event CandidatesRevoked(address[] consensusAddrs)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseCandidatesRevoked(log types.Log) (*RoninValidatorSetCandidatesRevoked, error) {
	event := new(RoninValidatorSetCandidatesRevoked)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "CandidatesRevoked", log); err != nil {
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
	ConsensusAddr common.Address
	Recipient     common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMiningRewardDistributed is a free log retrieval operation binding the contract event 0x1ce7a1c4702402cd393500acb1de5bd927727a54e144a587d328f1b679abe4ec.
//
// Solidity: event MiningRewardDistributed(address indexed consensusAddr, address indexed recipient, uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterMiningRewardDistributed(opts *bind.FilterOpts, consensusAddr []common.Address, recipient []common.Address) (*RoninValidatorSetMiningRewardDistributedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "MiningRewardDistributed", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetMiningRewardDistributedIterator{contract: _RoninValidatorSet.contract, event: "MiningRewardDistributed", logs: logs, sub: sub}, nil
}

// WatchMiningRewardDistributed is a free log subscription operation binding the contract event 0x1ce7a1c4702402cd393500acb1de5bd927727a54e144a587d328f1b679abe4ec.
//
// Solidity: event MiningRewardDistributed(address indexed consensusAddr, address indexed recipient, uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchMiningRewardDistributed(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetMiningRewardDistributed, consensusAddr []common.Address, recipient []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "MiningRewardDistributed", consensusAddrRule, recipientRule)
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

// ParseMiningRewardDistributed is a log parse operation binding the contract event 0x1ce7a1c4702402cd393500acb1de5bd927727a54e144a587d328f1b679abe4ec.
//
// Solidity: event MiningRewardDistributed(address indexed consensusAddr, address indexed recipient, uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseMiningRewardDistributed(log types.Log) (*RoninValidatorSetMiningRewardDistributed, error) {
	event := new(RoninValidatorSetMiningRewardDistributed)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "MiningRewardDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetMiningRewardDistributionFailedIterator is returned from FilterMiningRewardDistributionFailed and is used to iterate over the raw logs and unpacked data for MiningRewardDistributionFailed events raised by the RoninValidatorSet contract.
type RoninValidatorSetMiningRewardDistributionFailedIterator struct {
	Event *RoninValidatorSetMiningRewardDistributionFailed // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetMiningRewardDistributionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetMiningRewardDistributionFailed)
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
		it.Event = new(RoninValidatorSetMiningRewardDistributionFailed)
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
func (it *RoninValidatorSetMiningRewardDistributionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetMiningRewardDistributionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetMiningRewardDistributionFailed represents a MiningRewardDistributionFailed event raised by the RoninValidatorSet contract.
type RoninValidatorSetMiningRewardDistributionFailed struct {
	ConsensusAddr   common.Address
	Recipient       common.Address
	Amount          *big.Int
	ContractBalance *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMiningRewardDistributionFailed is a free log retrieval operation binding the contract event 0x6c69e09ee5c5ac33c0cd57787261c5bade070a392ab34a4b5487c6868f723f6e.
//
// Solidity: event MiningRewardDistributionFailed(address indexed consensusAddr, address indexed recipient, uint256 amount, uint256 contractBalance)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterMiningRewardDistributionFailed(opts *bind.FilterOpts, consensusAddr []common.Address, recipient []common.Address) (*RoninValidatorSetMiningRewardDistributionFailedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "MiningRewardDistributionFailed", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetMiningRewardDistributionFailedIterator{contract: _RoninValidatorSet.contract, event: "MiningRewardDistributionFailed", logs: logs, sub: sub}, nil
}

// WatchMiningRewardDistributionFailed is a free log subscription operation binding the contract event 0x6c69e09ee5c5ac33c0cd57787261c5bade070a392ab34a4b5487c6868f723f6e.
//
// Solidity: event MiningRewardDistributionFailed(address indexed consensusAddr, address indexed recipient, uint256 amount, uint256 contractBalance)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchMiningRewardDistributionFailed(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetMiningRewardDistributionFailed, consensusAddr []common.Address, recipient []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "MiningRewardDistributionFailed", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetMiningRewardDistributionFailed)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "MiningRewardDistributionFailed", log); err != nil {
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

// ParseMiningRewardDistributionFailed is a log parse operation binding the contract event 0x6c69e09ee5c5ac33c0cd57787261c5bade070a392ab34a4b5487c6868f723f6e.
//
// Solidity: event MiningRewardDistributionFailed(address indexed consensusAddr, address indexed recipient, uint256 amount, uint256 contractBalance)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseMiningRewardDistributionFailed(log types.Log) (*RoninValidatorSetMiningRewardDistributionFailed, error) {
	event := new(RoninValidatorSetMiningRewardDistributionFailed)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "MiningRewardDistributionFailed", log); err != nil {
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

// RoninValidatorSetRoninTrustedOrganizationContractUpdatedIterator is returned from FilterRoninTrustedOrganizationContractUpdated and is used to iterate over the raw logs and unpacked data for RoninTrustedOrganizationContractUpdated events raised by the RoninValidatorSet contract.
type RoninValidatorSetRoninTrustedOrganizationContractUpdatedIterator struct {
	Event *RoninValidatorSetRoninTrustedOrganizationContractUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetRoninTrustedOrganizationContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetRoninTrustedOrganizationContractUpdated)
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
		it.Event = new(RoninValidatorSetRoninTrustedOrganizationContractUpdated)
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
func (it *RoninValidatorSetRoninTrustedOrganizationContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetRoninTrustedOrganizationContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetRoninTrustedOrganizationContractUpdated represents a RoninTrustedOrganizationContractUpdated event raised by the RoninValidatorSet contract.
type RoninValidatorSetRoninTrustedOrganizationContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRoninTrustedOrganizationContractUpdated is a free log retrieval operation binding the contract event 0xfd6f5f93d69a07c593a09be0b208bff13ab4ffd6017df3b33433d63bdc59b4d7.
//
// Solidity: event RoninTrustedOrganizationContractUpdated(address arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterRoninTrustedOrganizationContractUpdated(opts *bind.FilterOpts) (*RoninValidatorSetRoninTrustedOrganizationContractUpdatedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "RoninTrustedOrganizationContractUpdated")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetRoninTrustedOrganizationContractUpdatedIterator{contract: _RoninValidatorSet.contract, event: "RoninTrustedOrganizationContractUpdated", logs: logs, sub: sub}, nil
}

// WatchRoninTrustedOrganizationContractUpdated is a free log subscription operation binding the contract event 0xfd6f5f93d69a07c593a09be0b208bff13ab4ffd6017df3b33433d63bdc59b4d7.
//
// Solidity: event RoninTrustedOrganizationContractUpdated(address arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchRoninTrustedOrganizationContractUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetRoninTrustedOrganizationContractUpdated) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "RoninTrustedOrganizationContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetRoninTrustedOrganizationContractUpdated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "RoninTrustedOrganizationContractUpdated", log); err != nil {
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

// ParseRoninTrustedOrganizationContractUpdated is a log parse operation binding the contract event 0xfd6f5f93d69a07c593a09be0b208bff13ab4ffd6017df3b33433d63bdc59b4d7.
//
// Solidity: event RoninTrustedOrganizationContractUpdated(address arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseRoninTrustedOrganizationContractUpdated(log types.Log) (*RoninValidatorSetRoninTrustedOrganizationContractUpdated, error) {
	event := new(RoninValidatorSetRoninTrustedOrganizationContractUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "RoninTrustedOrganizationContractUpdated", log); err != nil {
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

// RoninValidatorSetStakingRewardDistributionFailedIterator is returned from FilterStakingRewardDistributionFailed and is used to iterate over the raw logs and unpacked data for StakingRewardDistributionFailed events raised by the RoninValidatorSet contract.
type RoninValidatorSetStakingRewardDistributionFailedIterator struct {
	Event *RoninValidatorSetStakingRewardDistributionFailed // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetStakingRewardDistributionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetStakingRewardDistributionFailed)
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
		it.Event = new(RoninValidatorSetStakingRewardDistributionFailed)
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
func (it *RoninValidatorSetStakingRewardDistributionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetStakingRewardDistributionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetStakingRewardDistributionFailed represents a StakingRewardDistributionFailed event raised by the RoninValidatorSet contract.
type RoninValidatorSetStakingRewardDistributionFailed struct {
	Amount          *big.Int
	ContractBalance *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakingRewardDistributionFailed is a free log retrieval operation binding the contract event 0x0752cb1e4b6fb7b2beb1cf423d908acaec7acfb7782e67a88d158351b1c0c4a5.
//
// Solidity: event StakingRewardDistributionFailed(uint256 amount, uint256 contractBalance)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterStakingRewardDistributionFailed(opts *bind.FilterOpts) (*RoninValidatorSetStakingRewardDistributionFailedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "StakingRewardDistributionFailed")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetStakingRewardDistributionFailedIterator{contract: _RoninValidatorSet.contract, event: "StakingRewardDistributionFailed", logs: logs, sub: sub}, nil
}

// WatchStakingRewardDistributionFailed is a free log subscription operation binding the contract event 0x0752cb1e4b6fb7b2beb1cf423d908acaec7acfb7782e67a88d158351b1c0c4a5.
//
// Solidity: event StakingRewardDistributionFailed(uint256 amount, uint256 contractBalance)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchStakingRewardDistributionFailed(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetStakingRewardDistributionFailed) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "StakingRewardDistributionFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetStakingRewardDistributionFailed)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "StakingRewardDistributionFailed", log); err != nil {
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

// ParseStakingRewardDistributionFailed is a log parse operation binding the contract event 0x0752cb1e4b6fb7b2beb1cf423d908acaec7acfb7782e67a88d158351b1c0c4a5.
//
// Solidity: event StakingRewardDistributionFailed(uint256 amount, uint256 contractBalance)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseStakingRewardDistributionFailed(log types.Log) (*RoninValidatorSetStakingRewardDistributionFailed, error) {
	event := new(RoninValidatorSetStakingRewardDistributionFailed)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "StakingRewardDistributionFailed", log); err != nil {
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

// RoninValidatorSetValidatorLiberatedIterator is returned from FilterValidatorLiberated and is used to iterate over the raw logs and unpacked data for ValidatorLiberated events raised by the RoninValidatorSet contract.
type RoninValidatorSetValidatorLiberatedIterator struct {
	Event *RoninValidatorSetValidatorLiberated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetValidatorLiberatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetValidatorLiberated)
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
		it.Event = new(RoninValidatorSetValidatorLiberated)
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
func (it *RoninValidatorSetValidatorLiberatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetValidatorLiberatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetValidatorLiberated represents a ValidatorLiberated event raised by the RoninValidatorSet contract.
type RoninValidatorSetValidatorLiberated struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorLiberated is a free log retrieval operation binding the contract event 0x028097aa2663c9c46d17e72793860e66036c00b17374e44ae5410327322f396e.
//
// Solidity: event ValidatorLiberated(address indexed validator)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterValidatorLiberated(opts *bind.FilterOpts, validator []common.Address) (*RoninValidatorSetValidatorLiberatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "ValidatorLiberated", validatorRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetValidatorLiberatedIterator{contract: _RoninValidatorSet.contract, event: "ValidatorLiberated", logs: logs, sub: sub}, nil
}

// WatchValidatorLiberated is a free log subscription operation binding the contract event 0x028097aa2663c9c46d17e72793860e66036c00b17374e44ae5410327322f396e.
//
// Solidity: event ValidatorLiberated(address indexed validator)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchValidatorLiberated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetValidatorLiberated, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "ValidatorLiberated", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetValidatorLiberated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "ValidatorLiberated", log); err != nil {
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

// ParseValidatorLiberated is a log parse operation binding the contract event 0x028097aa2663c9c46d17e72793860e66036c00b17374e44ae5410327322f396e.
//
// Solidity: event ValidatorLiberated(address indexed validator)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseValidatorLiberated(log types.Log) (*RoninValidatorSetValidatorLiberated, error) {
	event := new(RoninValidatorSetValidatorLiberated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "ValidatorLiberated", log); err != nil {
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
	ConsensusAddr                  common.Address
	Period                         *big.Int
	JailedUntil                    *big.Int
	DeductedStakingAmount          *big.Int
	BlockProducerRewardDeprecated  bool
	BridgeOperatorRewardDeprecated bool
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterValidatorPunished is a free log retrieval operation binding the contract event 0x54ce99c5ce1fc9f61656d4a0fb2697974d0c973ac32eecaefe06fcf18b8ef68a.
//
// Solidity: event ValidatorPunished(address indexed consensusAddr, uint256 indexed period, uint256 jailedUntil, uint256 deductedStakingAmount, bool blockProducerRewardDeprecated, bool bridgeOperatorRewardDeprecated)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterValidatorPunished(opts *bind.FilterOpts, consensusAddr []common.Address, period []*big.Int) (*RoninValidatorSetValidatorPunishedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "ValidatorPunished", consensusAddrRule, periodRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetValidatorPunishedIterator{contract: _RoninValidatorSet.contract, event: "ValidatorPunished", logs: logs, sub: sub}, nil
}

// WatchValidatorPunished is a free log subscription operation binding the contract event 0x54ce99c5ce1fc9f61656d4a0fb2697974d0c973ac32eecaefe06fcf18b8ef68a.
//
// Solidity: event ValidatorPunished(address indexed consensusAddr, uint256 indexed period, uint256 jailedUntil, uint256 deductedStakingAmount, bool blockProducerRewardDeprecated, bool bridgeOperatorRewardDeprecated)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchValidatorPunished(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetValidatorPunished, consensusAddr []common.Address, period []*big.Int) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "ValidatorPunished", consensusAddrRule, periodRule)
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

// ParseValidatorPunished is a log parse operation binding the contract event 0x54ce99c5ce1fc9f61656d4a0fb2697974d0c973ac32eecaefe06fcf18b8ef68a.
//
// Solidity: event ValidatorPunished(address indexed consensusAddr, uint256 indexed period, uint256 jailedUntil, uint256 deductedStakingAmount, bool blockProducerRewardDeprecated, bool bridgeOperatorRewardDeprecated)
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
	Period         *big.Int
	ConsensusAddrs []common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetUpdated is a free log retrieval operation binding the contract event 0x3d0eea40644a206ec25781dd5bb3b60eb4fa1264b993c3bddf3c73b14f29ef5e.
//
// Solidity: event ValidatorSetUpdated(uint256 indexed period, address[] consensusAddrs)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterValidatorSetUpdated(opts *bind.FilterOpts, period []*big.Int) (*RoninValidatorSetValidatorSetUpdatedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "ValidatorSetUpdated", periodRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetValidatorSetUpdatedIterator{contract: _RoninValidatorSet.contract, event: "ValidatorSetUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorSetUpdated is a free log subscription operation binding the contract event 0x3d0eea40644a206ec25781dd5bb3b60eb4fa1264b993c3bddf3c73b14f29ef5e.
//
// Solidity: event ValidatorSetUpdated(uint256 indexed period, address[] consensusAddrs)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchValidatorSetUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetValidatorSetUpdated, period []*big.Int) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "ValidatorSetUpdated", periodRule)
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

// ParseValidatorSetUpdated is a log parse operation binding the contract event 0x3d0eea40644a206ec25781dd5bb3b60eb4fa1264b993c3bddf3c73b14f29ef5e.
//
// Solidity: event ValidatorSetUpdated(uint256 indexed period, address[] consensusAddrs)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseValidatorSetUpdated(log types.Log) (*RoninValidatorSetValidatorSetUpdated, error) {
	event := new(RoninValidatorSetValidatorSetUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "ValidatorSetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetWrappedUpEpochIterator is returned from FilterWrappedUpEpoch and is used to iterate over the raw logs and unpacked data for WrappedUpEpoch events raised by the RoninValidatorSet contract.
type RoninValidatorSetWrappedUpEpochIterator struct {
	Event *RoninValidatorSetWrappedUpEpoch // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetWrappedUpEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetWrappedUpEpoch)
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
		it.Event = new(RoninValidatorSetWrappedUpEpoch)
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
func (it *RoninValidatorSetWrappedUpEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetWrappedUpEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetWrappedUpEpoch represents a WrappedUpEpoch event raised by the RoninValidatorSet contract.
type RoninValidatorSetWrappedUpEpoch struct {
	PeriodNumber *big.Int
	EpochNumber  *big.Int
	PeriodEnding bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWrappedUpEpoch is a free log retrieval operation binding the contract event 0x0195462033384fec211477c56217da64a58bd405e0bed331ba4ded67e4ae4ce7.
//
// Solidity: event WrappedUpEpoch(uint256 indexed periodNumber, uint256 indexed epochNumber, bool periodEnding)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterWrappedUpEpoch(opts *bind.FilterOpts, periodNumber []*big.Int, epochNumber []*big.Int) (*RoninValidatorSetWrappedUpEpochIterator, error) {

	var periodNumberRule []interface{}
	for _, periodNumberItem := range periodNumber {
		periodNumberRule = append(periodNumberRule, periodNumberItem)
	}
	var epochNumberRule []interface{}
	for _, epochNumberItem := range epochNumber {
		epochNumberRule = append(epochNumberRule, epochNumberItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "WrappedUpEpoch", periodNumberRule, epochNumberRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetWrappedUpEpochIterator{contract: _RoninValidatorSet.contract, event: "WrappedUpEpoch", logs: logs, sub: sub}, nil
}

// WatchWrappedUpEpoch is a free log subscription operation binding the contract event 0x0195462033384fec211477c56217da64a58bd405e0bed331ba4ded67e4ae4ce7.
//
// Solidity: event WrappedUpEpoch(uint256 indexed periodNumber, uint256 indexed epochNumber, bool periodEnding)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchWrappedUpEpoch(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetWrappedUpEpoch, periodNumber []*big.Int, epochNumber []*big.Int) (event.Subscription, error) {

	var periodNumberRule []interface{}
	for _, periodNumberItem := range periodNumber {
		periodNumberRule = append(periodNumberRule, periodNumberItem)
	}
	var epochNumberRule []interface{}
	for _, epochNumberItem := range epochNumber {
		epochNumberRule = append(epochNumberRule, epochNumberItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "WrappedUpEpoch", periodNumberRule, epochNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetWrappedUpEpoch)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "WrappedUpEpoch", log); err != nil {
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

// ParseWrappedUpEpoch is a log parse operation binding the contract event 0x0195462033384fec211477c56217da64a58bd405e0bed331ba4ded67e4ae4ce7.
//
// Solidity: event WrappedUpEpoch(uint256 indexed periodNumber, uint256 indexed epochNumber, bool periodEnding)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseWrappedUpEpoch(log types.Log) (*RoninValidatorSetWrappedUpEpoch, error) {
	event := new(RoninValidatorSetWrappedUpEpoch)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "WrappedUpEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
