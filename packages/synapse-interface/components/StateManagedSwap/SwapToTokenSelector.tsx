import React from 'react'
import { useDispatch } from 'react-redux'
import { setShowSwapToTokenListOverlay } from '@/slices/swapDisplaySlice'
import { useSwapState } from '@/slices/swap/hooks'
import { BridgeTokenSelector } from '../ui/BridgeCard'

export const SwapToTokenSelector = () => {
  const dispatch = useDispatch()

  const { swapToToken } = useSwapState()

  return (
    <BridgeTokenSelector
      dataTestId="bridge-destination-token"
      token={swapToToken}
      placeholder="Out"
      onClick={() => dispatch(setShowSwapToTokenListOverlay(true))}
    />
  )
}
