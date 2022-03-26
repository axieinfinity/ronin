package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/consortium"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/ethdb/httpdb"
	"github.com/ethereum/go-ethereum/params"
	"testing"
)

func TestReprocessTransaction(t *testing.T) {
	db := httpdb.NewDBWithLRU("https://api-archived.roninchain.com/rpc", "https://api-archived.roninchain.com/rpc", 0, 0)
	chainConfig, _, err := core.SetupGenesisBlockWithOverride(db, nil, nil)
	blockHash := rawdb.ReadCanonicalHash(db, 12185907)
	println(blockHash.Hex())
	block := rawdb.ReadBlock(db, common.HexToHash("0xa1fa3122c021db882f8dff2e389bd1b0f20d43e4c38ab0a4c239f28e5d795ed4"), 12208992)
	if block == nil {
		t.Fatal("cannot find block")
	}
	println(block.Hash().Hex())
	parentBlock := rawdb.ReadBlock(db, block.ParentHash(),12208991)
	if parentBlock == nil {
		t.Fatal("cannot find block")
	}
	txResult := make(chan *TransactionResult, 100)
	//balance := statedb.GetBalance(common.HexToAddress("0x3D20380A3815Ff52CB41c032A4Fe93877a2AD614"))
	//println(balance.String())
	go func() {
		for {
			select {
			case result := <-txResult:
				println(fmt.Sprintf("hash:%s UsedGas:%d Error:%s Data:%s", result.TransactionHash.Hex(), result.UsedGas, result.Err, result.ReturnData.String()))
			}
		}
	}()
	err = reprocessBlock(&chainContext{db}, parentBlock.Root(), block.Header(), block.Transactions(), chainConfig, state.NewDatabase(db), txResult)
	if err != nil {
		t.Fatal(err)
	}
	//println(result)
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
