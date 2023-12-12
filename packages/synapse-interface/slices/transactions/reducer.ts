import { createSlice, PayloadAction } from '@reduxjs/toolkit'

import {
  PendingBridgeTransaction,
  addPendingBridgeTransaction,
  removePendingBridgeTransaction,
  updatePendingBridgeTransaction,
  updatePendingBridgeTransactions,
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
  pendingBridgeTransactions: PendingBridgeTransaction[]
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
  pendingBridgeTransactions: [],
}

export const transactionsSlice = createSlice({
  name: 'transactions',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(
        addPendingBridgeTransaction,
        (state, action: PayloadAction<PendingBridgeTransaction>) => {
          state.pendingBridgeTransactions = [
            action.payload,
            ...state.pendingBridgeTransactions,
          ]
        }
      )
      .addCase(
        updatePendingBridgeTransaction,
        (
          state,
          action: PayloadAction<{
            id: number
            timestamp: number
            transactionHash: string
            isSubmitted: boolean
          }>
        ) => {
          const { id, timestamp, transactionHash, isSubmitted } = action.payload
          const transactionIndex = state.pendingBridgeTransactions.findIndex(
            (transaction) => transaction.id === id
          )

          if (transactionIndex !== -1) {
            state.pendingBridgeTransactions =
              state.pendingBridgeTransactions.map((transaction, index) =>
                index === transactionIndex
                  ? { ...transaction, transactionHash, isSubmitted, timestamp }
                  : transaction
              )
          }
        }
      )
      .addCase(
        removePendingBridgeTransaction,
        (state, action: PayloadAction<number>) => {
          const idTimestampToRemove = action.payload
          state.pendingBridgeTransactions =
            state.pendingBridgeTransactions.filter(
              (transaction) => transaction.id !== idTimestampToRemove
            )
        }
      )
      .addCase(
        updatePendingBridgeTransactions,
        (state, action: PayloadAction<PendingBridgeTransaction[]>) => {
          state.pendingBridgeTransactions = action.payload
        }
      )
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
              (transaction: BridgeTransaction) => transaction?.kappa !== kappa
            )
        }
      )
      .addCase(
        addFallbackQueryPendingTransaction,
        (state, action: PayloadAction<BridgeTransaction>) => {
          const newfallbackTransaction = action.payload
          const filtered = state.fallbackQueryPendingTransactions.filter(
            (transaction: BridgeTransaction) =>
              transaction?.kappa !== newfallbackTransaction?.kappa
          )

          state.fallbackQueryPendingTransactions = [action.payload, ...filtered]
        }
      )
      .addCase(
        removeFallbackQueryPendingTransaction,
        (state, action: PayloadAction<string>) => {
          const kappa: string = action.payload

          state.fallbackQueryPendingTransactions = [
            ...state.fallbackQueryPendingTransactions.filter(
              (transaction: BridgeTransaction) => transaction.kappa !== kappa
            ),
          ]
        }
      )
      .addCase(
        updateFallbackQueryPendingTransaction,
        (state, action: PayloadAction<BridgeTransaction>) => {
          const { kappa }: BridgeTransaction = action.payload

          state.fallbackQueryPendingTransactions = [
            action.payload,
            ...state.fallbackQueryPendingTransactions.filter(
              (transaction: BridgeTransaction) => transaction?.kappa !== kappa
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
                transaction?.kappa !== newFallbackTransaction?.kappa
            ),
          ]
        }
      )
      .addCase(
        removeFallbackQueryHistoricalTransaction,
        (state, action: PayloadAction<BridgeTransaction>) => {
          const fallbackTransaction = action.payload

          state.fallbackQueryHistoricalTransactions = [
            ...state.fallbackQueryHistoricalTransactions.filter(
              (transaction: BridgeTransaction) => {
                transaction !== fallbackTransaction
              }
            ),
          ]
        }
      )
      .addCase(resetTransactionsState, (state) => {
        state.pendingBridgeTransactions = initialState.pendingBridgeTransactions
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
        state.fallbackQueryHistoricalTransactions =
          initialState.fallbackQueryHistoricalTransactions
      })
  },
})

export default transactionsSlice.reducer
