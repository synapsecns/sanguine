#!/usr/bin/env bash
echo "Building synapse-interface with local sdk-router"
cd ../sdk-router || exit
yarn install
cd ../synapse-interface || exit
pwd
yarn add @synapsecns/synapse-interface@0.1.15
echo "Added ./packages/sdk-router, building..."

