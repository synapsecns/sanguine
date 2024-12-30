import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber } from '@ethersproject/bignumber'
import invariant from 'tiny-invariant'
import { AddressZero, Zero } from '@ethersproject/constants'
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
import { logger } from '../utils/logger'
import { isSameAddress } from '../utils/addressUtils'
import { marshallTicker, Ticker } from './ticker'
import { getAllQuotes, getBestRelayerQuote, RelayerQuote } from './api'
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
  getEmptyRoute,
  RouteInput,
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
}

type DestQueryData = {
  tokenOut: string
  minAmountOut: BigNumber
  deadline: BigNumber
  paramsV1?: SavedParamsV1
  paramsV2: BridgeParamsV2
}

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
    const originIntents = await this.getOriginQuotes(
      originChainId,
      tickers,
      tokenIn,
      amountIn
    )

    // Get routes for swaps on the destination chain from the "RFQ-supported token" into tokenOut
    const destIntents = await this.getDestinationQuotes(originIntents, tokenOut)
    // Apply the quotes from the RFQ API
    const intents = await Promise.all(
      destIntents.map(async (intent) => {
        const fullQuote = await this.getFullQuote(
          intent,
          tokenOut,
          originUserAddress
        )
        const { originRoute, destRoute } = await this.generateRoutes(fullQuote)
        return {
          ...fullQuote,
          originRoute,
          destRoute,
        }
      })
    )
    return intents
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
    const decodedZapData = decodeZapData(hexlify(paramsV2.zapData))
    // Do nothing if there is no Zap on the destination chain.
    if (!decodedZapData.target) {
      return destQueryPrecise
    }
    // Regenarate ZapData with the new minAmountOut
    const minAmountOut = applySlippage(destQueryPrecise.minAmountOut, slippage)
    const zapData = encodeZapData({
      ...decodedZapData,
      minFwdAmount: minAmountOut,
    })
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

  private async getOriginQuotes(
    originChainId: number,
    tickers: Ticker[],
    tokenIn: string,
    amountIn: BigintIsh
  ): Promise<OriginIntent[]> {
    const protocolFeeRate = await this.getSynapseIntentRouter(
      originChainId
    ).getProtocolFeeRate()
    const finalRecipient: Recipient = {
      entity: RecipientEntity.Self,
      address: this.engineSet.getTokenZap(originChainId),
    }
    const inputs: RouteInput[] = tickers.map(({ originToken }) => ({
      chainId: originToken.chainId,
      tokenIn,
      tokenOut: originToken.token,
      amountIn,
      finalRecipient,
    }))
    const allQuotes = await this.engineSet.getQuotes(inputs, {
      allowMultiStep: true,
    })
    // Note: tickers.length === allQuotes.length
    // Zip the tickers and routes together, apply the protocol fee, and filter out "no amount out" routes
    return tickers
      .map((ticker, index) => ({
        ticker,
        originInput: inputs[index],
        originQuote: allQuotes[index],
        originAmountOut: this.applyProtocolFeeRate(
          allQuotes[index].expectedAmountOut,
          protocolFeeRate
        ),
      }))
      .filter(({ originQuote }) => originQuote.expectedAmountOut.gt(Zero))
  }

  private async getDestinationQuotes(
    originIntents: OriginIntent[],
    tokenOut: string
  ): Promise<FullIntent[]> {
    const finalRecipient: Recipient = {
      entity: RecipientEntity.UserSimulated,
      address: USER_SIMULATED_ADDRESS,
    }
    const inputs: RouteInput[] = originIntents.map(
      ({ ticker, originAmountOut }) => ({
        chainId: ticker.destToken.chainId,
        tokenIn: ticker.destToken.token,
        tokenOut,
        amountIn: originAmountOut,
        finalRecipient,
      })
    )
    const allQuotes = await this.engineSet.getQuotes(inputs, {
      allowMultiStep: false,
    })
    // Note: originIntents.length === allQuotes.length
    // Zip the intents and quotes together, filter out "no amount out" quotes
    return originIntents
      .map((intent, index) => ({
        ...intent,
        destInput: inputs[index],
        destQuote: allQuotes[index],
      }))
      .filter(({ destQuote }) => destQuote.expectedAmountOut.gt(Zero))
  }

  private async getFullQuote(
    intent: FullIntent,
    tokenOut: string,
    originUserAddress?: string
  ): Promise<FullQuote> {
    // Note: zap data will still be using `USER_SIMULATED_ADDRESS` address - this will be overwritten
    // when the bridge calldata is generated (until then we don't know the final recipient).
    const finalRecipient: Recipient = {
      entity: RecipientEntity.UserSimulated,
      address: USER_SIMULATED_ADDRESS,
    }
    // intent.destQuote is generated by using `originAmountOut` as the input amount on the destination chain.
    // The Relayers will also use `originAmountOut` as the input amount for simulatiion purposes as per RFQ API spec.
    const destRoute = await this.engineSet.generateRoute(
      {
        chainId: intent.ticker.destToken.chainId,
        tokenIn: intent.ticker.destToken.token,
        tokenOut,
        amountIn: intent.originAmountOut,
        finalRecipient,
      },
      intent.destQuote
    )
    // FastBridge will use TokenZap as the recipient if there are any Zap steps to perform
    const destRelayRecipient =
      destRoute.steps.length === 0
        ? USER_SIMULATED_ADDRESS
        : this.engineSet.getTokenZap(intent.ticker.destToken.chainId)
    const encodedZapDataSimulation = extractSingleZapData(destRoute.steps)
    const relayerQuote = await getBestRelayerQuote(
      intent.ticker,
      intent.originAmountOut,
      {
        originSender: originUserAddress,
        destRecipient: destRelayRecipient,
        zapData: encodedZapDataSimulation,
      }
    )
    // Update `destInput` and `destQuote` with the actual values
    const destInput = {
      ...intent.destInput,
      amountIn: relayerQuote.destAmount,
    }
    const destQuote = await this.engineSet.getQuote(
      intent.destQuote.engineID,
      destInput,
      {
        allowMultiStep: false,
      }
    )
    return {
      ...intent,
      destInput,
      destQuote,
      destRelayRecipient,
      relayerQuote,
    }
  }

  private async generateRoutes(
    fullQuote: FullQuote
  ): Promise<{ originRoute: SwapEngineRoute; destRoute: SwapEngineRoute }> {
    // Do nothing if the final amount is 0
    if (fullQuote.destQuote.expectedAmountOut.eq(Zero)) {
      return {
        originRoute: getEmptyRoute(fullQuote.originQuote.engineID),
        destRoute: getEmptyRoute(fullQuote.destQuote.engineID),
      }
    }
    const [originRoute, destRoute] = await Promise.all([
      this.engineSet.generateRoute(
        fullQuote.originInput,
        fullQuote.originQuote
      ),
      this.engineSet.generateRoute(fullQuote.destInput, fullQuote.destQuote),
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
