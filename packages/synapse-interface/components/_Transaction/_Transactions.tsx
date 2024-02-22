import _ from 'lodash'
import { use_TransactionsState } from '@/slices/_transactions/hooks'
import { _TransactionDetails } from '@/slices/_transactions/reducer'
import { _Transaction } from './_Transaction'
import { checkTransactionsExist } from '@/utils/checkTransactionsExist'
import { useIntervalTimer } from './helpers/useIntervalTimer'

/** TODO: Update naming once refactoring of previous Activity/Tx flow is done */
export const _Transactions = ({
  connectedAddress,
}: {
  connectedAddress: string
}) => {
  const { transactions } = use_TransactionsState()
  const hasTransactions: boolean = checkTransactionsExist(transactions)

  const currentTime = useIntervalTimer(5000)

  if (hasTransactions) {
    const address = connectedAddress.toLowerCase()
    const filteredTransactions = transactions.filter(
      (txn) => txn.address?.toLowerCase() === address
    )

    const sortedTransactions = _.orderBy(
      filteredTransactions,
      ['timestamp'],
      ['desc']
    )
    return (
      <TransactionsContainer>
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
            // isStoredComplete={tx.isComplete}
            // isStoredReverted={tx.isReverted}
            status={getTransactionStatus(tx.isComplete, tx.isReverted)}
          />
        ))}
      </TransactionsContainer>
    )
  }

  return null
}

const getTransactionStatus = (
  isComplete: boolean,
  isReverted: boolean
): 'pending' | 'complete' | 'reverted' => {
  if (isComplete) {
    return 'complete'
  } else if (isReverted) {
    return 'reverted'
  } else {
    return 'pending'
  }
}

const TransactionsContainer = ({ children }) => {
  return (
    <div id="transaction-container" className="flex flex-col mt-3 space-y-3">
      {children}
    </div>
  )
}
