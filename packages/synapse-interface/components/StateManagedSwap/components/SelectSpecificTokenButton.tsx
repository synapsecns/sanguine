import {
  getBorderStyleForCoin,
  getBorderStyleForCoinHover,
  getMenuItemBgForCoin,
  getMenuItemStyleForCoin,
} from '@styles/tokens'
import { memo, useEffect, useRef } from 'react'
import { Token } from '@/utils/types'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { useSwapState } from '@/slices/swap/hooks'

import { AvailableChains } from '@/components/bridgeSwap/TokenButton/AvailableChains'
import { TokenBalance } from '@/components/bridgeSwap/TokenButton/TokenBalance'

export const SelectSpecificTokenButton = ({
  showAllChains,
  isOrigin,
  token,
  active,
  selectedToken,
  onClick,
  alternateBackground = false,
}: {
  showAllChains?: boolean
  isOrigin: boolean
  token: Token
  active: boolean
  selectedToken: Token
  onClick: () => void
  alternateBackground?: boolean
}) => {
  const ref = useRef<any>(null)
  const isCurrentlySelected = selectedToken?.routeSymbol === token?.routeSymbol
  const { swapChainId, swapFromToken, swapToToken } = useSwapState()

  useEffect(() => {
    if (active) {
      ref?.current?.focus()
    }
  }, [active])

  const chainId = swapChainId

  let bgClassName

  const classNameForMenuItemStyle = getMenuItemStyleForCoin(token?.color)

  if (isCurrentlySelected) {
    bgClassName = `${getMenuItemBgForCoin(
      token?.color
    )} ${getBorderStyleForCoin(token?.color)}`
  } else {
    bgClassName = getBorderStyleForCoinHover(token?.color)
  }

  return (
    <button
      ref={ref}
      tabIndex={active ? 1 : 0}
      onClick={onClick}
      className={`
        flex items-center
        transition-all duration-75
        w-full
        px-2 py-1
        cursor-pointer
        rounded-md
        border border-slate-400/10
        mb-1
        ${alternateBackground ? '' : !isCurrentlySelected && 'bg-slate-400/10' }
        ${bgClassName}
        ${classNameForMenuItemStyle}
      `}
    >
      <ButtonContent
        token={token}
        chainId={chainId}
        isOrigin={isOrigin}
        showAllChains={showAllChains}
      />
    </button>
  )
}

const ButtonContent = memo(
  ({
    token,
    chainId,
    isOrigin,
    showAllChains,
  }: {
    token: Token
    chainId: number
    isOrigin: boolean
    showAllChains: boolean
  }) => {
    const portfolioBalances = usePortfolioBalances()

    const parsedBalance = portfolioBalances[chainId]?.find(
      (tb) => tb.token.addresses[chainId] === token.addresses[chainId]
    )?.parsedBalance

    return (
      <div className="flex items-center w-full">
        <img
          alt="token image"
          className="w-8 h-8 ml-2 mr-4 rounded-full"
          src={token?.icon?.src}
        />
        <Coin token={token} showAllChains={showAllChains} />
        {isOrigin && (
          <TokenBalance
            token={token}
            parsedBalance={parsedBalance}
          />
        )}
      </div>
    )
  }
)

const Coin = ({ token, showAllChains }: { token; showAllChains: boolean }) => {
  return (
    <div className="flex-col text-left">
      <div className="text-lg text-primaryTextColor">{token?.symbol}</div>
      <div className="flex items-center space-x-2 text-xs text-secondaryTextColor">
        <div>{token?.name}</div>
        {showAllChains &&
          <AvailableChains
            token={token}
          />
        }
      </div>
    </div>
  )
}

