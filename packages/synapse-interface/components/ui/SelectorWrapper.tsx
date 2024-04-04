import React, { useEffect, useRef } from 'react'

import { getHoverStyleForButton } from '@/styles/hover'
import { DropDownArrowSvg } from '@/components/icons/DropDownArrowSvg'
import { SlideSearchBox } from '@/components/ui/SlideSearchBox'
import { CloseButton } from '@/components/ui/CloseButton'
import { SearchResults } from '@/components/ui/SearchResults'
import { useKeyPress } from '@/utils/hooks/useKeyPress'
import { joinClassNames } from '@/utils/joinClassNames'

export const SelectorWrapper = ({
  dataTestId,
  label,
  placeholder,
  selectedItem,
  children,
  searchStr,
  onClose,
  onSearch,
  hover,
  setHover,
}) => {
  const escPressed = useKeyPress('Escape', hover)

  const popoverRef = useRef(null)

  useEffect(() => {
    const ref = popoverRef?.current
    if (!ref) return

    const { y, height } = ref.getBoundingClientRect()
    const screen = window.innerHeight

    if (y + height * 0.67 > screen) {
      ref.style.position = 'fixed'
      ref.style.bottom = '4px'
      document.addEventListener('scroll', () => setHover(false), {
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

  const buttonClassName = joinClassNames({
    flex: 'flex items-center gap-2',
    space: 'px-2 py-1.5 rounded',
    border: 'border border-zinc-200 dark:border-transparent',
    text: 'leading-tight',
    hover: getHoverStyleForButton(selectedItem?.color),
    active: 'active:opacity-80',
    custom: label ? 'bg-transparent' : 'bg-white dark:bg-separator text-lg',
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
        onClick={() => setHover(true)}
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
      {hover && (
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
