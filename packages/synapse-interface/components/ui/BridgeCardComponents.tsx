import { Token } from '@/utils/types'
import { useState } from 'react'
import { DropDownArrowSvg } from '../icons/DropDownArrowSvg'
import {
  getBorderStyleForCoinHover,
  getMenuItemHoverBgForCoin,
} from '@/styles/tokens'
import LoadingDots from './tailwind/LoadingDots'

type TokenSelectorTypes = {
  dataTestId?: string
  token: Token
  placeholder: string
  onClick: () => void
}

type AmountInputTypes = {
  inputRef?: React.RefObject<HTMLInputElement>
  disabled?: boolean
  hasMounted?: boolean
  isConnected?: boolean
  isLoading?: boolean
  showValue: string
  handleFromValueChange?: (event: React.ChangeEvent<HTMLInputElement>) => void
  parsedBalance?: string
  onMaxBalance?: () => void
}

const join = (a) => Object.values(a).join(' ')

export function BridgeCard({ children }) {
  /* TODOs
   * Lift margin value up to parent
   * Remove need for popoverDependencies styles
   */
  const className = join({
    space: 'p-4 mt-5 rounded-[.75rem]',
    bgColor: 'bg-bgBase', // NEW: 'bg-zinc-100 dark:bg-zinc-900/95 shadow-xl',
    popoverDependencies: 'overflow-hidden transform',
  })

  return <div className={className}>{children}</div>
}

export function BridgeSectionContainer({ children }) {
  const className = join({
    space: 'grid gap-2 p-2 rounded-md',
    bgColor: 'bg-bgLight', // NEW: 'bg-zinc-50 dark:bg-zinc-800',
    borderColor: 'border border-zinc-300 dark:border-transparent',
  })

  return <section className={className}>{children}</section>
}

export function BridgeAmountContainer({ children }) {
  const className = join({
    space: 'flex items-center gap-4 p-2 rounded-md',
    bgColor: 'bg-white dark:bg-inherit',
    borderColor: 'border border-zinc-200 dark:border-zinc-700',
  })

  return <div className={className}>{children}</div>
}

export function TokenSelector({
  dataTestId,
  token,
  placeholder,
  onClick,
}: TokenSelectorTypes) {
  const className = join({
    space: 'flex items-center gap-2 p-2 rounded flex-none',
    text: 'text-lg',
    bgColor: 'bg-[#565058]', // NEW: 'bg-inherit dark:bg-zinc-700',
    borderColor: `border border-transparent`,
    bgHover: getMenuItemHoverBgForCoin(token?.color),
    borderHover: getBorderStyleForCoinHover(token?.color),
  })

  return (
    <button data-test-id={dataTestId} className={className} onClick={onClick}>
      {token && (
        <img
          src={token?.icon?.src ?? ''}
          alt={token?.symbol ?? ''}
          width="24"
          height="24"
        />
      )}
      {token?.symbol ?? placeholder}
      <DropDownArrowSvg />
    </button>
  )
}

export function AmountInput({
  inputRef,
  disabled = false,
  hasMounted,
  isConnected,
  isLoading = false,
  showValue,
  handleFromValueChange,
  parsedBalance,
  onMaxBalance,
}: AmountInputTypes) {
  const inputClassName = join({
    space: 'p-0 w-full',
    bgColor: 'bg-transparent',
    borderColor: 'border-none',
    textColor: 'placeholder:text-[#88818C] text-white text-opacity-80',
    font: 'text-xl md:text-2xl font-medium',
    focus: 'focus:outline-none focus:ring-0 focus:border-none',
  })

  const labelClassName = join({
    space: 'block',
    textColor: 'text-xs text-white',
    animation: 'transition-all duration-150 transform-gpu',
    hover: 'hover:text-opacity-70 hover:cursor-pointer',
  })

  return (
    <div className="flex-1">
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
          onChange={handleFromValueChange}
          value={showValue}
          name="inputRow"
          autoComplete="off"
          minLength={1}
          maxLength={79}
        />
      )}
      {hasMounted && isConnected && !disabled && (
        <label
          htmlFor="inputRow"
          className={labelClassName}
          onClick={onMaxBalance}
        >
          {parsedBalance ?? '0.0'}
          <span className="opacity-50"> available</span>
        </label>
      )}
    </div>
  )
}