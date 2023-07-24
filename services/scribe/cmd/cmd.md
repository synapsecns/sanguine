# Scribe
[![Go Reference](https://pkg.go.dev/badge/github.com/synapsecns/sanguine/services/scribe.svg)](https://pkg.go.dev/github.com/synapsecns/sanguine/services/scribe)
[![Go Report Card](https://goreportcard.com/badge/github.com/synapsecns/sanguine/services/scribe)](https://goreportcard.com/report/github.com/synapsecns/sanguine/services/scribe)

Scribe is a multi-chain indexing service. Scribe is designed to take a list of contracts specified by chain id and store logs, receipts, and txs for every event, past to present, in a mysql database.

Use cases
- Analytics for on chain events
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


A full list can be found at graphql/server/graph/schema/queries.graphql


### Scribe Indexer
The scribe indexer supports indexing on any number of contracts on any chain. With a list of contracts Scribe indexes from the
specified start of the contract (from the config)
## Schema
The schema for each table in the database can be found at db/datastore/sql/base


## Usage

Run the following command to start the Scribe:

```bash
$ go run main.go
```
Then the Scribe command line will be exposed. You can use the following commands:

```bash
# Start the Scribe
$ scribe --config </Full/Path/To/Config.yaml> --db <sqlite or mysql> --path <path/to/database or database url>
# Call a single backfill with the scribe
$ backfill --config </Full/Path/To/Config.yaml> --db <sqlite or mysql> --path <path/to/database or database url>
# Start the Scribe server
$ server --port <port> --db <sqlite or mysql> --path <path/to/database or database url>
```

### Deploy
See /charts/scribe for the deployment helm chart for this service

### Configuration
````
chain_id: The ID of the chain
required_confirmations: the number of confirmations required for a block to be finalized
contracts: stores all the contract information for the chain.
get_logs_range: is the number of blocks to request in a single getLogs request.
get_logs_batch_amount: is the number of getLogs requests to include in a single batch request.
store_concurrency: is the number of goroutines to use when storing data.
concurrency_threshold: is the max number of block from head in which concurrent operations (store, getlogs) is allowed.
````




## Understanding Scribe Indexer
The scribe indexer is composed of three components
1. `Fetcher`: Takes a list of contracts, a block range, and fetches and feeds logs into a channel to be consumed by an indexer.
2. `Indexer`: Takes a list of contracts, a block range, and a config and stores logs, receipts, and txs for all events in that range.
3. `ChainIndexer`: Runs 2+ indexers per chain.
   1. 1 indexer for each contract behind the specified livefill threshold (backfill)
   2. 1 indexer for every contract within the specified livefill threshold (livefill)
   3. 1 indexer for every contract within the specified "unconfirmed" range at the chain tip. (unconfirmed livefill)


### Flow
1. Scribe initializes with config
2. Each chain has a `ChainIndexer` spun up and runs it in a go routine
3. The `ChainIndexer` checks all contracts for their `lastIndexed` block. Contracts without a `lastIndexed` or with a `lastIndexed` block outside the
specified livefill block range are put into individual indexers (backfill). All other contracts are collected into a single indexer (livefill).
4. A contract in an individual indexer (backfill) reaches the livefill threshold, it is passed into a channel where it will be picked up by the go routine running the
indexer for the livefill contracts.
5. While contracts are being livefilled, there is another indexer with all inputted chains. This indexer is used to livefill the unconfirmed range at the chain tip. This range is set by the config
and stores data in seperate tables than the other indexers. This table has stale rows (old rows) deleted every few hours (set in config).



### Directory Structure

<pre>
scribe
├── <a href="./api">api</a>: Contains the Swagger API definition and GQL Client tests.
├── <a href="./backend">backend</a>: The backend implementation for the Scribe
├── <a href="./client">client</a>: Client implementation for Scribe (embedded/remote)
├── <a href="./cmd">cmd</a>: The command line interface functions for running the Scribe and GraphQL server
├── <a href="./config">config</a>: Configuration files for Scribe
├── <a href="./db">db</a>: The database schema and functions for interacting with the database
├── <a href="./graphql">graphql</a>: GraphQL implementation for the Scribe's recorded data
│   ├── <a href="./graphql/client">client</a>: The client interface for the GraphQL server
│   ├── <a href="./graphql/contrib">contrib</a>: The GraphQL generators for the Scribe
│   └── <a href="./graphql/server">server</a>: The server implementation for GraphQL
│       └── <a href="./graphql/server/graph">graph</a>: The server's models, resolvers, and schemas
├── <a href="./grpc">grpc</a>: The gRPC client implementation for the Scribe
├── <a href="./internal">internal</a>: Internal packages for the Scribe
├── <a href="./logger">logger</a>: Handles logging for various events in Scribe.
├── <a href="./metadata">metadata</a>: Provides metadata for building .
├── <a href="./scripts">scripts</a>: Scripts for Scribe
└── <a href="./testutil">testutil</a>: Test utilities for the Scribe
</pre>

## Regenerating protobuf definitions:

`make generate`


## Flow

