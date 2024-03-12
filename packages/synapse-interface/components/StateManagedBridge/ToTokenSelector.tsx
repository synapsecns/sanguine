import React from 'react'
import { useDispatch } from 'react-redux'
import { setShowToTokenListOverlay } from '@/slices/bridgeDisplaySlice'
import { useBridgeState } from '@/slices/bridge/hooks'
import { TokenSelector } from '../ui/BridgeCardComponents'

export const ToTokenSelector = () => {
  const dispatch = useDispatch()

  const { toToken } = useBridgeState()

  return (
    <TokenSelector
      dataTestId="bridge-destination-token"
      token={toToken}
      placeholder="Out"
      onClick={() => dispatch(setShowToTokenListOverlay(true))}
    />
  )
}
