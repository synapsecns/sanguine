import { useEffect, useMemo, useCallback, useState } from 'react'
import { useAccount, Address } from 'wagmi'
import { useRouter } from 'next/router'
import Link from 'next/link'
import Image from 'next/image'
import { useTransactionsState } from '@/slices/transactions/hooks'
import { PartialInfo, BridgeTransaction } from '@/slices/api/generated'
import {
  convertUnixTimestampToMonthAndDate,
  isTimestampToday,
} from '@/utils/time'
import { CHAINS_BY_ID } from '@/constants/chains'
import { Chain, Token } from '@/utils/types'
import {
  BRIDGABLE_TOKENS,
  tokenAddressToToken,
  tokenSymbolToToken,
} from '@/constants/tokens'
import { ANALYTICS_KAPPA, ANALYTICS_PATH } from '@/constants/urls'
import EtherscanIcon from '../icons/EtherscanIcon'
import { TransactionsState } from '@/slices/transactions/reducer'
import { PendingBridgeTransaction } from '@/slices/bridge/actions'
import { BridgeState } from '@/slices/bridge/reducer'
import { useBridgeState } from '@/slices/bridge/hooks'
import { getExplorerTxUrl } from '@/constants/urls'
import { PortfolioState } from '@/slices/portfolio/reducer'
import { usePortfolioState } from '@/slices/portfolio/hooks'
import { PortfolioTabs } from '@/slices/portfolio/actions'
import { shortenAddress } from '@/utils/shortenAddress'
import { BRIDGE_REQUIRED_CONFIRMATIONS } from '@/constants/bridge'
import {
  Transaction as UpdatedTransaction,
  TransactionType,
} from './Transaction'

export const Activity = ({ visibility }: { visibility: boolean }) => {
  const { address } = useAccount()
  const {
    userHistoricalTransactions,
    userPendingTransactions,
    isUserHistoricalTransactionsLoading,
    isUserPendingTransactionsLoading,
  }: TransactionsState = useTransactionsState()
  const { pendingBridgeTransactions }: BridgeState = useBridgeState()

  const hasPendingTransactions: boolean = useMemo(() => {
    if (userPendingTransactions && userPendingTransactions.length > 0)
      return true
    if (pendingBridgeTransactions && pendingBridgeTransactions.length > 0)
      return true
    return false
  }, [userPendingTransactions, pendingBridgeTransactions])

  const hasHistoricalTransactions: boolean = useMemo(
    () => userHistoricalTransactions && userHistoricalTransactions.length > 0,
    [userHistoricalTransactions]
  )
  const hasNoTransactions: boolean = useMemo(() => {
    return !hasPendingTransactions && !hasHistoricalTransactions
  }, [hasPendingTransactions, hasHistoricalTransactions, address])

  const isLoading: boolean =
    isUserHistoricalTransactionsLoading && isUserPendingTransactionsLoading

  return (
    <div
      data-test-id="activity"
      className={`${visibility ? 'block' : 'hidden'}`}
    >
      {!address && (
        <div className="text-[#C2C2D6]">
          Your pending and recent transactions will appear here.
        </div>
      )}

      {address && isLoading && (
        <div className="text-[#C2C2D6]">Loading activity...</div>
      )}

      {address && !isLoading && hasNoTransactions && (
        <div className="text-[#C2C2D6]">
          Your pending and recent transactions will appear here.
          <ExplorerLink connectedAddress={address} />
        </div>
      )}

      {address && !isLoading && hasPendingTransactions && (
        <ActivitySection title="Pending" twClassName="flex flex-col gap-2 mb-5">
          <PendingTransactionAwaitingIndexing />
          {userPendingTransactions &&
            userPendingTransactions.map((transaction: BridgeTransaction) => (
              <Transaction
                connectedAddress={address}
                bridgeTransaction={transaction}
                transactionType={ActivityType.PENDING}
                key={transaction.kappa}
              />
            ))}
        </ActivitySection>
      )}

      {/* {address && !isLoading && hasHistoricalTransactions && (
        <ActivitySection title="Recent">
          {userHistoricalTransactions &&
            userHistoricalTransactions
              // .slice(0, 7) //temporarily only show recent 5ß
              .map((transaction: BridgeTransaction) => (
                <Transaction
                  connectedAddress={address}
                  bridgeTransaction={transaction}
                  transactionType={ActivityType.RECENT}
                  key={transaction.kappa}
                />
              ))}
          <ExplorerLink connectedAddress={address} />
        </ActivitySection>
      )} */}

      {address && !isLoading && hasHistoricalTransactions && (
        <ActivitySection title="Recent">
          {userHistoricalTransactions &&
            userHistoricalTransactions
              // .slice(0, 7) //temporarily only show recent 5ß
              .map((transaction: BridgeTransaction) => {
                const {
                  address: destinationAddress,
                  chainID: originChainId,
                  destinationChainID: destinationChainId,
                  value: originRawValue,
                  formattedValue: originFormattedValue,
                  tokenAddress: originTokenAddress,
                  tokenSymbol: originTokenSymbol,
                  blockNumber: bridgeOriginBlockNumber,
                  time: bridgeOriginTime,
                  txnHash: originTxnHash,
                }: PartialInfo = transaction?.fromInfo

                const originChain: Chain = CHAINS_BY_ID[originChainId]
                const originToken: Token = tokenAddressToToken(
                  originChainId,
                  originTokenAddress
                )
                const {
                  value: destinationRawValue,
                  formattedValue: destinationFormattedValue,
                  tokenAddress: destinationTokenAddress,
                  tokenSymbol: destinationTokenSymbol,
                  blockNumber: bridgeDestinationBlockNumber,
                  time: bridgeDestinationTime,
                }: PartialInfo = transaction?.toInfo

                const destinationChain: Chain = CHAINS_BY_ID[destinationChainId]
                const destinationToken: Token = tokenAddressToToken(
                  destinationChainId,
                  destinationTokenAddress
                )
                return (
                  <UpdatedTransaction
                    connectedAddress={address}
                    destinationAddress={destinationAddress as Address}
                    originChain={originChain}
                    originToken={originToken}
                    originValue={originFormattedValue}
                    destinationChain={destinationChain}
                    destinationToken={destinationToken}
                    startedTimestamp={bridgeOriginTime}
                    completedTimestamp={bridgeDestinationTime}
                    transactionType={TransactionType.HISTORICAL}
                    key={transaction.kappa}
                  />
                )
              })}
          <ExplorerLink connectedAddress={address} />
        </ActivitySection>
      )}
    </div>
  )
}

export const MostRecentPendingTransaction = () => {
  const { pendingBridgeTransactions }: BridgeState = useBridgeState()
  const { userPendingTransactions }: TransactionsState = useTransactionsState()
  const { activeTab }: PortfolioState = usePortfolioState()

  let mostRecentPendingTransaction = null

  if (pendingBridgeTransactions && pendingBridgeTransactions.length > 0) {
    mostRecentPendingTransaction = pendingBridgeTransactions[0]
    return (
      <div className="relative mt-3">
        <div
          className={`
          border border-[#3D3D5C] rounded-md box-arrow relative
          ${activeTab !== PortfolioTabs.ACTIVITY ? 'block' : 'hidden'}
          `}
        >
          <RecentlyBridgedPendingTransaction
            recentlyBridgedTransaction={mostRecentPendingTransaction}
          />
        </div>
      </div>
    )
  } else if (userPendingTransactions && userPendingTransactions.length > 0) {
    mostRecentPendingTransaction = userPendingTransactions[0]
    return (
      <div className="relative mt-3">
        <div
          className={`
          border border-[#3D3D5C] rounded-md box-arrow
          ${activeTab !== PortfolioTabs.ACTIVITY ? 'block' : 'hidden'}
          `}
        >
          <Transaction
            bridgeTransaction={mostRecentPendingTransaction}
            transactionType={ActivityType.PENDING}
            key={mostRecentPendingTransaction.kappa}
          />
        </div>
      </div>
    )
  }
  return null
}

const RecentlyBridgedPendingTransaction = ({
  recentlyBridgedTransaction,
}: {
  recentlyBridgedTransaction: PendingBridgeTransaction
}) => {
  const [delayed, setDelayed] = useState<boolean>(false)
  const {
    originChain,
    originToken,
    originValue,
    destinationChain,
    destinationToken,
    transactionHash,
    timestamp,
  }: PendingBridgeTransaction = recentlyBridgedTransaction

  const handlePendingTransactionClick: () => void = useCallback(() => {
    if (transactionHash) {
      const explorerLink: string = getExplorerTxUrl({
        chainId: originChain.id,
        hash: transactionHash,
      })
      window.open(explorerLink, '_blank')
    }
  }, [transactionHash])

  return (
    <div
      data-test-id="recently-bridged-pending-transaction"
      className={`
      grid grid-cols-10 bg-[#1B1B29]
      py-3 px-2 text-sm text-white
      rounded-md hover:cursor-pointer
      `}
      onClick={handlePendingTransactionClick}
    >
      <div className="flex col-span-4 my-auto">
        <TransactionPayloadDetail
          chain={originChain}
          token={originToken}
          tokenSymbol={originToken?.symbol}
          tokenAmount={Number(originValue)}
        />
      </div>
      <div className="col-span-4 my-auto">
        <TransactionPayloadDetail chain={destinationChain} />
      </div>
      <div className="flex justify-end col-span-2 my-auto">
        <TimeElapsed
          startTime={timestamp}
          bridgeOriginChain={originChain}
          delayed={delayed}
          setDelayed={setDelayed}
        />
      </div>

      {delayed && (
        <div className="text-[#FFD966] text-sm mt-2 whitespace-nowrap">
          Confirmation taking longer than expected
        </div>
      )}
    </div>
  )
}

export const PendingTransactionAwaitingIndexing = () => {
  const { pendingBridgeTransactions }: BridgeState = useBridgeState()
  return (
    <>
      {pendingBridgeTransactions.map(
        (transaction: PendingBridgeTransaction) => (
          <RecentlyBridgedPendingTransaction
            recentlyBridgedTransaction={transaction}
          />
        )
      )}
    </>
  )
}

export const ExplorerLink = ({
  connectedAddress,
}: {
  connectedAddress?: Address | string
}) => {
  const explorerLink: string = connectedAddress
    ? `${ANALYTICS_PATH}address/${connectedAddress}`
    : ANALYTICS_PATH
  return (
    <div data-test-id="explorer-link" className="text-[#99E6FF] my-3">
      <Link href={explorerLink} target="_blank">
        <span className="hover:underline">Explorer</span> →
      </Link>
    </div>
  )
}

export const ActivitySection = ({
  title,
  children,
  twClassName,
}: {
  title: string
  children?: React.ReactNode
  twClassName?: string
}) => {
  return (
    <div data-test-id="activity-section" className={twClassName}>
      <h3 className="mb-2 text-xl text-white">{title}</h3>
      {children}
    </div>
  )
}

export enum ActivityType {
  PENDING,
  STUCK,
  RECENT,
}

export const getTransactionExplorerLink = ({
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
  connectedAddress,
  bridgeTransaction,
  transactionType,
}: {
  connectedAddress?: Address | string
  bridgeTransaction: BridgeTransaction
  transactionType: ActivityType
}) => {
  const [delayed, setDelayed] = useState<boolean>(false)
  const {
    fromInfo,
    toInfo,
    kappa,
  }: { fromInfo?: PartialInfo; toInfo?: PartialInfo; kappa?: string } =
    bridgeTransaction || {}

  const {
    address: destinationAddress,
    chainID: originChainId,
    destinationChainID: destinationChainId,
    value: originRawValue,
    formattedValue: originFormattedValue,
    tokenAddress: originTokenAddress,
    tokenSymbol: originTokenSymbol,
    blockNumber: bridgeOriginBlockNumber,
    time: bridgeOriginTime,
    txnHash: originTxnHash,
  }: PartialInfo = fromInfo || {}

  const originChain: Chain = CHAINS_BY_ID[originChainId]
  const originToken: Token = tokenAddressToToken(
    originChainId,
    originTokenAddress
  )

  const {
    value: destinationRawValue,
    formattedValue: destinationFormattedValue,
    tokenAddress: destinationTokenAddress,
    tokenSymbol: destinationTokenSymbol,
    blockNumber: bridgeDestinationBlockNumber,
    time: bridgeDestinationTime,
  }: PartialInfo = toInfo || {}

  const destinationChain: Chain = CHAINS_BY_ID[destinationChainId]
  const destinationToken: Token = tokenAddressToToken(
    destinationChainId,
    destinationTokenAddress
  )

  const handleTransactionClick: () => void = useCallback(() => {
    if (kappa && originChainId && transactionType === ActivityType.RECENT) {
      const explorerLink: string = getTransactionExplorerLink({
        kappa,
        fromChainId: originChainId,
        toChainId: destinationChainId,
      })
      window.open(explorerLink, '_blank')
    } else {
      const explorerLink: string = getExplorerTxUrl({
        chainId: originChainId,
        hash: originTxnHash,
      })
      window.open(explorerLink, '_blank')
    }
  }, [kappa, originChainId, destinationChainId, transactionType])

  const estimatedCompletionInSeconds: number =
    (BRIDGE_REQUIRED_CONFIRMATIONS[originChain.id] * originChain.blockTime) /
      1000 +
    30 // Add 30 seconds to account for indexing

  return (
    <div data-test-id="transaction" className="flex flex-col">
      <div
        onClick={handleTransactionClick}
        className={`
        grid grid-cols-10 mt-auto py-3 px-2
        text-sm text-white
        items-end hover:cursor-pointer
        ${
          transactionType === ActivityType.RECENT && 'border-b border-[#565058]'
        }
        ${transactionType === ActivityType.PENDING && 'bg-[#1B1B29] rounded-md'}
        `}
      >
        <div className="flex col-span-4 my-auto">
          <TransactionPayloadDetail
            chain={originChain}
            token={originToken}
            tokenSymbol={originToken?.symbol}
            tokenAmount={originFormattedValue}
          />
        </div>
        <div className="col-span-4 my-auto">
          <TransactionPayloadDetail
            chain={destinationChain}
            token={destinationToken}
            tokenSymbol={destinationToken?.symbol}
            tokenAmount={destinationFormattedValue}
          />
        </div>
        <div className="flex justify-end col-span-2 my-auto">
          {transactionType === ActivityType.RECENT && (
            <Completed
              transactionCompletedTime={bridgeDestinationTime}
              connectedAddress={connectedAddress}
              destinationAddress={destinationAddress}
            />
          )}
          {transactionType === ActivityType.PENDING && (
            <TimeElapsed
              startTime={bridgeOriginTime}
              bridgeOriginChain={originChain}
              delayed={delayed}
              setDelayed={setDelayed}
            />
          )}
        </div>
        {delayed && (
          <div className="text-[#FFD966] text-sm mt-2 whitespace-nowrap">
            Confirmation taking longer than expected
          </div>
        )}
      </div>
    </div>
  )
}

export const TimeElapsed = ({
  startTime,
  bridgeOriginChain,
  delayed,
  setDelayed,
}: {
  startTime: number
  bridgeOriginChain: Chain
  delayed: boolean
  setDelayed: React.Dispatch<React.SetStateAction<boolean>>
}) => {
  const [elapsedTime, setElapsedTime] = useState<number>(0)

  useEffect(() => {
    const interval = setInterval(() => {
      const currentTime: number = Math.floor(Date.now() / 1000)
      const elapsedSeconds: number = currentTime - startTime
      setElapsedTime(elapsedSeconds)
    }, 1000)

    return () => {
      clearInterval(interval)
    }
  }, [startTime])

  const hours: number = Math.floor(elapsedTime / 3600)
  const minutes: number = Math.floor((elapsedTime % 3600) / 60)
  const seconds: number = elapsedTime % 60

  const formattedMinutes: string = String(minutes).padStart(2, '0')
  const formattedSeconds: string = String(seconds).padStart(2, '0')

  const estimatedCompletionInSeconds: number =
    (BRIDGE_REQUIRED_CONFIRMATIONS[bridgeOriginChain.id] *
      bridgeOriginChain.blockTime) /
      1000 +
    30 // Add 30 seconds to account for indexing

  const estimatedMinutes = Math.floor(estimatedCompletionInSeconds / 60)
  const estimatedSeconds = estimatedCompletionInSeconds % 60

  const formattedEstimatedMinutes = String(estimatedMinutes).padStart(2, '0')
  const formattedEstimatedSeconds = String(estimatedSeconds).padStart(2, '0')

  const estimatedCompletionTime: string = useMemo(() => {
    const firstDelayTimeInSeconds: number = 60 * 15 // 15 minutes
    const secondDelayTimeInSeconds: number = 60 * 25 // 25 minutes
    if (
      elapsedTime > estimatedCompletionInSeconds &&
      elapsedTime < firstDelayTimeInSeconds
    ) {
      return `15:00`
    } else if (
      elapsedTime > estimatedCompletionInSeconds &&
      elapsedTime > secondDelayTimeInSeconds
    ) {
      return `25:00`
    } else return `${formattedEstimatedMinutes}:${formattedEstimatedSeconds}`
  }, [estimatedCompletionInSeconds, elapsedTime])

  useEffect(() => {
    if (!delayed && elapsedTime > estimatedCompletionInSeconds) {
      setDelayed(true)
    }
  }, [delayed, estimatedCompletionInSeconds, elapsedTime, setDelayed])

  return (
    <div
      data-test-id="time-elapsed"
      className="flex items-center whitespace-nowrap"
    >
      {hours > 0 ? `${hours}:` : ''}
      {formattedMinutes}:{formattedSeconds} / {estimatedCompletionTime}
      <EtherscanIcon className="ml-1" />
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

export const EstimatedDuration = ({
  estimatedCompletionInSeconds,
}: {
  estimatedCompletionInSeconds: number
}) => {
  const estimatedMinutes = Math.floor(estimatedCompletionInSeconds / 60)
  return (
    <div className="text-[#C2C2D6] text-sm">
      {estimatedMinutes}-{estimatedMinutes + 1} min
    </div>
  )
}

export const Completed = ({
  transactionCompletedTime,
  connectedAddress,
  destinationAddress,
}: {
  transactionCompletedTime: number
  connectedAddress?: Address | string
  destinationAddress: string
}) => {
  const formattedTime: string =
    transactionCompletedTime &&
    convertUnixTimestampToMonthAndDate(transactionCompletedTime)

  const isToday: boolean = isTimestampToday(transactionCompletedTime)

  const destinationIsSender: boolean =
    String(connectedAddress) === String(destinationAddress)

  return (
    <div
      data-test-id="completed"
      className="flex flex-col text-right text-[#C2C2D6] gap-1"
    >
      {!destinationIsSender && (
        <div>to {shortenAddress(destinationAddress, 3)} </div>
      )}
      {isToday ? (
        <div className="text-[#3BDD77]">Today</div>
      ) : (
        <div>{formattedTime}</div>
      )}
    </div>
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
            className="w-4 h-4 mr-3 rounded-full"
            alt={`${chain.name} icon`}
          />
          <div>{chain.name}</div>
        </div>
      )}

      {tokenSymbol && (
        <div
          data-test-id="transaction-payload-token"
          className="flex flex-row items-center"
        >
          <Image
            src={token?.icon}
            className="items-center w-4 h-4 mr-3 rounded-full"
            alt={`${token?.name} icon`}
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
