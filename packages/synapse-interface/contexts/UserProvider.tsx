import { createContext, useContext, useEffect, useRef } from 'react'
import { Chain, useAccount, useNetwork } from 'wagmi'
import { segmentAnalyticsEvent } from './SegmentAnalyticsProvider'
import { useRouter } from 'next/router'
import { setSwapChainId } from '@/slices/swap/reducer'

import { fetchAndStorePortfolioBalances } from '@/slices/portfolio/hooks'
import { useAppDispatch } from '@/store/hooks'

const WalletStatusContext = createContext(undefined)

export const UserProvider = ({ children }) => {
  const dispatch = useAppDispatch()
  const { chain } = useNetwork()
  const router = useRouter()
  const { query, pathname } = router
  const { address, connector } = useAccount({
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

  useEffect(() => {
    ;(async () => {
      if (address && chain?.id) {
        try {
          await dispatch(fetchAndStorePortfolioBalances(address))
        } catch (error) {
          console.error('Failed to fetch and store portfolio balances:', error)
        }
      }
    })()
  }, [chain, address])

  return (
    <WalletStatusContext.Provider value={null}>
      {children}
    </WalletStatusContext.Provider>
  )
}

export const useUserStatus = () => useContext(WalletStatusContext)
