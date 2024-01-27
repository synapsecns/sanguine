import { useEffect } from 'react'
import { useAccount } from 'wagmi'
import { Address } from '@wagmi/core'

import { useAppDispatch } from '@/store/hooks'
import { useTransactionsState } from '@/slices/transactions/hooks'
import { TransactionsState } from '@/slices/transactions/reducer'
import { fetchAndStoreSingleNetworkPortfolioBalances } from '@/slices/portfolio/hooks'
import { PendingBridgeTransaction } from '@/slices/transactions/actions'
import { BridgeTransaction } from '@/slices/api/generated'
import { checkTransactionsExist } from '@/utils/checkTransactionsExist'

export const usePortfolioListener = () => {
  const dispatch = useAppDispatch()
  const { address } = useAccount()
  const {
    pendingBridgeTransactions,
    userHistoricalTransactions,
    isUserHistoricalTransactionsLoading,
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

  return null
}
