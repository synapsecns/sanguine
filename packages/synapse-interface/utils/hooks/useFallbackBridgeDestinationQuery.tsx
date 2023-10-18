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
  return null
}
