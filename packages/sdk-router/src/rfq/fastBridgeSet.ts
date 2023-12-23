import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber } from '@ethersproject/bignumber'
import invariant from 'tiny-invariant'

import {
  BigintIsh,
  FAST_BRIDGE_ADDRESS_MAP,
  MEDIAN_TIME_RFQ,
} from '../constants'
import {
  BridgeRoute,
  FeeConfig,
  SynapseModule,
  SynapseModuleSet,
  Query,
  createNoSwapQuery,
} from '../module'
import { ChainProvider } from '../router'
import { FastBridge } from './fastBridge'
import { marshallTicker } from './ticker'
import { FastBridgeQuote, applyQuote } from './quote'
import { getAllQuotes } from './api'

export class FastBridgeSet extends SynapseModuleSet {
  public readonly bridgeModuleName = 'SynapseRFQ'
  public readonly allEvents = ['BridgeRequestedEvent', 'BridgeRelayedEvent']

  public fastBridges: {
    [chainId: number]: FastBridge
  }
  public providers: {
    [chainId: number]: Provider
  }

  constructor(chains: ChainProvider[]) {
    super()
    this.fastBridges = {}
    this.providers = {}
    chains.forEach(({ chainId, provider }) => {
      const address = FAST_BRIDGE_ADDRESS_MAP[chainId]
      // Skip chains without a FastBridge address
      if (address) {
        this.fastBridges[chainId] = new FastBridge(chainId, provider, address)
        this.providers[chainId] = provider
      }
    })
  }

  /**
   * @inheritdoc SynapseModuleSet.getModule
   */
  public getModule(chainId: number): SynapseModule | undefined {
    return this.fastBridges[chainId]
  }

  /**
   * @inheritdoc RouterSet.getOriginAmountOut
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
    const originQueries = await this.getOriginQueries(
      originChainId,
      tokenIn,
      amountIn,
      allQuotes.map((quote) => quote.ticker.originToken.token)
    )
    return allQuotes
      .map((quote, index) => ({
        quote,
        originQuery: originQueries[index],
        // Apply quote to the proceeds of the origin swap
        destAmountOut: applyQuote(quote, originQueries[index].minAmountOut),
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
  async getFeeData(bridgeRoute: BridgeRoute): Promise<{
    feeAmount: BigNumber
    feeConfig: FeeConfig
  }> {
    // TODO: implement
    console.log(bridgeRoute)
    return null as any
  }

  /**
   * @inheritdoc SynapseModuleSet.getDefaultPeriods
   */
  getDefaultPeriods(): {
    originPeriod: number
    destPeriod: number
  } {
    // TODO: implement
    return {
      originPeriod: 0,
      destPeriod: 0,
    }
  }

  /**
   * Returns the origin query for every "RFQ token" on the origin chain,
   * that contains the information for tokenIn -> RFQ token swaps.
   */
  private async getOriginQueries(
    originChainId: number,
    tokenIn: string,
    amountIn: BigintIsh,
    rfqTokens: string[]
  ): Promise<Query[]> {
    // Check if the RFQ token matches tokenIn
    // TODO: change this to "find best path" once swaps on the origin chain are supported
    invariant(originChainId, 'Origin chain ID is required')
    return rfqTokens
      .filter((rfqToken) => rfqToken.toLowerCase() === tokenIn.toLowerCase())
      .map((rfqToken) => createNoSwapQuery(rfqToken, BigNumber.from(amountIn)))
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
    const allQuotes = await getAllQuotes()
    return allQuotes.filter(
      (quote) =>
        quote.ticker.originToken.chainId === originChainId &&
        quote.ticker.destToken.chainId === destChainId &&
        quote.ticker.destToken.token.toLowerCase() === tokenOut.toLowerCase()
    )
  }
}
