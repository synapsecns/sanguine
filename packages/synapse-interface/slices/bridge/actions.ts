import { createAction } from '@reduxjs/toolkit'

import { Chain, Token } from '@/utils/types'

export interface PendingBridgeTransaction {
  originChain: Chain
  originToken: Token
  originValue: string
  destinationChain: Chain
  destinationToken: Token
  transactionHash?: string
  timestamp: number
}

export const addPendingBridgeTransaction =
  createAction<PendingBridgeTransaction>('bridge/addRecentBridgeTransaction')
export const updatePendingBridgeTransaction = createAction<{
  timestamp: number
  transactionHash: string
}>('bridge/updateRecentBridgeTransaction')
export const updatePendingBridgeTransactions = createAction<
  PendingBridgeTransaction[]
>('bridge/updateRecentBridgeTransactions')
