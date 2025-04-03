#!/usr/bin/env bash
# Usage: ./src/tests/gasZip.sh <rest-api-URL>
# Exit immediately if a command exits with a non-zero status
set -e
# Trap SIGINT (Ctrl+C) and exit the script
trap "echo 'Script terminated by user'; exit" INT

TESTED_MODULE='Gas.zip'
ADDR='0x0000000000000000000000000000000000000001'
GAS_TOKEN='0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE'
# List of possible inbound chains from https://www.gas.zip (without Ethereum)
# Must match chains names in https://backend.gas.zip/v2/chains
CHAIN_NAMES=(
  'Arbitrum One'
  'Avalanche'
  'Base Mainnet'
  'Berachain'
  'Blast'
  'BSC Mainnet'
  'Cronos'
  'Fantom'
  'Hyperliquid EVM'
  'Metis'
  'OP Mainnet'
  'Polygon'
  'zkScroll'
  'Unichain'
  'WorldChain'
)
ETH_CHAIN_ID=1
AMOUNT=0.001
AMOUNT_HEX=0x038d7ea4c68000
DEPOSIT_ADDRESS='0x391E7C679d29bD940d63be94AD22A25d25b5A604'

# Extract a field from JSON and check if it exists.
extract_json_field() {
  local json_object=$1
  local field_name=$2
  local parent_name=$3

  local value
  value=$(echo "$json_object" | jq -r ".$field_name")

  if [ -z "$value" ] || [ "$value" = "null" ]; then
    echo "    ❌ $field_name not found in $parent_name for $chainName" >&2
    echo "       $parent_name: $json_object" >&2
    exit 1
  fi

  echo "$value"
}

# Check if a field matches the expected value.
check_field() {
  local field_name=$1
  local actual_value=$2
  local expected_value=$3
  local quote=$4

  if [ "$actual_value" != "$expected_value" ]; then
    echo "    ❌ $TESTED_MODULE calldata.$field_name INCORRECT for $chainName" >&2
    echo "       Expected: $expected_value" >&2
    echo "         Actual: $actual_value" >&2
    echo "       Quote: $quote" >&2
    exit 1
  fi
}

restApiURL=$1
# Remove trailing slash
restApiURL="${restApiURL%%/}"
if [ -z "$restApiURL" ]; then
  echo "Usage: ./src/tests/gasZip.sh <rest-api-URL>" >&2
  exit 1
fi

chains=$(curl -s "https://backend.gas.zip/v2/chains" | jq '.chains')

echo "Testing $AMOUNT ETH from Ethereum Mainnet (chain ID: $ETH_CHAIN_ID) via $TESTED_MODULE"
for chainName in "${CHAIN_NAMES[@]}"; do
  # Find element in chains array that has "name" equal to chainName
  chain=$(echo "$chains" | jq --arg chainName "$chainName" '.[] | select(.name == $chainName)')
  if [ -z "$chain" ]; then
    echo "  ❌ Chain not found: $chainName" >&2
    exit 1
  fi
  chainId=$(echo "$chain" | jq -r '.chain')
  short=$(echo "$chain" | jq -r '.short')
  expectedData="0x01$(printf "%04x" "$short")"
  echo "  To $chainName (chain ID: $chainId)"
  quotes=$(curl -s "$restApiURL/bridge?fromChain=$ETH_CHAIN_ID&toChain=$chainId&fromToken=$GAS_TOKEN&toToken=$GAS_TOKEN&amount=$AMOUNT&originUserAddress=$ADDR&destAddress=$ADDR")
  quotesType=$(echo "$quotes" | jq -r 'type')
  # Check if the response is not an array
  if [ "$quotesType" != "array" ]; then
    echo "    ❌ Response is not an array for $chainName: $quotes" >&2
    exit 1
  fi
  # Check that module quote exists
  quote=$(echo "$quotes" | jq --arg TESTED_MODULE "$TESTED_MODULE" '.[] | select(.bridgeModuleName == $TESTED_MODULE)')
  if [ -z "$quote" ]; then
    echo "    ❌ $TESTED_MODULE quote not found for $chainName" >&2
    exit 1
  fi
  # Extract fields and check if they exist
  maxAmountOut=$(extract_json_field "$quote" "maxAmountOutStr" "quote")
  calldata=$(extract_json_field "$quote" "callData" "quote")
  # Extract calldata fields
  to=$(extract_json_field "$calldata" "to" "callData")
  value_hex=$(extract_json_field "$calldata" "value.hex" "callData")
  data=$(extract_json_field "$calldata" "data" "callData")
  # Check if the extracted fields are correct
  check_field "to" "$to" "$DEPOSIT_ADDRESS" "$quote"
  check_field "value.hex" "$value_hex" "$AMOUNT_HEX" "$quote"
  check_field "data" "$data" "$expectedData" "$quote"
  echo "    ✅ $TESTED_MODULE quote found for $chainName: $maxAmountOut"
done

echo "✅ $TESTED_MODULE: all tests passed"
