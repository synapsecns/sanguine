import { createSlice, PayloadAction } from '@reduxjs/toolkit'

export interface TransactionDetails {
  originAmount: string
  originChainId: number
  destinationChainId: number
  originTxHash: string
  bridgeModuleName: string
  estimatedTime: number
  timestamp: number
  kappa?: string
  isComplete?: boolean
}

export interface TransactionState {
  transactions: TransactionDetails[]
}

export const initialState: TransactionState = {
  transactions: [],
}

export const transactionsSlice = createSlice({
  name: 'transactions',
  initialState,
  reducers: {
    addTransaction: (state, action: PayloadAction<TransactionDetails>) => {
      state.transactions.push(action.payload)
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
