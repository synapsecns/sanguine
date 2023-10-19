import { createSlice, PayloadAction } from '@reduxjs/toolkit'

import {
  updateUserHistoricalTransactions,
  updateIsUserHistoricalTransactionsLoading,
  updateUserPendingTransactions,
  updateIsUserPendingTransactionsLoading,
  addSeenHistoricalTransaction,
  addPendingAwaitingCompletionTransaction,
  removePendingAwaitingCompletionTransaction,
  addFallbackQueryPendingTransaction,
  removeFallbackQueryPendingTransaction,
  resetTransactionsState,
  updateFallbackQueryPendingTransaction,
  addFallbackQueryHistoricalTransaction,
  removeFallbackQueryHistoricalTransaction,
} from './actions'
import { BridgeTransaction } from '../api/generated'

export interface TransactionsState {
  userHistoricalTransactions: BridgeTransaction[]
  isUserHistoricalTransactionsLoading: boolean
  userPendingTransactions: BridgeTransaction[]
  isUserPendingTransactionsLoading: boolean
  seenHistoricalTransactions: BridgeTransaction[]
  pendingAwaitingCompletionTransactions: BridgeTransaction[]
  fallbackQueryPendingTransactions: BridgeTransaction[]
  fallbackQueryHistoricalTransactions: BridgeTransaction[]
}

const initialState: TransactionsState = {
  userHistoricalTransactions: [],
  isUserHistoricalTransactionsLoading: true,
  userPendingTransactions: [],
  isUserPendingTransactionsLoading: true,
  seenHistoricalTransactions: [],
  pendingAwaitingCompletionTransactions: [],
  fallbackQueryPendingTransactions: [],
  fallbackQueryHistoricalTransactions: [],
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
          const newSeenHistoricalTransaction = action.payload

          state.seenHistoricalTransactions = [
            newSeenHistoricalTransaction,
            ...state.seenHistoricalTransactions,
          ]
        }
      )
      .addCase(
        addPendingAwaitingCompletionTransaction,
        (state, action: PayloadAction<BridgeTransaction>) => {
          const newPendingTransaction = action.payload

          state.pendingAwaitingCompletionTransactions = [
            newPendingTransaction,
            ...state.pendingAwaitingCompletionTransactions,
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
      .addCase(
        addFallbackQueryPendingTransaction,
        (state, action: PayloadAction<BridgeTransaction>) => {
          const newfallbackTransaction = action.payload
          const filtered = state.fallbackQueryPendingTransactions.filter(
            (transaction: BridgeTransaction) =>
              transaction.kappa !== newfallbackTransaction.kappa
          )

          state.fallbackQueryPendingTransactions = [action.payload, ...filtered]
        }
      )
      .addCase(
        removeFallbackQueryPendingTransaction,
        (state, action: PayloadAction<string>) => {
          const kappa: string = action.payload

          state.fallbackQueryPendingTransactions =
            state.fallbackQueryPendingTransactions.filter(
              (transaction: BridgeTransaction) => transaction.kappa !== kappa
            )
        }
      )
      .addCase(
        updateFallbackQueryPendingTransaction,
        (state, action: PayloadAction<BridgeTransaction>) => {
          const { kappa }: BridgeTransaction = action.payload

          state.fallbackQueryPendingTransactions = [
            action.payload,
            ...state.fallbackQueryPendingTransactions.filter(
              (transaction: BridgeTransaction) => transaction.kappa !== kappa
            ),
          ]
        }
      )
      .addCase(
        addFallbackQueryHistoricalTransaction,
        (state, action: PayloadAction<BridgeTransaction>) => {
          const newFallbackTransaction = action.payload

          state.fallbackQueryHistoricalTransactions = [
            newFallbackTransaction,
            ...state.fallbackQueryHistoricalTransactions.filter(
              (transaction: BridgeTransaction) =>
                transaction.kappa !== newFallbackTransaction.kappa
            ),
          ]
        }
      )
      .addCase(
        removeFallbackQueryHistoricalTransaction,
        (state, action: PayloadAction<BridgeTransaction>) => {
          const fallbackTransaction = action.payload

          state.fallbackQueryHistoricalTransactions =
            state.fallbackQueryHistoricalTransactions.filter(
              (transaction: BridgeTransaction) => {
                transaction !== fallbackTransaction
              }
            )
        }
      )
      .addCase(resetTransactionsState, (state) => {
        state.userHistoricalTransactions =
          initialState.userHistoricalTransactions
        state.isUserHistoricalTransactionsLoading =
          initialState.isUserHistoricalTransactionsLoading
        state.userPendingTransactions = initialState.userPendingTransactions
        state.seenHistoricalTransactions =
          initialState.seenHistoricalTransactions
        state.pendingAwaitingCompletionTransactions =
          initialState.pendingAwaitingCompletionTransactions
        state.fallbackQueryPendingTransactions =
          initialState.fallbackQueryPendingTransactions
      })
  },
})

export default transactionsSlice.reducer
