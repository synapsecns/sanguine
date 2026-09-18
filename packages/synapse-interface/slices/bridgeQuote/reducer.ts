import { createSlice, PayloadAction } from '@reduxjs/toolkit'

import { EMPTY_BRIDGE_QUOTE } from '@/constants/bridge'
import { type BridgeQuote } from '@/utils/types'
import { fetchBridgeQuote } from './thunks'

export interface BridgeQuoteState {
  bridgeQuote: BridgeQuote
  previousBridgeQuote: BridgeQuote | null
  isLoading: boolean
  currentRequestId?: string
}

export const initialState: BridgeQuoteState = {
  bridgeQuote: EMPTY_BRIDGE_QUOTE,
  previousBridgeQuote: null,
  isLoading: false,
}

export const bridgeQuoteSlice = createSlice({
  name: 'bridgeQuote',
  initialState,
  reducers: {
    setIsLoading: (state, action: PayloadAction<boolean>) => {
      state.isLoading = action.payload
    },
    resetBridgeQuote: (state) => {
      state.bridgeQuote = initialState.bridgeQuote
      state.currentRequestId = undefined
      state.isLoading = false
    },
    setPreviousBridgeQuote: (state, action: PayloadAction<any>) => {
      state.previousBridgeQuote = action.payload
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchBridgeQuote.pending, (state, action) => {
        state.isLoading = true
        state.currentRequestId = action.meta.requestId
      })
      .addCase(fetchBridgeQuote.fulfilled, (state, action) => {
        if (state.currentRequestId !== action.meta.requestId) return
        state.bridgeQuote = action.payload
        state.isLoading = false
      })
      .addCase(fetchBridgeQuote.rejected, (state, action) => {
        if (state.currentRequestId !== action.meta.requestId) return
        state.bridgeQuote = EMPTY_BRIDGE_QUOTE
        state.isLoading = false
      })
  },
})

export const { resetBridgeQuote, setIsLoading, setPreviousBridgeQuote } =
  bridgeQuoteSlice.actions

export default bridgeQuoteSlice.reducer
