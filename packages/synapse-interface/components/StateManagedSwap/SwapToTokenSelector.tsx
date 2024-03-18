import React from 'react'
import { useDispatch } from 'react-redux'
import { setShowSwapToTokenListOverlay } from '@/slices/swapDisplaySlice'
import { useSwapState } from '@/slices/swap/hooks'
import { TokenSelector } from '../ui/BridgeCardComponents'

export const SwapToTokenSelector = () => {
  const dispatch = useDispatch()

  const { swapToToken } = useSwapState()

  return <div>TODO: Replace</div>

  // return (
  //   <TokenSelector
  //     dataTestId="bridge-destination-token"
  //     selectedItem={swapToToken}
  //     placeholder="Out"
  //     onClick={() => dispatch(setShowSwapToTokenListOverlay(true))}
  //   />
  // )
}
