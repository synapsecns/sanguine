import { useState } from 'react'
import { useAppDispatch } from '@/store/hooks'
import { isValidAddress } from '@/utils/isValidAddress'
import { shortenAddress } from '@/utils/shortenAddress'
import { useBridgeState } from '@/slices/bridge/hooks'
import { setDestinationAddress } from '@/slices/bridge/reducer'
import { Address } from 'viem'

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

  const isInputValidAddress = isValidAddress(destinationAddress)

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
        onChange={(e) =>
          dispatch(setDestinationAddress(e.target.value as Address))
        }
        onFocus={handleInputFocus}
        onBlur={handleInputBlur}
        placeholder={placeholder}
        value={
          isInputValidAddress && !isInputFocused
            ? shortenAddress(destinationAddress)
            : destinationAddress
        }
        className={`
          text-md rounded-sm bg-bgBase border-separator text-secondary py-1 px-2
          focus:text-white focus:outline-none focus:ring-0
          ${
            isInputValidAddress
              ? 'border-synapsePurple focus:border-synapsePurple'
              : 'focus:border-separator'
          }
          ${connectedAddress ? 'w-32' : 'w-36'}
          ${isInputFocused ? 'text-left' : 'text-center'}
        `}
      />
      {/* {showWarning && (
        <DestinationInputWarning
          onEnableInput={onEnableInput}
          onReset={onReset}
        />
      )} */}
      <div className="text-white">
        {isInputValidAddress ? 'Valid Address' : 'Invalid Address'}
      </div>
    </div>
  )
}

const DestinationInputWarning = ({ onEnableInput, onReset }) => {
  return (
    <div>
      <h3>Warning</h3>
      <p>Do not send your funds to a custodial wallet or exchange address!</p>
      <p>It may be impossible to recover your funds</p>
      <div className="flex">
        <button onClick={onReset}>Cancel</button>
        <button onClick={onEnableInput}>Okay</button>
      </div>
    </div>
  )
}
