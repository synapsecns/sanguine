import React, { useState, useEffect } from 'react'
import { ChevronUpIcon } from '@heroicons/react/outline'
import Card from '@tw/Card'

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
      setIsExpanded(portfolioChainId === selectedFromChainId)
    }
  }, [portfolioChainId, selectedFromChainId, hasNoTokenBalance])

  return (
    <Card
      id="portfolio-accordion"
      className={
        `${isExpanded ? 'shadow-[0px_4px_4px_0px_rgba(0,0,0,0.25)]' : 'shadow-none'} !p-0 !from-transparent !to-transparent rounded-lg`
      }
    >
      <div
        id="portfolio-accordion-header"
        className={`
          group
          flex items-center justify-between border border-transparent pr-2 select-none
           hover:bg-bgBase/20
          active:border-[#3D3D5C] active:opacity-[67%]
          ${
            isExpanded
              ? 'bg-bgBase/10 rounded-t-lg hover:rounded-t-lg'
              : 'bg-transparent rounded-lg'
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
    </Card>
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
        p-1 mx-2 border border-transparent rounded-full
        cursor-pointer
        hover:bg-bgBase/10
        hover:border-white/10 active:border-white/80 transition-all
      `}
    >
      <ChevronUpIcon
          className={`
            w-4 h-4 stroke-[3] stroke-white/20 group-hover:stroke-white/70 transition-all
            ${isExpanded ? 'rotate-180' : 'rotate-0'}
          `}
      />
    </div>
  )
}
