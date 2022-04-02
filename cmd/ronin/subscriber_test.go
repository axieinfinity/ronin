package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/consortium"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/ethdb/httpdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestReprocessTransaction(t *testing.T) {
	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.LvlDebug)
	log.Root().SetHandler(glogger)
	db := httpdb.NewDBWithLRU("https://api-archived.roninchain.com/rpc", "https://api-archived.roninchain.com/rpc", 0)
	chainConfig, _, _ := core.SetupGenesisBlockWithOverride(db, nil, nil)
	block := rawdb.ReadBlock(db, common.HexToHash("0xa1fa3122c021db882f8dff2e389bd1b0f20d43e4c38ab0a4c239f28e5d795ed4"), 12208992)
	if block == nil {
		t.Fatal("cannot find block")
	}
	println(block.Hash().Hex())
	parentBlock := rawdb.ReadBlock(db, block.ParentHash(),12208991)
	if parentBlock == nil {
		t.Fatal("cannot find block")
	}
	errChan := make(chan error, 1)
	txResult, internalTxs := reprocessBlock(&chainContext{db}, parentBlock.Root(), block.Header(), block.Transactions(), chainConfig, state.NewDatabase(db), errChan)
	err := <-errChan
	require.NoError(t, err)
	require.Greater(t, len(txResult), 0)
	require.Len(t, len(internalTxs), 9)
}

type chainContext struct {
	db ethdb.Database
}

func (c *chainContext) Engine() consensus.Engine {
	return consortium.New(&params.ConsortiumConfig{
		Epoch: 600,
		Period: 3,
	}, c.db)
}

func (c *chainContext) GetHeader(hash common.Hash, number uint64) *types.Header {
	return rawdb.ReadHeader(c.db, hash, number)
}
