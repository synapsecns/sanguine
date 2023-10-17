import { useState, useEffect, useMemo } from 'react'
import {
  useLazyGetOriginBridgeTxFallbackQuery,
  BridgeType,
} from '@/slices/api/generated'
import { useTransactionsState } from '@/slices/transactions/hooks'
import { TransactionsState } from '@/slices/transactions/reducer'

interface fallbackBridgeOriginQueryProps {
  chainId?: number
  txnHash?: string
  bridgeType?: BridgeType
}
interface useFallbackBridgeOriginQueryProps
  extends fallbackBridgeOriginQueryProps {
  useFallback: boolean
}

export const useFallbackBridgeOriginQuery = ({
  useFallback,
  chainId,
  txnHash,
  bridgeType,
}: useFallbackBridgeOriginQueryProps) => {
  const { fallbackQueryTransactions }: TransactionsState =
    useTransactionsState()

  const [fetchFallbackBridgeOriginQuery, fetchedFallbackQueries] =
    useLazyGetOriginBridgeTxFallbackQuery({ pollingInterval: 30000 })

  const validQueryParams: fallbackBridgeOriginQueryProps | null =
    useMemo(() => {
      if (typeof chainId !== 'number') return null
      if (!txnHash) return null
      if (!bridgeType) return null

      return { chainId, txnHash, bridgeType }
    }, [chainId, txnHash, bridgeType])

  useEffect(() => {
    if (useFallback && validQueryParams) {
      fetchFallbackBridgeOriginQuery({
        chainId: validQueryParams.chainId,
        txnHash: validQueryParams.txnHash,
        bridgeType: validQueryParams.bridgeType,
      })
    }
  }, [useFallback, validQueryParams])

  return null
}
