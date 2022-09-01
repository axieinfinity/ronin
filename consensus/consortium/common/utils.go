package common

import (
	"github.com/ethereum/go-ethereum/common"
)

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
