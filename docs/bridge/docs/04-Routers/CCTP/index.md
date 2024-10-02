---
sidebar_label: CCTP
---

import { CCTPFlow } from '@site/src/components/CCTPFlow'

# CCTP Router

A [Synapse Router](../Synapse-Router) bridge module which uses Circle's [Cross-Chain Transfer Protocol](https://www.circle.com/en/cross-chain-transfer-protocol) to natively mint & burn USDC.

<figure>
    <CCTPFlow />
    <figcaption>User assets are sent to a Circle contract, moved to the `destChain`, and returned to the user.</figcaption>
</figure>

## Architecture

### Contracts

[Synapse CCTP contracts](/docs/Contracts/CCTP) overlay Circle CCTP contracts to mint and burn USDC and fulfill CCTP transactions.

### Configuration
CCTP can be configured to bridge through any supported liquidity source, such as [Uniswap](https://github.com/synapsecns/synapse-contracts/blob/master/contracts/router/modules/pool/uniswap/UniswapV3Module.sol), [Curve](https://github.com/synapsecns/synapse-contracts/blob/master/contracts/router/modules/pool/curve/CurveV1Module.sol), [Algebra](https://github.com/synapsecns/synapse-contracts/blob/master/contracts/router/modules/pool/algebra/AlgebraModule.sol), [DAI PSM](https://github.com/synapsecns/synapse-contracts/blob/master/contracts/router/modules/pool/dss/DssPsmModule.sol),  and others.

### Relayer

CCTP Relayers allow anyone to coordinate on-chain events and stored message states to send native USDC through SynapseCCTP and Circle's CCPT contracts.

:::tip

While the Synapse CCTP Golang relayer can be run by anyone, and is easily observable, you can also build and run your own relayer permissionlessly in any programming language.

:::

## Behavior

CCTP Relayers poll for new transactions and state updates from CCTP contracts on-chain, to store in an off-chain database.

Attestations from the [Circle API](https://developers.circle.com/stablecoin/reference) are submitted to the destination contract, and marked `Complete` when a transaction receipt is received.

| State       | Description |
|-------------|-------------|
| `Pending`   | Initiated on origin chain, and pending attestation |
| `Attested`  | Waiting for submission on destination chain |
| `Submitted` | Confirmed on destination chain |
| `Complete`  | Completed on destination chain |

<!-- [Message states ↗](https://pkg.go.dev/github.com/synapsecns/sanguine/services/cctp-relayer@v0.10.0/types#MessageState) -->
<!-- [combines](https://medium.com/@matt.denobrega/state-vs-event-based-web-architectures-59ab1f47656b) -->

## Configure

CCTP Relayers require a YAML configuration file path to be provided at run time.

:::note cctp_type

* **`synapse`** (recommended): Uses events & metadata from [Synapse CCTP contracts](/docs/Contracts/CCTP), and `synapse_cctp_address` when configuring `chains`.

* **`circle`** (USDC to USDC only): Uses raw [TokenMessenger](https://github.com/circlefin/evm-cctp-contracts/blob/817397db0a12963accc08ff86065491577bbc0e5/src/TokenMessenger.sol) events, and `token_messenger_address` when configuring `chains`.

:::

### Parameters

* `cctp_type`: Determines which event types and contracts are used.
* `chains`: `chain_id` list
* `base_omnirpc_url`: [OmniRPC service](/docs/Services/Omnirpc) base URL
* `unbonded_signer`: [Signer service](/docs/Services/Signer) — *should be a mounted secret*
* `port`: Relayer port (e.g. 8080)
* `host`: Relayer host (e.g. localhost) — *do not publicly expose*.
* `http_backoff_initial_interval_ms`: Initial backoff interval in milliseconds.
* `retry_interval_ms`: Retry interval between attestation requests in milliseconds — *[CCTP API Rate Limit](https://developers.circle.com/stablecoins/docs/limits)*.

 ### Example

```yaml
cctp_type: "synapse"
# prod contract addresses
chains:
  - chain_id: 1
    synapse_cctp_address: "0x12715a66773BD9C54534a01aBF01d05F6B4Bd35E"
  - chain_id: 42161
    synapse_cctp_address: "0x12715a66773BD9C54534a01aBF01d05F6B4Bd35E"
base_omnirpc_url: "http://omnrpc-url/"
unbonded_signer:
  type: "AWS"
  # should be a mounted secret
  file: "/config/aws.txt"
http_backoff_initial_interval_ms: 1000
http_backoff_max_elapsed_time_ms: 300000
# submitter config for cctp
submitter_config:
  chains:
    1:
      supports_eip_1559: true
      gas_estimate: 1000000
    42161:
      gas_estimate: 30000000
      max_gas_price: 10000000000
      supports_eip_1559: true
```


## Run

### From Docker

Run the Docker [image](https://github.com/synapsecns/sanguine/pkgs/container/sanguine%2Fcctp-relayer) along with the path to your [YAML configuration file](#configure).

1. `docker run ghcr.io/synapsecns/sanguine/cctp-relayer:latest --config /path/to/config.yaml`

### From Source

:::note Requires Go 1.21 or higher

Not generally recommended for end-users.

:::

Clone the Sanguine repository, then run the main.go file along with the path to your [YAML configuration file](#configure).

1. `git clone https://github.com/synapsecns/sanguine --recursive`
2. `cd sanguine/services/cctp-relayer`
3. `go run main.go --config /path/to/config.yaml`

### With Helm

There is a helm chart available for the CCTP Relayer [here](https://artifacthub.io/packages/helm/synapse/cctp/0.2.0), but it is recommended you create your own.

### Recommended services

CCTP Relayer uses open telemetry for tracing and metrics. See the [Observability](/docs/Services/Observability) page for details. We highly recommend setting up the [Submitter Dashboard](/docs/Services/Submitter) as well.
