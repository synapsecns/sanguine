import { useEffect } from 'react'

import { useAppDispatch } from '@/store/hooks'
import {
  fetchAvaxPrice,
  fetchEthPrice,
  fetchGmxPrice,
  fetchMetisPrice,
  fetchSynPrices,
} from '@/slices/priceDataSlice'
import { fetchGasAirdropPrices } from '@/slices/gasAirdropSlice'

export const useFetchPricesOnInterval = () => {
  const dispatch = useAppDispatch()

  useEffect(() => {
    const fetchPrices = () => {
      dispatch(fetchSynPrices())
      dispatch(fetchEthPrice())
      dispatch(fetchAvaxPrice())
      dispatch(fetchMetisPrice())
      dispatch(fetchGmxPrice())
      dispatch(fetchGasAirdropPrices())
    }

    // Fetch on mount
    fetchPrices()

    // Fetch every five minutes
    const interval = setInterval(fetchPrices, 300000)

    return () => clearInterval(interval)
  }, [dispatch])
}
