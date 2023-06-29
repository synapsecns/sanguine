import { createContext, useContext, useEffect, useRef } from 'react'
import { Chain, useAccount, useNetwork } from 'wagmi'
import { segmentAnalyticsEvent } from './SegmentAnalyticsProvider'
import { useRouter } from 'next/router'

const WalletStatusContext = createContext(undefined)

export const WalletAnalyticsProvider = ({ children }) => {
  const { chain } = useNetwork()
  const router = useRouter()
  const { query, pathname } = router
  const { connector } = useAccount({
    onConnect() {
      segmentAnalyticsEvent(`[Wallet Analytics] connects`, {
        walletId: connector?.id,
        chainId: chain?.id,
        query,
        pathname,
      })
    },
    onDisconnect() {
      segmentAnalyticsEvent('[Wallet Analytics] disconnect', {})
    },
  })

  const prevChainRef = useRef<Chain | null>(null)
  useEffect(() => {
    prevChainRef.current = chain
  }, [chain])
  const prevChain = prevChainRef.current

  useEffect(() => {
    if (!chain) {
      return
    }
    if (prevChain && chain !== prevChain) {
      segmentAnalyticsEvent(`[Wallet Analytics] connected to new chain`, {
        previousNetworkName: prevChain.name,
        previousChainId: prevChain.id,
        walletId: connector?.id,
        newChainName: chain.name,
        newChainId: chain.id,
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
