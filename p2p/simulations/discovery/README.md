# devp2p simulation for discovery benchmark

## Overview

In this simulation, we will focus on benchmarking the discovery process by simulating the network with a number of nodes and bootnodes. We aim to measure peer quality when bypassing and not bypassing the ENR filter when an ENR request fails, as well as adjusting the DHT bucket size from 16 to 256.

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

## Run simulation

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
