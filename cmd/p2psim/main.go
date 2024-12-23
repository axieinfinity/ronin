// Copyright 2017 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

// p2psim provides a command-line client for a simulation HTTP API.
//
// Here is an example of creating a 2 node network with the first node
// connected to the second:
//
//	$ p2psim node create
//	Created node01
//
//	$ p2psim node start node01
//	Started node01
//
//	$ p2psim node create
//	Created node02
//
//	$ p2psim node start node02
//	Started node02
//
//	$ p2psim node connect node01 node02
//	Connected node01 to node02
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"text/tabwriter"
	"time"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/simulations"
	"github.com/ethereum/go-ethereum/p2p/simulations/adapters"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/urfave/cli/v2"
)

var client *simulations.Client

func main() {
	app := cli.NewApp()
	app.Usage = "devp2p simulation command-line client"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "api",
			Value:   "http://localhost:8888",
			Usage:   "simulation API URL",
			EnvVars: []string{"P2PSIM_API_URL"},
		},
	}
	app.Before = func(ctx *cli.Context) error {
		client = simulations.NewClient(ctx.String("api"))
		return nil
	}
	app.Commands = []*cli.Command{
		{
			Name:   "show",
			Usage:  "show network information",
			Action: showNetwork,
		},
		{
			Name:   "events",
			Usage:  "stream network events",
			Action: streamNetwork,
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "current",
					Usage: "get existing nodes and conns first",
				},
				&cli.StringFlag{
					Name:  "filter",
					Value: "",
					Usage: "message filter",
				},
			},
		},
		{
			Name:   "snapshot",
			Usage:  "create a network snapshot to stdout",
			Action: createSnapshot,
		},
		{
			Name:   "load",
			Usage:  "load a network snapshot from stdin",
			Action: loadSnapshot,
		},
		{
			Name:  "network",
			Usage: "manage the simulation network",
			Subcommands: []*cli.Command{
				{
					Name:   "start",
					Usage:  "start all nodes in the network",
					Action: startNetwork,
				},
				{
					Name:   "peer-stats",
					Usage:  "show peer stats",
					Action: getNetworkPeerStats,
				},
				{
					Name:   "dht",
					Usage:  "Get all nodes in the DHT of all nodes",
					Action: getAllDHT,
				},
				{
					Name:   "peers",
					Usage:  "Get all peers of all nodes",
					Action: getAllNodePeersInfo,
				},
			},
		},
		{
			Name:   "node",
			Usage:  "manage simulation nodes",
			Action: listNodes,
			Subcommands: []*cli.Command{
				{
					Name:   "list",
					Usage:  "list nodes",
					Action: listNodes,
				},
				{
					Name:   "create",
					Usage:  "create a node",
					Action: createNode,
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Value: "",
							Usage: "node name",
						},
						&cli.StringFlag{
							Name:  "services",
							Value: "",
							Usage: "node services (comma separated)",
						},
						&cli.StringFlag{
							Name:  "key",
							Value: "",
							Usage: "node private key (hex encoded)",
						},
						&cli.BoolFlag{
							Name:  "sim.dialer",
							Usage: "Use the simulation dialer",
						},
						&cli.BoolFlag{
							Name:  "fake.iplistener",
							Usage: "Use the fake listener to random remote ip when accepting connections",
						},
						&cli.BoolFlag{
							Name:  "start",
							Usage: "start the node after creating successfully",
						},
						&cli.BoolFlag{
							Name:  "autofill.bootnodes",
							Usage: "autofill bootnodes with existing bootnodes from manager",
						},
						&cli.StringFlag{
							Name:  "node.type",
							Value: "default",
							Usage: "Set node type (default, outbound, dirty, bootnode)",
						},
						&cli.BoolFlag{
							Name:  "enable.enrfilter",
							Usage: "Enable ENR filter when adding nodes to the DHT",
						},
						&cli.BoolFlag{
							Name:  "only.outbound",
							Usage: "Only allow outbound connections",
						},
						utils.NoDiscoverFlag,
						utils.BootnodesFlag,
						utils.MaxPeersFlag,
					},
				},
				{
					Name:   "create-multi",
					Usage:  "create a node",
					Action: createMultiNode,
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Value: "",
							Usage: "node name",
						},
						&cli.IntFlag{
							Name:  "count",
							Value: 1,
							Usage: "number of nodes to create",
						},
						&cli.StringFlag{
							Name:  "services",
							Value: "",
							Usage: "node services (comma separated)",
						},
						&cli.BoolFlag{
							Name:  "sim.dialer",
							Usage: "Use the simulation dialer",
						},
						&cli.BoolFlag{
							Name:  "fake.iplistener",
							Usage: "Use the fake listener to random remote ip when accepting connections",
						},
						&cli.BoolFlag{
							Name:  "start",
							Usage: "start the node after creating successfully",
						},
						&cli.BoolFlag{
							Name:  "autofill.bootnodes",
							Usage: "autofill bootnodes with existing bootnodes from manager",
						},
						&cli.StringFlag{
							Name:  "node.type",
							Value: "default",
							Usage: "Set node type (default, outbound, dirty, bootnode)",
						},
						&cli.BoolFlag{
							Name:  "enable.enrfilter",
							Usage: "Enable ENR filter when adding nodes to the DHT",
						},
						&cli.DurationFlag{
							Name:  "interval",
							Usage: "create interval",
						},
						&cli.IntFlag{
							Name:  "dirty.rate",
							Usage: "Rate of dirty nodes",
						},
						&cli.IntFlag{
							Name:  "only.outbound.rate",
							Usage: "Rate of nodes that only allow outbound connections",
						},
						&cli.BoolFlag{
							Name:  "only.outbound",
							Usage: "Only allow outbound connections",
						},
						utils.NoDiscoverFlag,
						utils.BootnodesFlag,
						utils.MaxPeersFlag,
					},
				},
				{
					Name:      "show",
					ArgsUsage: "<node>",
					Usage:     "show node information",
					Action:    showNode,
				},
				{
					Name:      "start",
					ArgsUsage: "<node>",
					Usage:     "start a node",
					Action:    startNode,
				},
				{
					Name:      "stop",
					ArgsUsage: "<node>",
					Usage:     "stop a node",
					Action:    stopNode,
				},
				{
					Name:      "connect",
					ArgsUsage: "<node> <peer>",
					Usage:     "connect a node to a peer node",
					Action:    connectNode,
				},
				{
					Name:      "disconnect",
					ArgsUsage: "<node> <peer>",
					Usage:     "disconnect a node from a peer node",
					Action:    disconnectNode,
				},
				{
					Name:      "rpc",
					ArgsUsage: "<node> <method> [<args>]",
					Usage:     "call a node RPC method",
					Action:    rpcNode,
					Flags: []cli.Flag{
						&cli.BoolFlag{
							Name:  "subscribe",
							Usage: "method is a subscription",
						},
					},
				},
				{
					Name:      "peer-stats",
					Usage:     "show peer stats",
					ArgsUsage: "<node>",
					Action:    getNodePeerStats,
				},
			},
		},
		{
			Name:   "log-stats",
			Usage:  "log peer stats to a CSV file",
			Action: startLogStats,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "file",
					Usage: "output file",
					Value: "stats.csv",
				},
				&cli.DurationFlag{
					Name:  "interval",
					Usage: "log interval",
					Value: 15 * time.Second,
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func showNetwork(ctx *cli.Context) error {
	if ctx.Args().Len() != 0 {
		return cli.ShowCommandHelp(ctx, ctx.Command.Name)
	}
	network, err := client.GetNetwork()
	if err != nil {
		return err
	}
	w := tabwriter.NewWriter(ctx.App.Writer, 1, 2, 2, ' ', 0)
	defer w.Flush()
	fmt.Fprintf(w, "NODES\t%d\n", len(network.Nodes))
	fmt.Fprintf(w, "CONNS\t%d\n", len(network.Conns))
	return nil
}

func streamNetwork(ctx *cli.Context) error {
	if ctx.Args().Len() != 0 {
		return cli.ShowCommandHelp(ctx, ctx.Command.Name)
	}
	events := make(chan *simulations.Event)
	sub, err := client.SubscribeNetwork(events, simulations.SubscribeOpts{
		Current: ctx.Bool("current"),
		Filter:  ctx.String("filter"),
	})
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()
	enc := json.NewEncoder(ctx.App.Writer)
	for {
		select {
		case event := <-events:
			if err := enc.Encode(event); err != nil {
				return err
			}
		case err := <-sub.Err():
			return err
		}
	}
}

func createSnapshot(ctx *cli.Context) error {
	if ctx.Args().Len() != 0 {
		return cli.ShowCommandHelp(ctx, ctx.Command.Name)
	}
	snap, err := client.CreateSnapshot()
	if err != nil {
		return err
	}
	return json.NewEncoder(os.Stdout).Encode(snap)
}

func loadSnapshot(ctx *cli.Context) error {
	if ctx.Args().Len() != 0 {
		return cli.ShowCommandHelp(ctx, ctx.Command.Name)
	}
	snap := &simulations.Snapshot{}
	if err := json.NewDecoder(os.Stdin).Decode(snap); err != nil {
		return err
	}
	return client.LoadSnapshot(snap)
}

func listNodes(ctx *cli.Context) error {
	if ctx.Args().Len() != 0 {
		return cli.ShowCommandHelp(ctx, ctx.Command.Name)
	}
	nodes, err := client.GetNodes()
	if err != nil {
		return err
	}
	w := tabwriter.NewWriter(ctx.App.Writer, 1, 2, 2, ' ', 0)
	defer w.Flush()
	fmt.Fprintf(w, "NAME\tPROTOCOLS\tID\n")
	for _, node := range nodes {
		fmt.Fprintf(w, "%s\t%s\t%s\n", node.Name, strings.Join(protocolList(node), ","), node.ID)
	}
	return nil
}

func protocolList(node *p2p.NodeInfo) []string {
	protos := make([]string, 0, len(node.Protocols))
	for name := range node.Protocols {
		protos = append(protos, name)
	}
	return protos
}

func createNode(ctx *cli.Context) error {
	if ctx.Args().Len() != 0 {
		return cli.ShowCommandHelp(ctx, ctx.Command.Name)
	}
	config := adapters.RandomNodeConfig()
	config.Name = ctx.String("name")
	if key := ctx.String("key"); key != "" {
		privKey, err := crypto.HexToECDSA(key)
		if err != nil {
			return err
		}
		config.ID = enode.PubkeyToIDV4(&privKey.PublicKey)
		config.PrivateKey = privKey
	}
	if ctx.Bool(utils.NoDiscoverFlag.Name) {
		config.NoDiscovery = true
	}
	if ctx.Bool("sim.dialer") {
		config.UseTCPDialer = false
	} else {
		config.UseTCPDialer = true
	}
	if ctx.Bool("fake.iplistener") {
		config.UseFakeIPListener = true
	}
	config.BootstrapNodeURLs = ctx.String(utils.BootnodesFlag.Name)
	if ctx.Bool("autofill.bootnodes") {
		bootnodeURLs, err := getBootnodes()
		if err != nil {
			return err
		}
		if bootnodeURLs != "" {
			config.BootstrapNodeURLs += "," + bootnodeURLs
		}
	}
	config.MaxPeers = ctx.Int(utils.MaxPeersFlag.Name)
	config.DisableTCPListener = ctx.Bool("only.outbound")
	config.EnableENRFilter = ctx.Bool("enable.enrfilter")
	if services := ctx.String("services"); services != "" {
		config.Lifecycles = strings.Split(services, ",")
	}
	node, err := client.CreateNode(config)
	if err != nil {
		return err
	}
	fmt.Fprintln(ctx.App.Writer, "Created", node.Name)

	// Start node if needed
	if ctx.Bool("start") {
		if err := client.StartNode(node.Name); err != nil {
			return err
		}
		fmt.Fprintln(ctx.App.Writer, "Started", node.Name)
	}

	return nil
}

func getBootnodes() (string, error) {
	nodes, err := client.GetNodes()
	if err != nil {
		return "", err
	}

	bootnodes := make([]string, 0)
	for _, node := range nodes {
		if strings.HasPrefix(node.Name, "bootnode") {
			bootnodes = append(bootnodes, node.Enode)
		}
	}

	return strings.Join(bootnodes, ","), nil
}

func createMultiNode(ctx *cli.Context) error {
	if ctx.Args().Len() != 0 {
		return cli.ShowCommandHelp(ctx, ctx.Command.Name)
	}

	t := time.Now()

	createInterval := ctx.Duration("interval")
	bootNodeURLs := ctx.String(utils.BootnodesFlag.Name)
	if ctx.Bool("autofill.bootnodes") {
		existedBootnodeURLs, err := getBootnodes()
		if err != nil {
			return err
		}
		if existedBootnodeURLs != "" {
			bootNodeURLs += "," + existedBootnodeURLs
		}
	}

	// Create nodes
	count := ctx.Int("count")
	outboundRate := ctx.Int("only.outbound.rate")
	dirtyRate := ctx.Int("dirty.rate")
	per := make([]int, 0)
	for i := 0; i < count; i++ {
		if i < outboundRate*count/100 {
			per = append(per, 1)
		} else if i < (outboundRate+dirtyRate)*count/100 {
			per = append(per, 2)
		} else {
			per = append(per, 0)
		}
	}
	rand.Shuffle(len(per), func(i, j int) { per[i], per[j] = per[j], per[i] })

	isBootnode := ctx.String("node.type") == "bootnode"

	for i := 0; i < count; i++ {
		var nodeName string
		if isBootnode {
			nodeName = fmt.Sprintf("bootnode-%d-%d", t.Unix(), i)
			ctx.Set(utils.BootnodesFlag.Name, "")
		} else {
			nodeType := per[i%len(per)]
			switch nodeType {
			case 1:
				ctx.Set("only.outbound", "true")
				ctx.Set("node.type", "outbound")
				ctx.Set("services", "valid")
				nodeName = fmt.Sprintf("outbound-%d-%d", t.Unix(), i)
			case 2:
				ctx.Set("only.outbound", "false")
				ctx.Set("node.type", "dirty")
				ctx.Set("services", "invalid")
				nodeName = fmt.Sprintf("dirty-%d-%d", t.Unix(), i)
			default:
				ctx.Set("only.outbound", "false")
				ctx.Set("node.type", "default")
				ctx.Set("services", "valid")
				nodeName = fmt.Sprintf("node-%d-%d", t.Unix(), i)
			}
		}
		ctx.Set("name", nodeName)
		for {
			if err := createNode(ctx); err != nil {
				fmt.Fprintln(ctx.App.Writer, "Failed to create node", nodeName, err)
				// Try to create the node again
				client.DeleteNode(nodeName)
				time.Sleep(500 * time.Millisecond)
			} else {
				break
			}
		}
		if createInterval > 0 {
			time.Sleep(createInterval)
		}
	}

	return nil
}

func showNode(ctx *cli.Context) error {

	if ctx.Args().Len() != 1 {
		return cli.ShowCommandHelp(ctx, ctx.Command.Name)
	}
	nodeName := ctx.Args().First()
	node, err := client.GetNode(nodeName)
	if err != nil {
		return err
	}
	w := tabwriter.NewWriter(ctx.App.Writer, 1, 2, 2, ' ', 0)
	defer w.Flush()
	fmt.Fprintf(w, "NAME\t%s\n", node.Name)
	fmt.Fprintf(w, "PROTOCOLS\t%s\n", strings.Join(protocolList(node), ","))
	fmt.Fprintf(w, "ID\t%s\n", node.ID)
	fmt.Fprintf(w, "ENODE\t%s\n", node.Enode)
	for name, proto := range node.Protocols {
		fmt.Fprintln(w)
		fmt.Fprintf(w, "--- PROTOCOL INFO: %s\n", name)
		fmt.Fprintf(w, "%v\n", proto)
		fmt.Fprintf(w, "---\n")
	}
	return nil
}

func startNode(ctx *cli.Context) error {
	args := ctx.Args()
	if ctx.Args().Len() != 1 {
		return cli.ShowCommandHelp(ctx, ctx.Command.Name)
	}
	nodeName := args.First()
	if err := client.StartNode(nodeName); err != nil {
		return err
	}
	fmt.Fprintln(ctx.App.Writer, "Started", nodeName)
	return nil
}

func stopNode(ctx *cli.Context) error {
	args := ctx.Args()
	if args.Len() != 1 {
		return cli.ShowCommandHelp(ctx, ctx.Command.Name)
	}
	nodeName := args.First()
	if err := client.StopNode(nodeName); err != nil {
		return err
	}
	fmt.Fprintln(ctx.App.Writer, "Stopped", nodeName)
	return nil
}

func connectNode(ctx *cli.Context) error {
	args := ctx.Args()
	if args.Len() != 2 {
		return cli.ShowCommandHelp(ctx, ctx.Command.Name)
	}
	nodeName := args.Get(0)
	peerName := args.Get(1)
	if err := client.ConnectNode(nodeName, peerName); err != nil {
		return err
	}
	fmt.Fprintln(ctx.App.Writer, "Connected", nodeName, "to", peerName)
	return nil
}

func disconnectNode(ctx *cli.Context) error {
	args := ctx.Args()
	if args.Len() != 2 {
		return cli.ShowCommandHelp(ctx, ctx.Command.Name)
	}
	nodeName := args.Get(0)
	peerName := args.Get(1)
	if err := client.DisconnectNode(nodeName, peerName); err != nil {
		return err
	}
	fmt.Fprintln(ctx.App.Writer, "Disconnected", nodeName, "from", peerName)
	return nil
}

func rpcNode(ctx *cli.Context) error {
	args := ctx.Args()
	if args.Len() < 2 {
		return cli.ShowCommandHelp(ctx, ctx.Command.Name)
	}
	nodeName := args.Get(0)
	method := args.Get(1)
	rpcClient, err := client.RPCClient(context.Background(), nodeName)
	if err != nil {
		return err
	}
	if ctx.Bool("subscribe") {
		return rpcSubscribe(rpcClient, ctx.App.Writer, method, args.Slice()[3:]...)
	}
	var result interface{}
	params := make([]interface{}, len(args.Slice()[3:]))
	for i, v := range args.Slice()[3:] {
		params[i] = v
	}
	if err := rpcClient.Call(&result, method, params...); err != nil {
		return err
	}
	return json.NewEncoder(ctx.App.Writer).Encode(result)
}

func rpcSubscribe(client *rpc.Client, out io.Writer, method string, args ...string) error {
	parts := strings.SplitN(method, "_", 2)
	namespace := parts[0]
	method = parts[1]
	ch := make(chan interface{})
	subArgs := make([]interface{}, len(args)+1)
	subArgs[0] = method
	for i, v := range args {
		subArgs[i+1] = v
	}
	sub, err := client.Subscribe(context.Background(), namespace, ch, subArgs...)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()
	enc := json.NewEncoder(out)
	for {
		select {
		case v := <-ch:
			if err := enc.Encode(v); err != nil {
				return err
			}
		case err := <-sub.Err():
			return err
		}
	}
}

func startNetwork(ctx *cli.Context) error {
	if ctx.Args().Len() != 0 {
		return cli.ShowCommandHelp(ctx, ctx.Command.Name)
	}
	if err := client.StartNetwork(); err != nil {
		return err
	}
	fmt.Fprintln(ctx.App.Writer, "Started network")
	return nil
}

func getNodePeerStats(ctx *cli.Context) error {
	if ctx.Args().Len() != 1 {
		return cli.ShowCommandHelp(ctx, ctx.Command.Name)
	}
	nodeName := ctx.Args().Get(0)
	stats, err := client.GetNodePeerStats(nodeName)
	if err != nil {
		return err
	}
	fmt.Fprintln(ctx.App.Writer, "Peer stats of", ctx.String("node"))
	fmt.Fprintln(ctx.App.Writer, "Peer count: ", stats.PeerCount)
	fmt.Fprintln(ctx.App.Writer, "Tried: ", stats.Tried)
	fmt.Fprintln(ctx.App.Writer, "Failed: ", stats.Failed)
	fmt.Fprintln(ctx.App.Writer, "Nodes count: ", stats.DifferentNodesDiscovered)
	fmt.Fprintln(ctx.App.Writer, "DHT: ", stats.DHTBuckets)
	return nil
}

func getNetworkPeerStats(ctx *cli.Context) error {
	if ctx.Args().Len() != 0 {
		return cli.ShowCommandHelp(ctx, ctx.Command.Name)
	}
	stats, err := client.GetAllNodePeerStats()
	if err != nil {
		return err
	}
	for nodeID, stats := range stats {
		fmt.Fprintln(ctx.App.Writer, "Peer stats of", nodeID)
		fmt.Fprintln(ctx.App.Writer, "Peer count: ", stats.PeerCount)
		fmt.Fprintln(ctx.App.Writer, "Tried: ", stats.Tried)
		fmt.Fprintln(ctx.App.Writer, "Failed: ", stats.Failed)
		fmt.Fprintln(ctx.App.Writer, "Nodes count: ", stats.DifferentNodesDiscovered)
		fmt.Fprintln(ctx.App.Writer, "DHT: ", stats.DHTBuckets)
	}
	return nil
}

func getAllDHT(ctx *cli.Context) error {
	if ctx.Args().Len() != 0 {
		return cli.ShowCommandHelp(ctx, ctx.Command.Name)
	}

	nodes, err := client.GetNodes()
	if err != nil {
		return err
	}
	nodeID2Name := make(map[string]string)
	for _, node := range nodes {
		nodeID2Name[node.ID] = node.Name
	}

	dht, err := client.GetAllNodeDHT()
	if err != nil {
		return err
	}
	for nodeName, buckets := range dht {
		fmt.Fprintf(ctx.App.Writer, "%s: ", nodeName)
		for _, bucket := range buckets {
			fmt.Fprintf(ctx.App.Writer, "[")
			for _, node := range bucket {
				fmt.Fprintf(ctx.App.Writer, "%s ", nodeID2Name[node.ID().String()])
			}
			fmt.Fprintf(ctx.App.Writer, "],")
		}
		fmt.Fprintf(ctx.App.Writer, "\n")
	}
	return nil
}

func getAllNodePeersInfo(ctx *cli.Context) error {
	if ctx.Args().Len() != 0 {
		return cli.ShowCommandHelp(ctx, ctx.Command.Name)
	}

	nodes, err := client.GetNodes()
	if err != nil {
		return err
	}
	nodeID2Name := make(map[string]string)
	for _, node := range nodes {
		nodeID2Name[node.ID] = node.Name
	}

	peers, err := client.GetAllNodePeersInfo()
	if err != nil {
		return err
	}
	for nodeName, peerInfos := range peers {
		fmt.Fprintf(ctx.App.Writer, "%s: ", nodeName)
		for _, peerInfo := range peerInfos {
			fmt.Fprintf(ctx.App.Writer, "(%s %v), ", nodeID2Name[peerInfo.ID], peerInfo.Network.Inbound)
		}
		fmt.Fprintf(ctx.App.Writer, "\n")
	}
	return nil
}

func startLogStats(ctx *cli.Context) error {
	if ctx.Args().Len() != 0 {
		return cli.ShowCommandHelp(ctx, ctx.Command.Name)
	}
	csvFile := ctx.String("file")
	f, err := os.OpenFile(csvFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	timer := time.NewTicker(ctx.Duration("interval"))

	f.WriteString("node,timestamp,type,value\n")

loop:
	for {
		select {
		case <-sig:
			return nil
		case <-timer.C:
			stats, err := client.GetAllNodePeerStats()
			if err != nil {
				fmt.Fprintln(ctx.App.Writer, err)
				goto loop
			}
			for nodeID, stats := range stats {
				t := time.Now()
				f.WriteString(fmt.Sprintf("%s,%d,%s,%d\n", nodeID, t.Unix(), "PeerCount", stats.PeerCount))
				f.WriteString(fmt.Sprintf("%s,%d,%s,%d\n", nodeID, t.Unix(), "Tried", stats.Tried))
				f.WriteString(fmt.Sprintf("%s,%d,%s,%d\n", nodeID, t.Unix(), "Failed", stats.Failed))
				f.WriteString(fmt.Sprintf("%s,%d,%s,%d\n", nodeID, t.Unix(), "DifferentNodesDiscovered", stats.DifferentNodesDiscovered))
				f.WriteString(fmt.Sprintf("%s,%d,%s,%d\n", nodeID, t.Unix(), "DHTBuckets", stats.DHTBuckets))
			}
		}
	}
}
