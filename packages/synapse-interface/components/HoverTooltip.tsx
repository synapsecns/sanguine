import React, { useState } from 'react'

export const HoverTooltip = ({
  children,
  hoverContent,
  isActive = true,
  align = 'center',
}: {
  children: React.ReactNode
  hoverContent: React.ReactNode
  isActive?: boolean
  align?: 'center' | 'start'
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
          <Tooltip isHovered={showTooltip} align={align}>
            {hoverContent}
          </Tooltip>
        ) : null}
      </div>
    )
  }
}

const Tooltip = ({
  isHovered,
  children,
  align = 'center',
}: {
  isHovered: boolean
  children: React.ReactNode
  align?: 'center' | 'start'
}) => {
  if (isHovered) {
    const positionClass =
      align === 'start' ? 'left-0' : 'left-1/2 translate-x-[-50%]'
    return (
      <div
        className={`
          absolute ${positionClass} bottom-full
          z-50 hover-content px-2 py-1 text-white mb-1
          border border-solid border-[#252537]
          bg-[#101018] rounded-md text-left text-sm whitespace-nowrap
        `}
      >
        <data>{children}</data>
      </div>
    )
  }
}
