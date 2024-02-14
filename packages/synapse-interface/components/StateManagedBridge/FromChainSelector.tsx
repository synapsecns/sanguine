import React from 'react'
import { useDispatch } from 'react-redux'

import { setShowFromChainListOverlay } from '@/slices/bridgeDisplaySlice'
import { useBridgeState } from '@/slices/bridge/hooks'
import { ChainSelector } from '@/components/bridgeSwap/ChainSelector'

export const FromChainSelector = () => {
  const dispatch = useDispatch()
  const { fromChainId } = useBridgeState()

  return (
    <ChainSelector
      data-test-id="bridge-origin-chain-list-button"
      chainId={fromChainId}
      label="From"
      onClick={() => dispatch(setShowFromChainListOverlay(true))}
    />
  )
}
