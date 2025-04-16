#!/usr/bin/env bash
# Usage: ./script/utils/get-mainnet-chains.sh

# Trap SIGINT (Ctrl+C) and exit the script
trap "echo 'Script terminated by user'; exit" INT

# Extract mainnet chain names from foundry.toml
# Steps:
# 1. Find section between [rpc_endpoints] and '# Testnets' markers
# 2. Filter lines containing '=' (each represents a chain)
# 3. Remove the equals sign and everything after it to get clean chain names
sed -n '/\[rpc_endpoints\]/,/# Testnets/p' foundry.toml |
  grep '=' |
  sed 's/[ \t]*=.*$//'
