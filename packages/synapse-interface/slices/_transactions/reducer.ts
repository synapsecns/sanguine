import { createSlice } from '@reduxjs/toolkit'

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
          originValue,
          bridgeModuleName,
          originChain,
          originToken,
          destinationChain,
          destinationToken,
          estimatedTime,
          timestamp,
        },
      }
    ) => {
      if (!originTxHash) return

      transactions[originTxHash] = {
        originTxHash,
        originValue,
        bridgeModuleName,
        originChain,
        originToken,
        destinationChain,
        destinationToken,
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
