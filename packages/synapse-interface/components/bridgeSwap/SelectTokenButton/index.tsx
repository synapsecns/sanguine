import _ from 'lodash'
import { memo, useEffect, useRef, useState } from 'react'

import {
  getBorderStyleForCoin,
  getBorderStyleForCoinHover,
  getMenuItemBgForCoin,
  getMenuItemStyleForCoin,
} from '@styles/tokens'
import type { Token } from '@/utils/types'

import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { TokenBalance } from '@/components/bridgeSwap/SelectTokenButton/TokenBalance'
import { Coin } from '@/components/bridgeSwap/SelectTokenButton/Coin'


export const SelectTokenButton = ({
  chainId,
  showAllChains,
  isOrigin,
  token,
  active,
  selectedToken,
  onClick,
  alternateBackground = false,
  isEligible,
  pausedChainIds,
  children
}: {
  chainId: number
  showAllChains?: boolean
  isOrigin: boolean
  token: Token
  active: boolean
  selectedToken: Token
  onClick: () => void
  alternateBackground?: boolean
  isEligible?: boolean
  pausedChainIds?: string[]
  children?: React.ReactNode
}) => {
  const ref = useRef<any>(null)
  const isCurrentlySelected = selectedToken?.routeSymbol === token?.routeSymbol

  useEffect(() => {
    if (active) {
      ref?.current?.focus()
    }
  }, [active])


  let bgClassName

  const classNameForMenuItemStyle = getMenuItemStyleForCoin(token?.color)

  if (isCurrentlySelected) {
    bgClassName = `${getMenuItemBgForCoin(
      token?.color
    )} ${getBorderStyleForCoin(token?.color)}`
  } else {
    bgClassName = getBorderStyleForCoinHover(token?.color)
  }

  const portfolioBalances = usePortfolioBalances()

  const parsedBalance = portfolioBalances[chainId]?.find(
    (tb) => tb.token.addresses[chainId] === token.addresses[chainId]
  )?.parsedBalance

  return (
    <button
      data-test-id="select-specific-token-button"
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
      {children}
    </button>
  )
}