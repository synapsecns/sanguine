import { useAccount } from 'wagmi'
import { useGetUserHistoricalActivityQuery } from '@/slices/api/generated'
import { getTimeMinutesFromNow } from '@/utils/time'

export const Activity = () => {
  const { address } = useAccount()
  const currentTime: number = getTimeMinutesFromNow(0)
  const shouldSkip: boolean = !address || !currentTime
  console.log('address: ', address)
  console.log('currentTime: ', currentTime)
  console.log('shouldSkip: ', shouldSkip)

  const {
    data: history,
    isLoading,
    isSuccess,
    isError,
    error,
  } = useGetUserHistoricalActivityQuery(
    {
      address: address,
      startTime: currentTime,
    },
    {
      skip: shouldSkip,
    }
  )

  console.log('history:', history)

  return <>Activity</>
}
