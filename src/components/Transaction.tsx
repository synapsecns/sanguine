import { useState, useEffect, useMemo, useCallback } from 'react'
import { useAppDispatch } from '@/state/hooks'
import { DownArrow } from './icons/DownArrow'
import { getTxBlockExplorerLink } from '@/utils/getTxBlockExplorerLink'
import { getExplorerAddressUrl } from '@/utils/getExplorerAddressLink'
import { getTxSynapseExplorerLink } from '@/utils/getTxSynapseExplorerLink'
import { useBridgeTxStatus } from '@/hooks/useBridgeTxStatus'
import { isNull } from '@/utils/isNull'
import {
  updateTransactionKappa,
  completeTransaction,
  removeTransaction,
} from '@/state/slices/transactions/reducer'
import { useTransactionsState } from '@/state/slices/transactions/hooks'
import { useSynapseContext } from '@/providers/SynapseProvider'

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

  const estTime = useMemo(() => {
    if (remainingTime > 60) {
      return Math.ceil(remainingTime / 60) + ' minutes'
    } else {
      return remainingTime + ' seconds'
    }
  }, [remainingTime])

  return <div>{estTime} min</div>
}

export const Transaction = ({
  connectedAddress,
  originChainId,
  destinationChainId,
  originTxHash,
  bridgeModuleName,
  estimatedTime,
  kappa,
  timestamp,
  currentTime,
  isStoredComplete,
}: {
  connectedAddress: string
  originChainId: number
  destinationChainId: number
  originTxHash: string
  bridgeModuleName: string
  estimatedTime: number // in seconds
  kappa?: string
  timestamp: number
  currentTime: number
  isStoredComplete: boolean
}) => {
  const dispatch = useAppDispatch()
  const transactions = useTransactionsState()

  const { synapseSDK } = useSynapseContext()

  const [originTxExplorerLink, originExplorerName] = getTxBlockExplorerLink(
    originChainId,
    originTxHash
  )
  const [destExplorerAddressLink, destExplorerName] = getExplorerAddressUrl(
    destinationChainId,
    connectedAddress
  )
  const synapseExplorerLink = getTxSynapseExplorerLink({
    originChainId,
    destinationChainId,
    txHash: originTxHash,
    kappa,
  })

  const elapsedTime: number = currentTime - timestamp // in seconds
  const remainingTime: number = estimatedTime - elapsedTime

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
    checkStatus: !isStoredComplete || isEstimatedTimeReached,
    currentTime: currentTime,
  })

  /** Check if store already marked tx as complete, otherwise check hook status */
  const isTxFinalized = isStoredComplete ?? isTxComplete

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
    const txKappa = kappa ?? _kappa

    if (isTxComplete && originTxHash && txKappa) {
      if (transactions[originTxHash]?.isComplete === false) {
        dispatch(
          completeTransaction({ originTxHash, kappa: txKappa as string })
        )
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
      {isTxFinalized ? (
        <TransactionStatus string="Complete" />
      ) : (
        <TransactionStatus string="Pending" />
      )}
      <div className="flex items-center justify-end gap-2 grow">
        <TimeRemaining
          isComplete={isTxFinalized as boolean}
          remainingTime={remainingTime}
          isDelayed={isEstimatedTimeReached}
        />

        <DropdownMenu>
          {!isNull(originTxExplorerLink) && (
            <MenuItem text={originExplorerName} link={originTxExplorerLink} />
          )}
          {!isNull(destExplorerAddressLink) && (
            <MenuItem text={destExplorerName} link={destExplorerAddressLink} />
          )}
          {!isNull(synapseExplorerLink) && (
            <MenuItem text="Synapse Explorer" link={synapseExplorerLink} />
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
