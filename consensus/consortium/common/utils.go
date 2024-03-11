package common

import (
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/common"
)

// ExtractAddressFromBytes extracts validators' address from extra data in header
// and return a list addresses
func ExtractAddressFromBytes(bytes []byte) []common.Address {
	if bytes != nil && len(bytes) < common.AddressLength {
		return []common.Address{}
	}
	results := make([]common.Address, len(bytes)/common.AddressLength)
	for i := 0; i < len(results); i++ {
		copy(results[i][:], bytes[i*common.AddressLength:])
	}
	return results
}

// CompareSignersLists compares 2 signers lists
// return true if they are same elements, otherwise return false
func CompareSignersLists(list1 []common.Address, list2 []common.Address) bool {
	if len(list1) != len(list2) {
		return false
	}
	for i := 0; i < len(list1); i++ {
		if list1[i].Hex() != list2[i].Hex() {
			return false
		}
	}
	return true
}

// SignerInList checks the given signer address is whether in the validators set or not
func SignerInList(signer common.Address, validators []common.Address) bool {
	for _, validator := range validators {
		if signer == validator {
			return true
		}
	}
	return false
}

// RemoveOutdatedRecents removes outdated recents list
func RemoveOutdatedRecents(recents map[uint64]common.Address, currentBlock uint64) map[uint64]common.Address {
	var blocks []uint64
	for n, _ := range recents {
		blocks = append(blocks, n)
	}
	sort.Slice(blocks, func(i, j int) bool { return blocks[i] > blocks[j] })

	newRecents := make(map[uint64]common.Address)
	for _, n := range blocks {
		if currentBlock == n {
			newRecents[n] = recents[n]
		}
		currentBlock -= 1
	}

	return newRecents
}

// 1. The vote weight of each validator is validator pool's staked amount / total staked of all validator's pools
// 2. If the vote weight of a validator is higher than 1 / n, then the vote weight is 1 / n with n is the number
// of validators
// 3. After the step 2, the total vote weight might be lower than 1. Normalize the vote weight to make total vote
// weight is 1 (new vote weight = current vote weight / current total vote weight) (after this step, the total vote
// weight might not be 1 due to precision problem, but it is neglectible with small n)
//
// For vote weight, we don't use floating pointer number but multiply the vote weight with MaxFinalityVotePercentage
// and store vote weight in integer type. The precision of calculation is based on MaxFinalityVotePercentage.
func NormalizeFinalityVoteWeight(stakedAmounts []*big.Int) []uint16 {
	var (
		totalStakedAmount  = big.NewInt(0)
		finalityVoteWeight []uint16
		maxVoteWeight      uint16
		totalVoteWeight    uint
	)

	// Calculate the maximum vote weight of each validator for step 2
	// 1 * MaxFinalityVotePercentage / n
	maxVoteWeight = MaxFinalityVotePercentage / uint16(len(stakedAmounts))

	for _, stakedAmount := range stakedAmounts {
		totalStakedAmount.Add(totalStakedAmount, stakedAmount)
	}

	// Step 1, 2
	for _, stakedAmount := range stakedAmounts {
		weight := new(big.Int).Mul(stakedAmount, big.NewInt(int64(MaxFinalityVotePercentage)))
		weight.Div(weight, totalStakedAmount)

		w := uint16(weight.Uint64())
		if w > maxVoteWeight {
			w = maxVoteWeight
		}
		totalVoteWeight += uint(w)
		finalityVoteWeight = append(finalityVoteWeight, w)
	}

	// Step 3
	for i, weight := range finalityVoteWeight {
		normalizedWeight := uint16(uint(weight) * uint(MaxFinalityVotePercentage) / totalVoteWeight)
		finalityVoteWeight[i] = normalizedWeight
	}

	return finalityVoteWeight
}
