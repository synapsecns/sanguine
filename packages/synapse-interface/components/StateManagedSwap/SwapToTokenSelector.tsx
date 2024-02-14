import React from 'react'
import { useDispatch } from 'react-redux'

import { DropDownArrowSvg } from '../icons/DropDownArrowSvg'
import {
  getBorderStyleForCoinHover,
  getMenuItemHoverBgForCoin,
} from '@/styles/tokens'
import { setShowSwapToTokenListOverlay } from '@/slices/swapDisplaySlice'
import { useSwapState } from '@/slices/swap/hooks'

export const SwapToTokenSelector = () => {
  const dispatch = useDispatch()

  const { swapToToken } = useSwapState()

  const buttonContent = swapToToken ? (
    <div className="flex items-center space-x-2">
      <div className="flex-none hidden md:inline-block">
        <img
          src={swapToToken?.icon?.src ?? ''}
          alt={swapToToken?.symbol ?? ''}
          className="w-6 h-6"
        />
      </div>
      <div className="text-left">
        <div className="text-lg text-primaryTextColor">
          {swapToToken?.symbol}
        </div>
      </div>
      <DropDownArrowSvg className="flex-none" />
    </div>
  ) : (
    <div className="flex items-center space-x-3">
      <div className="text-left">
        <div className="text-lg text-primaryTextColor">Out</div>
      </div>
      <DropDownArrowSvg className="flex-none" />
    </div>
  )

  return (
    <button
      data-test-id="bridge-destination-token"
      className={`
        p-md rounded-sm min-w-[80px]
        bg-slate-400/10 hover:bg-slate-400/20
        border border-white/10
        ${getMenuItemHoverBgForCoin(swapToToken?.color)}
        ${getBorderStyleForCoinHover(swapToToken?.color)}
      `}
      onClick={() => dispatch(setShowSwapToTokenListOverlay(true))}
    >
      {buttonContent}
    </button>
  )
}
