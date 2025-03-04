#!/usr/bin/env bash
# Usage: ./src/tests/allModules.sh <rest-api-URL>
./src/tests/gasZip.sh "$@"
./src/tests/rfq.sh "$@"
