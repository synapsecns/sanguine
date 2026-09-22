import { useTranslations } from 'next-intl'

import { useAppDispatch } from '@/store/hooks'
import { useBridgeDisplayState } from '@/slices/bridge/hooks'
import { setIsDestinationWarningAccepted } from '@/slices/bridgeDisplaySlice'
import { WarningMessage } from '@/components/Warning'

export const SynHyperCoreRecipientWarning = ({
  accepted,
  onAccept,
}: {
  accepted: boolean
  onAccept: (accepted: boolean) => void
}) => {
  const t = useTranslations('Bridge')

  return (
    <WarningMessage
      twClassName="mb-2 border border-amber-500/30 !bg-amber-50 !text-amber-950 dark:!bg-amber-950/30 dark:!text-amber-100"
      header={t('SynHyperCoreRecipientWarningTitle')}
      message={
        <>
          <p>{t('SynHyperCoreRecipientWarningMessage')}</p>
          <label className="flex items-start gap-2 mt-3 cursor-pointer">
            <input
              type="checkbox"
              checked={accepted}
              onChange={(event) => onAccept(event.target.checked)}
              className="mt-0.5 rounded border-amber-500/50 text-amber-500 bg-transparent focus:ring-amber-500"
            />
            <span>{t('SynHyperCoreRecipientWarningAccept')}</span>
          </label>
        </>
      }
    />
  )
}

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
