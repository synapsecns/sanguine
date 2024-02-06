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
      id="portfolio-accordion"
      className={
        isExpanded ? 'shadow-[0px_4px_4px_0px_rgba(0,0,0,0.25)]' : 'shadow-none'
      }
    >
      <div
        id="portfolio-accordion-header"
        className={`
          flex items-center justify-between border border-transparent pr-2 select-none
          hover:border-[#3D3D5C] hover:bg-[#272731]
          active:border-[#3D3D5C] active:opacity-[67%]
          ${
            isExpanded
              ? 'bg-tint rounded-t-md hover:rounded-t-md'
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
            id="portfolio-accordion-clickable"
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
      <div id="portfolio-accordion-contents" className="flex flex-col">
        {isExpanded && <React.Fragment>{children}</React.Fragment>}
      </div>
    </div>
  )
}

const AccordionIcon = ({
  isExpanded,
  onClick,
  isHovered,
}: {
  isExpanded: boolean
  onClick: () => void
  isHovered: boolean
}) => {
  return (
    <div
      id="accordion-icon"
      onClick={onClick}
      className={`
        p-1 mx-2 border border-surface rounded-full
        cursor-pointer hover:border-transparent active:border-transparent
        ${isHovered ? 'border-transparent' : 'border-surface'}
      `}
    >
      {isExpanded ? (
        <ChevronDoubleUpIcon className="w-4 h-4 stroke-[3] stroke-separator" />
      ) : (
        <ChevronDoubleDownIcon className="w-4 h-4 stroke-[3] stroke-separator" />
      )}
    </div>
  )
}
