---
sidebar_position: 1
---

# SDK Intro

The Synapse SDK is a Javascript sdk that wraps Synapse Router. It allows routing from any token to any token as long as a route is supported by a provider.

<!-- TODO: link to synapse router page -->
<!-- TODO: link to document providers -->

## Getting Started

To install the sdk, run `npx install @synapsecns/sdk-router`.

### Usage

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
