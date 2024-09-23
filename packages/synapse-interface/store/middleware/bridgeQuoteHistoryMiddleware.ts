import {
  Middleware,
  MiddlewareAPI,
  Dispatch,
  AnyAction,
} from '@reduxjs/toolkit'

export const bridgeQuoteHistoryMiddleware: Middleware =
  (store: MiddlewareAPI) => (next: Dispatch) => (action: AnyAction) => {
    const previousState = store.getState()
    const result = next(action)
    const currentState = store.getState()

    if (
      previousState.bridgeQuote.bridgeQuote !==
      currentState.bridgeQuote.bridgeQuote
    ) {
      store.dispatch({
        type: 'bridgeQuote/setPreviousBridgeQuote',
        payload: previousState.bridgeQuote.bridgeQuote,
      })
    }

    return result
  }
