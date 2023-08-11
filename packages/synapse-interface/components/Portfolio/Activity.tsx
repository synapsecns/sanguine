import { useEffect, useMemo } from 'react'
import { useAccount } from 'wagmi'
import Image from 'next/image'
import {
  useLazyGetUserHistoricalActivityQuery,
  useLazyGetUserPendingTransactionsQuery,
  PartialInfo,
  BridgeTransaction,
  GetUserHistoricalActivityQuery,
} from '@/slices/api/generated'
import {
  getTimeMinutesBeforeNow,
  convertUnixTimestampToMonthAndDate,
} from '@/utils/time'
import { CHAINS_BY_ID } from '@/constants/chains'
import { Chain, Token } from '@/utils/types'
import { tokenSymbolToToken } from '@/constants/tokens'

export const Activity = () => {
  const { address } = useAccount()
  const oneMonthInMinutes: number = 43200
  const queryTime: number = getTimeMinutesBeforeNow(oneMonthInMinutes)

  const [fetchUserHistoricalActivity, historicalActivity, lastPromiseInfo] =
    useLazyGetUserHistoricalActivityQuery()

  const [fetchUserPendingActivity, pendingActivity, lastPendingPromiseInfo] =
    useLazyGetUserPendingTransactionsQuery()

  const userHistoricalActivity: BridgeTransaction[] = useMemo(() => {
    return historicalActivity?.data?.bridgeTransactions || []
  }, [historicalActivity?.data?.bridgeTransactions])

  const userPendingActivity: BridgeTransaction[] = useMemo(() => {
    return pendingActivity?.data?.bridgeTransactions || []
  }, [pendingActivity?.data?.bridgeTransactions])

  useEffect(() => {
    address &&
      queryTime &&
      fetchUserHistoricalActivity({ address: address, startTime: queryTime })
  }, [address])

  console.log('userHistoricalActivity: ', userHistoricalActivity)
  console.log('userPendingActivity: ', userPendingActivity)

  return (
    <div data-test-id="activity">
      <ActivitySection title="Pending">
        <TransactionHeader transactionType={ActivityType.PENDING} />
      </ActivitySection>
      <ActivitySection title="Recent">
        <TransactionHeader transactionType={ActivityType.RECENT} />
        {userHistoricalActivity.map((transaction: BridgeTransaction) => (
          <Transaction
            bridgeTransaction={transaction}
            transactionType={ActivityType.RECENT}
          />
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
    <div data-test-id="activity-section">
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
    <div
      data-test-id="transaction-header"
      className="grid grid-cols-10 gap-2 text-[#C0BCC2] text-sm mt-4 mb-2"
    >
      <div className="col-span-3">From</div>
      <div className="col-span-3">To</div>
      <div className="flex justify-end col-span-2">
        {transactionType === ActivityType.PENDING && 'Blocks'}
        {transactionType === ActivityType.RECENT && 'Rate'}
      </div>
      <div className="flex justify-end col-span-2">
        {transactionType === ActivityType.PENDING && 'Elapsed'}
        {transactionType === ActivityType.RECENT && 'Completed'}
      </div>
    </div>
  )
}

export const Transaction = ({
  bridgeTransaction,
  transactionType,
}: {
  bridgeTransaction: BridgeTransaction
  transactionType: ActivityType
}) => {
  const { fromInfo, toInfo }: { fromInfo?: PartialInfo; toInfo?: PartialInfo } =
    bridgeTransaction || {}

  const {
    chainID: originChainId,
    value: originRawValue,
    formattedValue: originFormattedValue,
    tokenAddress: originTokenAddress,
    tokenSymbol: originTokenSymbol,
    blockNumber: bridgeOriginBlockNumber,
    time: bridgeOriginTime,
  } = fromInfo || {}

  const originChain: Chain = CHAINS_BY_ID[originChainId]
  const originToken: Token = tokenSymbolToToken(
    originChainId,
    originTokenSymbol
  )

  const {
    chainID: destinationChainId,
    value: destinationRawValue,
    formattedValue: destinationFormattedValue,
    tokenAddress: destinationTokenAddress,
    tokenSymbol: destinationTokenSymbol,
    blockNumber: bridgeDestinationBlockNumber,
    time: bridgeDestinationTime,
  } = toInfo || {}

  const destinationChain: Chain = CHAINS_BY_ID[destinationChainId]
  const destinationToken: Token = tokenSymbolToToken(
    destinationChainId,
    destinationTokenSymbol
  )

  return (
    <div
      data-test-id="transaction"
      className={`
        grid grid-cols-10 mt-auto py-3
        text-sm text-white border-b border-[#565058]
        items-end
        `}
    >
      <div className="flex col-span-3">
        <TransactionPayloadDetail
          chain={originChain}
          token={originToken}
          tokenAmount={originFormattedValue}
        />
        <div className="flex items-end mb-[3px] ml-auto px-4">→</div>
      </div>
      <div className="col-span-3">
        <TransactionPayloadDetail
          chain={destinationChain}
          token={destinationToken}
          tokenAmount={destinationFormattedValue}
        />
      </div>
      <div className="flex justify-end col-span-2">
        {transactionType === ActivityType.RECENT && (
          <ExchangeRate
            originValue={originFormattedValue}
            destinationValue={destinationFormattedValue}
          />
        )}
      </div>
      <div className="flex justify-end col-span-2">
        {transactionType === ActivityType.RECENT && (
          <Completed transactionCompletedTime={bridgeDestinationTime} />
        )}
      </div>
    </div>
  )
}

export const ExchangeRate = ({
  originValue,
  destinationValue,
}: {
  originValue: number
  destinationValue: number
}) => {
  const exchangeRate: number = originValue / destinationValue
  const formattedExchangeRate: string = exchangeRate.toFixed(4)
  return (
    <span>
      <span className="text-[#C0BCC2]">{`1 : `}</span>
      <span className="text-white">{formattedExchangeRate}</span>
    </span>
  )
}

export const Completed = ({
  transactionCompletedTime,
}: {
  transactionCompletedTime: number
}) => {
  const formattedTime: string =
    transactionCompletedTime &&
    convertUnixTimestampToMonthAndDate(transactionCompletedTime)
  return (
    <span>
      <span className="w-4 pt-3 mb-auto font-bold text-green-500"> ✓ </span>{' '}
      {formattedTime}
    </span>
  )
}

export const TransactionPayloadDetail = ({
  chain,
  token,
  tokenAmount,
}: {
  chain?: Chain
  token?: Token
  tokenAmount?: number
}) => {
  return (
    <div
      data-test-id="transaction-payload-detail"
      className="flex flex-col space-y-1"
    >
      {chain && (
        <div
          data-test-id="transaction-payload-network"
          className="flex flex-row items-center"
        >
          <Image
            src={chain.chainImg}
            className="w-6 h-6 mr-3 rounded-full"
            alt={`${chain.name} icon`}
          />
          <div>{chain.name}</div>
        </div>
      )}

      {token && (
        <div
          data-test-id="transaction-payload-token"
          className="flex flex-row items-center"
        >
          <Image
            src={token.icon}
            className="items-center w-6 h-6 mr-3 rounded-full"
            alt={`${token.name} icon`}
          />
          {typeof tokenAmount === 'number' && (
            <div className="mr-1">{tokenAmount}</div>
          )}
          <div>{token.description}</div>
        </div>
      )}
    </div>
  )
}
