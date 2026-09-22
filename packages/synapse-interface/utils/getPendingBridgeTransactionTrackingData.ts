import { zeroAddress } from 'viem'

import { HYPERLIQUID } from '@/constants/chains/master'
import { isHyperliquidUsdcDeposit } from '@/utils/hyperliquid'
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
  const isHyperliquidDeposit =
    !tx.bridgeModuleName &&
    isHyperliquidUsdcDeposit(
      destinationChain?.id,
      tx.destinationToken ?? tx.originToken
    )
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
    destinationChain,
    destinationToken: tx.destinationToken ?? tx.originToken,
    bridgeModuleName: tx.bridgeModuleName ?? '',
    routerAddress: tx.routerAddress ?? zeroAddress,
    estimatedTime: tx.estimatedTime ?? 0,
    timestamp: tx.timestamp ?? tx.id,
    status: 'pending' as const,
  }
}
