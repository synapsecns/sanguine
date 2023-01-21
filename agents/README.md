# Agents

The Agents are the off-chain components of Sanguine. They are responsible for signing, reporting, verifying, and executing messages across chains.

## Use

From `sanguine/agents`, run `go run main.go` to expose the CLI for all agents. Visit each agents' respective directory's README.md for information on individual agent usage.

## Directory Structure

<pre>
root
├── <a href="./agents">agents</a>: Contain all the agents used in optimistic messaging
│   ├── <a href="./agents/agentsintegration">agentsintegration</a>: Testing all the agents working together
│   ├── <a href="./agents/executor">executor</a>: Responsible for verifying and executing cross-chain messages
│   ├── <a href="./agents/guard">guard</a>: Signs attestations and posts to attestation collector and destination
│   └── <a href="./agents/notary">notary</a>: Scans origin chains for messages and signs attestations then posts to AttestationCollector
├── <a href="./cmd">cmd</a>: The command line entrypoint for all agents
├── <a href="./config">config</a>: Configs for agents
├── <a href="./contracts">contracts</a>: Go interfaces for messaging contracts
├── <a href="./db">db</a>: Agents datastore
├── <a href="./domains">domains</a>: Adapters for each domain
├── <a href="./indexer">indexer</a>: Periodically reads from the db and stores data in the db
├── <a href="./internal">internal</a>: Dev dependencies
├── <a href="./testutil">testutil</a>: Contains mock deployers for interacting with the mock backend
└── <a href="./types">types</a>: Common agents types
</pre>
