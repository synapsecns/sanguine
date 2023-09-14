import Fuse from 'fuse.js'
import SlideSearchBox from '@/pages/bridge/SlideSearchBox'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { NetworkTokenBalancesAndAllowances } from '@/utils/actions/fetchPortfolioBalances'

export const SearchBar = () => {
  const userTokenBalances: NetworkTokenBalancesAndAllowances =
    usePortfolioBalances()

  return (
    <div data-test-id="portfolio-search-bar" className="ml-auto">
      <SlideSearchBox placeholder="Filter" searchStr="" onSearch={() => null} />
    </div>
  )
}
