#!/bin/bash

# Takes the raw transaction as input
rawTx=$1

# Decode the transaction using cast dt and store the result in a variable
decodedTx=$(cast dt $rawTx)

# Use jq to pretty print the decoded transaction
echo "$decodedTx" | jq '.'

# Extract the type of the transaction
type=$(echo "$decodedTx" | jq -r '.type')

# Convert the type from hexadecimal to decimal
txType=$(cast --to-dec $type)

# Extract common fields
gas=$(echo "$decodedTx" | jq -r '.gas')
value=$(echo "$decodedTx" | jq -r '.value')
nonce=$(echo "$decodedTx" | jq -r '.nonce')

echo "Converted Values:"
echo "Type: $txType"
echo "Gas: $(cast --to-dec $gas)"
echo "Value: $(cast --to-unit $value ether)"
echo "Nonce: $(cast --to-dec $nonce)"

# Conditional logic based on transaction type
if [[ "$txType" == "0" ]]; then
    # Transaction type 0 specific field
    gasPrice=$(echo "$decodedTx" | jq -r '.gasPrice')
    echo "GasPrice: $(cast --to-unit $gasPrice gwei)"
elif [[ "$txType" == "2" ]]; then
    # Transaction type 2 specific fields
    maxPriorityFeePerGas=$(echo "$decodedTx" | jq -r '.maxPriorityFeePerGas')
    maxFeePerGas=$(echo "$decodedTx" | jq -r '.maxFeePerGas')
    echo "MaxPriorityFeePerGas: $(cast --to-unit $maxPriorityFeePerGas gwei)"
    echo "MaxFeePerGas: $(cast --to-unit $maxFeePerGas gwei)"
fi
