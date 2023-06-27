import { configureStore } from '@reduxjs/toolkit'

import bridgeReducer, { tokenDecimalMiddleware } from '@/slices/bridgeSlice'
import bridgeDisplayReducer from '@/slices/bridgeDisplaySlice'

export const store = configureStore({
  reducer: {
    bridge: bridgeReducer,
    bridgeDisplay: bridgeDisplayReducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(tokenDecimalMiddleware),
})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
