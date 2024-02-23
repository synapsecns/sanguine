import React, { useEffect } from 'react'
import type { Address } from 'viem'
import { useAppDispatch } from '@/store/hooks'
import {
  usePortfolioActionHandlers,
  usePortfolioState,
  fetchAndStoreSearchInputPortfolioBalances,
} from '@/slices/portfolio/hooks'
import { PortfolioTabs } from '@/slices/portfolio/actions'
import { PortfolioState } from '@/slices/portfolio/reducer'
import { getValidAddress } from '@/utils/isValidAddress'
import { isTransactionHash } from '@/utils/validators'
import { getTransactionHashExplorerLink } from '../Transaction/components/TransactionExplorerLink'
import { ClearSearchButton } from './ClearSearchButton'
import { useSearchInputState } from '../helpers/useSearchInputStatus'

export const SearchBar = () => {
  const dispatch = useAppDispatch()
  const { onSearchInput, clearSearchInput } = usePortfolioActionHandlers()
  const { activeTab, searchInput, searchedBalances }: PortfolioState =
    usePortfolioState()
  const { isSearchInputActive, isMasqueradeActive } = useSearchInputState()

  const placeholder = getFilterPlaceholder(activeTab)

  const checksumValidAddress = getValidAddress(searchInput)
  const isSearchInputTransactionHash = isTransactionHash(searchInput)

  useEffect(() => {
    if (checksumValidAddress) {
      isMasqueradeActive
        ? clearSearchInput()
        : dispatch(
            fetchAndStoreSearchInputPortfolioBalances(
              checksumValidAddress as Address
            )
          )
    }
  }, [checksumValidAddress, searchedBalances])

  /** Trigger opening new browser window for tx on block explorer */
  useEffect(() => {
    if (isSearchInputTransactionHash) {
      const explorerLink: string = getTransactionHashExplorerLink({
        transactionHash: searchInput,
      })
      window.open(explorerLink, '_blank', 'noopener,noreferrer')
    }
  }, [isSearchInputTransactionHash])


  return (
    <div
      id="portfolio-search-bar"
      className={`
        relative flex items-center ml-auto
        border rounded-xl
       border-bgBase/10 bg-transparent
        focus-within:border-synapsePurple focus-within:bg-bgBase/10
        ${isSearchInputActive && 'border-synapsePurple bg-bgBase/10'}
      `}
    >
      <FilterInput
        placeholder={placeholder}
        searchStr={searchInput}
        onSearch={onSearchInput}
      />
      <ClearSearchButton
        show={isSearchInputActive}
        onClick={clearSearchInput}
      />
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
  disabled?: boolean
}) => {
  return (
    <input
      id="filter-input"
      autoComplete="off"
      placeholder={placeholder}
      onChange={(e) => onSearch(e.target.value)}
      value={searchStr}
      disabled={disabled}
      tabIndex={0}
      className={`
        flex-grow py-2.5 pl-4 pr-1
        font-normal text-sm text-primaryTextColor
        border h-full w-6/12 rounded-xl bg-transparent custom-shadow
        placeholder-white placeholder-opacity-40
        border-transparent outline-none ring-0
        focus:outline-none focus:ring-0 focus:border-transparent
        ${disabled && 'opacity-30'}
      `}
    />
  )
}

function getFilterPlaceholder(activeTab: PortfolioTabs | undefined) {
  switch (activeTab) {
    case PortfolioTabs.PORTFOLIO:
      return 'Tokens, chains...'
    case PortfolioTabs.ACTIVITY:
      return 'Bridge txs...'
    default:
      return 'Search...'
  }
}
