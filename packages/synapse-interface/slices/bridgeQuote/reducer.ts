import { createSlice, PayloadAction } from '@reduxjs/toolkit'

import { EMPTY_BRIDGE_QUOTE } from '@/constants/bridge'
import { type BridgeQuote } from '@/utils/types'

export interface BridgeQuoteState {
  bridgeQuote: BridgeQuote
  isLoading: boolean
}

export const initialState: BridgeQuoteState = {
  bridgeQuote: EMPTY_BRIDGE_QUOTE,
  isLoading: false,
}

export const bridgeQuoteSlice = createSlice({
  name: 'bridgeQuote',
  initialState,
  reducers: {
    setIsLoading: (state, action: PayloadAction<boolean>) => {
      state.isLoading = action.payload
    },
    setBridgeQuote: (state, action: PayloadAction<BridgeQuote>) => {
      state.bridgeQuote = action.payload
    },
    resetBridgeQuote: (state) => {
      state.bridgeQuote = initialState.bridgeQuote
    },
  },
})

export const { setBridgeQuote, resetBridgeQuote, setIsLoading } =
  bridgeQuoteSlice.actions

export default bridgeQuoteSlice.reducer
