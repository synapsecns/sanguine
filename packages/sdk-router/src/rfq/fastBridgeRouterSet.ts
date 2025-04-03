import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber } from '@ethersproject/bignumber'
import { Zero } from '@ethersproject/constants'
import { BigNumberish } from 'ethers'
import NodeCache from 'node-cache'
import invariant from 'tiny-invariant'

import {
  FAST_BRIDGE_INTERCEPTOR_ADDRESS_MAP,
  FAST_BRIDGE_ROUTER_ADDRESS_MAP,
  MEDIAN_TIME_RFQ,
} from '../constants'
import {
  BridgeRoute,
  FeeConfig,
  Query,
  SynapseModule,
  SynapseModuleSet,
  createNoSwapQuery,
  applySlippageToQuery,
  BridgeTokenCandidate,
  BridgeRouteV2,
  GetBridgeTokenCandidatesParameters,
  GetBridgeRouteV2Parameters,
} from '../module'
import { FastBridgeRouter } from './fastBridgeRouter'
import { ChainProvider } from '../router'
import { applySlippage, encodeZapData, USER_SIMULATED_ADDRESS } from '../swap'
import {
  calculateDeadline,
  ONE_HOUR,
  TEN_MINUTES,
  isSameAddress,
  logger,
} from '../utils'
import { getAllQuotes } from './api'
import { FastBridgeQuote, applyQuote } from './quote'
import { marshallTicker } from './ticker'
import {
  IFastBridge,
  IFastBridgeInterceptor,
} from '../typechain/FastBridgeInterceptor'

export class FastBridgeRouterSet extends SynapseModuleSet {
  static readonly MAX_QUOTE_AGE_MILLISECONDS = 5 * 60 * 1000 // 5 minutes
  static readonly ALL_QUOTES_CACHE_TTL = 10 // 10 seconds cache for getAllQuotes results

  public readonly moduleName = 'SynapseRFQ'
  public readonly allEvents = ['BridgeRequestedEvent', 'BridgeRelayedEvent']
  public readonly isBridgeV2Supported = true

  public routers: {
    [chainId: number]: FastBridgeRouter
  }
  public providers: {
    [chainId: number]: Provider
  }

  private quotesCache: NodeCache

  constructor(chains: ChainProvider[]) {
    super()
    this.routers = {}
    this.providers = {}
    this.quotesCache = new NodeCache({
      stdTTL: FastBridgeRouterSet.ALL_QUOTES_CACHE_TTL,
    })
    chains.forEach(({ chainId, provider }) => {
      const address = FAST_BRIDGE_ROUTER_ADDRESS_MAP[chainId]
      const interceptor = FAST_BRIDGE_INTERCEPTOR_ADDRESS_MAP[chainId]
      // Skip chains without a FastBridgeRouter address
      if (address) {
        this.routers[chainId] = new FastBridgeRouter(
          chainId,
          provider,
          address,
          interceptor
        )
        this.providers[chainId] = provider
      }
    })
  }

  /**
   * @inheritdoc SynapseModuleSet.getModule
   */
  public getModule(chainId: number): SynapseModule | undefined {
    return this.routers[chainId]
  }

  /**
   * @inheritdoc SynapseModuleSet.getOriginAmountOut
   */
  public getEstimatedTime(chainId: number): number {
    const medianTime = MEDIAN_TIME_RFQ[chainId as keyof typeof MEDIAN_TIME_RFQ]
    invariant(medianTime, `No estimated time for chain ${chainId}`)
    return medianTime
  }

  /**
   * @inheritdoc SynapseModuleSet.getGasDropAmount
   */
  public async getGasDropAmount(): Promise<BigNumber> {
    return Zero
  }

  public async getBridgeTokenCandidates({
    fromChainId,
    toChainId,
    toToken,
  }: GetBridgeTokenCandidatesParameters): Promise<BridgeTokenCandidate[]> {
    if (!this.getModule(fromChainId) || !this.getModule(toChainId)) {
      return []
    }
    const quotes = await this.getQuotes(fromChainId, toChainId, toToken)
    // Filter out duplicates of the bridge token
    return Array.from(
      new Map(
        quotes.map((quote) => [marshallTicker(quote.ticker), quote])
      ).values()
    ).map((quote) => ({
      originChainId: fromChainId,
      destChainId: toChainId,
      originToken: quote.ticker.originToken.token,
      destToken: quote.ticker.destToken.token,
    }))
  }

  public async getBridgeRouteV2({
    originSwapRoute,
    bridgeToken,
    toToken,
    fromSender,
    toRecipient,
    slippage,
    allowMultipleTxs,
  }: GetBridgeRouteV2Parameters): Promise<BridgeRouteV2 | undefined> {
    if (
      !this.getModule(bridgeToken.originChainId) ||
      !this.getModule(bridgeToken.destChainId)
    ) {
      return undefined
    }
    if (!allowMultipleTxs && !isSameAddress(bridgeToken.destToken, toToken)) {
      return undefined
    }
    const originChainId = bridgeToken.originChainId
    const protocolFeeRate = await this.getFastBridgeRouter(
      originChainId
    ).getProtocolFeeRate()
    const bridgedAmount = this.applyProtocolFeeRate(
      originSwapRoute.expectedToAmount,
      protocolFeeRate
    )
    const quotes = (
      await this.getQuotes(
        originChainId,
        bridgeToken.destChainId,
        bridgeToken.destToken
      )
    ).filter((quote) =>
      isSameAddress(quote.ticker.originToken.token, bridgeToken.originToken)
    )
    const expectedToAmount = quotes
      .map((quote) => applyQuote(quote, bridgedAmount))
      .reduce((a, b) => (a.gt(b) ? a : b), Zero)
    if (expectedToAmount.isZero()) {
      return undefined
    }
    // With no slippage or no swap on origin, the minToAmount is the same as expectedToAmount.
    const minToAmount =
      !slippage ||
      originSwapRoute.expectedToAmount.eq(originSwapRoute.minToAmount)
        ? expectedToAmount
        : applySlippage(expectedToAmount, slippage)
    const route: BridgeRouteV2 = {
      bridgeToken,
      toToken: bridgeToken.destToken,
      expectedToAmount,
      minToAmount,
      zapData: await this.getBridgeZapData(
        bridgeToken,
        originSwapRoute.expectedToAmount,
        expectedToAmount,
        fromSender,
        toRecipient
      ),
    }
    return route
  }

  /**
   * @inheritdoc SynapseModuleSet.getBridgeRoutes
   */
  public async getBridgeRoutes(
    originChainId: number,
    destChainId: number,
    tokenIn: string,
    tokenOut: string,
    amountIn: BigNumberish,
    originUserAddress?: string
  ): Promise<BridgeRoute[]> {
    // Check that Routers exist on both chains
    if (!this.getModule(originChainId) || !this.getModule(destChainId)) {
      return []
    }
    // Get all quotes that result in the final token
    const allQuotes: FastBridgeQuote[] = await this.getQuotes(
      originChainId,
      destChainId,
      tokenOut
    )
    // Get queries for swaps on the origin chain into the "RFQ-supported token"
    const filteredQuotes = await this.filterOriginQuotes(
      originChainId,
      tokenIn,
      amountIn,
      allQuotes
    )
    const protocolFeeRate = await this.getFastBridgeRouter(
      originChainId
    ).getProtocolFeeRate()
    return filteredQuotes
      .map(({ quote, originQuery }) => ({
        quote,
        originQuery,
        // Apply quote to the proceeds of the origin swap with protocol fee applied
        // TODO: handle optional gas airdrop pricing
        destAmountOut: applyQuote(
          quote,
          this.applyProtocolFeeRate(originQuery.minAmountOut, protocolFeeRate)
        ),
      }))
      .filter(({ destAmountOut }) => destAmountOut.gt(0))
      .map(({ quote, originQuery, destAmountOut }) => ({
        originChainId,
        destChainId,
        bridgeToken: {
          symbol: marshallTicker(quote.ticker),
          token: quote.ticker.destToken.token,
        },
        originQuery,
        destQuery: FastBridgeRouterSet.createRFQDestQuery(
          tokenOut,
          destAmountOut,
          originUserAddress
        ),
        bridgeModuleName: this.moduleName,
      }))
  }

  /**
   * @inheritdoc SynapseModuleSet.getFeeData
   */
  public async getFeeData(bridgeRoute: BridgeRoute): Promise<{
    feeAmount: BigNumber
    feeConfig: FeeConfig
  }> {
    // TODO: do we actually need to return non-zero alues here?
    // Origin Out vs Dest Out is the effective fee if amountOut is within 1% of amountIn.
    // Otherwise origin and destination tokens are different, so the SDK has no means to determine the effective fee.
    const amountIn = bridgeRoute.originQuery.minAmountOut
    const amountOut = bridgeRoute.destQuery.minAmountOut
    const feeAmount =
      amountOut.gte(amountIn.mul(99).div(100)) && amountOut.lte(amountIn)
        ? amountIn.sub(amountOut)
        : Zero
    return {
      feeAmount,
      feeConfig: {
        bridgeFee: 0,
        minFee: BigNumber.from(0),
        maxFee: BigNumber.from(0),
      },
    }
  }

  /**
   * @inheritdoc SynapseModuleSet.getDefaultPeriods
   */
  public getDefaultPeriods(): {
    originPeriod: number
    destPeriod: number
  } {
    return {
      originPeriod: TEN_MINUTES,
      destPeriod: 2 * ONE_HOUR,
    }
  }

  /**
   * @inheritdoc SynapseModuleSet.applySlippage
   */
  public applySlippage(
    originQueryPrecise: Query,
    destQueryPrecise: Query,
    slipNumerator: number,
    slipDenominator: number
  ): { originQuery: Query; destQuery: Query } {
    // Max slippage for origin swap is 5% of the fixed fee
    // Relayer is using a 10% buffer for the fixed fee, so if origin swap slippage
    // is under 5% of the fixed fee, the relayer will still honor the quote.
    let maxOriginSlippage = originQueryPrecise.minAmountOut
      .sub(destQueryPrecise.minAmountOut)
      .div(20)
    // TODO: figure out a better way to handle destAmount > originAmount
    if (maxOriginSlippage.isNegative()) {
      maxOriginSlippage = BigNumber.from(0)
    }
    const originQuery = applySlippageToQuery(
      originQueryPrecise,
      slipNumerator,
      slipDenominator
    )
    if (
      originQuery.minAmountOut
        .add(maxOriginSlippage)
        .lt(originQueryPrecise.minAmountOut)
    ) {
      originQuery.minAmountOut =
        originQueryPrecise.minAmountOut.sub(maxOriginSlippage)
    }
    // Never modify the dest query, as the exact amount from it will always be used by the Relayer
    // So applying slippage there will only reduce the user proceeds on the destination chain
    return {
      originQuery,
      destQuery: destQueryPrecise,
    }
  }

  /**
   * Returns the existing FastBridgeRouter instance for the given chain.
   *
   * @throws Will throw an error if FastBridgeRouter is not deployed on the given chain.
   */
  public getFastBridgeRouter(chainId: number): FastBridgeRouter {
    return this.getExistingModule(chainId) as FastBridgeRouter
  }

  /**
   * Returns the address of the FastBridge contract for the given chain.
   */
  public async getFastBridgeAddress(chainId: number): Promise<string> {
    const fastBridgeContract = await this.getFastBridgeRouter(
      chainId
    ).getFastBridgeContract()
    return fastBridgeContract.address
  }

  /**
   * Applies the protocol fee to the amount.
   *
   * @returns The amount after the fee.
   */
  public applyProtocolFeeRate(
    amount: BigNumber,
    protocolFeeRate: BigNumber
  ): BigNumber {
    const protocolFee = amount.mul(protocolFeeRate).div(1_000_000)
    return amount.sub(protocolFee)
  }

  /**
   * Filters the list of quotes to only include those that can be used for given amount of input token.
   * For every filtered quote, the origin query is returned with the information for tokenIn -> RFQ token swaps.
   */
  private async filterOriginQuotes(
    originChainId: number,
    tokenIn: string,
    amountIn: BigNumberish,
    allQuotes: FastBridgeQuote[]
  ): Promise<{ quote: FastBridgeQuote; originQuery: Query }[]> {
    // Get queries for swaps on the origin chain into the "RFQ-supported token"
    const originQueries = await this.getFastBridgeRouter(
      originChainId
    ).getOriginAmountOut(
      tokenIn,
      allQuotes.map((quote) => quote.ticker.originToken.token),
      amountIn
    )
    // Note: allQuotes.length === originQueries.length
    // Zip the quotes and queries together, filter out "no path found" queries
    return allQuotes
      .map((quote, index) => ({
        quote,
        originQuery: originQueries[index],
      }))
      .filter(({ originQuery }) => originQuery.minAmountOut.gt(0))
  }

  /**
   * Retrieves all quotes with caching.
   *
   * @returns A promise that resolves to all available quotes.
   */
  private async getCachedAllQuotes(): Promise<FastBridgeQuote[]> {
    const cacheKey = 'all_quotes'
    const cachedQuotes = this.quotesCache.get<FastBridgeQuote[]>(cacheKey)
    if (cachedQuotes) {
      return cachedQuotes
    }
    const allQuotes = await getAllQuotes()
    this.quotesCache.set(cacheKey, allQuotes)
    return allQuotes
  }

  /**
   * Get the list of quotes between two chains for a given final token.
   *
   * @param originChainId - The ID of the origin chain.
   * @param destChainId - The ID of the destination chain.
   * @param tokenOut - The final token of the cross-chain swap.
   * @returns A promise that resolves to the list of supported tickers.
   */
  private async getQuotes(
    originChainId: number,
    destChainId: number,
    tokenOut?: string
  ): Promise<FastBridgeQuote[]> {
    const allQuotes = await this.getCachedAllQuotes()
    const originFB = await this.getFastBridgeAddress(originChainId)
    const destFB = await this.getFastBridgeAddress(destChainId)
    // Apply optional filtering by the final token
    return allQuotes
      .filter(
        (quote) =>
          quote.ticker.originToken.chainId === originChainId &&
          quote.ticker.destToken.chainId === destChainId &&
          (!tokenOut || isSameAddress(quote.ticker.destToken.token, tokenOut))
      )
      .filter(
        (quote) =>
          isSameAddress(quote.originFastBridge, originFB) &&
          isSameAddress(quote.destFastBridge, destFB)
      )
      .filter((quote) => {
        const age = Date.now() - quote.updatedAt
        return 0 <= age && age < FastBridgeRouterSet.MAX_QUOTE_AGE_MILLISECONDS
      })
  }

  public static createRFQDestQuery(
    tokenOut: string,
    amountOut: BigNumber,
    originUserAddress?: string
  ): Query {
    // On-chain swaps are not supported for RFQ on the destination chain
    const destQuery = createNoSwapQuery(tokenOut, amountOut)
    // Don't modify the Query if user address is undefined
    if (!originUserAddress) {
      return destQuery
    }
    // Make sure the rebate flag is always included if user address is defined.
    // 0x00 is a single byte that indicates the rebate flag is turned off.
    // Concatenate the originUserAddress (without 0x prefix) to the end of the rawParams.
    destQuery.rawParams = '0x00' + originUserAddress.slice(2)
    return destQuery
  }

  private async getBridgeZapData(
    bridgeToken: BridgeTokenCandidate,
    fromAmount: BigNumber,
    expectedToAmount: BigNumber,
    fromSender?: string,
    toRecipient?: string
  ): Promise<string | undefined> {
    if (
      expectedToAmount.isZero() ||
      !fromSender ||
      !toRecipient ||
      isSameAddress(fromSender, USER_SIMULATED_ADDRESS) ||
      isSameAddress(toRecipient, USER_SIMULATED_ADDRESS)
    ) {
      return undefined
    }
    const bridgeParams: IFastBridge.BridgeParamsStruct = {
      dstChainId: bridgeToken.destChainId,
      sender: fromSender,
      to: toRecipient,
      originToken: bridgeToken.originToken,
      destToken: bridgeToken.destToken,
      // Will be set in encodeZapData below
      originAmount: 0,
      destAmount: expectedToAmount,
      sendChainGas: false,
      deadline: calculateDeadline(this.getDefaultPeriods().destPeriod),
    }
    const fastBridge = await this.getFastBridgeAddress(
      bridgeToken.originChainId
    )
    const fastBridgeInterceptor = this.getFastBridgeRouter(
      bridgeToken.originChainId
    ).interceptorContract
    if (!fastBridgeInterceptor) {
      logger.error(
        `FastBridgeInterceptor not found for chainId ${bridgeToken.originChainId}`
      )
      return undefined
    }
    const interceptorParams: IFastBridgeInterceptor.InterceptorParamsStruct = {
      fastBridge,
      quoteOriginAmount: fromAmount,
    }
    const fbiCalldata = (
      await fastBridgeInterceptor.populateTransaction.bridgeWithInterception(
        bridgeParams,
        interceptorParams
      )
    ).data
    return encodeZapData({
      target: fastBridgeInterceptor.address,
      payload: fbiCalldata,
      amountPosition: 4 + 32 * 5,
    })
  }
}
