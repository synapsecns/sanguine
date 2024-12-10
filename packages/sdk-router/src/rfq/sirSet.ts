import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber } from '@ethersproject/bignumber'
import invariant from 'tiny-invariant'
import { AddressZero, Zero } from '@ethersproject/constants'

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
import { BridgeParamsV2, encodeSavedBridgeParams } from './paramsV2'
import { applyDefaultValues, ZapDataV1 } from './zapData'

type OriginIntent = {
  ticker: Ticker
  originQuery: Query
  originAmountOut: BigNumber
  destQuery?: Query
}

type Intent = Required<OriginIntent>

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
    // Get all tickers that can be used between the two chains
    const tickers = await this.getAllTickers(originChainId, destChainId)
    // Get queries for swaps on the origin chain from tokenIn into the "RFQ-supported token"
    const originQuotes = await this.filterTickersWithOriginSwap(
      originChainId,
      tokenIn,
      amountIn,
      tickers
    )
    // Get queries for swaps on the destination chain from the "RFQ-supported token" into tokenOut
    const intents: Intent[] = await this.filterQuotesWithDestSwap(
      destChainId,
      tokenOut,
      originQuotes,
      originUserAddress
    )
    // Apply the quotes from the RFQ API
    const fullQuotes = await this.applyQuotes(intents, originUserAddress)
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
    originChainId: number,
    tokenIn: string,
    amountIn: BigintIsh,
    tickers: Ticker[]
  ): Promise<OriginIntent[]> {
    const protocolFeeRate = await this.getSynapseIntentRouter(
      originChainId
    ).getProtocolFeeRate()
    // Get queries for swaps on the origin chain into the "RFQ-supported token"
    const originQueries = await this.getSynapseIntentRouter(
      originChainId
    ).getOriginAmountOut(
      tokenIn,
      tickers.map((ticker) => ticker.originToken.token),
      amountIn
    )
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
    destChainId: number,
    tokenOut: string,
    quotes: OriginIntent[],
    originUserAddress?: string
  ): Promise<Intent[]> {
    // TODO: actually find swaps, for now filter by tokenOut
    return quotes
      .filter((quote) => quote.ticker.destToken.chainId === destChainId)
      .filter((quote) => isSameAddress(quote.ticker.destToken.token, tokenOut))
      .map((quote) => ({
        ...quote,
        destQuery: SynapseIntentRouterSet.createRFQDestQuery(
          tokenOut,
          quote.originAmountOut,
          {
            // TODO: non-default values
            originUserAddress,
            paramsV2: {
              quoteRelayer: AddressZero,
              quoteExclusivitySeconds: Zero,
              quoteId: '0x',
              zapNative: Zero,
              zapData: '0x',
            },
            zapData: {},
          }
        ),
      }))
  }

  /**
   * Given a list of intents, fetches the best quote for each intent. THe quotes are recorded
   * as `destQuery.minAmountOut`, while the other fields are left unchanged.
   */
  private async applyQuotes(
    intents: Intent[],
    originUserAddress?: string
  ): Promise<Intent[]> {
    return Promise.all(
      intents.map(
        async ({ ticker, originQuery, destQuery, originAmountOut }) => {
          const quote = await getBestRelayerQuote(ticker, originAmountOut, {
            originSender: originUserAddress,
          })
          const destQueryAdjusted: Query = {
            ...destQuery,
            minAmountOut: quote.destAmount,
          }
          return {
            ticker,
            originQuery,
            originAmountOut,
            destQuery: destQueryAdjusted,
          }
        }
      )
    )
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
    tokenOut: string,
    amountOut: BigNumber,
    savedParams: {
      originUserAddress?: string
      paramsV2?: BridgeParamsV2
      zapData: Partial<ZapDataV1>
    }
  ): Query {
    // On-chain swaps are not supported for RFQ on the destination chain
    const destQuery = createNoSwapQuery(tokenOut, amountOut)
    // Don't modify the Query if there are no params to be saved
    if (!savedParams.originUserAddress || !savedParams.paramsV2) {
      return destQuery
    }
    destQuery.rawParams = encodeSavedBridgeParams(
      savedParams.originUserAddress,
      savedParams.paramsV2,
      applyDefaultValues(savedParams.zapData)
    )
    return destQuery
  }
}
