package core

import (
	"bytes"
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/rlp"
)

const (
	JournalSectionSize = 256                    // size of indexed section
	journalConfirm     = 256                    // the number of confirmation blocks before indexing
	journalThrottling  = 100 * time.Millisecond // time to wait for next index iteration
)

// JournalIndexer implements a core.ChainIndexer, building up the batch of
// journal from journal in each block to reduce the disk read operations
// when reading journal. This indexer assumes there is no reorg after the
// journalConfirm.
type JournalIndexer struct {
	db            ethdb.Database // database to write indexed section to
	size          uint64         // size of indexed section
	section       uint64         // currently indexed section number
	current       uint64         // the current number of stored journal in indexed section
	buffer        []byte         // the current data in indexed section
	metadata      []uint64       // metadata of the batch journal
	sectionHashes []common.Hash  // the block hash of indexed block journals that is deleted when batch journal is finalized
}

// NewJournalIndexer returns a chain indexer that generates batch journal from
// canonical chain.
func NewJournalIndexer(db ethdb.Database) *ChainIndexer {
	journalIndexer := &JournalIndexer{
		db:   db,
		size: JournalSectionSize,
	}
	table := rawdb.NewTable(db, string(rawdb.JournalIndexPrefix))
	return NewChainIndexer(db, table, journalIndexer, JournalSectionSize, journalConfirm, journalThrottling, "journalIndexer")
}

// Reset implements core.ChainIndexerBackend, starting a new journal index
// section.
func (indexer *JournalIndexer) Reset(ctx context.Context, section uint64, lastSectionHead common.Hash) error {
	indexer.buffer = make([]byte, 0)
	indexer.metadata = make([]uint64, indexer.size)
	indexer.section = section
	indexer.current = 0
	indexer.sectionHashes = make([]common.Hash, 0)
	return nil
}

// Process implements core.ChainIndexerBackend, reads the size from block
// journal to build metadata of batch journal and merges the journal data
// to index section without rlp decoding
func (indexer *JournalIndexer) Process(ctx context.Context, header *types.Header) error {
	journal := rawdb.ReadStoredJournal(indexer.db, header.Hash())
	if len(journal) == 0 {
		indexer.metadata[header.Number.Uint64()-indexer.section*indexer.size] = indexer.current
		return nil
	}

	reader := bytes.NewReader(journal)
	size, err := state.BlockJournalSize(reader)
	if err != nil {
		return err
	}

	rawBytes := make([]byte, reader.Len())
	_, err = reader.Read(rawBytes)
	if err != nil {
		return err
	}
	indexer.buffer = append(indexer.buffer, rawBytes...)
	indexer.metadata[header.Number.Uint64()-indexer.section*indexer.size] = indexer.current
	indexer.current += uint64(size)
	indexer.sectionHashes = append(indexer.sectionHashes, header.Hash())

	return nil
}

// Commit implements core.ChainIndexerBackend, writes the metadata and
// index section data to disk, deletes the indexed block journal
func (indexer *JournalIndexer) Commit() error {
	if len(indexer.buffer) == 0 {
		return nil
	}

	batchJournal, err := rlp.EncodeToBytes(state.BatchJournalRLP{
		Metadata: indexer.metadata,
		Data:     indexer.buffer,
	})
	if err != nil {
		return err
	}

	batch := indexer.db.NewBatch()
	rawdb.WriteBatchJournal(batch, indexer.section, batchJournal)

	for _, hash := range indexer.sectionHashes {
		rawdb.DeleteStoredJournal(batch, hash)
	}

	return batch.Write()
}

// Prune implements core.ChainIndexerBackend, this indexer does not
// support pruning
func (indexer *JournalIndexer) Prune(threshold uint64) error {
	return nil
}
