import { useMemo } from 'react'
import { useAccount, Address } from 'wagmi'
import { useGetUserHistoricalActivityQuery } from '@/slices/api/generated'
import { getTimeMinutesBeforeNow } from '@/utils/time'

export const Activity = () => {
  const { address } = useAccount()
  const oneMonthInMinutes: number = 43200
  const queryTime: number = getTimeMinutesBeforeNow(oneMonthInMinutes)

  const currentAddress: Address = useMemo(() => address, [address])
  const shouldSkip: boolean = !address || !queryTime

  console.log('currentAddress: ', currentAddress)
  console.log('queryTime: ', queryTime)
  console.log('shouldSkip: ', shouldSkip)

  const { data, isLoading, isSuccess, isError, error } =
    useGetUserHistoricalActivityQuery(
      {
        address: currentAddress ?? '',
        startTime: queryTime,
        // address: '0xF080B794AbF6BB905F2330d25DF545914e6027F8',
        // startTime: 1689015547,
      },
      {
        skip: shouldSkip,
      }
    )

  const historicalActivityData = useMemo(() => {
    console.log('data: ', data)
    console.log('isSuccess: ', isSuccess)
  }, [data, shouldSkip])

  return <>Activity</>
}
