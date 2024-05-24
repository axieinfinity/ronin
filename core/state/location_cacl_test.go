package state

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestGetLocMappingAtKey(t *testing.T) {
	hash := GetLocMappingAtKey(common.BigToHash(big.NewInt(10)), 12)

	expect := common.HexToHash("0x9e6c92d7be355807bd948171438a5e65aaf9e4c36f1405c1b9ca25d27c4ea3a0")
	if hash != expect {
		t.Fatalf("Hash mismatches, got %s expect %s", hash, expect)
	}
}

func BenchmarkGetLocMappingAtKey(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GetLocMappingAtKey(common.BigToHash(big.NewInt(10)), 12)
	}
}
