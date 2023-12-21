import { createSlice, PayloadAction } from '@reduxjs/toolkit'

export interface TransactionDetails {
  originChainId: number
  destinationChainId: number
  originTxHash: string
  bridgeModuleName: string
  estimatedTime: number
  timestamp: number
  kappa: string
  isComplete: boolean
}

export interface TransactionState {
  [transactionHash: string]: TransactionDetails
}

export const initialState: TransactionState = {}

export const transactionsSlice = createSlice({
  name: 'transactions',
  initialState,
  reducers: {
    addTransaction: (
      transactions: TransactionState,
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
      transactions: TransactionState,
      { payload: { originTxHash } }
    ) => {
      if (transactions[originTxHash]) {
        delete transactions[originTxHash]
      }
    },
    updateTransactionKappa: (
      transactions: TransactionState,
      { payload: { originTxHash, kappa } }
    ) => {
      const tx = transactions[originTxHash]
      if (!tx) return

      tx.kappa = kappa
    },
    completeTransaction: (
      transactions: TransactionState,
      { payload: { originTxHash, kappa } }
    ) => {
      const tx = transactions[originTxHash]
      if (!tx) return

      tx.isComplete = true
    },
    clearTransactions: (transactions: TransactionState) => {
      transactions = {}
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
