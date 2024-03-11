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

  const space = 'p-2 rounded flex-none flex items-center gap-2'
  const bgColor = `bg-[#565058]`
  const bgHover = getMenuItemHoverBgForCoin(swapToToken?.color)
  const borderColor = `border border-transparent`
  const borderHover = getBorderStyleForCoinHover(swapToToken?.color)
  const textStyle = 'text-lg text-primaryTextColor'

  return (
    <button
      data-test-id="bridge-destination-token"
      className={`${space} ${bgColor} ${bgHover} ${borderColor} ${borderHover} ${textStyle}`}
      onClick={() => dispatch(setShowSwapToTokenListOverlay(true))}
    >
      {swapToToken && (
        <img
          src={swapToToken?.icon?.src ?? ''}
          alt={swapToToken?.symbol ?? ''}
          className="w-6 h-6"
        />
      )}
      {swapToToken?.symbol ?? 'Out'}
      <DropDownArrowSvg />
    </button>
  )
}
