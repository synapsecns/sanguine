import { useState } from 'react'
import { isValidAddress } from '@/utils/isValidAddress'

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

  const isInputValidAddress = isValidAddress(inputValue)

  return (
    <div id="destination-address-input">
      <input
        placeholder={connectedAddress ?? 'Connect Wallet'}
        value={inputValue}
        onChange={handleInput}
      />
      <div>{isInputValidAddress ? 'Valid Address' : 'Invalid Address'}</div>
    </div>
  )
}
