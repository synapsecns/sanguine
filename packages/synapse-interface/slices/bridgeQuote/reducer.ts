import { createSlice, PayloadAction } from '@reduxjs/toolkit'

import { EMPTY_BRIDGE_QUOTE } from '@/constants/bridge'
import { type BridgeQuote } from '@/utils/types'
import { fetchBridgeQuote } from './thunks'

export interface BridgeQuoteState {
  bridgeQuote: BridgeQuote
  isLoading: boolean
  error: any
}

export const initialState: BridgeQuoteState = {
  bridgeQuote: EMPTY_BRIDGE_QUOTE,
  isLoading: false,
  error: null,
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
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchBridgeQuote.pending, (state) => {
        // state.status = FetchState.LOADING
        state.isLoading = true
      })
      .addCase(
        fetchBridgeQuote.fulfilled,
        (state, action: PayloadAction<BridgeQuote>) => {
          state.bridgeQuote = action.payload
          // state.status = FetchState.VALID
          state.isLoading = false
        }
      )
      .addCase(fetchBridgeQuote.rejected, (state, action) => {
        // state.error = action.payload
        state.bridgeQuote = EMPTY_BRIDGE_QUOTE
        // state.status = FetchState.INVALID
        state.isLoading = false
      })
  },
})

export const { resetBridgeQuote, setIsLoading } = bridgeQuoteSlice.actions

export default bridgeQuoteSlice.reducer
