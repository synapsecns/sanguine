import React from 'react'
import { useDispatch } from 'react-redux'

import { setShowToChainListOverlay } from '@/slices/bridgeDisplaySlice'
import { useBridgeState } from '@/slices/bridge/hooks'
import { CHAINS_BY_ID } from '@/constants/chains'
import { DropDownArrowSvg } from '../icons/DropDownArrowSvg'
import {
  getNetworkButtonBgClassNameActive,
  getNetworkButtonBorderActive,
  getNetworkButtonBorderHover,
  getNetworkHover,
} from '@/styles/chains'

export const ToChainSelector = () => {
  const dispatch = useDispatch()
  const { toChainId } = useBridgeState()
  const toChain = CHAINS_BY_ID[toChainId]
  const BASE_BUTTON_PROPERTIES = 'bg-transparent p-md border rounded-sm'

  let buttonContent
  let buttonClassName

  if (toChainId) {
    buttonClassName = `
      ${BASE_BUTTON_PROPERTIES} 
      border-transparent
      ${getNetworkHover(toChain?.color)}
      ${getNetworkButtonBgClassNameActive(toChain?.color)}
      ${getNetworkButtonBorderActive(toChain?.color)}
      ${getNetworkButtonBorderHover(toChain?.color)}
    `

    buttonContent = (
      <div className="flex items-center space-x-3">
        <div className="">
          <img
            src={toChain?.chainImg?.src}
            alt={toChain?.name}
            className="w-6 h-6 rounded-sm"
          />
        </div>
        <div className="text-left">
          <div className="text-sm text-secondaryTextColor">To</div>
          <div className="text-md text-primaryTextColor">{toChain?.name}</div>
        </div>
        <DropDownArrowSvg />
      </div>
    )
  } else {
    buttonClassName = `
      ${BASE_BUTTON_PROPERTIES}
      border border-transparent hover:border-secondary
    `

    buttonContent = (
      <div className="flex items-center space-x-3">
        <div className="text-left">
          <div className="text-sm text-secondaryTextColor">To</div>
          <div className="text-md text-primaryTextColor">Network</div>
        </div>
        <DropDownArrowSvg />
      </div>
    )
  }

  return (
    <button
      className={buttonClassName}
      onClick={() => dispatch(setShowToChainListOverlay(true))}
    >
      {buttonContent}
    </button>
  )
}
