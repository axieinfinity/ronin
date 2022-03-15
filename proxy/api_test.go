package proxy

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/rpc"
	"os"
	"sync/atomic"
	"testing"
	"time"
)

func TestEthCall(t *testing.T) {
	ethConfig := ethconfig.Defaults
	ethConfig.RPCEVMTimeout = time.Second*50
	backend, err := newBackend(&Config{ArchiveUrl: "https://api-archived.roninchain.com/rpc", RpcUrl: "https://api-archived.roninchain.com/rpc"}, &ethConfig)
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
	backend, err := newBackend(&Config{ArchiveUrl: "https://api-archived.roninchain.com/rpc", RpcUrl: "http://34.121.216.144:8545"}, &ethConfig)
	if err != nil {
		t.Fatal(err)
	}
	//api := newAPI(backend)
	//to := common.HexToAddress("0xc99a6a985ed2cac1ef41640596c5a5f9f4e19ef5")
	//blockNumber := rpc.LatestBlockNumber
	//data := hexutil.Bytes(common.Hex2Bytes("0x70a082310000000000000000000000006755b9c63779d5b16cfa434e9e1c018689b03e45"))
	//rs, err := api.Call(context.Background(), ethapi.TransactionArgs{
	//	Data: &data,
	//	To: &to,
	//}, rpc.BlockNumberOrHash{BlockNumber: &blockNumber}, nil)
	//if err != nil && err.Error() != "execution reverted" {
	//	t.Fatal(err)
	//}
	block := backend.CurrentBlock()
	println(block.NumberU64())

	time.Sleep(3*time.Second)

	// second call
	block = backend.CurrentBlock()
	println(block.NumberU64())
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

func TestViaPublicAPI(t *testing.T) {
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlDebug, log.StreamHandler(os.Stdout, log.TerminalFormat(true))))
	log.New()
	metrics.Enabled = true
	ethConfig := ethconfig.Defaults
	backend, err := newBackend(&Config{ArchiveUrl: "https://api-archived.roninchain.com/rpc", RpcUrl: "http://34.121.216.144:8545"}, &ethConfig)
	if err != nil {
		t.Fatal(err)
	}
	var (
		requestCounter = metrics.GetOrRegisterCounter("cache/request", nil)
		cacheHitCounter = metrics.GetOrRegisterCounter("cache/request/hit", nil)
		cacheItemsCounter = metrics.GetOrRegisterCounter("cache/items", nil)
		cacheItemsSizeCounter = metrics.GetOrRegisterCounter("cache/items/size", nil)
	)
	api := newAPI(backend)
	var counter int32
	ch := make(chan interface{})
	go func() {
		for {
			api.BlockNumber()
			ch <- 1
		}
	} ()
	go func () {
		for {
			latest := rpc.LatestBlockNumber
			address := common.HexToAddress("0x213073989821f738a7ba3520c3d31a1f9ad31bbd")
			data := hexutil.Bytes(common.Hex2Bytes("0x4d51bfc40000000000000000000000002b83d45a4ce40da980fec4eb2749f171acd410ee000000000000000000000000c99a6a985ed2cac1ef41640596c5a5f9f4e19ef5000000000000000000000000000000000000000000000000005fec5b60ef800000000000000000000000000000000000000000000000000000000000006e151899afb0f91b262eede7980e6e61131f9592cd39e4be82bd5a1fdc8d8c39918e5f"))
			api.Call(context.Background(), ethapi.TransactionArgs{
				To: &address,
				Data: &data,
			}, rpc.BlockNumberOrHash{BlockNumber: &latest}, nil)
			ch <- 1
		}
	} ()
	ticker := time.NewTicker(3*time.Second)
	for {
		select {
		case <-ch:
			//log.Info(fmt.Sprintf("new counter: %d", atomic.AddInt32(&counter, 1)))
			atomic.AddInt32(&counter, 1)
			if counter > 1000099999 {
				return
			}
		case <-ticker.C:
			log.Info(fmt.Sprintf("============================================"))
			log.Info(fmt.Sprintf("requestCounter: %d", requestCounter.Count()))
			log.Info(fmt.Sprintf("cacheHitCounter: %d", cacheHitCounter.Count()))
			log.Info(fmt.Sprintf("%v %s", (float64(cacheHitCounter.Count())/float64(requestCounter.Count()))*100, "%"))
			log.Info(fmt.Sprintf("cacheItemsCounter: %d", cacheItemsCounter.Count()))
			log.Info(fmt.Sprintf("cacheItemsSizeCounter: %d", cacheItemsSizeCounter.Count()))
			log.Info(fmt.Sprintf("============================================"))
		}
	}
}

func TestCurrentBlock(t *testing.T) {
	//ethConfig := ethconfig.Defaults
	//backend, err := newBackend(&Config{ArchiveUrl: "https://api-archived.roninchain.com/rpc", RpcUrl: "http://34.121.216.144:8545"}, &ethConfig)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//var counter int32
	ch := make(chan int, 100)
	errCh := make(chan error, 1)
	for {
		go func() {
			ch <- 1
			client, err := rpc.DialHTTP("http://localhost:8560")
			if err != nil {
				return
			}
			defer client.Close()
			<- ch
			var lastBlock string
			err = client.CallContext(context.Background(), &lastBlock, "eth_blockNumber")
			if err != nil {
				errCh <- err
				return
			}
			height := hexutil.MustDecodeUint64(lastBlock)
			println(height)
		}()
	}
}

func TestLoopCall(t *testing.T) {
	client, err := rpc.DialHTTP("http://localhost:8545")
	if err != nil {
		t.Fatal(err)
	}
	//for i:=0; i < 1000000; i++ {
	var hex hexutil.Bytes
	err = client.CallContext(context.Background(), &hex, "eth_call", map[string]interface{}{
		"to": "0x213073989821f738a7ba3520c3d31a1f9ad31bbd",
		"data": "0x4d51bfc40000000000000000000000002b83d45a4ce40da980fec4eb2749f171acd410ee000000000000000000000000c99a6a985ed2cac1ef41640596c5a5f9f4e19ef5000000000000000000000000000000000000000000000000005fec5b60ef800000000000000000000000000000000000000000000000000000000000006e151899afb0f91b262eede7980e6e61131f9592cd39e4be82bd5a1fdc8d8c39918e5f",
	}, "latest")
	if err != nil {
		println(err.Error())
		return
	}
	println(fmt.Sprintf("result: %s", common.Bytes2Hex(hex)))
	//}

}
