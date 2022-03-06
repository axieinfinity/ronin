package proxy

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/rpc"
	"testing"
	"time"
)

func TestEthCall(t *testing.T) {
	ethConfig := ethconfig.Defaults
	ethConfig.RPCEVMTimeout = time.Second*50
	backend, err := newBackend(&Config{ArchiveUrl: "https://api-archived.roninchain.com/rpc", RpcUrl: "http://localhost:8000"}, &ethConfig)
	if err != nil {
		t.Fatal(err)
	}
	api := newAPI(backend)
	to := common.HexToAddress("0xc99a6a985ed2cac1ef41640596c5a5f9f4e19ef5")
	blockNumber := rpc.LatestBlockNumber
	data := hexutil.Bytes(common.Hex2Bytes("0x70a082310000000000000000000000006755b9c63779d5b16cfa434e9e1c018689b03e45"))
	rs, err := api.Call(context.Background(), ethapi.TransactionArgs{
		Data: &data,
		To: &to,
	}, rpc.BlockNumberOrHash{BlockNumber: &blockNumber}, nil)
	if err != nil && err.Error() != "execution reverted" {
		t.Fatal(err)
	}
	println(rs.String())
}

func TestGetStateAtBlock(t *testing.T) {
	ethConfig := ethconfig.Defaults
	backend, err := newBackend(&Config{ArchiveUrl: "http://localhost:8545", RpcUrl: "http://localhost:8549"}, &ethConfig)
	if err != nil {
		t.Fatal(err)
	}
	println("start getting state by number")
	_, _, err = backend.StateAndHeaderByNumber(context.Background(), rpc.BlockNumber(1000))
	if err != nil{
		t.Fatal(err)
	}
}

func TestGetStateAtBehindBlock1024(t *testing.T) {
	ethConfig := ethconfig.Defaults
	backend, err := newBackend(&Config{ArchiveUrl: "http://localhost:8545", RpcUrl: "http://localhost:8549"}, &ethConfig)
	if err != nil {
		t.Fatal(err)
	}
	block := backend.CurrentBlock()
	blockNum := block.NumberU64() - 10000
	println("TestGetStateAtBehindBlock1024 ", "number:", blockNum, ", current:", block.NumberU64())
	_, _, err = backend.StateAndHeaderByNumber(context.Background(), rpc.BlockNumber(blockNum))
	if err != nil{
		t.Fatal(err)
	}
}
