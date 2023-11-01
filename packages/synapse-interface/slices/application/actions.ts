import { createAction } from '@reduxjs/toolkit'

import { RootActions } from '@/store/reducer'

export const resetReduxCache = createAction(RootActions.RESET_REDUX_CACHE)
