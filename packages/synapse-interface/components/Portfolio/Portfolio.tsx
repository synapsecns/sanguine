import React, { useEffect, useState } from 'react'
import { useAccount, useNetwork } from 'wagmi'
import { useAppDispatch } from '@/store/hooks'
import { setFromChainId } from '@/slices/bridge/reducer'
import { PortfolioTabManager } from './PortfolioTabManager'
import {
  NetworkTokenBalancesAndAllowances,
  TokenWithBalanceAndAllowance,
} from '@/utils/actions/fetchPortfolioBalances'
import { PortfolioContent } from './PortfolioContent/PortfolioContent'
import {
  useFetchPortfolioBalances,
  fetchAndStorePortfolioBalances,
  usePortfolioState,
} from '@/slices/portfolio/hooks'
import {
  PortfolioTabs,
  resetPortfolioState,
  setActiveTab,
} from '@/slices/portfolio/actions'
import { resetTransactionsState } from '@/slices/transactions/actions'
import { Activity } from './Activity'
import { PortfolioState } from '@/slices/portfolio/reducer'
import { useBridgeState } from '@/slices/bridge/hooks'
import { BridgeState } from '@/slices/bridge/reducer'

export const Portfolio = () => {
  const dispatch = useAppDispatch()
  const { fromChainId }: BridgeState = useBridgeState()
  const { activeTab }: PortfolioState = usePortfolioState()
  const { chain } = useNetwork()
  const { address } = useAccount({
    onDisconnect() {
      dispatch(resetPortfolioState())
      dispatch(resetTransactionsState())
    },
  })

  const { balancesAndAllowances: portfolioData, status: fetchState } =
    useFetchPortfolioBalances()

  const filteredPortfolioDataForBalances: NetworkTokenBalancesAndAllowances =
    filterPortfolioBalancesWithBalances(portfolioData)

  useEffect(() => {
    dispatch(resetPortfolioState())
  }, [address])

  useEffect(() => {
    ;(async () => {
      if (address && chain?.id) {
        await dispatch(setFromChainId(chain.id))
        await dispatch(fetchAndStorePortfolioBalances(address))
      }
    })()
  }, [chain, address])

  const [mounted, setMounted] = useState<boolean>(false)
  useEffect(() => setMounted(true), [])

  return (
    <div
      data-test-id="portfolio"
      className="flex flex-col w-full max-w-lg mx-auto lg:mx-0"
    >
      <PortfolioTabManager />
      <div className="mt-3">
        {mounted && (
          <>
            <PortfolioContent
              connectedAddress={address}
              connectedChainId={chain?.id}
              selectedFromChainId={fromChainId}
              networkPortfolioWithBalances={filteredPortfolioDataForBalances}
              fetchState={fetchState}
              visibility={activeTab === PortfolioTabs.PORTFOLIO}
            />
            <Activity visibility={activeTab === PortfolioTabs.ACTIVITY} />
          </>
        )}
      </div>
    </div>
  )
}

export function filterPortfolioBalancesWithBalances(
  balancesAndAllowances: NetworkTokenBalancesAndAllowances
): NetworkTokenBalancesAndAllowances {
  return Object.entries(balancesAndAllowances).reduce(
    (
      filteredBalances: NetworkTokenBalancesAndAllowances,
      [key, tokenWithBalances]
    ) => {
      const filtered = tokenWithBalances.filter(
        (token: TokenWithBalanceAndAllowance) => token.balance > 0n
      )
      if (filtered.length > 0) {
        filteredBalances[key] = filtered
      }
      return filteredBalances
    },
    {}
  )
}
