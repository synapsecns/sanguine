import { createAction } from '@reduxjs/toolkit'

import { Token } from '@/utils/types'

export interface RecentBridgeTransaction {
  fromChainId: number
  fromToken: Token
  fromValue: string
  toChainId: number
  toToken: Token
  transactionHash: string
  timestamp: number
}

export const addRecentBridgeTransaction = createAction<RecentBridgeTransaction>(
  'bridge/addRecentBridgeTransaction'
)
export const updateRecentBridgeTransactions = createAction<
  RecentBridgeTransaction[]
>('bridge/updateRecentBridgeTransactions')
