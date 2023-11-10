import { createAction } from '@reduxjs/toolkit'

export const updateDebouncedFromValue = createAction<string>(
  'bridge/updateDebouncedFromValue'
)
export const updateDebouncedToTokensFromValue = createAction<string>(
  'bridge/updateDebouncedToTokensFromValue'
)
export const resetBridgeInputs = createAction<void>('bridge/resetBridgeInputs')
export const resetFetchedBridgeQuotes = createAction<void>(
  'bridge/resetFetchedBridgeQuotes'
)
