import Image from 'next/image'
import { TokenAndBalance } from '@/utils/actions/fetchPortfolioBalances'
import { HoverTooltip } from '@/components/HoverTooltip'
import { getParsedBalance } from '@/utils/getParsedBalance'
import { formatAmount } from '@/utils/formatAmount'

export const PortfolioTokenVisualizer = ({
  portfolioTokens,
  portfolioChainId,
}: {
  portfolioTokens: TokenAndBalance[]
  portfolioChainId: number
}) => {
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

  const firstTokenBalance = getParsedBalance(
    portfolioTokens?.[0]?.balance,
    portfolioTokens?.[0]?.token?.decimals[portfolioChainId]
  )
  const secondTokenBalance = getParsedBalance(
    portfolioTokens?.[1]?.balance,
    portfolioTokens?.[1]?.token?.decimals[portfolioChainId]
  )

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
          <HoverTooltip
            hoverContent={
              <div className="whitespace-nowrap">
                {formatAmount(firstTokenBalance)}{' '}
                {portfolioTokens[0]?.token.symbol}
              </div>
            }
          >
            <Image
              loading="lazy"
              className="w-6 h-6 rounded-md"
              alt={`${portfolioTokens[0].token.symbol} img`}
              src={portfolioTokens[0].token.icon}
            />
          </HoverTooltip>
        </div>
      )}

      {hasOnlyOneToken && (
        <div className="text-white whitespace-nowrap">
          {formatAmount(firstTokenBalance)} {portfolioTokens[0].token.symbol}
        </div>
      )}

      {hasTwoTokens && (
        <div>
          <HoverTooltip
            hoverContent={
              <div className="whitespace-nowrap">
                {formatAmount(secondTokenBalance)}{' '}
                {portfolioTokens[1]?.token.symbol}
              </div>
            }
          >
            <Image
              loading="lazy"
              className="w-6 h-6 rounded-md"
              alt={`${portfolioTokens[1].token.symbol} img`}
              src={portfolioTokens[1].token.icon}
            />
          </HoverTooltip>
        </div>
      )}

      {numOverTwoTokens > 0 && (
        <div className="relative inline-block">
          <HoverTooltip
            hoverContent={portfolioTokens?.map(
              (token: TokenAndBalance, key: number) => {
                if (key > 1) {
                  const tokenSymbol = token.token.symbol
                  const parsedBalance = getParsedBalance(
                    token.balance,
                    token.token.decimals[portfolioChainId]
                  )
                  return (
                    <div className="whitespace-nowrap" key={key}>
                      {formatAmount(parsedBalance)} {tokenSymbol}
                    </div>
                  )
                }
              }
            )}
          >
            <div className="text-white">+ {numOverTwoTokens}</div>
          </HoverTooltip>
        </div>
      )}
    </div>
  )
}
