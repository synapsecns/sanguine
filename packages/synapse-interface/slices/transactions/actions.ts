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
export const resetTransactionsState = createAction<void>(
  'transactions/resetTransactionsState'
)
