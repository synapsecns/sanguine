import React from 'react'
import { useDispatch } from 'react-redux'

import { useSwapState } from '@/slices/swap/hooks'
import { setShowSwapFromTokenListOverlay } from '@/slices/swapDisplaySlice'

export const SwapFromTokenSelector = () => {
  const dispatch = useDispatch()

  const { swapFromToken } = useSwapState()

  return <div>TODO: Replace</div>

  // return (
  //   <TokenSelector
  //     dataTestId="bridge-origin-token"
  //     selectedItem={swapFromToken}
  //     placeholder="In"
  //     onClick={() => dispatch(setShowSwapFromTokenListOverlay(true))}
  //   />
  // )
}
