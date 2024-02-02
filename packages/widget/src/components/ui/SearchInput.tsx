import React, { Dispatch, SetStateAction } from 'react'

export const SearchInput = ({
  inputValue,
  setInputValue,
  placeholder,
}: {
  inputValue: string
  setInputValue: Dispatch<SetStateAction<string>>
  placeholder: string
}) => {
  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(event.target.value)
  }
  return (
    <div
      data-test-id="search-input"
      className="p-2 border border-solid rounded-lg border-[--synapse-focus]"
      style={{ background: 'var(--synapse-select-bg)' }}
    >
      <input
        type="text"
        placeholder={placeholder}
        value={inputValue}
        onChange={handleInputChange}
        className={`
          text-[--synapse-secondary]
          w-full border-none shadow-none
          focus:ring-0 focus:border-none focus:outline-none
        `}
        style={{ background: 'var(--synapse-select-bg)' }}
      />
    </div>
  )
}
