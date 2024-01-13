import React, { useState, useEffect } from 'react'
import {
  ChevronDoubleDownIcon,
  ChevronDoubleUpIcon,
} from '@heroicons/react/outline'

type PortfolioAccordionProps = {
  header: React.ReactNode
  expandedProps: React.ReactNode
  collapsedProps: React.ReactNode
  children: React.ReactNode
  initializeExpanded: boolean
  portfolioChainId: number
  connectedChainId: number
  selectedFromChainId: number
  hasNoTokenBalance: boolean
}

export const PortfolioAccordion = ({
  header,
  expandedProps,
  collapsedProps,
  children,
  initializeExpanded = false,
  portfolioChainId,
  connectedChainId,
  selectedFromChainId,
  hasNoTokenBalance,
}: PortfolioAccordionProps) => {
  const [isHovered, setIsHovered] = useState<boolean>(false)
  const [isExpanded, setIsExpanded] = useState(false)
  const handleToggle = () => setIsExpanded((prevExpanded) => !prevExpanded)

  useEffect(() => {
    if (!hasNoTokenBalance) {
      portfolioChainId === selectedFromChainId
        ? setIsExpanded(true)
        : setIsExpanded(false)
    }
  }, [portfolioChainId, selectedFromChainId, hasNoTokenBalance])

  return (
    <div
      data-test-id="portfolio-accordion"
      className={
        isExpanded ? 'rounded-md shadow-md' : 'shadow-none'
      }
    >
      <div
        data-test-id="portfolio-accordion-header"
        className={`
          flex items-center justify-between
          border border-transparent
          pr-2 select-none
          hover:bg-zinc-100 hover:dark:bg-zinc-800
          active:opacity-70
          ${
            isExpanded
              ? 'bg-zinc-100 dark:bg-zinc-800 hover:border-b-zinc-200 hover:dark:border-b-zinc-700 rounded-t-md'
              : 'bg-transparent rounded-md'
          }
        `}
      >
        <div
          onClick={handleToggle}
          onMouseEnter={() => setIsHovered(true)}
          onMouseLeave={() => setIsHovered(false)}
          className="flex-1"
        >
          <div
            data-test-id="portfolio-accordion-clickable"
            className="flex flex-row justify-between"
          >
            {header}
            {!isExpanded && collapsedProps}
          </div>
        </div>
        {isExpanded && expandedProps}
        <AccordionIcon
          isExpanded={isExpanded}
          onClick={handleToggle}
          isHovered={isHovered}
        />
      </div>
      <div data-test-id="portfolio-accordion-contents">
        {isExpanded && <React.Fragment>{children}</React.Fragment>}
      </div>
    </div>
  )
}

export const AccordionIcon = ({
  isExpanded,
  onClick,
  isHovered,
}: {
  isExpanded: boolean
  onClick: () => void
  isHovered: boolean
}) => {

  const chevronStyles = "w-4 h-4 stroke-[3] stroke-separator stroke-zinc-500 opacity-70 group-hover:opacity-100"

  return (
    <div
      data-test-id="accordion-icon"
      onClick={onClick}
      className="group p-1 mx-2 border border-zinc-500 rounded-full cursor-pointer"
    >
      {isExpanded ? (
        <ChevronDoubleUpIcon className={chevronStyles} />
      ) : (
        <ChevronDoubleDownIcon className={chevronStyles} />
      )}
    </div>
  )
}
