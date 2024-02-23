import React, { useMemo } from 'react'
import _ from 'lodash'
import Fuse from 'fuse.js'
import { useAccount, Address } from 'wagmi'
import { useTransactionsState } from '@/slices/transactions/hooks'
import { usePortfolioState } from '@/slices/portfolio/hooks'
import { BridgeTransaction } from '@/slices/api/generated'
import { CHAINS_BY_ID } from '@/constants/chains'
import type { Chain } from '@/utils/types'
import { tokenAddressToToken } from '@/constants/tokens'
import { TransactionsState } from '@/slices/transactions/reducer'
import { PortfolioState } from '@/slices/portfolio/reducer'
import { Transaction, TransactionType } from './Transaction/Transaction'
import { UserExplorerLink } from './Transaction/components/TransactionExplorerLink'
import { NoSearchResultsContent } from './components/NoSearchResultContent'
import { checkTransactionsExist } from '@/utils/checkTransactionsExist'
import { AirdropRewards } from '../Activity/AirdropRewards'

export const Activity = ({ visibility }: { visibility: boolean }) => {
  const { address } = useAccount()
  const {
    userHistoricalTransactions,
    isUserHistoricalTransactionsLoading,
  }: TransactionsState = useTransactionsState()
  const { searchInput, searchedBalances }: PortfolioState = usePortfolioState()

  const isLoading: boolean = isUserHistoricalTransactionsLoading
  const isSearchInputActive = Boolean(searchInput.length > 0)
  const isMasqueradeActive = Object.keys(searchedBalances).length > 0
  const masqueradeAddress = Object.keys(searchedBalances)[0] as Address

  const filteredHistoricalTransactions = filterTransactionsBySearch(
    userHistoricalTransactions,
    isLoading,
    searchInput,
    isSearchInputActive,
    masqueradeAddress
  )

  const hasHistoricalTransactions = !_.isEmpty(userHistoricalTransactions)
  const hasFilteredSearchResults = !_.isEmpty(filteredHistoricalTransactions)

  const viewingAddress: string | null = useMemo(() => {
    if (isMasqueradeActive) {
      return masqueradeAddress
    } else if (address) {
      return address
    } else return null
  }, [isMasqueradeActive, masqueradeAddress, address])

  return (
    <div
      data-test-id="activity"
      className={`${visibility ? 'block' : 'hidden'}`}
    >
      {!viewingAddress && (
        <div className="text-secondary">
          Your pending and recent transactions will appear here.
        </div>
      )}

      {viewingAddress && isLoading && (
        <div className="text-secondary">Loading activity...</div>
      )}

      {viewingAddress && !isLoading && !hasHistoricalTransactions && (
        <div className="text-secondary">
          No transactions in last 30 days.
          <UserExplorerLink connectedAddress={viewingAddress} />
        </div>
      )}

      {viewingAddress && !isLoading && hasHistoricalTransactions && (
        <ActivitySection title="Recent">
          {/* TODO: Update AirdropRewards to work for masquerade */}
          <AirdropRewards />

          {userHistoricalTransactions &&
            filteredHistoricalTransactions
              .slice(0, isSearchInputActive ? 100 : 6)
              .map((transaction: BridgeTransaction) =>
                renderTransaction(transaction, address)
              )}
          {isSearchInputActive && !hasFilteredSearchResults && (
            <NoSearchResultsContent searchStr={searchInput} />
          )}
          <UserExplorerLink connectedAddress={viewingAddress} />
        </ActivitySection>
      )}
    </div>
  )
}

const renderTransaction = (
  transaction: BridgeTransaction,
  viewingAddress: Address
) => {
  return (
    <Transaction
      key={transaction?.kappa}
      connectedAddress={viewingAddress}
      destinationAddress={transaction?.fromInfo?.address as Address}
      startedTimestamp={transaction?.fromInfo?.time}
      completedTimestamp={transaction?.toInfo?.time}
      transactionHash={transaction?.fromInfo?.txnHash}
      kappa={transaction?.kappa}
      isCompleted={true}
      transactionType={TransactionType.HISTORICAL}
      originValue={transaction?.fromInfo?.value}
      destinationValue={transaction?.toInfo?.value}
      originChain={CHAINS_BY_ID[transaction?.fromInfo?.chainID]}
      originToken={tokenAddressToToken(
        transaction?.fromInfo?.chainID,
        transaction?.fromInfo?.tokenAddress
      )}
      destinationChain={CHAINS_BY_ID[transaction?.toInfo?.chainID]}
      destinationToken={tokenAddressToToken(
        transaction?.toInfo?.chainID,
        transaction?.toInfo?.tokenAddress
      )}
    />
  )
}

const filterTransactionsBySearch = (
  transactions: BridgeTransaction[],
  isLoading: boolean,
  searchInput: string,
  isSearchActive: boolean,
  masqueradeAddress: string
) => {
  let searchFiltered = []

  const fuseOptions = {
    includeScore: true,
    threshold: 0.33,
    distance: 20,
    keys: [
      'originChain.name',
      'originToken.symbol',
      'destinationChain.name',
      'destinationToken.symbol',
      'originTokenAddresses',
      'destinationTokenAddresses',
      'fromInfo.txnHash',
      'toInfo.txnHash',
    ],
  }

  if (!isLoading && checkTransactionsExist(transactions)) {
    const formatted = transactions.map((transaction) => {
      const originToken = tokenAddressToToken(
        transaction?.fromInfo?.chainID,
        transaction?.fromInfo?.tokenAddress
      )
      const destinationToken = tokenAddressToToken(
        transaction?.toInfo?.chainID,
        transaction?.toInfo?.tokenAddress
      )
      return {
        ...transaction,
        originChain: CHAINS_BY_ID[transaction?.fromInfo?.chainID] as Chain,
        originToken: originToken,
        originTokenAddresses:
          originToken && Object.values(originToken?.addresses),
        destinationChain: CHAINS_BY_ID[transaction?.toInfo?.chainID] as Chain,
        destinationToken: destinationToken,
        destinationTokenAddresses:
          destinationToken && Object.values(destinationToken?.addresses),
      }
    })

    const fuse = new Fuse(formatted, fuseOptions)
    if (isSearchActive) {
      searchFiltered = fuse.search(searchInput).map((i) => i.item)
    }

    const inputIsMasqueradeAddress = searchInput === masqueradeAddress
    return isSearchActive && !inputIsMasqueradeAddress
      ? searchFiltered
      : transactions
  }

  return searchFiltered
}

export const ActivitySection = ({
  title,
  children,
  twClassName,
}: {
  title: string
  children?: React.ReactNode
  twClassName?: string
}) => {
  return (
    <div data-test-id="activity-section" className={twClassName}>
      <h3 className="mb-2 text-xl text-white">{title}</h3>
      {children}
    </div>
  )
}
