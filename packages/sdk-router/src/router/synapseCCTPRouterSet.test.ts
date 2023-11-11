import { Provider } from '@ethersproject/abstract-provider'
import { providers } from 'ethers'
import { describe, expect, it } from 'vitest'

import {
  PUBLIC_PROVIDER_URLS,
  ROUTER_ADDRESS_MAP,
  CCTP_ROUTER_ADDRESS_MAP,
  MEDIAN_TIME_CCTP,
  CCTP_SUPPORTED_CHAIN_IDS,
  SupportedChainId,
} from '../constants'
import { ChainProvider } from './routerSet'
import { SynapseCCTPRouterSet } from './synapseCCTPRouterSet'
import { SynapseCCTPRouter } from './synapseCCTPRouter'

describe('SynapseCCTPRouterSet', () => {
  const ethProvider: Provider = new providers.JsonRpcProvider(
    PUBLIC_PROVIDER_URLS[SupportedChainId.ETH]
  )

  const arbProvider: Provider = new providers.JsonRpcProvider(
    PUBLIC_PROVIDER_URLS[SupportedChainId.ARBITRUM]
  )

  // Chain where CCTP is unlikely to be deployed
  const moonbeamProvider: Provider = new providers.JsonRpcProvider(
    PUBLIC_PROVIDER_URLS[SupportedChainId.MOONBEAM]
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
    {
      chainId: SupportedChainId.MOONBEAM,
      provider: moonbeamProvider,
    },
  ]

  const routerSet = new SynapseCCTPRouterSet(testProviders)

  describe('#constructor', () => {
    it('Creates SynapseCCTPRouter instances for chains with CCTP', () => {
      expect(routerSet.routers[SupportedChainId.ETH]).toBeDefined()
      expect(routerSet.routers[SupportedChainId.ARBITRUM]).toBeDefined()
    })

    it('Uses correct addresses for each chain', () => {
      expect(routerSet.routers[SupportedChainId.ETH].address).toEqual(
        CCTP_ROUTER_ADDRESS_MAP[SupportedChainId.ETH]
      )
      expect(routerSet.routers[SupportedChainId.ARBITRUM].address).toEqual(
        CCTP_ROUTER_ADDRESS_MAP[SupportedChainId.ARBITRUM]
      )
    })

    it('Does not create SynapseCCTPRouter instances for chains without CCTP', () => {
      expect(routerSet.routers[SupportedChainId.MOONBEAM]).toBeUndefined()
    })

    it('Does not create SynapseCCTPRouter instances for chains without providers', () => {
      expect(routerSet.routers[SupportedChainId.AVALANCHE]).toBeUndefined()
    })

    it('Correct bridge module name', () => {
      expect(routerSet.bridgeModuleName).toEqual('SynapseCCTP')
    })
  })

  describe('getEstimatedTime', () => {
    it('Returns the correct estimated time for all supported chains', () => {
      CCTP_SUPPORTED_CHAIN_IDS.forEach((chainId) => {
        expect(routerSet.getEstimatedTime(Number(chainId))).toEqual(
          MEDIAN_TIME_CCTP[chainId as keyof typeof MEDIAN_TIME_CCTP]
        )
      })
    })

    it('Throws error for unsupported chain with a provider', () => {
      expect(() =>
        routerSet.getEstimatedTime(SupportedChainId.MOONBEAM)
      ).toThrow('No estimated time for chain 1284')
    })

    it('Throws error for unsupported chain without a provider', () => {
      // 5 is the chain ID for Goerli testnet
      expect(() => routerSet.getEstimatedTime(5)).toThrow(
        'No estimated time for chain 5'
      )
    })
  })

  describe('getRouter', () => {
    it('Returns the correct router given correct address', () => {
      expect(
        routerSet.getRouter(
          SupportedChainId.ETH,
          CCTP_ROUTER_ADDRESS_MAP[SupportedChainId.ETH]
        )
      ).toEqual(routerSet.routers[SupportedChainId.ETH])
    })

    it('Returns undefined given incorrect address', () => {
      expect(
        routerSet.getRouter(
          SupportedChainId.ETH,
          ROUTER_ADDRESS_MAP[SupportedChainId.ETH]
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

    it('Returns undefined given chainId without CCTP', () => {
      expect(
        routerSet.getRouter(
          SupportedChainId.MOONBEAM,
          CCTP_ROUTER_ADDRESS_MAP[SupportedChainId.ETH]
        )
      ).toBeUndefined()
    })
  })

  describe('getSynapseCCTPRouter', () => {
    it('Returns the correct router for supported chain', () => {
      const router = routerSet.getSynapseCCTPRouter(SupportedChainId.ETH)
      expect(router).toEqual(routerSet.routers[SupportedChainId.ETH])
      expect(router).toBeInstanceOf(SynapseCCTPRouter)
    })

    it('Throws error for unsupported chain', () => {
      expect(() =>
        routerSet.getSynapseCCTPRouter(SupportedChainId.AVALANCHE)
      ).toThrow('No SynapseCCTPRouter deployed on chain 43114')
    })
  })

  // TODO (Chi): more tests
})
