import { Provider } from '@ethersproject/abstract-provider'

import { SynapseCCTPRouter } from './synapseCCTPRouter'
import { RouterSet } from './routerSet'

export class SynapseCCTPRouterSet extends RouterSet {
  public readonly routerName = 'SynapseCCTPRouter'

  /**
   * Creates a new Router instance for the given chain.
   */
  public instantiateRouter(
    chainId: number,
    provider: Provider,
    address: string
  ): SynapseCCTPRouter {
    return new SynapseCCTPRouter(chainId, provider, address)
  }

  /**
   * Returns the existing SynapseCCTPRouter instance for the given chain.
   */
  public getSynapseCCTPRouter(chainId: number): SynapseCCTPRouter {
    return this.routers[chainId] as SynapseCCTPRouter
  }
}
