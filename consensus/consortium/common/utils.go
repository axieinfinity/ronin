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

//  1. The vote weight of each validator candidate is validator pool's staked amount / total staked of all validator's pools.
//  2. If the vote weight of a validator candidate is higher than 1 / 22, then the vote weight is 1 / 22.
//  3. The total vote weight may not equal to 100%, the cut off weight is split across remaining candidate based on their
//     weight / remaining total weight. Repeatedly do this until the cut off weight cannot be split anymore.
//  4. After step 1, some candidates may get 0 weight and cannot receive cut off weight from step 3. Split the remaining cut
//     off weight among 0 weight candidates.
//  5. Due to the imprecision of division, the remaining weight for the total to reach 100% is split equally across cnadidates.
//     After this step, the total weight may still not reach 100% but the imprecision is neglectible (lower than the length of
//     candidate list)
func NormalizeFinalityVoteWeight(stakedAmounts []*big.Int) []uint16 {
	var (
		finalityVoteWeight = make([]uint16, len(stakedAmounts))
		totalStakedAmount  = big.NewInt(0)
		weightTheshold     = MaxFinalityVotePercentage / VoteWeightThreshold
		cutOffWeight       uint16
		numOfZeroWeight    uint16
	)

	for _, stakedAmount := range stakedAmounts {
		totalStakedAmount.Add(totalStakedAmount, stakedAmount)
	}

	// The candidate list is too small, so weight is equal among the candidates
	if len(stakedAmounts) <= VoteWeightThreshold {
		for i := range finalityVoteWeight {
			finalityVoteWeight[i] = MaxFinalityVotePercentage / uint16(len(finalityVoteWeight))
		}

		return finalityVoteWeight
	}

	// Step 1, 2
	for i, stakedAmount := range stakedAmounts {
		weight := new(big.Int).Mul(stakedAmount, big.NewInt(int64(MaxFinalityVotePercentage)))
		weight.Div(weight, totalStakedAmount)

		w := uint16(weight.Uint64())
		if w == 0 {
			numOfZeroWeight++
		} else if w > weightTheshold {
			cutOffWeight += w - weightTheshold
			w = weightTheshold
		}

		finalityVoteWeight[i] = w
	}

	// Step 3
	for cutOffWeight > 0 {
		prevCutOffWeight := cutOffWeight

		var totalRemainingWeight uint16
		for _, weight := range finalityVoteWeight {
			if weight < weightTheshold {
				totalRemainingWeight += weight
			}
		}

		if totalRemainingWeight > 0 {
			for i, weight := range finalityVoteWeight {
				if weight < weightTheshold {
					topUpWeight := uint16(uint64(weight) * uint64(prevCutOffWeight) / uint64(totalRemainingWeight))

					if finalityVoteWeight[i]+topUpWeight > weightTheshold {
						topUpWeight = weightTheshold - finalityVoteWeight[i]
					}

					cutOffWeight -= topUpWeight
					finalityVoteWeight[i] += topUpWeight
				}
			}
		}

		if cutOffWeight == prevCutOffWeight {
			break
		}
	}

	// Step 4
	if cutOffWeight != 0 && numOfZeroWeight != 0 {
		topUpWeight := cutOffWeight / numOfZeroWeight
		for i, weight := range finalityVoteWeight {
			if weight == 0 {
				finalityVoteWeight[i] += topUpWeight
			}
		}
	}

	// Step 5
	var totalFinalityWeight uint16
	for _, weight := range finalityVoteWeight {
		totalFinalityWeight += weight
	}
	cutOffWeight = MaxFinalityVotePercentage - totalFinalityWeight
	topUpWeight := cutOffWeight / uint16(len(finalityVoteWeight))
	for i := range finalityVoteWeight {
		finalityVoteWeight[i] += topUpWeight
	}

	return finalityVoteWeight
}
