import { useState } from 'react'

export const Tooltip = ({
  hoverText,
  children,
}: {
  hoverText: string
  children: React.ReactNode
}) => {
  const [isHovered, setIsHovered] = useState<boolean>(false)
  return (
    <div
      data-test-id="tool-tip"
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => setIsHovered(false)}
      className="flex justify-center relative"
    >
      {isHovered && (
        <div
          className={`
          absolute -top-7 z-10 px-2 py-1 rounded-md text-sm 
          bg-[--synapse-bg-root] border border-solid border-[--synapse-border] cursor-default shadow
          `}
        >
          {hoverText}
        </div>
      )}
      {children}
    </div>
  )
}
