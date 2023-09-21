import React, { useEffect, useRef, useMemo, useState } from 'react'
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
import { isValidAddress } from '@/utils/isValidAddress'
import { shortenAddress } from '@/utils/shortenAddress'
import { isTransactionHash } from '@/utils/validators'
import { getTransactionHashExplorerLink } from './Transaction/components/TransactionExplorerLink'

export const inputRef = React.createRef<HTMLInputElement>()

export const SearchBar = () => {
  const dispatch = useAppDispatch()
  const { onSearchInput, clearSearchInput, clearSearchResults } =
    usePortfolioActionHandlers()
  const {
    activeTab,
    searchInput,
    searchedBalancesAndAllowances,
  }: PortfolioState = usePortfolioState()

  const [isFocused, setIsFocused] = useState(false)

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

  const isActive: boolean = searchInput !== portfolioInitialState.searchInput

  const placeholder: string = useMemo(() => {
    switch (activeTab) {
      case PortfolioTabs.PORTFOLIO:
        return 'Search tokens & chains'
      case PortfolioTabs.ACTIVITY:
        return 'Search transactions'
      default:
        return 'Filter'
    }
  }, [activeTab])

  const searchInputIsAddress: boolean = useMemo(() => {
    return isValidAddress(searchInput)
  }, [searchInput])

  const searchInputIsTransactionHash: boolean = useMemo(() => {
    return isTransactionHash(searchInput)
  }, [searchInput])

  useEffect(() => {
    const masqueradeActive: boolean =
      Object.keys(searchedBalancesAndAllowances).length > 0
    if (searchInputIsAddress && !masqueradeActive) {
      dispatch(
        fetchAndStoreSearchInputPortfolioBalances(searchInput as Address)
      )
    }
  }, [searchInputIsAddress, searchedBalancesAndAllowances])

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
        border bg-[#252226] rounded-xl
        ${isActive ? 'border-synapsePurple' : 'border-transparent'}
      `}
    >
      <FilterInput
        placeholder={placeholder}
        searchStr={searchInput}
        onSearch={onSearchInput}
      />
      <ClearSearchButton show={isActive} onClick={clearSearchInput} />
    </div>
  )
}

export default function FilterInput({
  searchStr,
  onSearch,
  placeholder,
}: {
  searchStr: string
  onSearch: (str: string) => void
  placeholder: string
}) {
  return (
    <input
      ref={inputRef}
      tabIndex={0}
      data-test-id="filter-input"
      className={`
        flex-grow py-2 px-4
        font-normal text-sm text-primaryTextColor
        border h-full w-6/12 rounded-xl bg-[#252226] custom-shadow
        placeholder-white placeholder-opacity-40
        border-transparent outline-none ring-0
        focus:outline-none focus:ring-0 focus:border-transparent
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
        ${show ? 'z-10' : 'z-[-10]'}
        flex w-6 h-6 mr-2
        items-center justify-center
        border border-separator rounded-full
        hover:cursor-pointer hover:border-secondary
      `}
      onClick={onClick}
    >
      <XIcon strokeWidth={3} className="inline w-4 text-secondary" />
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
        border border-synapsePurple rounded-sm
      `}
      style={{
        background:
          'linear-gradient(310.65deg, rgba(172, 143, 255, 0.2) -17.9%, rgba(255, 0, 255, 0.2) 86.48%)',
      }}
    >
      <div className="flex space-x-1">
        <div className="text-secondary ">Viewing</div>
        <div className="font-bold text-primary">{shortened}</div>
      </div>
      <ClearSearchButton onClick={clearSearchResults} show={true} />
    </div>
  )
}
