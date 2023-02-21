package common

import (
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
