import { createContext, useContext, useEffect, useRef } from 'react'
import { Chain, useAccount, useNetwork } from 'wagmi'
import { segmentAnalyticsEvent } from './SegmentAnalyticsProvider'
import { useRouter } from 'next/router'
import { shortenAddress } from '@/utils/shortenAddress'

const WalletStatusContext = createContext(undefined)

export const WalletAnalyticsProvider = ({ children }) => {
  const { address, connector } = useAccount()
  const { chain } = useNetwork()
  const router = useRouter()
  const { query, pathname } = router

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
      segmentAnalyticsEvent(
        `[Wallet Analytics] ${shortenAddress(address)} connected`,
        {
          walletId,
          networkName,
          networkId,
          query,
          pathname,
        }
      )
    } else {
      segmentAnalyticsEvent(`[Wallet Analytics] User not connected`, {
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
      segmentAnalyticsEvent(
        `[Wallet Analytics] ${shortenAddress(address)} connected to chain`,
        {
          walletId,
          networkName,
          networkId,
          query,
          pathname,
        }
      )
    }
    if (prevChain && chain !== prevChain) {
      segmentAnalyticsEvent(
        `[Wallet Analytics] ${shortenAddress(address)} connected to new chain`,
        {
          previousNetworkName: prevChain?.name,
          previousNetworkId: prevChain?.id,
          walletId,
          networkName,
          networkId,
          query,
          pathname,
        }
      )
    }
  }, [chain])

  return (
    <WalletStatusContext.Provider value={null}>
      {children}
    </WalletStatusContext.Provider>
  )
}

export const useWalletStatus = () => useContext(WalletStatusContext)
