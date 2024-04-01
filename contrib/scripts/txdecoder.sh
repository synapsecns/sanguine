#!/bin/bash

# Takes the transaction hash as input
rawTx=$1

# Decode the transaction using cast dt and store the result in a variable
decodedTx=$(cast dt $rawTx)

# Use jq to pretty print the decoded transaction
echo "$decodedTx" | jq "."

# Use jq to extract the values from the decoded transaction
type=$(echo "$decodedTx" | jq -r ".type")
gas=$(echo "$decodedTx" | jq -r ".gas")
value=$(echo "$decodedTx" | jq -r ".value")
nonce=$(echo "$decodedTx" | jq -r ".nonce")
maxPriorityFeePerGas=$(echo "$decodedTx" | jq -r ".maxPriorityFeePerGas")
maxFeePerGas=$(echo "$decodedTx" | jq -r ".maxFeePerGas")

# Note: The actual conversion commands will depend on the capabilities of "cast"
echo "Converted Values:"
echo "Type: $(cast --to-dec $type)" # This specific conversion might not make sense for "type", as it"s not a value field.
echo "Gas: $(cast --to-dec $gas)"
echo "Value: $(cast --to-unit $value ether)"
echo "Nonce: $(cast --to-dec $nonce)"
echo "MaxPriorityFeePerGas: $(cast --to-unit $maxPriorityFeePerGas gwei)"
echo "MaxFeePerGas: $(cast --to-unit $maxFeePerGas gwei)"
