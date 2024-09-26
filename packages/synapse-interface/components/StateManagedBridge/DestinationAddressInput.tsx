import React, { useState, useRef, useEffect } from 'react'
import { isNull, isString } from 'lodash'
import { useTranslations } from 'next-intl'

import { useAppDispatch } from '@/store/hooks'
import { isValidAddress } from '@/utils/isValidAddress'
import { shortenAddress } from '@/utils/shortenAddress'
import { useBridgeState, useBridgeDisplayState } from '@/slices/bridge/hooks'
import {
  setDestinationAddress,
  clearDestinationAddress,
} from '@/slices/bridge/reducer'
import {
  setShowDestinationWarning,
  setIsDestinationWarningAccepted,
} from '@/slices/bridgeDisplaySlice'
import { Address } from 'viem'
import { isEmptyString } from '@/utils/isEmptyString'
import { CloseButton } from '@/components/ui/CloseButton'
import { useTransactionsState } from '@/slices/transactions/hooks'
import { TransactionsState } from '@/slices/transactions/reducer'
import { BridgeTransaction } from '@/slices/api/generated'
import { getValidAddress } from '@/utils/isValidAddress'
import { useKeyPress } from '@/utils/hooks/useKeyPress'
import useCloseOnOutsideClick from '@/utils/hooks/useCloseOnOutsideClick'

export const DestinationAddressInput = ({
  connectedAddress,
}: {
  connectedAddress: string
}) => {
  const dispatch = useAppDispatch()
  const t = useTranslations('Bridge')
  const { destinationAddress } = useBridgeState()
  const { showDestinationWarning } = useBridgeDisplayState()
  const { userHistoricalTransactions }: TransactionsState =
    useTransactionsState()

  const recipientList = filterTxsByRecipient(
    userHistoricalTransactions,
    connectedAddress
  )
  const filteredRecipientList = filterNewestTxByRecipient(recipientList)

  const isInputValidAddress: boolean = destinationAddress
    ? isValidAddress(destinationAddress)
    : false

  const isInputInvalid: boolean =
    (destinationAddress &&
      isString(destinationAddress) &&
      isEmptyString(destinationAddress)) ||
    (destinationAddress && !isInputValidAddress)

  const isSameAddress =
    connectedAddress &&
    isInputValidAddress &&
    getValidAddress(destinationAddress) === getValidAddress(connectedAddress)

  useEffect(() => {
    const showWarning = isInputValidAddress && !isSameAddress

    if (showWarning && !showDestinationWarning) {
      dispatch(setShowDestinationWarning(true))
    }

    if (!isInputValidAddress && showDestinationWarning) {
      dispatch(setShowDestinationWarning(false))
      dispatch(setIsDestinationWarningAccepted(false))
    }
  }, [
    dispatch,
    connectedAddress,
    destinationAddress,
    showDestinationWarning,
    isInputValidAddress,
  ])

  const inputRef = useRef<HTMLInputElement>(null)
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

    if (destinationAddress && isInputInvalid) {
      handleClearInput()
    }
  }

  const handleClearInput = () => {
    dispatch(clearDestinationAddress())

    if (inputRef.current) {
      inputRef.current.value = ''
    }
  }

  const onClearUserInput = () => {
    handleClearInput()
    handleInputBlur()
  }

  useEffect(() => {
    dispatch(clearDestinationAddress())
    handleClearInput()
  }, [connectedAddress])

  useEffect(() => {
    if (!isInputFocused && isSameAddress) {
      handleClearInput()
    }
  }, [isSameAddress, destinationAddress, connectedAddress, isInputFocused])

  let placeholder

  if (isInputFocused) {
    placeholder = ''
  } else {
    placeholder = connectedAddress ? shortenAddress(connectedAddress) : '0x...'
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
  const [currentIdx, setCurrentIdx] = useState<number>(null)

  const handleCloseList = () => {
    if (showRecipientList) {
      setShowRecipientList(false)
    }
    setCurrentIdx(null)
  }

  const handlePaste = async () => {
    const pastedValue = await navigator.clipboard.readText()
    dispatch(setDestinationAddress(pastedValue as Address))
  }

  const onSelectRecipient = (address) => {
    dispatch(setDestinationAddress(address))
    handleInputBlur()
    handleCloseList()
  }

  useCloseOnOutsideClick(listRef, handleCloseList)

  const escPressed = useKeyPress('Escape', true)
  const arrowUp = useKeyPress('ArrowUp', true)
  const arrowDown = useKeyPress('ArrowDown', true)
  const enterPressed = useKeyPress('Enter', true)

  function escFunc() {
    if (!showRecipientList) return

    if (escPressed) {
      handleCloseList()
      handleClearInput()
      handleInputBlur()
    }
  }

  function arrowDownFunc() {
    if (filteredRecipientList.length === 0) return
    if (!showRecipientList) return

    const nextIdx = currentIdx + 1
    if (currentIdx === null) {
      setCurrentIdx(0)
    } else if (arrowDown && nextIdx < filteredRecipientList.length) {
      setCurrentIdx(nextIdx)
    }
  }

  function arrowUpFunc() {
    if (filteredRecipientList.length === 0) return
    if (!showRecipientList) return

    const nextIdx = currentIdx - 1

    if (arrowUp && -1 < nextIdx) {
      setCurrentIdx(nextIdx)
    }
  }

  function enterPressedFunc() {
    if (enterPressed && isNull(currentIdx)) {
      onSelectRecipient(destinationAddress)
    }

    if (enterPressed && !isNull(currentIdx) && currentIdx > -1) {
      onSelectRecipient(filteredRecipientList[currentIdx]?.toAddress)
    }
  }

  useEffect(enterPressedFunc, [enterPressed])
  useEffect(escFunc, [escPressed])
  useEffect(arrowDownFunc, [arrowDown])
  useEffect(arrowUpFunc, [arrowUp])

  const adjustInputSize = () => {
    const addressInput: HTMLElement = document.getElementById('address-input')

    if (!addressInput) return

    if (isInputFocused || isInputInvalid) {
      addressInput.style.width = '12rem'
    } else if (inputValue.length > 0) {
      addressInput.style.width = inputValue.length + 2 + 'ch'
    } else {
      addressInput.style.width = placeholder.length + 'ch'
    }
  }

  useEffect(() => {
    adjustInputSize()
  }, [inputValue, placeholder, isInputFocused, showRecipientList])

  return (
    <div id="destination-address-input">
      <div className="relative flex items-center">
        <div className="mr-1.5 text-secondary text-sm">{t('To')}: </div>
        <div
          className={`
           flex border text-md rounded-sm
           ${isInputFocused ? ' bg-bgBase' : 'bg-transparent hover:opacity-80'}
          ${
            isInputInvalid
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
              transform-gpu transition-all duration-175 cursor-pointer max-w-36 md:max-w-48
              text-sm rounded-sm text-strong py-0.5 pl-1.5 z-0 border-0 bg-transparent
              focus:text-white focus:border-transparent focus:outline-none focus:ring-0
              ${isInputFocused || isInputInvalid ? 'text-left ' : 'text-center'}
              ${destinationAddress ? 'pr-6' : 'pr-1.5'}
            `}
          />
          {destinationAddress ? (
            <CloseButton
              onClick={onClearUserInput}
              className="!right-0 !w-5 !h-5 mr-0.5 mt-0.5 hover:opacity-70"
            />
          ) : isInputFocused ? (
            <PasteButton onPaste={handlePaste} />
          ) : null}
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
            {filteredRecipientList?.map((recipient, index) => {
              return (
                <ListRecipient
                  key={recipient?.toAddress}
                  index={index}
                  address={recipient?.toAddress}
                  daysAgo={recipient?.daysAgo}
                  selectedListItem={currentIdx}
                  onSelectRecipient={onSelectRecipient}
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
  index,
  address,
  daysAgo,
  selectedListItem,
  onSelectRecipient,
}: {
  index: number
  address: string
  daysAgo: number
  selectedListItem: number | null
  onSelectRecipient?: (destinationAddress: Address) => void
}) => {
  const handleMouseDown = (event: React.MouseEvent<HTMLDivElement>) => {
    event.preventDefault()
    onSelectRecipient && onSelectRecipient(address as Address)
  }

  const t = useTranslations('Time')

  return (
    <div
      onMouseDown={handleMouseDown}
      className={`
        flex justify-between px-2 py-1 space-x-3.5
        cursor-pointer text-strong text-sm
        hover:bg-separator
        ${selectedListItem === index && 'bg-separator'}
      `}
    >
      <div>{shortenAddress(address)}</div>
      <div>
        {daysAgo}
        {t('d')}
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

const PasteButton = ({ onPaste }: { onPaste: () => Promise<void> }) => {
  return (
    <svg
      width="19"
      height="19"
      viewBox="0 0 24 24"
      xmlns="http://www.w3.org/2000/svg"
      onClick={onPaste}
      onMouseDown={(e) => e.preventDefault()}
      className={`
        absolute border-transparent cursor-pointer
        right-0.5 mt-0.5 justify-self-end
        fill-zinc-100 stroke-zinc-100 hover:opacity-70
      `}
    >
      <rect x="5.5" y="5.5" width="13" height="16" rx="2" fill="none" />
      <path d="M9 7.5C8.72386 7.5 8.5 7.27614 8.5 7C8.5 5.067 10.067 3.5 12 3.5C13.933 3.5 15.5 5.067 15.5 7C15.5 7.27614 15.2761 7.5 15 7.5H9Z" />
    </svg>
  )
}
