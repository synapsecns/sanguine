import { createContext, useContext, useEffect, useRef } from 'react'
import { useAccount, useNetwork } from 'wagmi'
import { Chain } from 'viem'
import { useRouter } from 'next/router'

import { setSwapChainId } from '@/slices/swap/reducer'
import { fetchAndStorePortfolioBalances } from '@/slices/portfolio/hooks'
import { resetPortfolioState } from '@/slices/portfolio/actions'
import { useAppDispatch } from '@/store/hooks'

import {
  fetchAllEthStablecoinPrices,
  fetchArbPrice,
  fetchAvaxPrice,
  fetchCoingeckoPrices,
  fetchDaiePrice,
  fetchEthPrice,
  fetchGmxPrice,
  fetchMetisPrice,
  fetchMusdcPrice,
  fetchSynPrices,
} from '@/slices/priceDataSlice'
import { isBlacklisted } from '@/utils/isBlacklisted'
import { screenAddress } from '@/utils/screenAddress'
import {
  fetchArbStipRewards,
  fetchFeeAndRebate,
} from '@/slices/feeAndRebateSlice'

import { segmentAnalyticsEvent } from '@/contexts/segmentAnalyticsEvent'


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
    dispatch(fetchSynPrices())
    dispatch(fetchEthPrice())
    dispatch(fetchAvaxPrice())
    dispatch(fetchMetisPrice())
    dispatch(fetchArbPrice())
    dispatch(fetchGmxPrice())
    dispatch(fetchAllEthStablecoinPrices())
    dispatch(fetchCoingeckoPrices())
    dispatch(fetchMusdcPrice())
    dispatch(fetchDaiePrice())
    dispatch(fetchFeeAndRebate())
  }, [])

  useEffect(() => {
    if (chain) {
      dispatch(setSwapChainId(chain.id))
    } else {
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
    (async () => {
      if (address && chain?.id) {
        try {
          await dispatch(fetchAndStorePortfolioBalances(address))
        } catch (error) {
          console.error('Failed to fetch and store portfolio balances:', error)
        }
      }

      if (address) {
        dispatch(fetchArbStipRewards(address))
      }

      if (!address) {
        dispatch(resetPortfolioState())
      }
    })()
  }, [chain, address])

  useEffect(() => {
    if (address) {
      if (!isBlacklisted(address)) {
        screenAddress(address)
      } else {
        document.body = document.createElement('body')
      }
    }
  }, [address])

  return (
    <WalletStatusContext.Provider value={null}>
      {children}
    </WalletStatusContext.Provider>
  )
}

/* THIS DOES NOT APPEAR TO DO ANYTHING?!?!?!?!? */
export const useUserStatus = () => useContext(WalletStatusContext)
