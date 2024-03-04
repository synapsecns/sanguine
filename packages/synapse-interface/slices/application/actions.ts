import { createAction } from '@reduxjs/toolkit'
import type { Address } from 'viem'

export enum RootActions {
  RESET_REDUX_CACHE = 'reset_redux_cache',
}

export const resetReduxCache = createAction(RootActions.RESET_REDUX_CACHE)
export const updateLastConnectedAddress = createAction<Address>(
  'application/updateLastConnectedAddress'
)
export const updateLastConnectedTime = createAction<number>(
  'application/updateLastConnectedTime'
)
