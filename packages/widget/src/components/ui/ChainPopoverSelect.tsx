import { useState, useEffect } from 'react'
import _ from 'lodash'
import { Chain } from 'types'

import usePopover from '@/hooks/usePopoverRef'
import { DownArrow } from '@/components/icons/DownArrow'
import { SearchInput } from './SearchInput'

type PopoverSelectProps = {
  options: Chain[]
  remaining: Chain[]
  onSelect: (selected: Chain) => void
  selected: Chain
  label: string
}

export const ChainPopoverSelect = ({
  options,
  remaining,
  onSelect,
  selected,
  label,
}: PopoverSelectProps) => {
  const { popoverRef, isOpen, togglePopover, closePopover } = usePopover()

  const handleSelect = (option: Chain) => {
    onSelect(option)
    closePopover()
  }

  const {
    filterValue,
    setFilterValue,
    filteredOptions,
    filteredRemaining,
    hasFilteredRemaining,
    hasFilteredResults,
  } = useChainInputFilter(options, remaining, isOpen)

  return (
    <div
      data-test-id="chain-popover-select"
      className="relative w-min col-span-full"
      ref={popoverRef}
    >
      <div
        onClick={() => togglePopover()}
        style={{ background: 'var(--synapse-select-bg)' }}
        className={`
          flex px-2.5 py-1.5 gap-2 items-center rounded
          text-[--synapse-select-text] whitespace-nowrap
          border border-solid border-[--synapse-select-border]
          cursor-pointer hover:border-[--synapse-focus]
        `}
      >
        {selected?.name || 'Network'}
        <DownArrow />
      </div>
      {isOpen && (
        <div
          style={{ background: 'var(--synapse-select-bg)' }}
          className={`
            absolute z-50 mt-1 p-1 max-h-60 min-w-48 rounded
            shadow popover text-left list-none overflow-y-auto
            border border-solid border-[--synapse-select-border]
          `}
        >
          <SearchInput
            inputValue={filterValue}
            setInputValue={setFilterValue}
            placeholder="Search Chains"
            isActive={isOpen}
          />
          {hasFilteredResults ? (
            <ul className="p-0 mt-px mb-0 space-y-px">
              {filteredOptions.map((option, i) => (
                <ChainOption
                  key={i}
                  option={option}
                  isSelected={option?.name === selected?.name}
                  onSelect={() => handleSelect(option)}
                />
              ))}
              {hasFilteredRemaining && (
                <div
                  style={{ background: 'var(--synapse-select-bg)' }}
                  className={`
                    sticky top-0 px-2.5 py-2 mt-2
                    text-sm text-[--synapse-secondary]
                  `}
                >
                  Other chains
                </div>
              )}
              {filteredRemaining.map((option, i) => (
                <ChainOption
                  key={i}
                  option={option}
                  isSelected={option?.name === selected?.name}
                  onSelect={() => handleSelect(option)}
                />
              ))}
            </ul>
          ) : (
            <div className="p-2 break-all">
              No chains found
              <br />
              matching '{filterValue}'.
            </div>
          )}
        </div>
      )}
    </div>
  )
}

const ChainOption = ({
  option,
  isSelected,
  onSelect,
}: {
  option: Chain
  isSelected: boolean
  onSelect: (option: Chain) => void
}) => (
  <li
    key={option.id}
    className={`
      pl-2.5 pr-8 py-2.5 rounded border border-solid
      hover:border-[--synapse-focus] active:opacity-40
      cursor-pointer whitespace-nowrap
      ${
        isSelected
          ? 'border-[--synapse-focus] hover:opacity-70'
          : 'border-transparent'
      }
    `}
    onClick={() => onSelect(option)}
  >
    {option.name}
  </li>
)

const useChainInputFilter = (
  options: Chain[],
  remaining: Chain[],
  isActive: boolean
) => {
  const [filterValue, setFilterValue] = useState('')

  useEffect(() => {
    if (!isActive) {
      setFilterValue('')
    }
  }, [isActive])

  const filterChains = (chains: Chain[], filter: string) => {
    const lowerFilter = filter.toLowerCase()
    return _.filter(chains, (option) => {
      const name = option.name.toLowerCase()
      return name.includes(lowerFilter) || name === lowerFilter
    })
  }

  const filteredOptions = filterChains(options, filterValue)
  const filteredRemaining = filterChains(remaining, filterValue)

  const hasFilteredOptions = !_.isEmpty(filteredOptions)
  const hasFilteredRemaining = !_.isEmpty(filteredRemaining)
  const hasFilteredResults = hasFilteredOptions || hasFilteredRemaining

  return {
    filterValue,
    setFilterValue,
    filteredOptions,
    filteredRemaining,
    hasFilteredOptions,
    hasFilteredRemaining,
    hasFilteredResults,
  }
}
