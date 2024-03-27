#!/usr/bin/env bash
# Usage: ./script/testnet-deploy.sh <chainName> <walletName>
# Parse command line arguments
chainName=$1
walletName=$2
# Check that both arguments are provided
if [ -z "$chainName" ] || [ -z "$walletName" ]; then
  echo "Usage: ./script/testnet-config.sh <chainName> <walletName>"
  exit 1
fi
# Get the rest of the options
shift 2

yarn fsr-str script/config/ConfigureClientV1.s.sol "$chainName" "$walletName" "testnet" "$@"
yarn fsr script/config/ConfigureExecutionFees.s.sol "$chainName" "$walletName" "$@"
yarn fsr-str script/config/ConfigureSynapseGasOracleV1.s.sol "$chainName" "$walletName" "testnet" "$@"
yarn fsr-str script/config/ConfigureSynapseExecutionServiceV1.s.sol "$chainName" "$walletName" "testnet" "$@"
yarn fsr-str script/config/ConfigurePingPongApp.s.sol "$chainName" "$walletName" "testnet" "$@"
yarn fsr-str script/config/ConfigureSynapseModule.s.sol "$chainName" "$walletName" "testnet" "$@"
