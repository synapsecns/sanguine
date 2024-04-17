//@ts-ignore
import { SynapseSDK } from '@synapsecns/sdk-router'
import { createContext, useContext, memo, useMemo } from 'react'
import {
  StaticJsonRpcProvider,
  FallbackProvider,
  FallbackProviderConfig,
} from '@ethersproject/providers'

export const SynapseContext = createContext(null)

export const SynapseProvider = memo(
  ({ children, chains }: { children: React.ReactNode; chains: any[] }) => {
    const synapseProviders = chains.map((chain) => {
      const providerUrls = [chain?.configRpc, chain?.fallbackRpc]

      // Set priority based on list order
      const providerConfigs: FallbackProviderConfig[] = providerUrls.map(
        (url, index) => ({
          provider: new StaticJsonRpcProvider(url, chain.id),
          priority: index,
          stallTimeout: 750,
        })
      )

      // Use quorum of 1
      return new FallbackProvider(providerConfigs, 1)
    })

    const providerMap = chains.reduce((map, chain) => {
      map[chain.id] = synapseProviders.find(
        (provider) => provider.network.chainId === chain.id
      )
      return map
    }, {})

    const chainIds = chains.map((chain) => chain.id)
    const synapseSDK = new SynapseSDK(chainIds, synapseProviders)

    return (
      <SynapseContext.Provider value={{ synapseSDK, providerMap }}>
        {children}
      </SynapseContext.Provider>
    )
  }
)

export const useSynapseContext = () => useContext(SynapseContext)
