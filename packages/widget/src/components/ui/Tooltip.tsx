import { useState } from 'react'

export const Tooltip = ({
  hoverText,
  children,
  positionStyles,
}: {
  hoverText: string
  children: React.ReactNode
  positionStyles?: string
}) => {
  const [isHovered, setIsHovered] = useState<boolean>(false)
  return (
    <div
      data-test-id="tool-tip"
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => setIsHovered(false)}
      className="relative flex justify-center"
    >
      {isHovered && hoverText && (
        <div
          className={`absolute w-max z-10 px-2 py-1 rounded-md text-sm border border-solid border-[--synapse-border] cursor-default shadow ${positionStyles}`}
          style={{ background: 'var(--synapse-select-bg)' }}
        >
          {hoverText}
        </div>
      )}
      {children}
    </div>
  )
}
