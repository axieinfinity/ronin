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
datadir="/ronin/data"
KEYSTORE_DIR="/ronin/keystore"
PASSWORD_FILE="/ronin/password"

# variables
genesisPath=""
params=""
syncmode="snap"
mine="true"

set -e

if [[ ! -z $DATA_DIR ]]; then
  datadir="$DATA_DIR"
fi

# networkid
if [[ ! -z $NETWORK_ID ]]; then
  case $NETWORK_ID in
    2020 )
      genesisPath="mainnet.json"
      ;;
    2021 )
      genesisPath="testnet.json"
      ;;
    * )
      # All other networkids use the devnet.json by default
      genesisPath="devnet.json"
      ;;
  esac
  params="$params $RONIN_PARAMS"
  params="$params --networkid $NETWORK_ID"
fi

# custom genesis path
if [[ ! -z $GENESIS_PATH ]]; then
  genesisPath="$GENESIS_PATH"
fi

# data dir
if [[ ! -d $datadir/ronin ]]; then
  echo "No blockchain data, creating genesis block."
  ronin init $genesisPath --datadir $datadir 2> /dev/null
elif [[ "$FORCE_INIT" = "true" && "$INIT_FORCE_OVERRIDE_CHAIN_CONFIG" = "true" ]]; then
  echo "Forcing update chain config with force overriding chain config."
  ronin init $genesisPath --overrideChainConfig --datadir $datadir 2> /dev/null
elif [ "$FORCE_INIT" = "true" ]; then
  echo "Forcing update chain config."
  ronin init $genesisPath --datadir $datadir 2> /dev/null
fi

# password file
if [[ ! -f $PASSWORD_FILE ]]; then
  mkdir -p $KEYSTORE_DIR
  if [[ ! -z $PASSWORD ]]; then
    echo "Password env is set. Writing into file."
    echo "$PASSWORD" > $PASSWORD_FILE
  else
    echo "No password set (or empty), generating a new one"
    $(< /dev/urandom tr -dc _A-Z-a-z-0-9 | head -c 32 > $PASSWORD_FILE)
  fi
fi

accountsCount=$(
  ronin account list --datadir $datadir  --keystore $KEYSTORE_DIR \
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
      --datadir $datadir \
      --keystore $KEYSTORE_DIR \
      --password $PASSWORD_FILE
    rm ./private_key
  else
    echo "Creating new account"
    ronin account new \
      --datadir $datadir \
      --keystore $KEYSTORE_DIR \
      --password $PASSWORD_FILE
  fi
fi

if [[ ! -z $KEYSTORE_DIR ]]; then
  account=$(
    ronin account list --datadir $datadir  --keystore $KEYSTORE_DIR \
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
  syncmode="$SYNC_MODE"
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
if [ "$SUBSCRIBER" = "true" ]; then
  params="$params --subscriber --subscriber.blockEventTopic subscriber.block"
  params="$params --subscriber.txEventTopic subscriber.txs"
  params="$params --subscriber.logsEventTopic subscriber.logs"
  params="$params --subscriber.reOrgBlockEventTopic subscriber.block.reorg"
  params="$params --subscriber.reorgTxEventTopic subscriber.txs.reorg"
  params="$params --subscriber.blockConfirmedEventTopic subscriber.block.confirmed"
  params="$params --subscriber.transactionConfirmedEventTopic subscriber.txs.confirmed"
  params="$params --subscriber.logsConfirmedEventTopic subscriber.logs.confirmed"
  params="$params --subscriber.internalTransactionEventTopic subscriber.txs.internal"

  if [[ ! -z $KAFKA_URL ]]; then
    params="$params --subscriber.kafka.url $KAFKA_URL"
  fi

  if [ ! -z $KAFKA_USERNAME ] && [ ! -z KAFKA_PASSWORD]; then
    params="$params --subscriber.kafka.username $KAFKA_USERNAME --subscriber.kafka.password $KAFKA_PASSWORD"
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

  if [[ ! -z $CONFIRM_BLOCK_AT ]]; then
    params="$params --subscriber.confirmBlockAt $CONFIRM_BLOCK_AT"
  fi
fi

if [[ ! -z $KEYSTORE_DIR ]]; then
  params="$params --keystore $KEYSTORE_DIR"
fi

if [[ ! -z $PASSWORD_FILE ]]; then
  params="$params --password $PASSWORD_FILE"
fi

if [[ ! -z $MINE ]]; then
  mine="$MINE"
fi

if [[ "$mine" = "true" ]]; then
  params="$params --mine"
fi

# dump
echo "dump: $account $BOOTNODES"

set -x

exec ronin $params \
  --syncmode $syncmode \
  --verbosity $VERBOSITY \
  --datadir $datadir \
  --port 30303 \
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
