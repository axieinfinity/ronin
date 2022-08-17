package proxy

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/consortium"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/go-redis/redis/v8"
	"time"
)

// Server is a proxy server that simulates rpc structures,
// it uses http database which remotely connects to other rpc nodes to get and cache data if needed
// it can be used to get data from rpc and cache them
// it also provides a VM that can process some calculation that needs VM such as: eth_estimateGas, eth_call.
type Server struct {
	config    *Config
	backend   *backend
	ethConfig *ethconfig.Config
	node      *node.Node
}

type Config struct {
	ArchiveUrl     string
	RpcUrl         string
	FreeGasProxy   string
	DBCachedSize   int
	SafeBlockRange uint
	Redis          *Redis
}

type Redis struct {
	Options    *redis.Options
	Expiration time.Duration
}

func NewServer(config *Config, ethConfig *ethconfig.Config, nodeConfig *node.Config) (*Server, error) {
	if config.RpcUrl == "" {
		return nil, errors.New("--proxy.rpcUrl must be set")
	}
	n, err := node.New(nodeConfig)
	if err != nil {
		return nil, err
	}
	backend, err := NewBackend(config, ethConfig)
	if err != nil {
		return nil, err
	}
	return &Server{
		config:    config,
		backend:   backend,
		ethConfig: ethConfig,
		node:      n,
	}, nil
}

func (s *Server) Start() {
	simBackend := backends.NewSimulatedBackendWithBC(s.backend.db)
	engine := consortium.New(&params.ChainConfig{}, s.backend.db, nil, simBackend, common.Hash{})
	var apis = []rpc.API{
		{
			Namespace: "eth",
			Version:   "1.0",
			Service:   newAPI(s.backend),
			Public:    true,
		},
		{
			Namespace: "net",
			Version:   "1.0",
			Service:   ethapi.NewPublicNetAPI(&p2p.Server{}, s.ethConfig.NetworkId),
			Public:    true,
		},
	}
	apis = append(apis, engine.APIs(s.backend)...)
	s.node.RegisterAPIs(apis)
	if err := s.node.StartRPC(); err != nil {
		panic(err)
	}
	s.node.Wait()
}

func (s *Server) Close() {
	s.backend.db.Close()
}
