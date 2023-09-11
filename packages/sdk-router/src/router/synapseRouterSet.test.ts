import { Provider } from '@ethersproject/abstract-provider'
import { providers } from 'ethers'

import {
  PUBLIC_PROVIDER_URLS,
  ROUTER_ADDRESS_MAP,
  CCTP_ROUTER_ADDRESS_MAP,
  SupportedChainId,
} from '../constants'
import { ChainProvider } from './routerSet'
import { SynapseRouterSet } from './synapseRouterSet'

describe('SynapseRouterSet', () => {
  const ethProvider: Provider = new providers.JsonRpcProvider(
    PUBLIC_PROVIDER_URLS[SupportedChainId.ETH]
  )

  const arbProvider: Provider = new providers.JsonRpcProvider(
    PUBLIC_PROVIDER_URLS[SupportedChainId.ARBITRUM]
  )

  const testProviders: ChainProvider[] = [
    {
      chainId: SupportedChainId.ETH,
      provider: ethProvider,
    },
    {
      chainId: SupportedChainId.ARBITRUM,
      provider: arbProvider,
    },
  ]

  const routerSet = new SynapseRouterSet(testProviders)

  describe('#constructor', () => {
    it('Creates SynapseRouter instances for each chain', () => {
      expect(routerSet.routers[SupportedChainId.ETH]).toBeDefined()
      expect(routerSet.routers[SupportedChainId.ARBITRUM]).toBeDefined()
    })

    it('Uses correct addresses for each chain', () => {
      expect(routerSet.routers[SupportedChainId.ETH].address).toEqual(
        ROUTER_ADDRESS_MAP[SupportedChainId.ETH]
      )
      expect(routerSet.routers[SupportedChainId.ARBITRUM].address).toEqual(
        ROUTER_ADDRESS_MAP[SupportedChainId.ARBITRUM]
      )
    })

    it('Does not create SynapseRouter instances for chains without providers', () => {
      expect(routerSet.routers[SupportedChainId.AVALANCHE]).toBeUndefined()
    })
  })

  describe('getRouter', () => {
    it('Returns the correct router given correct address', () => {
      expect(
        routerSet.getRouter(
          SupportedChainId.ETH,
          ROUTER_ADDRESS_MAP[SupportedChainId.ETH]
        )
      ).toEqual(routerSet.routers[SupportedChainId.ETH])
    })

    it('Returns undefined given incorrect address', () => {
      expect(
        routerSet.getRouter(
          SupportedChainId.ETH,
          CCTP_ROUTER_ADDRESS_MAP[SupportedChainId.ETH]
        )
      ).toBeUndefined()
    })

    it('Returns undefined given unknown chainId', () => {
      expect(
        routerSet.getRouter(
          SupportedChainId.AVALANCHE,
          ROUTER_ADDRESS_MAP[SupportedChainId.ETH]
        )
      ).toBeUndefined()
    })
  })

  // TODO (Chi): more tests
})
