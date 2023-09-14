import SlideSearchBox from '@/pages/bridge/SlideSearchBox'

export const SearchBar = () => {
  return (
    <div data-test-id="portfolio-search-bar" className="ml-auto">
      <SlideSearchBox placeholder="Filter" searchStr="" onSearch={() => null} />
    </div>
  )
}
