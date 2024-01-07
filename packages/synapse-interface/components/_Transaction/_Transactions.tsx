import { useState, useEffect } from 'react'
import { use_TransactionsState } from '@/slices/_transactions/hooks'
import { _TransactionDetails } from '@/slices/_transactions/reducer'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import { _Transaction } from './_Transaction'
import { getTimeMinutesFromNow } from '@/utils/time'

/** TODO: Update naming once refactoring of previous Activity/Tx flow is done */
export const _Transactions = ({
  connectedAddress,
}: {
  connectedAddress: string
}) => {
  const { synapseSDK } = useSynapseContext()
  const transactions = use_TransactionsState()

  const transactionsArray: _TransactionDetails[] = Object.values(transactions)

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
    return (
      <div className="mt-3">
        {transactionsArray.map((tx: _TransactionDetails) => (
          <_Transaction
            synapseSDK={synapseSDK}
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
            isComplete={tx.isComplete}
          />
        ))}
      </div>
    )
  }
}
