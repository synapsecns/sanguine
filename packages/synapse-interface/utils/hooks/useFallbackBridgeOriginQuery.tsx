import { useState, useEffect, useMemo } from 'react'
import { useAppDispatch } from '@/store/hooks'
import { addFallbackQueryTransaction } from '@/slices/transactions/actions'
import {
  useLazyGetOriginBridgeTxFallbackQuery,
  BridgeTransaction,
  BridgeType,
} from '@/slices/api/generated'
import { useTransactionsState } from '@/slices/transactions/hooks'
import { TransactionsState } from '@/slices/transactions/reducer'

interface FallbackBridgeOriginQueryProps {
  chainId?: number
  txnHash?: string
  bridgeType?: BridgeType
}
interface useFallbackBridgeOriginQueryProps
  extends FallbackBridgeOriginQueryProps {
  useFallback: boolean
}

export const useFallbackBridgeOriginQuery = ({
  useFallback,
  chainId,
  txnHash,
  bridgeType,
}: useFallbackBridgeOriginQueryProps) => {
  const dispatch = useAppDispatch()
  const { fallbackQueryTransactions }: TransactionsState =
    useTransactionsState()

  const [fetchFallbackBridgeOriginQuery, fetchedFallbackQuery] =
    useLazyGetOriginBridgeTxFallbackQuery({ pollingInterval: 30000 })

  const validQueryParams: FallbackBridgeOriginQueryProps | null =
    useMemo(() => {
      if (typeof chainId !== 'number') return null
      if (!txnHash) return null
      if (!bridgeType) return null

      return { chainId, txnHash, bridgeType }
    }, [chainId, txnHash, bridgeType])

  // Start fallback query
  useEffect(() => {
    if (useFallback && validQueryParams) {
      console.log('start fetch')
      fetchFallbackBridgeOriginQuery({
        chainId: validQueryParams.chainId,
        txnHash: validQueryParams.txnHash,
        bridgeType: validQueryParams.bridgeType,
      })
    } else if (!useFallback) {
      console.log('end fetch')
      fetchFallbackBridgeOriginQuery({
        chainId: null,
        txnHash: null,
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

    const { bridgeTx: originInfo, kappa } =
      fallbackQueryData?.getOriginBridgeTx || {}

    if (originInfo && kappa) {
      const constructedBridgeTransaction: BridgeTransaction = {
        fromInfo: originInfo,
        toInfo: null,
        kappa: kappa,
      }

      const alreadyExists: boolean = fallbackQueryTransactions.some(
        (transaction) => {
          return (
            transaction.kappa === constructedBridgeTransaction.kappa ||
            transaction.fromInfo === constructedBridgeTransaction.fromInfo
          )
        }
      )

      if (!alreadyExists) {
        dispatch(addFallbackQueryTransaction(constructedBridgeTransaction))
      }
    }
  }, [fetchedFallbackQuery])

  return null
}
