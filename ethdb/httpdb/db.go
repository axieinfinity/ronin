package httpdb

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/go-redis/redis/v8"
	"strings"
	"time"
)

const (
	GET                = "consortium_getDBValue"
	ANCIENT            = "consortium_getAncientValue"
	defaultCachedItems = 1024
	allowedMaxSize     = 64 * 1024 * 1024 // 64 MB
)

type IClient interface {
	Call(result interface{}, method string, args ...interface{}) error
}

type Cache interface {
	ethdb.KeyValueWriter
	ethdb.KeyValueReader
}

var (
	notfoundErr           = errors.New("not found")
	requestCounter        metrics.Counter
	cacheHitCounter       metrics.Counter
	cacheItemsCounter     metrics.Counter
	cacheItemsSizeCounter metrics.Counter
	requestRpcCounter     metrics.Counter
	requestArchiveCounter metrics.Counter
)

func initMetrics() {
	requestCounter = metrics.GetOrRegisterCounter("cache/request", nil)
	cacheHitCounter = metrics.GetOrRegisterCounter("cache/request/hit", nil)
	requestRpcCounter = metrics.GetOrRegisterCounter("cache/request/rpc", nil)
	requestArchiveCounter = metrics.GetOrRegisterCounter("cache/request/archive", nil)
	cacheItemsCounter = metrics.GetOrRegisterCounter("cache/items", nil)
	cacheItemsSizeCounter = metrics.GetOrRegisterCounter("cache/items/size", nil)
}

func getAncientKey(kind string, number uint64) []byte {
	return []byte(fmt.Sprintf("ancient-%s-%d", kind, number))
}

func PutAncient(db ethdb.Database, kind string, number uint64, value []byte) error {
	return db.Put(getAncientKey(kind, number), value)
}

func RemoveAncient(db ethdb.Database, kind string, number uint64) error {
	return db.Delete(getAncientKey(kind, number))
}

func query(client IClient, method string, params ...interface{}) ([]byte, error) {
	var res string
	if err := client.Call(&res, method, params...); err != nil {
		return nil, err
	}
	return common.Hex2Bytes(res), nil
}

// DB is a read only database which is used to query data from other nodes by using consortium's rpc 'consortium_getDBValue'
// it also caches return values using lru cache to get data immediately without RPC calling everytime.
// DB only supports Get and Has function which query data from other nodes.
type DB struct {
	client, archive IClient
	cache           Cache
}

func (db *DB) AncientRange(kind string, start, count, maxBytes uint64) ([][]byte, error) {
	panic("implement me")
}

func (db *DB) ReadAncients(fn func(ethdb.AncientReader) error) (err error) {
	return fn(db)
}

func (db *DB) ModifyAncients(f func(ethdb.AncientWriteOp) error) (int64, error) {
	panic("implement me")
}

func NewDBWithLRU(rpcUrl, archive string, cachedSize int) *DB {
	return NewDB(rpcUrl, archive, NewLRUCache(cachedSize))
}

func NewDBWithRedis(rpcUrl, archive string, expiration time.Duration, options *redis.Options) *DB {
	addresses := strings.Split(options.Addr, ",")
	if len(addresses) == 0 {
		addresses = append(addresses, "")
	}
	return NewDB(rpcUrl, archive, NewRedisCache(addresses, expiration, options))
}

func NewDB(rpcUrl, archive string, cache Cache) *DB {
	client, err := rpc.DialHTTP(rpcUrl)
	if err != nil {
		log.Error("[httpdb][NewDB] Dial RPC", "err", err)
		return nil
	}
	db := &DB{
		client: client,
		cache:  cache,
	}
	if archive != "" {
		if db.archive, err = rpc.DialHTTP(archive); err != nil {
			log.Error("[httpdb][NewDB] Dial Archive", "err", err)
			return nil
		}
	}
	initMetrics()
	return db
}

func (db *DB) Close() error {
	return nil
}

func (db *DB) Has(key []byte) (bool, error) {
	if ok, _ := db.cache.Has(key); ok {
		return true, nil
	}
	// try to get data from rpc if data cannot be found in cached items
	_, err := db.Get(key)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (db *DB) Get(key []byte) (val []byte, err error) {
	requestCounter.Inc(1)
	res, err := db.cache.Get(key)
	if err != nil {
		log.Error("[httpdb] getting data from cache", "err", err)
	}
	if res != nil {
		// increase hit counter
		cacheHitCounter.Inc(1)
		return res, nil
	}
	hexKey := common.Bytes2Hex(key)
	log.Debug("calling getDbValue via rpc", "key", hexKey)
	// try to get data from rpc if res is nil
	requestRpcCounter.Inc(1)
	val, err = query(db.client, GET, hexKey)
	if err != nil {
		// try to get data from archive if it is not nil
		if db.archive == nil {
			return nil, err
		}
		log.Debug("calling getDbValue via archive", "key", hexKey, "err", err.Error())
		requestArchiveCounter.Inc(1)
		val, err = query(db.archive, GET, hexKey)
		if err != nil {
			return nil, err
		}
	}
	if len(val) == 0 {
		log.Error("value not found", "key", hexKey)
		return nil, notfoundErr
	}
	log.Debug("getDbValue found", "key", hexKey)
	// store val to memory db for later use
	go db.Put(key, val)
	return val, nil
}

func (db *DB) Put(key, value []byte) error {
	if err := db.cache.Put(key, value); err != nil {
		return err
	}

	cacheItemsSizeCounter.Inc(int64(len(value)))
	cacheItemsCounter.Inc(1)

	return nil
}

func (db *DB) Delete(key []byte) error {
	if err := db.cache.Delete(key); err != nil {
		return err
	}
	return nil
}

func (db *DB) NewBatch() ethdb.Batch {
	return nil
}

func (db *DB) NewIterator(prefix []byte, start []byte) ethdb.Iterator {
	return nil
}

// Stat returns a particular internal stat of the database.
func (db *DB) Stat(property string) (string, error) {
	return "", errors.New("unknown property")
}

// Compact is not supported on a http db
func (db *DB) Compact(start []byte, limit []byte) error {
	return nil
}

// HasAncient returns an indicator whether the specified data exists in the
// ancient store.
func (db *DB) HasAncient(kind string, number uint64) (bool, error) {
	return false, nil
}

// Ancient retrieves an ancient binary blob from the append-only immutable files.
func (db *DB) Ancient(kind string, number uint64) ([]byte, error) {
	key := getAncientKey(kind, number)
	requestCounter.Inc(1)
	res, err := db.cache.Get(key)
	if err != nil {
		log.Error("[httpdb] getting ancient data from cache", "err", err)
	}
	if res != nil {
		// increase hit counter
		cacheHitCounter.Inc(1)
		return res, nil
	}
	hexData := hexutil.EncodeUint64(number)
	log.Debug("calling ancient via rpc", "kind", kind, "number", number)
	val, err := query(db.client, ANCIENT, kind, hexData)
	if err != nil {
		if strings.Contains(err.Error(), "out of bounds") {
			log.Debug("out of bounds, don't need to retry archive", "ancientKey", string(key))
			return nil, err
		}
		// try to get data from archive if it is not nil
		if db.archive == nil {
			return nil, err
		}
		log.Debug("calling ancient via archive", "kind", kind, "number", number)
		val, err = query(db.archive, ANCIENT, kind, hexData)
		if err != nil {
			log.Debug("value not found via get ancient", "kind", kind, "number", number, "err", err)
			return nil, err
		}
	}
	log.Debug("saving ancient data", "kind", kind, "number", number)
	if err = db.Put(key, val); err != nil {
		return nil, err
	}
	return val, nil
}

// Ancients returns the ancient item numbers in the ancient store.
func (db *DB) Ancients() (uint64, error) { return 0, nil }

// AncientSize returns the ancient size of the specified category.
func (db *DB) AncientSize(kind string) (uint64, error) { return 0, nil }

// AppendAncient injects all binary blobs belong to block at the end of the
// append-only immutable table files.
func (db *DB) AppendAncient(number uint64, hash, header, body, receipt, td []byte) error { return nil }

// TruncateAncients discards all but the first n ancient data from the ancient store.
func (db *DB) TruncateAncients(n uint64) error { return nil }

// Sync flushes all in-memory ancient store data to disk.
func (db *DB) Sync() error { return nil }
