import { createSlice, PayloadAction } from '@reduxjs/toolkit'

import { EMPTY_BRIDGE_QUOTE } from '@/constants/bridge'
import { type BridgeQuote } from '@/utils/types'
import { fetchBridgeQuote } from './thunks'

export interface BridgeQuoteState {
  bridgeQuote: BridgeQuote
  previousBridgeQuote: BridgeQuote | null
  isLoading: boolean
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
    },
    setPreviousBridgeQuote: (state, action: PayloadAction<any>) => {
      state.previousBridgeQuote = action.payload
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchBridgeQuote.pending, (state) => {
        state.isLoading = true
      })
      .addCase(
        fetchBridgeQuote.fulfilled,
        (state, action: PayloadAction<BridgeQuote>) => {
          state.bridgeQuote = action.payload
          state.isLoading = false
        }
      )
      .addCase(fetchBridgeQuote.rejected, (state) => {
        state.bridgeQuote = EMPTY_BRIDGE_QUOTE
        state.isLoading = false
      })
  },
})

export const { resetBridgeQuote, setIsLoading, setPreviousBridgeQuote } =
  bridgeQuoteSlice.actions

export default bridgeQuoteSlice.reducer
