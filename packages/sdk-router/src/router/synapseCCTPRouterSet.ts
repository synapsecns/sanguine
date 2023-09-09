import { Provider } from '@ethersproject/abstract-provider'

import { Abi } from '../utils/types'
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
    address: string,
    abi: Abi
  ): SynapseCCTPRouter {
    return new SynapseCCTPRouter(chainId, provider, address, abi)
  }

  /**
   * Returns the existing SynapseCCTPRouter instance for the given chain.
   */
  public getSynapseRouter(chainId: number): SynapseCCTPRouter {
    return this.routers[chainId] as SynapseCCTPRouter
  }
}
