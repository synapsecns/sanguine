import { combineReducers } from '@reduxjs/toolkit'
import { PersistConfig, persistReducer } from 'redux-persist'
import localForage from 'localforage'
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
import swapDisplaySlice from '@/slices/swapDisplaySlice'
import { api as dataApi } from '@/slices/api/slice'

const persistedReducers = {
  bridge,
  transactions,
}

const appReducer = combineReducers({
  portfolio,
  swap,
  bridgeDisplay,
  poolData,
  poolDeposit,
  poolUserData,
  poolWithdraw,
  swapDisplaySlice,
  [dataApi.reducerPath]: dataApi.reducer,
  ...persistedReducers,
})

export type AppState = ReturnType<typeof appReducer>

const persistConfig: PersistConfig<AppState> = {
  key: 'synapse-interface',
  storage,
  whitelist: Object.keys(persistedReducers),
}

const persistedReducer = persistReducer(persistConfig, appReducer)

export default persistedReducer
