import { PayloadAction, createSlice } from '@reduxjs/toolkit'

import { Chain, Token } from '@/utils/types'
import StateManagedBridge from '@/pages/state-managed-bridge'

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
      state.transactions = [...action.payload]
    },
    removeTransaction: (
      transactions: _TransactionsState,
      { payload: { originTxHash } }
    ) => {
      if (transactions[originTxHash]) {
        delete transactions[originTxHash]
      }
    },
    updateTransactionKappa: (
      transactions: _TransactionsState,
      { payload: { originTxHash, kappa } }
    ) => {
      const tx = transactions[originTxHash]
      if (!tx) return

      tx.kappa = kappa
    },
    completeTransaction: (
      transactions: _TransactionsState,
      { payload: { originTxHash } }
    ) => {
      const tx = transactions[originTxHash]
      if (!tx) return

      tx.isComplete = true
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
