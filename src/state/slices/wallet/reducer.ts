import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { fetchAndStoreTokenBalances } from './hooks'
import { TokenBalance } from '@/utils/actions/fetchTokenBalances'

export enum FetchState {
  IDLE = 'idle',
  LOADING = 'loading',
  VALID = 'valid',
  INVALID = 'invalid',
}

export interface WalletState {
  balances: TokenBalance[]
  status: FetchState
  error?: any
}

const initialState: WalletState = {
  balances: [],
  status: FetchState.IDLE,
  error: null,
}

export const walletSlice = createSlice({
  name: 'wallet',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchAndStoreTokenBalances.pending, (state) => {
        state.status = FetchState.LOADING
      })
      .addCase(
        fetchAndStoreTokenBalances.fulfilled,
        (state, action: PayloadAction<TokenBalance[]>) => {
          state.balances = action.payload
        }
      )
      .addCase(fetchAndStoreTokenBalances.rejected, (state, action) => {
        state.error = action.payload
      })
  },
})

export default walletSlice.reducer
