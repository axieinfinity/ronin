#!/bin/bash

# A tool to create/clean up the overlayfs over the chaindata directory.
# overlayfs can be used to quickly rollback the chaindata directory
# after testing.

unset DATADIR
unset CLEANUP

PROGNAME=$(basename "$0")
USAGE="$PROGNAME -d datadir [-c]

where:
    -d datadir: the chaindata's parent directory
    -c:         clean up mode
"

while getopts :hd:c option; do
  case $option in
    h)
      echo "$USAGE"
      exit 0
      ;;
    d) DATADIR=$OPTARG ;;
    c) CLEANUP=true ;;
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

if [[ -z $DATADIR ]]; then
  echo "$PROGNAME: missing mandatory option" >&2
  echo "Try '$(basename "$0") -h' for more information" >&2
  exit 1
fi

cleanup_overlay () {
  set -ex

  cd $DATADIR
  umount ./chaindata
  rm -rf chaindata
  rm -rf upper
  rm -rf workdir
  mv orig_chaindata chaindata
}

setup_overlay () {
  set -ex

  cd $DATADIR
  mkdir -p upper
  mkdir -p workdir
  mv chaindata orig_chaindata
  mkdir -p chaindata

  mount -t overlay overlay \
    -olowerdir=./orig_chaindata,upperdir=./upper,workdir=./workdir chaindata
}

if [[ "$CLEANUP" = "true" ]]; then
  cleanup_overlay
else
  setup_overlay
fi
