import React, { useEffect, useMemo, useState } from 'react'
import { Address } from 'viem'
import { XIcon } from '@heroicons/react/outline'

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
import { isValidAddress, getValidAddress } from '@/utils/isValidAddress'
import { shortenAddress } from '@/utils/shortenAddress'
import { isTransactionHash } from '@/utils/validators'
import { getTransactionHashExplorerLink } from './Transaction/components/TransactionExplorerLink'

export const inputRef = React.createRef<HTMLInputElement>()

export const SearchBar = () => {
  const dispatch = useAppDispatch()
  const { onSearchInput, clearSearchInput, clearSearchResults } =
    usePortfolioActionHandlers()
  const { activeTab, searchInput, searchedBalances }: PortfolioState =
    usePortfolioState()

  const [mounted, setMounted] = useState<boolean>(false)
  const [isFocused, setIsFocused] = useState<boolean>(false)
  const isActive: boolean = searchInput !== portfolioInitialState.searchInput

  useEffect(() => {
    setMounted(true)
  }, [])

  useEffect(() => {
    const handleFocus = () => setIsFocused(true)
    const handleBlur = () => setIsFocused(false)
    const input = inputRef.current

    if (input) {
      input.addEventListener('focus', handleFocus)
      input.addEventListener('blur', handleBlur)
      return () => {
        input.removeEventListener('focus', handleFocus)
        input.removeEventListener('blur', handleBlur)
      }
    }
  }, [inputRef])

  const placeholder: string = useMemo(() => {
    switch (activeTab) {
      case PortfolioTabs.PORTFOLIO:
        return 'Tokens, chains…'
      case PortfolioTabs.ACTIVITY:
        return 'Bridge txs…'
      default:
        return 'Search…'
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
        flex items-center border rounded-full ml-auto pr-2 relative
        ${!mounted && 'border-opacity-30'}
        ${
          isFocused || isActive
            ? 'border-fuchsia-500 bg-zinc-100 dark:bg-zinc-800'
            : 'border-zinc-500'
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

export default function FilterInput({
  searchStr,
  onSearch,
  placeholder,
  disabled = false,
}: {
  searchStr: string
  onSearch: (str: string) => void
  placeholder: string
  disabled: boolean
}) {
  return (
    <input
      disabled={disabled}
      ref={inputRef}
      tabIndex={0}
      data-test-id="filter-input"
      className={`
        flex-grow py-2.5 pl-4 pr-1
        font-normal text-sm w-1/2
        bg-transparent custom-shadow
        border-transparent outline-none ring-0
        focus:outline-none focus:ring-0 focus:border-transparent
        ${disabled && 'opacity-40'}
      `}
      placeholder={placeholder}
      onChange={(e) => onSearch(e.target.value)}
      value={searchStr}
    />
  )
}

export const ClearSearchButton = ({
  show,
  onClick,
}: {
  show: boolean
  onClick: () => void
}) => {
  return (
    <button
      data-test-id="clear-search-button"
      className={`
        ${show ? 'visible' : 'invisible'}
        w-6 h-6
        flex items-center justify-center
        border border-current rounded-full
        opacity-70
        hover:cursor-pointer hover:opacity-100
      `}
      onClick={onClick}
    >
      <XIcon strokeWidth={3} className="w-4" />
    </button>
  )
}

export const ViewSearchAddressBanner = ({
  viewingAddress,
}: {
  viewingAddress: Address
}) => {
  const { clearSearchResults } = usePortfolioActionHandlers()
  const shortened: string = shortenAddress(viewingAddress, 4)
  return (
    <div
      data-test-id="view-search-address-banner"
      className={`
        flex justify-between p-3 mb-3
        border border-purple-500 rounded
      `}
      style={{
        background:
          'linear-gradient(310.65deg, rgba(172, 143, 255, 0.2) -17.9%, rgba(255, 0, 255, 0.2) 86.48%)',
      }}
    >
      <div>
        Viewing <span className="font-medium">{shortened}</span>
      </div>
      <ClearSearchButton onClick={clearSearchResults} show={true} />
    </div>
  )
}
