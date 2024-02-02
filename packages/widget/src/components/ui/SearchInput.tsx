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
    <div
      data-test-id="search-input"
      className="p-2 border border-solid rounded border-[--synapse-focus]"
      style={{ background: 'var(--synapse-select-bg)' }}
    >
      <input
        type="text"
        ref={inputRef}
        placeholder={placeholder}
        value={inputValue}
        onChange={handleInputChange}
        style={{ background: 'var(--synapse-select-bg)' }}
        className={`
          text-[--primary] placeholder:text-[--synapse-secondary]
          w-full border-none shadow-none text-base
          focus:ring-0 focus:border-none focus:outline-none
        `}
      />
    </div>
  )
}
