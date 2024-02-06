import { configureStore } from '@reduxjs/toolkit'

import bridge from '@/state/slices/bridge/reducer'
import wallet from '@/state/slices/wallet/reducer'
import transactions from '@/state/slices/transactions/reducer'
import bridgeQuote from '@/state/slices/bridgeQuote/reducer'
import approveTransaction from '@/state/slices/approveTransaction/reducer'
import bridgeTransaction from '@/state/slices/bridgeTransaction/reducer'

// This allows Redux Devtools to have access to BigInts which it cannot otherwise serialize
declare global {
  interface BigInt {
    toJSON(): string
  }
}

BigInt.prototype.toJSON = function () {
  return this.toString()
}

export const store = configureStore({
  reducer: {
    approveTransaction,
    bridge,
    bridgeQuote,
    bridgeTransaction,
    wallet,
    transactions,
  },

  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: false,
    }),
})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
