import { createAction } from '@reduxjs/toolkit'

import { Token } from '@/utils/types'

export interface RecentBridgeTransaction {
  originChainId: number
  originChainName: string
  originToken: Token
  originValue: string
  destinationChainId: number
  destinationChainName: string
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
