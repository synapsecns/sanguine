#!/usr/bin/env bash
# Usage: ./script/wire.sh <walletName> [...options]

trap "echo 'Wiring script terminated by user'; exit" INT
set -o pipefail

# Parse command line arguments
walletName=$1
# Check that the walletName is not empty
if [[ -z "$walletName" ]]; then
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
  log_file=$(mktemp)
  # Pass the rest of the options to forge-script-run utility
  if ! npx fsr ./script/WireSBA.s.sol "$chain" "$walletName" "$@" 2>&1 | tee "$log_file"; then
    status=1
    for pipeline_status in "${PIPESTATUS[@]}"; do
      if [[ "$pipeline_status" -ne 0 ]]; then
        status=$pipeline_status
        break
      fi
    done
    rm -f "$log_file"
    exit "$status"
  fi
  if grep -q "data: 0x" "$log_file"; then
    read -r -p "Multisig calldata printed for $chain. Press Enter to continue or Ctrl-C to stop..." _
  fi
  rm -f "$log_file"
done
