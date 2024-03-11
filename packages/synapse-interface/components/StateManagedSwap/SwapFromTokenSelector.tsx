import React from 'react'
import { useDispatch } from 'react-redux'

import { BridgeTokenSelector } from '../ui/BridgeCard'
import { useSwapState } from '@/slices/swap/hooks'
import { setShowSwapFromTokenListOverlay } from '@/slices/swapDisplaySlice'

export const SwapFromTokenSelector = () => {
  const dispatch = useDispatch()

  const { swapFromToken } = useSwapState()

  return (
    <BridgeTokenSelector
      dataTestId="bridge-origin-token"
      token={swapFromToken}
      placeholder="In"
      onClick={() => dispatch(setShowSwapFromTokenListOverlay(true))}
    />
  )
}
