package proxy

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rpc"
	"runtime"
	"runtime/debug"
	"time"
)

var memStats  = &runtime.MemStats{}

// Server is a proxy server that simulates rpc structures,
// it uses http database which remotely connects to other rpc nodes to get and cache data if needed
// it can be used to get data from rpc and cache them
// it also provides a VM that can process some calculation that needs VM such as: eth_estimateGas, eth_call.
type Server struct {
	config *Config
	backend *backend
	ethConfig *ethconfig.Config
	node *node.Node
}

type Config struct {
	ArchiveUrl     string
	RpcUrl         string
	FreeGasProxy   string
	DBCachedSize   int
	SafeBlockRange uint
	ResetThreshold int
	Redis          bool
	Addresses      string
	Expiration     time.Duration
}

func NewServer(config *Config, ethConfig *ethconfig.Config, nodeConfig *node.Config) (*Server, error) {
	if config.RpcUrl == "" {
		return nil, errors.New("--proxy.rpcUrl must be set")
	}
	n, err := node.New(nodeConfig)
	if err != nil {
		return nil, err
	}
	backend, err := newBackend(config, ethConfig)
	if err != nil {
		return nil, err
	}
	return &Server{
		config: config,
		backend: backend,
		ethConfig: ethConfig,
		node: n,
	}, nil
}

func (s *Server) Start() {
	var apis = []rpc.API{
		{
			Namespace: "eth",
			Version:   "1.0",
			Service:   newAPI(s.backend),
			Public:    true,
		},
		{
			Namespace: "net",
			Version: "1.0",
			Service: ethapi.NewPublicNetAPI(&p2p.Server{}, s.ethConfig.NetworkId),
			Public: true,
		},
	}
	s.node.RegisterAPIs(apis)
	if err := s.node.StartRPC(); err != nil {
		panic(err)
	}
	go func() {
		for {
			select {
			case <-time.Tick(30*time.Second):
				PrintMemUsage()
			case <-time.Tick(time.Minute):
				debug.FreeOSMemory()
			}
		}
	} ()
	s.node.Wait()
}

func (s *Server) Close() {
	s.backend.db.Close()
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	runtime.ReadMemStats(memStats)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	log.Info(fmt.Sprintf("%s%s%s%s",
		fmt.Sprintf("Alloc = %v MiB", bToMb(memStats.Alloc)),
		fmt.Sprintf("\n\tTotalAlloc = %v MiB", bToMb(memStats.TotalAlloc)),
		fmt.Sprintf("\n\tSys = %v MiB", bToMb(memStats.Sys)),
		fmt.Sprintf("\n\tNumGC = %v\n", memStats.NumGC),
	))

}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
