# Verifier Devnet

This module provides everything that's needed to spin up a Verifier development network in order to test functionality of Verifier nodes and Module verification requests.

# Provisioner

Package `provisioner` sets up the development network with the addresses of your Verifier nodes and changes the threshold for an accepted verification in the network.

# Sender

Package `sender` sends verification requests to the `SynapseModule`, which will then be picked up by Verifier nodes in the network for verification in the network.

# Manager

Package `manager` wraps both `Provisioner` and `Sender` objects .

# Setup

## Node setup

1. `cd docker/committee-devnet`

2. `docker compose build`

3. `docker compose up -d` to run the devnet chains, which are

1. Anvil Chain A `ChainID=42` (default RPC endpoint: `http://localhost:8042`)
2. Anvil Chain B `ChainID=43` (default RPC url: `http://localhost:8042`)
3. Anvil Chain C: `ChainID=44` (default RPC url: `http://localhost:8042`)

OmniRPC will set up a proxy for these three chains so they're all accessible under one endpoint (`http://localhost:9001/rpc/<chain_id>`), please refer to the OmniRPC documentation for more.

## Provisioner setup

1. ` cd committee/devnet`

2. run `go run main.go`

3. Once the interactive shell comes up, run
`run --config=<PATH_TO_CONFIG> `

## Sender setup (optional)

Simply add the additional flags `--sender --senderconfig=<PATH_TO_SENDER_CONFIG>` to the command above in the Provisioner setup.

Example: `run --config=./config.yml --sender --senderconfig=./senderconfig.yml`
