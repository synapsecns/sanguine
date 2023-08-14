import { useEffect, useMemo, useCallback, useState } from 'react'
import { useAccount } from 'wagmi'
import { useRouter } from 'next/router'
import Link from 'next/link'
import Image from 'next/image'
import { useAppDispatch } from '@/store/hooks'
import { useTransactionsState } from '@/slices/transactions/hooks'
import {
  updateUserHistoricalTransactions,
  updateIsUserHistoricalTransactionsLoading,
  updateUserPendingTransactions,
} from '@/slices/transactions/actions'
import {
  useLazyGetUserHistoricalActivityQuery,
  useLazyGetUserPendingTransactionsQuery,
  PartialInfo,
  BridgeTransaction,
} from '@/slices/api/generated'
import {
  getTimeMinutesBeforeNow,
  convertUnixTimestampToMonthAndDate,
} from '@/utils/time'
import { CHAINS_BY_ID } from '@/constants/chains'
import { Chain, Token } from '@/utils/types'
import { tokenSymbolToToken } from '@/constants/tokens'
import { ANALYTICS_KAPPA, ANALYTICS_PATH } from '@/constants/urls'

export const Activity = () => {
  const dispatch = useAppDispatch()
  const {
    userHistoricalTransactions,
    userPendingTransactions,
    isUserHistoricalTransactionsLoading,
  } = useTransactionsState()
  const { address } = useAccount()

  const [
    fetchUserHistoricalActivity,
    fetchedHistoricalActivity,
    lastFetchArgs,
  ] = useLazyGetUserHistoricalActivityQuery()

  const [fetchUserPendingActivity, pendingActivity] =
    useLazyGetUserPendingTransactionsQuery()

  const userHistoricalActivity: BridgeTransaction[] = useMemo(() => {
    return fetchedHistoricalActivity?.data?.bridgeTransactions || []
  }, [fetchedHistoricalActivity?.data?.bridgeTransactions])

  const userPendingActivity: BridgeTransaction[] = useMemo(() => {
    return pendingActivity?.data?.bridgeTransactions || []
  }, [pendingActivity?.data?.bridgeTransactions])

  const oneMonthInMinutes: number = 43200
  const oneDayInMinutes: number = 1440
  const queryHistoricalTime: number = getTimeMinutesBeforeNow(oneMonthInMinutes)
  const queryPendingTime: number = getTimeMinutesBeforeNow(oneDayInMinutes)

  useEffect(() => {
    address &&
      queryHistoricalTime &&
      fetchUserHistoricalActivity({
        address: address,
        startTime: queryHistoricalTime,
      })

    address &&
      queryPendingTime &&
      fetchUserPendingActivity({
        address: address,
        startTime: queryPendingTime,
      })
  }, [address])

  useEffect(() => {
    const { isLoading, isUninitialized } = fetchedHistoricalActivity

    if (isUserHistoricalTransactionsLoading) {
      !isLoading &&
        !isUninitialized &&
        dispatch(updateIsUserHistoricalTransactionsLoading(false))
    }
  }, [fetchedHistoricalActivity, isUserHistoricalTransactionsLoading])

  useEffect(() => {
    if (userHistoricalActivity.length > 0) {
      dispatch(updateUserHistoricalTransactions(userHistoricalActivity))
    }
  }, [userHistoricalActivity])

  useEffect(() => {
    if (userPendingActivity.length > 0) {
      dispatch(updateUserPendingTransactions(userPendingActivity))
    }
  }, [userPendingActivity])

  const hasPendingTransactions: boolean = useMemo(
    () => userPendingTransactions.length > 0,
    [userPendingTransactions]
  )

  const hasHistoricalTransactions: boolean = useMemo(
    () => userHistoricalActivity.length > 0,
    [userHistoricalActivity]
  )

  return (
    <div data-test-id="activity">
      {hasPendingTransactions && (
        <ActivitySection title="Pending">
          <TransactionHeader transactionType={ActivityType.PENDING} />
          {userPendingTransactions.map((transaction: BridgeTransaction) => (
            <Transaction
              bridgeTransaction={transaction}
              transactionType={ActivityType.PENDING}
            />
          ))}
        </ActivitySection>
      )}
      {isUserHistoricalTransactionsLoading ? (
        <div className="text-[#A3A3C2]"> Loading activity... </div>
      ) : (
        <ActivitySection title="Recent">
          <TransactionHeader transactionType={ActivityType.RECENT} />
          {userHistoricalTransactions.map((transaction: BridgeTransaction) => (
            <Transaction
              bridgeTransaction={transaction}
              transactionType={ActivityType.RECENT}
            />
          ))}
          <ExplorerLink />
        </ActivitySection>
      )}
    </div>
  )
}

export const ExplorerLink = () => {
  return (
    <div data-test-id="explorer-link" className="text-[#99E6FF] my-3">
      <Link href={ANALYTICS_PATH} target="_blank">
        Explorer →
      </Link>
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
    <div data-test-id="activity-section" className="py-6">
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
        {transactionType === ActivityType.PENDING && 'Block Initiated'}
        {transactionType === ActivityType.RECENT && 'Rate'}
      </div>
      <div className="flex justify-end col-span-2">
        {transactionType === ActivityType.PENDING && 'Elapsed'}
        {transactionType === ActivityType.RECENT && 'Completed'}
      </div>
    </div>
  )
}

export const getExplorerLink = ({
  kappa,
  fromChainId,
  toChainId,
}: {
  kappa: string
  fromChainId: number
  toChainId: number
}): string => {
  return `${ANALYTICS_KAPPA}${kappa}?chainIdFrom=${fromChainId}&chainIdTo=${toChainId}`
}

export const Transaction = ({
  bridgeTransaction,
  transactionType,
}: {
  bridgeTransaction: BridgeTransaction
  transactionType: ActivityType
}) => {
  const router = useRouter()
  const {
    fromInfo,
    toInfo,
    kappa,
  }: { fromInfo?: PartialInfo; toInfo?: PartialInfo; kappa?: string } =
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

  const handleTransactionClick: () => void = useCallback(() => {
    if (kappa && originChainId) {
      const explorerLink = getExplorerLink({
        kappa,
        fromChainId: originChainId,
        toChainId: destinationChainId,
      })
      window.open(explorerLink, '_blank')
    }
  }, [kappa, originChainId, destinationChainId])

  return (
    <div
      data-test-id="transaction"
      className={`
        grid grid-cols-10 mt-auto py-3
        text-sm text-white border-b border-[#565058]
        items-end hover:cursor-pointer hover:bg-[#272731]
        `}
      onClick={handleTransactionClick}
    >
      <div className="flex col-span-3">
        <TransactionPayloadDetail
          chain={originChain}
          token={originToken}
          tokenSymbol={originTokenSymbol}
          tokenAmount={originFormattedValue}
        />
        <div className="flex items-end mb-[3px] ml-auto px-4">→</div>
      </div>
      <div className="col-span-3">
        <TransactionPayloadDetail
          chain={destinationChain}
          token={destinationToken}
          tokenSymbol={destinationTokenSymbol}
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
        {transactionType === ActivityType.PENDING && (
          <div>{bridgeOriginBlockNumber}</div>
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

export const TimeElapsed = ({ startTime }: { startTime: number }) => {
  const [elapsedTime, setElapsedTime] = useState(0)

  useEffect(() => {
    const interval = setInterval(() => {
      const currentTime = Math.floor(Date.now() / 1000)
      const elapsedSeconds = currentTime - startTime
      setElapsedTime(elapsedSeconds)
    }, 1000)

    return () => {
      clearInterval(interval)
    }
  }, [startTime])

  const hours = Math.floor(elapsedTime / 3600)
  const minutes = Math.floor((elapsedTime % 3600) / 60)
  const seconds = elapsedTime % 60

  return (
    <div>
      {hours > 0 ? `${hours}:` : ''}
      {minutes}:{seconds}
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
  const exchangeRate: number = destinationValue / originValue
  const formattedExchangeRate: string = exchangeRate.toFixed(4)
  return (
    <span data-test-id="exchange-rate">
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
    <span data-test-id="completed">
      <span className="w-4 pt-3 mb-auto font-bold text-green-500"> ✓ </span>{' '}
      {formattedTime}
    </span>
  )
}

export const TransactionPayloadDetail = ({
  chain,
  token,
  tokenSymbol,
  tokenAmount,
}: {
  chain?: Chain
  token?: Token
  tokenSymbol?: string
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
          <div>{tokenSymbol}</div>
        </div>
      )}
    </div>
  )
}
