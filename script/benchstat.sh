#!/bin/bash

unset COMMAND
unset OLD_COMMIT
unset NEW_COMMIT
PROGNAME=$(basename "$0")
USAGE="$PROGNAME -c command -o old_commit -n new_commit

where:
    command:     benchmark command
    old_commit:  the old commit
    new_commit:  the new commit
"

while getopts :hc:o:n: option; do
  case $option in
    h)
      echo "$USAGE"
      exit 0
      ;;
    c) COMMAND=$OPTARG ;;
    o) OLD_COMMIT=$OPTARG ;;
    n) NEW_COMMIT=$OPTARG ;;
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

if [[ -z $COMMAND || -z $OLD_COMMIT || -z $NEW_COMMIT ]]; then
  echo "$PROGNAME: missing mandatory option" >&2
  echo "Try '$(basename "$0") -h' for more information" >&2
  exit 1
fi

CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
if [[ $OLD_COMMIT =~ "HEAD" ]]; then
	OLD_COMMIT=$(git rev-parse $OLD_COMMIT)
fi

if [[ $NEW_COMMIT =~ "HEAD" ]]; then
	NEW_COMMIT=$(git rev-parse $NEW_COMMIT)
fi

set -ex

git checkout $OLD_COMMIT
$COMMAND | tee old.txt
git checkout $NEW_COMMIT
$COMMAND | tee new.txt

benchstat old.txt new.txt

git checkout $CURRENT_BRANCH
rm old.txt
rm new.txt
