#!/bin/bash
docker run --rm -v $(pwd):/src ethereum/solc:0.8.20 --base-path /src --include-path /src/node_modules/ "$@"
