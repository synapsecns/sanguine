#!/usr/bin/env bash
# Usage: ./script/testnet-fb-deploy.sh [...options]

trap "echo 'Deployment script terminated by user'; exit" INT

walletName="testnet_deployer"

npx fsr script/ConfigureFastBridge.s.sol arb_sepolia "$walletName" --sig runTestnet "$@"
npx fsr script/ConfigureFastBridge.s.sol base_sepolia "$walletName" --sig runTestnet "$@"
npx fsr script/ConfigureFastBridge.s.sol bnb_testnet "$walletName" --sig runTestnet "$@"
npx fsr script/ConfigureFastBridge.s.sol eth_sepolia "$walletName" --sig runTestnet "$@"
npx fsr script/ConfigureFastBridge.s.sol op_sepolia "$walletName" --sig runTestnet "$@"
npx fsr script/ConfigureFastBridge.s.sol scroll_sepolia "$walletName" --sig runTestnet "$@"
