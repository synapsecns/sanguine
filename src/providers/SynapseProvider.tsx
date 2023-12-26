import { SynapseSDK } from '@synapsecns/sdk-router'
import { createContext, useContext, memo, useMemo } from 'react'
import { StaticJsonRpcProvider } from '@ethersproject/providers'
import { Chain } from 'types'

export const SynapseContext = createContext(null)

export const SynapseProvider = memo(
  ({ children, chains }: { children: React.ReactNode; chains: Chain[] }) => {
    const synapseProviders = useMemo(() => {
      return chains.map(
        (chain) => new StaticJsonRpcProvider(chain.rpcUrls.primary, chain.id)
      )
    }, [chains])

    const providerMap = useMemo(() => {
      return chains.reduce((map, chain) => {
        map[chain.id] = synapseProviders.find(
          (provider) => provider.connection.url === chain.rpcUrls.primary
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
