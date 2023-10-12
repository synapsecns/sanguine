# Guard
The Guard is an agent responsible for verifying actions from other Agents in the optimistic messaging model. This includes polling for invalid states and attestations as well as submitting fraud reports.

## Components
The Guard operates with four main components:
### Run
`streamLogs` is the data-getter for the Guard. It works by instantiating a gRPC connection to Scribe, and puts logs in a channel for the origin and destination contracts on each chain in the config. From here, it verifies the logs' order since the order of logs are very important for merkle tree construction.
<br /> Additionally, if the Guard unexpectedly stops in the middle of streaming logs, it will use the current database state to reconstruct the tree up to where the last log was, then continue to use gRPC.
<br /> <br > `receiveLogs` is the data-processor for the Guard. It works by taking the logs streamed from `streamLogs` and parsing the logs into either a `Message` on the `Origin.sol` contract, or a `Attestation` on the `Destination.sol` contract. It then stores the data into the Guard's database and builds the tree accordingly.
<br /> <br > `loadOriginLatestStates` polls Origin states and caches them in order to make the latest data available.
<br /> <br > `submitLatestSnapshot` fetches the latest snapshot from Origin and submits it on Summit.
<br /> <br > `updateAgentStatuses` polls the database for `RelayableAgentStatus` entries and calls `updateAgentStatus()` once a sufficient agent root is passed to the given remote chain.

### Fraud Reporting
The fraud reporting logic can be found in `fraud.go`, which consists mostly of handlers for various logs. The two major handlers are `handleSnapshotAccepted` and `handleAttestationAccepted`, both of which verify states corresponding to the incoming snapshot/attestation, initiate slashing if applicable, and submit state reports to eligible chains.

## Usage

Navigate to `sanguine/agents/agents/guard/main` and run the following command to start the Guard:

```bash
$ go run main.go
```
Then the Guard command line will be exposed. The Guard requires a gRPC connection to a Scribe instance to stream logs. This can be done with either a remote Scribe or an embedded Scribe.

For more information on how to interact with Scribe see the Executor README.

## Directory Structure

<pre>
Guard
├── <a href="./api">main</a>: API server
├── <a href="./cmd">cmd</a>: CLI commands
├── <a href="./db">db</a>: Database interface
│   └── <a href="db/sql">sql</a>: Database writer, reader, and migrations
├── <a href="./main">main</a>: CLI entrypoint
</pre>
