import { configureStore } from '@reduxjs/toolkit'
import { serializeBalance } from '@/utils/bigint/serialization'

import bridgeReducer, { tokenDecimalMiddleware } from '@/slices/bridgeSlice'
import bridgeDisplayReducer from '@/slices/bridgeDisplaySlice'
import { serialize } from 'wagmi'

const bigIntSerializationMiddleware = ({ dispatch, getState }) => next => action => {
  // Check if the action is setSupportedFromTokenBalances and its payload exists
  if (action.type === 'bridge/setSupportedFromTokenBalances' && action.payload) {
    action.payload = serializeBalance(action.payload);
  }

  if(action.type=== 'bridge/updateFromValue'  && action.payload) {
    action.payload = serialize(action.payload);
  }

  // Pass action to the next middleware/reducer
  let result = next(action);

  // Return the result
  return result;
};


export const store = configureStore({
  reducer: {
    bridge: bridgeReducer,
    bridgeDisplay: bridgeDisplayReducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: false,
    }).concat(tokenDecimalMiddleware).concat(bigIntSerializationMiddleware),
})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
