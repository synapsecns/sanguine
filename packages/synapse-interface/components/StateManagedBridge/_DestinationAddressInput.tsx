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

  const handleInputFocus = () => setIsInputFocused(true)
  const handleInputBlur = () => setIsInputFocused(false)

  const isInputValidAddress = isValidAddress(destinationAddress)

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
    setShowWarning(!showWarning)
  }
  const handleAcceptWarning = () => {
    setShowWarning(false)
  }
  const handleRejectWarning = () => {
    setShowWarning(false)
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
      <DestinationInputWarning
        show={showWarning}
        onAccept={() => handleAcceptWarning()}
        onCancel={() => handleRejectWarning()}
      />
      <div className="text-white">
        {isInputValidAddress ? 'Valid Address' : 'Invalid Address'}
      </div>
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
      top-0 left-0 w-full h-full
      ${show ? 'absolute' : 'hidden'}
      `}
    >
      <h3>Warning</h3>
      <p>Do not send your funds to a custodial wallet or exchange address!</p>
      <p>It may be impossible to recover your funds</p>
      <div className="flex">
        <button onClick={onCancel}>Cancel</button>
        <button onClick={onAccept}>Okay</button>
      </div>
    </div>
  )
}
