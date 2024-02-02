import { useState } from 'react'

export const InputFilter = () => {
  const [inputValue, setInputValue] = useState('')

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(event.target.value)
  }

  return (
    <div data-test-id="input-filter">
      <input type="text" value={inputValue} onChange={handleInputChange} />
    </div>
  )
}
