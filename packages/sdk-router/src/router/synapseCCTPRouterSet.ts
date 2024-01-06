import invariant from 'tiny-invariant'
import { BigNumber } from '@ethersproject/bignumber'

import { SynapseCCTPRouter } from './synapseCCTPRouter'
import { ChainProvider, RouterSet } from './routerSet'
import { CCTP_ROUTER_ADDRESS_MAP, MEDIAN_TIME_CCTP } from '../constants'
import { BridgeRoute } from '../module'

/**
 * Wrapper class for interacting with a SynapseCCTPRouter contracts deployed on multiple chains.
 */
export class SynapseCCTPRouterSet extends RouterSet {
  public readonly bridgeModuleName = 'SynapseCCTP'
  public readonly allEvents = [
    'CircleRequestSentEvent',
    'CircleRequestFulfilledEvent',
  ]

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
  getGasDropAmount(bridgeRoute: BridgeRoute): Promise<BigNumber> {
    return this.getSynapseCCTPRouter(bridgeRoute.destChainId).chainGasAmount()
  }

  /**
   * Returns the existing SynapseCCTPRouter instance for the given chain.
   *
   * @throws Will throw an error if SynapseCCTPRouter is not deployed on the given chain.
   */
  public getSynapseCCTPRouter(chainId: number): SynapseCCTPRouter {
    return this.getExistingModule(chainId) as SynapseCCTPRouter
  }
}
