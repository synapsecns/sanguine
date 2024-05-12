import React, { useState } from 'react'

export const HoverTooltip = ({
  children,
  hoverContent,
  isActive = true,
}: {
  children: React.ReactNode
  hoverContent: React.ReactNode
  isActive?: boolean
}) => {
  const [showTooltip, setShowTooltip] = useState(false)

  const activateTooltip = () => setShowTooltip(true)
  const hideTooltip = () => setShowTooltip(false)

  if (!isActive) {
    return <div>{children}</div>
  } else {
    return (
      <div
        onMouseEnter={activateTooltip}
        onMouseLeave={hideTooltip}
        className="relative w-fit"
      >
        {children}
        {hoverContent ? (
          <Tooltip isHovered={showTooltip}>{hoverContent}</Tooltip>
        ) : null}
      </div>
    )
  }
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
