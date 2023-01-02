# Executor
The Executor is an agent responsible for executing messages on the `Destination.sol` contract. It verifies that the message has been accepted by verifying a merkle proof, as well as the message's optimistic period.

## Components
The Executor operates with two main components:
### Start/Listen
`Start` is the data-getter for the Executor. It works by instantiating a gRPC connection to Scribe, and streams logs for the origin and destination contracts on each chain in the config. From here, it verifies the logs' order since the order of logs are very important for merkle tree construction.
<br /> Additionally, if the Executor unexpectedly stops in the middle of streaming logs, it will use the current database state to reconstruct the tree up to where the last log was, then continue to use gRPC.
<br /> <br > `Listen` is the data-processor for the Executor. It works by taking the logs streamed from `Start` and parsing the logs into either a `Message` on the `Origin.sol` contract, or a `Attestation` on the `Destination.sol` contract. It then stores the data into the Executor's database and builds the tree accordingly.

### Execute
`Execute` works in 3 parts.
1. Verify that the optimistic period for the message has passed. This is done by checking the timestamp that the attestation was accepted on the Destination, then checking the timestamp of the last block.
2. Verify that the message is in the merkle tree.
3. Call the `Execute` function on the `Destination.sol` contract.

## Directory Structure

TODO: fill this out

<pre>
Executor
├── <a href="./config">config</a>
├── <a href="./db">db</a>
│   ├── <a href="./db/datastore/sql">sql</a>
├── <a href="./types">types</a>
├── <a href="./config">config</a>
└── <a href="./consumer">consumer</a>
</pre>
