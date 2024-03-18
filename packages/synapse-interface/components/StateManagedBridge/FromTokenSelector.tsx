import React from 'react'
import { useDispatch } from 'react-redux'

import { setShowFromTokenListOverlay } from '@/slices/bridgeDisplaySlice'
import { useBridgeState } from '@/slices/bridge/hooks'
import { TokenSelector } from '../ui/BridgeCardComponents'

export const FromTokenSelector = () => {
  const dispatch = useDispatch()

  const { fromToken } = useBridgeState()

  // return (
  //   <TokenSelector
  //     dataTestId="bridge-origin-token"
  //     selectedItem={fromToken}
  //     placeholder="In"
  //     onClick={() => dispatch(setShowFromTokenListOverlay(true))}
  //   />
  // )
}
