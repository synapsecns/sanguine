import { createAction } from '@reduxjs/toolkit'

import { BridgeTransaction } from '../api/generated'

export const updateUserHistoricalTransactions = createAction<
  BridgeTransaction[]
>('transactions/updateUserHistoricalTransactions')
export const updateIsUserHistoricalTransactionsLoading = createAction<boolean>(
  'transactions/updateIsUserHistoricalTransactionsLoading'
)
export const updateUserPendingTransactions = createAction<BridgeTransaction[]>(
  'transactions/updateUserPendingTransactions'
)
export const updateIsUserPendingTransactionsLoading = createAction<boolean>(
  'transactions/updateIsUserPendingTransactionsLoading'
)
export const addSeenHistoricalTransaction = createAction<BridgeTransaction>(
  'transactions/addSeenHistoricalTransaction'
)
export const addPendingAwaitingCompletionTransaction =
  createAction<BridgeTransaction>(
    'transactions/addPendingAwaitingCompletionTransaction'
  )
export const addFallbackQueryTransaction = createAction<BridgeTransaction>(
  'transactions/addFallbackQueryTransaction'
)
export const removePendingAwaitingCompletionTransaction = createAction<string>(
  'transactions/removePendingAwaitingCompletionTransaction'
)
export const resetTransactionsState = createAction<void>(
  'transactions/resetTransactionsState'
)
