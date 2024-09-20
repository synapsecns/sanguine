---
sidebar_position: 2
sidebar_label: CCTP
---

# CCTP

Synapse CCTP Router uses Circle's [Cross-Chain Transfer Protocol](https://www.circle.com/en/cross-chain-transfer-protocol) to natively mint & burn USDC on [supported chains](/docs/Contracts/CCTP). It can be run by anyone, and is easily observable.

## Architecture

### Contracts

[Synapse CCTP contracts](/docs/Contracts/CCTP) overlay Circle CCTP contracts to mint and burn USDC and fulfill CCTP transactions.

### Liquidity
As a modular component of [Synapse Router](../Synapse-Router), CCTP can be configured to bridge through any supported liquidity source, such as [Curve](https://github.com/synapsecns/synapse-contracts/blob/885cbe06a960591b1bdef330f3d3d57c49dba8e2/contracts/router/modules/pool/curve/CurveV1Module.sol), [Algebra](https://github.com/synapsecns/synapse-contracts/blob/885cbe06a960591b1bdef330f3d3d57c49dba8e2/contracts/router/modules/pool/algebra/AlgebraModule.sol), [DAI PSM](https://github.com/synapsecns/synapse-contracts/blob/885cbe06a960591b1bdef330f3d3d57c49dba8e2/contracts/router/modules/pool/dss/DssPsmModule.sol),  and others.

### Relayer

Go application which uses on-chain events with off-chain message state to process transactions. CCTP Relayer can be run by anyone, and is easily observable.

## Behavior

CCTP Relayer polls for new transactions and and [message state](https://pkg.go.dev/github.com/synapsecns/sanguine/services/cctp-relayer@v0.10.0/types#MessageState) updates from CCTP contracts on-chain, which are stored in an off-chain database.

Successful attestations fetched from the [Circle API](https://developers.circle.com/stablecoin/reference) are submitted to the destination CCTP contract, and marked as `Complete` once a transaction receipt is received.

| State       | Description |
|-------------|-------------|
| `Pending`   | USDC transfer has been initiated on the origin chain and is pending attestation |
| `Attested`  | USDC transfer is waiting for submittion on the destination chain |
| `Submitted` | USDC transfer has been confirmed on the destination chain |
| `Complete`  | USDC transfer has been completed on the destination chain |

<!-- [Message states ↗](https://pkg.go.dev/github.com/synapsecns/sanguine/services/cctp-relayer@v0.10.0/types#MessageState) -->
<!-- [combines](https://medium.com/@matt.denobrega/state-vs-event-based-web-architectures-59ab1f47656b) -->

## Run a Relayer

### Configure

To run CCTP Relayer, first create your [YAML configuration file](#configuration-file), Which is required as part of the `run` command.

### From Docker

Run the Docker [image](https://github.com/synapsecns/sanguine/pkgs/container/sanguine%2Fcctp-relayer) along with the path to your [YAML configuration file](#configuration-file).

1. `docker run ghcr.io/synapsecns/sanguine/cctp-relayer:latest --config /path/to/config.yaml`

### From Source

:::note Requires Go 1.21 or higher

Not generally recommended for end-users.

:::

Clone the sanguine repository, then run the main.go file along with the path to your [YAML configuration file](#configuration-file).

1. `git clone https://github.com/synapsecns/sanguine --recursive`
2. `cd sanguine/services/cctp-relayer`
3. `go run main.go --config /path/to/config.yaml`

### With Helm

There is a helm chart available for the CCTP Relayer [here](https://artifacthub.io/packages/helm/synapse/cctp/0.2.0), but it is recommended you create your own.

### Recommended services

CCTP Relayer uses open telemetry for tracing and metrics. See the [Observability](/docs/Services/Observability) page for details. We highly recommend setting up the [Submitter Dashboard](/docs/Services/Submitter) as well.

## Configuration

Relayer requires a YAML configuration file at run time.

* `cctp_type`
   * `synapse`: (Recommended) Follows and relays events & metadata from [Synapse CCTP contracts](docs/Contracts/CCTP)
   * `circle`: Relays raw [TokenMessenger](https://github.com/circlefin/evm-cctp-contracts/blob/817397db0a12963accc08ff86065491577bbc0e5/src/TokenMessenger.sol) events — *USDC to USDC only*
 * `chains`: `chain_id` list
   * `synapse`: Use `synapse_cctp_address`
   * `circle`: Use `token_messenger_address`
 * `base_omnirpc_url`: Base url for the [OmniRPC](/docs/Services/Omnirpc) service
 * `unbonded_signer`: [Signer](/docs/Services/Signer) to use for transactions — *should be a mounted secret*
   * `AWS`
   * `File`
   * `GCP`
 * `port`: Relayer port (e.g. 8080)
 * `host`: Relayer host (e.g. localhost) — *do not publicly expose*.
 * `http_backoff_initial_interval_ms`: Initial backoff interval in milliseconds.
 * `retry_interval_ms`: Retry interval between attestation requests in milliseconds — *see [CCTP API Rate Limit](https://developers.circle.com/stablecoins/docs/limits)*.

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
