# Sanguine

[![Go Workflows](https://github.com/synapsecns/sanguine/actions/workflows/go.yml/badge.svg)](https://github.com/synapsecns/sanguine/actions/workflows/go.yml)
[![Foundry Tests](https://github.com/synapsecns/sanguine/actions/workflows/foundry-tests.yml/badge.svg)](https://github.com/synapsecns/sanguine/actions/workflows/foundry-tests.yml)


## Contributing

Read through [CONTRIBUTING.md](./CONTRIBUTING.md) for a general overview of our contribution process.
Then check out our list of [good first issues](https://github.com/ethereum-optimism/optimism/contribute) to find something fun to work on!

## Directory Structure

<pre>
root
├── <a href="./packages">packages</a>
│   ├── <a href="./packages/contracts">contracts</a>: Contracts used for synapse
├── <a href="./tools">tools</a>
│   ├── <a href="./tools/abigen">abigen</a>: Used to generate abigen bindings for go
├── <a href="./core">core</a>: Core service contains the code used by different agents
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
