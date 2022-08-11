package systemcontracts

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
)

type UpgradeConfig struct {
	BeforeUpgrade upgradeHook
	AfterUpgrade  upgradeHook
	ContractAddr  common.Address
	CommitUrl     string
	Code          string
}

type Upgrade struct {
	UpgradeName string
	Configs     []*UpgradeConfig
}

type upgradeHook func(blockNumber *big.Int, contractAddr common.Address, statedb *state.StateDB) error

const (
	mainnet = "mainnet"
	testnet = "testnet"
)

var (
	GenesisHash common.Hash
	//upgrade config
	consortiumV2 = make(map[string]*Upgrade)
)

func init() {
	consortiumV2[mainnet] = &Upgrade{
		UpgradeName: "consortium v2",
		Configs: []*UpgradeConfig{
			{
				ContractAddr: common.HexToAddress(ValidatorContract),
				CommitUrl:    "",
				Code:         ConsortiumV2ValidatorContractCode,
			},
			{
				ContractAddr: common.HexToAddress(SlashContract),
				CommitUrl:    "",
				Code:         ConsortiumV2SlashContractCode,
			},
		},
	}
	consortiumV2[testnet] = consortiumV2[mainnet]
}

func UpgradeBuildInSystemContract(config *params.ChainConfig, blockNumber *big.Int, statedb *state.StateDB) {
	if config == nil || blockNumber == nil || statedb == nil {
		return
	}
	var network string
	switch GenesisHash {
	/* Add mainnet genesis hash */
	case params.RoninMainnetGenesisHash:
		network = mainnet
	default:
		network = testnet
	}

	if config.IsOnConsortiumV2(blockNumber) {
		applySystemContractUpgrade(consortiumV2[network], blockNumber, statedb)
	}

}

func applySystemContractUpgrade(upgrade *Upgrade, blockNumber *big.Int, statedb *state.StateDB) {
	if upgrade == nil {
		log.Info("Empty upgrade config", "height", blockNumber.String())
		return
	}

	log.Info(fmt.Sprintf("Apply upgrade %s at height %d", upgrade.UpgradeName, blockNumber.Int64()))
	for _, cfg := range upgrade.Configs {
		log.Info(fmt.Sprintf("Upgrade contract %s to commit %s", cfg.ContractAddr.String(), cfg.CommitUrl))

		if cfg.BeforeUpgrade != nil {
			if err := cfg.BeforeUpgrade(blockNumber, cfg.ContractAddr, statedb); err != nil {
				panic(fmt.Errorf("contract address: %s, execute beforeUpgrade error: %s", cfg.ContractAddr.String(), err.Error()))
			}
		}

		newContractCode, err := hex.DecodeString(cfg.Code)
		if err != nil {
			panic(fmt.Errorf("failed to decode new contract code: %s", err.Error()))
		}
		statedb.SetCode(cfg.ContractAddr, newContractCode)

		if cfg.AfterUpgrade != nil {
			if err := cfg.AfterUpgrade(blockNumber, cfg.ContractAddr, statedb); err != nil {
				panic(fmt.Errorf("contract address: %s, execute afterUpgrade error: %s", cfg.ContractAddr.String(), err.Error()))
			}
		}
	}
}
