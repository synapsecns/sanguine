import { useEffect } from 'react'
import { useNetwork } from 'wagmi'

import { useAppDispatch } from '@/store/hooks'
import { fetchGasData } from '@/slices/gasDataSlice'

export const useFetchGasDataOnInterval = () => {
  const dispatch = useAppDispatch()
  const { chain } = useNetwork()

  const fetchData = (chainId: number) => {
    dispatch(fetchGasData(chainId))
  }

  useEffect(() => {
    // Fetch when chainId available
    if (chain?.id) {
      fetchData(chain?.id)
    }

    // Fetch every 60 seconds
    const interval = setInterval(fetchGasData, 60000)

    return () => clearInterval(interval)
  }, [dispatch, chain?.id])
}
