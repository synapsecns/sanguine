#!/bin/bash
echo "Building synapse-interface with local sdk-router"
cd ../sdk-router
yarn install
cd ../synapse-interface
yarn add ../sdk-router

