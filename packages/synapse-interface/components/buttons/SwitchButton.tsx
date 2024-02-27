import { useState } from 'react'

export function SwitchButton({ onClick }: { onClick: () => void }) {
  const [isActive, setIsActive] = useState(false)
  const handleClick = () => {
    onClick()
    setIsActive(true)
    setTimeout(() => setIsActive(false), 200)
  }

  return (
    <div
      className={`
        flex items-center justify-center
        -mt-2 -mb-2
        group transform-gpu transition-all duration-200
        ${isActive ? 'rotate-90' : ''}
      `}
    >
      <SwitchButtonSvg onClick={handleClick} />
    </div>
  )
}

function SwitchButtonSvg({ onClick }: { onClick: () => void }) {
  return (
    <svg
      onClick={onClick}
      className="hover:opacity-80 hover:cursor-pointer"
      width="32"
      height="32"
      viewBox="0 0 32 32"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <rect x="0.5" y="0.5" width="31" height="31" rx="3.5" fill="#353038" />
      <rect x="0.5" y="0.5" width="31" height="31" rx="3.5" stroke="#252028" />
      <path
        d="M11 22V8M11 8L7 12M11 8L15 12"
        stroke="#C0BCC2"
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
      />
      <path
        d="M21 9V23M21 23L25 19M21 23L17 19"
        stroke="#C0BCC2"
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
      />
    </svg>
  )
}
