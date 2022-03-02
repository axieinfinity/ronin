package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/consortium"
	"github.com/ethereum/go-ethereum/core/rawdb"
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
	block := rawdb.ReadBlock(backend.ChainDb(), common.HexToHash("0x280978a4362eb0391547a28b46a8ac1df344d024d0b2a5c2d0ff653167a3b7a0"), 10000000)
	if block == nil {
		t.Fatal("cannot find block")
	}
	result, err := reprocessBlock(&chainContext{backend.ChainDb()}, block.Header(), block.Transactions(), backend.ChainConfig(), backend.ChainDb())
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
