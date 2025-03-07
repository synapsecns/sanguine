#!/usr/bin/env bash
# Usage: ./script/utils/run-mainnets.sh <scriptPath> <walletName> [...options]

# Trap SIGINT (Ctrl+C) and exit the script
trap "echo 'Script terminated by user'; exit" INT

# Parse command line arguments
scriptPath=$1
walletName=$2
# Check that all arguments are not empty
if [ -z "$scriptPath" ] || [ -z "$walletName" ]; then
  echo "Usage: ./script/utils/run-mainnets.sh <scriptPath> <walletName> [...options]"
  exit 1
fi
# Get the rest of the options
shift 2

# Get the mainnet chains
chains=$(./script/utils/get-mainnet-chains.sh)
# Transform the chains separated by newlines into a comma separated list
chainsPretty="[$(echo "$chains" | tr '\n' ',' | sed 's/,$//' | sed 's/,/, /g')]"
echo "Running $scriptPath as $walletName on chains: $chainsPretty"
for chain in $chains; do
  npx fsr "$scriptPath" "$chain" "$walletName" "$@"
done
