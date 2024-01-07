import { useEffect } from 'react'
import { useAppDispatch } from '@/store/hooks'
import { use_TransactionsState } from './hooks'
import { addTransaction } from './reducer'
import { useTransactionsState } from '../transactions/hooks'
import { checkTransactionsExist } from '@/utils/checkTransactionsExist'
import { PendingBridgeTransaction } from '../transactions/actions'

export default function Updater() {
  const dispatch = useAppDispatch()
  const { pendingBridgeTransactions } = useTransactionsState()
  const transactions = use_TransactionsState()

  /** Add transaction if not in _transactions store */
  useEffect(() => {
    if (checkTransactionsExist(pendingBridgeTransactions)) {
      const txnExists =
        transactions && transactions.find((tx) => !!tx.transactionHash)
      pendingBridgeTransactions.forEach((tx: PendingBridgeTransaction) => {
        if (!txnExists) {
          dispatch(
            addTransaction({
              originTxHash: tx.transactionHash,
              originValue: tx.originValue,
              originChain: tx.originChain,
              originToken: tx.originToken,
              destinationChain: tx.destinationChain,
              destinationToken: tx.destinationToken,
              bridgeModuleName: tx.bridgeModuleName,
              estimatedTime: tx.estimatedTime,
              timestamp: tx.id,
            })
          )
        }
      })
    }
  }, [pendingBridgeTransactions])

  return null
}
