import React, { useEffect, useState, useMemo } from 'react'
import Fuse from 'fuse.js'
import { Address, useAccount, useNetwork } from 'wagmi'
import { useAppDispatch } from '@/store/hooks'
import { setFromChainId } from '@/slices/bridge/reducer'
import { PortfolioTabManager } from './components/PortfolioTabManager'
import {
  NetworkTokenBalances,
  TokenAndBalance,
} from '@/utils/actions/fetchPortfolioBalances'
import { PortfolioContent } from './components/PortfolioContent'
import {
  useFetchPortfolioBalances,
  usePortfolioState,
} from '@/slices/portfolio/hooks'
import {
  PortfolioTabs,
  resetPortfolioState,
  FetchState,
} from '@/slices/portfolio/actions'
import { resetTransactionsState } from '@/slices/transactions/actions'
import { PortfolioState } from '@/slices/portfolio/reducer'
import { useBridgeState } from '@/slices/bridge/hooks'
import { BridgeState } from '@/slices/bridge/reducer'
import { resetBridgeInputs } from '@/slices/bridge/actions'
import { ViewSearchAddressBanner } from './SearchBar'
import { Activity } from './Activity'
import { useSearchInputState } from './helpers/useSearchInputStatus'

export const Portfolio = () => {
  const dispatch = useAppDispatch()
  const { fromChainId }: BridgeState = useBridgeState()
  const {
    activeTab,
    searchInput,
    searchStatus,
    searchedBalances,
  }: PortfolioState = usePortfolioState()
  const { chain } = useNetwork()
  const { address } = useAccount({
    onDisconnect() {
      dispatch(resetPortfolioState())
      dispatch(resetTransactionsState())
      dispatch(resetBridgeInputs())
    },
  })

  const { balances: portfolioData, status: fetchState } =
    useFetchPortfolioBalances()

  const filteredPortfolioDataForBalances: NetworkTokenBalances =
    filterPortfolioBalancesWithBalances(portfolioData)

  const {
    isSearchInputActive,
    isSearchInputAddress,
    isMasqueradeActive,
    masqueradeAddress,
  } = useSearchInputState()

  const filteredSearchedPortfolioDataForBalances = useMemo(() => {
    if (isMasqueradeActive) {
      return {
        balances: filterPortfolioBalancesWithBalances(
          searchedBalances[masqueradeAddress]
        ),
        address: masqueradeAddress,
      }
    }
    return {
      balances: {},
      address: '',
    }
  }, [searchedBalances, isMasqueradeActive, searchInput])

  const flattenedPortfolioData = flattenData(portfolioData)

  const filteredBySearchInput = filterBySearchInput(
    flattenedPortfolioData,
    searchInput
  )

  useEffect(() => {
    if (address && chain?.id) {
      dispatch(setFromChainId(chain.id))
    }
  }, [chain])

  const [mounted, setMounted] = useState<boolean>(false)
  useEffect(() => setMounted(true), [])

  return (
    <div
      data-test-id="portfolio"
      className="flex flex-col w-full max-w-lg mx-auto lg:mx-0"
    >
      <PortfolioTabManager />
      <div className="mt-6">
        {mounted && (
          <>
            {searchStatus === FetchState.LOADING && (
              <div className="pb-3 text-secondary">Loading new address...</div>
            )}
            {isMasqueradeActive ? (
              <>
                <ViewSearchAddressBanner
                  viewingAddress={
                    filteredSearchedPortfolioDataForBalances.address as Address
                  }
                />
                <PortfolioContent
                  connectedAddress={
                    filteredSearchedPortfolioDataForBalances.address
                  }
                  connectedChainId={chain?.id}
                  selectedFromChainId={fromChainId}
                  networkPortfolioWithBalances={
                    isSearchInputActive && !isSearchInputAddress
                      ? filteredBySearchInput
                      : filteredSearchedPortfolioDataForBalances.balances
                  }
                  fetchState={searchStatus}
                  visibility={activeTab === PortfolioTabs.PORTFOLIO}
                  searchInputActive={isSearchInputActive}
                  searchStatus={searchStatus}
                  searchInput={searchInput}
                />
              </>
            ) : (
              <div
                className={
                  searchStatus === FetchState.LOADING
                    ? 'opacity-30 cursor-not-allowed'
                    : 'opacity-100'
                }
              >
                <PortfolioContent
                  connectedAddress={address}
                  connectedChainId={chain?.id}
                  selectedFromChainId={fromChainId}
                  networkPortfolioWithBalances={
                    isSearchInputActive
                      ? filteredBySearchInput
                      : filteredPortfolioDataForBalances
                  }
                  fetchState={fetchState}
                  visibility={activeTab === PortfolioTabs.PORTFOLIO}
                  searchInputActive={isSearchInputActive}
                  searchStatus={searchStatus}
                  searchInput={searchInput}
                />
              </div>
            )}
            <Activity visibility={activeTab === PortfolioTabs.ACTIVITY} />
          </>
        )}
      </div>
    </div>
  )
}

function flattenData(portfolioData: NetworkTokenBalances): TokenAndBalance[] {
  const flattened: TokenAndBalance[] = []
  Object.entries(portfolioData).forEach(([chainId, tokens]) => {
    tokens.forEach((token: TokenAndBalance) => {
      flattened.push({ ...token })
    })
  })
  return flattened
}

function filterBySearchInput(
  portfolioData: TokenAndBalance[],
  searchInput: string
) {
  const searchFiltered: NetworkTokenBalances = {}
  const fuseOptions = {
    includeScore: true,
    threshold: 0.33,
    distance: 20,
    keys: ['queriedChain.name', 'token.name', 'token.symbol'],
  }
  const fuse = new Fuse(portfolioData, fuseOptions)

  if (searchInput.length > 0) {
    fuse
      .search(searchInput)
      .map((i: Fuse.FuseResult<TokenAndBalance>) => i.item)
      .forEach((item) => {
        const chainId = item.queriedChain.id
        if (!searchFiltered[chainId]) {
          searchFiltered[chainId] = []
        }
        searchFiltered[chainId].push(item)
      })
  }

  return searchFiltered
}

export function filterPortfolioBalancesWithBalances(
  balances: NetworkTokenBalances
): NetworkTokenBalances {
  return Object.entries(balances).reduce(
    (filteredBalances: NetworkTokenBalances, [key, tokenWithBalances]) => {
      const filtered = tokenWithBalances.filter(
        (token: TokenAndBalance) => token.balance > 0n
      )
      if (filtered.length > 0) {
        filteredBalances[key] = filtered
      }
      return filteredBalances
    },
    {}
  )
}
