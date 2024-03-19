import React, { useState, useRef, useEffect } from 'react'
import { isEmpty, isString } from 'lodash'
import { useAppDispatch } from '@/store/hooks'
import { isValidAddress } from '@/utils/isValidAddress'
import { shortenAddress } from '@/utils/shortenAddress'
import { useBridgeState, useBridgeDisplayState } from '@/slices/bridge/hooks'
import {
  setDestinationAddress,
  clearDestinationAddress,
} from '@/slices/bridge/reducer'
import { setShowDestinationWarning } from '@/slices/bridgeDisplaySlice'
import { Address } from 'viem'
import { isEmptyString } from '@/utils/isEmptyString'
import { CloseButton } from './components/CloseButton'
import { useTransactionsState } from '@/slices/transactions/hooks'
import { TransactionsState } from '@/slices/transactions/reducer'
import { BridgeTransaction } from '@/slices/api/generated'
import { getValidAddress } from '@/utils/isValidAddress'
import useCloseOnOutsideClick from '@/utils/hooks/useCloseOnOutsideClick'
import { useKeyPress } from '@/utils/hooks/useKeyPress'

const inputRef = React.createRef<HTMLInputElement>()

export const DestinationAddressInput = ({
  connectedAddress,
}: {
  connectedAddress: string
}) => {
  const dispatch = useAppDispatch()
  const { destinationAddress } = useBridgeState()
  const { showDestinationWarning } = useBridgeDisplayState()
  const { userHistoricalTransactions }: TransactionsState =
    useTransactionsState()

  const recipientList = filterTxsByRecipient(
    userHistoricalTransactions,
    connectedAddress
  )
  const filteredRecipientList = filterNewestTxByRecipient(recipientList)

  const [isInputFocused, setIsInputFocused] = useState<boolean>(false)

  const handleInputFocus = () => {
    setIsInputFocused(true)
    setShowRecipientList(true)
    if (inputRef.current) {
      inputRef.current.focus()
    }
  }

  const handleInputBlur = () => {
    setIsInputFocused(false)

    if (inputRef.current) {
      inputRef.current.blur()
    }
  }

  const handleClearInput = () => {
    if (inputRef.current) {
      inputRef.current.value = ''
    }
  }

  const onClearUserInput = () => {
    dispatch(clearDestinationAddress())
    handleClearInput()
    handleInputFocus()
  }

  const isInputValidAddress: boolean = destinationAddress
    ? isValidAddress(destinationAddress)
    : false

  const isInputInvalid: boolean =
    (destinationAddress &&
      isString(destinationAddress) &&
      isEmptyString(destinationAddress)) ||
    (destinationAddress && !isInputValidAddress)

  useEffect(() => {
    const isSameAddress =
      connectedAddress &&
      isInputValidAddress &&
      getValidAddress(destinationAddress) === getValidAddress(connectedAddress)

    const showWarning = isInputValidAddress && !isSameAddress

    if (showWarning && !showDestinationWarning) {
      dispatch(setShowDestinationWarning(true))
    }

    if (!isInputValidAddress && showDestinationWarning) {
      dispatch(setShowDestinationWarning(false))
    }
  }, [
    destinationAddress,
    connectedAddress,
    showDestinationWarning,
    dispatch,
    isInputValidAddress,
  ])

  let placeholder

  if (isInputFocused) {
    placeholder = ''
  } else {
    placeholder = connectedAddress
      ? // ? shortenAddress(connectedAddress)
        '0x...'
      : 'Wallet address'
  }

  let inputValue

  if (!destinationAddress) {
    inputValue = ''
  } else {
    inputValue =
      isInputValidAddress && !isInputFocused
        ? shortenAddress(destinationAddress)
        : destinationAddress
  }

  const listRef = useRef(null)
  const [showRecipientList, setShowRecipientList] = useState<boolean>(false)

  const handleCloseList = () => {
    if (showRecipientList) {
      setShowRecipientList(false)
    }
  }

  useCloseOnOutsideClick(listRef, handleCloseList)

  const escPressed = useKeyPress('Escape')

  function escFunc() {
    if (escPressed) {
      handleCloseList()
      handleClearInput()
      handleInputBlur()
    }
  }

  useEffect(escFunc, [escPressed])

  useEffect(() => {
    dispatch(clearDestinationAddress())
    handleClearInput()
  }, [connectedAddress])

  const adjustInputSize = () => {
    const addressInput: HTMLElement = document.getElementById('address-input')

    if (isInputFocused || isInputInvalid || showRecipientList) {
      addressInput.style.width = '8rem'
    } else if (inputValue.length > 0) {
      addressInput.style.width = inputValue.length + 2 + 'ch'
    } else {
      addressInput.style.width = placeholder.length + 1 + 'ch'
    }
  }

  useEffect(() => {
    adjustInputSize()
  }, [inputValue, placeholder, isInputFocused, showRecipientList])

  return (
    <div id="destination-address-input">
      <div className="flex items-center">
        <div className="mr-1.5 text-secondary">To: </div>
        <div
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
            id="address-input"
            ref={inputRef}
            onChange={(e) =>
              dispatch(setDestinationAddress(e.target.value as Address))
            }
            onFocus={handleInputFocus}
            onBlur={handleInputBlur}
            placeholder={placeholder}
            value={inputValue}
            className={`
              transform-gpu transition-all duration-75 cursor-pointer
              text-md rounded-sm text-strong py-0.5 pl-2 z-0 border-0 bg-transparent max-w-32
              focus:text-white focus:border-transparent focus:outline-none focus:ring-0
              ${isInputFocused || isInputInvalid ? 'text-left ' : 'text-center'}
              ${destinationAddress ? 'pr-6' : 'pr-1'}
            `}
          />
          {destinationAddress && (
            <div>
              <CloseButton
                onClick={onClearUserInput}
                className="!w-5 !h-5 mr-1 mt-1"
              />
            </div>
          )}
        </div>
      </div>
      <div className="relative">
        {showRecipientList && (
          <ul
            ref={listRef}
            className={`
              absolute right-0 z-50 p-0 top-1 bg-surface
              border border-solid border-tint rounded shadow
              popover list-none text-left overflow-hidden
            `}
          >
            {filteredRecipientList?.map((recipient) => {
              return (
                <ListRecipient
                  key={recipient?.toAddress}
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
        flex justify-between px-1.5 py-1 space-x-2
        cursor-pointer text-strong
        hover:bg-separator
      `}
    >
      <div>{shortenAddress(address)}</div>
      <div>{daysAgo}d</div>
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
