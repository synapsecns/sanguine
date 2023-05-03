import { ChevronUpIcon } from '@heroicons/react/outline'

export const DrawerButton = ({
  className,
  onClick,
  isOrigin,
}: {
  className?: string
  onClick: () => void
  isOrigin: boolean
}) => {
  const dataId = isOrigin
    ? 'bridge-origin-list-close-button'
    : 'bridge-destination-list-close-button'

  return (
    <div
      data-test-id={dataId}
      className={`
        flex items-center justify-center
        w-8 h-8
        float-right
        group
        hover:cursor-pointer
        rounded-full
        bg-white bg-opacity-10
        ${className}
      `}
      onClick={onClick}
    >
      <ChevronUpIcon
        className={`
        inline w-6 text-white transition-all ease-in
        duration-75 transform-gpu group-hover:opacity-50
        group-active:-rotate-180
        `}
      />
    </div>
  )
}
