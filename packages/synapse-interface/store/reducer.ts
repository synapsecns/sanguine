import { combineReducers } from '@reduxjs/toolkit'
import { PersistConfig, persistReducer } from 'redux-persist'
import storage from 'redux-persist/lib/storage'

import bridge from '@/slices/bridge/reducer'
import portfolio from '@/slices/portfolio/reducer'
import swap from '@/slices/swap/reducer'
import transactions from '@/slices/transactions/reducer'
import bridgeDisplay from '@/slices/bridgeDisplaySlice'
import poolData from '@/slices/poolDataSlice'
import poolDeposit from '@/slices/poolDepositSlice'
import poolUserData from '@/slices/poolUserDataSlice'
import poolWithdraw from '@/slices/poolWithdrawSlice'
import priceData from '@/slices/priceDataSlice'
import swapDisplaySlice from '@/slices/swapDisplaySlice'
import { api } from '@/slices/api/slice'

const persistedReducers = {
  bridge,
  transactions,
}

const persistConfig: PersistConfig<AppState> = {
  key: 'synapse-interface',
  storage,
  whitelist: Object.keys(persistedReducers),
}

const appReducer = combineReducers({
  portfolio,
  swap,
  bridgeDisplay,
  poolData,
  poolDeposit,
  poolUserData,
  poolWithdraw,
  priceData,
  swapDisplaySlice,
  [api.reducerPath]: api.reducer,
  ...persistedReducers,
})

export type AppState = ReturnType<typeof appReducer>

const persistedReducer = persistReducer(persistConfig, appReducer)

export default persistedReducer
