import React, { useEffect, useMemo, useState } from 'react'
import { Address } from 'viem'
import { useAppDispatch } from '@/store/hooks'
import {
  usePortfolioActionHandlers,
  usePortfolioState,
  fetchAndStoreSearchInputPortfolioBalances,
} from '@/slices/portfolio/hooks'
import { PortfolioTabs } from '@/slices/portfolio/actions'
import {
  initialState as portfolioInitialState,
  PortfolioState,
} from '@/slices/portfolio/reducer'
import { getValidAddress } from '@/utils/isValidAddress'
import { isTransactionHash } from '@/utils/validators'
import { getTransactionHashExplorerLink } from '../Transaction/components/TransactionExplorerLink'
import { ClearSearchButton } from './ClearSearchButton'
import { useIsFocused } from '../helpers/useIsFocused'

export const inputRef = React.createRef<HTMLInputElement>()

export const SearchBar = () => {
  const dispatch = useAppDispatch()
  const { onSearchInput, clearSearchInput, clearSearchResults } =
    usePortfolioActionHandlers()
  const { activeTab, searchInput, searchedBalances }: PortfolioState =
    usePortfolioState()

  const [mounted, setMounted] = useState<boolean>(false)
  const isActive: boolean = searchInput !== portfolioInitialState.searchInput

  useEffect(() => {
    setMounted(true)
  }, [])

  const isFocused = useIsFocused(inputRef)

  const placeholder: string = useMemo(() => {
    switch (activeTab) {
      case PortfolioTabs.PORTFOLIO:
        return 'Tokens, chains...'
      case PortfolioTabs.ACTIVITY:
        return 'Bridge txs...'
      default:
        return 'Search...'
    }
  }, [activeTab])

  const searchInputIsTransactionHash: boolean = useMemo(() => {
    return isTransactionHash(searchInput)
  }, [searchInput])

  const checksumValidAddress: Address | null = useMemo(() => {
    return getValidAddress(searchInput)
  }, [searchInput])

  useEffect(() => {
    const masqueradeActive: boolean = Object.keys(searchedBalances).length > 0
    if (checksumValidAddress && !masqueradeActive) {
      dispatch(
        fetchAndStoreSearchInputPortfolioBalances(
          checksumValidAddress as Address
        )
      )
    }

    if (masqueradeActive && checksumValidAddress) {
      clearSearchInput()
    }
  }, [checksumValidAddress, searchedBalances])

  useEffect(() => {
    if (searchInputIsTransactionHash) {
      const explorerLink: string = getTransactionHashExplorerLink({
        transactionHash: searchInput,
      })
      window.open(explorerLink, '_blank', 'noopener,noreferrer')
    }
  }, [searchInputIsTransactionHash])

  return (
    <div
      data-test-id="portfolio-search-bar"
      className={`
        relative flex items-center ml-auto
        border rounded-xl
        ${!mounted && 'border-opacity-30'}
        ${
          isFocused || isActive
            ? 'border-synapsePurple bg-tint'
            : 'border-separator bg-transparent'
        }
      `}
    >
      <FilterInput
        placeholder={placeholder}
        searchStr={searchInput}
        onSearch={onSearchInput}
        disabled={mounted ? false : true}
      />
      <ClearSearchButton show={isActive} onClick={clearSearchInput} />
    </div>
  )
}

const FilterInput = ({
  searchStr,
  onSearch,
  placeholder,
  disabled = false,
}: {
  searchStr: string
  onSearch: (str: string) => void
  placeholder: string
  disabled: boolean
}) => {
  return (
    <input
      disabled={disabled}
      ref={inputRef}
      tabIndex={0}
      data-test-id="filter-input"
      className={`
        flex-grow py-2.5 pl-4 pr-1
        font-normal text-sm text-primaryTextColor
        border h-full w-6/12 rounded-xl bg-transparent custom-shadow
        placeholder-white placeholder-opacity-40
        border-transparent outline-none ring-0
        focus:outline-none focus:ring-0 focus:border-transparent
        ${disabled && 'opacity-30'}
      `}
      placeholder={placeholder}
      onChange={(e) => onSearch(e.target.value)}
      value={searchStr}
    />
  )
}
