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
  const { onSearchInput, clearSearchInput } = usePortfolioActionHandlers()
  const { searchInput }: PortfolioState = usePortfolioState()

  const isSearchActive: boolean = searchInput.length > 0

  const inputIsAddress: boolean = useMemo(() => {
    return isValidAddress(searchInput)
  }, [searchInput])

  useEffect(() => {
    if (inputIsAddress) {
      dispatch(
        fetchAndStoreSearchInputPortfolioBalances(searchInput as Address)
      )
    }
  }, [inputIsAddress])

  return (
    <div
      data-test-id="portfolio-search-bar"
      className="relative flex items-center ml-auto"
    >
      <FilterInput
        placeholder="Filter"
        searchStr={searchInput}
        onSearch={onSearchInput}
      />
      <ClearSearchButton show={isSearchActive} onClick={clearSearchInput} />
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
        focus:border-[#D747FF] focus:outline-none focus:ring-0
        placeholder-white placeholder-opacity-40
        ${isActive ? 'border-[#D747FF]' : 'border-transparent'}
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
      className={`
        ${show ? 'absolute' : 'hidden'}
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
