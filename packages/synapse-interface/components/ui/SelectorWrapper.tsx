import React, { useEffect, useRef } from 'react'

import { getActiveStyleForButton, getHoverStyleForButton } from '@/styles/hover'
import { DropDownArrowSvg } from '@/components/icons/DropDownArrowSvg'
import { SlideSearchBox } from '@/components/ui/SlideSearchBox'
import { CloseButton } from '@/components/ui/CloseButton'
import { SearchResults } from '@/components/ui/SearchResults'
import { useKeyPress } from '@/utils/hooks/useKeyPress'
import { joinClassNames } from '@/utils/joinClassNames'
import useCloseOnOutsideClick from '@/utils/hooks/useCloseOnOutsideClick'

export const SelectorWrapper = ({
  dataTestId,
  label,
  placeholder,
  selectedItem,
  children,
  searchStr,
  onClose,
  onSearch,
  open,
  setOpen,
  disabled,
}) => {
  const escPressed = useKeyPress('Escape', open)

  const parentRef = useRef(null)
  const popoverRef = useRef(null)

  useEffect(() => {
    const ref = popoverRef?.current
    if (!ref) return

    const { y, height } = ref.getBoundingClientRect()
    const screen = window.innerHeight

    if (y + height * 0.67 > screen) {
      ref.style.position = 'fixed'
      ref.style.bottom = '4px'
      document.addEventListener('scroll', () => setOpen(false), {
        once: true,
      })
    }
    if (ref.getBoundingClientRect().y < 0) {
      ref.style.position = 'fixed'
      ref.style.top = '4px'
      const search = ref.firstChild.firstChild
      const inner = search.nextSibling
      inner.style.height = `${screen - search.offsetHeight - 16}px`
    }
  })

  const escFunc = () => {
    if (escPressed) {
      onClose()
    }
  }

  useEffect(escFunc, [escPressed])

  useCloseOnOutsideClick(parentRef, () => onClose())

  const buttonClassNames = {
    flex: 'flex items-center gap-2',
    space: 'px-2 py-1.5 rounded',
    border: 'border border-zinc-200 dark:border-transparent',
    text: `leading-tight ${!label && 'text-lg'}`,
    open: `${getHoverStyleForButton(selectedItem?.color)} ${
      open && getActiveStyleForButton(selectedItem?.color)
    }`,
    active: 'active:opacity-80',
    custom: label || open ? 'bg-transparent' : 'bg-white dark:bg-separator',
    disabled: `${disabled ? 'hover:cursor-not-allowed' : ''}`,
  }

  // TODO: Unify chainImg/icon properties between Chain and Token types
  const imgSrc =
    selectedItem?.['chainImg' in selectedItem ? 'chainImg' : 'icon']?.src

  const itemName = selectedItem?.['symbol' in selectedItem ? 'symbol' : 'name']

  return (
    <div className="relative min-w-fit" ref={parentRef}>
      <button
        data-test-id={`${dataTestId}-button`}
        className={joinClassNames(buttonClassNames)}
        onClick={() => {
          if (!disabled) {
            !open ? setOpen(true) : onClose()
          }
        }}
      >
        {itemName && (
          <img
            src={imgSrc}
            alt={itemName}
            width="24"
            height="24"
            className="py-0.5 block"
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
      {open && (
        <div
          ref={popoverRef}
          data-test-id={`${dataTestId}-overlay`}
          className="absolute z-20 pt-1 animate-slide-down"
        >
          <div className="relative">
            <div className="absolute border rounded shadow-md bg-bgLight border-separator">
              <div className="flex items-center p-1 font-medium">
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
              <SearchResults searchStr={searchStr} />
            </div>
          </div>
        </div>
      )}
    </div>
  )
}
