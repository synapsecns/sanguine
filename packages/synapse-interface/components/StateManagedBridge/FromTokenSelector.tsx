import React from 'react'
import { useDispatch } from 'react-redux'

import { setShowFromTokenSlideOver } from '@/slices/bridgeDisplaySlice'
import { useBridgeState } from '@/slices/bridge/hooks'
import { DropDownArrowSvg } from './components/DropDownArrowSvg'

export const FromTokenSelector = () => {
  const dispatch = useDispatch()

  const { fromToken } = useBridgeState()

  let buttonContent

  if (fromToken) {
    const src = fromToken.icon.src
    const symbol = fromToken.symbol

    buttonContent = (
      <div className="flex items-center space-x-2">
        <div>
          <img src={src} alt={symbol} className="w-5 h-5" />
        </div>
        <div className="text-left">
          <div className="text-lg text-primaryTextColor">{symbol}</div>
        </div>
        <DropDownArrowSvg />
      </div>
    )
  } else {
    buttonContent = (
      <div className="flex items-center space-x-3">
        <div className="text-left">
          <div className="text-lg text-primaryTextColor">In</div>
        </div>
        <DropDownArrowSvg />
      </div>
    )
  }

  return (
    <button
      data-test-id="bridge-origin-token"
      className="bg-[#565058] pl-2 pr-2 pt-1 pb-1 rounded-sm min-w-[70px]"
      onClick={() => dispatch(setShowFromTokenSlideOver(true))}
    >
      {buttonContent}
    </button>
  )
}
