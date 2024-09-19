import { createContext, useContext, useEffect, useRef, useState } from 'react'
import { useAccount, useAccountEffect } from 'wagmi'
import { Chain } from 'viem'
import { segmentAnalyticsEvent } from './SegmentAnalyticsProvider'
import { useRouter } from 'next/router'
import { setSwapChainId } from '@/slices/swap/reducer'

import { fetchAndStorePortfolioBalances } from '@/slices/portfolio/hooks'
import { useAppDispatch } from '@/store/hooks'
import { resetPortfolioState } from '@/slices/portfolio/actions'
import { screenAddress } from '@/utils/screenAddress'

const WalletStatusContext = createContext(undefined)

export const UserProvider = ({ children }) => {
  const dispatch = useAppDispatch()
  const { chain, address, connector } = useAccount()
  const [isClient, setIsClient] = useState(false)
  const router = useRouter()
  const { query, pathname } = router

  useAccountEffect({
    onConnect() {
      segmentAnalyticsEvent(
        `[Wallet Analytics] connects`,
        {
          walletId: connector?.id,
          chainId: chain?.id,
          query,
          pathname,
        },
        true
      )
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
    setIsClient(true)
  }, [])

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
      if (isClient && address) {
        try {
          await dispatch(fetchAndStorePortfolioBalances(address))
        } catch (error) {
          console.error('Failed to fetch and store portfolio balances:', error)
        }
      }

      if (!address) {
        dispatch(resetPortfolioState())
      }
    })()
  }, [chain, address, isClient])

  useEffect(() => {
    if (address) {
      ;(async () => {
        const isRisky = await screenAddress(address)
        if (isRisky) {
          return
        }
      })()
    }
  }, [address])

  return (
    <WalletStatusContext.Provider value={null}>
      {children}
    </WalletStatusContext.Provider>
  )
}

export const useUserStatus = () => useContext(WalletStatusContext)
