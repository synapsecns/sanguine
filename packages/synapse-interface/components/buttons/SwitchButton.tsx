import { useState } from 'react'
import { joinClassNames } from '@/utils/joinClassNames'

const ms = 300

export function SwitchButton({
  onClick,
  disabled,
}: {
  onClick: () => void
  disabled: boolean
}) {
  const [isActive, setIsActive] = useState(false)

  const handleClick = () => {
    if (!disabled) {
      onClick()
      setIsActive(true)
      setTimeout(() => setIsActive(false), ms)
    }
  }

  const className = joinClassNames({
    space: '-mt-4 -mb-4 rounded z-10 justify-self-center',
    background: 'bg-zinc-50 dark:bg-bgLight', // TODO: Remove
    // background: 'bg-zinc-50 dark:bg-zinc-800/50',
    border: 'border border-zinc-200 dark:border-bgBase', // TODO: Remove
    // border: 'border border-zinc-200 dark:border-zinc-900/95',
    stroke: 'stroke-2 stroke-zinc-500 dark:stroke-secondary', // TODO: Remove
    // stroke: 'stroke-2 stroke-zinc-500 dark:stroke-zinc-400',
    transition: `hover:opacity-80 cursor-pointer transition-transform ${
      isActive ? `duration-${ms} rotate-180 ease-in-out` : 'ease-out' // 'duration-0'
    }`,
  })

  return (
    <svg
      onClick={handleClick}
      className={className}
      width="32"
      height="32"
      viewBox="0 0 32 32"
      xmlns="http://www.w3.org/2000/svg"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M11,22V8M11,8L7,12M11,8L15,12" />
      <path d="M21,9V23M21,23L25,19M21,23L17,19" />
    </svg>
  )
}
