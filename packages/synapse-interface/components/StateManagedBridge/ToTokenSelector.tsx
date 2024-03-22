import React from 'react'
import { useDispatch } from 'react-redux'
import { useBridgeState } from '@/slices/bridge/hooks'

export const ToTokenSelector = () => {
  const dispatch = useDispatch()

  const { toToken } = useBridgeState()

  // return (
  //   <TokenSelector
  //     dataTestId="bridge-destination-token"
  //     selectedItem={toToken}
  //     placeholder="Out"
  //     onClick={() => dispatch(setShowToTokenListOverlay(true))}
  //   />
  // )
}
