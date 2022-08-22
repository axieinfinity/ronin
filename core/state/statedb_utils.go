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
	GIFTTICKET    = "giftTicket"
)

var (
	slotWhitelistDeployerMapping = map[string]uint64{
		WHITELISTED:   1,
		WHITELIST_ALL: 2,
	}
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
	slotGiftTicketMapping = map[string]uint64{
		GIFTTICKET: 3,
	}
)

// IsWhitelistedDeployer reads the contract storage to check if an address is allow to deploy
func IsWhitelistedDeployer(statedb *StateDB, address common.Address) bool {
	contract := common.HexToAddress(common.WhitelistDeployerSC)
	whitelistAllSlot := slotWhitelistDeployerMapping[WHITELIST_ALL]
	whitelistAll := statedb.GetState(contract, GetLocSimpleVariable(new(big.Int).SetUint64(whitelistAllSlot)))
	if whitelistAll.Big().Cmp(big.NewInt(1)) == 0 {
		return true
	}

	whitelistedSlot := slotWhitelistDeployerMapping[WHITELISTED]
	valueLoc := GetLocMappingAtKey(address.Hash(), new(big.Int).SetUint64(whitelistedSlot))
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
	disabled := statedb.GetState(contract, GetLocSimpleVariable(new(big.Int).SetUint64(disabledSlot)))
	if disabled.Big().Cmp(big.NewInt(1)) == 0 {
		return false
	}

	blacklistedSlot := slotBlacklistContractMapping[BLACKLISTED]
	valueLoc := GetLocMappingAtKey(address.Hash(), new(big.Int).SetUint64(blacklistedSlot))
	blacklisted := statedb.GetState(contract, valueLoc)
	return blacklisted.Big().Cmp(big.NewInt(1)) == 0
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

// The mapping in transaction pass contract is
// mapping payer => (nonce => ticket's remaining uses)
// mapping(address => mapping(uint256 => uint256))
func IsGiftTicketUsable(statedb *StateDB, transactionPassContract *common.Address, payer *common.Address, nonce *big.Int) bool {
	slot := slotGiftTicketMapping[GIFTTICKET]
	payerMap := GetLocMappingAtKey(payer.Hash(), new(big.Int).SetUint64(slot))
	remainingUseSlot := GetLocMappingAtKey(common.BytesToHash(nonce.Bytes()), payerMap.Big())
	remainingUse := statedb.GetState(*transactionPassContract, remainingUseSlot)

	return remainingUse.Big().Cmp(big.NewInt(0)) == 1
}

func SubGiftTicketRemainingUse(statedb *StateDB, transactionPassContract *common.Address, payer *common.Address, nonce *big.Int) {
	slot := slotGiftTicketMapping[GIFTTICKET]
	payerMap := GetLocMappingAtKey(payer.Hash(), new(big.Int).SetUint64(slot))
	remainingUseSlot := GetLocMappingAtKey(common.BytesToHash(nonce.Bytes()), payerMap.Big())
	remainingUse := statedb.GetState(*transactionPassContract, remainingUseSlot)

	remainingUse = common.BytesToHash(new(big.Int).Sub(remainingUse.Big(), big.NewInt(1)).Bytes())
	statedb.SetState(*transactionPassContract, remainingUseSlot, remainingUse)
}
