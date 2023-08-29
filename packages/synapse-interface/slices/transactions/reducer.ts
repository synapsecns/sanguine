import { createSlice, PayloadAction } from '@reduxjs/toolkit'

import {
  updateUserHistoricalTransactions,
  updateIsUserHistoricalTransactionsLoading,
  updateUserPendingTransactions,
  updateIsUserPendingTransactionsLoading,
  addSeenHistoricalTransaction,
  resetTransactionsState,
} from './actions'
import { BridgeTransaction } from '../api/generated'

export interface TransactionsState {
  userHistoricalTransactions: BridgeTransaction[]
  isUserHistoricalTransactionsLoading: boolean
  userPendingTransactions: BridgeTransaction[]
  isUserPendingTransactionsLoading: boolean
  seenHistoricalTransaction: BridgeTransaction[]
}

const initialState: TransactionsState = {
  userHistoricalTransactions: [],
  isUserHistoricalTransactionsLoading: true,
  userPendingTransactions: [],
  isUserPendingTransactionsLoading: true,
  seenHistoricalTransaction: [],
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
      .addCase(
        updateIsUserPendingTransactionsLoading,
        (state, action: PayloadAction<boolean>) => {
          state.isUserPendingTransactionsLoading = action.payload
        }
      )
      .addCase(
        addSeenHistoricalTransaction,
        (state, action: PayloadAction<BridgeTransaction>) => {
          state.seenHistoricalTransaction = [
            ...state.seenHistoricalTransaction,
            action.payload,
          ]
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
