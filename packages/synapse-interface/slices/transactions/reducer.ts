import { createSlice, PayloadAction } from '@reduxjs/toolkit'

import { updateUserHistoricalTransactions } from './actions'
import { BridgeTransaction } from '../api/generated'

export interface TransactionsState {
  userHistoricalTransactions: BridgeTransaction[]
  userPendingTransactions: BridgeTransaction[]
}

const initialState: TransactionsState = {
  userHistoricalTransactions: [],
  userPendingTransactions: [],
}

export const transactionsSlice = createSlice({
  name: 'transactions',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(
      updateUserHistoricalTransactions,
      (state, action: PayloadAction<BridgeTransaction[]>) => {
        state.userHistoricalTransactions = action.payload
      }
    )
  },
})

export default transactionsSlice.reducer
