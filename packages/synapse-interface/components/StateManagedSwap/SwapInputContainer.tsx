import { useRef } from 'react'
import { useDispatch } from 'react-redux'

import { useSwapState } from '@/slices/swap/hooks'
import { updateSwapFromValue } from '@/slices/swap/reducer'
import {
  setShowSwapChainListOverlay,
  setShowSwapFromTokenListOverlay
} from '@/slices/swapDisplaySlice'

import { ChainSelector } from '@/components/bridgeSwap/ChainSelector'
import { TokenSelector } from '@/components/bridgeSwap/TokenSelector'
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
      chainSelector={
        <ChainSelector
          data-test-id="swap-chain-list-button"
          chainId={swapChainId}
          label="On"
          onClick={() => dispatch(setShowSwapChainListOverlay(true))}
        />
      }
      tokenSelector={
        <TokenSelector
          data-test-id="swap-origin-token"
          token={swapFromToken}
          label="In"
          onClick={() => dispatch(setShowSwapFromTokenListOverlay(true))}
        />
      }
    />
  )
}
