import { useEffect } from 'react'
import { useAppDispatch } from '@/store/hooks'
import { useBridgeState } from '../bridge/hooks'
import { BridgeState } from '../bridge/reducer'
import { fetchAndStoreSingleNetworkPortfolioBalances } from './hooks'
import { useAccount } from 'wagmi'

export default function Updater(): null {
  const { address } = useAccount()
  const dispatch = useAppDispatch()
  const { recentBridgeTransactions }: BridgeState = useBridgeState()

  useEffect(() => {
    if (
      address &&
      recentBridgeTransactions &&
      recentBridgeTransactions.length > 0
    ) {
      const updateChainId: number = recentBridgeTransactions[0]?.originChain?.id
      dispatch(
        fetchAndStoreSingleNetworkPortfolioBalances({
          address: address,
          chainId: updateChainId,
        })
      )
    }
  }, [recentBridgeTransactions, address])

  return null
}
