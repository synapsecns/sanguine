#!/usr/bin/env bash
echo "Building synapse-interface with local sdk-router"
cd ../sdk-router || exit
yarn install
cd ../synapse-interface || exit
pwd
yarn add ./packages/sdk-router
echo "added ./packages/sdk-router, building"
npx next build
