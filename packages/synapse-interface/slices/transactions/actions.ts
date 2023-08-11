import { createAction } from '@reduxjs/toolkit'

import { BridgeTransaction } from '../api/generated'

export const updateUserHistoricalTransactions = createAction<{
  userHistoricalTransactions: BridgeTransaction[]
}>('transactions/updateUserHistoricalTransactions')
