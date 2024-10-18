package core

import (
	"errors"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
)

const (
	testnetChainId = 2021
)

type blacklistedAddress struct{}

func (c *blacklistedAddress) RequiredGas(input []byte) uint64 {
	return params.CallValueTransferGas + params.SloadGas*2
}

func (c *blacklistedAddress) Run(input []byte) ([]byte, error) {
	return nil, errors.New("address is blacklisted")
}

type TestnetHook struct{}

func (h TestnetHook) PrecompileHook(evm *vm.EVM, addr common.Address) (bool, vm.PrecompiledContract, bool) {
	var (
		badBlocks = []int64{
			3312683,
			3315192,
			3315194,
			3315196,
			3315198,
		}
		blacklistAddress = common.HexToAddress("0x3a3012ec6812a00811A435c47eD79346318c8C92")
	)

	blockNumber := evm.Context.BlockNumber.Int64()
	index := sort.Search(len(badBlocks), func(i int) bool { return badBlocks[i] >= blockNumber })

	if index < len(badBlocks) && badBlocks[index] == blockNumber {
		if blacklistAddress == addr {
			return true, &blacklistedAddress{}, true
		}
	}

	return false, nil, false
}

func (h TestnetHook) CreateHook(evm *vm.EVM, gas uint64, addr common.Address) (bool, []byte, common.Address, uint64, error) {
	var (
		badBlocks = []int64{
			3281085,
			3283414,
			3283446,
			3283452,
			3283454,
			3283456,
			3283458,
			3284174,
			3284176,
			3284178,
			3284180,
			3284188,
			3284190,
			3284275,
			3284277,
			3284279,
			3284281,
			3284307,
			3284309,
			3284311,
			3284313,
			3284617,
			3284619,
			3284621,
			3284623,
			3284752,
			3284754,
			3284756,
			3284758,
			3285191,
			3285193,
			3285195,
			3285197,
			3285607,
			3285609,
			3285611,
			3285613,
		}

		blacklistAddress = common.HexToAddress("0x4e59b44847b379578588920ca78fbf26c0b4956c")
	)

	blockNumber := evm.Context.BlockNumber.Int64()
	index := sort.Search(len(badBlocks), func(i int) bool { return badBlocks[i] >= blockNumber })

	if index < len(badBlocks) && badBlocks[index] == blockNumber {
		if addr == blacklistAddress {
			return true, nil, common.Address{}, gas, vm.ErrExecutionReverted
		}
	}

	return false, nil, common.Address{}, 0, nil
}
