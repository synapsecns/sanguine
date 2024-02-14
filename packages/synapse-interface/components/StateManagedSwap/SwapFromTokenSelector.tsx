import { useDispatch } from 'react-redux'

import { useSwapState } from '@/slices/swap/hooks'
import { setShowSwapFromTokenListOverlay } from '@/slices/swapDisplaySlice'

import { TokenSelector } from '@/components/bridgeSwap/TokenSelector'

export const SwapFromTokenSelector = () => {
  const dispatch = useDispatch()

  const { swapFromToken } = useSwapState()

  return (
    <TokenSelector
      data-test-id="swap-origin-token"
      token={swapFromToken}
      label="In"
      onClick={() => dispatch(setShowSwapFromTokenListOverlay(true))}
    />
  )
}
