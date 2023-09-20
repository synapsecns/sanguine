import { useEffect, useRef, useMemo } from 'react'
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

export const SearchBar = () => {
  const dispatch = useAppDispatch()
  const { onSearchInput, clearSearchInput, clearSearchResults } =
    usePortfolioActionHandlers()
  const {
    activeTab,
    searchInput,
    searchedBalancesAndAllowances,
  }: PortfolioState = usePortfolioState()

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
        border bg-[#252226] rounded-sm
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
  const { clearSearchInput } = usePortfolioActionHandlers()
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
      <ClearSearchButton onClick={clearSearchInput} show={true} />
    </div>
  )
}
