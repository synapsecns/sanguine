import Dropdown from 'react-dropdown'
import 'react-dropdown/style.css'
import { useState } from 'react'

const options = ['one', 'two', 'three']
const defaultOption = options[0]

export const TransactionOptions = () => {
  const [selectedOption, setSelectedOption] = useState(defaultOption)

  const onSelect = (option) => {
    setSelectedOption(option.value)
    // You can perform additional actions here if needed
  }

  return (
    <Dropdown
      options={options.map((option) => ({ value: option, label: option }))} // Adapt the options format to match what Dropdown expects
      onChange={onSelect}
      value={selectedOption}
      placeholder="Select an option"
    />
  )
}
