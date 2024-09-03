import React, { useState, useEffect, useCallback } from 'react'
import { debounce } from 'lodash'
import LoadingDots from './tailwind/LoadingDots'
import { joinClassNames } from '@/utils/joinClassNames'
import { formatNumberWithCommas } from '@/utils/formatNumberWithCommas'

interface AmountInputTypes {
  inputRef?: React.RefObject<HTMLInputElement>
  disabled?: boolean
  isLoading?: boolean
  showValue: string
  handleFromValueChange?: (event: React.ChangeEvent<HTMLInputElement>) => void
  setIsTyping?: (isTyping: boolean) => void
}

export function AmountInput({
  inputRef,
  disabled = false,
  isLoading = false,
  showValue,
  handleFromValueChange,
  setIsTyping,
}: AmountInputTypes) {
  const debouncedSetIsTyping = useCallback(
    debounce((value: boolean) => setIsTyping?.(value), 600),
    [setIsTyping]
  )

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setIsTyping?.(true)
    debouncedSetIsTyping(false)
    handleFromValueChange?.(event)
  }

  const inputClassNames = {
    unset: 'bg-transparent border-none p-0',
    layout: 'w-full',
    placeholder: 'placeholder:text-zinc-500 placeholder:dark:text-zinc-400',
    font: 'text-xl md:text-2xl font-medium',
    focus: 'focus:outline-none focus:ring-0 focus:border-none',
  }

  return (
    <>
      {isLoading ? (
        <LoadingDots className="opacity-50" />
      ) : (
        <input
          ref={inputRef}
          pattern={disabled ? '[0-9.]+' : '^[0-9]*[.,]?[0-9]*$'}
          disabled={disabled}
          readOnly={disabled}
          className={joinClassNames(inputClassNames)}
          placeholder="0.0000"
          onChange={handleInputChange}
          value={formatNumberWithCommas(showValue)}
          name="inputRow"
          autoComplete="off"
          minLength={1}
          maxLength={79}
        />
      )}
    </>
  )
}
