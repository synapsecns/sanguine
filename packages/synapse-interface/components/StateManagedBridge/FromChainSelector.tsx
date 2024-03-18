import React, { useState } from 'react'
import { useDispatch } from 'react-redux'

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
    <div
      className="relative"
      onMouseEnter={() => setHover(true)}
      onMouseLeave={() => setHover(false)}
      onMouseDown={(e) => e.stopPropagation()}
      onClick={() => setHover(!hover)}
    >
      <ChainSelector
        dataTestId="bridge-origin-chain-list-button"
        selectedItem={fromChain}
        label="From"
        onClick={() => setHover(true)}
        // onClick={() => dispatch(setShowFromChainListOverlay(true))}
      />
      {hover && <FromChainListOverlay />}
    </div>
  )
}
