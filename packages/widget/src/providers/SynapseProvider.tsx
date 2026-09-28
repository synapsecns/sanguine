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
    const configurationKey = JSON.stringify(
      chains.map((chain) => ({
        id: chain.id,
        urls: customRpcs?.[chain.id]
          ? [
              customRpcs[chain.id],
              chain.rpcUrls.primary,
              chain.rpcUrls.fallback,
            ]
          : [chain.rpcUrls.primary, chain.rpcUrls.fallback],
      }))
    )
    const configuration = useMemo(
      () => JSON.parse(configurationKey),
      [configurationKey]
    )
    const synapseProviders = useMemo(() => {
      return configuration.map((chain) => {
        const providerConfigs: FallbackProviderConfig[] = chain.urls.map(
          (url, index) => ({
            provider: new StaticJsonRpcProvider(url, chain.id),
            priority: index,
            stallTimeout: 750,
          })
        )

        return new FallbackProvider(providerConfigs, 1)
      })
    }, [configuration])

    const providerMap = useMemo(() => {
      return configuration.reduce((map, chain) => {
        map[chain.id] = synapseProviders.find(
          (provider) => provider.network.chainId === chain.id
        )
        return map
      }, {})
    }, [configuration, synapseProviders])

    const synapseSDK = useMemo(
      () =>
        new SynapseSDK(
          configuration.map((chain) => chain.id),
          synapseProviders
        ),
      [configuration, synapseProviders]
    )
    const context = useMemo(
      () => ({ synapseSDK, providerMap, synapseProviders }),
      [synapseSDK, providerMap, synapseProviders]
    )
    return (
      <SynapseContext.Provider value={context}>
        {children}
      </SynapseContext.Provider>
    )
  }
)

export const useSynapseContext = () => useContext(SynapseContext)
