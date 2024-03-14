import { Chain, Token } from '@/utils/types'
import { DropDownArrowSvg } from '../icons/DropDownArrowSvg'
import { getHoverStyleForButton } from '@/styles/hover'
import LoadingDots from './tailwind/LoadingDots'

const join = (a) => Object.values(a).join(' ')

interface BridgeCardTypes {
  ref: React.RefObject<HTMLDivElement>
  children: React.ReactNode
}

interface SelectorTypes {
  dataTestId?: string
  placeholder?: string
  selectedItem: Token | Chain
  onClick: React.MouseEventHandler<HTMLButtonElement>
}

interface TokenSelectorTypes extends SelectorTypes {
  selectedItem: Token
}

interface ChainSelectorTypes extends SelectorTypes {
  label: string
  selectedItem: Chain
}

interface AmountInputTypes {
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

export function BridgeCard({ ref, children }: BridgeCardTypes) {
  /* TODOs
   * Lift margin value up to parent
   * Remove need for popoverDependencies styles
   */
  const className = join({
    grid: 'grid gap-2',
    space: 'p-3 mt-5 rounded-[.75rem]',
    background: 'bg-zinc-100 dark:bg-bgBase', // TODO: Remove
    // background: 'bg-zinc-100 dark:bg-zinc-900/95 shadow-xl',
    popoverDependencies: 'overflow-hidden transform',
  })

  return (
    <div ref={ref} className={className}>
      {children}
    </div>
  )
}

export function BridgeSectionContainer({ children }) {
  const className = join({
    space: 'grid gap-2 p-2 rounded-md',
    background: 'bg-zinc-50 dark:bg-bgLight', // TODO: Remove
    // background: 'bg-zinc-50 dark:bg-zinc-800',
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
  selectedItem,
  placeholder,
  onClick,
}: TokenSelectorTypes) {
  const className = join({
    flex: 'flex items-center gap-2',
    background: 'bg-white dark:bg-separator',
    // background: 'bg-white dark:bg-zinc-700',
    border: 'border border-zinc-200 dark:border-transparent',
    space: 'p-2 rounded flex-none',
    font: 'text-lg',
    hover: getHoverStyleForButton(selectedItem?.color),
    active: 'active:opacity-75',
  })

  return (
    <button data-test-id={dataTestId} className={className} onClick={onClick}>
      {selectedItem && (
        <img
          src={selectedItem?.icon?.src ?? ''}
          alt={selectedItem?.symbol ?? ''}
          width="24"
          height="24"
        />
      )}
      {selectedItem?.symbol ?? placeholder ?? 'Token'}
      <DropDownArrowSvg />
    </button>
  )
}

export function ChainSelector({
  dataTestId,
  selectedItem,
  label,
  placeholder,
  onClick,
}: ChainSelectorTypes) {
  const className = join({
    unset: 'text-left',
    flex: 'flex items-center gap-2.5',
    space: 'px-2 py-1.5 mx-0.5 rounded flex-none',
    background: 'bg-transparent',
    border: 'border border-zinc-200 dark:border-transparent',
    font: 'leading-tight',
    hover: getHoverStyleForButton(selectedItem?.color),
    active: 'active:opacity-70',
  })

  return (
    <button data-test-id={dataTestId} className={className} onClick={onClick}>
      {selectedItem && (
        <img
          src={selectedItem?.chainImg?.src}
          alt={selectedItem?.name}
          width="24"
          height="24"
        />
      )}
      <span>
        <div className="text-sm text-zinc-500 dark:text-zinc-400">{label}</div>
        {selectedItem?.name ?? placeholder ?? 'Network'}
      </span>
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
    unset: 'bg-transparent border-none p-0',
    layout: 'w-full',
    placeholder: 'placeholder:text-zinc-500 placeholder:dark:text-zinc-400',
    font: 'text-xl md:text-2xl font-medium',
    focus: 'focus:outline-none focus:ring-0 focus:border-none',
  })

  const labelClassName = join({
    space: 'block',
    textColor: 'text-xs',
    animation: 'transition-all duration-150 transform-gpu',
    hover: 'hover:opacity-70 cursor-pointer',
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
          <span className="text-zinc-500 dark:text-zinc-400"> available</span>
        </label>
      )}
    </div>
  )
}
