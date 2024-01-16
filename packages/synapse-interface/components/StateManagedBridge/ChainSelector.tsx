import React from 'react'
import { useDispatch } from 'react-redux'

import {
  setShowFromChainListOverlay,
  setShowToChainListOverlay,
} from '@/slices/bridgeDisplaySlice'
import { useBridgeState } from '@/slices/bridge/hooks'
import { CHAINS_BY_ID } from '@/constants/chains'
import { DropDownArrowSvg } from '../icons/DropDownArrowSvg'
import {
  getNetworkButtonBgClassNameActive,
  getNetworkButtonBorderActive,
  getNetworkButtonBorderHover,
  getNetworkHover,
} from '@/styles/chains'

export const ChainSelector = ({ side }: { side: string }) => {
  console.log('side', side)
  const dispatch = useDispatch()
  const { fromChainId, toChainId } = useBridgeState()

  const chainId = side === 'from'
    ? fromChainId
    : toChainId

  const chain = CHAINS_BY_ID[chainId]

  return (
    <button
      data-test-id="bridge-origin-chain-list-button"
      className={`
        flex items-center gap-1.5
        p-2 rounded
        border border-transparent
        active:opacity-70
        ${getNetworkHover(chain?.color)}
        ${getNetworkButtonBorderHover(chain?.color)}
      `}
      onClick={() => side === 'from'
        ? dispatch(setShowFromChainListOverlay(true))
        : dispatch(setShowToChainListOverlay(true))
      }
    >
      {chain && (
        <img
          src={chain?.chainImg?.src}
          alt={chain?.name}
          className="w-6 h-6 rounded-sm"
        />
      )}
      <dl className="text-left">
        <dt className="text-sm opacity-50 capitalize">{side}</dt>
        <dd>{chain?.name || 'Network'}</dd>
      </dl>
      <DropDownArrowSvg className="mx-0.5" />
    </button>
  )
}
