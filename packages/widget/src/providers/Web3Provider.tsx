import { createContext, useEffect, useState } from 'react'

export interface Web3ProviderTypes {
  connectedAddress: string
  networkId: number | null
  signer: any | null
  provider: any | null
}

export interface ContextType {
  web3Provider: Web3ProviderTypes | null
  setWeb3Provider: React.Dispatch<React.SetStateAction<Web3ProviderTypes>>
}

export const Web3Context = createContext<ContextType | null>(null)

export const Web3Provider = ({
  children,
  config,
}: {
  children: React.ReactNode
  config: any
}) => {
  const [web3Provider, setWeb3Provider] = useState<Web3ProviderTypes>({
    connectedAddress: '',
    networkId: null,
    signer: null,
    provider: null,
  })

  useEffect(() => {
    async function fetchData() {
      try {
        const signer = await config.getSigner()
        const address = await signer.getAddress()
        const network = await config.getNetwork()

        setWeb3Provider({
          connectedAddress: address,
          networkId: Number(network?.chainId),
          signer,
          provider: config,
        })
      } catch (e) {
        console.log('Error', e)
      }
    }

    if (config) {
      fetchData()
    }
  }, [config])

  return (
    <Web3Context.Provider value={{ web3Provider, setWeb3Provider }}>
      {children}
    </Web3Context.Provider>
  )
}
