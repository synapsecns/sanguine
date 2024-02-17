import { useState } from 'react'
import Image from 'next/image'
import { TokenAndBalance } from '@/utils/actions/fetchPortfolioBalances'
import { HoverTokenAndBalance } from './HoverTokenAndBalance'

export const PortfolioTokenVisualizer = ({
  portfolioTokens,
}: {
  portfolioTokens: TokenAndBalance[]
}) => {
  const [isT1Hovered, setIsT1Hovered] = useState<boolean>(false)
  const [isT2Hovered, setIsT2Hovered] = useState<boolean>(false)
  const [isT3Hovered, setIsT3Hovered] = useState<boolean>(false)

  const hasNoTokens: boolean =
    !portfolioTokens || (portfolioTokens && portfolioTokens.length === 0)
  const hasOneToken: boolean = portfolioTokens && portfolioTokens.length > 0
  const hasTwoTokens: boolean = portfolioTokens && portfolioTokens.length > 1
  const numOverTwoTokens: number =
    portfolioTokens && portfolioTokens.length - 2 > 0
      ? portfolioTokens.length - 2
      : 0
  const hasOnlyOneToken: boolean =
    portfolioTokens && portfolioTokens.length === 1

  if (hasNoTokens) {
    return (
      <div
        id="portfolio-token-visualizer"
        className="flex flex-row items-center mr-4 cursor-pointer hover-trigger text-secondary"
      >
        -
      </div>
    )
  }
  return (
    <div
      id="portfolio-token-visualizer"
      className="flex flex-row items-center space-x-2 cursor-pointer hover-trigger"
    >
      {hasOneToken && (
        <div>
          <Image
            loading="lazy"
            className="w-6 h-6 rounded-md"
            alt={`${portfolioTokens[0].token.symbol} img`}
            src={portfolioTokens[0].token.icon}
            onMouseEnter={() => setIsT1Hovered(true)}
            onMouseLeave={() => setIsT1Hovered(false)}
          />
          <HoverTokenAndBalance
            isHovered={isT1Hovered}
            tokens={[portfolioTokens[0]]}
          />
        </div>
      )}
      {hasOnlyOneToken && (
        <div className="text-white whitespace-nowrap">
          {portfolioTokens[0].parsedBalance} {portfolioTokens[0].token.symbol}
        </div>
      )}
      {hasTwoTokens && (
        <div>
          <Image
            loading="lazy"
            className="w-6 h-6 rounded-md"
            alt={`${portfolioTokens[1].token.symbol} img`}
            src={portfolioTokens[1].token.icon}
            onMouseEnter={() => setIsT2Hovered(true)}
            onMouseLeave={() => setIsT2Hovered(false)}
          />
          <HoverTokenAndBalance
            isHovered={isT2Hovered}
            tokens={[portfolioTokens[1]]}
          />
        </div>
      )}
      {numOverTwoTokens > 0 && (
        <div
          className="text-white"
          onMouseEnter={() => setIsT3Hovered(true)}
          onMouseLeave={() => setIsT3Hovered(false)}
        >
          + {numOverTwoTokens}
        </div>
      )}
      <HoverTokenAndBalance
        isHovered={isT3Hovered}
        tokens={portfolioTokens}
        startFrom={2}
      />
    </div>
  )
}

