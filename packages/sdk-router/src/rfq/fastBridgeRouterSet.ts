import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber } from '@ethersproject/bignumber'
import { Zero } from '@ethersproject/constants'
import { BigNumberish } from 'ethers'
import NodeCache from 'node-cache'

import {
  FAST_BRIDGE_INTERCEPTOR_ADDRESS_MAP,
  FAST_BRIDGE_ROUTER_ADDRESS_MAP,
  MEDIAN_TIME_RFQ,
  MEDIAN_TIME_RFQ_ETHEREUM,
  SupportedChainId,
} from '../constants'
import {
  BridgeRoute,
  FeeConfig,
  Query,
  SynapseModule,
  SynapseModuleSet,
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
  logExecutionTime,
} from '../utils'
import { getAllQuotes } from './api'
import { FastBridgeQuote, applyQuote } from './quote'
import { marshallTicker } from './ticker'
import {
  IFastBridge,
  IFastBridgeInterceptor,
} from '../typechain/FastBridgeInterceptor'

enum CacheDuration {
  Short = 'short',
  Long = 'long',
}

export class FastBridgeRouterSet extends SynapseModuleSet {
  static readonly MAX_QUOTE_AGE_MILLISECONDS = 5 * 60 * 1000 // 5 minutes
  static readonly QUOTES_TTL: Record<CacheDuration, number> = {
    [CacheDuration.Short]: 10, // 10 seconds
    [CacheDuration.Long]: 60 * 60, // 1 hour
  }

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
    this.quotesCache = new NodeCache()
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
  public getEstimatedTime(originChainId: number, destChainId?: number): number {
    return originChainId === SupportedChainId.ETH ||
      destChainId === SupportedChainId.ETH
      ? MEDIAN_TIME_RFQ_ETHEREUM
      : MEDIAN_TIME_RFQ
  }

  /**
   * @inheritdoc SynapseModuleSet.getGasDropAmount
   */
  public async getGasDropAmount(): Promise<BigNumber> {
    return Zero
  }

  @logExecutionTime('FastBridgeRouterSet.getBridgeTokenCandidates')
  public async getBridgeTokenCandidates({
    fromChainId,
    toChainId,
    toToken,
  }: GetBridgeTokenCandidatesParameters): Promise<BridgeTokenCandidate[]> {
    if (!this.getModule(fromChainId) || !this.getModule(toChainId)) {
      return []
    }
    // Use long cache duration for token candidates
    const quotes = await this.getQuotes(
      CacheDuration.Long,
      fromChainId,
      toChainId,
      toToken
    )
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

  @logExecutionTime('FastBridgeRouterSet.getBridgeRouteV2')
  public async getBridgeRouteV2(
    params: GetBridgeRouteV2Parameters
  ): Promise<BridgeRouteV2 | undefined> {
    if (!this.validateBridgeRouteV2Params(params)) {
      return undefined
    }
    const { originSwapRoute, bridgeToken, fromSender, toRecipient, slippage } =
      params
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
        CacheDuration.Short, // Use short cache duration for most recent quotes
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
    const hasOriginSlippage = !originSwapRoute.expectedToAmount.eq(
      originSwapRoute.minToAmount
    )
    const minToAmount =
      hasOriginSlippage && slippage
        ? applySlippage(expectedToAmount, slippage)
        : expectedToAmount
    const route: BridgeRouteV2 = {
      bridgeToken,
      toToken: bridgeToken.destToken,
      expectedToAmount,
      minToAmount,
      nativeFee: Zero,
      zapData: await this.getBridgeZapData(
        bridgeToken,
        expectedToAmount,
        hasOriginSlippage ? originSwapRoute.expectedToAmount : undefined,
        fromSender,
        toRecipient
      ),
    }
    return route
  }

  /**
   * @inheritdoc SynapseModuleSet.getBridgeRoutes
   */
  @logExecutionTime('FastBridgeRouterSet.getBridgeRoutes')
  public async getBridgeRoutes(
    _originChainId: number,
    _destChainId: number,
    _tokenIn: string,
    _tokenOut: string,
    _amountIn: BigNumberish,
    _originUserAddress?: string
  ): Promise<BridgeRoute[]> {
    // Bridge V1 is not supported.
    return []
  }

  /**
   * @inheritdoc SynapseModuleSet.getFeeData
   */
  public async getFeeData(_bridgeRoute: BridgeRoute): Promise<{
    feeAmount: BigNumber
    feeConfig: FeeConfig
  }> {
    return {
      feeAmount: Zero,
      feeConfig: {
        bridgeFee: 0,
        minFee: Zero,
        maxFee: Zero,
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
    _slipNumerator: number,
    _slipDenominator: number
  ): { originQuery: Query; destQuery: Query } {
    return {
      originQuery: originQueryPrecise,
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
   * Retrieves all quotes with caching.
   *
   * @returns A promise that resolves to all available quotes.
   */
  private async getAllQuotes(
    cacheDuration: CacheDuration
  ): Promise<FastBridgeQuote[]> {
    const cacheKey = `all_quotes_${cacheDuration}`
    const cachedQuotes = this.quotesCache.get<FastBridgeQuote[]>(cacheKey)
    if (cachedQuotes) {
      return cachedQuotes
    }
    const allQuotes = await getAllQuotes()
    // Update both long and short caches
    this.quotesCache.set(
      `all_quotes_${CacheDuration.Long}`,
      allQuotes,
      FastBridgeRouterSet.QUOTES_TTL.long
    )
    this.quotesCache.set(
      `all_quotes_${CacheDuration.Short}`,
      allQuotes,
      FastBridgeRouterSet.QUOTES_TTL.short
    )
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
    cacheDuration: CacheDuration,
    originChainId: number,
    destChainId: number,
    tokenOut?: string
  ): Promise<FastBridgeQuote[]> {
    const allQuotes = await this.getAllQuotes(cacheDuration)
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

  private async getBridgeZapData(
    bridgeToken: BridgeTokenCandidate,
    expectedToAmount: BigNumber,
    fromAmount?: BigNumber,
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
    const fastBridge = await this.getFastBridgeRouter(
      bridgeToken.originChainId
    ).getFastBridgeContract()
    if (fromAmount) {
      // Quote origin amount was supplied - use Interceptor to adjust the quote in flight
      const fastBridgeInterceptor = this.getFastBridgeRouter(
        bridgeToken.originChainId
      ).interceptorContract
      if (!fastBridgeInterceptor) {
        logger.error(
          `FastBridgeInterceptor not found for chainId ${bridgeToken.originChainId}`
        )
        return undefined
      }
      const interceptorParams: IFastBridgeInterceptor.InterceptorParamsStruct =
        {
          fastBridge: fastBridge.address,
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
    // Quote origin amount was not supplied - use FastBridge to bridge directly
    const fastBridgeCalldata = (
      await fastBridge.populateTransaction.bridge(bridgeParams)
    ).data
    return encodeZapData({
      target: fastBridge.address,
      payload: fastBridgeCalldata,
      amountPosition: 4 + 32 * 5,
    })
  }
}
