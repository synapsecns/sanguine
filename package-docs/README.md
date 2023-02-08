sanguine / [Exports](modules.md)

# Sanguine

[![Go Workflows](https://github.com/synapsecns/sanguine/actions/workflows/go.yml/badge.svg)](https://github.com/synapsecns/sanguine/actions/workflows/go.yml)
[![Foundry Tests](https://github.com/synapsecns/sanguine/actions/workflows/foundry-tests.yml/badge.svg)](https://github.com/synapsecns/sanguine/actions/workflows/foundry-tests.yml)

## Contributing

Read through [CONTRIBUTING.md](./CONTRIBUTING.md) for a general overview of our contribution process.
Then check out our list of [good first issues](https://github.com/ethereum-optimism/optimism/contribute) to find something fun to work on!

## Directory Structure

<pre>
root
├── <a href="./agents">agents</a>: agents contain all the agents used in optimistic messaging
├── <a href="./charts">charts</a>: The helm charts used for deploying sanguine related services
├── <a href="./contrib">contrib</a>: Devops related tools
│   ├── <a href="./contrib/git-changes-action">git-changes-action</a>: Github action for identifying changes in dependent modules in a go workspace
│   ├── <a href="./contrib/release-copier-action">release-copier-action</a>: Github action for copying releases from one repo to another
│   ├── <a href="./contrib/terraform-provider-iap">terraform-provider-iap</a>: Terraform provider used for bastion proxy tunneling
│   ├── <a href="./contrib/terraform-provider-helmproxy">terraform-provider-helmproxy</a>: Terraform provider that allows helm to be proxied through an iap bastion proxy
│   ├── <a href="./contrib/terraform-provider-kubeproxy">terraform-provider-kubeproxy</a>: Terraform provider that allows kube to be proxied through an iap bastion proxy
│   ├── <a href="./contrib/tfcore">tfcore</a>: Terraform core utilities + iap utilities
├── <a href="./core">core</a>: The Go core library with common utilities for use across the monorepo
├── <a href="./ethergo">ethergo</a>: Go-based ethereum testing + common library
├── <a href="./packages">packages</a>
│   ├── <a href="./packages/contracts-core">contracts-core</a>: Core contracts used for synapse, powered by <a href="https://github.com/foundry-rs/foundry">Foundry</a>
│   ├── <a href="./packages/explorer-ui">explorer-ui</a>: Explorer UI
├── <a href="./tools">services</a>
│   ├── <a href="./services/explorer">explorer</a>: Bridge/messaging explorer backend
│   ├── <a href="./services/omnirpc">omnirpc</a>: Latency aware RPC Client used across multiple-chains at once
│   ├── <a href="./services/scribe">scribe</a>: Generalized ethereum event logger
├── <a href="./tools">tools</a>
│   ├── <a href="./tools/abigen">abigen</a>: Used to generate abigen bindings for go
│   ├── <a href="./tools/bundle">bundle</a>: Modified version of <a href="https://pkg.go.dev/golang.org/x/tools@v0.5.0/cmd/bundle"> go bundler </a> with improved shadowing support
│   ├── <a href="./tools/modulecopier">module copier</a>: Used to copy internal modules and export methods for testing
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
