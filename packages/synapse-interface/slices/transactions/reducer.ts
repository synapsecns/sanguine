import { createSlice, PayloadAction } from '@reduxjs/toolkit'

import {
  updateUserHistoricalTransactions,
  updateIsUserHistoricalTransactionsLoading,
  updateUserPendingTransactions,
  resetTransactionsState,
} from './actions'
import { BridgeTransaction } from '../api/generated'

export interface TransactionsState {
  userHistoricalTransactions: BridgeTransaction[]
  isUserHistoricalTransactionsLoading: boolean
  userPendingTransactions: BridgeTransaction[]
}

const initialState: TransactionsState = {
  userHistoricalTransactions: [],
  isUserHistoricalTransactionsLoading: true,
  userPendingTransactions: [],
}

export const transactionsSlice = createSlice({
  name: 'transactions',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(
        updateUserHistoricalTransactions,
        (state, action: PayloadAction<BridgeTransaction[]>) => {
          state.userHistoricalTransactions = action.payload
        }
      )
      .addCase(
        updateIsUserHistoricalTransactionsLoading,
        (state, action: PayloadAction<boolean>) => {
          state.isUserHistoricalTransactionsLoading = action.payload
        }
      )
      .addCase(
        updateUserPendingTransactions,
        (state, action: PayloadAction<BridgeTransaction[]>) => {
          state.userPendingTransactions = action.payload
        }
      )
      .addCase(resetTransactionsState, (state) => {
        state.userHistoricalTransactions =
          initialState.userHistoricalTransactions
        state.isUserHistoricalTransactionsLoading =
          initialState.isUserHistoricalTransactionsLoading
        state.userPendingTransactions = initialState.userPendingTransactions
      })
  },
})

export default transactionsSlice.reducer
