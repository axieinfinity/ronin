package state

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

const (
	WHITELISTED   = "whitelisted"
	WHITELIST_ALL = "whitelistAll"
	BLACKLISTED   = "_blacklisted"
	DISABLED      = "disabled"
	VALIDATORS    = "validators"
)

var (
	slotWhitelistDeployerMapping = map[string]uint64{
		WHITELISTED:   1,
		WHITELIST_ALL: 2,
	}
	slotWhitelistDeployerMappingV2 = map[string]uint64{
		WHITELISTED:   53,
		WHITELIST_ALL: 58,
	} // Contract Infinity
	slotBlacklistContractMapping = map[string]uint64{
		BLACKLISTED: 1,
		DISABLED:    2,
	}
	slotSCValidatorMapping = map[string]uint64{
		VALIDATORS: 1,
	}
	slotRoninValidatorMapping = map[string]uint64{
		VALIDATORS: 6,
	}
	valueOne common.Hash
)

func init() {
	valueOne.SetBytes([]byte{0x1})
}

// IsWhitelistedDeployer reads the contract storage to check if an address is allow to deploy
func IsWhitelistedDeployerV2(statedb *StateDB, address common.Address, blockTime uint64, whiteListContract *common.Address) bool {
	contract := *whiteListContract
	whitelistAllSlot := slotWhitelistDeployerMappingV2[WHITELIST_ALL]
	whitelistAll := statedb.GetState(contract, GetLocSimpleVariable(whitelistAllSlot))

	if whitelistAll.Big().Cmp(common.Big1) == 0 {
		return true
	}

	whitelistedSlot := slotWhitelistDeployerMappingV2[WHITELISTED]
	// WhiteListInfo have 2 fields, so we need to plus 1.
	// struct WhiteListInfo {
	// 	uint256 expiryTimestamp;
	// 	bool activated;
	//   }
	expiredLoc := GetLocMappingAtKey(address.Hash(), whitelistedSlot)
	activatedLoc := common.BigToHash(expiredLoc.Big().Add(expiredLoc.Big(), common.Big1))
	expiredHash := statedb.GetState(contract, expiredLoc)

	activatedHash := statedb.GetState(contract, activatedLoc)

	// (whiteListInfo.activated && block.timestamp < whiteListInfo.expiryTimestamp)
	// Compare expiredHash with Blockheader timestamp.
	if activatedHash.Big().Cmp(common.Big1) == 0 {
		if expiredHash.Big().Cmp(big.NewInt(int64(blockTime))) > 0 {
			// Block time still is in expiredTime
			return true
		}
	}
	return false
}

// IsWhitelistedDeployer reads the contract storage to check if an address is allow to deploy
func IsWhitelistedDeployer(statedb *StateDB, address common.Address) bool {
	contract := common.HexToAddress(common.WhitelistDeployerSC)
	whitelistAllSlot := slotWhitelistDeployerMapping[WHITELIST_ALL]
	whitelistAll := statedb.GetState(contract, GetLocSimpleVariable(whitelistAllSlot))

	if whitelistAll.Big().Cmp(big.NewInt(1)) == 0 {
		return true
	}

	whitelistedSlot := slotWhitelistDeployerMapping[WHITELISTED]
	valueLoc := GetLocMappingAtKey(address.Hash(), whitelistedSlot)
	whitelisted := statedb.GetState(contract, valueLoc)

	return whitelisted.Big().Cmp(big.NewInt(1)) == 0
}

// IsAddressBlacklisted reads the contract storage to check if an address is blacklisted or not
func IsAddressBlacklisted(statedb *StateDB, blacklistAddr *common.Address, address *common.Address) bool {
	if blacklistAddr == nil || address == nil {
		return false
	}

	contract := *blacklistAddr
	disabledSlot := slotBlacklistContractMapping[DISABLED]
	disabledStateValue := statedb.GetState(contract, GetLocSimpleVariable(disabledSlot))

	if disabledStateValue == valueOne {
		return false
	}

	blacklistedSlot := slotBlacklistContractMapping[BLACKLISTED]
	valueLoc := GetLocMappingAtKey(address.Hash(), blacklistedSlot)
	blacklistedStateValue := statedb.GetState(contract, valueLoc)
	return blacklistedStateValue == valueOne
}

func GetSCValidators(statedb *StateDB) []common.Address {
	slot := slotSCValidatorMapping[VALIDATORS]
	slotHash := common.BigToHash(new(big.Int).SetUint64(slot))
	arrLength := statedb.GetState(common.HexToAddress(common.ValidatorSC), slotHash)
	keys := []common.Hash{}
	for i := uint64(0); i < arrLength.Big().Uint64(); i++ {
		key := GetLocDynamicArrAtElement(slotHash, i, 1)
		keys = append(keys, key)
	}
	rets := []common.Address{}
	for _, key := range keys {
		ret := statedb.GetState(common.HexToAddress(common.ValidatorSC), key)
		rets = append(rets, common.HexToAddress(ret.Hex()))
	}
	return rets
}

func GetFenixValidators(statedb *StateDB, roninValidatorContract *common.Address) []common.Address {
	if roninValidatorContract == nil {
		log.Crit("Cannot get Ronin Validator contract")
		return GetSCValidators(statedb)
	}

	slot := slotRoninValidatorMapping[VALIDATORS]
	slotHash := common.BigToHash(new(big.Int).SetUint64(slot))
	arrLength := statedb.GetState(*roninValidatorContract, slotHash)
	keys := []common.Hash{}
	for i := uint64(0); i < arrLength.Big().Uint64(); i++ {
		key := GetLocDynamicArrAtElement(slotHash, i, 1)
		keys = append(keys, key)
	}
	rets := []common.Address{}
	for _, key := range keys {
		ret := statedb.GetState(*roninValidatorContract, key)
		rets = append(rets, common.HexToAddress(ret.Hex()))
	}
	return rets
}
