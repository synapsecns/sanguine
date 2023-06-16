import { configureStore } from '@reduxjs/toolkit'

import bridgeReducer from '../slices/bridgeSlice'

export const store = configureStore({
  reducer: {
    bridge: bridgeReducer,
  },
})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
