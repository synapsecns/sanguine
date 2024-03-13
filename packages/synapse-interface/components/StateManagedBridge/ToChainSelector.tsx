import React from 'react'
import { useDispatch } from 'react-redux'

import { setShowToChainListOverlay } from '@/slices/bridgeDisplaySlice'
import { useBridgeState } from '@/slices/bridge/hooks'
import { CHAINS_BY_ID } from '@/constants/chains'
import { ChainSelector } from '../ui/BridgeCardComponents'

export const ToChainSelector = () => {
  const dispatch = useDispatch()
  const { toChainId } = useBridgeState()
  const toChain = CHAINS_BY_ID[toChainId]

  return (
    <ChainSelector
      dataTestId="bridge-origin-chain-list-button"
      chain={toChain}
      label="To"
      placeholder="Network"
      onClick={() => dispatch(setShowToChainListOverlay(true))}
    />
  )
}