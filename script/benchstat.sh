#!/bin/bash

unset COMMAND
unset OLD_COMMIT
unset NEW_COMMIT
unset COUNT
unset INTERLEAVE
PROGNAME=$(basename "$0")
USAGE="$PROGNAME -f flags -c count -o old_commit -n new_commit [-i]

where:
    -f flags:       benchmark flags (e.g. -test.v -test.run=^$ -test.bench=BenchmarkGetAccount)
    -o old_commit:  the old commit
    -n new_commit:  the new commit
    -c count:       the number of benchmark runs
    -i:             interleave mode
"

while getopts :hf:o:n:c:i option; do
  case $option in
    h)
      echo "$USAGE"
      exit 0
      ;;
    f) COMMAND=$OPTARG ;;
    o) OLD_COMMIT=$OPTARG ;;
    n) NEW_COMMIT=$OPTARG ;;
    c) COUNT=$OPTARG ;;
    i) INTERLEAVE=true ;;
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

if [[ -z $COMMAND || -z $OLD_COMMIT || -z $NEW_COMMIT || -z $COUNT ]]; then
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

if [[ "$INTERLEAVE" = "true" ]]; then
  git checkout $OLD_COMMIT
  go test -c -o old.test
  OLD_COMMAND="./old.test $COMMAND"

  git checkout $NEW_COMMIT
  go test -c -o new.test
  NEW_COMMAND="./new.test $COMMAND"

  for i in $(seq 1 $COUNT); do
    $OLD_COMMAND | tee -a old.txt
    $NEW_COMMAND | tee -a new.txt
  done

else
  COMMAND="go test $COMMAND -test.count=$COUNT"
  git checkout $OLD_COMMIT
  $COMMAND | tee old.txt
  git checkout $NEW_COMMIT
  $COMMAND | tee new.txt
fi

benchstat old.txt new.txt

git checkout $CURRENT_BRANCH
rm old.txt
rm new.txt

if [[ "$INTERLEAVE" = "true" ]]; then
  rm old.test new.test
fi
