import React, { useState } from 'react'
import AccordionIcon from '../icons/AccordionIcon'

type PortfolioAccordionProps = {
  header: React.ReactNode
  children: React.ReactNode
}

export const PortfolioAccordion = ({
  header,
  children,
}: PortfolioAccordionProps) => {
  const [isExpanded, setIsExpanded] = useState(false)

  const handleToggle = () => {
    setIsExpanded((prevExpanded) => !prevExpanded)
  }

  return (
    <div>
      <div className="flex flex-row" data-test-id="Accordion">
        <button onClick={handleToggle} className="mb-4 mr-2">
          <AccordionIcon className="w-4 h-4" />
        </button>
        {header}
      </div>
      <div>{isExpanded && <React.Fragment>{children}</React.Fragment>}</div>
    </div>
  )
}

export default PortfolioAccordion
