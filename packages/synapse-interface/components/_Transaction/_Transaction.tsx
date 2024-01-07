import { useMemo, useEffect, useState, useCallback } from 'react'
import { useAppDispatch } from '@/store/hooks'
import { use_TransactionsState } from '@/slices/_transactions/hooks'
import { getTxBlockExplorerLink } from './helpers/getTxBlockExplorerLink'
import { getExplorerAddressLink } from './helpers/getExplorerAddressLink'
import { useBridgeTxStatus } from './helpers/useBridgeTxStatus'
import { isNull } from 'lodash'
import { DownArrow } from '../icons/DownArrow'
import {
  updateTransactionKappa,
  completeTransaction,
  removeTransaction,
} from '@/slices/_transactions/reducer'
import { TransactionPayloadDetail } from '../Portfolio/Transaction/components/TransactionPayloadDetail'
import { Chain, Token } from '@/utils/types'
import TransactionArrow from '../icons/TransactionArrow'

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

interface _TransactionProps {
  synapseSDK: any
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
  // isComplete: boolean
}

/** TODO: Update naming after refactoring existing Activity / Transaction flow */
export const _Transaction = ({
  synapseSDK,
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
}: // isComplete,
_TransactionProps) => {
  const dispatch = useAppDispatch()
  const { transactions } = use_TransactionsState()

  const [originTxExplorerLink, originExplorerName] = getTxBlockExplorerLink(
    originChain.id,
    originTxHash
  )
  const [destExplorerAddressLink, destExplorerName] = getExplorerAddressLink(
    destinationChain.id,
    connectedAddress
  )

  const elapsedTime: number = currentTime - timestamp // in seconds
  const remainingTime: number = estimatedTime - elapsedTime
  const remainingTimeInMinutes: number = Math.ceil(remainingTime / 60) // add additional min for buffer

  const isEstimatedTimeReached: boolean = useMemo(() => {
    // Define the interval in minutes before the estimated completion when we should start checking
    const intervalBeforeCompletion = 1 // X minutes before completion
    // Calculate the time in seconds when we should start checking
    const startCheckingTime =
      currentTime + estimatedTime - intervalBeforeCompletion * 60

    // if current time is above startCheckingTime, return true to begin calling the SDK
    return currentTime >= startCheckingTime

    // TODO: OLD CODE BELOW:
    // if (!currentTime || !estimatedTime || !timestamp) return false
    // return currentTime - timestamp > estimatedTime
  }, [estimatedTime, currentTime, timestamp])

  const [isTxComplete, _kappa] = useBridgeTxStatus({
    synapseSDK,
    originChainId: originChain.id,
    destinationChainId: destinationChain.id,
    originTxHash,
    bridgeModuleName,
    kappa: kappa,
    checkStatus: isEstimatedTimeReached,
    currentTime: currentTime,
  })

  /** Update tx kappa when available */
  useEffect(() => {
    if (_kappa && originTxHash) {
      dispatch(
        updateTransactionKappa({ originTxHash, kappa: _kappa as string })
      )
    }
  }, [_kappa, dispatch])

  /** Update tx for completion */
  /** Check that we have not already marked tx as complete */
  useEffect(() => {
    const txKappa = _kappa

    if (isTxComplete && originTxHash && txKappa) {
      const txn = transactions.find((tx) => tx.originTxHash === originTxHash)
      if (!txn.isComplete) {
        dispatch(
          completeTransaction({ originTxHash, kappa: txKappa as string })
        )
      }
    }
  }, [isTxComplete, dispatch, transactions, _kappa])

  const handleClearTransaction = useCallback(() => {
    dispatch(removeTransaction({ originTxHash }))
  }, [dispatch])

  return (
    <div
      data-test-id="_transaction"
      className={`
        flex flex-col gap-1 justify-end items-center my-2
        bg-tint fill-surface text-primary
        border border-solid border-surface rounded-md
      `}
    >
      <div className="flex items-center w-full">
        <div className="flex rounded bg-surface fill-surface">
          <TransactionPayloadDetail
            chain={originChain}
            token={originToken}
            tokenAmount={originValue}
            isOrigin={true}
          />
          <TransactionArrow className="bg-tint fill-surface" />
        </div>
        <TransactionPayloadDetail
          chain={destinationChain}
          token={destinationToken}
          tokenAmount={null}
          isOrigin={false}
        />
        <div>
          <div className="text-xs">
            {new Date(timestamp * 1000).toLocaleString('en-US', {
              hour: '2-digit',
              minute: '2-digit',
              second: '2-digit',
              hour12: true,
            })}
          </div>
          <div>{typeof _kappa === 'string' && _kappa?.substring(0, 15)}</div>
        </div>
        {/* TODO: Update visual format */}
        <div className="flex justify-between gap-2 pr-2 ml-auto">
          {isTxComplete ? (
            <TransactionStatus string="Complete" />
          ) : (
            <TransactionStatus string="Pending" />
          )}
          <div className="flex items-center justify-end gap-2 grow">
            <TimeRemaining
              isComplete={isTxComplete as boolean}
              remainingTime={remainingTimeInMinutes}
              isDelayed={isEstimatedTimeReached}
            />

            <DropdownMenu>
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
              {isTxComplete && (
                <MenuItem
                  text="Clear Transaction"
                  link={null}
                  onClick={handleClearTransaction}
                />
              )}
            </DropdownMenu>
          </div>
        </div>
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
          bg-surface
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
            absolute z-50 mt-1 p-0 bg-surface border border-solid border-tint rounded shadow popover -right-1 list-none text-left text-sm
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
