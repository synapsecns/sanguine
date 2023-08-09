import { useEffect, useMemo } from 'react'
import { useAccount } from 'wagmi'
import {
  useLazyGetUserHistoricalActivityQuery,
  PartialInfo,
  BridgeTransaction,
  GetUserHistoricalActivityQuery,
} from '@/slices/api/generated'
import { getTimeMinutesBeforeNow } from '@/utils/time'

export const Activity = () => {
  const { address } = useAccount()
  const oneMonthInMinutes: number = 43200
  const queryTime: number = getTimeMinutesBeforeNow(oneMonthInMinutes)

  const [fetchUserHistoricalActivity, historicalActivity, lastPromiseInfo] =
    useLazyGetUserHistoricalActivityQuery()

  const userHistoricalActivity: GetUserHistoricalActivityQuery = useMemo(() => {
    return historicalActivity?.data
  }, [historicalActivity?.data])

  useEffect(() => {
    address &&
      queryTime &&
      fetchUserHistoricalActivity({ address: address, startTime: queryTime })
  }, [address])

  console.log('userHistoricalActivity: ', userHistoricalActivity)
  return <>Activity</>
}
