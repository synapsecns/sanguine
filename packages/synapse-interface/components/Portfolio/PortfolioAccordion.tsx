import React, { useState } from 'react'
import AccordionIcon from '../icons/AccordionIcon'

type PortfolioAccordionProps = {
  expandedHeader: React.ReactNode
  collapsedHeader: React.ReactNode
  children: React.ReactNode
  initializeExpanded: boolean
}

export const PortfolioAccordion = ({
  expandedHeader,
  collapsedHeader,
  children,
  initializeExpanded = false,
}: PortfolioAccordionProps) => {
  const [isExpanded, setIsExpanded] = useState(initializeExpanded)

  const handleToggle = () => {
    setIsExpanded((prevExpanded) => !prevExpanded)
  }

  const expanded = 'rotate-0'
  const collapsed = '-rotate-90'

  return (
    <div>
      <div
        className="flex flex-row items-center"
        data-test-id="portfolio-accordion"
      >
        <button onClick={handleToggle} className="mb-3 mr-3">
          <AccordionIcon
            className={`
              ${isExpanded ? expanded : collapsed}
              w-4 h-4
            `}
          />
        </button>
        {isExpanded ? expandedHeader : collapsedHeader}
      </div>
      <div>{isExpanded && <React.Fragment>{children}</React.Fragment>}</div>
    </div>
  )
}

export default PortfolioAccordion
