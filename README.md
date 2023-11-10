<br/>
<p align="center">
<a href="https://synapseprotoocol.com" target="_blank">
<img src="https://raw.githubusercontent.com/synapsecns/sanguine/master/assets/logo.svg" width="225" alt="Synapse logo">
</a>
</p>
<br/>

[![GitHub](https://img.shields.io/github/license/synapsecns/sanguine)](https://github.com/synapsecns/sanguine/blob/master/LICENSE)
[![GitHub commit activity (branch)](https://img.shields.io/github/commit-activity/y/synapsecns/sanguine?style=flat-square)](https://github.com/synapsecns/sanguine/commits/)
![GitHub pull requests](https://img.shields.io/github/issues-pr/synapsecns/sanguine)
[![Codecov](https://img.shields.io/codecov/c/github/synapsecns/sanguine?style=flat-square&logo=codecov&link=https%3A%2F%2Fapp.codecov.io%2Fgh%2Fsynapsecns%2Fsanguine)](https://app.codecov.io/gh/synapsecns/sanguine)
[![Code Climate maintainability](https://img.shields.io/codeclimate/maintainability-percentage/synapsecns/sanguine)](https://codeclimate.com/github/synapsecns/sanguine/progress/maintainability)


[![Gitbook](https://img.shields.io/badge/docs-grey?logo=gitbook&style=flat-square)](https://docs.synapseprotocol.com/)
[![Discord](https://img.shields.io/discord/887411327696511027?style=flat-square&logo=discord)](https://discord.com/invite/3AHETJmrXW)
[![Telegram](https://img.shields.io/badge/chat-telegram-blue?logo=telegram&style=flat-square)](https://t.me/synapseprotocol)
[![X (formerly Twitter) Follow](https://img.shields.io/twitter/follow/synapseprotocol?style=flat-square&logo=twitter)](https://x.com/synapseprotocol)



[Synapse](https://synapseprotocol.com/) is a interchain messaging protocol & bridge that allows for the creation of decentralized applications that can be deployed across multiple chains. This repo contains a variety of tools, interfaces and services that form the basis of, or are built on top of Synapse.

This repo serves as (or will serve) as the monrepo for all synapse code. Currently, only the [liquidity/bridge contracts](https://github.com/synapsecns/synapse-contracts) still need to be migrated.

## Contributing
Read through [CONTRIBUTING.md](./CONTRIBUTING.md) for a general overview of our contribution process.
Then check out our list of [good first issues](https://github.com/synapsecns/sanguine/contribute) to find something fun to work on!

## Projects

There are a variety of different packages in this repo, covered comprehensively in the [Directory Structure](#directory-structure) section of this document. However, for ease of use to new contributors, a comprehensive list of user-facing components are listed below, along with their source code components.

- Synapse Interchain Network: An optimistic proof-of-stake interchain messaging network.  Please see docs [here](https://docs.synapseprotocol.com/synapse-interchain-network-sin/synapse-interchain-network)
  - [Contracts](packages/contracts-core)
  - [Off Chain Agents](agents)
- Synapse Bridge: A cross-chain liquidity + canonical token bridge. Please see docs [here](https://docs.synapseprotocol.com/synapse-bridge/synapse-bridge)
  - [Contracts](https://github.com/synapsecns/synapse-contracts): The contracts used to power the synapse bridge.
  - [Frontend](packages/synapse-interface): The frontend used to interact with the bridge.
  - [SDK](packages/sdk-router): The SDK used to interact with the bridge.
- Explorer: A bridge explorer that allows users to view on-chain data state about the bridge.
  - [Backend](services/explorer): The backend used to power the explorer.
  - [Frontend](packages/explorer-ui): The frontend used to interact with the explorer.

## Directory Structure

<pre>
root
├── <a href="./agents">agents</a>: agents contain all the agents used in optimistic messaging
├── <a href="./charts">charts</a>: The helm charts used for deploying sanguine related services
├── <a href="./contrib">contrib</a>: Devops related tools
│   ├── <a href="./contrib/git-changes-action">git-changes-action</a>: Github action for identifying changes in dependent modules in a go workspace
│   ├── <a href="./contrib/promexporter">promexporter</a>: Multi-service prometheus exporter
│   ├── <a href="./contrib/release-copier-action">release-copier-action</a>: Github action for copying releases from one repo to another
│   ├── <a href="./contrib/terraform-provider-iap">terraform-provider-iap</a>: Terraform provider used for bastion proxy tunneling
│   ├── <a href="./contrib/terraform-provider-helmproxy">terraform-provider-helmproxy</a>: Terraform provider that allows helm to be proxied through an iap bastion proxy
│   ├── <a href="./contrib/terraform-provider-kubeproxy">terraform-provider-kubeproxy</a>: Terraform provider that allows kube to be proxied through an iap bastion proxy
│   ├── <a href="./contrib/tfcore">tfcore</a>: Terraform core utilities + iap utilities
├── <a href="./core">core</a>: The Go core library with common utilities for use across the monorepo
├── <a href="./ethergo">ethergo</a>: Go-based ethereum testing + common library
├── <a href="./packages">packages</a>
│   ├── <a href="./packages/contracts-core">contracts-core</a>: Core contracts used for synapse, powered by <a href="https://github.com/foundry-rs/foundry">Foundry</a>
│   ├── <a href="./packages/coverage-aggregator">coverage-aggregator</a>: Javascript coverage aggregator based on <a href="https://www.npmjs.com/package/nyc">nyc</a>
│   ├── <a href="./packages/docs">docs</a>: Docasaurus documentation. Note: this is not yet in use, and docs are still maintained on gitbook
│   ├── <a href="./packages/explorer-ui">explorer-ui</a>: Explorer UI
│   ├── <a href="./packages/sdk-router">sdk-router</a>: SDK router
│   ├── <a href="./packages/sdk-router">synapse-interface</a>: Synapse frontend code
├── <a href="./services">services</a>
│   ├── <a href="./services/cctp-relayer">CCTP Relayer</a>: CCTP message relayer
│   ├── <a href="./services/explorer">explorer</a>: Bridge/messaging explorer backend
│   ├── <a href="./services/scribe">scribe</a>: Generalized ethereum event logger
│   ├── <a href="./services/omnirpc">omnirpc</a>: Latency aware RPC Client used across multiple-chains at once
│   ├── <a href="./services/sinner">sinner</a>: [Synapse Interchain Network](https://interchain.synapseprotocol.com/) indexer & query interface
├── <a href="./tools">tools</a>
│   ├── <a href="./tools/abigen">abigen</a>: Used to generate abigen bindings for go
│   ├── <a href="./tools/bundle">bundle</a>: Modified version of <a href="https://pkg.go.dev/golang.org/x/tools@v0.5.0/cmd/bundle"> go bundler </a> with improved shadowing support
│   ├── <a href="./tools/modulecopier">module copier</a>: Used to copy internal modules and export methods for testing
</pre>

## Setup

Clone the repository, open it, and install nodejs packages with `yarn`:

```bash
git clone https://github.com/synapsecns/sanguine --recurse-submodules -j10
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

## Dealing with submodules

This repo make use of [multiple](.gitattributes) [submodules](https://git-scm.com/book/en/v2/Git-Tools-Submodules). To avoid issues when checking out different branches, you can use `git submodule update --init --recursive` after switching to a branch or `git checkout feat/branch-name --recurse-submodules` when switching branches.

# Building Agents Locally

<!-- TODO: we need to move this thing into an ops docs package. Given that the docs are still a work in progress, I'm leaving this here for now. -->
<!-- Actually, it's unclear if this belongs in a contributing.md file, the docs or both. Maybe a symlink? -->

In order to minimize risks coming from extraneous dependencies or supply chain attacks in a production like enviornment, all distributed images are built as [scratch](https://hub.docker.com/_/scratch) or [distroless](https://github.com/GoogleContainerTools/distroless#distroless-container-images) images. Builder containers are also not used to restrict the build enviornment to the [goreleaser container](https://github.com/synapsecns/sanguine/pkgs/container/sanguine-goreleaser). All production images are kept in the `docker/` file as `[dir].Dockerfile`. Local
