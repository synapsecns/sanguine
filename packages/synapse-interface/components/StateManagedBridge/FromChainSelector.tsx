import React from 'react'
import { useDispatch } from 'react-redux'

import { setShowFromChainListOverlay } from '@/slices/bridgeDisplaySlice'
import { useBridgeState } from '@/slices/bridge/hooks'
import { CHAINS_BY_ID } from '@/constants/chains'
import { ChainSelector } from '../ui/BridgeCardComponents'

export const FromChainSelector = () => {
  const dispatch = useDispatch()
  const { fromChainId } = useBridgeState()
  const fromChain = CHAINS_BY_ID[fromChainId]

  return (
    <ChainSelector
      dataTestId="bridge-origin-chain-list-button"
      selectedItem={fromChain}
      label="From"
      onClick={() => dispatch(setShowFromChainListOverlay(true))}
    />
  )
}
