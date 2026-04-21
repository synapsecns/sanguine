import { zeroAddress } from 'viem'

import { HYPERLIQUID } from '@/constants/chains/master'
import { PendingBridgeTransaction } from '@/slices/transactions/actions'

export const getPendingBridgeTransactionTrackingData = (
  tx: PendingBridgeTransaction,
  address: string
) => {
  if (!tx.transactionHash) {
    return null
  }

  const isHyperliquidDeposit = tx.destinationChain?.id === HYPERLIQUID.id
  const hasTrackedTimestamp = typeof tx.timestamp === 'number'

  if (!hasTrackedTimestamp && !isHyperliquidDeposit) {
    return null
  }

  return {
    address,
    destinationAddress: tx.destinationAddress ?? null,
    originTxHash: tx.transactionHash,
    originValue: tx.originValue,
    originChain: tx.originChain,
    originToken: tx.originToken,
    destinationChain: tx.destinationChain,
    destinationToken: tx.destinationToken ?? tx.originToken,
    bridgeModuleName: tx.bridgeModuleName ?? '',
    routerAddress: tx.routerAddress ?? zeroAddress,
    estimatedTime: tx.estimatedTime ?? 0,
    timestamp: tx.timestamp ?? tx.id,
    status: 'pending' as const,
  }
}
