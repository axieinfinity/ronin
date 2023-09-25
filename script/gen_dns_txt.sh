#!/bin/bash

# Create the DNS TXT record from the list of nodes

unset BOOTNODES
unset NETWORK
unset DOMAIN
unset IP_LIST
PROGNAME=$(basename "$0")

USAGE="$PROGNAME -b bootnodes -n network -d domain [-i ip_list]

where:
    bootnodes:  list of bootnodes to crawl from (comma separated)
    network:    ronin-mainnet | ronin-testnet to choose the network
    domain:     the DNS domain name
    ip_list:    the list of IPs to include in ENR tree (comma separated)
"

set -e

while getopts :hb:n:d:i: option; do
  case $option in
    h)
      echo "$USAGE"
      exit 0
      ;;
    b) BOOTNODES=$OPTARG ;;
    n) NETWORK=$OPTARG ;;
    d) DOMAIN=$OPTARG ;;
    i) IP_LIST=$OPTARG ;;
    :)
      echo "$PROGNAME: option requires an argument -- '$OPTARG'" >&2
      exit 1
      ;;
    *)
      echo "$PROGNAME: bad option -$OPTARG" >&2
      echo "Try '$PROGNAME -h' for more information" >&2
      exit 1
      ;;
  esac
done
shift $((OPTIND - 1))

if [[ -z $BOOTNODES || -z $NETWORK || -z $DOMAIN ]]; then
  echo "$PROGNAME: missing mandatory option" >&2
  echo "Try '$(basename "$0") -h' for more information" >&2
  exit 1
fi

DNS_DIR=dns_record

if [[ -z $PRIVATE_KEY ]]; then
  echo "The PRIVATE_KEY environment for signing DNS record must be provided"
  exit 1
fi

if [[ -z $PASSWORD ]]; then
  echo "The PASSWORD environment for decrypting the private key must be provided"
  exit 1
fi

echo "$PASSWORD" > password
echo "$PRIVATE_KEY" > private_key
unset PASSWORD
unset PRIVATE_KEY

set -x
# Create an encrypted keyfile.json
ethkey generate --passwordfile password --privatekey private_key

devp2p discv4 crawl --timeout 1m --bootnodes $BOOTNODES all_nodes.json

mkdir -p $DNS_DIR

set +x
IP_LIST_PARAMS=""
if [[ ! -z $IP_LIST ]]; then
  IP_LIST_PARAMS="-ip-list $IP_LIST"
fi
set -x

devp2p nodeset filter all_nodes.json -eth-network $NETWORK $IP_LIST_PARAMS > $DNS_DIR/nodes.json
devp2p dns sign $DNS_DIR keyfile.json password --domain $DOMAIN
devp2p dns to-txt $DNS_DIR $DNS_DIR/txt_record.json

echo "Cleanup files"
rm private_key
rm password
rm keyfile.json
