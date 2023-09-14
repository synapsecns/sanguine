import React, { useEffect, useState, useMemo } from 'react'
import Fuse from 'fuse.js'
import { useAccount, useNetwork } from 'wagmi'
import { useAppDispatch } from '@/store/hooks'
import { setFromChainId } from '@/slices/bridge/reducer'
import { PortfolioTabManager } from './PortfolioTabManager'
import {
  NetworkTokenBalancesAndAllowances,
  TokenWithBalanceAndAllowance,
  TokenWithBalanceAndAllowances,
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
import { CHAINS_BY_ID } from '@/constants/chains'

export const Portfolio = () => {
  const dispatch = useAppDispatch()
  const { fromChainId }: BridgeState = useBridgeState()
  const { activeTab, searchInput }: PortfolioState = usePortfolioState()
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

  const filteredBySearchInput = useMemo(() => {
    if (filteredPortfolioDataForBalances) {
      const flattened: TokenWithBalanceAndAllowances[] = []
      const reconstructedFilteredData: NetworkTokenBalancesAndAllowances = {}

      Object.entries(filteredPortfolioDataForBalances).forEach(
        ([chainId, tokens]) => {
          tokens.forEach((token: TokenWithBalanceAndAllowances) => {
            flattened.push({ ...token })
          })
        }
      )

      const fuseOptions = {
        includeScore: true,
        threshold: 0.0,
        keys: ['queriedChain.name', 'token.name', 'token.symbol'],
      }

      const fuse = new Fuse(flattened, fuseOptions)

      if (searchInput.length > 0) {
        const results = fuse
          .search(searchInput)
          .map((i: Fuse.FuseResult<TokenWithBalanceAndAllowances>) => i.item)

        results.forEach((item: TokenWithBalanceAndAllowances) => {
          const chainId = item.queriedChain.id

          if (reconstructedFilteredData[chainId]) {
            reconstructedFilteredData[chainId] = [
              ...reconstructedFilteredData[chainId],
              item,
            ]
          } else {
            reconstructedFilteredData[chainId] = [item]
          }
        })

        console.log('reconstructedFilteredData:', reconstructedFilteredData)
      }
    }
  }, [searchInput, filteredPortfolioDataForBalances])

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
