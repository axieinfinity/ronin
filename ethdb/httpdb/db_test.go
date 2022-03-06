package httpdb

import (
	"fmt"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/stretchr/testify/require"
	"testing"
)

type mockRpc struct {}

func (m *mockRpc) Call(result interface{}, method string, args ...interface{}) error {
	result = []byte("0x123")
	return nil
}

func TestEvict(t *testing.T) {
	metrics.Enabled = true
	db := NewDB("", "", 0)
	db.client = &mockRpc{}

	var totalSize int64
	for i := 0; i <= defaultCachedSize; i++ {
		key := []byte(fmt.Sprintf("key-%d", i))
		val := []byte(fmt.Sprintf("val-%d", i))
		if err := db.Put(key, val); err != nil {
			t.Fatal(err)
		}
		if i > 0 {
			totalSize += int64(len(val))
		}
	}
	require.Equal(t, int64(1), evictedCallCounter.Count())
	require.Equal(t, totalSize, cacheItemsSizeCounter.Count())
}