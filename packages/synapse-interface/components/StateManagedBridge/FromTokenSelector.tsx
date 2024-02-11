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

  let buttonContent

  if (fromToken) {
    const src = fromToken?.icon?.src
    const symbol = fromToken?.symbol

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
      className={`
        p-md rounded-sm min-w-[80px]
        bg-slate-400/20
        ${getMenuItemHoverBgForCoin(fromToken?.color)}
        border border-transparent
        ${getBorderStyleForCoinHover(fromToken?.color)}
      `}
      onClick={() => dispatch(setShowFromTokenListOverlay(true))}
    >
      {buttonContent}
    </button>
  )
}
