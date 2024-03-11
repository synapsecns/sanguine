import React from 'react'
import { useDispatch } from 'react-redux'

import { setShowFromTokenListOverlay } from '@/slices/bridgeDisplaySlice'
import { useBridgeState } from '@/slices/bridge/hooks'
import { DropDownArrowSvg } from '../icons/DropDownArrowSvg'
import {
  getBorderStyleForCoinHover,
  getMenuItemHoverBgForCoin,
} from '@/styles/tokens'
import { BridgeTokenSelector } from '../ui/BridgeCard'

export const FromTokenSelector = () => {
  const dispatch = useDispatch()

  const { fromToken } = useBridgeState()

  return (
    <BridgeTokenSelector
      dataTestId="bridge-origin-token"
      token={fromToken}
      placeholder="In"
      onClick={() => dispatch(setShowFromTokenListOverlay(true))}
    />
  )
}
