#!/usr/bin/env bash
# This script verifies Messaging Contracts on a given chain
# Usage: ./script/sh/verify-contract.sh <chainName>

RED='\033[0;31m'
NC='\033[0m' # No Color

CHAIN_NAME=$1

if [ -z "$CHAIN_NAME" ]; then
  echo -e "${RED}Error: Please provide a chain name as the first argument.${NC}"
  exit 1
fi

./script/sh/verify-contract.sh "$CHAIN_NAME" "BondingManager"
./script/sh/verify-contract.sh "$CHAIN_NAME" "Destination"
./script/sh/verify-contract.sh "$CHAIN_NAME" "GasOracle"
./script/sh/verify-contract.sh "$CHAIN_NAME" "Inbox"
./script/sh/verify-contract.sh "$CHAIN_NAME" "Origin"
./script/sh/verify-contract.sh "$CHAIN_NAME" "Summit"
