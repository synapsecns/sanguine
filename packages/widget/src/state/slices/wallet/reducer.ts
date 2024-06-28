import { createSlice, PayloadAction } from '@reduxjs/toolkit'

import { fetchAndStoreTokenBalances, fetchAndStoreAllowance } from './hooks'
import { TokenBalance } from '@/utils/actions/fetchTokenBalances'

export enum FetchState {
  IDLE = 'idle',
  LOADING = 'loading',
  VALID = 'valid',
  INVALID = 'invalid',
}

export interface WalletState {
  balances: TokenBalance[]
  allowance: string
  status: FetchState
  error?: any
  isWalletPending: boolean
}

const initialState: WalletState = {
  balances: [],
  allowance: null,
  status: FetchState.IDLE,
  error: null,
  isWalletPending: false,
}

export const walletSlice = createSlice({
  name: 'wallet',
  initialState,
  reducers: {
    setIsWalletPending: (state, action: PayloadAction<boolean>) => {
      state.isWalletPending = action.payload
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchAndStoreTokenBalances.pending, (state) => {
        state.balances = []
        state.status = FetchState.LOADING
        state.error = null
      })
      .addCase(
        fetchAndStoreTokenBalances.fulfilled,
        (state, action: PayloadAction<TokenBalance[]>) => {
          state.balances = action.payload
          state.status = FetchState.VALID
          state.error = null
        }
      )
      .addCase(fetchAndStoreTokenBalances.rejected, (state, action) => {
        state.error = action.payload
        state.status = FetchState.INVALID
      })
      .addCase(fetchAndStoreAllowance.pending, (state) => {
        state.status = FetchState.LOADING
        state.error = null
      })
      .addCase(
        fetchAndStoreAllowance.fulfilled,
        (state, action: PayloadAction<string>) => {
          state.allowance = action.payload
          state.status = FetchState.VALID
          state.error = null
        }
      )
      .addCase(fetchAndStoreAllowance.rejected, (state, action) => {
        state.error = action.payload
        state.status = FetchState.INVALID
      })
  },
})

export const { setIsWalletPending } = walletSlice.actions

export default walletSlice.reducer
