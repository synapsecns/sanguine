import { useState, useEffect, useMemo } from 'react'
import { Address } from 'viem'
import { useAppDispatch } from '@/store/hooks'
import {
  useLazyGetDestinationBridgeTxFallbackQuery,
  BridgeTransaction,
  BridgeType,
} from '@/slices/api/generated'
import { useTransactionsState } from '@/slices/transactions/hooks'
import { TransactionsState } from '@/slices/transactions/reducer'
import {
  addFallbackQueryHistoricalTransaction,
  addFallbackQueryPendingTransaction,
  removeFallbackQueryPendingTransaction,
  updateFallbackQueryPendingTransaction,
} from '@/slices/transactions/actions'

interface FallbackBridgeDestinationQueryProps {
  chainId?: number
  address?: Address
  kappa?: string
  timestamp?: number
  bridgeType?: BridgeType
}

interface useFallbackBridgeDestinationQueryProps
  extends FallbackBridgeDestinationQueryProps {
  useFallback: boolean
}

export const useFallbackBridgeDestinationQuery = ({
  chainId,
  address,
  kappa,
  timestamp,
  bridgeType,
  useFallback,
}: useFallbackBridgeDestinationQueryProps) => {
  const dispatch = useAppDispatch()

  const {
    fallbackQueryPendingTransactions,
    pendingAwaitingCompletionTransactions,
  }: TransactionsState = useTransactionsState()

  const [
    fetchFallbackBridgeDestinationQuery,
    fetchedFallbackQuery,
    lastFetchedQueryParams,
  ] = useLazyGetDestinationBridgeTxFallbackQuery({ pollingInterval: 30000 })

  const validQueryParams: FallbackBridgeDestinationQueryProps | null =
    useMemo(() => {
      if (typeof chainId !== 'number') return null
      if (!address) return null
      if (!kappa) return null
      if (!timestamp) return null
      if (!bridgeType) return null

      return { chainId, address, kappa, timestamp, bridgeType }
    }, [chainId, address, kappa, timestamp, bridgeType])

  // Start fallback query
  useEffect(() => {
    const lastFetchedAddress: boolean = Boolean(
      lastFetchedQueryParams?.lastArg?.address
    )
    if (useFallback && validQueryParams) {
      console.log('start dest fallback subscription, kappa: ', kappa)
      fetchFallbackBridgeDestinationQuery({
        chainId: validQueryParams.chainId,
        address: validQueryParams.address,
        kappa: validQueryParams.kappa,
        timestamp: validQueryParams.timestamp,
        bridgeType: validQueryParams.bridgeType,
      })
    } else if (!useFallback && lastFetchedAddress) {
      console.log('end dest fallback subscription, kappa: ', kappa)
      fetchFallbackBridgeDestinationQuery({
        chainId: null,
        address: null,
        kappa: null,
        timestamp: null,
        bridgeType: null,
      }).unsubscribe()
    }
  }, [useFallback, validQueryParams])

  useEffect(() => {
    const {
      isLoading,
      isUninitialized,
      isSuccess,
      data: fallbackQueryData,
    } = fetchedFallbackQuery

    const {
      bridgeTx: destinationInfo,
      kappa,
      pending,
    } = fallbackQueryData?.getDestinationBridgeTx || {}

    const isCompleted: boolean =
      Boolean(!pending) || Boolean(destinationInfo?.txnHash)

    // Update bridge transaction in either Pending or Fallback
    if (destinationInfo && kappa && isCompleted) {
      const originQueryTransaction: BridgeTransaction | undefined =
        fallbackQueryPendingTransactions.find(
          (transaction: BridgeTransaction) => transaction.kappa === kappa
        ) ??
        pendingAwaitingCompletionTransactions.find(
          (transaction: BridgeTransaction) => transaction.kappa === kappa
        )

      const destinationQueryAlreadySaved: boolean =
        fallbackQueryPendingTransactions.some(
          (transaction: BridgeTransaction) =>
            transaction?.toInfo === destinationInfo
        )

      if (originQueryTransaction && !destinationQueryAlreadySaved) {
        const constructedBridgeTransaction: BridgeTransaction = {
          fromInfo: originQueryTransaction?.fromInfo,
          toInfo: destinationInfo,
          kappa: kappa,
        }
        console.log(
          'complete fallback transaction: ',
          constructedBridgeTransaction
        )
        dispatch(removeFallbackQueryPendingTransaction(kappa))
        dispatch(
          addFallbackQueryHistoricalTransaction(constructedBridgeTransaction)
        )
      }
    }
  }, [fetchedFallbackQuery, fallbackQueryPendingTransactions])

  return null
}
