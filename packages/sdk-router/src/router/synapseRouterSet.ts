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
   */
  public getSynapseRouter(chainId: number): SynapseRouter {
    return this.routers[chainId] as SynapseRouter
  }
}
