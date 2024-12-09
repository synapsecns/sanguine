---
title: Bridge SDK
---

# Bridge SDK

The Synapse Bridge SDK is the easiest way to integrate cross-chain token & liquidity transfers into your application. It is fully isomorphic and can be used on both client and server sides.

The Synapse Bridge SDK is built on top of the [Synapse Router](/docs/Routers/Synapse-Router) contract.

### Use cases

- Integrate your front-end application with the Synapse Bridge.
- Provide bridge liquidity.
- Perform cross-chain arbitrage.
- Integrate the Synapse Javascript SDK with your non-Javascript application.

## Install

:::note requires Node v16+

The SDK has only been fully tested on Node 16+ or greater. Earlier versions are not guaranteed to work.

:::

Requires either the `npm` or `yarn` package manager.

| Options
|-
| `npm install @synapsecns/sdk-router`
| `yarn add @synapsecns/sdk-router`

## Configure Ethers

The SDK package relies on the `@ethersproject` and `ethers` dependencies, installed from `npm`.

To begin constructing bridge-related transactions, first set up your environment with your providers, and format them, along with the `chainIds` you will be using, to set up a `SynapseSDK` instance.

#### Ethers v5

```js
//Set up providers (RPCs) for each chain desired
const arbitrumProvider: Provider = new ethers.providers.JsonRpcProvider(
  'https://arb1.arbitrum.io/rpc'
)
const avalancheProvider: Provider = new ethers.providers.JsonRpcProvider(
  'https://api.avax.network/ext/bc/C/rpc'
)

//Structure arguments properly
const chainIds = [42161, 43114]
const providers = [arbitrumProvider, avalancheProvider]

//Set up a SynapseSDK instance
const Synapse = new SynapseSDK(chainIds, providers)
```

#### Ethers v6

:::tip Ethers v6

Use of Ethers v6 requires the `@ethersproject/providers` dependency to be installed via `npm` or `yarn`:

- `npm install @ethersproject/providers@^5.7.2`
- `yarn add @ethersproject/providers@^5.7.2`

:::

```js
import { JsonRpcProvider } from '@ethersproject/providers'

//Set up providers (RPCs) for each chain desired
const arbitrumProvider: Provider = new JsonRpcProvider(
  'https://arb1.arbitrum.io/rpc'
)
const avalancheProvider: Provider = new JsonRpcProvider(
  'https://api.avax.network/ext/bc/C/rpc'
)

//Structure arguments properly
const chainIds = [42161, 43114]
const providers = [arbitrumProvider, avalancheProvider]

//Set up a SynapseSDK instance
const Synapse = new SynapseSDK(chainIds, providers)
```

## Functions

:::tip Query Data Type

`originQuery` and `destQuery`, returned by `bridgeQuote()` and required for `bridge()`, are [`Query`](https://synapserouter.gitbook.io/untitled/) objects, which contain:

- `swapAdapter`: (string): 0x address of the swap adapter.
- `tokenOut`: (string): 0x address of the outputted token on that chain.
- `minAmountOut`: (Ethers BigNumber): The min amount of value exiting the transaction.
- `deadline`: (Ethers BigNumber): The deadline for the potential transaction.
- `rawParams`: (string): 0x params for the potential transaction.

:::

### `bridgeQuote()`

Get all relevant information regarding a possible transaction.

#### Parameters

`bridgeQuote()` requires the following arguments:

- `fromChain` (number): Origin chain id.
- `toChain` (number): Destination chain id.
- `fromToken` (string): 0x token address on the origin chain.
- `toToken` (string): 0x token address on the destination chain.
- `amount` (Ethers BigNumber): The amount (with the correct amount of decimals specified by the token on the origin chain)
- An `object` with three separate args:
  - `deadline` (Ethers BigNumber): Deadline for the transaction to be initiated on the origin chain, in seconds (optional)
  - `originUserAddress` (string): Address of the user on the origin chain, optional, mandatory if a smart contract is going to initiate the bridge operation
  - `excludedModules` (array): (optional) List of bridge modules to exclude from the result

#### Return value

`bridgeQuote` returns the following information

- `feeAmount` (Ethers BigNumber): The calculated amount of fee to be taken.
- `bridgeFee` (number): The percentage of fee to be taken.
- `maxAmountOut` (Ethers BigNumber): The maximum output amount resulting from the bridge transaction.
- `originQuery` (`Query`): The query to be executed on the origin chain.
- `destQuery` (`Query`): The query to be executed on the destination chain.

### `bridge()`

Use `bridgeQuote` to request a Bridge transaction

#### Parameters

- `toAddress` (number): The 0x wallet address on the destination chain.
- `routerAddress` (string): The 0x contract address on the origin chain of the bridge router contract.
- `fromChain` (number): The origin chain id.
- `toChain` (number): The destination chain id.
- `fromToken` (string): The 0x token address on the origin chain.
- `amount` (Ethers BigNumber): The amount (with the correct amount of decimals specified by the token on the origin chain)
- `originQuery` (`Query`): The query to be executed on the origin chain.
- `destQuery` (`Query`): The query to be executed on the destination chain.

#### Return value

- `to` (string): 0x wallet address on the destination chain.
- `data` (string): Output data in 0x hex format

### `allBridgeQuotes()`

#### Return value

Returns an array all possible bridge quotes, with the first item in the array being the cheapest route.

Quotes are returned from various Bridge "types" such as [CCTP](/docs/Routers/CCTP) or [RFQ](/docs/RFQ). More information is available on the [Synapse Router](/docs/Routers/Synapse-Router) page, or [SynapseCNS](https://github.com/synapsecns/sdk-router) Github repository.

## Examples

### Get a quote

```js
const quotes = await Synapse.bridgeQuote(
  routerAddress
  42161, // Origin Chain
  43114, // Destination Chain
  '0xff970a61a04b1ca14834a43f5de4533ebddb5cc8', // Origin Token Address
  '0xa7d7079b0fead91f3e65f86e8915cb59c1a4c664', // Destination Token Address
  BigNumber.from('20000000') // Amount in
  {
    // Deadline for the transaction to be initiated on the origin chain, in seconds (optional)
    deadline: 1234567890,
    // List of bridge modules to exclude from the result, optional.
    // Empty list means that all modules are included.
    excludedModules: ['SynapseBridge', 'SynapseCCTP', 'SynapseRFQ'],
    // Address of the user on the origin chain, optional.
    // MANDATORY if a smart contract is going to initiate the bridge operation on behalf of the user.
    originUserAddress: '0x1234567890abcdef1234567890abcdef12345678',
  }
)
```

### Request a transaction

```js
await Synapse.bridge(
  '0x0AF91FA049A7e1894F480bFE5bBa20142C6c29a9', // To Address
  bridgeQuote.routerAddress, // address of the contract to route the txn
  42161, // Origin Chain
  43114, // Destination Chain
  '0xff970a61a04b1ca14834a43f5de4533ebddb5cc8', // Origin Token Address
  BigNumber.from('20000000'), // Amount
  quote.originQuery, // Origin query from bridgeQuote()
  quote.destQuery // Destination query from bridgeQuote()
)
```

## Version 0.10.0 breaking changes

### Options object

- `deadline`, `excludeCCTP` (now `excludedModules`), and `originUserAddress` parameters are now found in an (optional) options object at the end of the arguments list for `bridgeQuote()`, and `allBridgeQuotes()`.
- `excludedModules` excludes one or more modules with an array of the module names. Supported names are `SynapseBridge`, `SynapseCCTP`, and `SynapseRFQ`.
- `originUserAddress` is required as part of the options object to initiate a bridge transaction on behalf of a user.

### Examples

```js
bridgeQuote(...arguments, {
  deadline: 1234567890,
  excludedModules: ['SynapseCCTP'],
  originUserAddress: '0x1234...',
})

allBridgeQuotes({
  deadline: 1234567890,
  excludedModules: ['SynapseCCTP'],
  originUserAddress: '0x1234...',
})
```

### `FastBridgeRouter`

The previous `FastBridgeRouter` deployment is deprecated, if your integration is using the hardcoded address, please see the router deployments/deprecated deployments table [here](https://github.com/synapsecns/sanguine/tree/master/packages/sdk-router#router-deployments)
