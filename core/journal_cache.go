package core

import (
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/ethdb"
	lru "github.com/hashicorp/golang-lru"
)

// journalCache is the caching layer when loading batch journal from database
type journalCache struct {
	db         ethdb.Database
	batchCache *lru.Cache
}

func newJournalCache(db ethdb.Database) *journalCache {
	lru, _ := lru.New(2)

	return &journalCache{
		db:         db,
		batchCache: lru,
	}
}

// loadFromBatchJournal looks up the block's journal from the cached batch journal
func (cache *journalCache) loadFromBatchJournal(blockNumber uint64) ([]state.StoredJournal, error) {
	var (
		batch *state.BatchJournal
		err   error
	)

	requestedSection := blockNumber / JournalSectionSize
	value, ok := cache.batchCache.Get(requestedSection)
	if !ok {
		batch, err = state.LoadBatchJournal(cache.db, requestedSection)
		if err != nil {
			return nil, err
		}
		if batch == nil {
			cache.batchCache.Add(requestedSection, state.BatchJournal{})
			batchValue := state.BatchJournal{}
			batch = &batchValue
		} else {
			cache.batchCache.Add(requestedSection, *batch)
		}
	} else {
		batchValue := value.(state.BatchJournal)
		batch = &batchValue
	}

	if len(batch.Data) == 0 {
		return nil, nil
	}

	return state.GetBlockJournalFromBatch(batch, blockNumber%JournalSectionSize, JournalSectionSize), nil
}
