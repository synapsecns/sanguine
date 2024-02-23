import _ from 'lodash'
import { memo } from 'react'

import { Token } from '@/utils/types'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { TokenBalance } from '@/components/bridgeSwap/SelectTokenButton/TokenBalance'
import { Coin } from '@/components/bridgeSwap/SelectTokenButton/Coin'

export const ButtonContent = memo(
  ({
    token,
    chainId,
    isOrigin,
    showAllChains,
    isEligible,
    pausedChainIds,
  }: {
    token: Token
    chainId: number
    isOrigin: boolean
    showAllChains: boolean
    isEligible?: boolean
    pausedChainIds?: string[]
  }) => {
    const portfolioBalances = usePortfolioBalances()

    const parsedBalance = portfolioBalances[chainId]?.find(
      (tb) => tb.token.addresses[chainId] === token.addresses[chainId]
    )?.parsedBalance

    return (
      <div data-test-id="button-content" className="flex items-center w-full">
        <img
          alt="token image"
          className="w-8 h-8 ml-2 mr-4 rounded-full"
          src={token?.icon?.src}
        />
        <Coin
          token={token}
          showAllChains={showAllChains}
          isOrigin={isOrigin}
          isEligible={isEligible}
          pausedChainIds={pausedChainIds}
        />
        {isOrigin && (
          <TokenBalance token={token} parsedBalance={parsedBalance} />
        )}
      </div>
    )
  }
)