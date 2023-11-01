import { Action, combineReducers } from '@reduxjs/toolkit'
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
import swapDisplay from '@/slices/swapDisplaySlice'
import { api } from '@/slices/api/slice'

const persistedReducers = {
  bridge,
  transactions,
}

const key: string = 'synapse-interface'

const persistConfig: PersistConfig<AppState> = {
  key,
  storage,
  whitelist: Object.keys(persistedReducers),
}

export const appReducer = combineReducers({
  portfolio,
  swap,
  bridgeDisplay,
  poolData,
  poolDeposit,
  poolUserData,
  poolWithdraw,
  priceData,
  swapDisplay,
  [api.reducerPath]: api.reducer,
  ...persistedReducers,
})

export enum RootActions {
  RESET_REDUX_CACHE = 'reset_redux_cache',
}

export const rootReducer = (
  state: AppState | undefined,
  action: Action<string>
) => {
  if (action.type === RootActions.RESET_REDUX_CACHE) {
    // persistor.purge()
    localStorage.removeItem(`persist:${key}`)

    state = undefined
  }
  return appReducer(state, action)
}

export type AppState = ReturnType<typeof appReducer>

const persistedReducer = persistReducer(persistConfig, appReducer)

export default persistedReducer
