# CCTP Relayer

The CCTP relayer is an off-chain service aimed at fulfilling transactions requested through the [CCTP Contracts](./Contracts.md). The relayer is responsible for fetching attestations from the [Circle API](https://developers.circle.com/stablecoin/reference) and submitting them to the CCTP contracts. Anyone can run a relayer.

### Architecture

The relayer is a Golang application that polls for events on chain and uses a combo state (db status) and event (on-chain logs) driven [architecture](https://medium.com/@matt.denobrega/state-vs-event-based-web-architectures-59ab1f47656b) to process transactions. The relayer is designed to be run by anyone and be easily observable.

At a high level, the relayer works like this:

1. Poll for new transactions from the CCTP contracts and add them to the database as [Pending](https://pkg.go.dev/github.com/synapsecns/sanguine/services/cctp-relayer@v0.10.0/types#MessageState)
2. Fetch the attestation from the Circle API. Once successful add attestation to the database and update status to be [Attested](https://pkg.go.dev/github.com/synapsecns/sanguine/services/cctp-relayer@v0.10.0/types#MessageState)
3. Submit the attestation to the CCTP contracts. Once the transaction has been added to [Submitter](../Services/Submitter#Observability), mark as [Submitted](https://pkg.go.dev/github.com/synapsecns/sanguine/services/cctp-relayer@v0.10.0/types#MessageState)
4. Poll for the transaction receipt and mark as [Confirmed](https://pkg.go.dev/github.com/synapsecns/sanguine/services/cctp-relayer@v0.10.0/types#MessageState)

### Modes

As specified by the [cctp_type](#Configuration), the CCTP relayer can be run in one of two modes. In [Synapse mode](https://pkg.go.dev/github.com/synapsecns/sanguine/services/cctp-relayer@v0.10.0/types#MessageType), the [Synapse CCTP](./Contracts.md)contracts are listened to and events relayed through there (including metadata). In [Circle Mode](https://pkg.go.dev/github.com/synapsecns/sanguine/services/cctp-relayer@v0.10.0/types#MessageType), raw [TokenMessenger](https://github.com/circlefin/evm-cctp-contracts/blob/817397db0a12963accc08ff86065491577bbc0e5/src/TokenMessenger.sol) events are relayed. This mode can only be used for USDC to USDC bridges and is not commonly used.

## Running the Relayer

### Building From Source

To build the CCTP Relayer from source, you will need to clone the repository and run the main.go file with the config file. Building from source requires go 1.21 or higher and is generally not recommended for end-users.

1. `git clone https://github.com/synapsecns/sanguine --recursive`
2. `cd sanguine/services/cctp-relayer`
3. `go run main.go --config /path/to/config.yaml`

### Running the Docker Image

The CCTP Relayer can also be run with docker. To do this, you will need to pull the [docker image](https://github.com/synapsecns/sanguine/pkgs/container/sanguine%2Fcctp-relayer) and run it with the config file:

```bash
docker run ghcr.io/synapsecns/sanguine/cctp-relayer:latest --config /path/to/config
```

There is also a helm chart available for the CCTP Relayer [here](https://artifacthub.io/packages/helm/synapse/cctp/0.2.0), but it is recommended you create your own.

### Configuration

The CCTP Relayer is configured with a yaml file. The following is an example configuration:

<details>
  <summary>example config</summary>
```yaml
    cctp_type: "synapse"
    # prod contract addresses
    chains:
      - chain_id: 1
        synapse_cctp_address: "0x12715a66773BD9C54534a01aBF01d05F6B4Bd35E"
      - chain_id: 43114
        synapse_cctp_address: "0x12715a66773BD9C54534a01aBF01d05F6B4Bd35E"
      - chain_id: 42161
        synapse_cctp_address: "0x12715a66773BD9C54534a01aBF01d05F6B4Bd35E"
      - chain_id: 10
        synapse_cctp_address: "0x12715a66773BD9C54534a01aBF01d05F6B4Bd35E"
      - chain_id: 8453
        synapse_cctp_address: "0x12715a66773BD9C54534a01aBF01d05F6B4Bd35E"
      - chain_id: 137
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
        43114:
          gas_estimate: 5000000
          max_gas_price: 1000000000000
          supports_eip_1559: true
        10:
          gas_estimate: 5000000
          max_gas_price: 2000000000
          supports_eip_1559: true
        8453:
          gas_estimate: 5000000
        137:
          gas_estimate: 5000000
          max_gas_price: 10000000000000
          supports_eip_1559: true
```
</details>

 - `cctp_type`: The type of CCTP to listen to. Can be either `synapse` or `circle`.
 - `chains`: A list of chain ids and their corresponding CCTP contract addresses. If synapse mode, this should be `synapse_cctp_address` and if circle mode, this should be `token_messenger_address`. Both modes cannot be used at once and the other will be ignored if both are set.
 - `base_omnirpc_url`: The base URL for the OmniRPC service.
 - `unbonded_signer`: The signer to use for transactions. Can be either `AWS`, `File` or `GCP`. The file should be a mounted secret. More details can be found [here](../Services/Signer).
 - `port`: The port to run the relayer on (e.g. 8080)
 - `host`: The host to run the relayer on (e.g. localhost). Note: this should not be publicly exposed
 - `http_backoff_initial_interval_ms`: The initial interval for the backoff in milliseconds.
 - `retry_interval_ms`: The interval to wait between attestation request retries in milliseconds. The [CCTP API Rate Limit](https://developers.circle.com/stablecoins/docs/limits) should be kept in mind.

### Observability

The CCTP relayer implements open telemetry for both tracing and metrics. Please see the [Observability](docs/Services/Observability) page for more info. We'd also highly recommend setting up the [submitter dashboard](../Services/Submitter) as well.
