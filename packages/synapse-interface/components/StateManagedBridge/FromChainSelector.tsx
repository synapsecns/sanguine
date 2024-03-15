import React from 'react'
import { useDispatch } from 'react-redux'
import { useState } from 'react'

import { setShowFromChainListOverlay } from '@/slices/bridgeDisplaySlice'
import { useBridgeState } from '@/slices/bridge/hooks'
import { CHAINS_BY_ID } from '@/constants/chains'
import { ChainSelector } from '../ui/BridgeCardComponents'

import { FromChainListOverlay } from './FromChainListOverlay'

export const FromChainSelector = () => {
  const dispatch = useDispatch()
  const { fromChainId } = useBridgeState()
  const fromChain = CHAINS_BY_ID[fromChainId]

  const [hover, setHover] = useState(false)

  return (
    <div className="group relative">
      <ChainSelector
        dataTestId="bridge-origin-chain-list-button"
        selectedItem={fromChain}
        label="From"
        onClick={() => dispatch(setShowFromChainListOverlay(true))}
      />
      <FromChainListOverlay />
    </div>
  )
}
