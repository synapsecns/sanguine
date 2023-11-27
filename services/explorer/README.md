# Explorer

[![Go Reference](https://pkg.go.dev/badge/github.com/synapsecns/sanguine/services/explorer.svg)](https://pkg.go.dev/github.com/synapsecns/sanguine/services/explorer)
[![Go Report Card](https://goreportcard.com/badge/github.com/synapsecns/sanguine/services/explorer)](https://goreportcard.com/report/github.com/synapsecns/sanguine/services/explorer)


<!-- TODO: explain the two parts: indexer and server, etc-->

See [#167](https://github.com/synapsecns/sanguine/issues/167) to learn more.

To access the clickhouse database, you can use the following command from the docker image:

```bash
$ clickhouse-client --database=clickhouse_test --user=clickhouse_test --password=clickhouse_test
```

## Directory Structure.

<pre>
explorer
├── <a href="./api">api</a>: API server
├── <a href="./backfill">backfill</a>: Chain level backfilling service to populate the database
├── <a href="./cmd">cmd</a>: CLI commands
├── <a href="./config">config</a>: Configuration files
├── <a href="./consumer">consumer</a>: Consumes data from Scribe and populates the Explorer database
│   ├── <a href="consumer/fetchers/scribe/client">client</a>: Client for the Scribe consumer
│   ├── <a href="./consumer/fetcher">fetcher</a>: Fetches data from Scribe, BridgeConfig contract, and Swap contract
│   └── <a href="./consumer/parser">parser</a>: Parses and stores events
├── <a href="./contracts">contracts</a>: Smart contracts and their generated interfaces/utils
│   ├── <a href="./contracts/bridge">bridge</a>: Bridge smart contract applications
│   ├── <a href="./contracts/bridgeconfig">bridgeconfig</a>: BridgeConfig smart contract applications
│   ├── <a href="./contracts/contracts">contracts</a>: Raw flattened smart contracts and test contracts
│   └── <a href="./contracts/swap">swap</a>: Swap smart contract applications
├── <a href="./db">db</a>: Database interface
│   └── <a href="./db/sql">sql</a>: Database writer, reader, and migrations
├── <a href="./graphql">graphql</a>: GraphQL implementation for the Explorer's recorded data
│   ├── <a href="./graphql/client">client</a>: The client interface for the GraphQL server
│   ├── <a href="./graphql/contrib">contrib</a>: Generator for the GraphQL schema
│   └── <a href="./graphql/server">server</a>: The server implementation for GraphQL
│       └── <a href="./graphql/server/graph">graph</a>: The server's models, resolvers, and schemas
├── <a href="./node">node</a>: Live Explorer node
├── <a href="./testutil">testutil</a>: Test utilities
└── <a href="./types">types</a>: Explorer specific types
</pre>
