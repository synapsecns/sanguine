import { createSlice, PayloadAction } from '@reduxjs/toolkit'

import {
  updateUserHistoricalTransactions,
  updateUserPendingTransactions,
} from './actions'
import { BridgeTransaction } from '../api/generated'

export interface TransactionsState {
  userHistoricalTransactions: BridgeTransaction[]
  isUserHistoricalTransactionsLoading: boolean
  userPendingTransactions: BridgeTransaction[]
}

const initialState: TransactionsState = {
  userHistoricalTransactions: [],
  isUserHistoricalTransactionsLoading: true,
  userPendingTransactions: [],
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
        updateUserPendingTransactions,
        (state, action: PayloadAction<BridgeTransaction[]>) => {
          state.userPendingTransactions = action.payload
        }
      )
  },
})

export default transactionsSlice.reducer
