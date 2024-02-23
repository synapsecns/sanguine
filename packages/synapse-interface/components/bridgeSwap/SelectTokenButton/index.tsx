import _ from 'lodash'
import { memo, useEffect, useRef, useState } from 'react'

import {
  getBorderStyleForCoin,
  getBorderStyleForCoinHover,
  getMenuItemBgForCoin,
  getMenuItemStyleForCoin,
} from '@styles/tokens'
import { Token } from '@/utils/types'

import { ButtonContent } from '@/components/bridgeSwap/SelectTokenButton/ButtonContent'

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
      <ButtonContent
        token={token}
        chainId={chainId}
        isOrigin={isOrigin}
        showAllChains={showAllChains}
        isEligible={isEligible}
        pausedChainIds={pausedChainIds}
      />
      {children}
    </button>
  )
}