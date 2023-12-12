import { configureStore } from '@reduxjs/toolkit'

import bridge from '@/state/slices/bridge/reducer'

export const store = configureStore({
  reducer: {
    bridge,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: false,
    }),
})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
