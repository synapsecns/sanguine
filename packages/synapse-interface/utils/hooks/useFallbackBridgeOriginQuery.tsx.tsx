import { useState, useEffect } from 'react'
import {
  useLazyGetOriginBridgeTxFallbackQuery,
  BridgeType,
} from '@/slices/api/generated'

interface useFallbackBridgeOriginQueryProps {
  useFallback: boolean
  chainId?: number
  txnHash?: string
  bridgeType?: BridgeType
}

export const useFallbackBridgeOriginQuery = ({
  useFallback,
  chainId,
  txnHash,
  bridgeType,
}: useFallbackBridgeOriginQueryProps) => {
  const [fetchFallbackBridgeOriginQuery, fetchedFallbackQueries] =
    useLazyGetOriginBridgeTxFallbackQuery({ pollingInterval: 30000 })

  return null
}
