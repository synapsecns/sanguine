import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber } from '@ethersproject/bignumber'
import invariant from 'tiny-invariant'
import { AddressZero, MaxUint256, Zero } from '@ethersproject/constants'
import { hexlify } from '@ethersproject/bytes'

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
} from '../module'
import { SynapseIntentRouter } from './sir'
import { ChainProvider } from '../router'
import { ONE_HOUR, TEN_MINUTES } from '../utils/deadlines'
import { isSameAddress } from '../utils/addressUtils'
import { marshallTicker, Ticker } from './ticker'
import { getAllQuotes, getBestRelayerQuote } from './api'
import {
  BridgeParamsV2,
  SavedParamsV1,
  encodeSavedBridgeParams,
} from './paramsV2'
import {
  applyDefaultValues,
  decodeZapData,
  FORWARD_TO_SIMULATED,
  ZapDataV1,
} from './zapData'

type OriginIntent = {
  ticker: Ticker
  originQuery: Query
  originAmountOut: BigNumber
}

type DestIntent = OriginIntent & {
  encodedZapData: string
  tokenOut: string
}

type FullIntent = DestIntent & {
  destQuery: Query
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
    const originSIR = this.getSynapseIntentRouter(originChainId)
    const destSIR = this.getSynapseIntentRouter(destChainId)
    // Get all tickers that can be used between the two chains
    const tickers = await this.getAllTickers(originChainId, destChainId)
    // Get queries for swaps on the origin chain from tokenIn into the "RFQ-supported token"
    const originIntents = await this.filterTickersWithOriginSwap(
      originSIR,
      tickers,
      tokenIn,
      amountIn
    )
    // Get queries for swaps on the destination chain from the "RFQ-supported token" into tokenOut
    const intents: DestIntent[] = await this.filterQuotesWithDestSwap(
      destSIR,
      originIntents,
      tokenOut
    )
    // Apply the quotes from the RFQ API
    const fullQuotes = await Promise.all(
      intents.map((intent) =>
        this.getFullIntentQuote(destSIR, intent, originUserAddress)
      )
    )
    return fullQuotes
      .filter(({ destQuery }) => destQuery.minAmountOut.gt(0))
      .map(({ ticker, originQuery, destQuery }) => ({
        originChainId,
        destChainId,
        bridgeToken: {
          symbol: marshallTicker(ticker),
          token: ticker.destToken.token,
        },
        originQuery,
        destQuery,
        bridgeModuleName: this.bridgeModuleName,
      }))
  }

  /**
   * @inheritdoc SynapseModuleSet.getFeeData
   */
  public async getFeeData(bridgeRoute: BridgeRoute): Promise<{
    feeAmount: BigNumber
    feeConfig: FeeConfig
  }> {
    // Origin Out vs Dest Out is the effective fee
    return {
      feeAmount: bridgeRoute.originQuery.minAmountOut.sub(
        bridgeRoute.destQuery.minAmountOut
      ),
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

  /**
   * Filters the list of tickers to only include those that can be used for given amount of input token.
   * For every filtered ticker, the origin query is returned with the information for tokenIn -> ticker swaps.
   */
  private async filterTickersWithOriginSwap(
    originSIR: SynapseIntentRouter,
    tickers: Ticker[],
    tokenIn: string,
    amountIn: BigintIsh
  ): Promise<OriginIntent[]> {
    // Get queries for swaps on the origin chain into the "RFQ-supported token"
    const originQueries = await originSIR.getOriginAmountOut(
      tokenIn,
      tickers.map((ticker) => ticker.originToken.token),
      amountIn
    )
    const protocolFeeRate = await originSIR.getProtocolFeeRate()
    // Note: tickers.length === originQueries.length
    // Zip the tickers and queries together, filter out "no path found" queries
    return tickers
      .map((ticker, index) => ({
        ticker,
        originQuery: originQueries[index],
        originAmountOut: this.applyProtocolFeeRate(
          originQueries[index].minAmountOut,
          protocolFeeRate
        ),
      }))
      .filter(({ originQuery }) => originQuery.minAmountOut.gt(0))
  }

  /**
   * Filters the list of quotes to only include those that can be used to receive teh given destination token.
   * For every filtered quote, the origin query is returned with the information for tokenIn -> ticker swaps,
   * and the dest query with the information for ticker -> tokenOut swap.
   */
  private async filterQuotesWithDestSwap(
    destSIR: SynapseIntentRouter,
    originIntents: OriginIntent[],
    tokenOut: string
  ): Promise<DestIntent[]> {
    const destIntents = await Promise.all(
      originIntents.map(async ({ ticker, originQuery, originAmountOut }) => {
        const { amountOut: destAmountOut, steps } = await destSIR.previewIntent(
          ticker.destToken.token,
          tokenOut,
          originAmountOut,
          FORWARD_TO_SIMULATED
        )
        return {
          ticker,
          originQuery,
          originAmountOut,
          destAmountOut,
          steps,
        }
      })
    )
    // Up to a single Zap is supported
    return destIntents
      .filter(({ destAmountOut }) => destAmountOut.gt(0))
      .filter(({ steps }) => steps.length <= 1)
      .map(({ ticker, originQuery, originAmountOut, steps }) => ({
        ticker,
        originQuery,
        originAmountOut,
        encodedZapData: steps.length > 0 ? hexlify(steps[0].zapData) : '0x',
        tokenOut,
      }))
  }

  /**
   * Gets the full quote for fulfilling the intent.
   */
  private async getFullIntentQuote(
    destSIR: SynapseIntentRouter,
    intent: DestIntent,
    originUserAddress?: string
  ): Promise<FullIntent> {
    // Unwrap the intent. Note that the zap data was generated by using `originAmountOut` as the amount.
    const {
      ticker,
      originQuery,
      originAmountOut,
      encodedZapData: encodedZapDataSimulated,
      tokenOut,
    } = intent
    const quote = await getBestRelayerQuote(ticker, originAmountOut, {
      originSender: originUserAddress,
      destRecipient:
        encodedZapDataSimulated === '0x'
          ? undefined
          : TOKEN_ZAP_V1_ADDRESS_MAP[ticker.destToken.chainId],
      zapData: encodedZapDataSimulated,
    })
    // Now that we got the quote, we need to get the final amount out and adjust the zap data.
    // Note: zap data will still be using `FORWARD_TO_SIMULATED` address - this will be overwritten
    // when the bridge calldata is generated (until then we don't know the final recipient).
    const { amountOut, steps } = await destSIR.previewIntent(
      ticker.destToken.token,
      tokenOut,
      quote.destAmount,
      FORWARD_TO_SIMULATED
    )
    // As previously mentioned, up to a single Zap is supported.
    const destAmountOut = steps.length > 1 ? Zero : amountOut
    const encodedZapData = steps.length > 0 ? hexlify(steps[0].zapData) : '0x'
    const paramsV1 = originUserAddress
      ? {
          sender: originUserAddress,
          destToken: ticker.destToken.token,
          destAmount: quote.destAmount,
        }
      : undefined
    return {
      ticker,
      originQuery,
      originAmountOut,
      encodedZapData,
      tokenOut,
      destQuery: SynapseIntentRouterSet.createRFQDestQuery(
        destSIR,
        ticker.destToken.token,
        destAmountOut,
        {
          paramsV1,
          paramsV2: {
            // TODO: exclusivity
            quoteRelayer: AddressZero,
            quoteExclusivitySeconds: Zero,
            // TODO: quote ID
            quoteId: '0x',
            zapNative: Zero,
            zapData: encodedZapData,
          },
          zapData: decodeZapData(encodedZapData),
        }
      ),
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

  public static createRFQDestQuery(
    destSIR: SynapseIntentRouter,
    tokenOut: string,
    amountOut: BigNumber,
    savedParams: {
      paramsV1?: SavedParamsV1
      paramsV2: BridgeParamsV2
      zapData: Partial<ZapDataV1>
    }
  ): Query {
    // To preserve consistency with other modules, router adapter is not set for a no-op intent.
    const destQuery = createNoSwapQuery(tokenOut, amountOut)
    // Don't modify the Query if there are no params to be saved.
    if (!savedParams.paramsV1) {
      return destQuery
    }
    destQuery.routerAdapter = destSIR.address
    destQuery.deadline = MaxUint256
    destQuery.rawParams = encodeSavedBridgeParams(
      savedParams.paramsV1,
      savedParams.paramsV2,
      applyDefaultValues(savedParams.zapData)
    )
    return destQuery
  }
}
