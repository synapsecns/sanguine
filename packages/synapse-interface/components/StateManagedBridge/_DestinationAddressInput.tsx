import React, { useState } from 'react'
import { useAppDispatch, useAppSelector } from '@/store/hooks'
import { isValidAddress } from '@/utils/isValidAddress'
import { shortenAddress } from '@/utils/shortenAddress'
import { useBridgeState } from '@/slices/bridge/hooks'
import { setDestinationAddress } from '@/slices/bridge/reducer'
import { setShowDestinationWarning } from '@/slices/bridgeDisplaySlice'
import { Address } from 'viem'
import { isEmptyString } from '@/utils/isEmptyString'
import { CloseButton } from './components/CloseButton'

export const inputRef = React.createRef<HTMLInputElement>()

export const _DestinationAddressInput = ({
  connectedAddress,
}: {
  connectedAddress: string
}) => {
  const dispatch = useAppDispatch()
  const { destinationAddress } = useBridgeState()
  const { showDestinationWarning } = useAppSelector(
    (state) => state.bridgeDisplay
  )

  const handleClearInput = () => dispatch(setDestinationAddress('' as Address))

  const [isInputFocused, setIsInputFocused] = useState<boolean>(false)

  const handleInputFocus = () => setIsInputFocused(true)
  const handleInputBlur = () => setIsInputFocused(false)

  const isInputValidAddress = isValidAddress(destinationAddress)
  const isInputInvalid =
    destinationAddress &&
    !isEmptyString(destinationAddress) &&
    !isInputValidAddress

  let placeholder

  if (isInputFocused) {
    placeholder = ''
  } else {
    placeholder = connectedAddress
      ? shortenAddress(connectedAddress)
      : 'Wallet address'
  }

  /** Warning State */
  const [showWarning, setShowWarning] = useState<boolean>(false)

  const handleActivateWarning = () => {
    if (!showWarning && showDestinationWarning) {
      setShowWarning(!showWarning)
    }
  }

  const handleAcceptWarning = () => {
    inputRef.current.focus()
    setShowWarning(false)
    dispatch(setShowDestinationWarning(false))
  }

  const handleRejectWarning = () => {
    setShowWarning(false)
  }

  return (
    <div id="destination-address-input" onClick={handleActivateWarning}>
      {/* <div className="text-xs text-center text-white">
        {isInputValidAddress ? 'Valid Address' : 'Invalid Address'}
      </div> */}

      <div className="flex">
        <input
          ref={inputRef}
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
          text-md rounded-sm text-secondary py-1 px-2 z-0
          focus:text-white focus:outline-none focus:ring-0
          ${
            isInputValidAddress
              ? 'border-synapsePurple focus:border-synapsePurple'
              : isInputInvalid
              ? 'border-red-500 focus:border-red-500'
              : 'border-separator focus:border-separator'
          }
          ${connectedAddress ? 'w-32' : 'w-36'}
          ${
            isInputFocused
              ? 'text-left bg-bgBase'
              : 'text-center bg-transparent'
          }
        `}
        />
        <CloseButton onClick={handleClearInput} />
      </div>
      <DestinationInputWarning
        show={showWarning}
        onAccept={() => handleAcceptWarning()}
        onCancel={() => handleRejectWarning()}
      />
    </div>
  )
}

const DestinationInputWarning = ({
  show,
  onAccept,
  onCancel,
}: {
  show: boolean
  onAccept: () => void
  onCancel: () => void
}) => {
  return (
    <div
      className={`
      p-2 border rounded-sm bg-surface border-separator text-secondary
      top-0 left-0 w-full space-y-2 z-50
      ${show ? 'absolute' : 'hidden'}
      `}
    >
      <h3 className="text-2xl text-white">Warning</h3>
      <p className="text-white">
        Do not send your funds to a custodial wallet or exchange address!
      </p>
      <p className="text-secondary">
        It may be impossible to recover your funds
      </p>
      <div className="flex space-x-2">
        <button onClick={onCancel} className="w-1/2 py-3 bg-separator">
          Cancel
        </button>
        <button
          onClick={onAccept}
          className="w-1/2 py-3 bg-transparent border border-separator"
        >
          Okay
        </button>
      </div>
    </div>
  )
}
