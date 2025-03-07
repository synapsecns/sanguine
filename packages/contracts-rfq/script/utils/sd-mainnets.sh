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
  npx sd "$chain" SynapseIntentPreviewer
  npx sd "$chain" SynapseIntentRouter
  npx sd "$chain" TokenZapV1
  npx fvc "$chain" SynapseIntentPreviewer
  npx fvc "$chain" SynapseIntentRouter
  npx fvc "$chain" TokenZapV1
done
