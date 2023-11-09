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
  const BASE_BUTTON_PROPERTIES = 'bg-transparent p-md border rounded-sm'

  let buttonContent
  let buttonClassName

  if (fromChainId) {
    buttonClassName = `
      ${BASE_BUTTON_PROPERTIES} 
      border-transparent
      ${getNetworkHover(fromChain?.color)}
      ${getNetworkButtonBgClassNameActive(fromChain?.color)}
      ${getNetworkButtonBorderActive(fromChain?.color)}
      ${getNetworkButtonBorderHover(fromChain?.color)}
    `
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
          <div className="text-sm text-secondaryTextColor">From</div>
          <div className="text-md text-primaryTextColor">{fromChain?.name}</div>
        </div>
        <DropDownArrowSvg />
      </div>
    )
  } else {
    buttonClassName = `
      ${BASE_BUTTON_PROPERTIES}
      border-transparent hover:border-secondary
    `

    buttonContent = (
      <div className="flex items-center space-x-3">
        <div className="text-left">
          <div className="text-sm text-secondaryTextColor">From</div>
          <div className="text-md text-primaryTextColor">Network</div>
        </div>
        <DropDownArrowSvg />
      </div>
    )
  }

  return (
    <button
      data-test-id="bridge-origin-chain-list-button"
      className={buttonClassName}
      onClick={() => dispatch(setShowFromChainListOverlay(true))}
    >
      <div className="flex items-center">{buttonContent}</div>
    </button>
  )
}
