import React from 'react'
import { useDispatch } from 'react-redux'

import { useBridgeState } from '@/slices/bridge/hooks'
import { initialState, updateFromValue } from '@/slices/bridge/reducer'
import {
  setShowFromChainListOverlay,
  setShowFromTokenListOverlay
 } from '@/slices/bridgeDisplaySlice'

import { ChainSelector } from '@/components/bridgeSwap/ChainSelector'
import { TokenSelector } from '@/components/bridgeSwap/TokenSelector'
import { GenericInputContainer } from '@/components/bridgeSwap/GenericInputContainer'

export const inputRef = React.createRef<HTMLInputElement>()

export const InputContainer = () => {
  const { fromChainId, fromToken, fromValue } = useBridgeState()

  const dispatch = useDispatch()

  return (
    <GenericInputContainer
      inputRef={inputRef}
      chainId={fromChainId}
      token={fromToken}
      value={fromValue}
      initialStateValue={initialState.fromValue}
      dispatchUpdateFunc={(val) => dispatch(updateFromValue(val))}
      chainSelector={
        <ChainSelector
          data-test-id="bridge-origin-chain-list-button"
          chainId={fromChainId}
          label="From"
          onClick={() => dispatch(setShowFromChainListOverlay(true))}
        />
      }
      tokenSelector={
        <TokenSelector
          data-test-id="bridge-origin-token"
          token={fromToken}
          label="In"
          onClick={() => dispatch(setShowFromTokenListOverlay(true))}
        />
      }
    />
  )
}
