import { useState, useEffect } from 'react'
import { Network, ethers } from 'ethers'

declare global {
  interface Window {
    ethereum: any
  }
}

export const useEthereumWallet = () => {
  const [connectedAddress, setConnectedAddress] = useState<string>('')
  const [web3Provider, setWeb3Provider] = useState<any>(null)
  const [connectedNetwork, setConnectedNetwork] = useState<Network | null>()

  /**
   * Subscribing to network and account changes
   * Every time the network or account is changed, the above states will be udpated
   */
  useEffect(() => {
    if (window.ethereum) {
      connectWallet()

      window.ethereum.on('chainChanged', () => {
        console.log('Consumer chain updated')
        connectWallet()
      })

      window.ethereum.on('accountsChanged', () => {
        console.log('Consumer account changed')
        connectWallet()
      })
    }
  }, [])

  /** Connects to the injected wallet */
  const connectWallet = async () => {
    try {
      if (window.ethereum) {
        // Getting the browser provider
        const _provider = new ethers.BrowserProvider(window.ethereum)
        /**
         * if using ethers v5, you can call the browser provider as follows
         * const _provider = new ethers.providers.Web3Provider(window.ethereum, "any");
         */

        const signer = await _provider.getSigner()
        const _address = await signer.getAddress()
        const _network = await _provider.getNetwork()

        setConnectedNetwork(_network)
        setWeb3Provider(_provider)
        setConnectedAddress(_address)
      } else {
        alert('Injected Ethereum wallet not found')
      }
    } catch (e) {
      console.log('Error while connecting to the wallet', e)
    }
  }

  return { web3Provider, connectedAddress, connectedNetwork }
}
