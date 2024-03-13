import { useState } from 'react'

const join = (a) => Object.values(a).join(' ')

export function SwitchButton({ onClick }: { onClick: () => void }) {
  const [isActive, setIsActive] = useState(false)
  const ms = 300
  const handleClick = () => {
    onClick()
    setIsActive(true)
    setTimeout(() => setIsActive(false), ms)
    console.log('click')
  }

  const className = join({
    space: '-mt-4 -mb-4 rounded z-10 justify-self-center',
    // background: 'bg-bgLight', // TODO: Remove
    background: 'bg-zinc-50 dark:bg-zinc-800',
    // border: 'border border-bgBase', // TODO: Remove
    border: 'border border-zinc-200 dark:border-zinc-900/95',
    // stroke: 'stroke-2 stroke-secondary', // TODO: Remove
    stroke: 'stroke-2 stroke-zinc-500',
    transition: `hover:border-zinc-400 hover:dark:border-zinc-500 cursor-pointer transition-transform ${
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
      fill="none"
      overflow="visible"
      xmlns="http://www.w3.org/2000/svg"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M11,22V8M11,8L7,12M11,8L15,12" />
      <path d="M21,9V23M21,23L25,19M21,23L17,19" />
    </svg>
  )
}
