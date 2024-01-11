import { Address } from 'viem'
import { createAction } from '@reduxjs/toolkit'

import { Token } from '@/utils/types'

export enum PortfolioTabs {
  PORTFOLIO = 'portfolio',
  ACTIVITY = 'activity',
}

export enum FetchState {
  IDLE = 'idle',
  LOADING = 'loading',
  VALID = 'valid',
  INVALID = 'invalid',
}

export const setActiveTab = createAction<PortfolioTabs>(
  'portfolio/setActiveTab'
)
export const typeSearchInput = createAction<{ searchInput: string }>(
  'portfolio/typeSearchInput'
)

export const resetPortfolioState = createAction<void>(
  'portfolio/resetPortfolioState'
)
export const resetSearchState = createAction<void>('portfolio/resetSearchState')
