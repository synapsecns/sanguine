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

  return (
    <button
      data-test-id="bridge-origin-chain-list-button"
      className={`
        flex items-center gap-1.5
        p-2 rounded
        border border-transparent
        active:opacity-70
        ${getNetworkHover(fromChain?.color)}
        ${getNetworkButtonBorderHover(fromChain?.color)}
      `}
      onClick={() => dispatch(setShowFromChainListOverlay(true))}
    >
      {fromChainId && (
        <img
          src={fromChain?.chainImg?.src}
          alt={fromChain?.name}
          className="w-6 h-6 rounded-sm"
        />
      )}
      <dl className="text-left">
        <dt className="text-sm opacity-50">From</dt>
        <dd>{fromChain?.name || 'Network'}</dd>
      </dl>
      <DropDownArrowSvg className="mx-0.5" />
    </button>
  )
}
