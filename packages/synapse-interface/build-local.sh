#!/bin/bash
echo "Building synapse-interface with local sdk-router"
cd ../sdk-router
yarn install
cd ../synapse-interface
pwd
yarn add ./packages/sdk-router

