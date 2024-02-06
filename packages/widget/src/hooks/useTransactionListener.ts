import { useEffect } from 'react'

import { useAppDispatch } from '@/state/hooks'
import { useBridgeTransactionState } from '@/state/slices/bridgeTransaction/hooks'
import { useTransactionsState } from '@/state/slices/transactions/hooks'
import { addTransaction } from '@/state/slices/transactions/reducer'
import { isNull } from '@/utils/isNull'

export const useTransactionListener = () => {
  const dispatch = useAppDispatch()
  const {
    txHash,
    originAmount,
    originChainId,
    destinationChainId,
    bridgeModuleName,
    estimatedTime,
    timestamp,
  } = useBridgeTransactionState()
  const transactions = useTransactionsState()

  /** Add transaction if not in transactions store */
  useEffect(() => {
    if (isNull(txHash)) return

    if (!transactions[txHash]) {
      dispatch(
        addTransaction({
          originAmount,
          originTxHash: txHash,
          originChainId,
          destinationChainId,
          bridgeModuleName,
          estimatedTime,
          timestamp,
        })
      )
    }
  }, [txHash])

  return null
}
