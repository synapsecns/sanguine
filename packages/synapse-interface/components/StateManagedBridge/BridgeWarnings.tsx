import { useAppDispatch } from '@/store/hooks'
import { useBridgeDisplayState } from '@/slices/bridge/hooks'
import { setIsDestinationWarningAccepted } from '@/slices/bridgeDisplaySlice'

export const ConfirmDestinationAddressWarning = () => {
  const dispatch = useAppDispatch()
  const {
    showDestinationWarning,
    isDestinationWarningAccepted,
    showDestinationAddress,
  } = useBridgeDisplayState()

  const handleCheckboxChange = () => {
    dispatch(setIsDestinationWarningAccepted(!isDestinationWarningAccepted))
  }

  if (showDestinationAddress && showDestinationWarning) {
    return (
      <div
        className="flex items-center mb-2 space-x-3 cursor-pointer"
        onClick={handleCheckboxChange}
      >
        <input
          type="checkbox"
          id="destination-warning"
          name="destinationWarning"
          value=""
          checked={isDestinationWarningAccepted}
          onChange={handleCheckboxChange}
          className={`
              cursor-pointer border rounded-[4px] border-secondary
             text-synapsePurple bg-transparent outline-none
              focus:!outline-0 focus:ring-0 focus:!border-0
              active:!outline-0 active:ring-0 active:!border-0
            `}
        />
        <div>
          <p className="text-sm text-secondary">
            <strong>Required:</strong> Verify your destination address to
            continue.
            <br />
            Do <strong>not</strong> send assets to a custodial or exchange
            address. It may be impossible to recover your funds.
          </p>
        </div>
      </div>
    )
  }
}
