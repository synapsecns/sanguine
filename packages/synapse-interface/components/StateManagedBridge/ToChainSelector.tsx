import { useDispatch } from 'react-redux'

import { setShowToChainListOverlay } from '@/slices/bridgeDisplaySlice'
import { useBridgeState } from '@/slices/bridge/hooks'

import { ChainSelector } from '@/components/bridgeSwap/ChainSelector'

export const ToChainSelector = () => {
  const dispatch = useDispatch()
  const { toChainId } = useBridgeState()

  return (
    <ChainSelector
      chainId={toChainId}
      label="To"
      onClick={() => dispatch(setShowToChainListOverlay(true))}
    />
  )
}
