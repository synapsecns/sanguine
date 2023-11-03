import { useState, useEffect, useMemo } from 'react'
import { useAppDispatch } from '@/store/hooks'
import { addFallbackQueryPendingTransaction } from '@/slices/transactions/actions'
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

  const { fallbackQueryPendingTransactions }: TransactionsState =
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

  const queryTransactionAlreadyStored: boolean = useMemo(() => {
    return fallbackQueryPendingTransactions.some((transaction) => {
      return transaction?.fromInfo?.txnHash === txnHash
    })
  }, [fallbackQueryPendingTransactions, txnHash])

  // Start fallback query
  useEffect(() => {
    if (useFallback && validQueryParams) {
      console.log('start origin fallback subscription, txnHash: ', txnHash)
      fetchFallbackBridgeOriginQuery({
        chainId: validQueryParams.chainId,
        txnHash: validQueryParams.txnHash,
        bridgeType: validQueryParams.bridgeType,
      })
    } else if (!useFallback || queryTransactionAlreadyStored) {
      console.log('end origin fallback subscription, txnHash: ', txnHash)
      fetchFallbackBridgeOriginQuery({
        chainId: null,
        txnHash: null,
        bridgeType: null,
      }).unsubscribe()
    }
  }, [useFallback, validQueryParams, queryTransactionAlreadyStored])

  useEffect(() => {
    const {
      isLoading,
      isUninitialized,
      isSuccess,
      data: fallbackQueryData,
    } = fetchedFallbackQuery

    const {
      bridgeTx: originInfo,
      kappa,
      pending,
    } = fallbackQueryData?.getOriginBridgeTx || {}

    if (originInfo && kappa && pending) {
      const constructedBridgeTransaction: BridgeTransaction = {
        fromInfo: originInfo,
        toInfo: null,
        kappa: kappa,
      }

      const alreadyExists: boolean = fallbackQueryPendingTransactions.some(
        (transaction) => {
          return (
            transaction.kappa === constructedBridgeTransaction.kappa ||
            transaction.fromInfo === constructedBridgeTransaction.fromInfo
          )
        }
      )

      if (!alreadyExists) {
        dispatch(
          addFallbackQueryPendingTransaction(constructedBridgeTransaction)
        )
      }
    }
  }, [fetchedFallbackQuery, fallbackQueryPendingTransactions])

  return null
}
