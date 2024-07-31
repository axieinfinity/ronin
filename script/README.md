# Scripts

## bootstrap_local_chain.sh

This script helps prepare environment and start scripts to bootstrap a 3 nodes Ronin devnet.

### Bootstrap the config
```bash
# run with default settings 
./script/bootstrap_local_chain.sh

# run with custom data directory and genesis file 
RONIN_NODE_PATH=./script/run/ronin GENESIS_FILE=./genesis/devnet.json ./script/bootstrap_local_chain.sh
```
### Run the nodes
```bash
./script/run_node1.sh
./script/run_node2.sh
./script/run_node3.sh
```
