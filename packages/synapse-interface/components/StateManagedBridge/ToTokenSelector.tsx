import { useDispatch } from 'react-redux'

import { useBridgeState } from '@/slices/bridge/hooks'
import { setShowToTokenListOverlay } from '@/slices/bridgeDisplaySlice'

import { TokenSelector } from '@/components/bridgeSwap/TokenSelector'

export const ToTokenSelector = () => {
  const dispatch = useDispatch()

  const { toToken } = useBridgeState()

  return (
    <TokenSelector
      data-test-id="bridge-destination-token"
      token={toToken}
      label="Out"
      onClick={() => dispatch(setShowToTokenListOverlay(true))}
    />
  )
}
