# Scribe

Scribe is a multi-chain indexing service designed to store logs, receipts, and transactions for every event from specified contracts across multiple blockchains. It provides a powerful tool for analytics, event streaming, and monitoring on-chain activities.

## Key Features

1. **Multi-Chain Support**: Currently indexes over 900 contracts across 18 different chains.
2. **Flexible Indexing**: Supports indexing any number of contracts on any chain.
3. **Real-Time and Historical Data**: Indexes from a specified start block up to real-time events.
4. **Reorg Protection**: Stores "unconfirmed" events near the chain tip in separate tables.
5. **GraphQL API**: Provides a GraphQL/IQL endpoint for easy data querying.

## Architecture

Scribe consists of two main components:

1. **Scribe Indexer**: Responsible for fetching and storing blockchain data.
2. **Scribe Server**: Provides a GraphQL API for querying indexed data.

The indexer architecture consists of three main components:

1. **Fetcher**: Retrieves logs for specified contracts and block ranges.
2. **Indexer**: Stores logs, receipts, and transactions for events for a given contract.
3. **ChainIndexer**: Manages multiple indexers per chain for different indexing stages (backfill, livefill, unconfirmed livefill).

## Running Scribe

### Building From Source

To build Scribe from source, you will need to have Go installed. You can install Go by following the instructions [here](https://golang.org/doc/install). Once you have Go installed, you can build Scribe by running the following commands:

1. `git clone https://github.com/synapsecns/sanguine --recursive`
2. `cd sanguine/services/scribe`
3. `go build`


There is also a [docker image](https://github.com/synapsecns/sanguine/pkgs/container/sanguine%2Fscribe) available for scribe.

### Running the Components

To start the Scribe indexer:

```bash
./scribe --config </Full/Path/To/Config.yaml> --db <sqlite or mysql> --path <path/to/database or database url>
```

To start the Scribe server:

```bash
./server --port <port> --db <sqlite or mysql> --path <path/to/database or database url>
```

### Deployment

Scribe can be deployed using Helm charts. You can see the scribe chart on [artifacthub](https://artifacthub.io/packages/helm/synapse/scribe)

### Configuration

Most deployments of scribe will run two components, the indexer and the server. The indexer is responsible for fetching and storing blockchain data, while the server provides a GraphQL API for querying indexed data. The configuration file for Scribe is a YAML file with the following structure:

indexer.yaml:
<details>
  <summary> example config</summary>
  ```yaml

    ```
</details>

```yaml
chains:
  - chain_id: 1
    get_logs_range: 500
    get_block_batch_amount: 1
    store_concurrency: 100
    concurrency_threshold: 50000
    livefill_threshold: 300
    livefill_range: 200
    livefill_flush_interval: 10000
    confirmations: 200
    contracts:
      - address: 0xAf41a65F786339e7911F4acDAD6BD49426F2Dc6b
        start_block: 18646320
```

Key configuration parameters include:

- `rpc_url`: The omnirpc url to use for querying chain data (no trailing slash). For more information on omnirpc, see [here](Omnirpc).
- `chains`: List of chains to index, including chain-specific parameters:
  - `chain_id`: The ID of the chain.
  - `get_logs_range`: The number of blocks to request in a single `getLogs` request.
  - `get_logs_batch_amount`: The number of `getLogs` requests to include in a batch request.
  - `store_concurrency`: The number of goroutines to use when storing data.
  - `concurrency_threshold`: The maximum number of blocks from the head in which concurrent operations (store, getLogs) are allowed.
  - `livefill_threshold`: Number of blocks away from the head to start livefilling.
  - `livefill_range`: Range in which the `getLogs` request for the livefill contracts will be requesting.
  - `livefill_flush_interval`: The interval in which the unconfirmed livefill table will be flushed.
  - `confirmations`: The number of blocks from the head that the livefiller will livefill up to (and where the unconfirmed livefill indexer will begin).
  - `contracts`: List of contracts to index per chain, including:
    - `address`: Address of the contract.
    - `start_block`: Block to start indexing the contract from (block with the first transaction).

## API

Scribe provides a GraphQL API for querying indexed data. Some basic queries include:

- `lastIndexed(chain_id, contract_address)`
- `logsRange(chain_id, contract_address, start_block, end_block, page)`
- `blockTime(chain_id, block_number)`
- `txSender(tx_hash, chain_id)`

For a full list of available queries, refer to the GraphQL schema.

## Observability

Scribe implements open telemetry for both tracing and metrics. Please see the [Observability](Observability) page for more info.
