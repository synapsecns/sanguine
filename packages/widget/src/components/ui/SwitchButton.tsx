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
        -mt-4 -mb-4 rounded z-10 justify-self-center
        bg-[--synapse-button-bg]
        border border-solid border-[--synapse-button-border]
        stroke-2 stroke-[--synapse-button-text]
        hover:opacity-80 cursor-pointer transition-transform
        ${isActive ? `duration-${ms} rotate-180 ease-in-out` : 'ease-out'}
      `}
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
