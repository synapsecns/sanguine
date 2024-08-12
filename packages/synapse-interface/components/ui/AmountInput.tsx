import React, { useState, useEffect, useCallback } from 'react'
import { debounce } from 'lodash'
import LoadingDots from './tailwind/LoadingDots'
import { joinClassNames } from '@/utils/joinClassNames'

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
  const [localValue, setLocalValue] = useState(showValue)

  const debouncedSetIsTyping = useCallback(
    debounce((value: boolean) => setIsTyping(value), 500),
    [setIsTyping]
  )

  useEffect(() => {
    setLocalValue(showValue)
  }, [showValue])

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const newValue = event.target.value
    setLocalValue(newValue)
    setIsTyping(true)
    debouncedSetIsTyping(false)
    handleFromValueChange?.(event)
  }

  const inputClassName = joinClassNames({
    unset: 'bg-transparent border-none p-0',
    layout: 'w-full',
    placeholder: 'placeholder:text-zinc-500 placeholder:dark:text-zinc-400',
    font: 'text-xl md:text-2xl font-medium',
    focus: 'focus:outline-none focus:ring-0 focus:border-none',
  })

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
          className={inputClassName}
          placeholder="0.0000"
          onChange={handleInputChange}
          value={localValue}
          name="inputRow"
          autoComplete="off"
          minLength={1}
          maxLength={79}
        />
      )}
    </>
  )
}
