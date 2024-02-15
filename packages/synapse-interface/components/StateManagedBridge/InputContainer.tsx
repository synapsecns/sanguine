import React, { useEffect, useState, useCallback } from 'react'
import { useDispatch } from 'react-redux'
import { useAccount } from 'wagmi'

import { initialState, updateFromValue } from '@/slices/bridge/reducer'
import MiniMaxButton from '../buttons/MiniMaxButton'
import { formatBigIntToString } from '@/utils/bigint/format'
import { cleanNumberInput } from '@/utils/cleanNumberInput'

import { FromChainSelector } from './FromChainSelector'
import { FromTokenSelector } from './FromTokenSelector'
import { useBridgeState } from '@/slices/bridge/hooks'
import { usePortfolioState } from '@/slices/portfolio/hooks'
import { ConnectStatusIndicator } from '@/components/bridgeSwap/ConnectStatusIndicator'
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
      value={fromValue} // fromValue
      initialStateValue={initialState.fromValue} // initialState.fromValue
      dispatchUpdateFunc={(val) => dispatch(updateFromValue(val))} // (inputValue) => dispatch(updateFromValue(inputValue))
      chainSelector={<FromChainSelector />}
      tokenSelector={<FromTokenSelector />}
    />
  )
}
