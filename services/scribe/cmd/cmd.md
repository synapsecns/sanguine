# Scribe
[![Go Reference](https://pkg.go.dev/badge/github.com/synapsecns/sanguine/services/scribe.svg)](https://pkg.go.dev/github.com/synapsecns/sanguine/services/scribe)
[![Go Report Card](https://goreportcard.com/badge/github.com/synapsecns/sanguine/services/scribe)](https://goreportcard.com/report/github.com/synapsecns/sanguine/services/scribe)

Scribe is a multi-chain indexing service. Scribe is designed to take a list of contracts specified by chain id and store logs, receipts, and txs for every event, past to present, in a mysql database.

Use cases
- Analytics for on-chain events
- Stream events occurring across chains
- Monitor activity on your contracts


Scribe is currently indexing over 900 contracts across 18 different chains.



## Core Features
### Scribe Server
Scribe comes with a graphql/iql endpoint with various queries to make interacting with the collected data easier. Here are some basic queries
- `lastIndexed(chain_id, contract_address)`
- `logsRange(chain_id, contract_address, start_block, end_block, page)`
- `blockTime(chain_id, block_number)`
- `txSender(tx_hash, chain_id)`


A full list can be found at <a href="./graphql/server/graph/schema/queries.graphql">graphql/server/graph/schema/queries.graphql</a>


### Scribe Indexer
Scribe indexer supports indexing on any number of contracts on any chain. For each contract Scribe indexes from the
specified start of the contract from the config. Scribe stores every log, receipt, tx, and timestamp for every event until
it reaches the specified livefill threshold where it will then continue to index events as they occur in real time.
For reorg protection, Scribe does not store any events that are more than the specified number of confirmations from the chain tip
into its primary tables. Scribe stores "unconfirmed" events near the tip of the chain in separate and transient tables.
The Scribe server has built in queries to query both the primary (confirmed) and transient (unconfirmed) tables at the same time if
realtime data needed for querying or streaming.


## Usage

Run the following command to start Scribe:

```bash
$ go run main.go
```
Then Scribe command line will be exposed. You can use the following commands:

```bash
# Start Scribe indexer
$ Scribe --config </Full/Path/To/Config.yaml> --db <sqlite or mysql> --path <path/to/database or database url>
# Start Scribe server
$ server --port <port> --db <sqlite or mysql> --path <path/to/database or database url>
```

### Deploy
See <a href="../../charts/scribe">/charts/scribe</a> for the deployment helm chart for this service

### Configuration
There are many ways to augment Scribe's behavior. Before running Scribe, please take a look at the different parameters.
```
rpc_url: The url of the rpc aggregator
chains: A list of chains to index
  chain_id: The ID of the chain
  get_logs_range: is the number of blocks to request in a single getLogs request.
  get_logs_batch_amount: is the number of getLogs requests to include in a batch request.
  store_concurrency: is the number of goroutines to use when storing data.
  concurrency_threshold: is the max number of block from head in which concurrent operations (store, getlogs) is allowed.
  livefill_threshold: number of blocks away from the head to start livefilling (see "Understanding Scribe Indexer" to understand this better)
  livefill_range: range in whcih the getLogs request for the livefill contracts will be requesting.
  livefill_flush_interval: the interval in which the unconfirmed livefill table will be flushed.
  confirmations: the number of blocks from head that the livefiller will livefill up to (and where the unconfirmed livefill indexer will begin)
  contracts: stores all the contract information for the chain
    address: address of the contract
    start_block: block to start indexing the contract from (block with the first tx)
```


#### Example Config
```yaml
rpc_url: <omnirpc>
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



## Understanding the Scribe Indexer
The Scribe indexer is composed of three components
1. `Fetcher`: Takes a list of contracts, a block range, and fetches and feeds logs into a channel to be consumed by an indexer.
2. `Indexer`: Takes a list of contracts, a block range, and a config and stores logs, receipts, and txs for all events in that range.
3. `ChainIndexer`: Runs 2+ indexers per chain.
   1. 1 indexer for each contract behind the specified livefill threshold (backfill)
   2. 1 indexer for every contract within the specified livefill threshold (livefill)
   3. 1 indexer for every contract within the specified "unconfirmed" range at the chain tip. (unconfirmed livefill)


### Chain level flow
1. Scribe initializes with config
2. Each chain has a `ChainIndexer` spun up and runs it in a go routine
3. The `ChainIndexer` checks all contracts for their `lastIndexed` block. Contracts without a `lastIndexed` or with a `lastIndexed` block outside the
specified livefill block range are put into individual indexers (backfill). All other contracts are collected into a single indexer (livefill).
4. A contract in an individual indexer (backfill) reaches the livefill threshold, it is passed into a channel where it will be picked up by the go routine running the
indexer for the livefill contracts.
5. While contracts are being livefilled, there is another indexer with all contracts listed on the given chain. This indexer is used to livefill the unconfirmed range at the chain tip. This range is set by the config
and stores data in separate tables than the other indexers. This table has stale rows (old rows) deleted every few hours (set in config).


### Indexer level flow
1. `Indexer` is initialized with config and a list of contracts
2. `Indexer` is started with a block range
3. The indexer creates a `Fetcher` that feeds logs into a channel that is then consumed by the indexer. The fetcher retrieves logs in chunks specified in the config
(`get_logs_range` and `get_logs_batch_amount`).
4. When a new log is retrieved the indexer gets the tx, block header (timestamp), and receipt (and all of its logs) for that log and the stores them in the database. Depending on the concurrency
settings, Scribe will spin up multiple concurrent processes for retrieving this data and storing. It is recommended that concurrency (`concurrency_threshold`) is set to at least a couple hours
from the head to ensure that data is inserted into the database in order (if streaming).
5. The indexer will continue to fetch and store data until it reaches the end of the block range.


### Directory Structure

<pre>
Scribe
├── <a href="./api">api</a>: Contains the Swagger API definition and GQL Client tests.
├── <a href="./backend">backend</a>: The backend implementation for Scribe
├── <a href="./client">client</a>: Client implementation for Scribe (embedded/remote)
├── <a href="./cmd">cmd</a>: The command line interface functions for running Scribe and GraphQL server
├── <a href="./config">config</a>: Configuration files for Scribe
├── <a href="./db">db</a>: The database schema and functions for interacting with the database
├── <a href="./graphql">graphql</a>: GraphQL implementation for Scribe's recorded data
│   ├── <a href="./graphql/client">client</a>: The client interface for the GraphQL server
│   ├── <a href="./graphql/contrib">contrib</a>: The GraphQL generators for Scribe
│   └── <a href="./graphql/server">server</a>: The server implementation for GraphQL
│       └── <a href="./graphql/server/graph">graph</a>: The server's models, resolvers, and schemas
├── <a href="./grpc">grpc</a>: The gRPC client implementation for Scribe
├── <a href="./internal">internal</a>: Internal packages for Scribe
├── <a href="./logger">logger</a>: Handles logging for various events in Scribe.
├── <a href="./metadata">metadata</a>: Provides metadata for building.
├── <a href="./scripts">scripts</a>: Scripts for Scribe
├── <a href="./service">service</a>: Service holds Scribe indexer code (Fetcher, Indexer, ChainIndexer)
├── <a href="./testhelper">testhelper</a>: Assists testing in downstream services.
├── <a href="./testutil">testutil</a>: Test utilities suite for Scribe
└── <a href="./types">types</a>: Holds various custom types for Scribe
</pre>



### Schema
The schema for each table in the database can be found at <a href="./db/datastore/sql/base">db/datastore/sql/base</a>

## Regenerating protobuf definitions:

`make generate`



