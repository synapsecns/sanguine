#!/usr/bin/env bash
echo "Building synapse-interface with local sdk-router"
cd ../sdk-router || exit
yarn install
cd ../synapse-interface || exit
pwd
yarn add file:./packages/sdk-router
echo "Added sdk-router, building site..."

