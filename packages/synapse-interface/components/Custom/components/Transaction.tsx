import { useEffect, useMemo, useCallback } from 'react'
import { type Chain } from '@/utils/types'

import { useAppDispatch } from '@/store/hooks'
import { getTxBlockExplorerLink } from '../utils/getTxBlockExplorerLink'
import { getExplorerAddressUrl } from '../utils/getExplorerAddressUrl'
import { useBridgeTxStatus } from '../hooks/useBridgeTxStatus'
import { isNull } from '../utils/isNull'
import {
  updateTransactionKappa,
  completeTransaction,
  removeTransaction,
} from '@/slices/_transactions/reducer'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import { TimeRemaining } from './TimeRemaining'
import { DropdownMenu } from './DropdownMenu'
import { MenuItem } from './MenuItem'
import { CHAINS_BY_ID } from '@/constants/chains'
import { AnimatedProgressBar } from './AnimatedProgressBar'
import { use_TransactionsState } from '@/slices/_transactions/hooks'

export const Transaction = ({
  connectedAddress,
  originAmount,
  originTokenSymbol,
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
  originAmount: string
  originTokenSymbol: string
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
  const { transactions } = use_TransactionsState()

  const { synapseSDK } = useSynapseContext()

  const [originTxExplorerLink, originExplorerName] = getTxBlockExplorerLink(
    originChainId,
    originTxHash
  )
  const [destExplorerAddressLink, destExplorerName] = getExplorerAddressUrl(
    destinationChainId,
    connectedAddress
  )

  const elapsedTime: number = currentTime - timestamp // in seconds
  const remainingTime: number = estimatedTime - elapsedTime

  const isEstimatedTimeReached: boolean = useMemo(() => {
    if (!currentTime || !estimatedTime || !timestamp) {
      return false
    }
    return currentTime - timestamp > estimatedTime
  }, [estimatedTime, currentTime, timestamp])

  const delayedTime = isEstimatedTimeReached ? remainingTime : null
  const delayedTimeInMin = remainingTime ? Math.floor(remainingTime / 60) : null

  const [isTxComplete, _kappa] = useBridgeTxStatus({
    synapseSDK,
    originChainId,
    destinationChainId,
    originTxHash,
    bridgeModuleName,
    kappa,
    checkStatus: !isStoredComplete || isEstimatedTimeReached,
    currentTime,
  })

  /** Check if store already marked tx as complete, otherwise check hook status */
  const isTxFinalized = isStoredComplete ?? isTxComplete

  const showTransactionSupport =
    !isTxFinalized && delayedTimeInMin ? delayedTimeInMin <= -5 : false

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

    if (!isStoredComplete && isTxComplete && originTxHash && txKappa) {
      dispatch(completeTransaction({ originTxHash, kappa: txKappa as string }))
    }
  }, [isStoredComplete, isTxComplete, dispatch, transactions])

  const handleClearTransaction = useCallback(() => {
    dispatch(removeTransaction({ originTxHash }))
  }, [dispatch])

  return (
    <div
      data-test-id="transaction"
      className={`
        flex flex-col relative
        rounded-md
        gap-1 justify-end items-center pl-2.5 pr-1.5 py-1
        ${isStoredComplete ? 'bg-[#1B3925]' : 'bg-zinc-100 dark:bg-bgBase'}
      `}
    >
      <div className="flex flex-wrap-reverse items-center justify-between w-full">
        <TransactionBridgeDetail
          tokenAmount={originAmount}
          originTokenSymbol={originTokenSymbol}
          destinationChain={CHAINS_BY_ID[destinationChainId]}
        />
        <div className="flex items-center justify-end gap-2 grow">
          <DropdownMenu
            menuTitleElement={
              <TimeRemaining
                isComplete={isTxFinalized as boolean}
                remainingTime={remainingTime}
                isDelayed={isEstimatedTimeReached}
                delayedTime={delayedTime}
              />
            }
          >
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
      <div className="absolute bottom-0 w-full px-1 text-[0px]">
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

const TransactionBridgeDetail = ({
  tokenAmount,
  originTokenSymbol,
  destinationChain,
}: {
  tokenAmount: string
  originTokenSymbol: string
  destinationChain: Chain
}) => {
  const showAmount = parseFloat(tokenAmount)?.toFixed(6)

  return (
    <div className="flex">
      {showAmount} {originTokenSymbol} to {destinationChain?.name}
    </div>
  )
}

const TRANSACTION_SUPPORT_URL =
  'https://docs.synapseprotocol.com/synapse-bridge/synapse-bridge/transaction-support-faq'

export const TransactionSupport = () => {
  return (
    <div
      id="transaction-support"
      className="flex items-center justify-between w-full"
    >
      <div>What's taking so long?</div>
      <div className="flex items-center">
        <a
          href={TRANSACTION_SUPPORT_URL}
          target="_blank"
          className={`
            px-2 py-1 underline text-[--synapse-text]
            hover:rounded hover:bg-[--synapse-select-bg] hover:no-underline
          `}
        >
          F.A.Q
        </a>
        <div>/</div>
        <a
          href="https://discord.gg/synapseprotocol"
          target="_blank"
          className={`
            px-2 py-1 underline text-[--synapse-text]
            hover:rounded hover:bg-[--synapse-select-bg] hover:no-underline
            min-[360px]:after:content-['_(Discord)']
          `}
        >
          Support
        </a>
      </div>
    </div>
  )
}
