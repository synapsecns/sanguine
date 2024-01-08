import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber } from '@ethersproject/bignumber'

import { BigintIsh, FAST_BRIDGE_ROUTER_ADDRESS_MAP } from '../constants'
import {
  BridgeRoute,
  FeeConfig,
  SynapseModule,
  SynapseModuleSet,
} from '../module'
import { FastBridgeRouter } from './fastBridgeRouter'
import { ChainProvider } from '../router'

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
    // TODO: implement
    console.log(chainId)
    return null as any
  }

  /**
   * @inheritdoc SynapseModuleSet.getOriginAmountOut
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
    // TODO: implement
    console.log(originChainId, destChainId, tokenIn, tokenOut, amountIn)
    return []
  }

  /**
   * @inheritdoc SynapseModuleSet.getFeeData
   */
  async getFeeData(): Promise<{
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
}
