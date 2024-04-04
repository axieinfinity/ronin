#!/bin/sh
# eg, RONIN = ronin/cmd/ronin
# eg, RONIN_NODE_PATH = ronin_nodes // where the nodes are stored
# eg, SANDBOX = ronin/genesis/sandbox.json // where the genesis file is stored
export RONIN=
export RONIN_NODE_PATH=
export SANDBOX=

# Ensure to mkdir node/keystore, node/ronin
rm -rf $RONIN_NODE_PATH/node1/keystore/UTC*
rm -rf $RONIN_NODE_PATH/node2/keystore/UTC*
rm -rf $RONIN_NODE_PATH/node3/keystore/UTC*
rm -rf $RONIN_NODE_PATH/node1/keystore/all*
rm -rf $RONIN_NODE_PATH/node2/keystore/all*
rm -rf $RONIN_NODE_PATH/node3/keystore/all*


go build $RONIN

chmod +x $RONIN/ronin

addr1=$(./ronin account new --datadir $RONIN_NODE_PATH/node1 --password $RONIN_NODE_PATH/node1/keystore/password | grep "0x" | cut -b 30-71)
addr2=$(./ronin account new --datadir $RONIN_NODE_PATH/node2 --password $RONIN_NODE_PATH/node2/keystore/password | grep "0x" | cut -b 30-71)
addr3=$(./ronin account new --datadir $RONIN_NODE_PATH/node3 --password $RONIN_NODE_PATH/node3/keystore/password | grep "0x" | cut -b 30-71)

echo "$addr1"
echo "$addr2"
echo "$addr3"

bls1=$(./ronin account generatebls --finality.blswalletpath $RONIN_NODE_PATH/node1/keystore --finality.blspasswordpath $RONIN_NODE_PATH/node1/keystore/password | grep "{" | cut -b 14-109)
bls2=$(./ronin account generatebls --finality.blswalletpath $RONIN_NODE_PATH/node2/keystore --finality.blspasswordpath $RONIN_NODE_PATH/node2/keystore/password | grep "{" | cut -b 14-109)
bls3=$(./ronin account generatebls --finality.blswalletpath $RONIN_NODE_PATH/node3/keystore --finality.blspasswordpath $RONIN_NODE_PATH/node3/keystore/password | grep "{" | cut -b 14-109)

stake1="1111"
stake2="1111"
stake3="1111"

echo "./ronin --http.api eth,net,web3,debug --networkid 2022 --verbosity 3 --datadir $RONIN_NODE_PATH/node1 --port 30303 --http --http.corsdomain '*' --http.addr 0.0.0.0 --http.port 8545 -allow-insecure-unlock --mine --keystore $RONIN_NODE_PATH/node1/keystore --password $RONIN_NODE_PATH/node1/keystore/password --unlock $addr1 --miner.gaslimit 100000000 --mock.validators $addr1,$addr2,$addr3 --mock.stakeamounts $stake1,$stake2,$stake3 --bootnodes enode://cb879c77c82dedb32072e48457f184dfe6d3e21b7f44130d48bd16b9a17d97062e8d48661d5af773ddfddaa89a3b9cd8849ee839c0b352a9283113baa759a2a8@127.0.0.1:30304 --mock.blspublickeys "$bls1,$bls2,$bls3" --finality.blswalletpath $RONIN_NODE_PATH/node1/keystore --finality.blspasswordpath $RONIN_NODE_PATH/node1/keystore/password --finality.enable --finality.enablesign"
echo "./ronin --http.api eth,net,web3,debug --networkid 2022 --verbosity 3 --datadir $RONIN_NODE_PATH/node2 --port 30304 --http --http.corsdomain '*' --http.addr 0.0.0.0 --http.port 8546 -allow-insecure-unlock --mine --keystore $RONIN_NODE_PATH/node2/keystore --password $RONIN_NODE_PATH/node2/keystore/password --unlock $addr2 --miner.gaslimit 100000000 --mock.validators $addr1,$addr2,$addr3 --mock.stakeamounts $stake1,$stake2,$stake3 --bootnodes enode://8c420102e7f5b0dfa6e78a2b8bb4f6e10f989e2181cbee9d6bd6a3a32cd2fb463747de49210f09f446489066fa1e0156e71e0283bc2b02473fc03debd2fbd2d7@127.0.0.1:30303 --mock.blspublickeys "$bls1,$bls2,$bls3" --finality.blswalletpath $RONIN_NODE_PATH/node2/keystore --finality.blspasswordpath $RONIN_NODE_PATH/node2/keystore/password --finality.enable --finality.enablesign"
echo "./ronin --http.api eth,net,web3,debug --networkid 2022 --verbosity 3 --datadir $RONIN_NODE_PATH/node3 --port 30306 --http --http.corsdomain '*' --http.addr 0.0.0.0 --http.port 8547 -allow-insecure-unlock --mine --keystore $RONIN_NODE_PATH/node3/keystore --password $RONIN_NODE_PATH/node3/keystore/password --unlock $addr3 --miner.gaslimit 100000000 --mock.validators $addr1,$addr2,$addr3 --mock.stakeamounts $stake1,$stake2,$stake3 --bootnodes enode://8c420102e7f5b0dfa6e78a2b8bb4f6e10f989e2181cbee9d6bd6a3a32cd2fb463747de49210f09f446489066fa1e0156e71e0283bc2b02473fc03debd2fbd2d7@127.0.0.1:30303 --mock.blspublickeys "$bls1,$bls2,$bls3" --finality.blswalletpath $RONIN_NODE_PATH/node3/keystore --finality.blspasswordpath $RONIN_NODE_PATH/node3/keystore/password --finality.enable --finality.enablesign"

rm -rf $RONIN_NODE_PATH/node3/ronin && \
rm -rf $RONIN_NODE_PATH/node2/ronin && \
rm -rf $RONIN_NODE_PATH/node1/ronin && \
./ronin init $SANDBOX --datadir $RONIN_NODE_PATH/node1 && \
./ronin init $SANDBOX --datadir $RONIN_NODE_PATH/node2 && \
./ronin init $SANDBOX --datadir $RONIN_NODE_PATH/node3 
