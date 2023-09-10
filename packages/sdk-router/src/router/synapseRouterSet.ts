import { Provider } from '@ethersproject/abstract-provider'

import { SynapseRouter } from './synapseRouter'
import { RouterSet } from './routerSet'

export class SynapseRouterSet extends RouterSet {
  public readonly routerName = 'SynapseRouter'

  /**
   * Creates a new Router instance for the given chain.
   */
  public instantiateRouter(
    chainId: number,
    provider: Provider,
    address: string
  ): SynapseRouter {
    return new SynapseRouter(chainId, provider, address)
  }

  /**
   * Returns the existing SynapseRouter instance for the given chain.
   */
  public getSynapseRouter(chainId: number): SynapseRouter {
    return this.routers[chainId] as SynapseRouter
  }
}
