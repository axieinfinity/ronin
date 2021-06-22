#!/bin/sh

# vars from docker env
# - PASSWORD (default to empty)
# - PRIVATE_KEY (default to empty)
# - BOOTNODES (default to empty)
# - VERBOSITY (default to 3)
# - SYNC_MODE (default to 'snap')
# - NETWORK_ID (default to 2021)
# - GASPRICE (default to 0)
# - FORCE_INIT (default to 'false')

# constants
DATA_DIR="/ronin/data"
KEYSTORE_DIR="/ronin/keystore"
PASSWORD_FILE="$KEYSTORE_DIR/password"

# variables
genesisPath=""
params=""

# networkid
if [[ ! -z $NETWORK_ID ]]; then
  case $NETWORK_ID in
    2020 )
      genesisPath="mainnet.json"
      ;;
    2021 )
      genesisPath="testnet.json"
      params="$params --gcmode archive --http.api eth,net,web3,debug,consortium"
      ;;
    2022 )
      genesisPath="devnet.json"
      params="$params --gcmode archive --http.api eth,net,web3,debug,consortium"
      ;;
    * )
      echo "network id not supported"
      ;;
  esac
  params="$params --networkid $NETWORK_ID"
fi

# custom genesis path
if [[ ! -z $GENESIS_PATH ]]; then
  genesisPath="$GENESIS_PATH"
fi

# data dir
if [[ ! -d $DATA_DIR/ronin ]]; then
  echo "No blockchain data, creating genesis block."
  ronin init $genesisPath --datadir $DATA_DIR 2> /dev/null
elif [[ $FORCE_INIT = 'true' ]]; then
  echo "Forcing update chain config."
  ronin init $genesisPath --datadir $DATA_DIR 2> /dev/null
fi

# password file
if [[ ! -f $PASSWORD_FILE ]]; then
  mkdir -p $KEYSTORE_DIR
  if [[ ! -z $PASSWORD ]]; then
    echo "Password env is set. Writing into file."
    echo "$PASSWORD" > $PASSWORD_FILE
  else
    echo "No password set (or empty), generating a new one"
    $(< /dev/urandom tr -dc _A-Z-a-z-0-9 | head -c${1:-32} > $PASSWORD_FILE)
  fi
fi

accountsCount=$(
  ronin account list --datadir $DATA_DIR  --keystore $KEYSTORE_DIR \
  2> /dev/null \
  | wc -l
)

# private key
if [[ $accountsCount -le 0 ]]; then
  echo "No accounts found"
  if [[ ! -z $PRIVATE_KEY ]]; then
    echo "Creating account from private key"
    echo "$PRIVATE_KEY" > ./private_key
    ronin account import ./private_key \
      --datadir $DATA_DIR \
      --keystore $KEYSTORE_DIR \
      --password $PASSWORD_FILE
    rm ./private_key
  else
    echo "Creating new account"
    ronin account new \
      --datadir $DATA_DIR \
      --keystore $KEYSTORE_DIR \
      --password $PASSWORD_FILE
  fi
fi

account=$(
  ronin account list --datadir $DATA_DIR  --keystore $KEYSTORE_DIR \
  2> /dev/null \
  | head -n 1 \
  | cut -d"{" -f 2 | cut -d"}" -f 1
)
echo "Using account $account"
params="$params --unlock $account"

# bootnodes
if [[ ! -z $BOOTNODES ]]; then
  params="$params --bootnodes $BOOTNODES"
fi

# syncmode
if [[ ! -z $SYNC_MODE ]]; then
  params="$params --syncmode ${SYNC_MODE}"
fi

# debug mode - enable rpc and disable local transactions
if [[ ! -z $RPC_NODE ]]; then
  params="$params --gcmode archive --http.api eth,net,web3,debug,consortium --txpool.nolocals"
fi

# ethstats
if [[ ! -z $ETHSTATS_ENDPOINT ]]; then
  params="$params --ethstats $ETHSTATS_ENDPOINT"
fi

# nodekey
if [[ ! -z $NODEKEY ]]; then
  echo $NODEKEY > $PWD/.nodekey
  params="$params --nodekey $PWD/.nodekey"
fi

# gasprice
if [[ ! -z $GASPRICE ]]; then
  params="$params --miner.gasprice $GASPRICE"
fi

# dump
echo "dump: $account $BOOTNODES"

set -x

exec ronin $params \
  --verbosity $VERBOSITY \
  --datadir $DATA_DIR \
  --keystore $KEYSTORE_DIR \
  --password $PASSWORD_FILE \
  --port 30303 \
  --txpool.globalqueue 10000 \
  --txpool.globalslots 10000 \
  --http \
  --http.corsdomain "*" \
  --http.addr 0.0.0.0 \
  --http.port 8545 \
  --http.vhosts "*" \
  --ws \
  --ws.addr 0.0.0.0 \
  --ws.port 8546 \
  --ws.origins "*" \
  --mine \
  --allow-insecure-unlock \
  --miner.gastarget "100000000" \
  "$@"
