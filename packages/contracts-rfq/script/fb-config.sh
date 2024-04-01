#!/usr/bin/env bash
# Usage: ./script/fb-config.sh <walletName> [...options]

# Trap SIGINT (Ctrl+C) and exit the script
trap "echo 'Script terminated by user'; exit" INT

# Parse command line arguments
walletName=$1
# Check that the walletName is not empty
if [ -z "$walletName" ]; then
  echo "Usage: ./script/fb-config.sh <walletName> [...options]"
  exit 1
fi
# Get the rest of the options
shift 1

yarn fsr script/ConfigureFastBridge.s.sol arbitrum "$walletName" "$@"
yarn fsr script/ConfigureFastBridge.s.sol base "$walletName" "$@"
yarn fsr script/ConfigureFastBridge.s.sol mainnet "$walletName" "$@"
yarn fsr script/ConfigureFastBridge.s.sol optimism "$walletName" "$@"
