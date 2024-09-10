import { useTranslations } from 'next-intl'

import { useAppDispatch } from '@/store/hooks'
import { useBridgeDisplayState } from '@/slices/bridge/hooks'
import { setIsDestinationWarningAccepted } from '@/slices/bridgeDisplaySlice'

export const ConfirmDestinationAddressWarning = () => {
  const dispatch = useAppDispatch()
  const t = useTranslations('Destination')

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
              text-fuchsia-400 bg-transparent outline-none
              focus:!outline-0 focus:ring-0 focus:!border-0
              active:!outline-0 active:ring-0 active:!border-0
            `}
        />
        <div>
          <p className="text-sm text-secondary">
            <b>{t('WarningMessage1')}</b> {t('WarningMessage2')} <br />
            <b>{t('WarningMessage3')}</b> {t('WarningMessage4')}
          </p>
        </div>
      </div>
    )
  }
}
