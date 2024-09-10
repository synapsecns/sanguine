import { useTranslations } from 'next-intl'

import { SettingsIcon } from '../icons/SettingsIcon'

export const SettingsToggle = ({
  showSettingsToggle,
}: {
  showSettingsToggle: boolean
}) => {
  const t = useTranslations('Settings')
  return (
    <>
      {showSettingsToggle ? (
        <>
          <SettingsIcon className="w-5 h-5 mr-2" />
          <span>{t('Settings')}</span>
        </>
      ) : (
        <span>{t('Close')}</span>
      )}
    </>
  )
}
