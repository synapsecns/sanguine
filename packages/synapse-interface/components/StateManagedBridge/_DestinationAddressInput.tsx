import { useState } from 'react'
import { isValidAddress } from '@/utils/isValidAddress'

export const _DestinationAddressInput = ({
  connectedAddress,
}: {
  connectedAddress: string
}) => {
  // TODO: Lift state up to Redux slice after working example
  const [enableInput, setEnableInput] = useState<boolean>(false)
  const [showWarning, setShowWarning] = useState<boolean>(false)
  const [inputValue, setInputValue] = useState<string>('')

  const handleInput = (value) => {
    setInputValue(value)
  }

  const handleActivateWarning = () => {
    setShowWarning(true)
  }

  const onEnableInput = () => {
    setEnableInput(true)
  }

  const onReset = () => {
    setShowWarning(false)
    setEnableInput(false)
  }

  const isInputValidAddress = isValidAddress(inputValue)

  return (
    <div id="destination-address-input" onClick={handleActivateWarning}>
      <input
        placeholder={connectedAddress ?? 'Connect Wallet'}
        value={inputValue}
        onChange={handleInput}
      />
      {showWarning && (
        <DestinationInputWarning
          onEnableInput={onEnableInput}
          onReset={onReset}
        />
      )}
      <div>{isInputValidAddress ? 'Valid Address' : 'Invalid Address'}</div>
    </div>
  )
}

const DestinationInputWarning = ({ onEnableInput, onReset }) => {
  return (
    <div>
      Warning!
      <div className="flex">
        <button onClick={onEnableInput}>Accept</button>
        <button onClick={onReset}>Cancel</button>
      </div>
    </div>
  )
}
