#!/bin/bash

# Define the base directories
ARTIFACTS_DIR="./out"
OUTPUT_DIR="./rfq-quoter/bindings"

# List of contract names to process
CONTRACT_NAMES=("FastBridge" "MockERC20") # Replace with your contract names

cd contracts
forge build --extra-output-files bin abi --force 
cd ..

# Loop over the list of contract names
for CONTRACT_NAME in "${CONTRACT_NAMES[@]}"; do
    ABI_FILE="${ARTIFACTS_DIR}/${CONTRACT_NAME}.sol/${CONTRACT_NAME}.abi.json"
    BIN_FILE="${ARTIFACTS_DIR}/${CONTRACT_NAME}.sol/${CONTRACT_NAME}.bin"

    # Check if ABI and BIN files exist
    if [[ -f "$ABI_FILE" && -f "$BIN_FILE" ]]; then
        # Generate Go bindings
        echo "Generating bindings for $CONTRACT_NAME..."
        abigen --abi="$ABI_FILE" --bin="$BIN_FILE" --pkg="bindings" --out="${OUTPUT_DIR}/${CONTRACT_NAME}.go" --type="$CONTRACT_NAME"
        echo "Bindings for $CONTRACT_NAME generated."
    else
        echo "ABI or BIN file for $CONTRACT_NAME not found."
    fi
done

echo "Binding generation complete."
