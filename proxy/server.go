package proxy

import (
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/ethdb/httpdb"
	"github.com/ethereum/go-ethereum/node"
	"time"
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
	Port int
	RPC  string
	FreeGasProxy string
	Interval time.Duration
	TTL time.Duration
}

func NewServer(config *Config, ethConfig *ethconfig.Config, nodeConfig *node.Config) (*Server, error) {
	n, err := node.New(nodeConfig)
	if err != nil {
		return nil, err
	}
	return &Server{
		rpc: config.RPC,
		freeGasProxy: config.FreeGasProxy,
		db: httpdb.NewDB(config.RPC, config.Interval, config.TTL),
		ethConfig: ethConfig,
		node: n,
	}, nil
}

func (s *Server) Start() {
	//var apis = []rpc.API{
	//	{
	//		Namespace: "eth",
	//		Version:   "1.0",
	//		Service:   NewPublicEthereumAPI(s),
	//		Public:    true,
	//	}
	//}
	//s.node.RegisterAPIs()

}

func (s *Server) Close() {
	s.db.Close()
}
