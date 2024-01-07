import React, { useMemo } from 'react'
import Fuse from 'fuse.js'
import { useAccount, Address } from 'wagmi'
import { useTransactionsState } from '@/slices/transactions/hooks'
import { usePortfolioState } from '@/slices/portfolio/hooks'
import { BridgeTransaction } from '@/slices/api/generated'
import { CHAINS_BY_ID } from '@/constants/chains'
import { Chain, Token } from '@/utils/types'
import { tokenAddressToToken } from '@/constants/tokens'
import { TransactionsState } from '@/slices/transactions/reducer'
import { PortfolioState } from '@/slices/portfolio/reducer'
import { PendingBridgeTransaction } from '@/slices/transactions/actions'
import { Transaction, TransactionType } from './Transaction/Transaction'
import { PendingTransaction } from './Transaction/PendingTransaction'
import { UserExplorerLink } from './Transaction/components/TransactionExplorerLink'
import { NoSearchResultsContent } from './PortfolioContent/PortfolioContent'
import { checkTransactionsExist } from '@/utils/checkTransactionsExist'

export const Activity = ({ visibility }: { visibility: boolean }) => {
  const { address } = useAccount()
  const {
    userHistoricalTransactions,
    isUserHistoricalTransactionsLoading,
    isUserPendingTransactionsLoading,
    pendingAwaitingCompletionTransactions,
    fallbackQueryPendingTransactions,
    fallbackQueryHistoricalTransactions,
    pendingBridgeTransactions,
  }: TransactionsState = useTransactionsState()
  const { searchInput, searchedBalancesAndAllowances }: PortfolioState =
    usePortfolioState()

  const pendingAwaitingCompletionTransactionsWithFallback: BridgeTransaction[] =
    useMemo(() => {
      let transactions: BridgeTransaction[] = []

      if (checkTransactionsExist(pendingAwaitingCompletionTransactions)) {
        transactions = [...pendingAwaitingCompletionTransactions]
      }

      if (checkTransactionsExist(fallbackQueryPendingTransactions)) {
        const fallbackTransactions = [...fallbackQueryPendingTransactions]
        const mergedTransactions = [...transactions, ...fallbackTransactions]

        const uniqueMergedTransactions = Array.from(
          new Set(mergedTransactions.map((transaction) => transaction?.kappa))
        ).map((kappa) =>
          mergedTransactions.find((item) => item?.kappa === kappa)
        )
        return uniqueMergedTransactions
      }

      return transactions
    }, [
      pendingAwaitingCompletionTransactions,
      fallbackQueryPendingTransactions,
    ])

  const hasPendingTransactions: boolean = useMemo(() => {
    if (checkTransactionsExist(pendingAwaitingCompletionTransactions)) {
      return true
    }
    if (checkTransactionsExist(pendingBridgeTransactions)) {
      return true
    }
    if (
      checkTransactionsExist(pendingAwaitingCompletionTransactionsWithFallback)
    ) {
      return true
    }
    return false
  }, [
    pendingBridgeTransactions,
    pendingAwaitingCompletionTransactions,
    pendingAwaitingCompletionTransactionsWithFallback,
  ])

  const hasHistoricalTransactions: boolean = useMemo(
    () => checkTransactionsExist(userHistoricalTransactions),
    [userHistoricalTransactions]
  )

  const hasNoTransactions: boolean = useMemo(() => {
    return !hasPendingTransactions && !hasHistoricalTransactions
  }, [hasPendingTransactions, hasHistoricalTransactions, address])

  const isLoading: boolean =
    isUserHistoricalTransactionsLoading && isUserPendingTransactionsLoading

  const searchInputActive: boolean = searchInput.length > 0

  const masqueradeActive: boolean = useMemo(() => {
    return Object.keys(searchedBalancesAndAllowances).length > 0
  }, [searchedBalancesAndAllowances])

  const masqueradeAddress: Address = useMemo(() => {
    return Object.keys(searchedBalancesAndAllowances)[0] as Address
  }, [searchedBalancesAndAllowances])

  const filteredHistoricalTransactionsBySearchInput: BridgeTransaction[] =
    useMemo(() => {
      let searchFiltered: BridgeTransaction[] = []
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

      if (
        !isUserHistoricalTransactionsLoading &&
        checkTransactionsExist(userHistoricalTransactions)
      ) {
        const formatted: BridgeTransaction[] = userHistoricalTransactions.map(
          (transaction: BridgeTransaction) => {
            const originToken: Token = tokenAddressToToken(
              transaction?.fromInfo?.chainID,
              transaction?.fromInfo?.tokenAddress
            )
            const destinationToken: Token = tokenAddressToToken(
              transaction?.toInfo?.chainID,
              transaction?.toInfo?.tokenAddress
            )
            return {
              ...transaction,
              originChain: CHAINS_BY_ID[
                transaction?.fromInfo?.chainID
              ] as Chain,
              originToken: originToken,
              originTokenAddresses:
                originToken && Object.values(originToken?.addresses),
              destinationChain: CHAINS_BY_ID[
                transaction?.toInfo?.chainID
              ] as Chain,
              destinationToken: destinationToken,
              destinationTokenAddresses:
                destinationToken && Object.values(destinationToken?.addresses),
            }
          }
        )
        const fuse = new Fuse(formatted, fuseOptions)
        if (searchInputActive) {
          searchFiltered = fuse
            .search(searchInput)
            .map((i: Fuse.FuseResult<BridgeTransaction>) => i.item)
        }
        const inputIsMasqueradeAddress: boolean =
          searchInput === masqueradeAddress

        return searchInputActive && !inputIsMasqueradeAddress
          ? searchFiltered
          : userHistoricalTransactions
      }
    }, [
      searchInput,
      masqueradeAddress,
      searchInputActive,
      userHistoricalTransactions,
      isUserHistoricalTransactionsLoading,
    ])

  const filteredHistoricalTransactionsBySearchInputWithFallback =
    useMemo(() => {
      let transactions: BridgeTransaction[] = []

      if (checkTransactionsExist(filteredHistoricalTransactionsBySearchInput)) {
        transactions = filteredHistoricalTransactionsBySearchInput
      }

      if (checkTransactionsExist(fallbackQueryHistoricalTransactions)) {
        const fallbackTransactions = [...fallbackQueryHistoricalTransactions]
        const mergedTransactions = [...fallbackTransactions, ...transactions]

        const uniqueMergedTransactions = Array.from(
          new Set(mergedTransactions.map((transaction) => transaction?.kappa))
        ).map((kappa) =>
          mergedTransactions.find((item) => item?.kappa === kappa)
        )
        return uniqueMergedTransactions
      }

      return transactions
    }, [
      filteredHistoricalTransactionsBySearchInput,
      fallbackQueryHistoricalTransactions,
    ])

  const hasFilteredSearchResults: boolean = useMemo(() => {
    if (filteredHistoricalTransactionsBySearchInputWithFallback) {
      return filteredHistoricalTransactionsBySearchInputWithFallback.length > 0
    } else {
      return false
    }
  }, [filteredHistoricalTransactionsBySearchInputWithFallback])

  const viewingAddress: string | null = useMemo(() => {
    if (masqueradeActive) {
      return masqueradeAddress
    } else if (address) {
      return address
    } else return null
  }, [masqueradeActive, masqueradeAddress, address])

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

      {viewingAddress && !isLoading && hasNoTransactions && (
        <div className="text-secondary">
          No transactions in last 30 days.
          <UserExplorerLink connectedAddress={viewingAddress} />
        </div>
      )}

      {/* {viewingAddress && !isLoading && hasPendingTransactions && (
        <ActivitySection title="Pending" twClassName="flex flex-col mb-5">
          {pendingAwaitingCompletionTransactionsWithFallback &&
            pendingAwaitingCompletionTransactionsWithFallback.map(
              (transaction: BridgeTransaction, key: number) => (
                <PendingTransaction
                  key={key}
                  connectedAddress={viewingAddress as Address}
                  destinationAddress={transaction?.fromInfo?.address as Address}
                  startedTimestamp={transaction?.fromInfo?.time}
                  transactionHash={transaction?.fromInfo?.txnHash}
                  eventType={transaction?.fromInfo?.eventType}
                  formattedEventType={transaction?.fromInfo?.formattedEventType}
                  isSubmitted={transaction?.fromInfo?.txnHash ? true : false}
                  isCompleted={transaction?.toInfo?.time ? true : false}
                  kappa={transaction?.kappa}
                  transactionType={TransactionType.PENDING}
                  originValue={transaction?.fromInfo?.value}
                  destinationValue={transaction?.toInfo?.value}
                  originChain={
                    CHAINS_BY_ID[transaction?.fromInfo?.chainID] as Chain
                  }
                  destinationChain={
                    CHAINS_BY_ID[
                      transaction?.fromInfo?.destinationChainID
                    ] as Chain
                  }
                  originToken={
                    tokenAddressToToken(
                      transaction?.fromInfo?.chainID,
                      transaction?.fromInfo?.tokenAddress
                    ) as Token
                  }
                  destinationToken={
                    tokenAddressToToken(
                      transaction?.toInfo?.chainID,
                      transaction?.toInfo?.tokenAddress
                    ) as Token
                  }
                />
              )
            )}
          <PendingTransactionAwaitingIndexing />
        </ActivitySection>
      )} */}

      {viewingAddress && !isLoading && hasHistoricalTransactions && (
        <ActivitySection title="Recent">
          {userHistoricalTransactions &&
            filteredHistoricalTransactionsBySearchInputWithFallback
              .slice(0, searchInputActive ? 100 : 6)
              .map((transaction: BridgeTransaction) => (
                <Transaction
                  key={transaction?.kappa}
                  connectedAddress={viewingAddress as Address}
                  destinationAddress={transaction?.fromInfo?.address as Address}
                  startedTimestamp={transaction?.fromInfo?.time}
                  completedTimestamp={transaction?.toInfo?.time}
                  transactionHash={transaction?.fromInfo?.txnHash}
                  kappa={transaction?.kappa}
                  isCompleted={true}
                  transactionType={TransactionType.HISTORICAL}
                  originValue={transaction?.fromInfo?.value}
                  destinationValue={transaction?.toInfo?.value}
                  originChain={
                    CHAINS_BY_ID[transaction?.fromInfo?.chainID] as Chain
                  }
                  originToken={
                    tokenAddressToToken(
                      transaction?.fromInfo?.chainID,
                      transaction?.fromInfo?.tokenAddress
                    ) as Token
                  }
                  destinationChain={
                    CHAINS_BY_ID[
                      transaction?.fromInfo?.destinationChainID
                    ] as Chain
                  }
                  destinationToken={
                    tokenAddressToToken(
                      transaction?.toInfo?.chainID,
                      transaction?.toInfo?.tokenAddress
                    ) as Token
                  }
                />
              ))}
          {searchInputActive && !hasFilteredSearchResults && (
            <NoSearchResultsContent searchStr={searchInput} />
          )}
          <UserExplorerLink connectedAddress={viewingAddress} />
        </ActivitySection>
      )}
    </div>
  )
}

export const PendingTransactionAwaitingIndexing = () => {
  const { address } = useAccount()
  const { pendingBridgeTransactions }: TransactionsState =
    useTransactionsState()
  return (
    <>
      {pendingBridgeTransactions.map(
        (transaction: PendingBridgeTransaction, key: number) => (
          <PendingTransaction
            key={key}
            connectedAddress={address as Address}
            originChain={transaction.originChain as Chain}
            originToken={transaction.originToken as Token}
            originValue={Number(transaction.originValue)}
            destinationChain={transaction.destinationChain as Chain}
            destinationToken={transaction.destinationToken as Token}
            estimatedDuration={transaction.estimatedTime}
            transactionHash={transaction.transactionHash}
            bridgeModuleName={transaction.bridgeModuleName}
            isSubmitted={transaction.isSubmitted as boolean}
            startedTimestamp={transaction.timestamp as number}
            transactionType={TransactionType.PENDING}
          />
        )
      )}
    </>
  )
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
