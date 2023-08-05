import React from 'react'
import { useDispatch } from 'react-redux'

import { setShowToTokenListOverlay } from '@/slices/bridgeDisplaySlice'
import { useBridgeState } from '@/slices/bridge/hooks'
import { DropDownArrowSvg } from './components/DropDownArrowSvg'
import {
  getBorderStyleForCoinHover,
  getMenuItemHoverBgForCoin,
} from '@/styles/tokens'

export const ToTokenSelector = () => {
  const dispatch = useDispatch()

  const { toToken } = useBridgeState()

  let buttonContent

  if (toToken) {
    const src = toToken.icon.src
    const symbol = toToken.symbol

    buttonContent = (
      <div className="flex items-center space-x-2">
        <div className="flex-none hidden md:inline-block">
          <img src={src} alt={symbol} className="w-5 h-5" />
        </div>
        <div className="text-left">
          <div className="text-lg text-primaryTextColor">{symbol}</div>
        </div>
        <DropDownArrowSvg className="flex-none" />
      </div>
    )
  } else {
    buttonContent = (
      <div className="flex items-center space-x-3">
        <div className="text-left">
          <div className="text-lg text-primaryTextColor">Out</div>
        </div>
        <DropDownArrowSvg className="flex-none" />
      </div>
    )
  }

  return (
    <button
      data-test-id="bridge-destination-token"
      className={`
        pl-2 pr-2 pt-1 pb-1 rounded-md min-w-[80px]
        bg-[#565058]
        ${getMenuItemHoverBgForCoin(toToken?.color)}
        border border-transparent
        ${getBorderStyleForCoinHover(toToken?.color)}
      `}
      onClick={() => dispatch(setShowToTokenListOverlay(true))}
    >
      {buttonContent}
    </button>
  )
}
