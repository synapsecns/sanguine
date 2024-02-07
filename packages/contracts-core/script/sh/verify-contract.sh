#!/usr/bin/env bash
# This script verifies a deployed contract on a given chain
# Usage: ./script/sh/verify-contract.sh <chainName> <contractName>

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
NC='\033[0m' # No Color

CHAIN_NAME=$1
CONTRACT_NAME=$2

if [ -z "$CHAIN_NAME" ]; then
  echo -e "${RED}Error: Please provide a chain name as the first argument.${NC}"
  exit 1
fi
if [ -z "$CONTRACT_NAME" ]; then
  echo -e "${RED}Error: Please provide a contract name as the second argument.${NC}"
  exit 1
fi

# https://www.shellcheck.net/wiki/SC1091
# shellcheck source=/dev/null
source .env
ETHERSCAN_KEY="ETHERSCAN_${CHAIN_NAME^^}_KEY"
ETHERSCAN_KEY=${!ETHERSCAN_KEY}
if [ -z "$ETHERSCAN_KEY" ]; then
  echo -e "${RED}Error: Please provide an etherscan key for the chain [${CHAIN_NAME}].${NC}"
  exit 1
fi

DEPLOYMENT_FN="deployments/$CHAIN_NAME/$CONTRACT_NAME.json"
if [ ! -e "$DEPLOYMENT_FN" ]; then
  echo -e "${RED}Error: Deployment file not found at [$DEPLOYMENT_FN].${NC}"
  exit 1
fi

ADDRESS=$(jq -r ".address" < "$DEPLOYMENT_FN")
if [ "$ADDRESS" == "null" ]; then
  echo -e "${RED}Error: Contract address not found in [$DEPLOYMENT_FN].${NC}"
  exit 1
fi
CONSTUCTOR_ARGS=$(jq -r ".args" < "$DEPLOYMENT_FN")
if [ "$CONSTUCTOR_ARGS" == "null" ]; then
  echo -e "${YELLOW}No constructor args found in [$DEPLOYMENT_FN].${NC}"
  CONSTUCTOR_ARGS=""
else
  echo -e "${GREEN}Constructor args found in [$DEPLOYMENT_FN].${NC}"
  CONSTUCTOR_ARGS="--constructor-args $CONSTUCTOR_ARGS"
fi

bash -x -c "forge verify-contract --chain $CHAIN_NAME -e $ETHERSCAN_KEY $CONSTUCTOR_ARGS  --watch $ADDRESS $CONTRACT_NAME"
