import { useEffect } from 'react'
import { Address } from '@wagmi/core'
import { watchPendingTransactions, waitForTransaction } from '@wagmi/core'
import { useAppDispatch } from '@/store/hooks'
import { useBridgeState } from '../bridge/hooks'
import { BridgeState } from '../bridge/reducer'
import { fetchAndStoreSingleNetworkPortfolioBalances } from './hooks'
import { useAccount } from 'wagmi'

export default function Updater(): null {
  const dispatch = useAppDispatch()
  const { address } = useAccount()
  const { recentBridgeTransactions }: BridgeState = useBridgeState()

  // Update balances when transaction resolves
  useEffect(() => {
    if (
      !address ||
      !recentBridgeTransactions ||
      recentBridgeTransactions.length === 0
    ) {
      return
    }

    ;(async () => {
      const updateChainId = recentBridgeTransactions[0]?.originChain?.id
      const transactionHash = recentBridgeTransactions[0]
        .transactionHash as Address

      const resolvedTransaction = await waitForTransaction({
        hash: transactionHash,
      })

      await dispatch(
        fetchAndStoreSingleNetworkPortfolioBalances({
          address: address,
          chainId: updateChainId,
        })
      )
    })()
  }, [recentBridgeTransactions, address, dispatch])

  return null
}
