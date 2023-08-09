import { useEffect } from 'react'
import { useAccount } from 'wagmi'
import { useLazyGetUserHistoricalActivityQuery } from '@/slices/api/generated'
import { getTimeMinutesBeforeNow } from '@/utils/time'

export const Activity = () => {
  const { address } = useAccount()
  const oneMonthInMinutes: number = 43200
  const queryTime: number = getTimeMinutesBeforeNow(oneMonthInMinutes)

  const [trigger, result, lastPromiseInfo] =
    useLazyGetUserHistoricalActivityQuery()

  useEffect(() => {
    address && queryTime && trigger({ address: address, startTime: queryTime })
  }, [address])

  console.log('result: ', result)

  return <>Activity</>
}
