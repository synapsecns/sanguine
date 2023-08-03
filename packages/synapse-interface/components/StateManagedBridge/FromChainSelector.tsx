import React from 'react'
import { useDispatch } from 'react-redux'

import { setShowFromChainListOverlay } from '@/slices/bridgeDisplaySlice'
import { useBridgeState } from '@/slices/bridge/hooks'
import { CHAINS_BY_ID } from '@/constants/chains'
import { DropDownArrowSvg } from './components/DropDownArrowSvg'
import { setFromChainId } from '@/slices/bridge/reducer'

export const FromChainSelector = () => {
  const dispatch = useDispatch()

  const { fromChainId } = useBridgeState()

  let buttonContent

  if (fromChainId) {
    const src = CHAINS_BY_ID[fromChainId].chainImg.src
    const name = CHAINS_BY_ID[fromChainId].name

    buttonContent = (
      <div className="flex items-center space-x-3">
        <div>
          <img src={src} alt={name} className="w-5 h-5" />
        </div>
        <div className="text-left">
          <div className="text-xs text-secondaryTextColor">From</div>
          <div className="text-md text-primaryTextColor">{name}</div>
        </div>
        <DropDownArrowSvg />
      </div>
    )
  } else {
    buttonContent = (
      <div className="flex items-center space-x-3">
        <div className="text-left">
          <div className="text-xs text-secondaryTextColor">From</div>
          <div className="text-md text-primaryTextColor">Network</div>
        </div>
        <DropDownArrowSvg />
      </div>
    )
  }

  return (
    <button
      data-test-id="bridge-origin-chain-list-button"
      className=""
      onClick={() => dispatch(setShowFromChainListOverlay(true))}
    >
      <div className="flex items-center">{buttonContent}</div>
    </button>
  )
}
