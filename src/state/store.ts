import { configureStore } from '@reduxjs/toolkit'

import bridge from '@/state/slices/bridge/reducer'
import wallet from '@/state/slices/wallet/reducer'

export const store = configureStore({
  reducer: {
    bridge,
    wallet,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: false,
    }),
})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
