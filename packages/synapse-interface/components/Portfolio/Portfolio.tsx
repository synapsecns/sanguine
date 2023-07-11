import React, { useState, useEffect, useMemo } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { useAccount, useNetwork, Address } from 'wagmi'
import { Zero } from '@ethersproject/constants'
import { RootState } from '@/store/store'
import { setFromChainId } from '@/slices/bridgeSlice'
import { ConnectWalletButton } from './ConnectWalletButton'
import { PortfolioTabManager } from './PortfolioTabManager'
import {
  usePortfolioBalancesAndAllowances,
  NetworkTokenBalancesAndAllowances,
  TokenWithBalanceAndAllowance,
  FetchState,
} from '@/utils/hooks/usePortfolioBalances'
import { PortfolioContent, HomeContent } from './PortfolioContent'

export enum PortfolioTabs {
  HOME = 'home',
  PORTFOLIO = 'portfolio',
}

export const Portfolio = () => {
  const [tab, setTab] = useState<PortfolioTabs>(PortfolioTabs.HOME)

  const dispatch = useDispatch()
  const { fromChainId, bridgeTxHashes } = useSelector(
    (state: RootState) => state.bridge
  )

  const {
    balancesAndAllowances: portfolioData,
    fetchPortfolioBalances,
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

  useEffect(() => {
    if (address) fetchPortfolioBalances()
  }, [address, bridgeTxHashes])

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
