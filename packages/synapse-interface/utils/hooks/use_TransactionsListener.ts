import { useEffect } from 'react'
import { useAccount } from 'wagmi'

import { useAppDispatch } from '@/store/hooks'
import { use_TransactionsState } from '@/slices/_transactions/hooks'
import { addTransaction } from '@/slices/_transactions/reducer'
import { useTransactionsState } from '@/slices/transactions/hooks'
import { checkTransactionsExist } from '@/utils/checkTransactionsExist'
import { getPendingBridgeTransactionTrackingData } from '@/utils/getPendingBridgeTransactionTrackingData'
import {
  PendingBridgeTransaction,
  removePendingBridgeTransaction,
} from '@/slices/transactions/actions'

export const use_TransactionsListener = () => {
  const dispatch = useAppDispatch()
  const { pendingBridgeTransactions } = useTransactionsState()
  const { transactions } = use_TransactionsState()
  const { address } = useAccount()

  /** Add transaction if not in _transactions store */
  useEffect(() => {
    if (checkTransactionsExist(pendingBridgeTransactions)) {
      pendingBridgeTransactions.forEach((tx: PendingBridgeTransaction) => {
        const trackedTransaction = address
          ? getPendingBridgeTransactionTrackingData(tx, address)
          : null
        const txnConfirmed = trackedTransaction !== null

        /** Check Transaction is already stored */
        const txnExists =
          transactions &&
          transactions.some(
            (storedTx) => tx.transactionHash === storedTx.originTxHash
          )

        /** Remove pendingBridgeTransaction once stored in transactions */
        if (txnExists) {
          dispatch(removePendingBridgeTransaction(tx.id))
        }

        if (txnConfirmed && !txnExists) {
          dispatch(addTransaction(trackedTransaction))
        }
      })
    }
  }, [pendingBridgeTransactions, transactions])

  return null
}
