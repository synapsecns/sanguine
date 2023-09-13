import { useState, useEffect, useMemo } from 'react'
import { useAccount, Address } from 'wagmi'
import { Chain, Token } from '@/utils/types'
import { useBridgeState } from '@/slices/bridge/hooks'
import { useTransactionsState } from '@/slices/transactions/hooks'
import { BridgeState } from '@/slices/bridge/reducer'
import { TransactionsState } from '@/slices/transactions/reducer'
import { PendingBridgeTransaction } from '@/slices/bridge/actions'
import { BridgeTransaction } from '@/slices/api/generated'
import { getTimeMinutesBeforeNow } from '@/utils/time'
import { TransactionType } from './Transaction'
import { tokenAddressToToken } from '@/constants/tokens'
import { CHAINS_BY_ID } from '@/constants/chains'
import { PendingTransaction } from './PendingTransaction'

export const MostRecentTransaction = () => {
  const { address } = useAccount()
  const { pendingBridgeTransactions }: BridgeState = useBridgeState()
  const {
    userHistoricalTransactions,
    isUserHistoricalTransactionsLoading,
    isUserPendingTransactionsLoading,
    seenHistoricalTransactions,
    pendingAwaitingCompletionTransactions,
  }: TransactionsState = useTransactionsState()

  const [currentTime, setCurrentTime] = useState<number>(
    getTimeMinutesBeforeNow(0)
  )

  useEffect(() => {
    const interval = setInterval(() => {
      setCurrentTime(getTimeMinutesBeforeNow(0))
    }, 60000)

    return () => clearInterval(interval)
  }, [])

  const lastPendingBridgeTransaction: PendingBridgeTransaction = useMemo(() => {
    return pendingBridgeTransactions?.[0]
  }, [pendingBridgeTransactions])

  const lastPendingTransaction: BridgeTransaction = useMemo(() => {
    return pendingAwaitingCompletionTransactions?.[0]
  }, [pendingAwaitingCompletionTransactions])

  const lastHistoricalTransaction: BridgeTransaction = useMemo(() => {
    return userHistoricalTransactions?.[0]
  }, [userHistoricalTransactions])

  const recentMinutesInUnix: number = 15 * 60

  const isLastHistoricalTransactionRecent: boolean = useMemo(
    () =>
      currentTime - lastHistoricalTransaction?.toInfo?.time <
      recentMinutesInUnix,
    [currentTime]
  )

  const seenLastHistoricalTransaction: boolean = useMemo(() => {
    if (!seenHistoricalTransactions || !userHistoricalTransactions) {
      return false
    } else {
      return seenHistoricalTransactions.some(
        (transaction: BridgeTransaction) =>
          transaction === (lastHistoricalTransaction as BridgeTransaction)
      )
    }
  }, [seenHistoricalTransactions, lastHistoricalTransaction])

  let transaction

  if (isUserHistoricalTransactionsLoading || isUserPendingTransactionsLoading) {
    return null
  }

  if (lastPendingBridgeTransaction) {
    transaction = lastPendingBridgeTransaction as PendingBridgeTransaction
    return (
      <div data-test-id="most-recent-transaction" className="mt-3">
        <PendingTransaction
          connectedAddress={address as Address}
          originChain={transaction.originChain as Chain}
          originToken={transaction.originToken as Token}
          originValue={Number(transaction.originValue)}
          destinationChain={transaction.destinationChain as Chain}
          destinationToken={transaction.destinationToken as Token}
          startedTimestamp={transaction.id ?? transaction.startedTimestamp}
          transactionHash={transaction.transactionHash as string}
          isSubmitted={transaction.isSubmitted as boolean}
          transactionType={TransactionType.PENDING}
        />
      </div>
    )
  }

  if (lastPendingTransaction) {
    transaction = lastPendingTransaction as BridgeTransaction
    return (
      <div data-test-id="most-recent-transaction" className="mt-3">
        <PendingTransaction
          connectedAddress={address as Address}
          startedTimestamp={transaction?.fromInfo?.time as number}
          transactionHash={transaction?.fromInfo?.txnHash as string}
          transactionType={TransactionType.PENDING}
          originValue={transaction?.fromInfo?.formattedValue as number}
          originChain={CHAINS_BY_ID[transaction?.fromInfo?.chainID] as Chain}
          destinationChain={
            CHAINS_BY_ID[transaction?.fromInfo?.destinationChainID] as Chain
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
          destinationAddress={transaction?.fromInfo?.address as Address}
          isSubmitted={transaction?.fromInfo?.txnHash ? true : false}
          isCompleted={transaction?.toInfo?.time ? true : false}
          kappa={transaction?.kappa}
        />
      </div>
    )
  }

  if (
    lastHistoricalTransaction &&
    isLastHistoricalTransactionRecent &&
    !seenLastHistoricalTransaction
  ) {
    transaction = lastHistoricalTransaction as BridgeTransaction
    return (
      <div data-test-id="most-recent-transaction" className="mt-3">
        <PendingTransaction
          connectedAddress={address as Address}
          destinationAddress={transaction?.fromInfo?.address as Address}
          startedTimestamp={transaction?.fromInfo?.time as number}
          completedTimestamp={transaction?.toInfo?.time as number}
          transactionHash={transaction?.fromInfo?.txnHash as string}
          kappa={transaction?.kappa as string}
          transactionType={TransactionType.PENDING}
          originValue={transaction?.fromInfo?.formattedValue as number}
          destinationValue={transaction?.toInfo?.formattedValue as number}
          originChain={CHAINS_BY_ID[transaction?.fromInfo?.chainID] as Chain}
          destinationChain={
            CHAINS_BY_ID[transaction?.fromInfo?.destinationChainID] as Chain
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
          isSubmitted={transaction?.fromInfo?.txnHash ? true : false}
          isCompleted={true}
        />
      </div>
    )
  }
}
