package httpdb

import (
	"context"
	"errors"
	"fmt"
	"github.com/eko/gocache/store"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/go-redis/redis/v8"
	"math"
	"sync"
	"time"
)

var (
	defaultTTL            = time.Minute
	defaultReadTimeout    = 15 * time.Second
	defaultWriteTimeout   = defaultReadTimeout
	defaultConnectTimeout = defaultReadTimeout

	defaultLimitSize = 128.0
)

const keyFmt = "k%sidx%d"

type RedisStoreInterface interface {
	Get(ctx context.Context, key interface{}) (interface{}, error)
	GetWithTTL(ctx context.Context, key interface{}) (interface{}, time.Duration, error)
	Set(ctx context.Context, key interface{}, value interface{}, options *store.Options) error
	Delete(ctx context.Context, key interface{}) error
	Invalidate(ctx context.Context, options store.InvalidateOptions) error
	GetType() string
	Clear(ctx context.Context) error
}

type redisCache struct {
	isCluster             bool
	readStore, writeStore RedisStoreInterface
}

func copyOptions(options *redis.Options) *redis.Options {
	return &redis.Options{
		Addr:         options.Addr,
		ReadTimeout:  options.ReadTimeout,
		WriteTimeout: options.WriteTimeout,
		PoolSize:     options.PoolSize,
		PoolTimeout:  options.PoolTimeout,
	}
}

func NewRedisCache(addresses []string, expiration time.Duration, options *redis.Options) *redisCache {
	var (
		readStore, writeStore RedisStoreInterface
	)
	opts := &store.Options{Expiration: defaultTTL}
	if expiration > 0 {
		opts.Expiration = expiration
	}
	if options.ReadTimeout == 0 {
		options.ReadTimeout = defaultReadTimeout
	}
	if options.WriteTimeout == 0 {
		options.WriteTimeout = defaultWriteTimeout
	}
	if options.PoolTimeout == 0 {
		options.PoolTimeout = defaultConnectTimeout
	}
	var client store.RedisClientInterface
	if len(addresses) == 1 {
		client = redis.NewClient(options)
		readStore = store.NewRedis(client, opts)
		writeStore = readStore
	} else if len(addresses) > 2 {
		client = redis.NewClusterClient(&redis.ClusterOptions{Addrs: addresses})
		readStore = store.NewRedisCluster(redis.NewClusterClient(&redis.ClusterOptions{Addrs: addresses}), opts)
		writeStore = readStore
	} else if len(addresses) == 2 {
		// set write store with first address as master
		writeClient := redis.NewClient(copyOptions(options)).WithContext(context.Background())
		writeClient.Options().Addr = addresses[0]
		writeStore = store.NewRedis(writeClient, opts)

		// set read store with second address as slave
		readClient := redis.NewClient(copyOptions(options)).WithContext(context.Background())
		readClient.Options().Addr = addresses[1]
		readStore = store.NewRedis(readClient, opts)
	} else {
		panic("cannot init new redisCache")
	}
	return &redisCache{readStore: readStore, writeStore: writeStore}
}

func (c *redisCache) Get(key []byte) ([]byte, error) {
	redisKey := common.Bytes2Hex(key)
	length, err := c.getInt(redisKey)
	if err != nil {
		return nil, err
	}
	valuesMap := make(map[int][]byte)
	size := int(math.Ceil(float64(length) / defaultLimitSize))
	var (
		wg   sync.WaitGroup
		lock sync.Mutex
	)
	wg.Add(size)
	for i := 0; i < size; i++ {
		go func(w *sync.WaitGroup, index int) {
			val, err := c.getBytes(fmt.Sprintf(keyFmt, redisKey, index))
			if err != nil {
				log.Error("[redisCache][Get] error while get bytes value from index", "err", err, "key", fmt.Sprintf(keyFmt, redisKey, index))
			}

			lock.Lock()
			valuesMap[index] = val
			lock.Unlock()

			w.Done()
		}(&wg, i)
	}
	wg.Wait()
	values := make([]byte, 0)
	for _, v := range valuesMap {
		values = append(values, v...)
	}
	return values, nil
}

func (c *redisCache) getInt(key string) (int, error) {
	val, err := c.get(key)
	if err != nil {
		return -1, err
	}
	if _, ok := val.(int); !ok {
		return -1, errors.New("cannot cast interface value into integer")
	}
	return val.(int), nil
}

func (c *redisCache) getBytes(key string) ([]byte, error) {
	val, err := c.get(key)
	if err != nil {
		return nil, err
	}
	if _, ok := val.([]byte); !ok {
		return nil, errors.New("cannot cast interface value into integer")
	}
	return val.([]byte), nil
}

func (c *redisCache) get(key string) (interface{}, error) {
	val, err := c.readStore.Get(context.Background(), key)
	if err != nil && err.Error() == "redis: nil" {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return val, err
}

func (c *redisCache) Has(key []byte) (bool, error) {
	val, err := c.get(common.Bytes2Hex(key))
	return val != nil, err
}

func (c *redisCache) Put(key, value []byte) error {
	if has, _ := c.Has(key); has {
		return nil
	}
	// the main key will contain the length of value
	redisKey := common.Bytes2Hex(key)
	size := int(math.Ceil(float64(len(value)) / defaultLimitSize))
	if err := c.put(redisKey, len(value)); err != nil {
		return err
	}
	// then loop through size and separate value based on limit size
	start, end := 0, 0
	for i := 0; i < size; i++ {
		if end+int(defaultLimitSize) < len(value) {
			end += int(defaultLimitSize)
		} else {
			end = len(value)
		}
		if err := c.put(fmt.Sprintf(keyFmt, redisKey, i), value[start:end]); err != nil {
			return err
		}
		start = end
	}
	return c.put(common.Bytes2Hex(key), value)
}

func (c *redisCache) put(key string, value interface{}) error {
	return c.writeStore.Set(context.Background(), key, value, nil)
}

func (c *redisCache) Delete(key []byte) error {
	redisKey := common.Bytes2Hex(key)
	length, err := c.get(redisKey)
	if err != nil {
		return err
	}
	if _, ok := length.(int); !ok {
		return errors.New("cannot cast size to int")
	}
	size := int(math.Ceil(float64(length.(int)) / defaultLimitSize))
	var wg sync.WaitGroup
	wg.Add(size)
	// delete all indexes
	for i := 0; i < size; i++ {
		go func(w *sync.WaitGroup, index int) {
			if err := c.writeStore.Delete(context.Background(), fmt.Sprintf(keyFmt, redisKey, index)); err != nil {
				log.Error("[redisCache][Delete] error while deleting data", "key", fmt.Sprintf(keyFmt, redisKey, index))
			}
			wg.Done()
		}(&wg, i)
	}
	wg.Wait()
	if err = c.writeStore.Delete(context.Background(), redisKey); err != nil {
		return err
	}
	cacheItemsCounter.Inc(-1)
	cacheItemsSizeCounter.Inc(-int64(length.(int)))
	return nil
}
