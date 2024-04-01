import React, { useState } from 'react'
import { HoverContent } from './PortfolioTokenVisualizer'

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

      <HoverContent isHovered={showTooltip}>{hoverContent}</HoverContent>
    </div>
  )
}
