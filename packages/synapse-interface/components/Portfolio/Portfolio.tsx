import React, { useState, useEffect, useMemo } from 'react'
import { useAccount, useNetwork, Address } from 'wagmi'
import { Zero } from '@ethersproject/constants'
import { ConnectWalletButton } from './ConnectWalletButton'
import { PortfolioTabManager } from './PortfolioTabManager'
import {
  usePortfolioBalancesAndAllowances,
  NetworkTokenBalancesAndAllowances,
  TokenWithBalanceAndAllowance,
  FetchState,
} from '@/utils/hooks/usePortfolioBalances'
import { PortfolioContent, HomeContent } from './PortfolioContent'
import { RootState } from '@/store/store'
import { useSelector } from 'react-redux'

export enum PortfolioTabs {
  HOME = 'home',
  PORTFOLIO = 'portfolio',
}

export const Portfolio = () => {
  const [tab, setTab] = useState<PortfolioTabs>(PortfolioTabs.HOME)

  const { fromChainId } = useSelector((state: RootState) => state.bridge)

  const {
    balancesAndAllowances: portfolioData,
    fetchPortfolioBalances: refetch,
    status: fetchState,
  } = usePortfolioBalancesAndAllowances()

  const filteredPortfolioDataForBalances: NetworkTokenBalancesAndAllowances =
    filterPortfolioBalancesWithBalances(portfolioData)

  const { address } = useAccount()
  const { chain } = useNetwork()

  useEffect(() => {
    if (address) {
      setTab(PortfolioTabs.PORTFOLIO)
    }
  }, [address])

  return (
    <div
      data-test-id="portfolio"
      className="flex flex-col w-full px-4 py-10 lg:py-2 lg:w-2/5"
    >
      <PortfolioTabManager activeTab={tab} setTab={setTab} />
      <div className="border-t border-solid border-[#28282F] mt-4">
        {tab === PortfolioTabs.HOME && <HomeContent />}
        {tab === PortfolioTabs.PORTFOLIO && (
          <PortfolioContent
            connectedAddress={address}
            connectedChainId={chain?.id}
            selectedFromChainId={fromChainId}
            networkPortfolioWithBalances={filteredPortfolioDataForBalances}
            fetchState={fetchState}
          />
        )}
      </div>
    </div>
  )
}

function filterPortfolioBalancesWithBalances(
  balancesAndAllowances: NetworkTokenBalancesAndAllowances
): NetworkTokenBalancesAndAllowances {
  const filteredBalances: NetworkTokenBalancesAndAllowances = {}

  Object.entries(balancesAndAllowances).forEach(([key, tokenWithBalances]) => {
    const filteredTokenWithBalances = tokenWithBalances.filter(
      (token: TokenWithBalanceAndAllowance) => token.balance > Zero
    )

    if (filteredTokenWithBalances.length > 0) {
      filteredBalances[key] = filteredTokenWithBalances
    }
  })

  return filteredBalances
}
