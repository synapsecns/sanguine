import { SettingsIcon } from '../icons/SettingsIcon'

export const SettingsToggle = ({
  showSettingsToggle,
}: {
  showSettingsToggle: boolean
}) => {
  return (
    <>
      {showSettingsToggle ? (
        <>
          <SettingsIcon className="w-5 h-5 mr-2" />
          <span>Settings</span>
        </>
      ) : (
        <span>Close</span>
      )}
    </>
  )
}
