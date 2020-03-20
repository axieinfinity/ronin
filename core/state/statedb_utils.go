package state

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var (
	slotWhitelistDeployerMapping = map[string]uint64{
		"whitelisted":  1,
		"whitelistAll": 2,
	}
)

// IsWhitelistedDeployer reads the contract storage to check if an address is allow to deploy
func IsWhitelistedDeployer(statedb *StateDB, address common.Address) bool {
	contract := common.HexToAddress(common.WhitelistDeployerSC)
	whitelistAllSlot := slotWhitelistDeployerMapping["whitelistAll"]
	whitelistAll := statedb.GetState(contract, GetLocSimpleVariable(whitelistAllSlot))
	if whitelistAll.Big().Cmp(big.NewInt(1)) == 0 {
		return true
	}

	whitelistedSlot := slotWhitelistDeployerMapping["whitelisted"]
	valueLoc := GetLocMappingAtKey(address.Hash(), whitelistedSlot)
	whitelisted := statedb.GetState(contract, valueLoc)

	return whitelisted.Big().Cmp(big.NewInt(1)) == 0
}
