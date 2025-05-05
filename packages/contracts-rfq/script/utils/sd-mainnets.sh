#!/usr/bin/env bash
# Usage: ./script/utils/sd-mainnets.sh

# Trap SIGINT (Ctrl+C) and exit the script
trap "echo 'Script terminated by user'; exit" INT

# Get the mainnet chains
chains=$(./script/utils/get-mainnet-chains.sh)
# Transform the chains separated by newlines into a comma separated list
chainsPretty="[$(echo "$chains" | tr '\n' ',' | sed 's/,$//' | sed 's/,/, /g')]"
echo "Saving deployments on chains: $chainsPretty"
for chain in $chains; do
  ./script/utils/sd-sir.sh "$chain"
done
