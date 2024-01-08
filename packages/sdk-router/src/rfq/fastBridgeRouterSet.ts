import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber } from '@ethersproject/bignumber'
import invariant from 'tiny-invariant'

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
} from '../module'
import { FastBridgeRouter } from './fastBridgeRouter'
import { ChainProvider } from '../router'
import { ONE_HOUR, TEN_MINUTES } from '../utils/deadlines'
import { FastBridgeQuote, applyQuote } from './quote'
import { marshallTicker } from './ticker'

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
   * @inheritdoc SynapseModuleSet.getBridgeRoutes
   */
  public async getBridgeRoutes(
    originChainId: number,
    destChainId: number,
    tokenIn: string,
    tokenOut: string,
    amountIn: BigintIsh
  ): Promise<BridgeRoute[]> {
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
    return filteredQuotes
      .map(({ quote, originQuery }) => ({
        quote,
        originQuery,
        // Apply quote to the proceeds of the origin swap
        destAmountOut: applyQuote(quote, originQuery.minAmountOut),
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
        // On-chain swaps are not supported for RFQ tokens
        destQuery: createNoSwapQuery(tokenOut, destAmountOut),
        bridgeModuleName: this.bridgeModuleName,
      }))
  }

  /**
   * @inheritdoc SynapseModuleSet.getFeeData
   */
  public async getFeeData(): Promise<{
    feeAmount: BigNumber
    feeConfig: FeeConfig
  }> {
    // TODO: figure out if we need to report anything here
    return {
      feeAmount: BigNumber.from(0),
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
      destPeriod: ONE_HOUR,
    }
  }

  /**
   * Filters the list of quotes to only include those that can be used for given amount of input token.
   * For every filtered quote, the origin query is returned with the information for tokenIn -> RFQ token swaps.
   */
  private async filterOriginQuotes(
    originChainId: number,
    tokenIn: string,
    amountIn: BigintIsh,
    allQuotes: FastBridgeQuote[]
  ): Promise<{ quote: FastBridgeQuote; originQuery: Query }[]> {
    // TODO: implement
    console.log(originChainId, tokenIn, amountIn, allQuotes)
    return []
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
    tokenOut: string
  ): Promise<FastBridgeQuote[]> {
    // TODO: implement
    console.log(originChainId, destChainId, tokenOut)
    return []
  }
}
