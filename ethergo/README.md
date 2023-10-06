# Ethergo

[![Go Reference](https://pkg.go.dev/badge/github.com/synapsecns/sanguine/ethergo.svg)](https://pkg.go.dev/github.com/synapsecns/sanguine/ethergo)
[![Go Report Card](https://goreportcard.com/badge/github.com/synapsecns/sanguine/ethergo)](https://goreportcard.com/report/github.com/synapsecns/sanguine/ethergo)


## Overview

Ethergo is a comprehensive embedded test suite for [Synapse](https://synapseprotocol.com), built using the Go programming language. It provides the functionality to deploy contracts to various backends, streamlining the testing process for different types of blockchain networks. The module is organized into several subdirectories, each dedicated to specific aspects of the testing process, such as backends, chain interaction, contract management, and signing utilities.

Some key components of Ethergo include:

 - **Backends**: Ethergo supports different test backends for testing against various types of chains, including Anvil, Ganache, Geth, Preset, and Simulated backends.
 - **Chain**: This subdirectory provides tools for chain interaction, such as a chain client, a chain watcher, a deterministic gas estimator, and a client interface for the chain watcher.
 - **Contracts**: This section contains interfaces for working with contracts, including deployer and manager functionalities.
 - **Manager**: The manager is responsible for handling contract deployments.
 -  **Signer**: Ethergo offers signing and transaction submission utilities, with various adapters for signing transactions and handling race conditions with nonces. It also includes a wallet interface capable of importing keys from seed phrases, private keys, and mnemonics.

 To see a working example of Ethergo in action, refer to the [`example`](./example) folder, which provides a complete demonstration of how to use the deployer and manager components. Please note that the APIs are subject to change.


## Directory Structure

<pre>
root
├── <a href="./backends">backends</a>: Contains different test backends for testing against different types of chains
│   ├── <a href="./backends/anvil">anvil</a>: Contains the anvil backend. This is useful for fork testing and requires docker. It aloso contains cheatcodes for testing detailed [here](https://book.getfoundry.sh/anvil/)
│   ├── <a href="./backends/ganache">ganache</a>: Contains the ganache backend. This is useful for fork testing and requires docker. This is currently using an older version of ganache, but will be updated to use the latest version. See [here](https://github.com/trufflesuite/ganache) for details.
│   ├── <a href="./backends/geth">geth</a>: Contains an embedded geth backend. This is useful for testing against a local geth instance without forking capabilities. This does not require docker and runs fully embedded in the go application, as such it is faster than the docker-based backends, but less versatile. Used when an rpc address is needed for a localnet.
│   ├── <a href="./backends/preset">preset</a>: Contains a number of preset backends for testing.
│   ├── <a href="./backends/simulated">simulated</a>: The fastest backend, this does not expose an rpc endpoint and uses geth's [simulated backend](https://goethereumbook.org/en/client-simulated/)
├── <a href="./chain">chain</a>: Contains a client for interacting with the chain. This will be removed in a future version. Please use [client](./client) going forward.
│   ├── <a href="./chain/chainwatcher">chainwatcher</a>: Watches the chain for events, blocks and logs
│   ├── <a href="./chain/client">client</a>: Contains eth clients w/ rate limiting, workarounds for bugs in some chains, etc.
│   ├── <a href="./chain/gas">gas</a>: Contains a deterministic gas estimator
│   ├── <a href="./chain/watcher">watcher</a>: Client interface for chain watcher.
├── <a href="./contracts">contracts</a>: Contains interfaces for using contracts with the deployer + manager
├── <a href="./client">client</a>: Contains an open tracing compatible ethclient with batching.
├── <a href="./example">example</a>: Contains a full featured example of how to use deployer + manager
├── <a href="./forker">forker</a>: Allows the use of fork tests in live chains without docker using an anvil binary.
├── <a href="./manager">manager</a>: Manages contract deployments.
├── <a href="./mocks">mocks</a>: Contains mocks for testing various data types (transactions, addresses, logs, etc)
├── <a href="./parser">parser</a>: Parse hardhat deployments
│   ├── <a href="./parser/hardhat">hardhat</a>: Parses hardhat deployments
│   ├── <a href="./parser/abiutil">abi</a>: Parses abi function input
│   ├── <a href="./parser/rpc">rpc</a>: Parses rpc requests/responses
├── <a href="./signer">signer</a>: Signing + tx submission utilities
│   ├── <a href="./signer/nonce">nonce</a>: Automatically handles race conditions with nonces, used in backends.
│   ├── <a href="./signer/signer">signer</a>: Various adapters for signing transactions
│     ├── <a href="./signer/signer/awssigner">awssigner</a>: Use AWS KMS to sign transactions
│     ├── <a href="./signer/signer/gcpsigner">gcpsigner</a>: Use GCP Cloud KMS to sign transactions
│     ├── <a href="./signer/signer/localsigner">localsigner</a>: Use a local signer to sign transactions
│   ├── <a href="./signer/wallet">wallet</a>: Wallet contains a wallet interface. It is capable of importing keys from seed phrases, private keys, and mnemonics.
├── <a href="./submitter">submitter</a>: Generic tx submitter, currently not stable enough for general use.
├── <a href="./util">util</a>: various utilities used throughout the library
</pre>

