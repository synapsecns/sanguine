#!/bin/bash

# Number of times to run the test
num_runs=$1

# Check if a number is provided
if [[ ! $num_runs =~ ^[0-9]+$ ]]; then
    echo "Usage: $0 <number-of-runs>"
    exit 1
fi

# Run the test specified number of times
for ((i = 1; i <= num_runs; i++)); do
    log_file="test-$i.log"
    echo "Running test iteration $i, logging to $log_file..."
    go test -v -run TestGuardSuite/TestFraudulentStateInSnapshot | tee "$log_file"
done

