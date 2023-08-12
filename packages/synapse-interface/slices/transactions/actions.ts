import { createAction } from '@reduxjs/toolkit'

import { BridgeTransaction } from '../api/generated'

export const updateUserHistoricalTransactions = createAction<
  BridgeTransaction[]
>('transactions/updateUserHistoricalTransactions')
export const updateUserPendingTransactions = createAction<BridgeTransaction[]>(
  'transactions/updateUserPendingTransactions'
)
