import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber } from '@ethersproject/bignumber'
import invariant from 'tiny-invariant'

import { Router } from './router'
import { AddressMap, BigintIsh } from '../constants'
import { BridgeQuote, BridgeRoute, DestRequest } from './types'
import { ONE_WEEK, TEN_MINUTES, calculateDeadline } from '../utils/deadlines'
import { SynapseModule, SynapseModuleSet } from '../module'
import { hasComplexBridgeAction } from './query'

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
  public getModule(
    chainId: number,
    moduleAddress: string
  ): SynapseModule | undefined {
    const router = this.routers[chainId]
    // Check if router exists on chain and that router address matches
    if (router?.address.toLowerCase() === moduleAddress.toLowerCase()) {
      return router
    } else {
      return undefined
    }
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
          originRouterAddress: originRouter.address,
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
   * @inheritdoc SynapseModuleSet.finalizeBridgeRoute
   */
  public async finalizeBridgeRoute(
    bridgeRoute: BridgeRoute,
    deadline?: BigNumber
  ): Promise<BridgeQuote> {
    const originRouter = this.routers[bridgeRoute.originChainId]
    const destRouter = this.routers[bridgeRoute.destChainId]
    invariant(originRouter && destRouter, 'Route not supported')
    const { originQuery, destQuery, bridgeToken } = bridgeRoute
    // Set origin deadline to 10 mins if not provided
    originQuery.deadline = deadline ?? calculateDeadline(TEN_MINUTES)
    // Destination deadline is always 1 week
    destQuery.deadline = calculateDeadline(ONE_WEEK)
    // Get fee data: for some Bridge contracts it will depend on the complexity of the bridge action
    const { feeAmount, feeConfig } = await destRouter.getBridgeFees(
      bridgeToken.token,
      originQuery.minAmountOut,
      hasComplexBridgeAction(destQuery)
    )
    const estimatedTime = this.getEstimatedTime(bridgeRoute.originChainId)
    return {
      feeAmount,
      feeConfig,
      routerAddress: originRouter.address,
      maxAmountOut: destQuery.minAmountOut,
      originQuery,
      destQuery,
      estimatedTime,
      bridgeModuleName: this.bridgeModuleName,
    }
  }
}
