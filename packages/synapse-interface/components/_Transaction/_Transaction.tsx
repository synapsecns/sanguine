import { useCallback } from 'react'
import { isNull } from 'lodash'
import { useAppDispatch } from '@/store/hooks'
import { getTxBlockExplorerLink } from './helpers/getTxBlockExplorerLink'
import { getExplorerAddressLink } from './helpers/getExplorerAddressLink'
import { useBridgeTxStatus } from './helpers/useBridgeTxStatus'
import { removeTransaction } from '@/slices/_transactions/reducer'
import { TransactionPayloadDetail } from '@/components/Portfolio/Transaction/components/TransactionPayloadDetail'
import { Chain, Token } from '@/utils/types'

import { TimeRemaining } from './components/TimeRemaining'

import { getEstimatedTimeStatus } from './helpers/getEstimatedTimeStatus'
import { DropdownMenu } from './components/DropdownMenu'
import { MenuItem } from './components/MenuItem'
import { useBridgeTxUpdater } from './helpers/useBridgeTxUpdater'
import { AnimatedProgressBar } from './components/AnimatedProgressBar'
import { TransactionSupport } from './components/TransactionSupport'

import TransactionArrow from '../icons/TransactionArrow'

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
    isStartCheckingTimeReached,
  } = getEstimatedTimeStatus(currentTime, timestamp, estimatedTime)

  const [isTxComplete, _kappa] = useBridgeTxStatus({
    originChainId: originChain.id,
    destinationChainId: destinationChain.id,
    originTxHash,
    bridgeModuleName,
    kappa: kappa,
    checkStatus: !isStoredComplete || isStartCheckingTimeReached,
    currentTime: currentTime,
  })

  /** Check if store already marked tx as complete, otherwise check hook status */
  const isTxFinalized = isStoredComplete ?? isTxComplete

  const showTransactionSupport =
    !isTxFinalized && delayedTimeInMin ? delayedTimeInMin <= -5 : false

  useBridgeTxUpdater(
    connectedAddress,
    destinationChain,
    _kappa,
    originTxHash,
    isTxComplete
  )
  console.log(kappa)
  console.log(isTxFinalized)
  return (
    <div
      data-test-id="_transaction"
      className={`
        bg-bgBase/10 fill-bgBase/20
        border border-white/10 rounded-md
        text-primary text-xs md:text-base
      `}
    >
      <div
        className={`
          flex items-center pr-2
          ${showTransactionSupport ? 'pb-0' : ''}
        `}
      >
        <TransactionPayloadDetail
          chain={originChain}
          token={originToken}
          tokenAmount={originValue}
          isOrigin={true}
          showChain={true}
          className="py-2 px-2 bg-bgBase/10 rounded-l-md"
        />
        <TransactionArrow
          className={`
              mr-2
              ${isTxFinalized
              ? 'stroke-white/10 fill-bgBase/10 '
              : 'stroke-white/10 fill-transparent'}
          `}
        />
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
                remainingTime={remainingTime}
                isDelayed={isEstimatedTimeReached}
                delayedTime={delayedTime}
              />
            }
          >
            <div className="p-2 text-xs cursor-default italic text-slate-300/90">
              Began{' '}
              {new Date(timestamp * 1000).toLocaleString('en-US', {
                month: 'short',
                day: 'numeric',
                hour: '2-digit',
                minute: '2-digit',
                hour12: true,
              })}
              {/* <div>{typeof _kappa === 'string' && _kappa?.substring(0, 15)}</div> */}
            </div>
            {!isNull(originTxExplorerLink) && (
              <MenuItem text={originExplorerName} link={originTxExplorerLink} />
            )}
            {!isNull(destExplorerAddressLink) && (
              <MenuItem
                text={destExplorerName}
                link={destExplorerAddressLink}
              />
            )}
            <MenuItem
              text="Contact Support"
              link="https://discord.gg/synapseprotocol"
            />
            {isTxFinalized && (
              <MenuItem
                text="Clear Transaction"
                link={null}
                onClick={handleClearTransaction}
              />
            )}
          </DropdownMenu>
        </div>
      </div>
      {showTransactionSupport && <TransactionSupport />}
      <div
        className={`
          px-1
          h-0 overflow-visible
          ${isTxFinalized ? "" : "-mt-[3px]"}
        `}
      >
        <AnimatedProgressBar
          id={originTxHash}
          startTime={timestamp}
          estDuration={estimatedTime * 2} // 2x buffer
          isComplete={isTxFinalized}
        />
      </div>
    </div>
  )
}
