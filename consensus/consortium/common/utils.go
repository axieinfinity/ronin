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

// NormalizeFinalityVoteWeight returns the finality vote weights based on staked amounts.
//
// Algorithm (assuming threshold = 22)
// 1. Sort stakedAmounts in descending order
// 2. Init: total = sum(stakedAmounts), threshold = total * 1 / 22
// 3. Loop through stakedAmounts: stakeAmounts[i] > threshold => stakeAmount[i] = threshold
// 4. If no change, break
// 5. total = sum(unchanged(stakedAmounts)), threshold = total * 1 / (22 - num(changed(stakedAmounts))), go to step 3
//
// Note: the stakedAmounts may be changed inside this function
func NormalizeFinalityVoteWeight(stakedAmounts []*big.Int, threshold int) []uint16 {
	weights := make([]uint16, 0, len(stakedAmounts))

	// The candidate list is too small, so weight is equal among the candidates
	if len(stakedAmounts) <= threshold {
		for range stakedAmounts {
			weights = append(weights, MaxFinalityVotePercentage/uint16(len(stakedAmounts)))
		}

		return weights
	}

	cpyStakedAmounts := make([]*big.Int, len(stakedAmounts))
	for i, stakedAmount := range stakedAmounts {
		cpyStakedAmounts[i] = new(big.Int).Set(stakedAmount)
	}

	// Sort staked amount in descending order
	for i := 0; i < len(cpyStakedAmounts)-1; i++ {
		for j := i + 1; j < len(cpyStakedAmounts); j++ {
			if cpyStakedAmounts[i].Cmp(cpyStakedAmounts[j]) < 0 {
				cpyStakedAmounts[i], cpyStakedAmounts[j] = cpyStakedAmounts[j], cpyStakedAmounts[i]
			}
		}
	}

	totalStakedAmount := new(big.Int)
	for _, stakedAmount := range cpyStakedAmounts {
		totalStakedAmount.Add(totalStakedAmount, stakedAmount)
	}
	weightThreshold := new(big.Int).Div(totalStakedAmount, big.NewInt(int64(threshold)))

	pointer := 0
	sumOfUnchangedElements := totalStakedAmount
	for {
		sumOfChangedElements := new(big.Int)
		shouldBreak := true
		for cpyStakedAmounts[pointer].Cmp(weightThreshold) > 0 {
			sumOfChangedElements.Add(sumOfChangedElements, cpyStakedAmounts[pointer])
			shouldBreak = false
			pointer++
		}

		if shouldBreak {
			break
		}

		sumOfUnchangedElements = new(big.Int).Sub(sumOfUnchangedElements, sumOfChangedElements)
		weightThreshold = new(big.Int).Div(
			sumOfUnchangedElements,
			new(big.Int).Sub(big.NewInt(int64(threshold)), big.NewInt(int64(pointer))),
		)
	}

	for i, stakedAmount := range stakedAmounts {
		if stakedAmount.Cmp(weightThreshold) > 0 {
			stakedAmounts[i] = weightThreshold
		}
	}

	totalStakedAmount.SetUint64(0)
	for _, stakedAmount := range stakedAmounts {
		totalStakedAmount.Add(totalStakedAmount, stakedAmount)
	}

	for _, stakedAmount := range stakedAmounts {
		weight := new(big.Int).Mul(stakedAmount, big.NewInt(int64(MaxFinalityVotePercentage)))
		weight.Div(weight, totalStakedAmount)

		weights = append(weights, uint16(weight.Uint64()))
	}

	// Due to the imprecision of division, the remaining weight for the total to reach 100% is
	// split equally across cnadidates. After this step, the total weight may still not reach
	// 100% but the imprecision is neglectible (lower than the length of candidate list)
	var totalFinalityWeight uint16
	for _, weight := range weights {
		totalFinalityWeight += weight
	}
	cutOffWeight := MaxFinalityVotePercentage - totalFinalityWeight
	topUpWeight := cutOffWeight / uint16(len(weights))
	for i := range weights {
		weights[i] += topUpWeight
	}

	return weights
}
