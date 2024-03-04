import _ from 'lodash'

import { setShowSwapChainListOverlay } from '@/slices/swapDisplaySlice'
import { setSwapChainId } from '@/slices/swap/reducer'
import { useSwapState } from '@/slices/swap/hooks'

import { ChainListOverlay } from '@/components/bridgeSwap/ChainListOverlay'

export const SwapChainListOverlay = () => {
  const { swapChainId, swapFromChainIds } = useSwapState()
  const dataId = 'swap-origin-chain-list'

  return (
    <ChainListOverlay
      isOrigin={true}
      primaryLabel={"From"}
      chainId={swapChainId}
      chainIds={swapFromChainIds}
      setChainId={setSwapChainId}
      setShowOverlay={setShowSwapChainListOverlay}
      filterPausedChains={false}
    />
  )
}
