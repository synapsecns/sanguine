import { useState, useEffect } from 'react'
import { useBridgeTransactionsState } from '@/slices/bridgeTransactions/hooks'
import { BridgeTransactionDetails } from '@/slices/bridgeTransactions/reducer'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import { _Transaction as Transaction } from './_Transaction'
import { getTimeMinutesFromNow } from '@/utils/time'

/** TODO: Pull synapseSDK from context vs passing in */
export const Transactions = ({
  connectedAddress,
}: {
  connectedAddress: string
}) => {
  const { synapseSDK } = useSynapseContext()
  const transactions = useBridgeTransactionsState()

  const transactionsArray: BridgeTransactionDetails[] =
    Object.values(transactions)

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
    return transactionsArray.map((transaction: BridgeTransactionDetails) => (
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
