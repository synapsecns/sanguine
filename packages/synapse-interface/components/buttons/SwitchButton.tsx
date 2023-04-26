import { SwitchVerticalIcon } from '@heroicons/react/outline'

export default function SwitchButton({
  className,
  innerClassName,
  onClick,
}: {
  className?: string
  innerClassName?: string
  onClick: () => void
}) {
  return (
    <div
      className={`
        rounded-full p-2 -mr-2 -ml-2 hover:cursor-pointer select-none
        ${className}
      `}
    >
      <div
        onClick={onClick}
        className={`
          group rounded-full inline-block p-2
          bg-bgLighter
          transform-gpu transition-all duration-100
          active:rotate-90

          ${className}
          ${innerClassName}
        `}
      >
        <SwitchVerticalIcon
          className={`
            w-6 h-6 transition-all
            text-white group-hover:text-opacity-50
          `}
        />
      </div>
    </div>
  )
}
