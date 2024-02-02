import React, { Dispatch, SetStateAction } from 'react'

export const InputFilter = ({
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
    <div data-test-id="input-filter" className="p-2 bg-white">
      <input
        type="text"
        placeholder={placeholder}
        value={inputValue}
        onChange={handleInputChange}
        className={`
          w-full
          border-none shadow-none
          focus:ring-0 focus:border-none focus:outline-none
        `}
      />
    </div>
  )
}
