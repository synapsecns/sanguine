import { useDispatch } from 'react-redux'

import { useBridgeState } from '@/slices/bridge/hooks'
import { setShowFromTokenListOverlay } from '@/slices/bridgeDisplaySlice'

import { TokenSelector } from '@/components/bridgeSwap/TokenSelector'

export const FromTokenSelector = () => {
  const dispatch = useDispatch()

  const { fromToken } = useBridgeState()

  return (
    <TokenSelector
      data-test-id="bridge-origin-token"
      token={fromToken}
      label="In"
      onClick={() => dispatch(setShowFromTokenListOverlay(true))}
    />

  )
}
