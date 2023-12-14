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
      className={`
          flex items-center justify-center
          relative group px-1
          hover:cursor-pointer
        `}
    >
      {isHovered && (
        <div
          className={`
          absolute -top-7 z-10 p-1 rounded-md text-[14px] 
          bg-[#F5F5F5] border border-[#DCDCDC]
          `}
        >
          <div>{hoverText}</div>
        </div>
      )}
      {children}
    </div>
  )
}
