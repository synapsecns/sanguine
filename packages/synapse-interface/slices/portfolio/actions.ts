import { Address } from 'viem'
import { createAction } from '@reduxjs/toolkit'

import { Token } from '@/utils/types'

export enum PortfolioTabs {
  HOME = 'home',
  PORTFOLIO = 'portfolio',
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
export const updateSingleTokenAllowance = createAction<{
  chainId: number
  allowance: bigint
  spender: Address
  token: Token
}>('portfolio/updateSingleTokenAllowance')
export const resetPortfolioState = createAction<void>(
  'portfolio/resetPortfolioState'
)
