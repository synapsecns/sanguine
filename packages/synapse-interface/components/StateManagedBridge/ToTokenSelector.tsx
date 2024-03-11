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

  const space = 'p-2 rounded flex-none flex items-center gap-2'
  const bgColor = `bg-[#565058]`
  const bgHover = getMenuItemHoverBgForCoin(toToken?.color)
  const borderColor = `border border-transparent`
  const borderHover = getBorderStyleForCoinHover(toToken?.color)
  const textStyle = 'text-lg text-primaryTextColor'

  return (
    <button
      data-test-id="bridge-destination-token"
      className={`${space} ${bgColor} ${bgHover} ${borderColor} ${borderHover} ${textStyle}`}
      onClick={() => dispatch(setShowToTokenListOverlay(true))}
    >
      {toToken && (
        <img
          src={toToken?.icon?.src ?? ''}
          alt={toToken?.symbol ?? ''}
          className="w-6 h-6"
        />
      )}
      {toToken?.symbol ?? 'Out'}
      <DropDownArrowSvg />
    </button>
  )
}
