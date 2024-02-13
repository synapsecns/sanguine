import { useState } from 'react'
import Image from 'next/image'
import { TokenAndBalance } from '@/utils/actions/fetchPortfolioBalances'

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
          <div className="relative">
            <HoverContent isHovered={isT1Hovered}>
              <div className="whitespace-nowrap">
                {portfolioTokens[0]?.parsedBalance}{' '}
                {portfolioTokens[0]?.token.symbol}
              </div>
            </HoverContent>
          </div>
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
          <div className="relative">
            <HoverContent isHovered={isT2Hovered}>
              <div className="whitespace-nowrap">
                {portfolioTokens[1]?.parsedBalance}{' '}
                {portfolioTokens[1]?.token.symbol}
              </div>
            </HoverContent>
          </div>
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
      <div className="relative inline-block">
        <HoverContent isHovered={isT3Hovered}>
          {portfolioTokens?.map((token: TokenAndBalance, key: number) => {
            if (key > 1) {
              const tokenSymbol = token.token.symbol
              const balance = token.parsedBalance
              return (
                <div className="whitespace-nowrap" key={key}>
                  {balance} {tokenSymbol}
                </div>
              )
            }
          })}
        </HoverContent>
      </div>
    </div>
  )
}

export const HoverContent = ({
  isHovered,
  children,
}: {
  isHovered: boolean
  children: React.ReactNode
}) => {
  if (isHovered) {
    return (
      <div
        className={`
          absolute z-50 hover-content py-2 px-3 text-white
          border border-white/20 bg-bgBase/10 backdrop-blur-xl
          rounded-md text-left min-w-[200px]
        `}
      >
        {children}
      </div>
    )
  }
}
