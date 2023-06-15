import { createContext, useContext, useEffect, useRef } from 'react'
import { Chain, useAccount, useNetwork } from 'wagmi'
import { useAnalytics } from './AnalyticsProvider'
import { useRouter } from 'next/router'

const WalletStatusContext = createContext(undefined)

export const WalletAnalyticsProvider = ({ children }) => {
  const { address, connector } = useAccount()
  const { chain } = useNetwork()
  const router = useRouter()
  const { query, pathname } = router
  const analytics = useAnalytics()

  const walletId: string = connector?.id
  const networkName: string = chain?.name
  const networkId: number = chain?.id

  const { isConnected } = useAccount()

  const prevChainRef = useRef<Chain | null>(null)
  useEffect(() => {
    prevChainRef.current = chain
  }, [chain])
  const prevChain = prevChainRef.current

  useEffect(() => {
    if (isConnected) {
      analytics.track(`[Wallet Analytics] User ${address} connected`, {
        walletId,
        networkName,
        networkId,
        query,
        pathname,
      })
    } else {
      analytics.track(`[Wallet Analytics] User not connected`, {
        query,
        pathname,
      })
    }
  }, [isConnected])

  useEffect(() => {
    if (!chain) {
      return
    }
    if (!prevChain) {
      analytics.track(`[Wallet Analytics] User ${address} connected to chain`, {
        walletId,
        networkName,
        networkId,
        query,
        pathname,
      })
    }
    if (chain !== prevChain) {
      analytics.track(`[Wallet Analytics] User ${address} switched chains`, {
        previousNetworkName: prevChain?.name,
        previousNetworkId: prevChain?.id,
        walletId,
        networkName,
        networkId,
        query,
        pathname,
      })
    }
  }, [chain])

  return (
    <WalletStatusContext.Provider value={null}>
      {children}
    </WalletStatusContext.Provider>
  )
}

export const useWalletStatus = () => useContext(WalletStatusContext)
