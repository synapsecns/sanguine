import { useState, useEffect } from 'react'

import { useTransactionsState } from '@/state/slices/transactions/hooks'
import { TransactionDetails } from '@/state/slices/transactions/reducer'
import { Transaction } from '@/components/Transaction'
import { getTimeMinutesFromNow } from '@/utils/getTimeMinutesFromNow'

export const Transactions = ({
  connectedAddress,
}: {
  connectedAddress: string
}) => {
  const { transactions } = useTransactionsState()

  const hasTransactions: boolean = transactions.length > 0

  const [currentTime, setCurrentTime] = useState<number>(
    getTimeMinutesFromNow(0)
  )

  /** Update time to trigger transactions to recheck tx status */
  useEffect(() => {
    const interval = setInterval(() => {
      setCurrentTime(getTimeMinutesFromNow(0))
    }, 5000) // 5000 milliseconds = 5 seconds

    return () => {
      clearInterval(interval) // Clear the interval when the component unmounts
    }
  }, [])

  if (hasTransactions) {
    return transactions.map((transaction: TransactionDetails) => (
      <Transaction
        key={transaction.originTxHash}
        connectedAddress={connectedAddress}
        originAmount={transaction.originAmount}
        originTokenSymbol={transaction.originTokenSymbol}
        originChainId={transaction.originChainId}
        destinationChainId={transaction.destinationChainId}
        originTxHash={transaction.originTxHash}
        bridgeModuleName={transaction.bridgeModuleName}
        estimatedTime={transaction.estimatedTime}
        kappa={transaction?.kappa}
        timestamp={transaction.timestamp}
        currentTime={currentTime}
        isStoredComplete={transaction.isComplete}
      />
    ))
  }
}
