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
options=$@

yarn sol-run-str script/deploy/DeployNoArgs.s.sol $chainName $walletName "InterchainDB" $options

yarn sol-run-str script/deploy/DeployWithMsgSender.s.sol $chainName $walletName "InterchainAppExample" $options

yarn sol-run script/deploy/DeployInterchainClientV1.s.sol $chainName $walletName $options
yarn sol-run script/deploy/DeploySynapseModule.s.sol $chainName $walletName $options

yarn sol-run-str script/deploy/DeployWithMsgSender.s.sol $chainName $walletName "ExecutionFees" $options
yarn sol-run-str script/deploy/DeployWithMsgSender.s.sol $chainName $walletName "ExecutionService" $options

yarn sol-run-str script/deploy/DeployNoArgs.s.sol $chainName $walletName "GasOracleMock" $options