import { Provider } from '@ethersproject/abstract-provider'

import { SynapseCCTPRouter } from './synapseCCTPRouter'
import { ChainProvider, RouterSet } from './routerSet'
import { CCTP_ROUTER_ADDRESS_MAP } from '../constants'

export class SynapseCCTPRouterSet extends RouterSet {
  public readonly routerName = 'SynapseCCTPRouter'

  constructor(chains: ChainProvider[]) {
    super(chains, CCTP_ROUTER_ADDRESS_MAP)
  }

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
