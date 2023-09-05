import { createSlice, PayloadAction } from '@reduxjs/toolkit'

import { DAI, NUSD, USDC } from '@/constants/tokens/bridgeable'
import { SwapQuote, Token } from '@/utils/types'
import { EMPTY_SWAP_QUOTE } from '@/constants/swap'
import { SWAP_CHAIN_IDS } from '@/constants/existingSwapRoutes'

export interface SwapState {
  swapChainId: number
  swapFromToken: Token
  swapToToken: Token
  swapFromChainIds: number[]
  swapFromTokens: Token[]
  swapToTokens: Token[]

  swapFromValue: string
  swapQuote: SwapQuote
  isLoading: boolean
}

export const initialState: SwapState = {
  swapChainId: 1,
  swapFromToken: USDC,
  swapToToken: DAI,
  swapFromChainIds: SWAP_CHAIN_IDS,
  swapFromTokens: [USDC, NUSD, DAI],
  swapToTokens: [USDC, NUSD, DAI],

  swapFromValue: '',
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
    setSwapChainId: (state, action: PayloadAction<number>) => {
      state.swapChainId = action.payload
    },
    setSwapFromToken: (state, action: PayloadAction<Token>) => {
      state.swapFromToken = action.payload
    },
    setSwapToToken: (state, action: PayloadAction<Token>) => {
      state.swapToToken = action.payload
    },
    setSwapQuote: (state, action: PayloadAction<SwapQuote>) => {
      state.swapQuote = action.payload
    },
    updateSwapFromValue: (state, action: PayloadAction<string>) => {
      state.swapFromValue = action.payload
    },
  },
})

export const {
  setSwapChainId,
  setSwapFromToken,
  setSwapToToken,
  updateSwapFromValue,
  setSwapQuote,
  setIsLoading,
} = swapSlice.actions

export default swapSlice.reducer
