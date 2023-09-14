import { createContext, useContext, useEffect, useRef } from 'react'
import { Chain, useAccount, useNetwork } from 'wagmi'
import { segmentAnalyticsEvent } from './SegmentAnalyticsProvider'
import { useRouter } from 'next/router'
import { setSwapChainId } from '@/slices/swap/reducer'

import { useDispatch } from 'react-redux'

const WalletStatusContext = createContext(undefined)

// Refactor as User Provider

export const WalletAnalyticsProvider = ({ children }) => {
  const dispatch = useDispatch()
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
    if (chain) {
      dispatch(setSwapChainId(chain.id))
    }

    if (!chain) {
      return
    }
    if (prevChain && chain !== prevChain) {
      dispatch(setSwapChainId(chain.id))

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
