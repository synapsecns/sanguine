# Synapse Router

The Synapse Router overhauls current Synapse Bridge contracts to abstract much of the complexity around liquidity based bridging to one simple [`bridge()`](/docs/Bridge/SDK/#bridge) function.

<!-- The new Router is comprised of one bridge() function and three supporting functions that help to construct a bridge transaction. All of the mentioned are organized by an important struct called a “Query”. Before diving into these functions, a deeper understanding of how the bridge actually works is fundamental. -->

<!-- Find out more about Synapse intermediary tokens nUSD and nETH in the [docs](https://docs.synapseprotocol.com/reference/faq#synapse-bridge). -->

## Queries

:::tip Intermediate tokens

Most Bridge transactions require an [“intermediary token”](/docs/Bridge/#pool-liquidity) which the protocol has mint/burn permissions over. When users Bridge native or wrapped tokens, their transaction is routed through a liquidity pool on both chains.

**Swaps to or from an intermediate token will return an empty query.**

:::

Determining an optimal route with minimal slippage and maximum output is not trivial. Synapse Router condenses this complexity into a [`Query`](https://github.com/synapsecns/synapse-contracts/blob/7bc3a579c08838d7f4c3e62954135901abb16183/contracts/bridge/libraries/BridgeStructs.sol#L12-L18), or data structure that describes generic swap instructions that will result in a `token`, and an `amountReceived`.



Given a bridgable token, a `Query` identifies on-chain “swaps” which allow the Bridge to identify the route(s) for an origin and intermediate token swap on the origin chain, and subsequent token swap on the destination chain.

Origin and destination queries are taken from the `getOriginAmountOut()` and `getDestinationAmountOut()` supporting functions, which are called for each transaction, from which your application can decide its preferred route.


<!-- See the Example Page for further information (and arguments) that the functions above require. It is imperative that the program uses the functions here to construct Queries instead of manually doing so -- this guarantees that the transaction won't be reverted for misconfigured parameters.  -->



## Constructing a Bridge transaction

:::tip Avoid manual construction

Attempting to manually construct queries may result in transactions being reverted for misconfigured parameters. Always use the Router functions below to construct your Bridge transactions.

:::

**1. Get output `tokenSymbols`**

Call `getConnectedBridgeTokens()` with your output token to receive a formatted `tokenSymbols` list.

**2. Get `originQuery` list**

Call `getOriginAmountOut()` with your input token, `tokenSymbols` list, and `amountIn`, to receive a `Query` list for the origin chain.

**3. Get `destRequest` list**

Convert each origin `Query` to a `destRequest` as seen in this example:

```js
let requests = symbols.map((value, index) => {
  let request: DestRequest = {
    symbol: value,
    amountIn: originQueries[index].minAmountOut,
  };
  return request;
});
```

**4. Get `destQuery` list**

Call `getDestinationAmoutOut()` with your `destRequest` list and output token to receive a `Query` list for the destination chain.

**5. select `originQuery` and `destQuery`**

Determine which `originQuery` and `destQuery` to use. This simple example selects the origin and destination pair with the highest output:

```js
let destQuery = maxBy(destQueries, (query) => query.minAmountOut);
let selectedIndex = destQueries.indexOf(destQuery)
let originQuery = originQueries[selectedIndex]
```

**6. Format queries and apply user settings**

Add any user settings such as `slippage`, and `deadline` to your queries, and specify a [`swapAdapter`](#swap-adapter) for the swap to use.

**7. `bridge()`**

Call `bridge()` with your selected `originQuery` and `destQuery` pair.

## Swap Adapter

:::note

SynapseRouterV1's swap adapter supports Synapse hosted pools. Future versions will allow additional adapters to support aggregators on different chains, allowing any-to-any Bridge transactions.

:::

The Synapse Adapter is a configurable wrapper that facilitates the "swap" action on the origin and destination chains, and exposes useful methods to get `Quote` and `Query` structs, supported pools, tokens, and more.

## Example

### Direct contract integration

```js
/**
 * Struct representing a bridge token.
 * @param symbol  Bridge token symbol: unique token ID consistent among all chains
 * @param token   Bridge token address
 */
type BridgeToken = {
  symbol: String;
  token: Address;
};

/**
 * Struct representing a request for a swap quote from a bridge token.
 * @param symbol    Bridge token symbol: unique token ID consistent among all chains
 * @param amountIn  Amount of bridge token to start with, before the bridge fee is applied
 */
type DestRequest = {
  symbol: String;
  amountIn: BigInt;
};

/**
 * Struct representing a request swap (list of instructions) for SynapseRouter.
 * @param swapAdapter   Adapter address that will perform the swap. Address(0) specifies a "no swap" query.
 * @param tokenOut      Token address to swap to.
 * @param minAmountOut  Minimum amount of tokens to receive after the swap, or tx will be reverted.
 * @param deadline      Latest timestamp for when the transaction needs to be executed, or tx will be reverted.
 * @param rawBytes      ABI-encoded params for the swap that will be passed to `swapAdapter`.
 */
type SwapQuery = {
  swapAdapter: Address;
  tokenOut: Address;
  minAmountOut: BigInt;
  deadline: BigInt;
  rawParams: BytesLike;
};

type UserSettings = {
  maxSlippage: BigInt;
  deadlineOrigin: BigInt;
  deadlineDest: BigInt;
};

interface SynapseRouter {
  /**
   * Initiate a bridge transaction with an optional swap on both origin and destination chains
   * @param to            Address to receive tokens on destination chain
   * @param chainId       Destination chain id
   * @param token         Initial token for the bridge transaction to be pulled from the user
   * @param amount        Amount of the initial tokens for the bridge transaction
   * @param originQuery   Origin swap query. Empty struct indicates no swap is required
   * @param destQuery     Destination swap query. Empty struct indicates no swap is required
   */
  bridge(
    to: Address,
    chainId: Number,
    token: Address,
    amount: BigInt,
    originQuery: SwapQuery,
    destQuery: SwapQuery
  ): null;

  /**
   * Gets the list of all bridge tokens (and their symbols), such that destination swap
   * from a bridge token to `tokenOut` is possible.
   * @param tokenOut  Token address to swap to on destination chain
   */
  getConnectedBridgeTokens(tokenOut: Address): BridgeToken[];

  /**
   * Finds the best path between `tokenIn` and every supported bridge token from the given list,
   * treating the swap as "origin swap", without putting any restrictions on the swap.
   * @param tokenIn       Initial token that user wants to bridge/swap
   * @param tokenSymbols  List of symbols representing bridge tokens
   * @param amountIn      Amount of tokens user wants to bridge/swap
   */
  getOriginAmountOut(
    tokenIn: Address,
    tokenSymbols: String[],
    amountIn: BigInt
  ): SwapQuery[];

  /**
   * Finds the best path between every supported bridge token from the given list and `tokenOut`,
   * treating the swap as "destination swap", limiting possible actions to those available for every bridge token.
   * Will take the bridge fee into account, when returning a quote for every bridge token.
   * @param requests  List of structs with following information:
   *                  - symbol: unique token ID consistent among all chains
   *                  - amountIn: amount of bridge token to start with, before the bridge fee is applied
   * @param tokenOut  Token user wants to receive on destination chain
   */
  getDestinationAmountOut(
    requests: DestRequest[],
    tokenOut: Address
  ): SwapQuery[];
}

/// Perform a cross-chain swap using Synapse:Bridge
/// Start from `amountIn` worth of `tokenIn` on origin chain
/// Receive `tokenOut` on destination chain
function synapseBridge(
  originChainId: Number,
  destChainId: Number,
  tokenIn: Address,
  tokenOut: Address,
  amountIn: BigInt,
  userOrigin: Address,
  userDest: Address,
  userSettings: UserSettings
) {
  // Every cross-chain swap via Synapse:Bridge is fueled by using one of the
  // supported "bridge tokens" as the intermediary token.
  // A following set of actions will be initiated by a single SynapseRouter.bridge() call:
  // - Origin chain: tokenIn -> bToken swap is performed
  // - Synapse: bridge bToken from origin chain to destination
  // - Destination chain: bToken -> tokenOut is performed

  // Here we describe a list of actions to perform such a cross-chain swap, knowing only
  // - tokenIn, tokenOut, amountIn
  // - SynapseRouter deployments
  // - User settings for maximum slippage and deadline
  // - User address on origin and destinaion chain (might be equal or different)

  // Beware: below is a TypeScript pseudocode.

  // 0. Fetch deployments of SynapseRouter on origin and destiantion chains
  let routerOrigin = getSynapseRouter(originChainId);
  let routerDest = getSynapseRouter(destChainId);

  // 1. Determine the set of bridge tokens that could enable "receive tokenOut on destination chain"
  // For that we pefrorm a static call to SynapseRouter on destination chain
  let bridgeTokens = routerDest.getConnectedBridgeTokens(tokenOut);
  // Then we get the list of bridge token symbols
  let symbols = bridgeTokens.map((token) => token.symbol);

  // 2. Get the list of Queries with possible swap instructions for origin chain
  // For that we pefrorm a static call to SynapseRouter on origin chain
  // This gets us the quotes from tokenIn to every bridge token (one quote per bridge token in the list)
  let originQueries = routerOrigin.getOriginAmountOut(
    tokenIn,
    symbols,
    amountIn
  );

  // 3. Get the list of Queries with possible swap instructions for destination chain
  // First, we form a list of "destiantion requests" by merging
  // list of token symbols with list of quotes obtained in step 2.
  let requests = symbols.map((value, index) => {
    let request: DestRequest = {
      symbol: value,
      amountIn: originQueries[index].minAmountOut,
    };
    return request;
  });
  // Then we perform a static call to SynapseRouter on destination chain
  // This gets us the quotes from every bridge token to tokenOut (one quote per bridge token in the list)
  // These quotes will take into account the fee for bridging the token to destination chain
  let destQueries = routerDest.getDestinationAmountOut(requests, tokenOut);

  // 4. Pick a pair of originQueries[i], destQueries[i] to pefrom the cross-chain swap
  // In this example we are picking the pair that yeilds the best overall quote
  let destQuery = maxBy(destQueries, (query) => query.minAmountOut);
  let selectedIndex = destQueries.indexOf(destQuery)
  let originQuery = originQueries[selectedIndex]

  // Now we apply user slippage and deadline settings
  originQuery = applyUserSettings(originQuery, userSettings)
  destQuery = applyUserSettings(destQuery, userSettings)

  // 5. Call SynapseRouter on origin chain to perform a swap
  let amountETH: BigInt;
  // 0xEeee address is used to represent native ETH
  if (tokenIn == "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE") {
    // If user selected "native ETH" as tokenIn, we would need to modify msg.value for the call
    amountETH = amountIn;
  } else {
    // If user selected an ERC-20 token as tokenIn, we would need to use msg.value=0
    amountETH = 0
    // We also need to check if user approved routerOrigin to spend `tokenIn`
    if (allowance(tokenIn, userOrigin, routerOrigin) < amountIn) {
      // Users needs to issue a token approval
      // tokenIn.approve(routerOrigin, amountIn)
    }
  }
  // Perform a call to Synapse Router with all the derevied parameters
  // Use previously determined msg.value for this call
  // (WETH wrapping is done by the Synapse Router)
  routerOrigin.bridge{value: amountETH}(
    userDest,
    destChainId,
    tokenIn,
    amountIn,
    originQuery,
    destQuery
  );
}
```
