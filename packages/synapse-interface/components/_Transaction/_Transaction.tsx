import { useCallback } from 'react'
import { useTranslations } from 'next-intl'

import { useAppDispatch } from '@/store/hooks'
import { getTxBlockExplorerLink } from './helpers/getTxBlockExplorerLink'
import { getExplorerAddressLink } from './helpers/getExplorerAddressLink'
import { useBridgeTxStatus } from './helpers/useBridgeTxStatus'
import { isNull } from 'lodash'
import { removeTransaction } from '@/slices/_transactions/reducer'
import { TransactionPayloadDetail } from '../Activity/Transaction/components/TransactionPayloadDetail'
import { Chain, Token } from '@/utils/types'
import { TimeRemaining } from './components/TimeRemaining'
import { calculateEstimatedTimeStatus } from './helpers/calculateEstimatedTimeStatus'
import { DropdownMenu } from './components/DropdownMenu'
import { MenuItem } from './components/MenuItem'
import { useBridgeTxUpdater } from './helpers/useBridgeTxUpdater'
import { AnimatedProgressBar } from './components/AnimatedProgressBar'
import { TransactionSupport } from './components/TransactionSupport'
import { RightArrow } from '@/components/icons/RightArrow'
import { Address } from 'viem'
import { useIsTxReverted } from './helpers/useIsTxReverted'
import { useTxRefundStatus } from './helpers/useTxRefundStatus'
import { HYPERLIQUID } from '@/constants/chains/master'

interface _TransactionProps {
  connectedAddress: string
  destinationAddress: Address | null
  originValue: number
  originChain: Chain
  originToken: Token
  destinationChain: Chain
  destinationToken: Token
  originTxHash: string
  bridgeModuleName: string
  routerAddress: string
  estimatedTime: number // in seconds
  timestamp: number
  currentTime: number
  kappa?: string
  status: 'pending' | 'completed' | 'reverted' | 'refunded'
  disabled: boolean
}

/** TODO: Update naming after refactoring existing Activity / Transaction flow */
export const _Transaction = ({
  connectedAddress,
  destinationAddress,
  originValue,
  originChain,
  originToken,
  destinationChain,
  destinationToken,
  originTxHash,
  bridgeModuleName,
  routerAddress,
  estimatedTime,
  timestamp,
  currentTime,
  kappa,
  status,
  disabled,
}: _TransactionProps) => {
  const dispatch = useAppDispatch()

  const t = useTranslations('Time')

  const handleClearTransaction = useCallback(() => {
    dispatch(removeTransaction({ originTxHash }))
  }, [dispatch])

  const [originTxExplorerLink, originExplorerName] = getTxBlockExplorerLink(
    originChain?.id,
    originTxHash
  )
  const [destExplorerAddressLink, destExplorerName] = getExplorerAddressLink(
    destinationChain?.id,
    destinationAddress ?? connectedAddress
  )

  const {
    remainingTime,
    delayedTime,
    delayedTimeInMin,
    isEstimatedTimeReached,
    isCheckTxStatus,
    isCheckTxForRevert,
    isCheckTxForRefund,
  } = calculateEstimatedTimeStatus(currentTime, timestamp, estimatedTime)

  const [isTxCompleted, _kappa] = useBridgeTxStatus({
    originChainId: originChain?.id,
    destinationChainId: destinationChain?.id,
    originTxHash,
    bridgeModuleName,
    kappa: kappa,
    checkStatus: isCheckTxStatus && status === 'pending',
    currentTime: currentTime,
  })

  const isTxReverted = useIsTxReverted(
    originTxHash as Address,
    originChain,
    isCheckTxForRevert && status === 'pending'
  )

  const isTxRefunded = useTxRefundStatus(
    kappa,
    routerAddress as Address,
    originChain,
    isCheckTxForRefund &&
      status === 'pending' &&
      bridgeModuleName === 'SynapseRFQ'
  )

  useBridgeTxUpdater(
    connectedAddress,
    destinationChain,
    _kappa,
    originTxHash,
    isTxCompleted,
    isTxReverted,
    isTxRefunded
  )

  // Show transaction support if the transaction is delayed by more than 5 minutes and not finalized or reverted
  const showTransactionSupport =
    status === 'reverted' ||
    status === 'refunded' ||
    (status === 'pending' && delayedTimeInMin && delayedTimeInMin <= -5)

  return (
    <div
      data-test-id="_transaction"
      className={`
        border border-surface rounded-md bg-tint
        text-primary text-xs md:text-base
      `}
    >
      <div
        className={`
          flex items-center px-1 pt-2
          ${showTransactionSupport ? 'pb-0' : 'pb-2'}
        `}
      >
        <TransactionPayloadDetail
          chain={originChain}
          token={originToken}
          tokenAmount={originValue}
          isOrigin={true}
          showChain={false}
          disabled={disabled}
        />
        <RightArrow className="stroke-secondaryTextColor mt-0.5 mx-1" />
        <TransactionPayloadDetail
          chain={destinationChain}
          token={destinationToken}
          tokenAmount={null}
          isOrigin={false}
          disabled={disabled}
        />
        <div className="flex items-center justify-end gap-2 grow">
          <DropdownMenu
            menuTitleElement={
              <TimeRemaining
                status={status}
                isDelayed={isEstimatedTimeReached}
                remainingTime={remainingTime}
                delayedTime={delayedTime}
              />
            }
          >
            <div className="p-2 mt-1 text-xs cursor-default text-zinc-300">
              {t('Began')}{' '}
              {new Date(timestamp * 1000).toLocaleString('en-US', {
                month: 'short',
                day: 'numeric',
                hour: '2-digit',
                minute: '2-digit',
                hour12: true,
              })}
            </div>
            {!isNull(originTxExplorerLink) && (
              <MenuItem
                text={originExplorerName}
                link={originTxExplorerLink}
                iconUrl={originChain?.explorerImg}
              />
            )}
            {destinationChain.id !== HYPERLIQUID.id &&
              !isNull(destExplorerAddressLink) &&
              !isTxReverted && (
                <MenuItem
                  text={destExplorerName}
                  link={destExplorerAddressLink}
                  iconUrl={destinationChain?.explorerImg}
                />
              )}
            <MenuItem
              text={t('Contact Support (Discord)')}
              link="https://discord.gg/synapseprotocol"
            />
            {status !== 'pending' && (
              <MenuItem
                text={
                  isTxReverted || isTxRefunded
                    ? t('Clear notification')
                    : t('Clear transaction')
                }
                link={null}
                onClick={handleClearTransaction}
              />
            )}
          </DropdownMenu>
        </div>
      </div>
      {showTransactionSupport ? <TransactionSupport status={status} /> : null}

      <div className="px-1">
        <AnimatedProgressBar
          id={originTxHash}
          startTime={timestamp}
          estDuration={estimatedTime * 2} // 2x buffer
          status={status}
        />
      </div>
    </div>
  )
}
