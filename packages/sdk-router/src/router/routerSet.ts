import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber } from '@ethersproject/bignumber'
import invariant from 'tiny-invariant'

import { Router } from './router'
import { AddressMap, BigintIsh } from '../constants'
import { DestRequest } from './types'
import {
  BridgeRoute,
  FeeConfig,
  SynapseModule,
  SynapseModuleSet,
} from '../module'
import {
  applySlippageToQuery,
  Query,
  hasComplexBridgeAction,
} from '../module/query'
import { ONE_WEEK, TEN_MINUTES } from '../utils/deadlines'

export type ChainProvider = {
  chainId: number
  provider: Provider
}

export type RouterConstructor = new (
  chainId: number,
  provider: Provider,
  address: string
) => Router

/**
 * Abstract class for a set of routers existing on a few chains. Handles Router interactions
 * on a set of chains: the RouterSet users don't need to know which Router contract is used.
 *
 * The class children should provide the router addresses for each chain, as well as the Router constructor.
 *
 * @property bridgeModuleName The name of the bridge module used by the routers.
 * @property routers Collection of Router instances indexed by chainId.
 * @property providers Collection of Provider instances indexed by chainId.
 */
export abstract class RouterSet extends SynapseModuleSet {
  public routers: {
    [chainId: number]: Router
  }
  public providers: {
    [chainId: number]: Provider
  }

  /**
   * Constructor for a RouterSet class.
   * It creates the Router instances for each chain that has both a provider and a router address.
   */
  constructor(
    chains: ChainProvider[],
    addressMap: AddressMap,
    ctor: RouterConstructor
  ) {
    super()
    this.routers = {}
    this.providers = {}
    chains.forEach(({ chainId, provider }) => {
      const address = addressMap[chainId]
      // Skip chains without a router address
      if (address) {
        this.routers[chainId] = new ctor(chainId, provider, address)
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
   * @inheritdoc SynapseModuleSet.getBridgeRoutes
   */
  public async getBridgeRoutes(
    originChainId: number,
    destChainId: number,
    tokenIn: string,
    tokenOut: string,
    amountIn: BigintIsh
  ): Promise<BridgeRoute[]> {
    const originRouter = this.routers[originChainId]
    const destRouter = this.routers[destChainId]
    // Check if routers exist on both chains
    if (!originRouter || !destRouter) {
      return []
    }

    // First, get the list of bridge tokens that could be used to receive the output token.
    const bridgeTokens = await destRouter.getConnectedBridgeTokens(tokenOut)
    try {
      // Fetch queries from origin router: originQueries.length === bridgeTokens.length
      const originQueries = await originRouter.getOriginAmountOut(
        tokenIn,
        bridgeTokens.map((bridgeToken) => bridgeToken.symbol),
        amountIn
      )
      // Zip origin queries with bridge tokens and filter out origin queries with zero minAmountOut
      const originRoutes = originQueries
        .map((originQuery, index) => ({
          originQuery,
          bridgeToken: bridgeTokens[index],
        }))
        .filter((originRoute) => originRoute.originQuery.minAmountOut.gt(0))
      // Exit early if no routes on origin chain found
      if (!originRoutes.length) {
        return []
      }

      // Build destination requests: requests.length === originRoutes.length
      const requests: DestRequest[] = originRoutes.map((originRoute) => ({
        symbol: originRoute.bridgeToken.symbol,
        amountIn: originRoute.originQuery.minAmountOut,
      }))
      // Fetch destination queries: destQueries.length === originRoutes.length
      const destQueries = await destRouter.getDestinationQueries(
        requests,
        tokenOut
      )
      // Zip origin routes with destination queries
      const bridgeRoutes: BridgeRoute[] = originRoutes.map(
        (originRoute, index) => ({
          originChainId,
          destChainId,
          originQuery: originRoute.originQuery,
          destQuery: destQueries[index],
          bridgeToken: originRoute.bridgeToken,
          bridgeModuleName: this.bridgeModuleName,
        })
      )
      // Return routes with non-zero minAmountOut
      return bridgeRoutes.filter((bridgeRoute) =>
        bridgeRoute.destQuery.minAmountOut.gt(0)
      )
    } catch (error) {
      console.error(
        `Error when trying to calculate the best quote with bridge tokens: ${bridgeTokens}`,
        error
      )
      return []
    }
  }

  /**
   * @inheritdoc SynapseModuleSet.getFeeData
   */
  async getFeeData(bridgeRoute: BridgeRoute): Promise<{
    feeAmount: BigNumber
    feeConfig: FeeConfig
  }> {
    const destRouter = this.routers[bridgeRoute.destChainId]
    invariant(destRouter, 'Router not found')
    // Get fee data: for some Bridge contracts it will depend on the complexity of the bridge action
    return destRouter.getBridgeFees(
      bridgeRoute.bridgeToken.token,
      bridgeRoute.originQuery.minAmountOut,
      hasComplexBridgeAction(bridgeRoute.destQuery)
    )
  }

  /**
   * @inheritdoc SynapseModuleSet.getDefaultPeriods
   */
  getDefaultPeriods(): {
    originPeriod: number
    destPeriod: number
  } {
    // Use the same default periods for SynapseBridge and SynapseCCTP modules
    return {
      originPeriod: TEN_MINUTES,
      destPeriod: ONE_WEEK,
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
    return {
      originQuery: applySlippageToQuery(
        originQueryPrecise,
        slipNumerator,
        slipDenominator
      ),
      destQuery: applySlippageToQuery(
        destQueryPrecise,
        slipNumerator,
        slipDenominator
      ),
    }
  }
}
