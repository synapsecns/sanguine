import { useEffect, useRef, useMemo } from 'react'
import { Address } from 'viem'

import { useAppDispatch } from '@/store/hooks'
import {
  usePortfolioActionHandlers,
  usePortfolioState,
  fetchAndStoreSearchInputPortfolioBalances,
} from '@/slices/portfolio/hooks'
import { PortfolioState } from '@/slices/portfolio/reducer'
import { XIcon } from '@heroicons/react/outline'
import { initialState as portfolioInitialState } from '@/slices/portfolio/reducer'
import { isValidAddress } from '@/utils/isValidAddress'

export const SearchBar = () => {
  const dispatch = useAppDispatch()
  const { onSearchInput, clearSearchInput, clearSearchResults } =
    usePortfolioActionHandlers()
  const { searchInput, searchedBalancesAndAllowances }: PortfolioState =
    usePortfolioState()

  // const isSearchActive: boolean = searchInput.length > 0

  const isActive: boolean = searchInput !== portfolioInitialState.searchInput

  const searchInputIsAddress: boolean = useMemo(() => {
    return isValidAddress(searchInput)
  }, [searchInput])

  useEffect(() => {
    const searchResultsExist: boolean =
      Object.keys(searchedBalancesAndAllowances).length !== 0

    if (searchInputIsAddress) {
      dispatch(
        fetchAndStoreSearchInputPortfolioBalances(searchInput as Address)
      )
    }

    if (!searchInputIsAddress && searchResultsExist) {
      clearSearchResults()
    }
  }, [searchInputIsAddress])

  return (
    <div
      data-test-id="portfolio-search-bar"
      className={`
        relative flex items-center ml-auto
        border border-transparent
        ${isActive ? 'border-synapsePurple' : 'border-transparent'}
      `}
    >
      <FilterInput
        placeholder="Filter"
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
  const isActive: boolean = searchStr !== portfolioInitialState.searchInput

  return (
    <input
      data-test-id="filter-input"
      className={`
        flex-grow py-2 p-2
        font-normal text-sm text-primaryTextColor
        border h-full w-6/12 rounded bg-[#252226] custom-shadow
        placeholder-white placeholder-opacity-40
        outline-none ring-0 focus:outline-none focus:ring-0 border-transparent focus:border-transparent
      `}
      placeholder={placeholder}
      onChange={(e) => onSearch(e.target.value)}
      value={searchStr}
    />
  )
}

// focus:border-synapsePurple focus:outline-none focus:ring-0
// ${isActive ? 'border-synapsePurple' : 'border-transparent'}

export const ClearSearchButton = ({
  show,
  onClick,
}: {
  show: boolean
  onClick: () => void
}) => {
  return (
    <button
      className={`
        ${show ? 'block' : 'hidden'}
        flex w-6 h-6 right-2
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
