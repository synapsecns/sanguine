#!/usr/bin/env bash
# Usage: ./script/testnet-deploy.sh <chainName> <walletName>
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

yarn fsr-str script/deploy/DeployNoArgs.s.sol "$chainName" "$walletName" "InterchainDB" "$@"

yarn fsr-str script/deploy/DeployWithMsgSender.s.sol "$chainName" "$walletName" "PingPongApp" "$@"

yarn fsr script/deploy/DeployInterchainClientV1.s.sol "$chainName" "$walletName" "$@"
yarn fsr script/deploy/DeploySynapseModule.s.sol "$chainName" "$walletName" "$@"

yarn fsr-str script/deploy/DeployWithMsgSender.s.sol "$chainName" "$walletName" "ExecutionFees" "$@"
yarn fsr script/deploy/DeploySynapseExecutionServiceV1.s.sol "$chainName" "$walletName" "$@"

yarn fsr-str script/deploy/DeployWithMsgSender.s.sol "$chainName" "$walletName" "SynapseGasOracleV1" "$@"
