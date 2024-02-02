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
    noFilteredResults,
  } = useChainInputFilter(options, remaining, isOpen)

  /** Filters chains based on User Input */
  // const [filterValue, setFilterValue] = useState('')

  // useEffect(() => {
  //   if (!isOpen) {
  //     setFilterValue('')
  //   }
  // }, [isOpen])

  // const filteredOptions = _.filter(options, (option) => {
  //   const name = option.name
  //   const lowerName = name.toLowerCase()
  //   const lowerFilter = filterValue.toLowerCase()
  //   return lowerName.includes(lowerFilter) || lowerName === lowerFilter
  // })
  // const filteredRemaining = _.filter(remaining, (option) => {
  //   const name = option.name
  //   const lowerName = name.toLowerCase()
  //   const lowerFilter = filterValue.toLowerCase()
  //   return lowerName.includes(lowerFilter) || lowerName === lowerFilter
  // })

  // const noFilteredOptions = _.isEmpty(filteredOptions)
  // const noFilteredRemaining = _.isEmpty(filteredRemaining)
  // const noFilteredResults = noFilteredOptions && noFilteredRemaining

  return (
    <div
      data-test-id="chain-popover-select"
      className="relative w-min col-span-full"
      ref={popoverRef}
    >
      <div
        className={`
          cursor-pointer flex px-2.5 py-1.5 gap-2 items-center rounded-[.1875rem]
          text-[--synapse-select-text]
          border border-solid border-[--synapse-select-border]
          hover:border-[--synapse-focus]
          whitespace-nowrap
        `}
        style={{ background: 'var(--synapse-select-bg)' }}
        onClick={() => togglePopover()}
      >
        {selected?.name || 'Network'}
        <DownArrow />
      </div>
      {isOpen && (
        <div
          className="absolute z-50 mt-1 p-0 border border-solid border-[--synapse-select-border] rounded shadow popover text-left list-none overflow-y-auto max-h-60  min-w-48"
          style={{ background: 'var(--synapse-select-bg)' }}
        >
          <SearchInput
            inputValue={filterValue}
            setInputValue={setFilterValue}
            placeholder="Search Chains"
          />
          {noFilteredResults ? (
            <div className="p-2 break-all">
              No chains found matching '{filterValue}'.
            </div>
          ) : (
            <ul className="p-0 m-0">
              {filteredOptions.map((option) => (
                <ChainOption
                  option={option}
                  isSelected={option?.name === selected?.name}
                  onSelect={() => handleSelect(option)}
                />
              ))}
              <div
                className="px-2.5 py-2 mt-2 text-sm text-[--synapse-secondary] cursor-default sticky top-0"
                style={{ background: 'var(--synapse-select-bg)' }}
              >
                Other chains
              </div>
              {filteredRemaining.map((option) => (
                <ChainOption
                  option={option}
                  isSelected={option?.name === selected?.name}
                  onSelect={() => handleSelect(option)}
                />
              ))}
            </ul>
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
    className={`cursor-pointer pl-2.5 pr-8 py-2.5 rounded border border-solid hover:border-[--synapse-focus] active:opacity-40 whitespace-nowrap ${
      isSelected
        ? 'border-[--synapse-focus] hover:opacity-70'
        : 'border-transparent'
    }`}
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

  const noFilteredOptions = _.isEmpty(filteredOptions)
  const noFilteredRemaining = _.isEmpty(filteredRemaining)
  const noFilteredResults = noFilteredOptions && noFilteredRemaining

  return {
    filterValue,
    setFilterValue,
    filteredOptions,
    filteredRemaining,
    noFilteredResults,
  }
}
