#!/bin/bash

# Takes the raw transaction as input
rawTx=$1

# Decode the transaction using cast dt and store the result in a variable
decodedTx=$(cast dt $rawTx)

# Check if the decoding was successful
if [[ -z "$decodedTx" ]]; then
    echo "Error: Decoding failed. Invalid raw transaction?"
    exit 1
fi

# Use jq to pretty print the decoded transaction
echo "$decodedTx" | jq '.'

# Extract the type of the transaction
type=$(echo "$decodedTx" | jq -r '.type')

# Check if type is present
if [[ -z "$type" ]]; then
    echo "Error: Transaction type not found."
    exit 1
fi

# Convert the type from hexadecimal to decimal
txType=$(cast --to-dec $type)

# Extract common fields
gas=$(echo "$decodedTx" | jq -r '.gas')
value=$(echo "$decodedTx" | jq -r '.value')
nonce=$(echo "$decodedTx" | jq -r '.nonce')

# Check if gas, value, or nonce are empty
if [[ -z "$gas" || -z "$value" || -z "$nonce" ]]; then
    echo "Error: Missing required fields (gas, value, or nonce)."
    exit 1
fi

echo "Converted Values:"
echo "Type: $txType"
echo "Gas: $(cast --to-dec $gas)"
echo "Value: $(cast --to-unit $value ether)"
echo "Nonce: $(cast --to-dec $nonce)"

# Conditional logic based on transaction type
if [[ "$txType" == "0" ]]; then
    # Transaction type 0 specific field
    gasPrice=$(echo "$decodedTx" | jq -r '.gasPrice')
    if [[ -z "$gasPrice" ]]; then
        echo "Error: Missing gasPrice for transaction type 0."
        exit 1
    fi
    echo "GasPrice: $(cast --to-unit $gasPrice gwei)"
elif [[ "$txType" == "2" ]]; then
    # Transaction type 2 specific fields
    maxPriorityFeePerGas=$(echo "$decodedTx" | jq -r '.maxPriorityFeePerGas')
    maxFeePerGas=$(echo "$decodedTx" | jq -r '.maxFeePerGas')

    if [[ -z "$maxPriorityFeePerGas" || -z "$maxFeePerGas" ]]; then
        echo "Error: Missing fee fields (maxPriorityFeePerGas or maxFeePerGas) for transaction type 2."
        exit 1
    fi

    echo "MaxPriorityFeePerGas: $(cast --to-unit $maxPriorityFeePerGas gwei)"
    echo "MaxFeePerGas: $(cast --to-unit $maxFeePerGas gwei)"
fi
