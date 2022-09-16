package common

import (
	"github.com/ethereum/go-ethereum/common"
	"sort"
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

func SignerInList(signer common.Address, validators []common.Address) bool {
	for _, validator := range validators {
		if signer == validator {
			return true
		}
	}
	return false
}

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
