import { useEffect } from 'react'
import { Address } from '@wagmi/core'
import { watchPendingTransactions, waitForTransaction } from '@wagmi/core'
import { useAppDispatch } from '@/store/hooks'
import { useBridgeState } from '../bridge/hooks'
import { useTransactionsState } from '../transactions/hooks'
import { TransactionsState } from '../transactions/reducer'
import { BridgeState } from '../bridge/reducer'
import { fetchAndStoreSingleNetworkPortfolioBalances } from './hooks'

import { useAccount } from 'wagmi'
import {
  PendingBridgeTransaction,
  updatePendingBridgeTransaction,
} from '../bridge/actions'
import { BridgeTransaction } from '../api/generated'

export default function Updater(): null {
  const dispatch = useAppDispatch()
  const { address } = useAccount()
  const { pendingBridgeTransactions }: BridgeState = useBridgeState()
  const {
    userHistoricalTransactions,
    isUserHistoricalTransactionsLoading,
  }: TransactionsState = useTransactionsState()

  // Update Origin balances when transaction resolves
  useEffect(() => {
    if (!address || !pendingBridgeTransactions) return
    if (pendingBridgeTransactions && pendingBridgeTransactions.length > 0) {
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

  // Update Destination balances for new historical transaction
  useEffect(() => {
    if (!address || isUserHistoricalTransactionsLoading) return
    if (userHistoricalTransactions && userHistoricalTransactions.length > 0) {
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
