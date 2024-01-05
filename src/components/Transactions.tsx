import { useState, useEffect } from 'react'
import { useTransactionsState } from '@/state/slices/transactions/hooks'
import { TransactionDetails } from '@/state/slices/transactions/reducer'
import { Transaction } from './Transaction'
import { getTimeMinutesFromNow } from '@/utils/getTimeMinutesFromNow'

/** TODO: Pull synapseSDK from context vs passing in */
export const Transactions = ({
  synapseSDK,
  connectedAddress,
}: {
  synapseSDK: any
  connectedAddress: string
}) => {
  const transactions = useTransactionsState()

  const transactionsArray: TransactionDetails[] = Object.values(transactions)

  const hasTransactions: boolean = transactionsArray.length > 0

  const [currentTime, setCurrentTime] = useState<number>(
    getTimeMinutesFromNow(0)
  )

  /** Update time to trigger transactions to recheck tx status */
  useEffect(() => {
    const interval = setInterval(() => {
      setCurrentTime(getTimeMinutesFromNow(0))
    }, 30000) // 30000 milliseconds = 30 seconds

    return () => {
      clearInterval(interval) // Clear the interval when the component unmounts
    }
  }, [])

  if (hasTransactions) {
    return transactionsArray.map((transaction: TransactionDetails) => (
      <Transaction
        synapseSDK={synapseSDK}
        connectedAddress={connectedAddress}
        originChainId={transaction.originChainId}
        destinationChainId={transaction.destinationChainId}
        originTxHash={transaction.originTxHash}
        bridgeModuleName={transaction.bridgeModuleName}
        estimatedTime={transaction.estimatedTime}
        kappa={transaction?.kappa}
        timestamp={transaction.timestamp}
        currentTime={currentTime}
        isComplete={transaction.isComplete}
      />
    ))
  }
}
