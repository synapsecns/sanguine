import React, { useEffect } from 'react'
import { useAccount, useNetwork } from 'wagmi'
import { useAppDispatch } from '@/store/hooks'
import { setFromChainId } from '@/slices/bridge/reducer'
import { PortfolioTabManager } from './PortfolioTabManager'
import {
  NetworkTokenBalancesAndAllowances,
  TokenWithBalanceAndAllowance,
} from '@/utils/actions/fetchPortfolioBalances'
import { PortfolioContent, HomeContent } from './PortfolioContent'
import {
  useFetchPortfolioBalances,
  fetchAndStorePortfolioBalances,
  usePortfolioState,
} from '@/slices/portfolio/hooks'
import { PortfolioTabs, setActiveTab } from '@/slices/portfolio/actions'
import { PortfolioState } from '@/slices/portfolio/reducer'
import { useBridgeState } from '@/slices/bridge/hooks'
import { BridgeState } from '@/slices/bridge/reducer'

export const Portfolio = () => {
  const { fromChainId, bridgeTxHashes }: BridgeState = useBridgeState()
  const { activeTab }: PortfolioState = usePortfolioState()
  const { address } = useAccount({
    async onConnect() {
      await dispatch(setActiveTab(PortfolioTabs.PORTFOLIO))
      await dispatch(setFromChainId(chain.id))
      await dispatch(fetchAndStorePortfolioBalances(address))
    },
  })
  const { chain } = useNetwork()
  const dispatch = useAppDispatch()

  const {
    balancesAndAllowances: portfolioData,
    fetchPortfolioBalances,
    status: fetchState,
  } = useFetchPortfolioBalances()

  const filteredPortfolioDataForBalances: NetworkTokenBalancesAndAllowances =
    filterPortfolioBalancesWithBalances(portfolioData)

  return (
    <div
      data-test-id="portfolio"
      className="flex flex-col w-full max-w-lg mx-auto lg:mx-0"
    >
      <PortfolioTabManager />
      <div className="mt-4">
        {activeTab === PortfolioTabs.HOME && <HomeContent />}
        {activeTab === PortfolioTabs.PORTFOLIO && (
          <PortfolioContent
            connectedAddress={address}
            connectedChainId={chain?.id}
            selectedFromChainId={fromChainId}
            networkPortfolioWithBalances={filteredPortfolioDataForBalances}
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
