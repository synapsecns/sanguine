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
  addFallbackQueryTransaction,
  updateFallbackQueryTransaction,
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

  const { fallbackQueryTransactions }: TransactionsState =
    useTransactionsState()

  const [fetchFallbackBridgeDestinationQuery, fetchedFallbackQuery] =
    useLazyGetDestinationBridgeTxFallbackQuery({ pollingInterval: 30000 })

  const validQueryParams: FallbackBridgeDestinationQueryProps | null =
    useMemo(() => {
      if (typeof chainId !== 'number') return null
      if (!address) return null
      if (!kappa) return null
      if (!timestamp) return null
      if (!bridgeType) return null

      return { chainId, kappa, timestamp, bridgeType }
    }, [chainId, kappa, timestamp, bridgeType])

  // Start fallback query
  useEffect(() => {
    if (useFallback && validQueryParams) {
      // console.log('start fetch')
      fetchFallbackBridgeDestinationQuery({
        chainId: validQueryParams.chainId,
        address: validQueryParams.address,
        kappa: validQueryParams.kappa,
        timestamp: validQueryParams.timestamp,
        bridgeType: validQueryParams.bridgeType,
      })
    } else if (!useFallback) {
      // console.log('end fetch')
      fetchFallbackBridgeDestinationQuery({
        chainId: null,
        address: null,
        kappa: null,
        timestamp: null,
        bridgeType: null,
      })
    }
  }, [useFallback, validQueryParams])

  useEffect(() => {
    const {
      isLoading,
      isUninitialized,
      isSuccess,
      data: fallbackQueryData,
    } = fetchedFallbackQuery

    const { bridgeTx: destinationInfo, kappa } =
      fallbackQueryData?.getDestinationBridgeTx || {}

    // Update bridge transaction in either Pending or Fallback
    if (destinationInfo && kappa) {
      const originQueryTransaction: BridgeTransaction | undefined =
        fallbackQueryTransactions.find(
          (transaction: BridgeTransaction) => transaction.kappa === kappa
        )

      if (originQueryTransaction) {
        const constructedBridgeTransaction: BridgeTransaction = {
          fromInfo: originQueryTransaction?.fromInfo,
          toInfo: destinationInfo,
          kappa: kappa,
        }
        dispatch(updateFallbackQueryTransaction(constructedBridgeTransaction))
      }
    }
  }, [fetchedFallbackQuery, fallbackQueryTransactions])

  return null
}
