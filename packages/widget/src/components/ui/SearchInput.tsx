import React, { Dispatch, SetStateAction, useRef, useEffect } from 'react'

export const SearchInput = ({
  inputValue,
  setInputValue,
  placeholder,
  isActive,
}: {
  inputValue: string
  setInputValue: Dispatch<SetStateAction<string>>
  placeholder: string
  isActive: boolean
}) => {
  const inputRef = useRef<HTMLInputElement>(null)

  // Focus on the input when isActive
  useEffect(() => {
    if (isActive && inputRef.current) {
      inputRef.current.focus()
    }
  }, [isActive])

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(event.target.value)
  }
  return (
    <div data-test-id="search-input">
      <input
        type="text"
        ref={inputRef}
        placeholder={placeholder}
        value={inputValue}
        onChange={handleInputChange}
        style={{ background: 'var(--synapse-select-bg)' }}
        className={`
          text-[--synapse-text] placeholder:text-[--synapse-secondary]
          w-full border border-solid border-[--synapse-focus] shadow-none text-sm
          focus:ring-0 focus:border-[--synapse-focus] focus:outline-none
          px-2 py-1.5 rounded
        `}
      />
    </div>
  )
}
