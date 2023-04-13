import { createContext, useContext } from 'react'
import { SynapseSDK } from '@synapsecns/sdk-router'
import { Provider } from '@ethersproject/abstract-provider'

export const SynapseContext = createContext<SynapseSDK>(null)

export const SynapseProvider = ({
  children,
  chainIds,
  providers,
}: {
  children: React.ReactNode
  chainIds: number[]
  providers: Provider[]
}) => {
  const sdk = new SynapseSDK(chainIds, providers)
  return (
    <SynapseContext.Provider value={sdk}>{children}</SynapseContext.Provider>
  )
}

export const useSynapseContext = () => useContext(SynapseContext)
