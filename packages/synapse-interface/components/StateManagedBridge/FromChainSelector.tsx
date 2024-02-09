import React from 'react'
import { useDispatch } from 'react-redux'

import { setShowFromChainListOverlay } from '@/slices/bridgeDisplaySlice'
import { useBridgeState } from '@/slices/bridge/hooks'
import { CHAINS_BY_ID } from '@/constants/chains'
import { DropDownArrowSvg } from '../icons/DropDownArrowSvg'



import {
  getNetworkButtonBgClassNameActive,
  getNetworkButtonBorderActive,
  getNetworkButtonBorderHover,
  getNetworkHover,
} from '@/styles/chains'

export const FromChainSelector = () => {
  const dispatch = useDispatch()
  const { fromChainId } = useBridgeState()
  const fromChain = CHAINS_BY_ID[fromChainId]

  let buttonContent

  if (fromChainId) {
    buttonContent = (
      <div className="flex items-center space-x-3">
        <div>
          <img
            src={fromChain?.chainImg?.src}
            alt={fromChain?.name}
            className="w-6 h-6 rounded-sm"
          />
        </div>
        <div className="text-left">
          <div className="text-xs text-secondaryTextColor">From</div>
          <div className="text-md text-primaryTextColor">{fromChain?.name}</div>
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
      className={`
        bg-transparent
        p-md
        ${getNetworkHover(fromChain?.color)}
        ${getNetworkButtonBgClassNameActive(fromChain?.color)}
        border border-transparent
        ${getNetworkButtonBorderActive(fromChain?.color)}
        ${getNetworkButtonBorderHover(fromChain?.color)}
        rounded-sm
      `}
      onClick={() => dispatch(setShowFromChainListOverlay(true))}
    >
      <div className="flex items-center">{buttonContent}</div>
    </button>
  )
}
