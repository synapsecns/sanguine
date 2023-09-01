import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { Address } from 'wagmi'

import { ETH, NETH } from '@/constants/tokens/bridgeable'
import { EMPTY_BRIDGE_QUOTE } from '@/constants/bridge'
import { ARBITRUM, ETH as ETHEREUM } from '@/constants/chains/master'
import { BridgeQuote, SwapQuote, Token } from '@/utils/types'
import {
  getRoutePossibilities,
  getSymbol,
} from '@/utils/routeMaker/generateRoutePossibilities'
import { getFromChainIds } from '@/utils/routeMaker/getFromChainIds'
import { getFromTokens } from '@/utils/routeMaker/getFromTokens'
import { getToChainIds } from '@/utils/routeMaker/getToChainIds'
import { getToTokens } from '@/utils/routeMaker/getToTokens'
import { findTokenByRouteSymbol } from '@/utils/findTokenByRouteSymbol'
import { EMPTY_SWAP_QUOTE } from '@/constants/swap'

export interface BridgeState {
  fromChainId: number
  fromToken: Token
  toToken: Token

  fromValue: string
  swapQuote: SwapQuote // should be SwapQuote
  isLoading: boolean
}

export const initialState: BridgeState = {
  fromChainId: 1,
  fromToken: ETH,
  toToken: NETH,

  fromValue: '',
  swapQuote: EMPTY_SWAP_QUOTE,
  isLoading: false,
}

export const swapSlice = createSlice({
  name: 'swap',
  initialState,
  reducers: {
    setIsLoading: (state, action: PayloadAction<boolean>) => {
      state.isLoading = action.payload
    },
    setFromChainId: (state, action: PayloadAction<number>) => {
      state.fromChainId = action.payload
    },
    setFromToken: (state, action: PayloadAction<Token>) => {
      state.fromToken = action.payload
    },
    setToToken: (state, action: PayloadAction<Token>) => {
      state.toToken = action.payload
    },
    setSwapQuote: (state, action: PayloadAction<SwapQuote>) => {
      state.swapQuote = action.payload
    },
    updateFromValue: (state, action: PayloadAction<string>) => {
      state.fromValue = action.payload
    },
  },
})

export const {
  setFromChainId,
  setFromToken,
  setToToken,
  updateFromValue,
  setSwapQuote,
  setIsLoading,
} = swapSlice.actions

export default swapSlice.reducer
