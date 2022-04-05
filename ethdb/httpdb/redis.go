package httpdb

import (
	"context"
	"github.com/eko/gocache/store"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	defaultTTL = time.Minute
	defaultReadTimeout = 15 * time.Second
	defaultWriteTimeout = defaultReadTimeout
	defaultConnectTimeout = defaultReadTimeout
)

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
	isCluster bool
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
	val, err := c.readStore.Get(context.Background(), common.Bytes2Hex(key))
	if err != nil && err.Error() == "redis: nil" {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return []byte(val.(string)), nil
}

func (c *redisCache) Has(key []byte) (bool, error) {
	val, err := c.Get(key)
	return val != nil, err
}

func (c *redisCache) Put(key, value []byte) error {
	if has, _ := c.Has(key); has {
		return nil
	}
	ctx := context.Background()
	return c.writeStore.Set(ctx, common.Bytes2Hex(key), value, nil)
}

func (c *redisCache) Delete(key []byte) error {
	val, err := c.Get(key)
	if err != nil {
		return err
	}
	if err = c.writeStore.Delete(context.Background(), common.Bytes2Hex(key)); err != nil {
		return err
	}
	cacheItemsCounter.Inc(-1)
	cacheItemsSizeCounter.Inc(-int64(len(val)))
	return nil
}
