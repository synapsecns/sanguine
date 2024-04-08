#!/usr/bin/env bash
# Usage: ./script/testnet-deploy.sh <chainName> <walletName>
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
# Permisionless DB w/o governance
yarn fsr-str script/deploy/DeployNoArgs.s.sol "$chainName" "$walletName" InterchainDB "$@"
# Synapse contracts
yarn fsr script/deploy/DeploySynapseModule.s.sol "$chainName" "$walletName" "$@"
yarn fsr script/deploy/DeploySynapseExecutionServiceV1.s.sol "$chainName" "$walletName" "$@"
yarn fsr-str script/deploy/DeployWithMsgSender.s.sol "$chainName" "$walletName" SynapseGasOracleV1 "$@"
# Verify Proxy on Etherscan
yarn vp "$chainName" SynapseExecutionServiceV1
# Client contracts
yarn fsr script/deploy/DeployInterchainClientV1.s.sol "$chainName" "$walletName" "$@"
yarn fsr-str script/deploy/DeployWithMsgSender.s.sol "$chainName" "$walletName" ExecutionFees "$@"
# Ping-Pong App
yarn fsr-str script/deploy/DeployWithMsgSender.s.sol "$chainName" "$walletName" PingPongApp "$@"
# Legacy contracts
yarn fsr-str script/deploy/DeployWithMsgSender.s.sol "$chainName" "$walletName" MessageBus "$@"
yarn fsr-str script/deploy/DeployWithMsgSender.s.sol "$chainName" "$walletName" LegacyPingPong "$@"