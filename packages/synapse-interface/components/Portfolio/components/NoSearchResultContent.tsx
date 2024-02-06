export const NoSearchResultsContent = ({
  searchStr,
}: {
  searchStr: string
}) => {
  return (
    <div id="no-search-results-content" className="text-white">
      <p className="mb-3 break-words">No results found for '{searchStr}'.</p>
    </div>
  )
}
