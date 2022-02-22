package httpdb

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	lru "github.com/hashicorp/golang-lru"
)

const (
	GET               = "consortium_getDBValue"
	defaultCachedSize = 1024
)

var notfoundErr = errors.New("not found")

func query(client *rpc.Client, method string, params ...interface{}) ([]byte, error) {
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
	client      *rpc.Client
	cachedItems *lru.Cache
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

func NewDB(rpcUrl string, cachedSize int) *DB {
	client, err := rpc.DialHTTP(rpcUrl)
	if err != nil {
		log.Error("[httpdb] NewDB", "err", err)
		return nil
	}
	db := &DB{
		client: client,
	}
	if cachedSize > 0 {
		db.cachedItems, _ = lru.New(cachedSize)
	} else {
		db.cachedItems, _ = lru.New(defaultCachedSize)
	}
	return db
}

func (db *DB) Close() error {
	return nil
}

func (db *DB) Has(key []byte) (bool, error) {
	if db.cachedItems.Contains(key) {
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
	if res, ok := db.cachedItems.Get(common.Bytes2Hex(key)); ok {
		return res.([]byte), nil
	}
	// try to get data from rpc if res is nil
	val, err = query(db.client, GET, common.Bytes2Hex(key))
	if err != nil {
		return nil, err
	}
	if len(val) == 0 {
		return nil, notfoundErr
	}
	// store val to memory db for later use
	return val, db.Put(key, val)
}

func (db *DB) Put(key, value []byte) error {
	db.cachedItems.Add(common.Bytes2Hex(key), value)
	return nil
}

func (db *DB) Delete(key []byte) error {
	db.cachedItems.Remove(common.Bytes2Hex(key))
	return nil
}

func (db *DB) Stats() {}

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
	return []byte{}, nil
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
