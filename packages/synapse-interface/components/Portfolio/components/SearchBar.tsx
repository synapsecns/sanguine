import React, { useEffect } from 'react'
import { Address } from 'viem'
import { useTranslations } from 'next-intl'

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
import { getTransactionHashExplorerLink } from '../../Activity/Transaction/components/TransactionExplorerLink'
import { ClearSearchButton } from './ClearSearchButton'
import { useIsFocused } from '@/utils/hooks/useIsFocused'
import { useIsMounted } from '@/utils/hooks/useIsMounted'
import { useSearchInputState } from '../hooks/useSearchInputStatus'

export const inputRef = React.createRef<HTMLInputElement>()

export const SearchBar = () => {
  const dispatch = useAppDispatch()
  const { onSearchInput, clearSearchInput } = usePortfolioActionHandlers()
  const { activeTab, searchInput, searchedBalances }: PortfolioState =
    usePortfolioState()
  const { isSearchInputActive, isMasqueradeActive } = useSearchInputState()

  const isMounted = useIsMounted()
  const isFocused = useIsFocused(inputRef)

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
        ${!isMounted && 'border-opacity-30'}
        ${
          isFocused || isSearchInputActive
            ? 'border-fuchsia-400 bg-tint'
            : 'border-separator bg-transparent'
        }
      `}
    >
      <FilterInput
        placeholder={placeholder}
        searchStr={searchInput}
        onSearch={onSearchInput}
        disabled={isMounted ? false : true}
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
  disabled: boolean
}) => {
  return (
    <input
      id="filter-input"
      ref={inputRef}
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
  const t = useTranslations('Portfolio')

  switch (activeTab) {
    case PortfolioTabs.PORTFOLIO:
      return `${t('Tokens, chains')}...`
    case PortfolioTabs.ACTIVITY:
      return `${t('Bridge txs')}...`
    default:
      return `${t('Search')}...`
  }
}
