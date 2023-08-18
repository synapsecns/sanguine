import { createAction } from '@reduxjs/toolkit'

import { Chain, Token } from '@/utils/types'

export interface RecentBridgeTransaction {
  originChain: Chain
  originToken: Token
  originValue: string
  destinationChain: Chain
  destinationToken: Token
  transactionHash: string
  timestamp: number
}

export const addRecentBridgeTransaction = createAction<RecentBridgeTransaction>(
  'bridge/addRecentBridgeTransaction'
)
export const updateRecentBridgeTransactions = createAction<
  RecentBridgeTransaction[]
>('bridge/updateRecentBridgeTransactions')
