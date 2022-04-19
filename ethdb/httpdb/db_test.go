package httpdb

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type mockRpc struct {}

func (m *mockRpc) Call(result interface{}, method string, args ...interface{}) error {
	result = []byte("0x123")
	return nil
}

func TestEvict(t *testing.T) {
	metrics.Enabled = true
	db := NewDBWithLRU("", "", 0)
	db.client = &mockRpc{}

	var totalSize int64
	for i := 0; i <= defaultCachedItems; i++ {
		key := []byte(fmt.Sprintf("key-%d", i))
		val := []byte(fmt.Sprintf("val-%d", i))
		if err := db.Put(key, val); err != nil {
			t.Fatal(err)
		}
		if i > 0 {
			totalSize += int64(len(val))
		}
	}
	time.Sleep(3*time.Second)
	require.Equal(t, totalSize, cacheItemsSizeCounter.Count())
}

func TestRedisCache(t *testing.T) {
	db := NewRedisCache([]string{"127.0.0.1:6379"}, 0, &redis.Options{})
	if err := db.Put(common.Hex2Bytes("0xmy-key-1"), common.Hex2Bytes("0xmy-value")); err != nil {
		t.Fatal(err)
	}
	val, err := db.Get(common.Hex2Bytes("0xmy-key-1"))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, common.Hex2Bytes("0xmy-value"), val)
}
