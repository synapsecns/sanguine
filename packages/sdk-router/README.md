# Synapse Router SDK

This package contains the Synapse Protocol Cross-Chain Swap and Bridging SDK

[See the Docs](https://synapserouter.gitbook.io/bridge-sdk-2)

# Synapse SDK

The Synapse SDK allows you to interact with [Synapse Protocol](https://synapseprotocol.com/) routers and bridges. It handles:

- Getting bridge quotes, performing bridging operations
- Getting swap quotes, performing swaps
- Getting liquidity pool information
- Calculating add/remove liquidity amounts

## Installation

```bash
npm install @synapsecs/sdk-router
```

## Usage

To use the SDK, first instantiate it with chain IDs and Ethereum providers:

```ts
import { SynapseSDK } from '@synapsecs/sdk-router'

const chainIds = [1, 42161, 10]
// Replace with JSON providers
const providers = [ethereumProvider, arbitrumProvider, optimismProvider]
const synapseSDK = new SynapseSDK(chainIds, providers)
```

### Bridging

Get a bridge quote from the Synapse Bridge Router:

```ts
const { maxAmountOut, originQuery, destQuery, feeConfig, routerAddress } =
  await Synapse.bridgeQuote(
    1,
    42161,
    '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48',
    '0xaf88d065e77c8cc2239327c5edb3a432268e5831',
    BigNumber.from('100000000')
  )
```

Perform a bridge through a Synapse Bridge Router or Synapse CCTP Router:

```ts
const { data, to } = await Synapse.bridge(
  addressTo,
  routerAddress,
  originChain,
  destinationChain,
  tokenIn,
  amountIn,
  originQuery,
  destQuery
)
```

The Synapse SDK allows quick and easy interaction with Synapse Protocol routers and bridges across multiple chains.
