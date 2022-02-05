package proxy

import (
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/ethdb/httpdb"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rpc"
)

// Server is a proxy server that simulates rpc structures,
// it uses http database which remotely connects to other rpc nodes to get and cache data if needed
// it can be used to get data from rpc and cache them with ttl
// it also provides a VM that can process some calculation that needs VM such as: eth_estimateGas, eth_call.
type Server struct {
	rpc string
	freeGasProxy string
	db ethdb.Database
	ethConfig *ethconfig.Config
	node *node.Node
}

type Config struct {
	RPC  string
	FreeGasProxy string
	DBCachedSize int
}

func NewServer(config *Config, ethConfig *ethconfig.Config, nodeConfig *node.Config) (*Server, error) {
	n, err := node.New(nodeConfig)
	if err != nil {
		return nil, err
	}
	return &Server{
		rpc: config.RPC,
		freeGasProxy: config.FreeGasProxy,
		db: httpdb.NewDB(config.RPC, config.DBCachedSize),
		ethConfig: ethConfig,
		node: n,
	}, nil
}

func (s *Server) Start() {
	backend, err := newBackend(s.db, s.ethConfig, s.rpc, s.freeGasProxy)
	if err != nil {
		panic(err)
	}
	var apis = []rpc.API{
		{
			Namespace: "eth",
			Version:   "1.0",
			Service:   newAPI(backend),
			Public:    true,
		},
	}
	s.node.RegisterAPIs(apis)
	if err = s.node.StartRPC(); err != nil {
		panic(err)
	}
	s.node.Wait()
}

func (s *Server) Close() {
	s.db.Close()
}
