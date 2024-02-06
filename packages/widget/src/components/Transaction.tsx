import { useEffect, useMemo, useCallback } from 'react'

import { useAppDispatch } from '@/state/hooks'
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
import { TimeRemaining } from '@/components/TimeRemaining'
import { DropdownMenu } from '@/components/ui/DropdownMenu'
import { MenuItem } from '@/components/ui/MenuItem'

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
    if (!currentTime || !estimatedTime || !timestamp) {
      return false
    }
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
    currentTime,
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
        flex flex-col
        gap-1 justify-end items-center pl-2.5 pr-1.5 py-1
        border border-solid border-[--synapse-border] rounded-md
      `}
      style={{ background: 'var(--synapse-surface' }}
    >
      <div className="flex flex-wrap-reverse justify-between w-full">
        {isTxFinalized ? 'Complete' : 'Pending'}
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
              <MenuItem
                text={destExplorerName}
                link={destExplorerAddressLink}
              />
            )}
            {/* {!isNull(synapseExplorerLink) && (
            <MenuItem text="Synapse Explorer" link={synapseExplorerLink} />
          )} */}
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
      <TransactionSupport />
    </div>
  )
}

const TransactionSupport = () => {
  return (
    <div id="transaction-support" className="flex items-center justify-between">
      <div>What's taking so long?</div>
      <div className="flex">
        <a
          href=""
          target="_blank"
          className="px-2 py-1 underline hover:rounded hover:bg-[--synapse-select-bg]"
        >
          Help page
        </a>
        <div>/</div>
        <a
          href="https://discord.gg/synapseprotocol"
          target="_blank"
          className="px-2 py-1 underline hover:rounded hover:bg-[--synapse-select-bg]"
        >
          Support (Discord)
        </a>
      </div>
    </div>
  )
}
