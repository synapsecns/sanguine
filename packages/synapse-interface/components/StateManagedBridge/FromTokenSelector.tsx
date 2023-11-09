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
  const { fromToken, toToken } = useBridgeState()
  const BASE_BUTTON_PROPERTIES = 'p-md rounded-sm bg-[#565058] border'

  let buttonContent
  let buttonClassName

  if (fromToken) {
    const src = fromToken?.icon?.src
    const symbol = fromToken?.symbol

    buttonClassName = `
      ${BASE_BUTTON_PROPERTIES}
      border-transparent
      ${getMenuItemHoverBgForCoin(fromToken?.color)}
      ${getBorderStyleForCoinHover(fromToken?.color)}
    `

    buttonContent = (
      <div className="flex items-center space-x-2">
        <div className="flex-none hidden md:inline-block">
          <img src={src} alt={symbol} className="w-6 h-6" />
        </div>
        <div className="text-left">
          <div className="text-lg text-primaryTextColor">{symbol}</div>
        </div>
        <DropDownArrowSvg className="flex-none" />
      </div>
    )
  } else {
    buttonClassName = `
      ${BASE_BUTTON_PROPERTIES}
      border-transparent hover:border-secondary
    `

    buttonContent = (
      <div className="flex items-center space-x-3">
        <div className="text-left">
          <div className="text-lg text-primaryTextColor">In</div>
        </div>
        <DropDownArrowSvg className="flex-none" />
      </div>
    )
  }

  return (
    <button
      data-test-id="bridge-origin-token"
      className={buttonClassName}
      onClick={() => dispatch(setShowFromTokenListOverlay(true))}
    >
      {buttonContent}
    </button>
  )
}
