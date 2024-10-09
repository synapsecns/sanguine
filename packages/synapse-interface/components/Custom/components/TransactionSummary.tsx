import _ from 'lodash'
import { useState, useEffect } from 'react'
import { useAccount } from 'wagmi'

import { Transaction } from './Transaction'
import { getTimeMinutesFromNow } from '../utils/getTimeMinutesFromNow'
import { use_TransactionsState } from '@/slices/_transactions/hooks'
import { SCROLL } from '@/constants/chains/master'

export const TransactionSummary = () => {
  const { address, isConnected } = useAccount()
  const { transactions } = use_TransactionsState()

  const hasTransactions: boolean = transactions.length > 0

  const [currentTime, setCurrentTime] = useState<number>(
    getTimeMinutesFromNow(0)
  )

  /** Update time to trigger transactions to recheck tx status */
  useEffect(() => {
    const interval = setInterval(() => {
      setCurrentTime(getTimeMinutesFromNow(0))
    }, 5000)

    return () => {
      clearInterval(interval)
    }
  }, [])

  if (hasTransactions && isConnected) {
    const sortedTransactions = _.orderBy(transactions, ['timestamp'], ['desc'])
      .filter((t) => t.destinationChain.id === SCROLL.id)
      .slice(0, 4)

    return sortedTransactions.map((transaction) => (
      <Transaction
        key={transaction.originTxHash}
        connectedAddress={address}
        originAmount={transaction.originValue}
        originTokenSymbol={transaction.originToken.symbol}
        originChainId={transaction.originChain.id}
        destinationChainId={transaction.destinationChain.id}
        originTxHash={transaction.originTxHash}
        bridgeModuleName={transaction.bridgeModuleName}
        estimatedTime={transaction.estimatedTime}
        kappa={transaction?.kappa}
        timestamp={transaction.timestamp}
        currentTime={currentTime}
        isStoredComplete={transaction.status === 'completed'}
      />
    ))
  }
}
