import { useDispatch } from 'react-redux'

import { useSwapState } from '@/slices/swap/hooks'
import { setShowSwapChainListOverlay } from '@/slices/swapDisplaySlice'

import { ChainSelector } from '@/components/bridgeSwap/ChainSelector'

export const SwapChainSelector = () => {
  const dispatch = useDispatch()
  const { swapChainId } = useSwapState()

  return (
    <ChainSelector
      data-test-id="swap-chain-list-button"
      chainId={swapChainId}
      label="On"
      onClick={() => dispatch(setShowSwapChainListOverlay(true))}
    />
  )
}
