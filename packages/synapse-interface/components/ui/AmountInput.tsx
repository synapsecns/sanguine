import { debounce } from 'lodash'
import React, { useCallback } from 'react'
import { NumericFormat } from 'react-number-format'
import { joinClassNames } from '@/utils/joinClassNames'
import LoadingDots from './tailwind/LoadingDots'

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
        <NumericFormat
          getInputRef={inputRef}
          placeholder="0.0000"
          value={showValue}
          pattern={disabled ? '[0-9.]+' : '^[0-9]+([.,]?[0-9]+)?$'}
          disabled={disabled}
          readOnly={disabled}
          onChange={handleInputChange}
          className={joinClassNames(inputClassNames)}
          name="inputRow"
          minLength={1}
          maxLength={79}
          autoComplete="off"
          thousandSeparator={true}
          allowNegative={false}
        />
      )}
    </>
  )
}
