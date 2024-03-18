export const SearchResults = ({
  searchStr,
  type,
}: {
  searchStr: string
  type: string
}) => {
  return (
    <div>
      {searchStr && (
        <div className="p-2 text-sm">
          No other results found for <q>{searchStr}</q>.
          <div className="pt-2 align-bottom text-primaryTextColor text-md">
            Want to see a {type} supported on Synapse? Let us know!
          </div>
        </div>
      )}
    </div>
  )
}
