import { useDispatch } from 'react-redux'

import { useSwapState } from '@/slices/swap/hooks'
import { setShowSwapToTokenListOverlay } from '@/slices/swapDisplaySlice'

import { TokenSelector } from '@/components/bridgeSwap/TokenSelector'

export const SwapToTokenSelector = () => {
  const dispatch = useDispatch()

  const { swapToToken } = useSwapState()


  return (
    <TokenSelector
      data-test-id="swap-destination-token"
      token={swapToToken}
      label="Out"
      onClick={() => dispatch(setShowSwapToTokenListOverlay(true))}
    />
  )
}
