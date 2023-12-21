import { Provider } from '@ethersproject/abstract-provider'
import { providers } from 'ethers'

import {
  getTestProviderUrl,
  ROUTER_ADDRESS_MAP,
  CCTP_ROUTER_ADDRESS_MAP,
  MEDIAN_TIME_BRIDGE,
  SUPPORTED_CHAIN_IDS,
  SupportedChainId,
} from '../constants'
import { ChainProvider } from './routerSet'
import { SynapseRouterSet } from './synapseRouterSet'
import { SynapseRouter } from './synapseRouter'

describe('SynapseRouterSet', () => {
  const ethProvider: Provider = new providers.JsonRpcProvider(
    getTestProviderUrl(SupportedChainId.ETH)
  )

  const arbProvider: Provider = new providers.JsonRpcProvider(
    getTestProviderUrl(SupportedChainId.ARBITRUM)
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

    it('Correct bridge module name', () => {
      expect(routerSet.bridgeModuleName).toEqual('SynapseBridge')
    })
  })

  describe('getEstimatedTime', () => {
    it('Returns the correct estimated time for all supported chains', () => {
      SUPPORTED_CHAIN_IDS.forEach((chainId) => {
        expect(routerSet.getEstimatedTime(Number(chainId))).toEqual(
          MEDIAN_TIME_BRIDGE[chainId as keyof typeof MEDIAN_TIME_BRIDGE]
        )
      })
    })

    it('Throws error for unsupported chain', () => {
      // 5 is the chain ID for Goerli testnet
      expect(() => routerSet.getEstimatedTime(5)).toThrow(
        'No estimated time for chain 5'
      )
    })
  })

  describe('getModuleWithAddress', () => {
    it('Returns the correct router given correct address', () => {
      expect(
        routerSet.getModuleWithAddress(
          SupportedChainId.ETH,
          ROUTER_ADDRESS_MAP[SupportedChainId.ETH]
        )
      ).toEqual(routerSet.routers[SupportedChainId.ETH])
    })

    it('Returns undefined given incorrect address', () => {
      expect(
        routerSet.getModuleWithAddress(
          SupportedChainId.ETH,
          CCTP_ROUTER_ADDRESS_MAP[SupportedChainId.ETH]
        )
      ).toBeUndefined()
    })

    it('Returns undefined given unknown chainId', () => {
      expect(
        routerSet.getModuleWithAddress(
          SupportedChainId.AVALANCHE,
          ROUTER_ADDRESS_MAP[SupportedChainId.ETH]
        )
      ).toBeUndefined()
    })
  })

  describe('getSynapseRouter', () => {
    it('Returns the correct router for supported chain', () => {
      const router = routerSet.getSynapseRouter(SupportedChainId.ETH)
      expect(router).toEqual(routerSet.routers[SupportedChainId.ETH])
      expect(router).toBeInstanceOf(SynapseRouter)
    })

    it('Throws error for unsupported chain', () => {
      expect(() =>
        routerSet.getSynapseRouter(SupportedChainId.AVALANCHE)
      ).toThrow('No module found for chain 43114')
    })
  })

  // TODO (Chi): more tests
})
