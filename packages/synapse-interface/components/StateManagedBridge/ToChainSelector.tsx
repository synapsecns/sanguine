import React from 'react'
import { useDispatch } from 'react-redux'

import { setShowToChainListOverlay } from '@/slices/bridgeDisplaySlice'
import { useBridgeState } from '@/slices/bridge/hooks'
import { CHAINS_BY_ID } from '@/constants/chains'
import { DropDownArrowSvg } from './components/DropDownArrowSvg'

export const ToChainSelector = () => {
  const dispatch = useDispatch()

  const { toChainId } = useBridgeState()

  let buttonContent

  if (toChainId) {
    const src = CHAINS_BY_ID[toChainId].chainImg.src
    const name = CHAINS_BY_ID[toChainId].name

    buttonContent = (
      <div className="flex items-center space-x-3">
        <div>
          <img src={src} alt={name} className="w-5 h-5" />
        </div>
        <div className="text-left">
          <div className="text-xs text-secondaryTextColor">To</div>
          <div className="text-md text-primaryTextColor">{name}</div>
        </div>
        <DropDownArrowSvg />
      </div>
    )
  } else {
    buttonContent = (
      <div className="flex items-center space-x-3">
        <div className="text-left">
          <div className="text-xs text-secondaryTextColor">To</div>
          <div className="text-md text-primaryTextColor">Network</div>
        </div>
        <DropDownArrowSvg />
      </div>
    )
  }

  return (
    <button
      className=""
      onClick={() => dispatch(setShowToChainListOverlay(true))}
    >
      {buttonContent}
    </button>
  )
}
