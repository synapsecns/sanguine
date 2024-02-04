import { useState, useEffect } from 'react'
import _ from 'lodash'
import { Chain } from 'types'

import usePopover from '@/hooks/usePopoverRef'
import { DownArrow } from '@/components/icons/DownArrow'
import { SearchInput } from './SearchInput'
import { TabOption, ToggleTabs } from './ToggleTabs'
import { ChainOption } from './ChainOption'
import { useChainInputFilter } from '@/hooks/useChainInputFilter'

type PopoverSelectProps = {
  options: Chain[]
  remaining: Chain[]
  targets: Chain[]
  onSelect: (selected: Chain) => void
  selected: Chain
  label: string
  isOrigin: boolean
}

export const ChainPopoverSelect = ({
  options,
  remaining,
  targets,
  onSelect,
  selected,
  label,
  isOrigin,
}: PopoverSelectProps) => {
  const { popoverRef, isOpen, togglePopover, closePopover } = usePopover()
  const [activeTab, setActiveTab] = useState<TabOption>(
    isOrigin ? 'All' : 'Target'
  )

  useEffect(() => {
    if (!targets || _.isEmpty(targets)) {
      setActiveTab('All')
    } else if (isOrigin) {
      setActiveTab('All')
    } else if (!isOrigin) {
      setActiveTab('Target')
    }
  }, [targets, isOrigin])

  const handleSelect = (option: Chain) => {
    onSelect(option)
    closePopover()
  }

  const handleTabSelect = (tab: TabOption) => {
    setActiveTab(tab)
  }

  const {
    filterValue,
    setFilterValue,
    filteredOptions,
    filteredRemaining,
    filteredTargets,
    hasFilteredRemaining,
    hasFilteredResults,
    hasFilteredTargets,
  } = useChainInputFilter(options, remaining, targets, isOpen)

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
            absolute z-50 mt-1 max-h-60 min-w-48 rounded
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
          {targets && targets.length > 0 && (
            <ToggleTabs
              selectedTab={activeTab}
              onTabSelect={handleTabSelect}
              isOrigin={isOrigin}
            />
          )}
          {activeTab === 'All' ? (
            hasFilteredResults ? (
              <ul className="p-0 m-0">
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
                    className={`
                    sticky top-0 px-2.5 py-2 mt-2 text-sm
                    text-[--synapse-secondary] bg-[--synapse-surface]
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
              <div className="p-2 text-sm break-all">
                No chains found
                <br />
                matching '{filterValue}'.
              </div>
            )
          ) : null}
          {activeTab === 'Target' ? (
            hasFilteredTargets ? (
              <ul className="p-0 m-0">
                {filteredTargets.map((option, i) => (
                  <ChainOption
                    key={i}
                    option={option}
                    isSelected={option?.name === selected?.name}
                    onSelect={() => handleSelect(option)}
                  />
                ))}
              </ul>
            ) : (
              <div className="p-2 text-sm break-all">
                No chains found
                <br />
                matching '{filterValue}'.
              </div>
            )
          ) : null}
        </div>
      )}
    </div>
  )
}
