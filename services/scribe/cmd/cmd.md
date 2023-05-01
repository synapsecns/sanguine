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
