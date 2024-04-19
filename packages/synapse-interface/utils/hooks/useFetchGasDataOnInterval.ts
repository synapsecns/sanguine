import { useEffect } from 'react'

import { useAppDispatch } from '@/store/hooks'
import { fetchGasData, resetGasData } from '@/slices/gasDataSlice'
import { useBridgeState } from '@/slices/bridge/hooks'

export const useFetchGasDataOnInterval = () => {
  const dispatch = useAppDispatch()
  const { fromChainId } = useBridgeState()

  const fetchData = (chainId: number) => {
    dispatch(fetchGasData(chainId))
  }

  useEffect(() => {
    // Reset gas data when selecting new chain
    dispatch(resetGasData())

    // Fetch when chainId available
    if (fromChainId) {
      fetchData(fromChainId)
    }

    // Fetch every 60 seconds
    const interval = setInterval(fetchGasData, 60000)

    return () => clearInterval(interval)
  }, [dispatch, fromChainId])
}
