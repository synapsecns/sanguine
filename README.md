# Sanguine

[![Go Workflows](https://github.com/synapsecns/sanguine/actions/workflows/go.yml/badge.svg)](https://github.com/synapsecns/sanguine/actions/workflows/go.yml)
[![Foundry Tests](https://github.com/synapsecns/sanguine/actions/workflows/foundry-tests.yml/badge.svg)](https://github.com/synapsecns/sanguine/actions/workflows/foundry-tests.yml)

## Contributing

Read through [CONTRIBUTING.md](./CONTRIBUTING.md) for a general overview of our contribution process.
Then check out our list of [good first issues](https://github.com/ethereum-optimism/optimism/contribute) to find something fun to work on!

## Directory Structure

<pre>
root
├── <a href="./core">core</a>: The Go core library with common utilities for use across the monorepo
├── <a href="./ethergo">ethergo</a>: Go-based ethereum testing + common library
├── <a href="./packages">packages</a>
│   ├── <a href="./packages/contracts-core">contracts-core</a>: Core contracts used for synapse, powered by <a href="https://github.com/foundry-rs/foundry">Foundry</a>
├── <a href="./tools">services</a>
│   ├── <a href="./services/omnirpc">omnirpc</a>: Latency aware RPC Client used across multiple-chains at once
│   ├── <a href="./services/scribe">scribe</a>: Generalized ethereum event logger
├── <a href="./tools">tools</a>
│   ├── <a href="./tools/abigen">abigen</a>: Used to generate abigen bindings for go
│   ├── <a href="./tools/modulecopier">module copier</a>: Used to copy internal modules and export methods for testing
├── <a href="./agents">agents</a>: agents contain all the agents used in optimistic messaging
</pre>


## Setup

Clone the repository, open it, and install nodejs packages with `yarn`:

```bash
git clone https://github.com/synapsecns/sanguine
cd sanguine
yarn install
```


### Install the Correct Version of NodeJS

Using `nvm`, install the correct version of NodeJS.

```
nvm use
```

### Building the TypeScript packages

To build all of the [TypeScript packages](./packages), run:

```bash
yarn clean
yarn build
```

Packages compiled when on one branch may not be compatible with packages on a different branch.
**You should recompile all packages whenever you move from one branch to another.**
Use the above commands to recompile the packages.
