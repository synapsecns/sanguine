# Scribe

The scribe is a rpc client that can assert the same data is coming from a range of rpcs.

`./scribe backfill --config x`: Backfills to the current block and then stops

## Directory Structure

<pre>
scribe
├── <a href="./backfill">backfill</a>: Used to fetch logs, receipts, and transactions to store in the database
├── <a href="./cmd">cmd</a>: The command line interface functions for running the Scribe and GraphQL server
├── <a href="./config">config</a>: Configuration files for the Scribe
├── <a href="./db">db</a>: The database schema and functions for interacting with the database
├── <a href="./graphql">graphql</a>: GraphQL implementation for the Scribe's recorded data
│   ├── <a href="./graphql/client">client</a>: The client interface for the GraphQL server
│   └── <a href="./graphql/server">server</a>: The server implementation for GraphQL
│       └── <a href="./graphql/server/graph">graph</a>: The server's models, resolvers, and schemas
├── <a href="./internal">internal</a>: Internal packages for the Scribe
└── <a href="./node">node</a>: The new block listener that calls backfill
</pre>
