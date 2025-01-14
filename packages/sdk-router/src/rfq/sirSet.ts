import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber } from '@ethersproject/bignumber'
import invariant from 'tiny-invariant'
import { AddressZero, Zero } from '@ethersproject/constants'
import { hexDataLength, hexlify } from '@ethersproject/bytes'
import NodeCache from 'node-cache'

import {
  BigintIsh,
  FAST_BRIDGE_V2_ADDRESS_MAP,
  SYNAPSE_INTENT_PREVIEWER_ADDRESS_MAP,
  SYNAPSE_INTENT_ROUTER_ADDRESS_MAP,
  SWAP_QUOTER_V2_ADDRESS_MAP,
  TOKEN_ZAP_V1_ADDRESS_MAP,
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
  CCTPRouterQuery,
} from '../module'
import { SynapseIntentRouter } from './sir'
import { ChainProvider } from '../router'
import { ONE_HOUR, TEN_MINUTES } from '../utils/deadlines'
import { logExecutionTime, logger } from '../utils/logger'
import { isSameAddress } from '../utils/addressUtils'
import { marshallTicker, Ticker } from './ticker'
import {
  getAllQuotes,
  getBestRelayerQuote,
  QuoteRequestOptions,
  RelayerQuote,
} from './api'
import {
  EngineSet,
  SwapEngineRoute,
  USER_SIMULATED_ADDRESS,
  Recipient,
  RecipientEntity,
  validateEngineID,
  Slippage,
  applySlippage,
  SwapEngineQuote,
  RouteInput,
  EngineTimeout,
} from './engine'
import {
  BridgeParamsV2,
  decodeSavedBridgeParams,
  encodeSavedBridgeParams,
  SavedParamsV1,
} from './paramsV2'
import { modifyMinFinalAmount } from './zapData'
import {
  decodeStepParams,
  encodeStepParams,
  extractSingleZapData,
} from './steps'
import { FastBridgeQuote } from './quote'

type OriginIntent = {
  ticker: Ticker
  originInput: RouteInput
  originQuote: SwapEngineQuote
  originAmountOut: BigNumber
}

type DestIntent = {
  destInput: RouteInput
  destQuote: SwapEngineQuote
}

type FullIntent = OriginIntent & DestIntent

type FullQuote = FullIntent & {
  destRelayRecipient: string
  relayerQuote: RelayerQuote
  originRoute?: SwapEngineRoute
  destRoute?: SwapEngineRoute
}

type DestQueryData = {
  tokenOut: string
  minAmountOut: BigNumber
  deadline: BigNumber
  paramsV1?: SavedParamsV1
  paramsV2: BridgeParamsV2
}

export class SynapseIntentRouterSet extends SynapseModuleSet {
  public readonly bridgeModuleName = 'SynapseIntents'
  public readonly allEvents = ['BridgeRequestedEvent', 'BridgeRelayedEvent']

  public routers: {
    [chainId: number]: SynapseIntentRouter
  }
  public providers: {
    [chainId: number]: Provider
  }

  private engineSet: EngineSet
  private cache = new NodeCache()

  constructor(chains: ChainProvider[]) {
    super()
    this.routers = {}
    this.providers = {}
    chains.forEach(({ chainId, provider }) => {
      const sirAddress = SYNAPSE_INTENT_ROUTER_ADDRESS_MAP[chainId]
      // Skip chains without a SynapseIntentRouter address
      if (sirAddress) {
        this.routers[chainId] = new SynapseIntentRouter(chainId, provider, {
          fastBridgeV2Address: FAST_BRIDGE_V2_ADDRESS_MAP[chainId],
          previewerAddress: SYNAPSE_INTENT_PREVIEWER_ADDRESS_MAP[chainId],
          sirAddress,
          swapQuoterAddress: SWAP_QUOTER_V2_ADDRESS_MAP[chainId],
          tokenZapAddress: TOKEN_ZAP_V1_ADDRESS_MAP[chainId],
        })
        this.providers[chainId] = provider
      }
    })
    this.engineSet = new EngineSet(chains)
  }

  /**
   * @inheritdoc SynapseModuleSet.getModule
   */
  public getModule(chainId: number): SynapseModule | undefined {
    return this.routers[chainId]
  }

  /**
   * @inheritdoc SynapseModuleSet.getEstimatedTime
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

  /**
   * @inheritdoc SynapseModuleSet.getBridgeRoutes
   */
  @logExecutionTime('SynapseIntents.getBridgeRoutes')
  public async getBridgeRoutes(
    originChainId: number,
    destChainId: number,
    tokenIn: string,
    tokenOut: string,
    amountIn: BigintIsh,
    originUserAddress?: string
  ): Promise<BridgeRoute[]> {
    // Check that Routers exist on both chains
    if (!this.getModule(originChainId) || !this.getModule(destChainId)) {
      return []
    }
    // Get all tickers that can be used between the two chains
    const tickers = await this.getAllTickers(originChainId, destChainId)
    const protocolFeeRate = await this.getSynapseIntentRouter(
      originChainId
    ).getProtocolFeeRate()
    const quotes = await Promise.all(
      tickers.map(async (ticker) =>
        this.getTickerQuote(
          ticker,
          tokenIn,
          tokenOut,
          amountIn,
          protocolFeeRate,
          originUserAddress
        )
      )
    )
    return quotes
      .filter((quote): quote is Required<FullQuote> => !!quote)
      .filter(({ destRoute }) => destRoute.expectedAmountOut.gt(Zero))
      .map((intent) => ({
        bridgeModuleName: this.bridgeModuleName,
        originChainId,
        destChainId,
        bridgeToken: {
          symbol: marshallTicker(intent.ticker),
          token: intent.ticker.destToken.token,
        },
        originQuery: this.engineSet.getOriginQuery(
          originChainId,
          intent.originRoute,
          intent.ticker.originToken.token
        ),
        originAmountOut: intent.originAmountOut,
        destQuery: this.getRFQDestinationQuery({
          tokenOut,
          minAmountOut: intent.destRoute.expectedAmountOut,
          // The default deadline will be overridden later in `finalizeBridgeRoute`
          deadline: Zero,
          paramsV1: this.getSavedParamsV1(intent, originUserAddress),
          paramsV2: this.getBridgeParamsV2(intent.destRoute),
        }),
        destAmountIn: intent.relayerQuote.destAmount,
        destAmountOut: intent.destRoute.expectedAmountOut,
      }))
  }

  /**
   * @inheritdoc SynapseModuleSet.getFeeData
   */
  public async getFeeData(bridgeRoute: BridgeRoute): Promise<{
    feeAmount: BigNumber
    feeConfig: FeeConfig
  }> {
    // Origin Out vs Dest In is the effective fee
    const feeAmount = bridgeRoute.originAmountOut.sub(bridgeRoute.destAmountIn)
    return {
      feeAmount: feeAmount.lt(Zero) ? Zero : feeAmount,
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
    // We should have saved neccessary params within dstQuery.rawParams
    if (hexDataLength(destQueryPrecise.rawParams) === 0) {
      logger.warn(
        'No params saved in destQuery.rawParams, slippage is not applied'
      )
      return {
        originQuery: originQueryPrecise,
        destQuery: destQueryPrecise,
      }
    }
    // Find out the quoted destAmount for the RFQ token
    const { paramsV1, paramsV2 } = decodeSavedBridgeParams(
      destQueryPrecise.rawParams
    )
    if (
      isSameAddress(paramsV1.destRelayToken, AddressZero) ||
      paramsV1.destRelayAmount.eq(0)
    ) {
      logger.warn(
        'No destToken or destAmount saved in destQuery.rawParams, slippage is not applied'
      )
      return {
        originQuery: originQueryPrecise,
        destQuery: destQueryPrecise,
      }
    }
    const slippage = {
      numerator: slipNumerator,
      denominator: slipDenominator,
    }
    return {
      originQuery: this.applyOriginSlippage(
        originQueryPrecise,
        paramsV1.destRelayAmount,
        slippage
      ),
      destQuery: this.applyDestinationSlippage(
        destQueryPrecise,
        paramsV1,
        paramsV2,
        slippage
      ),
    }
  }

  private applyOriginSlippage(
    originQueryPrecise: Query,
    destRelayAmount: BigNumber,
    slippage: Slippage
  ): Query {
    // Do nothing if there are no Zap steps.
    if (hexDataLength(originQueryPrecise.rawParams) === 0) {
      return originQueryPrecise
    }
    // Max slippage for the origin swap is 5% of the (destAmount - originAmount).
    // Anything over that might lead to quote that the Relayers will not process.
    const maxOriginSlippage = originQueryPrecise.minAmountOut
      .sub(destRelayAmount)
      .div(20)
    // TODO: figure out a better way to handle destAmount > originAmount
    const minAmountFinalAmount = maxOriginSlippage.isNegative()
      ? originQueryPrecise.minAmountOut
      : originQueryPrecise.minAmountOut.sub(maxOriginSlippage)
    const originQuery = applySlippageToQuery(
      originQueryPrecise,
      slippage.numerator,
      slippage.denominator
    )
    if (originQuery.minAmountOut.lt(minAmountFinalAmount)) {
      originQuery.minAmountOut = minAmountFinalAmount
    }
    // Adjust the slippage in the last origin step.
    const originSteps = decodeStepParams(originQueryPrecise.rawParams)
    if (originSteps.length === 0) {
      logger.error({ originQueryPrecise }, 'No steps in originQueryPrecise')
      return originQuery
    }
    originSteps[originSteps.length - 1].zapData = modifyMinFinalAmount(
      hexlify(originSteps[originSteps.length - 1].zapData),
      originQuery.minAmountOut
    )
    originQuery.rawParams = encodeStepParams(originSteps)
    return originQuery
  }

  private applyDestinationSlippage(
    destQueryPrecise: Query,
    paramsV1: SavedParamsV1,
    paramsV2: BridgeParamsV2,
    slippage: Slippage
  ): Query {
    // Check that engineID is within range
    if (!validateEngineID(paramsV1.destEngineID)) {
      throw new Error(`Invalid engineID: ${paramsV1.destEngineID}`)
    }
    const oldZapData = hexlify(paramsV2.zapData)
    // Do nothing if there is no Zap on the destination chain.
    if (hexDataLength(oldZapData) === 0) {
      return destQueryPrecise
    }
    // Regenarate ZapData with the new minAmountOut
    const minAmountOut = applySlippage(destQueryPrecise.minAmountOut, slippage)
    const zapData = modifyMinFinalAmount(oldZapData, minAmountOut)
    return this.getRFQDestinationQuery({
      tokenOut: destQueryPrecise.tokenOut,
      minAmountOut,
      deadline: destQueryPrecise.deadline,
      paramsV1,
      paramsV2: {
        ...paramsV2,
        zapData,
      },
    })
  }

  /**
   * Returns the existing SynapseIntentRouter instance for the given chain.
   *
   * @throws Will throw an error if SynapseIntentRouter is not deployed on the given chain.
   */
  public getSynapseIntentRouter(chainId: number): SynapseIntentRouter {
    return this.getExistingModule(chainId) as SynapseIntentRouter
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

  @logExecutionTime('SynapseIntents.getTickerQuote')
  private async getTickerQuote(
    ticker: Ticker,
    tokenIn: string,
    tokenOut: string,
    amountIn: BigintIsh,
    protocolFeeRate: BigNumber,
    originUserAddress?: string
  ): Promise<Required<FullQuote> | undefined> {
    const originIntent = await this.getOriginQuote(
      ticker,
      tokenIn,
      amountIn,
      protocolFeeRate
    )
    if (!originIntent) {
      return
    }
    const destIntent = await this.getDestinationQuote(originIntent, tokenOut)
    if (!destIntent) {
      return
    }
    const fullQuote = await this.getFullQuote(destIntent, originUserAddress)
    if (!fullQuote) {
      return
    }
    const { originRoute, destRoute } = await this.generateRoutes(fullQuote)
    if (!originRoute || !destRoute) {
      return
    }
    return {
      ...fullQuote,
      originRoute,
      destRoute,
    }
  }

  @logExecutionTime('SynapseIntents.getOriginQuote')
  private async getOriginQuote(
    ticker: Ticker,
    tokenIn: string,
    amountIn: BigintIsh,
    protocolFeeRate: BigNumber
  ): Promise<OriginIntent | undefined> {
    const finalRecipient: Recipient = {
      entity: RecipientEntity.Self,
      address: this.engineSet.getTokenZap(ticker.originToken.chainId),
    }
    // Swap complexity is not restricted on the origin chain, where execution is done by the user at the time of bridging.
    const input: RouteInput = {
      chainId: ticker.originToken.chainId,
      tokenIn,
      tokenOut: ticker.originToken.token,
      amountIn,
      finalRecipient,
      restrictComplexity: false,
    }
    const quote = await this.engineSet.getBestQuote(input, {
      allowMultiStep: true,
    })
    if (!quote) {
      return
    }
    return {
      ticker,
      originInput: input,
      originQuote: quote,
      originAmountOut: this.applyProtocolFeeRate(
        quote.expectedAmountOut,
        protocolFeeRate
      ),
    }
  }

  @logExecutionTime('SynapseIntents.getDestinationQuote')
  private async getDestinationQuote(
    originIntent: OriginIntent,
    tokenOut: string
  ): Promise<FullIntent | undefined> {
    const finalRecipient: Recipient = {
      entity: RecipientEntity.UserSimulated,
      address: USER_SIMULATED_ADDRESS,
    }
    // Swap complexity is restricted on the destination chain, where execution is done by the Relayers with a delay.
    const input: RouteInput = {
      chainId: originIntent.ticker.destToken.chainId,
      tokenIn: originIntent.ticker.destToken.token,
      tokenOut,
      amountIn: originIntent.originAmountOut,
      finalRecipient,
      restrictComplexity: true,
    }
    const quote = await this.engineSet.getBestQuote(input, {
      allowMultiStep: false,
    })
    if (!quote) {
      return
    }
    return {
      ...originIntent,
      destInput: input,
      destQuote: quote,
    }
  }

  @logExecutionTime('SynapseIntents.getFullQuote')
  private async getFullQuote(
    intent: FullIntent,
    originUserAddress?: string
  ): Promise<FullQuote | undefined> {
    // Note: we leave the default max slippage from `generateRoute` here to ensure that the Relayer simulation
    // suceeds even in the even that on-chain price moves. We will overwrite this later in `generateRoutes`.
    const destRoute = await this.engineSet.generateRoute(
      intent.destInput,
      intent.destQuote,
      { allowMultiStep: false, useZeroSlippage: false }
    )
    if (!destRoute) {
      return
    }
    // FastBridge will use TokenZap as the recipient if there are any Zap steps to perform
    const destRelayRecipient =
      destRoute.steps.length === 0
        ? USER_SIMULATED_ADDRESS
        : this.engineSet.getTokenZap(intent.ticker.destToken.chainId)
    const encodedZapDataSimulation = extractSingleZapData(destRoute.steps)
    // intent.destQuote is generated by using `originAmountOut` as the input amount on the destination chain.
    // The Relayers will also use `originAmountOut` as the input amount for simulatiion purposes as per RFQ API spec.
    const relayerQuote = await this.apiGetBestRelayerQuote(
      intent.ticker,
      intent.originAmountOut,
      {
        originSender: originUserAddress,
        destRecipient: destRelayRecipient,
        zapData: encodedZapDataSimulation,
      }
    )
    if (!relayerQuote) {
      return
    }
    return {
      ...intent,
      destRelayRecipient,
      relayerQuote,
    }
  }

  @logExecutionTime('SynapseIntents.generateRoutes')
  private async generateRoutes(
    fullQuote: FullQuote
  ): Promise<{ originRoute?: SwapEngineRoute; destRoute?: SwapEngineRoute }> {
    // Update `destInput` and `destQuote` with the actual values.
    const destInput = {
      ...fullQuote.destInput,
      amountIn: fullQuote.relayerQuote.destAmount,
    }
    // Use longer timeout for finalizing the route.
    const destQuote = await this.engineSet.getQuote(
      fullQuote.destQuote.engineID,
      destInput,
      {
        allowMultiStep: false,
        timeout: EngineTimeout.Long,
      }
    )
    if (!destQuote) {
      return {}
    }
    // Final rotures will be returned with a zero slippage by default, and could be then modified
    // by the SDK consumer.
    const [originRoute, destRoute] = await Promise.all([
      this.engineSet.generateRoute(
        fullQuote.originInput,
        fullQuote.originQuote,
        {
          allowMultiStep: true,
          useZeroSlippage: true,
        }
      ),
      this.engineSet.generateRoute(destInput, destQuote, {
        allowMultiStep: false,
        useZeroSlippage: true,
      }),
    ])
    return {
      originRoute,
      destRoute,
    }
  }

  /**
   * Get all unique tickers for a given origin and destination chains. In other words,
   * this is the list of all (originToken, destToken) pairs that can be used to create a quote
   * for a swap between the two chains, without duplicates.
   *
   * @param originChainId - The ID of the origin chain.
   * @param destChainId - The ID of the destination chain.
   * @returns A promise that resolves to the list of tickers.
   */
  private async getAllTickers(
    originChainId: number,
    destChainId: number
  ): Promise<Ticker[]> {
    const allQuotes = await this.apiGetAllQuotes()
    const originFB = FAST_BRIDGE_V2_ADDRESS_MAP[originChainId]
    const destFB = FAST_BRIDGE_V2_ADDRESS_MAP[destChainId]
    // First, we filter out quotes for other chainIDs and bridge addresses.
    // Finally, we remove the duplicates of the origin token.
    return allQuotes
      .filter(
        (quote) =>
          quote.ticker.originToken.chainId === originChainId &&
          quote.ticker.destToken.chainId === destChainId &&
          isSameAddress(quote.originFastBridge, originFB) &&
          isSameAddress(quote.destFastBridge, destFB)
      )
      .map((quote) => quote.ticker)
      .filter(
        (ticker, index, self) =>
          index ===
          self.findIndex((t) =>
            isSameAddress(t.originToken.token, ticker.originToken.token)
          )
      )
  }

  @logExecutionTime('API/quotes')
  private async apiGetAllQuotes(): Promise<FastBridgeQuote[]> {
    // Try getting cached quotes first.
    const cacheKey = 'getAllQuotes'
    const cachedQuotes = this.cache.get<FastBridgeQuote[]>(cacheKey)
    if (cachedQuotes) {
      return cachedQuotes
    }
    // If not cached, fetch new quotes and cache them.
    const data = await getAllQuotes()
    this.cache.set(cacheKey, data, ONE_HOUR)
    return data
  }

  @logExecutionTime('API/rfq')
  private async apiGetBestRelayerQuote(
    ticker: Ticker,
    originAmount: BigNumber,
    options: QuoteRequestOptions = {}
  ): Promise<RelayerQuote | undefined> {
    const quote = await getBestRelayerQuote(ticker, originAmount, options)
    return quote.destAmount.gt(Zero) ? quote : undefined
  }

  private getSavedParamsV1(
    intent: FullQuote,
    originUserAddress?: string
  ): SavedParamsV1 | undefined {
    return originUserAddress
      ? {
          originSender: originUserAddress,
          destChainId: intent.ticker.destToken.chainId,
          destEngineID: intent.destQuote.engineID,
          destRelayRecipient: intent.destRelayRecipient,
          destRelayToken: intent.ticker.destToken.token,
          destRelayAmount: intent.relayerQuote.destAmount,
        }
      : undefined
  }

  private getBridgeParamsV2(destRoute: SwapEngineRoute): BridgeParamsV2 {
    return {
      // TODO: exclusivity
      quoteRelayer: AddressZero,
      quoteExclusivitySeconds: Zero,
      // TODO: quoteId
      quoteId: '0x',
      zapNative: Zero,
      zapData: extractSingleZapData(destRoute.steps),
    }
  }

  private getRFQDestinationQuery(data: DestQueryData): CCTPRouterQuery {
    // Use no-swap query by default.
    const destQuery = createNoSwapQuery(data.tokenOut, data.minAmountOut)
    destQuery.deadline = data.deadline
    if (!data.paramsV1) {
      return destQuery
    }
    destQuery.routerAdapter = this.engineSet.getTokenZap(
      data.paramsV1.destChainId
    )
    destQuery.rawParams = encodeSavedBridgeParams(data.paramsV1, data.paramsV2)
    return destQuery
  }
}
