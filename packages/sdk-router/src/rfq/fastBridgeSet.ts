import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber } from '@ethersproject/bignumber'

import { AddressMap, BigintIsh } from '../constants'
import {
  BridgeRoute,
  BridgeToken,
  FeeConfig,
  SynapseModule,
  SynapseModuleSet,
  Query,
  createNoSwapQuery,
} from '../module'
import { ChainProvider } from '../router'
import { FastBridge } from './fastBridge'
import { Ticker, marshallTicker } from './ticker'

export class FastBridgeSet extends SynapseModuleSet {
  public readonly bridgeModuleName = 'SynapseRFQ'
  public readonly allEvents = ['BridgeRequestedEvent', 'BridgeRelayedEvent']

  public fastBridges: {
    [chainId: number]: FastBridge
  }
  public providers: {
    [chainId: number]: Provider
  }

  constructor(chains: ChainProvider[], addressMap: AddressMap) {
    super()
    this.fastBridges = {}
    this.providers = {}
    chains.forEach(({ chainId, provider }) => {
      const address = addressMap[chainId]
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
    // TODO: implement
    console.log(chainId)
    return 0
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
    const allTickers: Ticker[] = await this.getSupportedTickers(
      originChainId,
      destChainId
    )
    // Get all tickets that could fulfill the tokenIn -> tokenOut cross-chain swap
    const originRoutes = allTickers
      .filter(
        (ticker) =>
          // Filter tickers that have destination token matching tokenOut
          ticker.destToken.token.toLowerCase() === tokenOut.toLowerCase() &&
          // Check if the origin token matches tokenIn
          // TODO: adjust this once swaps on the origin chain are supported
          ticker.originToken.token.toLowerCase() === tokenIn.toLowerCase()
      )
      .map((ticker) => ({
        // TODO: Create Query struct for tokenIn -> ticker.originToken swap
        originQuery: createNoSwapQuery(tokenIn, BigNumber.from(amountIn)),
        ticker,
      }))
      .filter(({ originQuery }) => !originQuery.minAmountOut.isZero())
    // Figure out RFQ quotes for each origin route
    const destAmountOuts = await Promise.all(
      originRoutes.map(({ originQuery, ticker }) =>
        this.getQuote(ticker, originQuery.minAmountOut)
      )
    )
    // Zip originRoutes with destAmountOuts
    return originRoutes.map((originRoute, index) => {
      // Swaps on destination chain are not supported in RFQ
      const destQuery: Query = createNoSwapQuery(
        tokenOut,
        destAmountOuts[index]
      )
      // Use the ticker as "bridge symbol"
      const bridgeToken: BridgeToken = {
        symbol: marshallTicker(originRoute.ticker),
        token: originRoute.ticker.destToken.token,
      }
      return {
        originChainId,
        destChainId,
        bridgeToken,
        originQuery: originRoute.originQuery,
        destQuery,
        bridgeModuleName: this.bridgeModuleName,
      }
    })
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
   * Get the list of tickers supported by the RFQ between the two chains.
   *
   * @param originChainId - The ID of the origin chain.
   * @param destChainId - The ID of the destination chain.
   * @returns A promise that resolves to the list of supported tickers.
   */
  private async getSupportedTickers(
    originChainId: number,
    destChainId: number
  ): Promise<Ticker[]> {
    // TODO: hit Quoter API to get supported tickers, then unmarshall them
    console.log(originChainId, destChainId)
    return []
  }

  /**
   * Get the quote for a given ticker and amount.
   *
   * @param ticker - The ticker to get the quote for.
   * @param amount - The amount to get the quote for.
   * @returns A promise that resolves to the quote.
   */
  private async getQuote(
    ticker: Ticker,
    amount: BigNumber
  ): Promise<BigNumber> {
    // TODO: hit Quoter API to get the quote for ticker.originToken -> ticker.destToken with the given amount
    console.log(ticker, amount)
    return BigNumber.from(0)
  }
}
