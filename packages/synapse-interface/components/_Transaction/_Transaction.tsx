import { useMemo, useEffect, useState, useCallback } from 'react'
import { useAppDispatch } from '@/store/hooks'
import { fetchAndStoreSingleNetworkPortfolioBalances } from '@/slices/portfolio/hooks'
import { use_TransactionsState } from '@/slices/_transactions/hooks'
import { getTxBlockExplorerLink } from './helpers/getTxBlockExplorerLink'
import { getExplorerAddressLink } from './helpers/getExplorerAddressLink'
import { useBridgeTxStatus } from './helpers/useBridgeTxStatus'
import { isNull } from 'lodash'
import {
  updateTransactionKappa,
  completeTransaction,
  removeTransaction,
} from '@/slices/_transactions/reducer'
import { TransactionPayloadDetail } from '../Portfolio/Transaction/components/TransactionPayloadDetail'
import { Chain, Token } from '@/utils/types'
import TransactionArrow from '../icons/TransactionArrow'
import { TimeRemaining } from './components/TimeRemaining'
import { TransactionStatus } from './components/TransactionStatus'
import { getEstimatedTimeStatus } from './helpers/getEstimatedTimeStatus'
import { DropdownMenu } from './components/DropdownMenu'
import { MenuItem } from './components/DropdownMenu'

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
  const { transactions } = use_TransactionsState()

  const [originTxExplorerLink, originExplorerName] = getTxBlockExplorerLink(
    originChain.id,
    originTxHash
  )
  const [destExplorerAddressLink, destExplorerName] = getExplorerAddressLink(
    destinationChain.id,
    connectedAddress
  )

  const {
    elapsedTime: estimatedElapsedTime,
    startCheckingTimeReached,
    remainingTimeInMinutes: EstimatedRemainingTimeInMinutes,
  } = getEstimatedTimeStatus(currentTime, timestamp, estimatedTime)

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

  console.log('estimatedElapsedTime: ', estimatedElapsedTime)
  console.log('startCheckingTimeReached: ', startCheckingTimeReached)
  console.log(
    'EstimatedRemainingTimeInMinutes:',
    EstimatedRemainingTimeInMinutes
  )

  console.log('isEstimatedTimeReached:', isEstimatedTimeReached)
  console.log('elapsedTime:', elapsedTime)
  console.log('remainingTimeInMinutes: ', remainingTimeInMinutes)

  const [isTxComplete, _kappa] = useBridgeTxStatus({
    originChainId: originChain.id,
    destinationChainId: destinationChain.id,
    originTxHash,
    bridgeModuleName,
    kappa: kappa,
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
    const txKappa = _kappa

    if (isTxComplete && originTxHash && txKappa) {
      const txn = transactions.find((tx) => tx.originTxHash === originTxHash)
      if (!txn.isComplete) {
        dispatch(
          completeTransaction({ originTxHash, kappa: txKappa as string })
        )
        /** Update Destination Chain token balances after tx is marked complete  */
        dispatch(
          fetchAndStoreSingleNetworkPortfolioBalances({
            address: connectedAddress,
            chainId: destinationChain.id,
          })
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
        <div className="flex items-center">
          <TransactionPayloadDetail
            chain={destinationChain}
            token={destinationToken}
            tokenAmount={null}
            isOrigin={false}
          />
          <div className="mt-1 text-xs">
            {new Date(timestamp * 1000).toLocaleString('en-US', {
              hour: '2-digit',
              minute: '2-digit',
              second: '2-digit',
              hour12: true,
            })}
            {/* <div>{typeof _kappa === 'string' && _kappa?.substring(0, 15)}</div> */}
          </div>
        </div>
        {/* TODO: Update visual format */}
        <div className="flex justify-between gap-2 pr-2 ml-auto">
          {isTxFinalized ? (
            <TransactionStatus string="Complete" />
          ) : (
            <TransactionStatus string="Pending" />
          )}
          <div className="flex items-center justify-end gap-2 grow">
            <TimeRemaining
              isComplete={isTxFinalized as boolean}
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
      </div>
    </div>
  )
}
