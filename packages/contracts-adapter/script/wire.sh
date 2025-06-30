#!/usr/bin/env bash
# Usage: ./script/wire.sh <walletName> [...options]

trap "echo 'Wiring script terminated by user'; exit" INT

# Parse command line arguments
walletName=$1
# Check that the walletName is not empty
if [ -z "$walletName" ]; then
  echo "Usage: ./script/wire.sh <walletName> [...options]"
  exit 1
fi
# Get the rest of the options
shift 1

chains=$(jq -r '.chains | keys[]' devops.json)
# Print the comma separated list of chain names
echo Will be wiring chains: "$(echo "$chains" | tr ' ' ',')"

for chain in $chains; do
  echo Wiring chain: "$chain"
  # Pass the rest of the options to forge-script-run utility
  npx fsr ./script/WireSBA.s.sol "$chain" "$walletName" "$@"
done
