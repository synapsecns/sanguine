import { configureStore } from '@reduxjs/toolkit'

import bridgeReducer, {
  segmentMiddleware,
  tokenDecimalMiddleware,
} from '@/slices/bridgeSlice'
import bridgeDisplayReducer from '@/slices/bridgeDisplaySlice'

export const store = configureStore({
  reducer: {
    bridge: bridgeReducer,
    bridgeDisplay: bridgeDisplayReducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(tokenDecimalMiddleware, segmentMiddleware),
})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
