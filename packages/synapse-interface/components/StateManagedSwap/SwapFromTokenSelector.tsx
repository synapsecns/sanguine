import React from 'react'
import { useDispatch } from 'react-redux'

import { DropDownArrowSvg } from '../icons/DropDownArrowSvg'
import {
  getBorderStyleForCoinHover,
  getMenuItemHoverBgForCoin,
} from '@/styles/tokens'
import { useSwapState } from '@/slices/swap/hooks'
import { setShowSwapFromTokenListOverlay } from '@/slices/swapDisplaySlice'

export const SwapFromTokenSelector = () => {
  const dispatch = useDispatch()

  const { swapFromToken } = useSwapState()

  const buttonContent = swapFromToken ? (
    <div className="flex items-center space-x-2">
      <div className="flex-none hidden md:inline-block">
        <img
          src={swapFromToken?.icon?.src}
          alt={`Icon for ${swapFromToken?.symbol} token`}
          className="w-6 h-6"
        />
      </div>
      <div className="text-left">
        <div className="text-lg text-primaryTextColor">
          {swapFromToken?.symbol}
        </div>
      </div>
      <DropDownArrowSvg className="flex-none" />
    </div>
  ) : (
    <div className="flex items-center space-x-3">
      <div className="text-left">
        <div className="text-lg text-primaryTextColor">In</div>
      </div>
      <DropDownArrowSvg className="flex-none" />
    </div>
  )

  return (
    <button
      data-test-id="bridge-origin-token"
      className={`
        p-md rounded-sm min-w-[80px]
        bg-slate-400/10 hover:bg-slate-400/20
        border border-white/10
        ${getMenuItemHoverBgForCoin(swapFromToken?.color)}
        ${getBorderStyleForCoinHover(swapFromToken?.color)}
      `}
      onClick={() => dispatch(setShowSwapFromTokenListOverlay(true))}
    >
      {buttonContent}
    </button>
  )
}
