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
  chainId: number
  allowance: bigint
  spender: Address
  token: Token
}>('portfolio/updateSingleTokenAllowance')
