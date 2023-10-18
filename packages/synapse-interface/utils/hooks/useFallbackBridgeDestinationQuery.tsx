import { useState, useEffect, useMemo } from 'react'
import { useAppDispatch } from '@/store/hooks'
import {
  useLazyGetDestinationBridgeTxFallbackQuery,
  BridgeTransaction,
  BridgeType,
} from '@/slices/api/generated'
import { useTransactionsState } from '@/slices/transactions/hooks'
import { TransactionsState } from '@/slices/transactions/reducer'

interface FallbackBridgeDestinationQueryProps {
  chainId?: number
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
  kappa,
  timestamp,
  bridgeType,
  useFallback,
}: useFallbackBridgeDestinationQueryProps) => {
  const [fetchFallbackBridgeDestinationQuery, fetchedFallbackQuery] =
    useLazyGetDestinationBridgeTxFallbackQuery({ pollingInterval: 30000 })

  const validQueryParams: FallbackBridgeDestinationQueryProps | null =
    useMemo(() => {
      if (typeof chainId !== 'number') return null
      if (!kappa) return null
      if (!timestamp) return null
      if (!bridgeType) return null

      return { chainId, kappa, timestamp, bridgeType }
    }, [chainId, kappa, timestamp, bridgeType])

  return null
}
