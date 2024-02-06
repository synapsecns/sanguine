import { SynapseSDK } from '@synapsecns/sdk-router'
import { createContext, useContext, memo, useMemo } from 'react'
import {
  StaticJsonRpcProvider,
  FallbackProvider,
  FallbackProviderConfig,
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
        let providerUrls

        if (customRpcs && customRpcs[chain.id]) {
          providerUrls = [
            customRpcs[chain.id],
            chain?.rpcUrls.primary,
            chain?.rpcUrls.fallback,
          ]
        } else {
          providerUrls = [chain?.rpcUrls.primary, chain?.rpcUrls.fallback]
        }

        const providerConfigs: FallbackProviderConfig[] = providerUrls.map(
          (url, index) => ({
            provider: new StaticJsonRpcProvider(url, chain.id),
            priority: index,
            stallTimeout: 750,
          })
        )

        return new FallbackProvider(providerConfigs, 1)
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
