# Executor
The Executor is an agent responsible for executing messages on the `Destination.sol` contract. It verifies that the message has been accepted by verifying a merkle proof, as well as the message's optimistic period.

## Components
The Executor operates with four main components:
### Run
`streamLogs` is the data-getter for the Executor. It works by instantiating a gRPC connection to Scribe, and puts logs in a channel for the origin and destination contracts on each chain in the config. From here, it verifies the logs' order since the order of logs are very important for merkle tree construction.
<br /> Additionally, if the Executor unexpectedly stops in the middle of streaming logs, it will use the current database state to reconstruct the tree up to where the last log was, then continue to use gRPC.
<br /> <br > `receiveLogs` is the data-processor for the Executor. It works by taking the logs streamed from `streamLogs` and parsing the logs into either a `Message` on the `Origin.sol` contract, or a `Attestation` on the `Destination.sol` contract. It then stores the data into the Executor's database and builds the tree accordingly.
<br /> <br > `setMinimumTime` sets the minimum unix time for a message before it has passed the optimistic period and can be executed. It works by going through all messages in the database that haven't been set, and checks for attestations to set each message's time.
<br /> <br > `executeExecutable` is the function that executes messages. It looks through all messages that have a time set from `setMinimumTime` and executes them if the minimum time is past the current time.

### Execute
`Execute` works in 3 parts.
1. Verify that the optimistic period for the message has passed. This is done by checking the timestamp that the attestation was accepted on the Destination, then checking the timestamp of the last block.
2. Verify that the message is in the merkle tree.
3. Call the `Execute` function on the `Destination.sol` contract.

## Usage

Navigate to `sanguine/agents/agents/executor/main` and run the following command to start the Executor:

```bash
$ go run main.go
```
Then the Executor command line will be exposed. The Executor requires a gRPC connection to a Scribe instance to stream logs. This can be done with either a remote Scribe or an embedded Scribe.

From the CLI, you specify whether you want to use a `remote` or `embedded` scribe via the `--scribe-type` flag. If using `remote`, you need to use the `--scribe-port`, `--scribe-grpc-port` and `--scribe-url` flags. If using `embedded`, you need to use the `--scribe-db` and `--scribe-path` flags.

### Remote Scribe

Run the following command to start the Executor with a remote Scribe:

```bash
# Start the Executor
$ run --config </Users/synapsecns/config.yaml> --db <sqlite> or <mysql>\
 --path <path/to/database> or <database url> --scribe-type remote\
 --scribe-port <port> --scribe-grpc-port <port> --scribe-url <url>
```

### Embedded Scribe

When using an embedded Scribe, a Scribe config must be included in the Executor config.

Run the following command to start the Executor with an embedded Scribe:

```bash
# Start the Executor and an embedded Scribe
$ run --config </Users/synapsecns/config.yaml> --db <sqlite> or <mysql>\
 --path <path/to/database> or <database url> --scribe-type embedded\
 --scribe-db <sqlite> or <mysql> --scribe-path <path/to/database> or <database url>
```



## Directory Structure

<pre>
Executor
├── <a href="./cmd">cmd</a>: CLI commands
├── <a href="./config">config</a>: Configuration files
├── <a href="./db">db</a>: Database interface
│   └── <a href="./db/datastore/sql">sql</a>: Database writer, reader, and migrations
├── <a href="./main">main</a>: CLI entrypoint
└── <a href="./types">types</a>: Executor types
</pre>
