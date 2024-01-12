import _ from 'lodash'
import { useState, useEffect, useMemo } from 'react'
import { use_TransactionsState } from '@/slices/_transactions/hooks'
import { _TransactionDetails } from '@/slices/_transactions/reducer'
import { _Transaction } from './_Transaction'
import { getTimeMinutesFromNow } from '@/utils/time'
import { checkTransactionsExist } from '@/utils/checkTransactionsExist'

/** TODO: Update naming once refactoring of previous Activity/Tx flow is done */
export const _Transactions = ({
  connectedAddress,
}: {
  connectedAddress: string
}) => {
  const { transactions } = use_TransactionsState()

  const hasTransactions: boolean = checkTransactionsExist(transactions)

  const [currentTime, setCurrentTime] = useState<number>(
    getTimeMinutesFromNow(0)
  )

  /** Update time to trigger transactions to recheck tx status */
  useEffect(() => {
    const interval = setInterval(() => {
      let newCurrentTime = getTimeMinutesFromNow(0)
      setCurrentTime(newCurrentTime)
    }, 5000) // 5000 milliseconds = 5 seconds

    return () => {
      clearInterval(interval) // Clear the interval when the component unmounts
    }
  }, [])

  if (hasTransactions) {
    const sortedTransactions = _.orderBy(transactions, ['timestamp'], ['desc'])
    return (
      <div className="flex flex-col mt-3">
        {sortedTransactions.slice(0, 5).map((tx: _TransactionDetails) => (
          <_Transaction
            key={tx.timestamp}
            connectedAddress={connectedAddress}
            originValue={Number(tx.originValue)}
            originChain={tx.originChain}
            originToken={tx.originToken}
            destinationChain={tx.destinationChain}
            destinationToken={tx.destinationToken}
            originTxHash={tx.originTxHash}
            bridgeModuleName={tx.bridgeModuleName}
            estimatedTime={tx.estimatedTime}
            kappa={tx?.kappa}
            timestamp={tx.timestamp}
            currentTime={currentTime}
            isStoredComplete={tx.isComplete}
          />
        ))}
      </div>
    )
  }
}
