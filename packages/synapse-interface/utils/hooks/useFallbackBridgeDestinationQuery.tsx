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

  return null
}
