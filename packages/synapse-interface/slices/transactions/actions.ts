import { createAction } from '@reduxjs/toolkit'

import { BridgeTransaction } from '../api/generated'
import { Chain, Token } from '@/utils/types'

export interface PendingBridgeTransaction {
  id: number
  originChain: Chain
  originToken: Token
  originValue: string
  destinationChain: Chain
  destinationToken: Token
  transactionHash?: string
  timestamp: number
  isSubmitted: boolean
  estimatedTime: number
  bridgeModuleName: string
}

export const addPendingBridgeTransaction =
  createAction<PendingBridgeTransaction>(
    'transactions/addPendingBridgeTransaction'
  )
export const updatePendingBridgeTransaction = createAction<{
  id: number
  timestamp: number
  transactionHash: string
  isSubmitted: boolean
}>('transactions/updatePendingBridgeTransaction')
export const removePendingBridgeTransaction = createAction<number>(
  'transactions/removePendingBridgeTransaction'
)
export const updatePendingBridgeTransactions = createAction<
  PendingBridgeTransaction[]
>('transactions/updatePendingBridgeTransactions')

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
export const removePendingAwaitingCompletionTransaction = createAction<string>(
  'transactions/removePendingAwaitingCompletionTransaction'
)
export const addFallbackQueryPendingTransaction =
  createAction<BridgeTransaction>(
    'transactions/addFallbackQueryPendingTransaction'
  )
export const updateFallbackQueryPendingTransaction =
  createAction<BridgeTransaction>(
    'transactions/updateFallbackQueryPendingTransaction'
  )
export const removeFallbackQueryPendingTransaction = createAction<string>(
  'transactions/removeFallbackQueryPendingTransaction'
)
export const addFallbackQueryHistoricalTransaction =
  createAction<BridgeTransaction>(
    'transactions/addFallbackQueryHistoricalTransaction'
  )
export const removeFallbackQueryHistoricalTransaction =
  createAction<BridgeTransaction>(
    'transaction/removeFallbackQueryHistoricalTransaction'
  )
export const resetTransactionsState = createAction<void>(
  'transactions/resetTransactionsState'
)
