# devp2p simulation for discovery benchmark

## Overview

In this simulation, we will focus on benchmarking the discovery process by simulating the network with a number of nodes and bootnodes. We aim to measure peer quality when bypassing and not bypassing the ENR filter when an ENR request fails, as well as adjusting the DHT bucket size from 16 to 256.

## Manual run

Run the p2psim server by `go run discovery.go`, and in another terminal, we can use `p2psim` cli to start, manage new nodes in the simulation network. Example:

``` bash
$ go run discovery.go
INFO [12-24|14:46:39.132] starting simulation server               port=8888
```

``` bash
$ p2psim node create-multi --count 2 --fake.iplistener --start -node.type bootnode --enable.enrfilter
Created bootnode-1735026417-0
Started bootnode-1735026417-0
Created bootnode-1735026417-1
Started bootnode-1735026417-1
```

``` bash
$ p2psim node create-multi --count 16 --fake.iplistener --start --autofill.bootnodes --dirty.rate 50 --enable.enrfilter
Created node-1735026508-0
Started node-1735026508-0
Created node-1735026508-1
Started node-1735026508-1
Created dirty-1735026508-2
Started dirty-1735026508-2
Created dirty-1735026508-3
Started dirty-1735026508-3
Created dirty-1735026508-4
Started dirty-1735026508-4
Created node-1735026508-5
Started node-1735026508-5
Created node-1735026508-6
Started node-1735026508-6
Created node-1735026508-7
Started node-1735026508-7
Created node-1735026508-8
Started node-1735026508-8
Created dirty-1735026508-9
Started dirty-1735026508-9
Created node-1735026508-10
Started node-1735026508-10
Created dirty-1735026508-11
Started dirty-1735026508-11
Created dirty-1735026508-12
Started dirty-1735026508-12
Created node-1735026508-13
Started node-1735026508-13
Created dirty-1735026508-14
Started dirty-1735026508-14
Created dirty-1735026508-15
Started dirty-1735026508-15
```

## Strategy

We have some types of nodes:
1. Dirty nodes: Nodes that are not compatible with the valid nodes
2. Valid nodes
3. Valid nodes that only accept outbound connections

The benchmark default will run with 350 nodes and 2 bootnodes (can be adjusted in the configuration), and will be rolled out in 3 batches following below steps:
1. Start the simulation server, 2 bootnodes and rolling out nodes in batch 1 and sleep for a while
2. Rolling out nodes in batch 2 and sleep for a while
3. Rolling out nodes in batch 3 and sleep for a while
4. Export the DHT and peers info

## Run benchmark

To run the simulation, run `./discovery.sh` to start both p2psim server and start the benchmark with default parameters.

### Configuration

To show the help message, run `./discovery.sh --help`. Besides the configurable parameters, we can modify the source code to change the behavior of the simulation:
- If we want to change the DHT bucket size, we can modify the const `bucketSize` in `p2p/discover/table.go`
- Or filter node if request ENR fails, we can modify the function `Table::filterNode` in `p2p/discover/table.go` to:

```go
func (tab *Table) filterNode(n *node) bool {
    ...
	if node, err := tab.net.RequestENR(unwrapNode(n)); err != nil {
		return true // modify here
	} else if !tab.enrFilter(node.Record()) {
        ...
	}
    ...
}
```

### Export data

After running the simulation, some files will be generated in the `results_dir` folder:
- `$test_name.log`: Log of all nodes in the simulation network
- `stats_$test_name.csv`: Statistics of the simulation, including the number of peers, distribution of nodes in the DHT, ...
- `peers_$test_name.log`: List peers of each node in the network
- `dht_$test_name.log`: List nodes in the DHT of each node

### Visualization

To visualize the data, we can use the `discovery.py` script to plot the data.
Supported types:
- `dht_peer`: Ratio between the number of peers (outbound) and the number of nodes in the DHT
- `PeerCount`: Number of peers of each node
- `DHTBuckets`: Size of DHT
- And more type can see in the stats file
