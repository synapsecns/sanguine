import { createSlice, PayloadAction } from '@reduxjs/toolkit'

import {
  updateUserHistoricalTransactions,
  updateIsUserHistoricalTransactionsLoading,
  updateUserPendingTransactions,
  updateIsUserPendingTransactionsLoading,
  addSeenHistoricalTransaction,
  addPendingAwaitingCompletionTransaction,
  removePendingAwaitingCompletionTransaction,
  resetTransactionsState,
} from './actions'
import { BridgeTransaction } from '../api/generated'

export interface TransactionsState {
  userHistoricalTransactions: BridgeTransaction[]
  isUserHistoricalTransactionsLoading: boolean
  userPendingTransactions: BridgeTransaction[]
  isUserPendingTransactionsLoading: boolean
  seenHistoricalTransactions: BridgeTransaction[]
  pendingAwaitingCompletionTransactions: BridgeTransaction[]
}

const initialState: TransactionsState = {
  userHistoricalTransactions: [],
  isUserHistoricalTransactionsLoading: true,
  userPendingTransactions: [],
  isUserPendingTransactionsLoading: true,
  seenHistoricalTransactions: [],
  pendingAwaitingCompletionTransactions: [],
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
          state.seenHistoricalTransactions = [
            ...state.seenHistoricalTransactions,
            action.payload,
          ]
        }
      )
      .addCase(
        addPendingAwaitingCompletionTransaction,
        (state, action: PayloadAction<BridgeTransaction>) => {
          state.pendingAwaitingCompletionTransactions = [
            ...state.pendingAwaitingCompletionTransactions,
            action.payload,
          ]
        }
      )
      .addCase(
        removePendingAwaitingCompletionTransaction,
        (state, action: PayloadAction<string>) => {
          const kappa: string = action.payload

          state.pendingAwaitingCompletionTransactions =
            state.pendingAwaitingCompletionTransactions.filter(
              (transaction: BridgeTransaction) => transaction.kappa !== kappa
            )
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
