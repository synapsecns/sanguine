import { useEffect } from 'react'

import { useAppDispatch } from '@/state/hooks'
import { useTransactionsState } from '@/state/slices/transactions/hooks'
import {
  updateTransactionKappa,
  completeTransaction,
} from '@/state/slices/transactions/reducer'

export const useBridgeTxUpdater = (
  kappa: string,
  originTxHash: string,
  isTxComplete: boolean
) => {
  const dispatch = useAppDispatch()
  const { transactions } = useTransactionsState()

  const storedTx = transactions.find((tx) => tx.originTxHash === originTxHash)

  /** Update tx kappa when available */
  useEffect(() => {
    if (!storedTx.kappa && kappa && originTxHash) {
      dispatch(updateTransactionKappa({ originTxHash, kappa }))
    }
  }, [dispatch, kappa])

  /** Update tx for completion in store */
  useEffect(() => {
    if (isTxComplete && originTxHash && kappa) {
      /** Check that we have not already marked tx as complete */
      if (!storedTx.isComplete) {
        dispatch(completeTransaction({ originTxHash, kappa }))
      }
    }
  }, [dispatch, isTxComplete, originTxHash, kappa, transactions])
}
