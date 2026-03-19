# Synapse Constants

[![GitHub Package](https://img.shields.io/github/v/release/synapsecns/sanguine?filter=@synapsecns/synapse-constants*&style=flat-square&label=github%20packages)](https://github.com/synapsecns/sanguine/packages)

This package contains the Synapse Protocol Token and Chain Constants.

## Installation

This package is published to [GitHub Packages](https://github.com/synapsecns/sanguine/packages). Configure your `.npmrc` to use the GitHub registry for `@synapsecns` packages:

```
@synapsecns:registry=https://npm.pkg.github.com
```

Then install with Yarn:

```bash
yarn add @synapsecns/synapse-constants
```

## Build

The following command will build the package locally

```
yarn build
```

## Usage

Importing supported tokens and chains:

```js
import { BRIDGABLE_TOKENS, CHAINS } from '@synapsecns/synapse-constants'
```

Importing a specific token:

```js
import { USDC } from '@synapsecns/synapse-constants'
```

## TODO

- [ ] Instructions on adding new chains
- [ ] Instructions on adding new tokens
- [ ] Instructions on generating new token route map
