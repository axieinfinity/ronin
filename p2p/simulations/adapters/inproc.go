// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package adapters

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/core/forkid"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/enr"
	"github.com/ethereum/go-ethereum/p2p/simulations/pipes"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gorilla/websocket"
)

// SimAdapter is a NodeAdapter which creates in-memory simulation nodes and
// connects them using net.Pipe
type SimAdapter struct {
	pipe       func() (net.Conn, net.Conn, error)
	mtx        sync.RWMutex
	nodes      map[enode.ID]*SimNode
	lifecycles LifecycleConstructors
}

// NewSimAdapter creates a SimAdapter which is capable of running in-memory
// simulation nodes running any of the given services (the services to run on a
// particular node are passed to the NewNode function in the NodeConfig)
// the adapter uses a net.Pipe for in-memory simulated network connections
func NewSimAdapter(services LifecycleConstructors) *SimAdapter {
	return &SimAdapter{
		pipe:       pipes.NetPipe,
		nodes:      make(map[enode.ID]*SimNode),
		lifecycles: services,
	}
}

// Name returns the name of the adapter for logging purposes
func (s *SimAdapter) Name() string {
	return "sim-adapter"
}

// NewNode returns a new SimNode using the given config
func (s *SimAdapter) NewNode(config *NodeConfig) (Node, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	id := config.ID
	// verify that the node has a private key in the config
	if config.PrivateKey == nil {
		return nil, fmt.Errorf("node is missing private key: %s", id)
	}

	// check a node with the ID doesn't already exist
	if _, exists := s.nodes[id]; exists {
		return nil, fmt.Errorf("node already exists: %s", id)
	}

	// check the services are valid
	if len(config.Lifecycles) == 0 {
		return nil, errors.New("node must have at least one service")
	}
	for _, service := range config.Lifecycles {
		if _, exists := s.lifecycles[service]; !exists {
			return nil, fmt.Errorf("unknown node service %q", service)
		}
	}

	err := config.initDummyEnode()
	if err != nil {
		return nil, err
	}

	p2pCfg := p2p.Config{
		PrivateKey:      config.PrivateKey,
		MaxPeers:        config.MaxPeers,
		NoDiscovery:     config.NoDiscovery,
		EnableMsgEvents: config.EnableMsgEvents,
	}
	if !config.DisableTCPListener {
		p2pCfg.ListenAddr = fmt.Sprintf(":%d", config.Port)
	} else {
		p2pCfg.ListenAddr = ""
	}
	if len(config.BootstrapNodeURLs) > 0 {
		for _, url := range strings.Split(config.BootstrapNodeURLs, ",") {
			if len(url) == 0 {
				continue
			}
			n, err := enode.Parse(enode.ValidSchemes, url)
			if err != nil {
				log.Warn("invalid bootstrap node URL", "url", url, "err", err)
				continue
			}
			p2pCfg.BootstrapNodes = append(p2pCfg.BootstrapNodes, n)
		}
	}

	n, err := node.New(&node.Config{
		P2P:            p2pCfg,
		ExternalSigner: config.ExternalSigner,
		Logger:         log.New("node.name", config.Name),
	})
	if err != nil {
		return nil, err
	}

	simNode := &SimNode{
		ID:      id,
		config:  config,
		node:    n,
		adapter: s,
		running: make(map[string]node.Lifecycle),
	}
	if !config.UseTCPDialer {
		n.Server().Dialer = s
	} else {
		simNode.dialer = &wrapTCPDialerStats{
			d:        &net.Dialer{Timeout: 15 * time.Second},
			resultCh: make(chan resultDial, 10000),
		}
		n.Server().Dialer = simNode.dialer
	}

	if config.EnableENRFilter {
		n.Server().SetFilter(func(id forkid.ID) error {
			var eth struct {
				ForkID forkid.ID
				Rest   []rlp.RawValue `rlp:"tail"`
			}
			if err := n.Server().Self().Record().Load(enr.WithEntry("eth", &eth)); err != nil {
				log.Warn("failed to load eth entry", "err", err)
				return err
			}

			if id == eth.ForkID {
				return nil
			}
			return forkid.ErrLocalIncompatibleOrStale
		})
	}

	s.nodes[id] = simNode
	return simNode, nil
}

// Dial implements the p2p.NodeDialer interface by connecting to the node using
// an in-memory net.Pipe
func (s *SimAdapter) Dial(ctx context.Context, dest *enode.Node) (conn net.Conn, err error) {
	node, ok := s.GetNode(dest.ID())
	if !ok {
		return nil, fmt.Errorf("unknown node: %s", dest.ID())
	}
	srv := node.Server()
	if srv == nil {
		return nil, fmt.Errorf("node not running: %s", dest.ID())
	}
	// SimAdapter.pipe is net.Pipe (NewSimAdapter)
	pipe1, pipe2, err := s.pipe()
	if err != nil {
		return nil, err
	}
	// this is simulated 'listening'
	// asynchronously call the dialed destination node's p2p server
	// to set up connection on the 'listening' side
	go srv.SetupConn(pipe1, 0, nil)
	return pipe2, nil
}

// DialRPC implements the RPCDialer interface by creating an in-memory RPC
// client of the given node
func (s *SimAdapter) DialRPC(id enode.ID) (*rpc.Client, error) {
	node, ok := s.GetNode(id)
	if !ok {
		return nil, fmt.Errorf("unknown node: %s", id)
	}
	return node.node.Attach()
}

// GetNode returns the node with the given ID if it exists
func (s *SimAdapter) GetNode(id enode.ID) (*SimNode, bool) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	node, ok := s.nodes[id]
	return node, ok
}

// SimNode is an in-memory simulation node which connects to other nodes using
// net.Pipe (see SimAdapter.Dial), running devp2p protocols directly over that
// pipe
type SimNode struct {
	ctx          context.Context
	cancel       context.CancelFunc
	lock         sync.RWMutex
	ID           enode.ID
	config       *NodeConfig
	adapter      *SimAdapter
	node         *node.Node
	running      map[string]node.Lifecycle
	client       *rpc.Client
	registerOnce sync.Once
	dialer       *wrapTCPDialerStats

	// Track different nodes discovered by the node
	discoveredNodes    sync.Map
	differentNodeCount int
}

// Close closes the underlaying node.Node to release
// acquired resources.
func (sn *SimNode) Close() error {
	return sn.node.Close()
}

// Addr returns the node's discovery address
func (sn *SimNode) Addr() []byte {
	return []byte(sn.Node().String())
}

// Node returns a node descriptor representing the SimNode
func (sn *SimNode) Node() *enode.Node {
	return sn.config.Node()
}

// Client returns an rpc.Client which can be used to communicate with the
// underlying services (it is set once the node has started)
func (sn *SimNode) Client() (*rpc.Client, error) {
	sn.lock.RLock()
	defer sn.lock.RUnlock()
	if sn.client == nil {
		return nil, errors.New("node not started")
	}
	return sn.client, nil
}

// ServeRPC serves RPC requests over the given connection by creating an
// in-memory client to the node's RPC server.
func (sn *SimNode) ServeRPC(conn *websocket.Conn) error {
	handler, err := sn.node.RPCHandler()
	if err != nil {
		return err
	}
	codec := rpc.NewFuncCodec(conn, conn.WriteJSON, conn.ReadJSON)
	handler.ServeCodec(codec, 0)
	return nil
}

// Snapshots creates snapshots of the services by calling the
// simulation_snapshot RPC method
func (sn *SimNode) Snapshots() (map[string][]byte, error) {
	sn.lock.RLock()
	services := make(map[string]node.Lifecycle, len(sn.running))
	for name, service := range sn.running {
		services[name] = service
	}
	sn.lock.RUnlock()
	if len(services) == 0 {
		return nil, errors.New("no running services")
	}
	snapshots := make(map[string][]byte)
	for name, service := range services {
		if s, ok := service.(interface {
			Snapshot() ([]byte, error)
		}); ok {
			snap, err := s.Snapshot()
			if err != nil {
				return nil, err
			}
			snapshots[name] = snap
		}
	}
	return snapshots, nil
}

// Start registers the services and starts the underlying devp2p node
func (sn *SimNode) Start(snapshots map[string][]byte) error {
	sn.lock.Lock()
	if sn.cancel != nil {
		sn.lock.Unlock()
		return errors.New("node already started")
	}

	sn.ctx, sn.cancel = context.WithCancel(context.Background())
	sn.lock.Unlock()

	// ensure we only register the services once in the case of the node
	// being stopped and then started again
	var regErr error
	sn.registerOnce.Do(func() {
		for _, name := range sn.config.Lifecycles {
			ctx := &ServiceContext{
				RPCDialer: sn.adapter,
				Config:    sn.config,
			}
			if snapshots != nil {
				ctx.Snapshot = snapshots[name]
			}
			serviceFunc := sn.adapter.lifecycles[name]
			service, err := serviceFunc(ctx, sn.node)
			if err != nil {
				regErr = err
				break
			}
			// if the service has already been registered, don't register it again.
			if _, ok := sn.running[name]; ok {
				continue
			}
			sn.running[name] = service
		}
	})
	if regErr != nil {
		return regErr
	}

	if err := sn.node.Start(); err != nil {
		return err
	}

	// create an in-process RPC client
	client, err := sn.node.Attach()
	if err != nil {
		return err
	}
	sn.lock.Lock()
	sn.client = client
	sn.lock.Unlock()

	go sn.trackDiscoveredNode()

	return nil
}

// Stop closes the RPC client and stops the underlying devp2p node
func (sn *SimNode) Stop() error {
	sn.lock.Lock()
	if sn.client != nil {
		sn.client.Close()
		sn.client = nil
	}
	if sn.cancel != nil {
		sn.cancel()
		sn.cancel = nil
	}
	sn.lock.Unlock()
	return sn.node.Close()
}

// Service returns a running service by name
func (sn *SimNode) Service(name string) node.Lifecycle {
	sn.lock.RLock()
	defer sn.lock.RUnlock()
	return sn.running[name]
}

// Services returns a copy of the underlying services
func (sn *SimNode) Services() []node.Lifecycle {
	sn.lock.RLock()
	defer sn.lock.RUnlock()
	services := make([]node.Lifecycle, 0, len(sn.running))
	for _, service := range sn.running {
		services = append(services, service)
	}
	return services
}

// ServiceMap returns a map by names of the underlying services
func (sn *SimNode) ServiceMap() map[string]node.Lifecycle {
	sn.lock.RLock()
	defer sn.lock.RUnlock()
	services := make(map[string]node.Lifecycle, len(sn.running))
	for name, service := range sn.running {
		services[name] = service
	}
	return services
}

// Server returns the underlying p2p.Server
func (sn *SimNode) Server() *p2p.Server {
	return sn.node.Server()
}

// SubscribeEvents subscribes the given channel to peer events from the
// underlying p2p.Server
func (sn *SimNode) SubscribeEvents(ch chan *p2p.PeerEvent) event.Subscription {
	srv := sn.Server()
	if srv == nil {
		panic("node not running")
	}
	return srv.SubscribeEvents(ch)
}

// NodeInfo returns information about the node
func (sn *SimNode) NodeInfo() *p2p.NodeInfo {
	server := sn.Server()
	if server == nil {
		return &p2p.NodeInfo{
			ID:    sn.ID.String(),
			Enode: sn.Node().String(),
		}
	}
	return server.NodeInfo()
}

// PeerStats returns statistics about the node's peers
func (sn *SimNode) PeerStats() *PeerStats {
	if sn.dialer == nil || sn.node.Server() == nil || sn.node.Server().UDPv4() == nil {
		return &PeerStats{}
	}

	nodesCount := 0
	sn.discoveredNodes.Range(func(_, _ interface{}) bool {
		nodesCount++
		return true
	})
	buckets := sn.node.Server().UDPv4().NodesInDHT()
	bucketSizes := make([]int, len(buckets))
	for i, bucket := range buckets {
		bucketSizes[i] = len(bucket)
	}
	return &PeerStats{
		PeerCount:                sn.node.Server().PeerCount(),
		Failed:                   sn.dialer.failed,
		Tried:                    sn.dialer.tried,
		DifferentNodesDiscovered: nodesCount,
		DHTBuckets:               bucketSizes,
	}
}

// NodesInDHT returns the nodes in the DHT buckets
func (sn *SimNode) NodesInDHT() [][]enode.Node {
	if sn.node.Server() == nil || sn.node.Server().UDPv4() == nil {
		return nil
	}
	return sn.node.Server().UDPv4().NodesInDHT()
}

// PeersInfo returns information about the node's peers
func (sn *SimNode) PeersInfo() []*p2p.PeerInfo {
	if sn.node.Server() == nil {
		return nil
	}
	return sn.node.Server().PeersInfo()
}

// trackDiscoveredNodes tracks all nodes discovered by the node and dial by wrapTCPDialerStats
func (sn *SimNode) trackDiscoveredNode() {
	if sn.dialer == nil {
		return
	}

	for {
		select {
		case <-sn.ctx.Done():
			return
		case r := <-sn.dialer.resultCh:
			if _, ok := sn.discoveredNodes.LoadOrStore(r.node, struct{}{}); !ok {
				sn.differentNodeCount++
			}
			if r.err != nil {
				log.Info("dial failed", "node", r.node, "err", r.err)
				sn.dialer.failed++
			}
			log.Info("dial tried", "from", sn.ID, "to", r.node)
			sn.dialer.tried++
		}
	}
}

// wrapTCPDialerStats is a wrapper around the net.Dialer which tracks nodes that have been tried to dial
type wrapTCPDialerStats struct {
	d        *net.Dialer
	failed   int
	tried    int
	resultCh chan resultDial
}

type resultDial struct {
	err  error
	node enode.ID
}

func (d wrapTCPDialerStats) Dial(ctx context.Context, dest *enode.Node) (net.Conn, error) {
	nodeAddr := &net.TCPAddr{IP: dest.IP(), Port: dest.TCP()}
	conn, err := d.d.DialContext(ctx, "tcp", nodeAddr.String())
	d.resultCh <- resultDial{err, dest.ID()}
	return conn, err
}
