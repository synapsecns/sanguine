import React, { useState } from 'react'
import { PlayIcon } from '@heroicons/react/outline'

type AccordionProps = {
  header: React.ReactNode
  children: React.ReactNode
}

export const Accordion = ({ header, children }: AccordionProps) => {
  const [isExpanded, setIsExpanded] = useState(false)

  const handleToggle = () => {
    setIsExpanded((prevExpanded) => !prevExpanded)
  }

  return (
    <div>
      <div className="flex flex-row" data-test-id="Accordion">
        <button onClick={handleToggle} className={`w-2`}>
          <PlayIcon className="w-4 h-4" />
        </button>
        {header}
      </div>
      <div>{isExpanded && <React.Fragment>{children}</React.Fragment>}</div>
    </div>
  )
}

export default Accordion
