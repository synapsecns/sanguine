import React, { useState, useEffect, useRef } from 'react'
import _ from 'lodash'
import { useAppDispatch, useAppSelector } from '@/store/hooks'
import { isValidAddress } from '@/utils/isValidAddress'
import { shortenAddress } from '@/utils/shortenAddress'
import { useBridgeState } from '@/slices/bridge/hooks'
import { setDestinationAddress } from '@/slices/bridge/reducer'
import { setShowDestinationWarning } from '@/slices/bridgeDisplaySlice'
import { Address } from 'viem'
import { isEmptyString } from '@/utils/isEmptyString'
import { CloseButton } from './components/CloseButton'
import { useTransactionsState } from '@/slices/transactions/hooks'
import { TransactionsState } from '@/slices/transactions/reducer'
import { BridgeTransaction } from '@/slices/api/generated'
import { getValidAddress } from '@/utils/isValidAddress'
import useCloseOnOutsideClick from '@/utils/hooks/useCloseOnOutsideClick'

export const inputRef = React.createRef<HTMLInputElement>()

export const DestinationAddressInput = ({
  connectedAddress,
}: {
  connectedAddress: string
}) => {
  const dispatch = useAppDispatch()
  const { destinationAddress } = useBridgeState()
  const { userHistoricalTransactions }: TransactionsState =
    useTransactionsState()
  const { showDestinationWarning } = useAppSelector(
    (state) => state.bridgeDisplay
  )
  const [isInputFocused, setIsInputFocused] = useState<boolean>(false)

  const recipientList = filterTxsByRecipient(
    userHistoricalTransactions,
    connectedAddress
  )

  const filteredRecipientList = filterNewestTxByRecipient(recipientList)

  /** TODO: Remove after testing */
  console.log('recipientList:', recipientList)
  console.log('filteredRecipientList:', filteredRecipientList)

  const handleInputFocus = () => {
    setIsInputFocused(true)
    setShowRecipientList(true)
  }

  const handleInputBlur = () => {
    setIsInputFocused(false)
  }

  const handleClearInput = () => {
    dispatch(setDestinationAddress('' as Address))
    inputRef.current.focus()
  }

  const isInputValidAddress = isValidAddress(destinationAddress)

  const isInputInvalid =
    destinationAddress &&
    !isEmptyString(destinationAddress) &&
    !isInputValidAddress

  let placeholder

  if (isInputFocused) {
    placeholder = ''
  } else {
    placeholder = connectedAddress
      ? shortenAddress(connectedAddress)
      : 'Wallet address'
  }

  const [isInputHovered, setIsInputHovered] = useState<boolean>(false)
  const [showRecipientList, setShowRecipientList] = useState<boolean>(false)
  const listRef = useRef(null)
  useCloseOnOutsideClick(listRef, () => setShowRecipientList(false))

  /** Warning State */
  const [showWarning, setShowWarning] = useState<boolean>(false)

  const handleActivateWarning = () => {
    if (!showWarning && showDestinationWarning) {
      setShowWarning(!showWarning)
    }
  }

  const handleAcceptWarning = () => {
    inputRef.current.focus()
    setShowWarning(false)
    dispatch(setShowDestinationWarning(false))
  }

  const handleRejectWarning = () => {
    setShowWarning(false)
  }

  return (
    <div id="destination-address-input" onClick={handleActivateWarning}>
      <HoverContent isHovered={!isInputFocused && isInputHovered}>
        {_.isEmpty(connectedAddress) ? 'Connect Wallet' : 'Destination address'}
      </HoverContent>
      <div
        onMouseEnter={() => setIsInputHovered(true)}
        onMouseLeave={() => setIsInputHovered(false)}
        className={`
           flex border text-md rounded-sm
           ${isInputFocused ? ' bg-bgBase' : 'bg-transparent hover:opacity-80'}
          ${
            isInputValidAddress
              ? 'border-synapsePurple focus:border-synapsePurple'
              : isInputInvalid
              ? 'border-red-500 focus:border-red-500'
              : 'border-separator focus:border-separator'
          }
        `}
      >
        <input
          ref={inputRef}
          onChange={(e) =>
            dispatch(setDestinationAddress(e.target.value as Address))
          }
          onFocus={handleInputFocus}
          onBlur={handleInputBlur}
          placeholder={placeholder}
          value={
            isInputValidAddress && !isInputFocused
              ? shortenAddress(destinationAddress)
              : destinationAddress
          }
          disabled={_.isEmpty(connectedAddress)}
          className={`
            text-md rounded-sm text-secondary py-1 px-2 z-0 border-0 bg-transparent
            focus:text-white focus:border-transparent focus:outline-none focus:ring-0
            ${connectedAddress ? 'w-32' : 'w-36'}
            ${
              isInputFocused || isInputInvalid
                ? 'text-left cursor-text'
                : 'text-center cursor-pointer'
            }
          `}
        />
        {(isInputInvalid || isInputValidAddress) && (
          <CloseButton
            onClick={handleClearInput}
            className="!static w-fit mr-1"
          />
        )}
      </div>
      <div className="relative">
        {showRecipientList && (
          <ul
            ref={listRef}
            className={`
              absolute right-0 z-50 p-0 top-1 bg-surface
              border border-solid border-tint rounded shadow
              popover list-none text-left text-sm
            `}
          >
            {filteredRecipientList?.map((recipient) => {
              return (
                <ListRecipient
                  address={recipient?.toAddress}
                  daysAgo={recipient?.daysAgo}
                  onSelectRecipient={(destinationAddress: Address) => {
                    dispatch(setDestinationAddress(destinationAddress))
                    setShowRecipientList(false)
                  }}
                />
              )
            })}
          </ul>
        )}
      </div>
      <DestinationInputWarning
        show={showWarning}
        onAccept={() => handleAcceptWarning()}
        onCancel={() => handleRejectWarning()}
      />
    </div>
  )
}

const ListRecipient = ({
  address,
  daysAgo,
  onSelectRecipient,
}: {
  address: string
  daysAgo: number
  onSelectRecipient?: (destinationAddress: Address) => void
}) => {
  return (
    <div
      onClick={() => onSelectRecipient(address as Address)}
      className={`
        flex justify-between p-1 space-x-2
        cursor-pointer text-strong
        hover:bg-separator
      `}
    >
      <div>{shortenAddress(address)}</div>
      <div>{daysAgo}d</div>
    </div>
  )
}

const DestinationInputWarning = ({
  show,
  onAccept,
  onCancel,
}: {
  show: boolean
  onAccept: () => void
  onCancel: () => void
}) => {
  return (
    <div
      className={`
        p-3 border rounded-sm bg-surface border-separator text-secondary
        top-0 left-0 w-full space-y-2 z-50
        ${show ? 'absolute' : 'hidden'}
      `}
    >
      <h3 className="text-2xl font-semibold tracking-wide text-white">
        Warning
      </h3>
      <p className="text-white">
        Do not send your funds to a{' '}
        <div className="inline text-yellowText">custodial wallet</div> or{' '}
        <div className="inline text-yellowText">exchange</div> address!
      </p>
      <p className="text-secondary">
        It may be impossible to recover your funds.
      </p>
      <div className="flex space-x-2">
        <button onClick={onCancel} className="w-1/2 py-3 bg-separator">
          Cancel
        </button>
        <button
          onClick={onAccept}
          className="w-1/2 py-3 bg-transparent border border-separator"
        >
          Okay
        </button>
      </div>
    </div>
  )
}

const filterTxsByRecipient = (
  transactions: BridgeTransaction[],
  connectedAddress?: string
): {
  toAddress: string | undefined
  time: string | undefined
  daysAgo: number
}[] => {
  const checkAddress = getValidAddress(connectedAddress)
  return transactions
    ?.filter(
      (transaction) =>
        getValidAddress(transaction.toInfo?.address) !== checkAddress
    )
    .map((transaction) => ({
      toAddress: getValidAddress(transaction.toInfo?.address),
      time: transaction.toInfo?.formattedTime,
      daysAgo: calculateDaysAgo(transaction.toInfo?.formattedTime),
    }))
}

const filterNewestTxByRecipient = (
  transactions: {
    toAddress: string | undefined
    time: string | undefined
    daysAgo: number
  }[]
) => {
  const newestTxMap = new Map()

  transactions?.forEach((tx) => {
    const existingTx = newestTxMap.get(tx.toAddress)

    if (!existingTx || tx.daysAgo < existingTx.daysAgo) {
      newestTxMap.set(tx.toAddress, tx)
    }
  })

  return Array.from(newestTxMap.values())
}

const calculateDaysAgo = (time?: string) => {
  if (time) {
    // Assuming timeString is in "YYYY-MM-DD HH:MM:SS +0000 UTC" format
    const formattedTimeString = time.replace(' +0000 UTC', 'Z')

    const timeDate = new Date(formattedTimeString)
    const now = new Date()

    return calculateDaysBetween(timeDate, now)
  } else {
    return null
  }
}

const calculateDaysBetween = (startDate: Date, endDate: Date) => {
  const msPerDay = 1000 * 60 * 60 * 24
  const utc1 = Date.UTC(
    startDate.getFullYear(),
    startDate.getMonth(),
    startDate.getDate()
  )
  const utc2 = Date.UTC(
    endDate.getFullYear(),
    endDate.getMonth(),
    endDate.getDate()
  )

  return Math.floor((utc2 - utc1) / msPerDay)
}

export const HoverContent = ({
  isHovered,
  children,
}: {
  isHovered: boolean
  children: React.ReactNode
}) => {
  if (isHovered) {
    return (
      <div
        className={`
          absolute top-[-0.5rem] z-50 hover-content py-1 px-2 text-secondary
          border border-solid border-separator text-xs
          bg-[#101018] rounded-sm text-center whitespace-nowrap
        `}
      >
        {children}
      </div>
    )
  }
}
