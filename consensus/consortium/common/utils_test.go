package common

import (
	"github.com/ethereum/go-ethereum/common"
	"reflect"
	"testing"
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
