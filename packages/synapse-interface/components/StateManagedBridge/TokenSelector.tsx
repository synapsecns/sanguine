import React from 'react'
import { useDispatch } from 'react-redux'

import {
  setShowFromTokenListOverlay,
  setShowToTokenListOverlay,
} from '@/slices/bridgeDisplaySlice'
import { useBridgeState } from '@/slices/bridge/hooks'
import { DropDownArrowSvg } from '../icons/DropDownArrowSvg'
import {
  getBorderStyleForCoinHover,
  getMenuItemHoverBgForCoin,
} from '@/styles/tokens'

export const TokenSelector = ({ side }: { side: string }) => {
  const dispatch = useDispatch()
  const { fromToken, toToken } = useBridgeState()

  const token = side === 'from' ? fromToken : toToken

  return (
    <button
      data-test-id="bridge-origin-token"
      className={`
        flex items-center gap-1.5
        p-2 rounded-sm
        bg-zinc-100 dark:bg-zinc-600
        border border-zinc-200 dark:border-transparent
        text-lg
        ${getMenuItemHoverBgForCoin(token?.color)}
        ${getBorderStyleForCoinHover(token?.color)}
      `}
      onClick={side === 'from'
        ? () => dispatch(setShowFromTokenListOverlay(true))
        : () => dispatch(setShowToTokenListOverlay(true))
      }
    >
      {token && <img
        src={token?.icon?.src}
        alt={token?.symbol}
        className="w-6 h-6 hidden xs:inline-block"
      />}
      {token?.symbol || (side === 'from' ? 'Out' : 'In')}
      <DropDownArrowSvg className="w-8" />
    </button>
  )
}
