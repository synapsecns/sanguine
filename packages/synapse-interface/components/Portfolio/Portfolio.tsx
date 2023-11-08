import React, { useEffect, useState, useMemo } from 'react'
import Fuse from 'fuse.js'
import { Address, useAccount, useNetwork } from 'wagmi'
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
  FetchState,
} from '@/slices/portfolio/actions'
import { resetTransactionsState } from '@/slices/transactions/actions'
import { PortfolioState } from '@/slices/portfolio/reducer'
import { useBridgeState } from '@/slices/bridge/hooks'
import { BridgeState } from '@/slices/bridge/reducer'
import { resetBridgeInputs } from '@/slices/bridge/actions'
import { resetReduxCache } from '@/slices/application/actions'
import { isValidAddress } from '@/utils/isValidAddress'
import { ViewSearchAddressBanner } from './SearchBar'
import { Activity } from './Activity'

export const Portfolio = () => {
  const dispatch = useAppDispatch()
  const { fromChainId }: BridgeState = useBridgeState()
  const {
    activeTab,
    searchInput,
    searchStatus,
    searchedBalancesAndAllowances,
  }: PortfolioState = usePortfolioState()
  const { chain } = useNetwork()
  const { address } = useAccount({
    onDisconnect() {
      dispatch(resetPortfolioState())
      dispatch(resetTransactionsState())
      dispatch(resetBridgeInputs())
      dispatch(resetReduxCache())
    },
  })

  const { balancesAndAllowances: portfolioData, status: fetchState } =
    useFetchPortfolioBalances()

  const filteredPortfolioDataForBalances: NetworkTokenBalancesAndAllowances =
    filterPortfolioBalancesWithBalances(portfolioData)

  const searchInputActive: boolean = useMemo(() => {
    return searchInput.length > 0
  }, [searchInput])

  const searchInputIsAddress: boolean = useMemo(() => {
    return isValidAddress(searchInput)
  }, [searchInput])

  const masqueradeActive: boolean = useMemo(() => {
    return Object.keys(searchedBalancesAndAllowances).length > 0
  }, [searchedBalancesAndAllowances])

  const filteredSearchedPortfolioDataForBalances = useMemo(() => {
    if (masqueradeActive) {
      const queriedAddress: Address = Object.keys(
        searchedBalancesAndAllowances
      )[0] as Address
      return {
        balances: filterPortfolioBalancesWithBalances(
          searchedBalancesAndAllowances[queriedAddress]
        ),
        address: queriedAddress,
      }
    }
    return {
      balances: {},
      address: '',
    }
  }, [searchedBalancesAndAllowances, masqueradeActive, searchInput])

  const flattenedPortfolioData: TokenWithBalanceAndAllowances[] =
    useMemo(() => {
      const flattened: TokenWithBalanceAndAllowances[] = []
      const portfolio: NetworkTokenBalancesAndAllowances = masqueradeActive
        ? filteredSearchedPortfolioDataForBalances.balances
        : filteredPortfolioDataForBalances
      Object.entries(portfolio).forEach(([chainId, tokens]) => {
        tokens.forEach((token: TokenWithBalanceAndAllowances) => {
          flattened.push({ ...token })
        })
      })
      return flattened
    }, [
      masqueradeActive,
      filteredPortfolioDataForBalances,
      filteredSearchedPortfolioDataForBalances,
    ])

  const filteredBySearchInput: NetworkTokenBalancesAndAllowances =
    useMemo(() => {
      const searchFiltered: NetworkTokenBalancesAndAllowances = {}
      const fuseOptions = {
        includeScore: true,
        threshold: 0.33,
        distance: 20,
        keys: ['queriedChain.name', 'token.name', 'token.symbol'],
      }
      const fuse = new Fuse(flattenedPortfolioData, fuseOptions)

      if (searchInput.length > 0) {
        const results = fuse
          .search(searchInput)
          .map((i: Fuse.FuseResult<TokenWithBalanceAndAllowances>) => i.item)
          .forEach((item: TokenWithBalanceAndAllowances) => {
            const chainId: number = item.queriedChain.id
            searchFiltered[chainId] = searchFiltered[chainId]
              ? [...searchFiltered[chainId], item]
              : [item]
          })
      }
      return searchFiltered
    }, [searchInput, flattenedPortfolioData])

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
      <div className="mt-6">
        {mounted && (
          <>
            {searchStatus === FetchState.LOADING && (
              <div className="pb-3 text-secondary">Loading new address...</div>
            )}
            {masqueradeActive ? (
              <>
                <ViewSearchAddressBanner
                  viewingAddress={
                    filteredSearchedPortfolioDataForBalances.address as Address
                  }
                />
                <PortfolioContent
                  connectedAddress={
                    filteredSearchedPortfolioDataForBalances.address as Address
                  }
                  connectedChainId={chain?.id}
                  selectedFromChainId={fromChainId}
                  networkPortfolioWithBalances={
                    searchInputActive && !searchInputIsAddress
                      ? filteredBySearchInput
                      : filteredSearchedPortfolioDataForBalances.balances
                  }
                  fetchState={searchStatus}
                  visibility={activeTab === PortfolioTabs.PORTFOLIO}
                  searchInputActive={searchInputActive}
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
                    searchInputActive
                      ? filteredBySearchInput
                      : filteredPortfolioDataForBalances
                  }
                  fetchState={fetchState}
                  visibility={activeTab === PortfolioTabs.PORTFOLIO}
                  searchInputActive={searchInputActive}
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
