import React, { Dispatch, SetStateAction } from 'react'

export const InputFilter = ({
  inputValue,
  setInputValue,
}: {
  inputValue: string
  setInputValue: Dispatch<SetStateAction<string>>
}) => {
  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(event.target.value)
  }
  return (
    <div data-test-id="input-filter">
      <input type="text" value={inputValue} onChange={handleInputChange} />
    </div>
  )
}
