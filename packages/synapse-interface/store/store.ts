import _ from 'lodash'
import { configureStore } from '@reduxjs/toolkit'
import { persistStore } from 'redux-persist'

import { api } from '@/slices/api/slice'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { storageKey, persistConfig, persistedReducer } from './reducer'
import { resetReduxCache } from '@/slices/application/actions'

const checkVersionAndResetCache = (): boolean => {
  if (typeof window !== 'undefined') {
    const persistedStateRaw = localStorage.getItem(`persist:${storageKey}`)
    if (persistedStateRaw) {
      const persistedState = JSON.parse(persistedStateRaw)
      const persistedVersion = JSON.parse(persistedState._persist)

      if (persistedVersion.version !== persistConfig.version) {
        return true
      }
    }
  }
  return false
}

export const store = configureStore({
  reducer: persistedReducer,
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: false,
    }).concat(api.middleware),
})

if (checkVersionAndResetCache()) {
  store.dispatch(resetReduxCache())
}

export const persistor = persistStore(store)

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch

let previousState = store.getState()

store.subscribe(() => {
  const currentState = store.getState()
  const bridgeState = currentState.bridge

  const address = currentState.application?.lastConnectedAddress

  let eventTitle
  let eventData

  if (
    !_.isEqual(
      previousState.bridge.bridgeQuote,
      currentState.bridge.bridgeQuote
    ) &&
    currentState.bridge.bridgeQuote.outputAmount !== 0n
  ) {
    const { outputAmountString, routerAddress, exchangeRate } =
      bridgeState.bridgeQuote
    const { fromChainId, toChainId, fromToken, toToken, debouncedFromValue } =
      bridgeState

    eventTitle = `[Bridge System Action] Generate bridge quote`
    eventData = {
      address,
      fromChainId,
      toChainId,
      fromToken: fromToken?.symbol,
      toToken: toToken?.symbol,
      inputAmountString: debouncedFromValue,
      outputAmountString,
      routerAddress,
      exchangeRate: BigInt(exchangeRate.toString()),
    }
    segmentAnalyticsEvent(eventTitle, eventData)
  }

  if (
    previousState.bridgeDisplay.showDestinationAddress === false &&
    currentState.bridgeDisplay.showDestinationAddress === true
  ) {
    eventTitle = `[Bridge User Action] Show destination address`
    eventData = {}

    segmentAnalyticsEvent(eventTitle, eventData)
  }

  if (
    previousState.bridgeDisplay.showDestinationAddress === true &&
    currentState.bridgeDisplay.showDestinationAddress === false
  ) {
    eventTitle = `[Bridge User Action] Hide destination address`
    eventData = {}

    segmentAnalyticsEvent(eventTitle, eventData)
  }

  if (
    previousState.bridgeDisplay.showSettingsSlideOver === false &&
    currentState.bridgeDisplay.showSettingsSlideOver === true
  ) {
    eventTitle = `[Bridge User Action] Show Settings`
    eventData = {}
    segmentAnalyticsEvent(eventTitle, eventData)
  }

  if (
    previousState.bridgeDisplay.showSettingsSlideOver === true &&
    currentState.bridgeDisplay.showSettingsSlideOver === false
  ) {
    eventTitle = `[Bridge User Action] Hide Settings`
    eventData = {}
    segmentAnalyticsEvent(eventTitle, eventData)
  }

  previousState = currentState
})
