# Executor

The Executor is an agent that verifies messages and executes them on the destination chain.

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
