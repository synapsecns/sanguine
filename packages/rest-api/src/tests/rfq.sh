#!/usr/bin/env bash
# Usage: ./src/tests/rfq.sh <rest-api-URL>
# Exit immediately if a command exits with a non-zero status
set -e
# Trap SIGINT (Ctrl+C) and exit the script
trap "echo 'Script terminated by user'; exit" INT

TESTED_MODULE='SynapseRFQ'
ADDR='0x0000000000000000000000000000000000000001'
GAS_TOKEN='0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE'
ETH_CHAIN_ID=1
AMOUNT=0.001
AMOUNT_RAW=1000000000000000
AMOUNT_HEX=0x038d7ea4c68000
ROUTER_ADDRESS='0x00cD000000003f7F682BE4813200893d4e690000'

RFQ_API_URL='https://rfq-api.omnirpc.io/quotes'

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
  local exact_match=${5:-true}

  # Convert to exact match regex pattern if needed
  if [ "$exact_match" = "true" ]; then
    # Wrap with ^ and $ for exact match
    expected_value="^${expected_value}$"
  fi

  # Use regex comparison
  if [[ ! "$actual_value" =~ $expected_value ]]; then
    echo "    ❌ $TESTED_MODULE calldata.$field_name INCORRECT for $chainName" >&2
    if [ "$exact_match" = "true" ]; then
      # Remove ^ and $ for display
      echo "       Expected: ${expected_value:1:-1}" >&2
    else
      echo "       Expected: $expected_value" >&2
    fi
    echo "         Actual: $actual_value" >&2
    echo "       Quote: $quote" >&2
    exit 1
  fi
}

encode_address() {
  address=$1
  # Add leading zeroes, remove 0x prefix, cast to lowercase
  echo "000000000000000000000000${address:2}" | tr '[:upper:]' '[:lower:]'
}

restApiURL=$1
# Remove trailing slash
restApiURL="${restApiURL%%/}"
if [ -z "$restApiURL" ]; then
  echo "Usage: ./src/tests/rfq.sh <rest-api-URL>" >&2
  exit 1
fi

chains=$(curl -s "https://backend.gas.zip/v2/chains" | jq '.chains')

# Get filtered quotes
quotes=$(curl -s "$RFQ_API_URL" | jq \
  --argjson eth_chain_id "$ETH_CHAIN_ID" \
  --arg gas_token "$GAS_TOKEN" \
  --argjson min_amount "$AMOUNT_RAW" \
  '.[] | select(
    .origin_chain_id == $eth_chain_id and
    .origin_token_addr == $gas_token and
    .dest_token_addr == $gas_token and
    (.max_origin_amount | tonumber) >= $min_amount and
    (.fixed_fee | tonumber) < $min_amount
  )')

# Extract unique destination chain IDs
dest_chains=$(echo "$quotes" | jq -r '.dest_chain_id' | sort -n | uniq)

echo "Testing $AMOUNT ETH from Ethereum Mainnet (chain ID: $ETH_CHAIN_ID) via $TESTED_MODULE"
for chainId in $dest_chains; do
  chain=$(echo "$chains" | jq --arg id "$chainId" '.[] | select(.chain == ($id | tonumber))')
  chainName=$(echo "$chain" | jq -r '.name')
  if [ -z "$chainName" ]; then
    chainName="Unknown chain"
  fi
  # bridge(recipient, chainId, token, amount, ...)
  expectedDataPrefix="0xc2288147$(encode_address "$ADDR")$(printf "%064x" "$chainId")$(encode_address "$GAS_TOKEN")$(printf "%064x" "$AMOUNT_RAW")"
  echo "  To $chainName (chain ID: $chainId)"
  quotes=$(curl -s "$restApiURL/bridge?fromChain=$ETH_CHAIN_ID&toChain=$chainId&fromToken=$GAS_TOKEN&toToken=$GAS_TOKEN&amount=$AMOUNT&originUserAddress=$ADDR&destAddress=$ADDR")
  quotesType=$(echo "$quotes" | jq -r 'type')
  # Check if the response is not an array
  if [ "$quotesType" != "array" ]; then
    echo "    ❌ Response is not an array for $chainName: $quotes" >&2
    exit 1
  fi
  # Check that module quote exists
  # Select quote with maximum amountOut value
  quote=$(echo "$quotes" | jq --arg TESTED_MODULE "$TESTED_MODULE" '[ .[] | select(.bridgeModuleName == $TESTED_MODULE) ] | sort_by(.maxAmountOutStr | tonumber) | last')
  # Extract fields and check if they exist
  maxAmountOut=$(extract_json_field "$quote" "maxAmountOutStr" "quote")
  calldata=$(extract_json_field "$quote" "callData" "quote")
  # Extract calldata fields
  to=$(extract_json_field "$calldata" "to" "callData")
  value_hex=$(extract_json_field "$calldata" "value.hex" "callData")
  data=$(extract_json_field "$calldata" "data" "callData")
  # Check if the extracted fields are correct
  check_field "to" "$to" "$ROUTER_ADDRESS" "$quote"
  check_field "value.hex" "$value_hex" "$AMOUNT_HEX" "$quote"
  check_field "data" "$data" "$expectedDataPrefix.*" "$quote" false
  echo "    ✅ $TESTED_MODULE quote found for $chainName: $maxAmountOut"
done

echo "✅ $TESTED_MODULE: all tests passed"
