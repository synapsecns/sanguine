export function SearchBox({
  searchField,
  setSearchField,
  inputType,
  placeholder,
  extraStyling,
}) {
  return (
    // @ts-expect-error TS(2304): Cannot find name 'form'.
    <form className="flex items-center">
      // @ts-expect-error TS(2304): Cannot find name 'div'.
      <div className="relative w-full group ">
        // @ts-expect-error TS(2304): Cannot find name 'input'.
        <input
          // @ts-expect-error TS(2304): Cannot find name 'type'.
          type="text"
          // @ts-expect-error TS(2304): Cannot find name 'id'.
          id="simple-search"
          // @ts-expect-error TS(2304): Cannot find name 'className'.
          className={`
            bg-white bg-opacity-5
            rounded-md
            border border-white border-opacity-20
            focus:outline-none focus-within:border-gray-500
            block w-full  px-4 py-2
            text-white
            placeholder:text-white placeholder:text-opacity-60

          ` +extraStyling}
          // @ts-expect-error TS(2304): Cannot find name 'placeholder'.
          placeholder={placeholder}
          // @ts-expect-error TS(2304): Cannot find name 'onChange'.
          onChange={(e) => {
            // @ts-expect-error TS(2304): Cannot find name 'setSearchField'.
            setSearchField(e.target.value)
          }}
          // @ts-expect-error TS(2304): Cannot find name 'value'.
          value={searchField}
        />
      </div>
    </form>
  )
}
