export function SearchResultsContainer({ label, children }) {
  return (
    <>
      <div className="px-2 pb-4 pt-2 text-sm text-primaryTextColor">
        {label}
      </div>
      <div className="px-2 pb-1 md:px-2">
        {children}
      </div>
    </>
  )
}

