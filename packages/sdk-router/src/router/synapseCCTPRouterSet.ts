import { BigNumber } from '@ethersproject/bignumber'
import invariant from 'tiny-invariant'

import { ChainProvider, RouterSet } from './routerSet'
import { SynapseCCTPRouter } from './synapseCCTPRouter'
import { CCTP_ROUTER_ADDRESS_MAP, MEDIAN_TIME_CCTP } from '../constants'
import { BridgeRouteV2, BridgeTokenCandidate } from '../module'
import { logExecutionTime } from '../utils'

/**
 * Wrapper class for interacting with a SynapseCCTPRouter contracts deployed on multiple chains.
 */
export class SynapseCCTPRouterSet extends RouterSet {
  public readonly moduleName = 'SynapseCCTP'
  public readonly allEvents = [
    'CircleRequestSentEvent',
    'CircleRequestFulfilledEvent',
  ]
  public readonly isBridgeV2Supported = false

  constructor(chains: ChainProvider[]) {
    super(chains, CCTP_ROUTER_ADDRESS_MAP, SynapseCCTPRouter)
  }

  /**
   * @inheritdoc RouterSet.getOriginAmountOut
   */
  public getEstimatedTime(chainId: number): number {
    const medianTime =
      MEDIAN_TIME_CCTP[chainId as keyof typeof MEDIAN_TIME_CCTP]
    invariant(medianTime, `No estimated time for chain ${chainId}`)
    return medianTime
  }

  /**
   * @inheritdoc SynapseModuleSet.getGasDropAmount
   */
  public async getGasDropAmount(destChainId: number): Promise<BigNumber> {
    return this.getSynapseCCTPRouter(destChainId).chainGasAmount()
  }

  @logExecutionTime('SynapseCCTPRouterSet.getBridgeRoutes')
  public async getBridgeRoutes(
    ...args: Parameters<RouterSet['getBridgeRoutes']>
  ): ReturnType<RouterSet['getBridgeRoutes']> {
    return super.getBridgeRoutes(...args)
  }

  /**
   * Returns the existing SynapseCCTPRouter instance for the given chain.
   *
   * @throws Will throw an error if SynapseCCTPRouter is not deployed on the given chain.
   */
  public getSynapseCCTPRouter(chainId: number): SynapseCCTPRouter {
    return this.getExistingModule(chainId) as SynapseCCTPRouter
  }

  public async getBridgeTokenCandidates(): Promise<BridgeTokenCandidate[]> {
    return []
  }

  public async getBridgeRouteV2(): Promise<BridgeRouteV2> {
    throw new Error('BridgeRouteV2 is not supported by ' + this.moduleName)
  }
}
