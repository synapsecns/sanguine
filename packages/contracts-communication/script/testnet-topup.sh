#!/usr/bin/env bash
# Usage: ./script/testnet-topup.sh <chainName> <walletName>
# Trap SIGINT (Ctrl+C) and exit the script
trap "echo 'Script terminated by user'; exit" INT
# Parse command line arguments
chainName=$1
walletName=$2
# Check that both arguments are provided
if [ -z "$chainName" ] || [ -z "$walletName" ]; then
  echo "Usage: ./script/testnet-deploy.sh <chainName> <walletName>"
  exit 1
fi
# Get the rest of the options
shift 2
# Top up PingPong apps
yarn fsr-str script/misc/TopUpGas.s.sol "$chainName" "$walletName" PingPongApp "$@"
yarn fsr-str script/misc/TopUpGas.s.sol "$chainName" "$walletName" LegacyPingPong "$@"
