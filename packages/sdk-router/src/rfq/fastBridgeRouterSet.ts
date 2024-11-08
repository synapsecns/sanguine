import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber } from '@ethersproject/bignumber'
import invariant from 'tiny-invariant'
import { Zero } from '@ethersproject/constants'

import {
  BigintIsh,
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
} from '../module'
import { FastBridgeRouter } from './fastBridgeRouter'
import { ChainProvider } from '../router'
import { ONE_HOUR, TEN_MINUTES } from '../utils/deadlines'
import { isSameAddress } from '../utils/addressUtils'
import { marshallTicker, Ticker } from './ticker'
import { getAllQuotes, getBestRFQQuote } from './api'

export class FastBridgeRouterSet extends SynapseModuleSet {
  static readonly MAX_QUOTE_AGE_MILLISECONDS = 5 * 60 * 1000 // 5 minutes

  public readonly bridgeModuleName = 'SynapseRFQ'
  public readonly allEvents = ['BridgeRequestedEvent', 'BridgeRelayedEvent']

  public routers: {
    [chainId: number]: FastBridgeRouter
  }
  public providers: {
    [chainId: number]: Provider
  }

  // The answer to life, the universe, and everything
  private readonly GAS_REBATE_FLAG = '0x2a'

  constructor(chains: ChainProvider[]) {
    super()
    this.routers = {}
    this.providers = {}
    chains.forEach(({ chainId, provider }) => {
      const address = FAST_BRIDGE_ROUTER_ADDRESS_MAP[chainId]
      // Skip chains without a FastBridgeRouter address
      if (address) {
        this.routers[chainId] = new FastBridgeRouter(chainId, provider, address)
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
  public async getGasDropAmount(bridgeRoute: BridgeRoute): Promise<BigNumber> {
    // TODO: test this once chainGasAmount is set to be non-zero
    if (
      bridgeRoute.destQuery.rawParams
        .toLowerCase()
        .startsWith(this.GAS_REBATE_FLAG)
    ) {
      return this.getFastBridgeRouter(bridgeRoute.destChainId).chainGasAmount()
    }
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
    // Get all tickers that can be used to fulfill the tokenIn -> tokenOut intent via RFQ
    const tickers = await this.getAllTickers(
      originChainId,
      destChainId,
      tokenOut
    )
    // Get queries for swaps on the origin chain from tokenIn into the "RFQ-supported token"
    const filteredTickers = await this.filterTickersWithPossibleSwap(
      originChainId,
      tokenIn,
      amountIn,
      tickers
    )
    const protocolFeeRate = await this.getFastBridgeRouter(
      originChainId
    ).getProtocolFeeRate()
    const quotes = await Promise.all(
      filteredTickers.map(async ({ ticker, originQuery }) => ({
        ticker,
        originQuery,
        quote: await getBestRFQQuote(
          ticker,
          // Get the quote for the proceeds of the origin swap with protocol fee applied
          this.applyProtocolFeeRate(originQuery.minAmountOut, protocolFeeRate),
          originUserAddress
        ),
      }))
    )
    return quotes
      .filter(({ quote }) => quote.destAmount.gt(0))
      .map(({ ticker, originQuery, quote }) => ({
        originChainId,
        destChainId,
        bridgeToken: {
          symbol: marshallTicker(ticker),
          token: ticker.destToken.token,
        },
        originQuery,
        destQuery: FastBridgeRouterSet.createRFQDestQuery(
          tokenOut,
          quote.destAmount,
          originUserAddress
        ),
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
   * Filters the list of tickers to only include those that can be used for given amount of input token.
   * For every filtered ticker, the origin query is returned with the information for tokenIn -> ticker swaps.
   */
  private async filterTickersWithPossibleSwap(
    originChainId: number,
    tokenIn: string,
    amountIn: BigintIsh,
    tickers: Ticker[]
  ): Promise<{ ticker: Ticker; originQuery: Query }[]> {
    // Get queries for swaps on the origin chain into the "RFQ-supported token"
    const originQueries = await this.getFastBridgeRouter(
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
      }))
      .filter(({ originQuery }) => originQuery.minAmountOut.gt(0))
  }

  /**
   * Get all unique tickers for a given origin chain and a destination token. In other words,
   * this is the list of all origin tokens that can be used to create a quote for a
   * swap to the given destination token, without duplicates.
   *
   * @param originChainId - The ID of the origin chain.
   * @param destChainId - The ID of the destination chain.
   * @param tokenOut - The final token of the cross-chain swap.
   * @returns A promise that resolves to the list of tickers.
   */
  private async getAllTickers(
    originChainId: number,
    destChainId: number,
    tokenOut: string
  ): Promise<Ticker[]> {
    const allQuotes = await getAllQuotes()
    const originFB = await this.getFastBridgeAddress(originChainId)
    const destFB = await this.getFastBridgeAddress(destChainId)
    // First, we filter out quotes for other chainIDs, bridges or destination token.
    // Then, we filter out quotes that are too old.
    // Finally, we remove the duplicates of the origin token.
    return allQuotes
      .filter((quote) => {
        const areSameChainsAndToken =
          quote.ticker.originToken.chainId === originChainId &&
          quote.ticker.destToken.chainId === destChainId &&
          isSameAddress(quote.originFastBridge, originFB) &&
          isSameAddress(quote.destFastBridge, destFB) &&
          isSameAddress(quote.ticker.destToken.token, tokenOut)
        const age = Date.now() - quote.updatedAt
        const isValidAge =
          0 <= age && age < FastBridgeRouterSet.MAX_QUOTE_AGE_MILLISECONDS
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
}
