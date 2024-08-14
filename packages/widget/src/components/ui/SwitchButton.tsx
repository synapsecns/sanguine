import { useState } from 'react'

const ms = 300

export const SwitchButton = ({ onClick }: { onClick: () => void }) => {
  const [isActive, setIsActive] = useState(false)

  const handleClick = () => {
    onClick()
    setIsActive(true)
    setTimeout(() => setIsActive(false), ms)
  }

  return (
    <svg
      onClick={handleClick}
      className={`
        -mt-3 -mb-3 rounded z-10 justify-self-center
        bg-[--synapse-surface]
        border-2 border-solid border-[--synapse-root]
        stroke-2 stroke-[--synapse-secondary]
        hover:stroke-[--synapse-text] cursor-pointer transition-transform
        ${isActive ? `duration-${ms} rotate-180 ease-in-out` : 'ease-out'}
      `}
      width="28"
      height="28"
      viewBox="0 0 28 28"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        d="M14 8V20M14 20L18 16M14 20L10 16"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
      />
    </svg>
  )
}
