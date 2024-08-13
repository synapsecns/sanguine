import _ from 'lodash'
import { use_TransactionsState } from '@/slices/_transactions/hooks'
import { _TransactionDetails } from '@/slices/_transactions/reducer'
import { _Transaction } from './_Transaction'
import { checkTransactionsExist } from '@/utils/checkTransactionsExist'
import { useIntervalTimer } from '@/utils/hooks/useIntervalTimer'
import { ALL_TOKENS } from '@/constants/tokens/master'
import { CHAINS_BY_ID } from '@/constants/chains'
import { useWalletState } from '@/slices/wallet/hooks'

/** TODO: Update naming once refactoring of previous Activity/Tx flow is done */
export const _Transactions = ({
  connectedAddress,
}: {
  connectedAddress: string
}) => {
  const { transactions } = use_TransactionsState()
  const { isWalletPending } = useWalletState()
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
        {sortedTransactions.slice(0, 5).map((tx: _TransactionDetails) => {
          const originChain = CHAINS_BY_ID[tx.originChain?.id]
          const originToken = _(ALL_TOKENS).find(
            (token) => token.routeSymbol === tx.originToken?.routeSymbol
          )

          const destinationChain = CHAINS_BY_ID[tx.destinationChain?.id]
          const destinationToken = _(ALL_TOKENS).find(
            (token) => token.routeSymbol === tx.destinationToken?.routeSymbol
          )

          return (
            <_Transaction
              key={tx.timestamp}
              connectedAddress={connectedAddress}
              destinationAddress={tx.destinationAddress}
              originValue={Number(tx.originValue)}
              originChain={originChain}
              originToken={originToken}
              destinationChain={destinationChain}
              destinationToken={destinationToken}
              originTxHash={tx.originTxHash}
              bridgeModuleName={tx.bridgeModuleName}
              estimatedTime={tx.estimatedTime}
              kappa={tx?.kappa}
              timestamp={tx.timestamp}
              currentTime={currentTime}
              status={tx.status}
              disabled={isWalletPending}
            />
          )
        })}
      </TransactionsContainer>
    )
  }

  return null
}

const TransactionsContainer = ({ children }) => {
  return (
    <div id="transaction-container" className="flex flex-col mt-3 space-y-3">
      {children}
    </div>
  )
}
