#!/usr/bin/env bash
# Usage: ./script/utils/get-mainnet-chains.sh

# Trap SIGINT (Ctrl+C) and exit the script
trap "echo 'Script terminated by user'; exit" INT

# Mainnet chains are defined in foundry.toml
# Everything between [rpc_endpoints] and '# Testnets' headers is mainnet
# Every non-empty line is structured as <chainName> = "<rpcUrl>"
# We extract the chain names
awk '/^\[rpc_endpoints\]/,/# Testnets/ { if ($0 ~ /^\S+ = /) print $1 }' foundry.toml | cut -d'=' -f1
