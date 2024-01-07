import { useMemo, useEffect, useState, useCallback } from 'react'
import { useAppDispatch } from '@/store/hooks'
import { useBridgeTransactionsState } from '@/slices/bridgeTransactions/hooks'
import { getTxBlockExplorerLink } from './helpers/getTxBlockExplorerLink'
import { getExplorerAddressLink } from './helpers/getExplorerAddressLink'
import { useBridgeTxStatus } from './helpers/useBridgeTxStatus'
import { isNull } from 'lodash'
import { DownArrow } from '../icons/DownArrow'
import {
  updateTransactionKappa,
  completeTransaction,
  removeTransaction,
} from '@/slices/bridgeTransactions/reducer'

const TransactionStatus = ({ string }) => {
  return <>{string}</>
}

const TimeRemaining = ({
  isComplete,
  remainingTime,
  isDelayed,
}: {
  isComplete: boolean
  remainingTime: number
  isDelayed: boolean
}) => {
  if (isComplete) return

  if (isDelayed) {
    return <div>Waiting...</div>
  }

  return <div>{remainingTime} min</div>
}

/** TODO: Update naming after refactoring existing Activity / Transaction flow */
export const _Transaction = ({
  synapseSDK,
  connectedAddress,
  originChainId,
  destinationChainId,
  originTxHash,
  bridgeModuleName,
  estimatedTime,
  kappa,
  timestamp,
  currentTime,
  isComplete,
}: {
  synapseSDK: any
  connectedAddress: string
  originChainId: number
  destinationChainId: number
  originTxHash: string
  bridgeModuleName: string
  estimatedTime: number // in seconds
  kappa?: string
  timestamp: number
  currentTime: number
  isComplete: boolean
}) => {
  const dispatch = useAppDispatch()
  const transactions = useBridgeTransactionsState()

  const [originTxExplorerLink, originExplorerName] = getTxBlockExplorerLink(
    originChainId,
    originTxHash
  )
  const [destExplorerAddressLink, destExplorerName] = getExplorerAddressLink(
    destinationChainId,
    connectedAddress
  )

  const elapsedTime: number = currentTime - timestamp // in seconds
  const remainingTime: number = estimatedTime - elapsedTime
  const remainingTimeInMinutes: number = Math.ceil(remainingTime / 60) // add additional min for buffer

  const isEstimatedTimeReached: boolean = useMemo(() => {
    if (!currentTime || !estimatedTime || !timestamp) return false
    return currentTime - timestamp > estimatedTime
  }, [estimatedTime, currentTime, timestamp])

  const [isTxComplete, _kappa] = useBridgeTxStatus({
    synapseSDK,
    originChainId,
    destinationChainId,
    originTxHash,
    bridgeModuleName,
    kappa,
    checkStatus: isEstimatedTimeReached,
    currentTime: currentTime,
  })

  /** Update tx kappa when available */
  useEffect(() => {
    if (_kappa && originTxHash) {
      dispatch(updateTransactionKappa({ originTxHash, kappa: _kappa }))
    }
  }, [_kappa, dispatch])

  /** Update tx for completion */
  /** Check that we have not already marked tx as complete */
  useEffect(() => {
    const txKappa = kappa ?? _kappa

    if (isTxComplete && originTxHash && txKappa) {
      if (transactions[originTxHash].isComplete === false) {
        dispatch(completeTransaction({ originTxHash, kappa: txKappa }))
      }
    }
  }, [isTxComplete, dispatch, transactions])

  const handleClearTransaction = useCallback(() => {
    dispatch(removeTransaction({ originTxHash }))
  }, [dispatch])

  return (
    <div
      data-test-id="transaction"
      className={`
        flex flex-wrap-reverse gap-1 justify-end items-center pl-2.5 pr-1.5 py-1
        bg-[--synapse-surface]
        border border-solid border-[--synapse-border] rounded-md
      `}
    >
      {isComplete ? (
        <TransactionStatus string="Complete" />
      ) : (
        <TransactionStatus string="Pending" />
      )}
      <div className="flex items-center justify-end gap-2 grow">
        <TimeRemaining
          isComplete={isComplete}
          remainingTime={remainingTimeInMinutes}
          isDelayed={isEstimatedTimeReached}
        />

        <DropdownMenu>
          {!isNull(originTxExplorerLink) && (
            <MenuItem text={originExplorerName} link={originTxExplorerLink} />
          )}
          {!isNull(destExplorerAddressLink) && (
            <MenuItem text={destExplorerName} link={destExplorerAddressLink} />
          )}
          <MenuItem
            text="Contact Support"
            link="https://discord.gg/synapseprotocol"
          />
          {isComplete && (
            <MenuItem
              text="Clear Transaction"
              link={null}
              onClick={handleClearTransaction}
            />
          )}
        </DropdownMenu>
      </div>
    </div>
  )
}

export const DropdownMenu = ({ children }) => {
  const [open, setOpen] = useState<boolean>(false)

  const handleClick = () => {
    setOpen(!open)
  }

  return (
    <div className="relative">
      <div
        onClick={handleClick}
        className={`
          rounded w-5 h-[21px] flex place-items-center justify-center
          bg-[--synapse-select-bg]
          border border-solid border-[--synapse-border]
          hover:border-[--synapse-focus]
          cursor-pointer
        `}
      >
        <DownArrow />
      </div>

      {open && (
        <ul
          className={`
            absolute z-50 mt-1 p-0 bg-[--synapse-surface] border border-solid border-[--synapse-border] rounded shadow popover -right-1 list-none text-left text-sm
          `}
        >
          {children}
        </ul>
      )}
    </div>
  )
}

export const MenuItem = ({
  text,
  link,
  onClick,
}: {
  text: string
  link: string
  onClick?: () => any
}) => {
  return (
    <li
      className={`
      rounded cursor-pointer
      border border-solid border-transparent
      hover:border-[--synapse-focus]
      active:opacity-40
    `}
    >
      {onClick ? (
        <div
          onClick={onClick}
          className={`
            block pl-2 pr-3 py-2 whitespace-nowrap text-[--synapse-text-primary] no-underline after:content-['_↗'] after:text-xs after:text-[--synapse-secondary]
          `}
        >
          {text}
        </div>
      ) : (
        <a
          href={link ?? ''}
          onClick={onClick}
          target="_blank"
          rel="noreferrer"
          className={`
            block pl-2 pr-3 py-2 whitespace-nowrap text-[--synapse-text-primary] no-underline after:content-['_↗'] after:text-xs after:text-[--synapse-secondary]
          `}
        >
          {text}
        </a>
      )}
    </li>
  )
}
