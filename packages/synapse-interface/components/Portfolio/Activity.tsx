import { useEffect, useMemo } from 'react'
import { useAccount } from 'wagmi'
import Image from 'next/image'
import {
  useLazyGetUserHistoricalActivityQuery,
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
      className="grid grid-cols-4 gap-2 text-[#C0BCC2] text-sm"
    >
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
      className="grid grid-cols-4 gap-2 text-sm text-white border-b border-[#565058]"
    >
      <TransactionPayloadDetail
        chain={originChain}
        token={originToken}
        tokenAmount={originFormattedValue}
      />
      <TransactionPayloadDetail
        chain={destinationChain}
        token={destinationToken}
        tokenAmount={destinationFormattedValue}
      />
      <div className="flex justify-end">
        {transactionType === ActivityType.RECENT && (
          <ExchangeRate
            originValue={originFormattedValue}
            destinationValue={destinationFormattedValue}
          />
        )}
      </div>
      <div className="flex justify-end">
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
    <div className="flex">
      <div className="text-[#C0BCC2]">{`1 : `}</div>
      <div className="text-white">{formattedExchangeRate}</div>
    </div>
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
  return <div>{formattedTime}</div>
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
    <div data-test-id="transaction-payload-detail" className="flex flex-col">
      {chain && (
        <div
          data-test-id="transaction-payload-network"
          className="flex flex-row"
        >
          <Image
            src={chain.chainImg}
            className="w-6 h-6 mr-3 rounded-md"
            alt={`${chain.name} icon`}
          />
          <div>{chain.name}</div>
        </div>
      )}

      {token && (
        <div data-test-id="transaction-payload-token" className="flex flex-row">
          <Image
            src={token.icon}
            className="w-6 h-6 mr-3 rounded-md"
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
