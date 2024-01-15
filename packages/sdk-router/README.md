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

#### Getting a bridge quote

Below is the example of how to get the list of quotes for sending 1000 USDC from Ethereum and receiving USDT on Arbitrum:

```ts
const bridgeQuotes: BridgeQuote[] = await synapseSDK.bridgeQuote(
  // 1
  originChainId,
  // 42161
  destChainId,
  // Address of the token to start from on origin chain: 0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48
  tokenIn,
  // Address of the token to end with on destination chain: 0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9
  tokenOut,
  // Amount of tokens to bridge, in origin token decimals: 1_000_000_000
  amountIn,
  // Deadline for the transaction to be initiated on the origin chain, in seconds (optional)
  deadline
)
```

The returned list is sorted by the `maxAmountOut` field, so the first quote is the one yielding the highest amount of tokens on the destination chain.

> **Note:** The `bridgeQuote` method is a wrapper around the `allBridgeQuotes` method. `bridgeQuote` returns only the first quote from the list, while `allBridgeQuotes` returns the entire list.

#### Applying slippage

Some of the returned quotes may contain information about the optional swaps on origin and destination chains. As the liquidity composition may change over time, it is recommended to apply slippage to the quotes to account for the possible price changes. If no slippage is applied, the user transaction might be reverted due to insufficient funds. The default value for the slippage is 10 basis points (0.1%).

```ts
const { originQuery, destQuery } = await synapseSDK.applyBridgeSlippage(
  // fields from the BridgeQuote object returned by the allBridgeQuotes method
  bridgeQuote.bridgeModuleName,
  bridgeQuote.originQuery,
  bridgeQuote.destQuery,
  // Numerator of the slippage percentage, optional (defaults to 10)
  slipNumerator,
  // Denominator of the slippage percentage, optional (defaults to 10000)
  slipDenominator
)
```

> **Note**: this method will not modify the original `Query` objects, but will return new ones. This allows to change the applied slippage without having to re-fetch the quotes.

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
