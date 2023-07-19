import React, { useState, useEffect } from 'react'
import { useAccount, useNetwork } from 'wagmi'
import { setFromChainId } from '@/slices/bridge/reducer'
import { PortfolioTabManager } from './PortfolioTabManager'
import {
  NetworkTokenBalancesAndAllowances,
  TokenWithBalanceAndAllowance,
} from '@/utils/hooks/usePortfolioBalances'
import { PortfolioContent, HomeContent } from './PortfolioContent'
import { useFetchPortfolioBalances } from '@/slices/portfolio/hooks'
import { fetchAndStorePortfolioBalances } from '@/slices/portfolio/reducer'
import { useAppDispatch } from '@/store/hooks'
import { useBridgeState } from '@/slices/bridge/hooks'

export enum PortfolioTabs {
  HOME = 'home',
  PORTFOLIO = 'portfolio',
}

export const Portfolio = () => {
  const [tab, setTab] = useState<PortfolioTabs>(PortfolioTabs.HOME)
  const { fromChainId, bridgeTxHashes } = useBridgeState()
  const { address } = useAccount()
  const { chain } = useNetwork()
  const dispatch = useAppDispatch()

  useEffect(() => {
    if (address) {
      dispatch(fetchAndStorePortfolioBalances(address))
    }
  }, [address])

  const {
    balancesAndAllowances: portfolioData,
    fetchPortfolioBalances,
    status: fetchState,
  } = useFetchPortfolioBalances()

  const filteredPortfolioDataForBalances: NetworkTokenBalancesAndAllowances =
    filterPortfolioBalancesWithBalances(portfolioData)

  useEffect(() => {
    if (address) {
      setTab(PortfolioTabs.PORTFOLIO)
    }
  }, [address])

  useEffect(() => {
    ;(async () => {
      if (address && chain.id) {
        await dispatch(setFromChainId(chain.id))
        await dispatch(fetchAndStorePortfolioBalances(address))
      }
    })()
  }, [address, chain])

  return (
    <div
      data-test-id="portfolio"
      className="flex flex-col w-full max-w-lg mx-auto lg:mx-0"
    >
      <PortfolioTabManager activeTab={tab} setTab={setTab} />
      <div className="mt-4">
        {tab === PortfolioTabs.HOME && <HomeContent />}
        {tab === PortfolioTabs.PORTFOLIO && (
          <PortfolioContent
            connectedAddress={address}
            connectedChainId={chain?.id}
            selectedFromChainId={fromChainId}
            networkPortfolioWithBalances={filteredPortfolioDataForBalances}
            fetchPortfolioBalancesCallback={fetchPortfolioBalances}
            fetchState={fetchState}
            bridgeTxHashes={bridgeTxHashes}
          />
        )}
      </div>
    </div>
  )
}

function filterPortfolioBalancesWithBalances(
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
