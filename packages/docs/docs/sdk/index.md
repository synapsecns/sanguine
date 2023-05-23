---
sidebar_position: 1
---

# Router SDK

The Synapse SDK is the easiest way for any developer to integrate cross-chain token & liquidity transfers into their application. The SDK is built to support full-fledged frontend applications, but is fully isomorphic, able to be used both client & server-side.

<!-- TODO: link to synapse router page -->
<!-- TODO: link to document providers -->

## Getting Started

### Installation

Pre-reqs: Node v16+. The SDK is only fully tested on Node 16+ or greater. Earlier versions may have errors.
Depending on the package manager of your choice, install the SDK using one of the following

import InstallSwitch from '@site/src/components/SdkFeatures/InstallSwitch'

<InstallSwitch />

### Quickstart

```typescript
import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber } from '@ethersproject/bignumber'
import { SynapseSDK } from '@synapsecns/sdk'

const arbitrumProvider: Provider = new etherProvider.JsonRpcProvider(
  'https://arb1.arbitrum.io/rpc'
)
const avalancheProvider: Provider = new etherProvider.JsonRpcProvider(
  'https://api.avax.network/ext/bc/C/rpc'
)

// get a bridge quote for asset bridge from chain a->b
// for token 0xf->0xa
const chainIds = [42161, 43114]
const providers = [arbitrumProvider, avalancheProvider]
const Synapse = new SynapseSDK(chainIds, providers)
const quotes = await Synapse.bridgeQuote(
  42161,
  43114,
  '0xff970a61a04b1ca14834a43f5de4533ebddb5cc8',
  '0xa7d7079b0fead91f3e65f86e8915cb59c1a4c664',
  BigNumber.from('20000000')
)
console.log(quotes)

// bridge it
await Synapse.bridge(
  '0x0AF91FA049A7e1894F480bFE5bBa20142C6c29a9',
  42161,
  43114,
  '0xff970a61a04b1ca14834a43f5de4533ebddb5cc8',
  BigNumber.from('20000000'),
  quotes.originQuery,
  quotes.destQuery
)
```
