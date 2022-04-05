package proxy

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/rpc"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"
)

func TestEthCall(t *testing.T) {
	ethConfig := ethconfig.Defaults
	ethConfig.RPCEVMTimeout = time.Second*50
	backend, err := NewBackend(&Config{ArchiveUrl: "https://api-archived.roninchain.com/rpc", RpcUrl: "https://api-archived.roninchain.com/rpc"}, &ethConfig)
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

func TestEthBlockNumber(t *testing.T) {
	ethConfig := ethconfig.Defaults
	ethConfig.RPCEVMTimeout = time.Second*50
	backend, err := NewBackend(&Config{ArchiveUrl: "https://api-archived.roninchain.com/rpc", RpcUrl: "http://34.121.216.144:8545"}, &ethConfig)
	if err != nil {
		t.Fatal(err)
	}
	block := backend.CurrentBlock()
	println(block.NumberU64())

	time.Sleep(3*time.Second)

	// second call
	block = backend.CurrentBlock()
	println(block.NumberU64())
}

func TestGetStateAtBlock(t *testing.T) {
	ethConfig := ethconfig.Defaults
	backend, err := NewBackend(&Config{ArchiveUrl: "http://localhost:8545", RpcUrl: "http://localhost:8549"}, &ethConfig)
	if err != nil {
		t.Fatal(err)
	}
	println("start getting state by number")
	_, _, err = backend.StateAndHeaderByNumber(context.Background(), rpc.BlockNumber(1000))
	if err != nil{
		t.Fatal(err)
	}
}

func blockNumber() {
	client, err := rpc.DialHTTP("http://localhost:8545")
	if err != nil {
		return
	}
	var lastBlock string
	err = client.CallContext(context.Background(), &lastBlock, "eth_blockNumber")
	if err != nil {
		println(err.Error())
		return
	}
	height := hexutil.MustDecodeUint64(lastBlock)
	println(height)
	client.Close()
}

func ethCall1() {
	client, err := rpc.DialHTTP("http://localhost:8545")
	if err != nil {
		return
	}
	defer client.Close()
	var hex hexutil.Bytes
	err = client.CallContext(context.Background(), &hex, "eth_call", map[string]interface{}{
		"to":   "0x213073989821f738a7ba3520c3d31a1f9ad31bbd",
		"data": "0x4d51bfc40000000000000000000000002b83d45a4ce40da980fec4eb2749f171acd410ee000000000000000000000000c99a6a985ed2cac1ef41640596c5a5f9f4e19ef5000000000000000000000000000000000000000000000000005fec5b60ef800000000000000000000000000000000000000000000000000000000000006e151899afb0f91b262eede7980e6e61131f9592cd39e4be82bd5a1fdc8d8c39918e5f",
	}, "latest")
	if err != nil {
		println(err.Error())
		return
	}
}

func ethCall2() {
	client, err := rpc.DialHTTP("http://localhost:8545")
	if err != nil {
		return
	}
	defer client.Close()
	var hex hexutil.Bytes
	err = client.CallContext(context.Background(), &hex, "eth_call", map[string]interface{}{
		"to":   "0x32950db2a7164aE833121501C797D79E7B79d74C",
		"data": "0x3e2156aa00000000000000000000000000000000000000000000000000000000005091fa",
	}, "latest")
	if err != nil {
		println(err.Error())
		return
	}
}

func TestMultipleCall(t *testing.T) {
	var counter int32
	ch := make(chan int, 150)
	countCh := make(chan int)
	exec := []func() {
		blockNumber, ethCall1, ethCall2,
	}
	go func () {
		for {
			ch <- 1
			go func() {
				i := rand.Int31n(3)
				exec[i]()
				<-ch
				countCh <- 1
			} ()
		}
	} ()
	for {
		select {
		case <-countCh:
			val := atomic.AddInt32(&counter, 1)
			if val >= 1000000 {
				return
			}
		}
	}
}
