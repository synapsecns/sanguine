import React from 'react'
import { useDispatch } from 'react-redux'

import { CHAINS_BY_ID } from '@/constants/chains'
import { DropDownArrowSvg } from '@/components/icons/DropDownArrowSvg'
import {
  getNetworkButtonBgClassNameActive,
  getNetworkButtonBorderActive,
  getNetworkButtonBorderHover,
  getNetworkHover,
} from '@/styles/chains'
import { useSwapState } from '@/slices/swap/hooks'
import { setShowSwapChainListOverlay } from '@/slices/swapDisplaySlice'

export const SwapChainSelector = () => {
  const dispatch = useDispatch()
  const { swapChainId } = useSwapState()
  const chain = CHAINS_BY_ID[swapChainId]

  const buttonContent = swapChainId ? (
    <div className="flex items-center space-x-3">
      <div>
        <img
          src={chain?.chainImg?.src}
          alt={chain?.name}
          className="w-6 h-6 rounded-sm"
        />
      </div>
      <div className="text-left">
        <div className="text-xs text-secondaryTextColor">From</div>
        <div className="text-md text-primaryTextColor">{chain.name}</div>
      </div>
      <DropDownArrowSvg />
    </div>
  ) : (
    <div className="flex items-center space-x-3">
      <div className="text-left">
        <div className="text-xs text-secondaryTextColor">From</div>
        <div className="text-md text-primaryTextColor">Network</div>
      </div>
      <DropDownArrowSvg />
    </div>
  )

  return (
    <button
      data-test-id="bridge-origin-chain-list-button"
      className={`
        bg-transparent
        p-md
        ${getNetworkHover(chain?.color)}
        ${getNetworkButtonBgClassNameActive(chain?.color)}
        border border-transparent
        ${getNetworkButtonBorderActive(chain?.color)}
        ${getNetworkButtonBorderHover(chain?.color)}
        rounded-sm
      `}
      onClick={() => dispatch(setShowSwapChainListOverlay(true))}
    >
      <div className="flex items-center">{buttonContent}</div>
    </button>
  )
}
