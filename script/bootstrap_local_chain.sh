#!/bin/bash -e

CURRENT_DIR=`pwd`
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )


cd $SCRIPT_DIR/..
echo "Changed working directory into: `pwd`"
trap "echo \"Restoring working directory to $CURRENT_DIR\"; cd $CURRENT_DIR" EXIT


export RONIN_NODE_PATH=${RONIN_NODE_PATH:-./script/run/ronin}
export GENESIS_FILE=${GENESIS_FILE:=./genesis/devnet.json}

NUM_NODES=3
for i in $(seq 1 $NUM_NODES); do
  # Ensure to mkdir node/keystore, node/ronin
  rm -rf $RONIN_NODE_PATH/node$i/keystore/UTC*
  rm -rf $RONIN_NODE_PATH/node$i/keystore/all*
  mkdir -p $RONIN_NODE_PATH/node$i/keystore/
  if [[ ! -f "$RONIN_NODE_PATH/node$i/keystore/password" ]]; then
    openssl rand -base64 20 > $RONIN_NODE_PATH/node$i/keystore/password
  fi
done


make ronin
make bootnode
RONIN_CMD=./build/bin/ronin
BOOTNODE_CMD=./build/bin/bootnode

declare -a addrs
declare -a bls


for i in $(seq 1 $NUM_NODES); do
  addr=$($RONIN_CMD account new --datadir $RONIN_NODE_PATH/node$i --password $RONIN_NODE_PATH/node$i/keystore/password | grep "0x" | cut -b 30-71)
  addrs+=($addr)
  echo "$addr"
  bls_key=$($RONIN_CMD account generatebls --finality.blswalletpath $RONIN_NODE_PATH/node$i/keystore --finality.blspasswordpath $RONIN_NODE_PATH/node$i/keystore/password | grep "{" | cut -b 14-109)
  bls+=($bls_key)
done


for i in $(seq 1 $NUM_NODES); do
  rm -rf $RONIN_NODE_PATH/node$i/ronin
  $RONIN_CMD init --datadir $RONIN_NODE_PATH/node$i $GENESIS_FILE
done

NODE1_NODEKEY=$(cat $RONIN_NODE_PATH/node1/ronin/nodekey)
BOOTNODE_ADDR=$($BOOTNODE_CMD -nodekeyhex ${NODE1_NODEKEY} -writeaddress)

stake="1111"

# Construct the mock validators, stake amounts, and BLS public keys strings
validators=$(IFS=,; echo "${addrs[*]}")
stake_amounts=$(IFS=,; echo "${stake},${stake},${stake}")
bls_public_keys=$(IFS=,; echo "${bls[*]}")

for i in $(seq 1 $NUM_NODES); do
  cat <<EOF > ${SCRIPT_DIR}/run_node$i.sh
#!/bin/bash

SCRIPT_DIR=\$( cd -- "\$( dirname -- "\${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd \$SCRIPT_DIR/..
echo "Changed working directory into: `pwd`"
trap "echo \"Restoring working directory to \$CURRENT_DIR\"; cd \$CURRENT_DIR" EXIT

$RONIN_CMD --http.api eth,net,web3,debug --networkid 2022 --verbosity 3 \\
  --rpc.allow-unprotected-txs \\
  --datadir $RONIN_NODE_PATH/node$i --port $((30303 + i)) \\
  --http --http.corsdomain '*' --http.addr 0.0.0.0 --http.port $((8545 + i)) \\
  --ws --ws.origins "*" --ws.addr 0.0.0.0 --ws.port $((8645 + i))  \\
  -allow-insecure-unlock --mine \\
  --keystore $RONIN_NODE_PATH/node$i/keystore --password $RONIN_NODE_PATH/node$i/keystore/password \\
  --unlock ${addrs[$((i-1))]} --miner.gaslimit 100000000 \\
  --mock.validators $validators --mock.stakeamounts $stake_amounts \\
  --bootnodes enode://$BOOTNODE_ADDR@127.0.0.1:30304 \\
  --mock.blspublickeys "$bls_public_keys" \\
  --finality.blswalletpath $RONIN_NODE_PATH/node$i/keystore \\
  --finality.blspasswordpath $RONIN_NODE_PATH/node$i/keystore/password --finality.enable --finality.enablesign
EOF
done


# Make the run scripts executable
for i in $(seq 1 $NUM_NODES); do
  chmod +x ${SCRIPT_DIR}/run_node$i.sh
done
