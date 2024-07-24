#!/bin/bash -e

CURRENT_DIR=`pwd`
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )


cd $SCRIPT_DIR/..
echo "Changed working directory into: `pwd`"
trap "echo \"Restoring working directory to $CURRENT_DIR\"; cd $CURRENT_DIR" EXIT


export RONIN_NODE_PATH=${RONIN_NODE_PATH:-./script/run/ronin}
export GENESIS_FILE=${GENESIS_FILE:=./genesis/devnet.json}

# Ensure to mkdir node/keystore, node/ronin
rm -rf $RONIN_NODE_PATH/node1/keystore/UTC*
rm -rf $RONIN_NODE_PATH/node2/keystore/UTC*
rm -rf $RONIN_NODE_PATH/node3/keystore/UTC*
rm -rf $RONIN_NODE_PATH/node1/keystore/all*
rm -rf $RONIN_NODE_PATH/node2/keystore/all*
rm -rf $RONIN_NODE_PATH/node3/keystore/all*

mkdir -p $RONIN_NODE_PATH/node1/keystore/
mkdir -p $RONIN_NODE_PATH/node2/keystore/
mkdir -p $RONIN_NODE_PATH/node3/keystore/
if [[ ! -f "$RONIN_NODE_PATH/node1/keystore/" ]]; then
  openssl rand -base64 20 >$RONIN_NODE_PATH/node1/keystore/password
fi
if [[ ! -f "$RONIN_NODE_PATH/node2/keystore/" ]]; then
  openssl rand -base64 20 >$RONIN_NODE_PATH/node2/keystore/password
fi
if [[ ! -f "$RONIN_NODE_PATH/node3/keystore/" ]]; then
  openssl rand -base64 20 >$RONIN_NODE_PATH/node3/keystore/password
fi

make ronin
make bootnode
RONIN_CMD=./build/bin/ronin
BOOTNODE_CMD=./build/bin/bootnode

addr1=$($RONIN_CMD account new --datadir $RONIN_NODE_PATH/node1 --password $RONIN_NODE_PATH/node1/keystore/password | grep "0x" | cut -b 30-71)
addr2=$($RONIN_CMD account new --datadir $RONIN_NODE_PATH/node2 --password $RONIN_NODE_PATH/node2/keystore/password | grep "0x" | cut -b 30-71)
addr3=$($RONIN_CMD account new --datadir $RONIN_NODE_PATH/node3 --password $RONIN_NODE_PATH/node3/keystore/password | grep "0x" | cut -b 30-71)

echo "$addr1"
echo "$addr2"
echo "$addr3"

bls1=$($RONIN_CMD account generatebls --finality.blswalletpath $RONIN_NODE_PATH/node1/keystore --finality.blspasswordpath $RONIN_NODE_PATH/node1/keystore/password | grep "{" | cut -b 14-109)
bls2=$($RONIN_CMD account generatebls --finality.blswalletpath $RONIN_NODE_PATH/node2/keystore --finality.blspasswordpath $RONIN_NODE_PATH/node2/keystore/password | grep "{" | cut -b 14-109)
bls3=$($RONIN_CMD account generatebls --finality.blswalletpath $RONIN_NODE_PATH/node3/keystore --finality.blspasswordpath $RONIN_NODE_PATH/node3/keystore/password | grep "{" | cut -b 14-109)


rm -rf $RONIN_NODE_PATH/node3/ronin
rm -rf $RONIN_NODE_PATH/node2/ronin
rm -rf $RONIN_NODE_PATH/node1/ronin

$RONIN_CMD init --datadir $RONIN_NODE_PATH/node1 $GENESIS_FILE
$RONIN_CMD init --datadir $RONIN_NODE_PATH/node2 $GENESIS_FILE
$RONIN_CMD init --datadir $RONIN_NODE_PATH/node3 $GENESIS_FILE


NODE1_NODEKEY=$(cat $RONIN_NODE_PATH/node1/ronin/nodekey)
BOOTNODE_ADDR=$($BOOTNODE_CMD -nodekeyhex ${NODE1_NODEKEY} -writeaddress)

stake1="1111"
stake2="1111"
stake3="1111"

cat <<EOF > ${SCRIPT_DIR}/run_node1.sh
#!/bin/bash

SCRIPT_DIR=\$( cd -- "\$( dirname -- "\${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd \$SCRIPT_DIR/..
echo "Changed working directory into: `pwd`"
trap "echo \"Restoring working directory to \$CURRENT_DIR\"; cd \$CURRENT_DIR" EXIT

$RONIN_CMD --http.api eth,net,web3,debug --networkid 2022 --verbosity 3 \\
  --rpc.allow-unprotected-txs \\
  --datadir $RONIN_NODE_PATH/node1 --port 30303 \\
  --http --http.corsdomain '*' --http.addr 0.0.0.0 --http.port 8545 \\
  --ws --ws.origins "*" --ws.addr 0.0.0.0 --ws.port 8645  \\
  -allow-insecure-unlock --mine \\
  --keystore $RONIN_NODE_PATH/node1/keystore --password $RONIN_NODE_PATH/node1/keystore/password \\
  --unlock $addr1 --miner.gaslimit 100000000 \\
  --mock.validators $addr1,$addr2,$addr3 --mock.stakeamounts $stake1,$stake2,$stake3 \\
  --bootnodes enode://$BOOTNODE_ADDR@127.0.0.1:30303 \\
  --mock.blspublickeys "$bls1,$bls2,$bls3" \\
  --finality.blswalletpath $RONIN_NODE_PATH/node1/keystore \\
  --finality.blspasswordpath $RONIN_NODE_PATH/node1/keystore/password --finality.enable --finality.enablesign
EOF

cat <<EOF > ${SCRIPT_DIR}/run_node2.sh
#!/bin/bash

SCRIPT_DIR=\$( cd -- "\$( dirname -- "\${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd \$SCRIPT_DIR/..
echo "Changed working directory into: `pwd`"
trap "echo \"Restoring working directory to \$CURRENT_DIR\"; cd \$CURRENT_DIR" EXIT

$RONIN_CMD --http.api eth,net,web3,debug --networkid 2022 --verbosity 3 \\
  --rpc.allow-unprotected-txs \\
  --datadir $RONIN_NODE_PATH/node2 --port 30304 \\
  --http --http.corsdomain '*' --http.addr 0.0.0.0 --http.port 8546 \\
  --ws --ws.origins "*" --ws.addr 0.0.0.0 --ws.port 8646  \\
  -allow-insecure-unlock --mine \\
  --keystore $RONIN_NODE_PATH/node2/keystore --password $RONIN_NODE_PATH/node2/keystore/password \\
  --unlock $addr2 --miner.gaslimit 100000000 \\
  --mock.validators $addr1,$addr2,$addr3 --mock.stakeamounts $stake1,$stake2,$stake3 \\
  --bootnodes enode://$BOOTNODE_ADDR@127.0.0.1:30303 \\
  --mock.blspublickeys "$bls1,$bls2,$bls3" \\
  --finality.blswalletpath $RONIN_NODE_PATH/node2/keystore \\
  --finality.blspasswordpath $RONIN_NODE_PATH/node2/keystore/password --finality.enable --finality.enablesign
EOF

cat <<EOF > ${SCRIPT_DIR}/run_node3.sh
#!/bin/bash

SCRIPT_DIR=\$( cd -- "\$( dirname -- "\${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd \$SCRIPT_DIR/..
echo "Changed working directory into: `pwd`"
trap "echo \"Restoring working directory to \$CURRENT_DIR\"; cd \$CURRENT_DIR" EXIT

$RONIN_CMD --http.api eth,net,web3,debug --networkid 2022 --verbosity 3 \\
  --rpc.allow-unprotected-txs \\
  --datadir $RONIN_NODE_PATH/node3 --port 30306 \\
  --http --http.corsdomain '*' --http.addr 0.0.0.0 --http.port 8547 \\
  --ws --ws.origins "*" --ws.addr 0.0.0.0 --ws.port 8647  \\
  -allow-insecure-unlock --mine \\
  --keystore $RONIN_NODE_PATH/node3/keystore --password $RONIN_NODE_PATH/node3/keystore/password \\
  --unlock $addr3 --miner.gaslimit 100000000 \\
  --mock.validators $addr1,$addr2,$addr3 --mock.stakeamounts $stake1,$stake2,$stake3 \\
  --bootnodes enode://$BOOTNODE_ADDR@127.0.0.1:30303 \\
  --mock.blspublickeys "$bls1,$bls2,$bls3" \\
  --finality.blswalletpath $RONIN_NODE_PATH/node3/keystore \\
  --finality.blspasswordpath $RONIN_NODE_PATH/node3/keystore/password --finality.enable --finality.enablesign
EOF

chmod +x run_node1.sh run_node2.sh run_node3.sh