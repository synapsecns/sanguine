import { useEffect } from 'react'

import { useAppDispatch } from '@/store/hooks'
import { useBridgeState } from '@/slices/bridge/hooks'
import { fetchGasData, resetGasData } from '@/slices/gasDataSlice'

export const useFetchGasDataOnInterval = () => {
  const dispatch = useAppDispatch()
  const { fromChainId } = useBridgeState()

  const fetchGas = () => {
    if (fromChainId) {
      dispatch(fetchGasData(fromChainId))
    }
  }

  useEffect(() => {
    // Reset gas data when selecting new chain
    dispatch(resetGasData())

    fetchGas()

    const interval = setInterval(fetchGas, 60000)

    return () => clearInterval(interval)
  }, [dispatch, fromChainId])
}
