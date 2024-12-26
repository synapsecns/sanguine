import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber } from '@ethersproject/bignumber'
import invariant from 'tiny-invariant'
import { AddressZero, MaxUint256, Zero } from '@ethersproject/constants'
import { hexDataLength, hexlify } from '@ethersproject/bytes'

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
import { isSameAddress } from '../utils/addressUtils'
import { marshallTicker, Ticker } from './ticker'
import { getAllQuotes, getBestRelayerQuote } from './api'
import {
  EngineSet,
  SwapEngineRoute,
  USER_SIMULATED_ADDRESS,
  Recipient,
  RecipientEntity,
  validateEngineID,
  Slippage,
  applySlippage,
} from './engine'
import {
  BridgeParamsV2,
  decodeSavedBridgeParams,
  encodeSavedBridgeParams,
  SavedParamsV1,
} from './paramsV2'
import { decodeZapData, encodeZapData } from './zapData'
import { extractSingleZapData } from './steps'

type OriginIntent = {
  ticker: Ticker
  originRoute: SwapEngineRoute
  originAmountOut: BigNumber
}

type DestIntent = {
  destRelayAmount: BigNumber
  destRelayRecipient: string
  destRelayToken: string
  destRoute: SwapEngineRoute
}

type FullIntent = OriginIntent & DestIntent

export class SynapseIntentRouterSet extends SynapseModuleSet {
  static readonly MAX_QUOTE_AGE_MILLISECONDS = 5 * 60 * 1000 // 5 minutes

  public readonly bridgeModuleName = 'SynapseIntents'
  public readonly allEvents = ['BridgeRequestedEvent', 'BridgeRelayedEvent']

  public routers: {
    [chainId: number]: SynapseIntentRouter
  }
  public providers: {
    [chainId: number]: Provider
  }

  private engineSet: EngineSet

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
    // Get routes for swaps on the origin chain from tokenIn into the "RFQ-supported token", and apply the protocol fees
    const originRoutes = await this.getIntentsWithOriginRoute(
      originChainId,
      tickers,
      tokenIn,
      amountIn
    )

    // Get routes for swaps on the destination chain from the "RFQ-supported token" into tokenOut
    const destRoutes = await this.getIntentsWithDestRoute(
      destChainId,
      originRoutes,
      tokenOut
    )
    // Apply the quotes from the RFQ API
    const fullQuotes = await Promise.all(
      destRoutes.map((route) =>
        this.getFinalIntentQuote(route, tokenOut, originUserAddress)
      )
    )
    return fullQuotes
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
        destQuery: this.getRFQDestinationQuery(
          destChainId,
          intent,
          tokenOut,
          intent.destRoute.expectedAmountOut,
          // The default deadline will be overridden later in `finalizeBridgeRoute`
          Zero,
          originUserAddress
        ),
        destAmountIn: intent.destRelayAmount,
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
      console.warn(
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
      console.warn(
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
        paramsV1.destChainId,
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
    return originQuery
  }

  private applyDestinationSlippage(
    destChainId: number,
    destQueryPrecise: Query,
    paramsV1: SavedParamsV1,
    paramsV2: BridgeParamsV2,
    slippage: Slippage
  ): Query {
    // Check that engineID is within range
    if (!validateEngineID(paramsV1.destEngineID)) {
      throw new Error(`Invalid engineID: ${paramsV1.destEngineID}`)
    }
    const decodedZapData = decodeZapData(hexlify(paramsV2.zapData))
    // Do nothing if there is no Zap on the destination chain.
    if (!decodedZapData.target) {
      return destQueryPrecise
    }
    const expectedAmountOut = destQueryPrecise.minAmountOut
    const minAmountOut = applySlippage(expectedAmountOut, slippage)
    const zapData = encodeZapData({
      ...decodedZapData,
      minFwdAmount: minAmountOut,
    })
    const destRoute = {
      engineID: paramsV1.destEngineID,
      expectedAmountOut,
      steps: [
        {
          token: paramsV1.destRelayToken,
          // Use the full balance for the Zap action
          amount: MaxUint256,
          msgValue: paramsV2.zapNative,
          zapData,
        },
      ],
    }
    return this.getRFQDestinationQuery(
      destChainId,
      {
        destRelayAmount: paramsV1.destRelayAmount,
        destRelayRecipient: paramsV1.destRelayRecipient,
        destRelayToken: paramsV1.destRelayToken,
        destRoute,
      },
      destQueryPrecise.tokenOut,
      minAmountOut,
      destQueryPrecise.deadline,
      paramsV1.originSender
    )
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

  private async getIntentsWithOriginRoute(
    originChainId: number,
    tickers: Ticker[],
    tokenIn: string,
    amountIn: BigintIsh
  ): Promise<OriginIntent[]> {
    const protocolFeeRate = await this.getSynapseIntentRouter(
      originChainId
    ).getProtocolFeeRate()
    const allRoutes = await this.engineSet.getOriginRoutes(
      originChainId,
      { address: tokenIn, amount: amountIn },
      tickers.map((ticker) => ticker.originToken.token)
    )
    // Note: tickers.length === allRoutes.length
    // Zip the tickers and routes together, apply the protocol fee, and filter out "no amount out" routes
    return tickers
      .map((ticker, index) => ({
        ticker,
        originRoute: allRoutes[index],
        originAmountOut: this.applyProtocolFeeRate(
          allRoutes[index].expectedAmountOut,
          protocolFeeRate
        ),
      }))
      .filter(({ originRoute }) => originRoute.expectedAmountOut.gt(Zero))
  }

  private async getIntentsWithDestRoute(
    destChainId: number,
    intents: OriginIntent[],
    tokenOut: string
  ): Promise<FullIntent[]> {
    const allRoutes = await this.engineSet.getDestinationRoutes(
      destChainId,
      intents.map(({ ticker, originAmountOut }) => ({
        address: ticker.destToken.token,
        amount: originAmountOut,
      })),
      tokenOut
    )
    // Note: originRoutes.length === allRoutes.length
    // Zip the origin routes and routes together, filter out "no amount out" routes
    return intents
      .map(({ ticker, originRoute, originAmountOut }, index) => ({
        ticker,
        originRoute,
        originAmountOut,
        // Will be filled in `getFinalQuote`
        destRelayAmount: Zero,
        // FastBridge will use TokenZap as the recipient if there are any Zap steps to perform
        destRelayRecipient:
          allRoutes[index].steps.length === 0
            ? USER_SIMULATED_ADDRESS
            : this.engineSet.getTokenZap(destChainId),
        destRelayToken: ticker.destToken.token,
        destRoute: allRoutes[index],
      }))
      .filter(({ destRoute }) => destRoute.expectedAmountOut.gt(Zero))
  }

  private async getFinalIntentQuote(
    route: FullIntent,
    tokenOut: string,
    originUserAddress?: string
  ): Promise<FullIntent> {
    // `encodedZapDataSimulated` was generated by using `originAmountOut` as the imput amount on the destination chain.
    const encodedZapDataSimulated = extractSingleZapData(route.destRoute.steps)
    const quote = await getBestRelayerQuote(
      route.ticker,
      route.originAmountOut,
      {
        originSender: originUserAddress,
        destRecipient: route.destRelayRecipient,
        zapData: encodedZapDataSimulated,
      }
    )
    // Now that we got the quote, we need to get the final amount out and adjust the zap data.
    // Note: zap data will still be using `USER_SIMULATED_ADDRESS` address - this will be overwritten
    // when the bridge calldata is generated (until then we don't know the final recipient).
    const destFinalRecipient: Recipient = {
      entity: RecipientEntity.UserSimulated,
      address: USER_SIMULATED_ADDRESS,
    }
    const finalDestRoute = await this.engineSet.findRoute(
      route.destRoute.engineID,
      route.ticker.destToken.chainId,
      { address: route.ticker.destToken.token, amount: quote.destAmount },
      tokenOut,
      destFinalRecipient
    )
    return {
      ...route,
      destRelayAmount: quote.destAmount,
      // Up to a single Zap is supported on the destination chain
      destRoute: this.engineSet.limitSingleZap(finalDestRoute),
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
    const allQuotes = await getAllQuotes()
    const originFB = FAST_BRIDGE_V2_ADDRESS_MAP[originChainId]
    const destFB = FAST_BRIDGE_V2_ADDRESS_MAP[destChainId]
    // First, we filter out quotes for other chainIDs and bridge addresses.
    // Then, we filter out quotes that are too old.
    // Finally, we remove the duplicates of the origin token.
    return allQuotes
      .filter((quote) => {
        const areSameChainsAndToken =
          quote.ticker.originToken.chainId === originChainId &&
          quote.ticker.destToken.chainId === destChainId &&
          isSameAddress(quote.originFastBridge, originFB) &&
          isSameAddress(quote.destFastBridge, destFB)
        // TODO: don't filter by age here
        const age = Date.now() - quote.updatedAt
        const isValidAge =
          0 <= age && age < SynapseIntentRouterSet.MAX_QUOTE_AGE_MILLISECONDS
        return areSameChainsAndToken && isValidAge
      })
      .map((quote) => quote.ticker)
      .filter(
        (ticker, index, self) =>
          index ===
          self.findIndex((t) =>
            isSameAddress(t.originToken.token, ticker.originToken.token)
          )
      )
  }

  private getRFQDestinationQuery(
    destChainId: number,
    intent: DestIntent,
    tokenOut: string,
    minAmountOut: BigNumber,
    deadline: BigNumber,
    originUserAddress?: string
  ): CCTPRouterQuery {
    // Use no-swap query by default.
    const destQuery = createNoSwapQuery(tokenOut, minAmountOut)
    destQuery.deadline = deadline
    if (!originUserAddress) {
      return destQuery
    }
    destQuery.routerAdapter = this.engineSet.getTokenZap(destChainId)
    // Encode neccessary params for invoking the FastBridgeV2 bridge function.
    const dstZapData = extractSingleZapData(intent.destRoute.steps)
    destQuery.rawParams = encodeSavedBridgeParams(
      {
        originSender: originUserAddress,
        destChainId,
        destEngineID: intent.destRoute.engineID,
        destRelayRecipient: intent.destRelayRecipient,
        destRelayToken: intent.destRelayToken,
        destRelayAmount: intent.destRelayAmount,
      },
      {
        // TODO: exclusivity
        quoteRelayer: AddressZero,
        quoteExclusivitySeconds: Zero,
        // TODO: quote ID
        quoteId: '0x',
        zapNative: Zero,
        zapData: dstZapData,
      }
    )
    return destQuery
  }
}
