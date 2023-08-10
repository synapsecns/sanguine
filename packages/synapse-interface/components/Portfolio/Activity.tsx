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
    return historicalActivity?.data?.bridgeTransactions || []
  }, [historicalActivity?.data?.bridgeTransactions])

  useEffect(() => {
    address &&
      queryTime &&
      fetchUserHistoricalActivity({ address: address, startTime: queryTime })
  }, [address])

  console.log('userHistoricalActivity: ', userHistoricalActivity)
  return (
    <div>
      <ActivitySection title="Pending">
        <TransactionHeader transactionType={ActivityType.PENDING} />
      </ActivitySection>
      <ActivitySection title="Recent">
        <TransactionHeader transactionType={ActivityType.RECENT} />
        {userHistoricalActivity.map((transaction: BridgeTransaction) => (
          <Transaction bridgeTransaction={transaction} />
        ))}
      </ActivitySection>
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

export enum ActivityType {
  PENDING,
  STUCK,
  RECENT,
}

export const TransactionHeader = ({
  transactionType,
}: {
  transactionType: ActivityType
}) => {
  return (
    <div className="grid grid-cols-4 gap-2 text-[#C0BCC2] text-sm">
      <div>From</div>
      <div>To</div>
      <div className="flex justify-end">
        {transactionType === ActivityType.PENDING && 'Blocks'}
        {transactionType === ActivityType.RECENT && 'Rate'}
      </div>
      <div className="flex justify-end">
        {transactionType === ActivityType.PENDING && 'Elapsed'}
        {transactionType === ActivityType.RECENT && 'Completed'}
      </div>
    </div>
  )
}

export const Transaction = ({
  bridgeTransaction,
}: {
  bridgeTransaction: BridgeTransaction
}) => {
  const { fromInfo, toInfo }: { fromInfo?: PartialInfo; toInfo?: PartialInfo } =
    bridgeTransaction || {}

  const {
    chainID: originChainId,
    formattedValue: originFormattedValue,
    tokenAddress: originTokenAddress,
    tokenSymbol: originTokenSymbol,
    blockNumber: bridgeOriginBlockNumber,
    time: bridgeOriginTime,
  } = fromInfo || {}

  const {
    chainID: destinationChainId,
    formattedValue: destinationFormattedValue,
    tokenAddress: destinationTokenAddress,
    tokenSymbol: destinationTokenSymbol,
    blockNumber: bridgeDestinationBlockNumber,
    time: bridgeDestinationTime,
  } = toInfo || {}

  return (
    <div className="grid grid-cols-4 gap-2 text-sm text-white">
      <div>{originChainId}</div>
      <div>{destinationChainId}</div>
      <div className="flex justify-end"></div>
      <div className="flex justify-end"></div>
    </div>
  )
}
