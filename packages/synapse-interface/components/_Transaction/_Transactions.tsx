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
  const currentTime = useIntervalTimer(5000)
  const { transactions } = use_TransactionsState()

  const sortedTransactions = _.orderBy(transactions, ['timestamp'], ['desc'])
  const hasTransactions: boolean = checkTransactionsExist(transactions)

  if (hasTransactions) {
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

  return null
}
