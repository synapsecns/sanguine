#!/usr/bin/env bash
# Usage: ./script/transfer-ownership.sh <walletName> [...options]

trap "echo 'Transfer ownership script terminated by user'; exit" INT

walletName=$1
if [[ -z "$walletName" ]]; then
  echo "Usage: ./script/transfer-ownership.sh <walletName> [...options]"
  exit 1
fi
shift 1

chains=$(jq -r '.chains | keys[]' devops.json)
echo Will transfer ownership for chains: "$(echo "$chains" | tr ' ' ',')"

for chain in $chains; do
  echo Transferring ownership on chain: "$chain"
  npx fsr ./script/TransferOwnershipSBA.s.sol "$chain" "$walletName" "$@"
done
