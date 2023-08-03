import React, { useState, useEffect } from 'react'
import { WarningMessage } from '../Warning'
import { TWITTER_URL, DISCORD_URL } from '@/constants/urls'

type PortfolioAccordionProps = {
  header: React.ReactNode
  expandedProps: React.ReactNode
  collapsedProps: React.ReactNode
  children: React.ReactNode
  initializeExpanded: boolean
  portfolioChainId: number
  connectedChainId: number
  selectedFromChainId: number
  isUnsupportedChain: boolean
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
  isUnsupportedChain,
}: PortfolioAccordionProps) => {
  const [isExpanded, setIsExpanded] = useState(initializeExpanded)
  const handleToggle = () => setIsExpanded((prevExpanded) => !prevExpanded)

  useEffect(() => {
    if (portfolioChainId === selectedFromChainId) {
      setIsExpanded(true)
    } else {
      setIsExpanded(false)
    }
  }, [portfolioChainId, selectedFromChainId])

  return (
    <div
      className={
        isExpanded ? 'border-b border-t border-solid border-[#3D3D5C]' : ''
      }
    >
      <div
        className={`
        flex flex-row
        items-center justify-between
        border border-transparent
        hover:border-[#3D3D5C] hover:bg-[#272731]
        active:border-[#3D3D5C] active:opacity-[67%]
        `}
        data-test-id="portfolio-accordion"
      >
        <button onClick={handleToggle} className="flex-1 mr-3">
          <div className="flex flex-row justify-between">
            {header}
            {!isExpanded && collapsedProps}
          </div>
        </button>
        {isExpanded && expandedProps}
      </div>
      <div className="flex flex-col">
        {isExpanded && isUnsupportedChain && (
          <WarningMessage
            message={
              <p>
                This chain is not yet supported. New chain or token support can
                be discussed on{' '}
                <a target="_blank" className="underline" href={TWITTER_URL}>
                  Twitter
                </a>{' '}
                or{' '}
                <a target="_blank" className="underline" href={DISCORD_URL}>
                  Discord
                </a>{' '}
                .
              </p>
            }
          />
        )}
        {isExpanded && <React.Fragment>{children}</React.Fragment>}
      </div>
    </div>
  )
}

export default PortfolioAccordion
