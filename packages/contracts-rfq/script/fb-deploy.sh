#!/usr/bin/env bash
# Usage: ./script/fb-deploy.sh <walletName> [...options]

trap "echo 'Deployment script terminated by user'; exit" INT

# Parse command line arguments
walletName=$1
# Check that the walletName is not empty
if [ -z "$walletName" ]; then
  echo "Usage: ./script/fb-deploy.sh <walletName> [...options]"
  exit 1
fi
# Get the rest of the options
shift 1

yarn fsr script/DeployFastBridge.CREATE2.s.sol arbitrum "$walletName" "$@"
yarn fsr script/DeployFastBridge.CREATE2.s.sol base "$walletName" "$@"
yarn fsr script/DeployFastBridge.CREATE2.s.sol mainnet "$walletName" "$@"
yarn fsr script/DeployFastBridge.CREATE2.s.sol optimism "$walletName" "$@"
