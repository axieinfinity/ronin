#!/bin/sh

# vars from docker env
# - PASSWORD (default to empty)
# - PRIVATE_KEY (default to empty)
# - BOOTNODES (default to empty)
# - VERBOSITY (default to 3)
# - SYNC_MODE (default to 'full')
# - NETWORK_ID (default to 2021)

# constants
DATA_DIR="data"
KEYSTORE_DIR="keystore"

# variables
genesisPath=""
params=""
accountsCount=$(
  geth account list --datadir $DATA_DIR  --keystore $KEYSTORE_DIR \
  2> /dev/null \
  | wc -l
)

# networkid
if [[ ! -z $NETWORK_ID ]]; then
  case $NETWORK_ID in
    2020 )
      genesisPath="mainnet.json"
      ;;
    2021 )
      genesisPath="testnet.json"
      params="$params --gcmode archive --rpcapi db,eth,net,web3,debug"
      ;;
    2022 )
      genesisPath="devnet.json"
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
if [[ ! -d $DATA_DIR/geth ]]; then
  echo "No blockchain data, creating genesis block."
  geth init $genesisPath --datadir $DATA_DIR 2> /dev/null
fi

# password file
if [[ ! -f ./password ]]; then
  if [[ ! -z $PASSWORD ]]; then
    echo "Password env is set. Writing into file."
    echo "$PASSWORD" > ./password
  else
    echo "No password set (or empty), generating a new one"
    $(< /dev/urandom tr -dc _A-Z-a-z-0-9 | head -c${1:-32} > password)
  fi
fi

# private key
if [[ $accountsCount -le 0 ]]; then
  echo "No accounts found"
  if [[ ! -z $PRIVATE_KEY ]]; then
    echo "Creating account from private key"
    echo "$PRIVATE_KEY" > ./private_key
    geth account import ./private_key \
      --datadir $DATA_DIR \
      --keystore $KEYSTORE_DIR \
      --password ./password
    rm ./private_key
  else
    echo "Creating new account"
    geth account new \
      --datadir $DATA_DIR \
      --keystore $KEYSTORE_DIR \
      --password ./password
  fi
fi

account=$(
  geth account list --datadir $DATA_DIR  --keystore $KEYSTORE_DIR \
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

# debug mode
if [[ ! -z $DEBUG_MODE ]]; then
  params="$params --gcmode archive --rpcapi db,eth,net,web3,debug,posv"
fi

# dump
echo "dump: $account $BOOTNODES"

set -x

exec geth $params \
  --verbosity $VERBOSITY \
  --datadir $DATA_DIR \
  --keystore $KEYSTORE_DIR \
  --password ./password \
  --port 30303 \
  --rpc \
  --rpccorsdomain "*" \
  --rpcaddr 0.0.0.0 \
  --rpcport 8545 \
  --ws \
  --wsaddr 0.0.0.0 \
  --wsport 8546 \
  --wsorigins "*" \
  --mine \
  --allow-insecure-unlock \
  --gasprice "0" \
  --targetgaslimit "20000000" \
  "$@"
