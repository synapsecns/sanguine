import _ from 'lodash'
import { configureStore } from '@reduxjs/toolkit'
import { getAccount } from '@wagmi/core'

import bridgeReducer from '@/slices/bridge/reducer'
import bridgeDisplayReducer from '@/slices/bridgeDisplaySlice'
import poolDataReducer from '@/slices/poolDataSlice'
import poolUserDataReducer from '@/slices/poolUserDataSlice'
import poolDepositReducer from '@/slices/poolDepositSlice'
import poolWithdrawReducer from '@/slices/poolWithdrawSlice'
import portfolioReducer from '@/slices/portfolio/reducer'
import { api } from '@/slices/api/slice'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'

export const store = configureStore({
  reducer: {
    bridge: bridgeReducer,
    bridgeDisplay: bridgeDisplayReducer,
    poolData: poolDataReducer,
    poolUserData: poolUserDataReducer,
    poolDeposit: poolDepositReducer,
    poolWithdraw: poolWithdrawReducer,
    portfolio: portfolioReducer,
    [api.reducerPath]: api.reducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: false,
    }).concat(api.middleware),
})

let previousState = store.getState()

store.subscribe(() => {
  const account = getAccount()
  const { address } = account

  const currentState = store.getState()
  const bridgeState = currentState.bridge

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
    const { fromChainId, toChainId, fromToken, toToken, fromValue } =
      bridgeState

    eventTitle = `[Bridge System Action] Generate bridge quote`
    eventData = {
      address,
      fromChainId,
      toChainId,
      fromToken: fromToken.symbol,
      toToken: toToken.symbol,
      inputAmountString: fromValue,
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

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
