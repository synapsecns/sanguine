import { Provider } from '@ethersproject/abstract-provider'
import invariant from 'tiny-invariant'
import { BigNumber } from '@ethersproject/bignumber'
import { PopulatedTransaction } from 'ethers'
import { AddressZero, Zero } from '@ethersproject/constants'
import { Interface } from '@ethersproject/abi'
import { Contract } from '@ethersproject/contracts'

import {
  handleNativeToken,
  ETH_NATIVE_TOKEN_ADDRESS,
} from './utils/handleNativeToken'
import { BigintIsh, CCTP_ROUTER_ADDRESS } from './constants'
import { SynapseRouter } from './synapseRouter'
import bridgeAbi from './abi/SynapseBridge.json'
import {
  Query,
  FeeConfig,
  PoolToken,
  SynapseCCTPRouterQuery,
  SynapseRouterQuery,
} from './utils/types'
import { SynapseCCTPRouter } from './SynapseCCTPRouter'
const ONE_WEEK_DEADLINE = BigNumber.from(Math.floor(Date.now() / 1000) + 604800) // one week in the future
const TEN_MIN_DEADLINE = BigNumber.from(Math.floor(Date.now() / 1000) + 600) // ten minutes in the future

type SynapseRouters = {
  [key: number]: SynapseRouter
}
type SynapseCCTPRouters = {
  [key: number]: SynapseCCTPRouter
}
type BridgeQuote = {
  feeAmount: BigNumber | undefined
  feeConfig: FeeConfig | undefined
  routerAddress: string | undefined
  maxAmountOut: BigNumber | undefined
  originQuery: Query | undefined
  destQuery: Query | undefined
}

/**
 * SynapseSDK class provides methods for interacting with Synapse protocol's routers
 * across different chains. It also provides methods for bridging tokens across these chains.
 *
 * @property {SynapseRouters} synapseRouters - Collection of SynapseRouters indexed by chain ID.
 * @property {SynapseCCTPRouters} synapseCCTPRouters - Collection of SynapseCCTPRouters indexed by chain ID.
 * @property {{[x: number]: Provider}} providers - Collection of Ethereum providers indexed by chain ID.
 */
class SynapseSDK {
  public synapseRouters: SynapseRouters
  public synapseCCTPRouters: SynapseCCTPRouters
  public providers: { [x: number]: Provider }
  public bridgeAbi: Interface = new Interface(bridgeAbi)
  public bridgeTokenCache: {
    [x: string]: { symbol: string; token: string }[]
  } = {}
  /**
   * Constructor for the SynapseSDK class.
   * It sets up the SynapseRouters and SynapseCCTPRouters for the specified chain IDs and providers.
   *
   * @param {number[]} chainIds - The IDs of the chains to initialize routers for.
   * @param {Provider[]} providers - The Ethereum providers for the respective chains.
   */
  constructor(chainIds: number[], providers: Provider[]) {
    invariant(
      chainIds.length === providers.length,
      `Amount of chains and providers does not equal`
    )
    this.synapseRouters = {}
    this.synapseCCTPRouters = {}
    this.providers = {}
    for (let i = 0; i < chainIds.length; i++) {
      this.synapseRouters[chainIds[i]] = new SynapseRouter(
        chainIds[i],
        providers[i]
      )
      this.providers[chainIds[i]] = providers[i]
      // check if the chain id is in the CCTP_ROUTER_ADDRESS object
      if (CCTP_ROUTER_ADDRESS.hasOwnProperty(chainIds[i])) {
        this.synapseCCTPRouters[chainIds[i]] = new SynapseCCTPRouter(
          chainIds[i],
          providers[i]
        )
      }
    }
  }

  /**
   * Fetches bridge tokens for a destination chain and output token.
   *
   * Checks the cache first, and fetches from the router if not cached. Filters invalid tokens and caches the result.
   *
   * @param destChainId - The destination chain ID.
   * @param tokenOut - The output token.
   * @param destRouter - The SynapseRouter or SynapseCCTPRouter to use.
   * @returns An array of { symbol, token } objects for valid bridge tokens.
   */
  public async getBridgeTokens(
    destChainId: number,
    tokenOut: string,
    destRouter: SynapseRouter | SynapseCCTPRouter
  ): Promise<{ symbol: string; token: string }[]> {
    // Check the cache first
    const cacheKey = `${destChainId}_${tokenOut}_${destRouter.routerContract.address}`
    if (this.bridgeTokenCache[cacheKey]) {
      return this.bridgeTokenCache[cacheKey]
    }

    // If not cached, get bridge tokens from the destination router
    const routerBridgeTokens =
      await destRouter.routerContract.getConnectedBridgeTokens(tokenOut)

    // Filter out invalid tokens
    const validBridgeTokens = routerBridgeTokens.filter(
      (token) => token.symbol && token.token !== AddressZero
    )

    // Store only the symbol and token fields
    const bridgeTokens = validBridgeTokens.map(({ symbol, token }) => ({
      symbol,
      token,
    }))

    // Cache the tokens for future use
    this.bridgeTokenCache[cacheKey] = bridgeTokens

    return bridgeTokens
  }

  /**
   * Fetches origin queries from either a SynapseRouter or SynapseCCTPRouter.
   *
   * @param router - The router to use (SynapseRouter or SynapseCCTPRouter)
   * @param tokenIn - The input token
   * @param tokenSymbols - The token symbols
   * @param amountIn - The input amount
   * @returns A promise that resolves to an array of Query objects
   * @throws Will throw an error if unable to fetch origin queries
   */
  public async getOriginQueries(
    router: SynapseRouter | SynapseCCTPRouter,
    tokenIn: string,
    tokenSymbols: string[],
    amountIn: BigintIsh
  ): Promise<Query[]> {
    try {
      if (router instanceof SynapseRouter) {
        const routerQueries = (await router.routerContract.getOriginAmountOut(
          tokenIn,
          tokenSymbols,
          amountIn
        )) as SynapseRouterQuery[]

        // Filter out 0 amounts and normalize the queries to the Query type
        return routerQueries
          .filter((query) => !query.minAmountOut.eq(0))
          .map((routerQuery) => ({
            swapAdapter: routerQuery.swapAdapter,
            tokenOut: routerQuery.tokenOut,
            minAmountOut: routerQuery.minAmountOut,
            deadline: routerQuery.deadline,
            rawParams: routerQuery.rawParams,
          }))
      } else {
        const routerQueries = (await router.routerContract.getOriginAmountOut(
          tokenIn,
          tokenSymbols,
          amountIn
        )) as SynapseCCTPRouterQuery[]

        // Filter out 0 amounts and normalize the queries to the Query type
        return routerQueries
          .filter((query) => !query.minAmountOut.eq(0))
          .map((routerQuery) => ({
            routerAdapter: routerQuery.routerAdapter,
            tokenOut: routerQuery.tokenOut,
            minAmountOut: routerQuery.minAmountOut,
            deadline: routerQuery.deadline,
            rawParams: routerQuery.rawParams,
          }))
      }
    } catch (error) {
      console.error('Failed to fetch origin queries', error)
      throw error
    }
  }

  /**
   * Fetches destination queries from either a SynapseRouter or SynapseCCTPRouter.
   *
   * @param router - The router to use (SynapseRouter or SynapseCCTPRouter)
   * @param requests - The requests with symbol and amount in.
   * @param tokenOut - The output token.
   * @returns A promise that resolves to an array of Query objects.
   * @throws Will throw an error if unable to fetch destination queries.
   */
  public async getDestinationQueries(
    router: SynapseRouter | SynapseCCTPRouter,
    requests: { symbol: string; amountIn: BigintIsh }[],
    tokenOut: string
  ): Promise<Query[]> {
    try {
      if (router instanceof SynapseRouter) {
        const routerQueries =
          (await router.routerContract.getDestinationAmountOut(
            requests,
            tokenOut
          )) as SynapseRouterQuery[]

        // Filter out 0 amounts and normalize the queries to the Query type
        return routerQueries
          .filter((query) => !query.minAmountOut.eq(0))
          .map((routerQuery) => ({
            ...routerQuery,
            swapAdapter: routerQuery.swapAdapter,
          }))
      } else {
        const routerQueries =
          (await router.routerContract.getDestinationAmountOut(
            requests,
            tokenOut
          )) as SynapseCCTPRouterQuery[]

        // Filter out 0 amounts and normalize the queries to the Query type
        return routerQueries
          .filter((query) => !query.minAmountOut.eq(0))
          .map((routerQuery) => ({
            ...routerQuery,
            swapAdapter: routerQuery.routerAdapter,
          }))
      }
    } catch (error) {
      console.error('Failed to fetch destination queries', error)
      throw error
    }
  }

  /**
   * Finds the best query pair from origin and destination queries.
   *
   * @param destQueries - The destination queries.
   * @param originQueries - The origin queries.
   * @param bridgeTokens - The bridge tokens.
   * @returns The best origin query, destination query, and bridge token.
   * @throws Will throw an error if no best queries are found.
   */
  public findBestQuery(
    destQueries: Query[],
    originQueries: Query[],
    bridgeTokens: { symbol: string; token: string }[]
  ): [Query, Query, { symbol: string; token: string }] {
    // avoid naive nested looped find best Query
    let maxAmountOut: BigNumber = BigNumber.from(0)
    let bestDestQuery: Query | null = null
    let bestOriginQuery: Query | null = null
    let bestBridgeToken: { symbol: string; token: string } | null = null

    for (let i = 0; i < destQueries.length; i++) {
      const destQuery = destQueries[i]
      if (!destQuery.minAmountOut.gt(maxAmountOut)) {
        continue
      }

      const originQuery = originQueries[i]
      const bridgeToken = bridgeTokens[i]

      maxAmountOut = destQuery.minAmountOut
      bestDestQuery = destQuery
      bestOriginQuery = originQuery
      bestBridgeToken = bridgeToken
    }

    if (!bestDestQuery || !bestOriginQuery || !bestBridgeToken) {
      throw new Error('No best queries found')
    }

    return [bestOriginQuery, bestDestQuery, bestBridgeToken]
  }

  /**
   * Finalizes a quote by getting fee data and setting default deadlines.
   *
   * @param bestQuery - The best origin query, destination query and bridge token.
   * @param router - The router to use (SynapseRouter or SynapseCCTPRouter).
   * @param deadline - The deadline to use (default 10 mins).
   * @param isCCTP - Whether the router is a SynapseCCTPRouter (default false).
   * @returns The finalized quote with fee data and deadlines.
   */
  public async finalizeQuote(
    bestQuery: [Query, Query, { symbol: string; token: string }],
    router: SynapseRouter | SynapseCCTPRouter,
    deadline?: BigNumber,
    isCCTP = false
  ): Promise<{
    feeAmount: BigNumber
    feeConfig: FeeConfig
    routerAddress: string
    maxAmountOut: BigNumber
    originQuery: Query
    destQuery: Query
  }> {
    const [originQuery, destQuery, bestBridgeToken] = bestQuery

    let formattedOriginQuery: Query
    let formattedDestQuery: Query

    // Set default deadlines
    if ((originQuery as SynapseCCTPRouterQuery).routerAdapter) {
      formattedOriginQuery = { ...(originQuery as SynapseCCTPRouterQuery) }
    } else {
      formattedOriginQuery = { ...(originQuery as SynapseRouterQuery) }
    }
    formattedOriginQuery.deadline = deadline ?? TEN_MIN_DEADLINE

    let isSwap = false
    if ((destQuery as SynapseCCTPRouterQuery).routerAdapter) {
      formattedDestQuery = { ...(destQuery as SynapseCCTPRouterQuery) }
      isSwap = formattedDestQuery.routerAdapter !== AddressZero
    } else {
      formattedDestQuery = { ...(destQuery as SynapseRouterQuery) }
    }
    formattedDestQuery.deadline = ONE_WEEK_DEADLINE

    let feeAmount!: BigNumber
    let feeConfig!: FeeConfig

    // Get fee data from the appropriate router
    if (isCCTP) {
      const cctpRouter = router as SynapseCCTPRouter

      feeAmount = await cctpRouter.routerContract.calculateFeeAmount(
        bestBridgeToken.token,
        formattedOriginQuery.minAmountOut,
        isSwap
      )

      const [relayerFee, minBaseFee, , maxFee] =
        await cctpRouter.routerContract.feeStructures(bestBridgeToken.token)
      feeConfig = {
        bridgeFee: relayerFee,
        minFee: minBaseFee,
        maxFee,
      }
    } else {
      const synapseRouter = router as SynapseRouter
      feeAmount = await synapseRouter.routerContract.calculateBridgeFee(
        bestBridgeToken.token,
        formattedOriginQuery.minAmountOut
      )
      feeConfig = await synapseRouter.routerContract.fee(bestBridgeToken.token)
    }

    return {
      feeAmount,
      feeConfig,
      routerAddress: router.routerContract.address,
      maxAmountOut: formattedDestQuery.minAmountOut,
      originQuery: formattedOriginQuery,
      destQuery: formattedDestQuery,
    }
  }

  /**
   * This method tries to fetch the best quote from either the Synapse Router or SynapseCCTP Router.
   * It first handles the native token, then fetches the best quote for both types of routers.
   * If the router addresses are valid for CCTP, it will fetch the quote from the CCTP routers, otherwise it will resolve to undefined.
   * It waits for both types of quotes, then determines the best one by comparing the maximum output amount.
   * If no best quote can be found, it will throw an error.
   *
   * @param originChainId - The ID of the original chain.
   * @param destChainId - The ID of the destination chain.
   * @param tokenIn - The input token.
   * @param tokenOut - The output token.
   * @param amountIn - The amount of input token.
   * @param deadline - The transaction deadline, optional.
   *
   * @returns - A promise that resolves to the best bridge quote, or undefined if no route is found.
   *
   * @throws - Will throw an error if no best quote could be determined.
   */
  public async bridgeQuote(
    originChainId: number,
    destChainId: number,
    tokenIn: string,
    tokenOut: string,
    amountIn: BigintIsh,
    deadline?: BigNumber
  ): Promise<BridgeQuote | undefined> {
    tokenOut = handleNativeToken(tokenOut)
    tokenIn = handleNativeToken(tokenIn)

    const originSynapseRouter = this.synapseRouters[originChainId]
    const destSynapseRouter = this.synapseRouters[destChainId]

    const synapseQuotePromise = this.calculateBestQuote(
      originSynapseRouter,
      destSynapseRouter,
      destChainId,
      tokenIn,
      tokenOut,
      amountIn,
      deadline
    ).catch((error) => {
      console.error('Error with synapseQuotePromise', error)
      return undefined
    })

    let cctpQuotePromise: Promise<BridgeQuote | undefined>
    if (
      CCTP_ROUTER_ADDRESS[originChainId] &&
      CCTP_ROUTER_ADDRESS[destChainId]
    ) {
      const originCCTPRouter = this.synapseCCTPRouters[originChainId]
      const destCCTPRouter = this.synapseCCTPRouters[destChainId]
      cctpQuotePromise = this.calculateBestQuote(
        originCCTPRouter,
        destCCTPRouter,
        destChainId,
        tokenIn,
        tokenOut,
        amountIn,
        deadline
      ).catch((error) => {
        console.error('Error with cctpQuotePromise', error)
        return undefined
      })
    } else {
      cctpQuotePromise = Promise.resolve(undefined)
    }

    const [synapseQuote, cctpQuote] = await Promise.all([
      synapseQuotePromise,
      cctpQuotePromise,
    ])

    const bestQuote = [synapseQuote, cctpQuote].reduce(
      (prev, current) =>
        !prev ||
        (current &&
          current.maxAmountOut &&
          prev.maxAmountOut &&
          current.maxAmountOut.gt(prev.maxAmountOut))
          ? current
          : prev,
      undefined
    )

    if (!bestQuote) {
      throw new Error('No route found')
    }

    return bestQuote
  }

  /**
   * This method calculates the best bridge quote by using either the SynapseRouter or SynapseCCTPRouter.
   * For each set of bridge tokens, it fetches queries from the origin and destination routers and finds the best query.
   * After that, it finalizes the quote and checks if this quote is better than the previous best.
   *
   * @param originRouter - The origin router (SynapseRouter or SynapseCCTPRouter or undefined).
   * @param destRouter - The destination router (SynapseRouter or SynapseCCTPRouter or undefined).
   * @param destChainId - The ID of the destination chain.
   * @param tokenIn - The input token.
   * @param tokenOut - The output token.
   * @param amountIn - The amount of input token.
   * @param deadline - The transaction deadline, optional.
   *
   * @returns - A promise that resolves to the best bridge quote, or undefined if no best quote could be determined.
   */
  private async calculateBestQuote(
    originRouter: SynapseRouter | SynapseCCTPRouter | undefined,
    destRouter: SynapseRouter | SynapseCCTPRouter | undefined,
    destChainId: number,
    tokenIn: string,
    tokenOut: string,
    amountIn: BigintIsh,
    deadline?: BigNumber
  ): Promise<BridgeQuote | undefined> {
    if (!originRouter || !destRouter) {
      return
    }

    let bestQuote: BridgeQuote | undefined

    // Getting bridge tokens from cache or fetch from destination router
    const bridgeTokensArray = await this.getBridgeTokens(
      destChainId,
      tokenOut,
      destRouter
    )

    // Iterate through each array of bridge tokens
    for (const bridgeTokens of bridgeTokensArray) {
      try {
        // Fetching queries from origin router
        const originQueries = await this.getOriginQueries(
          originRouter,
          tokenIn,
          [bridgeTokens.symbol],
          amountIn
        )
        if (!originQueries.length) {
          continue
        } // Skip if no origin queries for these bridge tokens

        // Building request for destination queries
        const requests = originQueries.map((query) => ({
          symbol: bridgeTokens.symbol,
          amountIn: query.minAmountOut,
        }))
        // Fetching queries from destination router
        const destQueries = await this.getDestinationQueries(
          destRouter,
          requests,
          tokenOut
        )
        if (!destQueries.length) {
          continue
        } // Skip if no destination queries for these requests

        // Finding the best query
        const bestQuery = this.findBestQuery(destQueries, originQueries, [
          bridgeTokens,
        ])

        // Finalizing quote
        const quote = await this.finalizeQuote(
          bestQuery,
          destRouter,
          deadline,
          destRouter instanceof SynapseCCTPRouter
        )
        // Check if this quote is better than previous best
        if (
          !bestQuote ||
          quote.maxAmountOut.gt(bestQuote.maxAmountOut ?? BigNumber.from(0))
        ) {
          bestQuote = quote
        }
      } catch (error) {
        console.error(
          `Error when trying to calculate the best quote with bridge tokens ${bridgeTokens.symbol}:`,
          error
        )
      }
    }

    return bestQuote
  }

  /**
   * Executes a bridge operation between two different chains. Depending on the origin router address, the operation
   * will use either a SynapseRouter or a SynapseCCTPRouter. This function creates a populated transaction ready
   * to be signed and sent to the origin chain.
   *
   * @param to - The recipient address of the bridged tokens.
   * @param originRouterAddress - The address of the origin router.
   * @param originChainId - The ID of the origin chain.
   * @param destChainId - The ID of the destination chain.
   * @param token - The token to bridge.
   * @param amount - The amount of token to bridge.
   * @param originQuery - The query for the origin chain.
   * @param destQuery - The query for the destination chain.
   *
   * @returns A promise that resolves to a populated transaction object which can be used to send the transaction.
   *
   * @throws Will throw an error if there's an issue with the bridge operation.
   */
  public async bridge(
    to: string,
    originRouterAddress: string,
    originChainId: number,
    destChainId: number,
    token: string,
    amount: BigintIsh,
    originQuery: Query,
    destQuery: Query
  ): Promise<PopulatedTransaction> {
    token = handleNativeToken(token)

    // Create new query objects and check the correct subtype to avoid type errors
    let bridgeOriginQuery: SynapseRouterQuery | SynapseCCTPRouterQuery
    if ((originQuery as SynapseCCTPRouterQuery).routerAdapter) {
      bridgeOriginQuery = { ...(originQuery as SynapseCCTPRouterQuery) }
    } else {
      bridgeOriginQuery = { ...(originQuery as SynapseRouterQuery) }
    }

    let bridgeDestQuery: SynapseRouterQuery | SynapseCCTPRouterQuery
    if ((destQuery as SynapseCCTPRouterQuery).routerAdapter) {
      bridgeDestQuery = { ...(destQuery as SynapseCCTPRouterQuery) }
    } else {
      bridgeDestQuery = { ...(destQuery as SynapseRouterQuery) }
    }

    const isCCTP =
      this.synapseCCTPRouters[originChainId] &&
      originRouterAddress.toLowerCase() ===
        this.synapseCCTPRouters[
          originChainId
        ].routerContract.address.toLowerCase()

    if (isCCTP) {
      // Call CCTP router bridge method
      return this.synapseCCTPRouters[
        originChainId
      ].routerContract.populateTransaction.bridge(
        to,
        destChainId,
        token,
        amount,
        { ...(bridgeOriginQuery as SynapseCCTPRouterQuery) },
        { ...(bridgeDestQuery as SynapseCCTPRouterQuery) }
      )
    } else {
      // Call Synapse router bridge method
      return this.synapseRouters[
        originChainId
      ].routerContract.populateTransaction.bridge(
        to,
        destChainId,
        token,
        amount,
        { ...(bridgeOriginQuery as SynapseRouterQuery) },
        { ...(bridgeDestQuery as SynapseRouterQuery) }
      )
    }
  }

  /**
   * Gets a swap quote from a Synapse Router.
   *
   * @param chainId The chain ID
   * @param tokenIn The input token
   * @param tokenOut The output token
   * @param amountIn The input amount
   * @param deadline The deadline
   * @returns The swap quote (query, max amount out, and router address)
   */
  public async swapQuote(
    chainId: number,
    tokenIn: string,
    tokenOut: string,
    amountIn: BigintIsh,
    deadline?: BigNumber
  ): Promise<{
    routerAddress: string | undefined
    maxAmountOut: BigNumber | undefined
    query: SynapseRouterQuery | undefined
  }> {
    tokenOut = handleNativeToken(tokenOut)
    tokenIn = handleNativeToken(tokenIn)

    const router: SynapseRouter = this.synapseRouters[chainId]
    const routerAddress = router.routerContract.address

    const rawQuery = await router.routerContract.getAmountOut(
      tokenIn,
      tokenOut,
      amountIn
    )

    // Check if call was unsuccessful.
    if (rawQuery?.length !== 5 || rawQuery.minAmountOut.isZero()) {
      throw Error('No queries found for this route')
    }

    const query = { ...(rawQuery as SynapseRouterQuery) }
    query.deadline = deadline ?? TEN_MIN_DEADLINE
    const maxAmountOut = query.minAmountOut

    return {
      routerAddress,
      maxAmountOut,
      query,
    }
  }

  /**
   * Performs a swap through a Synapse Router.
   *
   * @param chainId The chain ID
   * @param to The recipient address
   * @param token The token to swap
   * @param amount The swap amount
   * @param query The swap quote query
   * @returns A populated transaction to perform the swap
   */
  public async swap(
    chainId: number,
    to: string,
    token: string,
    amount: BigintIsh,
    query: SynapseRouterQuery
  ): Promise<PopulatedTransaction> {
    token = handleNativeToken(token)
    const originRouter: SynapseRouter = this.synapseRouters[chainId]
    return originRouter.routerContract.populateTransaction.swap(
      to,
      token,
      amount,
      query
    )
  }

  /**
   * Gets the chain gas amount for the Synapse bridge.
   *
   * @param chainId The chain ID
   * @returns The chain gas amount
   */
  public async getBridgeGas(chainId: number): Promise<BigintIsh> {
    const router: SynapseRouter = this.synapseRouters[chainId]
    const bridgeAddress = await router.routerContract.synapseBridge()
    const bridgeContract = new Contract(
      bridgeAddress,
      this.bridgeAbi,
      this.providers[chainId]
    )
    return bridgeContract.chainGasAmount()
  }

  /**
   * Gets pool tokens for a pool address.
   *
   * @param chainId The chain ID
   * @param poolAddress The pool address
   * @returns The pool tokens
   */
  public async getPoolTokens(
    chainId: number,
    poolAddress: string
  ): Promise<PoolToken[]> {
    const router: SynapseRouter = this.synapseRouters[chainId]
    const poolTokens = await router.routerContract.poolTokens(poolAddress)
    return poolTokens.map((token) => {
      return { token: token.token, isWeth: token?.isWeth }
    })
  }

  /**
   * Gets info for a pool (number of tokens and LP token).
   *
   * @param chainId The chain ID
   * @param poolAddress The pool address
   * @returns The pool info (number of tokens and LP token)
   */
  public async getPoolInfo(
    chainId: number,
    poolAddress: string
  ): Promise<{ tokens: BigNumber | undefined; lpToken: string | undefined }> {
    const router: SynapseRouter = this.synapseRouters[chainId]
    const poolInfo = await router.routerContract.poolInfo(poolAddress)
    return { tokens: poolInfo?.[0], lpToken: poolInfo?.[1] }
  }

  /**
   * Gets all pools for a chain ID.
   *
   * @param chainId The chain ID
   * @returns An array of all pools (address, tokens, LP token)
   */
  public async getAllPools(chainId: number): Promise<
    {
      poolAddress: string | undefined
      tokens: PoolToken[] | undefined
      lpToken: string | undefined
    }[]
  > {
    const router: SynapseRouter = this.synapseRouters[chainId]
    const pools = await router.routerContract.allPools()
    const res = pools.map((pool) => {
      return {
        poolAddress: pool?.pool,
        tokens: pool?.tokens.map((token) => {
          return { token: token.token, isWeth: token?.isWeth }
        }),
        lpToken: pool?.lpToken,
      }
    })
    return res
  }

  /**
   * Calculates the amount required to add liquidity for amounts of each token.
   *
   * @param chainId The chain ID
   * @param poolAddress The pool address
   * @param amounts The amounts of each token to add
   * @returns The amount of LP tokens needed and router address
   */
  public async calculateAddLiquidity(
    chainId: number,
    poolAddress: string,
    amounts: Record<string, BigNumber>
  ): Promise<{ amount: BigNumber; routerAddress: string }> {
    const router: SynapseRouter = this.synapseRouters[chainId]
    const poolTokens = await router.routerContract.poolTokens(poolAddress)
    const amountArr: BigNumber[] = []
    poolTokens.map((token) => {
      amountArr.push(amounts[token.token] ?? Zero)
    })
    if (amountArr.filter((amount) => !amount.isZero()).length === 0) {
      return { amount: Zero, routerAddress: router.routerContract.address }
    }
    return {
      amount: await router.routerContract.calculateAddLiquidity(
        poolAddress,
        amountArr
      ),
      routerAddress: router.routerContract.address,
    }
  }

  /**
   * Calculates the amounts received when removing liquidity.
   *
   * @param chainId The chain ID
   * @param poolAddress The pool address
   * @param amount The amount of LP tokens to remove
   * @returns The amounts of each token received and router address
   */
  public async calculateRemoveLiquidity(
    chainId: number,
    poolAddress: string,
    amount: BigNumber
  ): Promise<{
    amounts: Array<{ value: BigNumber; index: number }>
    routerAddress: string
  }> {
    const router: SynapseRouter = this.synapseRouters[chainId]
    const amounts = await router.routerContract.calculateRemoveLiquidity(
      poolAddress,
      amount
    )
    const amountsOut: Array<{ value: BigNumber; index: number }> = amounts.map(
      (respAmount, index) => ({
        value: respAmount,
        index,
      })
    )

    return {
      amounts: amountsOut,
      routerAddress: router.routerContract.address,
    }
  }

  /**
   * Calculates the amount of one token received when removing liquidity.
   *
   * @param chainId The chain ID
   * @param poolAddress The pool address
   * @param amount The amount of LP tokens to remove
   * @param poolIndex The index of the token to receive
   * @returns The amount received and router address
   */
  public async calculateRemoveLiquidityOne(
    chainId: number,
    poolAddress: string,
    amount: BigNumber,
    poolIndex: number
  ): Promise<{
    amount: { value: BigNumber; index: number }
    routerAddress: string
  }> {
    const router: SynapseRouter = this.synapseRouters[chainId]

    const outAmount = await router.routerContract.calculateWithdrawOneToken(
      poolAddress,
      amount,
      poolIndex
    )

    return {
      amount: { value: outAmount, index: poolIndex },
      routerAddress: router.routerContract.address,
    }
  }
}

export { SynapseSDK, ETH_NATIVE_TOKEN_ADDRESS, Query, PoolToken }
