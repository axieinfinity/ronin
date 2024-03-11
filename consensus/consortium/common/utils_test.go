package common

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestExtractAddressFromBytes(t *testing.T) {
	sampleAddresses := []common.Address{
		common.HexToAddress("0x93b8eed0a1e082ae2f478fd7f8c14b1fc0261bb1"),
		common.HexToAddress("0x93b8eed0a1e082ae2f478fd7f8c14b1fc0261bb1"),
		common.HexToAddress("0x93b8eed0a1e082ae2f478fd7f8c14b1fc0261bb1"),
	}
	sampleBytes := make([]byte, 0, common.AddressLength*len(sampleAddresses))
	for _, address := range sampleAddresses {
		sampleBytes = append(sampleBytes, address.Bytes()...)
	}

	result := ExtractAddressFromBytes(sampleBytes)

	if !reflect.DeepEqual(result, sampleAddresses) {
		t.Errorf("Output %q not equal to expected %q", result, sampleAddresses)
	}
}

func TestCompareSignersLists(t *testing.T) {
	a := []common.Address{
		common.HexToAddress("0x93b8eed0a1e082ae2f478fd7f8c14b1fc0261bb1"),
		common.HexToAddress("0x93b8eed0a1e082ae2f478fd7f8c14b1fc0261bb1"),
		common.HexToAddress("0x93b8eed0a1e082ae2f478fd7f8c14b1fc0261bb1"),
	}
	b := []common.Address{
		common.HexToAddress("0x93b8eed0a1e082ae2f478fd7f8c14b1fc0261bb1"),
		common.HexToAddress("0x93b8eed0a1e082ae2f478fd7f8c14b1fc0261bb1"),
		common.HexToAddress("0x93b8eed0a1e082ae2f478fd7f8c14b1fc0261bb1"),
	}
	c := []common.Address{
		common.HexToAddress("0x93b8eed0a1e082ae2f478fd7f8c14b1fc0261bb1"),
	}

	if CompareSignersLists(a, b) == false {
		t.Errorf("Output %t not equal to expected %t", false, true)
	}

	if CompareSignersLists(a, c) == true {
		t.Errorf("Output %t not equal to expected %t", false, true)
	}
}

func TestRemoveInvalidRecents(t *testing.T) {
	recents := map[uint64]common.Address{
		7302958: common.HexToAddress("0xAfB9554299491a34d303f2C5A91bebB162f6B2Cf"),
		7305269: common.HexToAddress("0x3B9F2587d55E96276B09b258ac909D809961F6C2"),
		7408557: common.HexToAddress("0xB6bc5bc0410773A3F86B1537ce7495C52e38f88B"),
		7408558: common.HexToAddress("0x3B9F2587d55E96276B09b258ac909D809961F6C2"),
	}
	actual := RemoveOutdatedRecents(recents, 7408558)
	expected := map[uint64]common.Address{
		7408557: common.HexToAddress("0xB6bc5bc0410773A3F86B1537ce7495C52e38f88B"),
		7408558: common.HexToAddress("0x3B9F2587d55E96276B09b258ac909D809961F6C2"),
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func TestNormalizeFinalityVoteWeight(t *testing.T) {
	// All staked amounts are equal
	var stakedAmounts []*big.Int
	for i := 0; i < 22; i++ {
		stakedAmounts = append(stakedAmounts, big.NewInt(1_000_000))
	}

	voteWeights := NormalizeFinalityVoteWeight(stakedAmounts)
	for _, voteWeight := range voteWeights {
		if voteWeight != 454 {
			t.Fatalf("Incorrect vote weight, expect %d got %d", 454, voteWeight)
		}
	}

	// All staked amount differs
	for i := 0; i < 22; i++ {
		stakedAmounts[i] = big.NewInt(int64(i) + 1)
	}
	voteWeights = NormalizeFinalityVoteWeight(stakedAmounts)
	expectedVoteWeights := []uint16{51, 103, 155, 207, 259, 311, 363, 415, 467, 519, 571, 597, 597, 597, 597, 597, 597, 597, 597, 597, 597, 597}

	for i := range voteWeights {
		if voteWeights[i] != expectedVoteWeights[i] {
			t.Fatalf("Incorrect vote weight, expect %d got %d", expectedVoteWeights[i], voteWeights[i])
		}
	}

	// Staked amount differences are small
	for i := 0; i < 22; i++ {
		stakedAmounts[i] = big.NewInt(int64(i) + 1_000_000)
	}
	voteWeights = NormalizeFinalityVoteWeight(stakedAmounts)
	for i := range voteWeights {
		if voteWeights[i] != 454 {
			t.Fatalf("Incorrect vote weight, expect %d got %d", 454, voteWeights[i])
		}
	}

	// Some staked amounts differ greatly
	for i := 0; i < 20; i++ {
		stakedAmounts[i] = big.NewInt(1_000_000)
	}
	stakedAmounts[20] = big.NewInt(1000)
	stakedAmounts[21] = big.NewInt(2500)
	voteWeights = NormalizeFinalityVoteWeight(stakedAmounts)
	for i := 0; i < 20; i++ {
		if voteWeights[i] != 499 {
			t.Fatalf("Incorrect vote weight, expect %d got %d", 499, voteWeights[i])
		}
	}
	if voteWeights[20] != 0 {
		t.Fatalf("Incorrect vote weight, expect %d got %d", 0, voteWeights[20])
	}
	if voteWeights[21] != 1 {
		t.Fatalf("Incorrect vote weight, expect %d got %d", 1, voteWeights[21])
	}
}
