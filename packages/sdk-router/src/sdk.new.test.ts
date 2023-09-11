import { Provider } from '@ethersproject/abstract-provider'
import { providers } from 'ethers'

import { SynapseSDK } from './sdk'
import { PUBLIC_PROVIDER_URLS, SupportedChainId } from './constants'

describe('SynapseSDK', () => {
  const ethProvider: Provider = new providers.JsonRpcProvider(
    PUBLIC_PROVIDER_URLS[SupportedChainId.ETH]
  )

  const arbProvider: Provider = new providers.JsonRpcProvider(
    PUBLIC_PROVIDER_URLS[SupportedChainId.ARBITRUM]
  )

  const avaxProvider: Provider = new providers.JsonRpcProvider(
    PUBLIC_PROVIDER_URLS[SupportedChainId.AVALANCHE]
  )

  const optProvider: Provider = new providers.JsonRpcProvider(
    PUBLIC_PROVIDER_URLS[SupportedChainId.OPTIMISM]
  )

  const bscProvider: Provider = new providers.JsonRpcProvider(
    PUBLIC_PROVIDER_URLS[SupportedChainId.BSC]
  )

  // Chain where CCTP is unlikely to be deployed
  const moonbeamProvider: Provider = new providers.JsonRpcProvider(
    PUBLIC_PROVIDER_URLS[SupportedChainId.MOONBEAM]
  )

  describe('#constructor', () => {
    const synapse = new SynapseSDK(
      [
        SupportedChainId.ETH,
        SupportedChainId.ARBITRUM,
        SupportedChainId.MOONBEAM,
      ],
      [ethProvider, arbProvider, moonbeamProvider]
    )

    it('fails with unequal amount of chains to providers', () => {
      const chainIds = [SupportedChainId.ETH, SupportedChainId.ARBITRUM]
      const testProviders = [ethProvider]
      expect(() => new SynapseSDK(chainIds, testProviders)).toThrow(
        'Amount of chains and providers does not equal'
      )
    })

    it('Instantiates SynapseRouters for each chain', () => {
      expect(synapse.synapseRouterSet).toBeDefined()
      expect(
        synapse.synapseRouterSet.routers[SupportedChainId.ETH]
      ).toBeDefined()
      expect(
        synapse.synapseRouterSet.routers[SupportedChainId.ARBITRUM]
      ).toBeDefined()
      expect(
        synapse.synapseRouterSet.routers[SupportedChainId.MOONBEAM]
      ).toBeDefined()
    })

    it('Does not instantiate SynapseRouters for chains without providers', () => {
      expect(
        synapse.synapseRouterSet.routers[SupportedChainId.AVALANCHE]
      ).toBeUndefined()
    })

    it('Instantiates SynapseCCTPRouters for each chain with CCTP', () => {
      expect(synapse.synapseCCTPRouterSet).toBeDefined()
      expect(
        synapse.synapseCCTPRouterSet.routers[SupportedChainId.ETH]
      ).toBeDefined()
      expect(
        synapse.synapseCCTPRouterSet.routers[SupportedChainId.ARBITRUM]
      ).toBeDefined()
    })

    it('Does not instantiate SynapseCCTPRouters for chains without CCTP', () => {
      expect(
        synapse.synapseCCTPRouterSet.routers[SupportedChainId.MOONBEAM]
      ).toBeUndefined()
    })

    it('Does not instantiate SynapseCCTPRouters for chains without providers', () => {
      expect(
        synapse.synapseCCTPRouterSet.routers[SupportedChainId.AVALANCHE]
      ).toBeUndefined()
    })

    it('Saves providers', () => {
      expect(synapse.providers[SupportedChainId.ETH]).toBe(ethProvider)
      expect(synapse.providers[SupportedChainId.ARBITRUM]).toBe(arbProvider)
      expect(synapse.providers[SupportedChainId.MOONBEAM]).toBe(
        moonbeamProvider
      )
    })
  })
})
