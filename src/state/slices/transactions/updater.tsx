import { useEffect } from 'react'
import { useAppDispatch } from '@/state/hooks'
import { useBridgeTransactionState } from '../bridgeTransaction/hooks'
import { useTransactionsState } from './hooks'
import { addTransaction } from './reducer'
import { isNull } from '@/utils/isNull'

export default function Updater() {
  const dispatch = useAppDispatch()
  const {
    txHash,
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
          originTxHash: txHash,
          originChainId: originChainId,
          destinationChainId: destinationChainId,
          bridgeModuleName: bridgeModuleName,
          estimatedTime: estimatedTime,
          timestamp: timestamp,
        })
      )
    }
  }, [txHash])

  return null
}
