import React from 'react'
import { useDispatch } from 'react-redux'

import { CHAINS_BY_ID } from '@/constants/chains'
import { useSwapState } from '@/slices/swap/hooks'
import { setShowSwapChainListOverlay } from '@/slices/swapDisplaySlice'
import { ChainSelector } from '../ui/BridgeCardComponents'

export const SwapChainSelector = () => {
  const dispatch = useDispatch()
  const { swapChainId } = useSwapState()
  const chain = CHAINS_BY_ID[swapChainId]

  return (
    <ChainSelector
      dataTestId="bridge-origin-chain-list-button"
      selectedItem={chain}
      label="From"
      onClick={() => dispatch(setShowSwapChainListOverlay(true))}
    />
  )
}
