import { debounce } from 'lodash'
import React, { useCallback } from 'react'
import { NumericFormat } from 'react-number-format'
import { joinClassNames } from '@/utils/joinClassNames'
import LoadingDots from './tailwind/LoadingDots'
import { HoverTooltip } from '@/components/HoverTooltip'

interface AmountInputTypes {
  inputRef?: React.RefObject<HTMLInputElement>
  disabled?: boolean
  isLoading?: boolean
  showValue: string
  handleFromValueChange?: (event: React.ChangeEvent<HTMLInputElement>) => void
  setIsTyping?: (isTyping: boolean) => void
  className?: string
  tooltipValue?: string
}

export function AmountInput({
  inputRef,
  disabled = false,
  isLoading = false,
  showValue,
  handleFromValueChange,
  setIsTyping,
  className,
  tooltipValue,
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
    custom: className,
    loading: isLoading ? 'invisible' : '',
  }

  return (
    <div className="flex items-center">
      {isLoading && <LoadingDots className="opacity-50" />}
      <HoverTooltip
        hoverContent={tooltipValue}
        isActive={!!(tooltipValue && !isLoading)}
        align="start"
      >
        <NumericFormat
          inputMode="numeric"
          getInputRef={inputRef}
          placeholder="0.0000"
          value={isLoading ? '0' : showValue}
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
      </HoverTooltip>
    </div>
  )
}
