package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/consortium"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/proxy"
	"testing"
)

func TestReprocessTransaction(t *testing.T) {
	ethConfig := ethconfig.Defaults
	ethConfig.NetworkId = 2020
	backend, err := proxy.NewBackend(&proxy.Config{RPC: "https://api-archived.roninchain.com/rpc"}, &ethConfig)
	if err != nil {
		t.Fatal(err)
	}
	block := rawdb.ReadBlock(backend.ChainDb(), common.HexToHash("0xccc6dc29e86b3ceaf2f08e04ccf71641bc4d2223fb142707b4e0699d63743698"), 10000017)
	if block == nil {
		t.Fatal("cannot find block")
	}
	parentBlock := rawdb.ReadBlock(backend.ChainDb(), block.ParentHash(), 10000016)
	if parentBlock == nil {
		t.Fatal("cannot find block")
	}
	statedb, err := state.New(parentBlock.Header().Root, state.NewDatabaseWithConfig(backend.ChainDb(), nil), nil)
	if err != nil {
		t.Fatal(err)
	}
	balance := statedb.GetBalance(common.HexToAddress("0x3D20380A3815Ff52CB41c032A4Fe93877a2AD614"))
	println(balance.String())
	result, err := reprocessBlock(&chainContext{backend.ChainDb()}, parentBlock.Root(), block.Header(), block.Transactions(), backend.ChainConfig(), backend.ChainDb())
	if err != nil {
		t.Fatal(err)
	}
	println(result)
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
