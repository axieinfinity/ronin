package httpdb

import (
	"context"
	"github.com/eko/gocache/store"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis/v8"
	"time"
)

var defaultTTL = time.Minute

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
	store RedisStoreInterface
	//locker *redislock.Client
}

func NewRedisCache(addresses []string, expiration time.Duration) *redisCache {
	var redisStore RedisStoreInterface
	opts := &store.Options{Expiration: defaultTTL}
	if expiration > 0 {
		opts.Expiration = expiration
	}
	var client store.RedisClientInterface
	if len(addresses) == 1 {
		client = redis.NewClient(&redis.Options{Addr: addresses[0]})
		redisStore = store.NewRedis(client, opts)
	} else if len(addresses) > 1 {
		client = redis.NewClusterClient(&redis.ClusterOptions{Addrs: addresses})
		redisStore = store.NewRedisCluster(redis.NewClusterClient(&redis.ClusterOptions{Addrs: addresses}), opts)
	} else {
		panic("cannot init new redisCache")
	}
	//return &redisCache{store: redisStore, locker: redislock.New(client.(redislock.RedisClient))}
	return &redisCache{store: redisStore}
}

func (c *redisCache) Get(key []byte) ([]byte, error) {
	val, err := c.store.Get(context.Background(), common.Bytes2Hex(key))
	if err != nil && err.Error() == "redis: nil" {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return []byte(val.(string)), nil
	//if ok {
	//	if len(rs) == 0 {
	//		return nil, nil
	//	}
	//	return rs, nil
	//}
	//return nil, errors.New(fmt.Sprintf("invalid return type expected string - got %v", reflect.TypeOf(val).String()))
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
	//locker, err := c.locker.Obtain(ctx, common.Bytes2Hex(key), time.Second, nil)
	//if err != nil {
	//	if err == redislock.ErrNotObtained {
	//		return nil
	//	}
	//	return err
	//}
	//defer locker.Release(ctx)
	return c.store.Set(ctx, common.Bytes2Hex(key), value, nil)
}

func (c *redisCache) Delete(key []byte) error {
	if has, _ := c.Has(key); !has {
		return nil
	}
	ctx := context.Background()
	//locker, err := c.locker.Obtain(ctx, common.Bytes2Hex(key), time.Second, nil)
	//if err != nil {
	//	if err == redislock.ErrNotObtained {
	//		return nil
	//	}
	//	return err
	//}
	//defer locker.Release(ctx)
	return c.store.Delete(ctx, common.Bytes2Hex(key))
}
