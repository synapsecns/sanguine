import { createSlice } from '@reduxjs/toolkit'

import { fetchBridgeQuote } from './hooks'

export enum FetchState {
  IDLE = 'idle',
  LOADING = 'loading',
  VALID = 'valid',
  INVALID = 'invalid',
}

export type BridgeQuoteTransaction = {
  data: string
  to: string
  value?: string | null
}

export type BridgeQuote = {
  outputAmount: bigint
  outputAmountString: string
  routerAddress: string
  exchangeRate: bigint
  feeAmount: bigint
  nativeFee: bigint
  delta: bigint
  estimatedTime: number | null
  bridgeModuleName: string | null
  tx: BridgeQuoteTransaction | null
  quoteAddress: string | null
  requestId: number | null
  timestamp: number | null
}

export const EMPTY_BRIDGE_QUOTE: BridgeQuote = {
  outputAmount: 0n,
  outputAmountString: '',
  routerAddress: '',
  exchangeRate: 0n,
  feeAmount: 0n,
  nativeFee: 0n,
  delta: 0n,
  estimatedTime: null,
  bridgeModuleName: null,
  tx: null,
  quoteAddress: null,
  requestId: null,
  timestamp: null,
}

export interface BridgeQuoteState {
  isLoading: boolean
  bridgeQuote: BridgeQuote
  currentRequestId: number | null
  status: string
  error: any
}

const initialState: BridgeQuoteState = {
  isLoading: false,
  bridgeQuote: EMPTY_BRIDGE_QUOTE,
  currentRequestId: null,
  status: FetchState.IDLE,
  error: null,
}

export const bridgeQuoteSlice = createSlice({
  name: 'bridgeQuote',
  initialState,
  reducers: {
    resetQuote: (state) => {
      state.bridgeQuote = EMPTY_BRIDGE_QUOTE
      state.currentRequestId = null
      state.status = FetchState.IDLE
      state.error = null
      state.isLoading = false
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchBridgeQuote.pending, (state, action) => {
        state.currentRequestId = action.meta.arg.requestId
        state.status = FetchState.LOADING
        state.isLoading = true
      })
      .addCase(fetchBridgeQuote.fulfilled, (state, action) => {
        if (state.currentRequestId !== action.payload.requestId) {
          return
        }

        state.bridgeQuote = action.payload
        state.currentRequestId = null
        state.status = FetchState.VALID
        state.isLoading = false
      })
      .addCase(fetchBridgeQuote.rejected, (state, action) => {
        if (state.currentRequestId !== action.meta.arg.requestId) {
          return
        }

        state.error = action.payload
        state.bridgeQuote = EMPTY_BRIDGE_QUOTE
        state.currentRequestId = null
        state.status = FetchState.INVALID
        state.isLoading = false
      })
  },
})

export const { resetQuote } = bridgeQuoteSlice.actions

export default bridgeQuoteSlice.reducer
