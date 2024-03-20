import React, { useCallback, useEffect, useRef, useState } from 'react'
import { Chain, Token } from '@/utils/types'
import { DropDownArrowSvg } from '../icons/DropDownArrowSvg'
import { getHoverStyleForButton, getActiveStyleForButton } from '@/styles/hover'
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
import SelectSpecificTokenButton from '../StateManagedBridge/components/SelectSpecificTokenButton'

const join = (a) => Object.values(a).join(' ')

interface BridgeCardTypes {
  bridgeRef: React.RefObject<HTMLDivElement>
  children: React.ReactNode
}

interface SelectorTypes {
  dataTestId?: string
  isOrigin: boolean
  label?: string
  placeholder?: string
  selectedItem: Token | Chain
  itemListFunction?: Function
  setFunction?: Function
}

interface TokenSelectorTypes extends SelectorTypes {
  selectedItem: Token
}

interface ChainSelectorTypes extends SelectorTypes {
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
  isOrigin,
  selectedItem,
  label,
  placeholder,
  itemListFunction,
  setFunction,
}: ChainSelectorTypes) {
  const [searchStr, setSearchStr] = useState('')

  const dispatch = useDispatch()

  const handleSetChainId = (chainId) => {
    if (selectedItem?.id !== chainId) {
      const eventTitle = `[Bridge User Action] Sets new fromChainId`
      const eventData = {
        previousFromChainId: selectedItem?.id,
        newFromChainId: chainId,
      }

      segmentAnalyticsEvent(eventTitle, eventData)
      dispatch(setFunction(chainId))
    }
  }

  const itemList = itemListFunction(searchStr)

  return (
    <SelectorWrapper
      dataTestId={dataTestId}
      label={label}
      placeholder={placeholder ?? 'Network'}
      selectedItem={selectedItem}
      searchStr={searchStr}
      setSearchStr={setSearchStr}
    >
      {Object.entries(itemList).map(([key, value]: [string, Chain[]]) => {
        return value.length ? (
          <ListSectionWrapper sectionKey={key}>
            {value.map((chain, idx) => (
              <SelectSpecificNetworkButton
                dataId={dataTestId}
                key={chain.id}
                itemChainId={chain.id}
                isOrigin={isOrigin}
                isCurrentChain={selectedItem?.id === chain.id}
                active={false}
                onClick={() => handleSetChainId(chain.id)}
              />
            ))}
          </ListSectionWrapper>
        ) : null
      })}
    </SelectorWrapper>
  )
}

const ListSectionWrapper = ({ sectionKey, children }) => (
  <section key={sectionKey} className="bg-bgBase first:bg-bgLight rounded">
    <header
      className="p-2 text-sm text-secondary sticky top-0 bg-inherit z-10 cursor-default"
      onClick={(e) => e.stopPropagation()}
    >
      {sectionKey}
    </header>
    {children}
  </section>
)

const SelectorWrapper = ({
  dataTestId,
  label,
  placeholder,
  selectedItem,
  children,
  searchStr,
  setSearchStr,
}) => {
  const [hover, setHover] = useState(false)

  const escPressed = useKeyPress('Escape')

  const popoverRef = useRef(null)

  useEffect(() => {
    const ref = popoverRef?.current
    if (!ref) return
    // if (window.innerWidth >= 1024) return

    if (searchStr) {
      ref.style.position = 'absolute'
      ref.style.top = 'auto'
      ref.style.bottom = 'auto'
    } else {
      const { y, height } = ref.getBoundingClientRect()
      const screen = window.innerHeight

      if (y + height * 0.67 > screen) {
        ref.style.position = 'fixed'
        ref.style.bottom = '4px'
        document.addEventListener('scroll', () => setHover(false), {
          once: true,
        })
      }
      if (y < 0) {
        ref.style.position = 'fixed'
        ref.style.top = '4px'
      }
    }
  })

  const onClose = useCallback(() => {
    setSearchStr('')
    setHover(false)
  }, [])

  const escFunc = () => {
    if (escPressed) {
      onClose()
    }
  }

  const onSearch = (str: string) => {
    setSearchStr(str)
  }

  useEffect(escFunc, [escPressed])

  function handleMouseMove(e) {
    if (
      (Math.round(e.movementX) < 1 && !e.movementY) ||
      (Math.round(e.movementY) < 1 && !e.movementX)
    )
      setHover(true)
  }

  function handleMouseLeave() {
    if (!searchStr) onClose()
  }

  const buttonClassName = join({
    flex: 'flex items-center gap-2',
    space: 'px-2 py-1.5 rounded',
    border: 'border border-zinc-200 dark:border-transparent',
    text: 'leading-tight',
    hover: getHoverStyleForButton(selectedItem?.color),
    active: 'active:opacity-80',
    custom: label ? 'bg-transparent' : 'bg-white dark:bg-separator text-lg',
    // bugfix: 'flex-none', // may not be needed any more
  })

  // TODO: Unify chainImg/icon properties between Chain and Token types
  const imgSrc =
    selectedItem?.['chainImg' in selectedItem ? 'chainImg' : 'icon']?.src

  const itemName = selectedItem?.['symbol' in selectedItem ? 'symbol' : 'name']

  return (
    <div className="relative" onMouseLeave={handleMouseLeave}>
      <button
        data-test-id={`${dataTestId}-button`}
        className={buttonClassName}
        onMouseMove={handleMouseMove}
        onClick={() => setHover(!hover)}
      >
        {itemName && (
          <img
            src={imgSrc}
            alt={itemName}
            width="24"
            height="24"
            className="py-0.5"
          />
        )}
        <span>
          {label && (
            <div className="text-sm text-left text-zinc-500 dark:text-zinc-400">
              {label}
            </div>
          )}
          {itemName ?? placeholder}
        </span>
        <DropDownArrowSvg />
      </button>
      {hover && (
        <div
          ref={popoverRef}
          data-test-id={`${dataTestId}-overlay`}
          className="z-20 absolute animate-slide-down pt-1"
        >
          <div className="bg-bgLight border border-separator rounded shadow-md">
            <div className="p-1 flex items-center font-medium">
              <SlideSearchBox
                placeholder="Find"
                searchStr={searchStr}
                onSearch={onSearch}
              />
              <CloseButton onClick={onClose} />
            </div>
            <div
              data-test-id={`${dataTestId}-list`}
              className="overflow-y-auto max-h-96"
              onClick={onClose}
            >
              {children}
            </div>
            <SearchResults searchStr={searchStr} type="chain" />
          </div>
        </div>
      )}
    </div>
  )
}

export function TokenSelector({
  dataTestId,
  selectedItem,
  label,
  placeholder,
  itemListFunction,
  setFunction,
  isOrigin,
}: TokenSelectorTypes) {
  const [currentIdx, setCurrentIdx] = useState(-1)
  const [searchStr, setSearchStr] = useState('')

  const dispatch = useDispatch()

  const handleSetFromToken = (token) => {
    if (selectedItem !== token) {
      const eventTitle = `[Bridge User Action] Sets new fromChainId`
      const eventData = {
        previousFromToken: selectedItem?.symbol,
        newFromToken: token?.symbol,
      }
      segmentAnalyticsEvent(eventTitle, eventData)
      dispatch(setFunction(token))
    }
  }

  const itemList = itemListFunction(searchStr) // TODO: Use result instead of variable in context?

  return (
    <SelectorWrapper
      dataTestId={dataTestId}
      label={label}
      placeholder={placeholder ?? 'Network'}
      selectedItem={selectedItem}
      searchStr={searchStr}
      setSearchStr={setSearchStr}
    >
      {Object.entries(itemList).map(([key, value]: [string, Token[]]) => {
        return value.length ? (
          <ListSectionWrapper sectionKey={key}>
            {value.map((token, idx) => (
              <SelectSpecificTokenButton
                isOrigin={isOrigin}
                key={idx}
                token={token}
                selectedToken={selectedItem}
                // active={
                //   idx +
                //     possibleTokens.length +
                //     remainingChainTokens.length ===
                //   currentIdx
                // }
                active={false}
                // showAllChains={true}
                onClick={() => handleSetFromToken(token)}
                alternateBackground={false}
              />
            ))}
          </ListSectionWrapper>
        ) : null
      })}
    </SelectorWrapper>
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
