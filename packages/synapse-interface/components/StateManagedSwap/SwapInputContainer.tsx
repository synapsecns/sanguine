import React, { useEffect, useState, useRef, useCallback } from 'react'
import { useDispatch } from 'react-redux'
import { useAccount, useNetwork } from 'wagmi'

import MiniMaxButton from '../buttons/MiniMaxButton'
import { formatBigIntToString, stringToBigInt } from '@/utils/bigint/format'
import { cleanNumberInput } from '@/utils/cleanNumberInput'

import { SwapChainSelector } from './SwapChainSelector'
import { SwapFromTokenSelector } from './SwapFromTokenSelector'
import { usePortfolioState } from '@/slices/portfolio/hooks'
import { updateSwapFromValue } from '@/slices/swap/reducer'
import { useSwapState } from '@/slices/swap/hooks'
import { ConnectStatusIndicator } from '@/components/bridgeSwap/ConnectStatusIndicator'
import { GenericInputContainer } from '@/components/bridgeSwap/GenericInputContainer'

export const SwapInputContainer = () => {
  const inputRef = useRef<HTMLInputElement>(null)
  const { swapChainId, swapFromToken, swapFromValue } = useSwapState()

  const dispatch = useDispatch()


  return (
    <GenericInputContainer
      inputRef={inputRef}
      chainId={swapChainId}
      token={swapFromToken}
      value={swapFromValue}
      initialStateValue={null} // initialState.fromValue
      dispatchUpdateFunc={(val) => dispatch(updateSwapFromValue(val))}
      chainSelector={<SwapChainSelector />}
      tokenSelector={<SwapFromTokenSelector />}
    />
  )
}
