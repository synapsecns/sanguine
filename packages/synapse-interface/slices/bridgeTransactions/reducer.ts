import { createSlice } from '@reduxjs/toolkit'

/** TODO: Possibly rename entire slice once done refactoring prior Activity flow */
export interface BridgeTransactionDetails {
  originChainId: number
  destinationChainId: number
  originTxHash: string
  bridgeModuleName: string
  estimatedTime: number
  timestamp: number
  kappa: string
  isComplete: boolean
}

export interface BridgeTransactionsState {
  [transactionHash: string]: BridgeTransactionDetails
}

export const initialState: BridgeTransactionsState = {}

export const transactionsSlice = createSlice({
  name: 'transactions',
  initialState,
  reducers: {
    addTransaction: (
      transactions: BridgeTransactionsState,
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
      transactions: BridgeTransactionsState,
      { payload: { originTxHash } }
    ) => {
      if (transactions[originTxHash]) {
        delete transactions[originTxHash]
      }
    },
    updateTransactionKappa: (
      transactions: BridgeTransactionsState,
      { payload: { originTxHash, kappa } }
    ) => {
      const tx = transactions[originTxHash]
      if (!tx) return

      tx.kappa = kappa
    },
    completeTransaction: (
      transactions: BridgeTransactionsState,
      { payload: { originTxHash } }
    ) => {
      const tx = transactions[originTxHash]
      if (!tx) return

      tx.isComplete = true
    },
    clearTransactions: (transactions: BridgeTransactionsState) => {
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
