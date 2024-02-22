import { useCallback } from 'react'
import { useAppDispatch } from '@/store/hooks'
import { getTxBlockExplorerLink } from './helpers/getTxBlockExplorerLink'
import { getExplorerAddressLink } from './helpers/getExplorerAddressLink'
import { useBridgeTxStatus } from './helpers/useBridgeTxStatus'
import { isNull } from 'lodash'
import { removeTransaction } from '@/slices/_transactions/reducer'
import { TransactionPayloadDetail } from '../Portfolio/Transaction/components/TransactionPayloadDetail'
import { Chain, Token } from '@/utils/types'
import { TimeRemaining } from './components/TimeRemaining'
import { getEstimatedTimeStatus } from './helpers/getEstimatedTimeStatus'
import { DropdownMenu } from './components/DropdownMenu'
import { MenuItem } from './components/MenuItem'
import { useBridgeTxUpdater } from './helpers/useBridgeTxUpdater'
import { AnimatedProgressBar } from './components/AnimatedProgressBar'
import { TransactionSupport } from './components/TransactionSupport'
import { RightArrow } from '@/components/icons/RightArrow'
import { Address } from 'viem'
import { useIsTxReverted } from './helpers/useIsTxReverted'

interface _TransactionProps {
  connectedAddress: string
  originValue: number
  originChain: Chain
  originToken: Token
  destinationChain: Chain
  destinationToken: Token
  originTxHash: string
  bridgeModuleName: string
  estimatedTime: number // in seconds
  timestamp: number
  currentTime: number
  kappa?: string
  isStoredComplete: boolean
  isStoredReverted: boolean
  status: 'pending' | 'complete' | 'reverted'
}

/** TODO: Update naming after refactoring existing Activity / Transaction flow */
export const _Transaction = ({
  connectedAddress,
  originValue,
  originChain,
  originToken,
  destinationChain,
  destinationToken,
  originTxHash,
  bridgeModuleName,
  estimatedTime,
  timestamp,
  currentTime,
  kappa,
  isStoredComplete,
  isStoredReverted,
  status,
}: _TransactionProps) => {
  const dispatch = useAppDispatch()

  const handleClearTransaction = useCallback(() => {
    dispatch(removeTransaction({ originTxHash }))
  }, [dispatch])

  const [originTxExplorerLink, originExplorerName] = getTxBlockExplorerLink(
    originChain.id,
    originTxHash
  )
  const [destExplorerAddressLink, destExplorerName] = getExplorerAddressLink(
    destinationChain.id,
    connectedAddress
  )

  const {
    targetTime,
    elapsedTime,
    remainingTime,
    delayedTime,
    delayedTimeInMin,
    isEstimatedTimeReached,
    isCheckTxStatus,
    isCheckTxForRevert,
  } = getEstimatedTimeStatus(currentTime, timestamp, estimatedTime)

  const [isTxComplete, _kappa] = useBridgeTxStatus({
    originChainId: originChain.id,
    destinationChainId: destinationChain.id,
    originTxHash,
    bridgeModuleName,
    kappa: kappa,
    checkStatus: isCheckTxStatus && status === 'pending',
    currentTime: currentTime,
  })
  const isTxFinalized = status === 'complete' || isTxComplete

  const isReverted = useIsTxReverted(
    originTxHash as Address,
    originChain,
    isCheckTxForRevert && status === 'pending'
  )
  const isTxReverted = status === 'reverted' || isReverted

  useBridgeTxUpdater(
    connectedAddress,
    destinationChain,
    _kappa,
    originTxHash,
    isTxComplete,
    isReverted
  )

  // Show transaction support if the transaction is delayed by more than 5 minutes and not finalized or reverted
  const showTransactionSupport =
    isTxReverted ||
    (!isTxFinalized && delayedTimeInMin && delayedTimeInMin <= -5)

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
        />
        <RightArrow className="stroke-secondaryTextColor mt-0.5 mx-1" />
        <TransactionPayloadDetail
          chain={destinationChain}
          token={destinationToken}
          tokenAmount={null}
          isOrigin={false}
        />
        <div className="flex items-center justify-end gap-2 grow">
          <DropdownMenu
            menuTitleElement={
              <TimeRemaining
                isComplete={isTxFinalized}
                isDelayed={isEstimatedTimeReached}
                isReverted={isStoredReverted}
                remainingTime={remainingTime}
                delayedTime={delayedTime}
              />
            }
          >
            <div className="p-2 mt-1 text-xs cursor-default text-zinc-300">
              Began{' '}
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
            {!isNull(destExplorerAddressLink) && !isTxReverted && (
              <MenuItem
                text={destExplorerName}
                link={destExplorerAddressLink}
                iconUrl={destinationChain?.explorerImg}
              />
            )}
            <MenuItem
              text="Contact Support (Discord)"
              link="https://discord.gg/synapseprotocol"
            />
            {(isTxFinalized || isTxReverted) && (
              <MenuItem
                text={isTxReverted ? 'Clear notification' : 'Clear transaction'}
                link={null}
                onClick={handleClearTransaction}
              />
            )}
          </DropdownMenu>
        </div>
      </div>
      {showTransactionSupport && (
        <TransactionSupport isReverted={isTxReverted} />
      )}
      <div className="px-1">
        <AnimatedProgressBar
          id={originTxHash}
          startTime={timestamp}
          estDuration={estimatedTime * 2} // 2x buffer
          isComplete={isTxFinalized}
          isError={isStoredReverted}
        />
      </div>
    </div>
  )
}
