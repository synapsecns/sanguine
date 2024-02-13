import { useCallback } from 'react'
import { isNull } from 'lodash'
import { useAppDispatch } from '@/store/hooks'
import { getTxBlockExplorerLink } from './helpers/getTxBlockExplorerLink'
import { getExplorerAddressLink } from './helpers/getExplorerAddressLink'
import { useBridgeTxStatus } from './helpers/useBridgeTxStatus'
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
import { AnimatedProgressBar } from './components/AnimatedProgressBar'
import { TransactionSupport } from './components/TransactionSupport'
import { RightArrow } from '@/components/icons/RightArrow'

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

  return (
    <div
      data-test-id="_transaction"
      className={`
        flex flex-col gap-1 justify-end items-center my-2
        bg-bgBase/10 fill-bgBase/20 text-primary
        border border-solid border-white/20 rounded-md
        text-xs md:text-base
      `}
    >
      <div className="flex items-center w-full">
        <div className="flex rounded  fill-bgBase/20">
          <div className="bg-bgBase/20">
            <TransactionPayloadDetail
              chain={originChain}
              token={originToken}
              tokenAmount={originValue}
              isOrigin={true}
            />
          </div>

          <TransactionArrow className="stroke-white/20 " />
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
          </div>
        </div>
        {/* TODO: Update visual format */}
        <div className="flex justify-between gap-2 pr-2 ml-auto">
          {isTxFinalized ? (
            <TransactionStatus
              string="Complete"
              className="text-green-300 text-sm"
            />
          ) : (
            <TransactionStatus string="Pending" className="text-sm" />
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
    </div>
  )
}
