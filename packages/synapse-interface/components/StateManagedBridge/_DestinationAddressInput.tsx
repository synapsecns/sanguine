import { useState } from 'react'
import { useAppDispatch } from '@/store/hooks'
import { isValidAddress } from '@/utils/isValidAddress'
import { shortenAddress } from '@/utils/shortenAddress'
import { useBridgeState } from '@/slices/bridge/hooks'

export const _DestinationAddressInput = ({
  connectedAddress,
}: {
  connectedAddress: string
}) => {
  const dispatch = useAppDispatch()

  const { destinationAddress } = useBridgeState()

  const [isInputFocused, setIsInputFocused] = useState<boolean>(false)

  const handleInputFocus = () => {
    setIsInputFocused(true)
  }

  const handleInputBlur = () => {
    setIsInputFocused(false)
  }

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

  let placeholder

  if (isInputFocused) {
    placeholder = ''
  } else {
    placeholder = connectedAddress
      ? shortenAddress(connectedAddress)
      : 'Wallet address'
  }

  return (
    <div id="destination-address-input" onClick={handleActivateWarning}>
      <input
        onChange={(e) => handleInput(e.target.value)}
        onFocus={handleInputFocus}
        onBlur={handleInputBlur}
        placeholder={placeholder}
        value={inputValue}
        className={`
          rounded-md bg-surface border-separator text-secondary
          active:text-white focus:outline-none focus:ring-0 focus:border-separator
        `}
      />
      {/* {showWarning && (
        <DestinationInputWarning
          onEnableInput={onEnableInput}
          onReset={onReset}
        />
      )} */}
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
