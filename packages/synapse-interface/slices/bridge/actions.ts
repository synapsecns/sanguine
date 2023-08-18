import { createAction } from '@reduxjs/toolkit'

import { Token } from '@/utils/types'

export interface RecentBridgeTransaction {
  fromChainId: number
  fromToken: Token
  fromValue: string
  toChainId: number
  toToken: Token
  transactionHash: string
}

export const addRecentBridgeTransaction = createAction<RecentBridgeTransaction>(
  'bridge/addRecentBridgeTransaction'
)
