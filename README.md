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
│   ├── <a href="./packages/contracts-rfq">rfq</a>: RFQ contracts
│   ├── <a href="./packages/coverage-aggregator">coverage-aggregator</a>: Javascript coverage aggregator based on <a href="https://www.npmjs.com/package/nyc">nyc</a>
│   ├── <a href="./packages/docs">docs</a>: Docasaurus documentation. Note: this is not yet in use, and docs are still maintained on gitbook
│   ├── <a href="./packages/explorer-ui">explorer-ui</a>: Explorer UI
│   ├── <a href="./packages/sdk-router">sdk-router</a>: SDK router
│   ├── <a href="./packages/sdk-router">synapse-interface</a>: Synapse frontend code
├── <a href="./services">services</a>
│   ├── <a href="./services/cctp-relayer">CCTP Relayer</a>: CCTP message relayer
│   ├── <a href="./services/explorer">explorer</a>: Bridge/messaging explorer ba
│   ├── <a href="./services/rfq">rfq</a>: RFQ contracts
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

## Contribution workflow
<!-- TODO: this actually belongs in a contributing.md file, but let's be honest, how many people actually read those? -->
<!-- Additionally, the two-branch approach is somewhat temporary, so having this section in the README is probably fine. -->

We use a two-branch strategy for development:

- `master`: This is the primary development branch where all development happens. All regular pull requests (new features, bug fixes, etc) should be opened against this branch.
  - Refer to [Implementing a New Feature](#scenario-1-implementing-a-new-feature) for more details.
- `fe-release`: This branch is used for production front-end releases and is the one that gets deployed to production. The production front-end build always uses the latest commit from this branch.
  - Refer to [Releasing a Front-end Update](#scenario-2-releasing-a-front-end-update) for more details.

> `master` should never be behind `fe-release`! The only exception is when a hotfix is needed on the production front-end.
> Refer to [Hotfixing the Production Front-end](#scenario-3-hotfixing-the-production-front-end) for more details.

We use the following merge strategies:

- **Squash merge**: All pull requests are squash merged into `master`. This keeps the commit history clean and makes it easier to revert changes if necessary.
  > On a very few occasions, we may use a regular merge when `master` needs to be updated with a few commits from `fe-release`. In this case, we will use a regular merge so that there are no merge conflicts when later merging into `fe-release`.
- **Regular merge**: Latest changes from `master` are **regularly merged** into `fe-release`. This ensures that the production front-end always uses the latest changes from `master`, and prevents merge conflicts when merging into `fe-release`.
  > `master` branch should never be the **source branch** when merging into `fe-release`. This is because the source branch is always deleted after merging, and we don't want to delete `master`.

### Scenario 1: Implementing a New Feature

In this scenario you are implementing a new feature that doesn't automatically trigger a front-end release.

1. **Create a new branch**: From the `master` branch, create a new branch for your feature. The branch name should be descriptive of the feature you're implementing.

```bash
git checkout master && git pull
# pkg stands for the package you're working on
git checkout -b pkg/feature-name
```

2. **Implement your feature**: Make your changes in this branch.
3. **Commit your changes**: Regularly commit your changes with clear, concise commit messages.
4. **Push your branch**: When you're ready to open a pull request, push your branch to the remote repository and open a pull request on GitHub. Add an overview of your changes and a link to the relevant issue (if applicable).

```bash
git push -u origin pkg/feature-name
```

5. **Sync with `master`**: If you need to use the latest changes from `master`, you can merge them into your branch. This is especially useful if you're working on a long-running feature branch (or if some tests are failing on your branch, which have been fixed on `master`).

```bash
git checkout master & git pull
git checkout pkg/feature-name
git merge master
```

6. Alternatively, you can rebase your branch on top of `master`. However, this will rewrite your commit history, which can be problematic if other contributors have already pulled your branch or started reviewing your code. The rule of thumb here is:

- If your branch hasn't been pushed yet, **always** rebase.
- If your branch has been pushed, but it is a draft that no one has reviewed yet, you **could** rebase.
- Otherwise, you **should** merge.

```bash
git checkout master & git pull
git checkout pkg/feature-name
# You might need to drop some of your commits here, if they are already in master
# Edit the rebase-todo file to squash, drop, reorder, etc your commits
git rebase -i master
```

7. **CI checks**: Once you've pushed your branch, the CI checks will run automatically. If any of the checks fail, you can fix the issues in your feature branch and push again. The CI checks will run again automatically.
8. **Review and merge**: The PR will be reviewed. If any changes are requested, make them in your feature branch and the PR will automatically update. Once the PR is approved by at least one maintainer, it can be **squash merged** into `master` and your feature branch will be deleted.

### Scenario 2: Releasing a Front-end Update

In this scenario you are releasing a front-end update using the latest changes from `master`.

1. **Create a new branch**: From the `fe-release` branch, create a new branch for the FE release.

```bash
git checkout fe-release && git pull
# Date format is YYYY-MM-DD (sorry my American friends)
git checkout -b release/date
```

2. **Merge `master` into your branch**: This ensures that the production front-end always uses the latest changes from `master`, and prevents merge conflicts when merging into `fe-release`.

```bash
git checkout master && git pull
git checkout release/date
# No merge conflicts should occur here
git merge master
```

3. **Push your branch**: Push your branch to the remote repository and open a pull request on GitHub. No overview is necessary, but you can add links to the PRs that are being released.

```bash
git push -u origin release/date
```

4. **CI checks**: Once you've pushed your branch, the CI checks will run automatically. Assuming that master is passing all checks, your branch should pass all checks as well.
5. **Review and merge**: The PR will be reviewed. Once the PR is approved by at **least two maintainers**, it can be **regularly merged** into `fe-release` and your release branch will be deleted.

### Scenario 3: Hotfixing the Production Front-end

Sometimes, a bug is discovered in the production front-end that needs to be fixed immediately. In this scenario, you are hotfixing the production front-end without using the latest changes from `master`.

1. **Create a new branch**: From the `fe-release` branch, create a new branch for the hotfix.

```bash
git checkout fe-release && git pull
# Date format is YYYY-MM-DD (sorry my American friends)
git checkout -b hotfix/date
```

2. **Implement your hotfix**: Make your changes in this branch.

3. **Push your branch**: Push your branch to the remote repository and open a pull request on GitHub.

```bash
git push -u origin hotfix/date
```

4. **CI checks**: Once you've pushed your branch, the CI checks will run automatically. Depending on the severity of the hotfix, you might want to merge the PR as soon or before it passes all checks. However, **you should wait for the linting checks to pass**, as these are the fastest and easiest to fix.

5. **Review and merge**: The PR will be reviewed. Once the PR is approved by at **least two maintainers**, it can be **squash merged** into `fe-release` and your hotfix branch will be deleted.

6. Take a deep breath and relax. You've just saved the day. 🦸‍♂️

7. **Catch up `master`**: now that the hotfix is released, you should catch up `master` with the latest changes from `fe-release`. This ensures that `master` branch is never behind `fe-release`.

```bash
git checkout fe-release && git pull
git checkout master && git pull
git checkout -b catchup/date
git merge fe-release
```

8. **Push your branch**: Push your `catchup/date` branch to the remote repository and open a pull request on GitHub.

```bash
git push -u origin catchup/date
```

9. **CI checks**: Once you've pushed your branch, the CI checks will run automatically. Assuming that `fe-release` is passing all checks, your branch should pass all checks as well.

10. **Review and merge**: The PR will be reviewed. Once the PR is approved by at **least one maintainer**, it can be **regularly merged** into `master` and your catchup branch will be deleted.

# Building Agents Locally

<!-- TODO: we need to move this thing into an ops docs package. Given that the docs are still a work in progress, I'm leaving this here for now. -->
<!-- Actually, it's unclear if this belongs in a contributing.md file, the docs or both. Maybe a symlink? -->

In order to minimize risks coming from extraneous dependencies or supply chain attacks in a production like enviornment, all distributed images are built as [scratch](https://hub.docker.com/_/scratch) or [distroless](https://github.com/GoogleContainerTools/distroless#distroless-container-images) images. Builder containers are also not used to restrict the build enviornment to the [goreleaser container](https://github.com/synapsecns/sanguine/pkgs/container/sanguine-goreleaser). All production images are kept in the `docker/` file as `[dir].Dockerfile`. Local
