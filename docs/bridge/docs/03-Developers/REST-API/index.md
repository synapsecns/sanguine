---
title: REST API
---

:::warning

The [previous version](https://synapse-rest-api-v2.herokuapp.com/) is no longer maintained and will be fully deprecated by October 2024.

:::

# REST API

Get read-only data from on-chain Synapse contracts, and generate Bridge and Swap quotes, plus additional transaction information.

## Use cases

* Integrate your front-end application with the Synapse Bridge.
* Provide bridge liquidity.
* Perform cross-chain arbitrage.
* Integrate the Synapse Javascript SDK with your non-Javascript application.

## Base URL

[`api.synapseprotocol.com/`](https://api.synapseprotocol.com/)

## GET Endpoints

### `/swap`

Get a token swap quote.

#### Parameters

* `chain` (int): `chainId` of the desired chain.
* `fromToken` (string): Address of token to swap from.
* `toToken` (string): Address of token to swap to.
* `amount` (int): Amount of `fromToken` to swap.

#### Returns

* `routerAddress` (string): The address of the router contract
* `maxAmountOut {` (object): The maximum amount of tokens that can be swapped out.
    * `type:` (string): The data type
    * `hex:` (string): The amount encoded in hexidecimal
* `query {` (object): Parameters for the swap query:
    * `0:` (string): Router contract address
    * `1:` (string): Address of `tokenIn`
    * `2:` (object): Amount of `tokenIn` to swap (same structure as `maxAmountOut`)
    * `3:` (object): Minimum amount of `tokenOut` requested (same structure as `maxAmountOut`)
    * `4:` (string): Encoded params for swap routing
    * `swapAdapter` (string): Address of the swap adapter contract
    * `tokenOut` (string): Address of `tokenOut`
    * `minAmountOut` (object): Minimum amount of `tokenOut` required (same structure as `maxAmountOut`)
    * `deadline` (object): Deadline parameter for the swap (same structure as `maxAmountOut`)
    * `rawParams` (string): Encoded hex string containing swap parameters
* `maxAmountOutStr` (string): The `maxAmountOut` value formatted as a decimal string

#### Example

[`https://api.synapseprotocol.com/swap?chain=1&fromToken=0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48&toToken=0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85&amount=100`](https://api.synapseprotocol.com/swap?chain=1&fromToken=0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48&toToken=0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85&amount=100)

### `/bridge`

Get a token bridge quote.

#### Parameters

* `fromChain` (int): `chainId` of origin chain
* `toChain` (int): `chainId` of destination chain
* `fromToken` (string): Token address to bridge
* `toToken` (string): Token address to bridge to
* `amount` (int): Amount to bridge
* **Optional**`originUserAddress` (string): Address of the user bridging the token

#### Returns

* `feeAmount` (object): The fee amount for the swap. Contains:
* `type` (string): Data type
* `hex` (string): Fee amount encoded in hex
* `feeConfig` (array): Fee configuration parameters, contains:
* `0` (number): Gas price
* `1` (object): Fee percentage denominator (hex encoded BigNumber)
* `2` (object): Protocol fee percentage numerator (hex encoded BigNumber)
* `routerAddress` (string): Address of the router contract
* `maxAmountOut` (object): Maximum amount receivable from swap, structure same as above
* `originQuery` (object): Original swap query parameters, contains:
* `swapAdapter` (string): Swap adapter address
* `tokenOut` (string): Address of output token
* `minAmountOut` (object): Minimum output token amount
* `deadline` (object): Expiry time
* `rawParams` (string): Encoded hex params
* `destQuery` (object): Destination swap query parameters, structure similar to `originQuery` above.
* `maxAmountOutStr` (string): `maxAmountOut` as a decimal string.
* `bridgeModuleName` (string): the bridge module the transaction will be routed through
* `gasDropAmount` (BigNumber): the amount of gas airdropped to the user on the dest chain.

#### Example

[`/bridge?fromChain=1&toChain=10&fromToken=0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48&toToken=0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85&amount=1000000`](https://api.synapseprotol.com/bridge?fromChain=1&toChain=10&fromToken=0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48&toToken=0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85&amount=1000000)

### `/swapTxInfo`

Get transaction data for executing a swap.

#### Parameters

* [`/swap`](#swap) parameters plus:
* `address`: Address attempting to swap assets.

#### Returns

* `data`: Binary data that forms the input to the transaction.
* `to`: Set as the [Synapse Router address](/docs/Contracts/Synapse-Router)

#### Example

[`/swapTxInfo?chain=1&fromToken=0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48&toToken=0x6b175474e89094c44da98b954eedeac495271d0f&amount=100&address=0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5`](https://api.synapseprotocol.com/swapTxInfo?chain=1&fromToken=0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48&toToken=0x6b175474e89094c44da98b954eedeac495271d0f&amount=100&address=0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5)

### `/bridgeTxInfo`
Used to get transaction data for executing a bridge.

#### Parameters

* [`/bridge`](#bridge) parameters plus:
  * `destAddress`: Address attempting to receive assets.

#### Returns

Returns txInfo for the best (highest expected output) quote.

* `data`: Binary data that forms the input to the transaction.
* `to`: Set as the [Synapse Router address](/docs/Contracts/Synapse-Router)

#### Example

[`/bridgeTxInfo?fromChain=1&toChain=10&fromToken=0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48&toToken=0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85&amount=1000000&destAddress=0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5`](https://api.synapseprotocol.com/bridgeTxInfo?fromChain=1&toChain=10&fromToken=0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48&toToken=0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85&amount=1000000&destAddress=0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5)

### `/synapseTxId`

Get the unique `synapseTxId` data needed to execute a bridge.

#### Parameters

* `originChainId` (int): `chainId` of the origin chain
* `bridgeModule` (string): String of the bridge module used (e.g. "SynapseRFQ")
* `txHash` (string): Transaction hash of the bridge transaction on the origin chain

#### Returns

* `synapseTxId` (Kappa)

#### Example

[`/synapseTxId?originChainId=8453&bridgeModule=SynapseRFQ&txHash=0x4acd82091b54cf584d50adcad9f57c61055beaca130016ecc3798d3d61f5487a`](https://api.synapseprotocol.com/synapseTxId?originChainId=8453&bridgeModule=SynapseRFQ&txHash=0x4acd82091b54cf584d50adcad9f57c61055beaca130016ecc3798d3d61f5487a)

### `/bridgeTxStatus`
Get the status of a bridge transaction, with destination transaction information if finalized.

#### Parameters
* `destChainId` (int): `chainId` of the destination chain
* `bridgeModule` (string): String of the bridge module used (e.g. "SynapseRFQ")
* `synapseTxId` (string): The unique `SynapseTxId` (Kappa)

#### Returns

* `status` (boolean): Returns `true` if transaction is complete.
* `toInfo` (object): Completed transaction information, or `null` if transaction is not complete.

#### Example

[`/bridgeTxStatus?destChainId=534352&bridgeModule=SynapseRFQ&synapseTxId=0xf4315cb818ad76305dc5fbd55181263688ffeb3fb3e1c6abc2b849a48b3a7c78`](https://api.synapseprotocol.com/bridgeTxStatus?destChainId=534352&bridgeModule=SynapseRFQ&synapseTxId=0xf4315cb818ad76305dc5fbd55181263688ffeb3fb3e1c6abc2b849a48b3a7c78)

### `/destinationTx`
Get the status of a bridge transaction, and the destination transaction information if the transaction is finalized. This is a simple implementation of the above two methods.

#### Parameters
* `originChainId` (int): `chainId` of the origin chain
* `txHash` (string): Transaction hash of the Bridge transaction on the origin chain

#### Returns
* `status` (string): Returns `completed` if transaction is complete
* `toInfo` (object): Completed transaction information, or `null` if transaction is not complete.

#### Example

[`/destinationTx?originChainId=1&txHash=0x93f9d78516ee5fbce2595519ec97e03ff03778af600acb1769d0ce6def32b804`](https://api.synapseprotocol.com/destinationTx?originChainId=1&txHash=0x93f9d78516ee5fbce2595519ec97e03ff03778af600acb1769d0ce6def32b804)

### `/destinationTokens`
Used to return which tokens you can bridge to, once an origin token is identified.

#### Parameters
* `fromChain` (int): `chainId` of the origin chain
* `fromToken` (string): Transaction hash of the bridge transaction on the origin chain

#### Returns

// This function returns a list of all tokens that the fromToken can be bridged to, along with data about that token. Each object in the list returns:

* Array of token objects the `fromToken` can be bridged to, containing:
  * `symbol`: The token symbol of the destination token identified
  * `chainId`: The `chainId` of the destination token identified
  * `addres`: The token address of the destination token identified


#### Example

[`/destinationTokens?fromChain=1&fromToken=0xdAC17F958D2ee523a2206206994597C13D831ec7`](https://api.synapseprotocol.com/destinationTokens?fromChain=1&fromToken=0xdAC17F958D2ee523a2206206994597C13D831ec7)

## Javascript examples

### Estimate bridge output

```js
async function estimateBridgeOutput(
  fromChain,
  toChain,
  fromToken,
  toToken,
  amountFrom
) {
  const query_string = `fromChain=${fromChain}&toChain=${toChain}&fromToken=${fromToken}&toToken=${toToken}&amountFrom=${amountFrom}`;
  const response = await fetch(
    `https://api.synapseprotocol.com/bridge?${query_string}`
  );

  const response_json = await response.json();
  // Check if the response is an array and has at least one item
  if (Array.isArray(response_json) && response_json.length > 0) {
    return response_json[0]; // Return the first item
  } else {
    throw new Error('No bridge quotes available');
  }
}

estimateBridgeOutput(
  1,     // Ethereum
  42161, // Arbitrum
  "USDC",
  "USDC",
  "1000"
).then(firstQuote => {
  console.log('First bridge quote:', firstQuote);
}).catch(error => {
  console.error('Error:', error.message);
});
```

### Generate unsigned bridge transaction

```js
async function generateUnsignedBridgeTxn(
  fromChain,
  toChain,
  fromToken,
  toToken,
  amountFrom,
  destAddress
) {
  const query_string = `fromChain=${fromChain}&toChain=${toChain}&fromToken=${fromToken}&toToken=${toToken}&amount=${amountFrom}&destAddress=${addressTo}`;
  const response = await fetch(
    `https://api.synapseprotocol.com/bridgeTxInfo?${query_string}`
  );
  const response_json = await response.json();
  return await response_json;
}

generateUnsignedBridgeTxn(
  1,     // Ethereum
  42161, // Arbitrum
  "USDC",
  "USDC",
  "1000"
  "0x2D2c027E0d1A899a1965910Dd272bcaE1cD03c22"
);
```

## Support

Please read the documentation and examples carefully before reaching out on [Discord](https://discord.gg/synapseprotocol) for questions.
