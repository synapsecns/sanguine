import { configureStore } from '@reduxjs/toolkit'

import bridgeReducer from '@/slices/bridgeSlice'
import bridgeDisplayReducer from '@/slices/bridgeDisplaySlice'
import poolDataReducer from '@/slices/poolDataSlice'
import poolUserDataReducer from '@/slices/poolUserDataSlice'
import poolDepositReducer from '@/slices/poolDepositSlice'
import poolWithdrawReducer from '@/slices/poolWithdrawSlice'

export const store = configureStore({
  reducer: {
    bridge: bridgeReducer,
    bridgeDisplay: bridgeDisplayReducer,
    poolData: poolDataReducer,
    poolUserData: poolUserDataReducer,
    poolDeposit: poolDepositReducer,
    poolWithdraw: poolWithdrawReducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: false,
    }),
})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
