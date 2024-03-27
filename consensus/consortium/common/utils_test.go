package common

import (
	"encoding/binary"
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
	tests := []struct {
		input  []int
		output []uint16
	}{
		// fewer or equal to 22 validator candidates
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22},
			[]uint16{454, 454, 454, 454, 454, 454, 454, 454, 454, 454, 454, 454, 454, 454, 454, 454, 454, 454, 454, 454, 454, 454},
		},
		// 23 validator candidates with different staked amounts
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23},
			[]uint16{150, 304, 454, 454, 454, 454, 454, 454, 454, 454, 454, 454, 454, 454, 454, 454, 454, 454, 454, 454, 454, 454, 454},
		},
		// 23 validator candidates with different staked amounts
		{
			[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			[]uint16{434, 434, 434, 434, 434, 434, 434, 434, 434, 434, 434, 434, 434, 434, 434, 434, 434, 434, 434, 434, 434, 434, 434},
		},
		// 23 validator candidates, some have very high staked amounts
		{
			[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 100, 200, 150, 50},
			[]uint16{430, 430, 430, 430, 430, 430, 430, 430, 430, 430, 430, 430, 430, 430, 430, 430, 430, 430, 430, 454, 454, 454, 454},
		},
	}

	for _, test := range tests {
		input := make([]*big.Int, 0, len(test.input))
		for _, in := range test.input {
			input = append(input, big.NewInt(int64(in)))
		}

		output := NormalizeFinalityVoteWeight(input)
		if !reflect.DeepEqual(output, test.output) {
			t.Fatalf("Input %v\nExpected output: %v\nGot: %v\n", test.input, test.output, output)
		}
	}
}

func FuzzNormalizeFinalityVoteWeight(f *testing.F) {
	f.Fuzz(func(t *testing.T, fuzzInput []byte) {
		input := make([]*big.Int, 0, len(fuzzInput))
		if len(input)%8 != 0 {
			return
		}
		for i := 0; i < len(fuzzInput)-7; i += 8 {
			in := binary.LittleEndian.Uint64(fuzzInput[i : i+8])
			if in == 0 {
				return
			}
			input = append(input, new(big.Int).SetUint64(in))
		}
		if len(input) == 0 || len(input) > 64 {
			return
		}

		output := NormalizeFinalityVoteWeight(input)
		totalWeight := uint16(0)
		for _, out := range output {
			if len(input) > VoteWeightThreshold {
				if out > MaxFinalityVotePercentage/VoteWeightThreshold+1 {
					t.Fatalf("Weight is higher than 1/22\nInput: %v\nOutput: %v\n", input, output)
				}
			}
			totalWeight += out
		}

		if totalWeight > MaxFinalityVotePercentage {
			t.Fatalf("Total weight is higher than 10_000\nInput: %v\nOutput: %v\n", input, output)
		}

		if MaxFinalityVotePercentage-totalWeight >= uint16(len(input)) {
			t.Fatalf("Total weight error is too high\nTotal weight: %d\nInput: %v\nOutput: %v\n", totalWeight, input, output)
		}
	})
}
