#!/usr/bin/env bash
# Usage: ./script/deploy.sh <walletName> [...options]

trap "echo 'Deployment script terminated by user'; exit" INT

# Parse command line arguments
walletName=$1
# Check that the walletName is not empty
if [ -z "$walletName" ]; then
  echo "Usage: ./script/deploy.sh <walletName> [...options]"
  exit 1
fi
# Get the rest of the options
shift 1

chains=$(cat devops.json | jq -r '.chains | keys[]')
# Print the comma separated list of chain names
echo Will be deploying to chains: $(echo $chains | tr ' ' ',')

for chain in $chains; do
  echo Deploying to chain: $chain
  # Pass the rest of the options to forge-script-run utility
  npx fsr ./script/DeploySBA.s.sol "$chain" "$walletName" "$@"
done
