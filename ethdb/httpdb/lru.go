package httpdb

import (
	"github.com/VictoriaMetrics/fastcache"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	lru "github.com/hashicorp/golang-lru"
	"sync"
	"time"
)

type lruCache struct {
	lock            sync.Mutex
	maxEntries      int
	lruCache        *lru.Cache
	cache           *fastcache.Cache
	cleanupInterval time.Duration
	resetThreshold  int
}

func (c *lruCache) onEvicted(key, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()

	log.Debug("onEvicted", "key", key)
	cacheItemsCounter.Inc(-1)
	cacheItemsSizeCounter.Inc(-int64(value.(int)))
	c.cache.Del(common.Hex2Bytes(key.(string)))
}

func NewLRUCache(cachedSize, resetThreshold int) *lruCache {
	cache := &lruCache{
		maxEntries:      defaultCachedItems,
		cache:           fastcache.New(allowedMaxSize),
		cleanupInterval: defaultCleanUp,
		resetThreshold:  defaultResetThreshold,
	}
	if cachedSize > 0 {
		cache.maxEntries = cachedSize
	}
	if resetThreshold > 0 && resetThreshold > cache.maxEntries {
		cache.resetThreshold = resetThreshold
	}
	cache.lruCache, _ = lru.NewWithEvict(cache.maxEntries, cache.onEvicted)

	go func() {
		for {
			select {
			case <-time.Tick(cache.cleanupInterval):
				cache.purge()
			}
		}
	}()
	return cache
}

func (c *lruCache) purge() {
	items := cacheItemsCounter.Count()
	if items == 0 {
		return
	}
	if items > int64(c.resetThreshold) {
		log.Debug("data is growing out of control, start purging")
		c.lruCache.Purge()
		c.cache.Reset()
		cacheItemsCounter.Clear()
		cacheItemsSizeCounter.Clear()
	}
}

func (c *lruCache) Get(key []byte) ([]byte, error) {
	hexKey := common.Bytes2Hex(key)
	if res, ok := c.cache.HasGet(nil, key); ok {
		// update recent-ness in lru cache
		c.lruCache.Get(hexKey)
		// increase hit counter
		cacheHitCounter.Inc(1)
		return res, nil
	}
	return nil, notfoundErr
}

func (c *lruCache) Has(key []byte) (bool, error) {
	return c.cache.Has(key), nil
}

func (c *lruCache) Put(key, value []byte) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.lruCache.Add(common.Bytes2Hex(key), len(value))
	c.cache.Set(key, value)
	return nil
}

func (c *lruCache) Delete(key []byte) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.lruCache.Remove(common.Bytes2Hex(key))
	c.cache.Del(key)
	return nil
}
