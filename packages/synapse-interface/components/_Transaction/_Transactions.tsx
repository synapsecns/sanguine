import _ from 'lodash'
import { use_TransactionsState } from '@/slices/_transactions/hooks'
import { _TransactionDetails } from '@/slices/_transactions/reducer'
import { _Transaction } from './_Transaction'
import { checkTransactionsExist } from '@/utils/checkTransactionsExist'
import { useIntervalTimer } from './helpers/useIntervalTimer'

import { useTransactionsState } from '@/slices/transactions/hooks'
import { TransactionsState } from '@/slices/transactions/reducer'
import { BridgeTransaction } from '@/slices/api/generated'
import { ARBITRUM } from '@/constants/chains/master'
import { ETH } from '@/constants/tokens/bridgeable'

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
            isStoredComplete={tx.isComplete}
          />
        ))}
      </TransactionsContainer>
    )
  }

  // return null

  /*
    Show historic transactions as 'Pending' for testing
    Requires useBridgeTxUpdater(…) to be commented out of _Transaction.tsx
  */

  const {
    userHistoricalTransactions,
    isUserHistoricalTransactionsLoading,
  }: TransactionsState = useTransactionsState()

  return (
    <TransactionsContainer>
        {userHistoricalTransactions.slice(0, 5).map((tx: BridgeTransaction) => (
          <_Transaction
            key={tx.fromInfo.time}
            connectedAddress={connectedAddress}
            originValue={Number(tx.fromInfo.chainID)}
            originChain={ARBITRUM}
            originToken={ETH}
            destinationChain={ARBITRUM}
            destinationToken={ETH}
            originTxHash={Math.random().toString()}
            bridgeModuleName={'Synapse Bridge'}
            estimatedTime={6000}
            kappa={tx?.kappa}
            timestamp={tx.fromInfo.time + 15000}
            currentTime={currentTime}
            isStoredComplete={false}
          />
        ))}
      </TransactionsContainer>
  )
}

const TransactionsContainer = ({ children }) => {
  return (
    <div id="transaction-container" className="flex flex-col mt-3">
      {children}
    </div>
  )
}
