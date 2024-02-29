<br/>
<p align="center">
<a href="https://interchain.synapseprotocol.com/" target="_blank">
<img src="https://raw.githubusercontent.com/synapsecns/sanguine/feat/readme-updates/assets/interchain-logo.svg" width="225" alt="Synapse Interchain logo">
</a>
</p>
<br/>

[![Go Reference](https://pkg.go.dev/badge/github.com/synapsecns/sanguine/agents.svg)](https://pkg.go.dev/github.com/synapsecns/sanguine/agents)
[![Go Report Card](https://goreportcard.com/badge/github.com/synapsecns/sanguine/agents)](https://goreportcard.com/report/github.com/synapsecns/sanguine/agents)
[![Image](https://ghcr-badge.egpl.dev/synapsecns/sanguine%2Fagents/size?color=%2344cc11&tag=latest&label=image+size&trim=)](https://github.com/synapsecns/sanguine/pkgs/container/sanguine%2Fagents)

# Agents


The Agents are the off-chain components of Synapse Interchain Network. They are responsible for signing, reporting, verifying, and executing messages across chains.

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

# Testing Suite

Tests for `agents` have setup hooks defined in `agents/testutil/simulated_backends_suite.go`. Any suite that embeds `SimulatedBackendsTestSuite` will have simulated backend and messaging contract scaffolding for Summit, Origin, and Desination chains. This includes `TestExecutorSuite`, `TestGuardSuite`, `TestNotarySuite`, `ExampleAgentSuite`, and `AgentsIntegrationSuite`.

To run all agent tests:

```bash
cd agents
go test -v ./...
```

To run an individual suite (for example, `TestExecutorSuite`):

```bash
cd agents/executor
go test -v
```

To run an individual test (for example, `TestVerifyState`):
```bash
cd agents/executor
go test -v -run TestExecutorSuite/TestVerifyState
```
