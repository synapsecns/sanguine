import React, { useState } from 'react'

type AccordionProps = {
  children: React.ReactNode
}

export const Accordion = ({ children }: AccordionProps) => {
  const [isExpanded, setIsExpanded] = useState(false)

  const handleToggle = () => {
    setIsExpanded((prevExpanded) => !prevExpanded)
  }

  return <div>{isExpanded && <React.Fragment>{children}</React.Fragment>}</div>
}

export default Accordion
