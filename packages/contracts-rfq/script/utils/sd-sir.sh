#!/usr/bin/env bash
# Usage: ./script/utils/sd-sir.sh <chain>

# Trap SIGINT (Ctrl+C) and exit the script
trap "echo 'Script terminated by user'; exit" INT

chain=$1
if [ -z "$chain" ]; then
  echo "Usage: ./script/utils/sd-sir.sh <chain>"
  exit 1
fi

npx sd "$chain" SynapseIntentPreviewer
npx sd "$chain" SynapseIntentRouter
npx sd "$chain" TokenZapV1
npx sd "$chain" FastBridgeInterceptor

npx fvc "$chain" SynapseIntentPreviewer
npx fvc "$chain" SynapseIntentRouter
npx fvc "$chain" TokenZapV1
npx fvc "$chain" FastBridgeInterceptor
