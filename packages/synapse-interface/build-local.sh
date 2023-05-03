#!/usr/bin/env bash
echo "Building synapse-interface with local sdk-router"
pwd
yarn install
yarn add @synapsecns/sdk-router
echo "Added sdk-router, building site..."
