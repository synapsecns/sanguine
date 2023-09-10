import { Provider } from '@ethersproject/abstract-provider'

import { SynapseRouter } from './synapseRouter'
import { RouterSet } from './routerSet'
import { ROUTER_ADDRESS_MAP } from '../constants'

export class SynapseRouterSet extends RouterSet {
  public readonly routerName = 'SynapseRouter'
  public readonly addressMap = ROUTER_ADDRESS_MAP

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
