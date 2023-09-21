import invariant from 'tiny-invariant'

import { SynapseCCTPRouter } from './synapseCCTPRouter'
import { ChainProvider, RouterSet } from './routerSet'
import { CCTP_ROUTER_ADDRESS_MAP } from '../constants'

/**
 * Wrapper class for interacting with a SynapseCCTPRouter contracts deployed on multiple chains.
 */
export class SynapseCCTPRouterSet extends RouterSet {
  public readonly routerName = 'SynapseCCTPRouter'

  constructor(chains: ChainProvider[]) {
    super(chains, CCTP_ROUTER_ADDRESS_MAP, SynapseCCTPRouter)
  }

  /**
   * Returns the existing SynapseCCTPRouter instance for the given chain.
   *
   * @throws Will throw an error if SynapseCCTPRouter is not deployed on the given chain.
   */
  public getSynapseCCTPRouter(chainId: number): SynapseCCTPRouter {
    invariant(
      this.routers[chainId],
      `No SynapseCCTPRouter deployed on chain ${chainId}`
    )
    return this.routers[chainId] as SynapseCCTPRouter
  }
}
