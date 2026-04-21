import { useEffect, useState } from 'react'
import type { BridgeProps } from '@synapsecns/widget'
import { ethers } from 'ethers'

type EthereumProvider = {
  request?: (args: { method: string; params?: unknown[] }) => Promise<unknown>
  on?: (eventName: string, listener: () => void) => void
  removeListener?: (eventName: string, listener: () => void) => void
}

type WalletNetwork = {
  chainId?: bigint | number
  name?: string
}

type WalletSigner = {
  getAddress: () => Promise<string>
}

type WalletProvider = {
  getNetwork: () => Promise<WalletNetwork>
  getSigner: () => Promise<WalletSigner> | WalletSigner
}

type WidgetWeb3Provider = BridgeProps['web3Provider']

declare global {
  interface Window {
    ethereum?: EthereumProvider
  }
}

export const useEthereumWallet = () => {
  const [connectedAddress, setConnectedAddress] = useState('')
  const [connectedNetwork, setConnectedNetwork] =
    useState<WalletNetwork | null>(null)
  const [error, setError] = useState<string | null>(null)
  const [hasInjectedWallet, setHasInjectedWallet] = useState(false)
  const [isInitializing, setIsInitializing] = useState(true)
  const [web3Provider, setWeb3Provider] = useState<WidgetWeb3Provider | null>(
    null
  )

  useEffect(() => {
    if (globalThis.window === undefined) {
      return undefined
    }

    const injectedProvider = globalThis.window.ethereum

    if (!injectedProvider) {
      setHasInjectedWallet(false)
      setIsInitializing(false)
      setWeb3Provider(null)
      setConnectedAddress('')
      setConnectedNetwork(null)
      setError(null)
      return undefined
    }

    setHasInjectedWallet(true)

    const createProvider = (): WalletProvider => {
      if (typeof ethers.BrowserProvider === 'function') {
        return new ethers.BrowserProvider(injectedProvider)
      }

      if (typeof ethers.providers?.Web3Provider === 'function') {
        return new ethers.providers.Web3Provider(injectedProvider, 'any')
      }

      throw new Error('No compatible ethers provider constructor was found.')
    }

    const syncWalletState = async () => {
      setIsInitializing(true)
      setError(null)

      try {
        const provider = createProvider()
        const network = await provider.getNetwork()
        const accounts = injectedProvider.request
          ? await injectedProvider.request({ method: 'eth_accounts' })
          : []

        setWeb3Provider(provider as unknown as WidgetWeb3Provider)
        setConnectedNetwork(network)

        if (Array.isArray(accounts) && typeof accounts[0] === 'string') {
          const signer = await Promise.resolve(provider.getSigner())
          const address = await signer.getAddress()
          setConnectedAddress(address)
          setError(null)
        } else {
          setConnectedAddress('')
          setError(
            'Wallet not connected. The widget is running in read-only mode until you connect an account.'
          )
        }
      } catch (err) {
        console.log('Error while connecting to the wallet', err)
        setWeb3Provider(null)
        setConnectedAddress('')
        setConnectedNetwork(null)
        setError(
          'Unable to initialize a browser wallet provider for the widget.'
        )
      } finally {
        setIsInitializing(false)
      }
    }

    const handleAccountsChanged = () => {
      void syncWalletState()
    }

    const handleChainChanged = () => {
      void syncWalletState()
    }

    void syncWalletState()

    injectedProvider.on?.('accountsChanged', handleAccountsChanged)
    injectedProvider.on?.('chainChanged', handleChainChanged)

    return () => {
      injectedProvider.removeListener?.(
        'accountsChanged',
        handleAccountsChanged
      )
      injectedProvider.removeListener?.('chainChanged', handleChainChanged)
    }
  }, [])

  return {
    connectedAddress,
    connectedNetwork,
    error,
    hasInjectedWallet,
    isInitializing,
    web3Provider,
  }
}
