import { createSlice, PayloadAction } from '@reduxjs/toolkit'

export interface TransactionDetails {
  txHash: string
  originChainId: number
  destinationChainId: number
  timestamp: number
  kappa?: string
}

export interface TransactionState {
  [txHash: string]: TransactionDetails
}

export const initialState: TransactionState = {}

export const transactionsSlice = createSlice({
  name: 'transactions',
  initialState,
  reducers: {
    addTransaction: (
      transactions: TransactionState,
      { payload: { txHash } }
    ) => {
      transactions[txHash] = txHash
    },
    removeTransaction: (
      transactions: TransactionState,
      { payload: { txHash } }
    ) => {
      if (transactions[txHash]) {
        delete transactions[txHash]
      }
    },
    finalizeTransaction: (
      transactions: TransactionState,
      { payload: { txHash, kappa } }
    ) => {
      const tx = transactions[txHash]
      if (!tx) return

      tx.kappa = kappa
    },
    clearTransactions: (transactions: TransactionState) => {
      transactions = {}
    },
  },
})

export default transactionsSlice.reducer
