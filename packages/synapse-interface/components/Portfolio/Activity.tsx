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
import { PendingBridgeTransaction } from '@/slices/bridge/actions'
import { BridgeState } from '@/slices/bridge/reducer'
import { useBridgeState } from '@/slices/bridge/hooks'
import { Transaction, TransactionType } from './Transaction/Transaction'
import { PendingTransaction } from './Transaction/PendingTransaction'
import { UserExplorerLink } from './Transaction/components/TransactionExplorerLink'

function checkTransactionsExist(
  transactions: any[] | undefined | null
): boolean {
  const exists: boolean =
    transactions && Array.isArray(transactions) && transactions.length > 0
  return exists
}

export const Activity = ({ visibility }: { visibility: boolean }) => {
  const { address } = useAccount()
  const {
    userHistoricalTransactions,
    isUserHistoricalTransactionsLoading,
    isUserPendingTransactionsLoading,
    pendingAwaitingCompletionTransactions,
  }: TransactionsState = useTransactionsState()
  const { pendingBridgeTransactions }: BridgeState = useBridgeState()
  const { searchInput, searchedBalancesAndAllowances }: PortfolioState =
    usePortfolioState()

  const hasPendingTransactions: boolean = useMemo(() => {
    if (checkTransactionsExist(pendingAwaitingCompletionTransactions)) {
      return true
    }
    if (checkTransactionsExist(pendingBridgeTransactions)) {
      return true
    }
    return false
  }, [pendingBridgeTransactions, pendingAwaitingCompletionTransactions])

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

      {viewingAddress && !isLoading && hasPendingTransactions && (
        <ActivitySection title="Pending" twClassName="flex flex-col mb-5">
          {pendingAwaitingCompletionTransactions &&
            pendingAwaitingCompletionTransactions.map(
              (transaction: BridgeTransaction) => (
                <PendingTransaction
                  connectedAddress={viewingAddress as Address}
                  destinationAddress={transaction?.fromInfo?.address as Address}
                  startedTimestamp={transaction?.fromInfo?.time as number}
                  transactionHash={transaction?.fromInfo?.txnHash as string}
                  eventType={transaction?.fromInfo?.eventType as number}
                  isSubmitted={transaction?.fromInfo?.txnHash ? true : false}
                  isCompleted={transaction?.toInfo?.time ? true : false}
                  kappa={transaction?.kappa as string}
                  transactionType={TransactionType.PENDING}
                  originValue={transaction?.fromInfo?.formattedValue as number}
                  destinationValue={
                    transaction?.toInfo?.formattedValue as number
                  }
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
      )}

      {viewingAddress && !isLoading && hasHistoricalTransactions && (
        <ActivitySection title="Recent">
          {userHistoricalTransactions &&
            filteredHistoricalTransactionsBySearchInput
              .slice(0, searchInputActive ? 100 : 6)
              .map((transaction: BridgeTransaction) => (
                <Transaction
                  key={transaction.kappa}
                  connectedAddress={viewingAddress as Address}
                  destinationAddress={transaction?.fromInfo?.address as Address}
                  startedTimestamp={transaction?.fromInfo?.time as number}
                  completedTimestamp={transaction?.toInfo?.time as number}
                  transactionHash={transaction?.fromInfo?.txnHash as string}
                  kappa={transaction?.kappa}
                  isCompleted={true}
                  transactionType={TransactionType.HISTORICAL}
                  originValue={transaction?.fromInfo?.formattedValue as number}
                  destinationValue={
                    transaction?.toInfo?.formattedValue as number
                  }
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
          <UserExplorerLink connectedAddress={viewingAddress} />
        </ActivitySection>
      )}
    </div>
  )
}

export const PendingTransactionAwaitingIndexing = () => {
  const { address } = useAccount()
  const { pendingBridgeTransactions }: BridgeState = useBridgeState()
  return (
    <>
      {pendingBridgeTransactions.map(
        (transaction: PendingBridgeTransaction) => (
          <PendingTransaction
            connectedAddress={address as Address}
            originChain={transaction.originChain as Chain}
            originToken={transaction.originToken as Token}
            originValue={Number(transaction.originValue)}
            destinationChain={transaction.destinationChain as Chain}
            destinationToken={transaction.destinationToken as Token}
            transactionHash={transaction.transactionHash}
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
