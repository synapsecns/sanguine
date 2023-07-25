import { Address } from 'viem'
import { createAction } from '@reduxjs/toolkit'

import { Token } from '@/utils/types'

export enum PortfolioTabs {
  HOME = 'home',
  PORTFOLIO = 'portfolio',
}

export const setActiveTab = createAction<PortfolioTabs>(
  'portfolio/setActiveTab'
)
export const updateSingleTokenAllowance = createAction<{
  allowance: bigint
  spender: Address
  owner: Address
  token: Token
}>('portfolio/updateSingleTokenAllowance')
