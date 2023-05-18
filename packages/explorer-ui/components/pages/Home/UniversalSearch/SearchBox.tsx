export function SearchBox({
  searchField,
  setSearchField,
  inputType,
  placeholder,
  extraStyling,
}) {
  return (
    <form className="flex items-center">
      <div className="relative w-full group ">
        <input
          type="text"
          id="simple-search"
          className={
            `
            bg-white bg-opacity-5
            rounded-md
            border border-white border-opacity-20
            focus:outline-none focus-within:border-gray-500
            block w-full  px-4 py-2
            text-white
            placeholder:text-white placeholder:text-opacity-60

          ` + extraStyling
          }
          placeholder={placeholder}
          onChange={(e) => {
            setSearchField(e.target.value)
          }}
          value={searchField}
        />
      </div>
    </form>
  )
}
