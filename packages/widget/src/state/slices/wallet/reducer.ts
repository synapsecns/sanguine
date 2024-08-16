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
  balancesFetchStatus: FetchState
  balancesFetchError?: any
  allowance: string
  allowancesFetchStatus: FetchState
  allowancesFetchError?: any
  isWalletPending: boolean
}

const initialState: WalletState = {
  balances: [],
  balancesFetchStatus: FetchState.IDLE,
  balancesFetchError: null,
  allowance: null,
  allowancesFetchStatus: FetchState.IDLE,
  allowancesFetchError: null,
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
        state.balancesFetchStatus = FetchState.LOADING
        state.balancesFetchError = null
      })
      .addCase(
        fetchAndStoreTokenBalances.fulfilled,
        (state, action: PayloadAction<TokenBalance[]>) => {
          state.balances = action.payload
          state.balancesFetchStatus = FetchState.VALID
          state.balancesFetchError = null
        }
      )
      .addCase(fetchAndStoreTokenBalances.rejected, (state, action) => {
        state.balancesFetchError = action.payload
        state.balancesFetchStatus = FetchState.INVALID
      })
      .addCase(fetchAndStoreAllowance.pending, (state) => {
        state.allowancesFetchStatus = FetchState.LOADING
        state.allowancesFetchError = null
      })
      .addCase(
        fetchAndStoreAllowance.fulfilled,
        (state, action: PayloadAction<string>) => {
          state.allowance = action.payload
          state.allowancesFetchStatus = FetchState.VALID
          state.allowancesFetchError = null
        }
      )
      .addCase(fetchAndStoreAllowance.rejected, (state, action) => {
        state.allowancesFetchError = action.payload
        state.allowancesFetchStatus = FetchState.INVALID
      })
  },
})

export const { setIsWalletPending } = walletSlice.actions

export default walletSlice.reducer
