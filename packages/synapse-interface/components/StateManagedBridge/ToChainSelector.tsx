import React, { useState } from 'react'
import { useDispatch } from 'react-redux'

import { setShowToChainListOverlay } from '@/slices/bridgeDisplaySlice'
import { useBridgeState } from '@/slices/bridge/hooks'
import { CHAINS_BY_ID } from '@/constants/chains'
import { ChainSelector } from '../ui/BridgeCardComponents'
import { ToChainListOverlay } from './ToChainListOverlay'

export const ToChainSelector = () => {
  const dispatch = useDispatch()
  const { toChainId } = useBridgeState()
  const toChain = CHAINS_BY_ID[toChainId]

  const [hover, setHover] = useState(false)

  return (
    <div
      className="relative"
      onMouseEnter={() => setHover(true)}
      onMouseLeave={() => setHover(false)}
      onMouseDown={(e) => e.stopPropagation()}
      onMouseUp={() => setHover(false)}
    >
      <ChainSelector
        dataTestId="bridge-origin-chain-list-button"
        selectedItem={toChain}
        label="To"
        onClick={() => setHover(true)}
        // onClick={() => dispatch(setShowToChainListOverlay(true))}
      />
      {hover && <ToChainListOverlay />}
    </div>
  )
}
