#!/usr/bin/env bash
echo "Building synapse-interface with local sdk-router"
cd ../sdk-router || exit
pwd
yarn install
cd ../synapse-interface || exit
pwd
yarn add ./packages/sdk-router
echo "Added sdk-router, building site..."
