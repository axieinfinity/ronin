#!/bin/sh

# vars from docker env
# - PASSWORD (default to empty)
# - PRIVATE_KEY (default to empty)
# - BOOTNODES (default to empty)
# - VERBOSITY (default to 3)
# - SYNC_MODE (default to 'snap')
# - NETWORK_ID (default to 2021)
# - GASPRICE (default to 0)
# - FORCE_INIT (default to 'true')

# constants
# DATA_DIR="/ronin/data"
# KEYSTORE_DIR="/ronin/keystore"
# PASSWORD_FILE="$KEYSTORE_DIR/password"

# variables
genesisPath=""
params=""

# networkid
if [[ ! -z $NETWORK_ID ]]; then
  case $NETWORK_ID in
    2020 )
      genesisPath="mainnet.json"
      params="$params $RONIN_PARAMS"
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

if [[ ! -z $KEYSTORE_DIR ]]; then
  account=$(
    ronin account list --datadir $DATA_DIR  --keystore $KEYSTORE_DIR \
    2> /dev/null \
    | head -n 1 \
    | cut -d"{" -f 2 | cut -d"}" -f 1
  )
  echo "Using account $account"
  params="$params --unlock $account"
fi

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

# subscriber
if [[ ! -z $SUBSCRIBER ]]; then
  params="$params --subscriber --subscriber.blockEventTopic block_event"
  params="$params --subscriber.txEventTopic txs_event"
  params="$params --subscriber.logsEventTopic logs_event"
  params="$params --subscriber.reOrgBlockEventTopic reorg_event"
  params="$params --subscriber.reorgTxEventTopic reorg_tx_event"

  if [[ ! -z $KAFKA_URL ]]; then
    params="$params --subscriber.kafka.url $KAFKA_URL"
  fi

  if [ ! -z $KAFKA_USERNAME ] && [ ! -z KAFKA_PASSWORD]; then
    params="$params --subscriber.kafka.username $KAFKA_USERNAME --subscriber.kafka.password $KAFKA_PASSWORD"
  fi

  if [[ ! -z $SUBSCRIBER_WORKERS ]]; then
    params="$params --subscriber.workers $SUBSCRIBER_WORKERS"
  fi

  if [[ ! -z $SUBSCRIBER_MAX_RETRY ]]; then
    params="$params --subscriber.maxRetry $SUBSCRIBER_MAX_RETRY"
  fi

  if [[ ! -z $SUBSCRIBER_BACK_OFF ]]; then
    params="$params --subscriber.backoff $SUBSCRIBER_BACK_OFF"
  fi

  if [[ ! -z $KAFKA_AUTHENTICATION_TYPE ]]; then
    case $KAFKA_AUTHENTICATION_TYPE in
      PLAIN|SCRAM-SHA-256|SCRAM-SHA-512 )
        params="$params --subscriber.kafka.authentication $KAFKA_AUTHENTICATION_TYPE"
        ;;
      * )
        params="$params --subscriber.kafka.authentication PLAIN"
        ;;
    esac
  fi

  if [[ ! -z $SUBSCRIBER_FROM_HEIGHT ]]; then
    params="$params --subscriber.fromHeight $SUBSCRIBER_FROM_HEIGHT"
  fi
fi

if [[ ! -z $KEYSTORE_DIR ]]; then
  params="$params --keystore $KEYSTORE_DIR"
fi

if [[ ! -z $PASSWORD_FILE ]]; then
  params="$params --password $PASSWORD_FILE"
fi

if [[ ! -z $MINE ]]; then
  params="$param --mine"
fi

# dump
echo "dump: $account $BOOTNODES"

set -x

echo "params: $params"

exec ronin $params \
  --verbosity $VERBOSITY \
  --datadir $DATA_DIR \
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
  --allow-insecure-unlock \
  --miner.gastarget "100000000" \
  "$@"
