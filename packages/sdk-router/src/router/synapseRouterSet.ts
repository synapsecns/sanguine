import invariant from 'tiny-invariant'

import { SynapseRouter } from './synapseRouter'
import { ChainProvider, RouterSet } from './routerSet'
import { ROUTER_ADDRESS_MAP } from '../constants'

/**
 * Wrapper class for interacting with a SynapseRouter contracts deployed on multiple chains.
 */
export class SynapseRouterSet extends RouterSet {
  public readonly routerName = 'SynapseRouter'

  constructor(chains: ChainProvider[]) {
    super(chains, ROUTER_ADDRESS_MAP, SynapseRouter)
  }

  /**
   * Returns the existing SynapseRouter instance for the given chain.
   *
   * @throws Will throw an error if SynapseRouter is not deployed on the given chain.
   */
  public getSynapseRouter(chainId: number): SynapseRouter {
    invariant(
      this.routers[chainId],
      `No SynapseRouter deployed on chain ${chainId}`
    )
    return this.routers[chainId] as SynapseRouter
  }
}
