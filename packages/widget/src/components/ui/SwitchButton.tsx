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
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        d="M12 6V18M12 18L16 14M12 18L8 14"
        stroke-linecap="round"
        stroke-linejoin="round"
      />
    </svg>
  )
}
