import { SynapseSDK } from '@synapsecns/sdk-router'
import { createContext, useContext, memo, useMemo } from 'react'
import {
  StaticJsonRpcProvider,
  FallbackProvider,
} from '@ethersproject/providers'
import { Chain, CustomFallbackRpcs } from 'types'

export const SynapseContext = createContext(null)

const STALL_TIMEOUT = 1000 // 1s or 1000ms

export const SynapseProvider = memo(
  ({
    children,
    chains,
    fallbackRpcs,
  }: {
    children: React.ReactNode
    chains: Chain[]
    fallbackRpcs?: CustomFallbackRpcs
  }) => {
    const synapseProviders = useMemo(() => {
      return chains.map((chain) => {
        let providers
        let fallbackProvider

        /** Include Consumer custom rpc if provided */
        /** Consumer provided rpc will have highest priority */
        if (fallbackRpcs && fallbackRpcs[chain.id]) {
          providers = [
            new StaticJsonRpcProvider(fallbackRpcs[chain.id], chain.id),
            new StaticJsonRpcProvider(chain.rpcUrls.primary, chain.id),
            new StaticJsonRpcProvider(chain.rpcUrls.fallback, chain.id),
          ]

          fallbackProvider = new FallbackProvider(
            [
              {
                provider: providers[0],
                priority: 1,
                stallTimeout: STALL_TIMEOUT,
              },
              {
                provider: providers[1],
                priority: 2,
                stallTimeout: STALL_TIMEOUT,
              },
              {
                provider: providers[2],
                priority: 2,
                stallTimeout: STALL_TIMEOUT,
              },
            ],
            1
          )
        } else {
          providers = [
            new StaticJsonRpcProvider(chain.rpcUrls.primary, chain.id),
            new StaticJsonRpcProvider(chain.rpcUrls.fallback, chain.id),
          ]

          fallbackProvider = new FallbackProvider(
            [
              {
                provider: providers[0],
                stallTimeout: STALL_TIMEOUT,
              },
              {
                provider: providers[1],
                stallTimeout: STALL_TIMEOUT,
              },
            ],
            1
          )
        }

        return fallbackProvider
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
