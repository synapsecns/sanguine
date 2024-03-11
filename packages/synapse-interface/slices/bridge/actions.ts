import { createAction } from '@reduxjs/toolkit'

export const updateDebouncedFromValue = createAction<string>(
  'bridge/updateDebouncedFromValue'
)
