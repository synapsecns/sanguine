import { useState } from 'react'

export const _DestinationAddressInput = ({
  connectedAddress,
}: {
  connectedAddress: string
}) => {
  const [showWarning, setShowWarning] = useState<boolean>(true)
  const [inputValue, setInputValue] = useState<string>('')

  const handleInput = (value) => {
    setInputValue(value)
  }

  return (
    <div id="destination-address-input">
      <input
        placeholder={connectedAddress ?? 'Connect Wallet'}
        value={inputValue}
        onChange={handleInput}
      />
    </div>
  )
}
