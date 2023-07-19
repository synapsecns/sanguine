import { configureStore } from '@reduxjs/toolkit'

import bridgeReducer, { tokenDecimalMiddleware } from '@/slices/bridge/reducer'
import bridgeDisplayReducer from '@/slices/bridgeDisplaySlice'
import poolDataReducer from '@/slices/poolDataSlice'
import poolUserDataReducer from '@/slices/poolUserDataSlice'
import poolDepositReducer from '@/slices/poolDepositSlice'
import poolWithdrawReducer from '@/slices/poolWithdrawSlice'
import portfolioReducer from '@/slices/portfolio/reducer'

export const store = configureStore({
  reducer: {
    bridge: bridgeReducer,
    bridgeDisplay: bridgeDisplayReducer,
    poolData: poolDataReducer,
    poolUserData: poolUserDataReducer,
    poolDeposit: poolDepositReducer,
    poolWithdraw: poolWithdrawReducer,
    portfolio: portfolioReducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: false,
    }).concat(tokenDecimalMiddleware),
})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
