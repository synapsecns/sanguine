import { PayloadAction, createSlice } from '@reduxjs/toolkit'

import { Chain, Token } from '@/utils/types'

/** TODO: Rename entire slice once done refactoring prior Activity flow */
export interface _TransactionDetails {
  originChain: Chain
  originToken: Token
  destinationChain: Chain
  destinationToken: Token
  originValue: string
  originTxHash: string
  bridgeModuleName: string
  estimatedTime: number
  timestamp: number
  kappa: string
  isComplete: boolean
}

export interface _TransactionsState {
  transactions: any[]
}

export const initialState: _TransactionsState = {
  transactions: [],
}

export const transactionsSlice = createSlice({
  name: '_transactions',
  initialState,
  reducers: {
    addTransaction: (state, action: PayloadAction<any>) => {
      if (!Array.isArray(state.transactions)) {
        state.transactions = []; // Initialize to an empty array if not already an array
      }
      state.transactions.push(action.payload);
    },
    removeTransaction: (
      state,
      action: PayloadAction<{ originTxHash: string }>
    ) => {
      const { originTxHash } = action.payload
      state.transactions = state.transactions.filter(
        (tx) => tx.originTxHash !== originTxHash
      )
    },
    updateTransactionKappa: (
      state,
      action: PayloadAction<{ originTxHash: string; kappa: string }>
    ) => {
      const { originTxHash, kappa } = action.payload

      const txIndex = state.transactions.findIndex(
        (tx) => tx.originTxHash === originTxHash
      )

      if (txIndex !== -1) {
        state.transactions[txIndex].kappa = kappa
      }
    },
    completeTransaction: (
      state,
      action: PayloadAction<{ originTxHash: string; kappa: string }>
    ) => {
      const { originTxHash } = action.payload

      const txIndex = state.transactions.findIndex(
        (tx) => tx.originTxHash === originTxHash
      )
      if (txIndex !== -1) {
        state.transactions[txIndex].isComplete = true
      }
    },
    clearTransactions: (state) => {
      state.transactions = []
    },
  },
})

export const {
  addTransaction,
  removeTransaction,
  updateTransactionKappa,
  completeTransaction,
  clearTransactions,
} = transactionsSlice.actions

export default transactionsSlice.reducer
