import { useEffect } from 'react'
import { useAccount } from 'wagmi'
import { Address } from '@wagmi/core'
import { useAppDispatch } from '@/store/hooks'
import { useTransactionsState } from '../transactions/hooks'
import { TransactionsState } from '../transactions/reducer'
import { fetchAndStoreSingleNetworkPortfolioBalances } from './hooks'
import { PendingBridgeTransaction } from '../transactions/actions'
import { BridgeTransaction } from '../api/generated'
import { checkTransactionsExist } from '@/utils/checkTransactionsExist'

export default function Updater(): null {
  const dispatch = useAppDispatch()
  const { address } = useAccount()
  const {
    pendingBridgeTransactions,
    userHistoricalTransactions,
    isUserHistoricalTransactionsLoading,
    fallbackQueryHistoricalTransactions,
  }: TransactionsState = useTransactionsState()

  // Update Origin balances when transaction resolves
  useEffect(() => {
    if (!address || !pendingBridgeTransactions) return
    if (checkTransactionsExist(pendingBridgeTransactions)) {
      const updateOriginBalancesForNewestTransaction = async () => {
        const newestTransaction: PendingBridgeTransaction =
          pendingBridgeTransactions[0]
        const updateChainId: number = newestTransaction.originChain?.id

        dispatch(
          fetchAndStoreSingleNetworkPortfolioBalances({
            address: address as Address,
            chainId: updateChainId,
          })
        )
      }
      updateOriginBalancesForNewestTransaction()
    }
  }, [pendingBridgeTransactions, address, dispatch])

  // Update Destination balances for new historical transactions
  useEffect(() => {
    if (!address || isUserHistoricalTransactionsLoading) return
    if (checkTransactionsExist(userHistoricalTransactions)) {
      const updateDestinationBalancesForLastTransaction = async () => {
        const lastTransaction: BridgeTransaction = userHistoricalTransactions[0]
        const destinationChainId: number = lastTransaction.toInfo?.chainID
        dispatch(
          fetchAndStoreSingleNetworkPortfolioBalances({
            address: address as Address,
            chainId: destinationChainId,
          })
        )
      }
      updateDestinationBalancesForLastTransaction()
    }
  }, [userHistoricalTransactions, address])

  // Update Destination balances for new fallback historical transactions
  useEffect(() => {
    if (
      !address ||
      !checkTransactionsExist(fallbackQueryHistoricalTransactions)
    )
      return
    if (checkTransactionsExist(fallbackQueryHistoricalTransactions)) {
      const updateDestinationBalancesForNewestTransaction = async () => {
        const newestTransaction: BridgeTransaction =
          fallbackQueryHistoricalTransactions[0]
        const updateChainId: number = newestTransaction?.fromInfo?.chainID

        dispatch(
          fetchAndStoreSingleNetworkPortfolioBalances({
            address: address as Address,
            chainId: updateChainId,
          })
        )
      }
      updateDestinationBalancesForNewestTransaction()
    }
  }, [fallbackQueryHistoricalTransactions])

  return null
}
