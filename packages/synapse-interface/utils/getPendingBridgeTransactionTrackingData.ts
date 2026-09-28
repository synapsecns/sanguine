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

  // Resolve saved Hyperliquid activity by its stable symbol after an ID change.
  const destinationChain =
    tx.destinationChain?.chainSymbol === HYPERLIQUID.chainSymbol
      ? HYPERLIQUID
      : tx.destinationChain
  const hasTrackedTimestamp = typeof tx.timestamp === 'number'

  if (!hasTrackedTimestamp) {
    return null
  }

  return {
    address,
    destinationAddress: tx.destinationAddress ?? null,
    originTxHash: tx.transactionHash,
    originValue: tx.originValue,
    originChain: tx.originChain,
    originToken: tx.originToken,
    destinationChain,
    destinationToken: tx.destinationToken ?? tx.originToken,
    bridgeModuleName: tx.bridgeModuleName ?? '',
    routerAddress: tx.routerAddress ?? zeroAddress,
    estimatedTime: tx.estimatedTime ?? 0,
    timestamp: tx.timestamp,
    status: 'pending' as const,
  }
}
