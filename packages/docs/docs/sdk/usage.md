---
sidebar_position: 1
---

# Usage

How to use the SDK to create quotes and bridge transactions.

<!-- TODO: link to synapse router page -->
<!-- TODO: link to document providers -->

## Setting up your Environment

In order to begin constructing bridge related transactions, developers need to set up their environment for the relevant chains. The first is configuring chains and providers.

```typescript
//Set up providers (RPCs) for each chain desired

const arbitrumProvider: Provider = new etherProvider.JsonRpcProvider(
  'https://arb1.arbitrum.io/rpc'
)
const avalancheProvider: Provider = new etherProvider.JsonRpcProvider(
  'https://api.avax.network/ext/bc/C/rpc'
)

//Structure arguments properly
const chainIds = [42161, 43114]
const providers = [arbitrumProvider, avalancheProvider]

//Set up a SynapseSDK instance
const Synapse = new SynapseSDK(chainIds, providers)
```

In the code snippet above, we set up the different providers and format them, along with the Ids of the chains we are bridging from/to for our quotes.

## Functions

### bridgeQuote()

Getting a bridge quote returns all relevant information regarding a possible transaction. Every bridge transaction requires this step.

#### Parameters

- `originChainId` - An integer representing the origin chain ID.
- `destChainId` - An integer representing the destination chain ID.
- `tokenIn` - A string representing the input token.
- `tokenOut` - A string representing the output token.
- `amountIn` - A BigintIsh representing the input token amount.

#### Return Value

A Promise object that resolves to an object with the following properties:

- `feeAmount` (BigNumber): A BigNumber object representing the fee amount for the transaction.
- `feeConfig` ([FeeConfig](#feeconfig)): A `FeeConfig` object representing the fee configuration for the transaction.
- `routerAddress` (string): A string representing the address of the router.
- `maxAmountOut` (BigNumber): A BigNumber object representing the maximum amount of tokens that can be received from the transaction (after fees).
- `originQuery` ([Query](#query)): A Query object representing the query to the origin SynapseRouter.
- `destQuery` ([Query](#query)): A Query object representing the query to the destination SynapseRouter.

<b>
  If the permutation of chains and tokens is not supported, the SDK will inform
  the user.
</b>

### bridge()

Returns [a prepared transaction](https://docs.ethers.org/v5/api/providers/types/#types--transactions) so that it may be executed by the client.

#### Parameters

- `to` (string): The recipient address on the destination chain.
- `originChainId` (number): The ID of the origin chain.
- `destChainId` (number): The ID of the destination chain.
- `token` (string): The token to be bridged.
- `amount` (BigintIsh): The amount of tokens to be bridged.
- `originQuery` ([Query](#query)): A `Query` object containing the origin SynapseRouter query (retrieved using `bridgeQuote`)
- `destQuery` ([Query](#query)): An `Query` object containing the destination SynapseRouter query (retrieved using `bridgeQuote`)

#### Return Value

A Promise [PopulatedTransaction](https://docs.ethers.org/v5/api/providers/types/#types--transactions) object is returned with the following attributes populated.

- `to` (string): The address the tx is to.
- `data` (string): The transaction data.

### swapQuote()

Returns a swap quote containing relevant information about a token swap transaction.

#### Parameters

- `chainId` (number): The ID of the chain.
- `tokenIn` (string): The input token for the swap.
- `tokenOut` (string): The output token for the swap.
- `amountIn` (BigintIsh): The amount of input tokens.

#### Return Value

A Promise object that resolves to an object with the following properties:

- `routerAddress` (string | undefined): The address of the router.
- `maxAmountOut` (BigNumber | undefined): The maximum amount of output tokens that can be obtained from the swap.
- `query` ([Query](#query) | undefined): A `Query` object representing the swap query.

### swap()

Executes a token swap transaction on the specified chain.

#### Parameters

- `chainId` (number): The ID of the chain.
- `to` (string): The recipient address for the swapped tokens.
- `token` (string): The token to be swapped.
- `amount` (BigintIsh): The amount of tokens to be swapped.
- `query` ([Query](#query)): An `Query` object containing the SynapseRouter swap query (retrieved using `swapQuote`)

#### Return Value

A Promise [PopulatedTransaction](https://docs.ethers.org/v5/api/providers/types/#types--transactions) object is returned with the following attributes populated.

- `to` (string): The address the tx is to.
- `data` (string): The transaction data.

### getBridgeGas()

Returns the gas amount for the bridge transaction on the specified chain.

#### Parameters

- `chainId` (number): The ID of the chain.

#### Return Value

A Promise object that resolves to a `BigintIsh` representing the gas amount for the bridge transaction.

### getPoolTokens()

Retrieves the pool tokens associated with a specific pool address on the specified chain.

#### Parameters

- `chainId` (number): The ID of the chain.
- `poolAddress` (string): The address of the pool.

#### Return Value

A Promise object that resolves to an array of `PoolToken` objects representing the pool tokens.

### getPoolInfo()

Retrieves information about a pool identified by its address on the specified chain.

#### Parameters

- `chainId` (number): The ID of the chain.
- `poolAddress` (string): The address of the pool.

#### Return Value

A Promise object that resolves to an object with the following properties:

- `tokens` (BigNumber | undefined): The number of tokens in the pool.
- `lpToken` (string | undefined): The address of the LP (liquidity provider) token associated with the pool.

### getAllPools()

Retrieves information about all pools on the specified chain.

#### Parameters

- `chainId` (number): The ID of the chain.

#### Return Value

A Promise object that resolves to an array of objects, each representing a pool. Each pool object has the following properties:

- `poolAddress` (string | undefined): The address of the pool.
- `tokens` ([PoolToken](#pooltoken)[] | undefined): An array of `PoolToken` objects representing the tokens in the pool.
- `lpToken` (string | undefined): The address of the LP (liquidity provider) token associated with the pool.

### calculateAddLiquidity()

Calculates the amount of liquidity tokens to be added when providing liquidity to a pool.

#### Parameters

- `chainId` (number): The ID of the chain.
- `poolAddress` (string): The address of the pool.
- `amounts` (object): A record of token addresses as keys and their corresponding amounts as values.

#### Return Value

A Promise object that resolves to an object with the following properties:

- `amount` (BigNumber): The calculated amount of liquidity tokens.
- `routerAddress` (string): The address of the router.

### calculateRemoveLiquidity()

Calculates the amounts of tokens to be received when removing liquidity from a pool.

#### Parameters

- `chainId` (number): The ID of the chain.
- `poolAddress` (string): The address of the pool.
- `amount` (BigNumber): The amount of liquidity tokens to be removed.

#### Return Value

A Promise object that resolves to an object with the following properties:

- `amounts` (Record<string, { value: BigNumber; index: number }>): A record of token addresses as keys and their corresponding amounts and index positions as values.
- `routerAddress` (string): The address of the router.

### calculateRemoveLiquidityOne()

Calculates the amount of a specific token to be received when removing liquidity from a pool.

#### Parameters

- `chainId` (number): The ID of the chain.
- `poolAddress` (string): The address of the pool.
- `amount` (BigNumber): The amount of liquidity tokens to be removed.
- `token` (string): The address of the token.

#### Return Value

A Promise object that resolves to an object with the following properties:

- `amount` ({ value: BigNumber; index: number }): The calculated amount of the specified token and its index position.
- `routerAddress` (string): The address of the router.

## Types

### Query

The `Query` type represents the instructions for the bridge or swap. This type has the following attributes.

- `swapAdapter` (string): The address of the swap adapter.
- `tokenOut` (string): The address of the outputted token on that chain.
- `minAmountOut` (BigNumber): The min amount of value exiting the transaction.
- `deadline` (BigNumber): The deadline for the potential transaction.
- `rawParams` (string): The raw params for the potential transaction.

### FeeConfig

The `FeeConfig` type represents the fee data stored in the bridge contract. This type has the following attributes.

- `bridgeFee` (number): The bridge fee taken for that query.
- `minFee` (BigNumber): The min fee possible for that bridge contract.
- `maxFee` (BigNumber): The max fee possible for that bridge contract.

### PoolToken

The `PoolToken` type represents a pool token with the following attributes.

- `token` (string): The token address.
- `isWeth` (boolean | undefined): Indicates whether the token is wrapped Ether (WETH).

### BigNumber/BigNumberish

BigNumber and BigNumberish are the [ethers v5 BigNumber and BigNumberish types](https://docs.ethers.org/v5/api/utils/bignumber/)
