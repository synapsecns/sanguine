import { useDispatch } from 'react-redux'

import { useSwapState } from '@/slices/swap/hooks'
import { setShowSwapToTokenListOverlay } from '@/slices/swapDisplaySlice'

import { GenericOutputContainer } from '@/components/bridgeSwap/GenericOutputContainer'
import { TokenSelector } from '@/components/bridgeSwap/TokenSelector'


export const SwapOutputContainer = () => {
  const dispatch = useDispatch()
  const { swapQuote, isLoading, swapToToken } = useSwapState()

  return (
    <GenericOutputContainer
      tokenSelector={
        <TokenSelector
          data-test-id="swap-destination-token"
          token={swapToToken}
          label="Out"
          onClick={() => dispatch(setShowSwapToTokenListOverlay(true))}
        />
      }
      isLoading={isLoading}
      quote={swapQuote}
    />
  )
}
