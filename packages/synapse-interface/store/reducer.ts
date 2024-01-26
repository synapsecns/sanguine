import { Action, combineReducers } from '@reduxjs/toolkit'
import { PersistConfig, persistReducer } from 'redux-persist'
import storage from 'redux-persist/lib/storage'

import application from '@/slices/application/reducer'
import _transactions from '@/slices/_transactions/reducer'
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
import swapDisplay from '@/slices/swapDisplaySlice'
import feeAndRebate from '@/slices/feeAndRebateSlice'
import { api } from '@/slices/api/slice'
import { RootActions } from '@/slices/application/actions'

const persistedReducers = {
  application,
  transactions,
  _transactions,
}

export const storageKey: string = 'synapse-interface'

export const persistConfig: PersistConfig<AppState> = {
  version: 1, // upgrade to reset cache when updated data structures throw errors
  key: storageKey,
  storage,
  whitelist: Object.keys(persistedReducers),
}

export const appReducer = combineReducers({
  bridge,
  portfolio,
  swap,
  bridgeDisplay,
  poolData,
  poolDeposit,
  poolUserData,
  poolWithdraw,
  priceData,
  swapDisplay,
  feeAndRebate,
  [api.reducerPath]: api.reducer,
  ...persistedReducers,
})

export const rootReducer = (
  state: AppState | undefined,
  action: Action<string>
) => {
  if (action.type === RootActions.RESET_REDUX_CACHE) {
    localStorage.removeItem(`persist:${storageKey}`)
    state = undefined
  }
  return appReducer(state, action)
}

export type AppState = ReturnType<typeof appReducer>

export const persistedReducer = persistReducer(persistConfig, rootReducer)
