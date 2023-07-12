import { SwitchVerticalIcon } from '@heroicons/react/outline'
import { useState } from 'react'

export default function SwitchButton({
  className,
  innerClassName,
  onClick,
}: {
  className?: string
  innerClassName?: string
  onClick: () => void
}) {
  const [isActive, setIsActive] = useState(false)

  const handleClick = () => {
    onClick()
    setIsActive(true)
    setTimeout(() => setIsActive(false), 200)
  }

  return (
    <div
      className={`
        rounded-full p-2 -mr-2 -ml-2 hover:cursor-pointer select-none
        ${className}
      `}
    >
      <div
        onClick={handleClick}
        className={`
          group rounded-full inline-block p-2
          bg-bgLighter
          transform-gpu transition-all duration-200
          ${isActive ? 'rotate-90' : ''}
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
