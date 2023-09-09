import { Provider } from '@ethersproject/abstract-provider'

import { Router } from './router'
import { Abi } from '../utils/types'

export type AddressMap = {
  [chainId: number]: string
}

export type ChainProvider = {
  chainId: number
  provider: Provider
}

/**
 * Abstract class for a set of routers existing on a few chains.
 *
 * @property routerName The name of the router set.
 * @property routers Collection of Router instances indexed by chainId.
 * @property providers Collection of Provider instances indexed by chainId.
 */
export abstract class RouterSet {
  abstract readonly routerName: string

  public routers: {
    [chainId: number]: Router
  }
  public providers: {
    [chainId: number]: Provider
  }

  /**
   * Constructor for a RouterSet class.
   * It creates the Router instances for each chain that has both a provider and a router address.
   */
  constructor(chains: ChainProvider[], addresses: AddressMap, abi: Abi) {
    this.routers = {}
    this.providers = {}
    chains.forEach(({ chainId, provider }) => {
      const address = addresses[chainId]
      // Skip chains without a router address
      if (address) {
        this.routers[chainId] = this.instantiateRouter(
          chainId,
          provider,
          address,
          abi
        )
        this.providers[chainId] = provider
      }
    })
  }

  /**
   * Creates a new Router instance for the given chain.
   */
  abstract instantiateRouter(
    chainId: number,
    provider: Provider,
    address: string,
    abi: Abi
  ): Router
}
