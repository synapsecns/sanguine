import { useCallback } from 'react'
import { useAppDispatch } from '@/store/hooks'
import { getTxBlockExplorerLink } from './helpers/getTxBlockExplorerLink'
import { getExplorerAddressLink } from './helpers/getExplorerAddressLink'
import { useBridgeTxStatus } from './helpers/useBridgeTxStatus'
import { isNull } from 'lodash'
import { removeTransaction } from '@/slices/_transactions/reducer'
import { TransactionPayloadDetail } from '../Portfolio/Transaction/components/TransactionPayloadDetail'
import { Chain, Token } from '@/utils/types'
import TransactionArrow from '../icons/TransactionArrow'
import { TimeRemaining } from './components/TimeRemaining'
import { TransactionStatus } from './components/TransactionStatus'
import { getEstimatedTimeStatus } from './helpers/getEstimatedTimeStatus'
// import { DropdownMenu } from './components/DropdownMenu'
import { MenuItem } from './components/MenuItem'
import { useBridgeTxUpdater } from './helpers/useBridgeTxUpdater'
import { AnimatedProgressBar } from './components/AnimatedProgressBar'
import { useState } from 'react'
import { DownArrow } from '@/components/icons/DownArrow'

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

  // useBridgeTxUpdater(
  //   connectedAddress,
  //   destinationChain,
  //   _kappa,
  //   originTxHash,
  //   isTxComplete
  // )

  const [open, setOpen] = useState<boolean>(false)

  /* TODO: Fix bug where e.stopProgagation() allows multiple dropdowns
           to be open at the same time
  */
  const handleClick = (e) => {
    setOpen(!open)

    if (!open) {
      document.addEventListener('click', (e) => {
        setOpen(false)
      }, { once: true })
    }

    e.stopPropagation()
  }

  return (
    <div
      data-test-id="_transaction"
      className={`
        my-2
        bg-tint text-primary
        border border-surface rounded-md
        text-xs md:text-base
      `}
    >
      <div className="flex items-center">
        <TransactionPayloadDetail
          // chain={originChain}
          token={originToken}
          tokenAmount={originValue}
          isOrigin={true}
          className='bg-surface px-0.5 py-1.5 rounded-l'
        />
        <TransactionArrow className="fill-surface" />
        <TransactionPayloadDetail
          chain={destinationChain}
          token={destinationToken}
          tokenAmount={null}
          isOrigin={false}
          className='p-1.5'
        />
        {/* TODO: QA visual format */}
        <div className="relative grow mr-1 text-sm">
          <div
            onClick={handleClick}
            className="flex items-center gap-1 hover:bg-zinc-700 cursor-pointer relative w-fit ml-auto px-2 py-1 rounded"
          >
            <TimeRemaining
              isComplete={isTxFinalized as boolean}
              remainingTime={remainingTime}
              isDelayed={isEstimatedTimeReached}
            />
            <DownArrow />
          </div>
          {open && (
            <div className='
              absolute z-50 mt-1 bg-surface
              border border-zinc-700 rounded shadow
              popover right-0
            '>
              <div className="text-xs p-2 mt-1 text-zinc-300 cursor-default">
                Began {new Date(timestamp * 1000).toLocaleString('en-US', {
                  // month: 'short',
                  // day: 'numeric',
                  hour: '2-digit',
                  minute: '2-digit',
                  hour12: true,
                })}
                {/* <div>{typeof _kappa === 'string' && _kappa?.substring(0, 15)}</div> */}
              </div>
              {!isNull(originTxExplorerLink) && (
                <MenuItem
                  text={originExplorerName}
                  link={originTxExplorerLink}
                />
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
            </div>
          )}
        </div>
      </div>
      <div className='px-1'>
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

