import React, { useState, useEffect } from 'react'
import AccordionIcon from '@icons/AccordionIcon'

type PortfolioAccordionProps = {
  header: React.ReactNode
  expandedProps: React.ReactNode
  collapsedProps: React.ReactNode
  children: React.ReactNode
  initializeExpanded: boolean
  portfolioChainId: number
  connectedChainId: number
}

export const PortfolioAccordion = ({
  header,
  expandedProps,
  collapsedProps,
  children,
  initializeExpanded = false,
  portfolioChainId,
  connectedChainId,
}: PortfolioAccordionProps) => {
  const [isExpanded, setIsExpanded] = useState(initializeExpanded)

  const handleToggle = () => {
    setIsExpanded((prevExpanded) => !prevExpanded)
  }

  const expanded = 'rotate-0'
  const collapsed = '-rotate-90'

  useEffect(() => {
    if (portfolioChainId === connectedChainId) {
      setIsExpanded(true)
    }
  }, [portfolioChainId, connectedChainId])

  return (
    <div className={isExpanded && 'pb-2'}>
      <div
        className="flex flex-row items-center py-3"
        data-test-id="portfolio-accordion"
      >
        <button onClick={handleToggle} className="mr-3">
          <AccordionIcon
            className={`
              ${isExpanded ? expanded : collapsed}
              w-4 h-4
            `}
          />
        </button>
        {header} {isExpanded ? expandedProps : collapsedProps}
      </div>
      <div>{isExpanded && <React.Fragment>{children}</React.Fragment>}</div>
    </div>
  )
}

export default PortfolioAccordion
