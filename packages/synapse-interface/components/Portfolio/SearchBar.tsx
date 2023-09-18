import SlideSearchBox from '@/pages/bridge/SlideSearchBox'
import {
  usePortfolioActionHandlers,
  usePortfolioState,
} from '@/slices/portfolio/hooks'
import { PortfolioState } from '@/slices/portfolio/reducer'
import { XIcon } from '@heroicons/react/outline'

export const SearchBar = () => {
  const { onSearchInput, clearSearchInput } = usePortfolioActionHandlers()
  const { searchInput }: PortfolioState = usePortfolioState()

  const isSearchActive: boolean = searchInput.length > 0

  return (
    <div
      data-test-id="portfolio-search-bar"
      className="relative flex items-center ml-auto"
    >
      <SlideSearchBox
        placeholder="Filter"
        searchStr={searchInput}
        onSearch={onSearchInput}
        focusOnMount={false}
      />
      <ClearSearchButton show={isSearchActive} onClick={clearSearchInput} />
    </div>
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
        flex w-6 h-6 right-1
        items-center justify-center
        border border-separator rounded-full
        hover:cursor-pointer hover:border-0
      `}
      onClick={onClick}
    >
      <XIcon strokeWidth={3} className="inline w-4 text-secondary" />
    </button>
  )
}
