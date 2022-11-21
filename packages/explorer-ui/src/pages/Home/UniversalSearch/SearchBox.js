export function SearchBox({ searchField, setSearchField, inputType }) {
  return (
    <form className="flex items-center mt-5">
      <div className="relative w-full group ">
        <input
          type="text"
          id="simple-search"
          className={`
            bg-white bg-opacity-10
            text-white text-opacity-40 font-medium text-xl
            rounded-lg
            border border-white border-opacity-10
            focus:outline-none focus-within:border-gray-500
            block w-full pl-5 p-3.5
          `}
          placeholder="Search by address or transaction..."
          onChange={(e) => {
            setSearchField(e.target.value)
          }}
          value={searchField}
        />
      </div>
    </form>
  )
}
