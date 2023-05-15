//@ts-ignore
import { SynapseSDK } from '@synapsecns/sdk-router'
import { Provider } from '@ethersproject/abstract-provider'
import { createContext, useContext, memo, useMemo } from 'react'
import { StaticJsonRpcProvider } from '@ethersproject/providers'
import { Provider as EthersProvider } from '@ethersproject/abstract-provider'

export const SynapseContext = createContext(null)

export const SynapseProvider = memo(
  ({ children, chains }: { children: React.ReactNode; chains: any[] }) => {
    const synapseProviders = useMemo(() => {
      return chains.map(
        (chain) => new StaticJsonRpcProvider(chain.configRpc, chain.id)
      )
    }, [chains])

    const providerMap = useMemo(() => {
      return chains.reduce((map, chain) => {
        map[chain.id] = synapseProviders.find(
          (provider) => provider.connection.url === chain.configRpc
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
      <SynapseContext.Provider value={{ synapseSDK, providerMap }}>
        {children}
      </SynapseContext.Provider>
    )
  }
)

export const useSynapseContext = () => useContext(SynapseContext)
