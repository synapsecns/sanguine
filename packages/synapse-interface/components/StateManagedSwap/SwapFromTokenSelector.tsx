import React from 'react'
import { useDispatch } from 'react-redux'

import { TokenSelector } from '../ui/BridgeCardComponents'
import { useSwapState } from '@/slices/swap/hooks'
import { setShowSwapFromTokenListOverlay } from '@/slices/swapDisplaySlice'

export const SwapFromTokenSelector = () => {
  const dispatch = useDispatch()

  const { swapFromToken } = useSwapState()

  return (
    <TokenSelector
      dataTestId="bridge-origin-token"
      selectedItem={swapFromToken}
      placeholder="In"
      onClick={() => dispatch(setShowSwapFromTokenListOverlay(true))}
    />
  )
}
