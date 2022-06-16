package httpdb

import (
	"github.com/VictoriaMetrics/fastcache"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	lru "github.com/hashicorp/golang-lru"
)

type lruCache struct {
	maxEntries int
	lruCache   *lru.Cache
	cache      *fastcache.Cache
}

func (c *lruCache) onEvicted(key, value interface{}) {
	log.Debug("onEvicted", "key", key)
	cacheItemsCounter.Inc(-1)
	cacheItemsSizeCounter.Inc(-int64(value.(int)))
	c.cache.Del(common.Hex2Bytes(key.(string)))
}

func NewLRUCache(cachedSize int) *lruCache {
	cache := &lruCache{
		maxEntries: defaultCachedItems,
		cache:      fastcache.New(allowedMaxSize),
	}
	if cachedSize > 0 {
		cache.maxEntries = cachedSize
	}
	cache.lruCache, _ = lru.NewWithEvict(cache.maxEntries, cache.onEvicted)
	return cache
}

func (c *lruCache) Get(key []byte) ([]byte, error) {
	hexKey := common.Bytes2Hex(key)
	if res, ok := c.cache.HasGet(nil, key); ok {
		// update recent-ness in lru cache
		c.lruCache.Get(hexKey)
		return res, nil
	}
	return nil, notfoundErr
}

func (c *lruCache) Has(key []byte) (bool, error) {
	return c.cache.Has(key), nil
}

func (c *lruCache) Put(key, value []byte) error {
	c.lruCache.Add(common.Bytes2Hex(key), len(value))
	c.cache.Set(key, value)
	return nil
}

func (c *lruCache) Delete(key []byte) error {
	c.lruCache.Remove(common.Bytes2Hex(key))
	c.cache.Del(key)
	return nil
}
