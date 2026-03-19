---
slug: npm-to-github-packages-migration
title: Synapse NPM Packages Are Moving to GitHub Packages
# authors: [synapse]
tags: [update, npm, packages]
---

All `@synapsecns` packages are now published to **GitHub Packages** instead of npmjs.org.

<!--truncate-->

## What Changed

We've migrated our package registry from npmjs.org to [GitHub Packages](https://github.com/orgs/synapsecns/packages). This affects the following packages:

- `@synapsecns/sdk-router`
- `@synapsecns/synapse-constants`
- `@synapsecns/contracts-core`
- `@synapsecns/solidity-devops`
- `@synapsecns/coverage-aggregator`
- `@synapsecns/widget`

## How to Update Your Project

To install `@synapsecns` packages from GitHub Packages, add a `.npmrc` file to the root of your project (next to your `package.json`):

```
@synapsecns:registry=https://npm.pkg.github.com
```

Then install packages as usual:

```bash
npm install @synapsecns/sdk-router
# or
yarn add @synapsecns/sdk-router
```

GitHub Packages requires authentication even for public packages. You can authenticate by adding your GitHub personal access token (with `read:packages` scope) to your `.npmrc`:

```
//npm.pkg.github.com/:_authToken=YOUR_GITHUB_TOKEN
```

Or set the environment variable:

```bash
export NPM_TOKEN=YOUR_GITHUB_TOKEN
```

## Why the Move

This migration consolidates our release infrastructure under GitHub, simplifying CI/CD and aligning package publishing with the repository where development happens.

Existing versions already published on npmjs.org will remain available. All **new** versions will be published exclusively to GitHub Packages.

If you have any questions, feel free to open an issue on [GitHub](https://github.com/synapsecns/sanguine/issues).
