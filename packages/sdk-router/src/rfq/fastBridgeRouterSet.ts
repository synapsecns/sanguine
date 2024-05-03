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
import { FastBridgeQuote, applyQuote } from './quote'
import { marshallTicker } from './ticker'
import { getAllQuotes } from './api'

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
    amountIn: BigintIsh
  ): Promise<BridgeRoute[]> {
    // Check that Routers exist on both chains
    if (!this.getModule(originChainId) || !this.getModule(destChainId)) {
      return []
    }
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
    const protocolFeeRate = await this.getFastBridgeRouter(
      originChainId
    ).getProtocolFeeRate()
    return filteredQuotes
      .map(({ quote, originQuery }) => ({
        quote,
        originQuery,
        // Apply quote to the proceeds of the origin swap with protocol fee applied
        // TODO: handle optional gas airdrop pricing
        destAmountOut: applyQuote(
          quote,
          this.applyProtocolFeeRate(originQuery.minAmountOut, protocolFeeRate)
        ),
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
        // TODO: signal optional gas airdrop
        destQuery: createNoSwapQuery(tokenOut, destAmountOut),
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
      destPeriod: 4 * ONE_HOUR,
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
   * Filters the list of quotes to only include those that can be used for given amount of input token.
   * For every filtered quote, the origin query is returned with the information for tokenIn -> RFQ token swaps.
   */
  private async filterOriginQuotes(
    originChainId: number,
    tokenIn: string,
    amountIn: BigintIsh,
    allQuotes: FastBridgeQuote[]
  ): Promise<{ quote: FastBridgeQuote; originQuery: Query }[]> {
    // Get queries for swaps on the origin chain into the "RFQ-supported token"
    const originQueries = await this.getFastBridgeRouter(
      originChainId
    ).getOriginAmountOut(
      tokenIn,
      allQuotes.map((quote) => quote.ticker.originToken.token),
      amountIn
    )
    // Note: allQuotes.length === originQueries.length
    // Zip the quotes and queries together, filter out "no path found" queries
    return allQuotes
      .map((quote, index) => ({
        quote,
        originQuery: originQueries[index],
      }))
      .filter(({ originQuery }) => originQuery.minAmountOut.gt(0))
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
    const originFB = await this.getFastBridgeAddress(originChainId)
    const destFB = await this.getFastBridgeAddress(destChainId)
    return allQuotes
      .filter(
        (quote) =>
          quote.ticker.originToken.chainId === originChainId &&
          quote.ticker.destToken.chainId === destChainId &&
          quote.ticker.destToken.token &&
          quote.ticker.destToken.token.toLowerCase() === tokenOut.toLowerCase()
      )
      .filter(
        (quote) =>
          quote.originFastBridge.toLowerCase() === originFB.toLowerCase() &&
          quote.destFastBridge.toLowerCase() === destFB.toLowerCase()
      )
      .filter((quote) => {
        const age = Date.now() - quote.updatedAt
        return 0 <= age && age < FastBridgeRouterSet.MAX_QUOTE_AGE_MILLISECONDS
      })
  }
}
