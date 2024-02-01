import { createSlice, PayloadAction } from '@reduxjs/toolkit'

import { fetchBridgeQuote } from './hooks'

export enum FetchState {
  IDLE = 'idle',
  LOADING = 'loading',
  VALID = 'valid',
  INVALID = 'invalid',
}

export type BridgeQuote = {
  outputAmount: bigint
  outputAmountString: string
  routerAddress: string
  exchangeRate: bigint
  feeAmount: bigint
  delta: bigint
  quotes: { originQuery: {}; destQuery: {} }
  estimatedTime: number
  bridgeModuleName: string
}

export const EMPTY_BRIDGE_QUOTE = {
  outputAmount: 0n,
  outputAmountString: '',
  routerAddress: '',
  exchangeRate: 0n,
  feeAmount: 0n,
  delta: 0n,
  quotes: { originQuery: null, destQuery: null },
  estimatedTime: null,
  bridgeModuleName: null,
}

export interface BridgeQuoteState {
  isLoading: boolean
  bridgeQuote: BridgeQuote
  status: string
  error: any
}

const initialState: BridgeQuoteState = {
  isLoading: false,
  bridgeQuote: EMPTY_BRIDGE_QUOTE,
  status: FetchState.IDLE,
  error: null,
}

export const bridgeQuoteSlice = createSlice({
  name: 'bridgeQuote',
  initialState,
  reducers: {
    resetQuote: (state) => {
      state.bridgeQuote = EMPTY_BRIDGE_QUOTE
      state.status = FetchState.IDLE
      state.error = null
      state.isLoading = false
    },
    setIsLoading: (state, action: PayloadAction<boolean>) => {
      state.isLoading = action.payload
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchBridgeQuote.pending, (state) => {
        state.status = FetchState.LOADING
        // state.isLoading = true
      })
      .addCase(
        fetchBridgeQuote.fulfilled,
        (state, action: PayloadAction<BridgeQuote>) => {
          state.bridgeQuote = action.payload
          state.status = FetchState.VALID
          state.isLoading = false
        }
      )
      .addCase(fetchBridgeQuote.rejected, (state, action) => {
        state.error = action.payload
        state.bridgeQuote = EMPTY_BRIDGE_QUOTE
        state.status = FetchState.INVALID
        state.isLoading = false
      })
  },
})

export const { resetQuote, setIsLoading } = bridgeQuoteSlice.actions

export default bridgeQuoteSlice.reducer
