import { SynapseSDK } from '@synapsecns/sdk-router'
import { createContext, useContext, memo, useMemo } from 'react'
import {
  StaticJsonRpcProvider,
  FallbackProvider,
} from '@ethersproject/providers'
import { Chain, CustomRpcs } from 'types'

export const SynapseContext = createContext(null)

export const SynapseProvider = memo(
  ({
    children,
    chains,
    customRpcs,
  }: {
    children: React.ReactNode
    chains: Chain[]
    customRpcs?: CustomRpcs
  }) => {
    const synapseProviders = useMemo(() => {
      return chains.map((chain) => {
        return configureFallbackProvider(chain, customRpcs)
      })
    }, [chains])

    const providerMap = useMemo(() => {
      return chains.reduce((map, chain) => {
        map[chain.id] = synapseProviders.find(
          (provider) => provider.network.chainId === chain.id
        )
        return map
      }, {})
    }, [chains, synapseProviders])

    const chainIds = chains.map((chain) => chain.id)
    const synapseSDK = useMemo(
      () => new SynapseSDK(chainIds, synapseProviders),
      [chainIds, synapseProviders]
    )

    return (
      <SynapseContext.Provider
        value={{ synapseSDK, providerMap, synapseProviders }}
      >
        {children}
      </SynapseContext.Provider>
    )
  }
)

export const useSynapseContext = () => useContext(SynapseContext)

/**
 * Configure providers based on custom RPCs and chain data.
 *
 * Set available rpcs to equivalent priorities to ensure
 * highest level of resilience in FallbackProvider if one fails
 */
const configureFallbackProvider = (chain: Chain, customRpcs?: CustomRpcs) => {
  const stallTime = 750 // in ms
  let providerConfigs

  if (customRpcs && customRpcs[chain.id]) {
    // Custom RPC available
    providerConfigs = [
      { url: customRpcs[chain.id], priority: 1 },
      { url: chain.rpcUrls.primary, priority: 1 },
      { url: chain.rpcUrls.fallback, priority: 1 },
    ]
  } else {
    // Default RPCs provided
    providerConfigs = [
      { url: chain.rpcUrls.primary, priority: 1 },
      { url: chain.rpcUrls.fallback, priority: 1 },
    ]
  }

  const providers = providerConfigs.map(
    (config) => new StaticJsonRpcProvider(config.url, chain.id)
  )

  const fallbackProviderConfig = providerConfigs.map((config, index) => ({
    provider: providers[index],
    priority: config.priority,
    stallTimeout: stallTime,
  }))

  return new FallbackProvider(fallbackProviderConfig, 1)
}
