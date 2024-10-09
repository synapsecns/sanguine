import { debounce } from 'lodash'
import React, { useCallback, useState } from 'react'
import { NumericFormat } from 'react-number-format'

import { joinClassNames } from '@/utils/joinClassNames'

interface CustomAmountInputTypes {
  inputRef?: React.RefObject<HTMLInputElement>
  disabled?: boolean
  showValue: string
  handleFromValueChange?: (value: string) => void
  setIsTyping?: (isTyping: boolean) => void
  className?: string
}

export function CustomAmountInput({
  inputRef,
  disabled = false,
  showValue,
  handleFromValueChange,
  setIsTyping,
  className,
}: CustomAmountInputTypes) {
  const [localValue, setLocalValue] = useState(showValue)

  const debouncedSetIsTyping = useCallback(
    debounce((value: boolean) => setIsTyping?.(value), 600),
    [setIsTyping]
  )

  const debouncedUpdateValue = useCallback(
    debounce((value: string) => handleFromValueChange?.(value), 600),
    [handleFromValueChange]
  )

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const newValue = event.target.value
    setLocalValue(newValue)
    setIsTyping?.(true)
    debouncedSetIsTyping(false)
    debouncedUpdateValue(newValue)
  }

  const inputClassNames = {
    unset: 'bg-transparent border-none p-0',
    layout: 'w-full',
    placeholder: 'placeholder:text-zinc-500 placeholder:dark:text-zinc-400',
    font: 'text-xl md:text-2xl font-medium',
    focus: 'focus:outline-none focus:ring-0 focus:border-none',
    custom: className,
  }

  return (
    <NumericFormat
      inputMode="numeric"
      getInputRef={inputRef}
      placeholder="0.0000"
      value={localValue}
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
  )
}
