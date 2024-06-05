import React, { useState } from 'react'

export const HoverTooltip = ({ children, hoverContent }) => {
  const [showTooltip, setShowTooltip] = useState(false)

  const activateTooltip = () => setShowTooltip(true)
  const hideTooltip = () => setShowTooltip(false)

  return (
    <div
      onMouseEnter={activateTooltip}
      onMouseLeave={hideTooltip}
      className="relative"
    >
      {children}
      <Tooltip isHovered={showTooltip}>{hoverContent}</Tooltip>
    </div>
  )
}

const Tooltip = ({
  isHovered,
  children,
}: {
  isHovered: boolean
  children: React.ReactNode
}) => {
  if (isHovered) {
    return (
      <div
        className={`
          absolute left-1/2 bottom-full translate-x-[-50%]
          z-50 hover-content px-2 py-1 text-white mb-1
          border border-solid border-[#252537]
          bg-[#101018] rounded-md text-left text-sm
        `}
      >
        <data>{children}</data>
      </div>
    )
  }
}
