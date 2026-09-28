//@ts-ignore
import { SynapseSDK } from '@synapsecns/sdk-router'
import { createContext, useContext, memo, useMemo } from 'react'
import {
  FallbackProvider,
  FallbackProviderConfig,
} from '@ethersproject/providers'

import { TransportAwareJsonRpcProvider } from './TransportAwareJsonRpcProvider'

export const SynapseContext = createContext(null)

export const SynapseProvider = memo(
  ({ children, chains }: { children: React.ReactNode; chains: any[] }) => {
    const configurationKey = JSON.stringify(
      chains.map(({ id, configRpc, fallbackRpc }) => ({
        id,
        configRpc,
        fallbackRpc,
      }))
    )
    const configuration = useMemo(
      () => JSON.parse(configurationKey),
      [configurationKey]
    )
    const synapseProviders = useMemo(
      () =>
        configuration.map((chain) => {
          const providerUrls = [
            `/api/rpc/${chain.id}`,
            chain?.configRpc,
            chain?.fallbackRpc,
          ]

          // Set priority based on list order
          const providerConfigs: FallbackProviderConfig[] = providerUrls.map(
            (url, index) => ({
              provider: new TransportAwareJsonRpcProvider(url, chain.id),
              priority: index,
              stallTimeout: 750,
            })
          )

          // Use quorum of 1
          return new FallbackProvider(providerConfigs, 1)
        }),
      [configuration]
    )

    const providerMap = useMemo(
      () =>
        configuration.reduce((map, chain) => {
          map[chain.id] = synapseProviders.find(
            (provider) => provider.network.chainId === chain.id
          )
          return map
        }, {}),
      [configuration, synapseProviders]
    )

    const synapseSDK = useMemo(
      () =>
        new SynapseSDK(
          configuration.map((chain) => chain.id),
          synapseProviders
        ),
      [configuration, synapseProviders]
    )
    const context = useMemo(
      () => ({ synapseSDK, providerMap }),
      [synapseSDK, providerMap]
    )

    return (
      <SynapseContext.Provider value={context}>
        {children}
      </SynapseContext.Provider>
    )
  }
)

export const useSynapseContext = () => useContext(SynapseContext)
