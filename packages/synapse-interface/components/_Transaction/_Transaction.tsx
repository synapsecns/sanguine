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
import { DropdownMenu } from './components/DropdownMenu'
import { MenuItem } from './components/MenuItem'
import { useBridgeTxUpdater } from './helpers/useBridgeTxUpdater'
import { ProgressBar } from './components/ProgressBar'

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
    isStartCheckingTimeReached,
    isEstimatedTimeReached,
    remainingTime,
    elapsedTime,
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

  useBridgeTxUpdater(
    connectedAddress,
    destinationChain,
    _kappa,
    originTxHash,
    isTxComplete
  )

  return (
    <div
      data-test-id="_transaction"
      className={`
        flex flex-col gap-1 justify-end items-center my-2
        bg-tint fill-surface text-primary
        border border-solid border-surface rounded-md
        text-xs md:text-base
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
        <div className="flex items-center space-x-4">
          <TransactionPayloadDetail
            chain={destinationChain}
            token={destinationToken}
            tokenAmount={null}
            isOrigin={false}
          />
          <div className="mt-1 text-xs">
            {new Date(timestamp * 1000).toLocaleString('en-US', {
              month: 'short',
              day: 'numeric',
              hour: '2-digit',
              minute: '2-digit',
              hour12: true,
            })}
            {/* <div>{typeof _kappa === 'string' && _kappa?.substring(0, 15)}</div> */}
          </div>
        </div>
        {/* TODO: Update visual format */}
        <div className="flex justify-between gap-2 pr-2 ml-auto">
          {isTxFinalized ? (
            <TransactionStatus string="Complete" className="text-green-300" />
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
      <ProgressBar
        elapsedTime={elapsedTime}
        totalTime={estimatedTime} // Double time for visual buffer
        isComplete={isTxFinalized}
      />
    </div>
  )
}
