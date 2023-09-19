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
import { shortenAddress } from '@/utils/shortenAddress'

export const SearchBar = () => {
  const dispatch = useAppDispatch()
  const { onSearchInput, clearSearchInput, clearSearchResults } =
    usePortfolioActionHandlers()
  const { searchInput, searchedBalancesAndAllowances }: PortfolioState =
    usePortfolioState()

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
        border border-transparent bg-[#252226] rounded-sm
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
  return (
    <input
      data-test-id="filter-input"
      className={`
        flex-grow py-2 p-2
        font-normal text-sm text-primaryTextColor
        border h-full w-6/12 rounded bg-[#252226] custom-shadow
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
        flex w-6 h-6 mr-1
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
  const shortened: string = shortenAddress(viewingAddress, 3)

  return <div data-test-id="view-search-address-banner"></div>
}
