import { useEffect } from 'react'
import { isNull } from 'lodash'
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
      pendingBridgeTransactions.forEach(
        (transaction: PendingBridgeTransaction) => {
          const {
            transactionHash,
            originChain,
            destinationChain,
            bridgeModuleName,
            estimatedTime,
            timestamp,
          } = transaction
          if (!transactions[transactionHash]) {
            dispatch(
              addTransaction({
                originTxHash: transactionHash,
                originChainId: originChain.id,
                destinationChainId: destinationChain.id,
                bridgeModuleName: bridgeModuleName,
                estimatedTime: estimatedTime,
                timestamp: timestamp,
              })
            )
          }
        }
      )
    }
  }, [pendingBridgeTransactions])

  return null
}
