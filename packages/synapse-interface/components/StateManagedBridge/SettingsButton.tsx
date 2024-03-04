import { XIcon, CogIcon } from '@heroicons/react/outline'
import Button from '@tw/Button'
export function SettingsButton({
  show,
  onClick,
}: {
  show: boolean
  onClick: () => void
}) {
  return (
    <Button
      className={`
        group flex items-center p-2 text-opacity-75 bg-bgBase/10 hover:bg-bgBase/20
        ring-1 ring-white/10 hover:ring-white/30 text-secondaryTextColor hover:text-white
       active:ring-white/60
        transition-all duration-100
        overflow-hidden
        ${show ? "w-[36px]" : "w-[96.27px]"}
      `}
      onClick={onClick}
    >
      {show ? (
        // the size-[36px] -m-2 p-[9px] is a hackfix for issue in click propagation bug in react/react-redux
        <XIcon
          key="settingsIconMorph"
          className="size-[36px] -m-2 p-[9px]"
        />
      ) : (
        <>
          <CogIcon
            key="settingsIconMorph"
            className="w-4 h-4 mr-2 group-hover:animate-spin"
          />
          <span className="text-sm mr-1">
            Settings
          </span>
        </>
      )}
    </Button>
  )
}