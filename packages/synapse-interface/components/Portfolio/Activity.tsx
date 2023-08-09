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

  const userHistoricalActivity: BridgeTransaction[] = useMemo(() => {
    return historicalActivity?.data?.bridgeTransactions
  }, [historicalActivity?.data?.bridgeTransactions])

  useEffect(() => {
    address &&
      queryTime &&
      fetchUserHistoricalActivity({ address: address, startTime: queryTime })
  }, [address])

  return (
    <div>
      <ActivitySection title="Stuck"></ActivitySection>
      <ActivitySection title="Recent"></ActivitySection>
    </div>
  )
}

export const ActivitySection = ({
  title,
  children,
}: {
  title: string
  children?: React.ReactNode
}) => {
  return (
    <div>
      <h3 className="text-lg text-white">{title}</h3>
      {children}
    </div>
  )
}
