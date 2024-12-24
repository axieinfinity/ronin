#!/bin/bash
# 
# Boot a simulation network and start benchmarking test
# Export the logs, peers info, and DHT info to files for visualization.

main_cmd="go run ."
p2psim_cmd="p2psim"

if ! which p2psim &>/dev/null; then
    fail "missing p2psim binary (you need to build cmd/p2psim and put it in \$PATH)"
fi

# Number of nodes to start for each batch
distribution=(150 100 100)

# Rate of dirty node that not compatible with the valid node
dirty_rate=60

# Rate of valid node but only accept outbound connection
only_outbound_rate=20

# Interval between each node creation
node_creation_interval=1s

# Sleep time between each batch
sleep_time=1200

# Number of bootnodes
num_bootnodes=2

# Other flags
other=""

# Test name
test_name="discovery_benchmark"

# Directory to store results
results_dir="./results"

# Parse the arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        --dirty.rate)
            dirty_rate=$2
            shift 2
            ;;
        --only.outbound.rate)
            only_outbound_rate=$2
            shift 2
            ;;
        --interval)
            node_creation_interval=$2
            shift 2
            ;;
        --distribution)
            IFS=',' read -r -a distribution <<< "$2"
            shift 2
            ;;
        --sleep)
            sleep_time=$2
            shift 2
            ;;
        --num.bootnodes)
            num_bootnodes=$2
            shift 2
            ;;
        --disable.enrfilter)
            other+=" --enable.enrfilter false"
            shift
            ;;
        --testname)
            test_name=$2
            shift 2
            ;;
        --results.dir)
            results_dir=$2
            shift 2
            ;;
        --help)
            echo "USAGE: $0 [OPTIONS]"
            echo "OPTIONS:"
            echo "  --dirty.rate <rate>              Rate of dirty node that not compatible with the valid node (default 60 means 60% dirty nodes)"
            echo "  --only.outbound.rate <rate>      Rate of valid node but only accept outbound connection (default 20 means 20% nodes only accept outbound connection)"
            echo "  --interval <interval>            Interval between each node creation (default 1s)"
            echo "  --distribution <distribution>    Number of nodes to start for each batch (default 150,100,100)"
            echo "  --sleep <time>                   Sleep time between each batch (default 1200)"
            echo "  --num.bootnodes <num>            Number of bootnodes (default 2)"
            echo "  --disable.enrfilter              Disable ENR filter"
            echo "  --testname <name>                Test name"
            echo "  --results.dir <dir>              Directory to store results"
            echo "  --help                           Show this help"
            exit 0
            ;;
        *)
            echo "Unknown argument $1"
            exit 1
            ;;
    esac
done

benchmark() {
    # Create results directory if not exists
    mkdir -p $results_dir

    # Output files
    log_file="$results_dir/$test_name.log"
    err_file="$results_dir/$test_name.err"
    stats_file="$results_dir/stats_$test_name.csv"
    dht_file="dht_$test_name.log"
    peers_file="peers_$test_name.log"

    # Start p2psim server and log tracker
    echo "Start server $test_name..."
    $main_cmd > $log_file 2> $err_file &
    echo "Start stats $test_name..."
    $p2psim_cmd log-stats --file $stats_file &

    # Start bootnodes
    echo "Start bootnodes $test_name..."
    $p2psim_cmd node create-multi --count $num_bootnodes --node.type bootnode $other

    # Roll out batches
    for num_node in ${distribution[@]}; do
        # Roll out nodes
        echo "Start $num_node nodes..."
        $p2psim_cmd node create-multi --count $num_node --autofill.bootnodes --interval $node_creation_interval --dirty.rate $dirty_rate --only.outbound.rate $only_outbound_rate $other

        # Wait for a while until the network is stable
        echo "Sleep $sleep_time..."
        sleep $sleep_time
    done

    # Export nodes in dht and peers info of all nodes
    $p2psim_cmd network-stats dht > $dht_file
    $p2psim_cmd network-stats peers > $peers_file

    # Kill p2psim server and log tracker
    echo "Kill server and stats $test_name..."
    kill -9 $(lsof -t -i:8888)
    ps aux | grep "p2psim" | grep -v "grep" |  awk '{print $2}' | xargs kill -9

    # Sleep for a while to make sure the server is killed
    sleep 10
}

benchmark
