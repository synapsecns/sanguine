import _ from 'lodash'
import { useEffect } from 'react'
import { useAccount } from 'wagmi'

import { useAppDispatch } from '@/store/hooks'
import { use_TransactionsState } from '@/slices/_transactions/hooks'
import { addTransaction } from '@/slices/_transactions/reducer'
import { useTransactionsState } from '@/slices/transactions/hooks'
import { checkTransactionsExist } from '@/utils/checkTransactionsExist'
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
        /** Check Transaction has been confirmed */
        const txnConfirmed =
          !_.isNull(tx.transactionHash) && !_.isUndefined(tx.transactionHash)

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
          dispatch(
            addTransaction({
              address,
              originTxHash: tx.transactionHash,
              originValue: tx.originValue,
              originChain: tx.originChain,
              originToken: tx.originToken,
              destinationChain: tx.destinationChain,
              destinationToken: tx.destinationToken,
              bridgeModuleName: tx.bridgeModuleName,
              estimatedTime: tx.estimatedTime,
              timestamp: tx.id,
              status: 'pending',
            })
          )
        }
      })
    }
  }, [pendingBridgeTransactions, transactions])

  return null
}
