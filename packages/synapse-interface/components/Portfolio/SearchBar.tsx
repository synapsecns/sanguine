import Fuse from 'fuse.js'
import SlideSearchBox from '@/pages/bridge/SlideSearchBox'
import { NetworkTokenBalancesAndAllowances } from '@/utils/actions/fetchPortfolioBalances'
import {
  usePortfolioActionHandlers,
  usePortfolioState,
} from '@/slices/portfolio/hooks'
import { PortfolioState } from '@/slices/portfolio/reducer'

export const SearchBar = () => {
  const { onSearchInput } = usePortfolioActionHandlers()

  const { searchInput }: PortfolioState = usePortfolioState()

  return (
    <div data-test-id="portfolio-search-bar" className="ml-auto">
      <SlideSearchBox
        placeholder="Filter"
        searchStr={searchInput}
        onSearch={onSearchInput}
      />
    </div>
  )
}
