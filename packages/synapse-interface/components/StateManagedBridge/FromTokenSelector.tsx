import React from 'react'
import { useDispatch } from 'react-redux'

import { setShowFromTokenListOverlay } from '@/slices/bridgeDisplaySlice'
import { useBridgeState } from '@/slices/bridge/hooks'
import { DropDownArrowSvg } from '../icons/DropDownArrowSvg'
import {
  getBorderStyleForCoinHover,
  getMenuItemHoverBgForCoin,
} from '@/styles/tokens'

export const FromTokenSelector = () => {
  const dispatch = useDispatch()

  const { fromToken } = useBridgeState()

  return (
    <button
      data-test-id="bridge-origin-token"
      className={`
        flex items-center gap-1.5
        p-md rounded-sm min-w-[80px]
        bg-zinc-100 dark:bg-zinc-600
        border border-transparent
        ${getMenuItemHoverBgForCoin(fromToken?.color)}
        ${getBorderStyleForCoinHover(fromToken?.color)}
      `}
      onClick={() => dispatch(setShowFromTokenListOverlay(true))}
    >
      {fromToken && <img
        src={fromToken?.icon?.src}
        alt={fromToken?.symbol}
        className="w-6 h-6 hidden xs:inline-block"
      />}
      <div className="text-lg">{fromToken?.symbol || 'In'}</div>
      <DropDownArrowSvg className="mx-0.5" />
    </button>
  )
}
