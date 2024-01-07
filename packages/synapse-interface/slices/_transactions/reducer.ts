import { createSlice } from '@reduxjs/toolkit'

/** TODO: Rename entire slice once done refactoring prior Activity flow */
export interface _TransactionDetails {
  originChainId: number
  destinationChainId: number
  originTxHash: string
  bridgeModuleName: string
  estimatedTime: number
  timestamp: number
  kappa: string
  isComplete: boolean
}

export interface _TransactionsState {
  [transactionHash: string]: _TransactionDetails
}

export const initialState: _TransactionsState = {}

export const transactionsSlice = createSlice({
  name: 'transactions',
  initialState,
  reducers: {
    addTransaction: (
      transactions: _TransactionsState,
      {
        payload: {
          originTxHash,
          bridgeModuleName,
          originChainId,
          destinationChainId,
          estimatedTime,
          timestamp,
        },
      }
    ) => {
      if (!originTxHash) return

      transactions[originTxHash] = {
        originChainId,
        destinationChainId,
        originTxHash,
        bridgeModuleName,
        estimatedTime,
        timestamp,
        kappa: null,
        isComplete: false,
      }
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
    clearTransactions: (transactions: _TransactionsState) => {
      transactions = {} // eslint-disable-line
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
