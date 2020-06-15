#!/bin/sh

# bootnode key
if [[ ! -z $BOOTNODE_KEY ]]; then
  echo $BOOTNODE_KEY > bootnode.key
elif [[ ! -f ./bootnode.key ]]; then
  bootnode -genkey bootnode.key
fi

# dump address
address="enode://$(bootnode -nodekey bootnode.key -writeaddress)@[$(hostname -i)]:30301"

set -x
echo "$address" > ./address

exec bootnode "$@"
