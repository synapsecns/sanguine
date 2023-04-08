import { ChevronUpIcon } from '@heroicons/react/outline'

export const DrawerButton = ({
  className,
  onClick,
}: {
  className?: string
  onClick: () => void
}) => {
  return (
    <div
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
      <ChevronUpIcon className="inline w-6 text-white transition transform-gpu group-hover:opacity-50 group-active:rotate-180" />
    </div>
  )
}
