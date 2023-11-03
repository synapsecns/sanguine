import { createAction } from '@reduxjs/toolkit'
import { Address } from 'viem'

import { RootActions } from '@/store/reducer'

export const resetReduxCache = createAction(RootActions.RESET_REDUX_CACHE)
export const updateLastConnectedAddress = createAction<Address>(
  'application/updateLastConnectedAddress'
)
