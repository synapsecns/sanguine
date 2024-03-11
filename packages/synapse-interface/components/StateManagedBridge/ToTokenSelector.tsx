import React from 'react'
import { useDispatch } from 'react-redux'
import { setShowToTokenListOverlay } from '@/slices/bridgeDisplaySlice'
import { useBridgeState } from '@/slices/bridge/hooks'
import { BridgeTokenSelector } from '../ui/BridgeCard'

export const ToTokenSelector = () => {
  const dispatch = useDispatch()

  const { toToken } = useBridgeState()

  return (
    <BridgeTokenSelector
      dataTestId="bridge-destination-token"
      token={toToken}
      placeholder="Out"
      onClick={() => dispatch(setShowToTokenListOverlay(true))}
    />
  )
}
