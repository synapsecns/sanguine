import React from 'react'
import { useDispatch } from 'react-redux'

import { setShowToTokenListOverlay } from '@/slices/bridgeDisplaySlice'
import { useBridgeState } from '@/slices/bridge/hooks'
import { DropDownArrowSvg } from '../icons/DropDownArrowSvg'
import {
  getBorderStyleForCoinHover,
  getMenuItemHoverBgForCoin,
} from '@/styles/tokens'

export const ToTokenSelector = () => {
  const dispatch = useDispatch()

  const { toToken } = useBridgeState()

  return (
    <button
      data-test-id="bridge-origin-token"
      className={`
        flex items-center gap-1.5
        p-md rounded-sm min-w-[80px]
        bg-zinc-100 dark:bg-zinc-600
        border border-transparent
        ${getMenuItemHoverBgForCoin(toToken?.color)}
        ${getBorderStyleForCoinHover(toToken?.color)}
      `}
      onClick={() => dispatch(setShowToTokenListOverlay(true))}
    >
      {toToken && <img
        src={toToken?.icon?.src}
        alt={toToken?.symbol}
        className="w-6 h-6 hidden xs:inline-block"
      />}
      <div className="text-lg">{toToken?.symbol || 'Out'}</div>
      <DropDownArrowSvg className="mx-0.5" />
    </button>
  )
}
