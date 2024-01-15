# Synapse Router SDK

[![npm](https://img.shields.io/npm/v/%40synapsecns%2Fsdk-router?style=flat-square)](https://www.npmjs.com/package/@synapsecns/sdk-router)

This package contains the Synapse Protocol Cross-Chain Swap and Bridging SDK

[See the Docs](https://synapserouter.gitbook.io/bridge-sdk-2)

# Synapse SDK

The Synapse SDK allows you to interact with [Synapse Protocol](https://synapseprotocol.com/) router contracts deployed on 19 chains. It handles:

- Bridging operations (cross-chain swaps):
  - Getting bridge quotes
  - Initiating bridge transactions
  - Tracking the status of bridge transactions
- On-chain swap operations:
  - Getting swap quotes
  - Initiating swap transactions
- Utilities for getting miscellaneous data related to protocol, fees, and chains

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

`BridgeQuote` objects are returned by the `bridgeQuote` and `allBridgeQuotes` methods. They contain the following fields:

```ts
export type BridgeQuote = {
  // Address of the Router contract that user will be interacting with
  routerAddress: string
  // Expected final amount of tokens to be received on the destination chain by the user,
  // if the bridge transaction is completed right after the quote is generated
  maxAmountOut: BigNumber
  // Query object for the origin chain
  originQuery: Query
  // Query object for the destination chain
  destQuery: Query
  // Estimated median time for the bridge transaction to be completed
  estimatedTime: number
  // Name of the "bridge module" that will be used to bridge the tokens.
  // Supported values are "SynapseBridge", "SynapseCCTP" and "SynapseRFQ"
  bridgeModuleName: string
  // Amount of native gas tokens that user will receive on the destination chain
  // on top of the token amount
  gasDropAmount: BigNumber
}
```

> **Note:** `Query` objects contain information about the optional swaps to be performed on behalf of the user on origin and destination chains. The exact composition of the `Query` object, as well as the concept of the optional swaps, is abstracted away from the SDK consumer.
> A collection of methods to modify the `Query` object is provided in the `SynapseSDK` class, allowing the consumer to be unaware of the underlying object structure.

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

The Synapse SDK allows quick and easy interaction with Synapse Protocol routers and bridges across multiple chains
