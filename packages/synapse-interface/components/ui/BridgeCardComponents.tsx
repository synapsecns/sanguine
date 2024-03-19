import { useCallback, useEffect, useRef, useState } from 'react'
import { Chain, Token } from '@/utils/types'
import { DropDownArrowSvg } from '../icons/DropDownArrowSvg'
import { getHoverStyleForButton } from '@/styles/hover'
import LoadingDots from './tailwind/LoadingDots'
import useCloseOnOutsideClick from '@/utils/hooks/useCloseOnOutsideClick'
import SlideSearchBox from '@/pages/bridge/SlideSearchBox'
import { CloseButton } from '../StateManagedBridge/components/CloseButton'
import { useKeyPress } from '@/utils/hooks/useKeyPress'
import { SelectSpecificNetworkButton } from '../StateManagedBridge/components/SelectSpecificNetworkButton'
import { SearchResults } from '../StateManagedBridge/components/SearchResults'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { useBridgeState } from '@/slices/bridge/hooks'
import { useDispatch } from 'react-redux'
import { FromChainListArray } from '../StateManagedBridge/FromChainListOverlay'

const join = (a) => Object.values(a).join(' ')

interface BridgeCardTypes {
  bridgeRef: React.RefObject<HTMLDivElement>
  children: React.ReactNode
}

interface SelectorTypes {
  dataTestId?: string
  label?: string
  placeholder?: string
  selectedItem: Token | Chain
  overlay?: React.ReactNode
}

interface TokenSelectorTypes extends SelectorTypes {
  selectedItem: Token
}

interface ChainSelectorTypes extends SelectorTypes {
  selectedItem: Chain
  itemListFunction: Function
  setFunction: Function
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

export function BridgeCard({ bridgeRef, children }: BridgeCardTypes) {
  /* TODOs
   * Lift margin value up to parent
   * Remove need for popoverDependencies styles (in progress)
   */
  const className = join({
    grid: 'grid gap-2',
    space: 'p-3 mt-5 rounded-[.75rem]',
    background: 'bg-zinc-100 dark:bg-bgBase', // TODO: Remove
    // background: 'bg-zinc-100 dark:bg-zinc-900/95 shadow-xl',
    // popoverDependencies: 'overflow-hidden transform',
  })

  return (
    <div ref={bridgeRef} className={className}>
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

export function ChainSelector({
  dataTestId,
  selectedItem,
  label,
  placeholder,
  overlay, // TODO: Remove
  itemListFunction,
  setFunction,
}: ChainSelectorTypes) {
  const { fromChainId } = useBridgeState()
  const [currentIdx, setCurrentIdx] = useState(-1)
  const [searchStr, setSearchStr] = useState('')
  const [hover, setHover] = useState(false)

  const dispatch = useDispatch()
  const overlayRef = useRef(null)

  const escPressed = useKeyPress('Escape')
  const arrowUp = useKeyPress('ArrowUp')
  const arrowDown = useKeyPress('ArrowDown')

  const isOrigin = label === 'From' // TODO: Improve

  const itemList = itemListFunction?.(searchStr) ?? {}

  const onClose = useCallback(
    () => {
      setCurrentIdx(-1)
      setSearchStr('')
      setHover(false)
    },
    [
      /*setShowFromChainListOverlay*/
    ]
  )

  const escFunc = () => {
    if (escPressed) {
      onClose()
    }
  }
  const arrowDownFunc = () => {
    const nextIdx = currentIdx + 1
    if (arrowDown && nextIdx < 0) {
      // masterList.length) {
      setCurrentIdx(nextIdx)
    }
  }

  const arrowUpFunc = () => {
    const nextIdx = currentIdx - 1
    if (arrowUp && -1 < nextIdx) {
      setCurrentIdx(nextIdx)
    }
  }

  const onSearch = (str: string) => {
    setSearchStr(str)
    setCurrentIdx(-1)
  }

  useEffect(arrowDownFunc, [arrowDown])
  useEffect(escFunc, [escPressed])
  useEffect(arrowUpFunc, [arrowUp])
  useCloseOnOutsideClick(overlayRef, onClose)

  const handleSetFromChainId = (chainId) => {
    if (fromChainId !== chainId) {
      const eventTitle = `[Bridge User Action] Sets new fromChainId`
      const eventData = {
        previousFromChainId: fromChainId,
        newFromChainId: chainId,
      }

      segmentAnalyticsEvent(eventTitle, eventData)
      dispatch(setFunction(chainId))
    }
    onClose()
  }

  useEffect(() => {
    const ref = overlayRef?.current
    if (!ref) return
    const { y, height } = ref.getBoundingClientRect()
    const screen = window.innerHeight
    if (y + height > screen) {
      ref.style.position = 'fixed'
      ref.style.bottom = '4px'
    }
    if (y < 0) {
      ref.style.position = 'fixed'
      ref.style.top = '4px'
    }
  })

  const buttonClassName = join({
    unset: 'text-left',
    flex: 'flex items-center gap-2.5',
    space: 'px-2 py-1.5 mx-0.5 mb-1 rounded flex-none',
    background: 'bg-transparent',
    border: 'border border-zinc-200 dark:border-transparent',
    font: 'leading-tight',
    hover: getHoverStyleForButton(selectedItem?.color),
    active: 'active:opacity-70',
  })

  return (
    <div
      className="relative"
      onMouseEnter={() => setHover(true)}
      onMouseLeave={() => setHover(false)}
      onMouseDown={(e) => e.stopPropagation()}
    >
      <button
        data-test-id={`${dataTestId}-button`}
        onClick={() => setHover(!hover)}
        className={buttonClassName}
      >
        {selectedItem && (
          <img
            src={selectedItem?.chainImg?.src}
            alt={selectedItem?.name}
            width="24"
            height="24"
          />
        )}
        <span>
          <div className="text-sm text-zinc-500 dark:text-zinc-400">
            {label}
          </div>
          {selectedItem?.name ?? placeholder ?? 'Network'}
        </span>
        <DropDownArrowSvg />
      </button>
      {hover && (
        <div
          ref={overlayRef}
          data-test-id="fromChain-list-overlay"
          className="z-20 absolute bg-bgLight border border-separator rounded overflow-y-auto max-h-96 animate-slide-down origin-top shadow-md"
        >
          <div className="p-1 flex items-center font-medium">
            <SlideSearchBox
              placeholder="Find"
              searchStr={searchStr}
              onSearch={onSearch}
            />
            <CloseButton onClick={onClose} />
          </div>
          <div data-test-id={dataTestId}>
            {Object.entries(itemList).map(([key, value]: [string, Chain[]]) => {
              return value.length ? (
                <div key={key} className="bg-bgBase first:bg-bgLight rounded">
                  <div className="p-2 text-sm text-secondary sticky top-0 bg-inherit z-10">
                    {key}
                  </div>
                  {value.map(({ id }, idx) => (
                    <SelectSpecificNetworkButton
                      key={id}
                      itemChainId={id}
                      isCurrentChain={selectedItem?.id === id}
                      isOrigin={isOrigin}
                      active={idx === currentIdx}
                      onClick={() => handleSetFromChainId(id)}
                      dataId={dataTestId}
                    />
                  ))}
                </div>
              ) : null
            })}
            <SearchResults searchStr={searchStr} type="chain" />
          </div>
        </div>
      )}
    </div>
  )
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
  overlay,
}: TokenSelectorTypes) {
  const [hover, setHover] = useState(false)

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
    <div
      className="relative"
      onMouseEnter={() => setHover(true)}
      onMouseLeave={() => setHover(false)}
      onMouseDown={(e) => e.stopPropagation()}
    >
      <button
        data-test-id={dataTestId}
        onClick={() => setHover(!hover)}
        className={className}
      >
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
      {hover && overlay}
    </div>
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
