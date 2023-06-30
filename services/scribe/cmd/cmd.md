# Scribe

Scribe is a service that goes through all chains and contracts specified in a config file and creates a database with logs, receipts, and transactions.

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

## Configuration

chain_id: The ID of the chain
required_confirmations: the number of confirmations required for a block to be finalized
contracts: stores all the contract information for the chain.
get_logs_range: is the number of blocks to request in a single getLogs request.
get_logs_batch_amount: is the number of getLogs requests to include in a single batch request.
store_concurrency: is the number of goroutines to use when storing data.
concurrency_threshold: is the max number of block from head in which concurrent operations (store, getlogs) is allowed.




## Directory Structure

<pre>
scribe
├── <a href="./api">api</a>: Contains the Swagger API definition and GQL Client tests.
├── <a href="./backfill">backfill</a>: Used to fetch logs, receipts, and transactions to store in the database
├── <a href="./cmd">cmd</a>: The command line interface functions for running the Scribe and GraphQL server
├── <a href="./config">config</a>: Configuration files for the Scribe
├── <a href="./db">db</a>: The database schema and functions for interacting with the database
├── <a href="./graphql">graphql</a>: GraphQL implementation for the Scribe's recorded data
│   ├── <a href="./graphql/client">client</a>: The client interface for the GraphQL server
│   ├── <a href="./graphql/contrib">contrib</a>: The GraphQL generators for the Scribe
│   └── <a href="./graphql/server">server</a>: The server implementation for GraphQL
│       └── <a href="./graphql/server/graph">graph</a>: The server's models, resolvers, and schemas
├── <a href="./grpc">grpc</a>: The gRPC client implementation for the Scribe
├── <a href="./internal">internal</a>: Internal packages for the Scribe
├── <a href="./node">node</a>: The new block listener that calls backfill
├── <a href="./scripts">scripts</a>: Scripts for Scribe
└── <a href="./testutil">testutil</a>: Test utilities for the Scribe
</pre>

## Regenerating protobuf definitions:

`make generate`


## Flow

1. Scribe initializes with config
2. A go routine is started for each chain in the config
3. Each go routine starts a concurrent backfill for each contract in the config
4. Each backfill fetches (eth_getLogs) for each contract in chunks. Chunk size is set in the config (GetLogsBatchAmount * GetLogsRange). The fetching flow is blocked by the channel size (15).
5. As the fetching channel is filled, the backfiller starts loads logs into another channel for processing.
6. As logs are taken from the processing channel, scribe does a eth_getTransactionReceipt for each log and a eth_getTransaction for each receipt.
7. Scribe stores the receipt, tx, blocktime for that tx, and all logs from the receipt.
