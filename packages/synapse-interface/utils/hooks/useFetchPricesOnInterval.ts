import { useEffect } from 'react'

import { useAppDispatch } from '@/store/hooks'
import {
  fetchAllEthStablecoinPrices,
  fetchArbPrice,
  fetchAvaxPrice,
  fetchCoingeckoPrices,
  fetchDaiePrice,
  fetchEthPrice,
  fetchGmxPrice,
  fetchMetisPrice,
  fetchMusdcPrice,
  fetchSynPrices,
} from '@/slices/priceDataSlice'

export const useFetchPricesOnInterval = () => {
  const dispatch = useAppDispatch()

  useEffect(() => {
    const fetchPrices = () => {
      dispatch(fetchSynPrices())
      dispatch(fetchEthPrice())
      dispatch(fetchAvaxPrice())
      dispatch(fetchMetisPrice())
      dispatch(fetchArbPrice())
      dispatch(fetchGmxPrice())
      dispatch(fetchAllEthStablecoinPrices())
      dispatch(fetchCoingeckoPrices())
      dispatch(fetchMusdcPrice())
      dispatch(fetchDaiePrice())
    }

    // Fetch on mount
    fetchPrices()

    // Fetch every five minutes
    const interval = setInterval(fetchPrices, 300000)

    return () => clearInterval(interval)
  }, [dispatch])
}
