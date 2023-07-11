import { configureStore } from '@reduxjs/toolkit'

import bridgeReducer, { tokenDecimalMiddleware } from '@/slices/bridgeSlice'
import bridgeDisplayReducer from '@/slices/bridgeDisplaySlice'
import poolDataSlice from '@/slices/poolDataSlice'
import poolUserDataSlice from '@/slices/poolUserDataSlice'

export const store = configureStore({
  reducer: {
    bridge: bridgeReducer,
    bridgeDisplay: bridgeDisplayReducer,
    poolData: poolDataSlice,
    poolUserData: poolUserDataSlice,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: false,
    }).concat(tokenDecimalMiddleware),
})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
