#!/usr/bin/env bash
# Usage: ./script/testnet-fb-deploy.sh [...options]

trap "echo 'Deployment script terminated by user'; exit" INT

walletName="testnet_deployer"

npx fsr script/DeployFastBridge.CREATE2.s.sol arb_sepolia "$walletName" --sig runTestnet "$@"
npx fsr script/DeployFastBridge.CREATE2.s.sol base_sepolia "$walletName" --sig runTestnet "$@"
npx fsr script/DeployFastBridge.CREATE2.s.sol bnb_testnet "$walletName" --sig runTestnet "$@"
npx fsr script/DeployFastBridge.CREATE2.s.sol eth_sepolia "$walletName" --sig runTestnet "$@"
npx fsr script/DeployFastBridge.CREATE2.s.sol op_sepolia "$walletName" --sig runTestnet "$@"
npx fsr script/DeployFastBridge.CREATE2.s.sol scroll_sepolia "$walletName" --sig runTestnet "$@"
