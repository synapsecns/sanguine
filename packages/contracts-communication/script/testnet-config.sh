#!/usr/bin/env bash
# Usage: ./script/testnet-deploy.sh <chainName> <walletName>
# Trap SIGINT (Ctrl+C) and exit the script
trap "echo 'Script terminated by user'; exit" INT
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
# Synapse contracts
yarn fsr-str script/config/ConfigureSynapseGasOracleV1.s.sol "$chainName" "$walletName" testnet "$@"
yarn fsr-str script/config/ConfigureSynapseModule.s.sol "$chainName" "$walletName" testnet "$@"
yarn fsr-str script/config/ConfigureSynapseExecutionServiceV1.s.sol "$chainName" "$walletName" testnet "$@"
# Client contracts
yarn fsr-str script/config/ConfigureClientV1.s.sol "$chainName" "$walletName" testnet "$@"
yarn fsr script/config/ConfigureExecutionFees.s.sol "$chainName" "$walletName" "$@"
# Ping-Pong App
yarn fsr-str script/config/ConfigurePingPongApp.s.sol "$chainName" "$walletName" testnet "$@"
# Legacy contracts
yarn fsr-str script/config/legacy/ConfigureMessageBus.s.sol "$chainName" "$walletName" testnet "$@"
yarn fsr-str script/config/legacy/ConfigureLegacyPingPong.s.sol "$chainName" "$walletName" testnet "$@"
